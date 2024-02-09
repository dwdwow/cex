package cex

// ReqConfig save some read-only info.
type ReqConfig[ReqData, RespData any] struct {
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

type ReqMaker[ReqData, RespData any] func(config ReqConfig[ReqData, RespData], reqData ReqData)

func reqMaker[ReqData, RespData any](config ReqConfig[ReqData, RespData], reqData ReqData) {

}

//func a()
