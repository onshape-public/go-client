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

// BTGetJsonPaths1697 struct for BTGetJsonPaths1697
type BTGetJsonPaths1697 struct {
	Paths []string `json:"paths,omitempty"`
}

// NewBTGetJsonPaths1697 instantiates a new BTGetJsonPaths1697 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGetJsonPaths1697() *BTGetJsonPaths1697 {
	this := BTGetJsonPaths1697{}
	return &this
}

// NewBTGetJsonPaths1697WithDefaults instantiates a new BTGetJsonPaths1697 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGetJsonPaths1697WithDefaults() *BTGetJsonPaths1697 {
	this := BTGetJsonPaths1697{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *BTGetJsonPaths1697) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonPaths1697) GetPathsOk() ([]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *BTGetJsonPaths1697) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *BTGetJsonPaths1697) SetPaths(v []string) {
	o.Paths = v
}

func (o BTGetJsonPaths1697) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableBTGetJsonPaths1697 struct {
	value *BTGetJsonPaths1697
	isSet bool
}

func (v NullableBTGetJsonPaths1697) Get() *BTGetJsonPaths1697 {
	return v.value
}

func (v *NullableBTGetJsonPaths1697) Set(val *BTGetJsonPaths1697) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGetJsonPaths1697) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGetJsonPaths1697) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGetJsonPaths1697(val *BTGetJsonPaths1697) *NullableBTGetJsonPaths1697 {
	return &NullableBTGetJsonPaths1697{value: val, isSet: true}
}

func (v NullableBTGetJsonPaths1697) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGetJsonPaths1697) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}