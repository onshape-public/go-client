# BufferView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buffer** | Pointer to **int32** |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**ByteOffset** | Pointer to **int32** |  | [optional] 
**ByteStride** | Pointer to **int32** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **int32** |  | [optional] 

## Methods

### NewBufferView

`func NewBufferView() *BufferView`

NewBufferView instantiates a new BufferView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBufferViewWithDefaults

`func NewBufferViewWithDefaults() *BufferView`

NewBufferViewWithDefaults instantiates a new BufferView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuffer

`func (o *BufferView) GetBuffer() int32`

GetBuffer returns the Buffer field if non-nil, zero value otherwise.

### GetBufferOk

`func (o *BufferView) GetBufferOk() (*int32, bool)`

GetBufferOk returns a tuple with the Buffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuffer

`func (o *BufferView) SetBuffer(v int32)`

SetBuffer sets Buffer field to given value.

### HasBuffer

`func (o *BufferView) HasBuffer() bool`

HasBuffer returns a boolean if a field has been set.

### GetByteLength

`func (o *BufferView) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *BufferView) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *BufferView) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *BufferView) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetByteOffset

`func (o *BufferView) GetByteOffset() int32`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *BufferView) GetByteOffsetOk() (*int32, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *BufferView) SetByteOffset(v int32)`

SetByteOffset sets ByteOffset field to given value.

### HasByteOffset

`func (o *BufferView) HasByteOffset() bool`

HasByteOffset returns a boolean if a field has been set.

### GetByteStride

`func (o *BufferView) GetByteStride() int32`

GetByteStride returns the ByteStride field if non-nil, zero value otherwise.

### GetByteStrideOk

`func (o *BufferView) GetByteStrideOk() (*int32, bool)`

GetByteStrideOk returns a tuple with the ByteStride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteStride

`func (o *BufferView) SetByteStride(v int32)`

SetByteStride sets ByteStride field to given value.

### HasByteStride

`func (o *BufferView) HasByteStride() bool`

HasByteStride returns a boolean if a field has been set.

### GetExtensions

`func (o *BufferView) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *BufferView) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *BufferView) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *BufferView) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *BufferView) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BufferView) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BufferView) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BufferView) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *BufferView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BufferView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BufferView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BufferView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *BufferView) GetTarget() int32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BufferView) GetTargetOk() (*int32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BufferView) SetTarget(v int32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BufferView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


