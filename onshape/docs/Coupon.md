# Coupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOff** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**DurationInMonths** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**MaxRedemptions** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**PercentOff** | Pointer to **int32** |  | [optional] 
**RedeemBy** | Pointer to **int64** |  | [optional] 
**TimesRedeemed** | Pointer to **int32** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewCoupon

`func NewCoupon() *Coupon`

NewCoupon instantiates a new Coupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponWithDefaults

`func NewCouponWithDefaults() *Coupon`

NewCouponWithDefaults instantiates a new Coupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOff

`func (o *Coupon) GetAmountOff() int64`

GetAmountOff returns the AmountOff field if non-nil, zero value otherwise.

### GetAmountOffOk

`func (o *Coupon) GetAmountOffOk() (*int64, bool)`

GetAmountOffOk returns a tuple with the AmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOff

`func (o *Coupon) SetAmountOff(v int64)`

SetAmountOff sets AmountOff field to given value.

### HasAmountOff

`func (o *Coupon) HasAmountOff() bool`

HasAmountOff returns a boolean if a field has been set.

### GetCreated

`func (o *Coupon) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Coupon) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Coupon) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Coupon) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *Coupon) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Coupon) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Coupon) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Coupon) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuration

`func (o *Coupon) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Coupon) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Coupon) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Coupon) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationInMonths

`func (o *Coupon) GetDurationInMonths() int32`

GetDurationInMonths returns the DurationInMonths field if non-nil, zero value otherwise.

### GetDurationInMonthsOk

`func (o *Coupon) GetDurationInMonthsOk() (*int32, bool)`

GetDurationInMonthsOk returns a tuple with the DurationInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMonths

`func (o *Coupon) SetDurationInMonths(v int32)`

SetDurationInMonths sets DurationInMonths field to given value.

### HasDurationInMonths

`func (o *Coupon) HasDurationInMonths() bool`

HasDurationInMonths returns a boolean if a field has been set.

### GetId

`func (o *Coupon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Coupon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Coupon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Coupon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *Coupon) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Coupon) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Coupon) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Coupon) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *Coupon) GetMaxRedemptions() int64`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *Coupon) GetMaxRedemptionsOk() (*int64, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *Coupon) SetMaxRedemptions(v int64)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *Coupon) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetMetadata

`func (o *Coupon) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Coupon) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Coupon) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Coupon) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObject

`func (o *Coupon) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Coupon) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Coupon) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Coupon) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPercentOff

`func (o *Coupon) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *Coupon) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *Coupon) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *Coupon) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetRedeemBy

`func (o *Coupon) GetRedeemBy() int64`

GetRedeemBy returns the RedeemBy field if non-nil, zero value otherwise.

### GetRedeemByOk

`func (o *Coupon) GetRedeemByOk() (*int64, bool)`

GetRedeemByOk returns a tuple with the RedeemBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBy

`func (o *Coupon) SetRedeemBy(v int64)`

SetRedeemBy sets RedeemBy field to given value.

### HasRedeemBy

`func (o *Coupon) HasRedeemBy() bool`

HasRedeemBy returns a boolean if a field has been set.

### GetTimesRedeemed

`func (o *Coupon) GetTimesRedeemed() int32`

GetTimesRedeemed returns the TimesRedeemed field if non-nil, zero value otherwise.

### GetTimesRedeemedOk

`func (o *Coupon) GetTimesRedeemedOk() (*int32, bool)`

GetTimesRedeemedOk returns a tuple with the TimesRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesRedeemed

`func (o *Coupon) SetTimesRedeemed(v int32)`

SetTimesRedeemed sets TimesRedeemed field to given value.

### HasTimesRedeemed

`func (o *Coupon) HasTimesRedeemed() bool`

HasTimesRedeemed returns a boolean if a field has been set.

### GetValid

`func (o *Coupon) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *Coupon) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *Coupon) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *Coupon) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


