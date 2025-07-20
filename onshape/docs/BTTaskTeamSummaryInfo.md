# BTTaskTeamSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acted** | Pointer to **bool** |  | [optional] 
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
**PredefinedTeam** | Pointer to **int32** |  | [optional] 
**PredefinedTeamMutable** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**UnparentHref** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTTaskTeamSummaryInfo

`func NewBTTaskTeamSummaryInfo() *BTTaskTeamSummaryInfo`

NewBTTaskTeamSummaryInfo instantiates a new BTTaskTeamSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTaskTeamSummaryInfoWithDefaults

`func NewBTTaskTeamSummaryInfoWithDefaults() *BTTaskTeamSummaryInfo`

NewBTTaskTeamSummaryInfoWithDefaults instantiates a new BTTaskTeamSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActed

`func (o *BTTaskTeamSummaryInfo) GetActed() bool`

GetActed returns the Acted field if non-nil, zero value otherwise.

### GetActedOk

`func (o *BTTaskTeamSummaryInfo) GetActedOk() (*bool, bool)`

GetActedOk returns a tuple with the Acted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActed

`func (o *BTTaskTeamSummaryInfo) SetActed(v bool)`

SetActed sets Acted field to given value.

### HasActed

`func (o *BTTaskTeamSummaryInfo) HasActed() bool`

HasActed returns a boolean if a field has been set.

### GetActive

`func (o *BTTaskTeamSummaryInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BTTaskTeamSummaryInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BTTaskTeamSummaryInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BTTaskTeamSummaryInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCanMove

`func (o *BTTaskTeamSummaryInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTTaskTeamSummaryInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTTaskTeamSummaryInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTTaskTeamSummaryInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetConnectionName

`func (o *BTTaskTeamSummaryInfo) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *BTTaskTeamSummaryInfo) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *BTTaskTeamSummaryInfo) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *BTTaskTeamSummaryInfo) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetConnectionNames

`func (o *BTTaskTeamSummaryInfo) GetConnectionNames() []string`

GetConnectionNames returns the ConnectionNames field if non-nil, zero value otherwise.

### GetConnectionNamesOk

`func (o *BTTaskTeamSummaryInfo) GetConnectionNamesOk() (*[]string, bool)`

GetConnectionNamesOk returns a tuple with the ConnectionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionNames

`func (o *BTTaskTeamSummaryInfo) SetConnectionNames(v []string)`

SetConnectionNames sets ConnectionNames field to given value.

### HasConnectionNames

`func (o *BTTaskTeamSummaryInfo) HasConnectionNames() bool`

HasConnectionNames returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTTaskTeamSummaryInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTTaskTeamSummaryInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTTaskTeamSummaryInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTTaskTeamSummaryInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTTaskTeamSummaryInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTTaskTeamSummaryInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTTaskTeamSummaryInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTTaskTeamSummaryInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTTaskTeamSummaryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTTaskTeamSummaryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTTaskTeamSummaryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTTaskTeamSummaryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *BTTaskTeamSummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTaskTeamSummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTaskTeamSummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTaskTeamSummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTaskTeamSummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTaskTeamSummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTaskTeamSummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTaskTeamSummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTTaskTeamSummaryInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTTaskTeamSummaryInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTTaskTeamSummaryInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTTaskTeamSummaryInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTTaskTeamSummaryInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTTaskTeamSummaryInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTTaskTeamSummaryInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTTaskTeamSummaryInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsExternalConnectionResource

`func (o *BTTaskTeamSummaryInfo) GetIsExternalConnectionResource() bool`

GetIsExternalConnectionResource returns the IsExternalConnectionResource field if non-nil, zero value otherwise.

### GetIsExternalConnectionResourceOk

`func (o *BTTaskTeamSummaryInfo) GetIsExternalConnectionResourceOk() (*bool, bool)`

GetIsExternalConnectionResourceOk returns a tuple with the IsExternalConnectionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalConnectionResource

`func (o *BTTaskTeamSummaryInfo) SetIsExternalConnectionResource(v bool)`

SetIsExternalConnectionResource sets IsExternalConnectionResource field to given value.

### HasIsExternalConnectionResource

`func (o *BTTaskTeamSummaryInfo) HasIsExternalConnectionResource() bool`

HasIsExternalConnectionResource returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTTaskTeamSummaryInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTTaskTeamSummaryInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTTaskTeamSummaryInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTTaskTeamSummaryInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTTaskTeamSummaryInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTTaskTeamSummaryInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTTaskTeamSummaryInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTTaskTeamSummaryInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTTaskTeamSummaryInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTTaskTeamSummaryInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTTaskTeamSummaryInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTTaskTeamSummaryInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTTaskTeamSummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTaskTeamSummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTaskTeamSummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTaskTeamSummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *BTTaskTeamSummaryInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTTaskTeamSummaryInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTTaskTeamSummaryInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTTaskTeamSummaryInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPredefinedTeam

`func (o *BTTaskTeamSummaryInfo) GetPredefinedTeam() int32`

GetPredefinedTeam returns the PredefinedTeam field if non-nil, zero value otherwise.

### GetPredefinedTeamOk

`func (o *BTTaskTeamSummaryInfo) GetPredefinedTeamOk() (*int32, bool)`

GetPredefinedTeamOk returns a tuple with the PredefinedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTeam

`func (o *BTTaskTeamSummaryInfo) SetPredefinedTeam(v int32)`

SetPredefinedTeam sets PredefinedTeam field to given value.

### HasPredefinedTeam

`func (o *BTTaskTeamSummaryInfo) HasPredefinedTeam() bool`

HasPredefinedTeam returns a boolean if a field has been set.

### GetPredefinedTeamMutable

`func (o *BTTaskTeamSummaryInfo) GetPredefinedTeamMutable() bool`

GetPredefinedTeamMutable returns the PredefinedTeamMutable field if non-nil, zero value otherwise.

### GetPredefinedTeamMutableOk

`func (o *BTTaskTeamSummaryInfo) GetPredefinedTeamMutableOk() (*bool, bool)`

GetPredefinedTeamMutableOk returns a tuple with the PredefinedTeamMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTeamMutable

`func (o *BTTaskTeamSummaryInfo) SetPredefinedTeamMutable(v bool)`

SetPredefinedTeamMutable sets PredefinedTeamMutable field to given value.

### HasPredefinedTeamMutable

`func (o *BTTaskTeamSummaryInfo) HasPredefinedTeamMutable() bool`

HasPredefinedTeamMutable returns a boolean if a field has been set.

### GetProjectId

`func (o *BTTaskTeamSummaryInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTTaskTeamSummaryInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTTaskTeamSummaryInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTTaskTeamSummaryInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceType

`func (o *BTTaskTeamSummaryInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTTaskTeamSummaryInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTTaskTeamSummaryInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTTaskTeamSummaryInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTTaskTeamSummaryInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTTaskTeamSummaryInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTTaskTeamSummaryInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTTaskTeamSummaryInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetUnparentHref

`func (o *BTTaskTeamSummaryInfo) GetUnparentHref() string`

GetUnparentHref returns the UnparentHref field if non-nil, zero value otherwise.

### GetUnparentHrefOk

`func (o *BTTaskTeamSummaryInfo) GetUnparentHrefOk() (*string, bool)`

GetUnparentHrefOk returns a tuple with the UnparentHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparentHref

`func (o *BTTaskTeamSummaryInfo) SetUnparentHref(v string)`

SetUnparentHref sets UnparentHref field to given value.

### HasUnparentHref

`func (o *BTTaskTeamSummaryInfo) HasUnparentHref() bool`

HasUnparentHref returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTaskTeamSummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTaskTeamSummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTaskTeamSummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTaskTeamSummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


