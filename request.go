package cex

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"reflect"
)

type RespCodeChecker func(int) error

// ReqDataTypeEmpty means that no request data.
type ReqDataTypeEmpty any

// ReqConfig save some read-only info.
// ReqDataType and RespDataType are not used in ReqConfig,
// but in practice, it is very useful.
// In practice, we call Request to query cex data,
// but we should know config, ReqDataType and RespDataType simultaneously.
// We have many config implementations in all cex packages.
// These config with patterns bind config, ReqDataType and RespDataType together.
// Set a config instance in Request as input, all patterns in Request are defined.
type ReqConfig[ReqDataType, RespDataType any] struct {
	// ex. https://www.example.com
	BaseUrl string

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
	HttpStatusCodeChecker RespCodeChecker

	// cex self custom codes and its meaning
	CexCustomCodeChecker RespCodeChecker
}

type ReqBaseConfig struct {
	// ex. https://www.example.com
	BaseUrl string

	// ex. /path/to/service
	Path string

	// http method, GET, POST...
	// better to use const method value in http package directly
	Method string

	// if true, should use api key
	IsUserData bool
}

func (rc ReqConfig[ReqDataType, RespDataType]) BaseConfig() ReqBaseConfig {
	return ReqBaseConfig{
		BaseUrl:    rc.BaseUrl,
		Path:       rc.Path,
		Method:     rc.Method,
		IsUserData: rc.IsUserData,
	}
}

type ReqOpt func(*resty.Client, *resty.Request)

// ReqHandler should be implemented in all cex package
type ReqHandler interface {
	MakeReq(config ReqBaseConfig, reqData any, opts ...ReqOpt) (*resty.Request, error)
	CheckResp(*resty.Response, *resty.Request) error
}

func Request[ReqDataType, RespDataType any](config ReqConfig[ReqDataType, RespDataType], reqData ReqDataType, handler ReqHandler, opts ...ReqOpt) (RespDataType, error) {
	respData := new(RespDataType)
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

	var anyRes any

	switch respType.Kind() {
	case reflect.String:
		anyRes = any(string(respBody))
	case reflect.Map, reflect.Struct:
		err = json.Unmarshal(respBody, respData)
		if err != nil {
			err = fmt.Errorf("cex: unmarshal response body, %w", err)
			return *respData, err
		}
		anyRes = any(*respData)
	default:
		return *respData, fmt.Errorf("cex: response data type %v is not supported", respType.Kind())
	}

	res, ok := anyRes.(RespDataType)

	if !ok {
		err = fmt.Errorf("cex: cannot convert to %T", res)
	}

	return res, err
}
