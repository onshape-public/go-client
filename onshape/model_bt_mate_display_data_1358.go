/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6820-1bef41f9cc03
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMateDisplayData1358 struct for BTMateDisplayData1358
type BTMateDisplayData1358 struct {
	BtType           *string         `json:"btType,omitempty"`
	Hidden           *bool           `json:"hidden,omitempty"`
	IsDerivedFeature *bool           `json:"isDerivedFeature,omitempty"`
	MateConnectorIds []string        `json:"mateConnectorIds,omitempty"`
	MateType         *string         `json:"mateType,omitempty"`
	NodeId           *string         `json:"nodeId,omitempty"`
	OwnerOccurrence  *BTOccurrence74 `json:"ownerOccurrence,omitempty"`
	Status           *string         `json:"status,omitempty"`
}

// NewBTMateDisplayData1358 instantiates a new BTMateDisplayData1358 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateDisplayData1358() *BTMateDisplayData1358 {
	this := BTMateDisplayData1358{}
	return &this
}

// NewBTMateDisplayData1358WithDefaults instantiates a new BTMateDisplayData1358 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateDisplayData1358WithDefaults() *BTMateDisplayData1358 {
	this := BTMateDisplayData1358{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMateDisplayData1358) SetBtType(v string) {
	o.BtType = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTMateDisplayData1358) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIsDerivedFeature returns the IsDerivedFeature field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetIsDerivedFeature() bool {
	if o == nil || o.IsDerivedFeature == nil {
		var ret bool
		return ret
	}
	return *o.IsDerivedFeature
}

// GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetIsDerivedFeatureOk() (*bool, bool) {
	if o == nil || o.IsDerivedFeature == nil {
		return nil, false
	}
	return o.IsDerivedFeature, true
}

// HasIsDerivedFeature returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasIsDerivedFeature() bool {
	if o != nil && o.IsDerivedFeature != nil {
		return true
	}

	return false
}

// SetIsDerivedFeature gets a reference to the given bool and assigns it to the IsDerivedFeature field.
func (o *BTMateDisplayData1358) SetIsDerivedFeature(v bool) {
	o.IsDerivedFeature = &v
}

// GetMateConnectorIds returns the MateConnectorIds field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetMateConnectorIds() []string {
	if o == nil || o.MateConnectorIds == nil {
		var ret []string
		return ret
	}
	return o.MateConnectorIds
}

// GetMateConnectorIdsOk returns a tuple with the MateConnectorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetMateConnectorIdsOk() ([]string, bool) {
	if o == nil || o.MateConnectorIds == nil {
		return nil, false
	}
	return o.MateConnectorIds, true
}

// HasMateConnectorIds returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasMateConnectorIds() bool {
	if o != nil && o.MateConnectorIds != nil {
		return true
	}

	return false
}

// SetMateConnectorIds gets a reference to the given []string and assigns it to the MateConnectorIds field.
func (o *BTMateDisplayData1358) SetMateConnectorIds(v []string) {
	o.MateConnectorIds = v
}

// GetMateType returns the MateType field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetMateType() string {
	if o == nil || o.MateType == nil {
		var ret string
		return ret
	}
	return *o.MateType
}

// GetMateTypeOk returns a tuple with the MateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetMateTypeOk() (*string, bool) {
	if o == nil || o.MateType == nil {
		return nil, false
	}
	return o.MateType, true
}

// HasMateType returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasMateType() bool {
	if o != nil && o.MateType != nil {
		return true
	}

	return false
}

// SetMateType gets a reference to the given string and assigns it to the MateType field.
func (o *BTMateDisplayData1358) SetMateType(v string) {
	o.MateType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMateDisplayData1358) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOwnerOccurrence returns the OwnerOccurrence field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetOwnerOccurrence() BTOccurrence74 {
	if o == nil || o.OwnerOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OwnerOccurrence
}

// GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetOwnerOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.OwnerOccurrence == nil {
		return nil, false
	}
	return o.OwnerOccurrence, true
}

// HasOwnerOccurrence returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasOwnerOccurrence() bool {
	if o != nil && o.OwnerOccurrence != nil {
		return true
	}

	return false
}

// SetOwnerOccurrence gets a reference to the given BTOccurrence74 and assigns it to the OwnerOccurrence field.
func (o *BTMateDisplayData1358) SetOwnerOccurrence(v BTOccurrence74) {
	o.OwnerOccurrence = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTMateDisplayData1358) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateDisplayData1358) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTMateDisplayData1358) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BTMateDisplayData1358) SetStatus(v string) {
	o.Status = &v
}

func (o BTMateDisplayData1358) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.IsDerivedFeature != nil {
		toSerialize["isDerivedFeature"] = o.IsDerivedFeature
	}
	if o.MateConnectorIds != nil {
		toSerialize["mateConnectorIds"] = o.MateConnectorIds
	}
	if o.MateType != nil {
		toSerialize["mateType"] = o.MateType
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
	return json.Marshal(toSerialize)
}

type NullableBTMateDisplayData1358 struct {
	value *BTMateDisplayData1358
	isSet bool
}

func (v NullableBTMateDisplayData1358) Get() *BTMateDisplayData1358 {
	return v.value
}

func (v *NullableBTMateDisplayData1358) Set(val *BTMateDisplayData1358) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateDisplayData1358) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateDisplayData1358) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateDisplayData1358(val *BTMateDisplayData1358) *NullableBTMateDisplayData1358 {
	return &NullableBTMateDisplayData1358{value: val, isSet: true}
}

func (v NullableBTMateDisplayData1358) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateDisplayData1358) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}