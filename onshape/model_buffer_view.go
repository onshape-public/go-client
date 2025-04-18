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

// BufferView struct for BufferView
type BufferView struct {
	Buffer     *int32                 `json:"buffer,omitempty"`
	ByteLength *int32                 `json:"byteLength,omitempty"`
	ByteOffset *int32                 `json:"byteOffset,omitempty"`
	ByteStride *int32                 `json:"byteStride,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	Extras     map[string]interface{} `json:"extras,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Target     *int32                 `json:"target,omitempty"`
}

// NewBufferView instantiates a new BufferView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBufferView() *BufferView {
	this := BufferView{}
	return &this
}

// NewBufferViewWithDefaults instantiates a new BufferView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBufferViewWithDefaults() *BufferView {
	this := BufferView{}
	return &this
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *BufferView) GetBuffer() int32 {
	if o == nil || o.Buffer == nil {
		var ret int32
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetBufferOk() (*int32, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *BufferView) HasBuffer() bool {
	if o != nil && o.Buffer != nil {
		return true
	}

	return false
}

// SetBuffer gets a reference to the given int32 and assigns it to the Buffer field.
func (o *BufferView) SetBuffer(v int32) {
	o.Buffer = &v
}

// GetByteLength returns the ByteLength field value if set, zero value otherwise.
func (o *BufferView) GetByteLength() int32 {
	if o == nil || o.ByteLength == nil {
		var ret int32
		return ret
	}
	return *o.ByteLength
}

// GetByteLengthOk returns a tuple with the ByteLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetByteLengthOk() (*int32, bool) {
	if o == nil || o.ByteLength == nil {
		return nil, false
	}
	return o.ByteLength, true
}

// HasByteLength returns a boolean if a field has been set.
func (o *BufferView) HasByteLength() bool {
	if o != nil && o.ByteLength != nil {
		return true
	}

	return false
}

// SetByteLength gets a reference to the given int32 and assigns it to the ByteLength field.
func (o *BufferView) SetByteLength(v int32) {
	o.ByteLength = &v
}

// GetByteOffset returns the ByteOffset field value if set, zero value otherwise.
func (o *BufferView) GetByteOffset() int32 {
	if o == nil || o.ByteOffset == nil {
		var ret int32
		return ret
	}
	return *o.ByteOffset
}

// GetByteOffsetOk returns a tuple with the ByteOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetByteOffsetOk() (*int32, bool) {
	if o == nil || o.ByteOffset == nil {
		return nil, false
	}
	return o.ByteOffset, true
}

// HasByteOffset returns a boolean if a field has been set.
func (o *BufferView) HasByteOffset() bool {
	if o != nil && o.ByteOffset != nil {
		return true
	}

	return false
}

// SetByteOffset gets a reference to the given int32 and assigns it to the ByteOffset field.
func (o *BufferView) SetByteOffset(v int32) {
	o.ByteOffset = &v
}

// GetByteStride returns the ByteStride field value if set, zero value otherwise.
func (o *BufferView) GetByteStride() int32 {
	if o == nil || o.ByteStride == nil {
		var ret int32
		return ret
	}
	return *o.ByteStride
}

// GetByteStrideOk returns a tuple with the ByteStride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetByteStrideOk() (*int32, bool) {
	if o == nil || o.ByteStride == nil {
		return nil, false
	}
	return o.ByteStride, true
}

// HasByteStride returns a boolean if a field has been set.
func (o *BufferView) HasByteStride() bool {
	if o != nil && o.ByteStride != nil {
		return true
	}

	return false
}

// SetByteStride gets a reference to the given int32 and assigns it to the ByteStride field.
func (o *BufferView) SetByteStride(v int32) {
	o.ByteStride = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *BufferView) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetExtensionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *BufferView) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *BufferView) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *BufferView) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *BufferView) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *BufferView) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BufferView) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BufferView) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BufferView) SetName(v string) {
	o.Name = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BufferView) GetTarget() int32 {
	if o == nil || o.Target == nil {
		var ret int32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferView) GetTargetOk() (*int32, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *BufferView) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given int32 and assigns it to the Target field.
func (o *BufferView) SetTarget(v int32) {
	o.Target = &v
}

func (o BufferView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.ByteLength != nil {
		toSerialize["byteLength"] = o.ByteLength
	}
	if o.ByteOffset != nil {
		toSerialize["byteOffset"] = o.ByteOffset
	}
	if o.ByteStride != nil {
		toSerialize["byteStride"] = o.ByteStride
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableBufferView struct {
	value *BufferView
	isSet bool
}

func (v NullableBufferView) Get() *BufferView {
	return v.value
}

func (v *NullableBufferView) Set(val *BufferView) {
	v.value = val
	v.isSet = true
}

func (v NullableBufferView) IsSet() bool {
	return v.isSet
}

func (v *NullableBufferView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBufferView(val *BufferView) *NullableBufferView {
	return &NullableBufferView{value: val, isSet: true}
}

func (v NullableBufferView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBufferView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
