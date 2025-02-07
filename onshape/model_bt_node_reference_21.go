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

// BTNodeReference21 struct for BTNodeReference21
type BTNodeReference21 struct {
	// Type of JSON object.
	BtType    *string     `json:"btType,omitempty"`
	NodeId    *string     `json:"nodeId,omitempty"`
	NodeIdRaw *BTObjectId `json:"nodeIdRaw,omitempty"`
}

// NewBTNodeReference21 instantiates a new BTNodeReference21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNodeReference21() *BTNodeReference21 {
	this := BTNodeReference21{}
	return &this
}

// NewBTNodeReference21WithDefaults instantiates a new BTNodeReference21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNodeReference21WithDefaults() *BTNodeReference21 {
	this := BTNodeReference21{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNodeReference21) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNodeReference21) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNodeReference21) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNodeReference21) SetBtType(v string) {
	o.BtType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTNodeReference21) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNodeReference21) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTNodeReference21) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTNodeReference21) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeIdRaw returns the NodeIdRaw field value if set, zero value otherwise.
func (o *BTNodeReference21) GetNodeIdRaw() BTObjectId {
	if o == nil || o.NodeIdRaw == nil {
		var ret BTObjectId
		return ret
	}
	return *o.NodeIdRaw
}

// GetNodeIdRawOk returns a tuple with the NodeIdRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNodeReference21) GetNodeIdRawOk() (*BTObjectId, bool) {
	if o == nil || o.NodeIdRaw == nil {
		return nil, false
	}
	return o.NodeIdRaw, true
}

// HasNodeIdRaw returns a boolean if a field has been set.
func (o *BTNodeReference21) HasNodeIdRaw() bool {
	if o != nil && o.NodeIdRaw != nil {
		return true
	}

	return false
}

// SetNodeIdRaw gets a reference to the given BTObjectId and assigns it to the NodeIdRaw field.
func (o *BTNodeReference21) SetNodeIdRaw(v BTObjectId) {
	o.NodeIdRaw = &v
}

func (o BTNodeReference21) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.NodeIdRaw != nil {
		toSerialize["nodeIdRaw"] = o.NodeIdRaw
	}
	return json.Marshal(toSerialize)
}

type NullableBTNodeReference21 struct {
	value *BTNodeReference21
	isSet bool
}

func (v NullableBTNodeReference21) Get() *BTNodeReference21 {
	return v.value
}

func (v *NullableBTNodeReference21) Set(val *BTNodeReference21) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNodeReference21) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNodeReference21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNodeReference21(val *BTNodeReference21) *NullableBTNodeReference21 {
	return &NullableBTNodeReference21{value: val, isSet: true}
}

func (v NullableBTNodeReference21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNodeReference21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
