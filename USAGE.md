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
	request := components.NewStreamPayload{
		Name: "test_stream",
		Pull: &components.Pull{
			Source: "https://myservice.com/live/stream.flv",
			Headers: map[string]string{
				"Authorization": "Bearer 123",
			},
			Location: &components.Location{
				Lat: 39.739,
				Lon: -104.988,
			},
		},
		PlaybackPolicy: &components.PlaybackPolicy{
			Type:      components.TypeWebhook,
			WebhookID: livepeergo.String("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Float64(600),
		},
		Profiles: []components.FfmpegProfile{
			components.FfmpegProfile{
				Width:   1280,
				Name:    "720p",
				Height:  486589,
				Bitrate: 3000000,
				Fps:     30,
				FpsDen:  livepeergo.Int64(1),
				Quality: livepeergo.Int64(23),
				Gop:     livepeergo.String("2"),
				Profile: components.ProfileH264Baseline.ToPointer(),
			},
		},
		Record: livepeergo.Bool(false),
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile:   "720p",
					VideoOnly: livepeergo.Bool(false),
					ID:        livepeergo.String("PUSH123"),
					Spec: &components.TargetSpec{
						Name: livepeergo.String("My target"),
						URL:  "rtmps://live.my-service.tv/channel/secretKey",
					},
				},
			},
		},
	}
	ctx := context.Background()
	res, err := s.Stream.Create(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->