// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"livepeer/internal/utils"
)

// TranscodePayloadType - Type of service. This is optional and defaults to `url` if
// ŚURL field is provided.
type TranscodePayloadType string

const (
	TranscodePayloadTypeS3 TranscodePayloadType = "s3"
)

func (e TranscodePayloadType) ToPointer() *TranscodePayloadType {
	return &e
}

func (e *TranscodePayloadType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "s3":
		*e = TranscodePayloadType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TranscodePayloadType: %v", v)
	}
}

// Credentials for the private input video storage
type Credentials struct {
	// Access Key ID
	AccessKeyID string `json:"accessKeyId"`
	// Secret Access Key
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *Credentials) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *Credentials) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// TranscodePayload2 - S3-like storage input video
type TranscodePayload2 struct {
	// Type of service. This is optional and defaults to `url` if
	// ŚURL field is provided.
	//
	Type TranscodePayloadType `json:"type"`
	// Service endpoint URL (AWS S3 endpoint list: https://docs.aws.amazon.com/general/latest/gr/s3.html, GCP S3 endpoint: https://storage.googleapis.com, Storj: https://gateway.storjshare.io)
	Endpoint string `json:"endpoint"`
	// Bucket with input file
	Bucket string `json:"bucket"`
	// Path to the input file inside the bucket
	Path string `json:"path"`
	// Credentials for the private input video storage
	Credentials Credentials `json:"credentials"`
}

func (o *TranscodePayload2) GetType() TranscodePayloadType {
	if o == nil {
		return TranscodePayloadType("")
	}
	return o.Type
}

func (o *TranscodePayload2) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *TranscodePayload2) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *TranscodePayload2) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *TranscodePayload2) GetCredentials() Credentials {
	if o == nil {
		return Credentials{}
	}
	return o.Credentials
}

// TranscodePayload1 - URL input video
type TranscodePayload1 struct {
	// URL of the video to transcode
	URL string `json:"url"`
}

func (o *TranscodePayload1) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type InputType string

const (
	InputTypeTranscodePayload1 InputType = "transcode-payload_1"
	InputTypeTranscodePayload2 InputType = "transcode-payload_2"
)

type Input struct {
	TranscodePayload1 *TranscodePayload1
	TranscodePayload2 *TranscodePayload2

	Type InputType
}

func CreateInputTranscodePayload1(transcodePayload1 TranscodePayload1) Input {
	typ := InputTypeTranscodePayload1

	return Input{
		TranscodePayload1: &transcodePayload1,
		Type:              typ,
	}
}

func CreateInputTranscodePayload2(transcodePayload2 TranscodePayload2) Input {
	typ := InputTypeTranscodePayload2

	return Input{
		TranscodePayload2: &transcodePayload2,
		Type:              typ,
	}
}

