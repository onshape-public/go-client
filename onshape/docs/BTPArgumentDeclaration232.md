# BTPArgumentDeclaration232

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atomic** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**DocumentationType** | Pointer to [**GBTPDefinitionType**](GBTPDefinitionType.md) |  | [optional] 
**EndSourceLocation** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**Name** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ShortDescriptor** | Pointer to **string** |  | [optional] 
**SpaceAfter** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBefore** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceDefault** | Pointer to **bool** |  | [optional] 
**StandardType** | Pointer to [**GBTPType**](GBTPType.md) |  | [optional] 
**StartSourceLocation** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**BTPTypeName290**](BTPTypeName290.md) |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPArgumentDeclaration232

`func NewBTPArgumentDeclaration232() *BTPArgumentDeclaration232`

NewBTPArgumentDeclaration232 instantiates a new BTPArgumentDeclaration232 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPArgumentDeclaration232WithDefaults

`func NewBTPArgumentDeclaration232WithDefaults() *BTPArgumentDeclaration232`

NewBTPArgumentDeclaration232WithDefaults instantiates a new BTPArgumentDeclaration232 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *BTPArgumentDeclaration232) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPArgumentDeclaration232) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPArgumentDeclaration232) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPArgumentDeclaration232) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBtType

`func (o *BTPArgumentDeclaration232) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPArgumentDeclaration232) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPArgumentDeclaration232) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPArgumentDeclaration232) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPArgumentDeclaration232) GetDocumentationType() GBTPDefinitionType`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPArgumentDeclaration232) GetDocumentationTypeOk() (*GBTPDefinitionType, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPArgumentDeclaration232) SetDocumentationType(v GBTPDefinitionType)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPArgumentDeclaration232) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPArgumentDeclaration232) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPArgumentDeclaration232) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPArgumentDeclaration232) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPArgumentDeclaration232) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetIdentifier

`func (o *BTPArgumentDeclaration232) GetIdentifier() BTPIdentifier8`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BTPArgumentDeclaration232) GetIdentifierOk() (*BTPIdentifier8, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BTPArgumentDeclaration232) SetIdentifier(v BTPIdentifier8)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BTPArgumentDeclaration232) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *BTPArgumentDeclaration232) GetName() BTPIdentifier8`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPArgumentDeclaration232) GetNameOk() (*BTPIdentifier8, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPArgumentDeclaration232) SetName(v BTPIdentifier8)`

SetName sets Name field to given value.

### HasName

`func (o *BTPArgumentDeclaration232) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPArgumentDeclaration232) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPArgumentDeclaration232) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPArgumentDeclaration232) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPArgumentDeclaration232) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPArgumentDeclaration232) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPArgumentDeclaration232) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPArgumentDeclaration232) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPArgumentDeclaration232) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPArgumentDeclaration232) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPArgumentDeclaration232) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPArgumentDeclaration232) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPArgumentDeclaration232) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPArgumentDeclaration232) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPArgumentDeclaration232) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPArgumentDeclaration232) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPArgumentDeclaration232) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPArgumentDeclaration232) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPArgumentDeclaration232) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPArgumentDeclaration232) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPArgumentDeclaration232) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetStandardType

`func (o *BTPArgumentDeclaration232) GetStandardType() GBTPType`

GetStandardType returns the StandardType field if non-nil, zero value otherwise.

### GetStandardTypeOk

`func (o *BTPArgumentDeclaration232) GetStandardTypeOk() (*GBTPType, bool)`

GetStandardTypeOk returns a tuple with the StandardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardType

`func (o *BTPArgumentDeclaration232) SetStandardType(v GBTPType)`

SetStandardType sets StandardType field to given value.

### HasStandardType

`func (o *BTPArgumentDeclaration232) HasStandardType() bool`

HasStandardType returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPArgumentDeclaration232) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPArgumentDeclaration232) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPArgumentDeclaration232) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPArgumentDeclaration232) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.

### GetType

`func (o *BTPArgumentDeclaration232) GetType() BTPTypeName290`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTPArgumentDeclaration232) GetTypeOk() (*BTPTypeName290, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTPArgumentDeclaration232) SetType(v BTPTypeName290)`

SetType sets Type field to given value.

### HasType

`func (o *BTPArgumentDeclaration232) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *BTPArgumentDeclaration232) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *BTPArgumentDeclaration232) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *BTPArgumentDeclaration232) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *BTPArgumentDeclaration232) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


