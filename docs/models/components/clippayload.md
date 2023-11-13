# ClipPayload


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `PlaybackID`                               | *string*                                   | :heavy_check_mark:                         | Playback ID of the stream or asset to clip |
| `StartTime`                                | *float64*                                  | :heavy_check_mark:                         | Start time of the clip in milliseconds     |
| `EndTime`                                  | **float64*                                 | :heavy_minus_sign:                         | End time of the clip in milliseconds       |
| `Name`                                     | **string*                                  | :heavy_minus_sign:                         | Name of the clip                           |
| `SessionID`                                | **string*                                  | :heavy_minus_sign:                         | Session ID of the stream to clip           |