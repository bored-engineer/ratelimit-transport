# ratelimit-transport [![Go Reference](https://pkg.go.dev/badge/github.com/bored-engineer/ratelimit-transport.svg)](https://pkg.go.dev/github.com/bored-engineer/ratelimit-transport)
A Golang [http.RoundTripper](https://pkg.go.dev/net/http#RoundTripper) that limits the rate (QPS) of HTTP requests using [github.com/uber-go/ratelimit](https://github.com/uber-go/ratelimit), ex:

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	ratelimit "github.com/bored-engineer/ratelimit-transport"
)

func main() {
	client := &http.Client{
		Transport: ratelimit.New(nil, 10, ratelimit.Per(time.Minute)),
	}
	for range 10 {
		resp, err := client.Get("https://example.com")
		if err != nil {
			panic(err)
		}
		fmt.Printf("(*http.Client).Get returned %d at %s\n", resp.StatusCode, time.Now())
	}
}
```
