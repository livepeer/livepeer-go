# Playback
(*Playback*)

### Available Operations

* [Get](#get) - Retrieve Playback Info

## Get

Retrieve Playback Info

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
    res, err := s.Playback.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.PlaybackInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The ID of the playback                                |


### Response

**[*operations.GetPlaybackInfoResponse](../../models/operations/getplaybackinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 404                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
