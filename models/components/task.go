// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
)

// TaskType - Type of the task
type TaskType string

const (
	TaskTypeUpload        TaskType = "upload"
	TaskTypeExport        TaskType = "export"
	TaskTypeExportData    TaskType = "export-data"
	TaskTypeTranscodeFile TaskType = "transcode-file"
	TaskTypeClip          TaskType = "clip"
)

func (e TaskType) ToPointer() *TaskType {
	return &e
}
func (e *TaskType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "upload":
		fallthrough
	case "export":
		fallthrough
	case "export-data":
		fallthrough
	case "transcode-file":
		fallthrough
	case "clip":
		*e = TaskType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskType: %v", v)
	}
}

// Upload - Parameters for the upload task
type Upload struct {
	// URL of the asset to "upload"
	URL        *string           `json:"url,omitempty"`
	Encryption *EncryptionOutput `json:"encryption,omitempty"`
	// Decides if the output video should include C2PA signature
	C2pa     *bool              `json:"c2pa,omitempty"`
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
	// How many seconds the duration of each output segment should be
	TargetSegmentSizeSecs *float64 `json:"targetSegmentSizeSecs,omitempty"`
}

func (o *Upload) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Upload) GetEncryption() *EncryptionOutput {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *Upload) GetC2pa() *bool {
	if o == nil {
		return nil
	}
	return o.C2pa
}

func (o *Upload) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *Upload) GetTargetSegmentSizeSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TargetSegmentSizeSecs
}

// Content - File content to store into IPFS
type Content struct {
}

// TaskExportData - Parameters for the export-data task
type TaskExportData struct {
	// File content to store into IPFS
	Content Content           `json:"content"`
	Ipfs    *IpfsExportParams `json:"ipfs,omitempty"`
	// Optional type of content
	Type *string `json:"type,omitempty"`
	// Optional ID of the content
	ID *string `json:"id,omitempty"`
}

func (o *TaskExportData) GetContent() Content {
	if o == nil {
		return Content{}
	}
	return o.Content
}

func (o *TaskExportData) GetIpfs() *IpfsExportParams {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

func (o *TaskExportData) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TaskExportData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TaskInput - Input video file to transcode
type TaskInput struct {
	// URL of a video to transcode, accepts object-store format
	// "s3+https"
	//
	URL *string `json:"url,omitempty"`
}

func (o *TaskInput) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// TaskStorage - Storage for the output files
type TaskStorage struct {
	// URL of the output storage, accepts object-store format
	// "s3+https"
	//
	URL *string `json:"url,omitempty"`
}

func (o *TaskStorage) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// TaskHls - HLS output format
type TaskHls struct {
	// Path for the HLS output
	Path *string `json:"path,omitempty"`
}

func (o *TaskHls) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// TaskMp4 - MP4 output format
type TaskMp4 struct {
	// Path for the MP4 output
	Path *string `json:"path,omitempty"`
}

func (o *TaskMp4) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// TaskOutputs - Output formats
type TaskOutputs struct {
	// HLS output format
	Hls *TaskHls `json:"hls,omitempty"`
	// MP4 output format
	Mp4 *TaskMp4 `json:"mp4,omitempty"`
}

func (o *TaskOutputs) GetHls() *TaskHls {
	if o == nil {
		return nil
	}
	return o.Hls
}

func (o *TaskOutputs) GetMp4() *TaskMp4 {
	if o == nil {
		return nil
	}
	return o.Mp4
}

// TranscodeFile - Parameters for the transcode-file task
type TranscodeFile struct {
	// Input video file to transcode
	Input *TaskInput `json:"input,omitempty"`
	// Storage for the output files
	Storage *TaskStorage `json:"storage,omitempty"`
	// Output formats
	Outputs  *TaskOutputs       `json:"outputs,omitempty"`
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
	// How many seconds the duration of each output segment should
	// be
	//
	TargetSegmentSizeSecs *float64        `json:"targetSegmentSizeSecs,omitempty"`
	CreatorID             *InputCreatorID `json:"creatorId,omitempty"`
	// Decides if the output video should include C2PA signature
	C2pa *bool `json:"c2pa,omitempty"`
}

func (o *TranscodeFile) GetInput() *TaskInput {
	if o == nil {
		return nil
	}
	return o.Input
}

func (o *TranscodeFile) GetStorage() *TaskStorage {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *TranscodeFile) GetOutputs() *TaskOutputs {
	if o == nil {
		return nil
	}
	return o.Outputs
}

func (o *TranscodeFile) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *TranscodeFile) GetTargetSegmentSizeSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TargetSegmentSizeSecs
}

