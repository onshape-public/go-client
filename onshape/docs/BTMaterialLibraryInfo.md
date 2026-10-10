# BTMaterialLibraryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalElementReference** | Pointer to [**BTExternalElementReferenceInfo**](BTExternalElementReferenceInfo.md) |  | [optional] 
**Materials** | Pointer to [**[]BTLibraryMaterialInfo**](BTLibraryMaterialInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PropertyDefinitions** | Pointer to [**[]BTMaterialPropertyDefinitionInfo**](BTMaterialPropertyDefinitionInfo.md) |  | [optional] 
**Versioned** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTMaterialLibraryInfo

`func NewBTMaterialLibraryInfo() *BTMaterialLibraryInfo`

NewBTMaterialLibraryInfo instantiates a new BTMaterialLibraryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMaterialLibraryInfoWithDefaults

`func NewBTMaterialLibraryInfoWithDefaults() *BTMaterialLibraryInfo`

NewBTMaterialLibraryInfoWithDefaults instantiates a new BTMaterialLibraryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTMaterialLibraryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTMaterialLibraryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTMaterialLibraryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTMaterialLibraryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *BTMaterialLibraryInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BTMaterialLibraryInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BTMaterialLibraryInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BTMaterialLibraryInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalElementReference

`func (o *BTMaterialLibraryInfo) GetExternalElementReference() BTExternalElementReferenceInfo`

GetExternalElementReference returns the ExternalElementReference field if non-nil, zero value otherwise.

### GetExternalElementReferenceOk

`func (o *BTMaterialLibraryInfo) GetExternalElementReferenceOk() (*BTExternalElementReferenceInfo, bool)`

GetExternalElementReferenceOk returns a tuple with the ExternalElementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalElementReference

`func (o *BTMaterialLibraryInfo) SetExternalElementReference(v BTExternalElementReferenceInfo)`

SetExternalElementReference sets ExternalElementReference field to given value.

### HasExternalElementReference

`func (o *BTMaterialLibraryInfo) HasExternalElementReference() bool`

HasExternalElementReference returns a boolean if a field has been set.

### GetMaterials

`func (o *BTMaterialLibraryInfo) GetMaterials() []BTLibraryMaterialInfo`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *BTMaterialLibraryInfo) GetMaterialsOk() (*[]BTLibraryMaterialInfo, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *BTMaterialLibraryInfo) SetMaterials(v []BTLibraryMaterialInfo)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *BTMaterialLibraryInfo) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetName

`func (o *BTMaterialLibraryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMaterialLibraryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMaterialLibraryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMaterialLibraryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyDefinitions

`func (o *BTMaterialLibraryInfo) GetPropertyDefinitions() []BTMaterialPropertyDefinitionInfo`

GetPropertyDefinitions returns the PropertyDefinitions field if non-nil, zero value otherwise.

### GetPropertyDefinitionsOk

`func (o *BTMaterialLibraryInfo) GetPropertyDefinitionsOk() (*[]BTMaterialPropertyDefinitionInfo, bool)`

GetPropertyDefinitionsOk returns a tuple with the PropertyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDefinitions

`func (o *BTMaterialLibraryInfo) SetPropertyDefinitions(v []BTMaterialPropertyDefinitionInfo)`

SetPropertyDefinitions sets PropertyDefinitions field to given value.

### HasPropertyDefinitions

`func (o *BTMaterialLibraryInfo) HasPropertyDefinitions() bool`

HasPropertyDefinitions returns a boolean if a field has been set.

### GetVersioned

`func (o *BTMaterialLibraryInfo) GetVersioned() bool`

GetVersioned returns the Versioned field if non-nil, zero value otherwise.

### GetVersionedOk

`func (o *BTMaterialLibraryInfo) GetVersionedOk() (*bool, bool)`

GetVersionedOk returns a tuple with the Versioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioned

`func (o *BTMaterialLibraryInfo) SetVersioned(v bool)`

SetVersioned sets Versioned field to given value.

### HasVersioned

`func (o *BTMaterialLibraryInfo) HasVersioned() bool`

HasVersioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


