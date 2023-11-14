# Metrics
(*Metrics*)

### Available Operations

* [GetViewership](#getviewership) - Query viewership metrics
* [GetCreatorViewership](#getcreatorviewership) - Query creator viewership metrics
* [GetPublicTotalViews](#getpublictotalviews) - Query public total views metrics
* [GetUsage](#getusage) - Query usage metrics

## GetViewership

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
    res, err := s.Metrics.GetViewership(ctx, operations.GetViewershipsMetricsRequest{
        From: operations.CreateFromInteger(
        980301,
        ),
        To: operations.CreateToInteger(
        366854,
        ),
        BreakdownBy: []operations.BreakdownBy{
            operations.BreakdownByPlaybackID,
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

## GetCreatorViewership

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
    res, err := s.Metrics.GetCreatorViewership(ctx, operations.GetCreatorMetricsRequest{
        From: operations.CreateQueryParamFromDateTime(
        types.MustTimeFromString("2021-06-16T23:48:30.007Z"),
        ),
        To: operations.CreateQueryParamToInteger(
        702371,
        ),
        BreakdownBy: []operations.QueryParamBreakdownBy{
            operations.QueryParamBreakdownByDeviceType,
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

## GetPublicTotalViews

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
    res, err := s.Metrics.GetPublicTotalViews(ctx, playbackID)
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

## GetUsage

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


    var from *int64 = 224089

    var to *int64 = 231125

    var timeStep *operations.GetUsageMetricsQueryParamTimeStep = operations.GetUsageMetricsQueryParamTimeStepDay

    var creatorID *string = "string"

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
| sdkerrors.SDKError | 400-600            | */*                |
