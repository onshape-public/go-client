# BTComputedAssemblyPropertyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatedPropertyId** | Pointer to **string** |  | [optional] 
**AggregationOperator** | Pointer to [**BTComputedAssemblyPropertyAggregationOperator**](BTComputedAssemblyPropertyAggregationOperator.md) |  | [optional] 
**ErrorValuePolicy** | Pointer to [**BTComputedAssemblyPropertyErrorPolicy**](BTComputedAssemblyPropertyErrorPolicy.md) |  | [optional] 
**FilterPropertyId** | Pointer to **string** |  | [optional] 
**IsFilterPropertyInverted** | Pointer to **bool** |  | [optional] 
**MissingValuePolicy** | Pointer to [**BTComputedAssemblyPropertyErrorPolicy**](BTComputedAssemblyPropertyErrorPolicy.md) |  | [optional] 
**SecondaryPropertyId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTComputedAssemblyPropertyConfig

`func NewBTComputedAssemblyPropertyConfig() *BTComputedAssemblyPropertyConfig`

NewBTComputedAssemblyPropertyConfig instantiates a new BTComputedAssemblyPropertyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTComputedAssemblyPropertyConfigWithDefaults

`func NewBTComputedAssemblyPropertyConfigWithDefaults() *BTComputedAssemblyPropertyConfig`

NewBTComputedAssemblyPropertyConfigWithDefaults instantiates a new BTComputedAssemblyPropertyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatedPropertyId

`func (o *BTComputedAssemblyPropertyConfig) GetAggregatedPropertyId() string`

GetAggregatedPropertyId returns the AggregatedPropertyId field if non-nil, zero value otherwise.

### GetAggregatedPropertyIdOk

`func (o *BTComputedAssemblyPropertyConfig) GetAggregatedPropertyIdOk() (*string, bool)`

GetAggregatedPropertyIdOk returns a tuple with the AggregatedPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedPropertyId

`func (o *BTComputedAssemblyPropertyConfig) SetAggregatedPropertyId(v string)`

SetAggregatedPropertyId sets AggregatedPropertyId field to given value.

### HasAggregatedPropertyId

`func (o *BTComputedAssemblyPropertyConfig) HasAggregatedPropertyId() bool`

HasAggregatedPropertyId returns a boolean if a field has been set.

### GetAggregationOperator

`func (o *BTComputedAssemblyPropertyConfig) GetAggregationOperator() BTComputedAssemblyPropertyAggregationOperator`

GetAggregationOperator returns the AggregationOperator field if non-nil, zero value otherwise.

### GetAggregationOperatorOk

`func (o *BTComputedAssemblyPropertyConfig) GetAggregationOperatorOk() (*BTComputedAssemblyPropertyAggregationOperator, bool)`

GetAggregationOperatorOk returns a tuple with the AggregationOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationOperator

`func (o *BTComputedAssemblyPropertyConfig) SetAggregationOperator(v BTComputedAssemblyPropertyAggregationOperator)`

SetAggregationOperator sets AggregationOperator field to given value.

### HasAggregationOperator

`func (o *BTComputedAssemblyPropertyConfig) HasAggregationOperator() bool`

HasAggregationOperator returns a boolean if a field has been set.

### GetErrorValuePolicy

`func (o *BTComputedAssemblyPropertyConfig) GetErrorValuePolicy() BTComputedAssemblyPropertyErrorPolicy`

GetErrorValuePolicy returns the ErrorValuePolicy field if non-nil, zero value otherwise.

### GetErrorValuePolicyOk

`func (o *BTComputedAssemblyPropertyConfig) GetErrorValuePolicyOk() (*BTComputedAssemblyPropertyErrorPolicy, bool)`

GetErrorValuePolicyOk returns a tuple with the ErrorValuePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValuePolicy

`func (o *BTComputedAssemblyPropertyConfig) SetErrorValuePolicy(v BTComputedAssemblyPropertyErrorPolicy)`

