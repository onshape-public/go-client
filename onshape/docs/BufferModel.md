# BufferModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BufferData** | Pointer to [**BufferModelBufferData**](BufferModelBufferData.md) |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewBufferModel

`func NewBufferModel() *BufferModel`

NewBufferModel instantiates a new BufferModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBufferModelWithDefaults

`func NewBufferModelWithDefaults() *BufferModel`

NewBufferModelWithDefaults instantiates a new BufferModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBufferData

`func (o *BufferModel) GetBufferData() BufferModelBufferData`

GetBufferData returns the BufferData field if non-nil, zero value otherwise.

### GetBufferDataOk

`func (o *BufferModel) GetBufferDataOk() (*BufferModelBufferData, bool)`

GetBufferDataOk returns a tuple with the BufferData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferData

`func (o *BufferModel) SetBufferData(v BufferModelBufferData)`

SetBufferData sets BufferData field to given value.

### HasBufferData

`func (o *BufferModel) HasBufferData() bool`

HasBufferData returns a boolean if a field has been set.

### GetByteLength

`func (o *BufferModel) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *BufferModel) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *BufferModel) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *BufferModel) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetName

`func (o *BufferModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BufferModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BufferModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BufferModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUri

`func (o *BufferModel) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BufferModel) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BufferModel) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BufferModel) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


