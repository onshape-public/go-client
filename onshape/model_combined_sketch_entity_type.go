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

// CombinedSketchEntityType struct for CombinedSketchEntityType
type CombinedSketchEntityType struct {
	ConstraintType *GBTConstraintType   `json:"constraintType,omitempty"`
	EntityType     *GBTSketchEntityType `json:"entityType,omitempty"`
}

// NewCombinedSketchEntityType instantiates a new CombinedSketchEntityType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombinedSketchEntityType() *CombinedSketchEntityType {
	this := CombinedSketchEntityType{}
	return &this
}

// NewCombinedSketchEntityTypeWithDefaults instantiates a new CombinedSketchEntityType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombinedSketchEntityTypeWithDefaults() *CombinedSketchEntityType {
	this := CombinedSketchEntityType{}
	return &this
}

// GetConstraintType returns the ConstraintType field value if set, zero value otherwise.
func (o *CombinedSketchEntityType) GetConstraintType() GBTConstraintType {
	if o == nil || o.ConstraintType == nil {
		var ret GBTConstraintType
		return ret
	}
	return *o.ConstraintType
}

// GetConstraintTypeOk returns a tuple with the ConstraintType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedSketchEntityType) GetConstraintTypeOk() (*GBTConstraintType, bool) {
	if o == nil || o.ConstraintType == nil {
		return nil, false
	}
	return o.ConstraintType, true
}

// HasConstraintType returns a boolean if a field has been set.
func (o *CombinedSketchEntityType) HasConstraintType() bool {
	if o != nil && o.ConstraintType != nil {
		return true
	}

	return false
}

// SetConstraintType gets a reference to the given GBTConstraintType and assigns it to the ConstraintType field.
func (o *CombinedSketchEntityType) SetConstraintType(v GBTConstraintType) {
	o.ConstraintType = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CombinedSketchEntityType) GetEntityType() GBTSketchEntityType {
	if o == nil || o.EntityType == nil {
		var ret GBTSketchEntityType
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedSketchEntityType) GetEntityTypeOk() (*GBTSketchEntityType, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CombinedSketchEntityType) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given GBTSketchEntityType and assigns it to the EntityType field.
func (o *CombinedSketchEntityType) SetEntityType(v GBTSketchEntityType) {
	o.EntityType = &v
}

func (o CombinedSketchEntityType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConstraintType != nil {
		toSerialize["constraintType"] = o.ConstraintType
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	return json.Marshal(toSerialize)
}

type NullableCombinedSketchEntityType struct {
	value *CombinedSketchEntityType
	isSet bool
}

func (v NullableCombinedSketchEntityType) Get() *CombinedSketchEntityType {
	return v.value
}

func (v *NullableCombinedSketchEntityType) Set(val *CombinedSketchEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableCombinedSketchEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableCombinedSketchEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombinedSketchEntityType(val *CombinedSketchEntityType) *NullableCombinedSketchEntityType {
	return &NullableCombinedSketchEntityType{value: val, isSet: true}
}

func (v NullableCombinedSketchEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombinedSketchEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
