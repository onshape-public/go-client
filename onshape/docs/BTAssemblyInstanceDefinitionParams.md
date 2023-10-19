# BTAssemblyInstanceDefinitionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **string** |  | [optional] 
**DocumentId** | **string** |  | 
**ElementId** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**IncludePartTypes** | Pointer to [**[]GBTInsertableType**](GBTInsertableType.md) |  | [optional] 
**IsAssembly** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsSuppressed** | Pointer to **bool** |  | [optional] 
**IsWholePartStudio** | Pointer to **bool** |  | [optional] 
**MicroversionId** | Pointer to **string** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAssemblyInstanceDefinitionParams

`func NewBTAssemblyInstanceDefinitionParams(documentId string, ) *BTAssemblyInstanceDefinitionParams`

NewBTAssemblyInstanceDefinitionParams instantiates a new BTAssemblyInstanceDefinitionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyInstanceDefinitionParamsWithDefaults

`func NewBTAssemblyInstanceDefinitionParamsWithDefaults() *BTAssemblyInstanceDefinitionParams`

NewBTAssemblyInstanceDefinitionParamsWithDefaults instantiates a new BTAssemblyInstanceDefinitionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTAssemblyInstanceDefinitionParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTAssemblyInstanceDefinitionParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTAssemblyInstanceDefinitionParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTAssemblyInstanceDefinitionParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTAssemblyInstanceDefinitionParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTAssemblyInstanceDefinitionParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTAssemblyInstanceDefinitionParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetElementId

`func (o *BTAssemblyInstanceDefinitionParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTAssemblyInstanceDefinitionParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTAssemblyInstanceDefinitionParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTAssemblyInstanceDefinitionParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTAssemblyInstanceDefinitionParams) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTAssemblyInstanceDefinitionParams) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTAssemblyInstanceDefinitionParams) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTAssemblyInstanceDefinitionParams) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetIncludePartTypes

`func (o *BTAssemblyInstanceDefinitionParams) GetIncludePartTypes() []GBTInsertableType`

GetIncludePartTypes returns the IncludePartTypes field if non-nil, zero value otherwise.

### GetIncludePartTypesOk

`func (o *BTAssemblyInstanceDefinitionParams) GetIncludePartTypesOk() (*[]GBTInsertableType, bool)`

GetIncludePartTypesOk returns a tuple with the IncludePartTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePartTypes

`func (o *BTAssemblyInstanceDefinitionParams) SetIncludePartTypes(v []GBTInsertableType)`

SetIncludePartTypes sets IncludePartTypes field to given value.

### HasIncludePartTypes

`func (o *BTAssemblyInstanceDefinitionParams) HasIncludePartTypes() bool`

HasIncludePartTypes returns a boolean if a field has been set.

### GetIsAssembly

`func (o *BTAssemblyInstanceDefinitionParams) GetIsAssembly() bool`

GetIsAssembly returns the IsAssembly field if non-nil, zero value otherwise.

### GetIsAssemblyOk

`func (o *BTAssemblyInstanceDefinitionParams) GetIsAssemblyOk() (*bool, bool)`

GetIsAssemblyOk returns a tuple with the IsAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssembly

`func (o *BTAssemblyInstanceDefinitionParams) SetIsAssembly(v bool)`

SetIsAssembly sets IsAssembly field to given value.

### HasIsAssembly

`func (o *BTAssemblyInstanceDefinitionParams) HasIsAssembly() bool`

HasIsAssembly returns a boolean if a field has been set.

### GetIsHidden

`func (o *BTAssemblyInstanceDefinitionParams) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *BTAssemblyInstanceDefinitionParams) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *BTAssemblyInstanceDefinitionParams) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *BTAssemblyInstanceDefinitionParams) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsSuppressed

`func (o *BTAssemblyInstanceDefinitionParams) GetIsSuppressed() bool`

GetIsSuppressed returns the IsSuppressed field if non-nil, zero value otherwise.

### GetIsSuppressedOk

`func (o *BTAssemblyInstanceDefinitionParams) GetIsSuppressedOk() (*bool, bool)`

GetIsSuppressedOk returns a tuple with the IsSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuppressed

`func (o *BTAssemblyInstanceDefinitionParams) SetIsSuppressed(v bool)`

SetIsSuppressed sets IsSuppressed field to given value.

### HasIsSuppressed

`func (o *BTAssemblyInstanceDefinitionParams) HasIsSuppressed() bool`

HasIsSuppressed returns a boolean if a field has been set.

### GetIsWholePartStudio

`func (o *BTAssemblyInstanceDefinitionParams) GetIsWholePartStudio() bool`

GetIsWholePartStudio returns the IsWholePartStudio field if non-nil, zero value otherwise.

### GetIsWholePartStudioOk

`func (o *BTAssemblyInstanceDefinitionParams) GetIsWholePartStudioOk() (*bool, bool)`

GetIsWholePartStudioOk returns a tuple with the IsWholePartStudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWholePartStudio

`func (o *BTAssemblyInstanceDefinitionParams) SetIsWholePartStudio(v bool)`

SetIsWholePartStudio sets IsWholePartStudio field to given value.

### HasIsWholePartStudio

`func (o *BTAssemblyInstanceDefinitionParams) HasIsWholePartStudio() bool`

HasIsWholePartStudio returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTAssemblyInstanceDefinitionParams) GetMicroversionId() string`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTAssemblyInstanceDefinitionParams) GetMicroversionIdOk() (*string, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTAssemblyInstanceDefinitionParams) SetMicroversionId(v string)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTAssemblyInstanceDefinitionParams) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetPartId

`func (o *BTAssemblyInstanceDefinitionParams) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTAssemblyInstanceDefinitionParams) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTAssemblyInstanceDefinitionParams) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTAssemblyInstanceDefinitionParams) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTAssemblyInstanceDefinitionParams) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTAssemblyInstanceDefinitionParams) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTAssemblyInstanceDefinitionParams) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTAssemblyInstanceDefinitionParams) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetRevision

`func (o *BTAssemblyInstanceDefinitionParams) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTAssemblyInstanceDefinitionParams) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTAssemblyInstanceDefinitionParams) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTAssemblyInstanceDefinitionParams) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetVersionId

`func (o *BTAssemblyInstanceDefinitionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTAssemblyInstanceDefinitionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTAssemblyInstanceDefinitionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTAssemblyInstanceDefinitionParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


