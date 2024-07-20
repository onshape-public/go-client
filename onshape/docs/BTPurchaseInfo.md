# BTPurchaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**ActualAmountPaidCents** | Pointer to **int64** |  | [optional] 
**AmountCents** | Pointer to **int64** |  | [optional] 
**Application** | Pointer to [**BTAPIApplicationSummaryInfo**](BTAPIApplicationSummaryInfo.md) |  | [optional] 
**CanceledAt** | Pointer to **JSONTime** |  | [optional] 
**Card** | Pointer to [**BTCardInfo**](BTCardInfo.md) |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**CouponAmountOff** | Pointer to **int64** |  | [optional] 
**CouponPercentOff** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**DurationMonths** | Pointer to **int32** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**LastModified** | Pointer to **JSONTime** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**LightSeats** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**NextCharge** | Pointer to [**NextCharge**](NextCharge.md) |  | [optional] 
**PaymentType** | Pointer to **int32** |  | [optional] 
**PendingCancelation** | Pointer to **bool** |  | [optional] 
**Plan** | Pointer to [**BTBillingPlanInfo**](BTBillingPlanInfo.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**PlanType** | Pointer to **int32** |  | [optional] 
**PreTrialPlanId** | Pointer to **string** |  | [optional] 
**ProratedCharges** | Pointer to [**[]ProratedCharges**](ProratedCharges.md) |  | [optional] 
**ProratedTotal** | Pointer to **int64** |  | [optional] 
**PurchaseDate** | Pointer to **JSONTime** |  | [optional] 
**ResellerName** | Pointer to **string** |  | [optional] 
**Seats** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**Subscribers** | Pointer to [**[]BTPlanSubscriberInfo**](BTPlanSubscriberInfo.md) |  | [optional] 
**SubscriptionBeginAt** | Pointer to **JSONTime** |  | [optional] 
**SubscriptionEndAt** | Pointer to **JSONTime** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SubscriptionType** | Pointer to **int32** |  | [optional] 
**TaxAmountCents** | Pointer to **int64** |  | [optional] 
**TrialEnd** | Pointer to **JSONTime** |  | [optional] 
**TrialInitiatedBy** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTPurchaseInfo

`func NewBTPurchaseInfo() *BTPurchaseInfo`

NewBTPurchaseInfo instantiates a new BTPurchaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPurchaseInfoWithDefaults

`func NewBTPurchaseInfoWithDefaults() *BTPurchaseInfo`

NewBTPurchaseInfoWithDefaults instantiates a new BTPurchaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BTPurchaseInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BTPurchaseInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BTPurchaseInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BTPurchaseInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActualAmountPaidCents

`func (o *BTPurchaseInfo) GetActualAmountPaidCents() int64`

GetActualAmountPaidCents returns the ActualAmountPaidCents field if non-nil, zero value otherwise.

### GetActualAmountPaidCentsOk

`func (o *BTPurchaseInfo) GetActualAmountPaidCentsOk() (*int64, bool)`

GetActualAmountPaidCentsOk returns a tuple with the ActualAmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualAmountPaidCents

`func (o *BTPurchaseInfo) SetActualAmountPaidCents(v int64)`

SetActualAmountPaidCents sets ActualAmountPaidCents field to given value.

### HasActualAmountPaidCents

`func (o *BTPurchaseInfo) HasActualAmountPaidCents() bool`

HasActualAmountPaidCents returns a boolean if a field has been set.

### GetAmountCents

`func (o *BTPurchaseInfo) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BTPurchaseInfo) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BTPurchaseInfo) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BTPurchaseInfo) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetApplication

`func (o *BTPurchaseInfo) GetApplication() BTAPIApplicationSummaryInfo`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *BTPurchaseInfo) GetApplicationOk() (*BTAPIApplicationSummaryInfo, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *BTPurchaseInfo) SetApplication(v BTAPIApplicationSummaryInfo)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *BTPurchaseInfo) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCanceledAt

`func (o *BTPurchaseInfo) GetCanceledAt() JSONTime`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *BTPurchaseInfo) GetCanceledAtOk() (*JSONTime, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *BTPurchaseInfo) SetCanceledAt(v JSONTime)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *BTPurchaseInfo) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCard

`func (o *BTPurchaseInfo) GetCard() BTCardInfo`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *BTPurchaseInfo) GetCardOk() (*BTCardInfo, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *BTPurchaseInfo) SetCard(v BTCardInfo)`

SetCard sets Card field to given value.

### HasCard

`func (o *BTPurchaseInfo) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetClientId

`func (o *BTPurchaseInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BTPurchaseInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BTPurchaseInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BTPurchaseInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCouponAmountOff

`func (o *BTPurchaseInfo) GetCouponAmountOff() int64`

GetCouponAmountOff returns the CouponAmountOff field if non-nil, zero value otherwise.

### GetCouponAmountOffOk

`func (o *BTPurchaseInfo) GetCouponAmountOffOk() (*int64, bool)`

GetCouponAmountOffOk returns a tuple with the CouponAmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAmountOff

`func (o *BTPurchaseInfo) SetCouponAmountOff(v int64)`

SetCouponAmountOff sets CouponAmountOff field to given value.

### HasCouponAmountOff

`func (o *BTPurchaseInfo) HasCouponAmountOff() bool`

HasCouponAmountOff returns a boolean if a field has been set.

### GetCouponPercentOff

`func (o *BTPurchaseInfo) GetCouponPercentOff() int32`

GetCouponPercentOff returns the CouponPercentOff field if non-nil, zero value otherwise.

### GetCouponPercentOffOk

`func (o *BTPurchaseInfo) GetCouponPercentOffOk() (*int32, bool)`

GetCouponPercentOffOk returns a tuple with the CouponPercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponPercentOff

`func (o *BTPurchaseInfo) SetCouponPercentOff(v int32)`

SetCouponPercentOff sets CouponPercentOff field to given value.

### HasCouponPercentOff

`func (o *BTPurchaseInfo) HasCouponPercentOff() bool`

HasCouponPercentOff returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTPurchaseInfo) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTPurchaseInfo) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTPurchaseInfo) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTPurchaseInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *BTPurchaseInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BTPurchaseInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BTPurchaseInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BTPurchaseInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuration

`func (o *BTPurchaseInfo) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BTPurchaseInfo) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BTPurchaseInfo) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BTPurchaseInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationMonths

`func (o *BTPurchaseInfo) GetDurationMonths() int32`

GetDurationMonths returns the DurationMonths field if non-nil, zero value otherwise.

### GetDurationMonthsOk

`func (o *BTPurchaseInfo) GetDurationMonthsOk() (*int32, bool)`

GetDurationMonthsOk returns a tuple with the DurationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMonths

`func (o *BTPurchaseInfo) SetDurationMonths(v int32)`

SetDurationMonths sets DurationMonths field to given value.

### HasDurationMonths

`func (o *BTPurchaseInfo) HasDurationMonths() bool`

HasDurationMonths returns a boolean if a field has been set.

### GetGroup

`func (o *BTPurchaseInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BTPurchaseInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BTPurchaseInfo) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BTPurchaseInfo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHref

`func (o *BTPurchaseInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTPurchaseInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTPurchaseInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTPurchaseInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTPurchaseInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPurchaseInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPurchaseInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPurchaseInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *BTPurchaseInfo) GetLastModified() JSONTime`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *BTPurchaseInfo) GetLastModifiedOk() (*JSONTime, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *BTPurchaseInfo) SetLastModified(v JSONTime)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *BTPurchaseInfo) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *BTPurchaseInfo) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BTPurchaseInfo) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BTPurchaseInfo) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *BTPurchaseInfo) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLightSeats

`func (o *BTPurchaseInfo) GetLightSeats() int64`

GetLightSeats returns the LightSeats field if non-nil, zero value otherwise.

### GetLightSeatsOk

`func (o *BTPurchaseInfo) GetLightSeatsOk() (*int64, bool)`

GetLightSeatsOk returns a tuple with the LightSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightSeats

`func (o *BTPurchaseInfo) SetLightSeats(v int64)`

SetLightSeats sets LightSeats field to given value.

### HasLightSeats

`func (o *BTPurchaseInfo) HasLightSeats() bool`

HasLightSeats returns a boolean if a field has been set.

### GetName

