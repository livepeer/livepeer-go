<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"livepeer"
	"livepeer/models/components"
	"log"
)

func main() {
	s := livepeer.New(
		livepeer.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->