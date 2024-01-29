# Stream
(*Stream*)

### Available Operations

* [GetAll](#getall) - Retrieve streams
* [Create](#create) - Create a stream
* [Delete](#delete) - Delete a stream
* [Get](#get) - Retrieve a stream
* [Update](#update) - Update a stream
* [Terminate](#terminate) - Terminates a live stream
* [CreateClip](#createclip) - Create a clip
* [GetAllClips](#getallclips) - Retrieve clips of a livestream
* [CreateMultistreamTarget](#createmultistreamtarget) - Add a multistream target
* [DeleteMultistreamTarget](#deletemultistreamtarget) - Remove a multistream target

## GetAll

Retrieve streams

### Example Usage

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


    var streamsonly *string = "string"

    ctx := context.Background()
    res, err := s.Stream.GetAll(ctx, streamsonly)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
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
| sdkerrors.SDKError | 400-600            | */*                |

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
    res, err := s.Stream.Create(ctx, components.NewStreamPayload{
        Name: "test_stream",
        CreatorID: components.CreateInputCreatorIDCreatorID(
                components.CreateCreatorIDCreatorID1(
                            components.CreatorID1{
                                Type: components.CreatorIDTypeUnverified,
                                Value: "string",
                            },
                ),
        ),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeJwt,
            WebhookContext: map[string]interface{}{
                "key": "string",
            },
        },
        Profiles: []components.FfmpegProfile{
            components.FfmpegProfile{
                Width: 638424,
                Name: "720p",
                Height: 859213,
                Bitrate: 417458,
                Fps: 288408,
            },
        },
        Record: livepeer.Bool(false),
        Multistream: &components.Multistream{
            Targets: []components.Target{
                components.Target{
                    Profile: "720p",
                    Spec: &components.TargetSpec{
                        URL: "rtmps://live.my-service.tv/channel/secretKey",
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
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
| sdkerrors.SDKError | 400-600            | */*                |

## Delete


This will also suspend any active stream sessions, so make sure to wait
until the stream has finished. To explicitly interrupt an active
session, consider instead updating the suspended field in the stream
using the PATCH stream API.


### Example Usage

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


    var id string = "string"

    ctx := context.Background()
    res, err := s.Stream.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

## Get

Retrieve a stream

### Example Usage

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


    var id string = "string"

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
| sdkerrors.SDKError | 400-600            | */*                |

## Update

Update a stream

### Example Usage

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


    var id string = "string"

    streamPatchPayload := components.StreamPatchPayload{
        CreatorID: components.CreateInputCreatorIDStr(
        "string",
        ),
        Record: livepeer.Bool(false),
        Multistream: &components.Multistream{
            Targets: []components.Target{
                components.Target{
                    Profile: "720p",
                    Spec: &components.TargetSpec{
                        URL: "rtmps://live.my-service.tv/channel/secretKey",
                    },
                },
            },
        },
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypePublic,
            WebhookContext: map[string]interface{}{
                "key": "string",
            },
        },
        Profiles: []components.FfmpegProfile{
            components.FfmpegProfile{
                Width: 597129,
                Name: "720p",
                Height: 15652,
                Bitrate: 344620,
                Fps: 708455,
            },
        },
    }

    ctx := context.Background()
    res, err := s.Stream.Update(ctx, id, streamPatchPayload)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

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
	"context"
	"log"
	"livepeer"
	"livepeer/models/components"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity(""),
    )


    var id string = "string"

    ctx := context.Background()
    res, err := s.Stream.Terminate(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

## CreateClip

Create a clip

### Example Usage

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
    res, err := s.Stream.CreateClip(ctx, components.ClipPayload{
        PlaybackID: "string",
        StartTime: 9418.72,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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

**[*operations.PostClipResponse](../../models/operations/postclipresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAllClips

Retrieve clips of a livestream

### Example Usage

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


    var id string = "string"

    ctx := context.Background()
    res, err := s.Stream.GetAllClips(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
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

**[*operations.GetStreamIDClipsResponse](../../models/operations/getstreamidclipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateMultistreamTarget

Add a multistream target

### Example Usage

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


    var id string = "string"

    targetAddPayload := components.TargetAddPayload{
        Profile: "720p",
        Spec: &components.TargetAddPayloadSpec{
            URL: "rtmps://live.my-service.tv/channel/secretKey",
        },
    }

    ctx := context.Background()
    res, err := s.Stream.CreateMultistreamTarget(ctx, id, targetAddPayload)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteMultistreamTarget

Remove a multistream target

### Example Usage

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


    var id string = "string"

    var targetID string = "string"

    ctx := context.Background()
    res, err := s.Stream.DeleteMultistreamTarget(ctx, id, targetID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |
