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

// BTWorkflowableTestObjectInfo struct for BTWorkflowableTestObjectInfo
type BTWorkflowableTestObjectInfo struct {
	CompanyId             *string                  `json:"companyId,omitempty"`
	CreatedAt             *JSONTime                `json:"createdAt,omitempty"`
	CreatedBy             *BTUserBasicSummaryInfo  `json:"createdBy,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	DescriptionAsProperty *string                  `json:"descriptionAsProperty,omitempty"`
	DocumentId            *string                  `json:"documentId,omitempty"`
	Href                  *string                  `json:"href,omitempty"`
	Id                    *string                  `json:"id,omitempty"`
	Info                  *map[string]string       `json:"info,omitempty"`
	ModifiedAt            *JSONTime                `json:"modifiedAt,omitempty"`
	ModifiedBy            *BTUserBasicSummaryInfo  `json:"modifiedBy,omitempty"`
	Name                  *string                  `json:"name,omitempty"`
	NameAsProperty        *string                  `json:"nameAsProperty,omitempty"`
	Properties            []BTWorkflowPropertyInfo `json:"properties,omitempty"`
	ViewRef               *string                  `json:"viewRef,omitempty"`
	Workflow              *BTWorkflowSnapshotInfo  `json:"workflow,omitempty"`
	WorkflowError         *string                  `json:"workflowError,omitempty"`
	WorkflowId            *BTPublishedWorkflowId   `json:"workflowId,omitempty"`
}

// NewBTWorkflowableTestObjectInfo instantiates a new BTWorkflowableTestObjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowableTestObjectInfo() *BTWorkflowableTestObjectInfo {
	this := BTWorkflowableTestObjectInfo{}
	return &this
}

// NewBTWorkflowableTestObjectInfoWithDefaults instantiates a new BTWorkflowableTestObjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowableTestObjectInfoWithDefaults() *BTWorkflowableTestObjectInfo {
	this := BTWorkflowableTestObjectInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTWorkflowableTestObjectInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTWorkflowableTestObjectInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTWorkflowableTestObjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWorkflowableTestObjectInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionAsProperty returns the DescriptionAsProperty field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetDescriptionAsProperty() string {
	if o == nil || o.DescriptionAsProperty == nil {
		var ret string
		return ret
	}
	return *o.DescriptionAsProperty
}

// GetDescriptionAsPropertyOk returns a tuple with the DescriptionAsProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetDescriptionAsPropertyOk() (*string, bool) {
	if o == nil || o.DescriptionAsProperty == nil {
		return nil, false
	}
	return o.DescriptionAsProperty, true
}

// HasDescriptionAsProperty returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasDescriptionAsProperty() bool {
	if o != nil && o.DescriptionAsProperty != nil {
		return true
	}

	return false
}

// SetDescriptionAsProperty gets a reference to the given string and assigns it to the DescriptionAsProperty field.
func (o *BTWorkflowableTestObjectInfo) SetDescriptionAsProperty(v string) {
	o.DescriptionAsProperty = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTWorkflowableTestObjectInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTWorkflowableTestObjectInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWorkflowableTestObjectInfo) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetInfo() map[string]string {
	if o == nil || o.Info == nil {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BTWorkflowableTestObjectInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTWorkflowableTestObjectInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTWorkflowableTestObjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkflowableTestObjectInfo) SetName(v string) {
	o.Name = &v
}

// GetNameAsProperty returns the NameAsProperty field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetNameAsProperty() string {
	if o == nil || o.NameAsProperty == nil {
		var ret string
		return ret
	}
	return *o.NameAsProperty
}

// GetNameAsPropertyOk returns a tuple with the NameAsProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetNameAsPropertyOk() (*string, bool) {
	if o == nil || o.NameAsProperty == nil {
		return nil, false
	}
	return o.NameAsProperty, true
}

// HasNameAsProperty returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasNameAsProperty() bool {
	if o != nil && o.NameAsProperty != nil {
		return true
	}

	return false
}

// SetNameAsProperty gets a reference to the given string and assigns it to the NameAsProperty field.
func (o *BTWorkflowableTestObjectInfo) SetNameAsProperty(v string) {
	o.NameAsProperty = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetProperties() []BTWorkflowPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTWorkflowPropertyInfo
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetPropertiesOk() ([]BTWorkflowPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTWorkflowPropertyInfo and assigns it to the Properties field.
func (o *BTWorkflowableTestObjectInfo) SetProperties(v []BTWorkflowPropertyInfo) {
	o.Properties = v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTWorkflowableTestObjectInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetWorkflow() BTWorkflowSnapshotInfo {
	if o == nil || o.Workflow == nil {
		var ret BTWorkflowSnapshotInfo
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowOk() (*BTWorkflowSnapshotInfo, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given BTWorkflowSnapshotInfo and assigns it to the Workflow field.
func (o *BTWorkflowableTestObjectInfo) SetWorkflow(v BTWorkflowSnapshotInfo) {
	o.Workflow = &v
}

// GetWorkflowError returns the WorkflowError field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowError() string {
	if o == nil || o.WorkflowError == nil {
		var ret string
		return ret
	}
	return *o.WorkflowError
}

// GetWorkflowErrorOk returns a tuple with the WorkflowError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowErrorOk() (*string, bool) {
	if o == nil || o.WorkflowError == nil {
		return nil, false
	}
	return o.WorkflowError, true
}

// HasWorkflowError returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasWorkflowError() bool {
	if o != nil && o.WorkflowError != nil {
		return true
	}

	return false
}

// SetWorkflowError gets a reference to the given string and assigns it to the WorkflowError field.
func (o *BTWorkflowableTestObjectInfo) SetWorkflowError(v string) {
	o.WorkflowError = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowId() BTPublishedWorkflowId {
	if o == nil || o.WorkflowId == nil {
		var ret BTPublishedWorkflowId
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowIdOk() (*BTPublishedWorkflowId, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given BTPublishedWorkflowId and assigns it to the WorkflowId field.
func (o *BTWorkflowableTestObjectInfo) SetWorkflowId(v BTPublishedWorkflowId) {
	o.WorkflowId = &v
}

func (o BTWorkflowableTestObjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DescriptionAsProperty != nil {
		toSerialize["descriptionAsProperty"] = o.DescriptionAsProperty
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameAsProperty != nil {
		toSerialize["nameAsProperty"] = o.NameAsProperty
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}
	if o.WorkflowError != nil {
		toSerialize["workflowError"] = o.WorkflowError
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowableTestObjectInfo struct {
	value *BTWorkflowableTestObjectInfo
	isSet bool
}

func (v NullableBTWorkflowableTestObjectInfo) Get() *BTWorkflowableTestObjectInfo {
	return v.value
}

func (v *NullableBTWorkflowableTestObjectInfo) Set(val *BTWorkflowableTestObjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowableTestObjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowableTestObjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowableTestObjectInfo(val *BTWorkflowableTestObjectInfo) *NullableBTWorkflowableTestObjectInfo {
	return &NullableBTWorkflowableTestObjectInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowableTestObjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowableTestObjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}