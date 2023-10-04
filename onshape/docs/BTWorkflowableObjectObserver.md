# BTWorkflowableObjectObserver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminOverride** | Pointer to **bool** |  | [optional] 
**ApprovalDate** | Pointer to **JSONTime** |  | [optional] 
**ApprovalState** | Pointer to [**BTWorkflowObserverState**](BTWorkflowObserverState.md) |  | [optional] 
**ApproverId** | Pointer to **string** |  | [optional] 
**ApproverName** | Pointer to **string** |  | [optional] 
**AssociatedStates** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**EntryType** | Pointer to [**BTWorkflowObserverEntryType**](BTWorkflowObserverEntryType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**New** | Pointer to **bool** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObservationType** | Pointer to **int32** |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**RejectionDate** | Pointer to **JSONTime** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTWorkflowableObjectObserver

`func NewBTWorkflowableObjectObserver() *BTWorkflowableObjectObserver`

NewBTWorkflowableObjectObserver instantiates a new BTWorkflowableObjectObserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowableObjectObserverWithDefaults

`func NewBTWorkflowableObjectObserverWithDefaults() *BTWorkflowableObjectObserver`

NewBTWorkflowableObjectObserverWithDefaults instantiates a new BTWorkflowableObjectObserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminOverride

`func (o *BTWorkflowableObjectObserver) GetAdminOverride() bool`

GetAdminOverride returns the AdminOverride field if non-nil, zero value otherwise.

### GetAdminOverrideOk

`func (o *BTWorkflowableObjectObserver) GetAdminOverrideOk() (*bool, bool)`

GetAdminOverrideOk returns a tuple with the AdminOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOverride

`func (o *BTWorkflowableObjectObserver) SetAdminOverride(v bool)`

SetAdminOverride sets AdminOverride field to given value.

### HasAdminOverride

`func (o *BTWorkflowableObjectObserver) HasAdminOverride() bool`

HasAdminOverride returns a boolean if a field has been set.

### GetApprovalDate

`func (o *BTWorkflowableObjectObserver) GetApprovalDate() JSONTime`

GetApprovalDate returns the ApprovalDate field if non-nil, zero value otherwise.

### GetApprovalDateOk

`func (o *BTWorkflowableObjectObserver) GetApprovalDateOk() (*JSONTime, bool)`

GetApprovalDateOk returns a tuple with the ApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalDate

`func (o *BTWorkflowableObjectObserver) SetApprovalDate(v JSONTime)`

SetApprovalDate sets ApprovalDate field to given value.

### HasApprovalDate

`func (o *BTWorkflowableObjectObserver) HasApprovalDate() bool`

HasApprovalDate returns a boolean if a field has been set.

### GetApprovalState

`func (o *BTWorkflowableObjectObserver) GetApprovalState() BTWorkflowObserverState`

GetApprovalState returns the ApprovalState field if non-nil, zero value otherwise.

### GetApprovalStateOk

`func (o *BTWorkflowableObjectObserver) GetApprovalStateOk() (*BTWorkflowObserverState, bool)`

GetApprovalStateOk returns a tuple with the ApprovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalState

`func (o *BTWorkflowableObjectObserver) SetApprovalState(v BTWorkflowObserverState)`

SetApprovalState sets ApprovalState field to given value.

### HasApprovalState

`func (o *BTWorkflowableObjectObserver) HasApprovalState() bool`

HasApprovalState returns a boolean if a field has been set.

### GetApproverId

`func (o *BTWorkflowableObjectObserver) GetApproverId() string`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *BTWorkflowableObjectObserver) GetApproverIdOk() (*string, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *BTWorkflowableObjectObserver) SetApproverId(v string)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *BTWorkflowableObjectObserver) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetApproverName

`func (o *BTWorkflowableObjectObserver) GetApproverName() string`

GetApproverName returns the ApproverName field if non-nil, zero value otherwise.

### GetApproverNameOk

`func (o *BTWorkflowableObjectObserver) GetApproverNameOk() (*string, bool)`

GetApproverNameOk returns a tuple with the ApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverName

`func (o *BTWorkflowableObjectObserver) SetApproverName(v string)`

SetApproverName sets ApproverName field to given value.

### HasApproverName

`func (o *BTWorkflowableObjectObserver) HasApproverName() bool`

HasApproverName returns a boolean if a field has been set.

### GetAssociatedStates

`func (o *BTWorkflowableObjectObserver) GetAssociatedStates() string`

