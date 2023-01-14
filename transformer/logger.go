package transformer

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// RequestLogger is deprecated. Please use ResponseLogger
var RequestLogger = ResponseLogger

// ResponseLogger lets you doing some output
func ResponseLogger(logger logrus.FieldLogger) Transformer {
	return func(res *http.Response) error {
		go func(r *http.Response) {
			req := r.Request
			logger.Infoln("response", req.Method, req.URL.String(), req.Header, res.StatusCode)
		}(res)
		return nil
	}
}
