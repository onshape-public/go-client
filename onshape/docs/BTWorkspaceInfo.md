# BTWorkspaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDelete** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Creator** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**LastModifier** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Microversion** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideDate** | Pointer to **JSONTime** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Parents** | Pointer to [**[]BTVersionInfo**](BTVersionInfo.md) |  | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkspaceInfo

`func NewBTWorkspaceInfo() *BTWorkspaceInfo`

NewBTWorkspaceInfo instantiates a new BTWorkspaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkspaceInfoWithDefaults

`func NewBTWorkspaceInfoWithDefaults() *BTWorkspaceInfo`

NewBTWorkspaceInfoWithDefaults instantiates a new BTWorkspaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDelete

`func (o *BTWorkspaceInfo) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *BTWorkspaceInfo) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *BTWorkspaceInfo) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *BTWorkspaceInfo) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTWorkspaceInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTWorkspaceInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTWorkspaceInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTWorkspaceInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BTWorkspaceInfo) GetCreator() BTUserBasicSummaryInfo`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BTWorkspaceInfo) GetCreatorOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BTWorkspaceInfo) SetCreator(v BTUserBasicSummaryInfo)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BTWorkspaceInfo) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *BTWorkspaceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWorkspaceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWorkspaceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWorkspaceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTWorkspaceInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTWorkspaceInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTWorkspaceInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTWorkspaceInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetHref

`func (o *BTWorkspaceInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTWorkspaceInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTWorkspaceInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTWorkspaceInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTWorkspaceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWorkspaceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWorkspaceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWorkspaceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *BTWorkspaceInfo) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *BTWorkspaceInfo) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *BTWorkspaceInfo) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *BTWorkspaceInfo) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetLastModifier

`func (o *BTWorkspaceInfo) GetLastModifier() BTUserBasicSummaryInfo`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *BTWorkspaceInfo) GetLastModifierOk() (*BTUserBasicSummaryInfo, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *BTWorkspaceInfo) SetLastModifier(v BTUserBasicSummaryInfo)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *BTWorkspaceInfo) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTWorkspaceInfo) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTWorkspaceInfo) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTWorkspaceInfo) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTWorkspaceInfo) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTWorkspaceInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTWorkspaceInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTWorkspaceInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTWorkspaceInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *BTWorkspaceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWorkspaceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWorkspaceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWorkspaceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideDate

`func (o *BTWorkspaceInfo) GetOverrideDate() JSONTime`

GetOverrideDate returns the OverrideDate field if non-nil, zero value otherwise.

### GetOverrideDateOk

`func (o *BTWorkspaceInfo) GetOverrideDateOk() (*JSONTime, bool)`

GetOverrideDateOk returns a tuple with the OverrideDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDate

`func (o *BTWorkspaceInfo) SetOverrideDate(v JSONTime)`

SetOverrideDate sets OverrideDate field to given value.

### HasOverrideDate

`func (o *BTWorkspaceInfo) HasOverrideDate() bool`

HasOverrideDate returns a boolean if a field has been set.

### GetParent

`func (o *BTWorkspaceInfo) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BTWorkspaceInfo) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BTWorkspaceInfo) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BTWorkspaceInfo) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParents

`func (o *BTWorkspaceInfo) GetParents() []BTVersionInfo`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *BTWorkspaceInfo) GetParentsOk() (*[]BTVersionInfo, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *BTWorkspaceInfo) SetParents(v []BTVersionInfo)`

SetParents sets Parents field to given value.

### HasParents

`func (o *BTWorkspaceInfo) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTWorkspaceInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTWorkspaceInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTWorkspaceInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTWorkspaceInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetType

`func (o *BTWorkspaceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTWorkspaceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTWorkspaceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTWorkspaceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetViewRef

`func (o *BTWorkspaceInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTWorkspaceInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTWorkspaceInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTWorkspaceInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


