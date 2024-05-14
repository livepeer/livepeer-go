// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
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

	return errors.New("could not unmarshal into supported union types")
}

func (u UserTags3) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UserTagsType string

const (
	UserTagsTypeStr              UserTagsType = "str"
	UserTagsTypeNumber           UserTagsType = "number"
	UserTagsTypeArrayOfuserTags3 UserTagsType = "arrayOfuserTags_3"
)

type UserTags struct {
	Str              *string
	Number           *float64
	ArrayOfuserTags3 []UserTags3

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

func CreateUserTagsArrayOfuserTags3(arrayOfuserTags3 []UserTags3) UserTags {
	typ := UserTagsTypeArrayOfuserTags3

	return UserTags{
		ArrayOfuserTags3: arrayOfuserTags3,
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

	var arrayOfuserTags3 []UserTags3 = []UserTags3{}
	if err := utils.UnmarshalJSON(data, &arrayOfuserTags3, "", true, true); err == nil {
		u.ArrayOfuserTags3 = arrayOfuserTags3
		u.Type = UserTagsTypeArrayOfuserTags3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UserTags) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.ArrayOfuserTags3 != nil {
		return utils.MarshalJSON(u.ArrayOfuserTags3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
