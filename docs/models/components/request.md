# Request


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `URL`                               | **string*                           | :heavy_minus_sign:                  | URL used for the request            | https://my-service.com/webhook      |
| `Method`                            | **string*                           | :heavy_minus_sign:                  | HTTP request method                 | POST                                |
| `Headers`                           | map[string]*string*                 | :heavy_minus_sign:                  | HTTP request headers                | {<br/>"User-Agent": "livepeer.studio"<br/>} |
| `Body`                              | **string*                           | :heavy_minus_sign:                  | request body                        | {"event": "stream.started"}         |