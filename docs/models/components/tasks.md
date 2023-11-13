# Tasks


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Pending`                                                                        | **string*                                                                        | :heavy_minus_sign:                                                               | ID of any currently running task that is exporting this<br/>asset to IPFS.<br/>  |
| `Last`                                                                           | **string*                                                                        | :heavy_minus_sign:                                                               | ID of the last task to run successfully, that created<br/>the currently saved data.<br/> |
| `Failed`                                                                         | **string*                                                                        | :heavy_minus_sign:                                                               | ID of the last task to fail execution.                                           |