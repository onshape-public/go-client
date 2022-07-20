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

// BTWorkflowableObjectObserver struct for BTWorkflowableObjectObserver
type BTWorkflowableObjectObserver struct {
	AdminOverride    *bool     `json:"adminOverride,omitempty"`
	ApprovalDate     *JSONTime `json:"approvalDate,omitempty"`
	ApprovalState    *string   `json:"approvalState,omitempty"`
	ApproverId       *string   `json:"approverId,omitempty"`
	ApproverName     *string   `json:"approverName,omitempty"`
	AssociatedStates *string   `json:"associatedStates,omitempty"`
	CompanyId        *string   `json:"companyId,omitempty"`
	CreatedAt        *JSONTime `json:"createdAt,omitempty"`
	CreatedBy        *string   `json:"createdBy,omitempty"`
	Description      *string   `json:"description,omitempty"`
	EntryId          *string   `json:"entryId,omitempty"`
	EntryType        *string   `json:"entryType,omitempty"`
	Id               *string   `json:"id,omitempty"`
	ModifiedAt       *JSONTime `json:"modifiedAt,omitempty"`
	ModifiedBy       *string   `json:"modifiedBy,omitempty"`
	Name             *string   `json:"name,omitempty"`
	New              *bool     `json:"new,omitempty"`
	ObjectId         *string   `json:"objectId,omitempty"`
	ObservationType  *int32    `json:"observationType,omitempty"`
	PropertyId       *string   `json:"propertyId,omitempty"`
	RejectionDate    *JSONTime `json:"rejectionDate,omitempty"`
	Removable        *bool     `json:"removable,omitempty"`
}

// NewBTWorkflowableObjectObserver instantiates a new BTWorkflowableObjectObserver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowableObjectObserver() *BTWorkflowableObjectObserver {
	this := BTWorkflowableObjectObserver{}
	return &this
}

// NewBTWorkflowableObjectObserverWithDefaults instantiates a new BTWorkflowableObjectObserver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowableObjectObserverWithDefaults() *BTWorkflowableObjectObserver {
	this := BTWorkflowableObjectObserver{}
	return &this
}

// GetAdminOverride returns the AdminOverride field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetAdminOverride() bool {
	if o == nil || o.AdminOverride == nil {
		var ret bool
		return ret
	}
	return *o.AdminOverride
}

// GetAdminOverrideOk returns a tuple with the AdminOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetAdminOverrideOk() (*bool, bool) {
	if o == nil || o.AdminOverride == nil {
		return nil, false
	}
	return o.AdminOverride, true
}

// HasAdminOverride returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasAdminOverride() bool {
	if o != nil && o.AdminOverride != nil {
		return true
	}

	return false
}

// SetAdminOverride gets a reference to the given bool and assigns it to the AdminOverride field.
func (o *BTWorkflowableObjectObserver) SetAdminOverride(v bool) {
	o.AdminOverride = &v
}

// GetApprovalDate returns the ApprovalDate field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetApprovalDate() JSONTime {
	if o == nil || o.ApprovalDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.ApprovalDate
}

// GetApprovalDateOk returns a tuple with the ApprovalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetApprovalDateOk() (*JSONTime, bool) {
	if o == nil || o.ApprovalDate == nil {
		return nil, false
	}
	return o.ApprovalDate, true
}

// HasApprovalDate returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasApprovalDate() bool {
	if o != nil && o.ApprovalDate != nil {
		return true
	}

	return false
}

// SetApprovalDate gets a reference to the given JSONTime and assigns it to the ApprovalDate field.
func (o *BTWorkflowableObjectObserver) SetApprovalDate(v JSONTime) {
	o.ApprovalDate = &v
}

// GetApprovalState returns the ApprovalState field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetApprovalState() string {
	if o == nil || o.ApprovalState == nil {
		var ret string
		return ret
	}
	return *o.ApprovalState
}

