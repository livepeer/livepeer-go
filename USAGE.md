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
	ctx := context.Background()

	s := livepeergo.New(
		livepeergo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Stream.Create(ctx, components.NewStreamPayload{
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
			WebhookID: livepeergo.Pointer("1bde4o2i6xycudoy"),
			WebhookContext: map[string]any{
				"streamerId": "my-custom-id",
			},
			RefreshInterval: livepeergo.Pointer[float64](600),
		},
		Profiles: []components.FfmpegProfile{},
		Record:   livepeergo.Pointer(false),
		RecordingSpec: &components.NewStreamPayloadRecordingSpec{
			Profiles: []components.TranscodeProfile{
				components.TranscodeProfile{
					Width:   livepeergo.Pointer[int64](1280),
					Name:    livepeergo.Pointer("720p"),
					Height:  livepeergo.Pointer[int64](720),
					Bitrate: 3000000,
					Quality: livepeergo.Pointer[int64](23),
					Fps:     livepeergo.Pointer[int64](30),
					FpsDen:  livepeergo.Pointer[int64](1),
					Gop:     livepeergo.Pointer("2"),
					Profile: components.TranscodeProfileProfileH264Baseline.ToPointer(),
					Encoder: components.TranscodeProfileEncoderH264.ToPointer(),
				},
			},
		},
		Multistream: &components.Multistream{
			Targets: []components.Target{
				components.Target{
					Profile: "720p",
					ID:      livepeergo.Pointer("PUSH123"),
				},
			},
		},
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