# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBalance** | Pointer to **int64** |  | [optional] 
**BusinessVatId** | Pointer to **string** |  | [optional] 
**Cards** | Pointer to [**CustomerCardCollection**](CustomerCardCollection.md) |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DefaultCard** | Pointer to **string** |  | [optional] 
**DefaultSource** | Pointer to **string** |  | [optional] 
**DefaultSourceObject** | Pointer to [**ExternalAccount**](ExternalAccount.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Delinquent** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**Discount**](Discount.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**NextRecurringCharge** | Pointer to [**NextRecurringCharge**](NextRecurringCharge.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Shipping** | Pointer to [**ShippingDetails**](ShippingDetails.md) |  | [optional] 
**Sources** | Pointer to [**ExternalAccountCollection**](ExternalAccountCollection.md) |  | [optional] 
**Subscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**Subscriptions** | Pointer to [**CustomerSubscriptionCollection**](CustomerSubscriptionCollection.md) |  | [optional] 
**TrialEnd** | Pointer to **int64** |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBalance

`func (o *Customer) GetAccountBalance() int64`

GetAccountBalance returns the AccountBalance field if non-nil, zero value otherwise.

### GetAccountBalanceOk

`func (o *Customer) GetAccountBalanceOk() (*int64, bool)`

GetAccountBalanceOk returns a tuple with the AccountBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalance

`func (o *Customer) SetAccountBalance(v int64)`

SetAccountBalance sets AccountBalance field to given value.

### HasAccountBalance

`func (o *Customer) HasAccountBalance() bool`

HasAccountBalance returns a boolean if a field has been set.

### GetBusinessVatId

`func (o *Customer) GetBusinessVatId() string`

GetBusinessVatId returns the BusinessVatId field if non-nil, zero value otherwise.

### GetBusinessVatIdOk

`func (o *Customer) GetBusinessVatIdOk() (*string, bool)`

GetBusinessVatIdOk returns a tuple with the BusinessVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessVatId

`func (o *Customer) SetBusinessVatId(v string)`

SetBusinessVatId sets BusinessVatId field to given value.

### HasBusinessVatId

`func (o *Customer) HasBusinessVatId() bool`

HasBusinessVatId returns a boolean if a field has been set.

### GetCards

`func (o *Customer) GetCards() CustomerCardCollection`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *Customer) GetCardsOk() (*CustomerCardCollection, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *Customer) SetCards(v CustomerCardCollection)`

SetCards sets Cards field to given value.

### HasCards

`func (o *Customer) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetCreated

`func (o *Customer) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Customer) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Customer) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Customer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *Customer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Customer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Customer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Customer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultCard

`func (o *Customer) GetDefaultCard() string`

GetDefaultCard returns the DefaultCard field if non-nil, zero value otherwise.

### GetDefaultCardOk

`func (o *Customer) GetDefaultCardOk() (*string, bool)`

GetDefaultCardOk returns a tuple with the DefaultCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCard

`func (o *Customer) SetDefaultCard(v string)`

SetDefaultCard sets DefaultCard field to given value.

### HasDefaultCard

`func (o *Customer) HasDefaultCard() bool`

HasDefaultCard returns a boolean if a field has been set.

### GetDefaultSource

`func (o *Customer) GetDefaultSource() string`

GetDefaultSource returns the DefaultSource field if non-nil, zero value otherwise.

### GetDefaultSourceOk

`func (o *Customer) GetDefaultSourceOk() (*string, bool)`

GetDefaultSourceOk returns a tuple with the DefaultSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSource

`func (o *Customer) SetDefaultSource(v string)`

SetDefaultSource sets DefaultSource field to given value.

### HasDefaultSource

`func (o *Customer) HasDefaultSource() bool`

HasDefaultSource returns a boolean if a field has been set.

### GetDefaultSourceObject

`func (o *Customer) GetDefaultSourceObject() ExternalAccount`

GetDefaultSourceObject returns the DefaultSourceObject field if non-nil, zero value otherwise.

### GetDefaultSourceObjectOk

`func (o *Customer) GetDefaultSourceObjectOk() (*ExternalAccount, bool)`

GetDefaultSourceObjectOk returns a tuple with the DefaultSourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSourceObject

`func (o *Customer) SetDefaultSourceObject(v ExternalAccount)`

SetDefaultSourceObject sets DefaultSourceObject field to given value.

### HasDefaultSourceObject

`func (o *Customer) HasDefaultSourceObject() bool`

HasDefaultSourceObject returns a boolean if a field has been set.

### GetDeleted

`func (o *Customer) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Customer) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Customer) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Customer) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDelinquent

`func (o *Customer) GetDelinquent() bool`

GetDelinquent returns the Delinquent field if non-nil, zero value otherwise.

### GetDelinquentOk

`func (o *Customer) GetDelinquentOk() (*bool, bool)`

GetDelinquentOk returns a tuple with the Delinquent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelinquent

`func (o *Customer) SetDelinquent(v bool)`

SetDelinquent sets Delinquent field to given value.

### HasDelinquent

`func (o *Customer) HasDelinquent() bool`

HasDelinquent returns a boolean if a field has been set.

### GetDescription

`func (o *Customer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Customer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Customer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Customer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscount

`func (o *Customer) GetDiscount() Discount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Customer) GetDiscountOk() (*Discount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Customer) SetDiscount(v Discount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Customer) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Customer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *Customer) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Customer) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Customer) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Customer) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMetadata

`func (o *Customer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Customer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Customer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Customer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNextRecurringCharge

`func (o *Customer) GetNextRecurringCharge() NextRecurringCharge`

GetNextRecurringCharge returns the NextRecurringCharge field if non-nil, zero value otherwise.

### GetNextRecurringChargeOk

`func (o *Customer) GetNextRecurringChargeOk() (*NextRecurringCharge, bool)`

GetNextRecurringChargeOk returns a tuple with the NextRecurringCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecurringCharge

`func (o *Customer) SetNextRecurringCharge(v NextRecurringCharge)`

SetNextRecurringCharge sets NextRecurringCharge field to given value.

### HasNextRecurringCharge

`func (o *Customer) HasNextRecurringCharge() bool`

HasNextRecurringCharge returns a boolean if a field has been set.

### GetObject

`func (o *Customer) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Customer) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Customer) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Customer) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetShipping

`func (o *Customer) GetShipping() ShippingDetails`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *Customer) GetShippingOk() (*ShippingDetails, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *Customer) SetShipping(v ShippingDetails)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *Customer) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSources

`func (o *Customer) GetSources() ExternalAccountCollection`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Customer) GetSourcesOk() (*ExternalAccountCollection, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Customer) SetSources(v ExternalAccountCollection)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Customer) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSubscription

`func (o *Customer) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Customer) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Customer) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Customer) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptions

`func (o *Customer) GetSubscriptions() CustomerSubscriptionCollection`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *Customer) GetSubscriptionsOk() (*CustomerSubscriptionCollection, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *Customer) SetSubscriptions(v CustomerSubscriptionCollection)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *Customer) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTrialEnd

`func (o *Customer) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *Customer) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *Customer) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *Customer) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


