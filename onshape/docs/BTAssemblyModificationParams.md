# BTAssemblyModificationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteInstances** | Pointer to **[]string** |  | [optional] 
**EditDescription** | Pointer to **string** |  | [optional] 
**SuppressInstances** | Pointer to **[]string** |  | [optional] 
**TransformDefinitions** | Pointer to [**[]BTAssemblyTransformDefinitionParams**](BTAssemblyTransformDefinitionParams.md) |  | [optional] 
**UnsuppressInstances** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBTAssemblyModificationParams

`func NewBTAssemblyModificationParams() *BTAssemblyModificationParams`

NewBTAssemblyModificationParams instantiates a new BTAssemblyModificationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyModificationParamsWithDefaults

`func NewBTAssemblyModificationParamsWithDefaults() *BTAssemblyModificationParams`

NewBTAssemblyModificationParamsWithDefaults instantiates a new BTAssemblyModificationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteInstances

`func (o *BTAssemblyModificationParams) GetDeleteInstances() []string`

GetDeleteInstances returns the DeleteInstances field if non-nil, zero value otherwise.

### GetDeleteInstancesOk

`func (o *BTAssemblyModificationParams) GetDeleteInstancesOk() (*[]string, bool)`

GetDeleteInstancesOk returns a tuple with the DeleteInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInstances

`func (o *BTAssemblyModificationParams) SetDeleteInstances(v []string)`

SetDeleteInstances sets DeleteInstances field to given value.

### HasDeleteInstances

`func (o *BTAssemblyModificationParams) HasDeleteInstances() bool`

HasDeleteInstances returns a boolean if a field has been set.

### GetEditDescription

`func (o *BTAssemblyModificationParams) GetEditDescription() string`

GetEditDescription returns the EditDescription field if non-nil, zero value otherwise.

### GetEditDescriptionOk

`func (o *BTAssemblyModificationParams) GetEditDescriptionOk() (*string, bool)`

GetEditDescriptionOk returns a tuple with the EditDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditDescription

`func (o *BTAssemblyModificationParams) SetEditDescription(v string)`

SetEditDescription sets EditDescription field to given value.

### HasEditDescription

`func (o *BTAssemblyModificationParams) HasEditDescription() bool`

HasEditDescription returns a boolean if a field has been set.

### GetSuppressInstances

`func (o *BTAssemblyModificationParams) GetSuppressInstances() []string`

GetSuppressInstances returns the SuppressInstances field if non-nil, zero value otherwise.

### GetSuppressInstancesOk

`func (o *BTAssemblyModificationParams) GetSuppressInstancesOk() (*[]string, bool)`

GetSuppressInstancesOk returns a tuple with the SuppressInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressInstances

`func (o *BTAssemblyModificationParams) SetSuppressInstances(v []string)`

SetSuppressInstances sets SuppressInstances field to given value.

### HasSuppressInstances

`func (o *BTAssemblyModificationParams) HasSuppressInstances() bool`

HasSuppressInstances returns a boolean if a field has been set.

### GetTransformDefinitions

`func (o *BTAssemblyModificationParams) GetTransformDefinitions() []BTAssemblyTransformDefinitionParams`

GetTransformDefinitions returns the TransformDefinitions field if non-nil, zero value otherwise.

### GetTransformDefinitionsOk

`func (o *BTAssemblyModificationParams) GetTransformDefinitionsOk() (*[]BTAssemblyTransformDefinitionParams, bool)`

GetTransformDefinitionsOk returns a tuple with the TransformDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformDefinitions

`func (o *BTAssemblyModificationParams) SetTransformDefinitions(v []BTAssemblyTransformDefinitionParams)`

SetTransformDefinitions sets TransformDefinitions field to given value.

### HasTransformDefinitions

`func (o *BTAssemblyModificationParams) HasTransformDefinitions() bool`

HasTransformDefinitions returns a boolean if a field has been set.

### GetUnsuppressInstances

`func (o *BTAssemblyModificationParams) GetUnsuppressInstances() []string`

GetUnsuppressInstances returns the UnsuppressInstances field if non-nil, zero value otherwise.

### GetUnsuppressInstancesOk

`func (o *BTAssemblyModificationParams) GetUnsuppressInstancesOk() (*[]string, bool)`

GetUnsuppressInstancesOk returns a tuple with the UnsuppressInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsuppressInstances

`func (o *BTAssemblyModificationParams) SetUnsuppressInstances(v []string)`

SetUnsuppressInstances sets UnsuppressInstances field to given value.

### HasUnsuppressInstances

`func (o *BTAssemblyModificationParams) HasUnsuppressInstances() bool`

HasUnsuppressInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


