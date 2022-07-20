# BTPStatementLoopForIn279

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**Container** | Pointer to [**BTPExpression9**](BTPExpression9.md) |  | [optional] 
**Identifiers** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**IsVarDeclaredHere** | Pointer to **bool** |  | [optional] 
**KeyVar** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**SpaceBeforeVar** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**StandardTypes** | Pointer to **[]string** |  | [optional] 
**TypeNames** | Pointer to **[]string** |  | [optional] 
**Var** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 

## Methods

### NewBTPStatementLoopForIn279

`func NewBTPStatementLoopForIn279() *BTPStatementLoopForIn279`

NewBTPStatementLoopForIn279 instantiates a new BTPStatementLoopForIn279 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPStatementLoopForIn279WithDefaults

`func NewBTPStatementLoopForIn279WithDefaults() *BTPStatementLoopForIn279`

NewBTPStatementLoopForIn279WithDefaults instantiates a new BTPStatementLoopForIn279 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPStatementLoopForIn279) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPStatementLoopForIn279) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPStatementLoopForIn279) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPStatementLoopForIn279) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetContainer

`func (o *BTPStatementLoopForIn279) GetContainer() BTPExpression9`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BTPStatementLoopForIn279) GetContainerOk() (*BTPExpression9, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BTPStatementLoopForIn279) SetContainer(v BTPExpression9)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BTPStatementLoopForIn279) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetIdentifiers

`func (o *BTPStatementLoopForIn279) GetIdentifiers() []BTPIdentifier8`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *BTPStatementLoopForIn279) GetIdentifiersOk() (*[]BTPIdentifier8, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *BTPStatementLoopForIn279) SetIdentifiers(v []BTPIdentifier8)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *BTPStatementLoopForIn279) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetIsVarDeclaredHere

`func (o *BTPStatementLoopForIn279) GetIsVarDeclaredHere() bool`

GetIsVarDeclaredHere returns the IsVarDeclaredHere field if non-nil, zero value otherwise.

### GetIsVarDeclaredHereOk

`func (o *BTPStatementLoopForIn279) GetIsVarDeclaredHereOk() (*bool, bool)`

GetIsVarDeclaredHereOk returns a tuple with the IsVarDeclaredHere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVarDeclaredHere

`func (o *BTPStatementLoopForIn279) SetIsVarDeclaredHere(v bool)`

SetIsVarDeclaredHere sets IsVarDeclaredHere field to given value.

### HasIsVarDeclaredHere

`func (o *BTPStatementLoopForIn279) HasIsVarDeclaredHere() bool`

HasIsVarDeclaredHere returns a boolean if a field has been set.

### GetKeyVar

`func (o *BTPStatementLoopForIn279) GetKeyVar() BTPIdentifier8`

GetKeyVar returns the KeyVar field if non-nil, zero value otherwise.

### GetKeyVarOk

`func (o *BTPStatementLoopForIn279) GetKeyVarOk() (*BTPIdentifier8, bool)`

GetKeyVarOk returns a tuple with the KeyVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVar

`func (o *BTPStatementLoopForIn279) SetKeyVar(v BTPIdentifier8)`

SetKeyVar sets KeyVar field to given value.

### HasKeyVar

`func (o *BTPStatementLoopForIn279) HasKeyVar() bool`

HasKeyVar returns a boolean if a field has been set.

### GetSpaceBeforeVar

`func (o *BTPStatementLoopForIn279) GetSpaceBeforeVar() BTPSpace10`

GetSpaceBeforeVar returns the SpaceBeforeVar field if non-nil, zero value otherwise.

### GetSpaceBeforeVarOk

`func (o *BTPStatementLoopForIn279) GetSpaceBeforeVarOk() (*BTPSpace10, bool)`

GetSpaceBeforeVarOk returns a tuple with the SpaceBeforeVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeVar

`func (o *BTPStatementLoopForIn279) SetSpaceBeforeVar(v BTPSpace10)`

SetSpaceBeforeVar sets SpaceBeforeVar field to given value.

### HasSpaceBeforeVar

`func (o *BTPStatementLoopForIn279) HasSpaceBeforeVar() bool`

HasSpaceBeforeVar returns a boolean if a field has been set.

### GetStandardTypes

`func (o *BTPStatementLoopForIn279) GetStandardTypes() []string`

GetStandardTypes returns the StandardTypes field if non-nil, zero value otherwise.

### GetStandardTypesOk

`func (o *BTPStatementLoopForIn279) GetStandardTypesOk() (*[]string, bool)`

GetStandardTypesOk returns a tuple with the StandardTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardTypes

`func (o *BTPStatementLoopForIn279) SetStandardTypes(v []string)`

SetStandardTypes sets StandardTypes field to given value.

### HasStandardTypes

`func (o *BTPStatementLoopForIn279) HasStandardTypes() bool`

HasStandardTypes returns a boolean if a field has been set.

### GetTypeNames

`func (o *BTPStatementLoopForIn279) GetTypeNames() []string`

GetTypeNames returns the TypeNames field if non-nil, zero value otherwise.

### GetTypeNamesOk

`func (o *BTPStatementLoopForIn279) GetTypeNamesOk() (*[]string, bool)`

GetTypeNamesOk returns a tuple with the TypeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNames

`func (o *BTPStatementLoopForIn279) SetTypeNames(v []string)`

SetTypeNames sets TypeNames field to given value.

### HasTypeNames

`func (o *BTPStatementLoopForIn279) HasTypeNames() bool`

HasTypeNames returns a boolean if a field has been set.

### GetVar

`func (o *BTPStatementLoopForIn279) GetVar() BTPIdentifier8`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *BTPStatementLoopForIn279) GetVarOk() (*BTPIdentifier8, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *BTPStatementLoopForIn279) SetVar(v BTPIdentifier8)`

SetVar sets Var field to given value.

### HasVar

`func (o *BTPStatementLoopForIn279) HasVar() bool`

HasVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


