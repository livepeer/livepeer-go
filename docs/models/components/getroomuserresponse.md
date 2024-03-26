# GetRoomUserResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ID`                                                    | **string*                                               | :heavy_minus_sign:                                      | The ID of the user                                      | d32ae9e6-c459-4931-9898-e86e2f5e7e16                    |
| `JoinedAt`                                              | **int64*                                                | :heavy_minus_sign:                                      | Timestamp (in milliseconds) at which the user joined    | 1687517025261                                           |
| `Name`                                                  | **string*                                               | :heavy_minus_sign:                                      | The display name of the user                            | name                                                    |
| `IsPublisher`                                           | **bool*                                                 | :heavy_minus_sign:                                      | Whether a user is allowed to publish audio/video tracks | true                                                    |
| `Metadata`                                              | **string*                                               | :heavy_minus_sign:                                      | User defined payload to store for the participant       |                                                         |