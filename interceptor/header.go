package interceptor

import (
	"net/http"
)

// Header will set/replace header with the given key and value in the request before request is sent
func Header(key string, value string) Interceptor {
	if key == "" {
		return nil
	}
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}
