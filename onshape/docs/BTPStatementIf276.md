# BTPStatementIf276

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to [**BTPExpression9**](BTPExpression-9.md) |  | [optional] 
**ElseBody** | Pointer to [**BTPStatement269**](BTPStatement-269.md) |  | [optional] 
**SpaceAfterIf** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**ThenBody** | Pointer to [**BTPStatement269**](BTPStatement-269.md) |  | [optional] 

## Methods

### NewBTPStatementIf276

`func NewBTPStatementIf276() *BTPStatementIf276`

NewBTPStatementIf276 instantiates a new BTPStatementIf276 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPStatementIf276WithDefaults

`func NewBTPStatementIf276WithDefaults() *BTPStatementIf276`

NewBTPStatementIf276WithDefaults instantiates a new BTPStatementIf276 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPStatementIf276) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPStatementIf276) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPStatementIf276) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPStatementIf276) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCondition

`func (o *BTPStatementIf276) GetCondition() BTPExpression9`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *BTPStatementIf276) GetConditionOk() (*BTPExpression9, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *BTPStatementIf276) SetCondition(v BTPExpression9)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *BTPStatementIf276) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetElseBody

`func (o *BTPStatementIf276) GetElseBody() BTPStatement269`

GetElseBody returns the ElseBody field if non-nil, zero value otherwise.

### GetElseBodyOk

`func (o *BTPStatementIf276) GetElseBodyOk() (*BTPStatement269, bool)`

GetElseBodyOk returns a tuple with the ElseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseBody

`func (o *BTPStatementIf276) SetElseBody(v BTPStatement269)`

SetElseBody sets ElseBody field to given value.

### HasElseBody

`func (o *BTPStatementIf276) HasElseBody() bool`

HasElseBody returns a boolean if a field has been set.

### GetSpaceAfterIf

`func (o *BTPStatementIf276) GetSpaceAfterIf() BTPSpace10`

GetSpaceAfterIf returns the SpaceAfterIf field if non-nil, zero value otherwise.

### GetSpaceAfterIfOk

`func (o *BTPStatementIf276) GetSpaceAfterIfOk() (*BTPSpace10, bool)`

GetSpaceAfterIfOk returns a tuple with the SpaceAfterIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterIf

`func (o *BTPStatementIf276) SetSpaceAfterIf(v BTPSpace10)`

SetSpaceAfterIf sets SpaceAfterIf field to given value.

### HasSpaceAfterIf

`func (o *BTPStatementIf276) HasSpaceAfterIf() bool`

HasSpaceAfterIf returns a boolean if a field has been set.

### GetThenBody

`func (o *BTPStatementIf276) GetThenBody() BTPStatement269`

GetThenBody returns the ThenBody field if non-nil, zero value otherwise.

### GetThenBodyOk

`func (o *BTPStatementIf276) GetThenBodyOk() (*BTPStatement269, bool)`

GetThenBodyOk returns a tuple with the ThenBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThenBody

`func (o *BTPStatementIf276) SetThenBody(v BTPStatement269)`

SetThenBody sets ThenBody field to given value.

### HasThenBody

`func (o *BTPStatementIf276) HasThenBody() bool`

HasThenBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


