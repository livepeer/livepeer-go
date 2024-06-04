# GetRecordedSessionsRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ParentID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | ID of the parent stream                                                |                                                                        |
| `Record`                                                               | [*operations.Record](../../models/operations/record.md)                | :heavy_minus_sign:                                                     | Flag indicating if the response should only include recorded<br/>sessions<br/> | true                                                                   |