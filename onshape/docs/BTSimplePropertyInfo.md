# BTSimplePropertyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Frozen** | Pointer to **bool** |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**PublishState** | Pointer to **int32** |  | [optional] 
**ValueType** | Pointer to [**BTMetadataValueType**](BTMetadataValueType.md) |  | [optional] 

## Methods

### NewBTSimplePropertyInfo

`func NewBTSimplePropertyInfo() *BTSimplePropertyInfo`

NewBTSimplePropertyInfo instantiates a new BTSimplePropertyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSimplePropertyInfoWithDefaults

`func NewBTSimplePropertyInfoWithDefaults() *BTSimplePropertyInfo`

NewBTSimplePropertyInfoWithDefaults instantiates a new BTSimplePropertyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *BTSimplePropertyInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BTSimplePropertyInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BTSimplePropertyInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BTSimplePropertyInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFrozen

`func (o *BTSimplePropertyInfo) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *BTSimplePropertyInfo) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *BTSimplePropertyInfo) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *BTSimplePropertyInfo) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTSimplePropertyInfo) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTSimplePropertyInfo) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTSimplePropertyInfo) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTSimplePropertyInfo) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPublishState

`func (o *BTSimplePropertyInfo) GetPublishState() int32`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *BTSimplePropertyInfo) GetPublishStateOk() (*int32, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *BTSimplePropertyInfo) SetPublishState(v int32)`

SetPublishState sets PublishState field to given value.

### HasPublishState

`func (o *BTSimplePropertyInfo) HasPublishState() bool`

HasPublishState returns a boolean if a field has been set.

### GetValueType

`func (o *BTSimplePropertyInfo) GetValueType() BTMetadataValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTSimplePropertyInfo) GetValueTypeOk() (*BTMetadataValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTSimplePropertyInfo) SetValueType(v BTMetadataValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BTSimplePropertyInfo) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


