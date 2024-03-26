# Livepeer SDK

## Overview

Livepeer API Reference: Welcome to the Livepeer API reference docs. Here you will find all the
endpoints exposed on the standard Livepeer API, learn how to use them and what they return.

### Available Operations

* [PostStream](#poststream) - Create a stream
* [GetStream](#getstream) - Retrieve streams
* [GetStreamID](#getstreamid) - Retrieve a stream
* [PatchStreamID](#patchstreamid) - Update a stream
* [DeleteStreamID](#deletestreamid) - Delete a stream
* [DeleteStreamIDTerminate](#deletestreamidterminate) - Terminates a live stream
* [PostStreamIDStartPull](#poststreamidstartpull) - Start ingest for a pull stream
* [GetMultistreamTarget](#getmultistreamtarget) - Retrieve Multistream Targets
* [PostMultistreamTarget](#postmultistreamtarget) - Create a multistream target
* [GetMultistreamTargetID](#getmultistreamtargetid) - Retrieve a multistream target
* [PatchMultistreamTargetID](#patchmultistreamtargetid) - Update Multistream Target
* [DeleteMultistreamTargetID](#deletemultistreamtargetid) - Delete a multistream target
* [GetWebhook](#getwebhook) - Retrieve a Webhook
* [PostWebhook](#postwebhook) - Create a webhook
* [GetWebhookID](#getwebhookid) - Retrieve a webhook
* [PutWebhookID](#putwebhookid) - Update a webhook
* [DeleteWebhookID](#deletewebhookid) - Delete a webhook
* [GetWebhookIDLog](#getwebhookidlog) - Retrieve webhook logs
* [GetWebhookIDLogLogID](#getwebhookidloglogid) - Retrieve a webhook log
* [PostWebhookIDLogLogIDResend](#postwebhookidloglogidresend) - Resend a webhook
* [GetAsset](#getasset) - Retrieve assets
* [PostAssetRequestUpload](#postassetrequestupload) - Upload an asset
* [PostAssetUploadURL](#postassetuploadurl) - Upload asset via URL
* [GetAssetAssetID](#getassetassetid) - Retrieves an asset
* [PatchAssetAssetID](#patchassetassetid) - Patch an asset
* [DeleteAssetAssetID](#deleteassetassetid) - Delete an asset
* [PostClip](#postclip) - Create a clip
* [GetStreamIDClips](#getstreamidclips) - Retrieve clips of a livestream
* [PostStreamIDCreateMultistreamTarget](#poststreamidcreatemultistreamtarget) - Add a multistream target
* [DeleteStreamIDMultistreamTargetID](#deletestreamidmultistreamtargetid) - Remove a multistream target
* [GetSessionIDClips](#getsessionidclips) - Retrieve clips of a session
* [PostRoom](#postroom) - Create a room
* [GetRoomID](#getroomid) - Retrieve a room
* [DeleteRoomID](#deleteroomid) - Delete a room
* [PostRoomIDEgress](#postroomidegress) - Start room RTMP egress
* [DeleteRoomIDEgress](#deleteroomidegress) - Stop room RTMP egress
* [PostRoomIDUser](#postroomiduser) - Create a room user
* [GetRoomIDUserUserID](#getroomiduseruserid) - Get user details
* [PutRoomIDUserUserID](#putroomiduseruserid) - Update a room user
* [DeleteRoomIDUserUserID](#deleteroomiduseruserid) - Remove a user from the room
* [GetDataViewsQuery](#getdataviewsquery) - Query viewership metrics
* [GetDataViewsQueryCreator](#getdataviewsquerycreator) - Query creator viewership metrics
* [GetDataViewsQueryTotalPlaybackID](#getdataviewsquerytotalplaybackid) - Query public total views metrics
* [GetDataUsageQuery](#getdatausagequery) - Query usage metrics
* [GetSession](#getsession) - Retrieve sessions
* [GetSessionID](#getsessionid) - Retrieve a session
* [GetStreamParentIDSessions](#getstreamparentidsessions) - Retrieve Recorded Sessions
* [PostAccessControlSigningKey](#postaccesscontrolsigningkey) - Create a signing key
* [GetAccessControlSigningKey](#getaccesscontrolsigningkey) - Retrieves signing keys
* [DeleteAccessControlSigningKeyKeyID](#deleteaccesscontrolsigningkeykeyid) - Delete Signing Key
* [GetAccessControlSigningKeyKeyID](#getaccesscontrolsigningkeykeyid) - Retrieves a signing key
* [PatchAccessControlSigningKeyKeyID](#patchaccesscontrolsigningkeykeyid) - Update a signing key
* [GetTask](#gettask) - Retrieve Tasks
* [GetTaskTaskID](#gettasktaskid) - Retrieve a Task
* [PostTranscode](#posttranscode) - Transcode a video
* [GetPlaybackID](#getplaybackid) - Retrieve Playback Info

## PostStream

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
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostStream(ctx, components.NewStreamPayload{
        Name: "test_stream",
        Record: livepeer.Bool(false),
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

### Response

**[*operations.PostStreamResponse](../../models/operations/poststreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStream

Retrieve streams

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var streamsonly *string = livepeer.String("<value>")

    ctx := context.Background()
    res, err := s.GetStream(ctx, streamsonly)
    if err != nil {
        log.Fatal(err)
    }
    if res.Streams != nil {
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

**[*operations.GetStreamResponse](../../models/operations/getstreamresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStreamID

Retrieve a stream

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetStreamID(ctx, id)
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

**[*operations.GetStreamIDResponse](../../models/operations/getstreamidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchStreamID

Update a stream

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    streamPatchPayload := components.StreamPatchPayload{
        Record: livepeer.Bool(false),
    }

    ctx := context.Background()
    res, err := s.PatchStreamID(ctx, id, streamPatchPayload)
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

**[*operations.PatchStreamIDResponse](../../models/operations/patchstreamidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteStreamID

This will also suspend any active stream sessions, so make sure to wait
until the stream has finished. To explicitly interrupt an active
session, consider instead updating the suspended field in the stream
using the PATCH stream API.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteStreamID(ctx, id)
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

**[*operations.DeleteStreamIDResponse](../../models/operations/deletestreamidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteStreamIDTerminate

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
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteStreamIDTerminate(ctx, id)
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

**[*operations.DeleteStreamIDTerminateResponse](../../models/operations/deletestreamidterminateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostStreamIDStartPull

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
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.PostStreamIDStartPull(ctx, id)
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

**[*operations.PostStreamIDStartPullResponse](../../models/operations/poststreamidstartpullresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMultistreamTarget

Retrieve Multistream Targets

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetMultistreamTarget(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MultistreamTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.GetMultistreamTargetResponse](../../models/operations/getmultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostMultistreamTarget

Create a multistream target

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostMultistreamTarget(ctx, components.MultistreamTargetInput{
        URL: "rtmps://live.my-service.tv/channel/secretKey",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MultistreamTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.MultistreamTargetInput](../../models/components/multistreamtargetinput.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |

### Response

**[*operations.PostMultistreamTargetResponse](../../models/operations/postmultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMultistreamTargetID

Retrieve a multistream target

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetMultistreamTargetID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.MultistreamTarget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the multistream target                          |

### Response

**[*operations.GetMultistreamTargetIDResponse](../../models/operations/getmultistreamtargetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchMultistreamTargetID

Update Multistream Target

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    multistreamTargetPatchPayload := components.MultistreamTargetPatchPayload{
        URL: "rtmps://live.my-service.tv/channel/secretKey",
    }

    ctx := context.Background()
    res, err := s.PatchMultistreamTargetID(ctx, id, multistreamTargetPatchPayload)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | ID of the multistream target                                                                         |
| `multistreamTargetPatchPayload`                                                                      | [components.MultistreamTargetPatchPayload](../../models/components/multistreamtargetpatchpayload.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |

### Response

**[*operations.PatchMultistreamTargetIDResponse](../../models/operations/patchmultistreamtargetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteMultistreamTargetID

Make sure to remove any references to the target on existing
streams before actually deleting it from the API.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteMultistreamTargetID(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the multistream target                          |

### Response

**[*operations.DeleteMultistreamTargetIDResponse](../../models/operations/deletemultistreamtargetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetWebhook

Retrieve a Webhook

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetWebhook(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.GetWebhookResponse](../../models/operations/getwebhookresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostWebhook

To create a new webhook, you need to make an API call with the events you want to listen for and the URL that will be called when those events occur.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostWebhook(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.PostWebhookResponse](../../models/operations/postwebhookresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetWebhookID

Retrieve a webhook

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetWebhookID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.GetWebhookIDResponse](../../models/operations/getwebhookidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutWebhookID

Update a webhook

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.PutWebhookID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.PutWebhookIDResponse](../../models/operations/putwebhookidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteWebhookID

Delete a webhook

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteWebhookID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.DeleteWebhookIDResponse](../../models/operations/deletewebhookidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetWebhookIDLog

Retrieve webhook logs

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetWebhookIDLog(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookLogs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.GetWebhookIDLogResponse](../../models/operations/getwebhookidlogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetWebhookIDLogLogID

Retrieve a webhook log

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    var logID string = "<value>"

    ctx := context.Background()
    res, err := s.GetWebhookIDLogLogID(ctx, id, logID)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookLog != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `logID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.GetWebhookIDLogLogIDResponse](../../models/operations/getwebhookidloglogidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostWebhookIDLogLogIDResend

Use this API to resend the same webhook request. This is useful when
developing and debugging, allowing you to easily repeat the same webhook
to check or fix the behaviour in your handler.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    var logID string = "<value>"

    ctx := context.Background()
    res, err := s.PostWebhookIDLogLogIDResend(ctx, id, logID)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookLog != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `logID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.PostWebhookIDLogLogIDResendResponse](../../models/operations/postwebhookidloglogidresendresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAsset

Retrieve assets

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetAsset(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Assets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.GetAssetResponse](../../models/operations/getassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAssetRequestUpload

To upload an asset, your first need to request for a direct upload URL
and only then actually upload the contents of the asset.
\
\
Once you created a upload link, you have 2 options, resumable or direct
upload. For a more reliable experience, you should use resumable uploads
which will work better for users with unreliable or slow network
connections. If you want a simpler implementation though, you should
just use a direct upload.

## Direct Upload

For a direct upload, make a PUT request to the URL received in the url
field of the response above, with the raw video file as the request
body. response above:

## Resumable Upload

Livepeer supports resumable uploads via Tus. This section provides a
simple example of how to use tus-js-client to upload a video file.
\
\
From the previous section, we generated a URL to upload a video file to
Livepeer on POST /api/asset/request-upload. You should use the
tusEndpoint field of the response to upload the video file and track the
progress:

```
# This assumes there is an `input` element of `type="file"` with id
`fileInput` in the HTML


const input = document.getElementById('fileInput');

const file = input.files[0];

const upload = new tus.Upload(file, {
  endpoint: tusEndpoint, // URL from `tusEndpoint` field in the
`/request-upload` response
  metadata: {
    filename,
    filetype: 'video/mp4',
  },
  uploadSize: file.size,
  onError(err) {
    console.error('Error uploading file:', err);
  },
  onProgress(bytesUploaded, bytesTotal) {
    const percentage = ((bytesUploaded / bytesTotal) * 100).toFixed(2);
    console.log('Uploaded ' + percentage + '%');
  },
  onSuccess() {
    console.log('Upload finished:', upload.url);
  },
});

const previousUploads = await upload.findPreviousUploads();

if (previousUploads.length > 0) {
  upload.resumeFromPreviousUpload(previousUploads[0]);
}

upload.start();

```

> Note: If you are using tus from node.js, you need to add a custom URL
storage to enable resuming from previous uploads. On the browser, this
is enabled by default using local storage. In node.js, add urlStorage:
new tus.FileUrlStorage("path/to/tmp/file"), to the UploadFile object
definition above.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostAssetRequestUpload(ctx, components.NewAssetPayload{
        Name: "filename.mp4",
        StaticMp4: livepeer.Bool(true),
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.NewAssetPayload](../../models/components/newassetpayload.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |

### Response

**[*operations.PostAssetRequestUploadResponse](../../models/operations/postassetrequestuploadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAssetUploadURL

Upload asset via URL

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostAssetUploadURL(ctx, components.NewAssetPayload{
        Name: "filename.mp4",
        StaticMp4: livepeer.Bool(true),
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.NewAssetPayload](../../models/components/newassetpayload.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |

### Response

**[*operations.PostAssetUploadURLResponse](../../models/operations/postassetuploadurlresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAssetAssetID

Retrieves an asset

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var assetID string = "<value>"

    ctx := context.Background()
    res, err := s.GetAssetAssetID(ctx, assetID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Asset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `assetID`                                             | *string*                                              | :heavy_check_mark:                                    | ID of the asset                                       |

### Response

**[*operations.GetAssetAssetIDResponse](../../models/operations/getassetassetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchAssetAssetID

Patch an asset

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var assetID string = "<value>"

    assetPatchPayload := components.AssetPatchPayload{
        Name: livepeer.String("filename.mp4"),
    }

    ctx := context.Background()
    res, err := s.PatchAssetAssetID(ctx, assetID, assetPatchPayload)
    if err != nil {
        log.Fatal(err)
    }
    if res.Asset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `assetID`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | ID of the asset                                                              |
| `assetPatchPayload`                                                          | [components.AssetPatchPayload](../../models/components/assetpatchpayload.md) | :heavy_check_mark:                                                           | N/A                                                                          |

### Response

**[*operations.PatchAssetAssetIDResponse](../../models/operations/patchassetassetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAssetAssetID

Delete an asset

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var assetID string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteAssetAssetID(ctx, assetID)
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
| `assetID`                                             | *string*                                              | :heavy_check_mark:                                    | ID of the asset                                       |

### Response

**[*operations.DeleteAssetAssetIDResponse](../../models/operations/deleteassetassetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostClip

Create a clip

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostClip(ctx, components.ClipPayload{
        PlaybackID: "<value>",
        StartTime: 2290.64,
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStreamIDClips

Retrieve clips of a livestream

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetStreamIDClips(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Assets != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostStreamIDCreateMultistreamTarget

Add a multistream target

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    targetAddPayload := components.TargetAddPayload{
        Profile: "720p",
    }

    ctx := context.Background()
    res, err := s.PostStreamIDCreateMultistreamTarget(ctx, id, targetAddPayload)
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

**[*operations.PostStreamIDCreateMultistreamTargetResponse](../../models/operations/poststreamidcreatemultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteStreamIDMultistreamTargetID

Remove a multistream target

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    var targetID string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteStreamIDMultistreamTargetID(ctx, id, targetID)
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

**[*operations.DeleteStreamIDMultistreamTargetIDResponse](../../models/operations/deletestreamidmultistreamtargetidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSessionIDClips

Retrieve clips of a session

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetSessionIDClips(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Assets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the parent session                              |

### Response

**[*operations.GetSessionIDClipsResponse](../../models/operations/getsessionidclipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRoom

Create a multiparticipant livestreaming room.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostRoom(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateRoomResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.PostRoomResponse](../../models/operations/postroomresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRoomID

Retrieve a room

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetRoomID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Room != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.GetRoomIDResponse](../../models/operations/getroomidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteRoomID

Delete a room

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteRoomID(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.DeleteRoomIDResponse](../../models/operations/deleteroomidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRoomIDEgress

Create a livestream for your room.
This allows you to leverage livestreaming features like recording and HLS output.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    roomEgressPayload := components.RoomEgressPayload{
        StreamID: "aac12556-4d65-4d34-9fb6-d1f0985eb0a9",
    }

    ctx := context.Background()
    res, err := s.PostRoomIDEgress(ctx, id, roomEgressPayload)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `roomEgressPayload`                                                          | [components.RoomEgressPayload](../../models/components/roomegresspayload.md) | :heavy_check_mark:                                                           | N/A                                                                          |

### Response

**[*operations.PostRoomIDEgressResponse](../../models/operations/postroomidegressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteRoomIDEgress

Stop room RTMP egress

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteRoomIDEgress(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.DeleteRoomIDEgressResponse](../../models/operations/deleteroomidegressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRoomIDUser

Call this endpoint to add a user to a room, specifying a display name at a minimum.
The response will contain a joining URL for Livepeer's default meeting app.
Alternatively the joining token can be used with a custom app.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    roomUserPayload := components.RoomUserPayload{
        Name: "name",
        CanPublish: livepeer.Bool(true),
        CanPublishData: livepeer.Bool(true),
    }

    ctx := context.Background()
    res, err := s.PostRoomIDUser(ctx, id, roomUserPayload)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `id`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `roomUserPayload`                                                        | [components.RoomUserPayload](../../models/components/roomuserpayload.md) | :heavy_check_mark:                                                       | N/A                                                                      |

### Response

**[*operations.PostRoomIDUserResponse](../../models/operations/postroomiduserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRoomIDUserUserID

Get user details

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    var userID string = "<value>"

    ctx := context.Background()
    res, err := s.GetRoomIDUserUserID(ctx, id, userID)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRoomUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `userID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.GetRoomIDUserUserIDResponse](../../models/operations/getroomiduseruseridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutRoomIDUserUserID

Update properties for a user.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    var userID string = "<value>"

    roomUserUpdatePayload := components.RoomUserUpdatePayload{
        CanPublish: livepeer.Bool(true),
        CanPublishData: livepeer.Bool(true),
    }

    ctx := context.Background()
    res, err := s.PutRoomIDUserUserID(ctx, id, userID, roomUserUpdatePayload)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `userID`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `roomUserUpdatePayload`                                                              | [components.RoomUserUpdatePayload](../../models/components/roomuserupdatepayload.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |

### Response

**[*operations.PutRoomIDUserUserIDResponse](../../models/operations/putroomiduseruseridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteRoomIDUserUserID

Remove a user from the room

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    var userID string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteRoomIDUserUserID(ctx, id, userID)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `userID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |

### Response

**[*operations.DeleteRoomIDUserUserIDResponse](../../models/operations/deleteroomiduseruseridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDataViewsQuery

Requires a private (non-CORS) API key to be used.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "github.com/livepeer/livepeer-go/models/operations"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetDataViewsQuery(ctx, operations.GetDataViewsQueryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ViewershipMetrics != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetDataViewsQueryRequest](../../models/operations/getdataviewsqueryrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |

### Response

**[*operations.GetDataViewsQueryResponse](../../models/operations/getdataviewsqueryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDataViewsQueryCreator

Requires a proof of ownership to be sent in the request, which for now is just the assetId or streamId parameters (1 of those must be in the query-string).

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "github.com/livepeer/livepeer-go/models/operations"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetDataViewsQueryCreator(ctx, operations.GetDataViewsQueryCreatorRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ViewershipMetrics != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetDataViewsQueryCreatorRequest](../../models/operations/getdataviewsquerycreatorrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |

### Response

**[*operations.GetDataViewsQueryCreatorResponse](../../models/operations/getdataviewsquerycreatorresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDataViewsQueryTotalPlaybackID

Allows querying for the public metrics for viewership about a video.
This can be called from the frontend with a CORS key, or even
unauthenticated.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var playbackID string = "<value>"

    ctx := context.Background()
    res, err := s.GetDataViewsQueryTotalPlaybackID(ctx, playbackID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |
| `playbackID`                                                                                                                                          | *string*                                                                                                                                              | :heavy_check_mark:                                                                                                                                    | The playback ID to filter the query results. This can be a canonical<br/>playback ID from Livepeer assets or streams, or dStorage identifiers<br/>for assets<br/> |

### Response

**[*operations.GetDataViewsQueryTotalPlaybackIDResponse](../../models/operations/getdataviewsquerytotalplaybackidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDataUsageQuery

Query usage metrics

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "github.com/livepeer/livepeer-go/models/operations"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var from *int64 = livepeer.Int64(511615)

    var to *int64 = livepeer.Int64(267511)

    var timeStep *operations.GetDataUsageQueryQueryParamTimeStep = operations.GetDataUsageQueryQueryParamTimeStepDay.ToPointer()

    var creatorID *string = livepeer.String("<value>")

    ctx := context.Background()
    res, err := s.GetDataUsageQuery(ctx, from, to, timeStep, creatorID)
    if err != nil {
        log.Fatal(err)
    }
    if res.UsageMetric != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |
| `from`                                                                                                            | **int64*                                                                                                          | :heavy_minus_sign:                                                                                                | Start millis timestamp for the query range (inclusive)<br/>                                                       |
| `to`                                                                                                              | **int64*                                                                                                          | :heavy_minus_sign:                                                                                                | End millis timestamp for the query range (exclusive)<br/>                                                         |
| `timeStep`                                                                                                        | [*operations.GetDataUsageQueryQueryParamTimeStep](../../models/operations/getdatausagequeryqueryparamtimestep.md) | :heavy_minus_sign:                                                                                                | The time step to aggregate viewership metrics by<br/>                                                             |
| `creatorID`                                                                                                       | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | The creator ID to filter the query results<br/>                                                                   |

### Response

**[*operations.GetDataUsageQueryResponse](../../models/operations/getdatausagequeryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSession

Retrieve sessions

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetSession(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Sessions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.GetSessionResponse](../../models/operations/getsessionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSessionID

Retrieve a session

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetSessionID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Session != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the session                                     |

### Response

**[*operations.GetSessionIDResponse](../../models/operations/getsessionidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStreamParentIDSessions

Retrieve Recorded Sessions

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var parentID string = "<value>"

    var record *int64 = livepeer.Int64(1)

    ctx := context.Background()
    res, err := s.GetStreamParentIDSessions(ctx, parentID, record)
    if err != nil {
        log.Fatal(err)
    }
    if res.Sessions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `parentID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | ID of the parent stream                                                |                                                                        |
| `record`                                                               | **int64*                                                               | :heavy_minus_sign:                                                     | Flag indicating if the response should only include recorded<br/>sessions<br/> | 1                                                                      |

### Response

**[*operations.GetStreamParentIDSessionsResponse](../../models/operations/getstreamparentidsessionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAccessControlSigningKey

The publicKey is a representation of the public key, encoded as base 64 and is passed as a string, and  the privateKey is displayed only on creation. This is the only moment where the client can save the private key, otherwise it will be lost. Remember to decode your string when signing JWTs.
Up to 10 signing keys can be generated, after that you must delete at least one signing key to create a new one.

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostAccessControlSigningKey(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.PostAccessControlSigningKeyResponse](../../models/operations/postaccesscontrolsigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAccessControlSigningKey

Retrieves signing keys

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetAccessControlSigningKey(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.GetAccessControlSigningKeyResponse](../../models/operations/getaccesscontrolsigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAccessControlSigningKeyKeyID

Delete Signing Key

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var keyID string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteAccessControlSigningKeyKeyID(ctx, keyID)
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
| `keyID`                                               | *string*                                              | :heavy_check_mark:                                    | ID of the signing key                                 |

### Response

**[*operations.DeleteAccessControlSigningKeyKeyIDResponse](../../models/operations/deleteaccesscontrolsigningkeykeyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAccessControlSigningKeyKeyID

Retrieves a signing key

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var keyID string = "<value>"

    ctx := context.Background()
    res, err := s.GetAccessControlSigningKeyKeyID(ctx, keyID)
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `keyID`                                               | *string*                                              | :heavy_check_mark:                                    | ID of the signing key                                 |

### Response

**[*operations.GetAccessControlSigningKeyKeyIDResponse](../../models/operations/getaccesscontrolsigningkeykeyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchAccessControlSigningKeyKeyID

Update a signing key

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "github.com/livepeer/livepeer-go/models/operations"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var keyID string = "<value>"

    requestBody := operations.PatchAccessControlSigningKeyKeyIDRequestBody{}

    ctx := context.Background()
    res, err := s.PatchAccessControlSigningKeyKeyID(ctx, keyID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `keyID`                                                                                                                            | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | ID of the signing key                                                                                                              |
| `requestBody`                                                                                                                      | [operations.PatchAccessControlSigningKeyKeyIDRequestBody](../../models/operations/patchaccesscontrolsigningkeykeyidrequestbody.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |

### Response

**[*operations.PatchAccessControlSigningKeyKeyIDResponse](../../models/operations/patchaccesscontrolsigningkeykeyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTask

Retrieve Tasks

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetTask(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Tasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |

### Response

**[*operations.GetTaskResponse](../../models/operations/gettaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTaskTaskID

Retrieve a Task

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var taskID string = "<value>"

    ctx := context.Background()
    res, err := s.GetTaskTaskID(ctx, taskID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `taskID`                                              | *string*                                              | :heavy_check_mark:                                    | ID of the task                                        |

### Response

**[*operations.GetTaskTaskIDResponse](../../models/operations/gettasktaskidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostTranscode

`POST /transcode` transcodes a video file and uploads the results to the
specified storage service.
\
\
Transcoding is asynchronous so you will need to check the status of the
task in order to determine when transcoding is complete. The `id` field
in the response is the unique ID for the transcoding `Task`. The task
status can be queried using the [GET tasks
endpoint](https://docs.livepeer.org/reference/api/get-tasks):
\
\
When `status.phase` is `completed`,  transcoding will be complete and
the results will be stored in the storage service and the specified
output location.
\
\
The results will be available under `params.outputs.hls.path` and
`params.outputs.mp4.path` in the specified storage service.

## Input

\
This endpoint currently supports the following inputs:

* HTTP
* S3 API Compatible Service
\
\
**HTTP**
\
A public HTTP URL can be used to read a video file.

```json
{
    "url": "https://www.example.com/video.mp4"
}
```

| Name | Type   | Description                          |
| ---- | ------ | ------------------------------------ |
| url  | string | A public HTTP URL for the video file. |

Note: For IPFS HTTP gateway URLs, the API currently only supports “path
style” URLs and does not support “subdomain style” URLs. The API will
support both styles of URLs in a future update.
\
\
**S3 API Compatible Service**
\
\
S3 credentials can be used to authenticate with a S3 API compatible
service to read a video file.

```json
{
    "type": "s3",
    "endpoint": "https://gateway.storjshare.io",
    "credentials": {
        "accessKeyId": "$ACCESS_KEY_ID",
        "secretAccessKey": "$SECRET_ACCESS_KEY"
    },
    "bucket": "inbucket",
    "path": "/video/source.mp4"
}
```

## Storage

\
This endpoint currently supports the following storage services:

* S3 API Compatible Service
* Web3 Storage
\
\
**S3 API Compatible Service**

```json
{
  "type": "s3",
    "endpoint": "https://gateway.storjshare.io",
    "credentials": {
        "accessKeyId": "$ACCESS_KEY_ID",
        "secretAccessKey": "$SECRET_ACCESS_KEY"
    },
    "bucket": "mybucket"
}
```

**Web3 Storage**

```json
{
  "type": "web3.storage",
    "credentials": {
        "proof": "$UCAN_DELEGATION_PROOF",
    }
}
```

## Outputs

\
This endpoint currently supports the following output types:

* HLS
* MP4

**HLS**

```json
{
  "hls": {
        "path": "/samplevideo/hls"
    }
}
```

**MP4**

```json
{
  "mp4": {
        "path": "/samplevideo/mp4"
    }
}
```

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostTranscode(ctx, components.TranscodePayload{
        Input: components.CreateInputInput1(
                components.Input1{
                    URL: "https://s3.amazonaws.com/bucket/file.mp4",
                },
        ),
        Storage: components.CreateTranscodePayloadStorageStorage1(
                components.Storage1{
                    Type: components.StorageTypeS3,
                    Endpoint: "https://gateway.storjshare.io",
                    Bucket: "outputbucket",
                    Credentials: components.StorageCredentials{
                        AccessKeyID: "AKIAIOSFODNN7EXAMPLE",
                        SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
                    },
                },
        ),
        Outputs: components.Outputs{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.TranscodePayload](../../models/components/transcodepayload.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |

### Response

**[*operations.PostTranscodeResponse](../../models/operations/posttranscoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPlaybackID

Retrieve Playback Info

### Example Usage

```go
package main

import(
 "github.com/livepeer/livepeer-go/models/components"
 livepeer "github.com/livepeer/livepeer-go"
 "context"
 "log"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.GetPlaybackID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaybackInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The playback ID from the asset or livestream, e.g. `eaw4nk06ts2d0mzb`. |

### Response

**[*operations.GetPlaybackIDResponse](../../models/operations/getplaybackidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
