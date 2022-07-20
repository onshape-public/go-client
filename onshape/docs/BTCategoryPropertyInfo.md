# BTCategoryPropertyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Array** | Pointer to **bool** |  | [optional] 
**Assignable** | Pointer to **bool** |  | [optional] 
**BlobMimeType** | Pointer to **string** |  | [optional] 
**CategoryPropertyConfigInfo** | Pointer to [**BTCategoryPropertyConfigInfo**](BTCategoryPropertyConfigInfo.md) |  | [optional] 
**CategorySummaryInfoList** | Pointer to [**[]BTMetadataCategorySummaryInfo**](BTMetadataCategorySummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EditableInMicroversion** | Pointer to **bool** |  | [optional] 
**EditableInVersion** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectDefName** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**UiReadonlyInMicroversion** | Pointer to **bool** |  | [optional] 
**UiReadonlyInVersion** | Pointer to **bool** |  | [optional] 
**ValueType** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTCategoryPropertyInfo

`func NewBTCategoryPropertyInfo() *BTCategoryPropertyInfo`

NewBTCategoryPropertyInfo instantiates a new BTCategoryPropertyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCategoryPropertyInfoWithDefaults

`func NewBTCategoryPropertyInfoWithDefaults() *BTCategoryPropertyInfo`

NewBTCategoryPropertyInfoWithDefaults instantiates a new BTCategoryPropertyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArray

`func (o *BTCategoryPropertyInfo) GetArray() bool`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *BTCategoryPropertyInfo) GetArrayOk() (*bool, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *BTCategoryPropertyInfo) SetArray(v bool)`

SetArray sets Array field to given value.

### HasArray

`func (o *BTCategoryPropertyInfo) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetAssignable

`func (o *BTCategoryPropertyInfo) GetAssignable() bool`

GetAssignable returns the Assignable field if non-nil, zero value otherwise.

### GetAssignableOk

`func (o *BTCategoryPropertyInfo) GetAssignableOk() (*bool, bool)`

GetAssignableOk returns a tuple with the Assignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignable

`func (o *BTCategoryPropertyInfo) SetAssignable(v bool)`

SetAssignable sets Assignable field to given value.

### HasAssignable

`func (o *BTCategoryPropertyInfo) HasAssignable() bool`

HasAssignable returns a boolean if a field has been set.

### GetBlobMimeType

`func (o *BTCategoryPropertyInfo) GetBlobMimeType() string`

GetBlobMimeType returns the BlobMimeType field if non-nil, zero value otherwise.

### GetBlobMimeTypeOk

`func (o *BTCategoryPropertyInfo) GetBlobMimeTypeOk() (*string, bool)`

GetBlobMimeTypeOk returns a tuple with the BlobMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobMimeType

`func (o *BTCategoryPropertyInfo) SetBlobMimeType(v string)`

SetBlobMimeType sets BlobMimeType field to given value.

### HasBlobMimeType

`func (o *BTCategoryPropertyInfo) HasBlobMimeType() bool`

HasBlobMimeType returns a boolean if a field has been set.

### GetCategoryPropertyConfigInfo

`func (o *BTCategoryPropertyInfo) GetCategoryPropertyConfigInfo() BTCategoryPropertyConfigInfo`

GetCategoryPropertyConfigInfo returns the CategoryPropertyConfigInfo field if non-nil, zero value otherwise.

### GetCategoryPropertyConfigInfoOk

`func (o *BTCategoryPropertyInfo) GetCategoryPropertyConfigInfoOk() (*BTCategoryPropertyConfigInfo, bool)`

GetCategoryPropertyConfigInfoOk returns a tuple with the CategoryPropertyConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPropertyConfigInfo

`func (o *BTCategoryPropertyInfo) SetCategoryPropertyConfigInfo(v BTCategoryPropertyConfigInfo)`

SetCategoryPropertyConfigInfo sets CategoryPropertyConfigInfo field to given value.

### HasCategoryPropertyConfigInfo

`func (o *BTCategoryPropertyInfo) HasCategoryPropertyConfigInfo() bool`

HasCategoryPropertyConfigInfo returns a boolean if a field has been set.

### GetCategorySummaryInfoList

`func (o *BTCategoryPropertyInfo) GetCategorySummaryInfoList() []BTMetadataCategorySummaryInfo`

GetCategorySummaryInfoList returns the CategorySummaryInfoList field if non-nil, zero value otherwise.

### GetCategorySummaryInfoListOk

`func (o *BTCategoryPropertyInfo) GetCategorySummaryInfoListOk() (*[]BTMetadataCategorySummaryInfo, bool)`

GetCategorySummaryInfoListOk returns a tuple with the CategorySummaryInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySummaryInfoList

`func (o *BTCategoryPropertyInfo) SetCategorySummaryInfoList(v []BTMetadataCategorySummaryInfo)`

SetCategorySummaryInfoList sets CategorySummaryInfoList field to given value.

### HasCategorySummaryInfoList

`func (o *BTCategoryPropertyInfo) HasCategorySummaryInfoList() bool`

HasCategorySummaryInfoList returns a boolean if a field has been set.

### GetDescription

