/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// EditType the model 'EditType'
type EditType string

// List of EditType
const (
	EditTypeNothing     EditType = "NOTHING"
	EditTypeNewRoot     EditType = "NEW_ROOT"
	EditTypeMove        EditType = "MOVE"
	EditTypeChange      EditType = "CHANGE"
	EditTypeChangeField EditType = "CHANGE_FIELD"
	EditTypeInsertion   EditType = "INSERTION"
	EditTypeDeletion    EditType = "DELETION"
	EditTypeList        EditType = "LIST"
)

// All allowed values of EditType enum
var AllowedEditTypeEnumValues = []EditType{
	"NOTHING",
	"NEW_ROOT",
	"MOVE",
	"CHANGE",
	"CHANGE_FIELD",
	"INSERTION",
	"DELETION",
	"LIST",
}

func (v *EditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EditType(value)
	for _, existing := range AllowedEditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EditType", value)
}

// NewEditTypeFromValue returns a pointer to a valid EditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEditTypeFromValue(v string) (*EditType, error) {
	ev := EditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EditType: valid values are %v", v, AllowedEditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EditType) IsValid() bool {
	for _, existing := range AllowedEditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EditType value
func (v EditType) Ptr() *EditType {
	return &v
}

type NullableEditType struct {
	value *EditType
	isSet bool
}

func (v NullableEditType) Get() *EditType {
	return v.value
}

func (v *NullableEditType) Set(val *EditType) {
	v.value = val
	v.isSet = true
}

func (v NullableEditType) IsSet() bool {
	return v.isSet
}

func (v *NullableEditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditType(val *EditType) *NullableEditType {
	return &NullableEditType{value: val, isSet: true}
}

func (v NullableEditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
