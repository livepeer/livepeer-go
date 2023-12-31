// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"livepeer/models/components"
	"net/http"
)

type PostClipTask struct {
	ID *string `json:"id,omitempty"`
}

func (o *PostClipTask) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PostClipData - Success
type PostClipData struct {
	Asset components.Asset `json:"asset"`
	Task  PostClipTask     `json:"task"`
}

func (o *PostClipData) GetAsset() components.Asset {
	if o == nil {
		return components.Asset{}
	}
	return o.Asset
}

func (o *PostClipData) GetTask() PostClipTask {
	if o == nil {
		return PostClipTask{}
	}
	return o.Task
}

type PostClipResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Data *PostClipData
}

func (o *PostClipResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostClipResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostClipResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostClipResponse) GetData() *PostClipData {
	if o == nil {
		return nil
	}
	return o.Data
}
