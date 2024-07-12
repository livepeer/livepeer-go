# Asset
(*Asset*)

## Overview

Operations related to asset/vod api

### Available Operations

* [GetAll](#getall) - Retrieve assets
* [Create](#create) - Upload an asset
* [CreateViaURL](#createviaurl) - Upload asset via URL
* [Get](#get) - Retrieves an asset
* [Update](#update) - Patch an asset
* [Delete](#delete) - Delete an asset

## GetAll

Retrieve assets

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

    ctx := context.Background()
    res, err := s.Asset.GetAll(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetAssetsResponse](../../models/operations/getassetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

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
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.NewAssetPayload{
        Name: "filename.mp4",
        StaticMp4: livepeergo.Bool(true),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Float64(600),
        },
        Profiles: []components.TranscodeProfile{
            components.TranscodeProfile{
                Width: livepeergo.Int64(1280),
                Name: livepeergo.String("720p"),
                Bitrate: 3000000,
                Quality: livepeergo.Int64(23),
                Fps: livepeergo.Int64(30),
                FpsDen: livepeergo.Int64(1),
                Gop: livepeergo.String("2"),
                Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
                Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Asset.Create(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.NewAssetPayload](../../models/components/newassetpayload.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |


### Response

**[*operations.RequestUploadResponse](../../models/operations/requestuploadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateViaURL

Upload asset via URL

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
    request := components.NewAssetFromURLPayload{
        Name: "filename.mp4",
        StaticMp4: livepeergo.Bool(true),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Float64(600),
        },
        URL: "https://s3.amazonaws.com/my-bucket/path/filename.mp4",
        Profiles: []components.TranscodeProfile{
            components.TranscodeProfile{
                Width: livepeergo.Int64(1280),
                Name: livepeergo.String("720p"),
                Bitrate: 3000000,
                Quality: livepeergo.Int64(23),
                Fps: livepeergo.Int64(30),
                FpsDen: livepeergo.Int64(1),
                Gop: livepeergo.String("2"),
                Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
                Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Asset.CreateViaURL(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.NewAssetFromURLPayload](../../models/components/newassetfromurlpayload.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.UploadAssetResponse](../../models/operations/uploadassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Retrieves an asset

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
    var assetID string = "<value>"
    ctx := context.Background()
    res, err := s.Asset.Get(ctx, assetID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Asset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | ID of the asset                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetAssetResponse](../../models/operations/getassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Patch an asset

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
    var assetID string = "<value>"

    assetPatchPayload := components.AssetPatchPayload{
        Name: livepeergo.String("filename.mp4"),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
            WebhookContext: map[string]any{
                "streamerId": "my-custom-id",
            },
            RefreshInterval: livepeergo.Float64(600),
        },
    }
    ctx := context.Background()
    res, err := s.Asset.Update(ctx, assetID, assetPatchPayload)
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
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.UpdateAssetResponse](../../models/operations/updateassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

Delete an asset

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
    var assetID string = "<value>"
    ctx := context.Background()
    res, err := s.Asset.Delete(ctx, assetID)
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
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | ID of the asset                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.DeleteAssetResponse](../../models/operations/deleteassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
