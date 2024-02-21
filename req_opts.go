package cex

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func CltOptRetryCount(count int, waitTime time.Duration) CltOpt {
	return func(client *resty.Client) {
		if client == nil {
			return
		}
		client.SetRetryCount(count)
		client.SetRetryWaitTime(waitTime)
		client.SetRetryMaxWaitTime(waitTime)
	}
}
