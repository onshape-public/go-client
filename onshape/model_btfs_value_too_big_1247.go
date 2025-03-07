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

// BTFSValueTooBig1247 struct for BTFSValueTooBig1247
type BTFSValueTooBig1247 struct {
	BTFSValue1888
	BtType  string  `json:"btType"`
	TypeTag *string `json:"typeTag,omitempty"`
}

// NewBTFSValueTooBig1247 instantiates a new BTFSValueTooBig1247 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueTooBig1247(btType string) *BTFSValueTooBig1247 {
	this := BTFSValueTooBig1247{}
	this.BtType = btType
	return &this
}

// NewBTFSValueTooBig1247WithDefaults instantiates a new BTFSValueTooBig1247 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueTooBig1247WithDefaults() *BTFSValueTooBig1247 {
	this := BTFSValueTooBig1247{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueTooBig1247) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueTooBig1247) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueTooBig1247) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueTooBig1247) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueTooBig1247) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueTooBig1247) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueTooBig1247) SetTypeTag(v string) {
	o.TypeTag = &v
}

func (o BTFSValueTooBig1247) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTFSValue1888, errBTFSValue1888 := json.Marshal(o.BTFSValue1888)
	if errBTFSValue1888 != nil {
		return []byte{}, errBTFSValue1888
	}
	errBTFSValue1888 = json.Unmarshal([]byte(serializedBTFSValue1888), &toSerialize)
	if errBTFSValue1888 != nil {
		return []byte{}, errBTFSValue1888
	}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueTooBig1247 struct {
	value *BTFSValueTooBig1247
	isSet bool
}

func (v NullableBTFSValueTooBig1247) Get() *BTFSValueTooBig1247 {
	return v.value
}

func (v *NullableBTFSValueTooBig1247) Set(val *BTFSValueTooBig1247) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueTooBig1247) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueTooBig1247) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueTooBig1247(val *BTFSValueTooBig1247) *NullableBTFSValueTooBig1247 {
	return &NullableBTFSValueTooBig1247{value: val, isSet: true}
}

func (v NullableBTFSValueTooBig1247) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueTooBig1247) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
