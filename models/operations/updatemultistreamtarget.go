// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"livepeer/models/components"
	"net/http"
)

type UpdateMultistreamTargetRequest struct {
	// ID of the multistream target
	ID                            string                                   `pathParam:"style=simple,explode=false,name=id"`
	MultistreamTargetPatchPayload components.MultistreamTargetPatchPayload `request:"mediaType=application/json"`
}

func (o *UpdateMultistreamTargetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateMultistreamTargetRequest) GetMultistreamTargetPatchPayload() components.MultistreamTargetPatchPayload {
	if o == nil {
		return components.MultistreamTargetPatchPayload{}
	}
	return o.MultistreamTargetPatchPayload
}

type UpdateMultistreamTargetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateMultistreamTargetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMultistreamTargetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMultistreamTargetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
