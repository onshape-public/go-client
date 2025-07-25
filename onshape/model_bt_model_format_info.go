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

// BTModelFormatInfo struct for BTModelFormatInfo
type BTModelFormatInfo struct {
	// Indicates if this format could be an assembly.
	CouldBeAssembly *bool `json:"couldBeAssembly,omitempty"`
	// Name of the format.
	Name *string `json:"name,omitempty"`
	// The name of the translator for the format.
	TranslatorName *string `json:"translatorName,omitempty"`
}

// NewBTModelFormatInfo instantiates a new BTModelFormatInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTModelFormatInfo() *BTModelFormatInfo {
	this := BTModelFormatInfo{}
	return &this
}

// NewBTModelFormatInfoWithDefaults instantiates a new BTModelFormatInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTModelFormatInfoWithDefaults() *BTModelFormatInfo {
	this := BTModelFormatInfo{}
	return &this
}

// GetCouldBeAssembly returns the CouldBeAssembly field value if set, zero value otherwise.
func (o *BTModelFormatInfo) GetCouldBeAssembly() bool {
	if o == nil || o.CouldBeAssembly == nil {
		var ret bool
		return ret
	}
	return *o.CouldBeAssembly
}

// GetCouldBeAssemblyOk returns a tuple with the CouldBeAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatInfo) GetCouldBeAssemblyOk() (*bool, bool) {
	if o == nil || o.CouldBeAssembly == nil {
		return nil, false
	}
	return o.CouldBeAssembly, true
}

// HasCouldBeAssembly returns a boolean if a field has been set.
func (o *BTModelFormatInfo) HasCouldBeAssembly() bool {
	if o != nil && o.CouldBeAssembly != nil {
		return true
	}

	return false
}

// SetCouldBeAssembly gets a reference to the given bool and assigns it to the CouldBeAssembly field.
func (o *BTModelFormatInfo) SetCouldBeAssembly(v bool) {
	o.CouldBeAssembly = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTModelFormatInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTModelFormatInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTModelFormatInfo) SetName(v string) {
	o.Name = &v
}

// GetTranslatorName returns the TranslatorName field value if set, zero value otherwise.
func (o *BTModelFormatInfo) GetTranslatorName() string {
	if o == nil || o.TranslatorName == nil {
		var ret string
		return ret
	}
	return *o.TranslatorName
}

// GetTranslatorNameOk returns a tuple with the TranslatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatInfo) GetTranslatorNameOk() (*string, bool) {
	if o == nil || o.TranslatorName == nil {
		return nil, false
	}
	return o.TranslatorName, true
}

// HasTranslatorName returns a boolean if a field has been set.
func (o *BTModelFormatInfo) HasTranslatorName() bool {
	if o != nil && o.TranslatorName != nil {
		return true
	}

	return false
}

// SetTranslatorName gets a reference to the given string and assigns it to the TranslatorName field.
func (o *BTModelFormatInfo) SetTranslatorName(v string) {
	o.TranslatorName = &v
}

func (o BTModelFormatInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CouldBeAssembly != nil {
		toSerialize["couldBeAssembly"] = o.CouldBeAssembly
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TranslatorName != nil {
		toSerialize["translatorName"] = o.TranslatorName
	}
	return json.Marshal(toSerialize)
}

type NullableBTModelFormatInfo struct {
	value *BTModelFormatInfo
	isSet bool
}

func (v NullableBTModelFormatInfo) Get() *BTModelFormatInfo {
	return v.value
}

func (v *NullableBTModelFormatInfo) Set(val *BTModelFormatInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTModelFormatInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTModelFormatInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTModelFormatInfo(val *BTModelFormatInfo) *NullableBTModelFormatInfo {
	return &NullableBTModelFormatInfo{value: val, isSet: true}
}

func (v NullableBTModelFormatInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTModelFormatInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
