<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	livepeergo "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostStream(ctx, components.NewStreamPayload{
		Name:   "test_stream",
		Record: livepeergo.Bool(false),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->