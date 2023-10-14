# GlTF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessors** | Pointer to [**[]Accessor**](Accessor.md) |  | [optional] 
**Animations** | Pointer to [**[]Animation**](Animation.md) |  | [optional] 
**Asset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**BufferViews** | Pointer to [**[]BufferView**](BufferView.md) |  | [optional] 
**Buffers** | Pointer to [**[]Buffer**](Buffer.md) |  | [optional] 
**Cameras** | Pointer to [**[]Camera**](Camera.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ExtensionsRequired** | Pointer to **[]string** |  | [optional] 
**ExtensionsUsed** | Pointer to **[]string** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**Materials** | Pointer to [**[]Material**](Material.md) |  | [optional] 
**Meshes** | Pointer to [**[]Mesh**](Mesh.md) |  | [optional] 
**Nodes** | Pointer to [**[]Node**](Node.md) |  | [optional] 
**Samplers** | Pointer to [**[]Sampler**](Sampler.md) |  | [optional] 
**Scene** | Pointer to **int32** |  | [optional] 
**Scenes** | Pointer to [**[]Scene**](Scene.md) |  | [optional] 
**Skins** | Pointer to [**[]Skin**](Skin.md) |  | [optional] 
**Textures** | Pointer to [**[]Texture**](Texture.md) |  | [optional] 

## Methods

### NewGlTF

`func NewGlTF() *GlTF`

NewGlTF instantiates a new GlTF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlTFWithDefaults

`func NewGlTFWithDefaults() *GlTF`

NewGlTFWithDefaults instantiates a new GlTF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessors

`func (o *GlTF) GetAccessors() []Accessor`

GetAccessors returns the Accessors field if non-nil, zero value otherwise.

### GetAccessorsOk

`func (o *GlTF) GetAccessorsOk() (*[]Accessor, bool)`

GetAccessorsOk returns a tuple with the Accessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessors

`func (o *GlTF) SetAccessors(v []Accessor)`

SetAccessors sets Accessors field to given value.

### HasAccessors

`func (o *GlTF) HasAccessors() bool`

HasAccessors returns a boolean if a field has been set.

### GetAnimations

`func (o *GlTF) GetAnimations() []Animation`

GetAnimations returns the Animations field if non-nil, zero value otherwise.

### GetAnimationsOk

`func (o *GlTF) GetAnimationsOk() (*[]Animation, bool)`

GetAnimationsOk returns a tuple with the Animations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimations

`func (o *GlTF) SetAnimations(v []Animation)`

SetAnimations sets Animations field to given value.

### HasAnimations

`func (o *GlTF) HasAnimations() bool`

HasAnimations returns a boolean if a field has been set.

### GetAsset

`func (o *GlTF) GetAsset() Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GlTF) GetAssetOk() (*Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GlTF) SetAsset(v Asset)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GlTF) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetBufferViews

`func (o *GlTF) GetBufferViews() []BufferView`

GetBufferViews returns the BufferViews field if non-nil, zero value otherwise.

### GetBufferViewsOk

`func (o *GlTF) GetBufferViewsOk() (*[]BufferView, bool)`

GetBufferViewsOk returns a tuple with the BufferViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferViews

`func (o *GlTF) SetBufferViews(v []BufferView)`

SetBufferViews sets BufferViews field to given value.

### HasBufferViews

`func (o *GlTF) HasBufferViews() bool`

HasBufferViews returns a boolean if a field has been set.

### GetBuffers

`func (o *GlTF) GetBuffers() []Buffer`

GetBuffers returns the Buffers field if non-nil, zero value otherwise.

### GetBuffersOk

`func (o *GlTF) GetBuffersOk() (*[]Buffer, bool)`

GetBuffersOk returns a tuple with the Buffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuffers

`func (o *GlTF) SetBuffers(v []Buffer)`

SetBuffers sets Buffers field to given value.

### HasBuffers

`func (o *GlTF) HasBuffers() bool`

HasBuffers returns a boolean if a field has been set.

### GetCameras

`func (o *GlTF) GetCameras() []Camera`

GetCameras returns the Cameras field if non-nil, zero value otherwise.

### GetCamerasOk

`func (o *GlTF) GetCamerasOk() (*[]Camera, bool)`

GetCamerasOk returns a tuple with the Cameras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameras

`func (o *GlTF) SetCameras(v []Camera)`

SetCameras sets Cameras field to given value.

### HasCameras

`func (o *GlTF) HasCameras() bool`

HasCameras returns a boolean if a field has been set.

### GetExtensions

`func (o *GlTF) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *GlTF) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *GlTF) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *GlTF) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtensionsRequired

`func (o *GlTF) GetExtensionsRequired() []string`

