# MultistreamTarget
(*MultistreamTarget*)

### Available Operations

* [GetMultistreamTargets](#getmultistreamtargets) - Retrieve Multistream Targets
* [CreateMultistreamTarget](#createmultistreamtarget) - Create a multistream target
* [DeleteMultistreamTarget](#deletemultistreamtarget) - Delete a multistream target
* [GetMultistreamTarget](#getmultistreamtarget) - Retrieve a multistream target
* [UpdateMultistreamTarget](#updatemultistreamtarget) - Update Multistream Target

## GetMultistreamTargets

Retrieve Multistream Targets

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

    ctx := context.Background()
    res, err := s.MultistreamTarget.GetMultistreamTargets(ctx)
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

**[*operations.GetMultistreamTargetsResponse](../../models/operations/getmultistreamtargetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateMultistreamTarget

Create a multistream target

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

    ctx := context.Background()
    res, err := s.MultistreamTarget.CreateMultistreamTarget(ctx, components.MultistreamTargetInput{
        Name: livepeer.String("My Multistream Target"),
        URL: "rtmps://live.my-service.tv/channel/secretKey",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.MultistreamTargetInput](../../models/components/multistreamtargetinput.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateMultistreamTargetResponse](../../models/operations/createmultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteMultistreamTarget

Delete a multistream target

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
    res, err := s.MultistreamTarget.DeleteMultistreamTarget(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the multistream target                          |


### Response

**[*operations.DeleteMultistreamTargetResponse](../../models/operations/deletemultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetMultistreamTarget

Retrieve a multistream target

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
    res, err := s.MultistreamTarget.GetMultistreamTarget(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.MultistreamTarget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the multistream target                          |


### Response

**[*operations.GetMultistreamTargetResponse](../../models/operations/getmultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateMultistreamTarget

Update Multistream Target

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

    multistreamTargetPatchPayload := components.MultistreamTargetPatchPayload{
        Name: livepeer.String("My Multistream Target"),
        URL: "rtmps://live.my-service.tv/channel/secretKey",
    }

    ctx := context.Background()
    res, err := s.MultistreamTarget.UpdateMultistreamTarget(ctx, id, multistreamTargetPatchPayload)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | ID of the multistream target                                                                         |
| `multistreamTargetPatchPayload`                                                                      | [components.MultistreamTargetPatchPayload](../../models/components/multistreamtargetpatchpayload.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |


### Response

**[*operations.UpdateMultistreamTargetResponse](../../models/operations/updatemultistreamtargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
