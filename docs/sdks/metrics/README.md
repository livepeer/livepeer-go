# Metrics
(*Metrics*)

## Overview

Operations related to metrics api

### Available Operations

* [GetRealtimeViewership](#getrealtimeviewership) - Query realtime viewership
* [GetViewership](#getviewership) - Query viewership metrics
* [GetCreatorViewership](#getcreatorviewership) - Query creator viewership metrics
* [GetPublicViewership](#getpublicviewership) - Query public total views metrics
* [GetUsage](#getusage) - Query usage metrics

## GetRealtimeViewership

Requires a private (non-CORS) API key to be used.


### Example Usage

<!-- UsageSnippet language="go" operationID="getRealtimeViewershipNow" method="get" path="/data/views/now" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Metrics.GetRealtimeViewership(ctx, nil, nil, nil)
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
| `playbackID`                                                                                                                                          | **string*                                                                                                                                             | :heavy_minus_sign:                                                                                                                                    | The playback ID to filter the query results. This can be a canonical<br/>playback ID from Livepeer assets or streams, or dStorage identifiers<br/>for assets<br/> |
| `creatorID`                                                                                                                                           | **string*                                                                                                                                             | :heavy_minus_sign:                                                                                                                                    | The creator ID to filter the query results                                                                                                            |
| `breakdownBy`                                                                                                                                         | [][operations.BreakdownBy](../../models/operations/breakdownby.md)                                                                                    | :heavy_minus_sign:                                                                                                                                    | The list of fields to break down the query results. Specify this<br/>query-string multiple times to break down by multiple fields.<br/>               |
| `opts`                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                              | :heavy_minus_sign:                                                                                                                                    | The options for this request.                                                                                                                         |

### Response

**[*operations.GetRealtimeViewershipNowResponse](../../models/operations/getrealtimeviewershipnowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetViewership

Requires a private (non-CORS) API key to be used.


### Example Usage

<!-- UsageSnippet language="go" operationID="getViewershipMetrics" method="get" path="/data/views/query" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Metrics.GetViewership(ctx, operations.GetViewershipMetricsRequest{})
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
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetViewershipMetricsResponse](../../models/operations/getviewershipmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCreatorViewership

Requires a proof of ownership to be sent in the request, which for now is just the assetId or streamId parameters (1 of those must be in the query-string).


### Example Usage

<!-- UsageSnippet language="go" operationID="getCreatorViewershipMetrics" method="get" path="/data/views/query/creator" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Metrics.GetCreatorViewership(ctx, operations.GetCreatorViewershipMetricsRequest{})
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
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetCreatorViewershipMetricsResponse](../../models/operations/getcreatorviewershipmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPublicViewership

Allows querying for the public metrics for viewership about a video.
This can be called from the frontend with a CORS key, or even
unauthenticated.


### Example Usage

<!-- UsageSnippet language="go" operationID="getPublicViewershipMetrics" method="get" path="/data/views/query/total/{playbackId}" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Metrics.GetPublicViewership(ctx, "<id>")
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
| `opts`                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                              | :heavy_minus_sign:                                                                                                                                    | The options for this request.                                                                                                                         |

### Response

**[*operations.GetPublicViewershipMetricsResponse](../../models/operations/getpublicviewershipmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUsage

Query usage metrics

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsageMetrics" method="get" path="/data/usage/query" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Metrics.GetUsage(ctx, operations.GetUsageMetricsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UsageMetric != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetUsageMetricsRequest](../../models/operations/getusagemetricsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetUsageMetricsResponse](../../models/operations/getusagemetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |