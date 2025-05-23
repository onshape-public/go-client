# BTInstance642

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Configuration** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**Configured** | Pointer to **bool** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementReference** | Pointer to [**BTElementReference725**](BTElementReference725.md) |  | [optional] 
**ExternalDocumentWithVersion** | Pointer to [**BTDocumentWithVersionId**](BTDocumentWithVersionId.md) |  | [optional] 
**ExternalDocumentWithVersionAndElementId** | Pointer to [**BTDocumentWithVersionAndElementId**](BTDocumentWithVersionAndElementId.md) |  | [optional] 
**LockedState** | Pointer to [**BTMParameter1**](BTMParameter1.md) |  | [optional] 
**MicroversionId** | Pointer to [**BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeWithReferenceList** | Pointer to [**[]BTNodeWithReference**](BTNodeWithReference.md) |  | [optional] 
**Parameters** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**ReferenceParameter** | Pointer to [**BTMParameterReferenceWithConfiguration3028**](BTMParameterReferenceWithConfiguration3028.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionIdIfExternal** | Pointer to **string** |  | [optional] 

## Methods

### NewBTInstance642

`func NewBTInstance642() *BTInstance642`

NewBTInstance642 instantiates a new BTInstance642 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInstance642WithDefaults

`func NewBTInstance642WithDefaults() *BTInstance642`

NewBTInstance642WithDefaults instantiates a new BTInstance642 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTInstance642) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTInstance642) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTInstance642) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTInstance642) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTInstance642) GetConfiguration() []BTMParameter1`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTInstance642) GetConfigurationOk() (*[]BTMParameter1, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTInstance642) SetConfiguration(v []BTMParameter1)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTInstance642) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConfigured

`func (o *BTInstance642) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *BTInstance642) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *BTInstance642) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *BTInstance642) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTInstance642) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTInstance642) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTInstance642) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTInstance642) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTInstance642) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTInstance642) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTInstance642) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTInstance642) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementReference

`func (o *BTInstance642) GetElementReference() BTElementReference725`

GetElementReference returns the ElementReference field if non-nil, zero value otherwise.

### GetElementReferenceOk

`func (o *BTInstance642) GetElementReferenceOk() (*BTElementReference725, bool)`

GetElementReferenceOk returns a tuple with the ElementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementReference

`func (o *BTInstance642) SetElementReference(v BTElementReference725)`

SetElementReference sets ElementReference field to given value.

### HasElementReference

`func (o *BTInstance642) HasElementReference() bool`

HasElementReference returns a boolean if a field has been set.

### GetExternalDocumentWithVersion

`func (o *BTInstance642) GetExternalDocumentWithVersion() BTDocumentWithVersionId`

GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionOk

`func (o *BTInstance642) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool)`

GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersion

`func (o *BTInstance642) SetExternalDocumentWithVersion(v BTDocumentWithVersionId)`

SetExternalDocumentWithVersion sets ExternalDocumentWithVersion field to given value.

### HasExternalDocumentWithVersion

`func (o *BTInstance642) HasExternalDocumentWithVersion() bool`

HasExternalDocumentWithVersion returns a boolean if a field has been set.

### GetExternalDocumentWithVersionAndElementId

`func (o *BTInstance642) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId`

GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionAndElementIdOk

`func (o *BTInstance642) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool)`

GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersionAndElementId

`func (o *BTInstance642) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId)`

SetExternalDocumentWithVersionAndElementId sets ExternalDocumentWithVersionAndElementId field to given value.

### HasExternalDocumentWithVersionAndElementId

`func (o *BTInstance642) HasExternalDocumentWithVersionAndElementId() bool`

HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.

### GetLockedState

`func (o *BTInstance642) GetLockedState() BTMParameter1`

GetLockedState returns the LockedState field if non-nil, zero value otherwise.

### GetLockedStateOk

`func (o *BTInstance642) GetLockedStateOk() (*BTMParameter1, bool)`

GetLockedStateOk returns a tuple with the LockedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedState

`func (o *BTInstance642) SetLockedState(v BTMParameter1)`

SetLockedState sets LockedState field to given value.

### HasLockedState

`func (o *BTInstance642) HasLockedState() bool`

HasLockedState returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTInstance642) GetMicroversionId() BTMicroversionId366`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTInstance642) GetMicroversionIdOk() (*BTMicroversionId366, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTInstance642) SetMicroversionId(v BTMicroversionId366)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTInstance642) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetName

`func (o *BTInstance642) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTInstance642) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTInstance642) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTInstance642) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeWithReferenceList

`func (o *BTInstance642) GetNodeWithReferenceList() []BTNodeWithReference`

GetNodeWithReferenceList returns the NodeWithReferenceList field if non-nil, zero value otherwise.

### GetNodeWithReferenceListOk

`func (o *BTInstance642) GetNodeWithReferenceListOk() (*[]BTNodeWithReference, bool)`

GetNodeWithReferenceListOk returns a tuple with the NodeWithReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWithReferenceList

`func (o *BTInstance642) SetNodeWithReferenceList(v []BTNodeWithReference)`

SetNodeWithReferenceList sets NodeWithReferenceList field to given value.

### HasNodeWithReferenceList

`func (o *BTInstance642) HasNodeWithReferenceList() bool`

HasNodeWithReferenceList returns a boolean if a field has been set.

### GetParameters

`func (o *BTInstance642) GetParameters() []BTMParameter1`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTInstance642) GetParametersOk() (*[]BTMParameter1, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTInstance642) SetParameters(v []BTMParameter1)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTInstance642) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReferenceParameter

`func (o *BTInstance642) GetReferenceParameter() BTMParameterReferenceWithConfiguration3028`

GetReferenceParameter returns the ReferenceParameter field if non-nil, zero value otherwise.

### GetReferenceParameterOk

`func (o *BTInstance642) GetReferenceParameterOk() (*BTMParameterReferenceWithConfiguration3028, bool)`

GetReferenceParameterOk returns a tuple with the ReferenceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceParameter

`func (o *BTInstance642) SetReferenceParameter(v BTMParameterReferenceWithConfiguration3028)`

SetReferenceParameter sets ReferenceParameter field to given value.

### HasReferenceParameter

`func (o *BTInstance642) HasReferenceParameter() bool`

HasReferenceParameter returns a boolean if a field has been set.

### GetVersionId

`func (o *BTInstance642) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTInstance642) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTInstance642) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTInstance642) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionIdIfExternal

`func (o *BTInstance642) GetVersionIdIfExternal() string`

GetVersionIdIfExternal returns the VersionIdIfExternal field if non-nil, zero value otherwise.

### GetVersionIdIfExternalOk

`func (o *BTInstance642) GetVersionIdIfExternalOk() (*string, bool)`

GetVersionIdIfExternalOk returns a tuple with the VersionIdIfExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIdIfExternal

`func (o *BTInstance642) SetVersionIdIfExternal(v string)`

SetVersionIdIfExternal sets VersionIdIfExternal field to given value.

### HasVersionIdIfExternal

`func (o *BTInstance642) HasVersionIdIfExternal() bool`

HasVersionIdIfExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


