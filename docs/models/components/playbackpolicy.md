# PlaybackPolicy

Whether the playback policy for a asset or stream is public or signed


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Type`                                             | [components.Type](../../models/components/type.md) | :heavy_check_mark:                                 | N/A                                                |                                                    |
| `WebhookID`                                        | **string*                                          | :heavy_minus_sign:                                 | ID of the webhook to use for playback policy       | 3e02c844-d364-4d48-b401-24b2773b5d6c               |
| `WebhookContext`                                   | map[string]*interface{}*                           | :heavy_minus_sign:                                 | User-defined webhook context                       | [object Object]                                    |