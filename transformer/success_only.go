package transformer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func SuccessOnly(res *http.Response) error {
	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		var buff bytes.Buffer
		io.Copy(&buff, res.Body)
		byt := bytes.NewReader(buff.Bytes())
		by, _ := ioutil.ReadAll(byt)
		return fmt.Errorf("status: %d; response: %s", res.StatusCode, string(by))
	}
	return nil
}
