# BTConfigurableTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ParameterLibraries** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**Parameters** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 
**SuppressionState** | Pointer to [**BTMSuppressionState1924**](BTMSuppressionState1924.md) |  | [optional] 

## Methods

### NewBTConfigurableTreeNode

`func NewBTConfigurableTreeNode() *BTConfigurableTreeNode`

NewBTConfigurableTreeNode instantiates a new BTConfigurableTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTConfigurableTreeNodeWithDefaults

`func NewBTConfigurableTreeNodeWithDefaults() *BTConfigurableTreeNode`

NewBTConfigurableTreeNodeWithDefaults instantiates a new BTConfigurableTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BTConfigurableTreeNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTConfigurableTreeNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTConfigurableTreeNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTConfigurableTreeNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameterLibraries

`func (o *BTConfigurableTreeNode) GetParameterLibraries() []BTMParameter1`

GetParameterLibraries returns the ParameterLibraries field if non-nil, zero value otherwise.

### GetParameterLibrariesOk

`func (o *BTConfigurableTreeNode) GetParameterLibrariesOk() (*[]BTMParameter1, bool)`

GetParameterLibrariesOk returns a tuple with the ParameterLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLibraries

`func (o *BTConfigurableTreeNode) SetParameterLibraries(v []BTMParameter1)`

SetParameterLibraries sets ParameterLibraries field to given value.

### HasParameterLibraries

`func (o *BTConfigurableTreeNode) HasParameterLibraries() bool`

HasParameterLibraries returns a boolean if a field has been set.

### GetParameters

`func (o *BTConfigurableTreeNode) GetParameters() []BTMParameter1`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTConfigurableTreeNode) GetParametersOk() (*[]BTMParameter1, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTConfigurableTreeNode) SetParameters(v []BTMParameter1)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTConfigurableTreeNode) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTConfigurableTreeNode) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTConfigurableTreeNode) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTConfigurableTreeNode) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTConfigurableTreeNode) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTConfigurableTreeNode) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTConfigurableTreeNode) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTConfigurableTreeNode) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTConfigurableTreeNode) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


