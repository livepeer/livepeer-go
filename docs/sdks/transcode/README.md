# Transcode
(*Transcode*)

## Overview

Operations related to transcode api

### Available Operations

* [Create](#create) - Transcode a video

## Create

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
- HTTP
- S3 API Compatible Service
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
- S3 API Compatible Service
- Web3 Storage
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
- HLS
- MP4

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
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.TranscodePayload{
        Input: components.CreateInputInput1(
                components.Input1{
                    URL: "https://s3.amazonaws.com/bucket/file.mp4",
                },
        ),
        Outputs: components.Outputs{
            Fmp4: &components.Fmp4{
                Path: "/samplevideo/fmp4",
            },
            Hls: &components.Hls{
                Path: "/samplevideo/hls",
            },
            Mp4: &components.Mp4{
                Path: "/samplevideo/mp4",
            },
        },
        Profiles: []components.TranscodeProfile{
            components.TranscodeProfile{
                Bitrate: 3000000,
                Encoder: components.EncoderH264.ToPointer(),
                Fps: livepeergo.Int64(30),
                FpsDen: livepeergo.Int64(1),
                Gop: livepeergo.String("2"),
                Name: livepeergo.String("720p"),
                Profile: components.ProfileH264Baseline.ToPointer(),
                Quality: livepeergo.Int64(23),
                Width: livepeergo.Int64(1280),
            },
        },
        Storage: components.CreateTranscodePayloadStorageStorage1(
                components.Storage1{
                    Bucket: "outputbucket",
                    Credentials: components.StorageCredentials{
                        AccessKeyID: "AKIAIOSFODNN7EXAMPLE",
                        SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
                    },
                    Endpoint: "https://gateway.storjshare.io",
                    Type: components.StorageTypeS3,
                },
        ),
    }
    ctx := context.Background()
    res, err := s.Transcode.Create(ctx, request)
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

**[*operations.TranscodeVideoResponse](../../models/operations/transcodevideoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
