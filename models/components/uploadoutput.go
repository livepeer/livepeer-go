// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"livepeer/internal/utils"
)

// UploadOutput - Output of the upload task
type UploadOutput struct {
	AssetSpec            *Asset                 `json:"assetSpec,omitempty"`
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
}

func (u UploadOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UploadOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UploadOutput) GetAssetSpec() *Asset {
	if o == nil {
		return nil
	}
	return o.AssetSpec
}

func (o *UploadOutput) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
