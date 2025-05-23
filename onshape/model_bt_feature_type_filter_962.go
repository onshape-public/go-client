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

// BTFeatureTypeFilter962 struct for BTFeatureTypeFilter962
type BTFeatureTypeFilter962 struct {
	BTQueryFilter183
	BtType      *string `json:"btType,omitempty"`
	FeatureType *string `json:"featureType,omitempty"`
}

// NewBTFeatureTypeFilter962 instantiates a new BTFeatureTypeFilter962 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureTypeFilter962() *BTFeatureTypeFilter962 {
	this := BTFeatureTypeFilter962{}
	return &this
}

// NewBTFeatureTypeFilter962WithDefaults instantiates a new BTFeatureTypeFilter962 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureTypeFilter962WithDefaults() *BTFeatureTypeFilter962 {
	this := BTFeatureTypeFilter962{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureTypeFilter962) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureTypeFilter962) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureTypeFilter962) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureTypeFilter962) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTFeatureTypeFilter962) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureTypeFilter962) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTFeatureTypeFilter962) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTFeatureTypeFilter962) SetFeatureType(v string) {
	o.FeatureType = &v
}

func (o BTFeatureTypeFilter962) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureTypeFilter962 struct {
	value *BTFeatureTypeFilter962
	isSet bool
}

func (v NullableBTFeatureTypeFilter962) Get() *BTFeatureTypeFilter962 {
	return v.value
}

func (v *NullableBTFeatureTypeFilter962) Set(val *BTFeatureTypeFilter962) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureTypeFilter962) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureTypeFilter962) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureTypeFilter962(val *BTFeatureTypeFilter962) *NullableBTFeatureTypeFilter962 {
	return &NullableBTFeatureTypeFilter962{value: val, isSet: true}
}

func (v NullableBTFeatureTypeFilter962) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureTypeFilter962) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
