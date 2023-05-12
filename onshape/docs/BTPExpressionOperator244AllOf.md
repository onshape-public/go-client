# BTPExpressionOperator244AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**ForExport** | Pointer to **bool** |  | [optional] 
**GlobalNamespace** | Pointer to **bool** |  | [optional] 
**ImportMicroversion** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**Operand1** | Pointer to [**BTPExpression9**](BTPExpression9.md) |  | [optional] 
**Operand2** | Pointer to [**BTPExpression9**](BTPExpression9.md) |  | [optional] 
**Operand3** | Pointer to [**BTPExpression9**](BTPExpression9.md) |  | [optional] 
**Operator** | Pointer to [**GBTPOperator**](GBTPOperator.md) |  | [optional] 
**SpaceAfterNamespace** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceAfterOperator** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBeforeOperator** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**WrittenAsFunctionCall** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTPExpressionOperator244AllOf

`func NewBTPExpressionOperator244AllOf() *BTPExpressionOperator244AllOf`

NewBTPExpressionOperator244AllOf instantiates a new BTPExpressionOperator244AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPExpressionOperator244AllOfWithDefaults

`func NewBTPExpressionOperator244AllOfWithDefaults() *BTPExpressionOperator244AllOf`

NewBTPExpressionOperator244AllOfWithDefaults instantiates a new BTPExpressionOperator244AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPExpressionOperator244AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPExpressionOperator244AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPExpressionOperator244AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPExpressionOperator244AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetForExport

`func (o *BTPExpressionOperator244AllOf) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *BTPExpressionOperator244AllOf) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *BTPExpressionOperator244AllOf) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *BTPExpressionOperator244AllOf) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetGlobalNamespace

`func (o *BTPExpressionOperator244AllOf) GetGlobalNamespace() bool`

GetGlobalNamespace returns the GlobalNamespace field if non-nil, zero value otherwise.

### GetGlobalNamespaceOk

`func (o *BTPExpressionOperator244AllOf) GetGlobalNamespaceOk() (*bool, bool)`

GetGlobalNamespaceOk returns a tuple with the GlobalNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalNamespace

`func (o *BTPExpressionOperator244AllOf) SetGlobalNamespace(v bool)`

SetGlobalNamespace sets GlobalNamespace field to given value.

### HasGlobalNamespace

`func (o *BTPExpressionOperator244AllOf) HasGlobalNamespace() bool`

HasGlobalNamespace returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTPExpressionOperator244AllOf) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTPExpressionOperator244AllOf) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTPExpressionOperator244AllOf) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTPExpressionOperator244AllOf) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetNamespace

`func (o *BTPExpressionOperator244AllOf) GetNamespace() []BTPIdentifier8`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTPExpressionOperator244AllOf) GetNamespaceOk() (*[]BTPIdentifier8, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTPExpressionOperator244AllOf) SetNamespace(v []BTPIdentifier8)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTPExpressionOperator244AllOf) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperand1

`func (o *BTPExpressionOperator244AllOf) GetOperand1() BTPExpression9`

GetOperand1 returns the Operand1 field if non-nil, zero value otherwise.

### GetOperand1Ok

`func (o *BTPExpressionOperator244AllOf) GetOperand1Ok() (*BTPExpression9, bool)`

GetOperand1Ok returns a tuple with the Operand1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand1

`func (o *BTPExpressionOperator244AllOf) SetOperand1(v BTPExpression9)`

SetOperand1 sets Operand1 field to given value.

### HasOperand1

`func (o *BTPExpressionOperator244AllOf) HasOperand1() bool`

HasOperand1 returns a boolean if a field has been set.

### GetOperand2

`func (o *BTPExpressionOperator244AllOf) GetOperand2() BTPExpression9`

GetOperand2 returns the Operand2 field if non-nil, zero value otherwise.

### GetOperand2Ok

`func (o *BTPExpressionOperator244AllOf) GetOperand2Ok() (*BTPExpression9, bool)`

GetOperand2Ok returns a tuple with the Operand2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand2

`func (o *BTPExpressionOperator244AllOf) SetOperand2(v BTPExpression9)`

SetOperand2 sets Operand2 field to given value.

### HasOperand2

`func (o *BTPExpressionOperator244AllOf) HasOperand2() bool`

HasOperand2 returns a boolean if a field has been set.

### GetOperand3

`func (o *BTPExpressionOperator244AllOf) GetOperand3() BTPExpression9`

GetOperand3 returns the Operand3 field if non-nil, zero value otherwise.

### GetOperand3Ok

`func (o *BTPExpressionOperator244AllOf) GetOperand3Ok() (*BTPExpression9, bool)`

GetOperand3Ok returns a tuple with the Operand3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand3

`func (o *BTPExpressionOperator244AllOf) SetOperand3(v BTPExpression9)`

SetOperand3 sets Operand3 field to given value.

### HasOperand3

`func (o *BTPExpressionOperator244AllOf) HasOperand3() bool`

HasOperand3 returns a boolean if a field has been set.

### GetOperator

`func (o *BTPExpressionOperator244AllOf) GetOperator() GBTPOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BTPExpressionOperator244AllOf) GetOperatorOk() (*GBTPOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BTPExpressionOperator244AllOf) SetOperator(v GBTPOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *BTPExpressionOperator244AllOf) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetSpaceAfterNamespace

