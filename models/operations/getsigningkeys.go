// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetSigningKeysResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Data []components.SigningKey
	// Error
	Error *sdkerrors.Error
}

func (o *GetSigningKeysResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSigningKeysResponse) GetData() []components.SigningKey {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSigningKeysResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
