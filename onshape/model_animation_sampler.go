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

// AnimationSampler struct for AnimationSampler
type AnimationSampler struct {
	Extensions    map[string]interface{} `json:"extensions,omitempty"`
	Extras        map[string]interface{} `json:"extras,omitempty"`
	Input         *int32                 `json:"input,omitempty"`
	Interpolation *string                `json:"interpolation,omitempty"`
	Output        *int32                 `json:"output,omitempty"`
}

// NewAnimationSampler instantiates a new AnimationSampler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimationSampler() *AnimationSampler {
	this := AnimationSampler{}
	return &this
}

// NewAnimationSamplerWithDefaults instantiates a new AnimationSampler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimationSamplerWithDefaults() *AnimationSampler {
	this := AnimationSampler{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AnimationSampler) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationSampler) GetExtensionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AnimationSampler) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *AnimationSampler) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *AnimationSampler) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationSampler) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *AnimationSampler) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *AnimationSampler) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *AnimationSampler) GetInput() int32 {
	if o == nil || o.Input == nil {
		var ret int32
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationSampler) GetInputOk() (*int32, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *AnimationSampler) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given int32 and assigns it to the Input field.
func (o *AnimationSampler) SetInput(v int32) {
	o.Input = &v
}

// GetInterpolation returns the Interpolation field value if set, zero value otherwise.
func (o *AnimationSampler) GetInterpolation() string {
	if o == nil || o.Interpolation == nil {
		var ret string
		return ret
	}
	return *o.Interpolation
}

// GetInterpolationOk returns a tuple with the Interpolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationSampler) GetInterpolationOk() (*string, bool) {
	if o == nil || o.Interpolation == nil {
		return nil, false
	}
	return o.Interpolation, true
}

// HasInterpolation returns a boolean if a field has been set.
func (o *AnimationSampler) HasInterpolation() bool {
	if o != nil && o.Interpolation != nil {
		return true
	}

	return false
}

// SetInterpolation gets a reference to the given string and assigns it to the Interpolation field.
func (o *AnimationSampler) SetInterpolation(v string) {
	o.Interpolation = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *AnimationSampler) GetOutput() int32 {
	if o == nil || o.Output == nil {
		var ret int32
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnimationSampler) GetOutputOk() (*int32, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *AnimationSampler) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given int32 and assigns it to the Output field.
func (o *AnimationSampler) SetOutput(v int32) {
	o.Output = &v
}

func (o AnimationSampler) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Interpolation != nil {
		toSerialize["interpolation"] = o.Interpolation
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	return json.Marshal(toSerialize)
}

type NullableAnimationSampler struct {
	value *AnimationSampler
	isSet bool
}

func (v NullableAnimationSampler) Get() *AnimationSampler {
	return v.value
}

func (v *NullableAnimationSampler) Set(val *AnimationSampler) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimationSampler) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimationSampler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimationSampler(val *AnimationSampler) *NullableAnimationSampler {
	return &NullableAnimationSampler{value: val, isSet: true}
}

func (v NullableAnimationSampler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimationSampler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
