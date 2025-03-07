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

// BTMatrix3x3340 struct for BTMatrix3x3340
type BTMatrix3x3340 struct {
	// Type of JSON object.
	BtType *string  `json:"btType,omitempty"`
	M00    *float64 `json:"m00,omitempty"`
	M01    *float64 `json:"m01,omitempty"`
	M02    *float64 `json:"m02,omitempty"`
	M10    *float64 `json:"m10,omitempty"`
	M11    *float64 `json:"m11,omitempty"`
	M12    *float64 `json:"m12,omitempty"`
	M20    *float64 `json:"m20,omitempty"`
	M21    *float64 `json:"m21,omitempty"`
	M22    *float64 `json:"m22,omitempty"`
}

// NewBTMatrix3x3340 instantiates a new BTMatrix3x3340 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMatrix3x3340() *BTMatrix3x3340 {
	this := BTMatrix3x3340{}
	return &this
}

// NewBTMatrix3x3340WithDefaults instantiates a new BTMatrix3x3340 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMatrix3x3340WithDefaults() *BTMatrix3x3340 {
	this := BTMatrix3x3340{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMatrix3x3340) SetBtType(v string) {
	o.BtType = &v
}

// GetM00 returns the M00 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM00() float64 {
	if o == nil || o.M00 == nil {
		var ret float64
		return ret
	}
	return *o.M00
}

// GetM00Ok returns a tuple with the M00 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM00Ok() (*float64, bool) {
	if o == nil || o.M00 == nil {
		return nil, false
	}
	return o.M00, true
}

// HasM00 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM00() bool {
	if o != nil && o.M00 != nil {
		return true
	}

	return false
}

// SetM00 gets a reference to the given float64 and assigns it to the M00 field.
func (o *BTMatrix3x3340) SetM00(v float64) {
	o.M00 = &v
}

// GetM01 returns the M01 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM01() float64 {
	if o == nil || o.M01 == nil {
		var ret float64
		return ret
	}
	return *o.M01
}

// GetM01Ok returns a tuple with the M01 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM01Ok() (*float64, bool) {
	if o == nil || o.M01 == nil {
		return nil, false
	}
	return o.M01, true
}

// HasM01 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM01() bool {
	if o != nil && o.M01 != nil {
		return true
	}

	return false
}

// SetM01 gets a reference to the given float64 and assigns it to the M01 field.
func (o *BTMatrix3x3340) SetM01(v float64) {
	o.M01 = &v
}

// GetM02 returns the M02 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM02() float64 {
	if o == nil || o.M02 == nil {
		var ret float64
		return ret
	}
	return *o.M02
}

// GetM02Ok returns a tuple with the M02 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM02Ok() (*float64, bool) {
	if o == nil || o.M02 == nil {
		return nil, false
	}
	return o.M02, true
}

// HasM02 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM02() bool {
	if o != nil && o.M02 != nil {
		return true
	}

	return false
}

// SetM02 gets a reference to the given float64 and assigns it to the M02 field.
func (o *BTMatrix3x3340) SetM02(v float64) {
	o.M02 = &v
}

// GetM10 returns the M10 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM10() float64 {
	if o == nil || o.M10 == nil {
		var ret float64
		return ret
	}
	return *o.M10
}

// GetM10Ok returns a tuple with the M10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM10Ok() (*float64, bool) {
	if o == nil || o.M10 == nil {
		return nil, false
	}
	return o.M10, true
}

// HasM10 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM10() bool {
	if o != nil && o.M10 != nil {
		return true
	}

	return false
}

// SetM10 gets a reference to the given float64 and assigns it to the M10 field.
func (o *BTMatrix3x3340) SetM10(v float64) {
	o.M10 = &v
}

// GetM11 returns the M11 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM11() float64 {
	if o == nil || o.M11 == nil {
		var ret float64
		return ret
	}
	return *o.M11
}

// GetM11Ok returns a tuple with the M11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM11Ok() (*float64, bool) {
	if o == nil || o.M11 == nil {
		return nil, false
	}
	return o.M11, true
}

// HasM11 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM11() bool {
	if o != nil && o.M11 != nil {
		return true
	}

	return false
}

