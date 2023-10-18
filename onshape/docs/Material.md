# Material

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlphaCutoff** | Pointer to **float32** |  | [optional] 
**AlphaMode** | Pointer to **string** |  | [optional] 
**DoubleSided** | Pointer to **bool** |  | [optional] 
**EmissiveFactor** | Pointer to **[]float32** |  | [optional] 
**EmissiveTexture** | Pointer to [**TextureInfo**](TextureInfo.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NormalTexture** | Pointer to [**MaterialNormalTextureInfo**](MaterialNormalTextureInfo.md) |  | [optional] 
**OcclusionTexture** | Pointer to [**MaterialOcclusionTextureInfo**](MaterialOcclusionTextureInfo.md) |  | [optional] 
**PbrMetallicRoughness** | Pointer to [**MaterialPbrMetallicRoughness**](MaterialPbrMetallicRoughness.md) |  | [optional] 

## Methods

### NewMaterial

`func NewMaterial() *Material`

NewMaterial instantiates a new Material object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaterialWithDefaults

`func NewMaterialWithDefaults() *Material`

NewMaterialWithDefaults instantiates a new Material object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlphaCutoff

`func (o *Material) GetAlphaCutoff() float32`

GetAlphaCutoff returns the AlphaCutoff field if non-nil, zero value otherwise.

### GetAlphaCutoffOk

`func (o *Material) GetAlphaCutoffOk() (*float32, bool)`

GetAlphaCutoffOk returns a tuple with the AlphaCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaCutoff

`func (o *Material) SetAlphaCutoff(v float32)`

SetAlphaCutoff sets AlphaCutoff field to given value.

### HasAlphaCutoff

`func (o *Material) HasAlphaCutoff() bool`

HasAlphaCutoff returns a boolean if a field has been set.

### GetAlphaMode

`func (o *Material) GetAlphaMode() string`

GetAlphaMode returns the AlphaMode field if non-nil, zero value otherwise.

### GetAlphaModeOk

`func (o *Material) GetAlphaModeOk() (*string, bool)`

GetAlphaModeOk returns a tuple with the AlphaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaMode

`func (o *Material) SetAlphaMode(v string)`

SetAlphaMode sets AlphaMode field to given value.

### HasAlphaMode

`func (o *Material) HasAlphaMode() bool`

HasAlphaMode returns a boolean if a field has been set.

### GetDoubleSided

`func (o *Material) GetDoubleSided() bool`

GetDoubleSided returns the DoubleSided field if non-nil, zero value otherwise.

### GetDoubleSidedOk

`func (o *Material) GetDoubleSidedOk() (*bool, bool)`

GetDoubleSidedOk returns a tuple with the DoubleSided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSided

`func (o *Material) SetDoubleSided(v bool)`

SetDoubleSided sets DoubleSided field to given value.

### HasDoubleSided

`func (o *Material) HasDoubleSided() bool`

HasDoubleSided returns a boolean if a field has been set.

### GetEmissiveFactor

`func (o *Material) GetEmissiveFactor() []float32`

GetEmissiveFactor returns the EmissiveFactor field if non-nil, zero value otherwise.

### GetEmissiveFactorOk

`func (o *Material) GetEmissiveFactorOk() (*[]float32, bool)`

GetEmissiveFactorOk returns a tuple with the EmissiveFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmissiveFactor

`func (o *Material) SetEmissiveFactor(v []float32)`

SetEmissiveFactor sets EmissiveFactor field to given value.

### HasEmissiveFactor

`func (o *Material) HasEmissiveFactor() bool`

HasEmissiveFactor returns a boolean if a field has been set.

### GetEmissiveTexture

`func (o *Material) GetEmissiveTexture() TextureInfo`

GetEmissiveTexture returns the EmissiveTexture field if non-nil, zero value otherwise.

### GetEmissiveTextureOk

`func (o *Material) GetEmissiveTextureOk() (*TextureInfo, bool)`

GetEmissiveTextureOk returns a tuple with the EmissiveTexture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmissiveTexture

`func (o *Material) SetEmissiveTexture(v TextureInfo)`

SetEmissiveTexture sets EmissiveTexture field to given value.

### HasEmissiveTexture

`func (o *Material) HasEmissiveTexture() bool`

HasEmissiveTexture returns a boolean if a field has been set.

### GetExtensions

`func (o *Material) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Material) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Material) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Material) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Material) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Material) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Material) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Material) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *Material) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Material) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Material) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Material) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNormalTexture

`func (o *Material) GetNormalTexture() MaterialNormalTextureInfo`

GetNormalTexture returns the NormalTexture field if non-nil, zero value otherwise.

### GetNormalTextureOk

`func (o *Material) GetNormalTextureOk() (*MaterialNormalTextureInfo, bool)`

GetNormalTextureOk returns a tuple with the NormalTexture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalTexture

`func (o *Material) SetNormalTexture(v MaterialNormalTextureInfo)`

SetNormalTexture sets NormalTexture field to given value.

### HasNormalTexture

`func (o *Material) HasNormalTexture() bool`

HasNormalTexture returns a boolean if a field has been set.

### GetOcclusionTexture

`func (o *Material) GetOcclusionTexture() MaterialOcclusionTextureInfo`

GetOcclusionTexture returns the OcclusionTexture field if non-nil, zero value otherwise.

### GetOcclusionTextureOk

`func (o *Material) GetOcclusionTextureOk() (*MaterialOcclusionTextureInfo, bool)`

GetOcclusionTextureOk returns a tuple with the OcclusionTexture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcclusionTexture

`func (o *Material) SetOcclusionTexture(v MaterialOcclusionTextureInfo)`

SetOcclusionTexture sets OcclusionTexture field to given value.

### HasOcclusionTexture

`func (o *Material) HasOcclusionTexture() bool`

HasOcclusionTexture returns a boolean if a field has been set.

### GetPbrMetallicRoughness

`func (o *Material) GetPbrMetallicRoughness() MaterialPbrMetallicRoughness`

GetPbrMetallicRoughness returns the PbrMetallicRoughness field if non-nil, zero value otherwise.

### GetPbrMetallicRoughnessOk

`func (o *Material) GetPbrMetallicRoughnessOk() (*MaterialPbrMetallicRoughness, bool)`

GetPbrMetallicRoughnessOk returns a tuple with the PbrMetallicRoughness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbrMetallicRoughness

`func (o *Material) SetPbrMetallicRoughness(v MaterialPbrMetallicRoughness)`

SetPbrMetallicRoughness sets PbrMetallicRoughness field to given value.

### HasPbrMetallicRoughness

`func (o *Material) HasPbrMetallicRoughness() bool`

HasPbrMetallicRoughness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


