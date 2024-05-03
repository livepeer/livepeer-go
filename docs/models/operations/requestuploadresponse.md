# RequestUploadResponse


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `HTTPMeta`                                                                    | [components.HTTPMetadata](../../models/components/httpmetadata.md)            | :heavy_check_mark:                                                            | N/A                                                                           |
| `Data`                                                                        | [*operations.RequestUploadData](../../models/operations/requestuploaddata.md) | :heavy_minus_sign:                                                            | Success                                                                       |
| `Error`                                                                       | **sdkerrors.Error*                                                            | :heavy_minus_sign:                                                            | Error                                                                         |