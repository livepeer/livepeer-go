# Task
(*Task*)

## Overview

Operations related to tasks api

### Available Operations

* [GetAll](#getall) - Retrieve Tasks
* [Get](#get) - Retrieve a Task

## GetAll

Retrieve Tasks

### Example Usage

<!-- UsageSnippet language="go" operationID="getTasks" method="get" path="/task" -->
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

    res, err := s.Task.GetAll(ctx)
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

**[*operations.GetTasksResponse](../../models/operations/gettasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a Task

### Example Usage

<!-- UsageSnippet language="go" operationID="getTask" method="get" path="/task/{taskId}" -->
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

    res, err := s.Task.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `taskID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the task                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTaskResponse](../../models/operations/gettaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |