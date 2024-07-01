# WebhookStatus

status of webhook


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `LastFailure`                                                        | [*components.LastFailure](../../models/components/lastfailure.md)    | :heavy_minus_sign:                                                   | failure timestamp and error message with status code                 |                                                                      |
| `LastTriggeredAt`                                                    | **float64*                                                           | :heavy_minus_sign:                                                   | Timestamp (in milliseconds) at which the webhook last was<br/>triggered<br/> | 1587667174725                                                        |