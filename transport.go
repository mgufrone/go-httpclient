package httpclient

import (
	"net"
	"net/http"
	"time"
)

type TransportConstructor func() *http.Transport

var (
	DefaultTransporter = defaultTransporter()
)

func defaultTransporter() TransportConstructor {
	return func() *http.Transport {
		return &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}
}

func HTTP2Transporter() TransportConstructor {
	return func() *http.Transport {
		tp := DefaultTransporter()
		tp.ForceAttemptHTTP2 = true
		return tp
	}
}
