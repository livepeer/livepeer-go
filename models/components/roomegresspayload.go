// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RoomEgressPayload struct {
	// The ID of the Livepeer Stream to stream to
	StreamID string `json:"streamId"`
}

func (o *RoomEgressPayload) GetStreamID() string {
	if o == nil {
		return ""
	}
	return o.StreamID
}