func (u *Input) UnmarshalJSON(data []byte) error {

	transcodePayload1 := TranscodePayload1{}
	if err := utils.UnmarshalJSON(data, &transcodePayload1, "", true, true); err == nil {
		u.TranscodePayload1 = &transcodePayload1
		u.Type = InputTypeTranscodePayload1
		return nil
	}

	transcodePayload2 := TranscodePayload2{}
	if err := utils.UnmarshalJSON(data, &transcodePayload2, "", true, true); err == nil {
		u.TranscodePayload2 = &transcodePayload2
		u.Type = InputTypeTranscodePayload2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Input) MarshalJSON() ([]byte, error) {
	if u.TranscodePayload1 != nil {
		return utils.MarshalJSON(u.TranscodePayload1, "", true)
	}

	if u.TranscodePayload2 != nil {
		return utils.MarshalJSON(u.TranscodePayload2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// TranscodePayloadSchemasStorageType - Type of service used for output files
type TranscodePayloadSchemasStorageType string

const (
	TranscodePayloadSchemasStorageTypeWeb3Storage TranscodePayloadSchemasStorageType = "web3.storage"
)

func (e TranscodePayloadSchemasStorageType) ToPointer() *TranscodePayloadSchemasStorageType {
	return &e
}

func (e *TranscodePayloadSchemasStorageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web3.storage":
		*e = TranscodePayloadSchemasStorageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TranscodePayloadSchemasStorageType: %v", v)
	}
}

// TranscodePayloadSchemasCredentials - Delegation proof for Livepeer to be able to upload to
// web3.storage
type TranscodePayloadSchemasCredentials struct {
	// Base64 encoded UCAN delegation proof
	Proof string `json:"proof"`
}

func (o *TranscodePayloadSchemasCredentials) GetProof() string {
	if o == nil {
		return ""
	}
	return o.Proof
}

// TranscodePayloadSchemas2 - Storage for the output files
type TranscodePayloadSchemas2 struct {
	// Type of service used for output files
	Type TranscodePayloadSchemasStorageType `json:"type"`
	// Delegation proof for Livepeer to be able to upload to
	// web3.storage
	//
	Credentials TranscodePayloadSchemasCredentials `json:"credentials"`
}

func (o *TranscodePayloadSchemas2) GetType() TranscodePayloadSchemasStorageType {
	if o == nil {
		return TranscodePayloadSchemasStorageType("")
	}
	return o.Type
}

func (o *TranscodePayloadSchemas2) GetCredentials() TranscodePayloadSchemasCredentials {
	if o == nil {
		return TranscodePayloadSchemasCredentials{}
	}
	return o.Credentials
}

// TranscodePayloadSchemasType - Type of service used for output files
type TranscodePayloadSchemasType string

const (
	TranscodePayloadSchemasTypeS3 TranscodePayloadSchemasType = "s3"
)

func (e TranscodePayloadSchemasType) ToPointer() *TranscodePayloadSchemasType {
	return &e
}

func (e *TranscodePayloadSchemasType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "s3":
		*e = TranscodePayloadSchemasType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TranscodePayloadSchemasType: %v", v)
	}
}

// TranscodePayloadCredentials - Credentials for the output video storage
type TranscodePayloadCredentials struct {
	// Access Key ID
	AccessKeyID string `json:"accessKeyId"`
	// Secret Access Key
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *TranscodePayloadCredentials) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *TranscodePayloadCredentials) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// TranscodePayloadSchemas1 - Storage for the output files
type TranscodePayloadSchemas1 struct {
	// Type of service used for output files
	Type TranscodePayloadSchemasType `json:"type"`
	// Service endpoint URL (AWS S3 endpoint list: https://docs.aws.amazon.com/general/latest/gr/s3.html, GCP S3 endpoint: https://storage.googleapis.com, Storj: https://gateway.storjshare.io)
	Endpoint string `json:"endpoint"`
	// Bucket with output files
	Bucket string `json:"bucket"`
	// Credentials for the output video storage
	Credentials TranscodePayloadCredentials `json:"credentials"`
}

func (o *TranscodePayloadSchemas1) GetType() TranscodePayloadSchemasType {
	if o == nil {
		return TranscodePayloadSchemasType("")
	}
	return o.Type
}

func (o *TranscodePayloadSchemas1) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *TranscodePayloadSchemas1) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *TranscodePayloadSchemas1) GetCredentials() TranscodePayloadCredentials {
	if o == nil {
		return TranscodePayloadCredentials{}
	}
	return o.Credentials
}

type TranscodePayloadStorageType string

const (
	TranscodePayloadStorageTypeTranscodePayloadSchemas1 TranscodePayloadStorageType = "transcode-payload_Schemas_1"
	TranscodePayloadStorageTypeTranscodePayloadSchemas2 TranscodePayloadStorageType = "transcode-payload_Schemas_2"
)

type TranscodePayloadStorage struct {
	TranscodePayloadSchemas1 *TranscodePayloadSchemas1
	TranscodePayloadSchemas2 *TranscodePayloadSchemas2

	Type TranscodePayloadStorageType
}

func CreateTranscodePayloadStorageTranscodePayloadSchemas1(transcodePayloadSchemas1 TranscodePayloadSchemas1) TranscodePayloadStorage {
	typ := TranscodePayloadStorageTypeTranscodePayloadSchemas1

	return TranscodePayloadStorage{
		TranscodePayloadSchemas1: &transcodePayloadSchemas1,
		Type:                     typ,
	}
}

