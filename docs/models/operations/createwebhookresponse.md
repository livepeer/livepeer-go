# CreateWebhookResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `Webhook`                                                          | [*components.Webhook](../../models/components/webhook.md)          | :heavy_minus_sign:                                                 | Success                                                            |
| `Error`                                                            | **sdkerrors.Error*                                                 | :heavy_minus_sign:                                                 | Error                                                              |