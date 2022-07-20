# BTProjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanMove** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsContainer** | Pointer to **bool** |  | [optional] 
**IsEnterpriseOwned** | Pointer to **bool** |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**PermissionScheme** | Pointer to [**BTRbacPermissionSchemeInfo**](BTRbacPermissionSchemeInfo.md) |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**RoleMapEntries** | Pointer to [**[]RoleMapEntry**](RoleMapEntry.md) |  | [optional] 
**Trash** | Pointer to **bool** |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTProjectInfo

`func NewBTProjectInfo() *BTProjectInfo`

NewBTProjectInfo instantiates a new BTProjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTProjectInfoWithDefaults

`func NewBTProjectInfoWithDefaults() *BTProjectInfo`

NewBTProjectInfoWithDefaults instantiates a new BTProjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanMove

`func (o *BTProjectInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTProjectInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTProjectInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTProjectInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTProjectInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTProjectInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTProjectInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTProjectInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTProjectInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTProjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTProjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTProjectInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTProjectInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTProjectInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTProjectInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTProjectInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *BTProjectInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTProjectInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTProjectInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTProjectInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTProjectInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTProjectInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTProjectInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTProjectInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTProjectInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTProjectInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTProjectInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTProjectInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTProjectInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTProjectInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTProjectInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTProjectInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTProjectInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTProjectInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTProjectInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTProjectInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTProjectInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTProjectInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTProjectInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTProjectInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTProjectInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTProjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTProjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTProjectInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTProjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTProjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTProjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTProjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *BTProjectInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTProjectInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTProjectInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTProjectInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPermissionScheme

`func (o *BTProjectInfo) GetPermissionScheme() BTRbacPermissionSchemeInfo`

GetPermissionScheme returns the PermissionScheme field if non-nil, zero value otherwise.

### GetPermissionSchemeOk

`func (o *BTProjectInfo) GetPermissionSchemeOk() (*BTRbacPermissionSchemeInfo, bool)`

GetPermissionSchemeOk returns a tuple with the PermissionScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionScheme

`func (o *BTProjectInfo) SetPermissionScheme(v BTRbacPermissionSchemeInfo)`

SetPermissionScheme sets PermissionScheme field to given value.

### HasPermissionScheme

`func (o *BTProjectInfo) HasPermissionScheme() bool`

HasPermissionScheme returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTProjectInfo) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTProjectInfo) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTProjectInfo) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTProjectInfo) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetProjectId

`func (o *BTProjectInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTProjectInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTProjectInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTProjectInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceType

`func (o *BTProjectInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTProjectInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTProjectInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTProjectInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRoleMapEntries

`func (o *BTProjectInfo) GetRoleMapEntries() []RoleMapEntry`

GetRoleMapEntries returns the RoleMapEntries field if non-nil, zero value otherwise.

### GetRoleMapEntriesOk

`func (o *BTProjectInfo) GetRoleMapEntriesOk() (*[]RoleMapEntry, bool)`

GetRoleMapEntriesOk returns a tuple with the RoleMapEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapEntries

`func (o *BTProjectInfo) SetRoleMapEntries(v []RoleMapEntry)`

SetRoleMapEntries sets RoleMapEntries field to given value.

### HasRoleMapEntries

`func (o *BTProjectInfo) HasRoleMapEntries() bool`

HasRoleMapEntries returns a boolean if a field has been set.

### GetTrash

`func (o *BTProjectInfo) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTProjectInfo) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTProjectInfo) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTProjectInfo) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTProjectInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTProjectInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTProjectInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTProjectInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetViewRef

`func (o *BTProjectInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTProjectInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTProjectInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTProjectInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


