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

// BTFSValueBoolean1195 struct for BTFSValueBoolean1195
type BTFSValueBoolean1195 struct {
	BTFSValue1888
	BtType  string  `json:"btType"`
	TypeTag *string `json:"typeTag,omitempty"`
	Value   *bool   `json:"value,omitempty"`
}

// NewBTFSValueBoolean1195 instantiates a new BTFSValueBoolean1195 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueBoolean1195(btType string) *BTFSValueBoolean1195 {
	this := BTFSValueBoolean1195{}
	this.BtType = btType
	return &this
}

// NewBTFSValueBoolean1195WithDefaults instantiates a new BTFSValueBoolean1195 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueBoolean1195WithDefaults() *BTFSValueBoolean1195 {
	this := BTFSValueBoolean1195{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueBoolean1195) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueBoolean1195) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueBoolean1195) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueBoolean1195) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueBoolean1195) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueBoolean1195) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueBoolean1195) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueBoolean1195) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueBoolean1195) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueBoolean1195) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *BTFSValueBoolean1195) SetValue(v bool) {
	o.Value = &v
}

func (o BTFSValueBoolean1195) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueBoolean1195 struct {
	value *BTFSValueBoolean1195
	isSet bool
}

func (v NullableBTFSValueBoolean1195) Get() *BTFSValueBoolean1195 {
	return v.value
}

func (v *NullableBTFSValueBoolean1195) Set(val *BTFSValueBoolean1195) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueBoolean1195) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueBoolean1195) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueBoolean1195(val *BTFSValueBoolean1195) *NullableBTFSValueBoolean1195 {
	return &NullableBTFSValueBoolean1195{value: val, isSet: true}
}

func (v NullableBTFSValueBoolean1195) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueBoolean1195) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
