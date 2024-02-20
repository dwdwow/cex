package cex

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func ReqOptRetryCount(count int, waitTime time.Duration) ReqOpt {
	return func(client *resty.Client, r *resty.Request) {
		if client == nil {
			return
		}
		client.SetRetryCount(count)
		client.SetRetryWaitTime(waitTime)
		client.SetRetryMaxWaitTime(waitTime)
	}
}
