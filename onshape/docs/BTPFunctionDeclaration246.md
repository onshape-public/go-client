# BTPFunctionDeclaration246

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**BTPAnnotation231**](BTPAnnotation231.md) |  | [optional] 
**Arguments** | Pointer to [**[]BTPArgumentDeclaration232**](BTPArgumentDeclaration232.md) |  | [optional] 
**ArgumentsToDocument** | Pointer to [**[]BTPArgumentDeclaration232**](BTPArgumentDeclaration232.md) |  | [optional] 
**Atomic** | Pointer to **bool** |  | [optional] 
**Body** | Pointer to [**BTPStatementBlock271**](BTPStatementBlock271.md) |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DeprecatedExplanation** | Pointer to **string** |  | [optional] 
**DocumentationType** | Pointer to [**GBTPDefinitionType**](GBTPDefinitionType.md) |  | [optional] 
**EndSourceLocation** | Pointer to **int32** |  | [optional] 
**ForExport** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Precondition** | Pointer to [**BTPStatement269**](BTPStatement269.md) |  | [optional] 
**ReturnType** | Pointer to [**BTPTypeName290**](BTPTypeName290.md) |  | [optional] 
**ShortDescriptor** | Pointer to **string** |  | [optional] 
**SpaceAfter** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceAfterArglist** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceAfterExport** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBefore** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceDefault** | Pointer to **bool** |  | [optional] 
**SpaceInEmptyList** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**StartSourceLocation** | Pointer to **int32** |  | [optional] 
**SymbolName** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 

## Methods

### NewBTPFunctionDeclaration246

`func NewBTPFunctionDeclaration246() *BTPFunctionDeclaration246`

NewBTPFunctionDeclaration246 instantiates a new BTPFunctionDeclaration246 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPFunctionDeclaration246WithDefaults

`func NewBTPFunctionDeclaration246WithDefaults() *BTPFunctionDeclaration246`

NewBTPFunctionDeclaration246WithDefaults instantiates a new BTPFunctionDeclaration246 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *BTPFunctionDeclaration246) GetAnnotation() BTPAnnotation231`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *BTPFunctionDeclaration246) GetAnnotationOk() (*BTPAnnotation231, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *BTPFunctionDeclaration246) SetAnnotation(v BTPAnnotation231)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *BTPFunctionDeclaration246) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetArguments

`func (o *BTPFunctionDeclaration246) GetArguments() []BTPArgumentDeclaration232`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *BTPFunctionDeclaration246) GetArgumentsOk() (*[]BTPArgumentDeclaration232, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *BTPFunctionDeclaration246) SetArguments(v []BTPArgumentDeclaration232)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *BTPFunctionDeclaration246) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetArgumentsToDocument

`func (o *BTPFunctionDeclaration246) GetArgumentsToDocument() []BTPArgumentDeclaration232`

GetArgumentsToDocument returns the ArgumentsToDocument field if non-nil, zero value otherwise.

### GetArgumentsToDocumentOk

`func (o *BTPFunctionDeclaration246) GetArgumentsToDocumentOk() (*[]BTPArgumentDeclaration232, bool)`

GetArgumentsToDocumentOk returns a tuple with the ArgumentsToDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgumentsToDocument

`func (o *BTPFunctionDeclaration246) SetArgumentsToDocument(v []BTPArgumentDeclaration232)`

SetArgumentsToDocument sets ArgumentsToDocument field to given value.

### HasArgumentsToDocument

`func (o *BTPFunctionDeclaration246) HasArgumentsToDocument() bool`

HasArgumentsToDocument returns a boolean if a field has been set.

### GetAtomic

`func (o *BTPFunctionDeclaration246) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPFunctionDeclaration246) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPFunctionDeclaration246) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPFunctionDeclaration246) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBody

`func (o *BTPFunctionDeclaration246) GetBody() BTPStatementBlock271`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BTPFunctionDeclaration246) GetBodyOk() (*BTPStatementBlock271, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BTPFunctionDeclaration246) SetBody(v BTPStatementBlock271)`

SetBody sets Body field to given value.

### HasBody

`func (o *BTPFunctionDeclaration246) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBtType

`func (o *BTPFunctionDeclaration246) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPFunctionDeclaration246) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPFunctionDeclaration246) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPFunctionDeclaration246) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDeprecated

`func (o *BTPFunctionDeclaration246) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *BTPFunctionDeclaration246) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *BTPFunctionDeclaration246) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *BTPFunctionDeclaration246) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDeprecatedExplanation

`func (o *BTPFunctionDeclaration246) GetDeprecatedExplanation() string`

GetDeprecatedExplanation returns the DeprecatedExplanation field if non-nil, zero value otherwise.

### GetDeprecatedExplanationOk

`func (o *BTPFunctionDeclaration246) GetDeprecatedExplanationOk() (*string, bool)`

GetDeprecatedExplanationOk returns a tuple with the DeprecatedExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedExplanation

`func (o *BTPFunctionDeclaration246) SetDeprecatedExplanation(v string)`

SetDeprecatedExplanation sets DeprecatedExplanation field to given value.

### HasDeprecatedExplanation

`func (o *BTPFunctionDeclaration246) HasDeprecatedExplanation() bool`

