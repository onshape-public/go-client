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

// GBTInsertableType the model 'GBTInsertableType'
type GBTInsertableType string

// List of GBTInsertableType
const (
	GBTInsertableTypeParts                 GBTInsertableType = "PARTS"
	GBTInsertableTypeSketches              GBTInsertableType = "SKETCHES"
	GBTInsertableTypeSurfaces              GBTInsertableType = "SURFACES"
	GBTInsertableTypeFlattenedParts        GBTInsertableType = "FLATTENED_PARTS"
	GBTInsertableTypeCompositeParts        GBTInsertableType = "COMPOSITE_PARTS"
	GBTInsertableTypePartStudios           GBTInsertableType = "PART_STUDIOS"
	GBTInsertableTypeWires                 GBTInsertableType = "WIRES"
	GBTInsertableTypeParametricPartStudios GBTInsertableType = "PARAMETRIC_PART_STUDIOS"
	GBTInsertableTypeUnknown               GBTInsertableType = "UNKNOWN"
)

// All allowed values of GBTInsertableType enum
var AllowedGBTInsertableTypeEnumValues = []GBTInsertableType{
	"PARTS",
	"SKETCHES",
	"SURFACES",
	"FLATTENED_PARTS",
	"COMPOSITE_PARTS",
	"PART_STUDIOS",
	"WIRES",
	"PARAMETRIC_PART_STUDIOS",
	"UNKNOWN",
}

func (v *GBTInsertableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTInsertableType(value)
	for _, existing := range AllowedGBTInsertableTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTInsertableType", value)
}

// NewGBTInsertableTypeFromValue returns a pointer to a valid GBTInsertableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTInsertableTypeFromValue(v string) (*GBTInsertableType, error) {
	ev := GBTInsertableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTInsertableType: valid values are %v", v, AllowedGBTInsertableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTInsertableType) IsValid() bool {
	for _, existing := range AllowedGBTInsertableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTInsertableType value
func (v GBTInsertableType) Ptr() *GBTInsertableType {
	return &v
}

type NullableGBTInsertableType struct {
	value *GBTInsertableType
	isSet bool
}

func (v NullableGBTInsertableType) Get() *GBTInsertableType {
	return v.value
}

func (v *NullableGBTInsertableType) Set(val *GBTInsertableType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTInsertableType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTInsertableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTInsertableType(val *GBTInsertableType) *NullableGBTInsertableType {
	return &NullableGBTInsertableType{value: val, isSet: true}
}

func (v NullableGBTInsertableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTInsertableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
