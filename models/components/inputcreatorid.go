// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type InputCreatorIDType string

const (
	InputCreatorIDTypeCreatorID InputCreatorIDType = "creator-id"
	InputCreatorIDTypeStr       InputCreatorIDType = "str"
)

type InputCreatorID struct {
	CreatorID *CreatorID
	Str       *string

	Type InputCreatorIDType
}

func CreateInputCreatorIDCreatorID(creatorID CreatorID) InputCreatorID {
	typ := InputCreatorIDTypeCreatorID

	return InputCreatorID{
		CreatorID: &creatorID,
		Type:      typ,
	}
}

func CreateInputCreatorIDStr(str string) InputCreatorID {
	typ := InputCreatorIDTypeStr

	return InputCreatorID{
		Str:  &str,
		Type: typ,
	}
}

func (u *InputCreatorID) UnmarshalJSON(data []byte) error {

	creatorID := CreatorID{}
	if err := utils.UnmarshalJSON(data, &creatorID, "", true, true); err == nil {
		u.CreatorID = &creatorID
		u.Type = InputCreatorIDTypeCreatorID
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = InputCreatorIDTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u InputCreatorID) MarshalJSON() ([]byte, error) {
	if u.CreatorID != nil {
		return utils.MarshalJSON(u.CreatorID, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
