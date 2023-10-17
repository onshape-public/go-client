# Mesh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Primitives** | Pointer to [**[]MeshPrimitive**](MeshPrimitive.md) |  | [optional] 
**Weights** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewMesh

`func NewMesh() *Mesh`

NewMesh instantiates a new Mesh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeshWithDefaults

`func NewMeshWithDefaults() *Mesh`

NewMeshWithDefaults instantiates a new Mesh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *Mesh) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Mesh) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Mesh) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Mesh) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Mesh) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Mesh) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Mesh) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Mesh) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *Mesh) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mesh) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mesh) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mesh) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimitives

`func (o *Mesh) GetPrimitives() []MeshPrimitive`

GetPrimitives returns the Primitives field if non-nil, zero value otherwise.

### GetPrimitivesOk

`func (o *Mesh) GetPrimitivesOk() (*[]MeshPrimitive, bool)`

GetPrimitivesOk returns a tuple with the Primitives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimitives

`func (o *Mesh) SetPrimitives(v []MeshPrimitive)`

SetPrimitives sets Primitives field to given value.

### HasPrimitives

`func (o *Mesh) HasPrimitives() bool`

HasPrimitives returns a boolean if a field has been set.

### GetWeights

`func (o *Mesh) GetWeights() []float32`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *Mesh) GetWeightsOk() (*[]float32, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *Mesh) SetWeights(v []float32)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *Mesh) HasWeights() bool`

HasWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


