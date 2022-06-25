package httpclient

import (
	"net/http"
	"time"
)

type RetryFunc func(req *http.Request, response *http.Response) bool

func ServerTimeout(req *http.Request, res *http.Response) bool {
	shouldRetryCodes := []int{http.StatusBadGateway, http.StatusGatewayTimeout, http.StatusServiceUnavailable}
	for _, c := range shouldRetryCodes {
		if c == res.StatusCode {
			return true
		}
	}
	return false
}
func defaultRetryFunc(_ *http.Request, _ *http.Response) bool {
	return true
}

type retryable struct {
	baseClient
	wait      func(reqTime int) time.Duration
	max       int
	retryWhen func(req *http.Request, response *http.Response) bool
}

// Retryable will create a retry-able http client. After the request is sent, it will execute retryFunc argument.
// when retryFunc returns true, it will retry again until the retry time is equals to max.
// if wait is defined, each retry will be delayed with the given duration from the wait value
func Retryable(max int, retryFunc RetryFunc, wait func(reqTime int) time.Duration) Client {
	if retryFunc == nil {
		retryFunc = defaultRetryFunc
	}
	cli := &retryable{wait: wait, max: max, retryWhen: retryFunc}
	return cli
}

// RoundTrip override the base client baseClient.RoundTrip method. it will internally evaluate when to retry.
func (r *retryable) RoundTrip(req *http.Request) (res *http.Response, err error) {
	for i := 0; i < r.max; i++ {
		res, err = r.baseClient.RoundTrip(req)
		if !r.retryWhen(req, res) {
			break
		}
		if r.wait != nil {
			tm := r.wait(i)
			if tm != 0 && i < r.max-1 {
				time.Sleep(tm)
			}
		}
	}
	return
}
