// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type RecordingSpec struct {
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
}

func (o *RecordingSpec) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}
