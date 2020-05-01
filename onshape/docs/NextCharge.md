# NextCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CurrentPeriodEnd** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewNextCharge

`func NewNextCharge() *NextCharge`

NewNextCharge instantiates a new NextCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextChargeWithDefaults

`func NewNextChargeWithDefaults() *NextCharge`

NewNextChargeWithDefaults instantiates a new NextCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *NextCharge) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NextCharge) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NextCharge) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *NextCharge) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *NextCharge) GetCurrentPeriodEnd() JSONTime`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *NextCharge) GetCurrentPeriodEndOk() (*JSONTime, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *NextCharge) SetCurrentPeriodEnd(v JSONTime)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *NextCharge) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetInterval

`func (o *NextCharge) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *NextCharge) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *NextCharge) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *NextCharge) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTotal

`func (o *NextCharge) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NextCharge) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NextCharge) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NextCharge) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


