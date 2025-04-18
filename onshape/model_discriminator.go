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

// Discriminator struct for Discriminator
type Discriminator struct {
	Extensions   map[string]interface{} `json:"extensions,omitempty"`
	Mapping      *map[string]string     `json:"mapping,omitempty"`
	PropertyName *string                `json:"propertyName,omitempty"`
}

// NewDiscriminator instantiates a new Discriminator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscriminator() *Discriminator {
	this := Discriminator{}
	return &this
}

// NewDiscriminatorWithDefaults instantiates a new Discriminator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscriminatorWithDefaults() *Discriminator {
	this := Discriminator{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Discriminator) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discriminator) GetExtensionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Discriminator) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *Discriminator) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *Discriminator) GetMapping() map[string]string {
	if o == nil || o.Mapping == nil {
		var ret map[string]string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discriminator) GetMappingOk() (*map[string]string, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *Discriminator) HasMapping() bool {
	if o != nil && o.Mapping != nil {
		return true
	}

	return false
}

// SetMapping gets a reference to the given map[string]string and assigns it to the Mapping field.
func (o *Discriminator) SetMapping(v map[string]string) {
	o.Mapping = &v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *Discriminator) GetPropertyName() string {
	if o == nil || o.PropertyName == nil {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discriminator) GetPropertyNameOk() (*string, bool) {
	if o == nil || o.PropertyName == nil {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *Discriminator) HasPropertyName() bool {
	if o != nil && o.PropertyName != nil {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *Discriminator) SetPropertyName(v string) {
	o.PropertyName = &v
}

func (o Discriminator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Mapping != nil {
		toSerialize["mapping"] = o.Mapping
	}
	if o.PropertyName != nil {
		toSerialize["propertyName"] = o.PropertyName
	}
	return json.Marshal(toSerialize)
}

type NullableDiscriminator struct {
	value *Discriminator
	isSet bool
}

func (v NullableDiscriminator) Get() *Discriminator {
	return v.value
}

func (v *NullableDiscriminator) Set(val *Discriminator) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscriminator) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscriminator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscriminator(val *Discriminator) *NullableDiscriminator {
	return &NullableDiscriminator{value: val, isSet: true}
}

func (v NullableDiscriminator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscriminator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
