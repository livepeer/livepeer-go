// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type UpdateAssetRequest struct {
	// ID of the asset
	AssetID           string                       `pathParam:"style=simple,explode=false,name=assetId"`
	AssetPatchPayload components.AssetPatchPayload `request:"mediaType=application/json"`
}

func (o *UpdateAssetRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

func (o *UpdateAssetRequest) GetAssetPatchPayload() components.AssetPatchPayload {
	if o == nil {
		return components.AssetPatchPayload{}
	}
	return o.AssetPatchPayload
}

type UpdateAssetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Asset *components.Asset
	// Error
	Error *sdkerrors.Error
}

func (o *UpdateAssetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAssetResponse) GetAsset() *components.Asset {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *UpdateAssetResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
