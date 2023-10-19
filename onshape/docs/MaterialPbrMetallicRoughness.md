# MaterialPbrMetallicRoughness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseColorFactor** | Pointer to **[]float32** |  | [optional] 
**BaseColorTexture** | Pointer to [**TextureInfo**](TextureInfo.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**MetallicFactor** | Pointer to **float32** |  | [optional] 
**MetallicRoughnessTexture** | Pointer to [**TextureInfo**](TextureInfo.md) |  | [optional] 
**RoughnessFactor** | Pointer to **float32** |  | [optional] 

## Methods

### NewMaterialPbrMetallicRoughness

`func NewMaterialPbrMetallicRoughness() *MaterialPbrMetallicRoughness`

NewMaterialPbrMetallicRoughness instantiates a new MaterialPbrMetallicRoughness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaterialPbrMetallicRoughnessWithDefaults

`func NewMaterialPbrMetallicRoughnessWithDefaults() *MaterialPbrMetallicRoughness`

NewMaterialPbrMetallicRoughnessWithDefaults instantiates a new MaterialPbrMetallicRoughness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseColorFactor

`func (o *MaterialPbrMetallicRoughness) GetBaseColorFactor() []float32`

GetBaseColorFactor returns the BaseColorFactor field if non-nil, zero value otherwise.

### GetBaseColorFactorOk

`func (o *MaterialPbrMetallicRoughness) GetBaseColorFactorOk() (*[]float32, bool)`

GetBaseColorFactorOk returns a tuple with the BaseColorFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseColorFactor

`func (o *MaterialPbrMetallicRoughness) SetBaseColorFactor(v []float32)`

SetBaseColorFactor sets BaseColorFactor field to given value.

### HasBaseColorFactor

`func (o *MaterialPbrMetallicRoughness) HasBaseColorFactor() bool`

HasBaseColorFactor returns a boolean if a field has been set.

### GetBaseColorTexture

`func (o *MaterialPbrMetallicRoughness) GetBaseColorTexture() TextureInfo`

GetBaseColorTexture returns the BaseColorTexture field if non-nil, zero value otherwise.

### GetBaseColorTextureOk

`func (o *MaterialPbrMetallicRoughness) GetBaseColorTextureOk() (*TextureInfo, bool)`

GetBaseColorTextureOk returns a tuple with the BaseColorTexture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseColorTexture

`func (o *MaterialPbrMetallicRoughness) SetBaseColorTexture(v TextureInfo)`

SetBaseColorTexture sets BaseColorTexture field to given value.

### HasBaseColorTexture

`func (o *MaterialPbrMetallicRoughness) HasBaseColorTexture() bool`

HasBaseColorTexture returns a boolean if a field has been set.

### GetExtensions

`func (o *MaterialPbrMetallicRoughness) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MaterialPbrMetallicRoughness) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MaterialPbrMetallicRoughness) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MaterialPbrMetallicRoughness) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *MaterialPbrMetallicRoughness) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *MaterialPbrMetallicRoughness) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *MaterialPbrMetallicRoughness) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *MaterialPbrMetallicRoughness) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetMetallicFactor

`func (o *MaterialPbrMetallicRoughness) GetMetallicFactor() float32`

GetMetallicFactor returns the MetallicFactor field if non-nil, zero value otherwise.

### GetMetallicFactorOk

`func (o *MaterialPbrMetallicRoughness) GetMetallicFactorOk() (*float32, bool)`

GetMetallicFactorOk returns a tuple with the MetallicFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetallicFactor

`func (o *MaterialPbrMetallicRoughness) SetMetallicFactor(v float32)`

SetMetallicFactor sets MetallicFactor field to given value.

### HasMetallicFactor

`func (o *MaterialPbrMetallicRoughness) HasMetallicFactor() bool`

HasMetallicFactor returns a boolean if a field has been set.

### GetMetallicRoughnessTexture

`func (o *MaterialPbrMetallicRoughness) GetMetallicRoughnessTexture() TextureInfo`

GetMetallicRoughnessTexture returns the MetallicRoughnessTexture field if non-nil, zero value otherwise.

### GetMetallicRoughnessTextureOk

`func (o *MaterialPbrMetallicRoughness) GetMetallicRoughnessTextureOk() (*TextureInfo, bool)`

GetMetallicRoughnessTextureOk returns a tuple with the MetallicRoughnessTexture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetallicRoughnessTexture

`func (o *MaterialPbrMetallicRoughness) SetMetallicRoughnessTexture(v TextureInfo)`

SetMetallicRoughnessTexture sets MetallicRoughnessTexture field to given value.

### HasMetallicRoughnessTexture

`func (o *MaterialPbrMetallicRoughness) HasMetallicRoughnessTexture() bool`

HasMetallicRoughnessTexture returns a boolean if a field has been set.

### GetRoughnessFactor

`func (o *MaterialPbrMetallicRoughness) GetRoughnessFactor() float32`

GetRoughnessFactor returns the RoughnessFactor field if non-nil, zero value otherwise.

### GetRoughnessFactorOk

`func (o *MaterialPbrMetallicRoughness) GetRoughnessFactorOk() (*float32, bool)`

GetRoughnessFactorOk returns a tuple with the RoughnessFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoughnessFactor

`func (o *MaterialPbrMetallicRoughness) SetRoughnessFactor(v float32)`

SetRoughnessFactor sets RoughnessFactor field to given value.

### HasRoughnessFactor

`func (o *MaterialPbrMetallicRoughness) HasRoughnessFactor() bool`

HasRoughnessFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


