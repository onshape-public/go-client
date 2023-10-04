# BTDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBalance** | Pointer to **int32** |  | [optional] 
**AmountOff** | Pointer to **int32** |  | [optional] 
**AmountOffCurrency** | Pointer to **string** |  | [optional] 
**CouponType** | Pointer to **int32** |  | [optional] 
**CouponValidMonths** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **JSONTime** |  | [optional] 
**Id** | Pointer to [**BTDiscountOwnerIdPlanId**](BTDiscountOwnerIdPlanId.md) |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**New** | Pointer to **bool** |  | [optional] 
**PercentOff** | Pointer to **int32** |  | [optional] 
**TrialEndDate** | Pointer to **string** |  | [optional] 
**UsedAt** | Pointer to **JSONTime** |  | [optional] 

## Methods

### NewBTDiscount

`func NewBTDiscount() *BTDiscount`

NewBTDiscount instantiates a new BTDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDiscountWithDefaults

`func NewBTDiscountWithDefaults() *BTDiscount`

NewBTDiscountWithDefaults instantiates a new BTDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBalance

`func (o *BTDiscount) GetAccountBalance() int32`

GetAccountBalance returns the AccountBalance field if non-nil, zero value otherwise.

### GetAccountBalanceOk

`func (o *BTDiscount) GetAccountBalanceOk() (*int32, bool)`

GetAccountBalanceOk returns a tuple with the AccountBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalance

`func (o *BTDiscount) SetAccountBalance(v int32)`

SetAccountBalance sets AccountBalance field to given value.

### HasAccountBalance

`func (o *BTDiscount) HasAccountBalance() bool`

HasAccountBalance returns a boolean if a field has been set.

### GetAmountOff

`func (o *BTDiscount) GetAmountOff() int32`

GetAmountOff returns the AmountOff field if non-nil, zero value otherwise.

### GetAmountOffOk

`func (o *BTDiscount) GetAmountOffOk() (*int32, bool)`

GetAmountOffOk returns a tuple with the AmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOff

`func (o *BTDiscount) SetAmountOff(v int32)`

SetAmountOff sets AmountOff field to given value.

### HasAmountOff

`func (o *BTDiscount) HasAmountOff() bool`

HasAmountOff returns a boolean if a field has been set.

### GetAmountOffCurrency

`func (o *BTDiscount) GetAmountOffCurrency() string`

GetAmountOffCurrency returns the AmountOffCurrency field if non-nil, zero value otherwise.

### GetAmountOffCurrencyOk

`func (o *BTDiscount) GetAmountOffCurrencyOk() (*string, bool)`

GetAmountOffCurrencyOk returns a tuple with the AmountOffCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOffCurrency

`func (o *BTDiscount) SetAmountOffCurrency(v string)`

SetAmountOffCurrency sets AmountOffCurrency field to given value.

### HasAmountOffCurrency

`func (o *BTDiscount) HasAmountOffCurrency() bool`

HasAmountOffCurrency returns a boolean if a field has been set.

### GetCouponType

`func (o *BTDiscount) GetCouponType() int32`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *BTDiscount) GetCouponTypeOk() (*int32, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *BTDiscount) SetCouponType(v int32)`

SetCouponType sets CouponType field to given value.

### HasCouponType

`func (o *BTDiscount) HasCouponType() bool`

HasCouponType returns a boolean if a field has been set.

### GetCouponValidMonths

`func (o *BTDiscount) GetCouponValidMonths() int32`

GetCouponValidMonths returns the CouponValidMonths field if non-nil, zero value otherwise.

### GetCouponValidMonthsOk

`func (o *BTDiscount) GetCouponValidMonthsOk() (*int32, bool)`

GetCouponValidMonthsOk returns a tuple with the CouponValidMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValidMonths

`func (o *BTDiscount) SetCouponValidMonths(v int32)`

SetCouponValidMonths sets CouponValidMonths field to given value.

### HasCouponValidMonths

`func (o *BTDiscount) HasCouponValidMonths() bool`

HasCouponValidMonths returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTDiscount) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTDiscount) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTDiscount) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTDiscount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTDiscount) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDiscount) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDiscount) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDiscount) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTDiscount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDiscount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDiscount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDiscount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BTDiscount) GetExpiresAt() JSONTime`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BTDiscount) GetExpiresAtOk() (*JSONTime, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BTDiscount) SetExpiresAt(v JSONTime)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BTDiscount) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *BTDiscount) GetId() BTDiscountOwnerIdPlanId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDiscount) GetIdOk() (*BTDiscountOwnerIdPlanId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDiscount) SetId(v BTDiscountOwnerIdPlanId)`

SetId sets Id field to given value.

### HasId

`func (o *BTDiscount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *BTDiscount) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BTDiscount) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BTDiscount) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BTDiscount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTDiscount) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTDiscount) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTDiscount) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTDiscount) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTDiscount) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTDiscount) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTDiscount) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTDiscount) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTDiscount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDiscount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDiscount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDiscount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNew

`func (o *BTDiscount) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *BTDiscount) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *BTDiscount) SetNew(v bool)`

SetNew sets New field to given value.

### HasNew

`func (o *BTDiscount) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetPercentOff

`func (o *BTDiscount) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *BTDiscount) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *BTDiscount) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *BTDiscount) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *BTDiscount) GetTrialEndDate() string`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *BTDiscount) GetTrialEndDateOk() (*string, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *BTDiscount) SetTrialEndDate(v string)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *BTDiscount) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUsedAt

`func (o *BTDiscount) GetUsedAt() JSONTime`

GetUsedAt returns the UsedAt field if non-nil, zero value otherwise.

### GetUsedAtOk

`func (o *BTDiscount) GetUsedAtOk() (*JSONTime, bool)`

GetUsedAtOk returns a tuple with the UsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAt

`func (o *BTDiscount) SetUsedAt(v JSONTime)`

SetUsedAt sets UsedAt field to given value.

### HasUsedAt

`func (o *BTDiscount) HasUsedAt() bool`

HasUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


