// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type DeleteAssetRequest struct {
	// ID of the asset
	AssetID string `pathParam:"style=simple,explode=false,name=assetId"`
}

func (o *DeleteAssetRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

type DeleteAssetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Error
	Error *sdkerrors.Error
}

func (o *DeleteAssetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteAssetResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
