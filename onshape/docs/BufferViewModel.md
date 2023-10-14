# BufferViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BufferModel** | Pointer to [**BufferModel**](BufferModel.md) |  | [optional] 
**BufferViewData** | Pointer to [**BufferViewModelBufferViewData**](BufferViewModelBufferViewData.md) |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**ByteOffset** | Pointer to **int32** |  | [optional] 
**ByteStride** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **int32** |  | [optional] 

## Methods

### NewBufferViewModel

`func NewBufferViewModel() *BufferViewModel`

NewBufferViewModel instantiates a new BufferViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBufferViewModelWithDefaults

`func NewBufferViewModelWithDefaults() *BufferViewModel`

NewBufferViewModelWithDefaults instantiates a new BufferViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBufferModel

`func (o *BufferViewModel) GetBufferModel() BufferModel`

GetBufferModel returns the BufferModel field if non-nil, zero value otherwise.

### GetBufferModelOk

`func (o *BufferViewModel) GetBufferModelOk() (*BufferModel, bool)`

GetBufferModelOk returns a tuple with the BufferModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferModel

`func (o *BufferViewModel) SetBufferModel(v BufferModel)`

SetBufferModel sets BufferModel field to given value.

### HasBufferModel

`func (o *BufferViewModel) HasBufferModel() bool`

HasBufferModel returns a boolean if a field has been set.

### GetBufferViewData

`func (o *BufferViewModel) GetBufferViewData() BufferViewModelBufferViewData`

GetBufferViewData returns the BufferViewData field if non-nil, zero value otherwise.

### GetBufferViewDataOk

`func (o *BufferViewModel) GetBufferViewDataOk() (*BufferViewModelBufferViewData, bool)`

GetBufferViewDataOk returns a tuple with the BufferViewData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferViewData

`func (o *BufferViewModel) SetBufferViewData(v BufferViewModelBufferViewData)`

SetBufferViewData sets BufferViewData field to given value.

### HasBufferViewData

`func (o *BufferViewModel) HasBufferViewData() bool`

HasBufferViewData returns a boolean if a field has been set.

### GetByteLength

`func (o *BufferViewModel) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *BufferViewModel) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *BufferViewModel) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *BufferViewModel) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetByteOffset

`func (o *BufferViewModel) GetByteOffset() int32`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *BufferViewModel) GetByteOffsetOk() (*int32, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *BufferViewModel) SetByteOffset(v int32)`

SetByteOffset sets ByteOffset field to given value.

### HasByteOffset

`func (o *BufferViewModel) HasByteOffset() bool`

HasByteOffset returns a boolean if a field has been set.

### GetByteStride

`func (o *BufferViewModel) GetByteStride() int32`

GetByteStride returns the ByteStride field if non-nil, zero value otherwise.

### GetByteStrideOk

`func (o *BufferViewModel) GetByteStrideOk() (*int32, bool)`

GetByteStrideOk returns a tuple with the ByteStride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteStride

`func (o *BufferViewModel) SetByteStride(v int32)`

SetByteStride sets ByteStride field to given value.

### HasByteStride

`func (o *BufferViewModel) HasByteStride() bool`

HasByteStride returns a boolean if a field has been set.

### GetName

`func (o *BufferViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BufferViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BufferViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BufferViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *BufferViewModel) GetTarget() int32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BufferViewModel) GetTargetOk() (*int32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BufferViewModel) SetTarget(v int32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BufferViewModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


