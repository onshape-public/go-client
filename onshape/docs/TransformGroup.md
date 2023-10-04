# TransformGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]BTAssemblyInstanceDefinitionParams**](BTAssemblyInstanceDefinitionParams.md) |  | [optional] 
**Transform** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewTransformGroup

`func NewTransformGroup() *TransformGroup`

NewTransformGroup instantiates a new TransformGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformGroupWithDefaults

`func NewTransformGroupWithDefaults() *TransformGroup`

NewTransformGroupWithDefaults instantiates a new TransformGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *TransformGroup) GetInstances() []BTAssemblyInstanceDefinitionParams`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *TransformGroup) GetInstancesOk() (*[]BTAssemblyInstanceDefinitionParams, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *TransformGroup) SetInstances(v []BTAssemblyInstanceDefinitionParams)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *TransformGroup) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetTransform

`func (o *TransformGroup) GetTransform() []float64`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *TransformGroup) GetTransformOk() (*[]float64, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *TransformGroup) SetTransform(v []float64)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *TransformGroup) HasTransform() bool`

HasTransform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


