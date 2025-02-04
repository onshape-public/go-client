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

// BTDiffJsonResponse2725 struct for BTDiffJsonResponse2725
type BTDiffJsonResponse2725 struct {
	// Type of JSON object.
	BtType         *string                 `json:"btType,omitempty"`
	Change         *BTJEdit3734            `json:"change,omitempty"`
	Patch          *map[string]interface{} `json:"patch,omitempty"`
	SourceChangeId *string                 `json:"sourceChangeId,omitempty"`
	TargetChangeId *string                 `json:"targetChangeId,omitempty"`
}

// NewBTDiffJsonResponse2725 instantiates a new BTDiffJsonResponse2725 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDiffJsonResponse2725() *BTDiffJsonResponse2725 {
	this := BTDiffJsonResponse2725{}
	return &this
}

// NewBTDiffJsonResponse2725WithDefaults instantiates a new BTDiffJsonResponse2725 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDiffJsonResponse2725WithDefaults() *BTDiffJsonResponse2725 {
	this := BTDiffJsonResponse2725{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDiffJsonResponse2725) SetBtType(v string) {
	o.BtType = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetChange() BTJEdit3734 {
	if o == nil || o.Change == nil {
		var ret BTJEdit3734
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetChangeOk() (*BTJEdit3734, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasChange() bool {
	if o != nil && o.Change != nil {
		return true
	}

	return false
}

// SetChange gets a reference to the given BTJEdit3734 and assigns it to the Change field.
func (o *BTDiffJsonResponse2725) SetChange(v BTJEdit3734) {
	o.Change = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetPatch() map[string]interface{} {
	if o == nil || o.Patch == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetPatchOk() (*map[string]interface{}, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given map[string]interface{} and assigns it to the Patch field.
func (o *BTDiffJsonResponse2725) SetPatch(v map[string]interface{}) {
	o.Patch = &v
}

// GetSourceChangeId returns the SourceChangeId field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetSourceChangeId() string {
	if o == nil || o.SourceChangeId == nil {
		var ret string
		return ret
	}
	return *o.SourceChangeId
}

// GetSourceChangeIdOk returns a tuple with the SourceChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetSourceChangeIdOk() (*string, bool) {
	if o == nil || o.SourceChangeId == nil {
		return nil, false
	}
	return o.SourceChangeId, true
}

// HasSourceChangeId returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasSourceChangeId() bool {
	if o != nil && o.SourceChangeId != nil {
		return true
	}

	return false
}

// SetSourceChangeId gets a reference to the given string and assigns it to the SourceChangeId field.
func (o *BTDiffJsonResponse2725) SetSourceChangeId(v string) {
	o.SourceChangeId = &v
}

// GetTargetChangeId returns the TargetChangeId field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetTargetChangeId() string {
	if o == nil || o.TargetChangeId == nil {
		var ret string
		return ret
	}
	return *o.TargetChangeId
}

// GetTargetChangeIdOk returns a tuple with the TargetChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetTargetChangeIdOk() (*string, bool) {
	if o == nil || o.TargetChangeId == nil {
		return nil, false
	}
	return o.TargetChangeId, true
}

// HasTargetChangeId returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasTargetChangeId() bool {
	if o != nil && o.TargetChangeId != nil {
		return true
	}

	return false
}

// SetTargetChangeId gets a reference to the given string and assigns it to the TargetChangeId field.
func (o *BTDiffJsonResponse2725) SetTargetChangeId(v string) {
	o.TargetChangeId = &v
}

func (o BTDiffJsonResponse2725) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	if o.SourceChangeId != nil {
		toSerialize["sourceChangeId"] = o.SourceChangeId
	}
	if o.TargetChangeId != nil {
		toSerialize["targetChangeId"] = o.TargetChangeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDiffJsonResponse2725 struct {
	value *BTDiffJsonResponse2725
	isSet bool
}

func (v NullableBTDiffJsonResponse2725) Get() *BTDiffJsonResponse2725 {
	return v.value
}

func (v *NullableBTDiffJsonResponse2725) Set(val *BTDiffJsonResponse2725) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDiffJsonResponse2725) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDiffJsonResponse2725) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDiffJsonResponse2725(val *BTDiffJsonResponse2725) *NullableBTDiffJsonResponse2725 {
	return &NullableBTDiffJsonResponse2725{value: val, isSet: true}
}

func (v NullableBTDiffJsonResponse2725) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDiffJsonResponse2725) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
