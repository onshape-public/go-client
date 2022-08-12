# BTWorkflowableTestObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]BTWorkflowPropertyInfo**](BTWorkflowPropertyInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 
**Workflow** | Pointer to [**BTWorkflowSnapshotInfo**](BTWorkflowSnapshotInfo.md) |  | [optional] 
**WorkflowError** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to [**BTPublishedWorkflowId**](BTPublishedWorkflowId.md) |  | [optional] 

## Methods

### NewBTWorkflowableTestObjectInfo

`func NewBTWorkflowableTestObjectInfo() *BTWorkflowableTestObjectInfo`

NewBTWorkflowableTestObjectInfo instantiates a new BTWorkflowableTestObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowableTestObjectInfoWithDefaults

`func NewBTWorkflowableTestObjectInfoWithDefaults() *BTWorkflowableTestObjectInfo`

NewBTWorkflowableTestObjectInfoWithDefaults instantiates a new BTWorkflowableTestObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTWorkflowableTestObjectInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTWorkflowableTestObjectInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTWorkflowableTestObjectInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTWorkflowableTestObjectInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTWorkflowableTestObjectInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTWorkflowableTestObjectInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTWorkflowableTestObjectInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTWorkflowableTestObjectInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTWorkflowableTestObjectInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTWorkflowableTestObjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTWorkflowableTestObjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTWorkflowableTestObjectInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTWorkflowableTestObjectInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWorkflowableTestObjectInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWorkflowableTestObjectInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWorkflowableTestObjectInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTWorkflowableTestObjectInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTWorkflowableTestObjectInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTWorkflowableTestObjectInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTWorkflowableTestObjectInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetHref

`func (o *BTWorkflowableTestObjectInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTWorkflowableTestObjectInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTWorkflowableTestObjectInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTWorkflowableTestObjectInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTWorkflowableTestObjectInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWorkflowableTestObjectInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWorkflowableTestObjectInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWorkflowableTestObjectInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *BTWorkflowableTestObjectInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BTWorkflowableTestObjectInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BTWorkflowableTestObjectInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BTWorkflowableTestObjectInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTWorkflowableTestObjectInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTWorkflowableTestObjectInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTWorkflowableTestObjectInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTWorkflowableTestObjectInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTWorkflowableTestObjectInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTWorkflowableTestObjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTWorkflowableTestObjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTWorkflowableTestObjectInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTWorkflowableTestObjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWorkflowableTestObjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWorkflowableTestObjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWorkflowableTestObjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *BTWorkflowableTestObjectInfo) GetProperties() []BTWorkflowPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTWorkflowableTestObjectInfo) GetPropertiesOk() (*[]BTWorkflowPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTWorkflowableTestObjectInfo) SetProperties(v []BTWorkflowPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTWorkflowableTestObjectInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetViewRef

`func (o *BTWorkflowableTestObjectInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTWorkflowableTestObjectInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTWorkflowableTestObjectInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTWorkflowableTestObjectInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkflow

`func (o *BTWorkflowableTestObjectInfo) GetWorkflow() BTWorkflowSnapshotInfo`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BTWorkflowableTestObjectInfo) GetWorkflowOk() (*BTWorkflowSnapshotInfo, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BTWorkflowableTestObjectInfo) SetWorkflow(v BTWorkflowSnapshotInfo)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *BTWorkflowableTestObjectInfo) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetWorkflowError

`func (o *BTWorkflowableTestObjectInfo) GetWorkflowError() string`

GetWorkflowError returns the WorkflowError field if non-nil, zero value otherwise.

### GetWorkflowErrorOk

`func (o *BTWorkflowableTestObjectInfo) GetWorkflowErrorOk() (*string, bool)`

GetWorkflowErrorOk returns a tuple with the WorkflowError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowError

`func (o *BTWorkflowableTestObjectInfo) SetWorkflowError(v string)`

SetWorkflowError sets WorkflowError field to given value.

### HasWorkflowError

`func (o *BTWorkflowableTestObjectInfo) HasWorkflowError() bool`

HasWorkflowError returns a boolean if a field has been set.

### GetWorkflowId

`func (o *BTWorkflowableTestObjectInfo) GetWorkflowId() BTPublishedWorkflowId`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *BTWorkflowableTestObjectInfo) GetWorkflowIdOk() (*BTPublishedWorkflowId, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *BTWorkflowableTestObjectInfo) SetWorkflowId(v BTPublishedWorkflowId)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *BTWorkflowableTestObjectInfo) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


