# TaskExportData

Parameters for the export-data task


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Content`                                                                   | [components.Content](../../models/components/content.md)                    | :heavy_check_mark:                                                          | File content to store into IPFS                                             |
| `Ipfs`                                                                      | [*components.IpfsExportParams](../../models/components/ipfsexportparams.md) | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Type`                                                                      | **string*                                                                   | :heavy_minus_sign:                                                          | Optional type of content                                                    |
| `ID`                                                                        | **string*                                                                   | :heavy_minus_sign:                                                          | Optional ID of the content                                                  |