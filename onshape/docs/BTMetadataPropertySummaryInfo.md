# BTMetadataPropertySummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EditableInMicroversion** | Pointer to **bool** |  | [optional] 
**EditableInVersion** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**PropertyConfigSummaryInfoList** | Pointer to [**[]BTMetadataPropertyConfigSummaryInfo**](BTMetadataPropertyConfigSummaryInfo.md) |  | [optional] 
**UiReadonlyInMicroversion** | Pointer to **bool** |  | [optional] 
**UiReadonlyInVersion** | Pointer to **bool** |  | [optional] 
**ValueType** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMetadataPropertySummaryInfo

`func NewBTMetadataPropertySummaryInfo() *BTMetadataPropertySummaryInfo`

NewBTMetadataPropertySummaryInfo instantiates a new BTMetadataPropertySummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataPropertySummaryInfoWithDefaults

`func NewBTMetadataPropertySummaryInfoWithDefaults() *BTMetadataPropertySummaryInfo`

NewBTMetadataPropertySummaryInfoWithDefaults instantiates a new BTMetadataPropertySummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTMetadataPropertySummaryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTMetadataPropertySummaryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTMetadataPropertySummaryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTMetadataPropertySummaryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditableInMicroversion

`func (o *BTMetadataPropertySummaryInfo) GetEditableInMicroversion() bool`

GetEditableInMicroversion returns the EditableInMicroversion field if non-nil, zero value otherwise.

### GetEditableInMicroversionOk

`func (o *BTMetadataPropertySummaryInfo) GetEditableInMicroversionOk() (*bool, bool)`

GetEditableInMicroversionOk returns a tuple with the EditableInMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInMicroversion

`func (o *BTMetadataPropertySummaryInfo) SetEditableInMicroversion(v bool)`

SetEditableInMicroversion sets EditableInMicroversion field to given value.

### HasEditableInMicroversion

`func (o *BTMetadataPropertySummaryInfo) HasEditableInMicroversion() bool`

HasEditableInMicroversion returns a boolean if a field has been set.

### GetEditableInVersion

`func (o *BTMetadataPropertySummaryInfo) GetEditableInVersion() bool`

GetEditableInVersion returns the EditableInVersion field if non-nil, zero value otherwise.

### GetEditableInVersionOk

`func (o *BTMetadataPropertySummaryInfo) GetEditableInVersionOk() (*bool, bool)`

GetEditableInVersionOk returns a tuple with the EditableInVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInVersion

`func (o *BTMetadataPropertySummaryInfo) SetEditableInVersion(v bool)`

SetEditableInVersion sets EditableInVersion field to given value.

### HasEditableInVersion

`func (o *BTMetadataPropertySummaryInfo) HasEditableInVersion() bool`

HasEditableInVersion returns a boolean if a field has been set.

### GetHref

`func (o *BTMetadataPropertySummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTMetadataPropertySummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTMetadataPropertySummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTMetadataPropertySummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTMetadataPropertySummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTMetadataPropertySummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTMetadataPropertySummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTMetadataPropertySummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataPropertySummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataPropertySummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataPropertySummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataPropertySummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTMetadataPropertySummaryInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTMetadataPropertySummaryInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTMetadataPropertySummaryInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTMetadataPropertySummaryInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTMetadataPropertySummaryInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTMetadataPropertySummaryInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTMetadataPropertySummaryInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTMetadataPropertySummaryInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTMetadataPropertySummaryInfo) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTMetadataPropertySummaryInfo) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTMetadataPropertySummaryInfo) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTMetadataPropertySummaryInfo) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPropertyConfigSummaryInfoList

`func (o *BTMetadataPropertySummaryInfo) GetPropertyConfigSummaryInfoList() []BTMetadataPropertyConfigSummaryInfo`

GetPropertyConfigSummaryInfoList returns the PropertyConfigSummaryInfoList field if non-nil, zero value otherwise.

### GetPropertyConfigSummaryInfoListOk

`func (o *BTMetadataPropertySummaryInfo) GetPropertyConfigSummaryInfoListOk() (*[]BTMetadataPropertyConfigSummaryInfo, bool)`

GetPropertyConfigSummaryInfoListOk returns a tuple with the PropertyConfigSummaryInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyConfigSummaryInfoList

`func (o *BTMetadataPropertySummaryInfo) SetPropertyConfigSummaryInfoList(v []BTMetadataPropertyConfigSummaryInfo)`

SetPropertyConfigSummaryInfoList sets PropertyConfigSummaryInfoList field to given value.

### HasPropertyConfigSummaryInfoList

`func (o *BTMetadataPropertySummaryInfo) HasPropertyConfigSummaryInfoList() bool`

HasPropertyConfigSummaryInfoList returns a boolean if a field has been set.

### GetUiReadonlyInMicroversion

`func (o *BTMetadataPropertySummaryInfo) GetUiReadonlyInMicroversion() bool`

GetUiReadonlyInMicroversion returns the UiReadonlyInMicroversion field if non-nil, zero value otherwise.

### GetUiReadonlyInMicroversionOk

`func (o *BTMetadataPropertySummaryInfo) GetUiReadonlyInMicroversionOk() (*bool, bool)`

GetUiReadonlyInMicroversionOk returns a tuple with the UiReadonlyInMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiReadonlyInMicroversion

`func (o *BTMetadataPropertySummaryInfo) SetUiReadonlyInMicroversion(v bool)`

SetUiReadonlyInMicroversion sets UiReadonlyInMicroversion field to given value.

### HasUiReadonlyInMicroversion

`func (o *BTMetadataPropertySummaryInfo) HasUiReadonlyInMicroversion() bool`

HasUiReadonlyInMicroversion returns a boolean if a field has been set.

### GetUiReadonlyInVersion

`func (o *BTMetadataPropertySummaryInfo) GetUiReadonlyInVersion() bool`

GetUiReadonlyInVersion returns the UiReadonlyInVersion field if non-nil, zero value otherwise.

### GetUiReadonlyInVersionOk

`func (o *BTMetadataPropertySummaryInfo) GetUiReadonlyInVersionOk() (*bool, bool)`

GetUiReadonlyInVersionOk returns a tuple with the UiReadonlyInVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiReadonlyInVersion

`func (o *BTMetadataPropertySummaryInfo) SetUiReadonlyInVersion(v bool)`

SetUiReadonlyInVersion sets UiReadonlyInVersion field to given value.

### HasUiReadonlyInVersion

`func (o *BTMetadataPropertySummaryInfo) HasUiReadonlyInVersion() bool`

HasUiReadonlyInVersion returns a boolean if a field has been set.

### GetValueType

`func (o *BTMetadataPropertySummaryInfo) GetValueType() int32`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTMetadataPropertySummaryInfo) GetValueTypeOk() (*int32, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTMetadataPropertySummaryInfo) SetValueType(v int32)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BTMetadataPropertySummaryInfo) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetViewRef

`func (o *BTMetadataPropertySummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTMetadataPropertySummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTMetadataPropertySummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTMetadataPropertySummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


