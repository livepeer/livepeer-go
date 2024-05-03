# WebhookInput


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Name`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | test_webhook                                             |
| `Events`                                                 | [][components.Events](../../models/components/events.md) | :heavy_minus_sign:                                       | N/A                                                      | [<br/>"stream.started",<br/>"stream.idle"<br/>]          |
| `URL`                                                    | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | https://my-service.com/webhook                           |
| `SharedSecret`                                           | **string*                                                | :heavy_minus_sign:                                       | shared secret used to sign the webhook payload           | my-secret                                                |
| `StreamID`                                               | **string*                                                | :heavy_minus_sign:                                       | streamId of the stream on which the webhook is applied   | de7818e7-610a-4057-8f6f-b785dc1e6f88                     |