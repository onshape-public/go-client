# BTTrialInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidCustomerInPast** | Pointer to **bool** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PlanInterval** | Pointer to **string** |  | [optional] 
**Seats** | Pointer to **int64** |  | [optional] 
**TrialDaysRemaining** | Pointer to **int64** |  | [optional] 
**TrialEndDate** | Pointer to **JSONTime** |  | [optional] 
**TrialStartDate** | Pointer to **JSONTime** |  | [optional] 

## Methods

### NewBTTrialInfo

`func NewBTTrialInfo() *BTTrialInfo`

NewBTTrialInfo instantiates a new BTTrialInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTrialInfoWithDefaults

`func NewBTTrialInfoWithDefaults() *BTTrialInfo`

NewBTTrialInfoWithDefaults instantiates a new BTTrialInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidCustomerInPast

`func (o *BTTrialInfo) GetPaidCustomerInPast() bool`

GetPaidCustomerInPast returns the PaidCustomerInPast field if non-nil, zero value otherwise.

### GetPaidCustomerInPastOk

`func (o *BTTrialInfo) GetPaidCustomerInPastOk() (*bool, bool)`

GetPaidCustomerInPastOk returns a tuple with the PaidCustomerInPast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCustomerInPast

`func (o *BTTrialInfo) SetPaidCustomerInPast(v bool)`

SetPaidCustomerInPast sets PaidCustomerInPast field to given value.

### HasPaidCustomerInPast

`func (o *BTTrialInfo) HasPaidCustomerInPast() bool`

HasPaidCustomerInPast returns a boolean if a field has been set.

### GetPlanId

`func (o *BTTrialInfo) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *BTTrialInfo) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *BTTrialInfo) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *BTTrialInfo) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanInterval

`func (o *BTTrialInfo) GetPlanInterval() string`

GetPlanInterval returns the PlanInterval field if non-nil, zero value otherwise.

### GetPlanIntervalOk

`func (o *BTTrialInfo) GetPlanIntervalOk() (*string, bool)`

GetPlanIntervalOk returns a tuple with the PlanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanInterval

`func (o *BTTrialInfo) SetPlanInterval(v string)`

SetPlanInterval sets PlanInterval field to given value.

### HasPlanInterval

`func (o *BTTrialInfo) HasPlanInterval() bool`

HasPlanInterval returns a boolean if a field has been set.

### GetSeats

`func (o *BTTrialInfo) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *BTTrialInfo) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *BTTrialInfo) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *BTTrialInfo) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetTrialDaysRemaining

`func (o *BTTrialInfo) GetTrialDaysRemaining() int64`

GetTrialDaysRemaining returns the TrialDaysRemaining field if non-nil, zero value otherwise.

### GetTrialDaysRemainingOk

`func (o *BTTrialInfo) GetTrialDaysRemainingOk() (*int64, bool)`

GetTrialDaysRemainingOk returns a tuple with the TrialDaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDaysRemaining

`func (o *BTTrialInfo) SetTrialDaysRemaining(v int64)`

SetTrialDaysRemaining sets TrialDaysRemaining field to given value.

### HasTrialDaysRemaining

`func (o *BTTrialInfo) HasTrialDaysRemaining() bool`

HasTrialDaysRemaining returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *BTTrialInfo) GetTrialEndDate() JSONTime`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *BTTrialInfo) GetTrialEndDateOk() (*JSONTime, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *BTTrialInfo) SetTrialEndDate(v JSONTime)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *BTTrialInfo) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetTrialStartDate

`func (o *BTTrialInfo) GetTrialStartDate() JSONTime`

GetTrialStartDate returns the TrialStartDate field if non-nil, zero value otherwise.

### GetTrialStartDateOk

`func (o *BTTrialInfo) GetTrialStartDateOk() (*JSONTime, bool)`

GetTrialStartDateOk returns a tuple with the TrialStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartDate

`func (o *BTTrialInfo) SetTrialStartDate(v JSONTime)`

SetTrialStartDate sets TrialStartDate field to given value.

### HasTrialStartDate

`func (o *BTTrialInfo) HasTrialStartDate() bool`

HasTrialStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


