# Texture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Sampler** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **int32** |  | [optional] 

## Methods

### NewTexture

`func NewTexture() *Texture`

NewTexture instantiates a new Texture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextureWithDefaults

`func NewTextureWithDefaults() *Texture`

NewTextureWithDefaults instantiates a new Texture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *Texture) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Texture) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Texture) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Texture) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Texture) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Texture) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Texture) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Texture) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *Texture) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Texture) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Texture) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Texture) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSampler

`func (o *Texture) GetSampler() int32`

GetSampler returns the Sampler field if non-nil, zero value otherwise.

### GetSamplerOk

`func (o *Texture) GetSamplerOk() (*int32, bool)`

GetSamplerOk returns a tuple with the Sampler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampler

`func (o *Texture) SetSampler(v int32)`

SetSampler sets Sampler field to given value.

### HasSampler

`func (o *Texture) HasSampler() bool`

HasSampler returns a boolean if a field has been set.

### GetSource

`func (o *Texture) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Texture) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Texture) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *Texture) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


