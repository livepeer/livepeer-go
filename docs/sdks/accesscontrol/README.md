# AccessControl
(*AccessControl*)

### Available Operations

* [GetSigningKeys](#getsigningkeys) - Retrieves signing keys
* [CreateSigningKey](#createsigningkey) - Create a signing key
* [DeleteSigningKey](#deletesigningkey) - Delete Signing Key
* [GetSigningKey](#getsigningkey) - Retrieves a signing key
* [UpdateSigningKey](#updatesigningkey) - Update a signing key

## GetSigningKeys

Retrieves signing keys

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
    res, err := s.AccessControl.GetSigningKeys(ctx)
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

**[*operations.GetSigningKeysResponse](../../models/operations/getsigningkeysresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateSigningKey

Create a signing key

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
    res, err := s.AccessControl.CreateSigningKey(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SigningKeyResponsePayload != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.CreateSigningKeyResponse](../../models/operations/createsigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteSigningKey

Delete Signing Key

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


    var keyID string = "string"

    ctx := context.Background()
    res, err := s.AccessControl.DeleteSigningKey(ctx, keyID)
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
| `keyID`                                               | *string*                                              | :heavy_check_mark:                                    | ID of the signing key                                 |


### Response

**[*operations.DeleteSigningKeyResponse](../../models/operations/deletesigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSigningKey

Retrieves a signing key

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


    var keyID string = "string"

    ctx := context.Background()
    res, err := s.AccessControl.GetSigningKey(ctx, keyID)
    if err != nil {
        log.Fatal(err)
    }

    if res.SigningKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `keyID`                                               | *string*                                              | :heavy_check_mark:                                    | ID of the signing key                                 |


### Response

**[*operations.GetSigningKeyResponse](../../models/operations/getsigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateSigningKey

Update a signing key

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


    var keyID string = "string"

    requestBody := operations.UpdateSigningKeyRequestBody{}

    ctx := context.Background()
    res, err := s.AccessControl.UpdateSigningKey(ctx, keyID, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `keyID`                                                                                          | *string*                                                                                         | :heavy_check_mark:                                                                               | ID of the signing key                                                                            |
| `requestBody`                                                                                    | [operations.UpdateSigningKeyRequestBody](../../models/operations/updatesigningkeyrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |


### Response

**[*operations.UpdateSigningKeyResponse](../../models/operations/updatesigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
