# BTWorkflowableObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Properties** | Pointer to [**[]BTWorkflowPropertyInfo**](BTWorkflowPropertyInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**Workflow** | Pointer to [**BTWorkflowSnapshotInfo**](BTWorkflowSnapshotInfo.md) |  | [optional] 
**WorkflowError** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to [**BTPublishedWorkflowId**](BTPublishedWorkflowId.md) |  | [optional] 

## Methods

### NewBTWorkflowableObjectInfo

`func NewBTWorkflowableObjectInfo() *BTWorkflowableObjectInfo`

NewBTWorkflowableObjectInfo instantiates a new BTWorkflowableObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowableObjectInfoWithDefaults

`func NewBTWorkflowableObjectInfoWithDefaults() *BTWorkflowableObjectInfo`

NewBTWorkflowableObjectInfoWithDefaults instantiates a new BTWorkflowableObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTWorkflowableObjectInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTWorkflowableObjectInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTWorkflowableObjectInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTWorkflowableObjectInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTWorkflowableObjectInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTWorkflowableObjectInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTWorkflowableObjectInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTWorkflowableObjectInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTWorkflowableObjectInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTWorkflowableObjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTWorkflowableObjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTWorkflowableObjectInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTWorkflowableObjectInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWorkflowableObjectInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWorkflowableObjectInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWorkflowableObjectInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTWorkflowableObjectInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTWorkflowableObjectInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTWorkflowableObjectInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTWorkflowableObjectInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetHref

`func (o *BTWorkflowableObjectInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTWorkflowableObjectInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTWorkflowableObjectInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTWorkflowableObjectInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTWorkflowableObjectInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWorkflowableObjectInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWorkflowableObjectInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWorkflowableObjectInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTWorkflowableObjectInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTWorkflowableObjectInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTWorkflowableObjectInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTWorkflowableObjectInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTWorkflowableObjectInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTWorkflowableObjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTWorkflowableObjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTWorkflowableObjectInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTWorkflowableObjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWorkflowableObjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWorkflowableObjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWorkflowableObjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *BTWorkflowableObjectInfo) GetProperties() []BTWorkflowPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTWorkflowableObjectInfo) GetPropertiesOk() (*[]BTWorkflowPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTWorkflowableObjectInfo) SetProperties(v []BTWorkflowPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTWorkflowableObjectInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetViewRef

`func (o *BTWorkflowableObjectInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTWorkflowableObjectInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTWorkflowableObjectInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTWorkflowableObjectInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkflow

`func (o *BTWorkflowableObjectInfo) GetWorkflow() BTWorkflowSnapshotInfo`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BTWorkflowableObjectInfo) GetWorkflowOk() (*BTWorkflowSnapshotInfo, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BTWorkflowableObjectInfo) SetWorkflow(v BTWorkflowSnapshotInfo)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *BTWorkflowableObjectInfo) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetWorkflowError

`func (o *BTWorkflowableObjectInfo) GetWorkflowError() string`

GetWorkflowError returns the WorkflowError field if non-nil, zero value otherwise.

### GetWorkflowErrorOk

`func (o *BTWorkflowableObjectInfo) GetWorkflowErrorOk() (*string, bool)`

GetWorkflowErrorOk returns a tuple with the WorkflowError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowError

`func (o *BTWorkflowableObjectInfo) SetWorkflowError(v string)`

SetWorkflowError sets WorkflowError field to given value.

### HasWorkflowError

`func (o *BTWorkflowableObjectInfo) HasWorkflowError() bool`

HasWorkflowError returns a boolean if a field has been set.

### GetWorkflowId

`func (o *BTWorkflowableObjectInfo) GetWorkflowId() BTPublishedWorkflowId`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *BTWorkflowableObjectInfo) GetWorkflowIdOk() (*BTPublishedWorkflowId, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *BTWorkflowableObjectInfo) SetWorkflowId(v BTPublishedWorkflowId)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *BTWorkflowableObjectInfo) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


