# MasksResponse

Response model for object segmentation.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Masks`                                                   | *string*                                                  | :heavy_check_mark:                                        | The generated masks.                                      |
| `Scores`                                                  | *string*                                                  | :heavy_check_mark:                                        | The model's confidence scores for each generated mask.    |
| `Logits`                                                  | *string*                                                  | :heavy_check_mark:                                        | The raw, unnormalized predictions (logits) for the masks. |