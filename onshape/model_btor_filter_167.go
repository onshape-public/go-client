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

// BTOrFilter167 struct for BTOrFilter167
type BTOrFilter167 struct {
	BTQueryFilter183
	BtType   *string           `json:"btType,omitempty"`
	Operand1 *BTQueryFilter183 `json:"operand1,omitempty"`
	Operand2 *BTQueryFilter183 `json:"operand2,omitempty"`
}

// NewBTOrFilter167 instantiates a new BTOrFilter167 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOrFilter167() *BTOrFilter167 {
	this := BTOrFilter167{}
	return &this
}

// NewBTOrFilter167WithDefaults instantiates a new BTOrFilter167 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOrFilter167WithDefaults() *BTOrFilter167 {
	this := BTOrFilter167{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOrFilter167) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOrFilter167) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOrFilter167) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOrFilter167) SetBtType(v string) {
	o.BtType = &v
}

// GetOperand1 returns the Operand1 field value if set, zero value otherwise.
func (o *BTOrFilter167) GetOperand1() BTQueryFilter183 {
	if o == nil || o.Operand1 == nil {
		var ret BTQueryFilter183
		return ret
	}
	return *o.Operand1
}

// GetOperand1Ok returns a tuple with the Operand1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOrFilter167) GetOperand1Ok() (*BTQueryFilter183, bool) {
	if o == nil || o.Operand1 == nil {
		return nil, false
	}
	return o.Operand1, true
}

// HasOperand1 returns a boolean if a field has been set.
func (o *BTOrFilter167) HasOperand1() bool {
	if o != nil && o.Operand1 != nil {
		return true
	}

	return false
}

// SetOperand1 gets a reference to the given BTQueryFilter183 and assigns it to the Operand1 field.
func (o *BTOrFilter167) SetOperand1(v BTQueryFilter183) {
	o.Operand1 = &v
}

// GetOperand2 returns the Operand2 field value if set, zero value otherwise.
func (o *BTOrFilter167) GetOperand2() BTQueryFilter183 {
	if o == nil || o.Operand2 == nil {
		var ret BTQueryFilter183
		return ret
	}
	return *o.Operand2
}

// GetOperand2Ok returns a tuple with the Operand2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOrFilter167) GetOperand2Ok() (*BTQueryFilter183, bool) {
	if o == nil || o.Operand2 == nil {
		return nil, false
	}
	return o.Operand2, true
}

// HasOperand2 returns a boolean if a field has been set.
func (o *BTOrFilter167) HasOperand2() bool {
	if o != nil && o.Operand2 != nil {
		return true
	}

	return false
}

// SetOperand2 gets a reference to the given BTQueryFilter183 and assigns it to the Operand2 field.
func (o *BTOrFilter167) SetOperand2(v BTQueryFilter183) {
	o.Operand2 = &v
}

func (o BTOrFilter167) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Operand1 != nil {
		toSerialize["operand1"] = o.Operand1
	}
	if o.Operand2 != nil {
		toSerialize["operand2"] = o.Operand2
	}
	return json.Marshal(toSerialize)
}

type NullableBTOrFilter167 struct {
	value *BTOrFilter167
	isSet bool
}

func (v NullableBTOrFilter167) Get() *BTOrFilter167 {
	return v.value
}

func (v *NullableBTOrFilter167) Set(val *BTOrFilter167) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOrFilter167) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOrFilter167) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOrFilter167(val *BTOrFilter167) *NullableBTOrFilter167 {
	return &NullableBTOrFilter167{value: val, isSet: true}
}

func (v NullableBTOrFilter167) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOrFilter167) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
