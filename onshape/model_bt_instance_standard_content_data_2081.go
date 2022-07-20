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

// BTInstanceStandardContentData2081 struct for BTInstanceStandardContentData2081
type BTInstanceStandardContentData2081 struct {
	BtType       *string `json:"btType,omitempty"`
	ParametersId *string `json:"parametersId,omitempty"`
}

// NewBTInstanceStandardContentData2081 instantiates a new BTInstanceStandardContentData2081 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInstanceStandardContentData2081() *BTInstanceStandardContentData2081 {
	this := BTInstanceStandardContentData2081{}
	return &this
}

// NewBTInstanceStandardContentData2081WithDefaults instantiates a new BTInstanceStandardContentData2081 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInstanceStandardContentData2081WithDefaults() *BTInstanceStandardContentData2081 {
	this := BTInstanceStandardContentData2081{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInstanceStandardContentData2081) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceStandardContentData2081) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInstanceStandardContentData2081) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInstanceStandardContentData2081) SetBtType(v string) {
	o.BtType = &v
}

// GetParametersId returns the ParametersId field value if set, zero value otherwise.
func (o *BTInstanceStandardContentData2081) GetParametersId() string {
	if o == nil || o.ParametersId == nil {
		var ret string
		return ret
	}
	return *o.ParametersId
}

// GetParametersIdOk returns a tuple with the ParametersId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceStandardContentData2081) GetParametersIdOk() (*string, bool) {
	if o == nil || o.ParametersId == nil {
		return nil, false
	}
	return o.ParametersId, true
}

// HasParametersId returns a boolean if a field has been set.
func (o *BTInstanceStandardContentData2081) HasParametersId() bool {
	if o != nil && o.ParametersId != nil {
		return true
	}

	return false
}

// SetParametersId gets a reference to the given string and assigns it to the ParametersId field.
func (o *BTInstanceStandardContentData2081) SetParametersId(v string) {
	o.ParametersId = &v
}

func (o BTInstanceStandardContentData2081) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ParametersId != nil {
		toSerialize["parametersId"] = o.ParametersId
	}
	return json.Marshal(toSerialize)
}

type NullableBTInstanceStandardContentData2081 struct {
	value *BTInstanceStandardContentData2081
	isSet bool
}

func (v NullableBTInstanceStandardContentData2081) Get() *BTInstanceStandardContentData2081 {
	return v.value
}

func (v *NullableBTInstanceStandardContentData2081) Set(val *BTInstanceStandardContentData2081) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInstanceStandardContentData2081) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInstanceStandardContentData2081) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInstanceStandardContentData2081(val *BTInstanceStandardContentData2081) *NullableBTInstanceStandardContentData2081 {
	return &NullableBTInstanceStandardContentData2081{value: val, isSet: true}
}

func (v NullableBTInstanceStandardContentData2081) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInstanceStandardContentData2081) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}