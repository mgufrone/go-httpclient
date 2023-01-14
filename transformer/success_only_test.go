package transformer

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func mockResponse(statusCode int) *http.Response {
	body := io.NopCloser(strings.NewReader("random content"))
	return &http.Response{
		Body:       body,
		StatusCode: statusCode,
	}
}
func TestSuccessOnly(t *testing.T) {
	cases := []struct {
		in       *http.Response
		wantsErr bool
	}{
		{
			nil,
			false,
		},
		{
			mockResponse(101),
			true,
		},
		{
			mockResponse(301),
			true,
		},
		{
			mockResponse(301),
			true,
		},
		{
			mockResponse(404),
			true,
		},
		{
			mockResponse(200),
			false,
		},
		{
			mockResponse(201),
			false,
		},
	}
	for _, c := range cases {
		err := SuccessOnly(c.in)
		if c.wantsErr {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
	}
}
