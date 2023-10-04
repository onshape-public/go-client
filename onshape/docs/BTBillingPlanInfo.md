# BTBillingPlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**CompanyPlan** | Pointer to **bool** |  | [optional] 
**ConsumableQuantity** | Pointer to **int32** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiscountInfo** | Pointer to [**BTDiscountInfo**](BTDiscountInfo.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**OnshapePlan** | Pointer to **bool** |  | [optional] 
**PlanType** | Pointer to **int32** |  | [optional] 
**TrialPeriodDays** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTBillingPlanInfo

`func NewBTBillingPlanInfo() *BTBillingPlanInfo`

NewBTBillingPlanInfo instantiates a new BTBillingPlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBillingPlanInfoWithDefaults

`func NewBTBillingPlanInfoWithDefaults() *BTBillingPlanInfo`

NewBTBillingPlanInfoWithDefaults instantiates a new BTBillingPlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *BTBillingPlanInfo) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BTBillingPlanInfo) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BTBillingPlanInfo) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BTBillingPlanInfo) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetApplicationId

`func (o *BTBillingPlanInfo) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *BTBillingPlanInfo) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *BTBillingPlanInfo) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *BTBillingPlanInfo) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetClientId

`func (o *BTBillingPlanInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BTBillingPlanInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BTBillingPlanInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BTBillingPlanInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCompanyPlan

`func (o *BTBillingPlanInfo) GetCompanyPlan() bool`

GetCompanyPlan returns the CompanyPlan field if non-nil, zero value otherwise.

### GetCompanyPlanOk

`func (o *BTBillingPlanInfo) GetCompanyPlanOk() (*bool, bool)`

GetCompanyPlanOk returns a tuple with the CompanyPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPlan

`func (o *BTBillingPlanInfo) SetCompanyPlan(v bool)`

SetCompanyPlan sets CompanyPlan field to given value.

### HasCompanyPlan

`func (o *BTBillingPlanInfo) HasCompanyPlan() bool`

HasCompanyPlan returns a boolean if a field has been set.

### GetConsumableQuantity

`func (o *BTBillingPlanInfo) GetConsumableQuantity() int32`

GetConsumableQuantity returns the ConsumableQuantity field if non-nil, zero value otherwise.

### GetConsumableQuantityOk

`func (o *BTBillingPlanInfo) GetConsumableQuantityOk() (*int32, bool)`

GetConsumableQuantityOk returns a tuple with the ConsumableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumableQuantity

`func (o *BTBillingPlanInfo) SetConsumableQuantity(v int32)`

SetConsumableQuantity sets ConsumableQuantity field to given value.

### HasConsumableQuantity

`func (o *BTBillingPlanInfo) HasConsumableQuantity() bool`

HasConsumableQuantity returns a boolean if a field has been set.

### GetDeprecated

`func (o *BTBillingPlanInfo) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *BTBillingPlanInfo) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *BTBillingPlanInfo) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *BTBillingPlanInfo) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *BTBillingPlanInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTBillingPlanInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTBillingPlanInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTBillingPlanInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountInfo

`func (o *BTBillingPlanInfo) GetDiscountInfo() BTDiscountInfo`

GetDiscountInfo returns the DiscountInfo field if non-nil, zero value otherwise.

### GetDiscountInfoOk

`func (o *BTBillingPlanInfo) GetDiscountInfoOk() (*BTDiscountInfo, bool)`

GetDiscountInfoOk returns a tuple with the DiscountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountInfo

`func (o *BTBillingPlanInfo) SetDiscountInfo(v BTDiscountInfo)`

SetDiscountInfo sets DiscountInfo field to given value.

### HasDiscountInfo

`func (o *BTBillingPlanInfo) HasDiscountInfo() bool`

HasDiscountInfo returns a boolean if a field has been set.

### GetGroup

`func (o *BTBillingPlanInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BTBillingPlanInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BTBillingPlanInfo) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BTBillingPlanInfo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHidden

`func (o *BTBillingPlanInfo) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BTBillingPlanInfo) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BTBillingPlanInfo) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BTBillingPlanInfo) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHref

`func (o *BTBillingPlanInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTBillingPlanInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTBillingPlanInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTBillingPlanInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTBillingPlanInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTBillingPlanInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTBillingPlanInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTBillingPlanInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *BTBillingPlanInfo) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BTBillingPlanInfo) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BTBillingPlanInfo) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BTBillingPlanInfo) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *BTBillingPlanInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTBillingPlanInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTBillingPlanInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTBillingPlanInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnshapePlan

`func (o *BTBillingPlanInfo) GetOnshapePlan() bool`

GetOnshapePlan returns the OnshapePlan field if non-nil, zero value otherwise.

### GetOnshapePlanOk

`func (o *BTBillingPlanInfo) GetOnshapePlanOk() (*bool, bool)`

GetOnshapePlanOk returns a tuple with the OnshapePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnshapePlan

`func (o *BTBillingPlanInfo) SetOnshapePlan(v bool)`

SetOnshapePlan sets OnshapePlan field to given value.

### HasOnshapePlan

`func (o *BTBillingPlanInfo) HasOnshapePlan() bool`

HasOnshapePlan returns a boolean if a field has been set.

### GetPlanType

`func (o *BTBillingPlanInfo) GetPlanType() int32`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *BTBillingPlanInfo) GetPlanTypeOk() (*int32, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *BTBillingPlanInfo) SetPlanType(v int32)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *BTBillingPlanInfo) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetTrialPeriodDays

`func (o *BTBillingPlanInfo) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *BTBillingPlanInfo) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *BTBillingPlanInfo) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *BTBillingPlanInfo) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### GetViewRef

`func (o *BTBillingPlanInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTBillingPlanInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTBillingPlanInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTBillingPlanInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


