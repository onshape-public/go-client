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

// BTBillOfMaterialsUniqueItemId2029 struct for BTBillOfMaterialsUniqueItemId2029
type BTBillOfMaterialsUniqueItemId2029 struct {
	ApiConfiguration *string `json:"apiConfiguration,omitempty"`
	// Type of JSON object.
	BtType                                 *string                `json:"btType,omitempty"`
	IsStandardContent                      *bool                  `json:"isStandardContent,omitempty"`
	ItemDefinitionId                       *string                `json:"itemDefinitionId,omitempty"`
	MetadataObjectType                     *BTMetadataObjectType  `json:"metadataObjectType,omitempty"`
	PartId                                 *string                `json:"partId,omitempty"`
	PartIdentity                           *BTPSOIdentity2741     `json:"partIdentity,omitempty"`
	SourceElement                          *BTElementReference725 `json:"sourceElement,omitempty"`
	VersionMetadataWorkspaceId             *string                `json:"versionMetadataWorkspaceId,omitempty"`
	VersionMetadataWorkspaceMicroversionId *string                `json:"versionMetadataWorkspaceMicroversionId,omitempty"`
}

// NewBTBillOfMaterialsUniqueItemId2029 instantiates a new BTBillOfMaterialsUniqueItemId2029 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsUniqueItemId2029() *BTBillOfMaterialsUniqueItemId2029 {
	this := BTBillOfMaterialsUniqueItemId2029{}
	return &this
}

// NewBTBillOfMaterialsUniqueItemId2029WithDefaults instantiates a new BTBillOfMaterialsUniqueItemId2029 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsUniqueItemId2029WithDefaults() *BTBillOfMaterialsUniqueItemId2029 {
	this := BTBillOfMaterialsUniqueItemId2029{}
	return &this
}

// GetApiConfiguration returns the ApiConfiguration field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetApiConfiguration() string {
	if o == nil || o.ApiConfiguration == nil {
		var ret string
		return ret
	}
	return *o.ApiConfiguration
}

// GetApiConfigurationOk returns a tuple with the ApiConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetApiConfigurationOk() (*string, bool) {
	if o == nil || o.ApiConfiguration == nil {
		return nil, false
	}
	return o.ApiConfiguration, true
}

// HasApiConfiguration returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasApiConfiguration() bool {
	if o != nil && o.ApiConfiguration != nil {
		return true
	}

	return false
}

// SetApiConfiguration gets a reference to the given string and assigns it to the ApiConfiguration field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetApiConfiguration(v string) {
	o.ApiConfiguration = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetBtType(v string) {
	o.BtType = &v
}

// GetIsStandardContent returns the IsStandardContent field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetIsStandardContent() bool {
	if o == nil || o.IsStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.IsStandardContent
}

// GetIsStandardContentOk returns a tuple with the IsStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetIsStandardContentOk() (*bool, bool) {
	if o == nil || o.IsStandardContent == nil {
		return nil, false
	}
	return o.IsStandardContent, true
}

// HasIsStandardContent returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasIsStandardContent() bool {
	if o != nil && o.IsStandardContent != nil {
		return true
	}

	return false
}

// SetIsStandardContent gets a reference to the given bool and assigns it to the IsStandardContent field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetIsStandardContent(v bool) {
	o.IsStandardContent = &v
}

// GetItemDefinitionId returns the ItemDefinitionId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetItemDefinitionId() string {
	if o == nil || o.ItemDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.ItemDefinitionId
}

// GetItemDefinitionIdOk returns a tuple with the ItemDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetItemDefinitionIdOk() (*string, bool) {
	if o == nil || o.ItemDefinitionId == nil {
		return nil, false
	}
	return o.ItemDefinitionId, true
}

// HasItemDefinitionId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasItemDefinitionId() bool {
	if o != nil && o.ItemDefinitionId != nil {
		return true
	}

	return false
}

// SetItemDefinitionId gets a reference to the given string and assigns it to the ItemDefinitionId field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetItemDefinitionId(v string) {
	o.ItemDefinitionId = &v
}

// GetMetadataObjectType returns the MetadataObjectType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetMetadataObjectType() BTMetadataObjectType {
	if o == nil || o.MetadataObjectType == nil {
		var ret BTMetadataObjectType
		return ret
	}
	return *o.MetadataObjectType
}

// GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetMetadataObjectTypeOk() (*BTMetadataObjectType, bool) {
	if o == nil || o.MetadataObjectType == nil {
		return nil, false
	}
	return o.MetadataObjectType, true
}

// HasMetadataObjectType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasMetadataObjectType() bool {
	if o != nil && o.MetadataObjectType != nil {
		return true
	}

	return false
}

