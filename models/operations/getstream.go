// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetStreamRequest struct {
	// ID of the stream
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetStreamRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Stream *components.Stream
	// Error
	Error *sdkerrors.Error
}

func (o *GetStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetStreamResponse) GetStream() *components.Stream {
	if o == nil {
		return nil
	}
	return o.Stream
}

func (o *GetStreamResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
