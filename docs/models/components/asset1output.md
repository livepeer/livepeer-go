# Asset1Output


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Type`                                                                      | [components.AssetSchemasType](../../models/components/assetschemastype.md)  | :heavy_check_mark:                                                          | N/A                                                                         |
| `URL`                                                                       | *string*                                                                    | :heavy_check_mark:                                                          | URL from which the asset was uploaded                                       |
| `GatewayURL`                                                                | **string*                                                                   | :heavy_minus_sign:                                                          | Gateway URL from asset if parsed from provided URL on upload.               |
| `Encryption`                                                                | [*components.EncryptionOutput](../../models/components/encryptionoutput.md) | :heavy_minus_sign:                                                          | N/A                                                                         |