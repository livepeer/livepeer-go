// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type Task struct {
	ID string `json:"id"`
}

func (o *Task) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// RequestUploadData - Success
type RequestUploadData struct {
	// The direct upload endpoint for which supports PUT requests. **It is recommended to use the Tus endpoint for a better upload experience.**
	URL string `json:"url"`
	// The [Tus-compatible](https://tus.io/) endpoint for resumable uploads. **This is the recommended way to upload assets.** See the [Tus-js](https://github.com/tus/tus-js-client) client for more information.
	TusEndpoint string           `json:"tusEndpoint"`
	Asset       components.Asset `json:"asset"`
	Task        Task             `json:"task"`
}

func (o *RequestUploadData) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *RequestUploadData) GetTusEndpoint() string {
	if o == nil {
		return ""
	}
	return o.TusEndpoint
}

func (o *RequestUploadData) GetAsset() components.Asset {
	if o == nil {
		return components.Asset{}
	}
	return o.Asset
}

func (o *RequestUploadData) GetTask() Task {
	if o == nil {
		return Task{}
	}
	return o.Task
}

type RequestUploadResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Data *RequestUploadData
	// Error
	Error *sdkerrors.Error
}

func (o *RequestUploadResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RequestUploadResponse) GetData() *RequestUploadData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *RequestUploadResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
