# GetDataViewsQueryTotalPlaybackIDResponseBody

A simplified metric object about aggregate viewership of an
asset. Either playbackId or dStorageUrl will be set.



## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `PlaybackID`                                          | **string*                                             | :heavy_minus_sign:                                    | The playback ID associated with the metric.           |
| `DStorageURL`                                         | **string*                                             | :heavy_minus_sign:                                    | The URL of the distributed storage used for the asset |
| `ViewCount`                                           | *int64*                                               | :heavy_check_mark:                                    | The number of views for the asset.                    |
| `PlaytimeMins`                                        | *float64*                                             | :heavy_check_mark:                                    | The total playtime in minutes for the asset.          |