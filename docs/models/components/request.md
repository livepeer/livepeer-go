# Request


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `Body`                              | **string*                           | :heavy_minus_sign:                  | request body                        | {"event": "stream.started"}         |
| `Headers`                           | map[string]*string*                 | :heavy_minus_sign:                  | HTTP request headers                | {<br/>"User-Agent": "livepeer.studio"<br/>} |
| `Method`                            | **string*                           | :heavy_minus_sign:                  | HTTP request method                 | POST                                |
| `URL`                               | **string*                           | :heavy_minus_sign:                  | URL used for the request            | https://my-service.com/webhook      |