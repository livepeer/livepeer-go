// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type CreatorIDType string

const (
	CreatorIDTypeUnverified CreatorIDType = "unverified"
)

func (e CreatorIDType) ToPointer() *CreatorIDType {
	return &e
}
func (e *CreatorIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unverified":
		*e = CreatorIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatorIDType: %v", v)
	}
}

type CreatorID1 struct {
	Type CreatorIDType `json:"type"`
	// Developer-managed ID of the user who created the resource.
	Value string `json:"value"`
}

func (o *CreatorID1) GetType() CreatorIDType {
	if o == nil {
		return CreatorIDType("")
	}
	return o.Type
}

func (o *CreatorID1) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type CreatorIDUnionType string

const (
	CreatorIDUnionTypeCreatorID1 CreatorIDUnionType = "creator-id_1"
)

type CreatorID struct {
	CreatorID1 *CreatorID1

	Type CreatorIDUnionType
}

func CreateCreatorIDCreatorID1(creatorID1 CreatorID1) CreatorID {
	typ := CreatorIDUnionTypeCreatorID1

	return CreatorID{
		CreatorID1: &creatorID1,
		Type:       typ,
	}
}

func (u *CreatorID) UnmarshalJSON(data []byte) error {

	var creatorID1 CreatorID1 = CreatorID1{}
	if err := utils.UnmarshalJSON(data, &creatorID1, "", true, true); err == nil {
		u.CreatorID1 = &creatorID1
		u.Type = CreatorIDUnionTypeCreatorID1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreatorID", string(data))
}

func (u CreatorID) MarshalJSON() ([]byte, error) {
	if u.CreatorID1 != nil {
		return utils.MarshalJSON(u.CreatorID1, "", true)
	}

	return nil, errors.New("could not marshal union type CreatorID: all fields are null")
}
