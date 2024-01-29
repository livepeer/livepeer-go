# Transcode
(*Transcode*)

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
    res, err := s.Transcode.Create(ctx, components.TranscodePayload{
        Input: components.CreateInputTranscodePayload1(
                components.TranscodePayload1{
                    URL: "https://s3.amazonaws.com/bucket/file.mp4",
                },
        ),
        Storage: components.CreateTranscodePayloadStorageTranscodePayloadSchemas1(
                components.TranscodePayloadSchemas1{
                    Type: components.TranscodePayloadSchemasTypeS3,
                    Endpoint: "https://gateway.storjshare.io",
                    Bucket: "outputbucket",
                    Credentials: components.TranscodePayloadCredentials{
                        AccessKeyID: "AKIAIOSFODNN7EXAMPLE",
                        SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
                    },
                },
        ),
        Outputs: components.Outputs{
            Hls: &components.Hls{
                Path: "/samplevideo/hls",
            },
            Mp4: &components.Mp4{
                Path: "/samplevideo/mp4",
            },
            Fmp4: &components.Fmp4{
                Path: "/samplevideo/fmp4",
            },
        },
        Profiles: []components.TranscodeProfile{
            components.TranscodeProfile{
                Name: livepeer.String("720p"),
                Bitrate: 638424,
            },
        },
        CreatorID: components.CreateInputCreatorIDStr(
        "string",
        ),
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

**[*operations.TranscodeResponse](../../models/operations/transcoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