GetAssociatedStates returns the AssociatedStates field if non-nil, zero value otherwise.

### GetAssociatedStatesOk

`func (o *BTWorkflowableObjectObserver) GetAssociatedStatesOk() (*string, bool)`

GetAssociatedStatesOk returns a tuple with the AssociatedStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedStates

`func (o *BTWorkflowableObjectObserver) SetAssociatedStates(v string)`

SetAssociatedStates sets AssociatedStates field to given value.

### HasAssociatedStates

`func (o *BTWorkflowableObjectObserver) HasAssociatedStates() bool`

HasAssociatedStates returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTWorkflowableObjectObserver) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTWorkflowableObjectObserver) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTWorkflowableObjectObserver) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTWorkflowableObjectObserver) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTWorkflowableObjectObserver) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTWorkflowableObjectObserver) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTWorkflowableObjectObserver) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTWorkflowableObjectObserver) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTWorkflowableObjectObserver) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTWorkflowableObjectObserver) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTWorkflowableObjectObserver) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTWorkflowableObjectObserver) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTWorkflowableObjectObserver) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWorkflowableObjectObserver) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWorkflowableObjectObserver) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWorkflowableObjectObserver) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntryId

`func (o *BTWorkflowableObjectObserver) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BTWorkflowableObjectObserver) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BTWorkflowableObjectObserver) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *BTWorkflowableObjectObserver) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEntryType

`func (o *BTWorkflowableObjectObserver) GetEntryType() BTWorkflowObserverEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BTWorkflowableObjectObserver) GetEntryTypeOk() (*BTWorkflowObserverEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BTWorkflowableObjectObserver) SetEntryType(v BTWorkflowObserverEntryType)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *BTWorkflowableObjectObserver) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetId

`func (o *BTWorkflowableObjectObserver) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWorkflowableObjectObserver) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWorkflowableObjectObserver) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWorkflowableObjectObserver) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTWorkflowableObjectObserver) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTWorkflowableObjectObserver) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTWorkflowableObjectObserver) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTWorkflowableObjectObserver) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTWorkflowableObjectObserver) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTWorkflowableObjectObserver) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTWorkflowableObjectObserver) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTWorkflowableObjectObserver) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTWorkflowableObjectObserver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWorkflowableObjectObserver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWorkflowableObjectObserver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWorkflowableObjectObserver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNew

`func (o *BTWorkflowableObjectObserver) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *BTWorkflowableObjectObserver) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *BTWorkflowableObjectObserver) SetNew(v bool)`

SetNew sets New field to given value.

### HasNew

`func (o *BTWorkflowableObjectObserver) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetObjectId

`func (o *BTWorkflowableObjectObserver) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTWorkflowableObjectObserver) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTWorkflowableObjectObserver) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTWorkflowableObjectObserver) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObservationType

`func (o *BTWorkflowableObjectObserver) GetObservationType() int32`

GetObservationType returns the ObservationType field if non-nil, zero value otherwise.

### GetObservationTypeOk

`func (o *BTWorkflowableObjectObserver) GetObservationTypeOk() (*int32, bool)`

GetObservationTypeOk returns a tuple with the ObservationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationType

`func (o *BTWorkflowableObjectObserver) SetObservationType(v int32)`

SetObservationType sets ObservationType field to given value.

### HasObservationType

`func (o *BTWorkflowableObjectObserver) HasObservationType() bool`

HasObservationType returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTWorkflowableObjectObserver) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTWorkflowableObjectObserver) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTWorkflowableObjectObserver) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTWorkflowableObjectObserver) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetRejectionDate

`func (o *BTWorkflowableObjectObserver) GetRejectionDate() JSONTime`

GetRejectionDate returns the RejectionDate field if non-nil, zero value otherwise.

### GetRejectionDateOk

`func (o *BTWorkflowableObjectObserver) GetRejectionDateOk() (*JSONTime, bool)`

GetRejectionDateOk returns a tuple with the RejectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionDate

`func (o *BTWorkflowableObjectObserver) SetRejectionDate(v JSONTime)`

SetRejectionDate sets RejectionDate field to given value.

### HasRejectionDate

`func (o *BTWorkflowableObjectObserver) HasRejectionDate() bool`

HasRejectionDate returns a boolean if a field has been set.

### GetRemovable

`func (o *BTWorkflowableObjectObserver) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *BTWorkflowableObjectObserver) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *BTWorkflowableObjectObserver) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *BTWorkflowableObjectObserver) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


