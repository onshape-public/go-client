# BTPExpressionFunction1325

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**[]BTPArgumentDeclaration232**](BTPArgumentDeclaration-232.md) |  | [optional] 
**Body** | Pointer to [**BTPStatementBlock271**](BTPStatementBlock-271.md) |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**Precondition** | Pointer to [**BTPStatement269**](BTPStatement-269.md) |  | [optional] 
**ReturnType** | Pointer to [**BTPTypeName290**](BTPTypeName-290.md) |  | [optional] 
**SpaceAfterArglist** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceAfterFunction** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceInEmptyList** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 

## Methods

### NewBTPExpressionFunction1325

`func NewBTPExpressionFunction1325() *BTPExpressionFunction1325`

NewBTPExpressionFunction1325 instantiates a new BTPExpressionFunction1325 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPExpressionFunction1325WithDefaults

`func NewBTPExpressionFunction1325WithDefaults() *BTPExpressionFunction1325`

NewBTPExpressionFunction1325WithDefaults instantiates a new BTPExpressionFunction1325 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *BTPExpressionFunction1325) GetArguments() []BTPArgumentDeclaration232`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *BTPExpressionFunction1325) GetArgumentsOk() (*[]BTPArgumentDeclaration232, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *BTPExpressionFunction1325) SetArguments(v []BTPArgumentDeclaration232)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *BTPExpressionFunction1325) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetBody

`func (o *BTPExpressionFunction1325) GetBody() BTPStatementBlock271`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BTPExpressionFunction1325) GetBodyOk() (*BTPStatementBlock271, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BTPExpressionFunction1325) SetBody(v BTPStatementBlock271)`

SetBody sets Body field to given value.

### HasBody

`func (o *BTPExpressionFunction1325) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBtType

`func (o *BTPExpressionFunction1325) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPExpressionFunction1325) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPExpressionFunction1325) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPExpressionFunction1325) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetPrecondition

`func (o *BTPExpressionFunction1325) GetPrecondition() BTPStatement269`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *BTPExpressionFunction1325) GetPreconditionOk() (*BTPStatement269, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *BTPExpressionFunction1325) SetPrecondition(v BTPStatement269)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *BTPExpressionFunction1325) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetReturnType

`func (o *BTPExpressionFunction1325) GetReturnType() BTPTypeName290`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *BTPExpressionFunction1325) GetReturnTypeOk() (*BTPTypeName290, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *BTPExpressionFunction1325) SetReturnType(v BTPTypeName290)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *BTPExpressionFunction1325) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetSpaceAfterArglist

`func (o *BTPExpressionFunction1325) GetSpaceAfterArglist() BTPSpace10`

GetSpaceAfterArglist returns the SpaceAfterArglist field if non-nil, zero value otherwise.

### GetSpaceAfterArglistOk

`func (o *BTPExpressionFunction1325) GetSpaceAfterArglistOk() (*BTPSpace10, bool)`

GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterArglist

`func (o *BTPExpressionFunction1325) SetSpaceAfterArglist(v BTPSpace10)`

SetSpaceAfterArglist sets SpaceAfterArglist field to given value.

### HasSpaceAfterArglist

`func (o *BTPExpressionFunction1325) HasSpaceAfterArglist() bool`

HasSpaceAfterArglist returns a boolean if a field has been set.

### GetSpaceAfterFunction

`func (o *BTPExpressionFunction1325) GetSpaceAfterFunction() BTPSpace10`

GetSpaceAfterFunction returns the SpaceAfterFunction field if non-nil, zero value otherwise.

### GetSpaceAfterFunctionOk

`func (o *BTPExpressionFunction1325) GetSpaceAfterFunctionOk() (*BTPSpace10, bool)`

GetSpaceAfterFunctionOk returns a tuple with the SpaceAfterFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterFunction

`func (o *BTPExpressionFunction1325) SetSpaceAfterFunction(v BTPSpace10)`

SetSpaceAfterFunction sets SpaceAfterFunction field to given value.

### HasSpaceAfterFunction

`func (o *BTPExpressionFunction1325) HasSpaceAfterFunction() bool`

HasSpaceAfterFunction returns a boolean if a field has been set.

### GetSpaceInEmptyList

`func (o *BTPExpressionFunction1325) GetSpaceInEmptyList() BTPSpace10`

GetSpaceInEmptyList returns the SpaceInEmptyList field if non-nil, zero value otherwise.

### GetSpaceInEmptyListOk

`func (o *BTPExpressionFunction1325) GetSpaceInEmptyListOk() (*BTPSpace10, bool)`

GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceInEmptyList

`func (o *BTPExpressionFunction1325) SetSpaceInEmptyList(v BTPSpace10)`

SetSpaceInEmptyList sets SpaceInEmptyList field to given value.

### HasSpaceInEmptyList

`func (o *BTPExpressionFunction1325) HasSpaceInEmptyList() bool`

HasSpaceInEmptyList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


