# ProratedCharges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**Date** | Pointer to **JSONTime** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewProratedCharges

`func NewProratedCharges() *ProratedCharges`

NewProratedCharges instantiates a new ProratedCharges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProratedChargesWithDefaults

`func NewProratedChargesWithDefaults() *ProratedCharges`

NewProratedChargesWithDefaults instantiates a new ProratedCharges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ProratedCharges) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProratedCharges) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProratedCharges) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ProratedCharges) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *ProratedCharges) GetDate() JSONTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProratedCharges) GetDateOk() (*JSONTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProratedCharges) SetDate(v JSONTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *ProratedCharges) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *ProratedCharges) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProratedCharges) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProratedCharges) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProratedCharges) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


