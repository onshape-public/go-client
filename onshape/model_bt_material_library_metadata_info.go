/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMaterialLibraryMetadataInfo struct for BTMaterialLibraryMetadataInfo
type BTMaterialLibraryMetadataInfo struct {
	DocumentId   *string `json:"documentId,omitempty"`
	DocumentName *string `json:"documentName,omitempty"`
	ElementId    *string `json:"elementId,omitempty"`
	IsPublic     *bool   `json:"isPublic,omitempty"`
	LibraryName  *string `json:"libraryName,omitempty"`
	OwnerName    *string `json:"ownerName,omitempty"`
	VersionId    *string `json:"versionId,omitempty"`
	WorkspaceId  *string `json:"workspaceId,omitempty"`
}

// NewBTMaterialLibraryMetadataInfo instantiates a new BTMaterialLibraryMetadataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMaterialLibraryMetadataInfo() *BTMaterialLibraryMetadataInfo {
	this := BTMaterialLibraryMetadataInfo{}
	return &this
}

// NewBTMaterialLibraryMetadataInfoWithDefaults instantiates a new BTMaterialLibraryMetadataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMaterialLibraryMetadataInfoWithDefaults() *BTMaterialLibraryMetadataInfo {
	this := BTMaterialLibraryMetadataInfo{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTMaterialLibraryMetadataInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetDocumentName() string {
	if o == nil || o.DocumentName == nil {
		var ret string
		return ret
	}
	return *o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetDocumentNameOk() (*string, bool) {
	if o == nil || o.DocumentName == nil {
		return nil, false
	}
	return o.DocumentName, true
}

// HasDocumentName returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasDocumentName() bool {
	if o != nil && o.DocumentName != nil {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given string and assigns it to the DocumentName field.
func (o *BTMaterialLibraryMetadataInfo) SetDocumentName(v string) {
	o.DocumentName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTMaterialLibraryMetadataInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *BTMaterialLibraryMetadataInfo) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetLibraryName returns the LibraryName field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetLibraryName() string {
	if o == nil || o.LibraryName == nil {
		var ret string
		return ret
	}
	return *o.LibraryName
}

// GetLibraryNameOk returns a tuple with the LibraryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetLibraryNameOk() (*string, bool) {
	if o == nil || o.LibraryName == nil {
		return nil, false
	}
	return o.LibraryName, true
}

// HasLibraryName returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasLibraryName() bool {
	if o != nil && o.LibraryName != nil {
		return true
	}

	return false
}

// SetLibraryName gets a reference to the given string and assigns it to the LibraryName field.
func (o *BTMaterialLibraryMetadataInfo) SetLibraryName(v string) {
	o.LibraryName = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetOwnerName() string {
	if o == nil || o.OwnerName == nil {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetOwnerNameOk() (*string, bool) {
	if o == nil || o.OwnerName == nil {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasOwnerName() bool {
	if o != nil && o.OwnerName != nil {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *BTMaterialLibraryMetadataInfo) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTMaterialLibraryMetadataInfo) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTMaterialLibraryMetadataInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibraryMetadataInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTMaterialLibraryMetadataInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTMaterialLibraryMetadataInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTMaterialLibraryMetadataInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentName != nil {
		toSerialize["documentName"] = o.DocumentName
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.LibraryName != nil {
		toSerialize["libraryName"] = o.LibraryName
	}
	if o.OwnerName != nil {
		toSerialize["ownerName"] = o.OwnerName
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMaterialLibraryMetadataInfo struct {
	value *BTMaterialLibraryMetadataInfo
	isSet bool
}

func (v NullableBTMaterialLibraryMetadataInfo) Get() *BTMaterialLibraryMetadataInfo {
	return v.value
}

func (v *NullableBTMaterialLibraryMetadataInfo) Set(val *BTMaterialLibraryMetadataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMaterialLibraryMetadataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMaterialLibraryMetadataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMaterialLibraryMetadataInfo(val *BTMaterialLibraryMetadataInfo) *NullableBTMaterialLibraryMetadataInfo {
	return &NullableBTMaterialLibraryMetadataInfo{value: val, isSet: true}
}

func (v NullableBTMaterialLibraryMetadataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMaterialLibraryMetadataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}