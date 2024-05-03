package main

import (
	"context"
	livepeer "github.com/livepeer/livepeer-go"
	"github.com/livepeer/livepeer-go/models/components"
	"log"
)

func main() {
	s := livepeer.New(
		livepeer.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Stream.Create(ctx, components.NewStreamPayload{
		Name: "stream from go sdk",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		log.Printf("Stream created successfully")
	} 
}
