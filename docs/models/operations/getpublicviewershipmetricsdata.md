# GetPublicViewershipMetricsData

A simplified metric object about aggregate viewership of an
asset. Either playbackId or dStorageUrl will be set.



## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `PlaybackID`                                          | **string*                                             | :heavy_minus_sign:                                    | The playback ID associated with the metric.           | 1bde4o2i6xycudoy                                      |
| `DStorageURL`                                         | **string*                                             | :heavy_minus_sign:                                    | The URL of the distributed storage used for the asset | ipfs://QmZ4                                           |
| `ViewCount`                                           | **int64*                                              | :heavy_minus_sign:                                    | The number of views for the asset.                    | 100                                                   |
| `PlaytimeMins`                                        | **float64*                                            | :heavy_minus_sign:                                    | The total playtime in minutes for the asset.          | 10                                                    |