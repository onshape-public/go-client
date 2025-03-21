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

// BTAllowedMateTypeFilter1511 struct for BTAllowedMateTypeFilter1511
type BTAllowedMateTypeFilter1511 struct {
	BTMateFilter162
	BtType               *string       `json:"btType,omitempty"`
	RequireMateQueryData *bool         `json:"requireMateQueryData,omitempty"`
	TopLevelMateOnly     *bool         `json:"topLevelMateOnly,omitempty"`
	AllowedMateTypes     []GBTMateType `json:"allowedMateTypes,omitempty"`
}

// NewBTAllowedMateTypeFilter1511 instantiates a new BTAllowedMateTypeFilter1511 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAllowedMateTypeFilter1511() *BTAllowedMateTypeFilter1511 {
	this := BTAllowedMateTypeFilter1511{}
	return &this
}

// NewBTAllowedMateTypeFilter1511WithDefaults instantiates a new BTAllowedMateTypeFilter1511 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAllowedMateTypeFilter1511WithDefaults() *BTAllowedMateTypeFilter1511 {
	this := BTAllowedMateTypeFilter1511{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAllowedMateTypeFilter1511) SetBtType(v string) {
	o.BtType = &v
}

// GetRequireMateQueryData returns the RequireMateQueryData field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetRequireMateQueryData() bool {
	if o == nil || o.RequireMateQueryData == nil {
		var ret bool
		return ret
	}
	return *o.RequireMateQueryData
}

// GetRequireMateQueryDataOk returns a tuple with the RequireMateQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetRequireMateQueryDataOk() (*bool, bool) {
	if o == nil || o.RequireMateQueryData == nil {
		return nil, false
	}
	return o.RequireMateQueryData, true
}

// HasRequireMateQueryData returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasRequireMateQueryData() bool {
	if o != nil && o.RequireMateQueryData != nil {
		return true
	}

	return false
}

// SetRequireMateQueryData gets a reference to the given bool and assigns it to the RequireMateQueryData field.
func (o *BTAllowedMateTypeFilter1511) SetRequireMateQueryData(v bool) {
	o.RequireMateQueryData = &v
}

// GetTopLevelMateOnly returns the TopLevelMateOnly field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetTopLevelMateOnly() bool {
	if o == nil || o.TopLevelMateOnly == nil {
		var ret bool
		return ret
	}
	return *o.TopLevelMateOnly
}

// GetTopLevelMateOnlyOk returns a tuple with the TopLevelMateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetTopLevelMateOnlyOk() (*bool, bool) {
	if o == nil || o.TopLevelMateOnly == nil {
		return nil, false
	}
	return o.TopLevelMateOnly, true
}

// HasTopLevelMateOnly returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasTopLevelMateOnly() bool {
	if o != nil && o.TopLevelMateOnly != nil {
		return true
	}

	return false
}

// SetTopLevelMateOnly gets a reference to the given bool and assigns it to the TopLevelMateOnly field.
func (o *BTAllowedMateTypeFilter1511) SetTopLevelMateOnly(v bool) {
	o.TopLevelMateOnly = &v
}

// GetAllowedMateTypes returns the AllowedMateTypes field value if set, zero value otherwise.
func (o *BTAllowedMateTypeFilter1511) GetAllowedMateTypes() []GBTMateType {
	if o == nil || o.AllowedMateTypes == nil {
		var ret []GBTMateType
		return ret
	}
	return o.AllowedMateTypes
}

// GetAllowedMateTypesOk returns a tuple with the AllowedMateTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowedMateTypeFilter1511) GetAllowedMateTypesOk() ([]GBTMateType, bool) {
	if o == nil || o.AllowedMateTypes == nil {
		return nil, false
	}
	return o.AllowedMateTypes, true
}

// HasAllowedMateTypes returns a boolean if a field has been set.
func (o *BTAllowedMateTypeFilter1511) HasAllowedMateTypes() bool {
	if o != nil && o.AllowedMateTypes != nil {
		return true
	}

	return false
}

// SetAllowedMateTypes gets a reference to the given []GBTMateType and assigns it to the AllowedMateTypes field.
func (o *BTAllowedMateTypeFilter1511) SetAllowedMateTypes(v []GBTMateType) {
	o.AllowedMateTypes = v
}

func (o BTAllowedMateTypeFilter1511) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMateFilter162, errBTMateFilter162 := json.Marshal(o.BTMateFilter162)
	if errBTMateFilter162 != nil {
		return []byte{}, errBTMateFilter162
	}
	errBTMateFilter162 = json.Unmarshal([]byte(serializedBTMateFilter162), &toSerialize)
	if errBTMateFilter162 != nil {
		return []byte{}, errBTMateFilter162
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.RequireMateQueryData != nil {
		toSerialize["requireMateQueryData"] = o.RequireMateQueryData
	}
	if o.TopLevelMateOnly != nil {
		toSerialize["topLevelMateOnly"] = o.TopLevelMateOnly
	}
	if o.AllowedMateTypes != nil {
		toSerialize["allowedMateTypes"] = o.AllowedMateTypes
	}
	return json.Marshal(toSerialize)
}

type NullableBTAllowedMateTypeFilter1511 struct {
	value *BTAllowedMateTypeFilter1511
	isSet bool
}

func (v NullableBTAllowedMateTypeFilter1511) Get() *BTAllowedMateTypeFilter1511 {
	return v.value
}

func (v *NullableBTAllowedMateTypeFilter1511) Set(val *BTAllowedMateTypeFilter1511) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAllowedMateTypeFilter1511) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAllowedMateTypeFilter1511) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAllowedMateTypeFilter1511(val *BTAllowedMateTypeFilter1511) *NullableBTAllowedMateTypeFilter1511 {
	return &NullableBTAllowedMateTypeFilter1511{value: val, isSet: true}
}

func (v NullableBTAllowedMateTypeFilter1511) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAllowedMateTypeFilter1511) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
