/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
)

// BTUserAppSettingsInfo struct for BTUserAppSettingsInfo
type BTUserAppSettingsInfo struct {
	Settings *[]BTSettingInfo `json:"settings,omitempty"`
}

// NewBTUserAppSettingsInfo instantiates a new BTUserAppSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserAppSettingsInfo() *BTUserAppSettingsInfo {
	this := BTUserAppSettingsInfo{}
	return &this
}

// NewBTUserAppSettingsInfoWithDefaults instantiates a new BTUserAppSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserAppSettingsInfoWithDefaults() *BTUserAppSettingsInfo {
	this := BTUserAppSettingsInfo{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BTUserAppSettingsInfo) GetSettings() []BTSettingInfo {
	if o == nil || o.Settings == nil {
		var ret []BTSettingInfo
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppSettingsInfo) GetSettingsOk() (*[]BTSettingInfo, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BTUserAppSettingsInfo) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []BTSettingInfo and assigns it to the Settings field.
func (o *BTUserAppSettingsInfo) SetSettings(v []BTSettingInfo) {
	o.Settings = &v
}

func (o BTUserAppSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserAppSettingsInfo struct {
	value *BTUserAppSettingsInfo
	isSet bool
}

func (v NullableBTUserAppSettingsInfo) Get() *BTUserAppSettingsInfo {
	return v.value
}

func (v *NullableBTUserAppSettingsInfo) Set(val *BTUserAppSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserAppSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserAppSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserAppSettingsInfo(val *BTUserAppSettingsInfo) *NullableBTUserAppSettingsInfo {
	return &NullableBTUserAppSettingsInfo{value: val, isSet: true}
}

func (v NullableBTUserAppSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserAppSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