`func (o *BTPExpressionOperator244AllOf) GetSpaceAfterNamespace() BTPSpace10`

GetSpaceAfterNamespace returns the SpaceAfterNamespace field if non-nil, zero value otherwise.

### GetSpaceAfterNamespaceOk

`func (o *BTPExpressionOperator244AllOf) GetSpaceAfterNamespaceOk() (*BTPSpace10, bool)`

GetSpaceAfterNamespaceOk returns a tuple with the SpaceAfterNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterNamespace

`func (o *BTPExpressionOperator244AllOf) SetSpaceAfterNamespace(v BTPSpace10)`

SetSpaceAfterNamespace sets SpaceAfterNamespace field to given value.

### HasSpaceAfterNamespace

`func (o *BTPExpressionOperator244AllOf) HasSpaceAfterNamespace() bool`

HasSpaceAfterNamespace returns a boolean if a field has been set.

### GetSpaceAfterOperator

`func (o *BTPExpressionOperator244AllOf) GetSpaceAfterOperator() BTPSpace10`

GetSpaceAfterOperator returns the SpaceAfterOperator field if non-nil, zero value otherwise.

### GetSpaceAfterOperatorOk

`func (o *BTPExpressionOperator244AllOf) GetSpaceAfterOperatorOk() (*BTPSpace10, bool)`

GetSpaceAfterOperatorOk returns a tuple with the SpaceAfterOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterOperator

`func (o *BTPExpressionOperator244AllOf) SetSpaceAfterOperator(v BTPSpace10)`

SetSpaceAfterOperator sets SpaceAfterOperator field to given value.

### HasSpaceAfterOperator

`func (o *BTPExpressionOperator244AllOf) HasSpaceAfterOperator() bool`

HasSpaceAfterOperator returns a boolean if a field has been set.

### GetSpaceBeforeOperator

`func (o *BTPExpressionOperator244AllOf) GetSpaceBeforeOperator() BTPSpace10`

GetSpaceBeforeOperator returns the SpaceBeforeOperator field if non-nil, zero value otherwise.

### GetSpaceBeforeOperatorOk

`func (o *BTPExpressionOperator244AllOf) GetSpaceBeforeOperatorOk() (*BTPSpace10, bool)`

GetSpaceBeforeOperatorOk returns a tuple with the SpaceBeforeOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeOperator

`func (o *BTPExpressionOperator244AllOf) SetSpaceBeforeOperator(v BTPSpace10)`

SetSpaceBeforeOperator sets SpaceBeforeOperator field to given value.

### HasSpaceBeforeOperator

`func (o *BTPExpressionOperator244AllOf) HasSpaceBeforeOperator() bool`

HasSpaceBeforeOperator returns a boolean if a field has been set.

### GetWrittenAsFunctionCall

`func (o *BTPExpressionOperator244AllOf) GetWrittenAsFunctionCall() bool`

GetWrittenAsFunctionCall returns the WrittenAsFunctionCall field if non-nil, zero value otherwise.

### GetWrittenAsFunctionCallOk

`func (o *BTPExpressionOperator244AllOf) GetWrittenAsFunctionCallOk() (*bool, bool)`

GetWrittenAsFunctionCallOk returns a tuple with the WrittenAsFunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenAsFunctionCall

`func (o *BTPExpressionOperator244AllOf) SetWrittenAsFunctionCall(v bool)`

SetWrittenAsFunctionCall sets WrittenAsFunctionCall field to given value.

### HasWrittenAsFunctionCall

`func (o *BTPExpressionOperator244AllOf) HasWrittenAsFunctionCall() bool`

HasWrittenAsFunctionCall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


