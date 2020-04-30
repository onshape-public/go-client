# BTVersionInfoBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Creator** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**LastModifier** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Microversion** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**OverrideDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Parents** | Pointer to [**[]BTVersionInfo**](BTVersionInfo.md) |  | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 

## Methods

### NewBTVersionInfoBase

`func NewBTVersionInfoBase() *BTVersionInfoBase`

NewBTVersionInfoBase instantiates a new BTVersionInfoBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVersionInfoBaseWithDefaults

`func NewBTVersionInfoBaseWithDefaults() *BTVersionInfoBase`

NewBTVersionInfoBaseWithDefaults instantiates a new BTVersionInfoBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BTVersionInfoBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTVersionInfoBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTVersionInfoBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTVersionInfoBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTVersionInfoBase) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTVersionInfoBase) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTVersionInfoBase) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTVersionInfoBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BTVersionInfoBase) GetCreator() BTUserBasicSummaryInfo`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BTVersionInfoBase) GetCreatorOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BTVersionInfoBase) SetCreator(v BTUserBasicSummaryInfo)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BTVersionInfoBase) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *BTVersionInfoBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTVersionInfoBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTVersionInfoBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTVersionInfoBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTVersionInfoBase) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTVersionInfoBase) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTVersionInfoBase) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTVersionInfoBase) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetLastModifier

`func (o *BTVersionInfoBase) GetLastModifier() BTUserBasicSummaryInfo`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *BTVersionInfoBase) GetLastModifierOk() (*BTUserBasicSummaryInfo, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *BTVersionInfoBase) SetLastModifier(v BTUserBasicSummaryInfo)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *BTVersionInfoBase) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTVersionInfoBase) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTVersionInfoBase) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTVersionInfoBase) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTVersionInfoBase) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTVersionInfoBase) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTVersionInfoBase) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTVersionInfoBase) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTVersionInfoBase) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetOverrideDate

`func (o *BTVersionInfoBase) GetOverrideDate() JSONTime`

GetOverrideDate returns the OverrideDate field if non-nil, zero value otherwise.

### GetOverrideDateOk

`func (o *BTVersionInfoBase) GetOverrideDateOk() (*JSONTime, bool)`

GetOverrideDateOk returns a tuple with the OverrideDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDate

`func (o *BTVersionInfoBase) SetOverrideDate(v JSONTime)`

SetOverrideDate sets OverrideDate field to given value.

### HasOverrideDate

`func (o *BTVersionInfoBase) HasOverrideDate() bool`

HasOverrideDate returns a boolean if a field has been set.

### GetParent

`func (o *BTVersionInfoBase) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BTVersionInfoBase) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BTVersionInfoBase) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BTVersionInfoBase) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParents

`func (o *BTVersionInfoBase) GetParents() []BTVersionInfo`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *BTVersionInfoBase) GetParentsOk() (*[]BTVersionInfo, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *BTVersionInfoBase) SetParents(v []BTVersionInfo)`

SetParents sets Parents field to given value.

### HasParents

`func (o *BTVersionInfoBase) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTVersionInfoBase) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTVersionInfoBase) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTVersionInfoBase) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTVersionInfoBase) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


