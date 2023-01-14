package httpclient

import (
	"github.com/mgufrone/go-httpclient/interceptor"
	"github.com/mgufrone/go-httpclient/transformer"
	"net/http"
)

type baseClient struct {
	transformers []transformer.Transformer
	interceptors []interceptor.Interceptor
	client       http.RoundTripper
}

func (b *baseClient) SetTransporter(transporter http.RoundTripper) Client {
	b.client = transporter
	return b
}

func (b *baseClient) ClearTransformers() Client {
	b.transformers = []transformer.Transformer{}
	return b
}

func (b *baseClient) ClearInterceptors() Client {
	b.interceptors = []interceptor.Interceptor{}
	return b
}

func (b *baseClient) RoundTrip(request *http.Request) (*http.Response, error) {
	for _, c := range b.interceptors {
		c(request)
	}
	if b.client == nil {
		b.client = DefaultTransporter()
	}
	res, err := b.client.RoundTrip(request)
	if err == nil {
		for _, t := range b.transformers {
			if err = t(res); err != nil {
				break
			}
		}
	}
	return res, err
}

// Do is a shorthand for RoundTrip
func (b *baseClient) Do(request *http.Request) (*http.Response, error) {
	return b.RoundTrip(request)
}

func (b *baseClient) AddInterceptor(interceptor interceptor.Interceptor) Client {
	b.interceptors = append(b.interceptors, interceptor)
	return b
}

func (b *baseClient) AddTransformer(transformer transformer.Transformer) Client {
	b.transformers = append(b.transformers, transformer)
	return b
}
