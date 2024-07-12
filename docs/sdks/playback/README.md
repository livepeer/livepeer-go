# Playback
(*Playback*)

## Overview

Operations related to playback api

### Available Operations

* [Get](#get) - Retrieve Playback Info

## Get

Retrieve Playback Info

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
    var id string = "<value>"
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The playback ID from the asset or livestream, e.g. `eaw4nk06ts2d0mzb`. |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |


### Response

**[*operations.GetPlaybackInfoResponse](../../models/operations/getplaybackinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
