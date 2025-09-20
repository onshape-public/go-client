# BTInstanceWithReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**CustomData** | Pointer to [**map[string]BTReferenceCustomData1551**](BTReferenceCustomData1551.md) |  | [optional] 
**DerivedAssemblyMirror** | Pointer to **bool** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementReference** | Pointer to [**BTElementReference725**](BTElementReference725.md) |  | [optional] 
**ExternalDocumentWithVersion** | Pointer to [**BTDocumentWithVersionId**](BTDocumentWithVersionId.md) |  | [optional] 
**ExternalDocumentWithVersionAndElementId** | Pointer to [**BTDocumentWithVersionAndElementId**](BTDocumentWithVersionAndElementId.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**LockedState** | Pointer to [**BTInstanceWithReference**](BTInstanceWithReference.md) |  | [optional] 
**MicroversionId** | Pointer to [**BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeWithReferenceList** | Pointer to [**[]BTNodeWithReference**](BTNodeWithReference.md) |  | [optional] 
**ReferenceParameter** | Pointer to [**BTMParameterReferenceWithConfiguration3028**](BTMParameterReferenceWithConfiguration3028.md) |  | [optional] 
**StandardContent** | Pointer to **bool** |  | [optional] 
**StandardContentParametersId** | Pointer to **string** |  | [optional] 
**ValidRevisionReference** | Pointer to **bool** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionIdIfExternal** | Pointer to **string** |  | [optional] 

## Methods

### NewBTInstanceWithReference

`func NewBTInstanceWithReference() *BTInstanceWithReference`

NewBTInstanceWithReference instantiates a new BTInstanceWithReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInstanceWithReferenceWithDefaults

`func NewBTInstanceWithReferenceWithDefaults() *BTInstanceWithReference`

NewBTInstanceWithReferenceWithDefaults instantiates a new BTInstanceWithReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTInstanceWithReference) GetConfiguration() []BTMParameter1`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTInstanceWithReference) GetConfigurationOk() (*[]BTMParameter1, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTInstanceWithReference) SetConfiguration(v []BTMParameter1)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTInstanceWithReference) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCustomData

`func (o *BTInstanceWithReference) GetCustomData() map[string]BTReferenceCustomData1551`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *BTInstanceWithReference) GetCustomDataOk() (*map[string]BTReferenceCustomData1551, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *BTInstanceWithReference) SetCustomData(v map[string]BTReferenceCustomData1551)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *BTInstanceWithReference) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetDerivedAssemblyMirror

`func (o *BTInstanceWithReference) GetDerivedAssemblyMirror() bool`

GetDerivedAssemblyMirror returns the DerivedAssemblyMirror field if non-nil, zero value otherwise.

### GetDerivedAssemblyMirrorOk

`func (o *BTInstanceWithReference) GetDerivedAssemblyMirrorOk() (*bool, bool)`

GetDerivedAssemblyMirrorOk returns a tuple with the DerivedAssemblyMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedAssemblyMirror

`func (o *BTInstanceWithReference) SetDerivedAssemblyMirror(v bool)`

SetDerivedAssemblyMirror sets DerivedAssemblyMirror field to given value.

### HasDerivedAssemblyMirror

`func (o *BTInstanceWithReference) HasDerivedAssemblyMirror() bool`

HasDerivedAssemblyMirror returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTInstanceWithReference) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTInstanceWithReference) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTInstanceWithReference) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTInstanceWithReference) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTInstanceWithReference) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTInstanceWithReference) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTInstanceWithReference) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTInstanceWithReference) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementReference

`func (o *BTInstanceWithReference) GetElementReference() BTElementReference725`

GetElementReference returns the ElementReference field if non-nil, zero value otherwise.

### GetElementReferenceOk

`func (o *BTInstanceWithReference) GetElementReferenceOk() (*BTElementReference725, bool)`

GetElementReferenceOk returns a tuple with the ElementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementReference

`func (o *BTInstanceWithReference) SetElementReference(v BTElementReference725)`

SetElementReference sets ElementReference field to given value.

### HasElementReference

`func (o *BTInstanceWithReference) HasElementReference() bool`

HasElementReference returns a boolean if a field has been set.

### GetExternalDocumentWithVersion

`func (o *BTInstanceWithReference) GetExternalDocumentWithVersion() BTDocumentWithVersionId`

GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionOk

`func (o *BTInstanceWithReference) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool)`

GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersion

`func (o *BTInstanceWithReference) SetExternalDocumentWithVersion(v BTDocumentWithVersionId)`

SetExternalDocumentWithVersion sets ExternalDocumentWithVersion field to given value.

### HasExternalDocumentWithVersion

`func (o *BTInstanceWithReference) HasExternalDocumentWithVersion() bool`

HasExternalDocumentWithVersion returns a boolean if a field has been set.

### GetExternalDocumentWithVersionAndElementId

`func (o *BTInstanceWithReference) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId`

GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionAndElementIdOk

`func (o *BTInstanceWithReference) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool)`

GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersionAndElementId

`func (o *BTInstanceWithReference) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId)`

SetExternalDocumentWithVersionAndElementId sets ExternalDocumentWithVersionAndElementId field to given value.

### HasExternalDocumentWithVersionAndElementId

`func (o *BTInstanceWithReference) HasExternalDocumentWithVersionAndElementId() bool`

HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.

### GetLocked

`func (o *BTInstanceWithReference) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BTInstanceWithReference) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BTInstanceWithReference) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *BTInstanceWithReference) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedState

