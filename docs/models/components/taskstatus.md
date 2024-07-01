# TaskStatus

Status of the task


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ErrorMessage`                                               | **string*                                                    | :heavy_minus_sign:                                           | Error message if the task failed                             | Failed to upload file                                        |
| `Phase`                                                      | [components.TaskPhase](../../models/components/taskphase.md) | :heavy_check_mark:                                           | Phase of the task                                            | pending                                                      |
| `Progress`                                                   | **float64*                                                   | :heavy_minus_sign:                                           | Current progress of the task in a 0-1 ratio                  | 0.5                                                          |
| `Retries`                                                    | **float64*                                                   | :heavy_minus_sign:                                           | Number of retries done on the task                           | 3                                                            |
| `UpdatedAt`                                                  | *float64*                                                    | :heavy_check_mark:                                           | Timestamp (in milliseconds) at which task was updated        | 1587667174725                                                |