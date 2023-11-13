# TaskStatus

Status of the task


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Phase`                                                      | [components.TaskPhase](../../models/components/taskphase.md) | :heavy_check_mark:                                           | Phase of the task                                            |                                                              |
| `UpdatedAt`                                                  | *float64*                                                    | :heavy_check_mark:                                           | Timestamp (in milliseconds) at which task was updated        | 1587667174725                                                |
| `Progress`                                                   | **float64*                                                   | :heavy_minus_sign:                                           | Current progress of the task in a 0-1 ratio                  |                                                              |
| `ErrorMessage`                                               | **string*                                                    | :heavy_minus_sign:                                           | Error message if the task failed                             |                                                              |
| `Retries`                                                    | **float64*                                                   | :heavy_minus_sign:                                           | Number of retries done on the task                           |                                                              |