func (o *TranscodeFile) GetCreatorID() *InputCreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *TranscodeFile) GetC2pa() *bool {
	if o == nil {
		return nil
	}
	return o.C2pa
}

// ClipStrategy - Strategy to use for clipping the asset. If not specified, the default strategy that Catalyst is configured for will be used. This field only available for admin users, and is only used for E2E testing.
type ClipStrategy struct {
	// The start timestamp of the clip in Unix milliseconds. _See the ClipTrigger in the UI Kit for an example of how this is calculated (for HLS, it uses `Program Date-Time` tags, and for WebRTC, it uses the latency from server to client at stream startup)._
	StartTime *float64 `json:"startTime,omitempty"`
	// The end timestamp of the clip in Unix milliseconds. _See the ClipTrigger in the UI Kit for an example of how this is calculated (for HLS, it uses `Program Date-Time` tags, and for WebRTC, it uses the latency from server to client at stream startup)._
	EndTime *float64 `json:"endTime,omitempty"`
	// The playback ID of the stream or stream recording to clip. Asset playback IDs are not supported yet.
	PlaybackID *string `json:"playbackId,omitempty"`
}

func (o *ClipStrategy) GetStartTime() *float64 {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *ClipStrategy) GetEndTime() *float64 {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ClipStrategy) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

// CatalystPipelineStrategy - Force to use a specific strategy in the Catalyst pipeline. If not specified, the default strategy that Catalyst is configured for will be used. This field only available for admin users, and is only used for E2E testing.
type CatalystPipelineStrategy string

const (
	CatalystPipelineStrategyCatalyst           CatalystPipelineStrategy = "catalyst"
	CatalystPipelineStrategyCatalystFfmpeg     CatalystPipelineStrategy = "catalyst_ffmpeg"
	CatalystPipelineStrategyBackgroundExternal CatalystPipelineStrategy = "background_external"
	CatalystPipelineStrategyBackgroundMist     CatalystPipelineStrategy = "background_mist"
	CatalystPipelineStrategyFallbackExternal   CatalystPipelineStrategy = "fallback_external"
	CatalystPipelineStrategyExternal           CatalystPipelineStrategy = "external"
)

func (e CatalystPipelineStrategy) ToPointer() *CatalystPipelineStrategy {
	return &e
}
func (e *CatalystPipelineStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "catalyst":
		fallthrough
	case "catalyst_ffmpeg":
		fallthrough
	case "background_external":
		fallthrough
	case "background_mist":
		fallthrough
	case "fallback_external":
		fallthrough
	case "external":
		*e = CatalystPipelineStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CatalystPipelineStrategy: %v", v)
	}
}

type Clip struct {
	// URL of the asset to "clip"
	URL *string `json:"url,omitempty"`
	// Strategy to use for clipping the asset. If not specified, the default strategy that Catalyst is configured for will be used. This field only available for admin users, and is only used for E2E testing.
	ClipStrategy *ClipStrategy `json:"clipStrategy,omitempty"`
	// Force to use a specific strategy in the Catalyst pipeline. If not specified, the default strategy that Catalyst is configured for will be used. This field only available for admin users, and is only used for E2E testing.
	CatalystPipelineStrategy *CatalystPipelineStrategy `json:"catalystPipelineStrategy,omitempty"`
	// ID of the session
	SessionID *string `json:"sessionId,omitempty"`
	// ID of the input asset or stream
	InputID *string `json:"inputId,omitempty"`
}

