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

// BTAssemblyItemMetadataRequestInfo struct for BTAssemblyItemMetadataRequestInfo
type BTAssemblyItemMetadataRequestInfo struct {
	ApiConfig        *string `json:"apiConfig,omitempty"`
	DocumentId       *string `json:"documentId,omitempty"`
	ElementId        *string `json:"elementId,omitempty"`
	ItemId           *string `json:"itemId,omitempty"`
	LinkedDocumentId *string `json:"linkedDocumentId,omitempty"`
	PartId           *string `json:"partId,omitempty"`
	WvmId            *string `json:"wvmId,omitempty"`
	WvmType          *string `json:"wvmType,omitempty"`
}

// NewBTAssemblyItemMetadataRequestInfo instantiates a new BTAssemblyItemMetadataRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyItemMetadataRequestInfo() *BTAssemblyItemMetadataRequestInfo {
	this := BTAssemblyItemMetadataRequestInfo{}
	return &this
}

// NewBTAssemblyItemMetadataRequestInfoWithDefaults instantiates a new BTAssemblyItemMetadataRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyItemMetadataRequestInfoWithDefaults() *BTAssemblyItemMetadataRequestInfo {
	this := BTAssemblyItemMetadataRequestInfo{}
	return &this
}

// GetApiConfig returns the ApiConfig field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetApiConfig() string {
	if o == nil || o.ApiConfig == nil {
		var ret string
		return ret
	}
	return *o.ApiConfig
}

// GetApiConfigOk returns a tuple with the ApiConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetApiConfigOk() (*string, bool) {
	if o == nil || o.ApiConfig == nil {
		return nil, false
	}
	return o.ApiConfig, true
}

// HasApiConfig returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasApiConfig() bool {
	if o != nil && o.ApiConfig != nil {
		return true
	}

	return false
}

// SetApiConfig gets a reference to the given string and assigns it to the ApiConfig field.
func (o *BTAssemblyItemMetadataRequestInfo) SetApiConfig(v string) {
	o.ApiConfig = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAssemblyItemMetadataRequestInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyItemMetadataRequestInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *BTAssemblyItemMetadataRequestInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetLinkedDocumentId returns the LinkedDocumentId field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetLinkedDocumentId() string {
	if o == nil || o.LinkedDocumentId == nil {
		var ret string
		return ret
	}
	return *o.LinkedDocumentId
}

// GetLinkedDocumentIdOk returns a tuple with the LinkedDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetLinkedDocumentIdOk() (*string, bool) {
	if o == nil || o.LinkedDocumentId == nil {
		return nil, false
	}
	return o.LinkedDocumentId, true
}

// HasLinkedDocumentId returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasLinkedDocumentId() bool {
	if o != nil && o.LinkedDocumentId != nil {
		return true
	}

	return false
}

// SetLinkedDocumentId gets a reference to the given string and assigns it to the LinkedDocumentId field.
func (o *BTAssemblyItemMetadataRequestInfo) SetLinkedDocumentId(v string) {
	o.LinkedDocumentId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTAssemblyItemMetadataRequestInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetWvmId returns the WvmId field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetWvmId() string {
	if o == nil || o.WvmId == nil {
		var ret string
		return ret
	}
	return *o.WvmId
}

// GetWvmIdOk returns a tuple with the WvmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetWvmIdOk() (*string, bool) {
	if o == nil || o.WvmId == nil {
		return nil, false
	}
	return o.WvmId, true
}

// HasWvmId returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasWvmId() bool {
	if o != nil && o.WvmId != nil {
		return true
	}

	return false
}

// SetWvmId gets a reference to the given string and assigns it to the WvmId field.
func (o *BTAssemblyItemMetadataRequestInfo) SetWvmId(v string) {
	o.WvmId = &v
}

// GetWvmType returns the WvmType field value if set, zero value otherwise.
func (o *BTAssemblyItemMetadataRequestInfo) GetWvmType() string {
	if o == nil || o.WvmType == nil {
		var ret string
		return ret
	}
	return *o.WvmType
}

// GetWvmTypeOk returns a tuple with the WvmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyItemMetadataRequestInfo) GetWvmTypeOk() (*string, bool) {
	if o == nil || o.WvmType == nil {
		return nil, false
	}
	return o.WvmType, true
}

// HasWvmType returns a boolean if a field has been set.
func (o *BTAssemblyItemMetadataRequestInfo) HasWvmType() bool {
	if o != nil && o.WvmType != nil {
		return true
	}

	return false
}

// SetWvmType gets a reference to the given string and assigns it to the WvmType field.
func (o *BTAssemblyItemMetadataRequestInfo) SetWvmType(v string) {
	o.WvmType = &v
}

func (o BTAssemblyItemMetadataRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiConfig != nil {
		toSerialize["apiConfig"] = o.ApiConfig
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ItemId != nil {
		toSerialize["itemId"] = o.ItemId
	}
	if o.LinkedDocumentId != nil {
		toSerialize["linkedDocumentId"] = o.LinkedDocumentId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.WvmId != nil {
		toSerialize["wvmId"] = o.WvmId
	}
	if o.WvmType != nil {
		toSerialize["wvmType"] = o.WvmType
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyItemMetadataRequestInfo struct {
	value *BTAssemblyItemMetadataRequestInfo
	isSet bool
}

func (v NullableBTAssemblyItemMetadataRequestInfo) Get() *BTAssemblyItemMetadataRequestInfo {
	return v.value
}

func (v *NullableBTAssemblyItemMetadataRequestInfo) Set(val *BTAssemblyItemMetadataRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyItemMetadataRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyItemMetadataRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyItemMetadataRequestInfo(val *BTAssemblyItemMetadataRequestInfo) *NullableBTAssemblyItemMetadataRequestInfo {
	return &NullableBTAssemblyItemMetadataRequestInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyItemMetadataRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyItemMetadataRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