SetErrorValuePolicy sets ErrorValuePolicy field to given value.

### HasErrorValuePolicy

`func (o *BTComputedAssemblyPropertyConfig) HasErrorValuePolicy() bool`

HasErrorValuePolicy returns a boolean if a field has been set.

### GetFilterPropertyId

`func (o *BTComputedAssemblyPropertyConfig) GetFilterPropertyId() string`

GetFilterPropertyId returns the FilterPropertyId field if non-nil, zero value otherwise.

### GetFilterPropertyIdOk

`func (o *BTComputedAssemblyPropertyConfig) GetFilterPropertyIdOk() (*string, bool)`

GetFilterPropertyIdOk returns a tuple with the FilterPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPropertyId

`func (o *BTComputedAssemblyPropertyConfig) SetFilterPropertyId(v string)`

SetFilterPropertyId sets FilterPropertyId field to given value.

### HasFilterPropertyId

`func (o *BTComputedAssemblyPropertyConfig) HasFilterPropertyId() bool`

HasFilterPropertyId returns a boolean if a field has been set.

### GetIsFilterPropertyInverted

`func (o *BTComputedAssemblyPropertyConfig) GetIsFilterPropertyInverted() bool`

GetIsFilterPropertyInverted returns the IsFilterPropertyInverted field if non-nil, zero value otherwise.

### GetIsFilterPropertyInvertedOk

`func (o *BTComputedAssemblyPropertyConfig) GetIsFilterPropertyInvertedOk() (*bool, bool)`

GetIsFilterPropertyInvertedOk returns a tuple with the IsFilterPropertyInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFilterPropertyInverted

`func (o *BTComputedAssemblyPropertyConfig) SetIsFilterPropertyInverted(v bool)`

SetIsFilterPropertyInverted sets IsFilterPropertyInverted field to given value.

### HasIsFilterPropertyInverted

`func (o *BTComputedAssemblyPropertyConfig) HasIsFilterPropertyInverted() bool`

HasIsFilterPropertyInverted returns a boolean if a field has been set.

### GetMissingValuePolicy

`func (o *BTComputedAssemblyPropertyConfig) GetMissingValuePolicy() BTComputedAssemblyPropertyErrorPolicy`

GetMissingValuePolicy returns the MissingValuePolicy field if non-nil, zero value otherwise.

### GetMissingValuePolicyOk

`func (o *BTComputedAssemblyPropertyConfig) GetMissingValuePolicyOk() (*BTComputedAssemblyPropertyErrorPolicy, bool)`

GetMissingValuePolicyOk returns a tuple with the MissingValuePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingValuePolicy

`func (o *BTComputedAssemblyPropertyConfig) SetMissingValuePolicy(v BTComputedAssemblyPropertyErrorPolicy)`

SetMissingValuePolicy sets MissingValuePolicy field to given value.

### HasMissingValuePolicy

`func (o *BTComputedAssemblyPropertyConfig) HasMissingValuePolicy() bool`

HasMissingValuePolicy returns a boolean if a field has been set.

### GetSecondaryPropertyId

`func (o *BTComputedAssemblyPropertyConfig) GetSecondaryPropertyId() string`

GetSecondaryPropertyId returns the SecondaryPropertyId field if non-nil, zero value otherwise.

### GetSecondaryPropertyIdOk

`func (o *BTComputedAssemblyPropertyConfig) GetSecondaryPropertyIdOk() (*string, bool)`

GetSecondaryPropertyIdOk returns a tuple with the SecondaryPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPropertyId

`func (o *BTComputedAssemblyPropertyConfig) SetSecondaryPropertyId(v string)`

SetSecondaryPropertyId sets SecondaryPropertyId field to given value.

### HasSecondaryPropertyId

`func (o *BTComputedAssemblyPropertyConfig) HasSecondaryPropertyId() bool`

HasSecondaryPropertyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


