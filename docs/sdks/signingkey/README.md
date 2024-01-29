# SigningKey
(*SigningKey*)

### Available Operations

* [Create](#create) - Create a signing key
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
    res, err := s.SigningKey.Create(ctx)
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
| sdkerrors.SDKError | 400-600            | */*                |

## Delete

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
    res, err := s.SigningKey.Delete(ctx, keyID)
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

## Get

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
    res, err := s.SigningKey.Get(ctx, keyID)
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

## Update

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
    res, err := s.SigningKey.Update(ctx, keyID, requestBody)
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
