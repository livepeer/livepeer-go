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

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for From", string(data))
}

func (u From) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type From: all fields are null")
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

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for To", string(data))
}

func (u To) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type To: all fields are null")
}

// QueryParamTimeStep - The time step to aggregate viewership metrics by
type QueryParamTimeStep string

const (
	QueryParamTimeStepHour  QueryParamTimeStep = "hour"
	QueryParamTimeStepDay   QueryParamTimeStep = "day"
	QueryParamTimeStepWeek  QueryParamTimeStep = "week"
	QueryParamTimeStepMonth QueryParamTimeStep = "month"
	QueryParamTimeStepYear  QueryParamTimeStep = "year"
)

func (e QueryParamTimeStep) ToPointer() *QueryParamTimeStep {
	return &e
}
func (e *QueryParamTimeStep) UnmarshalJSON(data []byte) error {
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
		*e = QueryParamTimeStep(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamTimeStep: %v", v)
	}
}

type GetViewershipMetricsQueryParamBreakdownBy string

const (
	GetViewershipMetricsQueryParamBreakdownByPlaybackID    GetViewershipMetricsQueryParamBreakdownBy = "playbackId"
	GetViewershipMetricsQueryParamBreakdownByDeviceType    GetViewershipMetricsQueryParamBreakdownBy = "deviceType"
	GetViewershipMetricsQueryParamBreakdownByDevice        GetViewershipMetricsQueryParamBreakdownBy = "device"
	GetViewershipMetricsQueryParamBreakdownByCPU           GetViewershipMetricsQueryParamBreakdownBy = "cpu"
	GetViewershipMetricsQueryParamBreakdownByOs            GetViewershipMetricsQueryParamBreakdownBy = "os"
	GetViewershipMetricsQueryParamBreakdownByBrowser       GetViewershipMetricsQueryParamBreakdownBy = "browser"
	GetViewershipMetricsQueryParamBreakdownByBrowserEngine GetViewershipMetricsQueryParamBreakdownBy = "browserEngine"
	GetViewershipMetricsQueryParamBreakdownByContinent     GetViewershipMetricsQueryParamBreakdownBy = "continent"
	GetViewershipMetricsQueryParamBreakdownByCountry       GetViewershipMetricsQueryParamBreakdownBy = "country"
	GetViewershipMetricsQueryParamBreakdownBySubdivision   GetViewershipMetricsQueryParamBreakdownBy = "subdivision"
	GetViewershipMetricsQueryParamBreakdownByTimezone      GetViewershipMetricsQueryParamBreakdownBy = "timezone"
	GetViewershipMetricsQueryParamBreakdownByGeohash       GetViewershipMetricsQueryParamBreakdownBy = "geohash"
	GetViewershipMetricsQueryParamBreakdownByViewerID      GetViewershipMetricsQueryParamBreakdownBy = "viewerId"
	GetViewershipMetricsQueryParamBreakdownByCreatorID     GetViewershipMetricsQueryParamBreakdownBy = "creatorId"
)

func (e GetViewershipMetricsQueryParamBreakdownBy) ToPointer() *GetViewershipMetricsQueryParamBreakdownBy {
	return &e
}
func (e *GetViewershipMetricsQueryParamBreakdownBy) UnmarshalJSON(data []byte) error {
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
		*e = GetViewershipMetricsQueryParamBreakdownBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetViewershipMetricsQueryParamBreakdownBy: %v", v)
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
	TimeStep *QueryParamTimeStep `queryParam:"style=form,explode=true,name=timeStep"`
	// The asset ID to filter metrics for
	AssetID *string `queryParam:"style=form,explode=true,name=assetId"`
	// The stream ID to filter metrics for
	StreamID *string `queryParam:"style=form,explode=true,name=streamId"`
	// The creator ID to filter the query results
	CreatorID *string `queryParam:"style=form,explode=true,name=creatorId"`
	// The list of fields to break down the query results. Specify this
	// query-string multiple times to break down by multiple fields.
	//
	BreakdownBy []GetViewershipMetricsQueryParamBreakdownBy `queryParam:"style=form,explode=true,name=breakdownBy[]"`
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

func (o *GetViewershipMetricsRequest) GetTimeStep() *QueryParamTimeStep {
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

func (o *GetViewershipMetricsRequest) GetBreakdownBy() []GetViewershipMetricsQueryParamBreakdownBy {
	if o == nil {
		return nil
	}
	return o.BreakdownBy
}

type GetViewershipMetricsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
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
