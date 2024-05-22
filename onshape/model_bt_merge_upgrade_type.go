/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMergeUpgradeType the model 'BTMergeUpgradeType'
type BTMergeUpgradeType string

// List of BTMergeUpgradeType
const (
	BTMergeUpgradeTypeSource          BTMergeUpgradeType = "SOURCE"
	BTMergeUpgradeTypeTarget          BTMergeUpgradeType = "TARGET"
	BTMergeUpgradeTypeSourceAndTarget BTMergeUpgradeType = "SOURCE_AND_TARGET"
)

// All allowed values of BTMergeUpgradeType enum
var AllowedBTMergeUpgradeTypeEnumValues = []BTMergeUpgradeType{
	"SOURCE",
	"TARGET",
	"SOURCE_AND_TARGET",
}

func (v *BTMergeUpgradeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTMergeUpgradeType(value)
	for _, existing := range AllowedBTMergeUpgradeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTMergeUpgradeType", value)
}

// NewBTMergeUpgradeTypeFromValue returns a pointer to a valid BTMergeUpgradeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTMergeUpgradeTypeFromValue(v string) (*BTMergeUpgradeType, error) {
	ev := BTMergeUpgradeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTMergeUpgradeType: valid values are %v", v, AllowedBTMergeUpgradeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTMergeUpgradeType) IsValid() bool {
	for _, existing := range AllowedBTMergeUpgradeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTMergeUpgradeType value
func (v BTMergeUpgradeType) Ptr() *BTMergeUpgradeType {
	return &v
}

type NullableBTMergeUpgradeType struct {
	value *BTMergeUpgradeType
	isSet bool
}

func (v NullableBTMergeUpgradeType) Get() *BTMergeUpgradeType {
	return v.value
}

func (v *NullableBTMergeUpgradeType) Set(val *BTMergeUpgradeType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMergeUpgradeType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMergeUpgradeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMergeUpgradeType(val *BTMergeUpgradeType) *NullableBTMergeUpgradeType {
	return &NullableBTMergeUpgradeType{value: val, isSet: true}
}

func (v NullableBTMergeUpgradeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMergeUpgradeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}