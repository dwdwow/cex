package cex

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"reflect"
)

// ReqConfig save some read-only info.
// ReqType and RespType are not used in ReqConfig,
// but in practice, it is very useful.
// In practice, we call Request to query cex data,
// but we should know config, ReqType and RespType simultaneously.
// We have many config implementations in all cex packages.
// These config with patterns bind config, ReqType and RespType together.
// Set a config instance in Request as input, all patterns in Request are defined.
type ReqConfig[ReqType, RespType any] struct {
	// ex. /path/to/service
	Path string

	// http method, GET, POST...
	// better to use const method value in http package directly
	Method string

	// if true, should use api key
	IsUserData bool

	// one user can rest every UserTimeInterval.
	// unit is millisecond
	UserTimeInterval int64

	// one ip can reset every IpTimeInterval
	// unit is millisecond
	IpTimeInterval int64

	// status code and its status message
	StatusCodes map[int]string

	// cex self custom codes and its meaning
	CexCustomCodes map[int]string
}

type BaseConfig struct {
	// ex. /path/to/service
	Path string

	// http method, GET, POST...
	// better to use const method value in http package directly
	Method string

	// if true, should use api key
	IsUserData bool
}

func (rc ReqConfig[ReqType, RespType]) BaseConfig() BaseConfig {
	return BaseConfig{
		Path:       rc.Path,
		Method:     rc.Method,
		IsUserData: rc.IsUserData,
	}
}

type ReqOpt func(*resty.Client, *resty.Request)

// ReqHandler should be implemented in all cex package
type ReqHandler interface {
	MakeReq(config BaseConfig, reqData any, opts ...ReqOpt) (*resty.Request, error)
	CheckResp(*resty.Response, *resty.Request) error
}

func Request[ReqType, RespType any](config ReqConfig[ReqType, RespType], reqData ReqType, handler ReqHandler, opts ...ReqOpt) (RespType, error) {
	respData := new(RespType)
	req, err := handler.MakeReq(config.BaseConfig(), reqData, opts...)
	if err != nil {
		return *respData, err
	}

	var resp *resty.Response
	switch config.Method {
	case http.MethodGet:
		resp, err = req.Get(config.Path)
	case http.MethodPost:
		resp, err = req.Post(config.Path)
	case http.MethodPut:
		resp, err = req.Put(config.Path)
	case http.MethodDelete:
		resp, err = req.Delete(config.Path)
	default:
	}
	if resp == nil {
		return *respData, fmt.Errorf("cex: http method %v is not supported", config.Method)
	}

	if err = handler.CheckResp(resp, req); err != nil {
		return *respData, fmt.Errorf("cex: check response, %w", err)
	}

	respBody := resp.Body()

	respType := reflect.TypeOf(respData).Elem()

	var result any

	switch respType.Kind() {
	case reflect.String:
		result = any(string(respBody))
	case reflect.Map, reflect.Struct:
		err = json.Unmarshal(respBody, respData)
		if err != nil {
			err = fmt.Errorf("cex: unmarshal response body, %w", err)
		}
		result = any(*respData)
	default:
		return *respData, fmt.Errorf("cex: response data type %v is not supported", respType.Kind())
	}

	return result.(RespType), err
}
