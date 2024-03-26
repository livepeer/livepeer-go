# Source1


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Type`                                                          | [components.SourceType](../../models/components/sourcetype.md)  | :heavy_check_mark:                                              | N/A                                                             |
| `URL`                                                           | *string*                                                        | :heavy_check_mark:                                              | URL from which the asset was uploaded.                          |
| `GatewayURL`                                                    | **string*                                                       | :heavy_minus_sign:                                              | Gateway URL from asset if parsed from provided URL on upload.   |
| `Encryption`                                                    | [*components.Encryption](../../models/components/encryption.md) | :heavy_minus_sign:                                              | N/A                                                             |