# Stream
(*Stream*)

## Overview

Operations related to livestream api

### Available Operations

* [Create](#create) - Create a stream
* [GetAll](#getall) - Retrieve streams
* [Get](#get) - Retrieve a stream
* [Update](#update) - Update a stream
* [Delete](#delete) - Delete a stream
* [Terminate](#terminate) - Terminates a live stream
* [StartPull](#startpull) - Start ingest for a pull stream
* [CreateClip](#createclip) - Create a clip
* [GetClips](#getclips) - Retrieve clips of a livestream
* [AddMultistreamTarget](#addmultistreamtarget) - Add a multistream target
* [RemoveMultistreamTarget](#removemultistreamtarget) - Remove a multistream target

## Create

The only parameter you are required to set is the name of your stream,
but we also highly recommend that you define transcoding profiles
parameter that suits your specific broadcasting configuration.
\
\
If you do not define transcoding rendition profiles when creating the
stream, a default set of profiles will be used. These profiles include
240p,  360p, 480p and 720p.
\
\
The playback policy is set to public by default for new streams. It can
also be added upon the creation of a new stream by adding
`"playbackPolicy": {"type": "jwt"}`


### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"context"
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
            Type: components.TypeWebhook,
            WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Float64(600),
        },
        Profiles: []components.FfmpegProfile{
            components.FfmpegProfile{
                Width: 1280,
                Name: "720p",
                Height: 486589,
                Bitrate: 3000000,
                Fps: 30,
                FpsDen: livepeergo.Int64(1),
                Quality: livepeergo.Int64(23),
                Gop: livepeergo.String("2"),
                Profile: components.ProfileH264Baseline.ToPointer(),
            },
        },
        Record: livepeergo.Bool(false),
        RecordingSpec: &components.RecordingSpec{
            Profiles: []components.FfmpegProfile{
                components.FfmpegProfile{
                    Width: 1280,
                    Name: "720p",
                    Height: 489382,
                    Bitrate: 3000000,
                    Fps: 30,
                    FpsDen: livepeergo.Int64(1),
                    Quality: livepeergo.Int64(23),
                    Gop: livepeergo.String("2"),
                    Profile: components.ProfileH264Baseline.ToPointer(),
                },
            },
        },
        Multistream: &components.Multistream{
            Targets: []components.Target{
                components.Target{
                    Profile: "720p",
                    VideoOnly: livepeergo.Bool(false),
                    ID: livepeergo.String("PUSH123"),
                    Spec: &components.TargetSpec{
                        Name: livepeergo.String("My target"),
                        URL: "rtmps://live.my-service.tv/channel/secretKey",
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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.NewStreamPayload](../../models/components/newstreampayload.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CreateStreamResponse](../../models/operations/createstreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAll

Retrieve streams

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var streamsonly *string = livepeergo.String("<value>")
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `streamsonly`                                         | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.GetStreamsResponse](../../models/operations/getstreamsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Retrieve a stream

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Stream.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the stream                                      |


### Response

**[*operations.GetStreamResponse](../../models/operations/getstreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update a stream

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"

    streamPatchPayload := components.StreamPatchPayload{
        Record: livepeergo.Bool(false),
        Multistream: &components.Multistream{
            Targets: []components.Target{
                components.Target{
                    Profile: "720p",
                    VideoOnly: livepeergo.Bool(false),
                    ID: livepeergo.String("PUSH123"),
                    Spec: &components.TargetSpec{
                        Name: livepeergo.String("My target"),
                        URL: "rtmps://live.my-service.tv/channel/secretKey",
                    },
                },
            },
        },
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Float64(600),
        },
        Profiles: []components.FfmpegProfile{
            components.FfmpegProfile{
                Width: 1280,
                Name: "720p",
                Height: 857478,
                Bitrate: 3000000,
                Fps: 30,
                FpsDen: livepeergo.Int64(1),
                Quality: livepeergo.Int64(23),
                Gop: livepeergo.String("2"),
                Profile: components.ProfileH264Baseline.ToPointer(),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Stream.Update(ctx, id, streamPatchPayload)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | ID of the stream                                                               |
| `streamPatchPayload`                                                           | [components.StreamPatchPayload](../../models/components/streampatchpayload.md) | :heavy_check_mark:                                                             | N/A                                                                            |


### Response

**[*operations.UpdateStreamResponse](../../models/operations/updatestreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

This will also suspend any active stream sessions, so make sure to wait
until the stream has finished. To explicitly interrupt an active
session, consider instead updating the suspended field in the stream
using the PATCH stream API.


### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Stream.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the stream                                      |


### Response

**[*operations.DeleteStreamResponse](../../models/operations/deletestreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Terminate

`DELETE /stream/{id}/terminate` can be used to terminate an ongoing
session on a live stream. Unlike suspending the stream, it allows the
streamer to restart streaming even immediately, but it will force
terminate the current session and stop the recording.
\
\
A 204 No Content status response indicates the stream was successfully
terminated.


### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Stream.Terminate(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the stream                                      |


### Response

**[*operations.TerminateStreamResponse](../../models/operations/terminatestreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## StartPull

`POST /stream/{id}/start-pull` can be used to start ingest for a stream
configured with a pull source. If the stream has recording configured,
it will also start recording.
\
\
A 204 No Content status response indicates the stream was successfully
started.


### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Stream.StartPull(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the stream                                      |


### Response

**[*operations.StartPullStreamResponse](../../models/operations/startpullstreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateClip

Create a clip

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.ClipPayload{
        PlaybackID: "eaw4nk06ts2d0mzb",
        StartTime: 1587667174725,
        EndTime: livepeergo.Float64(1587667174725),
        Name: livepeergo.String("My Clip"),
        SessionID: livepeergo.String("de7818e7-610a-4057-8f6f-b785dc1e6f88"),
    }
    ctx := context.Background()
    res, err := s.Stream.CreateClip(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.ClipPayload](../../models/components/clippayload.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.CreateClipResponse](../../models/operations/createclipresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetClips

Retrieve clips of a livestream

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Stream.GetClips(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `id`                                                   | *string*                                               | :heavy_check_mark:                                     | ID of the parent stream or playbackId of parent stream |


### Response

**[*operations.GetClipsResponse](../../models/operations/getclipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AddMultistreamTarget

Add a multistream target

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"

    targetAddPayload := components.TargetAddPayload{
        Profile: "720p0",
        VideoOnly: livepeergo.Bool(false),
        ID: livepeergo.String("PUSH123"),
        Spec: &components.TargetAddPayloadSpec{
            Name: livepeergo.String("My target"),
            URL: "rtmps://live.my-service.tv/channel/secretKey",
        },
    }
    ctx := context.Background()
    res, err := s.Stream.AddMultistreamTarget(ctx, id, targetAddPayload)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | ID of the parent stream                                                    |
| `targetAddPayload`                                                         | [components.TargetAddPayload](../../models/components/targetaddpayload.md) | :heavy_check_mark:                                                         | N/A                                                                        |


### Response

**[*operations.AddMultistreamTargetResponse](../../models/operations/addmultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveMultistreamTarget

Remove a multistream target

### Example Usage

```go
package main

import(
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var id string = "<value>"

    var targetID string = "<value>"
    ctx := context.Background()
    res, err := s.Stream.RemoveMultistreamTarget(ctx, id, targetID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the parent stream                               |
| `targetID`                                            | *string*                                              | :heavy_check_mark:                                    | ID of the multistream target                          |


### Response

**[*operations.RemoveMultistreamTargetResponse](../../models/operations/removemultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
