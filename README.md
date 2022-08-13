# go-write

[![codecov](https://codecov.io/gh/aidenwallis/go-write/branch/main/graph/badge.svg?token=BQOADNY183)](https://codecov.io/gh/aidenwallis/go-write) [![Go Reference](https://pkg.go.dev/badge/github.com/aidenwallis/go-write.svg)](https://pkg.go.dev/github.com/aidenwallis/go-write)

A simple, clean HTTP response builder for [net/http](https://pkg.go.dev/net/http). This package is deliberately very light, and simple. There are no external dependencies, just a really clean way to write handlers.

## Usage

You may interact with this package using the builder, or the code-generated functions. They behave the same.

The code-generated functions simply call the builder for you, they're just there for your convenience.

### Response types

There are 4 different ways you can currently end responses with go-write.

#### Bytes

You may send raw bytes using go-write, it will simply copy the bytes ot the body, and send your HTTP status code with a `Content-Length` header:

```go
err := write.New(w, http.StatusTeapot).Bytes([]byte("abc123"))
```

#### Empty

You may send empty responses using go-write, the body will be empty, and will only send a HTTP status code:

```go
err := write.New(w, http.StatusTeapot).Empty()
```

#### JSON

**`Content-Type`: `application/json; charset=utf-8`**

You may send JSON responses using go-write, the body will be marshalled JSON, it will set the corresponding `Content-Type` and `Content-Length` header, and send your HTTP status code:

```go

type myBody struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

err := write.New(w, http.StatusTeapot).JSON(&myBody{
    Key:   "foo",
    Value: "bar",
})
```

#### Text

**`Content-Type`: `text/plain; charset=utf-8`**

You may send plain text responses using go-write, it will simply copy the text to the body, set the `Content-Type` and `Content-Length` header, and send your HTTP status code:

```go
err := write.New(w, http.StatusTeapot).Text("example text!")
```

### Using the builder

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aidenwallis/go-write/write"
)

func main() {
	srv := &http.Server{
		Addr: ":1234",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			if err := write.New(w, http.StatusTeapot).Text("This handler is indeed a teapot..."); err != nil {
				log.Println("Failed to handle HTTP response: " + err.Error())
			}
		}),
	}
	defer srv.Shutdown(context.Background())

	go func() {
		log.Println("Server running.")
		_ = srv.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
```

### Using the code-generated functions

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aidenwallis/go-write/write"
)

func main() {
	srv := &http.Server{
		Addr: ":1234",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			if err := write.Teapot(w).Text("This handler is indeed a teapot..."); err != nil {
				log.Println("Failed to handle HTTP response: " + err.Error())
			}
		}),
	}
	defer srv.Shutdown(context.Background())

	go func() {
		log.Println("Server running.")
		_ = srv.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
```
