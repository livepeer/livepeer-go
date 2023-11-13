# TranscodeFile

Parameters for the transcode-file task


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Input`                                                                 | [*components.Input](../../models/components/input.md)                   | :heavy_minus_sign:                                                      | Input video file to transcode                                           |
| `Storage`                                                               | [*components.TaskStorage](../../models/components/taskstorage.md)       | :heavy_minus_sign:                                                      | Storage for the output files                                            |
| `Outputs`                                                               | [*components.Outputs](../../models/components/outputs.md)               | :heavy_minus_sign:                                                      | Output formats                                                          |
| `Profiles`                                                              | [][components.FfmpegProfile](../../models/components/ffmpegprofile.md)  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `TargetSegmentSizeSecs`                                                 | **float64*                                                              | :heavy_minus_sign:                                                      | How many seconds the duration of each output segment should<br/>be<br/> |
| `CreatorID`                                                             | [*components.InputCreatorID](../../models/components/inputcreatorid.md) | :heavy_minus_sign:                                                      | N/A                                                                     |