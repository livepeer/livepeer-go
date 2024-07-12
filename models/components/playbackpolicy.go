// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypePublic  Type = "public"
	TypeJwt     Type = "jwt"
	TypeWebhook Type = "webhook"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "jwt":
		fallthrough
	case "webhook":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// PlaybackPolicy - Whether the playback policy for an asset or stream is public or signed
type PlaybackPolicy struct {
	Type Type `json:"type"`
	// ID of the webhook to use for playback policy
	WebhookID *string `json:"webhookId,omitempty"`
	// User-defined webhook context
	WebhookContext map[string]any `json:"webhookContext,omitempty"`
	// Interval (in seconds) at which the playback policy should be
	// refreshed (default 600 seconds)
	//
	RefreshInterval *float64 `json:"refreshInterval,omitempty"`
	// List of allowed origins for CORS playback (<scheme>://<hostname>:<port>, <scheme>://<hostname>)
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

func (o *PlaybackPolicy) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *PlaybackPolicy) GetWebhookID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

func (o *PlaybackPolicy) GetWebhookContext() map[string]any {
	if o == nil {
		return nil
	}
	return o.WebhookContext
}

func (o *PlaybackPolicy) GetRefreshInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshInterval
}

func (o *PlaybackPolicy) GetAllowedOrigins() []string {
	if o == nil {
		return nil
	}
	return o.AllowedOrigins
}