// SetM11 gets a reference to the given float64 and assigns it to the M11 field.
func (o *BTMatrix3x3340) SetM11(v float64) {
	o.M11 = &v
}

// GetM12 returns the M12 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM12() float64 {
	if o == nil || o.M12 == nil {
		var ret float64
		return ret
	}
	return *o.M12
}

// GetM12Ok returns a tuple with the M12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM12Ok() (*float64, bool) {
	if o == nil || o.M12 == nil {
		return nil, false
	}
	return o.M12, true
}

// HasM12 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM12() bool {
	if o != nil && o.M12 != nil {
		return true
	}

	return false
}

// SetM12 gets a reference to the given float64 and assigns it to the M12 field.
func (o *BTMatrix3x3340) SetM12(v float64) {
	o.M12 = &v
}

// GetM20 returns the M20 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM20() float64 {
	if o == nil || o.M20 == nil {
		var ret float64
		return ret
	}
	return *o.M20
}

// GetM20Ok returns a tuple with the M20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM20Ok() (*float64, bool) {
	if o == nil || o.M20 == nil {
		return nil, false
	}
	return o.M20, true
}

// HasM20 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM20() bool {
	if o != nil && o.M20 != nil {
		return true
	}

	return false
}

// SetM20 gets a reference to the given float64 and assigns it to the M20 field.
func (o *BTMatrix3x3340) SetM20(v float64) {
	o.M20 = &v
}

// GetM21 returns the M21 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM21() float64 {
	if o == nil || o.M21 == nil {
		var ret float64
		return ret
	}
	return *o.M21
}

// GetM21Ok returns a tuple with the M21 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM21Ok() (*float64, bool) {
	if o == nil || o.M21 == nil {
		return nil, false
	}
	return o.M21, true
}

// HasM21 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM21() bool {
	if o != nil && o.M21 != nil {
		return true
	}

	return false
}

// SetM21 gets a reference to the given float64 and assigns it to the M21 field.
func (o *BTMatrix3x3340) SetM21(v float64) {
	o.M21 = &v
}

// GetM22 returns the M22 field value if set, zero value otherwise.
func (o *BTMatrix3x3340) GetM22() float64 {
	if o == nil || o.M22 == nil {
		var ret float64
		return ret
	}
	return *o.M22
}

// GetM22Ok returns a tuple with the M22 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMatrix3x3340) GetM22Ok() (*float64, bool) {
	if o == nil || o.M22 == nil {
		return nil, false
	}
	return o.M22, true
}

// HasM22 returns a boolean if a field has been set.
func (o *BTMatrix3x3340) HasM22() bool {
	if o != nil && o.M22 != nil {
		return true
	}

	return false
}

// SetM22 gets a reference to the given float64 and assigns it to the M22 field.
func (o *BTMatrix3x3340) SetM22(v float64) {
	o.M22 = &v
}

func (o BTMatrix3x3340) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.M00 != nil {
		toSerialize["m00"] = o.M00
	}
	if o.M01 != nil {
		toSerialize["m01"] = o.M01
	}
	if o.M02 != nil {
		toSerialize["m02"] = o.M02
	}
	if o.M10 != nil {
		toSerialize["m10"] = o.M10
	}
	if o.M11 != nil {
		toSerialize["m11"] = o.M11
	}
	if o.M12 != nil {
		toSerialize["m12"] = o.M12
	}
	if o.M20 != nil {
		toSerialize["m20"] = o.M20
	}
	if o.M21 != nil {
		toSerialize["m21"] = o.M21
	}
	if o.M22 != nil {
		toSerialize["m22"] = o.M22
	}
	return json.Marshal(toSerialize)
}

type NullableBTMatrix3x3340 struct {
	value *BTMatrix3x3340
	isSet bool
}

func (v NullableBTMatrix3x3340) Get() *BTMatrix3x3340 {
	return v.value
}

func (v *NullableBTMatrix3x3340) Set(val *BTMatrix3x3340) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMatrix3x3340) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMatrix3x3340) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMatrix3x3340(val *BTMatrix3x3340) *NullableBTMatrix3x3340 {
	return &NullableBTMatrix3x3340{value: val, isSet: true}
}

func (v NullableBTMatrix3x3340) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMatrix3x3340) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
