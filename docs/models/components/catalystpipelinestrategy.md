# CatalystPipelineStrategy

Force to use a specific strategy in the Catalyst pipeline. If not specified, the default strategy that Catalyst is configured for will be used. This field only available for admin users, and is only used for E2E testing.


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CatalystPipelineStrategyCatalyst`           | catalyst                                     |
| `CatalystPipelineStrategyCatalystFfmpeg`     | catalyst_ffmpeg                              |
| `CatalystPipelineStrategyBackgroundExternal` | background_external                          |
| `CatalystPipelineStrategyBackgroundMist`     | background_mist                              |
| `CatalystPipelineStrategyFallbackExternal`   | fallback_external                            |
| `CatalystPipelineStrategyExternal`           | external                                     |