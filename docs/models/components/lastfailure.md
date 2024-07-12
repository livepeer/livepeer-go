# LastFailure

failure timestamp and error message with status code


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Timestamp`                                                  | **float64*                                                   | :heavy_minus_sign:                                           | Timestamp (in milliseconds) at which the webhook last failed | 1587667174725                                                |
| `Error`                                                      | **string*                                                    | :heavy_minus_sign:                                           | Webhook failure error message                                | Error message                                                |
| `Response`                                                   | **string*                                                    | :heavy_minus_sign:                                           | Webhook failure response                                     | Response body                                                |
| `StatusCode`                                                 | **float64*                                                   | :heavy_minus_sign:                                           | Webhook failure status code                                  | 500                                                          |