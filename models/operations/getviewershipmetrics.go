// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
	"time"
)

type FromType string

const (
	FromTypeDateTime FromType = "date-time"
	FromTypeInteger  FromType = "integer"
)

// From - Start timestamp for the query range (inclusive)
type From struct {
	DateTime *time.Time
	Integer  *int64

	Type FromType
}

func CreateFromDateTime(dateTime time.Time) From {
	typ := FromTypeDateTime

	return From{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateFromInteger(integer int64) From {
	typ := FromTypeInteger

	return From{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *From) UnmarshalJSON(data []byte) error {

	var dateTime time.Time = time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = FromTypeDateTime
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = FromTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u From) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ToType string

const (
	ToTypeDateTime ToType = "date-time"
	ToTypeInteger  ToType = "integer"
)

// To - End timestamp for the query range (exclusive)
type To struct {
	DateTime *time.Time
	Integer  *int64

	Type ToType
}

func CreateToDateTime(dateTime time.Time) To {
	typ := ToTypeDateTime

	return To{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateToInteger(integer int64) To {
	typ := ToTypeInteger

	return To{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *To) UnmarshalJSON(data []byte) error {

	var dateTime time.Time = time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = ToTypeDateTime
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = ToTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u To) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// TimeStep - The time step to aggregate viewership metrics by
type TimeStep string

const (
	TimeStepHour  TimeStep = "hour"
	TimeStepDay   TimeStep = "day"
	TimeStepWeek  TimeStep = "week"
	TimeStepMonth TimeStep = "month"
	TimeStepYear  TimeStep = "year"
)

func (e TimeStep) ToPointer() *TimeStep {
	return &e
}

func (e *TimeStep) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hour":
		fallthrough
	case "day":
		fallthrough
	case "week":
		fallthrough
	case "month":
		fallthrough
	case "year":
		*e = TimeStep(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TimeStep: %v", v)
	}
}

type BreakdownBy string

const (
	BreakdownByPlaybackID    BreakdownBy = "playbackId"
	BreakdownByDeviceType    BreakdownBy = "deviceType"
	BreakdownByDevice        BreakdownBy = "device"
	BreakdownByCPU           BreakdownBy = "cpu"
	BreakdownByOs            BreakdownBy = "os"
	BreakdownByBrowser       BreakdownBy = "browser"
	BreakdownByBrowserEngine BreakdownBy = "browserEngine"
	BreakdownByContinent     BreakdownBy = "continent"
	BreakdownByCountry       BreakdownBy = "country"
	BreakdownBySubdivision   BreakdownBy = "subdivision"
	BreakdownByTimezone      BreakdownBy = "timezone"
	BreakdownByGeohash       BreakdownBy = "geohash"
	BreakdownByViewerID      BreakdownBy = "viewerId"
	BreakdownByCreatorID     BreakdownBy = "creatorId"
)

func (e BreakdownBy) ToPointer() *BreakdownBy {
	return &e
}

func (e *BreakdownBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "playbackId":
		fallthrough
	case "deviceType":
		fallthrough
	case "device":
		fallthrough
	case "cpu":
		fallthrough
	case "os":
		fallthrough
	case "browser":
		fallthrough
	case "browserEngine":
		fallthrough
	case "continent":
		fallthrough
	case "country":
		fallthrough
	case "subdivision":
		fallthrough
	case "timezone":
		fallthrough
	case "geohash":
		fallthrough
	case "viewerId":
		fallthrough
	case "creatorId":
		*e = BreakdownBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BreakdownBy: %v", v)
	}
}

type GetViewershipMetricsRequest struct {
	// The playback ID to filter the query results. This can be a canonical
	// playback ID from Livepeer assets or streams, or dStorage identifiers
	// for assets
	//
	PlaybackID *string `queryParam:"style=form,explode=true,name=playbackId"`
	// Start timestamp for the query range (inclusive)
	From *From `queryParam:"style=form,explode=true,name=from"`
	// End timestamp for the query range (exclusive)
	To *To `queryParam:"style=form,explode=true,name=to"`
	// The time step to aggregate viewership metrics by
	TimeStep *TimeStep `queryParam:"style=form,explode=true,name=timeStep"`
	// The asset ID to filter metrics for
	AssetID *string `queryParam:"style=form,explode=true,name=assetId"`
	// The stream ID to filter metrics for
	StreamID *string `queryParam:"style=form,explode=true,name=streamId"`
	// The creator ID to filter the query results
	CreatorID *string `queryParam:"style=form,explode=true,name=creatorId"`
	// The list of fields to break down the query results. Specify this
	// query-string multiple times to break down by multiple fields.
	//
	BreakdownBy []BreakdownBy `queryParam:"style=form,explode=true,name=breakdownBy[]"`
}

func (o *GetViewershipMetricsRequest) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

func (o *GetViewershipMetricsRequest) GetFrom() *From {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetViewershipMetricsRequest) GetTo() *To {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *GetViewershipMetricsRequest) GetTimeStep() *TimeStep {
	if o == nil {
		return nil
	}
	return o.TimeStep
}

func (o *GetViewershipMetricsRequest) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *GetViewershipMetricsRequest) GetStreamID() *string {
	if o == nil {
		return nil
	}
	return o.StreamID
}

func (o *GetViewershipMetricsRequest) GetCreatorID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *GetViewershipMetricsRequest) GetBreakdownBy() []BreakdownBy {
	if o == nil {
		return nil
	}
	return o.BreakdownBy
}

type GetViewershipMetricsResponse struct {
	HTTPMeta components.HTTPMetadata
	// A list of Metric objects
	Data []components.ViewershipMetric
	// Error
	Error *sdkerrors.Error
}

func (o *GetViewershipMetricsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetViewershipMetricsResponse) GetData() []components.ViewershipMetric {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetViewershipMetricsResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
