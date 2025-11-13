# Livepeer Go SDK

The Livepeer Go library provides convenient access to the Livepeer Studio API from applications written in Golang.

## Documentation

For full documentation and examples, please visit [docs.livepeer.org](https://docs.livepeer.org/sdks/go/).

<!-- Start Summary [summary] -->
## Summary

Livepeer API Reference: Welcome to the Livepeer API reference docs. Here you will find all the
endpoints exposed on the standard Livepeer API, learn how to use them and
what they return.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Livepeer Go SDK](#livepeer-go-sdk)
  * [Documentation](#documentation)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Retries](#retries)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/livepeer/livepeer-go
```
<!-- End SDK Installation [installation] -->

<!-- No SDK Example Usage [usage] -->
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


<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccessControl](docs/sdks/accesscontrol/README.md)

* [Create](docs/sdks/accesscontrol/README.md#create) - Create a signing key
* [GetAll](docs/sdks/accesscontrol/README.md#getall) - Retrieves signing keys
* [Delete](docs/sdks/accesscontrol/README.md#delete) - Delete Signing Key
* [Get](docs/sdks/accesscontrol/README.md#get) - Retrieves a signing key
* [Update](docs/sdks/accesscontrol/README.md#update) - Update a signing key

### [Asset](docs/sdks/asset/README.md)

* [GetAll](docs/sdks/asset/README.md#getall) - Retrieve assets
* [Create](docs/sdks/asset/README.md#create) - Upload an asset
* [CreateViaURL](docs/sdks/asset/README.md#createviaurl) - Upload asset via URL
* [Get](docs/sdks/asset/README.md#get) - Retrieves an asset
* [Update](docs/sdks/asset/README.md#update) - Patch an asset
* [Delete](docs/sdks/asset/README.md#delete) - Delete an asset

### [Generate](docs/sdks/generate/README.md)

* [TextToImage](docs/sdks/generate/README.md#texttoimage) - Text To Image
* [ImageToImage](docs/sdks/generate/README.md#imagetoimage) - Image To Image
* [ImageToVideo](docs/sdks/generate/README.md#imagetovideo) - Image To Video
* [Upscale](docs/sdks/generate/README.md#upscale) - Upscale
* [AudioToText](docs/sdks/generate/README.md#audiototext) - Audio To Text
* [SegmentAnything2](docs/sdks/generate/README.md#segmentanything2) - Segment Anything 2
* [Llm](docs/sdks/generate/README.md#llm) - LLM
* [ImageToText](docs/sdks/generate/README.md#imagetotext) - Image To Text
* [LiveVideoToVideo](docs/sdks/generate/README.md#livevideotovideo) - Live Video To Video
* [TextToSpeech](docs/sdks/generate/README.md#texttospeech) - Text To Speech

### [Metrics](docs/sdks/metrics/README.md)

* [GetRealtimeViewership](docs/sdks/metrics/README.md#getrealtimeviewership) - Query realtime viewership
* [GetViewership](docs/sdks/metrics/README.md#getviewership) - Query viewership metrics
* [GetCreatorViewership](docs/sdks/metrics/README.md#getcreatorviewership) - Query creator viewership metrics
* [GetPublicViewership](docs/sdks/metrics/README.md#getpublicviewership) - Query public total views metrics
* [GetUsage](docs/sdks/metrics/README.md#getusage) - Query usage metrics

### [Multistream](docs/sdks/multistream/README.md)

* [GetAll](docs/sdks/multistream/README.md#getall) - Retrieve Multistream Targets
* [Create](docs/sdks/multistream/README.md#create) - Create a multistream target
* [Get](docs/sdks/multistream/README.md#get) - Retrieve a multistream target
* [Update](docs/sdks/multistream/README.md#update) - Update Multistream Target
* [Delete](docs/sdks/multistream/README.md#delete) - Delete a multistream target

### [Playback](docs/sdks/playback/README.md)

* [Get](docs/sdks/playback/README.md#get) - Retrieve Playback Info

### [~~Room~~](docs/sdks/room/README.md)

* [~~Create~~](docs/sdks/room/README.md#create) - Create a room :warning: **Deprecated**
* [~~Get~~](docs/sdks/room/README.md#get) - Retrieve a room :warning: **Deprecated**
* [~~Delete~~](docs/sdks/room/README.md#delete) - Delete a room :warning: **Deprecated**
* [~~StartEgress~~](docs/sdks/room/README.md#startegress) - Start room RTMP egress :warning: **Deprecated**
* [~~StopEgress~~](docs/sdks/room/README.md#stopegress) - Stop room RTMP egress :warning: **Deprecated**
* [~~CreateUser~~](docs/sdks/room/README.md#createuser) - Create a room user :warning: **Deprecated**
* [~~GetUser~~](docs/sdks/room/README.md#getuser) - Get user details :warning: **Deprecated**
* [~~UpdateUser~~](docs/sdks/room/README.md#updateuser) - Update a room user :warning: **Deprecated**
* [~~DeleteUser~~](docs/sdks/room/README.md#deleteuser) - Remove a user from the room :warning: **Deprecated**

### [Session](docs/sdks/session/README.md)

* [GetClips](docs/sdks/session/README.md#getclips) - Retrieve clips of a session
* [GetAll](docs/sdks/session/README.md#getall) - Retrieve sessions
* [Get](docs/sdks/session/README.md#get) - Retrieve a session
* [GetRecorded](docs/sdks/session/README.md#getrecorded) - Retrieve Recorded Sessions

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

### [Task](docs/sdks/task/README.md)

* [GetAll](docs/sdks/task/README.md#getall) - Retrieve Tasks
* [Get](docs/sdks/task/README.md#get) - Retrieve a Task

### [Transcode](docs/sdks/transcode/README.md)

* [Create](docs/sdks/transcode/README.md#create) - Transcode a video

### [Webhook](docs/sdks/webhook/README.md)

* [GetAll](docs/sdks/webhook/README.md#getall) - Retrieve a Webhook
* [Create](docs/sdks/webhook/README.md#create) - Create a webhook
* [Get](docs/sdks/webhook/README.md#get) - Retrieve a webhook
* [Update](docs/sdks/webhook/README.md#update) - Update a webhook
* [Delete](docs/sdks/webhook/README.md#delete) - Delete a webhook
* [GetLogs](docs/sdks/webhook/README.md#getlogs) - Retrieve webhook logs
* [GetLog](docs/sdks/webhook/README.md#getlog) - Retrieve a webhook log
* [ResendLog](docs/sdks/webhook/README.md#resendlog) - Resend a webhook

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Get` function may return the following errors:

| Error Type         | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| sdkerrors.Error    | 404         | application/json |
| sdkerrors.SDKError | 4XX, 5XX    | \*/\*            |

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
	ctx := context.Background()

	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Playback.Get(ctx, "<id>")
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

<!-- No Server Selection [server] -->

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

	"github.com/livepeer/livepeer-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = livepeergo.New(livepeergo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type | Scheme      |
| -------- | ---- | ----------- |
| `APIKey` | http | HTTP Bearer |

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
	ctx := context.Background()

	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Stream.Create(ctx, components.NewStreamPayload{
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
			WebhookID: livepeergo.Pointer("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Pointer[float64](600),
		},
		Profiles: []components.FfmpegProfile{},
		Record:   livepeergo.Pointer(false),
		RecordingSpec: &components.NewStreamPayloadRecordingSpec{
			Profiles: []components.TranscodeProfile{
				components.TranscodeProfile{
					Width:   livepeergo.Pointer[int64](1280),
					Name:    livepeergo.Pointer("720p"),
					Height:  livepeergo.Pointer[int64](720),
					Bitrate: 3000000,
					Quality: livepeergo.Pointer[int64](23),
					Fps:     livepeergo.Pointer[int64](30),
					FpsDen:  livepeergo.Pointer[int64](1),
					Gop:     livepeergo.Pointer("2"),
					Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
					Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
				},
			},
		},
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile: "720p",
					ID:      livepeergo.Pointer("PUSH123"),
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Stream.Create(ctx, components.NewStreamPayload{
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
			WebhookID: livepeergo.Pointer("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Pointer[float64](600),
		},
		Profiles: []components.FfmpegProfile{},
		Record:   livepeergo.Pointer(false),
		RecordingSpec: &components.NewStreamPayloadRecordingSpec{
			Profiles: []components.TranscodeProfile{
				components.TranscodeProfile{
					Width:   livepeergo.Pointer[int64](1280),
					Name:    livepeergo.Pointer("720p"),
					Height:  livepeergo.Pointer[int64](720),
					Bitrate: 3000000,
					Quality: livepeergo.Pointer[int64](23),
					Fps:     livepeergo.Pointer[int64](30),
					FpsDen:  livepeergo.Pointer[int64](1),
					Gop:     livepeergo.Pointer("2"),
					Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
					Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
				},
			},
		},
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile: "720p",
					ID:      livepeergo.Pointer("PUSH123"),
				},
			},
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := livepeergo.New(
		livepeergo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Stream.Create(ctx, components.NewStreamPayload{
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
			WebhookID: livepeergo.Pointer("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Pointer[float64](600),
		},
		Profiles: []components.FfmpegProfile{},
		Record:   livepeergo.Pointer(false),
		RecordingSpec: &components.NewStreamPayloadRecordingSpec{
			Profiles: []components.TranscodeProfile{
				components.TranscodeProfile{
					Width:   livepeergo.Pointer[int64](1280),
					Name:    livepeergo.Pointer("720p"),
					Height:  livepeergo.Pointer[int64](720),
					Bitrate: 3000000,
					Quality: livepeergo.Pointer[int64](23),
					Fps:     livepeergo.Pointer[int64](30),
					FpsDen:  livepeergo.Pointer[int64](1),
					Gop:     livepeergo.Pointer("2"),
					Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
					Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
				},
			},
		},
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile: "720p",
					ID:      livepeergo.Pointer("PUSH123"),
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


