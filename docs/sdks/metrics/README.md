# Metrics
(*Metrics*)

### Available Operations

* [GetViewershipsMetrics](#getviewershipsmetrics) - Query viewership metrics
* [GetCreatorMetrics](#getcreatormetrics) - Query creator viewership metrics
* [GetPublicTotalViewsMetrics](#getpublictotalviewsmetrics) - Query public total views metrics
* [GetUsageMetrics](#getusagemetrics) - Query usage metrics

## GetViewershipsMetrics

Query viewership metrics

### Example Usage

```go
package main

import(
	"context"
	"log"
	"livepeer"
	"livepeer/models/components"
	"livepeer/models/operations"
	"livepeer/types"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Metrics.GetViewershipsMetrics(ctx, operations.GetViewershipsMetricsRequest{
        From: operations.CreateFromInteger(
        599370,
        ),
        To: operations.CreateToInteger(
        750430,
        ),
        BreakdownBy: []operations.BreakdownBy{
            operations.BreakdownByCountry,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetViewershipsMetricsRequest](../../models/operations/getviewershipsmetricsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetViewershipsMetricsResponse](../../models/operations/getviewershipsmetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetCreatorMetrics

Query creator viewership metrics

### Example Usage

```go
package main

import(
	"context"
	"log"
	"livepeer"
	"livepeer/models/components"
	"livepeer/models/operations"
	"livepeer/types"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Metrics.GetCreatorMetrics(ctx, operations.GetCreatorMetricsRequest{
        From: operations.CreateQueryParamFromDateTime(
        types.MustTimeFromString("2023-01-09T08:08:10.790Z"),
        ),
        To: operations.CreateQueryParamToDateTime(
        types.MustTimeFromString("2021-02-11T10:47:49.402Z"),
        ),
        BreakdownBy: []operations.QueryParamBreakdownBy{
            operations.QueryParamBreakdownByContinent,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetCreatorMetricsRequest](../../models/operations/getcreatormetricsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetCreatorMetricsResponse](../../models/operations/getcreatormetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPublicTotalViewsMetrics

Query public total views metrics

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


    var playbackID string = "string"

    ctx := context.Background()
    res, err := s.Metrics.GetPublicTotalViewsMetrics(ctx, playbackID)
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

**[*operations.GetPublicTotalViewsMetricsResponse](../../models/operations/getpublictotalviewsmetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetUsageMetrics

Query usage metrics

### Example Usage

```go
package main

import(
	"context"
	"log"
	"livepeer"
	"livepeer/models/components"
	"livepeer/models/operations"
)

func main() {
    s := livepeer.New(
        livepeer.WithSecurity(""),
    )


    var from *int64 = 169019

    var to *int64 = 623790

    var timeStep *operations.GetUsageMetricsQueryParamTimeStep = operations.GetUsageMetricsQueryParamTimeStepHour

    var creatorID *string = "string"

    ctx := context.Background()
    res, err := s.Metrics.GetUsageMetrics(ctx, from, to, timeStep, creatorID)
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
| sdkerrors.SDKError | 400-600            | */*                |
