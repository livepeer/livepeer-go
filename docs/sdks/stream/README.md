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

<!-- UsageSnippet language="go" operationID="createStream" method="post" path="/stream" -->
```go
package main

import(
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
            Type: components.TypeWebhook,
            WebhookID: livepeergo.Pointer("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Pointer[float64](600),
        },
        Profiles: []components.FfmpegProfile{},
        Record: livepeergo.Pointer(false),
        RecordingSpec: &components.NewStreamPayloadRecordingSpec{
            Profiles: []components.TranscodeProfile{
                components.TranscodeProfile{
                    Width: livepeergo.Pointer[int64](1280),
                    Name: livepeergo.Pointer("720p"),
                    Height: livepeergo.Pointer[int64](720),
                    Bitrate: 3000000,
                    Quality: livepeergo.Pointer[int64](23),
                    Fps: livepeergo.Pointer[int64](30),
                    FpsDen: livepeergo.Pointer[int64](1),
                    Gop: livepeergo.Pointer("2"),
                    Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
                    Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
                },
            },
        },
        Multistream: &components.Multistream{
            Targets: []components.Target{
                components.Target{
                    Profile: "720p",
                    ID: livepeergo.Pointer("PUSH123"),
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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.NewStreamPayload](../../models/components/newstreampayload.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateStreamResponse](../../models/operations/createstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAll

Retrieve streams

### Example Usage

<!-- UsageSnippet language="go" operationID="getStreams" method="get" path="/stream" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.GetAll(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `streamsonly`                                            | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStreamsResponse](../../models/operations/getstreamsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a stream

### Example Usage

<!-- UsageSnippet language="go" operationID="getStream" method="get" path="/stream/{id}" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the stream                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStreamResponse](../../models/operations/getstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a stream

### Example Usage

<!-- UsageSnippet language="go" operationID="updateStream" method="patch" path="/stream/{id}" -->
```go
package main

import(
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

    res, err := s.Stream.Update(ctx, "<id>", components.StreamPatchPayload{
        Record: livepeergo.Pointer(false),
        Multistream: &components.Multistream{
            Targets: []components.Target{
                components.Target{
                    Profile: "720p",
                    ID: livepeergo.Pointer("PUSH123"),
                },
            },
        },
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeergo.Pointer("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Pointer[float64](600),
        },
        Profiles: nil,
        RecordingSpec: &components.RecordingSpec{
            Profiles: []components.TranscodeProfile{
                components.TranscodeProfile{
                    Width: livepeergo.Pointer[int64](1280),
                    Name: livepeergo.Pointer("720p"),
                    Height: livepeergo.Pointer[int64](720),
                    Bitrate: 3000000,
                    Quality: livepeergo.Pointer[int64](23),
                    Fps: livepeergo.Pointer[int64](30),
                    FpsDen: livepeergo.Pointer[int64](1),
                    Gop: livepeergo.Pointer("2"),
                    Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
                    Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
                },
            },
        },
        Name: livepeergo.Pointer("test_stream"),
    })
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
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateStreamResponse](../../models/operations/updatestreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

This will also suspend any active stream sessions, so make sure to wait
until the stream has finished. To explicitly interrupt an active
session, consider instead updating the suspended field in the stream
using the PATCH stream API.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteStream" method="delete" path="/stream/{id}" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the stream                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteStreamResponse](../../models/operations/deletestreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

<!-- UsageSnippet language="go" operationID="terminateStream" method="delete" path="/stream/{id}/terminate" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.Terminate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the stream                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TerminateStreamResponse](../../models/operations/terminatestreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## StartPull

`POST /stream/{id}/start-pull` can be used to start ingest for a stream
configured with a pull source. If the stream has recording configured,
it will also start recording.
\
\
A 204 No Content status response indicates the stream was successfully
started.


### Example Usage

<!-- UsageSnippet language="go" operationID="startPullStream" method="post" path="/stream/{id}/start-pull" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.StartPull(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the stream                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.StartPullStreamResponse](../../models/operations/startpullstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateClip

Create a clip

### Example Usage

<!-- UsageSnippet language="go" operationID="createClip" method="post" path="/clip" -->
```go
package main

import(
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

    res, err := s.Stream.CreateClip(ctx, components.ClipPayload{
        PlaybackID: "eaw4nk06ts2d0mzb",
        StartTime: 1587667174725,
        EndTime: livepeergo.Pointer[float64](1587667174725),
        Name: livepeergo.Pointer("My Clip"),
        SessionID: livepeergo.Pointer("de7818e7-610a-4057-8f6f-b785dc1e6f88"),
    })
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
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreateClipResponse](../../models/operations/createclipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetClips

Retrieve clips of a livestream

### Example Usage

<!-- UsageSnippet language="go" operationID="getClips" method="get" path="/stream/{id}/clips" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.GetClips(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the parent stream or playbackId of parent stream   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetClipsResponse](../../models/operations/getclipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddMultistreamTarget

Add a multistream target

### Example Usage

<!-- UsageSnippet language="go" operationID="addMultistreamTarget" method="post" path="/stream/{id}/create-multistream-target" -->
```go
package main

import(
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

    res, err := s.Stream.AddMultistreamTarget(ctx, "<id>", components.TargetAddPayload{
        Profile: "720p0",
        ID: livepeergo.Pointer("PUSH123"),
        Spec: &components.TargetAddPayloadSpec{
            Name: livepeergo.Pointer("My target"),
            URL: "rtmps://live.my-service.tv/channel/secretKey",
        },
    })
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
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.AddMultistreamTargetResponse](../../models/operations/addmultistreamtargetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMultistreamTarget

Remove a multistream target

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMultistreamTarget" method="delete" path="/stream/{id}/multistream/{targetId}" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Stream.RemoveMultistreamTarget(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the parent stream                                  |
| `targetID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the multistream target                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RemoveMultistreamTargetResponse](../../models/operations/removemultistreamtargetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |