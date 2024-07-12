# UsageMetric

An individual metric about usage of a user.



## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `UserID`                                  | **string*                                 | :heavy_minus_sign:                        | The user ID associated with the metric    | 1bde4o2i6xycudoy                          |
| `CreatorID`                               | **string*                                 | :heavy_minus_sign:                        | The creator ID associated with the metric | john@doe.com                              |
| `DeliveryUsageMins`                       | **float64*                                | :heavy_minus_sign:                        | Total minutes of delivery usage.          | 100                                       |
| `TotalUsageMins`                          | **float64*                                | :heavy_minus_sign:                        | Total transcoded minutes.                 | 100                                       |
| `StorageUsageMins`                        | **float64*                                | :heavy_minus_sign:                        | Total minutes of storage usage.           | 100                                       |