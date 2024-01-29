# Upload

Parameters for the upload task


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `URL`                                                           | **string*                                                       | :heavy_minus_sign:                                              | URL of the asset to "upload"                                    | https://cdn.livepeer.com/ABC123/filename.mp4                    |
| `Encryption`                                                    | [*components.Encryption](../../models/components/encryption.md) | :heavy_minus_sign:                                              | N/A                                                             |                                                                 |
| `C2pa`                                                          | **bool*                                                         | :heavy_minus_sign:                                              | Decides if the output video should include C2PA signature       |                                                                 |