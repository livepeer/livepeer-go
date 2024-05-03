# AccessControl
(*AccessControl*)

## Overview

Operations related to access control/signing keys api

### Available Operations

* [Create](#create) - Create a signing key
* [GetAll](#getall) - Retrieves signing keys
* [Delete](#delete) - Delete Signing Key
* [Get](#get) - Retrieves a signing key
* [Update](#update) - Update a signing key

## Create

The publicKey is a representation of the public key, encoded as base 64 and is passed as a string, and  the privateKey is displayed only on creation. This is the only moment where the client can save the private key, otherwise it will be lost. Remember to decode your string when signing JWTs.
Up to 10 signing keys can be generated, after that you must delete at least one signing key to create a new one.


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
    res, err := s.AccessControl.Create(ctx)
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


### Response

**[*operations.CreateSigningKeyResponse](../../models/operations/createsigningkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAll

Retrieves signing keys

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
    res, err := s.AccessControl.GetAll(ctx)
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

Delete Signing Key

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


    var keyID string = "<value>"

    ctx := context.Background()
    res, err := s.AccessControl.Delete(ctx, keyID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Retrieves a signing key

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


    var keyID string = "<value>"

    ctx := context.Background()
    res, err := s.AccessControl.Get(ctx, keyID)
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update a signing key

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


    var keyID string = "<value>"

    requestBody := operations.UpdateSigningKeyRequestBody{}

    ctx := context.Background()
    res, err := s.AccessControl.Update(ctx, keyID, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |
