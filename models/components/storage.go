// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type One struct {
	Spec *Spec `json:"spec,omitempty"`
}

func (o *One) GetSpec() *Spec {
	if o == nil {
		return nil
	}
	return o.Spec
}

type IpfsType string

const (
	IpfsTypeOne     IpfsType = "1"
	IpfsTypeBoolean IpfsType = "boolean"
)

// Ipfs - Set to true to make default export to IPFS. To customize the
// pinned files, specify an object with a spec field. False or null
// means to unpin from IPFS, but it's unsupported right now.
type Ipfs struct {
	One     *One
	Boolean *bool

	Type IpfsType
}

func CreateIpfsOne(one One) Ipfs {
	typ := IpfsTypeOne

	return Ipfs{
		One:  &one,
		Type: typ,
	}
}

func CreateIpfsBoolean(boolean bool) Ipfs {
	typ := IpfsTypeBoolean

	return Ipfs{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *Ipfs) UnmarshalJSON(data []byte) error {

	one := One{}
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = IpfsTypeOne
		return nil
	}

	boolean := false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = IpfsTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Ipfs) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Storage struct {
	// Set to true to make default export to IPFS. To customize the
	// pinned files, specify an object with a spec field. False or null
	// means to unpin from IPFS, but it's unsupported right now.
	//
	Ipfs *Ipfs `json:"ipfs,omitempty"`
}

func (o *Storage) GetIpfs() *Ipfs {
	if o == nil {
		return nil
	}
	return o.Ipfs
}
