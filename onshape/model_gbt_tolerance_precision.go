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

// GBTTolerancePrecision the model 'GBTTolerancePrecision'
type GBTTolerancePrecision string

// List of GBTTolerancePrecision
const (
	GBTTolerancePrecisionDefault            GBTTolerancePrecision = "DEFAULT"
	GBTTolerancePrecisionOnes               GBTTolerancePrecision = "ONES"
	GBTTolerancePrecisionTenths             GBTTolerancePrecision = "TENTHS"
	GBTTolerancePrecisionHundredths         GBTTolerancePrecision = "HUNDREDTHS"
	GBTTolerancePrecisionThousandths        GBTTolerancePrecision = "THOUSANDTHS"
	GBTTolerancePrecisionTenThousandths     GBTTolerancePrecision = "TEN_THOUSANDTHS"
	GBTTolerancePrecisionHundredThousandths GBTTolerancePrecision = "HUNDRED_THOUSANDTHS"
	GBTTolerancePrecisionMillionths         GBTTolerancePrecision = "MILLIONTHS"
	GBTTolerancePrecisionUnknown            GBTTolerancePrecision = "UNKNOWN"
)

// All allowed values of GBTTolerancePrecision enum
var AllowedGBTTolerancePrecisionEnumValues = []GBTTolerancePrecision{
	"DEFAULT",
	"ONES",
	"TENTHS",
	"HUNDREDTHS",
	"THOUSANDTHS",
	"TEN_THOUSANDTHS",
	"HUNDRED_THOUSANDTHS",
	"MILLIONTHS",
	"UNKNOWN",
}

func (v *GBTTolerancePrecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTTolerancePrecision(value)
	for _, existing := range AllowedGBTTolerancePrecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTTolerancePrecision", value)
}

// NewGBTTolerancePrecisionFromValue returns a pointer to a valid GBTTolerancePrecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTTolerancePrecisionFromValue(v string) (*GBTTolerancePrecision, error) {
	ev := GBTTolerancePrecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTTolerancePrecision: valid values are %v", v, AllowedGBTTolerancePrecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTTolerancePrecision) IsValid() bool {
	for _, existing := range AllowedGBTTolerancePrecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTTolerancePrecision value
func (v GBTTolerancePrecision) Ptr() *GBTTolerancePrecision {
	return &v
}

type NullableGBTTolerancePrecision struct {
	value *GBTTolerancePrecision
	isSet bool
}

func (v NullableGBTTolerancePrecision) Get() *GBTTolerancePrecision {
	return v.value
}

func (v *NullableGBTTolerancePrecision) Set(val *GBTTolerancePrecision) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTTolerancePrecision) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTTolerancePrecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTTolerancePrecision(val *GBTTolerancePrecision) *NullableGBTTolerancePrecision {
	return &NullableGBTTolerancePrecision{value: val, isSet: true}
}

func (v NullableGBTTolerancePrecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTTolerancePrecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
