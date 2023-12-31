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

// ExportTaskParamsSchemas1 - Parameters for the export task
type ExportTaskParamsSchemas1 struct {
	// custom URL parameters for the export task
	Custom Custom `json:"custom"`
}

func (o *ExportTaskParamsSchemas1) GetCustom() Custom {
	if o == nil {
		return Custom{}
	}
	return o.Custom
}

type ExportTaskParamsType string

const (
	ExportTaskParamsTypeExportTaskParamsSchemas1 ExportTaskParamsType = "export-task-params_Schemas_1"
	ExportTaskParamsTypeExportTaskParams2        ExportTaskParamsType = "export-task-params_2"
)

type ExportTaskParams struct {
	ExportTaskParamsSchemas1 *ExportTaskParamsSchemas1
	ExportTaskParams2        *ExportTaskParams2

	Type ExportTaskParamsType
}

func CreateExportTaskParamsExportTaskParamsSchemas1(exportTaskParamsSchemas1 ExportTaskParamsSchemas1) ExportTaskParams {
	typ := ExportTaskParamsTypeExportTaskParamsSchemas1

	return ExportTaskParams{
		ExportTaskParamsSchemas1: &exportTaskParamsSchemas1,
		Type:                     typ,
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

	exportTaskParamsSchemas1 := ExportTaskParamsSchemas1{}
	if err := utils.UnmarshalJSON(data, &exportTaskParamsSchemas1, "", true, true); err == nil {
		u.ExportTaskParamsSchemas1 = &exportTaskParamsSchemas1
		u.Type = ExportTaskParamsTypeExportTaskParamsSchemas1
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
	if u.ExportTaskParamsSchemas1 != nil {
		return utils.MarshalJSON(u.ExportTaskParamsSchemas1, "", true)
	}

	if u.ExportTaskParams2 != nil {
		return utils.MarshalJSON(u.ExportTaskParams2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// ExportTaskParamsSchemas2 - Parameters for the export task
type ExportTaskParamsSchemas2 struct {
	Ipfs IpfsExportParams1 `json:"ipfs"`
}

func (o *ExportTaskParamsSchemas2) GetIpfs() IpfsExportParams1 {
	if o == nil {
		return IpfsExportParams1{}
	}
	return o.Ipfs
}

type ExportTaskParams1Type string

const (
	ExportTaskParams1TypeExportTaskParamsSchemas1 ExportTaskParams1Type = "export-task-params_Schemas_1"
	ExportTaskParams1TypeExportTaskParamsSchemas2 ExportTaskParams1Type = "export-task-params_Schemas_2"
)

type ExportTaskParams1 struct {
	ExportTaskParamsSchemas1 *ExportTaskParamsSchemas1
	ExportTaskParamsSchemas2 *ExportTaskParamsSchemas2

	Type ExportTaskParams1Type
}

func CreateExportTaskParams1ExportTaskParamsSchemas1(exportTaskParamsSchemas1 ExportTaskParamsSchemas1) ExportTaskParams1 {
	typ := ExportTaskParams1TypeExportTaskParamsSchemas1

	return ExportTaskParams1{
		ExportTaskParamsSchemas1: &exportTaskParamsSchemas1,
		Type:                     typ,
	}
}

func CreateExportTaskParams1ExportTaskParamsSchemas2(exportTaskParamsSchemas2 ExportTaskParamsSchemas2) ExportTaskParams1 {
	typ := ExportTaskParams1TypeExportTaskParamsSchemas2

	return ExportTaskParams1{
		ExportTaskParamsSchemas2: &exportTaskParamsSchemas2,
		Type:                     typ,
	}
}

func (u *ExportTaskParams1) UnmarshalJSON(data []byte) error {

	exportTaskParamsSchemas1 := ExportTaskParamsSchemas1{}
	if err := utils.UnmarshalJSON(data, &exportTaskParamsSchemas1, "", true, true); err == nil {
		u.ExportTaskParamsSchemas1 = &exportTaskParamsSchemas1
		u.Type = ExportTaskParams1TypeExportTaskParamsSchemas1
		return nil
	}

	exportTaskParamsSchemas2 := ExportTaskParamsSchemas2{}
	if err := utils.UnmarshalJSON(data, &exportTaskParamsSchemas2, "", true, true); err == nil {
		u.ExportTaskParamsSchemas2 = &exportTaskParamsSchemas2
		u.Type = ExportTaskParams1TypeExportTaskParamsSchemas2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ExportTaskParams1) MarshalJSON() ([]byte, error) {
	if u.ExportTaskParamsSchemas1 != nil {
		return utils.MarshalJSON(u.ExportTaskParamsSchemas1, "", true)
	}

	if u.ExportTaskParamsSchemas2 != nil {
		return utils.MarshalJSON(u.ExportTaskParamsSchemas2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
