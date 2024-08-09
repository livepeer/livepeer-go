// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AssetPatchPayload struct {
	// The name of the asset. This is not necessarily the filename - it can be a custom name or title.
	//
	Name      *string         `json:"name,omitempty"`
	CreatorID *InputCreatorID `json:"creatorId,omitempty"`
	// Whether the playback policy for an asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy `json:"playbackPolicy,omitempty"`
	Storage        *Storage        `json:"storage,omitempty"`
}

func (o *AssetPatchPayload) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AssetPatchPayload) GetCreatorID() *InputCreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *AssetPatchPayload) GetPlaybackPolicy() *PlaybackPolicy {
	if o == nil {
		return nil
	}
	return o.PlaybackPolicy
}

func (o *AssetPatchPayload) GetStorage() *Storage {
	if o == nil {
		return nil
	}
	return o.Storage
}
