// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetWebhookIDLogRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetWebhookIDLogRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetWebhookIDLogResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	WebhookLogs []components.WebhookLog
	// Error
	Error *sdkerrors.Error
}

func (o *GetWebhookIDLogResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWebhookIDLogResponse) GetWebhookLogs() []components.WebhookLog {
	if o == nil {
		return nil
	}
	return o.WebhookLogs
}

func (o *GetWebhookIDLogResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
