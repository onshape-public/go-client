# Scene

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewScene

`func NewScene() *Scene`

NewScene instantiates a new Scene object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSceneWithDefaults

`func NewSceneWithDefaults() *Scene`

NewSceneWithDefaults instantiates a new Scene object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *Scene) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Scene) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Scene) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Scene) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Scene) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Scene) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Scene) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Scene) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *Scene) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Scene) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Scene) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Scene) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *Scene) GetNodes() []int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Scene) GetNodesOk() (*[]int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Scene) SetNodes(v []int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Scene) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


