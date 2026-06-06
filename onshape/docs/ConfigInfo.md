# ConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayValue** | Pointer to **string** | The formatted display value of the configuration parameter. | [optional] 
**DisplayValueAbbrUnit** | Pointer to **string** | The abbreviated display value with unit for the configuration parameter. | [optional] 
**Id** | Pointer to **string** | The configuration parameter ID. | [optional] 
**Name** | Pointer to **string** | The configuration parameter name. | [optional] 
**Type** | Pointer to **int32** | The configuration parameter type. &#x60;0: ENUM | 1: BOOLEAN | 2: STRING | 3: QUANTITY&#x60; | [optional] 
**Value** | Pointer to **string** | The raw value of the configuration parameter. | [optional] 

## Methods

### NewConfigInfo

`func NewConfigInfo() *ConfigInfo`

NewConfigInfo instantiates a new ConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigInfoWithDefaults

`func NewConfigInfoWithDefaults() *ConfigInfo`

NewConfigInfoWithDefaults instantiates a new ConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayValue

`func (o *ConfigInfo) GetDisplayValue() string`

GetDisplayValue returns the DisplayValue field if non-nil, zero value otherwise.

### GetDisplayValueOk

`func (o *ConfigInfo) GetDisplayValueOk() (*string, bool)`

GetDisplayValueOk returns a tuple with the DisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValue

`func (o *ConfigInfo) SetDisplayValue(v string)`

SetDisplayValue sets DisplayValue field to given value.

### HasDisplayValue

`func (o *ConfigInfo) HasDisplayValue() bool`

HasDisplayValue returns a boolean if a field has been set.

### GetDisplayValueAbbrUnit

`func (o *ConfigInfo) GetDisplayValueAbbrUnit() string`

GetDisplayValueAbbrUnit returns the DisplayValueAbbrUnit field if non-nil, zero value otherwise.

### GetDisplayValueAbbrUnitOk

`func (o *ConfigInfo) GetDisplayValueAbbrUnitOk() (*string, bool)`

GetDisplayValueAbbrUnitOk returns a tuple with the DisplayValueAbbrUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueAbbrUnit

`func (o *ConfigInfo) SetDisplayValueAbbrUnit(v string)`

SetDisplayValueAbbrUnit sets DisplayValueAbbrUnit field to given value.

### HasDisplayValueAbbrUnit

`func (o *ConfigInfo) HasDisplayValueAbbrUnit() bool`

HasDisplayValueAbbrUnit returns a boolean if a field has been set.

### GetId

`func (o *ConfigInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConfigInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ConfigInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ConfigInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


