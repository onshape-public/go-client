# BTPAnnotation231

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atomic** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**DocumentationType** | Pointer to [**GBTPDefinitionType**](GBTPDefinitionType.md) |  | [optional] 
**EndSourceLocation** | Pointer to **int32** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ShortDescriptor** | Pointer to **string** |  | [optional] 
**SpaceAfter** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBefore** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceDefault** | Pointer to **bool** |  | [optional] 
**StartSourceLocation** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to [**BTPLiteralMap256**](BTPLiteralMap256.md) |  | [optional] 

## Methods

### NewBTPAnnotation231

`func NewBTPAnnotation231() *BTPAnnotation231`

NewBTPAnnotation231 instantiates a new BTPAnnotation231 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPAnnotation231WithDefaults

`func NewBTPAnnotation231WithDefaults() *BTPAnnotation231`

NewBTPAnnotation231WithDefaults instantiates a new BTPAnnotation231 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *BTPAnnotation231) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPAnnotation231) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPAnnotation231) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPAnnotation231) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBtType

`func (o *BTPAnnotation231) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPAnnotation231) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPAnnotation231) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPAnnotation231) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPAnnotation231) GetDocumentationType() GBTPDefinitionType`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPAnnotation231) GetDocumentationTypeOk() (*GBTPDefinitionType, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPAnnotation231) SetDocumentationType(v GBTPDefinitionType)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPAnnotation231) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPAnnotation231) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPAnnotation231) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPAnnotation231) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPAnnotation231) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPAnnotation231) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPAnnotation231) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPAnnotation231) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPAnnotation231) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPAnnotation231) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPAnnotation231) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPAnnotation231) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPAnnotation231) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPAnnotation231) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPAnnotation231) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPAnnotation231) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPAnnotation231) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPAnnotation231) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPAnnotation231) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPAnnotation231) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPAnnotation231) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPAnnotation231) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPAnnotation231) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPAnnotation231) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPAnnotation231) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPAnnotation231) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPAnnotation231) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPAnnotation231) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPAnnotation231) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.

### GetValue

`func (o *BTPAnnotation231) GetValue() BTPLiteralMap256`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTPAnnotation231) GetValueOk() (*BTPLiteralMap256, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTPAnnotation231) SetValue(v BTPLiteralMap256)`

SetValue sets Value field to given value.

### HasValue

`func (o *BTPAnnotation231) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


