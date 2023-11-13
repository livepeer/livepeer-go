# StorageStatus


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Phase`                                              | [components.Phase](../../models/components/phase.md) | :heavy_check_mark:                                   | Phase of the asset storage                           | ready                                                |
| `Progress`                                           | **float64*                                           | :heavy_minus_sign:                                   | Current progress of the task updating the storage.   | 0.5                                                  |
| `ErrorMessage`                                       | **string*                                            | :heavy_minus_sign:                                   | Error message if the last storage changed failed.    | Failed to export to IPFS                             |
| `Tasks`                                              | [components.Tasks](../../models/components/tasks.md) | :heavy_check_mark:                                   | N/A                                                  |                                                      |