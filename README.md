# go-write

[![Go Reference](https://pkg.go.dev/badge/github.com/aidenwallis/go-write.svg)](https://pkg.go.dev/github.com/aidenwallis/go-write)

A simple, clean HTTP response builder for [net/http](https://pkg.go.dev/net/http). This package is deliberately very light, and simple. There are no external dependencies, just a really clean way to write handlers.

## Usage

You may interact with this package using the builder, or the code-generated functions. They behave the same.

The code-generated functions simply call the builder for you, they're just there for your convenience.

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
