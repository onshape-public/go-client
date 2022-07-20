# BTPNode7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atomic** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**DocumentationType** | Pointer to **string** |  | [optional] 
**EndSourceLocation** | Pointer to **int32** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ShortDescriptor** | Pointer to **string** |  | [optional] 
**SpaceAfter** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBefore** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceDefault** | Pointer to **bool** |  | [optional] 
**StartSourceLocation** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTPNode7

`func NewBTPNode7() *BTPNode7`

NewBTPNode7 instantiates a new BTPNode7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPNode7WithDefaults

`func NewBTPNode7WithDefaults() *BTPNode7`

NewBTPNode7WithDefaults instantiates a new BTPNode7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *BTPNode7) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPNode7) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPNode7) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPNode7) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBtType

`func (o *BTPNode7) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPNode7) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPNode7) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPNode7) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPNode7) GetDocumentationType() string`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPNode7) GetDocumentationTypeOk() (*string, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPNode7) SetDocumentationType(v string)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPNode7) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPNode7) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPNode7) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPNode7) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPNode7) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPNode7) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPNode7) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPNode7) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPNode7) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPNode7) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPNode7) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPNode7) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPNode7) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPNode7) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPNode7) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPNode7) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPNode7) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPNode7) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPNode7) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPNode7) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPNode7) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPNode7) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPNode7) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPNode7) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPNode7) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPNode7) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPNode7) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPNode7) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPNode7) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


