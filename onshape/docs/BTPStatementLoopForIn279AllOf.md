# BTPStatementLoopForIn279AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**Container** | Pointer to [**BTPExpression9**](BTPExpression9.md) |  | [optional] 
**Identifiers** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**IsVarDeclaredHere** | Pointer to **bool** |  | [optional] 
**KeyVar** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**SpaceBeforeVar** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**StandardTypes** | Pointer to [**[]GBTPType**](GBTPType.md) |  | [optional] 
**TypeNames** | Pointer to **[]string** |  | [optional] 
**Var** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 

## Methods

### NewBTPStatementLoopForIn279AllOf

`func NewBTPStatementLoopForIn279AllOf() *BTPStatementLoopForIn279AllOf`

NewBTPStatementLoopForIn279AllOf instantiates a new BTPStatementLoopForIn279AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPStatementLoopForIn279AllOfWithDefaults

`func NewBTPStatementLoopForIn279AllOfWithDefaults() *BTPStatementLoopForIn279AllOf`

NewBTPStatementLoopForIn279AllOfWithDefaults instantiates a new BTPStatementLoopForIn279AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPStatementLoopForIn279AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPStatementLoopForIn279AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPStatementLoopForIn279AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPStatementLoopForIn279AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetContainer

`func (o *BTPStatementLoopForIn279AllOf) GetContainer() BTPExpression9`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BTPStatementLoopForIn279AllOf) GetContainerOk() (*BTPExpression9, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BTPStatementLoopForIn279AllOf) SetContainer(v BTPExpression9)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BTPStatementLoopForIn279AllOf) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetIdentifiers

`func (o *BTPStatementLoopForIn279AllOf) GetIdentifiers() []BTPIdentifier8`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *BTPStatementLoopForIn279AllOf) GetIdentifiersOk() (*[]BTPIdentifier8, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *BTPStatementLoopForIn279AllOf) SetIdentifiers(v []BTPIdentifier8)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *BTPStatementLoopForIn279AllOf) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetIsVarDeclaredHere

`func (o *BTPStatementLoopForIn279AllOf) GetIsVarDeclaredHere() bool`

GetIsVarDeclaredHere returns the IsVarDeclaredHere field if non-nil, zero value otherwise.

### GetIsVarDeclaredHereOk

`func (o *BTPStatementLoopForIn279AllOf) GetIsVarDeclaredHereOk() (*bool, bool)`

GetIsVarDeclaredHereOk returns a tuple with the IsVarDeclaredHere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVarDeclaredHere

`func (o *BTPStatementLoopForIn279AllOf) SetIsVarDeclaredHere(v bool)`

SetIsVarDeclaredHere sets IsVarDeclaredHere field to given value.

### HasIsVarDeclaredHere

`func (o *BTPStatementLoopForIn279AllOf) HasIsVarDeclaredHere() bool`

HasIsVarDeclaredHere returns a boolean if a field has been set.

### GetKeyVar

`func (o *BTPStatementLoopForIn279AllOf) GetKeyVar() BTPIdentifier8`

GetKeyVar returns the KeyVar field if non-nil, zero value otherwise.

### GetKeyVarOk

`func (o *BTPStatementLoopForIn279AllOf) GetKeyVarOk() (*BTPIdentifier8, bool)`

GetKeyVarOk returns a tuple with the KeyVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVar

`func (o *BTPStatementLoopForIn279AllOf) SetKeyVar(v BTPIdentifier8)`

SetKeyVar sets KeyVar field to given value.

### HasKeyVar

`func (o *BTPStatementLoopForIn279AllOf) HasKeyVar() bool`

HasKeyVar returns a boolean if a field has been set.

### GetSpaceBeforeVar

`func (o *BTPStatementLoopForIn279AllOf) GetSpaceBeforeVar() BTPSpace10`

GetSpaceBeforeVar returns the SpaceBeforeVar field if non-nil, zero value otherwise.

### GetSpaceBeforeVarOk

`func (o *BTPStatementLoopForIn279AllOf) GetSpaceBeforeVarOk() (*BTPSpace10, bool)`

GetSpaceBeforeVarOk returns a tuple with the SpaceBeforeVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeVar

`func (o *BTPStatementLoopForIn279AllOf) SetSpaceBeforeVar(v BTPSpace10)`

SetSpaceBeforeVar sets SpaceBeforeVar field to given value.

### HasSpaceBeforeVar

`func (o *BTPStatementLoopForIn279AllOf) HasSpaceBeforeVar() bool`

HasSpaceBeforeVar returns a boolean if a field has been set.

### GetStandardTypes

`func (o *BTPStatementLoopForIn279AllOf) GetStandardTypes() []GBTPType`

GetStandardTypes returns the StandardTypes field if non-nil, zero value otherwise.

### GetStandardTypesOk

`func (o *BTPStatementLoopForIn279AllOf) GetStandardTypesOk() (*[]GBTPType, bool)`

GetStandardTypesOk returns a tuple with the StandardTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardTypes

`func (o *BTPStatementLoopForIn279AllOf) SetStandardTypes(v []GBTPType)`

SetStandardTypes sets StandardTypes field to given value.

### HasStandardTypes

`func (o *BTPStatementLoopForIn279AllOf) HasStandardTypes() bool`

HasStandardTypes returns a boolean if a field has been set.

### GetTypeNames

`func (o *BTPStatementLoopForIn279AllOf) GetTypeNames() []string`

GetTypeNames returns the TypeNames field if non-nil, zero value otherwise.

### GetTypeNamesOk

`func (o *BTPStatementLoopForIn279AllOf) GetTypeNamesOk() (*[]string, bool)`

GetTypeNamesOk returns a tuple with the TypeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNames

`func (o *BTPStatementLoopForIn279AllOf) SetTypeNames(v []string)`

SetTypeNames sets TypeNames field to given value.

### HasTypeNames

`func (o *BTPStatementLoopForIn279AllOf) HasTypeNames() bool`

HasTypeNames returns a boolean if a field has been set.

### GetVar

`func (o *BTPStatementLoopForIn279AllOf) GetVar() BTPIdentifier8`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *BTPStatementLoopForIn279AllOf) GetVarOk() (*BTPIdentifier8, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *BTPStatementLoopForIn279AllOf) SetVar(v BTPIdentifier8)`

SetVar sets Var field to given value.

### HasVar

`func (o *BTPStatementLoopForIn279AllOf) HasVar() bool`

HasVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


