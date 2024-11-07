# GetCreatorViewershipMetricsResponse


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `HTTPMeta`                                                                   | [components.HTTPMetadata](../../models/components/httpmetadata.md)           | :heavy_check_mark:                                                           | N/A                                                                          |
| `Data`                                                                       | [][components.ViewershipMetric](../../models/components/viewershipmetric.md) | :heavy_minus_sign:                                                           | A list of Metric objects                                                     |
| `Error`                                                                      | [*components.Error](../../models/components/error.md)                        | :heavy_minus_sign:                                                           | Error                                                                        |