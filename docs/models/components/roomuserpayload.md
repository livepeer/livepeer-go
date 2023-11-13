# RoomUserPayload


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Name`                                                  | *string*                                                | :heavy_check_mark:                                      | Display name                                            | name                                                    |
| `CanPublish`                                            | **bool*                                                 | :heavy_minus_sign:                                      | Whether a user is allowed to publish audio/video tracks | true                                                    |
| `Metadata`                                              | **string*                                               | :heavy_minus_sign:                                      | User defined payload to store for the participant       | host user                                               |