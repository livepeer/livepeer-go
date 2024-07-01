<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.AccessControl.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Data != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->