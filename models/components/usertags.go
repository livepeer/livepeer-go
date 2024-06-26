// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type UserTags3Type string

const (
	UserTags3TypeStr    UserTags3Type = "str"
	UserTags3TypeNumber UserTags3Type = "number"
)

type UserTags3 struct {
	Str    *string
	Number *float64

	Type UserTags3Type
}

func CreateUserTags3Str(str string) UserTags3 {
	typ := UserTags3TypeStr

	return UserTags3{
		Str:  &str,
		Type: typ,
	}
}

func CreateUserTags3Number(number float64) UserTags3 {
	typ := UserTags3TypeNumber

	return UserTags3{
		Number: &number,
		Type:   typ,
	}
}

func (u *UserTags3) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UserTags3TypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = UserTags3TypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UserTags3", string(data))
}

func (u UserTags3) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type UserTags3: all fields are null")
}

type UserTagsType string

const (
	UserTagsTypeStr              UserTagsType = "str"
	UserTagsTypeNumber           UserTagsType = "number"
	UserTagsTypeArrayOfUserTags3 UserTagsType = "arrayOfUserTags3"
)

type UserTags struct {
	Str              *string
	Number           *float64
	ArrayOfUserTags3 []UserTags3

	Type UserTagsType
}

func CreateUserTagsStr(str string) UserTags {
	typ := UserTagsTypeStr

	return UserTags{
		Str:  &str,
		Type: typ,
	}
}

func CreateUserTagsNumber(number float64) UserTags {
	typ := UserTagsTypeNumber

	return UserTags{
		Number: &number,
		Type:   typ,
	}
}

func CreateUserTagsArrayOfUserTags3(arrayOfUserTags3 []UserTags3) UserTags {
	typ := UserTagsTypeArrayOfUserTags3

	return UserTags{
		ArrayOfUserTags3: arrayOfUserTags3,
		Type:             typ,
	}
}

func (u *UserTags) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UserTagsTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = UserTagsTypeNumber
		return nil
	}

	var arrayOfUserTags3 []UserTags3 = []UserTags3{}
	if err := utils.UnmarshalJSON(data, &arrayOfUserTags3, "", true, true); err == nil {
		u.ArrayOfUserTags3 = arrayOfUserTags3
		u.Type = UserTagsTypeArrayOfUserTags3
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UserTags", string(data))
}

func (u UserTags) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.ArrayOfUserTags3 != nil {
		return utils.MarshalJSON(u.ArrayOfUserTags3, "", true)
	}

	return nil, errors.New("could not marshal union type UserTags: all fields are null")
}
