# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFeePercent** | Pointer to **float64** |  | [optional] 
**Billing** | Pointer to **string** |  | [optional] 
**CancelAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**CanceledAt** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **int64** |  | [optional] 
**CurrentPeriodStart** | Pointer to **int64** |  | [optional] 
**Customer** | Pointer to **string** |  | [optional] 
**CustomerObject** | Pointer to [**Customer**](Customer.md) |  | [optional] 
**DaysUntilDue** | Pointer to **int32** |  | [optional] 
**Discount** | Pointer to [**Discount**](Discount.md) |  | [optional] 
**EndedAt** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscriptionItems** | Pointer to [**SubscriptionItemCollection**](SubscriptionItemCollection.md) |  | [optional] 
**TaxPercent** | Pointer to **float64** |  | [optional] 
**TrialEnd** | Pointer to **int64** |  | [optional] 
**TrialStart** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFeePercent

`func (o *Subscription) GetApplicationFeePercent() float64`

GetApplicationFeePercent returns the ApplicationFeePercent field if non-nil, zero value otherwise.

### GetApplicationFeePercentOk

`func (o *Subscription) GetApplicationFeePercentOk() (*float64, bool)`

GetApplicationFeePercentOk returns a tuple with the ApplicationFeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFeePercent

`func (o *Subscription) SetApplicationFeePercent(v float64)`

SetApplicationFeePercent sets ApplicationFeePercent field to given value.

### HasApplicationFeePercent

`func (o *Subscription) HasApplicationFeePercent() bool`

HasApplicationFeePercent returns a boolean if a field has been set.

### GetBilling

`func (o *Subscription) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Subscription) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Subscription) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Subscription) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *Subscription) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *Subscription) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *Subscription) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *Subscription) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *Subscription) GetCanceledAt() int64`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Subscription) GetCanceledAtOk() (*int64, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Subscription) SetCanceledAt(v int64)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Subscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCreated

`func (o *Subscription) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Subscription) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Subscription) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Subscription) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *Subscription) GetCurrentPeriodEnd() int64`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *Subscription) GetCurrentPeriodEndOk() (*int64, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *Subscription) SetCurrentPeriodEnd(v int64)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *Subscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *Subscription) GetCurrentPeriodStart() int64`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *Subscription) GetCurrentPeriodStartOk() (*int64, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *Subscription) SetCurrentPeriodStart(v int64)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *Subscription) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetCustomer

`func (o *Subscription) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Subscription) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Subscription) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Subscription) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerObject

`func (o *Subscription) GetCustomerObject() Customer`

GetCustomerObject returns the CustomerObject field if non-nil, zero value otherwise.

### GetCustomerObjectOk

`func (o *Subscription) GetCustomerObjectOk() (*Customer, bool)`

GetCustomerObjectOk returns a tuple with the CustomerObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerObject

`func (o *Subscription) SetCustomerObject(v Customer)`

SetCustomerObject sets CustomerObject field to given value.

### HasCustomerObject

`func (o *Subscription) HasCustomerObject() bool`

HasCustomerObject returns a boolean if a field has been set.

### GetDaysUntilDue

`func (o *Subscription) GetDaysUntilDue() int32`

GetDaysUntilDue returns the DaysUntilDue field if non-nil, zero value otherwise.

### GetDaysUntilDueOk

`func (o *Subscription) GetDaysUntilDueOk() (*int32, bool)`

GetDaysUntilDueOk returns a tuple with the DaysUntilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilDue

`func (o *Subscription) SetDaysUntilDue(v int32)`

SetDaysUntilDue sets DaysUntilDue field to given value.

### HasDaysUntilDue

`func (o *Subscription) HasDaysUntilDue() bool`

HasDaysUntilDue returns a boolean if a field has been set.

### GetDiscount

`func (o *Subscription) GetDiscount() Discount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Subscription) GetDiscountOk() (*Discount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Subscription) SetDiscount(v Discount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Subscription) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetEndedAt

`func (o *Subscription) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Subscription) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Subscription) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Subscription) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *Subscription) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Subscription) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Subscription) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Subscription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObject

`func (o *Subscription) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Subscription) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Subscription) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Subscription) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPlan

`func (o *Subscription) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Subscription) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Subscription) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Subscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *Subscription) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Subscription) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Subscription) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Subscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStart

`func (o *Subscription) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Subscription) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Subscription) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *Subscription) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionItems

`func (o *Subscription) GetSubscriptionItems() SubscriptionItemCollection`

GetSubscriptionItems returns the SubscriptionItems field if non-nil, zero value otherwise.

### GetSubscriptionItemsOk

`func (o *Subscription) GetSubscriptionItemsOk() (*SubscriptionItemCollection, bool)`

GetSubscriptionItemsOk returns a tuple with the SubscriptionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionItems

`func (o *Subscription) SetSubscriptionItems(v SubscriptionItemCollection)`

SetSubscriptionItems sets SubscriptionItems field to given value.

### HasSubscriptionItems

`func (o *Subscription) HasSubscriptionItems() bool`

HasSubscriptionItems returns a boolean if a field has been set.

### GetTaxPercent

`func (o *Subscription) GetTaxPercent() float64`

GetTaxPercent returns the TaxPercent field if non-nil, zero value otherwise.

### GetTaxPercentOk

`func (o *Subscription) GetTaxPercentOk() (*float64, bool)`

GetTaxPercentOk returns a tuple with the TaxPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercent

`func (o *Subscription) SetTaxPercent(v float64)`

SetTaxPercent sets TaxPercent field to given value.

### HasTaxPercent

`func (o *Subscription) HasTaxPercent() bool`

HasTaxPercent returns a boolean if a field has been set.

### GetTrialEnd

`func (o *Subscription) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *Subscription) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *Subscription) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *Subscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *Subscription) GetTrialStart() int64`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *Subscription) GetTrialStartOk() (*int64, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *Subscription) SetTrialStart(v int64)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *Subscription) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


