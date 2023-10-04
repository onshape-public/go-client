# BTDiscountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBalance** | Pointer to **int32** |  | [optional] 
**AmountOff** | Pointer to **int32** |  | [optional] 
**CouponType** | Pointer to **int32** |  | [optional] 
**CouponValidMonths** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**ExpiresAt** | Pointer to **JSONTime** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PercentOff** | Pointer to **int32** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**TrialEndDate** | Pointer to **string** |  | [optional] 
**UsedAt** | Pointer to **JSONTime** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTDiscountInfo

`func NewBTDiscountInfo() *BTDiscountInfo`

NewBTDiscountInfo instantiates a new BTDiscountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDiscountInfoWithDefaults

`func NewBTDiscountInfoWithDefaults() *BTDiscountInfo`

NewBTDiscountInfoWithDefaults instantiates a new BTDiscountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBalance

`func (o *BTDiscountInfo) GetAccountBalance() int32`

GetAccountBalance returns the AccountBalance field if non-nil, zero value otherwise.

### GetAccountBalanceOk

`func (o *BTDiscountInfo) GetAccountBalanceOk() (*int32, bool)`

GetAccountBalanceOk returns a tuple with the AccountBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalance

`func (o *BTDiscountInfo) SetAccountBalance(v int32)`

SetAccountBalance sets AccountBalance field to given value.

### HasAccountBalance

`func (o *BTDiscountInfo) HasAccountBalance() bool`

HasAccountBalance returns a boolean if a field has been set.

### GetAmountOff

`func (o *BTDiscountInfo) GetAmountOff() int32`

GetAmountOff returns the AmountOff field if non-nil, zero value otherwise.

### GetAmountOffOk

`func (o *BTDiscountInfo) GetAmountOffOk() (*int32, bool)`

GetAmountOffOk returns a tuple with the AmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOff

`func (o *BTDiscountInfo) SetAmountOff(v int32)`

SetAmountOff sets AmountOff field to given value.

### HasAmountOff

`func (o *BTDiscountInfo) HasAmountOff() bool`

HasAmountOff returns a boolean if a field has been set.

### GetCouponType

`func (o *BTDiscountInfo) GetCouponType() int32`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *BTDiscountInfo) GetCouponTypeOk() (*int32, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *BTDiscountInfo) SetCouponType(v int32)`

SetCouponType sets CouponType field to given value.

### HasCouponType

`func (o *BTDiscountInfo) HasCouponType() bool`

HasCouponType returns a boolean if a field has been set.

### GetCouponValidMonths

`func (o *BTDiscountInfo) GetCouponValidMonths() int32`

GetCouponValidMonths returns the CouponValidMonths field if non-nil, zero value otherwise.

### GetCouponValidMonthsOk

`func (o *BTDiscountInfo) GetCouponValidMonthsOk() (*int32, bool)`

GetCouponValidMonthsOk returns a tuple with the CouponValidMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValidMonths

`func (o *BTDiscountInfo) SetCouponValidMonths(v int32)`

SetCouponValidMonths sets CouponValidMonths field to given value.

### HasCouponValidMonths

`func (o *BTDiscountInfo) HasCouponValidMonths() bool`

HasCouponValidMonths returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTDiscountInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTDiscountInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTDiscountInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTDiscountInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTDiscountInfo) GetCreatedBy() BTUserSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDiscountInfo) GetCreatedByOk() (*BTUserSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDiscountInfo) SetCreatedBy(v BTUserSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDiscountInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BTDiscountInfo) GetExpiresAt() JSONTime`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BTDiscountInfo) GetExpiresAtOk() (*JSONTime, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BTDiscountInfo) SetExpiresAt(v JSONTime)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BTDiscountInfo) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHref

`func (o *BTDiscountInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTDiscountInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTDiscountInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTDiscountInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTDiscountInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDiscountInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDiscountInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDiscountInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTDiscountInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDiscountInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDiscountInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDiscountInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTDiscountInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTDiscountInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTDiscountInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTDiscountInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPercentOff

`func (o *BTDiscountInfo) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *BTDiscountInfo) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *BTDiscountInfo) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *BTDiscountInfo) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetPlanId

`func (o *BTDiscountInfo) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BTDiscountInfo) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BTDiscountInfo) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BTDiscountInfo) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *BTDiscountInfo) GetTrialEndDate() string`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *BTDiscountInfo) GetTrialEndDateOk() (*string, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *BTDiscountInfo) SetTrialEndDate(v string)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *BTDiscountInfo) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUsedAt

`func (o *BTDiscountInfo) GetUsedAt() JSONTime`

GetUsedAt returns the UsedAt field if non-nil, zero value otherwise.

### GetUsedAtOk

`func (o *BTDiscountInfo) GetUsedAtOk() (*JSONTime, bool)`

GetUsedAtOk returns a tuple with the UsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAt

`func (o *BTDiscountInfo) SetUsedAt(v JSONTime)`

SetUsedAt sets UsedAt field to given value.

### HasUsedAt

`func (o *BTDiscountInfo) HasUsedAt() bool`

HasUsedAt returns a boolean if a field has been set.

### GetViewRef

`func (o *BTDiscountInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTDiscountInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTDiscountInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTDiscountInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


