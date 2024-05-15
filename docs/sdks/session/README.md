# Session
(*Session*)

## Overview

Operations related to session api

### Available Operations

* [GetClips](#getclips) - Retrieve clips of a session
* [GetAll](#getall) - Retrieve sessions
* [Get](#get) - Retrieve a session
* [GetRecorded](#getrecorded) - Retrieve Recorded Sessions

## GetClips

Retrieve clips of a session

### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
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
    res, err := s.Session.GetClips(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the parent session                              |


### Response

**[*operations.GetSessionClipsResponse](../../models/operations/getsessionclipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAll

Retrieve sessions

### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    
    ctx := context.Background()
    res, err := s.Session.GetAll(ctx)
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

**[*operations.GetSessionsResponse](../../models/operations/getsessionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Retrieve a session

### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
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
    res, err := s.Session.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Session != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the session                                     |


### Response

**[*operations.GetSessionResponse](../../models/operations/getsessionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRecorded

Retrieve Recorded Sessions

### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
	livepeergo "github.com/livepeer/livepeer-go"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var parentID string = "<value>"

    var record *int64 = livepeergo.Int64(1)
    
    ctx := context.Background()
    res, err := s.Session.GetRecorded(ctx, parentID, record)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `parentID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | ID of the parent stream                                                |                                                                        |
| `record`                                                               | **int64*                                                               | :heavy_minus_sign:                                                     | Flag indicating if the response should only include recorded<br/>sessions<br/> | 1                                                                      |


### Response

**[*operations.GetRecordedSessionsResponse](../../models/operations/getrecordedsessionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
