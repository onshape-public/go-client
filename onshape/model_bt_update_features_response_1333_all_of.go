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

// BTUpdateFeaturesResponse1333AllOf struct for BTUpdateFeaturesResponse1333AllOf
type BTUpdateFeaturesResponse1333AllOf struct {
	BtType *string `json:"btType,omitempty"`
	FeatureStates *map[string]BTFeatureState1688 `json:"featureStates,omitempty"`
	Features *[]BTMFeature134 `json:"features,omitempty"`
}

// NewBTUpdateFeaturesResponse1333AllOf instantiates a new BTUpdateFeaturesResponse1333AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateFeaturesResponse1333AllOf() *BTUpdateFeaturesResponse1333AllOf {
	this := BTUpdateFeaturesResponse1333AllOf{}
	return &this
}

// NewBTUpdateFeaturesResponse1333AllOfWithDefaults instantiates a new BTUpdateFeaturesResponse1333AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateFeaturesResponse1333AllOfWithDefaults() *BTUpdateFeaturesResponse1333AllOf {
	this := BTUpdateFeaturesResponse1333AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTUpdateFeaturesResponse1333AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesResponse1333AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTUpdateFeaturesResponse1333AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTUpdateFeaturesResponse1333AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureStates returns the FeatureStates field value if set, zero value otherwise.
func (o *BTUpdateFeaturesResponse1333AllOf) GetFeatureStates() map[string]BTFeatureState1688 {
	if o == nil || o.FeatureStates == nil {
		var ret map[string]BTFeatureState1688
		return ret
	}
	return *o.FeatureStates
}

// GetFeatureStatesOk returns a tuple with the FeatureStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesResponse1333AllOf) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool) {
	if o == nil || o.FeatureStates == nil {
		return nil, false
	}
	return o.FeatureStates, true
}

// HasFeatureStates returns a boolean if a field has been set.
func (o *BTUpdateFeaturesResponse1333AllOf) HasFeatureStates() bool {
	if o != nil && o.FeatureStates != nil {
		return true
	}

	return false
}

// SetFeatureStates gets a reference to the given map[string]BTFeatureState1688 and assigns it to the FeatureStates field.
func (o *BTUpdateFeaturesResponse1333AllOf) SetFeatureStates(v map[string]BTFeatureState1688) {
	o.FeatureStates = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTUpdateFeaturesResponse1333AllOf) GetFeatures() []BTMFeature134 {
	if o == nil || o.Features == nil {
		var ret []BTMFeature134
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesResponse1333AllOf) GetFeaturesOk() (*[]BTMFeature134, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTUpdateFeaturesResponse1333AllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTMFeature134 and assigns it to the Features field.
func (o *BTUpdateFeaturesResponse1333AllOf) SetFeatures(v []BTMFeature134) {
	o.Features = &v
}

func (o BTUpdateFeaturesResponse1333AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureStates != nil {
		toSerialize["featureStates"] = o.FeatureStates
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateFeaturesResponse1333AllOf struct {
	value *BTUpdateFeaturesResponse1333AllOf
	isSet bool
}

func (v NullableBTUpdateFeaturesResponse1333AllOf) Get() *BTUpdateFeaturesResponse1333AllOf {
	return v.value
}

func (v *NullableBTUpdateFeaturesResponse1333AllOf) Set(val *BTUpdateFeaturesResponse1333AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateFeaturesResponse1333AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateFeaturesResponse1333AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateFeaturesResponse1333AllOf(val *BTUpdateFeaturesResponse1333AllOf) *NullableBTUpdateFeaturesResponse1333AllOf {
	return &NullableBTUpdateFeaturesResponse1333AllOf{value: val, isSet: true}
}

func (v NullableBTUpdateFeaturesResponse1333AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateFeaturesResponse1333AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
