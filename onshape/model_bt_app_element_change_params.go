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

// BTAppElementChangeParams struct for BTAppElementChangeParams
type BTAppElementChangeParams struct {
	// This base64-encoded data is treated as a full representation of the sub-element's data and all deltas previously added will no longer be returned.
	BaseContent *string `json:"baseContent,omitempty"`
	// This base64-encoded data is a delta which the application can apply to the last added baseContent data.
	Delta *string `json:"delta,omitempty"`
	// The id of the subelement to edit. The value is determined by the application.
	SubelementId string `json:"subelementId"`
}

// NewBTAppElementChangeParams instantiates a new BTAppElementChangeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementChangeParams(subelementId string) *BTAppElementChangeParams {
	this := BTAppElementChangeParams{}
	this.SubelementId = subelementId
	return &this
}

// NewBTAppElementChangeParamsWithDefaults instantiates a new BTAppElementChangeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementChangeParamsWithDefaults() *BTAppElementChangeParams {
	this := BTAppElementChangeParams{}
	return &this
}

// GetBaseContent returns the BaseContent field value if set, zero value otherwise.
func (o *BTAppElementChangeParams) GetBaseContent() string {
	if o == nil || o.BaseContent == nil {
		var ret string
		return ret
	}
	return *o.BaseContent
}

// GetBaseContentOk returns a tuple with the BaseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementChangeParams) GetBaseContentOk() (*string, bool) {
	if o == nil || o.BaseContent == nil {
		return nil, false
	}
	return o.BaseContent, true
}

// HasBaseContent returns a boolean if a field has been set.
func (o *BTAppElementChangeParams) HasBaseContent() bool {
	if o != nil && o.BaseContent != nil {
		return true
	}

	return false
}

// SetBaseContent gets a reference to the given string and assigns it to the BaseContent field.
func (o *BTAppElementChangeParams) SetBaseContent(v string) {
	o.BaseContent = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *BTAppElementChangeParams) GetDelta() string {
	if o == nil || o.Delta == nil {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementChangeParams) GetDeltaOk() (*string, bool) {
	if o == nil || o.Delta == nil {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *BTAppElementChangeParams) HasDelta() bool {
	if o != nil && o.Delta != nil {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *BTAppElementChangeParams) SetDelta(v string) {
	o.Delta = &v
}

// GetSubelementId returns the SubelementId field value
func (o *BTAppElementChangeParams) GetSubelementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubelementId
}

// GetSubelementIdOk returns a tuple with the SubelementId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementChangeParams) GetSubelementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubelementId, true
}

// SetSubelementId sets field value
func (o *BTAppElementChangeParams) SetSubelementId(v string) {
	o.SubelementId = v
}

func (o BTAppElementChangeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseContent != nil {
		toSerialize["baseContent"] = o.BaseContent
	}
	if o.Delta != nil {
		toSerialize["delta"] = o.Delta
	}
	if true {
		toSerialize["subelementId"] = o.SubelementId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementChangeParams struct {
	value *BTAppElementChangeParams
	isSet bool
}

func (v NullableBTAppElementChangeParams) Get() *BTAppElementChangeParams {
	return v.value
}

func (v *NullableBTAppElementChangeParams) Set(val *BTAppElementChangeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementChangeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementChangeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementChangeParams(val *BTAppElementChangeParams) *NullableBTAppElementChangeParams {
	return &NullableBTAppElementChangeParams{value: val, isSet: true}
}

func (v NullableBTAppElementChangeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementChangeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
