package interceptor

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// RequestLogger interceptor allows you to log the request you are about to send
func RequestLogger(logger logrus.FieldLogger) Interceptor {
	return func(req *http.Request) {
		go func(r *http.Request) {
			logger.Infoln("sending request", r.Method, r.URL.String(), r.Header)
			if req.Body == nil {
				return
			}
			body, err := r.GetBody()
			if err != nil {
				return
			}
			read, err := io.ReadAll(body)
			if err != nil {
				return
			}
			logger.Debugln("body", string(read))
		}(req)
	}
}
