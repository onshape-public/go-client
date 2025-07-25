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

// BTMParameterBoolean144 struct for BTMParameterBoolean144
type BTMParameterBoolean144 struct {
	BTMParameter1
	BtType *string `json:"btType,omitempty"`
	// Microversion that resulted from the import.
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	// ID of the parameter's node.
	NodeId *string `json:"nodeId,omitempty"`
	// Unique ID of the parameter.
	ParameterId *string `json:"parameterId,omitempty"`
	ValueString *string `json:"valueString,omitempty"`
	Value       *bool   `json:"value,omitempty"`
}

// NewBTMParameterBoolean144 instantiates a new BTMParameterBoolean144 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterBoolean144() *BTMParameterBoolean144 {
	this := BTMParameterBoolean144{}
	return &this
}

// NewBTMParameterBoolean144WithDefaults instantiates a new BTMParameterBoolean144 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterBoolean144WithDefaults() *BTMParameterBoolean144 {
	this := BTMParameterBoolean144{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterBoolean144) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBoolean144) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterBoolean144) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterBoolean144) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterBoolean144) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBoolean144) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterBoolean144) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterBoolean144) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterBoolean144) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBoolean144) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterBoolean144) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterBoolean144) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterBoolean144) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBoolean144) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterBoolean144) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterBoolean144) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *BTMParameterBoolean144) GetValueString() string {
	if o == nil || o.ValueString == nil {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBoolean144) GetValueStringOk() (*string, bool) {
	if o == nil || o.ValueString == nil {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *BTMParameterBoolean144) HasValueString() bool {
	if o != nil && o.ValueString != nil {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *BTMParameterBoolean144) SetValueString(v string) {
	o.ValueString = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMParameterBoolean144) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBoolean144) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMParameterBoolean144) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *BTMParameterBoolean144) SetValue(v bool) {
	o.Value = &v
}

func (o BTMParameterBoolean144) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMParameter1, errBTMParameter1 := json.Marshal(o.BTMParameter1)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	errBTMParameter1 = json.Unmarshal([]byte(serializedBTMParameter1), &toSerialize)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ValueString != nil {
		toSerialize["valueString"] = o.ValueString
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterBoolean144 struct {
	value *BTMParameterBoolean144
	isSet bool
}

func (v NullableBTMParameterBoolean144) Get() *BTMParameterBoolean144 {
	return v.value
}

func (v *NullableBTMParameterBoolean144) Set(val *BTMParameterBoolean144) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterBoolean144) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterBoolean144) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterBoolean144(val *BTMParameterBoolean144) *NullableBTMParameterBoolean144 {
	return &NullableBTMParameterBoolean144{value: val, isSet: true}
}

func (v NullableBTMParameterBoolean144) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterBoolean144) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
