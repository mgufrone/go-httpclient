package httpclient

import (
	"github.com/mgufrone/go-httpclient/interceptor"
	"github.com/mgufrone/go-httpclient/transformer"
	"net/http"
)

// Client interface
type Client interface {
	// AddInterceptor allows to intercept request before it is sent.
	AddInterceptor(interceptor interceptor.Interceptor) Client
	// AddTransformer allows to parser the http response after being called
	AddTransformer(transformer transformer.Transformer) Client
	// ClearTransformers should clear all registered transformers
	ClearTransformers() Client
	// ClearInterceptors should clear all registered transformers
	ClearInterceptors() Client
	// SetTransporter will change the http RoundTripper
	SetTransporter(transporter http.RoundTripper) Client
	// RoundTrip is similar to `RoundTrip` method inside http.RoundTripper
	// but it will apply registered  interceptors then transform the response
	RoundTrip(r *http.Request) (*http.Response, error)
}
