// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type CreateStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Stream *components.Stream
	// Error
	Error *sdkerrors.Error
}

func (o *CreateStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateStreamResponse) GetStream() *components.Stream {
	if o == nil {
		return nil
	}
	return o.Stream
}

func (o *CreateStreamResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
