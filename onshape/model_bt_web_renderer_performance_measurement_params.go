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

// BTWebRendererPerformanceMeasurementParams struct for BTWebRendererPerformanceMeasurementParams
type BTWebRendererPerformanceMeasurementParams struct {
	LinesPerSecond     *float32 `json:"linesPerSecond,omitempty"`
	Renderer           *string  `json:"renderer,omitempty"`
	TrianglesPerSecond *float32 `json:"trianglesPerSecond,omitempty"`
	Vendor             *string  `json:"vendor,omitempty"`
}

// NewBTWebRendererPerformanceMeasurementParams instantiates a new BTWebRendererPerformanceMeasurementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebRendererPerformanceMeasurementParams() *BTWebRendererPerformanceMeasurementParams {
	this := BTWebRendererPerformanceMeasurementParams{}
	return &this
}

// NewBTWebRendererPerformanceMeasurementParamsWithDefaults instantiates a new BTWebRendererPerformanceMeasurementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebRendererPerformanceMeasurementParamsWithDefaults() *BTWebRendererPerformanceMeasurementParams {
	this := BTWebRendererPerformanceMeasurementParams{}
	return &this
}

// GetLinesPerSecond returns the LinesPerSecond field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetLinesPerSecond() float32 {
	if o == nil || o.LinesPerSecond == nil {
		var ret float32
		return ret
	}
	return *o.LinesPerSecond
}

// GetLinesPerSecondOk returns a tuple with the LinesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetLinesPerSecondOk() (*float32, bool) {
	if o == nil || o.LinesPerSecond == nil {
		return nil, false
	}
	return o.LinesPerSecond, true
}

// HasLinesPerSecond returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasLinesPerSecond() bool {
	if o != nil && o.LinesPerSecond != nil {
		return true
	}

	return false
}

// SetLinesPerSecond gets a reference to the given float32 and assigns it to the LinesPerSecond field.
func (o *BTWebRendererPerformanceMeasurementParams) SetLinesPerSecond(v float32) {
	o.LinesPerSecond = &v
}

// GetRenderer returns the Renderer field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetRenderer() string {
	if o == nil || o.Renderer == nil {
		var ret string
		return ret
	}
	return *o.Renderer
}

// GetRendererOk returns a tuple with the Renderer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetRendererOk() (*string, bool) {
	if o == nil || o.Renderer == nil {
		return nil, false
	}
	return o.Renderer, true
}

// HasRenderer returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasRenderer() bool {
	if o != nil && o.Renderer != nil {
		return true
	}

	return false
}

// SetRenderer gets a reference to the given string and assigns it to the Renderer field.
func (o *BTWebRendererPerformanceMeasurementParams) SetRenderer(v string) {
	o.Renderer = &v
}

// GetTrianglesPerSecond returns the TrianglesPerSecond field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetTrianglesPerSecond() float32 {
	if o == nil || o.TrianglesPerSecond == nil {
		var ret float32
		return ret
	}
	return *o.TrianglesPerSecond
}

// GetTrianglesPerSecondOk returns a tuple with the TrianglesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetTrianglesPerSecondOk() (*float32, bool) {
	if o == nil || o.TrianglesPerSecond == nil {
		return nil, false
	}
	return o.TrianglesPerSecond, true
}

// HasTrianglesPerSecond returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasTrianglesPerSecond() bool {
	if o != nil && o.TrianglesPerSecond != nil {
		return true
	}

	return false
}

// SetTrianglesPerSecond gets a reference to the given float32 and assigns it to the TrianglesPerSecond field.
func (o *BTWebRendererPerformanceMeasurementParams) SetTrianglesPerSecond(v float32) {
	o.TrianglesPerSecond = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BTWebRendererPerformanceMeasurementParams) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebRendererPerformanceMeasurementParams) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BTWebRendererPerformanceMeasurementParams) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BTWebRendererPerformanceMeasurementParams) SetVendor(v string) {
	o.Vendor = &v
}

func (o BTWebRendererPerformanceMeasurementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinesPerSecond != nil {
		toSerialize["linesPerSecond"] = o.LinesPerSecond
	}
	if o.Renderer != nil {
		toSerialize["renderer"] = o.Renderer
	}
	if o.TrianglesPerSecond != nil {
		toSerialize["trianglesPerSecond"] = o.TrianglesPerSecond
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebRendererPerformanceMeasurementParams struct {
	value *BTWebRendererPerformanceMeasurementParams
	isSet bool
}

func (v NullableBTWebRendererPerformanceMeasurementParams) Get() *BTWebRendererPerformanceMeasurementParams {
	return v.value
}

func (v *NullableBTWebRendererPerformanceMeasurementParams) Set(val *BTWebRendererPerformanceMeasurementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebRendererPerformanceMeasurementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebRendererPerformanceMeasurementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebRendererPerformanceMeasurementParams(val *BTWebRendererPerformanceMeasurementParams) *NullableBTWebRendererPerformanceMeasurementParams {
	return &NullableBTWebRendererPerformanceMeasurementParams{value: val, isSet: true}
}

func (v NullableBTWebRendererPerformanceMeasurementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebRendererPerformanceMeasurementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
