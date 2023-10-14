# AccessorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessorData** | Pointer to [**AccessorData**](AccessorData.md) |  | [optional] 
**BufferViewModel** | Pointer to [**BufferViewModel**](BufferViewModel.md) |  | [optional] 
**ByteOffset** | Pointer to **int32** |  | [optional] 
**ByteStride** | Pointer to **int32** |  | [optional] 
**ComponentSizeInBytes** | Pointer to **int32** |  | [optional] 
**ComponentType** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**ElementSizeInBytes** | Pointer to **int32** |  | [optional] 
**ElementType** | Pointer to [**ElementType**](ElementType.md) |  | [optional] 
**Max** | Pointer to **[]float32** |  | [optional] 
**Min** | Pointer to **[]float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessorModel

`func NewAccessorModel() *AccessorModel`

NewAccessorModel instantiates a new AccessorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessorModelWithDefaults

`func NewAccessorModelWithDefaults() *AccessorModel`

NewAccessorModelWithDefaults instantiates a new AccessorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessorData

`func (o *AccessorModel) GetAccessorData() AccessorData`

GetAccessorData returns the AccessorData field if non-nil, zero value otherwise.

### GetAccessorDataOk

`func (o *AccessorModel) GetAccessorDataOk() (*AccessorData, bool)`

GetAccessorDataOk returns a tuple with the AccessorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessorData

`func (o *AccessorModel) SetAccessorData(v AccessorData)`

SetAccessorData sets AccessorData field to given value.

### HasAccessorData

`func (o *AccessorModel) HasAccessorData() bool`

HasAccessorData returns a boolean if a field has been set.

### GetBufferViewModel

`func (o *AccessorModel) GetBufferViewModel() BufferViewModel`

GetBufferViewModel returns the BufferViewModel field if non-nil, zero value otherwise.

### GetBufferViewModelOk

`func (o *AccessorModel) GetBufferViewModelOk() (*BufferViewModel, bool)`

GetBufferViewModelOk returns a tuple with the BufferViewModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferViewModel

`func (o *AccessorModel) SetBufferViewModel(v BufferViewModel)`

SetBufferViewModel sets BufferViewModel field to given value.

### HasBufferViewModel

`func (o *AccessorModel) HasBufferViewModel() bool`

HasBufferViewModel returns a boolean if a field has been set.

### GetByteOffset

`func (o *AccessorModel) GetByteOffset() int32`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *AccessorModel) GetByteOffsetOk() (*int32, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *AccessorModel) SetByteOffset(v int32)`

SetByteOffset sets ByteOffset field to given value.

### HasByteOffset

`func (o *AccessorModel) HasByteOffset() bool`

HasByteOffset returns a boolean if a field has been set.

### GetByteStride

`func (o *AccessorModel) GetByteStride() int32`

GetByteStride returns the ByteStride field if non-nil, zero value otherwise.

### GetByteStrideOk

`func (o *AccessorModel) GetByteStrideOk() (*int32, bool)`

GetByteStrideOk returns a tuple with the ByteStride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteStride

`func (o *AccessorModel) SetByteStride(v int32)`

SetByteStride sets ByteStride field to given value.

### HasByteStride

`func (o *AccessorModel) HasByteStride() bool`

HasByteStride returns a boolean if a field has been set.

### GetComponentSizeInBytes

`func (o *AccessorModel) GetComponentSizeInBytes() int32`

GetComponentSizeInBytes returns the ComponentSizeInBytes field if non-nil, zero value otherwise.

### GetComponentSizeInBytesOk

`func (o *AccessorModel) GetComponentSizeInBytesOk() (*int32, bool)`

GetComponentSizeInBytesOk returns a tuple with the ComponentSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentSizeInBytes

`func (o *AccessorModel) SetComponentSizeInBytes(v int32)`

SetComponentSizeInBytes sets ComponentSizeInBytes field to given value.

### HasComponentSizeInBytes

`func (o *AccessorModel) HasComponentSizeInBytes() bool`

HasComponentSizeInBytes returns a boolean if a field has been set.

### GetComponentType

`func (o *AccessorModel) GetComponentType() int32`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *AccessorModel) GetComponentTypeOk() (*int32, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *AccessorModel) SetComponentType(v int32)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *AccessorModel) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetCount

`func (o *AccessorModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AccessorModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AccessorModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AccessorModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetElementSizeInBytes

`func (o *AccessorModel) GetElementSizeInBytes() int32`

GetElementSizeInBytes returns the ElementSizeInBytes field if non-nil, zero value otherwise.

### GetElementSizeInBytesOk

`func (o *AccessorModel) GetElementSizeInBytesOk() (*int32, bool)`

GetElementSizeInBytesOk returns a tuple with the ElementSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementSizeInBytes

`func (o *AccessorModel) SetElementSizeInBytes(v int32)`

SetElementSizeInBytes sets ElementSizeInBytes field to given value.

### HasElementSizeInBytes

`func (o *AccessorModel) HasElementSizeInBytes() bool`

HasElementSizeInBytes returns a boolean if a field has been set.

### GetElementType

`func (o *AccessorModel) GetElementType() ElementType`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *AccessorModel) GetElementTypeOk() (*ElementType, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *AccessorModel) SetElementType(v ElementType)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *AccessorModel) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetMax

`func (o *AccessorModel) GetMax() []float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *AccessorModel) GetMaxOk() (*[]float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *AccessorModel) SetMax(v []float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *AccessorModel) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *AccessorModel) GetMin() []float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *AccessorModel) GetMinOk() (*[]float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *AccessorModel) SetMin(v []float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *AccessorModel) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetName

`func (o *AccessorModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessorModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessorModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessorModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


