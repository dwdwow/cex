package cex

type User interface {
	Api() Api
	ReqHandler
}
