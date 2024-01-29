# PlaybackPolicy

Whether the playback policy for a asset or stream is public or signed


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Type`                                             | [components.Type](../../models/components/type.md) | :heavy_check_mark:                                 | N/A                                                |
| `WebhookID`                                        | **string*                                          | :heavy_minus_sign:                                 | ID of the webhook to use for playback policy       |
| `WebhookContext`                                   | map[string]*interface{}*                           | :heavy_minus_sign:                                 | User-defined webhook context                       |