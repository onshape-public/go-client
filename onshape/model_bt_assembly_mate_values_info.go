/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyMateValuesInfo Get a list of mate values for all the mates of a specified assembly.
type BTAssemblyMateValuesInfo struct {
	// Describes the relative position of the first mate connector with respect to the second along the designated degrees of freedom (DOF) for mates in the specified assembly.
	MateValues []BTAssemblyMateValueInfo `json:"mateValues,omitempty"`
}

// NewBTAssemblyMateValuesInfo instantiates a new BTAssemblyMateValuesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyMateValuesInfo() *BTAssemblyMateValuesInfo {
	this := BTAssemblyMateValuesInfo{}
	return &this
}

// NewBTAssemblyMateValuesInfoWithDefaults instantiates a new BTAssemblyMateValuesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyMateValuesInfoWithDefaults() *BTAssemblyMateValuesInfo {
	this := BTAssemblyMateValuesInfo{}
	return &this
}

// GetMateValues returns the MateValues field value if set, zero value otherwise.
func (o *BTAssemblyMateValuesInfo) GetMateValues() []BTAssemblyMateValueInfo {
	if o == nil || o.MateValues == nil {
		var ret []BTAssemblyMateValueInfo
		return ret
	}
	return o.MateValues
}

// GetMateValuesOk returns a tuple with the MateValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyMateValuesInfo) GetMateValuesOk() ([]BTAssemblyMateValueInfo, bool) {
	if o == nil || o.MateValues == nil {
		return nil, false
	}
	return o.MateValues, true
}

// HasMateValues returns a boolean if a field has been set.
func (o *BTAssemblyMateValuesInfo) HasMateValues() bool {
	if o != nil && o.MateValues != nil {
		return true
	}

	return false
}

// SetMateValues gets a reference to the given []BTAssemblyMateValueInfo and assigns it to the MateValues field.
func (o *BTAssemblyMateValuesInfo) SetMateValues(v []BTAssemblyMateValueInfo) {
	o.MateValues = v
}

func (o BTAssemblyMateValuesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MateValues != nil {
		toSerialize["mateValues"] = o.MateValues
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyMateValuesInfo struct {
	value *BTAssemblyMateValuesInfo
	isSet bool
}

func (v NullableBTAssemblyMateValuesInfo) Get() *BTAssemblyMateValuesInfo {
	return v.value
}

func (v *NullableBTAssemblyMateValuesInfo) Set(val *BTAssemblyMateValuesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyMateValuesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyMateValuesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyMateValuesInfo(val *BTAssemblyMateValuesInfo) *NullableBTAssemblyMateValuesInfo {
	return &NullableBTAssemblyMateValuesInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyMateValuesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyMateValuesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
