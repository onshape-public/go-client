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

// BTJEditInsert2523 Inserts a value using the specified path.
type BTJEditInsert2523 struct {
	BtType *string                 `json:"btType,omitempty"`
	Path   *BTJPath3073            `json:"path,omitempty"`
	Value  *map[string]interface{} `json:"value,omitempty"`
}

// NewBTJEditInsert2523 instantiates a new BTJEditInsert2523 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJEditInsert2523() *BTJEditInsert2523 {
	this := BTJEditInsert2523{}
	return &this
}

// NewBTJEditInsert2523WithDefaults instantiates a new BTJEditInsert2523 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJEditInsert2523WithDefaults() *BTJEditInsert2523 {
	this := BTJEditInsert2523{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJEditInsert2523) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditInsert2523) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJEditInsert2523) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJEditInsert2523) SetBtType(v string) {
	o.BtType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTJEditInsert2523) GetPath() BTJPath3073 {
	if o == nil || o.Path == nil {
		var ret BTJPath3073
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditInsert2523) GetPathOk() (*BTJPath3073, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTJEditInsert2523) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given BTJPath3073 and assigns it to the Path field.
func (o *BTJEditInsert2523) SetPath(v BTJPath3073) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTJEditInsert2523) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditInsert2523) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTJEditInsert2523) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *BTJEditInsert2523) SetValue(v map[string]interface{}) {
	o.Value = &v
}

func (o BTJEditInsert2523) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTJEditInsert2523 struct {
	value *BTJEditInsert2523
	isSet bool
}

func (v NullableBTJEditInsert2523) Get() *BTJEditInsert2523 {
	return v.value
}

func (v *NullableBTJEditInsert2523) Set(val *BTJEditInsert2523) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJEditInsert2523) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJEditInsert2523) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJEditInsert2523(val *BTJEditInsert2523) *NullableBTJEditInsert2523 {
	return &NullableBTJEditInsert2523{value: val, isSet: true}
}

func (v NullableBTJEditInsert2523) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJEditInsert2523) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}