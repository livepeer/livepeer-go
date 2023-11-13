# UsageMetric

An individual metric about usage of a user.



## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `UserID`                                  | *string*                                  | :heavy_check_mark:                        | The user ID associated with the metric    | 3e02c844-d364-4d48-b401-24b2773b5d6c      |
| `CreatorID`                               | *string*                                  | :heavy_check_mark:                        | The creator ID associated with the metric | 3e02c844-d364-4d48-b401-24b2773b5d6c      |
| `DeliveryUsageMins`                       | *float64*                                 | :heavy_check_mark:                        | The number of minutes of delivery usage   | 10                                        |
| `TotalUsageMins`                          | *float64*                                 | :heavy_check_mark:                        | The number of minutes of total usage      | 10                                        |
| `StorageUsageMins`                        | *float64*                                 | :heavy_check_mark:                        | The number of minutes of storage usage    | 10                                        |