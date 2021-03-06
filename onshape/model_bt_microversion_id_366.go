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

// BTMicroversionId366 struct for BTMicroversionId366
type BTMicroversionId366 struct {
	BtType *string `json:"btType,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	TheId *string `json:"theId,omitempty"`
}

// NewBTMicroversionId366 instantiates a new BTMicroversionId366 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMicroversionId366() *BTMicroversionId366 {
	this := BTMicroversionId366{}
	return &this
}

// NewBTMicroversionId366WithDefaults instantiates a new BTMicroversionId366 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMicroversionId366WithDefaults() *BTMicroversionId366 {
	this := BTMicroversionId366{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMicroversionId366) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionId366) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMicroversionId366) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMicroversionId366) SetBtType(v string) {
	o.BtType = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *BTMicroversionId366) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionId366) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *BTMicroversionId366) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *BTMicroversionId366) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetTheId returns the TheId field value if set, zero value otherwise.
func (o *BTMicroversionId366) GetTheId() string {
	if o == nil || o.TheId == nil {
		var ret string
		return ret
	}
	return *o.TheId
}

// GetTheIdOk returns a tuple with the TheId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionId366) GetTheIdOk() (*string, bool) {
	if o == nil || o.TheId == nil {
		return nil, false
	}
	return o.TheId, true
}

// HasTheId returns a boolean if a field has been set.
func (o *BTMicroversionId366) HasTheId() bool {
	if o != nil && o.TheId != nil {
		return true
	}

	return false
}

// SetTheId gets a reference to the given string and assigns it to the TheId field.
func (o *BTMicroversionId366) SetTheId(v string) {
	o.TheId = &v
}

func (o BTMicroversionId366) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.TheId != nil {
		toSerialize["theId"] = o.TheId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMicroversionId366 struct {
	value *BTMicroversionId366
	isSet bool
}

func (v NullableBTMicroversionId366) Get() *BTMicroversionId366 {
	return v.value
}

func (v *NullableBTMicroversionId366) Set(val *BTMicroversionId366) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMicroversionId366) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMicroversionId366) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMicroversionId366(val *BTMicroversionId366) *NullableBTMicroversionId366 {
	return &NullableBTMicroversionId366{value: val, isSet: true}
}

func (v NullableBTMicroversionId366) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMicroversionId366) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
