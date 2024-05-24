// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type DeleteWebhookRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteWebhookRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteWebhookResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Webhook *components.Webhook
	// Error
	Error *sdkerrors.Error
}

func (o *DeleteWebhookResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteWebhookResponse) GetWebhook() *components.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

func (o *DeleteWebhookResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
