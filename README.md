# Livepeer Go SDK

The Livepeer Go library provides convenient access to the Livepeer Studio API from applications written in Golang.

## Documentation

For full documentation and examples, please visit [docs.livepeer.org](https://docs.livepeer.org/sdks/go/).

## SDK Installation

```bash
go get github.com/livepeer/livepeer-go
```

## SDK Example Usage

### Example

```go
package main

import (
	"context"
	livepeer "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeer.New(
		livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Stream.Create(ctx, components.NewStreamPayload{
		Name: "test_stream",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```

## Available Resources and Operations

### [Stream](docs/sdks/stream/README.md)

- [Create](docs/sdks/stream/README.md#create) - Create a stream
- [GetAll](docs/sdks/stream/README.md#getall) - Retrieve streams
- [Get](docs/sdks/stream/README.md#get) - Retrieve a stream
- [Update](docs/sdks/stream/README.md#update) - Update a stream
- [Delete](docs/sdks/stream/README.md#delete) - Delete a stream
- [Terminate](docs/sdks/stream/README.md#terminate) - Terminates a live stream
- [StartPull](docs/sdks/stream/README.md#startpull) - Start ingest for a pull stream
- [CreateClip](docs/sdks/stream/README.md#createclip) - Create a clip
- [GetClips](docs/sdks/stream/README.md#getclips) - Retrieve clips of a livestream
- [AddMultistreamTarget](docs/sdks/stream/README.md#addmultistreamtarget) - Add a multistream target
- [RemoveMultistreamTarget](docs/sdks/stream/README.md#removemultistreamtarget) - Remove a multistream target

### [Multistream](docs/sdks/multistream/README.md)

- [GetAll](docs/sdks/multistream/README.md#getall) - Retrieve Multistream Targets
- [Create](docs/sdks/multistream/README.md#create) - Create a multistream target
- [Get](docs/sdks/multistream/README.md#get) - Retrieve a multistream target
- [Update](docs/sdks/multistream/README.md#update) - Update Multistream Target
- [Delete](docs/sdks/multistream/README.md#delete) - Delete a multistream target

### [Webhook](docs/sdks/webhook/README.md)

- [GetAll](docs/sdks/webhook/README.md#getall) - Retrieve a Webhook
- [Create](docs/sdks/webhook/README.md#create) - Create a webhook
- [Get](docs/sdks/webhook/README.md#get) - Retrieve a webhook
- [Update](docs/sdks/webhook/README.md#update) - Update a webhook
- [Delete](docs/sdks/webhook/README.md#delete) - Delete a webhook
- [GetLogs](docs/sdks/webhook/README.md#getlogs) - Retrieve webhook logs
- [GetLog](docs/sdks/webhook/README.md#getlog) - Retrieve a webhook log
- [ResendLog](docs/sdks/webhook/README.md#resendlog) - Resend a webhook

### [Asset](docs/sdks/asset/README.md)

- [GetAll](docs/sdks/asset/README.md#getall) - Retrieve assets
- [Create](docs/sdks/asset/README.md#create) - Upload an asset
- [CreateViaURL](docs/sdks/asset/README.md#createviaurl) - Upload asset via URL
- [Get](docs/sdks/asset/README.md#get) - Retrieves an asset
- [Update](docs/sdks/asset/README.md#update) - Patch an asset
- [Delete](docs/sdks/asset/README.md#delete) - Delete an asset

### [Session](docs/sdks/session/README.md)

- [GetClips](docs/sdks/session/README.md#getclips) - Retrieve clips of a session
- [GetAll](docs/sdks/session/README.md#getall) - Retrieve sessions
- [Get](docs/sdks/session/README.md#get) - Retrieve a session
- [GetRecorded](docs/sdks/session/README.md#getrecorded) - Retrieve Recorded Sessions

### [Room](docs/sdks/room/README.md)

- [~~Create~~](docs/sdks/room/README.md#create) - Create a room :warning: **Deprecated**
- [~~Get~~](docs/sdks/room/README.md#get) - Retrieve a room :warning: **Deprecated**
- [~~Delete~~](docs/sdks/room/README.md#delete) - Delete a room :warning: **Deprecated**
- [~~StartEgress~~](docs/sdks/room/README.md#startegress) - Start room RTMP egress :warning: **Deprecated**
- [~~StopEgress~~](docs/sdks/room/README.md#stopegress) - Stop room RTMP egress :warning: **Deprecated**
- [~~CreateUser~~](docs/sdks/room/README.md#createuser) - Create a room user :warning: **Deprecated**
- [~~GetUser~~](docs/sdks/room/README.md#getuser) - Get user details :warning: **Deprecated**
- [~~UpdateUser~~](docs/sdks/room/README.md#updateuser) - Update a room user :warning: **Deprecated**
- [~~DeleteUser~~](docs/sdks/room/README.md#deleteuser) - Remove a user from the room :warning: **Deprecated**

### [Metrics](docs/sdks/metrics/README.md)

- [GetViewership](docs/sdks/metrics/README.md#getviewership) - Query viewership metrics
- [GetCreatorViewership](docs/sdks/metrics/README.md#getcreatorviewership) - Query creator viewership metrics
- [GetPublicViewership](docs/sdks/metrics/README.md#getpublicviewership) - Query public total views metrics
- [GetUsage](docs/sdks/metrics/README.md#getusage) - Query usage metrics

### [AccessControl](docs/sdks/accesscontrol/README.md)

- [Create](docs/sdks/accesscontrol/README.md#create) - Create a signing key
- [GetAll](docs/sdks/accesscontrol/README.md#getall) - Retrieves signing keys
- [Delete](docs/sdks/accesscontrol/README.md#delete) - Delete Signing Key
- [Get](docs/sdks/accesscontrol/README.md#get) - Retrieves a signing key
- [Update](docs/sdks/accesscontrol/README.md#update) - Update a signing key

### [Task](docs/sdks/task/README.md)

- [GetAll](docs/sdks/task/README.md#getall) - Retrieve Tasks
- [Get](docs/sdks/task/README.md#get) - Retrieve a Task

### [Transcode](docs/sdks/transcode/README.md)

- [Create](docs/sdks/transcode/README.md#create) - Transcode a video

### [Playback](docs/sdks/playback/README.md)

- [Get](docs/sdks/playback/README.md#get) - Retrieve Playback Info
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var id string = "<value>"
	ctx := context.Background()
	res, err := s.Playback.Get(ctx, id)
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://livepeer.studio/api` | None |

#### Example

```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithServerIndex(0),
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := components.NewStreamPayload{
		Name: "test_stream",
		Pull: &components.Pull{
			Source: "https://myservice.com/live/stream.flv",
			Headers: map[string]string{
				"Authorization": "Bearer 123",
			},
			Location: &components.Location{
				Lat: 39.739,
				Lon: -104.988,
			},
		},
		PlaybackPolicy: &components.PlaybackPolicy{
			Type:      components.TypeWebhook,
			WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Float64(600),
		},
		Profiles: []components.FfmpegProfile{
			components.FfmpegProfile{
				Width:   1280,
				Name:    "720p",
				Height:  486589,
				Bitrate: 3000000,
				Fps:     30,
				FpsDen:  livepeergo.Int64(1),
				Quality: livepeergo.Int64(23),
				Gop:     livepeergo.String("2"),
				Profile: components.ProfileH264Baseline.ToPointer(),
			},
		},
		Record: livepeergo.Bool(false),
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile:   "720p",
					VideoOnly: livepeergo.Bool(false),
					ID:        livepeergo.String("PUSH123"),
					Spec: &components.TargetSpec{
						Name: livepeergo.String("My target"),
						URL:  "rtmps://live.my-service.tv/channel/secretKey",
					},
				},
			},
		},
	}
	ctx := context.Background()
	res, err := s.Stream.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithServerURL("https://livepeer.studio/api"),
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := components.NewStreamPayload{
		Name: "test_stream",
		Pull: &components.Pull{
			Source: "https://myservice.com/live/stream.flv",
			Headers: map[string]string{
				"Authorization": "Bearer 123",
			},
			Location: &components.Location{
				Lat: 39.739,
				Lon: -104.988,
			},
		},
		PlaybackPolicy: &components.PlaybackPolicy{
			Type:      components.TypeWebhook,
			WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Float64(600),
		},
		Profiles: []components.FfmpegProfile{
			components.FfmpegProfile{
				Width:   1280,
				Name:    "720p",
				Height:  486589,
				Bitrate: 3000000,
				Fps:     30,
				FpsDen:  livepeergo.Int64(1),
				Quality: livepeergo.Int64(23),
				Gop:     livepeergo.String("2"),
				Profile: components.ProfileH264Baseline.ToPointer(),
			},
		},
		Record: livepeergo.Bool(false),
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile:   "720p",
					VideoOnly: livepeergo.Bool(false),
					ID:        livepeergo.String("PUSH123"),
					Spec: &components.TargetSpec{
						Name: livepeergo.String("My target"),
						URL:  "rtmps://live.my-service.tv/channel/secretKey",
					},
				},
			},
		},
	}
	ctx := context.Background()
	res, err := s.Stream.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
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
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `APIKey`    | http        | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := components.NewStreamPayload{
		Name: "test_stream",
		Pull: &components.Pull{
			Source: "https://myservice.com/live/stream.flv",
			Headers: map[string]string{
				"Authorization": "Bearer 123",
			},
			Location: &components.Location{
				Lat: 39.739,
				Lon: -104.988,
			},
		},
		PlaybackPolicy: &components.PlaybackPolicy{
			Type:      components.TypeWebhook,
			WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Float64(600),
		},
		Profiles: []components.FfmpegProfile{
			components.FfmpegProfile{
				Width:   1280,
				Name:    "720p",
				Height:  486589,
				Bitrate: 3000000,
				Fps:     30,
				FpsDen:  livepeergo.Int64(1),
				Quality: livepeergo.Int64(23),
				Gop:     livepeergo.String("2"),
				Profile: components.ProfileH264Baseline.ToPointer(),
			},
		},
		Record: livepeergo.Bool(false),
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile:   "720p",
					VideoOnly: livepeergo.Bool(false),
					ID:        livepeergo.String("PUSH123"),
					Spec: &components.TargetSpec{
						Name: livepeergo.String("My target"),
						URL:  "rtmps://live.my-service.tv/channel/secretKey",
					},
				},
			},
		},
	}
	ctx := context.Background()
	res, err := s.Stream.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/livepeer/livepeer-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := components.NewStreamPayload{
		Name: "test_stream",
		Pull: &components.Pull{
			Source: "https://myservice.com/live/stream.flv",
			Headers: map[string]string{
				"Authorization": "Bearer 123",
			},
			Location: &components.Location{
				Lat: 39.739,
				Lon: -104.988,
			},
		},
		PlaybackPolicy: &components.PlaybackPolicy{
			Type:      components.TypeWebhook,
			WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Float64(600),
		},
		Profiles: []components.FfmpegProfile{
			components.FfmpegProfile{
				Width:   1280,
				Name:    "720p",
				Height:  486589,
				Bitrate: 3000000,
				Fps:     30,
				FpsDen:  livepeergo.Int64(1),
				Quality: livepeergo.Int64(23),
				Gop:     livepeergo.String("2"),
				Profile: components.ProfileH264Baseline.ToPointer(),
			},
		},
		Record: livepeergo.Bool(false),
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile:   "720p",
					VideoOnly: livepeergo.Bool(false),
					ID:        livepeergo.String("PUSH123"),
					Spec: &components.TargetSpec{
						Name: livepeergo.String("My target"),
						URL:  "rtmps://live.my-service.tv/channel/secretKey",
					},
				},
			},
		},
	}
	ctx := context.Background()
	res, err := s.Stream.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Stream](docs/sdks/stream/README.md)

* [Create](docs/sdks/stream/README.md#create) - Create a stream
* [GetAll](docs/sdks/stream/README.md#getall) - Retrieve streams
* [Get](docs/sdks/stream/README.md#get) - Retrieve a stream
* [Update](docs/sdks/stream/README.md#update) - Update a stream
* [Delete](docs/sdks/stream/README.md#delete) - Delete a stream
* [Terminate](docs/sdks/stream/README.md#terminate) - Terminates a live stream
* [StartPull](docs/sdks/stream/README.md#startpull) - Start ingest for a pull stream
* [CreateClip](docs/sdks/stream/README.md#createclip) - Create a clip
* [GetClips](docs/sdks/stream/README.md#getclips) - Retrieve clips of a livestream
* [AddMultistreamTarget](docs/sdks/stream/README.md#addmultistreamtarget) - Add a multistream target
* [RemoveMultistreamTarget](docs/sdks/stream/README.md#removemultistreamtarget) - Remove a multistream target

### [Multistream](docs/sdks/multistream/README.md)

* [GetAll](docs/sdks/multistream/README.md#getall) - Retrieve Multistream Targets
* [Create](docs/sdks/multistream/README.md#create) - Create a multistream target
* [Get](docs/sdks/multistream/README.md#get) - Retrieve a multistream target
* [Update](docs/sdks/multistream/README.md#update) - Update Multistream Target
* [Delete](docs/sdks/multistream/README.md#delete) - Delete a multistream target

### [Webhook](docs/sdks/webhook/README.md)

* [GetAll](docs/sdks/webhook/README.md#getall) - Retrieve a Webhook
* [Create](docs/sdks/webhook/README.md#create) - Create a webhook
* [Get](docs/sdks/webhook/README.md#get) - Retrieve a webhook
* [Update](docs/sdks/webhook/README.md#update) - Update a webhook
* [Delete](docs/sdks/webhook/README.md#delete) - Delete a webhook
* [GetLogs](docs/sdks/webhook/README.md#getlogs) - Retrieve webhook logs
* [GetLog](docs/sdks/webhook/README.md#getlog) - Retrieve a webhook log
* [ResendLog](docs/sdks/webhook/README.md#resendlog) - Resend a webhook

### [Asset](docs/sdks/asset/README.md)

* [GetAll](docs/sdks/asset/README.md#getall) - Retrieve assets
* [Create](docs/sdks/asset/README.md#create) - Upload an asset
* [CreateViaURL](docs/sdks/asset/README.md#createviaurl) - Upload asset via URL
* [Get](docs/sdks/asset/README.md#get) - Retrieves an asset
* [Update](docs/sdks/asset/README.md#update) - Patch an asset
* [Delete](docs/sdks/asset/README.md#delete) - Delete an asset

### [Session](docs/sdks/session/README.md)

* [GetClips](docs/sdks/session/README.md#getclips) - Retrieve clips of a session
* [GetAll](docs/sdks/session/README.md#getall) - Retrieve sessions
* [Get](docs/sdks/session/README.md#get) - Retrieve a session
* [GetRecorded](docs/sdks/session/README.md#getrecorded) - Retrieve Recorded Sessions

### [Room](docs/sdks/room/README.md)

* [~~Create~~](docs/sdks/room/README.md#create) - Create a room :warning: **Deprecated**
* [~~Get~~](docs/sdks/room/README.md#get) - Retrieve a room :warning: **Deprecated**
* [~~Delete~~](docs/sdks/room/README.md#delete) - Delete a room :warning: **Deprecated**
* [~~StartEgress~~](docs/sdks/room/README.md#startegress) - Start room RTMP egress :warning: **Deprecated**
* [~~StopEgress~~](docs/sdks/room/README.md#stopegress) - Stop room RTMP egress :warning: **Deprecated**
* [~~CreateUser~~](docs/sdks/room/README.md#createuser) - Create a room user :warning: **Deprecated**
* [~~GetUser~~](docs/sdks/room/README.md#getuser) - Get user details :warning: **Deprecated**
* [~~UpdateUser~~](docs/sdks/room/README.md#updateuser) - Update a room user :warning: **Deprecated**
* [~~DeleteUser~~](docs/sdks/room/README.md#deleteuser) - Remove a user from the room :warning: **Deprecated**

### [Metrics](docs/sdks/metrics/README.md)

* [GetViewership](docs/sdks/metrics/README.md#getviewership) - Query viewership metrics
* [GetCreatorViewership](docs/sdks/metrics/README.md#getcreatorviewership) - Query creator viewership metrics
* [GetPublicViewership](docs/sdks/metrics/README.md#getpublicviewership) - Query public total views metrics
* [GetUsage](docs/sdks/metrics/README.md#getusage) - Query usage metrics

### [AccessControl](docs/sdks/accesscontrol/README.md)

* [Create](docs/sdks/accesscontrol/README.md#create) - Create a signing key
* [GetAll](docs/sdks/accesscontrol/README.md#getall) - Retrieves signing keys
* [Delete](docs/sdks/accesscontrol/README.md#delete) - Delete Signing Key
* [Get](docs/sdks/accesscontrol/README.md#get) - Retrieves a signing key
* [Update](docs/sdks/accesscontrol/README.md#update) - Update a signing key

### [Task](docs/sdks/task/README.md)

* [GetAll](docs/sdks/task/README.md#getall) - Retrieve Tasks
* [Get](docs/sdks/task/README.md#get) - Retrieve a Task

### [Transcode](docs/sdks/transcode/README.md)

* [Create](docs/sdks/transcode/README.md#create) - Transcode a video

### [Playback](docs/sdks/playback/README.md)

* [Get](docs/sdks/playback/README.md#get) - Retrieve Playback Info
<!-- End Available Resources and Operations [operations] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