`func (o *BTInstanceWithReference) GetLockedState() BTInstanceWithReference`

GetLockedState returns the LockedState field if non-nil, zero value otherwise.

### GetLockedStateOk

`func (o *BTInstanceWithReference) GetLockedStateOk() (*BTInstanceWithReference, bool)`

GetLockedStateOk returns a tuple with the LockedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedState

`func (o *BTInstanceWithReference) SetLockedState(v BTInstanceWithReference)`

SetLockedState sets LockedState field to given value.

### HasLockedState

`func (o *BTInstanceWithReference) HasLockedState() bool`

HasLockedState returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTInstanceWithReference) GetMicroversionId() BTMicroversionId366`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTInstanceWithReference) GetMicroversionIdOk() (*BTMicroversionId366, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTInstanceWithReference) SetMicroversionId(v BTMicroversionId366)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTInstanceWithReference) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetNodeId

`func (o *BTInstanceWithReference) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTInstanceWithReference) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTInstanceWithReference) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTInstanceWithReference) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeWithReferenceList

`func (o *BTInstanceWithReference) GetNodeWithReferenceList() []BTNodeWithReference`

GetNodeWithReferenceList returns the NodeWithReferenceList field if non-nil, zero value otherwise.

### GetNodeWithReferenceListOk

`func (o *BTInstanceWithReference) GetNodeWithReferenceListOk() (*[]BTNodeWithReference, bool)`

GetNodeWithReferenceListOk returns a tuple with the NodeWithReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWithReferenceList

`func (o *BTInstanceWithReference) SetNodeWithReferenceList(v []BTNodeWithReference)`

SetNodeWithReferenceList sets NodeWithReferenceList field to given value.

### HasNodeWithReferenceList

`func (o *BTInstanceWithReference) HasNodeWithReferenceList() bool`

HasNodeWithReferenceList returns a boolean if a field has been set.

### GetReferenceParameter

`func (o *BTInstanceWithReference) GetReferenceParameter() BTMParameterReferenceWithConfiguration3028`

GetReferenceParameter returns the ReferenceParameter field if non-nil, zero value otherwise.

### GetReferenceParameterOk

`func (o *BTInstanceWithReference) GetReferenceParameterOk() (*BTMParameterReferenceWithConfiguration3028, bool)`

GetReferenceParameterOk returns a tuple with the ReferenceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceParameter

`func (o *BTInstanceWithReference) SetReferenceParameter(v BTMParameterReferenceWithConfiguration3028)`

SetReferenceParameter sets ReferenceParameter field to given value.

### HasReferenceParameter

`func (o *BTInstanceWithReference) HasReferenceParameter() bool`

HasReferenceParameter returns a boolean if a field has been set.

### GetStandardContent

`func (o *BTInstanceWithReference) GetStandardContent() bool`

GetStandardContent returns the StandardContent field if non-nil, zero value otherwise.

### GetStandardContentOk

`func (o *BTInstanceWithReference) GetStandardContentOk() (*bool, bool)`

GetStandardContentOk returns a tuple with the StandardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContent

`func (o *BTInstanceWithReference) SetStandardContent(v bool)`

SetStandardContent sets StandardContent field to given value.

### HasStandardContent

`func (o *BTInstanceWithReference) HasStandardContent() bool`

HasStandardContent returns a boolean if a field has been set.

### GetStandardContentParametersId

`func (o *BTInstanceWithReference) GetStandardContentParametersId() string`

GetStandardContentParametersId returns the StandardContentParametersId field if non-nil, zero value otherwise.

### GetStandardContentParametersIdOk

`func (o *BTInstanceWithReference) GetStandardContentParametersIdOk() (*string, bool)`

GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentParametersId

`func (o *BTInstanceWithReference) SetStandardContentParametersId(v string)`

SetStandardContentParametersId sets StandardContentParametersId field to given value.

### HasStandardContentParametersId

`func (o *BTInstanceWithReference) HasStandardContentParametersId() bool`

HasStandardContentParametersId returns a boolean if a field has been set.

### GetValidRevisionReference

`func (o *BTInstanceWithReference) GetValidRevisionReference() bool`

GetValidRevisionReference returns the ValidRevisionReference field if non-nil, zero value otherwise.

### GetValidRevisionReferenceOk

`func (o *BTInstanceWithReference) GetValidRevisionReferenceOk() (*bool, bool)`

GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRevisionReference

`func (o *BTInstanceWithReference) SetValidRevisionReference(v bool)`

SetValidRevisionReference sets ValidRevisionReference field to given value.

### HasValidRevisionReference

`func (o *BTInstanceWithReference) HasValidRevisionReference() bool`

HasValidRevisionReference returns a boolean if a field has been set.

### GetVersionId

`func (o *BTInstanceWithReference) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTInstanceWithReference) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTInstanceWithReference) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTInstanceWithReference) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionIdIfExternal

`func (o *BTInstanceWithReference) GetVersionIdIfExternal() string`

GetVersionIdIfExternal returns the VersionIdIfExternal field if non-nil, zero value otherwise.

### GetVersionIdIfExternalOk

`func (o *BTInstanceWithReference) GetVersionIdIfExternalOk() (*string, bool)`

GetVersionIdIfExternalOk returns a tuple with the VersionIdIfExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIdIfExternal

`func (o *BTInstanceWithReference) SetVersionIdIfExternal(v string)`

SetVersionIdIfExternal sets VersionIdIfExternal field to given value.

### HasVersionIdIfExternal

`func (o *BTInstanceWithReference) HasVersionIdIfExternal() bool`

HasVersionIdIfExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