`func (o *BTCategoryPropertyInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTCategoryPropertyInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTCategoryPropertyInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTCategoryPropertyInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditableInMicroversion

`func (o *BTCategoryPropertyInfo) GetEditableInMicroversion() bool`

GetEditableInMicroversion returns the EditableInMicroversion field if non-nil, zero value otherwise.

### GetEditableInMicroversionOk

`func (o *BTCategoryPropertyInfo) GetEditableInMicroversionOk() (*bool, bool)`

GetEditableInMicroversionOk returns a tuple with the EditableInMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInMicroversion

`func (o *BTCategoryPropertyInfo) SetEditableInMicroversion(v bool)`

SetEditableInMicroversion sets EditableInMicroversion field to given value.

### HasEditableInMicroversion

`func (o *BTCategoryPropertyInfo) HasEditableInMicroversion() bool`

HasEditableInMicroversion returns a boolean if a field has been set.

### GetEditableInVersion

`func (o *BTCategoryPropertyInfo) GetEditableInVersion() bool`

GetEditableInVersion returns the EditableInVersion field if non-nil, zero value otherwise.

### GetEditableInVersionOk

`func (o *BTCategoryPropertyInfo) GetEditableInVersionOk() (*bool, bool)`

GetEditableInVersionOk returns a tuple with the EditableInVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInVersion

`func (o *BTCategoryPropertyInfo) SetEditableInVersion(v bool)`

SetEditableInVersion sets EditableInVersion field to given value.

### HasEditableInVersion

`func (o *BTCategoryPropertyInfo) HasEditableInVersion() bool`

HasEditableInVersion returns a boolean if a field has been set.

### GetHref

`func (o *BTCategoryPropertyInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTCategoryPropertyInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTCategoryPropertyInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTCategoryPropertyInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTCategoryPropertyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTCategoryPropertyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTCategoryPropertyInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTCategoryPropertyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTCategoryPropertyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTCategoryPropertyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTCategoryPropertyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTCategoryPropertyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectDefName

`func (o *BTCategoryPropertyInfo) GetObjectDefName() string`

GetObjectDefName returns the ObjectDefName field if non-nil, zero value otherwise.

### GetObjectDefNameOk

`func (o *BTCategoryPropertyInfo) GetObjectDefNameOk() (*string, bool)`

GetObjectDefNameOk returns a tuple with the ObjectDefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDefName

`func (o *BTCategoryPropertyInfo) SetObjectDefName(v string)`

SetObjectDefName sets ObjectDefName field to given value.

### HasObjectDefName

`func (o *BTCategoryPropertyInfo) HasObjectDefName() bool`

HasObjectDefName returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTCategoryPropertyInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTCategoryPropertyInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTCategoryPropertyInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTCategoryPropertyInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTCategoryPropertyInfo) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTCategoryPropertyInfo) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTCategoryPropertyInfo) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTCategoryPropertyInfo) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetUiReadonlyInMicroversion

`func (o *BTCategoryPropertyInfo) GetUiReadonlyInMicroversion() bool`

GetUiReadonlyInMicroversion returns the UiReadonlyInMicroversion field if non-nil, zero value otherwise.

### GetUiReadonlyInMicroversionOk

`func (o *BTCategoryPropertyInfo) GetUiReadonlyInMicroversionOk() (*bool, bool)`

GetUiReadonlyInMicroversionOk returns a tuple with the UiReadonlyInMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiReadonlyInMicroversion

`func (o *BTCategoryPropertyInfo) SetUiReadonlyInMicroversion(v bool)`

SetUiReadonlyInMicroversion sets UiReadonlyInMicroversion field to given value.

### HasUiReadonlyInMicroversion

`func (o *BTCategoryPropertyInfo) HasUiReadonlyInMicroversion() bool`

HasUiReadonlyInMicroversion returns a boolean if a field has been set.

### GetUiReadonlyInVersion

`func (o *BTCategoryPropertyInfo) GetUiReadonlyInVersion() bool`

GetUiReadonlyInVersion returns the UiReadonlyInVersion field if non-nil, zero value otherwise.

### GetUiReadonlyInVersionOk

`func (o *BTCategoryPropertyInfo) GetUiReadonlyInVersionOk() (*bool, bool)`

GetUiReadonlyInVersionOk returns a tuple with the UiReadonlyInVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiReadonlyInVersion

`func (o *BTCategoryPropertyInfo) SetUiReadonlyInVersion(v bool)`

SetUiReadonlyInVersion sets UiReadonlyInVersion field to given value.

### HasUiReadonlyInVersion

`func (o *BTCategoryPropertyInfo) HasUiReadonlyInVersion() bool`

HasUiReadonlyInVersion returns a boolean if a field has been set.

### GetValueType

`func (o *BTCategoryPropertyInfo) GetValueType() int32`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTCategoryPropertyInfo) GetValueTypeOk() (*int32, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTCategoryPropertyInfo) SetValueType(v int32)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BTCategoryPropertyInfo) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetViewRef

`func (o *BTCategoryPropertyInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTCategoryPropertyInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTCategoryPropertyInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTCategoryPropertyInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


