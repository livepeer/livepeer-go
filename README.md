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
    res, err := s.Stream.GetAll(ctx)
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	"livepeer"
	"livepeer/models/components"
	"log"
)

func main() {
	s := livepeer.New(
		livepeer.WithServerURL("https://livepeer.studio/api"),
		livepeer.WithSecurity(""),
	)

	var streamsonly *string = "string"

	ctx := context.Background()
	res, err := s.Stream.GetAll(ctx, streamsonly)
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```

### [Stream](docs/sdks/stream/README.md)

- [GetAll](docs/sdks/stream/README.md#getall) - Retrieve streams
- [Create](docs/sdks/stream/README.md#create) - Create a stream
- [Delete](docs/sdks/stream/README.md#delete) - Delete a stream
- [Get](docs/sdks/stream/README.md#get) - Retrieve a stream
- [Update](docs/sdks/stream/README.md#update) - Update a stream
- [CreateClip](docs/sdks/stream/README.md#createclip) - Create a clip
- [GetAllClips](docs/sdks/stream/README.md#getallclips) - Retrieve clips of a livestream

### [MultistreamTarget](docs/sdks/multistreamtarget/README.md)

- [GetAll](docs/sdks/multistreamtarget/README.md#getall) - Retrieve Multistream Targets
- [Create](docs/sdks/multistreamtarget/README.md#create) - Create a multistream target
- [Delete](docs/sdks/multistreamtarget/README.md#delete) - Delete a multistream target
- [Get](docs/sdks/multistreamtarget/README.md#get) - Retrieve a multistream target
- [Update](docs/sdks/multistreamtarget/README.md#update) - Update Multistream Target

### [Webhook](docs/sdks/webhook/README.md)

- [GetAll](docs/sdks/webhook/README.md#getall) - Retrieve a Webhook
- [Create](docs/sdks/webhook/README.md#create) - Create a webhook
- [Delete](docs/sdks/webhook/README.md#delete) - Delete a webhook
- [Get](docs/sdks/webhook/README.md#get) - Retrieve a webhook
- [Update](docs/sdks/webhook/README.md#update) - Update a webhook

### [Asset](docs/sdks/asset/README.md)

- [GetAll](docs/sdks/asset/README.md#getall) - Retrieve assets
- [Create](docs/sdks/asset/README.md#create) - Upload an asset
- [CreateViaURL](docs/sdks/asset/README.md#createviaurl) - Upload asset via URL
- [Delete](docs/sdks/asset/README.md#delete) - Delete an asset
- [Get](docs/sdks/asset/README.md#get) - Retrieves an asset
- [Update](docs/sdks/asset/README.md#update) - Update an asset

### [Metrics](docs/sdks/metrics/README.md)

- [GetViewership](docs/sdks/metrics/README.md#getviewership) - Query viewership metrics
- [GetCreatorViewership](docs/sdks/metrics/README.md#getcreatorviewership) - Query creator viewership metrics
- [GetPublicTotalViews](docs/sdks/metrics/README.md#getpublictotalviews) - Query public total views metrics
- [GetUsage](docs/sdks/metrics/README.md#getusage) - Query usage metrics

### [Session](docs/sdks/session/README.md)

- [GetAll](docs/sdks/session/README.md#getall) - Retrieve sessions
- [Get](docs/sdks/session/README.md#get) - Retrieve a session
- [GetRecorded](docs/sdks/session/README.md#getrecorded) - Retrieve Recorded Sessions
- [GetAllClips](docs/sdks/session/README.md#getallclips) - Retrieve clips of a session

### [AccessControl](docs/sdks/accesscontrol/README.md)

- [GetSigningKeys](docs/sdks/accesscontrol/README.md#getsigningkeys) - Retrieves signing keys
- [CreateSigningKey](docs/sdks/accesscontrol/README.md#createsigningkey) - Create a signing key
- [DeleteSigningKey](docs/sdks/accesscontrol/README.md#deletesigningkey) - Delete Signing Key
- [GetSigningKey](docs/sdks/accesscontrol/README.md#getsigningkey) - Retrieves a signing key
- [UpdateSigningKey](docs/sdks/accesscontrol/README.md#updatesigningkey) - Update a signing key

### [Task](docs/sdks/task/README.md)

- [GetAll](docs/sdks/task/README.md#getall) - Retrieve Tasks
- [Get](docs/sdks/task/README.md#get) - Retrieve a Task

### [Transcode](docs/sdks/transcode/README.md)

- [Create](docs/sdks/transcode/README.md#create) - Transcode a video

### [Playback](docs/sdks/playback/README.md)

- [Get](docs/sdks/playback/README.md#get) - Retrieve Playback Info
