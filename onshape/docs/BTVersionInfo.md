# BTVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Creator** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**LastModifier** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**MetadataWorkspaceId** | Pointer to **string** |  | [optional] 
**Microversion** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**OverrideDate** | Pointer to **JSONTime** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Purpose** | Pointer to **int32** |  | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTVersionInfo

`func NewBTVersionInfo() *BTVersionInfo`

NewBTVersionInfo instantiates a new BTVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVersionInfoWithDefaults

`func NewBTVersionInfoWithDefaults() *BTVersionInfo`

NewBTVersionInfoWithDefaults instantiates a new BTVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BTVersionInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTVersionInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTVersionInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTVersionInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BTVersionInfo) GetCreator() BTUserBasicSummaryInfo`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BTVersionInfo) GetCreatorOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BTVersionInfo) SetCreator(v BTUserBasicSummaryInfo)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BTVersionInfo) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *BTVersionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTVersionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTVersionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTVersionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTVersionInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTVersionInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTVersionInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTVersionInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetHref

`func (o *BTVersionInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTVersionInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTVersionInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTVersionInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTVersionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTVersionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTVersionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTVersionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifier

`func (o *BTVersionInfo) GetLastModifier() BTUserBasicSummaryInfo`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *BTVersionInfo) GetLastModifierOk() (*BTUserBasicSummaryInfo, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *BTVersionInfo) SetLastModifier(v BTUserBasicSummaryInfo)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *BTVersionInfo) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetMetadataWorkspaceId

`func (o *BTVersionInfo) GetMetadataWorkspaceId() string`

GetMetadataWorkspaceId returns the MetadataWorkspaceId field if non-nil, zero value otherwise.

### GetMetadataWorkspaceIdOk

`func (o *BTVersionInfo) GetMetadataWorkspaceIdOk() (*string, bool)`

GetMetadataWorkspaceIdOk returns a tuple with the MetadataWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataWorkspaceId

`func (o *BTVersionInfo) SetMetadataWorkspaceId(v string)`

SetMetadataWorkspaceId sets MetadataWorkspaceId field to given value.

### HasMetadataWorkspaceId

`func (o *BTVersionInfo) HasMetadataWorkspaceId() bool`

HasMetadataWorkspaceId returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTVersionInfo) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTVersionInfo) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTVersionInfo) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTVersionInfo) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTVersionInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTVersionInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTVersionInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTVersionInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *BTVersionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTVersionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTVersionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTVersionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideDate

`func (o *BTVersionInfo) GetOverrideDate() JSONTime`

GetOverrideDate returns the OverrideDate field if non-nil, zero value otherwise.

### GetOverrideDateOk

`func (o *BTVersionInfo) GetOverrideDateOk() (*JSONTime, bool)`

GetOverrideDateOk returns a tuple with the OverrideDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDate

`func (o *BTVersionInfo) SetOverrideDate(v JSONTime)`

SetOverrideDate sets OverrideDate field to given value.

### HasOverrideDate

`func (o *BTVersionInfo) HasOverrideDate() bool`

HasOverrideDate returns a boolean if a field has been set.

### GetParent

`func (o *BTVersionInfo) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BTVersionInfo) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BTVersionInfo) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BTVersionInfo) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPurpose

`func (o *BTVersionInfo) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *BTVersionInfo) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *BTVersionInfo) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *BTVersionInfo) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTVersionInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTVersionInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTVersionInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTVersionInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetType

`func (o *BTVersionInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTVersionInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTVersionInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTVersionInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetViewRef

`func (o *BTVersionInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTVersionInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTVersionInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTVersionInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


