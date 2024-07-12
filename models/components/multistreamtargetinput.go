// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type MultistreamTargetInput struct {
	Name *string `json:"name,omitempty"`
	// Livepeer-compatible multistream target URL (RTMP(S) or SRT)
	URL string `json:"url"`
	// If true then this multistream target will not be used for pushing
	// even if it is configured in a stream object.
	//
	Disabled *bool `json:"disabled,omitempty"`
}

func (o *MultistreamTargetInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MultistreamTargetInput) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *MultistreamTargetInput) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}
