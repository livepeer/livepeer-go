// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type PostAssetUploadURLTask struct {
	ID *string `json:"id,omitempty"`
}

func (o *PostAssetUploadURLTask) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PostAssetUploadURLResponseBody - Success
type PostAssetUploadURLResponseBody struct {
	Asset components.Asset       `json:"asset"`
	Task  PostAssetUploadURLTask `json:"task"`
}

func (o *PostAssetUploadURLResponseBody) GetAsset() components.Asset {
	if o == nil {
		return components.Asset{}
	}
	return o.Asset
}

func (o *PostAssetUploadURLResponseBody) GetTask() PostAssetUploadURLTask {
	if o == nil {
		return PostAssetUploadURLTask{}
	}
	return o.Task
}

type PostAssetUploadURLResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	Object *PostAssetUploadURLResponseBody
	// Error
	Error *sdkerrors.Error
}

func (o *PostAssetUploadURLResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostAssetUploadURLResponse) GetObject() *PostAssetUploadURLResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *PostAssetUploadURLResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
