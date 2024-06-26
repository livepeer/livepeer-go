// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/livepeer/livepeer-go/internal/utils"
)

type GetRoomUserResponse struct {
	// The ID of the user
	ID *string `json:"id,omitempty"`
	// Whether a user is allowed to publish audio/video tracks
	IsPublisher *bool `default:"true" json:"isPublisher"`
	// Timestamp (in milliseconds) at which the user joined
	JoinedAt *int64 `json:"joinedAt,omitempty"`
	// User defined payload to store for the participant
	Metadata *string `json:"metadata,omitempty"`
	// The display name of the user
	Name *string `json:"name,omitempty"`
}

func (g GetRoomUserResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRoomUserResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRoomUserResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetRoomUserResponse) GetIsPublisher() *bool {
	if o == nil {
		return nil
	}
	return o.IsPublisher
}

func (o *GetRoomUserResponse) GetJoinedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.JoinedAt
}

func (o *GetRoomUserResponse) GetMetadata() *string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *GetRoomUserResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