// GetApprovalStateOk returns a tuple with the ApprovalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetApprovalStateOk() (*string, bool) {
	if o == nil || o.ApprovalState == nil {
		return nil, false
	}
	return o.ApprovalState, true
}

// HasApprovalState returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasApprovalState() bool {
	if o != nil && o.ApprovalState != nil {
		return true
	}

	return false
}

// SetApprovalState gets a reference to the given string and assigns it to the ApprovalState field.
func (o *BTWorkflowableObjectObserver) SetApprovalState(v string) {
	o.ApprovalState = &v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetApproverId() string {
	if o == nil || o.ApproverId == nil {
		var ret string
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetApproverIdOk() (*string, bool) {
	if o == nil || o.ApproverId == nil {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasApproverId() bool {
	if o != nil && o.ApproverId != nil {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given string and assigns it to the ApproverId field.
func (o *BTWorkflowableObjectObserver) SetApproverId(v string) {
	o.ApproverId = &v
}

// GetApproverName returns the ApproverName field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetApproverName() string {
	if o == nil || o.ApproverName == nil {
		var ret string
		return ret
	}
	return *o.ApproverName
}

// GetApproverNameOk returns a tuple with the ApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetApproverNameOk() (*string, bool) {
	if o == nil || o.ApproverName == nil {
		return nil, false
	}
	return o.ApproverName, true
}

// HasApproverName returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasApproverName() bool {
	if o != nil && o.ApproverName != nil {
		return true
	}

	return false
}

// SetApproverName gets a reference to the given string and assigns it to the ApproverName field.
func (o *BTWorkflowableObjectObserver) SetApproverName(v string) {
	o.ApproverName = &v
}

// GetAssociatedStates returns the AssociatedStates field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetAssociatedStates() string {
	if o == nil || o.AssociatedStates == nil {
		var ret string
		return ret
	}
	return *o.AssociatedStates
}

// GetAssociatedStatesOk returns a tuple with the AssociatedStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetAssociatedStatesOk() (*string, bool) {
	if o == nil || o.AssociatedStates == nil {
		return nil, false
	}
	return o.AssociatedStates, true
}

// HasAssociatedStates returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasAssociatedStates() bool {
	if o != nil && o.AssociatedStates != nil {
		return true
	}

	return false
}

// SetAssociatedStates gets a reference to the given string and assigns it to the AssociatedStates field.
func (o *BTWorkflowableObjectObserver) SetAssociatedStates(v string) {
	o.AssociatedStates = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTWorkflowableObjectObserver) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTWorkflowableObjectObserver) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BTWorkflowableObjectObserver) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWorkflowableObjectObserver) SetDescription(v string) {
	o.Description = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *BTWorkflowableObjectObserver) SetEntryId(v string) {
	o.EntryId = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetEntryType() string {
	if o == nil || o.EntryType == nil {
		var ret string
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetEntryTypeOk() (*string, bool) {
	if o == nil || o.EntryType == nil {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasEntryType() bool {
	if o != nil && o.EntryType != nil {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given string and assigns it to the EntryType field.
func (o *BTWorkflowableObjectObserver) SetEntryType(v string) {
	o.EntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWorkflowableObjectObserver) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTWorkflowableObjectObserver) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *BTWorkflowableObjectObserver) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkflowableObjectObserver) SetName(v string) {
	o.Name = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetNew() bool {
	if o == nil || o.New == nil {
		var ret bool
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetNewOk() (*bool, bool) {
	if o == nil || o.New == nil {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasNew() bool {
	if o != nil && o.New != nil {
		return true
	}

	return false
}

// SetNew gets a reference to the given bool and assigns it to the New field.
func (o *BTWorkflowableObjectObserver) SetNew(v bool) {
	o.New = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTWorkflowableObjectObserver) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObservationType returns the ObservationType field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetObservationType() int32 {
	if o == nil || o.ObservationType == nil {
		var ret int32
		return ret
	}
	return *o.ObservationType
}

// GetObservationTypeOk returns a tuple with the ObservationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetObservationTypeOk() (*int32, bool) {
	if o == nil || o.ObservationType == nil {
		return nil, false
	}
	return o.ObservationType, true
}

// HasObservationType returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasObservationType() bool {
	if o != nil && o.ObservationType != nil {
		return true
	}

	return false
}

// SetObservationType gets a reference to the given int32 and assigns it to the ObservationType field.
func (o *BTWorkflowableObjectObserver) SetObservationType(v int32) {
	o.ObservationType = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTWorkflowableObjectObserver) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetRejectionDate returns the RejectionDate field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetRejectionDate() JSONTime {
	if o == nil || o.RejectionDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.RejectionDate
}

// GetRejectionDateOk returns a tuple with the RejectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetRejectionDateOk() (*JSONTime, bool) {
	if o == nil || o.RejectionDate == nil {
		return nil, false
	}
	return o.RejectionDate, true
}

// HasRejectionDate returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasRejectionDate() bool {
	if o != nil && o.RejectionDate != nil {
		return true
	}

	return false
}

// SetRejectionDate gets a reference to the given JSONTime and assigns it to the RejectionDate field.
func (o *BTWorkflowableObjectObserver) SetRejectionDate(v JSONTime) {
	o.RejectionDate = &v
}

// GetRemovable returns the Removable field value if set, zero value otherwise.
func (o *BTWorkflowableObjectObserver) GetRemovable() bool {
	if o == nil || o.Removable == nil {
		var ret bool
		return ret
	}
	return *o.Removable
}

// GetRemovableOk returns a tuple with the Removable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableObjectObserver) GetRemovableOk() (*bool, bool) {
	if o == nil || o.Removable == nil {
		return nil, false
	}
	return o.Removable, true
}

// HasRemovable returns a boolean if a field has been set.
func (o *BTWorkflowableObjectObserver) HasRemovable() bool {
	if o != nil && o.Removable != nil {
		return true
	}

	return false
}

// SetRemovable gets a reference to the given bool and assigns it to the Removable field.
func (o *BTWorkflowableObjectObserver) SetRemovable(v bool) {
	o.Removable = &v
}

func (o BTWorkflowableObjectObserver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminOverride != nil {
		toSerialize["adminOverride"] = o.AdminOverride
	}
	if o.ApprovalDate != nil {
		toSerialize["approvalDate"] = o.ApprovalDate
	}
	if o.ApprovalState != nil {
		toSerialize["approvalState"] = o.ApprovalState
	}
	if o.ApproverId != nil {
		toSerialize["approverId"] = o.ApproverId
	}
	if o.ApproverName != nil {
		toSerialize["approverName"] = o.ApproverName
	}
	if o.AssociatedStates != nil {
		toSerialize["associatedStates"] = o.AssociatedStates
	}
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
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.EntryType != nil {
		toSerialize["entryType"] = o.EntryType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.New != nil {
		toSerialize["new"] = o.New
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObservationType != nil {
		toSerialize["observationType"] = o.ObservationType
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.RejectionDate != nil {
		toSerialize["rejectionDate"] = o.RejectionDate
	}
	if o.Removable != nil {
		toSerialize["removable"] = o.Removable
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowableObjectObserver struct {
	value *BTWorkflowableObjectObserver
	isSet bool
}

func (v NullableBTWorkflowableObjectObserver) Get() *BTWorkflowableObjectObserver {
	return v.value
}

func (v *NullableBTWorkflowableObjectObserver) Set(val *BTWorkflowableObjectObserver) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowableObjectObserver) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowableObjectObserver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowableObjectObserver(val *BTWorkflowableObjectObserver) *NullableBTWorkflowableObjectObserver {
	return &NullableBTWorkflowableObjectObserver{value: val, isSet: true}
}

func (v NullableBTWorkflowableObjectObserver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowableObjectObserver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}