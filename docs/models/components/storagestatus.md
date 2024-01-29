# StorageStatus


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Phase`                                              | [components.Phase](../../models/components/phase.md) | :heavy_check_mark:                                   | Phase of the asset storage                           |
| `Progress`                                           | **float64*                                           | :heavy_minus_sign:                                   | Current progress of the task updating the storage.   |
| `ErrorMessage`                                       | **string*                                            | :heavy_minus_sign:                                   | Error message if the last storage changed failed.    |
| `Tasks`                                              | [components.Tasks](../../models/components/tasks.md) | :heavy_check_mark:                                   | N/A                                                  |