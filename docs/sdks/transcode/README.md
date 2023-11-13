# Transcode
(*Transcode*)

### Available Operations

* [Transcode](#transcode) - Transcode a video

## Transcode

Transcode a video

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
    res, err := s.Transcode.Transcode(ctx, components.TaskInput{
        InputAssetID: livepeer.String("09F8B46C-61A0-4254-9875-F71F4C605BC7"),
        OutputAssetID: livepeer.String("09F8B46C-61A0-4254-9875-F71F4C605BC7"),
        Params: &components.Params{
            Upload: &components.TaskUpload{
                URL: livepeer.String("https://cdn.livepeer.com/ABC123/filename.mp4"),
                Encryption: &components.Encryption{
                    EncryptedKey: "string",
                },
                RecordedSessionID: livepeer.String("78df0075-b5f3-4683-a618-1086faca35dc"),
            },
            Import: &components.Upload{
                URL: livepeer.String("https://cdn.livepeer.com/ABC123/filename.mp4"),
                Encryption: &components.Encryption{
                    EncryptedKey: "string",
                },
                RecordedSessionID: livepeer.String("78df0075-b5f3-4683-a618-1086faca35dc"),
            },
            Export: components.CreateExportTaskParams1ExportTaskParamsSchemas1(
                    components.ExportTaskParamsSchemas1{
                        Custom: components.Custom{
                            URL: "https://s3.amazonaws.com/my-bucket/path/filename.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LLMMB",
                            Headers: map[string]string{
                                "key": "string",
                            },
                        },
                    },
            ),
            ExportData: &components.TaskExportData{
                Content: components.Content{},
                Ipfs: &components.IpfsExportParams1{
                    NftMetadata: &components.NftMetadata{},
                    Pinata: components.CreatePinataIpfsExportParamsSchemas1(
                            components.IpfsExportParamsSchemas1{
                                Jwt: "string",
                            },
                    ),
                },
            },
            Transcode: &components.Transcode{
                Profile: &components.FfmpegProfile{
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
            TranscodeFile: &components.TranscodeFile{
                Input: &components.Input{
                    URL: livepeer.String("https://cdn.livepeer.com/ABC123/filename.mp4"),
                },
                Storage: &components.TaskStorage{
                    URL: livepeer.String("s3+https://accessKeyId:secretAccessKey@s3Endpoint/bucket"),
                },
                Outputs: &components.Outputs{
                    Hls: &components.Hls{
                        Path: livepeer.String("/samplevideo/hls"),
                    },
                    Mp4: &components.Mp4{
                        Path: livepeer.String("/samplevideo/mp4"),
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
                CreatorID: components.CreateInputCreatorIDStr(
                "string",
                ),
            },
        },
        Clip: &components.Clip{
            ClipStrategy: &components.ClipStrategy{},
        },
        Output: &components.TaskOutput{
            Upload: &components.TaskUploadInput{
                AssetSpec: &components.AssetInput{
                    Type: components.AssetTypeVideo.ToPointer(),
                    PlaybackID: livepeer.String("eaw4nk06ts2d0mzb"),
                    PlaybackPolicy: &components.PlaybackPolicy{
                        Type: components.TypeJwt,
                        WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
                        WebhookContext: map[string]interface{}{
                            "foo": "string",
                        },
                    },
                    Source: components.CreateSourceAsset1(
                            components.Asset1{
                                Type: components.AssetSchemasTypeURL,
                                URL: "http://thorough-supermarket.net",
                                Encryption: &components.Encryption{
                                    EncryptedKey: "string",
                                },
                            },
                    ),
                    CreatorID: components.CreateCreatorIDCreatorID1(
                            components.CreatorID1{
                                Type: components.CreatorIDTypeUnverified,
                                Value: "string",
                            },
                    ),
                    Storage: &components.AssetStorageInput{
                        Ipfs: &components.AssetIpfsInput{
                            Spec: &components.AssetSpec{
                                NftMetadata: &components.AssetNftMetadata{},
                            },
                            Cid: livepeer.String("bafybeihoqtemwitqajy6d654tmghqqvxmzgblddj2egst6yilplr5num6u"),
                            NftMetadata: &components.IpfsFileInfoInput{
                                Cid: "bafybeihoqtemwitqajy6d654tmghqqvxmzgblddj2egst6yilplr5num6u",
                            },
                        },
                    },
                    Name: "filename.mp4",
                    Hash: []components.Hash{
                        components.Hash{
                            Hash: livepeer.String("9b560b28b85378a5004117539196ab24e21bbd75b0e9eb1a8bc7c5fd80dc5b57"),
                            Algorithm: livepeer.String("sha256"),
                        },
                    },
                },
                AdditionalProperties: map[string]interface{}{
                    "key": "string",
                },
            },
            Import: &components.UploadInput{
                AssetSpec: &components.AssetInput{
                    Type: components.AssetTypeVideo.ToPointer(),
                    PlaybackID: livepeer.String("eaw4nk06ts2d0mzb"),
                    PlaybackPolicy: &components.PlaybackPolicy{
                        Type: components.TypeJwt,
                        WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
                        WebhookContext: map[string]interface{}{
                            "foo": "string",
                        },
                    },
                    Source: components.CreateSourceAsset1(
                            components.Asset1{
                                Type: components.AssetSchemasTypeURL,
                                URL: "http://doting-decongestant.biz",
                                Encryption: &components.Encryption{
                                    EncryptedKey: "string",
                                },
                            },
                    ),
                    CreatorID: components.CreateCreatorIDCreatorID1(
                            components.CreatorID1{
                                Type: components.CreatorIDTypeUnverified,
                                Value: "string",
                            },
                    ),
                    Storage: &components.AssetStorageInput{
                        Ipfs: &components.AssetIpfsInput{
                            Spec: &components.AssetSpec{
                                NftMetadata: &components.AssetNftMetadata{},
                            },
                            Cid: livepeer.String("bafybeihoqtemwitqajy6d654tmghqqvxmzgblddj2egst6yilplr5num6u"),
                            NftMetadata: &components.IpfsFileInfoInput{
                                Cid: "bafybeihoqtemwitqajy6d654tmghqqvxmzgblddj2egst6yilplr5num6u",
                            },
                        },
                    },
                    Name: "filename.mp4",
                    Hash: []components.Hash{
                        components.Hash{
                            Hash: livepeer.String("9b560b28b85378a5004117539196ab24e21bbd75b0e9eb1a8bc7c5fd80dc5b57"),
                            Algorithm: livepeer.String("sha256"),
                        },
                    },
                },
                AdditionalProperties: map[string]interface{}{
                    "key": "string",
                },
            },
            Export: &components.TaskExport{
                Ipfs: &components.TaskIpfsInput{
                    VideoFileCid: "string",
                },
            },
            ExportData: &components.TaskSchemasExportData{
                Ipfs: &components.TaskSchemasIpfs{
                    Cid: "string",
                },
            },
            Transcode: &components.TaskTranscodeInput{
                Asset: &components.TaskAssetInput{
                    AssetSpec: &components.AssetInput{
                        Type: components.AssetTypeVideo.ToPointer(),
                        PlaybackID: livepeer.String("eaw4nk06ts2d0mzb"),
                        PlaybackPolicy: &components.PlaybackPolicy{
                            Type: components.TypeJwt,
                            WebhookID: livepeer.String("3e02c844-d364-4d48-b401-24b2773b5d6c"),
                            WebhookContext: map[string]interface{}{
                                "foo": "string",
                            },
                        },
                        Source: components.CreateSourceAsset1(
                                components.Asset1{
                                    Type: components.AssetSchemasTypeURL,
                                    URL: "https://spotted-emergent.name",
                                    Encryption: &components.Encryption{
                                        EncryptedKey: "string",
                                    },
                                },
                        ),
                        CreatorID: components.CreateCreatorIDCreatorID1(
                                components.CreatorID1{
                                    Type: components.CreatorIDTypeUnverified,
                                    Value: "string",
                                },
                        ),
                        Storage: &components.AssetStorageInput{
                            Ipfs: &components.AssetIpfsInput{
                                Spec: &components.AssetSpec{
                                    NftMetadata: &components.AssetNftMetadata{},
                                },
                                Cid: livepeer.String("bafybeihoqtemwitqajy6d654tmghqqvxmzgblddj2egst6yilplr5num6u"),
                                NftMetadata: &components.IpfsFileInfoInput{
                                    Cid: "bafybeihoqtemwitqajy6d654tmghqqvxmzgblddj2egst6yilplr5num6u",
                                },
                            },
                        },
                        Name: "filename.mp4",
                        Hash: []components.Hash{
                            components.Hash{
                                Hash: livepeer.String("9b560b28b85378a5004117539196ab24e21bbd75b0e9eb1a8bc7c5fd80dc5b57"),
                                Algorithm: livepeer.String("sha256"),
                            },
                        },
                    },
                    AdditionalProperties: map[string]interface{}{
                        "key": "string",
                    },
                },
            },
        },
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [components.TaskInput](../../models/components/taskinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.TranscodeResponse](../../models/operations/transcoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
