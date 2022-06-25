package transformer

import (
	"net/http"
)

type Transformer func(res *http.Response) error
