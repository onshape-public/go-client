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

// OpenAPIPaths struct for OpenAPIPaths
type OpenAPIPaths struct {
	Extensions           map[string]interface{} `json:"extensions,omitempty"`
	Empty                *bool                  `json:"empty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenAPIPaths OpenAPIPaths

// NewOpenAPIPaths instantiates a new OpenAPIPaths object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAPIPaths() *OpenAPIPaths {
	this := OpenAPIPaths{}
	return &this
}

// NewOpenAPIPathsWithDefaults instantiates a new OpenAPIPaths object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAPIPathsWithDefaults() *OpenAPIPaths {
	this := OpenAPIPaths{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *OpenAPIPaths) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPIPaths) GetExtensionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *OpenAPIPaths) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *OpenAPIPaths) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *OpenAPIPaths) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPIPaths) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *OpenAPIPaths) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *OpenAPIPaths) SetEmpty(v bool) {
	o.Empty = &v
}

func (o OpenAPIPaths) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenAPIPaths) UnmarshalJSON(bytes []byte) (err error) {
	varOpenAPIPaths := _OpenAPIPaths{}

	if err = json.Unmarshal(bytes, &varOpenAPIPaths); err == nil {
		*o = OpenAPIPaths(varOpenAPIPaths)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "empty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenAPIPaths struct {
	value *OpenAPIPaths
	isSet bool
}

func (v NullableOpenAPIPaths) Get() *OpenAPIPaths {
	return v.value
}

func (v *NullableOpenAPIPaths) Set(val *OpenAPIPaths) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAPIPaths) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAPIPaths) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAPIPaths(val *OpenAPIPaths) *NullableOpenAPIPaths {
	return &NullableOpenAPIPaths{value: val, isSet: true}
}

func (v NullableOpenAPIPaths) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAPIPaths) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
