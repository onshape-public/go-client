# BTVariableTableInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableStudioReference** | Pointer to [**BTVariableStudioReferenceInfo**](BTVariableStudioReferenceInfo.md) |  | [optional] 
**Variables** | [**[]BTVariableInfo**](BTVariableInfo.md) | Variables in the VariableTable | 

## Methods

### NewBTVariableTableInfo

`func NewBTVariableTableInfo(variables []BTVariableInfo, ) *BTVariableTableInfo`

NewBTVariableTableInfo instantiates a new BTVariableTableInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVariableTableInfoWithDefaults

`func NewBTVariableTableInfoWithDefaults() *BTVariableTableInfo`

NewBTVariableTableInfoWithDefaults instantiates a new BTVariableTableInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariableStudioReference

`func (o *BTVariableTableInfo) GetVariableStudioReference() BTVariableStudioReferenceInfo`

GetVariableStudioReference returns the VariableStudioReference field if non-nil, zero value otherwise.

### GetVariableStudioReferenceOk

`func (o *BTVariableTableInfo) GetVariableStudioReferenceOk() (*BTVariableStudioReferenceInfo, bool)`

GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudioReference

`func (o *BTVariableTableInfo) SetVariableStudioReference(v BTVariableStudioReferenceInfo)`

SetVariableStudioReference sets VariableStudioReference field to given value.

### HasVariableStudioReference

`func (o *BTVariableTableInfo) HasVariableStudioReference() bool`

HasVariableStudioReference returns a boolean if a field has been set.

### GetVariables

`func (o *BTVariableTableInfo) GetVariables() []BTVariableInfo`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *BTVariableTableInfo) GetVariablesOk() (*[]BTVariableInfo, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *BTVariableTableInfo) SetVariables(v []BTVariableInfo)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


