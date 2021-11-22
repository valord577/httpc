Httpc
======

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/valord577/httpc?t=2)](https://goreportcard.com/report/github.com/valord577/httpc)
[![Go Reference](https://pkg.go.dev/badge/github.com/valord577/httpc.svg)](https://pkg.go.dev/github.com/valord577/httpc)
[![GitHub](https://img.shields.io/github/license/valord577/httpc?t=0)](LICENSE)

A customizable and simple HTTP client library. Only depend on the stdlib HTTP client.

Requirements
------

- Go 1.14 or higher.

Features
------

- Simple and easy to use
- Make HTTP calls customizable

Installing
------

go mod:

```shell
go get github.com/valord577/httpc
```

Example
------

<details>
<summary>
- Do HTTP calls
</summary>

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/valord577/httpc"
)

func main() {
    c := httpc.PackedReq{
        URL:              "https://www.google.com",
        Method:           http.MethodGet,
        ReqBodyPublisher: httpc.PublisherNoBody{},
        RespBodyHandler:  httpc.RespBodyAsByteArray{},
    }

    bs, err := c.Send()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s", bs)
}
```
</details>

<details>
<summary>
- Customize the processing of response body
</summary>

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    
    "github.com/valord577/httpc"
)

type RespBodyAsString struct {}

func (r RespBodyAsString) Apply(body io.ReadCloser) (interface{}, error) {
    bs, err := io.ReadAll(body)
    if err != nil {
        return nil, err
    }
    return string(bs), nil
}

func main() {
    c := httpc.PackedReq{
        URL:              "https://www.google.com",
        Method:           http.MethodGet,
        ReqBodyPublisher: httpc.PublisherNoBody{},
        RespBodyHandler:  RespBodyAsString{},
    }

    bs, err := c.Send()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s", bs)
}
```
</details>

Changes
------

See the [CHANGES](CHANGE.md) for changes.

License
------

See the [LICENSE](LICENSE) for Rights and Limitations (MIT).
