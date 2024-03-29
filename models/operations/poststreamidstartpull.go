// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type PostStreamIDStartPullRequest struct {
	// ID of the stream
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PostStreamIDStartPullRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PostStreamIDStartPullResponse struct {
	HTTPMeta components.HTTPMetadata
	// Error
	Error *sdkerrors.Error
}

func (o *PostStreamIDStartPullResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostStreamIDStartPullResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
