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

// BTWidthMateDisplayData2888 struct for BTWidthMateDisplayData2888
type BTWidthMateDisplayData2888 struct {
	BTAssemblyFeatureDisplayData1783
	BtType           *string                          `json:"btType,omitempty"`
	Hidden           *bool                            `json:"hidden,omitempty"`
	IsDerivedFeature *bool                            `json:"isDerivedFeature,omitempty"`
	NodeId           *string                          `json:"nodeId,omitempty"`
	OwnerOccurrence  *BTOccurrence74                  `json:"ownerOccurrence,omitempty"`
	Status           *GBTAssemblyFeatureDisplayStatus `json:"status,omitempty"`
	Location         *BTCoordinateSystem387           `json:"location,omitempty"`
	MateConnectorIds []string                         `json:"mateConnectorIds,omitempty"`
}

// NewBTWidthMateDisplayData2888 instantiates a new BTWidthMateDisplayData2888 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWidthMateDisplayData2888() *BTWidthMateDisplayData2888 {
	this := BTWidthMateDisplayData2888{}
	return &this
}

// NewBTWidthMateDisplayData2888WithDefaults instantiates a new BTWidthMateDisplayData2888 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWidthMateDisplayData2888WithDefaults() *BTWidthMateDisplayData2888 {
	this := BTWidthMateDisplayData2888{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTWidthMateDisplayData2888) SetBtType(v string) {
	o.BtType = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTWidthMateDisplayData2888) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIsDerivedFeature returns the IsDerivedFeature field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetIsDerivedFeature() bool {
	if o == nil || o.IsDerivedFeature == nil {
		var ret bool
		return ret
	}
	return *o.IsDerivedFeature
}

// GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetIsDerivedFeatureOk() (*bool, bool) {
	if o == nil || o.IsDerivedFeature == nil {
		return nil, false
	}
	return o.IsDerivedFeature, true
}

// HasIsDerivedFeature returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasIsDerivedFeature() bool {
	if o != nil && o.IsDerivedFeature != nil {
		return true
	}

	return false
}

// SetIsDerivedFeature gets a reference to the given bool and assigns it to the IsDerivedFeature field.
func (o *BTWidthMateDisplayData2888) SetIsDerivedFeature(v bool) {
	o.IsDerivedFeature = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTWidthMateDisplayData2888) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOwnerOccurrence returns the OwnerOccurrence field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetOwnerOccurrence() BTOccurrence74 {
	if o == nil || o.OwnerOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OwnerOccurrence
}

// GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetOwnerOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.OwnerOccurrence == nil {
		return nil, false
	}
	return o.OwnerOccurrence, true
}

// HasOwnerOccurrence returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasOwnerOccurrence() bool {
	if o != nil && o.OwnerOccurrence != nil {
		return true
	}

	return false
}

// SetOwnerOccurrence gets a reference to the given BTOccurrence74 and assigns it to the OwnerOccurrence field.
func (o *BTWidthMateDisplayData2888) SetOwnerOccurrence(v BTOccurrence74) {
	o.OwnerOccurrence = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetStatus() GBTAssemblyFeatureDisplayStatus {
	if o == nil || o.Status == nil {
		var ret GBTAssemblyFeatureDisplayStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetStatusOk() (*GBTAssemblyFeatureDisplayStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GBTAssemblyFeatureDisplayStatus and assigns it to the Status field.
func (o *BTWidthMateDisplayData2888) SetStatus(v GBTAssemblyFeatureDisplayStatus) {
	o.Status = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetLocation() BTCoordinateSystem387 {
	if o == nil || o.Location == nil {
		var ret BTCoordinateSystem387
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetLocationOk() (*BTCoordinateSystem387, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTCoordinateSystem387 and assigns it to the Location field.
func (o *BTWidthMateDisplayData2888) SetLocation(v BTCoordinateSystem387) {
	o.Location = &v
}

// GetMateConnectorIds returns the MateConnectorIds field value if set, zero value otherwise.
func (o *BTWidthMateDisplayData2888) GetMateConnectorIds() []string {
	if o == nil || o.MateConnectorIds == nil {
		var ret []string
		return ret
	}
	return o.MateConnectorIds
}

// GetMateConnectorIdsOk returns a tuple with the MateConnectorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWidthMateDisplayData2888) GetMateConnectorIdsOk() ([]string, bool) {
	if o == nil || o.MateConnectorIds == nil {
		return nil, false
	}
	return o.MateConnectorIds, true
}

// HasMateConnectorIds returns a boolean if a field has been set.
func (o *BTWidthMateDisplayData2888) HasMateConnectorIds() bool {
	if o != nil && o.MateConnectorIds != nil {
		return true
	}

	return false
}

// SetMateConnectorIds gets a reference to the given []string and assigns it to the MateConnectorIds field.
func (o *BTWidthMateDisplayData2888) SetMateConnectorIds(v []string) {
	o.MateConnectorIds = v
}

func (o BTWidthMateDisplayData2888) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTAssemblyFeatureDisplayData1783, errBTAssemblyFeatureDisplayData1783 := json.Marshal(o.BTAssemblyFeatureDisplayData1783)
	if errBTAssemblyFeatureDisplayData1783 != nil {
		return []byte{}, errBTAssemblyFeatureDisplayData1783
	}
	errBTAssemblyFeatureDisplayData1783 = json.Unmarshal([]byte(serializedBTAssemblyFeatureDisplayData1783), &toSerialize)
	if errBTAssemblyFeatureDisplayData1783 != nil {
		return []byte{}, errBTAssemblyFeatureDisplayData1783
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.IsDerivedFeature != nil {
		toSerialize["isDerivedFeature"] = o.IsDerivedFeature
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.OwnerOccurrence != nil {
		toSerialize["ownerOccurrence"] = o.OwnerOccurrence
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.MateConnectorIds != nil {
		toSerialize["mateConnectorIds"] = o.MateConnectorIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTWidthMateDisplayData2888 struct {
	value *BTWidthMateDisplayData2888
	isSet bool
}

func (v NullableBTWidthMateDisplayData2888) Get() *BTWidthMateDisplayData2888 {
	return v.value
}

func (v *NullableBTWidthMateDisplayData2888) Set(val *BTWidthMateDisplayData2888) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWidthMateDisplayData2888) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWidthMateDisplayData2888) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWidthMateDisplayData2888(val *BTWidthMateDisplayData2888) *NullableBTWidthMateDisplayData2888 {
	return &NullableBTWidthMateDisplayData2888{value: val, isSet: true}
}

func (v NullableBTWidthMateDisplayData2888) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWidthMateDisplayData2888) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
