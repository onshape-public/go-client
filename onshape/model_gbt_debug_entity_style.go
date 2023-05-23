/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16301-d273853a12e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTDebugEntityStyle the model 'GBTDebugEntityStyle'
type GBTDebugEntityStyle string

// List of GBTDebugEntityStyle
const (
	GBTDebugEntityStyleDefault   GBTDebugEntityStyle = "DEFAULT"
	GBTDebugEntityStyleStar      GBTDebugEntityStyle = "STAR"
	GBTDebugEntityStyleWithStart GBTDebugEntityStyle = "WITH_START"
	GBTDebugEntityStyleError     GBTDebugEntityStyle = "ERROR"
	GBTDebugEntityStyleUnknown   GBTDebugEntityStyle = "UNKNOWN"
)

// All allowed values of GBTDebugEntityStyle enum
var AllowedGBTDebugEntityStyleEnumValues = []GBTDebugEntityStyle{
	"DEFAULT",
	"STAR",
	"WITH_START",
	"ERROR",
	"UNKNOWN",
}

func (v *GBTDebugEntityStyle) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTDebugEntityStyle(value)
	for _, existing := range AllowedGBTDebugEntityStyleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTDebugEntityStyle", value)
}

// NewGBTDebugEntityStyleFromValue returns a pointer to a valid GBTDebugEntityStyle
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTDebugEntityStyleFromValue(v string) (*GBTDebugEntityStyle, error) {
	ev := GBTDebugEntityStyle(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTDebugEntityStyle: valid values are %v", v, AllowedGBTDebugEntityStyleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTDebugEntityStyle) IsValid() bool {
	for _, existing := range AllowedGBTDebugEntityStyleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTDebugEntityStyle value
func (v GBTDebugEntityStyle) Ptr() *GBTDebugEntityStyle {
	return &v
}

type NullableGBTDebugEntityStyle struct {
	value *GBTDebugEntityStyle
	isSet bool
}

func (v NullableGBTDebugEntityStyle) Get() *GBTDebugEntityStyle {
	return v.value
}

func (v *NullableGBTDebugEntityStyle) Set(val *GBTDebugEntityStyle) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTDebugEntityStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTDebugEntityStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTDebugEntityStyle(val *GBTDebugEntityStyle) *NullableGBTDebugEntityStyle {
	return &NullableGBTDebugEntityStyle{value: val, isSet: true}
}

func (v NullableGBTDebugEntityStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTDebugEntityStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}