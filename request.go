package cex

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"reflect"
)

// RespCodeChecker checks http and cex custom codes.
// All cex package should have two implementations of this function type,
// one is http's, another is cex's.
type RespCodeChecker func(int) error

// EmptyReqData means that no request data.
// If a ReqConfig ReqDataType is this,
// reqData should be nil.
type EmptyReqData any

// ReqBaseConfig save some read-only info.
// This struct is the real contain of ReqConfig.
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

// ReqConfig is wrapper of ReqBaseConfig.
// This struct makes it convenient to call Request.
// ReqDataType and RespDataType are not used in ReqConfig,
// but in practice, it is very useful.
// In practice, we call Request to query cex data,
// but we should know config, ReqDataType and RespDataType simultaneously.
// We have many config implementations in all cex packages.
// These config with patterns bind config, ReqDataType and RespDataType together.
// Set a config instance in Request as input, all patterns in Request are defined.
type ReqConfig[ReqDataType, RespDataType any] struct {
	ReqBaseConfig
}

// ReqOpt is function option that can custom request.
type ReqOpt func(*resty.Client, *resty.Request)

// Requester should be implemented in all cex package
type Requester interface {
	MakeReq(config ReqBaseConfig, reqData any, opts ...ReqOpt) (*resty.Request, error)
	CheckResp(*resty.Response, *resty.Request) error
}

// Request is the core method in cex.
func Request[ReqDataType, RespDataType any](handler Requester, config ReqConfig[ReqDataType, RespDataType], reqData ReqDataType, opts ...ReqOpt) (*resty.Response, RespDataType, error) {
	respData := new(RespDataType)
	req, err := handler.MakeReq(config.ReqBaseConfig, reqData, opts...)
	if err != nil {
		return nil, *respData, err
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

	if err != nil {
		return resp, *respData, fmt.Errorf("cex: response err: %w", err)
	}

	if resp == nil {
		// should not be here
		// if getting here, resty has bug
		return resp, *respData, fmt.Errorf("cex: http method %v is not supported", config.Method)
	}

	if err = handler.CheckResp(resp, req); err != nil {
		return resp, *respData, fmt.Errorf("cex: check response, %w", err)
	}

	respBody := resp.Body()

	respType := reflect.TypeOf(respData).Elem()

	var anyRes any

	switch respType.Kind() {
	case reflect.String:
		anyRes = any(string(respBody))
	case reflect.Slice, reflect.Struct, reflect.Map:
		err = json.Unmarshal(respBody, respData)
		if err != nil {
			err = fmt.Errorf("cex: unmarshal response body, %w", err)
			return resp, *respData, err
		}
		anyRes = any(*respData)
	default:
		return resp, *respData, fmt.Errorf("cex: response data type %v is not supported", respType.Kind())
	}

	res, ok := anyRes.(RespDataType)

	if !ok {
		err = fmt.Errorf("cex: cannot convert to %T", res)
	}

	return resp, res, err
}
