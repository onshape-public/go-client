# BTAssemblyModificationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteInstances** | Pointer to **[]string** | Node ids of the instances to delete. | [optional] 
**EditDescription** | Pointer to **string** | Short description of the modification. | [optional] 
**SuppressInstances** | Pointer to **[]string** | Deprecated in API v16. Use &#x60;suppressionStates&#x60; instead.Node ids of the instances to suppress. | [optional] 
**SuppressionStates** | Pointer to [**map[string]BTOptionallyConfiguredValue**](BTOptionallyConfiguredValue.md) | Suppression states keyed by node id. Each value is either a plain suppression state (&#x60;value&#x60;: &#x60;\&quot;true\&quot;&#x60; or &#x60;\&quot;false\&quot;&#x60;) or a configuration-controlled suppression state (&#x60;configuredValue&#x60;: a configuration parameter id and a map from configuration option id to &#x60;\&quot;true\&quot;&#x60; or &#x60;\&quot;false\&quot;&#x60;). | [optional] 
**TransformDefinitions** | Pointer to [**[]BTAssemblyTransformDefinitionParams**](BTAssemblyTransformDefinitionParams.md) | Occurrence transform definitions. | [optional] 
**UnsuppressInstances** | Pointer to **[]string** | Deprecated in API v16. Use &#x60;suppressionStates&#x60; instead.Node ids of the instances to unsuppress. | [optional] 

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

### GetSuppressionStates

`func (o *BTAssemblyModificationParams) GetSuppressionStates() map[string]BTOptionallyConfiguredValue`

GetSuppressionStates returns the SuppressionStates field if non-nil, zero value otherwise.

### GetSuppressionStatesOk

`func (o *BTAssemblyModificationParams) GetSuppressionStatesOk() (*map[string]BTOptionallyConfiguredValue, bool)`

GetSuppressionStatesOk returns a tuple with the SuppressionStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionStates

`func (o *BTAssemblyModificationParams) SetSuppressionStates(v map[string]BTOptionallyConfiguredValue)`

SetSuppressionStates sets SuppressionStates field to given value.

### HasSuppressionStates

`func (o *BTAssemblyModificationParams) HasSuppressionStates() bool`

HasSuppressionStates returns a boolean if a field has been set.

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


