# LiveVideoToVideoResponse

Response model for live video-to-video generation.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `SubscribeURL`                                             | *string*                                                   | :heavy_check_mark:                                         | Source URL of the incoming stream to subscribe to          |
| `PublishURL`                                               | *string*                                                   | :heavy_check_mark:                                         | Destination URL of the outgoing stream to publish to       |
| `ControlURL`                                               | **string*                                                  | :heavy_minus_sign:                                         | URL for updating the live video-to-video generation        |
| `EventsURL`                                                | **string*                                                  | :heavy_minus_sign:                                         | URL for subscribing to events for pipeline status and logs |
| `RequestID`                                                | **string*                                                  | :heavy_minus_sign:                                         | The ID generated for this request                          |
| `ManifestID`                                               | **string*                                                  | :heavy_minus_sign:                                         | Orchestrator manifest ID for this request                  |