func (o *Clip) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Clip) GetClipStrategy() *ClipStrategy {
	if o == nil {
		return nil
	}
	return o.ClipStrategy
}

func (o *Clip) GetCatalystPipelineStrategy() *CatalystPipelineStrategy {
	if o == nil {
		return nil
	}
	return o.CatalystPipelineStrategy
}

func (o *Clip) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

func (o *Clip) GetInputID() *string {
	if o == nil {
		return nil
	}
	return o.InputID
}

// Params - Parameters of the task
type Params struct {
	// Parameters for the upload task
	Upload *Upload `json:"upload,omitempty"`
	// Parameters for the export task
	Export *ExportTaskParams `json:"export,omitempty"`
	// Parameters for the export-data task
	ExportData *TaskExportData `json:"exportData,omitempty"`
	// Parameters for the transcode-file task
	TranscodeFile *TranscodeFile `json:"transcode-file,omitempty"`
	Clip          *Clip          `json:"clip,omitempty"`
}

func (o *Params) GetUpload() *Upload {
	if o == nil {
		return nil
	}
	return o.Upload
}

func (o *Params) GetExport() *ExportTaskParams {
	if o == nil {
		return nil
	}
	return o.Export
}

func (o *Params) GetExportData() *TaskExportData {
	if o == nil {
		return nil
	}
	return o.ExportData
}

func (o *Params) GetTranscodeFile() *TranscodeFile {
	if o == nil {
		return nil
	}
	return o.TranscodeFile
}

func (o *Params) GetClip() *Clip {
	if o == nil {
		return nil
	}
	return o.Clip
}

// TaskPhase - Phase of the task
type TaskPhase string

const (
	TaskPhasePending   TaskPhase = "pending"
	TaskPhaseWaiting   TaskPhase = "waiting"
	TaskPhaseRunning   TaskPhase = "running"
	TaskPhaseFailed    TaskPhase = "failed"
	TaskPhaseCompleted TaskPhase = "completed"
	TaskPhaseCancelled TaskPhase = "cancelled"
)

func (e TaskPhase) ToPointer() *TaskPhase {
	return &e
}
func (e *TaskPhase) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "waiting":
		fallthrough
	case "running":
		fallthrough
	case "failed":
		fallthrough
	case "completed":
		fallthrough
	case "cancelled":
		*e = TaskPhase(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskPhase: %v", v)
	}
}

// TaskStatus - Status of the task
type TaskStatus struct {
	// Phase of the task
	Phase TaskPhase `json:"phase"`
	// Timestamp (in milliseconds) at which task was updated
	UpdatedAt float64 `json:"updatedAt"`
	// Current progress of the task in a 0-1 ratio
	Progress *float64 `json:"progress,omitempty"`
	// Error message if the task failed
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Number of retries done on the task
	Retries *float64 `json:"retries,omitempty"`
}

func (o *TaskStatus) GetPhase() TaskPhase {
	if o == nil {
		return TaskPhase("")
	}
	return o.Phase
}

func (o *TaskStatus) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *TaskStatus) GetProgress() *float64 {
	if o == nil {
		return nil
	}
	return o.Progress
}

func (o *TaskStatus) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *TaskStatus) GetRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

