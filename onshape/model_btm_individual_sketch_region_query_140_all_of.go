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

// BTMIndividualSketchRegionQuery140AllOf struct for BTMIndividualSketchRegionQuery140AllOf
type BTMIndividualSketchRegionQuery140AllOf struct {
	BtType *string `json:"btType,omitempty"`
	FeatureId *string `json:"featureId,omitempty"`
	FilterInnerLoops *bool `json:"filterInnerLoops,omitempty"`
}

// NewBTMIndividualSketchRegionQuery140AllOf instantiates a new BTMIndividualSketchRegionQuery140AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualSketchRegionQuery140AllOf() *BTMIndividualSketchRegionQuery140AllOf {
	this := BTMIndividualSketchRegionQuery140AllOf{}
	return &this
}

// NewBTMIndividualSketchRegionQuery140AllOfWithDefaults instantiates a new BTMIndividualSketchRegionQuery140AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualSketchRegionQuery140AllOfWithDefaults() *BTMIndividualSketchRegionQuery140AllOf {
	this := BTMIndividualSketchRegionQuery140AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualSketchRegionQuery140AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140AllOf) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140AllOf) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140AllOf) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMIndividualSketchRegionQuery140AllOf) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFilterInnerLoops returns the FilterInnerLoops field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140AllOf) GetFilterInnerLoops() bool {
	if o == nil || o.FilterInnerLoops == nil {
		var ret bool
		return ret
	}
	return *o.FilterInnerLoops
}

// GetFilterInnerLoopsOk returns a tuple with the FilterInnerLoops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140AllOf) GetFilterInnerLoopsOk() (*bool, bool) {
	if o == nil || o.FilterInnerLoops == nil {
		return nil, false
	}
	return o.FilterInnerLoops, true
}

// HasFilterInnerLoops returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140AllOf) HasFilterInnerLoops() bool {
	if o != nil && o.FilterInnerLoops != nil {
		return true
	}

	return false
}

// SetFilterInnerLoops gets a reference to the given bool and assigns it to the FilterInnerLoops field.
func (o *BTMIndividualSketchRegionQuery140AllOf) SetFilterInnerLoops(v bool) {
	o.FilterInnerLoops = &v
}

func (o BTMIndividualSketchRegionQuery140AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FilterInnerLoops != nil {
		toSerialize["filterInnerLoops"] = o.FilterInnerLoops
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualSketchRegionQuery140AllOf struct {
	value *BTMIndividualSketchRegionQuery140AllOf
	isSet bool
}

func (v NullableBTMIndividualSketchRegionQuery140AllOf) Get() *BTMIndividualSketchRegionQuery140AllOf {
	return v.value
}

func (v *NullableBTMIndividualSketchRegionQuery140AllOf) Set(val *BTMIndividualSketchRegionQuery140AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualSketchRegionQuery140AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualSketchRegionQuery140AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualSketchRegionQuery140AllOf(val *BTMIndividualSketchRegionQuery140AllOf) *NullableBTMIndividualSketchRegionQuery140AllOf {
	return &NullableBTMIndividualSketchRegionQuery140AllOf{value: val, isSet: true}
}

func (v NullableBTMIndividualSketchRegionQuery140AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualSketchRegionQuery140AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
