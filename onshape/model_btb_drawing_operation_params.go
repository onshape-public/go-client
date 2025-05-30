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

// BTBDrawingOperationParams A single drawing entity creation or modification definition
type BTBDrawingOperationParams struct {
	// Operation description
	Description *string `json:"description,omitempty"`
	// Version of drawing entity format.
	FormatVersion string `json:"formatVersion"`
	// Type of drawing modification operation: `onshapeCreateAnnotations` | `onshapeEditAnnotations`
	MessageName string `json:"messageName"`
}

// NewBTBDrawingOperationParams instantiates a new BTBDrawingOperationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBDrawingOperationParams(formatVersion string, messageName string) *BTBDrawingOperationParams {
	this := BTBDrawingOperationParams{}
	this.FormatVersion = formatVersion
	this.MessageName = messageName
	return &this
}

// NewBTBDrawingOperationParamsWithDefaults instantiates a new BTBDrawingOperationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBDrawingOperationParamsWithDefaults() *BTBDrawingOperationParams {
	this := BTBDrawingOperationParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTBDrawingOperationParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBDrawingOperationParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTBDrawingOperationParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTBDrawingOperationParams) SetDescription(v string) {
	o.Description = &v
}

// GetFormatVersion returns the FormatVersion field value
func (o *BTBDrawingOperationParams) GetFormatVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value
// and a boolean to check if the value has been set.
func (o *BTBDrawingOperationParams) GetFormatVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormatVersion, true
}

// SetFormatVersion sets field value
func (o *BTBDrawingOperationParams) SetFormatVersion(v string) {
	o.FormatVersion = v
}

// GetMessageName returns the MessageName field value
func (o *BTBDrawingOperationParams) GetMessageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageName
}

// GetMessageNameOk returns a tuple with the MessageName field value
// and a boolean to check if the value has been set.
func (o *BTBDrawingOperationParams) GetMessageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageName, true
}

// SetMessageName sets field value
func (o *BTBDrawingOperationParams) SetMessageName(v string) {
	o.MessageName = v
}

func (o BTBDrawingOperationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["formatVersion"] = o.FormatVersion
	}
	if true {
		toSerialize["messageName"] = o.MessageName
	}
	return json.Marshal(toSerialize)
}

type NullableBTBDrawingOperationParams struct {
	value *BTBDrawingOperationParams
	isSet bool
}

func (v NullableBTBDrawingOperationParams) Get() *BTBDrawingOperationParams {
	return v.value
}

func (v *NullableBTBDrawingOperationParams) Set(val *BTBDrawingOperationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBDrawingOperationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBDrawingOperationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBDrawingOperationParams(val *BTBDrawingOperationParams) *NullableBTBDrawingOperationParams {
	return &NullableBTBDrawingOperationParams{value: val, isSet: true}
}

func (v NullableBTBDrawingOperationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBDrawingOperationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
