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

// BTEditingLogic2350 struct for BTEditingLogic2350
type BTEditingLogic2350 struct {
	FunctionName             *string `json:"functionName,omitempty"`
	WantsHiddenBodies        *bool   `json:"wantsHiddenBodies,omitempty"`
	WantsIsCreating          *bool   `json:"wantsIsCreating,omitempty"`
	WantsSpecifiedParameters *bool   `json:"wantsSpecifiedParameters,omitempty"`
}

// NewBTEditingLogic2350 instantiates a new BTEditingLogic2350 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEditingLogic2350() *BTEditingLogic2350 {
	this := BTEditingLogic2350{}
	return &this
}

// NewBTEditingLogic2350WithDefaults instantiates a new BTEditingLogic2350 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEditingLogic2350WithDefaults() *BTEditingLogic2350 {
	this := BTEditingLogic2350{}
	return &this
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *BTEditingLogic2350) GetFunctionName() string {
	if o == nil || o.FunctionName == nil {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEditingLogic2350) GetFunctionNameOk() (*string, bool) {
	if o == nil || o.FunctionName == nil {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *BTEditingLogic2350) HasFunctionName() bool {
	if o != nil && o.FunctionName != nil {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *BTEditingLogic2350) SetFunctionName(v string) {
	o.FunctionName = &v
}

// GetWantsHiddenBodies returns the WantsHiddenBodies field value if set, zero value otherwise.
func (o *BTEditingLogic2350) GetWantsHiddenBodies() bool {
	if o == nil || o.WantsHiddenBodies == nil {
		var ret bool
		return ret
	}
	return *o.WantsHiddenBodies
}

// GetWantsHiddenBodiesOk returns a tuple with the WantsHiddenBodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEditingLogic2350) GetWantsHiddenBodiesOk() (*bool, bool) {
	if o == nil || o.WantsHiddenBodies == nil {
		return nil, false
	}
	return o.WantsHiddenBodies, true
}

// HasWantsHiddenBodies returns a boolean if a field has been set.
func (o *BTEditingLogic2350) HasWantsHiddenBodies() bool {
	if o != nil && o.WantsHiddenBodies != nil {
		return true
	}

	return false
}

// SetWantsHiddenBodies gets a reference to the given bool and assigns it to the WantsHiddenBodies field.
func (o *BTEditingLogic2350) SetWantsHiddenBodies(v bool) {
	o.WantsHiddenBodies = &v
}

// GetWantsIsCreating returns the WantsIsCreating field value if set, zero value otherwise.
func (o *BTEditingLogic2350) GetWantsIsCreating() bool {
	if o == nil || o.WantsIsCreating == nil {
		var ret bool
		return ret
	}
	return *o.WantsIsCreating
}

// GetWantsIsCreatingOk returns a tuple with the WantsIsCreating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEditingLogic2350) GetWantsIsCreatingOk() (*bool, bool) {
	if o == nil || o.WantsIsCreating == nil {
		return nil, false
	}
	return o.WantsIsCreating, true
}

// HasWantsIsCreating returns a boolean if a field has been set.
func (o *BTEditingLogic2350) HasWantsIsCreating() bool {
	if o != nil && o.WantsIsCreating != nil {
		return true
	}

	return false
}

// SetWantsIsCreating gets a reference to the given bool and assigns it to the WantsIsCreating field.
func (o *BTEditingLogic2350) SetWantsIsCreating(v bool) {
	o.WantsIsCreating = &v
}

// GetWantsSpecifiedParameters returns the WantsSpecifiedParameters field value if set, zero value otherwise.
func (o *BTEditingLogic2350) GetWantsSpecifiedParameters() bool {
	if o == nil || o.WantsSpecifiedParameters == nil {
		var ret bool
		return ret
	}
	return *o.WantsSpecifiedParameters
}

// GetWantsSpecifiedParametersOk returns a tuple with the WantsSpecifiedParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEditingLogic2350) GetWantsSpecifiedParametersOk() (*bool, bool) {
	if o == nil || o.WantsSpecifiedParameters == nil {
		return nil, false
	}
	return o.WantsSpecifiedParameters, true
}

// HasWantsSpecifiedParameters returns a boolean if a field has been set.
func (o *BTEditingLogic2350) HasWantsSpecifiedParameters() bool {
	if o != nil && o.WantsSpecifiedParameters != nil {
		return true
	}

	return false
}

// SetWantsSpecifiedParameters gets a reference to the given bool and assigns it to the WantsSpecifiedParameters field.
func (o *BTEditingLogic2350) SetWantsSpecifiedParameters(v bool) {
	o.WantsSpecifiedParameters = &v
}

func (o BTEditingLogic2350) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FunctionName != nil {
		toSerialize["functionName"] = o.FunctionName
	}
	if o.WantsHiddenBodies != nil {
		toSerialize["wantsHiddenBodies"] = o.WantsHiddenBodies
	}
	if o.WantsIsCreating != nil {
		toSerialize["wantsIsCreating"] = o.WantsIsCreating
	}
	if o.WantsSpecifiedParameters != nil {
		toSerialize["wantsSpecifiedParameters"] = o.WantsSpecifiedParameters
	}
	return json.Marshal(toSerialize)
}

type NullableBTEditingLogic2350 struct {
	value *BTEditingLogic2350
	isSet bool
}

func (v NullableBTEditingLogic2350) Get() *BTEditingLogic2350 {
	return v.value
}

func (v *NullableBTEditingLogic2350) Set(val *BTEditingLogic2350) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEditingLogic2350) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEditingLogic2350) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEditingLogic2350(val *BTEditingLogic2350) *NullableBTEditingLogic2350 {
	return &NullableBTEditingLogic2350{value: val, isSet: true}
}

func (v NullableBTEditingLogic2350) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEditingLogic2350) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}