# Webhook
(*Webhook*)

## Overview

Operations related to webhook api

### Available Operations

* [GetAll](#getall) - Retrieve a Webhook
* [Create](#create) - Create a webhook
* [Get](#get) - Retrieve a webhook
* [Update](#update) - Update a webhook
* [Delete](#delete) - Delete a webhook
* [GetLogs](#getlogs) - Retrieve webhook logs
* [GetLog](#getlog) - Retrieve a webhook log
* [ResendLog](#resendlog) - Resend a webhook

## GetAll

Retrieve a Webhook

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebhooks" method="get" path="/webhook" -->
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

    res, err := s.Webhook.GetAll(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWebhooksResponse](../../models/operations/getwebhooksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

To create a new webhook, you need to make an API call with the events you want to listen for and the URL that will be called when those events occur.


### Example Usage

<!-- UsageSnippet language="go" operationID="createWebhook" method="post" path="/webhook" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Webhook.Create(ctx, components.WebhookInput{
        Name: "test_webhook",
        ProjectID: livepeergo.Pointer("aac12556-4d65-4d34-9fb6-d1f0985eb0a9"),
        Events: []components.Events{
            components.EventsStreamStarted,
            components.EventsStreamIdle,
        },
        URL: "https://my-service.com/webhook",
        SharedSecret: livepeergo.Pointer("my-secret"),
        StreamID: livepeergo.Pointer("de7818e7-610a-4057-8f6f-b785dc1e6f88"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.WebhookInput](../../models/components/webhookinput.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreateWebhookResponse](../../models/operations/createwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a webhook

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebhook" method="get" path="/webhook/{id}" -->
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

    res, err := s.Webhook.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWebhookResponse](../../models/operations/getwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a webhook

### Example Usage

<!-- UsageSnippet language="go" operationID="updateWebhook" method="put" path="/webhook/{id}" -->
```go
package main

import(
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := livepeergo.New(
        livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Webhook.Update(ctx, "<id>", components.WebhookInput{
        Name: "test_webhook",
        ProjectID: livepeergo.Pointer("aac12556-4d65-4d34-9fb6-d1f0985eb0a9"),
        Events: []components.Events{
            components.EventsStreamStarted,
            components.EventsStreamIdle,
        },
        URL: "https://my-service.com/webhook",
        SharedSecret: livepeergo.Pointer("my-secret"),
        StreamID: livepeergo.Pointer("de7818e7-610a-4057-8f6f-b785dc1e6f88"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `webhook`                                                          | [components.WebhookInput](../../models/components/webhookinput.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.UpdateWebhookResponse](../../models/operations/updatewebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a webhook

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteWebhook" method="delete" path="/webhook/{id}" -->
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

    res, err := s.Webhook.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteWebhookResponse](../../models/operations/deletewebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLogs

Retrieve webhook logs

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebhookLogs" method="get" path="/webhook/{id}/log" -->
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

    res, err := s.Webhook.GetLogs(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWebhookLogsResponse](../../models/operations/getwebhooklogsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLog

Retrieve a webhook log

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebhookLog" method="get" path="/webhook/{id}/log/{logId}" -->
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

    res, err := s.Webhook.GetLog(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookLog != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `logID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWebhookLogResponse](../../models/operations/getwebhooklogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ResendLog

Use this API to resend the same webhook request. This is useful when
developing and debugging, allowing you to easily repeat the same webhook
to check or fix the behaviour in your handler.


### Example Usage

<!-- UsageSnippet language="go" operationID="resendWebhook" method="post" path="/webhook/{id}/log/{logId}/resend" -->
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

    res, err := s.Webhook.ResendLog(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookLog != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `logID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ResendWebhookResponse](../../models/operations/resendwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |