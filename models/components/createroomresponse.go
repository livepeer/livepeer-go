// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CreateRoomResponse struct {
	// The ID of the room
	ID *string `json:"id,omitempty"`
}

func (o *CreateRoomResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
