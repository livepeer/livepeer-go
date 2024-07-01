# Status

Status of the asset


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ErrorMessage`                                                  | **string*                                                       | :heavy_minus_sign:                                              | Error message if the asset creation failed.                     |                                                                 |
| `Phase`                                                         | [components.AssetPhase](../../models/components/assetphase.md)  | :heavy_check_mark:                                              | Phase of the asset                                              |                                                                 |
| `Progress`                                                      | **float64*                                                      | :heavy_minus_sign:                                              | Current progress of the task creating this asset.               |                                                                 |
| `UpdatedAt`                                                     | *float64*                                                       | :heavy_check_mark:                                              | Timestamp (in milliseconds) at which the asset was last updated | 1587667174725                                                   |