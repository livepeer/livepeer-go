// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PlaybackInfoType string

const (
	PlaybackInfoTypeLive      PlaybackInfoType = "live"
	PlaybackInfoTypeVod       PlaybackInfoType = "vod"
	PlaybackInfoTypeRecording PlaybackInfoType = "recording"
)

func (e PlaybackInfoType) ToPointer() *PlaybackInfoType {
	return &e
}
func (e *PlaybackInfoType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "live":
		fallthrough
	case "vod":
		fallthrough
	case "recording":
		*e = PlaybackInfoType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaybackInfoType: %v", v)
	}
}

type Live int64

const (
	LiveZero Live = 0
	LiveOne  Live = 1
)

func (e Live) ToPointer() *Live {
	return &e
}
func (e *Live) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = Live(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Live: %v", v)
	}
}

// Hrn - Human Readable Name
type Hrn string

const (
	HrnHlsTs         Hrn = "HLS (TS)"
	HrnMp4           Hrn = "MP4"
	HrnWebRtcH264    Hrn = "WebRTC (H264)"
	HrnFlvH264       Hrn = "FLV (H264)"
	HrnThumbnailJpeg Hrn = "Thumbnail (JPEG)"
	HrnThumbnails    Hrn = "Thumbnails"
)

func (e Hrn) ToPointer() *Hrn {
	return &e
}
func (e *Hrn) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HLS (TS)":
		fallthrough
	case "MP4":
		fallthrough
	case "WebRTC (H264)":
		fallthrough
	case "FLV (H264)":
		fallthrough
	case "Thumbnail (JPEG)":
		fallthrough
	case "Thumbnails":
		*e = Hrn(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hrn: %v", v)
	}
}

type PlaybackInfoMetaType string

const (
	PlaybackInfoMetaTypeHtml5ApplicationVndAppleMpegurl PlaybackInfoMetaType = "html5/application/vnd.apple.mpegurl"
	PlaybackInfoMetaTypeHtml5VideoMp4                   PlaybackInfoMetaType = "html5/video/mp4"
	PlaybackInfoMetaTypeHtml5VideoH264                  PlaybackInfoMetaType = "html5/video/h264"
	PlaybackInfoMetaTypeVideoXFlv                       PlaybackInfoMetaType = "video/x-flv"
	PlaybackInfoMetaTypeImageJpeg                       PlaybackInfoMetaType = "image/jpeg"
	PlaybackInfoMetaTypeTextVtt                         PlaybackInfoMetaType = "text/vtt"
)

func (e PlaybackInfoMetaType) ToPointer() *PlaybackInfoMetaType {
	return &e
}
func (e *PlaybackInfoMetaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "html5/application/vnd.apple.mpegurl":
		fallthrough
	case "html5/video/mp4":
		fallthrough
	case "html5/video/h264":
		fallthrough
	case "video/x-flv":
		fallthrough
	case "image/jpeg":
		fallthrough
	case "text/vtt":
		*e = PlaybackInfoMetaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaybackInfoMetaType: %v", v)
	}
}

type PlaybackInfoSource struct {
	// Human Readable Name
	Hrn     Hrn                  `json:"hrn"`
	Type    PlaybackInfoMetaType `json:"type"`
	URL     string               `json:"url"`
	Size    *float64             `json:"size,omitempty"`
	Width   *float64             `json:"width,omitempty"`
	Height  *float64             `json:"height,omitempty"`
	Bitrate *float64             `json:"bitrate,omitempty"`
}

func (o *PlaybackInfoSource) GetHrn() Hrn {
	if o == nil {
		return Hrn("")
	}
	return o.Hrn
}

func (o *PlaybackInfoSource) GetType() PlaybackInfoMetaType {
	if o == nil {
		return PlaybackInfoMetaType("")
	}
	return o.Type
}

func (o *PlaybackInfoSource) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *PlaybackInfoSource) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *PlaybackInfoSource) GetWidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *PlaybackInfoSource) GetHeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *PlaybackInfoSource) GetBitrate() *float64 {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

type PlaybackInfoHrn string

const (
	PlaybackInfoHrnHlsTs PlaybackInfoHrn = "HLS (TS)"
)

func (e PlaybackInfoHrn) ToPointer() *PlaybackInfoHrn {
	return &e
}
func (e *PlaybackInfoHrn) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HLS (TS)":
		*e = PlaybackInfoHrn(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaybackInfoHrn: %v", v)
	}
}

type PlaybackInfoMetaDvrPlaybackType string

const (
	PlaybackInfoMetaDvrPlaybackTypeHtml5ApplicationVndAppleMpegurl PlaybackInfoMetaDvrPlaybackType = "html5/application/vnd.apple.mpegurl"
)

func (e PlaybackInfoMetaDvrPlaybackType) ToPointer() *PlaybackInfoMetaDvrPlaybackType {
	return &e
}
func (e *PlaybackInfoMetaDvrPlaybackType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "html5/application/vnd.apple.mpegurl":
		*e = PlaybackInfoMetaDvrPlaybackType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaybackInfoMetaDvrPlaybackType: %v", v)
	}
}

type DvrPlayback struct {
	Hrn   *PlaybackInfoHrn                 `json:"hrn,omitempty"`
	Type  *PlaybackInfoMetaDvrPlaybackType `json:"type,omitempty"`
	URL   *string                          `json:"url,omitempty"`
	Error *string                          `json:"error,omitempty"`
}

func (o *DvrPlayback) GetHrn() *PlaybackInfoHrn {
	if o == nil {
		return nil
	}
	return o.Hrn
}

func (o *DvrPlayback) GetType() *PlaybackInfoMetaDvrPlaybackType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *DvrPlayback) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *DvrPlayback) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

type Meta struct {
	Live *Live `json:"live,omitempty"`
	// Whether the playback policy for an asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy      `json:"playbackPolicy,omitempty"`
	Source         []PlaybackInfoSource `json:"source"`
	DvrPlayback    []DvrPlayback        `json:"dvrPlayback,omitempty"`
	Attestation    *Attestation         `json:"attestation,omitempty"`
}

func (o *Meta) GetLive() *Live {
	if o == nil {
		return nil
	}
	return o.Live
}

func (o *Meta) GetPlaybackPolicy() *PlaybackPolicy {
	if o == nil {
		return nil
	}
	return o.PlaybackPolicy
}

func (o *Meta) GetSource() []PlaybackInfoSource {
	if o == nil {
		return []PlaybackInfoSource{}
	}
	return o.Source
}

func (o *Meta) GetDvrPlayback() []DvrPlayback {
	if o == nil {
		return nil
	}
	return o.DvrPlayback
}

func (o *Meta) GetAttestation() *Attestation {
	if o == nil {
		return nil
	}
	return o.Attestation
}

type PlaybackInfo struct {
	Type PlaybackInfoType `json:"type"`
	Meta Meta             `json:"meta"`
}

func (o *PlaybackInfo) GetType() PlaybackInfoType {
	if o == nil {
		return PlaybackInfoType("")
	}
	return o.Type
}

func (o *PlaybackInfo) GetMeta() Meta {
	if o == nil {
		return Meta{}
	}
	return o.Meta
}
