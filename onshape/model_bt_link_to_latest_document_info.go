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

// BTLinkToLatestDocumentInfo struct for BTLinkToLatestDocumentInfo
type BTLinkToLatestDocumentInfo struct {
	ChangedElements []string `json:"changedElements,omitempty"`
}

// NewBTLinkToLatestDocumentInfo instantiates a new BTLinkToLatestDocumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLinkToLatestDocumentInfo() *BTLinkToLatestDocumentInfo {
	this := BTLinkToLatestDocumentInfo{}
	return &this
}

// NewBTLinkToLatestDocumentInfoWithDefaults instantiates a new BTLinkToLatestDocumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLinkToLatestDocumentInfoWithDefaults() *BTLinkToLatestDocumentInfo {
	this := BTLinkToLatestDocumentInfo{}
	return &this
}

// GetChangedElements returns the ChangedElements field value if set, zero value otherwise.
func (o *BTLinkToLatestDocumentInfo) GetChangedElements() []string {
	if o == nil || o.ChangedElements == nil {
		var ret []string
		return ret
	}
	return o.ChangedElements
}

// GetChangedElementsOk returns a tuple with the ChangedElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLinkToLatestDocumentInfo) GetChangedElementsOk() ([]string, bool) {
	if o == nil || o.ChangedElements == nil {
		return nil, false
	}
	return o.ChangedElements, true
}

// HasChangedElements returns a boolean if a field has been set.
func (o *BTLinkToLatestDocumentInfo) HasChangedElements() bool {
	if o != nil && o.ChangedElements != nil {
		return true
	}

	return false
}

// SetChangedElements gets a reference to the given []string and assigns it to the ChangedElements field.
func (o *BTLinkToLatestDocumentInfo) SetChangedElements(v []string) {
	o.ChangedElements = v
}

func (o BTLinkToLatestDocumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangedElements != nil {
		toSerialize["changedElements"] = o.ChangedElements
	}
	return json.Marshal(toSerialize)
}

type NullableBTLinkToLatestDocumentInfo struct {
	value *BTLinkToLatestDocumentInfo
	isSet bool
}

func (v NullableBTLinkToLatestDocumentInfo) Get() *BTLinkToLatestDocumentInfo {
	return v.value
}

func (v *NullableBTLinkToLatestDocumentInfo) Set(val *BTLinkToLatestDocumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLinkToLatestDocumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLinkToLatestDocumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLinkToLatestDocumentInfo(val *BTLinkToLatestDocumentInfo) *NullableBTLinkToLatestDocumentInfo {
	return &NullableBTLinkToLatestDocumentInfo{value: val, isSet: true}
}

func (v NullableBTLinkToLatestDocumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLinkToLatestDocumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
