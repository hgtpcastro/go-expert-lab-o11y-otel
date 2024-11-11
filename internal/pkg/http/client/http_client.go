package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	timeOut       = 5 * time.Second
	retryCount    = 3
	retryWaitTime = 300 * time.Millisecond
)

func NewHttpClient() *resty.Client {
	client := resty.New().
		SetTimeout(timeOut).
		SetRetryCount(retryCount).
		SetRetryWaitTime(retryWaitTime)
	return client
}
