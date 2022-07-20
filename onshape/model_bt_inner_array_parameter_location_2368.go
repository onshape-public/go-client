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

// BTInnerArrayParameterLocation2368 struct for BTInnerArrayParameterLocation2368
type BTInnerArrayParameterLocation2368 struct {
	BtType           *string `json:"btType,omitempty"`
	Index            *int32  `json:"index,omitempty"`
	OuterParameterId *string `json:"outerParameterId,omitempty"`
}

// NewBTInnerArrayParameterLocation2368 instantiates a new BTInnerArrayParameterLocation2368 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInnerArrayParameterLocation2368() *BTInnerArrayParameterLocation2368 {
	this := BTInnerArrayParameterLocation2368{}
	return &this
}

// NewBTInnerArrayParameterLocation2368WithDefaults instantiates a new BTInnerArrayParameterLocation2368 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInnerArrayParameterLocation2368WithDefaults() *BTInnerArrayParameterLocation2368 {
	this := BTInnerArrayParameterLocation2368{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInnerArrayParameterLocation2368) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInnerArrayParameterLocation2368) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInnerArrayParameterLocation2368) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInnerArrayParameterLocation2368) SetBtType(v string) {
	o.BtType = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BTInnerArrayParameterLocation2368) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInnerArrayParameterLocation2368) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BTInnerArrayParameterLocation2368) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *BTInnerArrayParameterLocation2368) SetIndex(v int32) {
	o.Index = &v
}

// GetOuterParameterId returns the OuterParameterId field value if set, zero value otherwise.
func (o *BTInnerArrayParameterLocation2368) GetOuterParameterId() string {
	if o == nil || o.OuterParameterId == nil {
		var ret string
		return ret
	}
	return *o.OuterParameterId
}

// GetOuterParameterIdOk returns a tuple with the OuterParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInnerArrayParameterLocation2368) GetOuterParameterIdOk() (*string, bool) {
	if o == nil || o.OuterParameterId == nil {
		return nil, false
	}
	return o.OuterParameterId, true
}

// HasOuterParameterId returns a boolean if a field has been set.
func (o *BTInnerArrayParameterLocation2368) HasOuterParameterId() bool {
	if o != nil && o.OuterParameterId != nil {
		return true
	}

	return false
}

// SetOuterParameterId gets a reference to the given string and assigns it to the OuterParameterId field.
func (o *BTInnerArrayParameterLocation2368) SetOuterParameterId(v string) {
	o.OuterParameterId = &v
}

func (o BTInnerArrayParameterLocation2368) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.OuterParameterId != nil {
		toSerialize["outerParameterId"] = o.OuterParameterId
	}
	return json.Marshal(toSerialize)
}

type NullableBTInnerArrayParameterLocation2368 struct {
	value *BTInnerArrayParameterLocation2368
	isSet bool
}

func (v NullableBTInnerArrayParameterLocation2368) Get() *BTInnerArrayParameterLocation2368 {
	return v.value
}

func (v *NullableBTInnerArrayParameterLocation2368) Set(val *BTInnerArrayParameterLocation2368) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInnerArrayParameterLocation2368) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInnerArrayParameterLocation2368) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInnerArrayParameterLocation2368(val *BTInnerArrayParameterLocation2368) *NullableBTInnerArrayParameterLocation2368 {
	return &NullableBTInnerArrayParameterLocation2368{value: val, isSet: true}
}

func (v NullableBTInnerArrayParameterLocation2368) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInnerArrayParameterLocation2368) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}