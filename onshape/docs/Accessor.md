# Accessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BufferView** | Pointer to **int32** |  | [optional] 
**ByteOffset** | Pointer to **int32** |  | [optional] 
**ComponentType** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Max** | Pointer to **[]float32** |  | [optional] 
**Min** | Pointer to **[]float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Normalized** | Pointer to **bool** |  | [optional] 
**Sparse** | Pointer to [**AccessorSparse**](AccessorSparse.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessor

`func NewAccessor() *Accessor`

NewAccessor instantiates a new Accessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessorWithDefaults

`func NewAccessorWithDefaults() *Accessor`

NewAccessorWithDefaults instantiates a new Accessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBufferView

`func (o *Accessor) GetBufferView() int32`

GetBufferView returns the BufferView field if non-nil, zero value otherwise.

### GetBufferViewOk

`func (o *Accessor) GetBufferViewOk() (*int32, bool)`

GetBufferViewOk returns a tuple with the BufferView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferView

`func (o *Accessor) SetBufferView(v int32)`

SetBufferView sets BufferView field to given value.

### HasBufferView

`func (o *Accessor) HasBufferView() bool`

HasBufferView returns a boolean if a field has been set.

### GetByteOffset

`func (o *Accessor) GetByteOffset() int32`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *Accessor) GetByteOffsetOk() (*int32, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *Accessor) SetByteOffset(v int32)`

SetByteOffset sets ByteOffset field to given value.

### HasByteOffset

`func (o *Accessor) HasByteOffset() bool`

HasByteOffset returns a boolean if a field has been set.

### GetComponentType

`func (o *Accessor) GetComponentType() int32`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *Accessor) GetComponentTypeOk() (*int32, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *Accessor) SetComponentType(v int32)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *Accessor) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetCount

`func (o *Accessor) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Accessor) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Accessor) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Accessor) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExtensions

`func (o *Accessor) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Accessor) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Accessor) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Accessor) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Accessor) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Accessor) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Accessor) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Accessor) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetMax

`func (o *Accessor) GetMax() []float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Accessor) GetMaxOk() (*[]float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Accessor) SetMax(v []float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *Accessor) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *Accessor) GetMin() []float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Accessor) GetMinOk() (*[]float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Accessor) SetMin(v []float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *Accessor) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetName

`func (o *Accessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Accessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Accessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Accessor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNormalized

`func (o *Accessor) GetNormalized() bool`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *Accessor) GetNormalizedOk() (*bool, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *Accessor) SetNormalized(v bool)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *Accessor) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### GetSparse

`func (o *Accessor) GetSparse() AccessorSparse`

GetSparse returns the Sparse field if non-nil, zero value otherwise.

### GetSparseOk

`func (o *Accessor) GetSparseOk() (*AccessorSparse, bool)`

GetSparseOk returns a tuple with the Sparse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparse

`func (o *Accessor) SetSparse(v AccessorSparse)`

SetSparse sets Sparse field to given value.

### HasSparse

`func (o *Accessor) HasSparse() bool`

HasSparse returns a boolean if a field has been set.

### GetType

`func (o *Accessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Accessor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Accessor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Accessor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


