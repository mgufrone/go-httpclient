package httpclient

import (
	"github.com/mgufrone/go-httpclient/interceptor"
	"github.com/mgufrone/go-httpclient/transformer"
	"net/http"
)

func New() Client {
	return &baseClient{
		transformers: []transformer.Transformer{},
		interceptors: []interceptor.Interceptor{},
		client:       DefaultTransporter(),
	}
}

func Standard() Client {
	cli := New()
	cli.AddTransformer(transformer.SuccessOnly)
	return cli
}

func WithTransporter(transporter http.RoundTripper) Client {
	return &baseClient{
		transformers: []transformer.Transformer{},
		interceptors: []interceptor.Interceptor{},
		client:       transporter,
	}
}
