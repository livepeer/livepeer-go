# UploadAssetResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `HTTPMeta`                                                                | [components.HTTPMetadata](../../models/components/httpmetadata.md)        | :heavy_check_mark:                                                        | N/A                                                                       |
| `Data`                                                                    | [*operations.UploadAssetData](../../models/operations/uploadassetdata.md) | :heavy_minus_sign:                                                        | Success                                                                   |
| `Error`                                                                   | **sdkerrors.Error*                                                        | :heavy_minus_sign:                                                        | Error                                                                     |