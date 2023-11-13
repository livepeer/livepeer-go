# Room
(*Room*)

### Available Operations

* [CreateRoom](#createroom) - Create a room
* [DeleteRoom](#deleteroom) - Delete a room
* [GetRoom](#getroom) - Retrieve a room
* [StopRoomEgress](#stoproomegress) - Stop room RTMP egress
* [StartRoomEgress](#startroomegress) - Start room RTMP egress
* [CreateRoomUser](#createroomuser) - Create a room user
* [DeleteRoomUser](#deleteroomuser) - Remove a user from the room
* [GetRoomUserDetails](#getroomuserdetails) - Get user details
* [UpdateRoomUserDetails](#updateroomuserdetails) - Update a room user

## CreateRoom

Create a room

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
    res, err := s.Room.CreateRoom(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateRoomResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.CreateRoomResponse](../../models/operations/createroomresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteRoom

Delete a room

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
    res, err := s.Room.DeleteRoom(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the room                                        |


### Response

**[*operations.DeleteRoomResponse](../../models/operations/deleteroomresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetRoom

Retrieve a room

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
    res, err := s.Room.GetRoom(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Room != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the room                                        |


### Response

**[*operations.GetRoomResponse](../../models/operations/getroomresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## StopRoomEgress

Stop room RTMP egress

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
    res, err := s.Room.StopRoomEgress(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the room                                        |


### Response

**[*operations.StopRoomEgressResponse](../../models/operations/stoproomegressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## StartRoomEgress

Start room RTMP egress

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

    roomEgressPayload := components.RoomEgressPayload{
        StreamID: "aac12556-4d65-4d34-9fb6-d1f0985eb0a9",
    }

    ctx := context.Background()
    res, err := s.Room.StartRoomEgress(ctx, id, roomEgressPayload)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | ID of the room                                                               |
| `roomEgressPayload`                                                          | [components.RoomEgressPayload](../../models/components/roomegresspayload.md) | :heavy_check_mark:                                                           | N/A                                                                          |


### Response

**[*operations.StartRoomEgressResponse](../../models/operations/startroomegressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateRoomUser

Create a room user

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

    roomUserPayload := components.RoomUserPayload{
        Name: "name",
        CanPublish: livepeer.Bool(true),
        Metadata: livepeer.String("host user"),
    }

    ctx := context.Background()
    res, err := s.Room.CreateRoomUser(ctx, id, roomUserPayload)
    if err != nil {
        log.Fatal(err)
    }

    if res.RoomUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `id`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | ID of the room                                                           |
| `roomUserPayload`                                                        | [components.RoomUserPayload](../../models/components/roomuserpayload.md) | :heavy_check_mark:                                                       | N/A                                                                      |


### Response

**[*operations.CreateRoomUserResponse](../../models/operations/createroomuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteRoomUser

Remove a user from the room

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

    var userID string = "string"

    ctx := context.Background()
    res, err := s.Room.DeleteRoomUser(ctx, id, userID)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the room                                        |
| `userID`                                              | *string*                                              | :heavy_check_mark:                                    | ID of the user                                        |


### Response

**[*operations.DeleteRoomUserResponse](../../models/operations/deleteroomuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetRoomUserDetails

Get user details

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

    var userID string = "string"

    ctx := context.Background()
    res, err := s.Room.GetRoomUserDetails(ctx, id, userID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRoomUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | ID of the room                                        |
| `userID`                                              | *string*                                              | :heavy_check_mark:                                    | ID of the user                                        |


### Response

**[*operations.GetRoomUserDetailsResponse](../../models/operations/getroomuserdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateRoomUserDetails

Update a room user

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

    var userID string = "string"

    roomUserUpdatePayload := components.RoomUserUpdatePayload{
        CanPublish: livepeer.Bool(true),
        Metadata: livepeer.String("host user"),
    }

    ctx := context.Background()
    res, err := s.Room.UpdateRoomUserDetails(ctx, id, userID, roomUserUpdatePayload)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | ID of the room                                                                       |
| `userID`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | ID of the user                                                                       |
| `roomUserUpdatePayload`                                                              | [components.RoomUserUpdatePayload](../../models/components/roomuserupdatepayload.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |


### Response

**[*operations.UpdateRoomUserDetailsResponse](../../models/operations/updateroomuserdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
