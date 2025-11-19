# BTTeamSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**CanMove** | Pointer to **bool** |  | [optional] 
**ConnectionName** | Pointer to **string** |  | [optional] 
**ConnectionNames** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsContainer** | Pointer to **bool** |  | [optional] 
**IsEnterpriseOwned** | Pointer to **bool** |  | [optional] 
**IsExternalConnectionResource** | Pointer to **bool** |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**PredefinedTeam** | Pointer to **int32** |  | [optional] 
**PredefinedTeamMutable** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**UnparentHref** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTTeamSummaryInfo

`func NewBTTeamSummaryInfo() *BTTeamSummaryInfo`

NewBTTeamSummaryInfo instantiates a new BTTeamSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTeamSummaryInfoWithDefaults

`func NewBTTeamSummaryInfoWithDefaults() *BTTeamSummaryInfo`

NewBTTeamSummaryInfoWithDefaults instantiates a new BTTeamSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BTTeamSummaryInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BTTeamSummaryInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BTTeamSummaryInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BTTeamSummaryInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCanMove

`func (o *BTTeamSummaryInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTTeamSummaryInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTTeamSummaryInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTTeamSummaryInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetConnectionName

`func (o *BTTeamSummaryInfo) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *BTTeamSummaryInfo) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *BTTeamSummaryInfo) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *BTTeamSummaryInfo) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetConnectionNames

`func (o *BTTeamSummaryInfo) GetConnectionNames() []string`

GetConnectionNames returns the ConnectionNames field if non-nil, zero value otherwise.

### GetConnectionNamesOk

`func (o *BTTeamSummaryInfo) GetConnectionNamesOk() (*[]string, bool)`

GetConnectionNamesOk returns a tuple with the ConnectionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionNames

`func (o *BTTeamSummaryInfo) SetConnectionNames(v []string)`

SetConnectionNames sets ConnectionNames field to given value.

### HasConnectionNames

`func (o *BTTeamSummaryInfo) HasConnectionNames() bool`

HasConnectionNames returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTTeamSummaryInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTTeamSummaryInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTTeamSummaryInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTTeamSummaryInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTTeamSummaryInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTTeamSummaryInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTTeamSummaryInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTTeamSummaryInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTTeamSummaryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTTeamSummaryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTTeamSummaryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTTeamSummaryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *BTTeamSummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTeamSummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTeamSummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTeamSummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTeamSummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTeamSummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTeamSummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTeamSummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTTeamSummaryInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTTeamSummaryInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTTeamSummaryInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTTeamSummaryInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTTeamSummaryInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTTeamSummaryInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTTeamSummaryInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTTeamSummaryInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsExternalConnectionResource

`func (o *BTTeamSummaryInfo) GetIsExternalConnectionResource() bool`

GetIsExternalConnectionResource returns the IsExternalConnectionResource field if non-nil, zero value otherwise.

### GetIsExternalConnectionResourceOk

`func (o *BTTeamSummaryInfo) GetIsExternalConnectionResourceOk() (*bool, bool)`

GetIsExternalConnectionResourceOk returns a tuple with the IsExternalConnectionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalConnectionResource

`func (o *BTTeamSummaryInfo) SetIsExternalConnectionResource(v bool)`

SetIsExternalConnectionResource sets IsExternalConnectionResource field to given value.

### HasIsExternalConnectionResource

`func (o *BTTeamSummaryInfo) HasIsExternalConnectionResource() bool`

HasIsExternalConnectionResource returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTTeamSummaryInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTTeamSummaryInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTTeamSummaryInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTTeamSummaryInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTTeamSummaryInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTTeamSummaryInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTTeamSummaryInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTTeamSummaryInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTTeamSummaryInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTTeamSummaryInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTTeamSummaryInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTTeamSummaryInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTTeamSummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTeamSummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTeamSummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTeamSummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *BTTeamSummaryInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTTeamSummaryInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTTeamSummaryInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTTeamSummaryInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentId

`func (o *BTTeamSummaryInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTTeamSummaryInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTTeamSummaryInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTTeamSummaryInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPredefinedTeam

`func (o *BTTeamSummaryInfo) GetPredefinedTeam() int32`

GetPredefinedTeam returns the PredefinedTeam field if non-nil, zero value otherwise.

### GetPredefinedTeamOk

`func (o *BTTeamSummaryInfo) GetPredefinedTeamOk() (*int32, bool)`

GetPredefinedTeamOk returns a tuple with the PredefinedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTeam

`func (o *BTTeamSummaryInfo) SetPredefinedTeam(v int32)`

SetPredefinedTeam sets PredefinedTeam field to given value.

### HasPredefinedTeam

`func (o *BTTeamSummaryInfo) HasPredefinedTeam() bool`

HasPredefinedTeam returns a boolean if a field has been set.

### GetPredefinedTeamMutable

`func (o *BTTeamSummaryInfo) GetPredefinedTeamMutable() bool`

GetPredefinedTeamMutable returns the PredefinedTeamMutable field if non-nil, zero value otherwise.

### GetPredefinedTeamMutableOk

`func (o *BTTeamSummaryInfo) GetPredefinedTeamMutableOk() (*bool, bool)`

GetPredefinedTeamMutableOk returns a tuple with the PredefinedTeamMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTeamMutable

`func (o *BTTeamSummaryInfo) SetPredefinedTeamMutable(v bool)`

SetPredefinedTeamMutable sets PredefinedTeamMutable field to given value.

### HasPredefinedTeamMutable

`func (o *BTTeamSummaryInfo) HasPredefinedTeamMutable() bool`

HasPredefinedTeamMutable returns a boolean if a field has been set.

### GetProjectId

`func (o *BTTeamSummaryInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTTeamSummaryInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTTeamSummaryInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTTeamSummaryInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceType

`func (o *BTTeamSummaryInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTTeamSummaryInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTTeamSummaryInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTTeamSummaryInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTTeamSummaryInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTTeamSummaryInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTTeamSummaryInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTTeamSummaryInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetUnparentHref

`func (o *BTTeamSummaryInfo) GetUnparentHref() string`

GetUnparentHref returns the UnparentHref field if non-nil, zero value otherwise.

### GetUnparentHrefOk

`func (o *BTTeamSummaryInfo) GetUnparentHrefOk() (*string, bool)`

GetUnparentHrefOk returns a tuple with the UnparentHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparentHref

`func (o *BTTeamSummaryInfo) SetUnparentHref(v string)`

SetUnparentHref sets UnparentHref field to given value.

### HasUnparentHref

`func (o *BTTeamSummaryInfo) HasUnparentHref() bool`

HasUnparentHref returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTeamSummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTeamSummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTeamSummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTeamSummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