`func (o *BTPurchaseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPurchaseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPurchaseInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTPurchaseInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextCharge

`func (o *BTPurchaseInfo) GetNextCharge() NextCharge`

GetNextCharge returns the NextCharge field if non-nil, zero value otherwise.

### GetNextChargeOk

`func (o *BTPurchaseInfo) GetNextChargeOk() (*NextCharge, bool)`

GetNextChargeOk returns a tuple with the NextCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCharge

`func (o *BTPurchaseInfo) SetNextCharge(v NextCharge)`

SetNextCharge sets NextCharge field to given value.

### HasNextCharge

`func (o *BTPurchaseInfo) HasNextCharge() bool`

HasNextCharge returns a boolean if a field has been set.

### GetPaymentType

`func (o *BTPurchaseInfo) GetPaymentType() int32`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *BTPurchaseInfo) GetPaymentTypeOk() (*int32, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *BTPurchaseInfo) SetPaymentType(v int32)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *BTPurchaseInfo) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPendingCancelation

`func (o *BTPurchaseInfo) GetPendingCancelation() bool`

GetPendingCancelation returns the PendingCancelation field if non-nil, zero value otherwise.

### GetPendingCancelationOk

`func (o *BTPurchaseInfo) GetPendingCancelationOk() (*bool, bool)`

GetPendingCancelationOk returns a tuple with the PendingCancelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCancelation

`func (o *BTPurchaseInfo) SetPendingCancelation(v bool)`

SetPendingCancelation sets PendingCancelation field to given value.

### HasPendingCancelation

`func (o *BTPurchaseInfo) HasPendingCancelation() bool`

HasPendingCancelation returns a boolean if a field has been set.

### GetPlan

`func (o *BTPurchaseInfo) GetPlan() BTBillingPlanInfo`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BTPurchaseInfo) GetPlanOk() (*BTBillingPlanInfo, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BTPurchaseInfo) SetPlan(v BTBillingPlanInfo)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BTPurchaseInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *BTPurchaseInfo) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BTPurchaseInfo) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BTPurchaseInfo) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BTPurchaseInfo) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *BTPurchaseInfo) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BTPurchaseInfo) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BTPurchaseInfo) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BTPurchaseInfo) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanType

`func (o *BTPurchaseInfo) GetPlanType() int32`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *BTPurchaseInfo) GetPlanTypeOk() (*int32, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *BTPurchaseInfo) SetPlanType(v int32)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *BTPurchaseInfo) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPreTrialPlanId

`func (o *BTPurchaseInfo) GetPreTrialPlanId() string`

GetPreTrialPlanId returns the PreTrialPlanId field if non-nil, zero value otherwise.

### GetPreTrialPlanIdOk

`func (o *BTPurchaseInfo) GetPreTrialPlanIdOk() (*string, bool)`

GetPreTrialPlanIdOk returns a tuple with the PreTrialPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreTrialPlanId

`func (o *BTPurchaseInfo) SetPreTrialPlanId(v string)`

SetPreTrialPlanId sets PreTrialPlanId field to given value.

### HasPreTrialPlanId

`func (o *BTPurchaseInfo) HasPreTrialPlanId() bool`

HasPreTrialPlanId returns a boolean if a field has been set.

### GetProratedCharges

`func (o *BTPurchaseInfo) GetProratedCharges() []ProratedCharges`

GetProratedCharges returns the ProratedCharges field if non-nil, zero value otherwise.

### GetProratedChargesOk

`func (o *BTPurchaseInfo) GetProratedChargesOk() (*[]ProratedCharges, bool)`

GetProratedChargesOk returns a tuple with the ProratedCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedCharges

`func (o *BTPurchaseInfo) SetProratedCharges(v []ProratedCharges)`

SetProratedCharges sets ProratedCharges field to given value.

### HasProratedCharges

`func (o *BTPurchaseInfo) HasProratedCharges() bool`

HasProratedCharges returns a boolean if a field has been set.

### GetProratedTotal

`func (o *BTPurchaseInfo) GetProratedTotal() int64`

GetProratedTotal returns the ProratedTotal field if non-nil, zero value otherwise.

### GetProratedTotalOk

`func (o *BTPurchaseInfo) GetProratedTotalOk() (*int64, bool)`

GetProratedTotalOk returns a tuple with the ProratedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedTotal

`func (o *BTPurchaseInfo) SetProratedTotal(v int64)`

SetProratedTotal sets ProratedTotal field to given value.

### HasProratedTotal

`func (o *BTPurchaseInfo) HasProratedTotal() bool`

HasProratedTotal returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *BTPurchaseInfo) GetPurchaseDate() JSONTime`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *BTPurchaseInfo) GetPurchaseDateOk() (*JSONTime, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *BTPurchaseInfo) SetPurchaseDate(v JSONTime)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *BTPurchaseInfo) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetResellerName

`func (o *BTPurchaseInfo) GetResellerName() string`

GetResellerName returns the ResellerName field if non-nil, zero value otherwise.

### GetResellerNameOk

`func (o *BTPurchaseInfo) GetResellerNameOk() (*string, bool)`

GetResellerNameOk returns a tuple with the ResellerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerName

`func (o *BTPurchaseInfo) SetResellerName(v string)`

SetResellerName sets ResellerName field to given value.

### HasResellerName

`func (o *BTPurchaseInfo) HasResellerName() bool`

