# BTPLiteralMap256

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atomic** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**DocumentationType** | Pointer to **string** |  | [optional] 
**EndSourceLocation** | Pointer to **int32** |  | [optional] 
**Entries** | Pointer to [**[]BTPLiteralMapEntry257**](BTPLiteralMapEntry-257.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ShortDescriptor** | Pointer to **string** |  | [optional] 
**SpaceAfter** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceBefore** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceDefault** | Pointer to **bool** |  | [optional] 
**SpaceInEmptyList** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**StartSourceLocation** | Pointer to **int32** |  | [optional] 
**TrailingComma** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTPLiteralMap256

`func NewBTPLiteralMap256() *BTPLiteralMap256`

NewBTPLiteralMap256 instantiates a new BTPLiteralMap256 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPLiteralMap256WithDefaults

`func NewBTPLiteralMap256WithDefaults() *BTPLiteralMap256`

NewBTPLiteralMap256WithDefaults instantiates a new BTPLiteralMap256 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *BTPLiteralMap256) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPLiteralMap256) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPLiteralMap256) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPLiteralMap256) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBtType

`func (o *BTPLiteralMap256) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPLiteralMap256) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPLiteralMap256) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPLiteralMap256) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPLiteralMap256) GetDocumentationType() string`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPLiteralMap256) GetDocumentationTypeOk() (*string, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPLiteralMap256) SetDocumentationType(v string)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPLiteralMap256) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPLiteralMap256) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPLiteralMap256) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPLiteralMap256) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPLiteralMap256) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetEntries

`func (o *BTPLiteralMap256) GetEntries() []BTPLiteralMapEntry257`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BTPLiteralMap256) GetEntriesOk() (*[]BTPLiteralMapEntry257, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BTPLiteralMap256) SetEntries(v []BTPLiteralMapEntry257)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BTPLiteralMap256) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPLiteralMap256) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPLiteralMap256) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPLiteralMap256) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPLiteralMap256) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPLiteralMap256) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPLiteralMap256) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPLiteralMap256) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPLiteralMap256) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPLiteralMap256) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPLiteralMap256) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPLiteralMap256) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPLiteralMap256) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPLiteralMap256) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPLiteralMap256) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPLiteralMap256) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPLiteralMap256) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPLiteralMap256) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPLiteralMap256) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPLiteralMap256) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPLiteralMap256) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetSpaceInEmptyList

`func (o *BTPLiteralMap256) GetSpaceInEmptyList() BTPSpace10`

GetSpaceInEmptyList returns the SpaceInEmptyList field if non-nil, zero value otherwise.

### GetSpaceInEmptyListOk

`func (o *BTPLiteralMap256) GetSpaceInEmptyListOk() (*BTPSpace10, bool)`

GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceInEmptyList

`func (o *BTPLiteralMap256) SetSpaceInEmptyList(v BTPSpace10)`

SetSpaceInEmptyList sets SpaceInEmptyList field to given value.

### HasSpaceInEmptyList

`func (o *BTPLiteralMap256) HasSpaceInEmptyList() bool`

HasSpaceInEmptyList returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPLiteralMap256) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPLiteralMap256) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPLiteralMap256) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPLiteralMap256) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.

### GetTrailingComma

`func (o *BTPLiteralMap256) GetTrailingComma() bool`

GetTrailingComma returns the TrailingComma field if non-nil, zero value otherwise.

### GetTrailingCommaOk

`func (o *BTPLiteralMap256) GetTrailingCommaOk() (*bool, bool)`

GetTrailingCommaOk returns a tuple with the TrailingComma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingComma

`func (o *BTPLiteralMap256) SetTrailingComma(v bool)`

SetTrailingComma sets TrailingComma field to given value.

### HasTrailingComma

`func (o *BTPLiteralMap256) HasTrailingComma() bool`

HasTrailingComma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


