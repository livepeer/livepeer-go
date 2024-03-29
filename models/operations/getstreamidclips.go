// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetStreamIDClipsRequest struct {
	// ID of the parent stream or playbackId of parent stream
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetStreamIDClipsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetStreamIDClipsResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	Assets []components.Asset
	// Error
	Error *sdkerrors.Error
}

func (o *GetStreamIDClipsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetStreamIDClipsResponse) GetAssets() []components.Asset {
	if o == nil {
		return nil
	}
	return o.Assets
}

func (o *GetStreamIDClipsResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
