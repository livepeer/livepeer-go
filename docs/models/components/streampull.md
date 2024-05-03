# StreamPull

Configuration for a stream that should be actively pulled from an
external source, rather than pushed to Livepeer. If specified, the
stream will not have a streamKey.


## Fields

| Field                                                                                                                           | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     | Example                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `Source`                                                                                                                        | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | URL from which to pull from.                                                                                                    | https://myservice.com/live/stream.flv                                                                                           |
| `Headers`                                                                                                                       | map[string]*string*                                                                                                             | :heavy_minus_sign:                                                                                                              | Headers to be sent with the request to the pull source.                                                                         | {<br/>"Authorization": "Bearer 123"<br/>}                                                                                       |
| `Location`                                                                                                                      | [*components.StreamLocation](../../models/components/streamlocation.md)                                                         | :heavy_minus_sign:                                                                                                              | Approximate location of the pull source. The location is used to<br/>determine the closest Livepeer region to pull the stream from. |                                                                                                                                 |