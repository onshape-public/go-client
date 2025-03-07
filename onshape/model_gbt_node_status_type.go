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

// GBTNodeStatusType the model 'GBTNodeStatusType'
type GBTNodeStatusType string

// List of GBTNodeStatusType
const (
	GBTNodeStatusTypeOk      GBTNodeStatusType = "OK"
	GBTNodeStatusTypeInfo    GBTNodeStatusType = "INFO"
	GBTNodeStatusTypeWarning GBTNodeStatusType = "WARNING"
	GBTNodeStatusTypeError   GBTNodeStatusType = "ERROR"
	GBTNodeStatusTypeUnknown GBTNodeStatusType = "UNKNOWN"
)

// All allowed values of GBTNodeStatusType enum
var AllowedGBTNodeStatusTypeEnumValues = []GBTNodeStatusType{
	"OK",
	"INFO",
	"WARNING",
	"ERROR",
	"UNKNOWN",
}

func (v *GBTNodeStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTNodeStatusType(value)
	for _, existing := range AllowedGBTNodeStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTNodeStatusType", value)
}

// NewGBTNodeStatusTypeFromValue returns a pointer to a valid GBTNodeStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTNodeStatusTypeFromValue(v string) (*GBTNodeStatusType, error) {
	ev := GBTNodeStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTNodeStatusType: valid values are %v", v, AllowedGBTNodeStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTNodeStatusType) IsValid() bool {
	for _, existing := range AllowedGBTNodeStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTNodeStatusType value
func (v GBTNodeStatusType) Ptr() *GBTNodeStatusType {
	return &v
}

type NullableGBTNodeStatusType struct {
	value *GBTNodeStatusType
	isSet bool
}

func (v NullableGBTNodeStatusType) Get() *GBTNodeStatusType {
	return v.value
}

func (v *NullableGBTNodeStatusType) Set(val *GBTNodeStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTNodeStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTNodeStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTNodeStatusType(val *GBTNodeStatusType) *NullableGBTNodeStatusType {
	return &NullableGBTNodeStatusType{value: val, isSet: true}
}

func (v NullableGBTNodeStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTNodeStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