// TaskUpload - Output of the upload task
type TaskUpload struct {
	AssetSpec            *Asset         `json:"assetSpec,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (t TaskUpload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskUpload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskUpload) GetAssetSpec() *Asset {
	if o == nil {
		return nil
	}
	return o.AssetSpec
}

func (o *TaskUpload) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

type TaskIpfs struct {
	// IPFS CID of the exported video file
	VideoFileCid string `json:"videoFileCid"`
	// URL for the file with the IPFS protocol
	VideoFileURL *string `json:"videoFileUrl,omitempty"`
	// URL to access file via HTTP through an IPFS gateway
	VideoFileGatewayURL *string `json:"videoFileGatewayUrl,omitempty"`
	// IPFS CID of the default metadata exported for the video
	NftMetadataCid *string `json:"nftMetadataCid,omitempty"`
	// URL for the metadata file with the IPFS protocol
	NftMetadataURL *string `json:"nftMetadataUrl,omitempty"`
	// URL to access metadata file via HTTP through an IPFS
	// gateway
	//
	NftMetadataGatewayURL *string `json:"nftMetadataGatewayUrl,omitempty"`
}

func (o *TaskIpfs) GetVideoFileCid() string {
	if o == nil {
		return ""
	}
	return o.VideoFileCid
}

func (o *TaskIpfs) GetVideoFileURL() *string {
	if o == nil {
		return nil
	}
	return o.VideoFileURL
}

func (o *TaskIpfs) GetVideoFileGatewayURL() *string {
	if o == nil {
		return nil
	}
	return o.VideoFileGatewayURL
}

func (o *TaskIpfs) GetNftMetadataCid() *string {
	if o == nil {
		return nil
	}
	return o.NftMetadataCid
}

func (o *TaskIpfs) GetNftMetadataURL() *string {
	if o == nil {
		return nil
	}
	return o.NftMetadataURL
}

func (o *TaskIpfs) GetNftMetadataGatewayURL() *string {
	if o == nil {
		return nil
	}
	return o.NftMetadataGatewayURL
}

// Export - Output of the export task
type Export struct {
	Ipfs *TaskIpfs `json:"ipfs,omitempty"`
}

func (o *Export) GetIpfs() *TaskIpfs {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

type TaskOutputIpfs struct {
	// IPFS CID of the exported data
	Cid string `json:"cid"`
}

func (o *TaskOutputIpfs) GetCid() string {
	if o == nil {
		return ""
	}
	return o.Cid
}

// ExportData - Output of the export data task
type ExportData struct {
	Ipfs *TaskOutputIpfs `json:"ipfs,omitempty"`
}

func (o *ExportData) GetIpfs() *TaskOutputIpfs {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

// Output of the task
type Output struct {
	// Output of the upload task
	Upload *TaskUpload `json:"upload,omitempty"`
	// Output of the export task
	Export *Export `json:"export,omitempty"`
	// Output of the export data task
	ExportData *ExportData `json:"exportData,omitempty"`
}

func (o *Output) GetUpload() *TaskUpload {
	if o == nil {
		return nil
	}
	return o.Upload
}

func (o *Output) GetExport() *Export {
	if o == nil {
		return nil
	}
	return o.Export
}

func (o *Output) GetExportData() *ExportData {
	if o == nil {
		return nil
	}
	return o.ExportData
}

type Task struct {
	// Task ID
	ID *string `json:"id,omitempty"`
	// Type of the task
	Type *TaskType `json:"type,omitempty"`
	// Timestamp (in milliseconds) at which task was created
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// Timestamp (in milliseconds) at which the task was scheduled for
	// execution (e.g. after file upload finished).
	//
	ScheduledAt *float64 `json:"scheduledAt,omitempty"`
	// ID of the input asset
	InputAssetID *string `json:"inputAssetId,omitempty"`
	// ID of the output asset
	OutputAssetID *string `json:"outputAssetId,omitempty"`
	// ID of the requester hash(IP + SALT + PlaybackId)
	RequesterID *string `json:"requesterId,omitempty"`
	// Parameters of the task
	Params *Params `json:"params,omitempty"`
	// Status of the task
	Status *TaskStatus `json:"status,omitempty"`
	// Output of the task
	Output *Output `json:"output,omitempty"`
}

func (o *Task) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Task) GetType() *TaskType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Task) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Task) GetScheduledAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ScheduledAt
}

func (o *Task) GetInputAssetID() *string {
	if o == nil {
		return nil
	}
	return o.InputAssetID
}

func (o *Task) GetOutputAssetID() *string {
	if o == nil {
		return nil
	}
	return o.OutputAssetID
}

func (o *Task) GetRequesterID() *string {
	if o == nil {
		return nil
	}
	return o.RequesterID
}

func (o *Task) GetParams() *Params {
	if o == nil {
		return nil
	}
	return o.Params
}

func (o *Task) GetStatus() *TaskStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Task) GetOutput() *Output {
	if o == nil {
		return nil
	}
	return o.Output
}
