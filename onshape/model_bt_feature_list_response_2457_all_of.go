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

// BTFeatureListResponse2457AllOf struct for BTFeatureListResponse2457AllOf
type BTFeatureListResponse2457AllOf struct {
	BtType *string `json:"btType,omitempty"`
	DefaultFeatures *[]BTMFeature134 `json:"defaultFeatures,omitempty"`
	FeatureStates *map[string]BTFeatureState1688 `json:"featureStates,omitempty"`
	Features *[]BTMFeature134 `json:"features,omitempty"`
	Imports *[]BTMImport136 `json:"imports,omitempty"`
	IsComplete *bool `json:"isComplete,omitempty"`
	RollbackIndex *int32 `json:"rollbackIndex,omitempty"`
}

// NewBTFeatureListResponse2457AllOf instantiates a new BTFeatureListResponse2457AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureListResponse2457AllOf() *BTFeatureListResponse2457AllOf {
	this := BTFeatureListResponse2457AllOf{}
	return &this
}

// NewBTFeatureListResponse2457AllOfWithDefaults instantiates a new BTFeatureListResponse2457AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureListResponse2457AllOfWithDefaults() *BTFeatureListResponse2457AllOf {
	this := BTFeatureListResponse2457AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureListResponse2457AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultFeatures returns the DefaultFeatures field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetDefaultFeatures() []BTMFeature134 {
	if o == nil || o.DefaultFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return *o.DefaultFeatures
}

// GetDefaultFeaturesOk returns a tuple with the DefaultFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetDefaultFeaturesOk() (*[]BTMFeature134, bool) {
	if o == nil || o.DefaultFeatures == nil {
		return nil, false
	}
	return o.DefaultFeatures, true
}

// HasDefaultFeatures returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasDefaultFeatures() bool {
	if o != nil && o.DefaultFeatures != nil {
		return true
	}

	return false
}

// SetDefaultFeatures gets a reference to the given []BTMFeature134 and assigns it to the DefaultFeatures field.
func (o *BTFeatureListResponse2457AllOf) SetDefaultFeatures(v []BTMFeature134) {
	o.DefaultFeatures = &v
}

// GetFeatureStates returns the FeatureStates field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetFeatureStates() map[string]BTFeatureState1688 {
	if o == nil || o.FeatureStates == nil {
		var ret map[string]BTFeatureState1688
		return ret
	}
	return *o.FeatureStates
}

// GetFeatureStatesOk returns a tuple with the FeatureStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool) {
	if o == nil || o.FeatureStates == nil {
		return nil, false
	}
	return o.FeatureStates, true
}

// HasFeatureStates returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasFeatureStates() bool {
	if o != nil && o.FeatureStates != nil {
		return true
	}

	return false
}

// SetFeatureStates gets a reference to the given map[string]BTFeatureState1688 and assigns it to the FeatureStates field.
func (o *BTFeatureListResponse2457AllOf) SetFeatureStates(v map[string]BTFeatureState1688) {
	o.FeatureStates = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetFeatures() []BTMFeature134 {
	if o == nil || o.Features == nil {
		var ret []BTMFeature134
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetFeaturesOk() (*[]BTMFeature134, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTMFeature134 and assigns it to the Features field.
func (o *BTFeatureListResponse2457AllOf) SetFeatures(v []BTMFeature134) {
	o.Features = &v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetImports() []BTMImport136 {
	if o == nil || o.Imports == nil {
		var ret []BTMImport136
		return ret
	}
	return *o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetImportsOk() (*[]BTMImport136, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []BTMImport136 and assigns it to the Imports field.
func (o *BTFeatureListResponse2457AllOf) SetImports(v []BTMImport136) {
	o.Imports = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetIsComplete() bool {
	if o == nil || o.IsComplete == nil {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetIsCompleteOk() (*bool, bool) {
	if o == nil || o.IsComplete == nil {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasIsComplete() bool {
	if o != nil && o.IsComplete != nil {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *BTFeatureListResponse2457AllOf) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetRollbackIndex returns the RollbackIndex field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457AllOf) GetRollbackIndex() int32 {
	if o == nil || o.RollbackIndex == nil {
		var ret int32
		return ret
	}
	return *o.RollbackIndex
}

// GetRollbackIndexOk returns a tuple with the RollbackIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457AllOf) GetRollbackIndexOk() (*int32, bool) {
	if o == nil || o.RollbackIndex == nil {
		return nil, false
	}
	return o.RollbackIndex, true
}

// HasRollbackIndex returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457AllOf) HasRollbackIndex() bool {
	if o != nil && o.RollbackIndex != nil {
		return true
	}

	return false
}

// SetRollbackIndex gets a reference to the given int32 and assigns it to the RollbackIndex field.
func (o *BTFeatureListResponse2457AllOf) SetRollbackIndex(v int32) {
	o.RollbackIndex = &v
}

func (o BTFeatureListResponse2457AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultFeatures != nil {
		toSerialize["defaultFeatures"] = o.DefaultFeatures
	}
	if o.FeatureStates != nil {
		toSerialize["featureStates"] = o.FeatureStates
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.IsComplete != nil {
		toSerialize["isComplete"] = o.IsComplete
	}
	if o.RollbackIndex != nil {
		toSerialize["rollbackIndex"] = o.RollbackIndex
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureListResponse2457AllOf struct {
	value *BTFeatureListResponse2457AllOf
	isSet bool
}

func (v NullableBTFeatureListResponse2457AllOf) Get() *BTFeatureListResponse2457AllOf {
	return v.value
}

func (v *NullableBTFeatureListResponse2457AllOf) Set(val *BTFeatureListResponse2457AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureListResponse2457AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureListResponse2457AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureListResponse2457AllOf(val *BTFeatureListResponse2457AllOf) *NullableBTFeatureListResponse2457AllOf {
	return &NullableBTFeatureListResponse2457AllOf{value: val, isSet: true}
}

func (v NullableBTFeatureListResponse2457AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureListResponse2457AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
