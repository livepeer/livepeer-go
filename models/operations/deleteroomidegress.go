// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type DeleteRoomIDEgressRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteRoomIDEgressRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteRoomIDEgressResponse struct {
	HTTPMeta components.HTTPMetadata
	// Error
	Error *sdkerrors.Error
}

func (o *DeleteRoomIDEgressResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteRoomIDEgressResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
