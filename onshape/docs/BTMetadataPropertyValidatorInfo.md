# BTMetadataPropertyValidatorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **float32** |  | [optional] 
**MaxCount** | Pointer to **int32** |  | [optional] 
**MaxDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**Min** | Pointer to **float32** |  | [optional] 
**MinCount** | Pointer to **int32** |  | [optional] 
**MinDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMetadataPropertyValidatorInfo

`func NewBTMetadataPropertyValidatorInfo() *BTMetadataPropertyValidatorInfo`

NewBTMetadataPropertyValidatorInfo instantiates a new BTMetadataPropertyValidatorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataPropertyValidatorInfoWithDefaults

`func NewBTMetadataPropertyValidatorInfoWithDefaults() *BTMetadataPropertyValidatorInfo`

NewBTMetadataPropertyValidatorInfoWithDefaults instantiates a new BTMetadataPropertyValidatorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *BTMetadataPropertyValidatorInfo) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *BTMetadataPropertyValidatorInfo) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *BTMetadataPropertyValidatorInfo) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *BTMetadataPropertyValidatorInfo) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxCount

`func (o *BTMetadataPropertyValidatorInfo) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *BTMetadataPropertyValidatorInfo) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *BTMetadataPropertyValidatorInfo) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *BTMetadataPropertyValidatorInfo) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetMaxDate

`func (o *BTMetadataPropertyValidatorInfo) GetMaxDate() JSONTime`

GetMaxDate returns the MaxDate field if non-nil, zero value otherwise.

### GetMaxDateOk

`func (o *BTMetadataPropertyValidatorInfo) GetMaxDateOk() (*JSONTime, bool)`

GetMaxDateOk returns a tuple with the MaxDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDate

`func (o *BTMetadataPropertyValidatorInfo) SetMaxDate(v JSONTime)`

SetMaxDate sets MaxDate field to given value.

### HasMaxDate

`func (o *BTMetadataPropertyValidatorInfo) HasMaxDate() bool`

HasMaxDate returns a boolean if a field has been set.

### GetMaxLength

`func (o *BTMetadataPropertyValidatorInfo) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *BTMetadataPropertyValidatorInfo) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *BTMetadataPropertyValidatorInfo) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *BTMetadataPropertyValidatorInfo) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMin

`func (o *BTMetadataPropertyValidatorInfo) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *BTMetadataPropertyValidatorInfo) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *BTMetadataPropertyValidatorInfo) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *BTMetadataPropertyValidatorInfo) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMinCount

`func (o *BTMetadataPropertyValidatorInfo) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *BTMetadataPropertyValidatorInfo) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *BTMetadataPropertyValidatorInfo) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *BTMetadataPropertyValidatorInfo) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMinDate

`func (o *BTMetadataPropertyValidatorInfo) GetMinDate() JSONTime`

GetMinDate returns the MinDate field if non-nil, zero value otherwise.

### GetMinDateOk

`func (o *BTMetadataPropertyValidatorInfo) GetMinDateOk() (*JSONTime, bool)`

GetMinDateOk returns a tuple with the MinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDate

`func (o *BTMetadataPropertyValidatorInfo) SetMinDate(v JSONTime)`

SetMinDate sets MinDate field to given value.

### HasMinDate

`func (o *BTMetadataPropertyValidatorInfo) HasMinDate() bool`

HasMinDate returns a boolean if a field has been set.

### GetMinLength

`func (o *BTMetadataPropertyValidatorInfo) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *BTMetadataPropertyValidatorInfo) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *BTMetadataPropertyValidatorInfo) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *BTMetadataPropertyValidatorInfo) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetPattern

`func (o *BTMetadataPropertyValidatorInfo) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *BTMetadataPropertyValidatorInfo) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *BTMetadataPropertyValidatorInfo) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *BTMetadataPropertyValidatorInfo) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


