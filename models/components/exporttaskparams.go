// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"livepeer/internal/utils"
)

// ExportTaskParams2 - Parameters for the export task
type ExportTaskParams2 struct {
	Ipfs IpfsExportParams `json:"ipfs"`
}

func (o *ExportTaskParams2) GetIpfs() IpfsExportParams {
	if o == nil {
		return IpfsExportParams{}
	}
	return o.Ipfs
}

// Custom - custom URL parameters for the export task
type Custom struct {
	// URL where to export the asset
	URL string `json:"url"`
	// Method to use on the export request
	Method *string `default:"PUT" json:"method"`
	// Headers to add to the export request
	Headers map[string]string `json:"headers,omitempty"`
}

func (c Custom) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Custom) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Custom) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *Custom) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *Custom) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

// ExportTaskParams1 - Parameters for the export task
type ExportTaskParams1 struct {
	// custom URL parameters for the export task
	Custom Custom `json:"custom"`
}

func (o *ExportTaskParams1) GetCustom() Custom {
	if o == nil {
		return Custom{}
	}
	return o.Custom
}

type ExportTaskParamsType string

const (
	ExportTaskParamsTypeExportTaskParams1 ExportTaskParamsType = "export-task-params_1"
	ExportTaskParamsTypeExportTaskParams2 ExportTaskParamsType = "export-task-params_2"
)

type ExportTaskParams struct {
	ExportTaskParams1 *ExportTaskParams1
	ExportTaskParams2 *ExportTaskParams2

	Type ExportTaskParamsType
}

func CreateExportTaskParamsExportTaskParams1(exportTaskParams1 ExportTaskParams1) ExportTaskParams {
	typ := ExportTaskParamsTypeExportTaskParams1

	return ExportTaskParams{
		ExportTaskParams1: &exportTaskParams1,
		Type:              typ,
	}
}

func CreateExportTaskParamsExportTaskParams2(exportTaskParams2 ExportTaskParams2) ExportTaskParams {
	typ := ExportTaskParamsTypeExportTaskParams2

	return ExportTaskParams{
		ExportTaskParams2: &exportTaskParams2,
		Type:              typ,
	}
}

func (u *ExportTaskParams) UnmarshalJSON(data []byte) error {

	exportTaskParams1 := ExportTaskParams1{}
	if err := utils.UnmarshalJSON(data, &exportTaskParams1, "", true, true); err == nil {
		u.ExportTaskParams1 = &exportTaskParams1
		u.Type = ExportTaskParamsTypeExportTaskParams1
		return nil
	}

	exportTaskParams2 := ExportTaskParams2{}
	if err := utils.UnmarshalJSON(data, &exportTaskParams2, "", true, true); err == nil {
		u.ExportTaskParams2 = &exportTaskParams2
		u.Type = ExportTaskParamsTypeExportTaskParams2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ExportTaskParams) MarshalJSON() ([]byte, error) {
	if u.ExportTaskParams1 != nil {
		return utils.MarshalJSON(u.ExportTaskParams1, "", true)
	}

	if u.ExportTaskParams2 != nil {
		return utils.MarshalJSON(u.ExportTaskParams2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
