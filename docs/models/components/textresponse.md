# TextResponse

Response model for text generation.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Text`                                                 | *string*                                               | :heavy_check_mark:                                     | The generated text.                                    |
| `Chunks`                                               | [][components.Chunk](../../models/components/chunk.md) | :heavy_check_mark:                                     | The generated text chunks.                             |