HasResellerName returns a boolean if a field has been set.

### GetSeats

`func (o *BTPurchaseInfo) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *BTPurchaseInfo) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *BTPurchaseInfo) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *BTPurchaseInfo) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetState

`func (o *BTPurchaseInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTPurchaseInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTPurchaseInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTPurchaseInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscribers

`func (o *BTPurchaseInfo) GetSubscribers() []BTPlanSubscriberInfo`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *BTPurchaseInfo) GetSubscribersOk() (*[]BTPlanSubscriberInfo, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *BTPurchaseInfo) SetSubscribers(v []BTPlanSubscriberInfo)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *BTPurchaseInfo) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetSubscriptionBeginAt

`func (o *BTPurchaseInfo) GetSubscriptionBeginAt() JSONTime`

GetSubscriptionBeginAt returns the SubscriptionBeginAt field if non-nil, zero value otherwise.

### GetSubscriptionBeginAtOk

`func (o *BTPurchaseInfo) GetSubscriptionBeginAtOk() (*JSONTime, bool)`

GetSubscriptionBeginAtOk returns a tuple with the SubscriptionBeginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionBeginAt

`func (o *BTPurchaseInfo) SetSubscriptionBeginAt(v JSONTime)`

SetSubscriptionBeginAt sets SubscriptionBeginAt field to given value.

### HasSubscriptionBeginAt

`func (o *BTPurchaseInfo) HasSubscriptionBeginAt() bool`

HasSubscriptionBeginAt returns a boolean if a field has been set.

### GetSubscriptionEndAt

`func (o *BTPurchaseInfo) GetSubscriptionEndAt() JSONTime`

GetSubscriptionEndAt returns the SubscriptionEndAt field if non-nil, zero value otherwise.

### GetSubscriptionEndAtOk

`func (o *BTPurchaseInfo) GetSubscriptionEndAtOk() (*JSONTime, bool)`

GetSubscriptionEndAtOk returns a tuple with the SubscriptionEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionEndAt

`func (o *BTPurchaseInfo) SetSubscriptionEndAt(v JSONTime)`

SetSubscriptionEndAt sets SubscriptionEndAt field to given value.

### HasSubscriptionEndAt

`func (o *BTPurchaseInfo) HasSubscriptionEndAt() bool`

HasSubscriptionEndAt returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *BTPurchaseInfo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *BTPurchaseInfo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *BTPurchaseInfo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *BTPurchaseInfo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *BTPurchaseInfo) GetSubscriptionType() int32`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *BTPurchaseInfo) GetSubscriptionTypeOk() (*int32, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *BTPurchaseInfo) SetSubscriptionType(v int32)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *BTPurchaseInfo) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetTaxAmountCents

`func (o *BTPurchaseInfo) GetTaxAmountCents() int64`

GetTaxAmountCents returns the TaxAmountCents field if non-nil, zero value otherwise.

### GetTaxAmountCentsOk

`func (o *BTPurchaseInfo) GetTaxAmountCentsOk() (*int64, bool)`

GetTaxAmountCentsOk returns a tuple with the TaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountCents

`func (o *BTPurchaseInfo) SetTaxAmountCents(v int64)`

SetTaxAmountCents sets TaxAmountCents field to given value.

### HasTaxAmountCents

`func (o *BTPurchaseInfo) HasTaxAmountCents() bool`

HasTaxAmountCents returns a boolean if a field has been set.

### GetTrialEnd

`func (o *BTPurchaseInfo) GetTrialEnd() JSONTime`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *BTPurchaseInfo) GetTrialEndOk() (*JSONTime, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *BTPurchaseInfo) SetTrialEnd(v JSONTime)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *BTPurchaseInfo) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialInitiatedBy

`func (o *BTPurchaseInfo) GetTrialInitiatedBy() string`

GetTrialInitiatedBy returns the TrialInitiatedBy field if non-nil, zero value otherwise.

### GetTrialInitiatedByOk

`func (o *BTPurchaseInfo) GetTrialInitiatedByOk() (*string, bool)`

GetTrialInitiatedByOk returns a tuple with the TrialInitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialInitiatedBy

`func (o *BTPurchaseInfo) SetTrialInitiatedBy(v string)`

SetTrialInitiatedBy sets TrialInitiatedBy field to given value.

### HasTrialInitiatedBy

`func (o *BTPurchaseInfo) HasTrialInitiatedBy() bool`

HasTrialInitiatedBy returns a boolean if a field has been set.

### GetViewRef

`func (o *BTPurchaseInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTPurchaseInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTPurchaseInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTPurchaseInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


