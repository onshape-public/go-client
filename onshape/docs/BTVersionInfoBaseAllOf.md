# BTVersionInfoBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBTVersionInfoBaseAllOf

`func NewBTVersionInfoBaseAllOf() *BTVersionInfoBaseAllOf`

NewBTVersionInfoBaseAllOf instantiates a new BTVersionInfoBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVersionInfoBaseAllOfWithDefaults

`func NewBTVersionInfoBaseAllOfWithDefaults() *BTVersionInfoBaseAllOf`

NewBTVersionInfoBaseAllOfWithDefaults instantiates a new BTVersionInfoBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BTVersionInfoBaseAllOf) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTVersionInfoBaseAllOf) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTVersionInfoBaseAllOf) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTVersionInfoBaseAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BTVersionInfoBaseAllOf) GetCreator() BTUserBasicSummaryInfo`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BTVersionInfoBaseAllOf) GetCreatorOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BTVersionInfoBaseAllOf) SetCreator(v BTUserBasicSummaryInfo)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BTVersionInfoBaseAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *BTVersionInfoBaseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTVersionInfoBaseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTVersionInfoBaseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTVersionInfoBaseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTVersionInfoBaseAllOf) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTVersionInfoBaseAllOf) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTVersionInfoBaseAllOf) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTVersionInfoBaseAllOf) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetLastModifier

`func (o *BTVersionInfoBaseAllOf) GetLastModifier() BTUserBasicSummaryInfo`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *BTVersionInfoBaseAllOf) GetLastModifierOk() (*BTUserBasicSummaryInfo, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *BTVersionInfoBaseAllOf) SetLastModifier(v BTUserBasicSummaryInfo)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *BTVersionInfoBaseAllOf) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTVersionInfoBaseAllOf) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTVersionInfoBaseAllOf) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTVersionInfoBaseAllOf) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTVersionInfoBaseAllOf) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTVersionInfoBaseAllOf) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTVersionInfoBaseAllOf) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTVersionInfoBaseAllOf) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTVersionInfoBaseAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetOverrideDate

`func (o *BTVersionInfoBaseAllOf) GetOverrideDate() JSONTime`

GetOverrideDate returns the OverrideDate field if non-nil, zero value otherwise.

### GetOverrideDateOk

`func (o *BTVersionInfoBaseAllOf) GetOverrideDateOk() (*JSONTime, bool)`

GetOverrideDateOk returns a tuple with the OverrideDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDate

`func (o *BTVersionInfoBaseAllOf) SetOverrideDate(v JSONTime)`

SetOverrideDate sets OverrideDate field to given value.

### HasOverrideDate

`func (o *BTVersionInfoBaseAllOf) HasOverrideDate() bool`

HasOverrideDate returns a boolean if a field has been set.

### GetParent

`func (o *BTVersionInfoBaseAllOf) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BTVersionInfoBaseAllOf) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BTVersionInfoBaseAllOf) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BTVersionInfoBaseAllOf) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParents

`func (o *BTVersionInfoBaseAllOf) GetParents() []BTVersionInfo`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *BTVersionInfoBaseAllOf) GetParentsOk() (*[]BTVersionInfo, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *BTVersionInfoBaseAllOf) SetParents(v []BTVersionInfo)`

SetParents sets Parents field to given value.

### HasParents

`func (o *BTVersionInfoBaseAllOf) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTVersionInfoBaseAllOf) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTVersionInfoBaseAllOf) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTVersionInfoBaseAllOf) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTVersionInfoBaseAllOf) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetType

`func (o *BTVersionInfoBaseAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTVersionInfoBaseAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTVersionInfoBaseAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTVersionInfoBaseAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


