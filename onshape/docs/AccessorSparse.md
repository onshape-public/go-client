# AccessorSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Indices** | Pointer to [**AccessorSparseIndices**](AccessorSparseIndices.md) |  | [optional] 
**Values** | Pointer to [**AccessorSparseValues**](AccessorSparseValues.md) |  | [optional] 

## Methods

### NewAccessorSparse

`func NewAccessorSparse() *AccessorSparse`

NewAccessorSparse instantiates a new AccessorSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessorSparseWithDefaults

`func NewAccessorSparseWithDefaults() *AccessorSparse`

NewAccessorSparseWithDefaults instantiates a new AccessorSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AccessorSparse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AccessorSparse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AccessorSparse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AccessorSparse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExtensions

`func (o *AccessorSparse) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *AccessorSparse) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *AccessorSparse) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *AccessorSparse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *AccessorSparse) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *AccessorSparse) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *AccessorSparse) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *AccessorSparse) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetIndices

`func (o *AccessorSparse) GetIndices() AccessorSparseIndices`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *AccessorSparse) GetIndicesOk() (*AccessorSparseIndices, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *AccessorSparse) SetIndices(v AccessorSparseIndices)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *AccessorSparse) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetValues

`func (o *AccessorSparse) GetValues() AccessorSparseValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AccessorSparse) GetValuesOk() (*AccessorSparseValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AccessorSparse) SetValues(v AccessorSparseValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *AccessorSparse) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


