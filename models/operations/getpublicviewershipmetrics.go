// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetPublicViewershipMetricsRequest struct {
	// The playback ID to filter the query results. This can be a canonical
	// playback ID from Livepeer assets or streams, or dStorage identifiers
	// for assets
	//
	PlaybackID string `pathParam:"style=simple,explode=false,name=playbackId"`
}

func (o *GetPublicViewershipMetricsRequest) GetPlaybackID() string {
	if o == nil {
		return ""
	}
	return o.PlaybackID
}

// GetPublicViewershipMetricsData - A simplified metric object about aggregate viewership of an
// asset. Either playbackId or dStorageUrl will be set.
type GetPublicViewershipMetricsData struct {
	// The URL of the distributed storage used for the asset
	DStorageURL *string `json:"dStorageUrl,omitempty"`
	// The playback ID associated with the metric.
	PlaybackID *string `json:"playbackId,omitempty"`
	// The total playtime in minutes for the asset.
	PlaytimeMins *float64 `json:"playtimeMins,omitempty"`
	// The number of views for the asset.
	ViewCount *int64 `json:"viewCount,omitempty"`
}

func (o *GetPublicViewershipMetricsData) GetDStorageURL() *string {
	if o == nil {
		return nil
	}
	return o.DStorageURL
}

func (o *GetPublicViewershipMetricsData) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

func (o *GetPublicViewershipMetricsData) GetPlaytimeMins() *float64 {
	if o == nil {
		return nil
	}
	return o.PlaytimeMins
}

func (o *GetPublicViewershipMetricsData) GetViewCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ViewCount
}

type GetPublicViewershipMetricsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A single Metric object with the viewCount and playtimeMins metrics.
	Data *GetPublicViewershipMetricsData
	// Error
	Error *sdkerrors.Error
}

func (o *GetPublicViewershipMetricsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPublicViewershipMetricsResponse) GetData() *GetPublicViewershipMetricsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetPublicViewershipMetricsResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
