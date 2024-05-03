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

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both. When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| sdkerrors.Error    | 404         | application/json |
| sdkerrors.SDKError | 4xx-5xx     | _/_              |

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
