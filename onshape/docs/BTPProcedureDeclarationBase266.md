# BTPProcedureDeclarationBase266

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**[]BTPArgumentDeclaration232**](BTPArgumentDeclaration232.md) |  | [optional] 
**Body** | Pointer to [**BTPStatementBlock271**](BTPStatementBlock271.md) |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**Precondition** | Pointer to [**BTPStatement269**](BTPStatement269.md) |  | [optional] 
**ReturnType** | Pointer to [**BTPTypeName290**](BTPTypeName290.md) |  | [optional] 
**SpaceAfterArglist** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceInEmptyList** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 

## Methods

### NewBTPProcedureDeclarationBase266

`func NewBTPProcedureDeclarationBase266() *BTPProcedureDeclarationBase266`

NewBTPProcedureDeclarationBase266 instantiates a new BTPProcedureDeclarationBase266 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPProcedureDeclarationBase266WithDefaults

`func NewBTPProcedureDeclarationBase266WithDefaults() *BTPProcedureDeclarationBase266`

NewBTPProcedureDeclarationBase266WithDefaults instantiates a new BTPProcedureDeclarationBase266 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *BTPProcedureDeclarationBase266) GetArguments() []BTPArgumentDeclaration232`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *BTPProcedureDeclarationBase266) GetArgumentsOk() (*[]BTPArgumentDeclaration232, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *BTPProcedureDeclarationBase266) SetArguments(v []BTPArgumentDeclaration232)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *BTPProcedureDeclarationBase266) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetBody

`func (o *BTPProcedureDeclarationBase266) GetBody() BTPStatementBlock271`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BTPProcedureDeclarationBase266) GetBodyOk() (*BTPStatementBlock271, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BTPProcedureDeclarationBase266) SetBody(v BTPStatementBlock271)`

SetBody sets Body field to given value.

### HasBody

`func (o *BTPProcedureDeclarationBase266) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBtType

`func (o *BTPProcedureDeclarationBase266) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPProcedureDeclarationBase266) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPProcedureDeclarationBase266) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPProcedureDeclarationBase266) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetPrecondition

`func (o *BTPProcedureDeclarationBase266) GetPrecondition() BTPStatement269`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *BTPProcedureDeclarationBase266) GetPreconditionOk() (*BTPStatement269, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *BTPProcedureDeclarationBase266) SetPrecondition(v BTPStatement269)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *BTPProcedureDeclarationBase266) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetReturnType

`func (o *BTPProcedureDeclarationBase266) GetReturnType() BTPTypeName290`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *BTPProcedureDeclarationBase266) GetReturnTypeOk() (*BTPTypeName290, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *BTPProcedureDeclarationBase266) SetReturnType(v BTPTypeName290)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *BTPProcedureDeclarationBase266) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetSpaceAfterArglist

`func (o *BTPProcedureDeclarationBase266) GetSpaceAfterArglist() BTPSpace10`

GetSpaceAfterArglist returns the SpaceAfterArglist field if non-nil, zero value otherwise.

### GetSpaceAfterArglistOk

`func (o *BTPProcedureDeclarationBase266) GetSpaceAfterArglistOk() (*BTPSpace10, bool)`

GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterArglist

`func (o *BTPProcedureDeclarationBase266) SetSpaceAfterArglist(v BTPSpace10)`

SetSpaceAfterArglist sets SpaceAfterArglist field to given value.

### HasSpaceAfterArglist

`func (o *BTPProcedureDeclarationBase266) HasSpaceAfterArglist() bool`

HasSpaceAfterArglist returns a boolean if a field has been set.

### GetSpaceInEmptyList

`func (o *BTPProcedureDeclarationBase266) GetSpaceInEmptyList() BTPSpace10`

GetSpaceInEmptyList returns the SpaceInEmptyList field if non-nil, zero value otherwise.

### GetSpaceInEmptyListOk

`func (o *BTPProcedureDeclarationBase266) GetSpaceInEmptyListOk() (*BTPSpace10, bool)`

GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceInEmptyList

`func (o *BTPProcedureDeclarationBase266) SetSpaceInEmptyList(v BTPSpace10)`

SetSpaceInEmptyList sets SpaceInEmptyList field to given value.

### HasSpaceInEmptyList

`func (o *BTPProcedureDeclarationBase266) HasSpaceInEmptyList() bool`

HasSpaceInEmptyList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


