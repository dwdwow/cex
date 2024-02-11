package cex

type User interface {
	Api() Api
	ReqMaker
}
