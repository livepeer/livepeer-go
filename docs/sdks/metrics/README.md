# Metrics
(*Metrics*)

## Overview

Operations related to metrics api

### Available Operations

* [GetViewership](#getviewership) - Query viewership metrics
* [GetCreatorViewership](#getcreatorviewership) - Query creator viewership metrics
* [GetPublicViewership](#getpublicviewership) - Query public total views metrics
* [GetUsage](#getusage) - Query usage metrics

## GetViewership

Requires a private (non-CORS) API key to be used.


### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/operations"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.GetViewershipMetricsRequest{}
    
    ctx := context.Background()
    res, err := s.Metrics.GetViewership(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetViewershipMetricsRequest](../../models/operations/getviewershipmetricsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetViewershipMetricsResponse](../../models/operations/getviewershipmetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCreatorViewership

Requires a proof of ownership to be sent in the request, which for now is just the assetId or streamId parameters (1 of those must be in the query-string).


### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/operations"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.GetCreatorViewershipMetricsRequest{}
    
    ctx := context.Background()
    res, err := s.Metrics.GetCreatorViewership(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCreatorViewershipMetricsRequest](../../models/operations/getcreatorviewershipmetricsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetCreatorViewershipMetricsResponse](../../models/operations/getcreatorviewershipmetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPublicViewership

Allows querying for the public metrics for viewership about a video.
This can be called from the frontend with a CORS key, or even
unauthenticated.


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

    var playbackID string = "<value>"
    
    ctx := context.Background()
    res, err := s.Metrics.GetPublicViewership(ctx, playbackID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |
| `playbackID`                                                                                                                                          | *string*                                                                                                                                              | :heavy_check_mark:                                                                                                                                    | The playback ID to filter the query results. This can be a canonical<br/>playback ID from Livepeer assets or streams, or dStorage identifiers<br/>for assets<br/> |


### Response

**[*operations.GetPublicViewershipMetricsResponse](../../models/operations/getpublicviewershipmetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUsage

Query usage metrics

### Example Usage

```go
package main

import(
	"github.com/livepeer/livepeer-go/models/components"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/operations"
	"context"
	"log"
)

func main() {
    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    var from *int64 = livepeergo.Int64(224089)

    var to *int64 = livepeergo.Int64(231125)

    var timeStep *operations.GetUsageMetricsQueryParamTimeStep = operations.GetUsageMetricsQueryParamTimeStepDay.ToPointer()

    var creatorID *string = livepeergo.String("<value>")
    
    ctx := context.Background()
    res, err := s.Metrics.GetUsage(ctx, from, to, timeStep, creatorID)
    if err != nil {
        log.Fatal(err)
    }
    if res.UsageMetric != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |
| `from`                                                                                                        | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | Start millis timestamp for the query range (inclusive)<br/>                                                   |
| `to`                                                                                                          | **int64*                                                                                                      | :heavy_minus_sign:                                                                                            | End millis timestamp for the query range (exclusive)<br/>                                                     |
| `timeStep`                                                                                                    | [*operations.GetUsageMetricsQueryParamTimeStep](../../models/operations/getusagemetricsqueryparamtimestep.md) | :heavy_minus_sign:                                                                                            | The time step to aggregate viewership metrics by<br/>                                                         |
| `creatorID`                                                                                                   | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | The creator ID to filter the query results<br/>                                                               |


### Response

**[*operations.GetUsageMetricsResponse](../../models/operations/getusagemetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
