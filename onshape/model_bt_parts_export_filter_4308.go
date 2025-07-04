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

// BTPartsExportFilter4308 Skip mesh/curve foreign data creation in individual parts export
type BTPartsExportFilter4308 struct {
	// Type of JSON object.
	BtType          *string `json:"btType,omitempty"`
	SkipAllMesh     *bool   `json:"skipAllMesh,omitempty"`
	SkipCurves      *bool   `json:"skipCurves,omitempty"`
	SkipPartialMesh *bool   `json:"skipPartialMesh,omitempty"`
}

// NewBTPartsExportFilter4308 instantiates a new BTPartsExportFilter4308 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartsExportFilter4308() *BTPartsExportFilter4308 {
	this := BTPartsExportFilter4308{}
	return &this
}

// NewBTPartsExportFilter4308WithDefaults instantiates a new BTPartsExportFilter4308 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartsExportFilter4308WithDefaults() *BTPartsExportFilter4308 {
	this := BTPartsExportFilter4308{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartsExportFilter4308) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartsExportFilter4308) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartsExportFilter4308) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartsExportFilter4308) SetBtType(v string) {
	o.BtType = &v
}

// GetSkipAllMesh returns the SkipAllMesh field value if set, zero value otherwise.
func (o *BTPartsExportFilter4308) GetSkipAllMesh() bool {
	if o == nil || o.SkipAllMesh == nil {
		var ret bool
		return ret
	}
	return *o.SkipAllMesh
}

// GetSkipAllMeshOk returns a tuple with the SkipAllMesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartsExportFilter4308) GetSkipAllMeshOk() (*bool, bool) {
	if o == nil || o.SkipAllMesh == nil {
		return nil, false
	}
	return o.SkipAllMesh, true
}

// HasSkipAllMesh returns a boolean if a field has been set.
func (o *BTPartsExportFilter4308) HasSkipAllMesh() bool {
	if o != nil && o.SkipAllMesh != nil {
		return true
	}

	return false
}

// SetSkipAllMesh gets a reference to the given bool and assigns it to the SkipAllMesh field.
func (o *BTPartsExportFilter4308) SetSkipAllMesh(v bool) {
	o.SkipAllMesh = &v
}

// GetSkipCurves returns the SkipCurves field value if set, zero value otherwise.
func (o *BTPartsExportFilter4308) GetSkipCurves() bool {
	if o == nil || o.SkipCurves == nil {
		var ret bool
		return ret
	}
	return *o.SkipCurves
}

// GetSkipCurvesOk returns a tuple with the SkipCurves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartsExportFilter4308) GetSkipCurvesOk() (*bool, bool) {
	if o == nil || o.SkipCurves == nil {
		return nil, false
	}
	return o.SkipCurves, true
}

// HasSkipCurves returns a boolean if a field has been set.
func (o *BTPartsExportFilter4308) HasSkipCurves() bool {
	if o != nil && o.SkipCurves != nil {
		return true
	}

	return false
}

// SetSkipCurves gets a reference to the given bool and assigns it to the SkipCurves field.
func (o *BTPartsExportFilter4308) SetSkipCurves(v bool) {
	o.SkipCurves = &v
}

// GetSkipPartialMesh returns the SkipPartialMesh field value if set, zero value otherwise.
func (o *BTPartsExportFilter4308) GetSkipPartialMesh() bool {
	if o == nil || o.SkipPartialMesh == nil {
		var ret bool
		return ret
	}
	return *o.SkipPartialMesh
}

// GetSkipPartialMeshOk returns a tuple with the SkipPartialMesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartsExportFilter4308) GetSkipPartialMeshOk() (*bool, bool) {
	if o == nil || o.SkipPartialMesh == nil {
		return nil, false
	}
	return o.SkipPartialMesh, true
}

// HasSkipPartialMesh returns a boolean if a field has been set.
func (o *BTPartsExportFilter4308) HasSkipPartialMesh() bool {
	if o != nil && o.SkipPartialMesh != nil {
		return true
	}

	return false
}

// SetSkipPartialMesh gets a reference to the given bool and assigns it to the SkipPartialMesh field.
func (o *BTPartsExportFilter4308) SetSkipPartialMesh(v bool) {
	o.SkipPartialMesh = &v
}

func (o BTPartsExportFilter4308) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SkipAllMesh != nil {
		toSerialize["skipAllMesh"] = o.SkipAllMesh
	}
	if o.SkipCurves != nil {
		toSerialize["skipCurves"] = o.SkipCurves
	}
	if o.SkipPartialMesh != nil {
		toSerialize["skipPartialMesh"] = o.SkipPartialMesh
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartsExportFilter4308 struct {
	value *BTPartsExportFilter4308
	isSet bool
}

func (v NullableBTPartsExportFilter4308) Get() *BTPartsExportFilter4308 {
	return v.value
}

func (v *NullableBTPartsExportFilter4308) Set(val *BTPartsExportFilter4308) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartsExportFilter4308) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartsExportFilter4308) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartsExportFilter4308(val *BTPartsExportFilter4308) *NullableBTPartsExportFilter4308 {
	return &NullableBTPartsExportFilter4308{value: val, isSet: true}
}

func (v NullableBTPartsExportFilter4308) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartsExportFilter4308) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
