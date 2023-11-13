# FfmpegProfile

LMPS ffmpeg profile


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Width`                                                   | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       | 1280                                                      |
| `Name`                                                    | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       | 720p                                                      |
| `Height`                                                  | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       | 720                                                       |
| `Bitrate`                                                 | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       | 4000                                                      |
| `Fps`                                                     | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       | 30                                                        |
| `FpsDen`                                                  | **int64*                                                  | :heavy_minus_sign:                                        | N/A                                                       | 1                                                         |
| `Gop`                                                     | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       | 60                                                        |
| `Profile`                                                 | [*components.Profile](../../models/components/profile.md) | :heavy_minus_sign:                                        | N/A                                                       | H264High                                                  |
| `Encoder`                                                 | [*components.Encoder](../../models/components/encoder.md) | :heavy_minus_sign:                                        | N/A                                                       | h264                                                      |