# GetUsageMetricsResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `UsageMetric`                                                      | [*components.UsageMetric](../../models/components/usagemetric.md)  | :heavy_minus_sign:                                                 | A Usage Metric object                                              |
| `Error`                                                            | **sdkerrors.Error*                                                 | :heavy_minus_sign:                                                 | Error                                                              |