# UsageMetric

An individual metric about usage of a user.



## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `UserID`                                  | *string*                                  | :heavy_check_mark:                        | The user ID associated with the metric    |
| `CreatorID`                               | *string*                                  | :heavy_check_mark:                        | The creator ID associated with the metric |
| `DeliveryUsageMins`                       | *float64*                                 | :heavy_check_mark:                        | The number of minutes of delivery usage   |
| `TotalUsageMins`                          | *float64*                                 | :heavy_check_mark:                        | The number of minutes of total usage      |
| `StorageUsageMins`                        | *float64*                                 | :heavy_check_mark:                        | The number of minutes of storage usage    |