# Stream
(*Stream*)

### Available Operations

* [GetAll](#getall) - Retrieve streams
* [Create](#create) - Create a stream
* [Delete](#delete) - Delete a stream
* [Get](#get) - Retrieve a stream
* [Update](#update) - Update a stream
* [CreateClip](#createclip) - Create a clip
* [GetAllClips](#getallclips) - Retrieve clips of a livestream

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

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `streamsonly`                                                                                      | **string*                                                                                          | :heavy_minus_sign:                                                                                 | Filter the API response and retrieve a specific subset of stream objects based on certain criteria |


### Response

**[*operations.GetStreamsResponse](../../models/operations/getstreamsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Create

Create a stream

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
            WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
            WebhookContext: map[string]interface{}{
                "foo": "string",
            },
        },
        Profiles: []components.FfmpegProfile{
            components.FfmpegProfile{
                Width: 1280,
                Name: "720p",
                Height: 720,
                Bitrate: 4000,
                Fps: 30,
                FpsDen: livepeer.Int64(1),
                Gop: livepeer.String("60"),
                Profile: components.ProfileH264High.ToPointer(),
                Encoder: components.EncoderH264.ToPointer(),
            },
        },
        Record: livepeer.Bool(false),
        Multistream: &components.Multistream{
            Targets: []components.Targets{
                components.Targets{
                    Profile: "720p",
                    Spec: &components.MultistreamSpec{
                        URL: "rtmps://live.my-service.tv/channel/secretKey",
                    },
                },
            },
        },
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

Delete a stream

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
            Targets: []components.Targets{
                components.Targets{
                    Profile: "720p",
                    Spec: &components.MultistreamSpec{
                        URL: "rtmps://live.my-service.tv/channel/secretKey",
                    },
                },
            },
        },
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypePublic,
            WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
            WebhookContext: map[string]interface{}{
                "foo": "string",
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

## CreateClip

Create a clip from a livestream


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

**[*operations.GetStreamIDClipsResponse](../../models/operations/getstreamidclipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
