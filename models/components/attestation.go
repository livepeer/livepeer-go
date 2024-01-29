// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// PrimaryType - Video Metadata EIP-712 primaryType
type PrimaryType string

const (
	PrimaryTypeVideoAttestation PrimaryType = "VideoAttestation"
)

func (e PrimaryType) ToPointer() *PrimaryType {
	return &e
}

func (e *PrimaryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VideoAttestation":
		*e = PrimaryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PrimaryType: %v", v)
	}
}

type Name string

const (
	NameVerifiableVideo Name = "Verifiable Video"
)

func (e Name) ToPointer() *Name {
	return &e
}

func (e *Name) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Verifiable Video":
		*e = Name(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Name: %v", v)
	}
}

type Version string

const (
	VersionOne Version = "1"
)

func (e Version) ToPointer() *Version {
	return &e
}

func (e *Version) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1":
		*e = Version(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Version: %v", v)
	}
}

// Domain - Video Metadata EIP-712 domain
type Domain struct {
	Name    Name    `json:"name"`
	Version Version `json:"version"`
}

func (o *Domain) GetName() Name {
	if o == nil {
		return Name("")
	}
	return o.Name
}

func (o *Domain) GetVersion() Version {
	if o == nil {
		return Version("")
	}
	return o.Version
}

type Attestations struct {
	Role    string `json:"role"`
	Address string `json:"address"`
}

func (o *Attestations) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *Attestations) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

// Message - Video Metadata EIP-712 message content
type Message struct {
	Video        string         `json:"video"`
	Attestations []Attestations `json:"attestations"`
	Signer       string         `json:"signer"`
	Timestamp    float64        `json:"timestamp"`
}

func (o *Message) GetVideo() string {
	if o == nil {
		return ""
	}
	return o.Video
}

func (o *Message) GetAttestations() []Attestations {
	if o == nil {
		return []Attestations{}
	}
	return o.Attestations
}

func (o *Message) GetSigner() string {
	if o == nil {
		return ""
	}
	return o.Signer
}

func (o *Message) GetTimestamp() float64 {
	if o == nil {
		return 0.0
	}
	return o.Timestamp
}

type SignatureType string

const (
	SignatureTypeEip712 SignatureType = "eip712"
	SignatureTypeFlow   SignatureType = "flow"
)

func (e SignatureType) ToPointer() *SignatureType {
	return &e
}

func (e *SignatureType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "eip712":
		fallthrough
	case "flow":
		*e = SignatureType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SignatureType: %v", v)
	}
}

type AttestationIpfs struct {
	DollarRef interface{} `json:"$ref,omitempty"`
	// Timestamp (in milliseconds) at which IPFS export task was updated
	//
	UpdatedAt *float64 `json:"updatedAt,omitempty"`
}

func (o *AttestationIpfs) GetDollarRef() interface{} {
	if o == nil {
		return nil
	}
	return o.DollarRef
}

func (o *AttestationIpfs) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type AttestationStorage struct {
	Ipfs   *AttestationIpfs `json:"ipfs,omitempty"`
	Status *StorageStatus   `json:"status,omitempty"`
}

func (o *AttestationStorage) GetIpfs() *AttestationIpfs {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

func (o *AttestationStorage) GetStatus() *StorageStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type Attestation struct {
	ID *string `json:"id,omitempty"`
	// Video Metadata EIP-712 primaryType
	PrimaryType PrimaryType `json:"primaryType"`
	// Video Metadata EIP-712 domain
	Domain Domain `json:"domain"`
	// Video Metadata EIP-712 message content
	Message Message `json:"message"`
	// Video Metadata EIP-712 message signature
	Signature string `json:"signature"`
	// Timestamp (in milliseconds) at which the object was created
	CreatedAt     *float64            `json:"createdAt,omitempty"`
	SignatureType *SignatureType      `json:"signatureType,omitempty"`
	Storage       *AttestationStorage `json:"storage,omitempty"`
}

func (o *Attestation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Attestation) GetPrimaryType() PrimaryType {
	if o == nil {
		return PrimaryType("")
	}
	return o.PrimaryType
}

func (o *Attestation) GetDomain() Domain {
	if o == nil {
		return Domain{}
	}
	return o.Domain
}

func (o *Attestation) GetMessage() Message {
	if o == nil {
		return Message{}
	}
	return o.Message
}

func (o *Attestation) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *Attestation) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Attestation) GetSignatureType() *SignatureType {
	if o == nil {
		return nil
	}
	return o.SignatureType
}

func (o *Attestation) GetStorage() *AttestationStorage {
	if o == nil {
		return nil
	}
	return o.Storage
}
