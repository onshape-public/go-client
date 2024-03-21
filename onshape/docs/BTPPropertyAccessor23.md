# BTPPropertyAccessor23

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

## Methods

### NewBTPPropertyAccessor23

`func NewBTPPropertyAccessor23() *BTPPropertyAccessor23`

NewBTPPropertyAccessor23 instantiates a new BTPPropertyAccessor23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPPropertyAccessor23WithDefaults

`func NewBTPPropertyAccessor23WithDefaults() *BTPPropertyAccessor23`

NewBTPPropertyAccessor23WithDefaults instantiates a new BTPPropertyAccessor23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *BTPPropertyAccessor23) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPPropertyAccessor23) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPPropertyAccessor23) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPPropertyAccessor23) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBtType

`func (o *BTPPropertyAccessor23) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPPropertyAccessor23) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPPropertyAccessor23) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPPropertyAccessor23) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPPropertyAccessor23) GetDocumentationType() GBTPDefinitionType`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPPropertyAccessor23) GetDocumentationTypeOk() (*GBTPDefinitionType, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPPropertyAccessor23) SetDocumentationType(v GBTPDefinitionType)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPPropertyAccessor23) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPPropertyAccessor23) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPPropertyAccessor23) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPPropertyAccessor23) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPPropertyAccessor23) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPPropertyAccessor23) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPPropertyAccessor23) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPPropertyAccessor23) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPPropertyAccessor23) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPPropertyAccessor23) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPPropertyAccessor23) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPPropertyAccessor23) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPPropertyAccessor23) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPPropertyAccessor23) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPPropertyAccessor23) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPPropertyAccessor23) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPPropertyAccessor23) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPPropertyAccessor23) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPPropertyAccessor23) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPPropertyAccessor23) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPPropertyAccessor23) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPPropertyAccessor23) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPPropertyAccessor23) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPPropertyAccessor23) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPPropertyAccessor23) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPPropertyAccessor23) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPPropertyAccessor23) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPPropertyAccessor23) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPPropertyAccessor23) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


