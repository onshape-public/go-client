# BTPName261

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atomic** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**DocumentationType** | Pointer to [**GBTPDefinitionType**](GBTPDefinitionType.md) |  | [optional] 
**EndSourceLocation** | Pointer to **int32** |  | [optional] 
**ForExport** | Pointer to **bool** |  | [optional] 
**GlobalNamespace** | Pointer to **bool** |  | [optional] 
**Identifier** | Pointer to [**BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Element microversion that is being imported. | [optional] 
**Namespace** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ShortDescriptor** | Pointer to **string** |  | [optional] 
**SpaceAfter** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceBefore** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 
**SpaceDefault** | Pointer to **bool** |  | [optional] 
**StartSourceLocation** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTPName261

`func NewBTPName261() *BTPName261`

NewBTPName261 instantiates a new BTPName261 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPName261WithDefaults

`func NewBTPName261WithDefaults() *BTPName261`

NewBTPName261WithDefaults instantiates a new BTPName261 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *BTPName261) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *BTPName261) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *BTPName261) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *BTPName261) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetBtType

`func (o *BTPName261) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPName261) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPName261) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPName261) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDocumentationType

`func (o *BTPName261) GetDocumentationType() GBTPDefinitionType`

GetDocumentationType returns the DocumentationType field if non-nil, zero value otherwise.

### GetDocumentationTypeOk

`func (o *BTPName261) GetDocumentationTypeOk() (*GBTPDefinitionType, bool)`

GetDocumentationTypeOk returns a tuple with the DocumentationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationType

`func (o *BTPName261) SetDocumentationType(v GBTPDefinitionType)`

SetDocumentationType sets DocumentationType field to given value.

### HasDocumentationType

`func (o *BTPName261) HasDocumentationType() bool`

HasDocumentationType returns a boolean if a field has been set.

### GetEndSourceLocation

`func (o *BTPName261) GetEndSourceLocation() int32`

GetEndSourceLocation returns the EndSourceLocation field if non-nil, zero value otherwise.

### GetEndSourceLocationOk

`func (o *BTPName261) GetEndSourceLocationOk() (*int32, bool)`

GetEndSourceLocationOk returns a tuple with the EndSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSourceLocation

`func (o *BTPName261) SetEndSourceLocation(v int32)`

SetEndSourceLocation sets EndSourceLocation field to given value.

### HasEndSourceLocation

`func (o *BTPName261) HasEndSourceLocation() bool`

HasEndSourceLocation returns a boolean if a field has been set.

### GetForExport

`func (o *BTPName261) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *BTPName261) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *BTPName261) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *BTPName261) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetGlobalNamespace

`func (o *BTPName261) GetGlobalNamespace() bool`

GetGlobalNamespace returns the GlobalNamespace field if non-nil, zero value otherwise.

### GetGlobalNamespaceOk

`func (o *BTPName261) GetGlobalNamespaceOk() (*bool, bool)`

GetGlobalNamespaceOk returns a tuple with the GlobalNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalNamespace

`func (o *BTPName261) SetGlobalNamespace(v bool)`

SetGlobalNamespace sets GlobalNamespace field to given value.

### HasGlobalNamespace

`func (o *BTPName261) HasGlobalNamespace() bool`

HasGlobalNamespace returns a boolean if a field has been set.

### GetIdentifier

`func (o *BTPName261) GetIdentifier() BTPIdentifier8`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BTPName261) GetIdentifierOk() (*BTPIdentifier8, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BTPName261) SetIdentifier(v BTPIdentifier8)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BTPName261) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTPName261) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTPName261) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTPName261) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTPName261) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetNamespace

`func (o *BTPName261) GetNamespace() []BTPIdentifier8`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTPName261) GetNamespaceOk() (*[]BTPIdentifier8, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTPName261) SetNamespace(v []BTPIdentifier8)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTPName261) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeId

`func (o *BTPName261) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTPName261) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTPName261) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTPName261) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetShortDescriptor

`func (o *BTPName261) GetShortDescriptor() string`

GetShortDescriptor returns the ShortDescriptor field if non-nil, zero value otherwise.

### GetShortDescriptorOk

`func (o *BTPName261) GetShortDescriptorOk() (*string, bool)`

GetShortDescriptorOk returns a tuple with the ShortDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescriptor

`func (o *BTPName261) SetShortDescriptor(v string)`

SetShortDescriptor sets ShortDescriptor field to given value.

### HasShortDescriptor

`func (o *BTPName261) HasShortDescriptor() bool`

HasShortDescriptor returns a boolean if a field has been set.

### GetSpaceAfter

`func (o *BTPName261) GetSpaceAfter() BTPSpace10`

GetSpaceAfter returns the SpaceAfter field if non-nil, zero value otherwise.

### GetSpaceAfterOk

`func (o *BTPName261) GetSpaceAfterOk() (*BTPSpace10, bool)`

GetSpaceAfterOk returns a tuple with the SpaceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfter

`func (o *BTPName261) SetSpaceAfter(v BTPSpace10)`

SetSpaceAfter sets SpaceAfter field to given value.

### HasSpaceAfter

`func (o *BTPName261) HasSpaceAfter() bool`

HasSpaceAfter returns a boolean if a field has been set.

### GetSpaceBefore

`func (o *BTPName261) GetSpaceBefore() BTPSpace10`

GetSpaceBefore returns the SpaceBefore field if non-nil, zero value otherwise.

### GetSpaceBeforeOk

`func (o *BTPName261) GetSpaceBeforeOk() (*BTPSpace10, bool)`

GetSpaceBeforeOk returns a tuple with the SpaceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBefore

`func (o *BTPName261) SetSpaceBefore(v BTPSpace10)`

SetSpaceBefore sets SpaceBefore field to given value.

### HasSpaceBefore

`func (o *BTPName261) HasSpaceBefore() bool`

HasSpaceBefore returns a boolean if a field has been set.

### GetSpaceDefault

`func (o *BTPName261) GetSpaceDefault() bool`

GetSpaceDefault returns the SpaceDefault field if non-nil, zero value otherwise.

### GetSpaceDefaultOk

`func (o *BTPName261) GetSpaceDefaultOk() (*bool, bool)`

GetSpaceDefaultOk returns a tuple with the SpaceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDefault

`func (o *BTPName261) SetSpaceDefault(v bool)`

SetSpaceDefault sets SpaceDefault field to given value.

### HasSpaceDefault

`func (o *BTPName261) HasSpaceDefault() bool`

HasSpaceDefault returns a boolean if a field has been set.

### GetStartSourceLocation

`func (o *BTPName261) GetStartSourceLocation() int32`

GetStartSourceLocation returns the StartSourceLocation field if non-nil, zero value otherwise.

### GetStartSourceLocationOk

`func (o *BTPName261) GetStartSourceLocationOk() (*int32, bool)`

GetStartSourceLocationOk returns a tuple with the StartSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSourceLocation

`func (o *BTPName261) SetStartSourceLocation(v int32)`

SetStartSourceLocation sets StartSourceLocation field to given value.

### HasStartSourceLocation

`func (o *BTPName261) HasStartSourceLocation() bool`

HasStartSourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


