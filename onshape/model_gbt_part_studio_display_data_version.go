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

// GBTPartStudioDisplayDataVersion the model 'GBTPartStudioDisplayDataVersion'
type GBTPartStudioDisplayDataVersion string

// List of GBTPartStudioDisplayDataVersion
const (
	GBTPartStudioDisplayDataVersionV0OriginalVersion             GBTPartStudioDisplayDataVersion = "V0_ORIGINAL_VERSION"
	GBTPartStudioDisplayDataVersionV1SmoothEdgesRenderingOptions GBTPartStudioDisplayDataVersion = "V1_SMOOTH_EDGES_RENDERING_OPTIONS"
	GBTPartStudioDisplayDataVersionV2SmoothEdgesToleranceChanged GBTPartStudioDisplayDataVersion = "V2_SMOOTH_EDGES_TOLERANCE_CHANGED"
	GBTPartStudioDisplayDataVersionUnknown                       GBTPartStudioDisplayDataVersion = "UNKNOWN"
)

// All allowed values of GBTPartStudioDisplayDataVersion enum
var AllowedGBTPartStudioDisplayDataVersionEnumValues = []GBTPartStudioDisplayDataVersion{
	"V0_ORIGINAL_VERSION",
	"V1_SMOOTH_EDGES_RENDERING_OPTIONS",
	"V2_SMOOTH_EDGES_TOLERANCE_CHANGED",
	"UNKNOWN",
}

func (v *GBTPartStudioDisplayDataVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPartStudioDisplayDataVersion(value)
	for _, existing := range AllowedGBTPartStudioDisplayDataVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPartStudioDisplayDataVersion", value)
}

// NewGBTPartStudioDisplayDataVersionFromValue returns a pointer to a valid GBTPartStudioDisplayDataVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPartStudioDisplayDataVersionFromValue(v string) (*GBTPartStudioDisplayDataVersion, error) {
	ev := GBTPartStudioDisplayDataVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPartStudioDisplayDataVersion: valid values are %v", v, AllowedGBTPartStudioDisplayDataVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPartStudioDisplayDataVersion) IsValid() bool {
	for _, existing := range AllowedGBTPartStudioDisplayDataVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPartStudioDisplayDataVersion value
func (v GBTPartStudioDisplayDataVersion) Ptr() *GBTPartStudioDisplayDataVersion {
	return &v
}

type NullableGBTPartStudioDisplayDataVersion struct {
	value *GBTPartStudioDisplayDataVersion
	isSet bool
}

func (v NullableGBTPartStudioDisplayDataVersion) Get() *GBTPartStudioDisplayDataVersion {
	return v.value
}

func (v *NullableGBTPartStudioDisplayDataVersion) Set(val *GBTPartStudioDisplayDataVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPartStudioDisplayDataVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPartStudioDisplayDataVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPartStudioDisplayDataVersion(val *GBTPartStudioDisplayDataVersion) *NullableGBTPartStudioDisplayDataVersion {
	return &NullableGBTPartStudioDisplayDataVersion{value: val, isSet: true}
}

func (v NullableGBTPartStudioDisplayDataVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPartStudioDisplayDataVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
