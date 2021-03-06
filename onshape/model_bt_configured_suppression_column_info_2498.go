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

// BTConfiguredSuppressionColumnInfo2498 struct for BTConfiguredSuppressionColumnInfo2498
type BTConfiguredSuppressionColumnInfo2498 struct {
	BTConfiguredValuesColumnInfo1025
	BtType *string `json:"btType,omitempty"`
}

// NewBTConfiguredSuppressionColumnInfo2498 instantiates a new BTConfiguredSuppressionColumnInfo2498 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredSuppressionColumnInfo2498() *BTConfiguredSuppressionColumnInfo2498 {
	this := BTConfiguredSuppressionColumnInfo2498{}
	return &this
}

// NewBTConfiguredSuppressionColumnInfo2498WithDefaults instantiates a new BTConfiguredSuppressionColumnInfo2498 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredSuppressionColumnInfo2498WithDefaults() *BTConfiguredSuppressionColumnInfo2498 {
	this := BTConfiguredSuppressionColumnInfo2498{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredSuppressionColumnInfo2498) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredSuppressionColumnInfo2498) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredSuppressionColumnInfo2498) SetBtType(v string) {
	o.BtType = &v
}

func (o BTConfiguredSuppressionColumnInfo2498) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTConfiguredValuesColumnInfo1025, errBTConfiguredValuesColumnInfo1025 := json.Marshal(o.BTConfiguredValuesColumnInfo1025)
	if errBTConfiguredValuesColumnInfo1025 != nil {
		return []byte{}, errBTConfiguredValuesColumnInfo1025
	}
	errBTConfiguredValuesColumnInfo1025 = json.Unmarshal([]byte(serializedBTConfiguredValuesColumnInfo1025), &toSerialize)
	if errBTConfiguredValuesColumnInfo1025 != nil {
		return []byte{}, errBTConfiguredValuesColumnInfo1025
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredSuppressionColumnInfo2498 struct {
	value *BTConfiguredSuppressionColumnInfo2498
	isSet bool
}

func (v NullableBTConfiguredSuppressionColumnInfo2498) Get() *BTConfiguredSuppressionColumnInfo2498 {
	return v.value
}

func (v *NullableBTConfiguredSuppressionColumnInfo2498) Set(val *BTConfiguredSuppressionColumnInfo2498) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredSuppressionColumnInfo2498) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredSuppressionColumnInfo2498) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredSuppressionColumnInfo2498(val *BTConfiguredSuppressionColumnInfo2498) *NullableBTConfiguredSuppressionColumnInfo2498 {
	return &NullableBTConfiguredSuppressionColumnInfo2498{value: val, isSet: true}
}

func (v NullableBTConfiguredSuppressionColumnInfo2498) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredSuppressionColumnInfo2498) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
