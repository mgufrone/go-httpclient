package interceptor

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func newReq() *http.Request {
	r, _ := http.NewRequest("GET", "/", nil)
	return r
}
func TestHeader(t *testing.T) {
	testCases := []struct {
		pre   *http.Request
		key   string
		value string
		eval  func(req *http.Request)
	}{
		{
			newReq(),
			"s",
			"",
			nil,
		},
		{
			newReq(),
			"a",
			"",
			func(req *http.Request) {
				require.Equal(t, "", req.Header.Get("a"))
			},
		},
	}
	for _, test := range testCases {
		inter := Header(test.key, test.value)
		inter(test.pre)
		if test.eval == nil {
			_, ok := test.pre.Header[test.key]
			require.False(t, ok)
			continue
		}
		test.eval(test.pre)
	}
}
