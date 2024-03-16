# BTPTopLevelImport285

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**CombinedNamespacePathAndVersion** | Pointer to **string** |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Element microversion that is being imported. | [optional] 
**ModuleId** | Pointer to [**BTPModuleId235**](BTPModuleId235.md) |  | [optional] 
**Namespace** | Pointer to [**[]BTPIdentifier8**](BTPIdentifier8.md) |  | [optional] 
**NamespaceString** | Pointer to **string** |  | [optional] 
**SpaceBeforeImport** | Pointer to [**BTPSpace10**](BTPSpace10.md) |  | [optional] 

## Methods

### NewBTPTopLevelImport285

`func NewBTPTopLevelImport285() *BTPTopLevelImport285`

NewBTPTopLevelImport285 instantiates a new BTPTopLevelImport285 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPTopLevelImport285WithDefaults

`func NewBTPTopLevelImport285WithDefaults() *BTPTopLevelImport285`

NewBTPTopLevelImport285WithDefaults instantiates a new BTPTopLevelImport285 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPTopLevelImport285) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPTopLevelImport285) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPTopLevelImport285) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPTopLevelImport285) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCombinedNamespacePathAndVersion

`func (o *BTPTopLevelImport285) GetCombinedNamespacePathAndVersion() string`

GetCombinedNamespacePathAndVersion returns the CombinedNamespacePathAndVersion field if non-nil, zero value otherwise.

### GetCombinedNamespacePathAndVersionOk

`func (o *BTPTopLevelImport285) GetCombinedNamespacePathAndVersionOk() (*string, bool)`

GetCombinedNamespacePathAndVersionOk returns a tuple with the CombinedNamespacePathAndVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedNamespacePathAndVersion

`func (o *BTPTopLevelImport285) SetCombinedNamespacePathAndVersion(v string)`

SetCombinedNamespacePathAndVersion sets CombinedNamespacePathAndVersion field to given value.

### HasCombinedNamespacePathAndVersion

`func (o *BTPTopLevelImport285) HasCombinedNamespacePathAndVersion() bool`

HasCombinedNamespacePathAndVersion returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTPTopLevelImport285) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTPTopLevelImport285) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTPTopLevelImport285) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTPTopLevelImport285) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetModuleId

`func (o *BTPTopLevelImport285) GetModuleId() BTPModuleId235`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *BTPTopLevelImport285) GetModuleIdOk() (*BTPModuleId235, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *BTPTopLevelImport285) SetModuleId(v BTPModuleId235)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *BTPTopLevelImport285) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetNamespace

`func (o *BTPTopLevelImport285) GetNamespace() []BTPIdentifier8`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTPTopLevelImport285) GetNamespaceOk() (*[]BTPIdentifier8, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTPTopLevelImport285) SetNamespace(v []BTPIdentifier8)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTPTopLevelImport285) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceString

`func (o *BTPTopLevelImport285) GetNamespaceString() string`

GetNamespaceString returns the NamespaceString field if non-nil, zero value otherwise.

### GetNamespaceStringOk

`func (o *BTPTopLevelImport285) GetNamespaceStringOk() (*string, bool)`

GetNamespaceStringOk returns a tuple with the NamespaceString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceString

`func (o *BTPTopLevelImport285) SetNamespaceString(v string)`

SetNamespaceString sets NamespaceString field to given value.

### HasNamespaceString

`func (o *BTPTopLevelImport285) HasNamespaceString() bool`

HasNamespaceString returns a boolean if a field has been set.

### GetSpaceBeforeImport

`func (o *BTPTopLevelImport285) GetSpaceBeforeImport() BTPSpace10`

GetSpaceBeforeImport returns the SpaceBeforeImport field if non-nil, zero value otherwise.

### GetSpaceBeforeImportOk

`func (o *BTPTopLevelImport285) GetSpaceBeforeImportOk() (*BTPSpace10, bool)`

GetSpaceBeforeImportOk returns a tuple with the SpaceBeforeImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeImport

`func (o *BTPTopLevelImport285) SetSpaceBeforeImport(v BTPSpace10)`

SetSpaceBeforeImport sets SpaceBeforeImport field to given value.

### HasSpaceBeforeImport

`func (o *BTPTopLevelImport285) HasSpaceBeforeImport() bool`

HasSpaceBeforeImport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


