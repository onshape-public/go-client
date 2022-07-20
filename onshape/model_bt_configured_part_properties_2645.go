/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConfiguredPartProperties2645 struct for BTConfiguredPartProperties2645
type BTConfiguredPartProperties2645 struct {
	BtType                       *string                                        `json:"btType,omitempty"`
	NodeId                       *string                                        `json:"nodeId,omitempty"`
	Parts                        []BTPartWithConfiguredProperties2163           `json:"parts,omitempty"`
	PropertyIdToConfiguredTable  *map[string]BTPartWithConfiguredProperties2163 `json:"propertyIdToConfiguredTable,omitempty"`
	SynchronizeToSingleEnumInput *bool                                          `json:"synchronizeToSingleEnumInput,omitempty"`
}

// NewBTConfiguredPartProperties2645 instantiates a new BTConfiguredPartProperties2645 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredPartProperties2645() *BTConfiguredPartProperties2645 {
	this := BTConfiguredPartProperties2645{}
	return &this
}

// NewBTConfiguredPartProperties2645WithDefaults instantiates a new BTConfiguredPartProperties2645 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredPartProperties2645WithDefaults() *BTConfiguredPartProperties2645 {
	this := BTConfiguredPartProperties2645{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredPartProperties2645) SetBtType(v string) {
	o.BtType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTConfiguredPartProperties2645) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetParts() []BTPartWithConfiguredProperties2163 {
	if o == nil || o.Parts == nil {
		var ret []BTPartWithConfiguredProperties2163
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetPartsOk() ([]BTPartWithConfiguredProperties2163, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasParts() bool {
	if o != nil && o.Parts != nil {
		return true
	}

	return false
}

// SetParts gets a reference to the given []BTPartWithConfiguredProperties2163 and assigns it to the Parts field.
func (o *BTConfiguredPartProperties2645) SetParts(v []BTPartWithConfiguredProperties2163) {
	o.Parts = v
}

// GetPropertyIdToConfiguredTable returns the PropertyIdToConfiguredTable field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetPropertyIdToConfiguredTable() map[string]BTPartWithConfiguredProperties2163 {
	if o == nil || o.PropertyIdToConfiguredTable == nil {
		var ret map[string]BTPartWithConfiguredProperties2163
		return ret
	}
	return *o.PropertyIdToConfiguredTable
}

// GetPropertyIdToConfiguredTableOk returns a tuple with the PropertyIdToConfiguredTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetPropertyIdToConfiguredTableOk() (*map[string]BTPartWithConfiguredProperties2163, bool) {
	if o == nil || o.PropertyIdToConfiguredTable == nil {
		return nil, false
	}
	return o.PropertyIdToConfiguredTable, true
}

// HasPropertyIdToConfiguredTable returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasPropertyIdToConfiguredTable() bool {
	if o != nil && o.PropertyIdToConfiguredTable != nil {
		return true
	}

	return false
}

// SetPropertyIdToConfiguredTable gets a reference to the given map[string]BTPartWithConfiguredProperties2163 and assigns it to the PropertyIdToConfiguredTable field.
func (o *BTConfiguredPartProperties2645) SetPropertyIdToConfiguredTable(v map[string]BTPartWithConfiguredProperties2163) {
	o.PropertyIdToConfiguredTable = &v
}

// GetSynchronizeToSingleEnumInput returns the SynchronizeToSingleEnumInput field value if set, zero value otherwise.
func (o *BTConfiguredPartProperties2645) GetSynchronizeToSingleEnumInput() bool {
	if o == nil || o.SynchronizeToSingleEnumInput == nil {
		var ret bool
		return ret
	}
	return *o.SynchronizeToSingleEnumInput
}

// GetSynchronizeToSingleEnumInputOk returns a tuple with the SynchronizeToSingleEnumInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartProperties2645) GetSynchronizeToSingleEnumInputOk() (*bool, bool) {
	if o == nil || o.SynchronizeToSingleEnumInput == nil {
		return nil, false
	}
	return o.SynchronizeToSingleEnumInput, true
}

// HasSynchronizeToSingleEnumInput returns a boolean if a field has been set.
func (o *BTConfiguredPartProperties2645) HasSynchronizeToSingleEnumInput() bool {
	if o != nil && o.SynchronizeToSingleEnumInput != nil {
		return true
	}

	return false
}

// SetSynchronizeToSingleEnumInput gets a reference to the given bool and assigns it to the SynchronizeToSingleEnumInput field.
func (o *BTConfiguredPartProperties2645) SetSynchronizeToSingleEnumInput(v bool) {
	o.SynchronizeToSingleEnumInput = &v
}

func (o BTConfiguredPartProperties2645) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
	}
	if o.PropertyIdToConfiguredTable != nil {
		toSerialize["propertyIdToConfiguredTable"] = o.PropertyIdToConfiguredTable
	}
	if o.SynchronizeToSingleEnumInput != nil {
		toSerialize["synchronizeToSingleEnumInput"] = o.SynchronizeToSingleEnumInput
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredPartProperties2645 struct {
	value *BTConfiguredPartProperties2645
	isSet bool
}

func (v NullableBTConfiguredPartProperties2645) Get() *BTConfiguredPartProperties2645 {
	return v.value
}

func (v *NullableBTConfiguredPartProperties2645) Set(val *BTConfiguredPartProperties2645) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredPartProperties2645) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredPartProperties2645) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredPartProperties2645(val *BTConfiguredPartProperties2645) *NullableBTConfiguredPartProperties2645 {
	return &NullableBTConfiguredPartProperties2645{value: val, isSet: true}
}

func (v NullableBTConfiguredPartProperties2645) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredPartProperties2645) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}