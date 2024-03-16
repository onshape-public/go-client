# BTPStatementTry1523

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Body** | Pointer to [**BTPStatementBlock271**](BTPStatementBlock271.md) |  | [optional] 
**CatchBlock** | Pointer to [**BTPStatementBlock271**](BTPStatementBlock271.md) |  | [optional] 
**CatchVariable** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**Identifier** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**Silent** | Pointer to **bool** |  | [optional] 
**SpaceAfterCatch** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBeforeSilent** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**StandardType** | Pointer to [**GBTPType**](GBTPType.md) |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPStatementTry1523

`func NewBTPStatementTry1523() *BTPStatementTry1523`

NewBTPStatementTry1523 instantiates a new BTPStatementTry1523 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPStatementTry1523WithDefaults

`func NewBTPStatementTry1523WithDefaults() *BTPStatementTry1523`

NewBTPStatementTry1523WithDefaults instantiates a new BTPStatementTry1523 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPStatementTry1523) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPStatementTry1523) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPStatementTry1523) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPStatementTry1523) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetBody

`func (o *BTPStatementTry1523) GetBody() BTPStatementBlock271`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BTPStatementTry1523) GetBodyOk() (*BTPStatementBlock271, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BTPStatementTry1523) SetBody(v BTPStatementBlock271)`

SetBody sets Body field to given value.

### HasBody

`func (o *BTPStatementTry1523) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCatchBlock

`func (o *BTPStatementTry1523) GetCatchBlock() BTPStatementBlock271`

GetCatchBlock returns the CatchBlock field if non-nil, zero value otherwise.

### GetCatchBlockOk

`func (o *BTPStatementTry1523) GetCatchBlockOk() (*BTPStatementBlock271, bool)`

GetCatchBlockOk returns a tuple with the CatchBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchBlock

`func (o *BTPStatementTry1523) SetCatchBlock(v BTPStatementBlock271)`

SetCatchBlock sets CatchBlock field to given value.

### HasCatchBlock

`func (o *BTPStatementTry1523) HasCatchBlock() bool`

HasCatchBlock returns a boolean if a field has been set.

### GetCatchVariable

`func (o *BTPStatementTry1523) GetCatchVariable() BTPIdentifier8`

GetCatchVariable returns the CatchVariable field if non-nil, zero value otherwise.

### GetCatchVariableOk

`func (o *BTPStatementTry1523) GetCatchVariableOk() (*BTPIdentifier8, bool)`

GetCatchVariableOk returns a tuple with the CatchVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchVariable

`func (o *BTPStatementTry1523) SetCatchVariable(v BTPIdentifier8)`

SetCatchVariable sets CatchVariable field to given value.

### HasCatchVariable

`func (o *BTPStatementTry1523) HasCatchVariable() bool`

HasCatchVariable returns a boolean if a field has been set.

### GetIdentifier

`func (o *BTPStatementTry1523) GetIdentifier() BTPIdentifier8`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BTPStatementTry1523) GetIdentifierOk() (*BTPIdentifier8, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BTPStatementTry1523) SetIdentifier(v BTPIdentifier8)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BTPStatementTry1523) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSilent

`func (o *BTPStatementTry1523) GetSilent() bool`

GetSilent returns the Silent field if non-nil, zero value otherwise.

### GetSilentOk

`func (o *BTPStatementTry1523) GetSilentOk() (*bool, bool)`

GetSilentOk returns a tuple with the Silent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilent

`func (o *BTPStatementTry1523) SetSilent(v bool)`

SetSilent sets Silent field to given value.

### HasSilent

`func (o *BTPStatementTry1523) HasSilent() bool`

HasSilent returns a boolean if a field has been set.

### GetSpaceAfterCatch

`func (o *BTPStatementTry1523) GetSpaceAfterCatch() BTPSpace10`

GetSpaceAfterCatch returns the SpaceAfterCatch field if non-nil, zero value otherwise.

### GetSpaceAfterCatchOk

`func (o *BTPStatementTry1523) GetSpaceAfterCatchOk() (*BTPSpace10, bool)`

GetSpaceAfterCatchOk returns a tuple with the SpaceAfterCatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterCatch

`func (o *BTPStatementTry1523) SetSpaceAfterCatch(v BTPSpace10)`

SetSpaceAfterCatch sets SpaceAfterCatch field to given value.

### HasSpaceAfterCatch

`func (o *BTPStatementTry1523) HasSpaceAfterCatch() bool`

HasSpaceAfterCatch returns a boolean if a field has been set.

### GetSpaceBeforeSilent

`func (o *BTPStatementTry1523) GetSpaceBeforeSilent() BTPSpace10`

GetSpaceBeforeSilent returns the SpaceBeforeSilent field if non-nil, zero value otherwise.

### GetSpaceBeforeSilentOk

`func (o *BTPStatementTry1523) GetSpaceBeforeSilentOk() (*BTPSpace10, bool)`

GetSpaceBeforeSilentOk returns a tuple with the SpaceBeforeSilent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeSilent

`func (o *BTPStatementTry1523) SetSpaceBeforeSilent(v BTPSpace10)`

SetSpaceBeforeSilent sets SpaceBeforeSilent field to given value.

### HasSpaceBeforeSilent

`func (o *BTPStatementTry1523) HasSpaceBeforeSilent() bool`

HasSpaceBeforeSilent returns a boolean if a field has been set.

### GetStandardType

`func (o *BTPStatementTry1523) GetStandardType() GBTPType`

GetStandardType returns the StandardType field if non-nil, zero value otherwise.

### GetStandardTypeOk

`func (o *BTPStatementTry1523) GetStandardTypeOk() (*GBTPType, bool)`

GetStandardTypeOk returns a tuple with the StandardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardType

`func (o *BTPStatementTry1523) SetStandardType(v GBTPType)`

SetStandardType sets StandardType field to given value.

### HasStandardType

`func (o *BTPStatementTry1523) HasStandardType() bool`

HasStandardType returns a boolean if a field has been set.

### GetTypeName

`func (o *BTPStatementTry1523) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *BTPStatementTry1523) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *BTPStatementTry1523) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *BTPStatementTry1523) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


