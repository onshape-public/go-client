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

// BTConfiguredPartPropertiesTable2740 struct for BTConfiguredPartPropertiesTable2740
type BTConfiguredPartPropertiesTable2740 struct {
	BTTable1825
	BtType *string `json:"btType,omitempty"`
	PartDeterministicId *string `json:"partDeterministicId,omitempty"`
	PartDeterministicIds *[]string `json:"partDeterministicIds,omitempty"`
	PropertyNodeId *string `json:"propertyNodeId,omitempty"`
}

// NewBTConfiguredPartPropertiesTable2740 instantiates a new BTConfiguredPartPropertiesTable2740 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredPartPropertiesTable2740() *BTConfiguredPartPropertiesTable2740 {
	this := BTConfiguredPartPropertiesTable2740{}
	return &this
}

// NewBTConfiguredPartPropertiesTable2740WithDefaults instantiates a new BTConfiguredPartPropertiesTable2740 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredPartPropertiesTable2740WithDefaults() *BTConfiguredPartPropertiesTable2740 {
	this := BTConfiguredPartPropertiesTable2740{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredPartPropertiesTable2740) SetBtType(v string) {
	o.BtType = &v
}

// GetPartDeterministicId returns the PartDeterministicId field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicId() string {
	if o == nil || o.PartDeterministicId == nil {
		var ret string
		return ret
	}
	return *o.PartDeterministicId
}

// GetPartDeterministicIdOk returns a tuple with the PartDeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicIdOk() (*string, bool) {
	if o == nil || o.PartDeterministicId == nil {
		return nil, false
	}
	return o.PartDeterministicId, true
}

// HasPartDeterministicId returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasPartDeterministicId() bool {
	if o != nil && o.PartDeterministicId != nil {
		return true
	}

	return false
}

// SetPartDeterministicId gets a reference to the given string and assigns it to the PartDeterministicId field.
func (o *BTConfiguredPartPropertiesTable2740) SetPartDeterministicId(v string) {
	o.PartDeterministicId = &v
}

// GetPartDeterministicIds returns the PartDeterministicIds field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicIds() []string {
	if o == nil || o.PartDeterministicIds == nil {
		var ret []string
		return ret
	}
	return *o.PartDeterministicIds
}

// GetPartDeterministicIdsOk returns a tuple with the PartDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicIdsOk() (*[]string, bool) {
	if o == nil || o.PartDeterministicIds == nil {
		return nil, false
	}
	return o.PartDeterministicIds, true
}

// HasPartDeterministicIds returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasPartDeterministicIds() bool {
	if o != nil && o.PartDeterministicIds != nil {
		return true
	}

	return false
}

// SetPartDeterministicIds gets a reference to the given []string and assigns it to the PartDeterministicIds field.
func (o *BTConfiguredPartPropertiesTable2740) SetPartDeterministicIds(v []string) {
	o.PartDeterministicIds = &v
}

// GetPropertyNodeId returns the PropertyNodeId field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetPropertyNodeId() string {
	if o == nil || o.PropertyNodeId == nil {
		var ret string
		return ret
	}
	return *o.PropertyNodeId
}

// GetPropertyNodeIdOk returns a tuple with the PropertyNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetPropertyNodeIdOk() (*string, bool) {
	if o == nil || o.PropertyNodeId == nil {
		return nil, false
	}
	return o.PropertyNodeId, true
}

// HasPropertyNodeId returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasPropertyNodeId() bool {
	if o != nil && o.PropertyNodeId != nil {
		return true
	}

	return false
}

// SetPropertyNodeId gets a reference to the given string and assigns it to the PropertyNodeId field.
func (o *BTConfiguredPartPropertiesTable2740) SetPropertyNodeId(v string) {
	o.PropertyNodeId = &v
}

func (o BTConfiguredPartPropertiesTable2740) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTable1825, errBTTable1825 := json.Marshal(o.BTTable1825)
	if errBTTable1825 != nil {
		return []byte{}, errBTTable1825
	}
	errBTTable1825 = json.Unmarshal([]byte(serializedBTTable1825), &toSerialize)
	if errBTTable1825 != nil {
		return []byte{}, errBTTable1825
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.PartDeterministicId != nil {
		toSerialize["partDeterministicId"] = o.PartDeterministicId
	}
	if o.PartDeterministicIds != nil {
		toSerialize["partDeterministicIds"] = o.PartDeterministicIds
	}
	if o.PropertyNodeId != nil {
		toSerialize["propertyNodeId"] = o.PropertyNodeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredPartPropertiesTable2740 struct {
	value *BTConfiguredPartPropertiesTable2740
	isSet bool
}

func (v NullableBTConfiguredPartPropertiesTable2740) Get() *BTConfiguredPartPropertiesTable2740 {
	return v.value
}

func (v *NullableBTConfiguredPartPropertiesTable2740) Set(val *BTConfiguredPartPropertiesTable2740) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredPartPropertiesTable2740) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredPartPropertiesTable2740) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredPartPropertiesTable2740(val *BTConfiguredPartPropertiesTable2740) *NullableBTConfiguredPartPropertiesTable2740 {
	return &NullableBTConfiguredPartPropertiesTable2740{value: val, isSet: true}
}

func (v NullableBTConfiguredPartPropertiesTable2740) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredPartPropertiesTable2740) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
