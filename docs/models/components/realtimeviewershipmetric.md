# RealtimeViewershipMetric

An individual metric about realtime viewership of a stream/asset.



## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `PlaybackID`                                | **string*                                   | :heavy_minus_sign:                          | The playback ID associated with the metric. | 1bde4o2i6xycudoy                            |
| `Device`                                    | **string*                                   | :heavy_minus_sign:                          | The device used by the viewer.              | iPhone                                      |
| `Browser`                                   | **string*                                   | :heavy_minus_sign:                          | The browser used by the viewer.             | Safari                                      |
| `Country`                                   | **string*                                   | :heavy_minus_sign:                          | The country where the viewer is located.    | United States                               |
| `ViewCount`                                 | *int64*                                     | :heavy_check_mark:                          | The number of views for the stream/asset.   | 100                                         |
| `ErrorRate`                                 | *float64*                                   | :heavy_check_mark:                          | The error rate for the stream/asset.        | 0.1                                         |