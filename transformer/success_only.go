package transformer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// SuccessOnly transformer lets you set the transaction becoming an error if the status code response is lower than 200 or higher than 299.
func SuccessOnly(res *http.Response) error {
	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		var buff bytes.Buffer
		io.Copy(&buff, res.Body)
		byt := bytes.NewReader(buff.Bytes())
		by, _ := io.ReadAll(byt)
		return fmt.Errorf("status: %d; response: %s", res.StatusCode, string(by))
	}
	return nil
}
