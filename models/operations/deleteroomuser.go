// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type DeleteRoomUserRequest struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *DeleteRoomUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteRoomUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type DeleteRoomUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Error
	Error *sdkerrors.Error
}

func (o *DeleteRoomUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteRoomUserResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
