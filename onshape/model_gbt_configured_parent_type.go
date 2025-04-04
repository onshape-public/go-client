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

// GBTConfiguredParentType the model 'GBTConfiguredParentType'
type GBTConfiguredParentType string

// List of GBTConfiguredParentType
const (
	GBTConfiguredParentTypeFeature       GBTConfiguredParentType = "FEATURE"
	GBTConfiguredParentTypeInstance      GBTConfiguredParentType = "INSTANCE"
	GBTConfiguredParentTypeMate          GBTConfiguredParentType = "MATE"
	GBTConfiguredParentTypeMateConnector GBTConfiguredParentType = "MATE_CONNECTOR"
	GBTConfiguredParentTypeUnknown       GBTConfiguredParentType = "UNKNOWN"
)

// All allowed values of GBTConfiguredParentType enum
var AllowedGBTConfiguredParentTypeEnumValues = []GBTConfiguredParentType{
	"FEATURE",
	"INSTANCE",
	"MATE",
	"MATE_CONNECTOR",
	"UNKNOWN",
}

func (v *GBTConfiguredParentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTConfiguredParentType(value)
	for _, existing := range AllowedGBTConfiguredParentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTConfiguredParentType", value)
}

// NewGBTConfiguredParentTypeFromValue returns a pointer to a valid GBTConfiguredParentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTConfiguredParentTypeFromValue(v string) (*GBTConfiguredParentType, error) {
	ev := GBTConfiguredParentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTConfiguredParentType: valid values are %v", v, AllowedGBTConfiguredParentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTConfiguredParentType) IsValid() bool {
	for _, existing := range AllowedGBTConfiguredParentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTConfiguredParentType value
func (v GBTConfiguredParentType) Ptr() *GBTConfiguredParentType {
	return &v
}

type NullableGBTConfiguredParentType struct {
	value *GBTConfiguredParentType
	isSet bool
}

func (v NullableGBTConfiguredParentType) Get() *GBTConfiguredParentType {
	return v.value
}

func (v *NullableGBTConfiguredParentType) Set(val *GBTConfiguredParentType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTConfiguredParentType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTConfiguredParentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTConfiguredParentType(val *GBTConfiguredParentType) *NullableGBTConfiguredParentType {
	return &NullableGBTConfiguredParentType{value: val, isSet: true}
}

func (v NullableGBTConfiguredParentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTConfiguredParentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
