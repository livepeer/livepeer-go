# TaskIpfs


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `VideoFileCid`                                                | *string*                                                      | :heavy_check_mark:                                            | IPFS CID of the exported video file                           |
| `VideoFileURL`                                                | **string*                                                     | :heavy_minus_sign:                                            | URL for the file with the IPFS protocol                       |
| `VideoFileGatewayURL`                                         | **string*                                                     | :heavy_minus_sign:                                            | URL to access file via HTTP through an IPFS gateway           |
| `NftMetadataCid`                                              | **string*                                                     | :heavy_minus_sign:                                            | IPFS CID of the default metadata exported for the video       |
| `NftMetadataURL`                                              | **string*                                                     | :heavy_minus_sign:                                            | URL for the metadata file with the IPFS protocol              |
| `NftMetadataGatewayURL`                                       | **string*                                                     | :heavy_minus_sign:                                            | URL to access metadata file via HTTP through an IPFS<br/>gateway<br/> |