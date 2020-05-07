# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**IntervalCount** | Pointer to **int32** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**StatementDescription** | Pointer to **string** |  | [optional] 
**StatementDescriptor** | Pointer to **string** |  | [optional] 
**TrialPeriodDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlan

`func NewPlan() *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Plan) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Plan) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Plan) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Plan) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreated

`func (o *Plan) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Plan) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Plan) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Plan) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *Plan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Plan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Plan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Plan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Plan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *Plan) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Plan) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Plan) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Plan) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIntervalCount

`func (o *Plan) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *Plan) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *Plan) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *Plan) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetLivemode

`func (o *Plan) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Plan) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Plan) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Plan) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMetadata

`func (o *Plan) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Plan) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Plan) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Plan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *Plan) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Plan) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Plan) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Plan) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatementDescription

`func (o *Plan) GetStatementDescription() string`

GetStatementDescription returns the StatementDescription field if non-nil, zero value otherwise.

### GetStatementDescriptionOk

`func (o *Plan) GetStatementDescriptionOk() (*string, bool)`

GetStatementDescriptionOk returns a tuple with the StatementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescription

`func (o *Plan) SetStatementDescription(v string)`

SetStatementDescription sets StatementDescription field to given value.

### HasStatementDescription

`func (o *Plan) HasStatementDescription() bool`

HasStatementDescription returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *Plan) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *Plan) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *Plan) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *Plan) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetTrialPeriodDays

`func (o *Plan) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *Plan) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *Plan) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *Plan) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


