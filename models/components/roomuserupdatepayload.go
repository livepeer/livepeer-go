// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/livepeer/livepeer-go/internal/utils"
)

type RoomUserUpdatePayload struct {
	// Whether a user is allowed to publish audio/video tracks (i.e. their microphone and webcam)
	CanPublish *bool `default:"true" json:"canPublish"`
	// Whether a user is allowed to publish data messages to the room
	CanPublishData *bool `default:"true" json:"canPublishData"`
	// User defined payload to store for the participant
	Metadata *string `json:"metadata,omitempty"`
}

func (r RoomUserUpdatePayload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RoomUserUpdatePayload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RoomUserUpdatePayload) GetCanPublish() *bool {
	if o == nil {
		return nil
	}
	return o.CanPublish
}

func (o *RoomUserUpdatePayload) GetCanPublishData() *bool {
	if o == nil {
		return nil
	}
	return o.CanPublishData
}

func (o *RoomUserUpdatePayload) GetMetadata() *string {
	if o == nil {
		return nil
	}
	return o.Metadata
}
