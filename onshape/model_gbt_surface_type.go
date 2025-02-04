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

// GBTSurfaceType the model 'GBTSurfaceType'
type GBTSurfaceType string

// List of GBTSurfaceType
const (
	GBTSurfaceTypePlane    GBTSurfaceType = "PLANE"
	GBTSurfaceTypeCylinder GBTSurfaceType = "CYLINDER"
	GBTSurfaceTypeCone     GBTSurfaceType = "CONE"
	GBTSurfaceTypeSphere   GBTSurfaceType = "SPHERE"
	GBTSurfaceTypeTorus    GBTSurfaceType = "TORUS"
	GBTSurfaceTypeOther    GBTSurfaceType = "OTHER"
	GBTSurfaceTypeRevolved GBTSurfaceType = "REVOLVED"
	GBTSurfaceTypeExtruded GBTSurfaceType = "EXTRUDED"
	GBTSurfaceTypeMesh     GBTSurfaceType = "MESH"
	GBTSurfaceTypeSpline   GBTSurfaceType = "SPLINE"
	GBTSurfaceTypeUnknown  GBTSurfaceType = "UNKNOWN"
)

// All allowed values of GBTSurfaceType enum
var AllowedGBTSurfaceTypeEnumValues = []GBTSurfaceType{
	"PLANE",
	"CYLINDER",
	"CONE",
	"SPHERE",
	"TORUS",
	"OTHER",
	"REVOLVED",
	"EXTRUDED",
	"MESH",
	"SPLINE",
	"UNKNOWN",
}

func (v *GBTSurfaceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTSurfaceType(value)
	for _, existing := range AllowedGBTSurfaceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTSurfaceType", value)
}

// NewGBTSurfaceTypeFromValue returns a pointer to a valid GBTSurfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTSurfaceTypeFromValue(v string) (*GBTSurfaceType, error) {
	ev := GBTSurfaceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTSurfaceType: valid values are %v", v, AllowedGBTSurfaceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTSurfaceType) IsValid() bool {
	for _, existing := range AllowedGBTSurfaceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTSurfaceType value
func (v GBTSurfaceType) Ptr() *GBTSurfaceType {
	return &v
}

type NullableGBTSurfaceType struct {
	value *GBTSurfaceType
	isSet bool
}

func (v NullableGBTSurfaceType) Get() *GBTSurfaceType {
	return v.value
}

func (v *NullableGBTSurfaceType) Set(val *GBTSurfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTSurfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTSurfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTSurfaceType(val *GBTSurfaceType) *NullableGBTSurfaceType {
	return &NullableGBTSurfaceType{value: val, isSet: true}
}

func (v NullableGBTSurfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTSurfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
