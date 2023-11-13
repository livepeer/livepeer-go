# Livepeer Go Library

The Livepeer Go library provides convenient access to the Livepeer Studio API from applications written in Golang.

## Documentation

For full documentation and examples, please visit [docs.livepeer.org](https://docs.livepeer.org/sdks/go/).

## Installation

Install the package:

```bash
go get github.com/livepeer/livepeer-go
```

## Usage

The library needs to be configured with your Livepeer Studio account's API key, which is available in the [Studio Dashboard](httpss://livepeer.studio)

```go
package main

import(
	"context"
	"log"
	"livepeer"
	"livepeer/models/components"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity(""),
    )


    ctx := context.Background()
    res, err := s.Stream.GetStreams(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/livepeer/livepeer-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = livepeer.New(livepeer.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
