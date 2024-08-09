// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Participants struct {
	// participant ID
	Identity *string `json:"identity,omitempty"`
	// user defined participant name
	Name *string `json:"name,omitempty"`
	// the time the participant joined
	JoinedAt *int64 `json:"joinedAt,omitempty"`
	// the time the participant left
	LeftAt *int64 `json:"leftAt,omitempty"`
}

func (o *Participants) GetIdentity() *string {
	if o == nil {
		return nil
	}
	return o.Identity
}

func (o *Participants) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Participants) GetJoinedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.JoinedAt
}

func (o *Participants) GetLeftAt() *int64 {
	if o == nil {
		return nil
	}
	return o.LeftAt
}

type Room struct {
	// room ID
	ID string `json:"id"`
	// Timestamp (in milliseconds) at which the room was created
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// Timestamp (in milliseconds) at which room was updated
	UpdatedAt *float64 `json:"updatedAt,omitempty"`
	// internal ID for egress output
	EgressID     *string                 `json:"egressId,omitempty"`
	Participants map[string]Participants `json:"participants"`
}

func (o *Room) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Room) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Room) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Room) GetEgressID() *string {
	if o == nil {
		return nil
	}
	return o.EgressID
}

func (o *Room) GetParticipants() map[string]Participants {
	if o == nil {
		return map[string]Participants{}
	}
	return o.Participants
}
