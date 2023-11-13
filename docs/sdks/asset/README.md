# Asset
(*Asset*)

### Available Operations

* [GetAssets](#getassets) - Retrieve assets
* [RequestUpload](#requestupload) - Upload an asset
* [UploadAssetViaURL](#uploadassetviaurl) - Upload asset via URL
* [DeleteAsset](#deleteasset) - Delete an asset
* [GetAsset](#getasset) - Retrieves an asset
* [PatchAssetAssetID](#patchassetassetid) - Update an asset

## GetAssets

Retrieve assets

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
    res, err := s.Asset.GetAssets(ctx)
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


### Response

**[*operations.GetAssetsResponse](../../models/operations/getassetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RequestUpload

Upload an asset

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
    res, err := s.Asset.RequestUpload(ctx, components.NewAssetPayload{
        Name: "filename.mp4",
        StaticMp4: livepeer.Bool(true),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
            WebhookContext: map[string]interface{}{
                "foo": "string",
            },
        },
        CreatorID: components.CreateInputCreatorIDCreatorID(
                components.CreateCreatorIDCreatorID1(
                            components.CreatorID1{
                                Type: components.CreatorIDTypeUnverified,
                                Value: "string",
                            },
                ),
        ),
        Storage: &components.NewAssetPayloadStorage{
            Ipfs: components.CreateNewAssetPayloadIpfsBoolean(
            false,
            ),
        },
        URL: livepeer.String("https://s3.amazonaws.com/my-bucket/path/filename.mp4"),
        Encryption: &components.NewAssetPayloadEncryption{
            EncryptedKey: "string",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.NewAssetPayload](../../models/components/newassetpayload.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.RequestUploadResponse](../../models/operations/requestuploadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UploadAssetViaURL

Upload asset via URL

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
    res, err := s.Asset.UploadAssetViaURL(ctx, components.NewAssetPayload{
        Name: "filename.mp4",
        StaticMp4: livepeer.Bool(true),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeJwt,
            WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
            WebhookContext: map[string]interface{}{
                "foo": "string",
            },
        },
        CreatorID: components.CreateInputCreatorIDStr(
        "string",
        ),
        Storage: &components.NewAssetPayloadStorage{
            Ipfs: components.CreateNewAssetPayloadIpfsNewAssetPayload1(
                    components.NewAssetPayload1{
                        Spec: &components.Spec{
                            NftMetadata: &components.SpecNftMetadata{},
                        },
                    },
            ),
        },
        URL: livepeer.String("https://s3.amazonaws.com/my-bucket/path/filename.mp4"),
        Encryption: &components.NewAssetPayloadEncryption{
            EncryptedKey: "string",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.NewAssetPayload](../../models/components/newassetpayload.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.UploadAssetViaURLResponse](../../models/operations/uploadassetviaurlresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteAsset

Delete an asset

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


    var assetID string = "string"

    ctx := context.Background()
    res, err := s.Asset.DeleteAsset(ctx, assetID)
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
| `assetID`                                             | *string*                                              | :heavy_check_mark:                                    | ID of the asset                                       |


### Response

**[*operations.DeleteAssetResponse](../../models/operations/deleteassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAsset

Retrieves an asset

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


    var assetID string = "string"

    ctx := context.Background()
    res, err := s.Asset.GetAsset(ctx, assetID)
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

**[*operations.GetAssetResponse](../../models/operations/getassetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PatchAssetAssetID

Update an asset

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


    var assetID string = "string"

    assetPatchPayload := components.AssetPatchPayload{
        Name: livepeer.String("filename.mp4"),
        CreatorID: components.CreateInputCreatorIDStr(
        "string",
        ),
        PlaybackPolicy: &components.PlaybackPolicy{
            Type: components.TypeWebhook,
            WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
            WebhookContext: map[string]interface{}{
                "foo": "string",
            },
        },
        Storage: &components.Storage{
            Ipfs: components.CreateIpfsOne(
                    components.One{
                        Spec: &components.Spec{
                            NftMetadata: &components.SpecNftMetadata{},
                        },
                    },
            ),
        },
    }

    ctx := context.Background()
    res, err := s.Asset.PatchAssetAssetID(ctx, assetID, assetPatchPayload)
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
| sdkerrors.SDKError | 400-600            | */*                |
