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

// BTSimulationTableColumnInfo1785 struct for BTSimulationTableColumnInfo1785
type BTSimulationTableColumnInfo1785 struct {
	BTTableColumnInfo1222
	BtType             *string                                `json:"btType,omitempty"`
	Id                 *string                                `json:"id,omitempty"`
	NodeId             *string                                `json:"nodeId,omitempty"`
	Specification      *BTTableColumnSpec1967                 `json:"specification,omitempty"`
	CrossHighlightData *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightData,omitempty"`
}

// NewBTSimulationTableColumnInfo1785 instantiates a new BTSimulationTableColumnInfo1785 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSimulationTableColumnInfo1785() *BTSimulationTableColumnInfo1785 {
	this := BTSimulationTableColumnInfo1785{}
	return &this
}

// NewBTSimulationTableColumnInfo1785WithDefaults instantiates a new BTSimulationTableColumnInfo1785 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSimulationTableColumnInfo1785WithDefaults() *BTSimulationTableColumnInfo1785 {
	this := BTSimulationTableColumnInfo1785{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSimulationTableColumnInfo1785) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationTableColumnInfo1785) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSimulationTableColumnInfo1785) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSimulationTableColumnInfo1785) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTSimulationTableColumnInfo1785) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationTableColumnInfo1785) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTSimulationTableColumnInfo1785) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTSimulationTableColumnInfo1785) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTSimulationTableColumnInfo1785) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationTableColumnInfo1785) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTSimulationTableColumnInfo1785) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTSimulationTableColumnInfo1785) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTSimulationTableColumnInfo1785) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationTableColumnInfo1785) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTSimulationTableColumnInfo1785) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTSimulationTableColumnInfo1785) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTSimulationTableColumnInfo1785) GetCrossHighlightData() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationTableColumnInfo1785) GetCrossHighlightDataOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTSimulationTableColumnInfo1785) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightData field.
func (o *BTSimulationTableColumnInfo1785) SetCrossHighlightData(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightData = &v
}

func (o BTSimulationTableColumnInfo1785) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTableColumnInfo1222, errBTTableColumnInfo1222 := json.Marshal(o.BTTableColumnInfo1222)
	if errBTTableColumnInfo1222 != nil {
		return []byte{}, errBTTableColumnInfo1222
	}
	errBTTableColumnInfo1222 = json.Unmarshal([]byte(serializedBTTableColumnInfo1222), &toSerialize)
	if errBTTableColumnInfo1222 != nil {
		return []byte{}, errBTTableColumnInfo1222
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTSimulationTableColumnInfo1785 struct {
	value *BTSimulationTableColumnInfo1785
	isSet bool
}

func (v NullableBTSimulationTableColumnInfo1785) Get() *BTSimulationTableColumnInfo1785 {
	return v.value
}

func (v *NullableBTSimulationTableColumnInfo1785) Set(val *BTSimulationTableColumnInfo1785) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSimulationTableColumnInfo1785) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSimulationTableColumnInfo1785) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSimulationTableColumnInfo1785(val *BTSimulationTableColumnInfo1785) *NullableBTSimulationTableColumnInfo1785 {
	return &NullableBTSimulationTableColumnInfo1785{value: val, isSet: true}
}

func (v NullableBTSimulationTableColumnInfo1785) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSimulationTableColumnInfo1785) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