// SetMetadataObjectType gets a reference to the given BTMetadataObjectType and assigns it to the MetadataObjectType field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetMetadataObjectType(v BTMetadataObjectType) {
	o.MetadataObjectType = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetPartIdentity() BTPSOIdentity2741 {
	if o == nil || o.PartIdentity == nil {
		var ret BTPSOIdentity2741
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetPartIdentityOk() (*BTPSOIdentity2741, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given BTPSOIdentity2741 and assigns it to the PartIdentity field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetPartIdentity(v BTPSOIdentity2741) {
	o.PartIdentity = &v
}

// GetSourceElement returns the SourceElement field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetSourceElement() BTElementReference725 {
	if o == nil || o.SourceElement == nil {
		var ret BTElementReference725
		return ret
	}
	return *o.SourceElement
}

// GetSourceElementOk returns a tuple with the SourceElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetSourceElementOk() (*BTElementReference725, bool) {
	if o == nil || o.SourceElement == nil {
		return nil, false
	}
	return o.SourceElement, true
}

// HasSourceElement returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasSourceElement() bool {
	if o != nil && o.SourceElement != nil {
		return true
	}

	return false
}

// SetSourceElement gets a reference to the given BTElementReference725 and assigns it to the SourceElement field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetSourceElement(v BTElementReference725) {
	o.SourceElement = &v
}

// GetVersionMetadataWorkspaceId returns the VersionMetadataWorkspaceId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetVersionMetadataWorkspaceId() string {
	if o == nil || o.VersionMetadataWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.VersionMetadataWorkspaceId
}

// GetVersionMetadataWorkspaceIdOk returns a tuple with the VersionMetadataWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetVersionMetadataWorkspaceIdOk() (*string, bool) {
	if o == nil || o.VersionMetadataWorkspaceId == nil {
		return nil, false
	}
	return o.VersionMetadataWorkspaceId, true
}

// HasVersionMetadataWorkspaceId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasVersionMetadataWorkspaceId() bool {
	if o != nil && o.VersionMetadataWorkspaceId != nil {
		return true
	}

	return false
}

// SetVersionMetadataWorkspaceId gets a reference to the given string and assigns it to the VersionMetadataWorkspaceId field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetVersionMetadataWorkspaceId(v string) {
	o.VersionMetadataWorkspaceId = &v
}

// GetVersionMetadataWorkspaceMicroversionId returns the VersionMetadataWorkspaceMicroversionId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsUniqueItemId2029) GetVersionMetadataWorkspaceMicroversionId() string {
	if o == nil || o.VersionMetadataWorkspaceMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.VersionMetadataWorkspaceMicroversionId
}

// GetVersionMetadataWorkspaceMicroversionIdOk returns a tuple with the VersionMetadataWorkspaceMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) GetVersionMetadataWorkspaceMicroversionIdOk() (*string, bool) {
	if o == nil || o.VersionMetadataWorkspaceMicroversionId == nil {
		return nil, false
	}
	return o.VersionMetadataWorkspaceMicroversionId, true
}

// HasVersionMetadataWorkspaceMicroversionId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsUniqueItemId2029) HasVersionMetadataWorkspaceMicroversionId() bool {
	if o != nil && o.VersionMetadataWorkspaceMicroversionId != nil {
		return true
	}

	return false
}

// SetVersionMetadataWorkspaceMicroversionId gets a reference to the given string and assigns it to the VersionMetadataWorkspaceMicroversionId field.
func (o *BTBillOfMaterialsUniqueItemId2029) SetVersionMetadataWorkspaceMicroversionId(v string) {
	o.VersionMetadataWorkspaceMicroversionId = &v
}

func (o BTBillOfMaterialsUniqueItemId2029) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiConfiguration != nil {
		toSerialize["apiConfiguration"] = o.ApiConfiguration
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsStandardContent != nil {
		toSerialize["isStandardContent"] = o.IsStandardContent
	}
	if o.ItemDefinitionId != nil {
		toSerialize["itemDefinitionId"] = o.ItemDefinitionId
	}
	if o.MetadataObjectType != nil {
		toSerialize["metadataObjectType"] = o.MetadataObjectType
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartIdentity != nil {
		toSerialize["partIdentity"] = o.PartIdentity
	}
	if o.SourceElement != nil {
		toSerialize["sourceElement"] = o.SourceElement
	}
	if o.VersionMetadataWorkspaceId != nil {
		toSerialize["versionMetadataWorkspaceId"] = o.VersionMetadataWorkspaceId
	}
	if o.VersionMetadataWorkspaceMicroversionId != nil {
		toSerialize["versionMetadataWorkspaceMicroversionId"] = o.VersionMetadataWorkspaceMicroversionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsUniqueItemId2029 struct {
	value *BTBillOfMaterialsUniqueItemId2029
	isSet bool
}

func (v NullableBTBillOfMaterialsUniqueItemId2029) Get() *BTBillOfMaterialsUniqueItemId2029 {
	return v.value
}

func (v *NullableBTBillOfMaterialsUniqueItemId2029) Set(val *BTBillOfMaterialsUniqueItemId2029) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsUniqueItemId2029) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsUniqueItemId2029) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsUniqueItemId2029(val *BTBillOfMaterialsUniqueItemId2029) *NullableBTBillOfMaterialsUniqueItemId2029 {
	return &NullableBTBillOfMaterialsUniqueItemId2029{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsUniqueItemId2029) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsUniqueItemId2029) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
