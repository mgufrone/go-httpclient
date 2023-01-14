## GO HTTP Client

A superset of go http client. It allows you to add interceptor and transformer.

### Table of contents
- [Features](#features)
- [Usage](#usage)
  - [Initialize Client](#initialize-client)
  - [Use Retryable Client](#use-retryable-client)
  - [Add Interceptor](#add-interceptor)
  - [Add Transformer](#add-transformer)

### Features
- Transformer
- Interceptor
- Retryable HTTP client


### Usage

#### Initialize client

To create a simple http client, you could use this snippet
```go
package main

import (
  "fmt"
  "github.com/mgufrone/go-httpclient"
  "net/http"
)

func main() {
  client := httpclient.Standard()
  req, _ := http.NewRequest("GET", "https://go.dev", nil)
  res, err := client.RoundTrip(req)
  // or
  res, err := client.Do(req) // it's a shorthand for RoundTrip
  fmt.Println(res, err)
}
```

If you would like to modify the transporter, you could use `WithTransporter` over `New`

```go
package main

import (
  "fmt"
  "github.com/mgufrone/go-httpclient"
  "net/http"
)

func main() {
  transporter := http.DefaultTransport
  // it needs an implementation of http.RoundTripper
  client := httpclient.New().SetTransporter(transporter)
  // or 
  client = httpclient.WithTransporter(transporter)
	
  req, _ := http.NewRequest("GET", "https://go.dev", nil)
  res, err := client.Do(req)
  fmt.Println(res, err)
}
```

#### Use Retryable client

Retryable client is an http client that will do retry 
under certain condition until it reached its maximum allowed retries. 
Here's how you'd define it if you want to use retryable client.

```go
package main

import (
  "fmt"
  "github.com/mgufrone/go-httpclient"
  "net/http"
  "time"
)

func main() {
  // it needs an implementation of http.RoundTripper
  client := httpclient.Retryable(10,  func(req *http.Request, response *http.Response) bool {
    return true
  }, func(reqTime int) time.Duration {
    return time.Millisecond
  })
  req, _ := http.NewRequest("GET", "https://go.dev", nil)
  res, err := client.Do(req)
  fmt.Println(res, err)
}
```

#### Add Interceptor

Interceptor is a hook that will be evaluated before sending the request. 
It could be used to inject some headers or extend the payload request.

For example, let's say you want to inject custom header in the request, you could do it this way

```go
package main

import (
  "fmt"
  "github.com/mgufrone/go-httpclient"
  "github.com/mgufrone/go-httpclient/interceptor"
  "net/http"
)

func main() {
  // it needs an implementation of http.RoundTripper
  client := httpclient.New().AddInterceptor(interceptor.Header("X-App-Id", "somelongid"))
  req, _ := http.NewRequest("GET", "https://go.dev", nil)
  res, err := client.Do(req)
  // the sent request will then contains header X-App-ID: somelongid
  fmt.Println(res, err)
}
```

#### Add Transformer

Transformer is a hook that will be evaluated after receiving response.
It could be used to verify the response or change the error message.

```go
package main

import (
  "fmt"
  "github.com/mgufrone/go-httpclient"
  "github.com/mgufrone/go-httpclient/transformer"
  "net/http"
)

func main() {
  // it needs an implementation of http.RoundTripper
  client := httpclient.New().AddInterceptor(transformer.SuccessOnly)
  req, _ := http.NewRequest("GET", "https://go.dev", nil)
  res, err := client.Do(req)
  // if the response is lower than 200 or higher than 299, the error would not be nil
  fmt.Println(res, err)
}
```