GetExtensionsRequired returns the ExtensionsRequired field if non-nil, zero value otherwise.

### GetExtensionsRequiredOk

`func (o *GlTF) GetExtensionsRequiredOk() (*[]string, bool)`

GetExtensionsRequiredOk returns a tuple with the ExtensionsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionsRequired

`func (o *GlTF) SetExtensionsRequired(v []string)`

SetExtensionsRequired sets ExtensionsRequired field to given value.

### HasExtensionsRequired

`func (o *GlTF) HasExtensionsRequired() bool`

HasExtensionsRequired returns a boolean if a field has been set.

### GetExtensionsUsed

`func (o *GlTF) GetExtensionsUsed() []string`

GetExtensionsUsed returns the ExtensionsUsed field if non-nil, zero value otherwise.

### GetExtensionsUsedOk

`func (o *GlTF) GetExtensionsUsedOk() (*[]string, bool)`

GetExtensionsUsedOk returns a tuple with the ExtensionsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionsUsed

`func (o *GlTF) SetExtensionsUsed(v []string)`

SetExtensionsUsed sets ExtensionsUsed field to given value.

### HasExtensionsUsed

`func (o *GlTF) HasExtensionsUsed() bool`

HasExtensionsUsed returns a boolean if a field has been set.

### GetExtras

`func (o *GlTF) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *GlTF) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *GlTF) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *GlTF) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetImages

`func (o *GlTF) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GlTF) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GlTF) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *GlTF) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetMaterials

`func (o *GlTF) GetMaterials() []Material`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *GlTF) GetMaterialsOk() (*[]Material, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *GlTF) SetMaterials(v []Material)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *GlTF) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetMeshes

`func (o *GlTF) GetMeshes() []Mesh`

GetMeshes returns the Meshes field if non-nil, zero value otherwise.

### GetMeshesOk

`func (o *GlTF) GetMeshesOk() (*[]Mesh, bool)`

GetMeshesOk returns a tuple with the Meshes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshes

`func (o *GlTF) SetMeshes(v []Mesh)`

SetMeshes sets Meshes field to given value.

### HasMeshes

`func (o *GlTF) HasMeshes() bool`

HasMeshes returns a boolean if a field has been set.

### GetNodes

`func (o *GlTF) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GlTF) GetNodesOk() (*[]Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GlTF) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GlTF) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSamplers

`func (o *GlTF) GetSamplers() []Sampler`

GetSamplers returns the Samplers field if non-nil, zero value otherwise.

### GetSamplersOk

`func (o *GlTF) GetSamplersOk() (*[]Sampler, bool)`

GetSamplersOk returns a tuple with the Samplers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplers

`func (o *GlTF) SetSamplers(v []Sampler)`

SetSamplers sets Samplers field to given value.

### HasSamplers

`func (o *GlTF) HasSamplers() bool`

HasSamplers returns a boolean if a field has been set.

### GetScene

`func (o *GlTF) GetScene() int32`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *GlTF) GetSceneOk() (*int32, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *GlTF) SetScene(v int32)`

SetScene sets Scene field to given value.

### HasScene

`func (o *GlTF) HasScene() bool`

HasScene returns a boolean if a field has been set.

### GetScenes

`func (o *GlTF) GetScenes() []Scene`

GetScenes returns the Scenes field if non-nil, zero value otherwise.

### GetScenesOk

`func (o *GlTF) GetScenesOk() (*[]Scene, bool)`

GetScenesOk returns a tuple with the Scenes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenes

`func (o *GlTF) SetScenes(v []Scene)`

SetScenes sets Scenes field to given value.

### HasScenes

`func (o *GlTF) HasScenes() bool`

HasScenes returns a boolean if a field has been set.

### GetSkins

`func (o *GlTF) GetSkins() []Skin`

GetSkins returns the Skins field if non-nil, zero value otherwise.

### GetSkinsOk

`func (o *GlTF) GetSkinsOk() (*[]Skin, bool)`

GetSkinsOk returns a tuple with the Skins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkins

`func (o *GlTF) SetSkins(v []Skin)`

SetSkins sets Skins field to given value.

### HasSkins

`func (o *GlTF) HasSkins() bool`

HasSkins returns a boolean if a field has been set.

### GetTextures

`func (o *GlTF) GetTextures() []Texture`

GetTextures returns the Textures field if non-nil, zero value otherwise.

### GetTexturesOk

`func (o *GlTF) GetTexturesOk() (*[]Texture, bool)`

GetTexturesOk returns a tuple with the Textures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextures

`func (o *GlTF) SetTextures(v []Texture)`

SetTextures sets Textures field to given value.

### HasTextures

`func (o *GlTF) HasTextures() bool`

HasTextures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


