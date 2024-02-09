package cex

type User interface {
	Api() Api
	Requester
}
