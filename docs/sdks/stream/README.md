# Stream
(*Stream*)

### Available Operations

* [GetStreams](#getstreams) - Retrieve streams
* [CreateStream](#createstream) - Create a stream
* [DeleteStream](#deletestream) - Delete a stream
* [GetStream](#getstream) - Retrieve a stream
* [UpdateStream](#updatestream) - Update a stream

## GetStreams

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
    res, err := s.Stream.GetStreams(ctx, streamsonly)
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

## CreateStream

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
    res, err := s.Stream.CreateStream(ctx, components.NewStreamPayload{
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
            Type: components.TypePublic,
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

## DeleteStream

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
    res, err := s.Stream.DeleteStream(ctx, id)
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

## GetStream

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
    res, err := s.Stream.GetStream(ctx, id)
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

## UpdateStream

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
        CreatorID: components.CreateInputCreatorIDCreatorID(
                components.CreateCreatorIDCreatorID1(
                            components.CreatorID1{
                                Type: components.CreatorIDTypeUnverified,
                                Value: "string",
                            },
                ),
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
    res, err := s.Stream.UpdateStream(ctx, id, streamPatchPayload)
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
