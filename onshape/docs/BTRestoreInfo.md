# BTRestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRestoreStrategy** | Pointer to [**BTRestoreStrategy**](BTRestoreStrategy.md) |  | [optional] 
**ElementIdToStrategyOverride** | Pointer to [**map[string]BTRestoreStrategy**](BTRestoreStrategy.md) |  | [optional] 

## Methods

### NewBTRestoreInfo

`func NewBTRestoreInfo() *BTRestoreInfo`

NewBTRestoreInfo instantiates a new BTRestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRestoreInfoWithDefaults

`func NewBTRestoreInfoWithDefaults() *BTRestoreInfo`

NewBTRestoreInfoWithDefaults instantiates a new BTRestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRestoreStrategy

`func (o *BTRestoreInfo) GetDefaultRestoreStrategy() BTRestoreStrategy`

GetDefaultRestoreStrategy returns the DefaultRestoreStrategy field if non-nil, zero value otherwise.

### GetDefaultRestoreStrategyOk

`func (o *BTRestoreInfo) GetDefaultRestoreStrategyOk() (*BTRestoreStrategy, bool)`

GetDefaultRestoreStrategyOk returns a tuple with the DefaultRestoreStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRestoreStrategy

`func (o *BTRestoreInfo) SetDefaultRestoreStrategy(v BTRestoreStrategy)`

SetDefaultRestoreStrategy sets DefaultRestoreStrategy field to given value.

### HasDefaultRestoreStrategy

`func (o *BTRestoreInfo) HasDefaultRestoreStrategy() bool`

HasDefaultRestoreStrategy returns a boolean if a field has been set.

### GetElementIdToStrategyOverride

`func (o *BTRestoreInfo) GetElementIdToStrategyOverride() map[string]BTRestoreStrategy`

GetElementIdToStrategyOverride returns the ElementIdToStrategyOverride field if non-nil, zero value otherwise.

### GetElementIdToStrategyOverrideOk

`func (o *BTRestoreInfo) GetElementIdToStrategyOverrideOk() (*map[string]BTRestoreStrategy, bool)`

GetElementIdToStrategyOverrideOk returns a tuple with the ElementIdToStrategyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIdToStrategyOverride

`func (o *BTRestoreInfo) SetElementIdToStrategyOverride(v map[string]BTRestoreStrategy)`

SetElementIdToStrategyOverride sets ElementIdToStrategyOverride field to given value.

### HasElementIdToStrategyOverride

`func (o *BTRestoreInfo) HasElementIdToStrategyOverride() bool`

HasElementIdToStrategyOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


