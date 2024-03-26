# Livepeer Go SDK

The Livepeer Go library provides convenient access to the Livepeer Studio API from applications written in Golang.

## Documentation

For full documentation and examples, please visit [docs.livepeer.org](https://docs.livepeer.org/sdks/go/).

## Installation

Install the package:

```bash
go get github.com/livepeer/livepeer-go
```

## Usage

The library needs to be configured with your Livepeer Studio account's API key, which is available in the [Studio Dashboard](https://livepeer.studio).

```go
package main

import(
 "context"
 "log"
 livepeer "github.com/livepeer/livepeer-go"
 "github.com/livepeer/livepeer-go/models/components"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Stream.GetAll(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## Resources and Operations

### [Livepeer SDK](docs/sdks/livepeer/README.md)

- [PostStream](docs/sdks/livepeer/README.md#poststream) - Create a stream
- [GetStream](docs/sdks/livepeer/README.md#getstream) - Retrieve streams
- [GetStreamID](docs/sdks/livepeer/README.md#getstreamid) - Retrieve a stream
- [PatchStreamID](docs/sdks/livepeer/README.md#patchstreamid) - Update a stream
- [DeleteStreamID](docs/sdks/livepeer/README.md#deletestreamid) - Delete a stream
- [DeleteStreamIDTerminate](docs/sdks/livepeer/README.md#deletestreamidterminate) - Terminates a live stream
- [PostStreamIDStartPull](docs/sdks/livepeer/README.md#poststreamidstartpull) - Start ingest for a pull stream
- [GetMultistreamTarget](docs/sdks/livepeer/README.md#getmultistreamtarget) - Retrieve Multistream Targets
- [PostMultistreamTarget](docs/sdks/livepeer/README.md#postmultistreamtarget) - Create a multistream target
- [GetMultistreamTargetID](docs/sdks/livepeer/README.md#getmultistreamtargetid) - Retrieve a multistream target
- [PatchMultistreamTargetID](docs/sdks/livepeer/README.md#patchmultistreamtargetid) - Update Multistream Target
- [DeleteMultistreamTargetID](docs/sdks/livepeer/README.md#deletemultistreamtargetid) - Delete a multistream target
- [GetWebhook](docs/sdks/livepeer/README.md#getwebhook) - Retrieve a Webhook
- [PostWebhook](docs/sdks/livepeer/README.md#postwebhook) - Create a webhook
- [GetWebhookID](docs/sdks/livepeer/README.md#getwebhookid) - Retrieve a webhook
- [PutWebhookID](docs/sdks/livepeer/README.md#putwebhookid) - Update a webhook
- [DeleteWebhookID](docs/sdks/livepeer/README.md#deletewebhookid) - Delete a webhook
- [GetWebhookIDLog](docs/sdks/livepeer/README.md#getwebhookidlog) - Retrieve webhook logs
- [GetWebhookIDLogLogID](docs/sdks/livepeer/README.md#getwebhookidloglogid) - Retrieve a webhook log
- [PostWebhookIDLogLogIDResend](docs/sdks/livepeer/README.md#postwebhookidloglogidresend) - Resend a webhook
- [GetAsset](docs/sdks/livepeer/README.md#getasset) - Retrieve assets
- [PostAssetRequestUpload](docs/sdks/livepeer/README.md#postassetrequestupload) - Upload an asset
- [PostAssetUploadURL](docs/sdks/livepeer/README.md#postassetuploadurl) - Upload asset via URL
- [GetAssetAssetID](docs/sdks/livepeer/README.md#getassetassetid) - Retrieves an asset
- [PatchAssetAssetID](docs/sdks/livepeer/README.md#patchassetassetid) - Patch an asset
- [DeleteAssetAssetID](docs/sdks/livepeer/README.md#deleteassetassetid) - Delete an asset
- [PostClip](docs/sdks/livepeer/README.md#postclip) - Create a clip
- [GetStreamIDClips](docs/sdks/livepeer/README.md#getstreamidclips) - Retrieve clips of a livestream
- [PostStreamIDCreateMultistreamTarget](docs/sdks/livepeer/README.md#poststreamidcreatemultistreamtarget) - Add a multistream target
- [DeleteStreamIDMultistreamTargetID](docs/sdks/livepeer/README.md#deletestreamidmultistreamtargetid) - Remove a multistream target
- [GetSessionIDClips](docs/sdks/livepeer/README.md#getsessionidclips) - Retrieve clips of a session
- [PostRoom](docs/sdks/livepeer/README.md#postroom) - Create a room
- [GetRoomID](docs/sdks/livepeer/README.md#getroomid) - Retrieve a room
- [DeleteRoomID](docs/sdks/livepeer/README.md#deleteroomid) - Delete a room
- [PostRoomIDEgress](docs/sdks/livepeer/README.md#postroomidegress) - Start room RTMP egress
- [DeleteRoomIDEgress](docs/sdks/livepeer/README.md#deleteroomidegress) - Stop room RTMP egress
- [PostRoomIDUser](docs/sdks/livepeer/README.md#postroomiduser) - Create a room user
- [GetRoomIDUserUserID](docs/sdks/livepeer/README.md#getroomiduseruserid) - Get user details
- [PutRoomIDUserUserID](docs/sdks/livepeer/README.md#putroomiduseruserid) - Update a room user
- [DeleteRoomIDUserUserID](docs/sdks/livepeer/README.md#deleteroomiduseruserid) - Remove a user from the room
- [GetDataViewsQuery](docs/sdks/livepeer/README.md#getdataviewsquery) - Query viewership metrics
- [GetDataViewsQueryCreator](docs/sdks/livepeer/README.md#getdataviewsquerycreator) - Query creator viewership metrics
- [GetDataViewsQueryTotalPlaybackID](docs/sdks/livepeer/README.md#getdataviewsquerytotalplaybackid) - Query public total views metrics
- [GetDataUsageQuery](docs/sdks/livepeer/README.md#getdatausagequery) - Query usage metrics
- [GetSession](docs/sdks/livepeer/README.md#getsession) - Retrieve sessions
- [GetSessionID](docs/sdks/livepeer/README.md#getsessionid) - Retrieve a session
- [GetStreamParentIDSessions](docs/sdks/livepeer/README.md#getstreamparentidsessions) - Retrieve Recorded Sessions
- [PostAccessControlSigningKey](docs/sdks/livepeer/README.md#postaccesscontrolsigningkey) - Create a signing key
- [GetAccessControlSigningKey](docs/sdks/livepeer/README.md#getaccesscontrolsigningkey) - Retrieves signing keys
- [DeleteAccessControlSigningKeyKeyID](docs/sdks/livepeer/README.md#deleteaccesscontrolsigningkeykeyid) - Delete Signing Key
- [GetAccessControlSigningKeyKeyID](docs/sdks/livepeer/README.md#getaccesscontrolsigningkeykeyid) - Retrieves a signing key
- [PatchAccessControlSigningKeyKeyID](docs/sdks/livepeer/README.md#patchaccesscontrolsigningkeykeyid) - Update a signing key
- [GetTask](docs/sdks/livepeer/README.md#gettask) - Retrieve Tasks
- [GetTaskTaskID](docs/sdks/livepeer/README.md#gettasktaskid) - Retrieve a Task
- [PostTranscode](docs/sdks/livepeer/README.md#posttranscode) - Transcode a video
- [GetPlaybackID](docs/sdks/livepeer/README.md#getplaybackid) - Retrieve Playback Info

## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
 "context"
 "errors"
 livepeer "github.com/livepeer/livepeer-go"
 "github.com/livepeer/livepeer-go/models/components"
 "github.com/livepeer/livepeer-go/models/sdkerrors"
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

  var e *sdkerrors.Error
  if errors.As(err, &e) {
   // handle error
   log.Fatal(e.Error())
  }

  var e *sdkerrors.SDKError
  if errors.As(err, &e) {
   // handle error
   log.Fatal(e.Error())
  }
 }
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
 "context"
 livepeer "github.com/livepeer/livepeer-go"
 "github.com/livepeer/livepeer-go/models/components"
 "log"
)

func main() {
 s := livepeer.New(
  livepeer.WithServerURL("https://livepeer.studio/api"),
  livepeer.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
 )

 ctx := context.Background()
 res, err := s.PostStream(ctx, components.NewStreamPayload{
  Name:   "test_stream",
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
 Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
 "net/http"
 "time"
 "github.com/myorg/your-go-sdk"
)

var (
 httpClient = &http.Client{Timeout: 30 * time.Second}
 sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
