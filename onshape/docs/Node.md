# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Camera** | Pointer to **int32** |  | [optional] 
**Children** | Pointer to **[]int32** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Matrix** | Pointer to **[]float32** |  | [optional] 
**Mesh** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rotation** | Pointer to **[]float32** |  | [optional] 
**Scale** | Pointer to **[]float32** |  | [optional] 
**Skin** | Pointer to **int32** |  | [optional] 
**Translation** | Pointer to **[]float32** |  | [optional] 
**Weights** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCamera

`func (o *Node) GetCamera() int32`

GetCamera returns the Camera field if non-nil, zero value otherwise.

### GetCameraOk

`func (o *Node) GetCameraOk() (*int32, bool)`

GetCameraOk returns a tuple with the Camera field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCamera

`func (o *Node) SetCamera(v int32)`

SetCamera sets Camera field to given value.

### HasCamera

`func (o *Node) HasCamera() bool`

HasCamera returns a boolean if a field has been set.

### GetChildren

`func (o *Node) GetChildren() []int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Node) GetChildrenOk() (*[]int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Node) SetChildren(v []int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Node) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetExtensions

`func (o *Node) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Node) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Node) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Node) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Node) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Node) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Node) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Node) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetMatrix

`func (o *Node) GetMatrix() []float32`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *Node) GetMatrixOk() (*[]float32, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *Node) SetMatrix(v []float32)`

SetMatrix sets Matrix field to given value.

### HasMatrix

`func (o *Node) HasMatrix() bool`

HasMatrix returns a boolean if a field has been set.

### GetMesh

`func (o *Node) GetMesh() int32`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *Node) GetMeshOk() (*int32, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *Node) SetMesh(v int32)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *Node) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRotation

`func (o *Node) GetRotation() []float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *Node) GetRotationOk() (*[]float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *Node) SetRotation(v []float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *Node) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetScale

`func (o *Node) GetScale() []float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *Node) GetScaleOk() (*[]float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *Node) SetScale(v []float32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *Node) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetSkin

`func (o *Node) GetSkin() int32`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *Node) GetSkinOk() (*int32, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkin

`func (o *Node) SetSkin(v int32)`

SetSkin sets Skin field to given value.

### HasSkin

`func (o *Node) HasSkin() bool`

HasSkin returns a boolean if a field has been set.

### GetTranslation

`func (o *Node) GetTranslation() []float32`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *Node) GetTranslationOk() (*[]float32, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *Node) SetTranslation(v []float32)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *Node) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### GetWeights

`func (o *Node) GetWeights() []float32`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *Node) GetWeightsOk() (*[]float32, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *Node) SetWeights(v []float32)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *Node) HasWeights() bool`

HasWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


