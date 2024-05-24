// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetRoomUserRequest struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *GetRoomUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRoomUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type GetRoomUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	GetRoomUserResponse *components.GetRoomUserResponse
	// Error
	Error *sdkerrors.Error
}

func (o *GetRoomUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRoomUserResponse) GetGetRoomUserResponse() *components.GetRoomUserResponse {
	if o == nil {
		return nil
	}
	return o.GetRoomUserResponse
}

func (o *GetRoomUserResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
