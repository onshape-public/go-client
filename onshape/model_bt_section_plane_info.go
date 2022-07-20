/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSectionPlaneInfo struct for BTSectionPlaneInfo
type BTSectionPlaneInfo struct {
	Center  []float64 `json:"center,omitempty"`
	Normal  []float64 `json:"normal,omitempty"`
	Tangent []float64 `json:"tangent,omitempty"`
}

// NewBTSectionPlaneInfo instantiates a new BTSectionPlaneInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSectionPlaneInfo() *BTSectionPlaneInfo {
	this := BTSectionPlaneInfo{}
	return &this
}

// NewBTSectionPlaneInfoWithDefaults instantiates a new BTSectionPlaneInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSectionPlaneInfoWithDefaults() *BTSectionPlaneInfo {
	this := BTSectionPlaneInfo{}
	return &this
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *BTSectionPlaneInfo) GetCenter() []float64 {
	if o == nil || o.Center == nil {
		var ret []float64
		return ret
	}
	return o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSectionPlaneInfo) GetCenterOk() ([]float64, bool) {
	if o == nil || o.Center == nil {
		return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *BTSectionPlaneInfo) HasCenter() bool {
	if o != nil && o.Center != nil {
		return true
	}

	return false
}

// SetCenter gets a reference to the given []float64 and assigns it to the Center field.
func (o *BTSectionPlaneInfo) SetCenter(v []float64) {
	o.Center = v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTSectionPlaneInfo) GetNormal() []float64 {
	if o == nil || o.Normal == nil {
		var ret []float64
		return ret
	}
	return o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSectionPlaneInfo) GetNormalOk() ([]float64, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTSectionPlaneInfo) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given []float64 and assigns it to the Normal field.
func (o *BTSectionPlaneInfo) SetNormal(v []float64) {
	o.Normal = v
}

// GetTangent returns the Tangent field value if set, zero value otherwise.
func (o *BTSectionPlaneInfo) GetTangent() []float64 {
	if o == nil || o.Tangent == nil {
		var ret []float64
		return ret
	}
	return o.Tangent
}

// GetTangentOk returns a tuple with the Tangent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSectionPlaneInfo) GetTangentOk() ([]float64, bool) {
	if o == nil || o.Tangent == nil {
		return nil, false
	}
	return o.Tangent, true
}

// HasTangent returns a boolean if a field has been set.
func (o *BTSectionPlaneInfo) HasTangent() bool {
	if o != nil && o.Tangent != nil {
		return true
	}

	return false
}

// SetTangent gets a reference to the given []float64 and assigns it to the Tangent field.
func (o *BTSectionPlaneInfo) SetTangent(v []float64) {
	o.Tangent = v
}

func (o BTSectionPlaneInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Center != nil {
		toSerialize["center"] = o.Center
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	if o.Tangent != nil {
		toSerialize["tangent"] = o.Tangent
	}
	return json.Marshal(toSerialize)
}

type NullableBTSectionPlaneInfo struct {
	value *BTSectionPlaneInfo
	isSet bool
}

func (v NullableBTSectionPlaneInfo) Get() *BTSectionPlaneInfo {
	return v.value
}

func (v *NullableBTSectionPlaneInfo) Set(val *BTSectionPlaneInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSectionPlaneInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSectionPlaneInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSectionPlaneInfo(val *BTSectionPlaneInfo) *NullableBTSectionPlaneInfo {
	return &NullableBTSectionPlaneInfo{value: val, isSet: true}
}

func (v NullableBTSectionPlaneInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSectionPlaneInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}