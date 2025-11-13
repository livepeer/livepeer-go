# Generate
(*Generate*)

## Overview

Operations related to AI generate api

### Available Operations

* [TextToImage](#texttoimage) - Text To Image
* [ImageToImage](#imagetoimage) - Image To Image
* [ImageToVideo](#imagetovideo) - Image To Video
* [Upscale](#upscale) - Upscale
* [AudioToText](#audiototext) - Audio To Text
* [SegmentAnything2](#segmentanything2) - Segment Anything 2
* [Llm](#llm) - LLM
* [ImageToText](#imagetotext) - Image To Text
* [LiveVideoToVideo](#livevideotovideo) - Live Video To Video
* [TextToSpeech](#texttospeech) - Text To Speech

## TextToImage

Generate images from text prompts.

### Example Usage

<!-- UsageSnippet language="go" operationID="genTextToImage" method="post" path="/api/generate/text-to-image" -->
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

    res, err := s.Generate.TextToImage(ctx, components.TextToImageParams{
        Prompt: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.TextToImageParams](../../models/components/texttoimageparams.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GenTextToImageResponse](../../models/operations/gentexttoimageresponse.md), error**

### Errors

| Error Type                                              | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| sdkerrors.GenTextToImageResponseBody                    | 400                                                     | application/json                                        |
| sdkerrors.GenTextToImageGenerateResponseBody            | 401                                                     | application/json                                        |
| sdkerrors.GenTextToImageGenerateResponseResponseBody    | 422                                                     | application/json                                        |
| sdkerrors.GenTextToImageGenerateResponse500ResponseBody | 500                                                     | application/json                                        |
| sdkerrors.SDKError                                      | 4XX, 5XX                                                | \*/\*                                                   |

## ImageToImage

Apply image transformations to a provided image.

### Example Usage

<!-- UsageSnippet language="go" operationID="genImageToImage" method="post" path="/api/generate/image-to-image" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"os"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Generate.ImageToImage(ctx, components.BodyGenImageToImage{
        Prompt: "<value>",
        Image: components.Image{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.BodyGenImageToImage](../../models/components/bodygenimagetoimage.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GenImageToImageResponse](../../models/operations/genimagetoimageresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.GenImageToImageResponseBody                    | 400                                                      | application/json                                         |
| sdkerrors.GenImageToImageGenerateResponseBody            | 401                                                      | application/json                                         |
| sdkerrors.GenImageToImageGenerateResponseResponseBody    | 422                                                      | application/json                                         |
| sdkerrors.GenImageToImageGenerateResponse500ResponseBody | 500                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## ImageToVideo

Generate a video from a provided image.

### Example Usage

<!-- UsageSnippet language="go" operationID="genImageToVideo" method="post" path="/api/generate/image-to-video" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"os"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Generate.ImageToVideo(ctx, components.BodyGenImageToVideo{
        Image: components.BodyGenImageToVideoImage{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VideoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.BodyGenImageToVideo](../../models/components/bodygenimagetovideo.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GenImageToVideoResponse](../../models/operations/genimagetovideoresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.GenImageToVideoResponseBody                    | 400                                                      | application/json                                         |
| sdkerrors.GenImageToVideoGenerateResponseBody            | 401                                                      | application/json                                         |
| sdkerrors.GenImageToVideoGenerateResponseResponseBody    | 422                                                      | application/json                                         |
| sdkerrors.GenImageToVideoGenerateResponse500ResponseBody | 500                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## Upscale

Upscale an image by increasing its resolution.

### Example Usage

<!-- UsageSnippet language="go" operationID="genUpscale" method="post" path="/api/generate/upscale" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"os"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Generate.Upscale(ctx, components.BodyGenUpscale{
        Prompt: "<value>",
        Image: components.BodyGenUpscaleImage{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.BodyGenUpscale](../../models/components/bodygenupscale.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.GenUpscaleResponse](../../models/operations/genupscaleresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.GenUpscaleResponseBody                    | 400                                                 | application/json                                    |
| sdkerrors.GenUpscaleGenerateResponseBody            | 401                                                 | application/json                                    |
| sdkerrors.GenUpscaleGenerateResponseResponseBody    | 422                                                 | application/json                                    |
| sdkerrors.GenUpscaleGenerateResponse500ResponseBody | 500                                                 | application/json                                    |
| sdkerrors.SDKError                                  | 4XX, 5XX                                            | \*/\*                                               |

## AudioToText

Transcribe audio files to text.

### Example Usage

<!-- UsageSnippet language="go" operationID="genAudioToText" method="post" path="/api/generate/audio-to-text" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"os"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Generate.AudioToText(ctx, components.BodyGenAudioToText{
        Audio: components.Audio{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TextResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BodyGenAudioToText](../../models/components/bodygenaudiototext.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GenAudioToTextResponse](../../models/operations/genaudiototextresponse.md), error**

### Errors

| Error Type                                              | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| sdkerrors.GenAudioToTextResponseBody                    | 400                                                     | application/json                                        |
| sdkerrors.GenAudioToTextGenerateResponseBody            | 401                                                     | application/json                                        |
| sdkerrors.GenAudioToTextGenerateResponseResponseBody    | 413                                                     | application/json                                        |
| sdkerrors.GenAudioToTextGenerateResponse415ResponseBody | 415                                                     | application/json                                        |
| sdkerrors.GenAudioToTextGenerateResponse422ResponseBody | 422                                                     | application/json                                        |
| sdkerrors.GenAudioToTextGenerateResponse500ResponseBody | 500                                                     | application/json                                        |
| sdkerrors.SDKError                                      | 4XX, 5XX                                                | \*/\*                                                   |

## SegmentAnything2

Segment objects in an image.

### Example Usage

<!-- UsageSnippet language="go" operationID="genSegmentAnything2" method="post" path="/api/generate/segment-anything-2" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"os"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Generate.SegmentAnything2(ctx, components.BodyGenSegmentAnything2{
        Image: components.BodyGenSegmentAnything2Image{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MasksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.BodyGenSegmentAnything2](../../models/components/bodygensegmentanything2.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GenSegmentAnything2Response](../../models/operations/gensegmentanything2response.md), error**

### Errors

| Error Type                                                   | Status Code                                                  | Content Type                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| sdkerrors.GenSegmentAnything2ResponseBody                    | 400                                                          | application/json                                             |
| sdkerrors.GenSegmentAnything2GenerateResponseBody            | 401                                                          | application/json                                             |
| sdkerrors.GenSegmentAnything2GenerateResponseResponseBody    | 422                                                          | application/json                                             |
| sdkerrors.GenSegmentAnything2GenerateResponse500ResponseBody | 500                                                          | application/json                                             |
| sdkerrors.SDKError                                           | 4XX, 5XX                                                     | \*/\*                                                        |

## Llm

Generate text using a language model.

### Example Usage

<!-- UsageSnippet language="go" operationID="genLLM" method="post" path="/api/generate/llm" -->
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

    res, err := s.Generate.Llm(ctx, components.LLMRequest{
        Messages: []components.LLMMessage{
            components.LLMMessage{
                Role: "<value>",
                Content: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LLMResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.LLMRequest](../../models/components/llmrequest.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.GenLLMResponse](../../models/operations/genllmresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.GenLLMResponseBody                    | 400                                             | application/json                                |
| sdkerrors.GenLLMGenerateResponseBody            | 401                                             | application/json                                |
| sdkerrors.GenLLMGenerateResponseResponseBody    | 422                                             | application/json                                |
| sdkerrors.GenLLMGenerateResponse500ResponseBody | 500                                             | application/json                                |
| sdkerrors.SDKError                              | 4XX, 5XX                                        | \*/\*                                           |

## ImageToText

Transform image files to text.

### Example Usage

<!-- UsageSnippet language="go" operationID="genImageToText" method="post" path="/api/generate/image-to-text" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"os"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Generate.ImageToText(ctx, components.BodyGenImageToText{
        Image: components.BodyGenImageToTextImage{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageToTextResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BodyGenImageToText](../../models/components/bodygenimagetotext.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GenImageToTextResponse](../../models/operations/genimagetotextresponse.md), error**

### Errors

| Error Type                                              | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| sdkerrors.GenImageToTextResponseBody                    | 400                                                     | application/json                                        |
| sdkerrors.GenImageToTextGenerateResponseBody            | 401                                                     | application/json                                        |
| sdkerrors.GenImageToTextGenerateResponseResponseBody    | 413                                                     | application/json                                        |
| sdkerrors.GenImageToTextGenerateResponse422ResponseBody | 422                                                     | application/json                                        |
| sdkerrors.GenImageToTextGenerateResponse500ResponseBody | 500                                                     | application/json                                        |
| sdkerrors.SDKError                                      | 4XX, 5XX                                                | \*/\*                                                   |

## LiveVideoToVideo

Apply transformations to a live video streamed to the returned endpoints.

### Example Usage

<!-- UsageSnippet language="go" operationID="genLiveVideoToVideo" method="post" path="/api/generate/live-video-to-video" -->
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

    res, err := s.Generate.LiveVideoToVideo(ctx, components.LiveVideoToVideoParams{
        SubscribeURL: "https://soulful-finding.biz",
        PublishURL: "https://monumental-representation.biz/",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LiveVideoToVideoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.LiveVideoToVideoParams](../../models/components/livevideotovideoparams.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GenLiveVideoToVideoResponse](../../models/operations/genlivevideotovideoresponse.md), error**

### Errors

| Error Type                                                   | Status Code                                                  | Content Type                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| sdkerrors.GenLiveVideoToVideoResponseBody                    | 400                                                          | application/json                                             |
| sdkerrors.GenLiveVideoToVideoGenerateResponseBody            | 401                                                          | application/json                                             |
| sdkerrors.GenLiveVideoToVideoGenerateResponseResponseBody    | 422                                                          | application/json                                             |
| sdkerrors.GenLiveVideoToVideoGenerateResponse500ResponseBody | 500                                                          | application/json                                             |
| sdkerrors.SDKError                                           | 4XX, 5XX                                                     | \*/\*                                                        |

## TextToSpeech

Generate a text-to-speech audio file based on the provided text input and speaker description.

### Example Usage

<!-- UsageSnippet language="go" operationID="genTextToSpeech" method="post" path="/api/generate/text-to-speech" -->
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

    res, err := s.Generate.TextToSpeech(ctx, components.TextToSpeechParams{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AudioResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.TextToSpeechParams](../../models/components/texttospeechparams.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GenTextToSpeechResponse](../../models/operations/gentexttospeechresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.GenTextToSpeechResponseBody                    | 400                                                      | application/json                                         |
| sdkerrors.GenTextToSpeechGenerateResponseBody            | 401                                                      | application/json                                         |
| sdkerrors.GenTextToSpeechGenerateResponseResponseBody    | 422                                                      | application/json                                         |
| sdkerrors.GenTextToSpeechGenerateResponse500ResponseBody | 500                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 4XX, 5XX                                                 | \*/\*                                                    |