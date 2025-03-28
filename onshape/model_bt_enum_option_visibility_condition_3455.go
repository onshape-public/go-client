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

// BTEnumOptionVisibilityCondition3455 - struct for BTEnumOptionVisibilityCondition3455
type BTEnumOptionVisibilityCondition3455 struct {
	implBTEnumOptionVisibilityCondition3455 interface{}
}

// BTEnumOptionVisibilityForRange4297AsBTEnumOptionVisibilityCondition3455 is a convenience function that returns BTEnumOptionVisibilityForRange4297 wrapped in BTEnumOptionVisibilityCondition3455
func (o *BTEnumOptionVisibilityForRange4297) AsBTEnumOptionVisibilityCondition3455() *BTEnumOptionVisibilityCondition3455 {
	return &BTEnumOptionVisibilityCondition3455{o}
}

// BTEnumOptionVisibilityForList1613AsBTEnumOptionVisibilityCondition3455 is a convenience function that returns BTEnumOptionVisibilityForList1613 wrapped in BTEnumOptionVisibilityCondition3455
func (o *BTEnumOptionVisibilityForList1613) AsBTEnumOptionVisibilityCondition3455() *BTEnumOptionVisibilityCondition3455 {
	return &BTEnumOptionVisibilityCondition3455{o}
}

// NewBTEnumOptionVisibilityCondition3455 instantiates a new BTEnumOptionVisibilityCondition3455 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEnumOptionVisibilityCondition3455() *BTEnumOptionVisibilityCondition3455 {
	this := BTEnumOptionVisibilityCondition3455{Newbase_BTEnumOptionVisibilityCondition3455()}
	return &this
}

// NewBTEnumOptionVisibilityCondition3455WithDefaults instantiates a new BTEnumOptionVisibilityCondition3455 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEnumOptionVisibilityCondition3455WithDefaults() *BTEnumOptionVisibilityCondition3455 {
	this := BTEnumOptionVisibilityCondition3455{Newbase_BTEnumOptionVisibilityCondition3455WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEnumOptionVisibilityCondition3455) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEnumOptionVisibilityCondition3455) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEnumOptionVisibilityCondition3455) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEnumOptionVisibilityCondition3455) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *BTEnumOptionVisibilityCondition3455) GetCondition() BTParameterVisibilityCondition177 {
	type getResult interface {
		GetCondition() BTParameterVisibilityCondition177
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCondition()
	} else {
		var de BTParameterVisibilityCondition177
		return de
	}
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEnumOptionVisibilityCondition3455) GetConditionOk() (*BTParameterVisibilityCondition177, bool) {
	type getResult interface {
		GetConditionOk() (*BTParameterVisibilityCondition177, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConditionOk()
	} else {
		return nil, false
	}
}

// HasCondition returns a boolean if a field has been set.
func (o *BTEnumOptionVisibilityCondition3455) HasCondition() bool {
	type getResult interface {
		HasCondition() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCondition()
	} else {
		return false
	}
}

// SetCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the Condition field.
func (o *BTEnumOptionVisibilityCondition3455) SetCondition(v BTParameterVisibilityCondition177) {
	type getResult interface {
		SetCondition(v BTParameterVisibilityCondition177)
	}

	o.GetActualInstance().(getResult).SetCondition(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTEnumOptionVisibilityCondition3455) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTEnumOptionVisibilityForList-1613'
	if jsonDict["btType"] == "BTEnumOptionVisibilityForList-1613" {
		// try to unmarshal JSON data into BTEnumOptionVisibilityForList1613
		var qr *BTEnumOptionVisibilityForList1613
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEnumOptionVisibilityCondition3455 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEnumOptionVisibilityCondition3455 = nil
			return fmt.Errorf("failed to unmarshal BTEnumOptionVisibilityCondition3455 as BTEnumOptionVisibilityForList1613: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEnumOptionVisibilityForRange-4297'
	if jsonDict["btType"] == "BTEnumOptionVisibilityForRange-4297" {
		// try to unmarshal JSON data into BTEnumOptionVisibilityForRange4297
		var qr *BTEnumOptionVisibilityForRange4297
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEnumOptionVisibilityCondition3455 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEnumOptionVisibilityCondition3455 = nil
			return fmt.Errorf("failed to unmarshal BTEnumOptionVisibilityCondition3455 as BTEnumOptionVisibilityForRange4297: %s", err.Error())
		}
	}

	var qtx *base_BTEnumOptionVisibilityCondition3455
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTEnumOptionVisibilityCondition3455 = qtx
		return nil // data stored in dst.base_BTEnumOptionVisibilityCondition3455, return on the first match
	} else {
		dst.implBTEnumOptionVisibilityCondition3455 = nil
		return fmt.Errorf("failed to unmarshal BTEnumOptionVisibilityCondition3455 as base_BTEnumOptionVisibilityCondition3455: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTEnumOptionVisibilityCondition3455) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTEnumOptionVisibilityCondition3455) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTEnumOptionVisibilityCondition3455
}

type NullableBTEnumOptionVisibilityCondition3455 struct {
	value *BTEnumOptionVisibilityCondition3455
	isSet bool
}

func (v NullableBTEnumOptionVisibilityCondition3455) Get() *BTEnumOptionVisibilityCondition3455 {
	return v.value
}

func (v *NullableBTEnumOptionVisibilityCondition3455) Set(val *BTEnumOptionVisibilityCondition3455) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEnumOptionVisibilityCondition3455) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEnumOptionVisibilityCondition3455) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEnumOptionVisibilityCondition3455(val *BTEnumOptionVisibilityCondition3455) *NullableBTEnumOptionVisibilityCondition3455 {
	return &NullableBTEnumOptionVisibilityCondition3455{value: val, isSet: true}
}

func (v NullableBTEnumOptionVisibilityCondition3455) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEnumOptionVisibilityCondition3455) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTEnumOptionVisibilityCondition3455 struct {
	// Type of JSON object.
	BtType    *string                            `json:"btType,omitempty"`
	Condition *BTParameterVisibilityCondition177 `json:"condition,omitempty"`
}

// Newbase_BTEnumOptionVisibilityCondition3455 instantiates a new base_BTEnumOptionVisibilityCondition3455 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTEnumOptionVisibilityCondition3455() *base_BTEnumOptionVisibilityCondition3455 {
	this := base_BTEnumOptionVisibilityCondition3455{}
	return &this
}

// Newbase_BTEnumOptionVisibilityCondition3455WithDefaults instantiates a new base_BTEnumOptionVisibilityCondition3455 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTEnumOptionVisibilityCondition3455WithDefaults() *base_BTEnumOptionVisibilityCondition3455 {
	this := base_BTEnumOptionVisibilityCondition3455{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTEnumOptionVisibilityCondition3455) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEnumOptionVisibilityCondition3455) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTEnumOptionVisibilityCondition3455) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTEnumOptionVisibilityCondition3455) SetBtType(v string) {
	o.BtType = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *base_BTEnumOptionVisibilityCondition3455) GetCondition() BTParameterVisibilityCondition177 {
	if o == nil || o.Condition == nil {
		var ret BTParameterVisibilityCondition177
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEnumOptionVisibilityCondition3455) GetConditionOk() (*BTParameterVisibilityCondition177, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *base_BTEnumOptionVisibilityCondition3455) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the Condition field.
func (o *base_BTEnumOptionVisibilityCondition3455) SetCondition(v BTParameterVisibilityCondition177) {
	o.Condition = &v
}

func (o base_BTEnumOptionVisibilityCondition3455) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	return json.Marshal(toSerialize)
}
