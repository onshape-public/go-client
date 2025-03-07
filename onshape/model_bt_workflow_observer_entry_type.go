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

// BTWorkflowObserverEntryType the model 'BTWorkflowObserverEntryType'
type BTWorkflowObserverEntryType string

// List of BTWorkflowObserverEntryType
const (
	BTWorkflowObserverEntryTypeUser  BTWorkflowObserverEntryType = "USER"
	BTWorkflowObserverEntryTypeTeam  BTWorkflowObserverEntryType = "TEAM"
	BTWorkflowObserverEntryTypeRole  BTWorkflowObserverEntryType = "ROLE"
	BTWorkflowObserverEntryTypeAlias BTWorkflowObserverEntryType = "ALIAS"
)

// All allowed values of BTWorkflowObserverEntryType enum
var AllowedBTWorkflowObserverEntryTypeEnumValues = []BTWorkflowObserverEntryType{
	"USER",
	"TEAM",
	"ROLE",
	"ALIAS",
}

func (v *BTWorkflowObserverEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTWorkflowObserverEntryType(value)
	for _, existing := range AllowedBTWorkflowObserverEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTWorkflowObserverEntryType", value)
}

// NewBTWorkflowObserverEntryTypeFromValue returns a pointer to a valid BTWorkflowObserverEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTWorkflowObserverEntryTypeFromValue(v string) (*BTWorkflowObserverEntryType, error) {
	ev := BTWorkflowObserverEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTWorkflowObserverEntryType: valid values are %v", v, AllowedBTWorkflowObserverEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTWorkflowObserverEntryType) IsValid() bool {
	for _, existing := range AllowedBTWorkflowObserverEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTWorkflowObserverEntryType value
func (v BTWorkflowObserverEntryType) Ptr() *BTWorkflowObserverEntryType {
	return &v
}

type NullableBTWorkflowObserverEntryType struct {
	value *BTWorkflowObserverEntryType
	isSet bool
}

func (v NullableBTWorkflowObserverEntryType) Get() *BTWorkflowObserverEntryType {
	return v.value
}

func (v *NullableBTWorkflowObserverEntryType) Set(val *BTWorkflowObserverEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowObserverEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowObserverEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowObserverEntryType(val *BTWorkflowObserverEntryType) *NullableBTWorkflowObserverEntryType {
	return &NullableBTWorkflowObserverEntryType{value: val, isSet: true}
}

func (v NullableBTWorkflowObserverEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowObserverEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
