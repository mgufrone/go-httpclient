package interceptor

import "net/http"

type Interceptor func(req *http.Request)