func CreateTranscodePayloadStorageTranscodePayloadSchemas2(transcodePayloadSchemas2 TranscodePayloadSchemas2) TranscodePayloadStorage {
	typ := TranscodePayloadStorageTypeTranscodePayloadSchemas2

	return TranscodePayloadStorage{
		TranscodePayloadSchemas2: &transcodePayloadSchemas2,
		Type:                     typ,
	}
}

func (u *TranscodePayloadStorage) UnmarshalJSON(data []byte) error {

	transcodePayloadSchemas2 := TranscodePayloadSchemas2{}
	if err := utils.UnmarshalJSON(data, &transcodePayloadSchemas2, "", true, true); err == nil {
		u.TranscodePayloadSchemas2 = &transcodePayloadSchemas2
		u.Type = TranscodePayloadStorageTypeTranscodePayloadSchemas2
		return nil
	}

	transcodePayloadSchemas1 := TranscodePayloadSchemas1{}
	if err := utils.UnmarshalJSON(data, &transcodePayloadSchemas1, "", true, true); err == nil {
		u.TranscodePayloadSchemas1 = &transcodePayloadSchemas1
		u.Type = TranscodePayloadStorageTypeTranscodePayloadSchemas1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TranscodePayloadStorage) MarshalJSON() ([]byte, error) {
	if u.TranscodePayloadSchemas1 != nil {
		return utils.MarshalJSON(u.TranscodePayloadSchemas1, "", true)
	}

	if u.TranscodePayloadSchemas2 != nil {
		return utils.MarshalJSON(u.TranscodePayloadSchemas2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Hls - HLS output format
type Hls struct {
	// Path for the HLS output
	Path string `json:"path"`
}

func (o *Hls) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// Mp4 - MP4 output format
type Mp4 struct {
	// Path for the MP4 output
	Path string `json:"path"`
}

func (o *Mp4) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// Fmp4 - FMP4 output format
type Fmp4 struct {
	// Path for the FMP4 output
	Path string `json:"path"`
}

func (o *Fmp4) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// Outputs - Output formats
type Outputs struct {
	// HLS output format
	Hls *Hls `json:"hls,omitempty"`
	// MP4 output format
	Mp4 *Mp4 `json:"mp4,omitempty"`
	// FMP4 output format
	Fmp4 *Fmp4 `json:"fmp4,omitempty"`
}

func (o *Outputs) GetHls() *Hls {
	if o == nil {
		return nil
	}
	return o.Hls
}

func (o *Outputs) GetMp4() *Mp4 {
	if o == nil {
		return nil
	}
	return o.Mp4
}

func (o *Outputs) GetFmp4() *Fmp4 {
	if o == nil {
		return nil
	}
	return o.Fmp4
}

type TranscodePayload struct {
	Input   Input                   `json:"input"`
	Storage TranscodePayloadStorage `json:"storage"`
	// Output formats
	Outputs  Outputs            `json:"outputs"`
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
	// How many seconds the duration of each output segment should be
	TargetSegmentSizeSecs *float64        `json:"targetSegmentSizeSecs,omitempty"`
	CreatorID             *InputCreatorID `json:"creatorId,omitempty"`
	// Decides if the output video should include C2PA signature
	C2pa *bool `json:"c2pa,omitempty"`
}

func (o *TranscodePayload) GetInput() Input {
	if o == nil {
		return Input{}
	}
	return o.Input
}

func (o *TranscodePayload) GetStorage() TranscodePayloadStorage {
	if o == nil {
		return TranscodePayloadStorage{}
	}
	return o.Storage
}

func (o *TranscodePayload) GetOutputs() Outputs {
	if o == nil {
		return Outputs{}
	}
	return o.Outputs
}

func (o *TranscodePayload) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *TranscodePayload) GetTargetSegmentSizeSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TargetSegmentSizeSecs
}

func (o *TranscodePayload) GetCreatorID() *InputCreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *TranscodePayload) GetC2pa() *bool {
	if o == nil {
		return nil
	}
	return o.C2pa
}
