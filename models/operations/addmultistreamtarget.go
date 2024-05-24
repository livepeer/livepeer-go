// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type AddMultistreamTargetRequest struct {
	// ID of the parent stream
	ID               string                      `pathParam:"style=simple,explode=false,name=id"`
	TargetAddPayload components.TargetAddPayload `request:"mediaType=application/json"`
}

func (o *AddMultistreamTargetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AddMultistreamTargetRequest) GetTargetAddPayload() components.TargetAddPayload {
	if o == nil {
		return components.TargetAddPayload{}
	}
	return o.TargetAddPayload
}

type AddMultistreamTargetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Error
	Error *sdkerrors.Error
}

func (o *AddMultistreamTargetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AddMultistreamTargetResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
