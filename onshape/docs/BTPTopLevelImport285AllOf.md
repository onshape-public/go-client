# BTPTopLevelImport285AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**CombinedNamespacePathAndVersion** | Pointer to **string** |  | [optional] 
**ImportMicroversion** | Pointer to **string** |  | [optional] 
**ModuleId** | Pointer to [**BTPModuleId235**](BTPModuleId235.md) |  | [optional] 
**Namespace** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**NamespaceString** | Pointer to **string** |  | [optional] 
**SpaceBeforeImport** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 

## Methods

### NewBTPTopLevelImport285AllOf

`func NewBTPTopLevelImport285AllOf() *BTPTopLevelImport285AllOf`

NewBTPTopLevelImport285AllOf instantiates a new BTPTopLevelImport285AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPTopLevelImport285AllOfWithDefaults

`func NewBTPTopLevelImport285AllOfWithDefaults() *BTPTopLevelImport285AllOf`

NewBTPTopLevelImport285AllOfWithDefaults instantiates a new BTPTopLevelImport285AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPTopLevelImport285AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPTopLevelImport285AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPTopLevelImport285AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPTopLevelImport285AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCombinedNamespacePathAndVersion

`func (o *BTPTopLevelImport285AllOf) GetCombinedNamespacePathAndVersion() string`

GetCombinedNamespacePathAndVersion returns the CombinedNamespacePathAndVersion field if non-nil, zero value otherwise.

### GetCombinedNamespacePathAndVersionOk

`func (o *BTPTopLevelImport285AllOf) GetCombinedNamespacePathAndVersionOk() (*string, bool)`

GetCombinedNamespacePathAndVersionOk returns a tuple with the CombinedNamespacePathAndVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedNamespacePathAndVersion

`func (o *BTPTopLevelImport285AllOf) SetCombinedNamespacePathAndVersion(v string)`

SetCombinedNamespacePathAndVersion sets CombinedNamespacePathAndVersion field to given value.

### HasCombinedNamespacePathAndVersion

`func (o *BTPTopLevelImport285AllOf) HasCombinedNamespacePathAndVersion() bool`

HasCombinedNamespacePathAndVersion returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTPTopLevelImport285AllOf) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTPTopLevelImport285AllOf) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTPTopLevelImport285AllOf) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTPTopLevelImport285AllOf) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetModuleId

`func (o *BTPTopLevelImport285AllOf) GetModuleId() BTPModuleId235`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *BTPTopLevelImport285AllOf) GetModuleIdOk() (*BTPModuleId235, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *BTPTopLevelImport285AllOf) SetModuleId(v BTPModuleId235)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *BTPTopLevelImport285AllOf) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetNamespace

`func (o *BTPTopLevelImport285AllOf) GetNamespace() []BTPIdentifier8`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTPTopLevelImport285AllOf) GetNamespaceOk() (*[]BTPIdentifier8, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTPTopLevelImport285AllOf) SetNamespace(v []BTPIdentifier8)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTPTopLevelImport285AllOf) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceString

`func (o *BTPTopLevelImport285AllOf) GetNamespaceString() string`

GetNamespaceString returns the NamespaceString field if non-nil, zero value otherwise.

### GetNamespaceStringOk

`func (o *BTPTopLevelImport285AllOf) GetNamespaceStringOk() (*string, bool)`

GetNamespaceStringOk returns a tuple with the NamespaceString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceString

`func (o *BTPTopLevelImport285AllOf) SetNamespaceString(v string)`

SetNamespaceString sets NamespaceString field to given value.

### HasNamespaceString

`func (o *BTPTopLevelImport285AllOf) HasNamespaceString() bool`

HasNamespaceString returns a boolean if a field has been set.

### GetSpaceBeforeImport

`func (o *BTPTopLevelImport285AllOf) GetSpaceBeforeImport() BTPSpace10`

GetSpaceBeforeImport returns the SpaceBeforeImport field if non-nil, zero value otherwise.

### GetSpaceBeforeImportOk

`func (o *BTPTopLevelImport285AllOf) GetSpaceBeforeImportOk() (*BTPSpace10, bool)`

GetSpaceBeforeImportOk returns a tuple with the SpaceBeforeImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeImport

`func (o *BTPTopLevelImport285AllOf) SetSpaceBeforeImport(v BTPSpace10)`

SetSpaceBeforeImport sets SpaceBeforeImport field to given value.

### HasSpaceBeforeImport

`func (o *BTPTopLevelImport285AllOf) HasSpaceBeforeImport() bool`

HasSpaceBeforeImport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


