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

// BTAnnotationElementDisplayData894 struct for BTAnnotationElementDisplayData894
type BTAnnotationElementDisplayData894 struct {
	AnnotationIdToDisplayObject *map[string]BTAnnotationDisplayData3225 `json:"annotationIdToDisplayObject,omitempty"`
	AnnotationIds               []string                                `json:"annotationIds,omitempty"`
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// NewBTAnnotationElementDisplayData894 instantiates a new BTAnnotationElementDisplayData894 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAnnotationElementDisplayData894() *BTAnnotationElementDisplayData894 {
	this := BTAnnotationElementDisplayData894{}
	return &this
}

// NewBTAnnotationElementDisplayData894WithDefaults instantiates a new BTAnnotationElementDisplayData894 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAnnotationElementDisplayData894WithDefaults() *BTAnnotationElementDisplayData894 {
	this := BTAnnotationElementDisplayData894{}
	return &this
}

// GetAnnotationIdToDisplayObject returns the AnnotationIdToDisplayObject field value if set, zero value otherwise.
func (o *BTAnnotationElementDisplayData894) GetAnnotationIdToDisplayObject() map[string]BTAnnotationDisplayData3225 {
	if o == nil || o.AnnotationIdToDisplayObject == nil {
		var ret map[string]BTAnnotationDisplayData3225
		return ret
	}
	return *o.AnnotationIdToDisplayObject
}

// GetAnnotationIdToDisplayObjectOk returns a tuple with the AnnotationIdToDisplayObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationElementDisplayData894) GetAnnotationIdToDisplayObjectOk() (*map[string]BTAnnotationDisplayData3225, bool) {
	if o == nil || o.AnnotationIdToDisplayObject == nil {
		return nil, false
	}
	return o.AnnotationIdToDisplayObject, true
}

// HasAnnotationIdToDisplayObject returns a boolean if a field has been set.
func (o *BTAnnotationElementDisplayData894) HasAnnotationIdToDisplayObject() bool {
	if o != nil && o.AnnotationIdToDisplayObject != nil {
		return true
	}

	return false
}

// SetAnnotationIdToDisplayObject gets a reference to the given map[string]BTAnnotationDisplayData3225 and assigns it to the AnnotationIdToDisplayObject field.
func (o *BTAnnotationElementDisplayData894) SetAnnotationIdToDisplayObject(v map[string]BTAnnotationDisplayData3225) {
	o.AnnotationIdToDisplayObject = &v
}

// GetAnnotationIds returns the AnnotationIds field value if set, zero value otherwise.
func (o *BTAnnotationElementDisplayData894) GetAnnotationIds() []string {
	if o == nil || o.AnnotationIds == nil {
		var ret []string
		return ret
	}
	return o.AnnotationIds
}

// GetAnnotationIdsOk returns a tuple with the AnnotationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationElementDisplayData894) GetAnnotationIdsOk() ([]string, bool) {
	if o == nil || o.AnnotationIds == nil {
		return nil, false
	}
	return o.AnnotationIds, true
}

// HasAnnotationIds returns a boolean if a field has been set.
func (o *BTAnnotationElementDisplayData894) HasAnnotationIds() bool {
	if o != nil && o.AnnotationIds != nil {
		return true
	}

	return false
}

// SetAnnotationIds gets a reference to the given []string and assigns it to the AnnotationIds field.
func (o *BTAnnotationElementDisplayData894) SetAnnotationIds(v []string) {
	o.AnnotationIds = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAnnotationElementDisplayData894) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAnnotationElementDisplayData894) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAnnotationElementDisplayData894) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAnnotationElementDisplayData894) SetBtType(v string) {
	o.BtType = &v
}

func (o BTAnnotationElementDisplayData894) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnnotationIdToDisplayObject != nil {
		toSerialize["annotationIdToDisplayObject"] = o.AnnotationIdToDisplayObject
	}
	if o.AnnotationIds != nil {
		toSerialize["annotationIds"] = o.AnnotationIds
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTAnnotationElementDisplayData894 struct {
	value *BTAnnotationElementDisplayData894
	isSet bool
}

func (v NullableBTAnnotationElementDisplayData894) Get() *BTAnnotationElementDisplayData894 {
	return v.value
}

func (v *NullableBTAnnotationElementDisplayData894) Set(val *BTAnnotationElementDisplayData894) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAnnotationElementDisplayData894) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAnnotationElementDisplayData894) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAnnotationElementDisplayData894(val *BTAnnotationElementDisplayData894) *NullableBTAnnotationElementDisplayData894 {
	return &NullableBTAnnotationElementDisplayData894{value: val, isSet: true}
}

func (v NullableBTAnnotationElementDisplayData894) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAnnotationElementDisplayData894) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
