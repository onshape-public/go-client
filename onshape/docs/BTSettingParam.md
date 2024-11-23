# BTSettingParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to [**BTUserAppSettingOperationType**](BTUserAppSettingOperationType.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBTSettingParam

`func NewBTSettingParam() *BTSettingParam`

NewBTSettingParam instantiates a new BTSettingParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSettingParamWithDefaults

`func NewBTSettingParamWithDefaults() *BTSettingParam`

NewBTSettingParamWithDefaults instantiates a new BTSettingParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *BTSettingParam) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *BTSettingParam) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *BTSettingParam) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *BTSettingParam) HasField() bool`

HasField returns a boolean if a field has been set.

### GetKey

`func (o *BTSettingParam) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BTSettingParam) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BTSettingParam) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BTSettingParam) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperation

`func (o *BTSettingParam) GetOperation() BTUserAppSettingOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BTSettingParam) GetOperationOk() (*BTUserAppSettingOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BTSettingParam) SetOperation(v BTUserAppSettingOperationType)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BTSettingParam) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *BTSettingParam) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTSettingParam) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTSettingParam) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BTSettingParam) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


