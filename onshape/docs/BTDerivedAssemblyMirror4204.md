# BTDerivedAssemblyMirror4204

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Configuration** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementReference** | Pointer to [**BTElementReference725**](BTElementReference725.md) |  | [optional] 
**ExternalDocumentWithVersion** | Pointer to [**BTDocumentWithVersionId**](BTDocumentWithVersionId.md) |  | [optional] 
**ExternalDocumentWithVersionAndElementId** | Pointer to [**BTDocumentWithVersionAndElementId**](BTDocumentWithVersionAndElementId.md) |  | [optional] 
**LockedState** | Pointer to [**BTInstanceWithReference**](BTInstanceWithReference.md) |  | [optional] 
**MicroversionId** | Pointer to [**BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**MirrorFeature** | Pointer to [**BTMDerivedAssemblyMirrorFeature5094**](BTMDerivedAssemblyMirrorFeature5094.md) |  | [optional] 
**MirrorFeatureInterface** | Pointer to [**BTAssemblyMirrorFeatureInterface**](BTAssemblyMirrorFeatureInterface.md) |  | [optional] 
**NodeWithReferenceList** | Pointer to [**[]BTNodeWithReference**](BTNodeWithReference.md) |  | [optional] 
**ReferenceParameter** | Pointer to [**BTMParameterReferenceWithConfiguration3028**](BTMParameterReferenceWithConfiguration3028.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionIdIfExternal** | Pointer to **string** |  | [optional] 

## Methods

### NewBTDerivedAssemblyMirror4204

`func NewBTDerivedAssemblyMirror4204() *BTDerivedAssemblyMirror4204`

NewBTDerivedAssemblyMirror4204 instantiates a new BTDerivedAssemblyMirror4204 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDerivedAssemblyMirror4204WithDefaults

`func NewBTDerivedAssemblyMirror4204WithDefaults() *BTDerivedAssemblyMirror4204`

NewBTDerivedAssemblyMirror4204WithDefaults instantiates a new BTDerivedAssemblyMirror4204 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTDerivedAssemblyMirror4204) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTDerivedAssemblyMirror4204) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTDerivedAssemblyMirror4204) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTDerivedAssemblyMirror4204) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTDerivedAssemblyMirror4204) GetConfiguration() []BTMParameter1`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTDerivedAssemblyMirror4204) GetConfigurationOk() (*[]BTMParameter1, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTDerivedAssemblyMirror4204) SetConfiguration(v []BTMParameter1)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTDerivedAssemblyMirror4204) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTDerivedAssemblyMirror4204) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTDerivedAssemblyMirror4204) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTDerivedAssemblyMirror4204) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTDerivedAssemblyMirror4204) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTDerivedAssemblyMirror4204) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTDerivedAssemblyMirror4204) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTDerivedAssemblyMirror4204) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTDerivedAssemblyMirror4204) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementReference

`func (o *BTDerivedAssemblyMirror4204) GetElementReference() BTElementReference725`

GetElementReference returns the ElementReference field if non-nil, zero value otherwise.

### GetElementReferenceOk

`func (o *BTDerivedAssemblyMirror4204) GetElementReferenceOk() (*BTElementReference725, bool)`

GetElementReferenceOk returns a tuple with the ElementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementReference

`func (o *BTDerivedAssemblyMirror4204) SetElementReference(v BTElementReference725)`

SetElementReference sets ElementReference field to given value.

### HasElementReference

`func (o *BTDerivedAssemblyMirror4204) HasElementReference() bool`

HasElementReference returns a boolean if a field has been set.

### GetExternalDocumentWithVersion

`func (o *BTDerivedAssemblyMirror4204) GetExternalDocumentWithVersion() BTDocumentWithVersionId`

GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionOk

`func (o *BTDerivedAssemblyMirror4204) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool)`

GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersion

`func (o *BTDerivedAssemblyMirror4204) SetExternalDocumentWithVersion(v BTDocumentWithVersionId)`

SetExternalDocumentWithVersion sets ExternalDocumentWithVersion field to given value.

### HasExternalDocumentWithVersion

`func (o *BTDerivedAssemblyMirror4204) HasExternalDocumentWithVersion() bool`

HasExternalDocumentWithVersion returns a boolean if a field has been set.

### GetExternalDocumentWithVersionAndElementId

`func (o *BTDerivedAssemblyMirror4204) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId`

GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionAndElementIdOk

`func (o *BTDerivedAssemblyMirror4204) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool)`

GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersionAndElementId

`func (o *BTDerivedAssemblyMirror4204) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId)`

SetExternalDocumentWithVersionAndElementId sets ExternalDocumentWithVersionAndElementId field to given value.

### HasExternalDocumentWithVersionAndElementId

`func (o *BTDerivedAssemblyMirror4204) HasExternalDocumentWithVersionAndElementId() bool`

HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.

### GetLockedState

`func (o *BTDerivedAssemblyMirror4204) GetLockedState() BTInstanceWithReference`

GetLockedState returns the LockedState field if non-nil, zero value otherwise.

### GetLockedStateOk

`func (o *BTDerivedAssemblyMirror4204) GetLockedStateOk() (*BTInstanceWithReference, bool)`

GetLockedStateOk returns a tuple with the LockedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedState

`func (o *BTDerivedAssemblyMirror4204) SetLockedState(v BTInstanceWithReference)`

SetLockedState sets LockedState field to given value.

### HasLockedState

`func (o *BTDerivedAssemblyMirror4204) HasLockedState() bool`

HasLockedState returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTDerivedAssemblyMirror4204) GetMicroversionId() BTMicroversionId366`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTDerivedAssemblyMirror4204) GetMicroversionIdOk() (*BTMicroversionId366, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTDerivedAssemblyMirror4204) SetMicroversionId(v BTMicroversionId366)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTDerivedAssemblyMirror4204) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetMirrorFeature

`func (o *BTDerivedAssemblyMirror4204) GetMirrorFeature() BTMDerivedAssemblyMirrorFeature5094`

GetMirrorFeature returns the MirrorFeature field if non-nil, zero value otherwise.

### GetMirrorFeatureOk

`func (o *BTDerivedAssemblyMirror4204) GetMirrorFeatureOk() (*BTMDerivedAssemblyMirrorFeature5094, bool)`

GetMirrorFeatureOk returns a tuple with the MirrorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorFeature

`func (o *BTDerivedAssemblyMirror4204) SetMirrorFeature(v BTMDerivedAssemblyMirrorFeature5094)`

SetMirrorFeature sets MirrorFeature field to given value.

### HasMirrorFeature

`func (o *BTDerivedAssemblyMirror4204) HasMirrorFeature() bool`

HasMirrorFeature returns a boolean if a field has been set.

### GetMirrorFeatureInterface

`func (o *BTDerivedAssemblyMirror4204) GetMirrorFeatureInterface() BTAssemblyMirrorFeatureInterface`

GetMirrorFeatureInterface returns the MirrorFeatureInterface field if non-nil, zero value otherwise.

### GetMirrorFeatureInterfaceOk

`func (o *BTDerivedAssemblyMirror4204) GetMirrorFeatureInterfaceOk() (*BTAssemblyMirrorFeatureInterface, bool)`

GetMirrorFeatureInterfaceOk returns a tuple with the MirrorFeatureInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorFeatureInterface

`func (o *BTDerivedAssemblyMirror4204) SetMirrorFeatureInterface(v BTAssemblyMirrorFeatureInterface)`

SetMirrorFeatureInterface sets MirrorFeatureInterface field to given value.

### HasMirrorFeatureInterface

`func (o *BTDerivedAssemblyMirror4204) HasMirrorFeatureInterface() bool`

HasMirrorFeatureInterface returns a boolean if a field has been set.

### GetNodeWithReferenceList

`func (o *BTDerivedAssemblyMirror4204) GetNodeWithReferenceList() []BTNodeWithReference`

GetNodeWithReferenceList returns the NodeWithReferenceList field if non-nil, zero value otherwise.

### GetNodeWithReferenceListOk

`func (o *BTDerivedAssemblyMirror4204) GetNodeWithReferenceListOk() (*[]BTNodeWithReference, bool)`

GetNodeWithReferenceListOk returns a tuple with the NodeWithReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWithReferenceList

`func (o *BTDerivedAssemblyMirror4204) SetNodeWithReferenceList(v []BTNodeWithReference)`

SetNodeWithReferenceList sets NodeWithReferenceList field to given value.

### HasNodeWithReferenceList

`func (o *BTDerivedAssemblyMirror4204) HasNodeWithReferenceList() bool`

HasNodeWithReferenceList returns a boolean if a field has been set.

### GetReferenceParameter

`func (o *BTDerivedAssemblyMirror4204) GetReferenceParameter() BTMParameterReferenceWithConfiguration3028`

GetReferenceParameter returns the ReferenceParameter field if non-nil, zero value otherwise.

### GetReferenceParameterOk

`func (o *BTDerivedAssemblyMirror4204) GetReferenceParameterOk() (*BTMParameterReferenceWithConfiguration3028, bool)`

GetReferenceParameterOk returns a tuple with the ReferenceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceParameter

`func (o *BTDerivedAssemblyMirror4204) SetReferenceParameter(v BTMParameterReferenceWithConfiguration3028)`

SetReferenceParameter sets ReferenceParameter field to given value.

### HasReferenceParameter

`func (o *BTDerivedAssemblyMirror4204) HasReferenceParameter() bool`

HasReferenceParameter returns a boolean if a field has been set.

### GetVersionId

`func (o *BTDerivedAssemblyMirror4204) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTDerivedAssemblyMirror4204) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTDerivedAssemblyMirror4204) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTDerivedAssemblyMirror4204) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionIdIfExternal

`func (o *BTDerivedAssemblyMirror4204) GetVersionIdIfExternal() string`

GetVersionIdIfExternal returns the VersionIdIfExternal field if non-nil, zero value otherwise.

### GetVersionIdIfExternalOk

`func (o *BTDerivedAssemblyMirror4204) GetVersionIdIfExternalOk() (*string, bool)`

GetVersionIdIfExternalOk returns a tuple with the VersionIdIfExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIdIfExternal

`func (o *BTDerivedAssemblyMirror4204) SetVersionIdIfExternal(v string)`

SetVersionIdIfExternal sets VersionIdIfExternal field to given value.

### HasVersionIdIfExternal

`func (o *BTDerivedAssemblyMirror4204) HasVersionIdIfExternal() bool`

HasVersionIdIfExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


