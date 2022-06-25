package transformer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RequestLogger(logger *logrus.Entry) Transformer {
	fmt.Println()
	return func(res *http.Response) error {
		go func(r *http.Response) {
			req := r.Request
			logger.Infoln("request result", req.Method, req.URL.String(), req.Header, res.StatusCode)
		}(res)
		return nil
	}
}
