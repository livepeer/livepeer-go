// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"livepeer/internal/utils"
)

type TargetOutput struct {
	// Name of transcoding profile that should be sent. Use
	// "source" for pushing source stream data
	//
	Profile string `json:"profile"`
	// If true, the stream audio will be muted and only silent
	// video will be pushed to the target.
	//
	VideoOnly *bool `default:"false" json:"videoOnly"`
	// ID of multistream target object where to push this stream
	ID *string `json:"id,omitempty"`
}

func (t TargetOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TargetOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TargetOutput) GetProfile() string {
	if o == nil {
		return ""
	}
	return o.Profile
}

func (o *TargetOutput) GetVideoOnly() *bool {
	if o == nil {
		return nil
	}
	return o.VideoOnly
}

func (o *TargetOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