HasDeprecatedExplanation returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPFunctionDeclaration246) GetDocumentationType() GBTPDefinitionType`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPFunctionDeclaration246) GetDocumentationTypeOk() (*GBTPDefinitionType, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPFunctionDeclaration246) SetDocumentationType(v GBTPDefinitionType)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPFunctionDeclaration246) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPFunctionDeclaration246) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPFunctionDeclaration246) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPFunctionDeclaration246) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPFunctionDeclaration246) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetForExport

`func (o *BTPFunctionDeclaration246) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *BTPFunctionDeclaration246) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *BTPFunctionDeclaration246) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *BTPFunctionDeclaration246) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetName

`func (o *BTPFunctionDeclaration246) GetName() BTPIdentifier8`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPFunctionDeclaration246) GetNameOk() (*BTPIdentifier8, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPFunctionDeclaration246) SetName(v BTPIdentifier8)`

SetName sets Name field to given value.

### HasName

`func (o *BTPFunctionDeclaration246) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPFunctionDeclaration246) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPFunctionDeclaration246) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPFunctionDeclaration246) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPFunctionDeclaration246) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetPrecondition

`func (o *BTPFunctionDeclaration246) GetPrecondition() BTPStatement269`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *BTPFunctionDeclaration246) GetPreconditionOk() (*BTPStatement269, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *BTPFunctionDeclaration246) SetPrecondition(v BTPStatement269)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *BTPFunctionDeclaration246) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetReturnType

`func (o *BTPFunctionDeclaration246) GetReturnType() BTPTypeName290`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *BTPFunctionDeclaration246) GetReturnTypeOk() (*BTPTypeName290, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *BTPFunctionDeclaration246) SetReturnType(v BTPTypeName290)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *BTPFunctionDeclaration246) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPFunctionDeclaration246) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPFunctionDeclaration246) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPFunctionDeclaration246) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPFunctionDeclaration246) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPFunctionDeclaration246) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPFunctionDeclaration246) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPFunctionDeclaration246) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPFunctionDeclaration246) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceAfterArglist

`func (o *BTPFunctionDeclaration246) GetSpaceAfterArglist() BTPSpace10`

GetSpaceAfterArglist returns the SpaceAfterArglist field if non-nil, zero value otherwise.

### GetSpaceAfterArglistOk

`func (o *BTPFunctionDeclaration246) GetSpaceAfterArglistOk() (*BTPSpace10, bool)`

GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterArglist

`func (o *BTPFunctionDeclaration246) SetSpaceAfterArglist(v BTPSpace10)`

SetSpaceAfterArglist sets SpaceAfterArglist field to given value.

### HasSpaceAfterArglist

`func (o *BTPFunctionDeclaration246) HasSpaceAfterArglist() bool`

HasSpaceAfterArglist returns a boolean if a field has been set.

### GetSpaceAfterExport

`func (o *BTPFunctionDeclaration246) GetSpaceAfterExport() BTPSpace10`

GetSpaceAfterExport returns the SpaceAfterExport field if non-nil, zero value otherwise.

### GetSpaceAfterExportOk

`func (o *BTPFunctionDeclaration246) GetSpaceAfterExportOk() (*BTPSpace10, bool)`

GetSpaceAfterExportOk returns a tuple with the SpaceAfterExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterExport

`func (o *BTPFunctionDeclaration246) SetSpaceAfterExport(v BTPSpace10)`

SetSpaceAfterExport sets SpaceAfterExport field to given value.

### HasSpaceAfterExport

`func (o *BTPFunctionDeclaration246) HasSpaceAfterExport() bool`

HasSpaceAfterExport returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPFunctionDeclaration246) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPFunctionDeclaration246) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPFunctionDeclaration246) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPFunctionDeclaration246) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPFunctionDeclaration246) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPFunctionDeclaration246) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPFunctionDeclaration246) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPFunctionDeclaration246) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetSpaceInEmptyList

`func (o *BTPFunctionDeclaration246) GetSpaceInEmptyList() BTPSpace10`

GetSpaceInEmptyList returns the SpaceInEmptyList field if non-nil, zero value otherwise.

### GetSpaceInEmptyListOk

`func (o *BTPFunctionDeclaration246) GetSpaceInEmptyListOk() (*BTPSpace10, bool)`

GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceInEmptyList

`func (o *BTPFunctionDeclaration246) SetSpaceInEmptyList(v BTPSpace10)`

SetSpaceInEmptyList sets SpaceInEmptyList field to given value.

### HasSpaceInEmptyList

`func (o *BTPFunctionDeclaration246) HasSpaceInEmptyList() bool`

HasSpaceInEmptyList returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPFunctionDeclaration246) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPFunctionDeclaration246) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPFunctionDeclaration246) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPFunctionDeclaration246) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.

### GetSymbolName

`func (o *BTPFunctionDeclaration246) GetSymbolName() BTPIdentifier8`

GetSymbolName returns the SymbolName field if non-nil, zero value otherwise.

### GetSymbolNameOk

`func (o *BTPFunctionDeclaration246) GetSymbolNameOk() (*BTPIdentifier8, bool)`

GetSymbolNameOk returns a tuple with the SymbolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolName

`func (o *BTPFunctionDeclaration246) SetSymbolName(v BTPIdentifier8)`

SetSymbolName sets SymbolName field to given value.

### HasSymbolName

`func (o *BTPFunctionDeclaration246) HasSymbolName() bool`

HasSymbolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


