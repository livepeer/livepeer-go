// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"livepeer/models/components"
	"net/http"
)

type Task struct {
	ID *string `json:"id,omitempty"`
}

func (o *Task) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestUploadResponseBody - Success
type RequestUploadResponseBody struct {
	URL         string           `json:"url"`
	TusEndpoint string           `json:"tusEndpoint"`
	Asset       components.Asset `json:"asset"`
	Task        Task             `json:"task"`
}

func (o *RequestUploadResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *RequestUploadResponseBody) GetTusEndpoint() string {
	if o == nil {
		return ""
	}
	return o.TusEndpoint
}

func (o *RequestUploadResponseBody) GetAsset() components.Asset {
	if o == nil {
		return components.Asset{}
	}
	return o.Asset
}

func (o *RequestUploadResponseBody) GetTask() Task {
	if o == nil {
		return Task{}
	}
	return o.Task
}

type RequestUploadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Object *RequestUploadResponseBody
}

func (o *RequestUploadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RequestUploadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RequestUploadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RequestUploadResponse) GetObject() *RequestUploadResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
