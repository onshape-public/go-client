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

// BTDocumentContentsInfo struct for BTDocumentContentsInfo
type BTDocumentContentsInfo struct {
	// The elements (tabs) in the document. This does not include folders.
	Elements []BTDocumentElementInfo `json:"elements,omitempty"`
	Folders  *BTElementGroup1458     `json:"folders,omitempty"`
}

// NewBTDocumentContentsInfo instantiates a new BTDocumentContentsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentContentsInfo() *BTDocumentContentsInfo {
	this := BTDocumentContentsInfo{}
	return &this
}

// NewBTDocumentContentsInfoWithDefaults instantiates a new BTDocumentContentsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentContentsInfoWithDefaults() *BTDocumentContentsInfo {
	this := BTDocumentContentsInfo{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTDocumentContentsInfo) GetElements() []BTDocumentElementInfo {
	if o == nil || o.Elements == nil {
		var ret []BTDocumentElementInfo
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentContentsInfo) GetElementsOk() ([]BTDocumentElementInfo, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTDocumentContentsInfo) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []BTDocumentElementInfo and assigns it to the Elements field.
func (o *BTDocumentContentsInfo) SetElements(v []BTDocumentElementInfo) {
	o.Elements = v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *BTDocumentContentsInfo) GetFolders() BTElementGroup1458 {
	if o == nil || o.Folders == nil {
		var ret BTElementGroup1458
		return ret
	}
	return *o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentContentsInfo) GetFoldersOk() (*BTElementGroup1458, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *BTDocumentContentsInfo) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given BTElementGroup1458 and assigns it to the Folders field.
func (o *BTDocumentContentsInfo) SetFolders(v BTElementGroup1458) {
	o.Folders = &v
}

func (o BTDocumentContentsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentContentsInfo struct {
	value *BTDocumentContentsInfo
	isSet bool
}

func (v NullableBTDocumentContentsInfo) Get() *BTDocumentContentsInfo {
	return v.value
}

func (v *NullableBTDocumentContentsInfo) Set(val *BTDocumentContentsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentContentsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentContentsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentContentsInfo(val *BTDocumentContentsInfo) *NullableBTDocumentContentsInfo {
	return &NullableBTDocumentContentsInfo{value: val, isSet: true}
}

func (v NullableBTDocumentContentsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentContentsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
