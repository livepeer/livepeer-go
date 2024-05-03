# GetWebhookLogResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `WebhookLog`                                                       | [*components.WebhookLog](../../models/components/webhooklog.md)    | :heavy_minus_sign:                                                 | Success                                                            |
| `Error`                                                            | **sdkerrors.Error*                                                 | :heavy_minus_sign:                                                 | Error                                                              |