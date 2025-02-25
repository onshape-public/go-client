# BTConfigurableTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


