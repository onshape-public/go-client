/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24257-687de06de652
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// AccessorSparseValues struct for AccessorSparseValues
type AccessorSparseValues struct {
	BufferView *int32                            `json:"bufferView,omitempty"`
	ByteOffset *int32                            `json:"byteOffset,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
}

// NewAccessorSparseValues instantiates a new AccessorSparseValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessorSparseValues() *AccessorSparseValues {
	this := AccessorSparseValues{}
	return &this
}

// NewAccessorSparseValuesWithDefaults instantiates a new AccessorSparseValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessorSparseValuesWithDefaults() *AccessorSparseValues {
	this := AccessorSparseValues{}
	return &this
}

// GetBufferView returns the BufferView field value if set, zero value otherwise.
func (o *AccessorSparseValues) GetBufferView() int32 {
	if o == nil || o.BufferView == nil {
		var ret int32
		return ret
	}
	return *o.BufferView
}

// GetBufferViewOk returns a tuple with the BufferView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseValues) GetBufferViewOk() (*int32, bool) {
	if o == nil || o.BufferView == nil {
		return nil, false
	}
	return o.BufferView, true
}

// HasBufferView returns a boolean if a field has been set.
func (o *AccessorSparseValues) HasBufferView() bool {
	if o != nil && o.BufferView != nil {
		return true
	}

	return false
}

// SetBufferView gets a reference to the given int32 and assigns it to the BufferView field.
func (o *AccessorSparseValues) SetBufferView(v int32) {
	o.BufferView = &v
}

// GetByteOffset returns the ByteOffset field value if set, zero value otherwise.
func (o *AccessorSparseValues) GetByteOffset() int32 {
	if o == nil || o.ByteOffset == nil {
		var ret int32
		return ret
	}
	return *o.ByteOffset
}

// GetByteOffsetOk returns a tuple with the ByteOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseValues) GetByteOffsetOk() (*int32, bool) {
	if o == nil || o.ByteOffset == nil {
		return nil, false
	}
	return o.ByteOffset, true
}

// HasByteOffset returns a boolean if a field has been set.
func (o *AccessorSparseValues) HasByteOffset() bool {
	if o != nil && o.ByteOffset != nil {
		return true
	}

	return false
}

// SetByteOffset gets a reference to the given int32 and assigns it to the ByteOffset field.
func (o *AccessorSparseValues) SetByteOffset(v int32) {
	o.ByteOffset = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AccessorSparseValues) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseValues) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AccessorSparseValues) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *AccessorSparseValues) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *AccessorSparseValues) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseValues) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *AccessorSparseValues) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *AccessorSparseValues) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

func (o AccessorSparseValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BufferView != nil {
		toSerialize["bufferView"] = o.BufferView
	}
	if o.ByteOffset != nil {
		toSerialize["byteOffset"] = o.ByteOffset
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	return json.Marshal(toSerialize)
}

type NullableAccessorSparseValues struct {
	value *AccessorSparseValues
	isSet bool
}

func (v NullableAccessorSparseValues) Get() *AccessorSparseValues {
	return v.value
}

func (v *NullableAccessorSparseValues) Set(val *AccessorSparseValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessorSparseValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessorSparseValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessorSparseValues(val *AccessorSparseValues) *NullableAccessorSparseValues {
	return &NullableAccessorSparseValues{value: val, isSet: true}
}

func (v NullableAccessorSparseValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessorSparseValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}