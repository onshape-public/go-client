# BTDocumentLabelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanMove** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsContainer** | Pointer to **bool** |  | [optional] 
**IsEnterpriseOwned** | Pointer to **bool** |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**ModifiedAt** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**ParentLabelId** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTDocumentLabelInfo

`func NewBTDocumentLabelInfo() *BTDocumentLabelInfo`

NewBTDocumentLabelInfo instantiates a new BTDocumentLabelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentLabelInfoWithDefaults

`func NewBTDocumentLabelInfoWithDefaults() *BTDocumentLabelInfo`

NewBTDocumentLabelInfoWithDefaults instantiates a new BTDocumentLabelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanMove

`func (o *BTDocumentLabelInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTDocumentLabelInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTDocumentLabelInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTDocumentLabelInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTDocumentLabelInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTDocumentLabelInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTDocumentLabelInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTDocumentLabelInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTDocumentLabelInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDocumentLabelInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDocumentLabelInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDocumentLabelInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTDocumentLabelInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDocumentLabelInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDocumentLabelInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDocumentLabelInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *BTDocumentLabelInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTDocumentLabelInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTDocumentLabelInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTDocumentLabelInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTDocumentLabelInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDocumentLabelInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDocumentLabelInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDocumentLabelInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTDocumentLabelInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTDocumentLabelInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTDocumentLabelInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTDocumentLabelInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTDocumentLabelInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTDocumentLabelInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTDocumentLabelInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTDocumentLabelInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTDocumentLabelInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTDocumentLabelInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTDocumentLabelInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTDocumentLabelInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTDocumentLabelInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTDocumentLabelInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTDocumentLabelInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTDocumentLabelInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTDocumentLabelInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTDocumentLabelInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTDocumentLabelInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTDocumentLabelInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentLabelInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentLabelInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentLabelInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentLabelInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *BTDocumentLabelInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTDocumentLabelInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTDocumentLabelInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTDocumentLabelInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentLabelId

`func (o *BTDocumentLabelInfo) GetParentLabelId() string`

GetParentLabelId returns the ParentLabelId field if non-nil, zero value otherwise.

### GetParentLabelIdOk

`func (o *BTDocumentLabelInfo) GetParentLabelIdOk() (*string, bool)`

GetParentLabelIdOk returns a tuple with the ParentLabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLabelId

`func (o *BTDocumentLabelInfo) SetParentLabelId(v string)`

SetParentLabelId sets ParentLabelId field to given value.

### HasParentLabelId

`func (o *BTDocumentLabelInfo) HasParentLabelId() bool`

HasParentLabelId returns a boolean if a field has been set.

### GetPath

`func (o *BTDocumentLabelInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BTDocumentLabelInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BTDocumentLabelInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BTDocumentLabelInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProjectId

`func (o *BTDocumentLabelInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTDocumentLabelInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTDocumentLabelInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTDocumentLabelInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceType

`func (o *BTDocumentLabelInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTDocumentLabelInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTDocumentLabelInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTDocumentLabelInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTDocumentLabelInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTDocumentLabelInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTDocumentLabelInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTDocumentLabelInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetViewRef

`func (o *BTDocumentLabelInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTDocumentLabelInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTDocumentLabelInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTDocumentLabelInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


