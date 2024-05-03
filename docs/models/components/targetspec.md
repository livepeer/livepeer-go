# TargetSpec

Inline multistream target object. Will automatically
create the target resource to be used by the created
stream.



## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Name`                                                      | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         | My target                                                   |
| `URL`                                                       | *string*                                                    | :heavy_check_mark:                                          | Livepeer-compatible multistream target URL (RTMP(S) or SRT) | rtmps://live.my-service.tv/channel/secretKey                |