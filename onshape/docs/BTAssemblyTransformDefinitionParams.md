# BTAssemblyTransformDefinitionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRelative** | Pointer to **bool** |  | [optional] 
**Occurrences** | Pointer to [**[]BTOccurrence74**](BTOccurrence74.md) |  | [optional] 
**Transform** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewBTAssemblyTransformDefinitionParams

`func NewBTAssemblyTransformDefinitionParams() *BTAssemblyTransformDefinitionParams`

NewBTAssemblyTransformDefinitionParams instantiates a new BTAssemblyTransformDefinitionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyTransformDefinitionParamsWithDefaults

`func NewBTAssemblyTransformDefinitionParamsWithDefaults() *BTAssemblyTransformDefinitionParams`

NewBTAssemblyTransformDefinitionParamsWithDefaults instantiates a new BTAssemblyTransformDefinitionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRelative

`func (o *BTAssemblyTransformDefinitionParams) GetIsRelative() bool`

GetIsRelative returns the IsRelative field if non-nil, zero value otherwise.

### GetIsRelativeOk

`func (o *BTAssemblyTransformDefinitionParams) GetIsRelativeOk() (*bool, bool)`

GetIsRelativeOk returns a tuple with the IsRelative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRelative

`func (o *BTAssemblyTransformDefinitionParams) SetIsRelative(v bool)`

SetIsRelative sets IsRelative field to given value.

### HasIsRelative

`func (o *BTAssemblyTransformDefinitionParams) HasIsRelative() bool`

HasIsRelative returns a boolean if a field has been set.

### GetOccurrences

`func (o *BTAssemblyTransformDefinitionParams) GetOccurrences() []BTOccurrence74`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *BTAssemblyTransformDefinitionParams) GetOccurrencesOk() (*[]BTOccurrence74, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *BTAssemblyTransformDefinitionParams) SetOccurrences(v []BTOccurrence74)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *BTAssemblyTransformDefinitionParams) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.

### GetTransform

`func (o *BTAssemblyTransformDefinitionParams) GetTransform() []float64`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *BTAssemblyTransformDefinitionParams) GetTransformOk() (*[]float64, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *BTAssemblyTransformDefinitionParams) SetTransform(v []float64)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *BTAssemblyTransformDefinitionParams) HasTransform() bool`

HasTransform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


