// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type NewAssetPayloadIpfs1 struct {
	Spec *Spec `json:"spec,omitempty"`
}

func (o *NewAssetPayloadIpfs1) GetSpec() *Spec {
	if o == nil {
		return nil
	}
	return o.Spec
}

type NewAssetPayloadIpfsType string

const (
	NewAssetPayloadIpfsTypeNewAssetPayloadIpfs1 NewAssetPayloadIpfsType = "new-asset-payload_ipfs_1"
	NewAssetPayloadIpfsTypeBoolean              NewAssetPayloadIpfsType = "boolean"
)

// NewAssetPayloadIpfs - Set to true to make default export to IPFS. To customize the
// pinned files, specify an object with a spec field. False or null
// means to unpin from IPFS, but it's unsupported right now.
type NewAssetPayloadIpfs struct {
	NewAssetPayloadIpfs1 *NewAssetPayloadIpfs1
	Boolean              *bool

	Type NewAssetPayloadIpfsType
}

func CreateNewAssetPayloadIpfsNewAssetPayloadIpfs1(newAssetPayloadIpfs1 NewAssetPayloadIpfs1) NewAssetPayloadIpfs {
	typ := NewAssetPayloadIpfsTypeNewAssetPayloadIpfs1

	return NewAssetPayloadIpfs{
		NewAssetPayloadIpfs1: &newAssetPayloadIpfs1,
		Type:                 typ,
	}
}

func CreateNewAssetPayloadIpfsBoolean(boolean bool) NewAssetPayloadIpfs {
	typ := NewAssetPayloadIpfsTypeBoolean

	return NewAssetPayloadIpfs{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *NewAssetPayloadIpfs) UnmarshalJSON(data []byte) error {

	var newAssetPayloadIpfs1 NewAssetPayloadIpfs1 = NewAssetPayloadIpfs1{}
	if err := utils.UnmarshalJSON(data, &newAssetPayloadIpfs1, "", true, true); err == nil {
		u.NewAssetPayloadIpfs1 = &newAssetPayloadIpfs1
		u.Type = NewAssetPayloadIpfsTypeNewAssetPayloadIpfs1
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = NewAssetPayloadIpfsTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for NewAssetPayloadIpfs", string(data))
}

func (u NewAssetPayloadIpfs) MarshalJSON() ([]byte, error) {
	if u.NewAssetPayloadIpfs1 != nil {
		return utils.MarshalJSON(u.NewAssetPayloadIpfs1, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type NewAssetPayloadIpfs: all fields are null")
}

type NewAssetPayloadStorage struct {
	// Set to true to make default export to IPFS. To customize the
	// pinned files, specify an object with a spec field. False or null
	// means to unpin from IPFS, but it's unsupported right now.
	//
	Ipfs *NewAssetPayloadIpfs `json:"ipfs,omitempty"`
}

func (o *NewAssetPayloadStorage) GetIpfs() *NewAssetPayloadIpfs {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

type NewAssetPayloadEncryption struct {
	// Encryption key used to encrypt the asset. Only writable in the upload asset endpoints and cannot be retrieved back.
	EncryptedKey string `json:"encryptedKey"`
}

func (o *NewAssetPayloadEncryption) GetEncryptedKey() string {
	if o == nil {
		return ""
	}
	return o.EncryptedKey
}

type NewAssetPayload struct {
	// The name of the asset. This is not necessarily the filename - it can be a custom name or title.
	//
	Name string `json:"name"`
	// Whether to generate MP4s for the asset.
	StaticMp4 *bool `json:"staticMp4,omitempty"`
	// Whether the playback policy for a asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy            `json:"playbackPolicy,omitempty"`
	CreatorID      *InputCreatorID            `json:"creatorId,omitempty"`
	Storage        *NewAssetPayloadStorage    `json:"storage,omitempty"`
	Encryption     *NewAssetPayloadEncryption `json:"encryption,omitempty"`
	// Decides if the output video should include C2PA signature
	C2pa     *bool              `json:"c2pa,omitempty"`
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
	// How many seconds the duration of each output segment should be
	TargetSegmentSizeSecs *float64 `json:"targetSegmentSizeSecs,omitempty"`
}

func (o *NewAssetPayload) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *NewAssetPayload) GetStaticMp4() *bool {
	if o == nil {
		return nil
	}
	return o.StaticMp4
}

func (o *NewAssetPayload) GetPlaybackPolicy() *PlaybackPolicy {
	if o == nil {
		return nil
	}
	return o.PlaybackPolicy
}

func (o *NewAssetPayload) GetCreatorID() *InputCreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *NewAssetPayload) GetStorage() *NewAssetPayloadStorage {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *NewAssetPayload) GetEncryption() *NewAssetPayloadEncryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *NewAssetPayload) GetC2pa() *bool {
	if o == nil {
		return nil
	}
	return o.C2pa
}

func (o *NewAssetPayload) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *NewAssetPayload) GetTargetSegmentSizeSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TargetSegmentSizeSecs
}
