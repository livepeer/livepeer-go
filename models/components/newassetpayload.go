// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"livepeer/internal/utils"
)

// NewAssetPayload1 - Set to true to make default export to IPFS. To customize the
// pinned files, specify an object with a spec field. False or null
// means to unpin from IPFS, but it's unsupported right now.
type NewAssetPayload1 struct {
	Spec *Spec `json:"spec,omitempty"`
}

func (o *NewAssetPayload1) GetSpec() *Spec {
	if o == nil {
		return nil
	}
	return o.Spec
}

type NewAssetPayloadIpfsType string

const (
	NewAssetPayloadIpfsTypeNewAssetPayload1 NewAssetPayloadIpfsType = "new-asset-payload_1"
	NewAssetPayloadIpfsTypeBoolean          NewAssetPayloadIpfsType = "boolean"
)

type NewAssetPayloadIpfs struct {
	NewAssetPayload1 *NewAssetPayload1
	Boolean          *bool

	Type NewAssetPayloadIpfsType
}

func CreateNewAssetPayloadIpfsNewAssetPayload1(newAssetPayload1 NewAssetPayload1) NewAssetPayloadIpfs {
	typ := NewAssetPayloadIpfsTypeNewAssetPayload1

	return NewAssetPayloadIpfs{
		NewAssetPayload1: &newAssetPayload1,
		Type:             typ,
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

	newAssetPayload1 := NewAssetPayload1{}
	if err := utils.UnmarshalJSON(data, &newAssetPayload1, "", true, true); err == nil {
		u.NewAssetPayload1 = &newAssetPayload1
		u.Type = NewAssetPayloadIpfsTypeNewAssetPayload1
		return nil
	}

	boolean := false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = NewAssetPayloadIpfsTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u NewAssetPayloadIpfs) MarshalJSON() ([]byte, error) {
	if u.NewAssetPayload1 != nil {
		return utils.MarshalJSON(u.NewAssetPayload1, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
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
	// Name of the asset. This is not necessarily the filename, can be a
	// custom name or title
	//
	Name string `json:"name"`
	// Whether to generate MP4s for the asset.
	StaticMp4 *bool `json:"staticMp4,omitempty"`
	// Whether the playback policy for a asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy         `json:"playbackPolicy,omitempty"`
	CreatorID      *InputCreatorID         `json:"creatorId,omitempty"`
	Storage        *NewAssetPayloadStorage `json:"storage,omitempty"`
	// URL where the asset contents can be retrieved. Only allowed (and
	// also required) in the upload asset via URL endpoint. For an IPFS
	// source, this should be similar to: `ipfs://{CID}`. For an Arweave
	// source: `ar://{CID}`.
	//
	URL        *string                    `json:"url,omitempty"`
	Encryption *NewAssetPayloadEncryption `json:"encryption,omitempty"`
	// Decides if the output video should include C2PA signature
	C2pa *bool `json:"c2pa,omitempty"`
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

func (o *NewAssetPayload) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
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
