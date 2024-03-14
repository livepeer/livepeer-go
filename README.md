# github.com/livepeer/livepeer-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/livepeer/livepeer-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostStream(ctx, components.NewStreamPayload{
		Name:   "test_stream",
		Record: livepeergo.Bool(false),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Livepeer SDK](docs/sdks/livepeer/README.md)

* [PostStream](docs/sdks/livepeer/README.md#poststream) - Create a stream
* [GetStream](docs/sdks/livepeer/README.md#getstream) - Retrieve streams
* [GetStreamID](docs/sdks/livepeer/README.md#getstreamid) - Retrieve a stream
* [PatchStreamID](docs/sdks/livepeer/README.md#patchstreamid) - Update a stream
* [DeleteStreamID](docs/sdks/livepeer/README.md#deletestreamid) - Delete a stream
* [DeleteStreamIDTerminate](docs/sdks/livepeer/README.md#deletestreamidterminate) - Terminates a live stream
* [PostStreamIDStartPull](docs/sdks/livepeer/README.md#poststreamidstartpull) - Start ingest for a pull stream
* [GetMultistreamTarget](docs/sdks/livepeer/README.md#getmultistreamtarget) - Retrieve Multistream Targets
* [PostMultistreamTarget](docs/sdks/livepeer/README.md#postmultistreamtarget) - Create a multistream target
* [GetMultistreamTargetID](docs/sdks/livepeer/README.md#getmultistreamtargetid) - Retrieve a multistream target
* [PatchMultistreamTargetID](docs/sdks/livepeer/README.md#patchmultistreamtargetid) - Update Multistream Target
* [DeleteMultistreamTargetID](docs/sdks/livepeer/README.md#deletemultistreamtargetid) - Delete a multistream target
* [GetWebhook](docs/sdks/livepeer/README.md#getwebhook) - Retrieve a Webhook
* [PostWebhook](docs/sdks/livepeer/README.md#postwebhook) - Create a webhook
* [GetWebhookID](docs/sdks/livepeer/README.md#getwebhookid) - Retrieve a webhook
* [PutWebhookID](docs/sdks/livepeer/README.md#putwebhookid) - Update a webhook
* [DeleteWebhookID](docs/sdks/livepeer/README.md#deletewebhookid) - Delete a webhook
* [GetWebhookIDLog](docs/sdks/livepeer/README.md#getwebhookidlog) - Retrieve webhook logs
* [GetWebhookIDLogLogID](docs/sdks/livepeer/README.md#getwebhookidloglogid) - Retrieve a webhook log
* [PostWebhookIDLogLogIDResend](docs/sdks/livepeer/README.md#postwebhookidloglogidresend) - Resend a webhook
* [GetAsset](docs/sdks/livepeer/README.md#getasset) - Retrieve assets
* [PostAssetRequestUpload](docs/sdks/livepeer/README.md#postassetrequestupload) - Upload an asset
* [PostAssetUploadURL](docs/sdks/livepeer/README.md#postassetuploadurl) - Upload asset via URL
* [GetAssetAssetID](docs/sdks/livepeer/README.md#getassetassetid) - Retrieves an asset
* [PatchAssetAssetID](docs/sdks/livepeer/README.md#patchassetassetid) - Patch an asset
* [DeleteAssetAssetID](docs/sdks/livepeer/README.md#deleteassetassetid) - Delete an asset
* [PostClip](docs/sdks/livepeer/README.md#postclip) - Create a clip
* [GetStreamIDClips](docs/sdks/livepeer/README.md#getstreamidclips) - Retrieve clips of a livestream
* [PostStreamIDCreateMultistreamTarget](docs/sdks/livepeer/README.md#poststreamidcreatemultistreamtarget) - Add a multistream target
* [DeleteStreamIDMultistreamTargetID](docs/sdks/livepeer/README.md#deletestreamidmultistreamtargetid) - Remove a multistream target
* [GetSessionIDClips](docs/sdks/livepeer/README.md#getsessionidclips) - Retrieve clips of a session
* [PostRoom](docs/sdks/livepeer/README.md#postroom) - Create a room
* [GetRoomID](docs/sdks/livepeer/README.md#getroomid) - Retrieve a room
* [DeleteRoomID](docs/sdks/livepeer/README.md#deleteroomid) - Delete a room
* [PostRoomIDEgress](docs/sdks/livepeer/README.md#postroomidegress) - Start room RTMP egress
* [DeleteRoomIDEgress](docs/sdks/livepeer/README.md#deleteroomidegress) - Stop room RTMP egress
* [PostRoomIDUser](docs/sdks/livepeer/README.md#postroomiduser) - Create a room user
* [GetRoomIDUserUserID](docs/sdks/livepeer/README.md#getroomiduseruserid) - Get user details
* [PutRoomIDUserUserID](docs/sdks/livepeer/README.md#putroomiduseruserid) - Update a room user
* [DeleteRoomIDUserUserID](docs/sdks/livepeer/README.md#deleteroomiduseruserid) - Remove a user from the room
* [GetDataViewsQuery](docs/sdks/livepeer/README.md#getdataviewsquery) - Query viewership metrics
* [GetDataViewsQueryCreator](docs/sdks/livepeer/README.md#getdataviewsquerycreator) - Query creator viewership metrics
* [GetDataViewsQueryTotalPlaybackID](docs/sdks/livepeer/README.md#getdataviewsquerytotalplaybackid) - Query public total views metrics
* [GetDataUsageQuery](docs/sdks/livepeer/README.md#getdatausagequery) - Query usage metrics
* [GetSession](docs/sdks/livepeer/README.md#getsession) - Retrieve sessions
* [GetSessionID](docs/sdks/livepeer/README.md#getsessionid) - Retrieve a session
* [GetStreamParentIDSessions](docs/sdks/livepeer/README.md#getstreamparentidsessions) - Retrieve Recorded Sessions
* [PostAccessControlSigningKey](docs/sdks/livepeer/README.md#postaccesscontrolsigningkey) - Create a signing key
* [GetAccessControlSigningKey](docs/sdks/livepeer/README.md#getaccesscontrolsigningkey) - Retrieves signing keys
* [DeleteAccessControlSigningKeyKeyID](docs/sdks/livepeer/README.md#deleteaccesscontrolsigningkeykeyid) - Delete Signing Key
* [GetAccessControlSigningKeyKeyID](docs/sdks/livepeer/README.md#getaccesscontrolsigningkeykeyid) - Retrieves a signing key
* [PatchAccessControlSigningKeyKeyID](docs/sdks/livepeer/README.md#patchaccesscontrolsigningkeykeyid) - Update a signing key
* [GetTask](docs/sdks/livepeer/README.md#gettask) - Retrieve Tasks
* [GetTaskTaskID](docs/sdks/livepeer/README.md#gettasktaskid) - Retrieve a Task
* [PostTranscode](docs/sdks/livepeer/README.md#posttranscode) - Transcode a video
* [GetPlaybackID](docs/sdks/livepeer/README.md#getplaybackid) - Retrieve Playback Info
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
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
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://livepeer.studio/api` | None |

#### Example

```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithServerIndex(0),
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostStream(ctx, components.NewStreamPayload{
		Name:   "test_stream",
		Record: livepeergo.Bool(false),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithServerURL("https://livepeer.studio/api"),
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostStream(ctx, components.NewStreamPayload{
		Name:   "test_stream",
		Record: livepeergo.Bool(false),
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
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `APIKey`    | http        | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostStream(ctx, components.NewStreamPayload{
		Name:   "test_stream",
		Record: livepeergo.Bool(false),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
