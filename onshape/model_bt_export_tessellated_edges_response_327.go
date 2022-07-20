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

// BTExportTessellatedEdgesResponse327 struct for BTExportTessellatedEdgesResponse327
type BTExportTessellatedEdgesResponse327 struct {
	Bodies    []BTExportTessellatedBody3398 `json:"bodies,omitempty"`
	ErrorEnum *string                       `json:"errorEnum,omitempty"`
}

// NewBTExportTessellatedEdgesResponse327 instantiates a new BTExportTessellatedEdgesResponse327 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedEdgesResponse327() *BTExportTessellatedEdgesResponse327 {
	this := BTExportTessellatedEdgesResponse327{}
	return &this
}

// NewBTExportTessellatedEdgesResponse327WithDefaults instantiates a new BTExportTessellatedEdgesResponse327 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedEdgesResponse327WithDefaults() *BTExportTessellatedEdgesResponse327 {
	this := BTExportTessellatedEdgesResponse327{}
	return &this
}

// GetBodies returns the Bodies field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesResponse327) GetBodies() []BTExportTessellatedBody3398 {
	if o == nil || o.Bodies == nil {
		var ret []BTExportTessellatedBody3398
		return ret
	}
	return o.Bodies
}

// GetBodiesOk returns a tuple with the Bodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesResponse327) GetBodiesOk() ([]BTExportTessellatedBody3398, bool) {
	if o == nil || o.Bodies == nil {
		return nil, false
	}
	return o.Bodies, true
}

// HasBodies returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesResponse327) HasBodies() bool {
	if o != nil && o.Bodies != nil {
		return true
	}

	return false
}

// SetBodies gets a reference to the given []BTExportTessellatedBody3398 and assigns it to the Bodies field.
func (o *BTExportTessellatedEdgesResponse327) SetBodies(v []BTExportTessellatedBody3398) {
	o.Bodies = v
}

// GetErrorEnum returns the ErrorEnum field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesResponse327) GetErrorEnum() string {
	if o == nil || o.ErrorEnum == nil {
		var ret string
		return ret
	}
	return *o.ErrorEnum
}

// GetErrorEnumOk returns a tuple with the ErrorEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesResponse327) GetErrorEnumOk() (*string, bool) {
	if o == nil || o.ErrorEnum == nil {
		return nil, false
	}
	return o.ErrorEnum, true
}

// HasErrorEnum returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesResponse327) HasErrorEnum() bool {
	if o != nil && o.ErrorEnum != nil {
		return true
	}

	return false
}

// SetErrorEnum gets a reference to the given string and assigns it to the ErrorEnum field.
func (o *BTExportTessellatedEdgesResponse327) SetErrorEnum(v string) {
	o.ErrorEnum = &v
}

func (o BTExportTessellatedEdgesResponse327) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bodies != nil {
		toSerialize["bodies"] = o.Bodies
	}
	if o.ErrorEnum != nil {
		toSerialize["errorEnum"] = o.ErrorEnum
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedEdgesResponse327 struct {
	value *BTExportTessellatedEdgesResponse327
	isSet bool
}

func (v NullableBTExportTessellatedEdgesResponse327) Get() *BTExportTessellatedEdgesResponse327 {
	return v.value
}

func (v *NullableBTExportTessellatedEdgesResponse327) Set(val *BTExportTessellatedEdgesResponse327) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedEdgesResponse327) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedEdgesResponse327) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedEdgesResponse327(val *BTExportTessellatedEdgesResponse327) *NullableBTExportTessellatedEdgesResponse327 {
	return &NullableBTExportTessellatedEdgesResponse327{value: val, isSet: true}
}

func (v NullableBTExportTessellatedEdgesResponse327) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedEdgesResponse327) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}