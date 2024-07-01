// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type NewStreamPayloadRecordingSpec struct {
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
}

func (o *NewStreamPayloadRecordingSpec) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

type NewStreamPayload struct {
	CreatorID   *InputCreatorID `json:"creatorId,omitempty"`
	Multistream *Multistream    `json:"multistream,omitempty"`
	Name        string          `json:"name"`
	// Whether the playback policy for an asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy `json:"playbackPolicy,omitempty"`
	Profiles       []FfmpegProfile `json:"profiles"`
	// Configuration for a stream that should be actively pulled from an
	// external source, rather than pushed to Livepeer. If specified, the
	// stream will not have a streamKey.
	Pull *Pull `json:"pull,omitempty"`
	// Should this stream be recorded? Uses default settings. For more
	// customization, create and configure an object store.
	//
	Record        *bool                          `json:"record,omitempty"`
	RecordingSpec *NewStreamPayloadRecordingSpec `json:"recordingSpec,omitempty"`
	// User input tags associated with the stream
	UserTags map[string]UserTags `json:"userTags,omitempty"`
}

func (o *NewStreamPayload) GetCreatorID() *InputCreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *NewStreamPayload) GetMultistream() *Multistream {
	if o == nil {
		return nil
	}
	return o.Multistream
}

func (o *NewStreamPayload) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NewStreamPayload) GetPlaybackPolicy() *PlaybackPolicy {
	if o == nil {
		return nil
	}
	return o.PlaybackPolicy
}

func (o *NewStreamPayload) GetProfiles() []FfmpegProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *NewStreamPayload) GetPull() *Pull {
	if o == nil {
		return nil
	}
	return o.Pull
}

func (o *NewStreamPayload) GetRecord() *bool {
	if o == nil {
		return nil
	}
	return o.Record
}

func (o *NewStreamPayload) GetRecordingSpec() *NewStreamPayloadRecordingSpec {
	if o == nil {
		return nil
	}
	return o.RecordingSpec
}

func (o *NewStreamPayload) GetUserTags() map[string]UserTags {
	if o == nil {
		return nil
	}
	return o.UserTags
}
