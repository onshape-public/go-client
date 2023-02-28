# BTInsertableInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyType** | Pointer to **string** |  | [optional] 
**ClassType** | Pointer to **int32** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**ConfigurationParameterValues** | Pointer to **[]string** |  | [optional] 
**ConfigurationParameters** | Pointer to **[]string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DeterministicId** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementName** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**FeatureName** | Pointer to **string** |  | [optional] 
**FeatureSpec** | Pointer to **[]string** |  | [optional] 
**FeatureType** | Pointer to **string** |  | [optional] 
**FsComputedPartPropertySpec** | Pointer to **[]string** |  | [optional] 
**FsTableSpec** | Pointer to **[]string** |  | [optional] 
**HasFaults** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**InsertableQuery** | Pointer to **string** |  | [optional] 
**IsFlattenedBody** | Pointer to **bool** |  | [optional] 
**IsMesh** | Pointer to **bool** |  | [optional] 
**MeshState** | Pointer to **string** |  | [optional] 
**MicroversionId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartName** | Pointer to **string** |  | [optional] 
**PredictableThumbnailId** | Pointer to **string** |  | [optional] 
**SourceFileExtension** | Pointer to **string** |  | [optional] 
**ThumbnailUri** | Pointer to **string** |  | [optional] 
**UnflattenedPartDeterministicId** | Pointer to **string** |  | [optional] 
**VariableName** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**WorkspaceName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTInsertableInfo

`func NewBTInsertableInfo() *BTInsertableInfo`

NewBTInsertableInfo instantiates a new BTInsertableInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInsertableInfoWithDefaults

`func NewBTInsertableInfoWithDefaults() *BTInsertableInfo`

NewBTInsertableInfoWithDefaults instantiates a new BTInsertableInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyType

`func (o *BTInsertableInfo) GetBodyType() string`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *BTInsertableInfo) GetBodyTypeOk() (*string, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *BTInsertableInfo) SetBodyType(v string)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *BTInsertableInfo) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetClassType

`func (o *BTInsertableInfo) GetClassType() int32`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *BTInsertableInfo) GetClassTypeOk() (*int32, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *BTInsertableInfo) SetClassType(v int32)`

SetClassType sets ClassType field to given value.

### HasClassType

`func (o *BTInsertableInfo) HasClassType() bool`

HasClassType returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTInsertableInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTInsertableInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTInsertableInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTInsertableInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConfigurationParameterValues

`func (o *BTInsertableInfo) GetConfigurationParameterValues() []string`

GetConfigurationParameterValues returns the ConfigurationParameterValues field if non-nil, zero value otherwise.

### GetConfigurationParameterValuesOk

`func (o *BTInsertableInfo) GetConfigurationParameterValuesOk() (*[]string, bool)`

GetConfigurationParameterValuesOk returns a tuple with the ConfigurationParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameterValues

`func (o *BTInsertableInfo) SetConfigurationParameterValues(v []string)`

SetConfigurationParameterValues sets ConfigurationParameterValues field to given value.

### HasConfigurationParameterValues

`func (o *BTInsertableInfo) HasConfigurationParameterValues() bool`

HasConfigurationParameterValues returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *BTInsertableInfo) GetConfigurationParameters() []string`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *BTInsertableInfo) GetConfigurationParametersOk() (*[]string, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *BTInsertableInfo) SetConfigurationParameters(v []string)`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *BTInsertableInfo) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### GetDataType

`func (o *BTInsertableInfo) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BTInsertableInfo) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BTInsertableInfo) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *BTInsertableInfo) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDeterministicId

`func (o *BTInsertableInfo) GetDeterministicId() string`

GetDeterministicId returns the DeterministicId field if non-nil, zero value otherwise.

### GetDeterministicIdOk

`func (o *BTInsertableInfo) GetDeterministicIdOk() (*string, bool)`

GetDeterministicIdOk returns a tuple with the DeterministicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicId

`func (o *BTInsertableInfo) SetDeterministicId(v string)`

SetDeterministicId sets DeterministicId field to given value.

### HasDeterministicId

`func (o *BTInsertableInfo) HasDeterministicId() bool`

HasDeterministicId returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTInsertableInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTInsertableInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTInsertableInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTInsertableInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTInsertableInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTInsertableInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTInsertableInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTInsertableInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementName

`func (o *BTInsertableInfo) GetElementName() string`

GetElementName returns the ElementName field if non-nil, zero value otherwise.

### GetElementNameOk

`func (o *BTInsertableInfo) GetElementNameOk() (*string, bool)`

GetElementNameOk returns a tuple with the ElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementName

`func (o *BTInsertableInfo) SetElementName(v string)`

SetElementName sets ElementName field to given value.

### HasElementName

`func (o *BTInsertableInfo) HasElementName() bool`

HasElementName returns a boolean if a field has been set.

### GetElementType

`func (o *BTInsertableInfo) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTInsertableInfo) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTInsertableInfo) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTInsertableInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTInsertableInfo) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTInsertableInfo) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTInsertableInfo) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTInsertableInfo) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureName

`func (o *BTInsertableInfo) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *BTInsertableInfo) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *BTInsertableInfo) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *BTInsertableInfo) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetFeatureSpec

`func (o *BTInsertableInfo) GetFeatureSpec() []string`

GetFeatureSpec returns the FeatureSpec field if non-nil, zero value otherwise.

### GetFeatureSpecOk

`func (o *BTInsertableInfo) GetFeatureSpecOk() (*[]string, bool)`

GetFeatureSpecOk returns a tuple with the FeatureSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSpec

`func (o *BTInsertableInfo) SetFeatureSpec(v []string)`

SetFeatureSpec sets FeatureSpec field to given value.

### HasFeatureSpec

`func (o *BTInsertableInfo) HasFeatureSpec() bool`

HasFeatureSpec returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTInsertableInfo) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTInsertableInfo) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTInsertableInfo) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTInsertableInfo) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFsComputedPartPropertySpec

`func (o *BTInsertableInfo) GetFsComputedPartPropertySpec() []string`

GetFsComputedPartPropertySpec returns the FsComputedPartPropertySpec field if non-nil, zero value otherwise.

### GetFsComputedPartPropertySpecOk

`func (o *BTInsertableInfo) GetFsComputedPartPropertySpecOk() (*[]string, bool)`

GetFsComputedPartPropertySpecOk returns a tuple with the FsComputedPartPropertySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsComputedPartPropertySpec

`func (o *BTInsertableInfo) SetFsComputedPartPropertySpec(v []string)`

SetFsComputedPartPropertySpec sets FsComputedPartPropertySpec field to given value.

### HasFsComputedPartPropertySpec

`func (o *BTInsertableInfo) HasFsComputedPartPropertySpec() bool`

HasFsComputedPartPropertySpec returns a boolean if a field has been set.

### GetFsTableSpec

`func (o *BTInsertableInfo) GetFsTableSpec() []string`

GetFsTableSpec returns the FsTableSpec field if non-nil, zero value otherwise.

### GetFsTableSpecOk

`func (o *BTInsertableInfo) GetFsTableSpecOk() (*[]string, bool)`

GetFsTableSpecOk returns a tuple with the FsTableSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsTableSpec

`func (o *BTInsertableInfo) SetFsTableSpec(v []string)`

SetFsTableSpec sets FsTableSpec field to given value.

### HasFsTableSpec

`func (o *BTInsertableInfo) HasFsTableSpec() bool`

HasFsTableSpec returns a boolean if a field has been set.

### GetHasFaults

`func (o *BTInsertableInfo) GetHasFaults() bool`

GetHasFaults returns the HasFaults field if non-nil, zero value otherwise.

### GetHasFaultsOk

`func (o *BTInsertableInfo) GetHasFaultsOk() (*bool, bool)`

GetHasFaultsOk returns a tuple with the HasFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFaults

`func (o *BTInsertableInfo) SetHasFaults(v bool)`

SetHasFaults sets HasFaults field to given value.

### HasHasFaults

`func (o *BTInsertableInfo) HasHasFaults() bool`

HasHasFaults returns a boolean if a field has been set.

### GetHref

`func (o *BTInsertableInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTInsertableInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTInsertableInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTInsertableInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTInsertableInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTInsertableInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTInsertableInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTInsertableInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertableQuery

`func (o *BTInsertableInfo) GetInsertableQuery() string`

GetInsertableQuery returns the InsertableQuery field if non-nil, zero value otherwise.

### GetInsertableQueryOk

`func (o *BTInsertableInfo) GetInsertableQueryOk() (*string, bool)`

GetInsertableQueryOk returns a tuple with the InsertableQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertableQuery

`func (o *BTInsertableInfo) SetInsertableQuery(v string)`

SetInsertableQuery sets InsertableQuery field to given value.

### HasInsertableQuery

`func (o *BTInsertableInfo) HasInsertableQuery() bool`

HasInsertableQuery returns a boolean if a field has been set.

### GetIsFlattenedBody

`func (o *BTInsertableInfo) GetIsFlattenedBody() bool`

GetIsFlattenedBody returns the IsFlattenedBody field if non-nil, zero value otherwise.

### GetIsFlattenedBodyOk

`func (o *BTInsertableInfo) GetIsFlattenedBodyOk() (*bool, bool)`

GetIsFlattenedBodyOk returns a tuple with the IsFlattenedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedBody

`func (o *BTInsertableInfo) SetIsFlattenedBody(v bool)`

SetIsFlattenedBody sets IsFlattenedBody field to given value.

### HasIsFlattenedBody

`func (o *BTInsertableInfo) HasIsFlattenedBody() bool`

HasIsFlattenedBody returns a boolean if a field has been set.

### GetIsMesh

`func (o *BTInsertableInfo) GetIsMesh() bool`

GetIsMesh returns the IsMesh field if non-nil, zero value otherwise.

### GetIsMeshOk

`func (o *BTInsertableInfo) GetIsMeshOk() (*bool, bool)`

GetIsMeshOk returns a tuple with the IsMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMesh

`func (o *BTInsertableInfo) SetIsMesh(v bool)`

SetIsMesh sets IsMesh field to given value.

### HasIsMesh

`func (o *BTInsertableInfo) HasIsMesh() bool`

HasIsMesh returns a boolean if a field has been set.

### GetMeshState

`func (o *BTInsertableInfo) GetMeshState() string`

GetMeshState returns the MeshState field if non-nil, zero value otherwise.

### GetMeshStateOk

`func (o *BTInsertableInfo) GetMeshStateOk() (*string, bool)`

GetMeshStateOk returns a tuple with the MeshState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshState

`func (o *BTInsertableInfo) SetMeshState(v string)`

SetMeshState sets MeshState field to given value.

### HasMeshState

`func (o *BTInsertableInfo) HasMeshState() bool`

HasMeshState returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTInsertableInfo) GetMicroversionId() string`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTInsertableInfo) GetMicroversionIdOk() (*string, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTInsertableInfo) SetMicroversionId(v string)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTInsertableInfo) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetName

`func (o *BTInsertableInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTInsertableInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTInsertableInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTInsertableInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *BTInsertableInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTInsertableInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTInsertableInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTInsertableInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTInsertableInfo) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTInsertableInfo) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTInsertableInfo) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTInsertableInfo) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartName

`func (o *BTInsertableInfo) GetPartName() string`

GetPartName returns the PartName field if non-nil, zero value otherwise.

### GetPartNameOk

`func (o *BTInsertableInfo) GetPartNameOk() (*string, bool)`

GetPartNameOk returns a tuple with the PartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartName

`func (o *BTInsertableInfo) SetPartName(v string)`

SetPartName sets PartName field to given value.

### HasPartName

`func (o *BTInsertableInfo) HasPartName() bool`

HasPartName returns a boolean if a field has been set.

### GetPredictableThumbnailId

`func (o *BTInsertableInfo) GetPredictableThumbnailId() string`

GetPredictableThumbnailId returns the PredictableThumbnailId field if non-nil, zero value otherwise.

### GetPredictableThumbnailIdOk

`func (o *BTInsertableInfo) GetPredictableThumbnailIdOk() (*string, bool)`

GetPredictableThumbnailIdOk returns a tuple with the PredictableThumbnailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictableThumbnailId

`func (o *BTInsertableInfo) SetPredictableThumbnailId(v string)`

SetPredictableThumbnailId sets PredictableThumbnailId field to given value.

### HasPredictableThumbnailId

`func (o *BTInsertableInfo) HasPredictableThumbnailId() bool`

HasPredictableThumbnailId returns a boolean if a field has been set.

### GetSourceFileExtension

`func (o *BTInsertableInfo) GetSourceFileExtension() string`

GetSourceFileExtension returns the SourceFileExtension field if non-nil, zero value otherwise.

### GetSourceFileExtensionOk

`func (o *BTInsertableInfo) GetSourceFileExtensionOk() (*string, bool)`

GetSourceFileExtensionOk returns a tuple with the SourceFileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFileExtension

`func (o *BTInsertableInfo) SetSourceFileExtension(v string)`

SetSourceFileExtension sets SourceFileExtension field to given value.

### HasSourceFileExtension

`func (o *BTInsertableInfo) HasSourceFileExtension() bool`

HasSourceFileExtension returns a boolean if a field has been set.

### GetThumbnailUri

`func (o *BTInsertableInfo) GetThumbnailUri() string`

GetThumbnailUri returns the ThumbnailUri field if non-nil, zero value otherwise.

### GetThumbnailUriOk

`func (o *BTInsertableInfo) GetThumbnailUriOk() (*string, bool)`

GetThumbnailUriOk returns a tuple with the ThumbnailUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUri

`func (o *BTInsertableInfo) SetThumbnailUri(v string)`

SetThumbnailUri sets ThumbnailUri field to given value.

### HasThumbnailUri

`func (o *BTInsertableInfo) HasThumbnailUri() bool`

HasThumbnailUri returns a boolean if a field has been set.

### GetUnflattenedPartDeterministicId

`func (o *BTInsertableInfo) GetUnflattenedPartDeterministicId() string`

GetUnflattenedPartDeterministicId returns the UnflattenedPartDeterministicId field if non-nil, zero value otherwise.

### GetUnflattenedPartDeterministicIdOk

`func (o *BTInsertableInfo) GetUnflattenedPartDeterministicIdOk() (*string, bool)`

GetUnflattenedPartDeterministicIdOk returns a tuple with the UnflattenedPartDeterministicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnflattenedPartDeterministicId

`func (o *BTInsertableInfo) SetUnflattenedPartDeterministicId(v string)`

SetUnflattenedPartDeterministicId sets UnflattenedPartDeterministicId field to given value.

### HasUnflattenedPartDeterministicId

`func (o *BTInsertableInfo) HasUnflattenedPartDeterministicId() bool`

HasUnflattenedPartDeterministicId returns a boolean if a field has been set.

### GetVariableName

`func (o *BTInsertableInfo) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *BTInsertableInfo) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *BTInsertableInfo) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *BTInsertableInfo) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.

### GetVersionId

`func (o *BTInsertableInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTInsertableInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTInsertableInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTInsertableInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *BTInsertableInfo) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *BTInsertableInfo) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *BTInsertableInfo) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *BTInsertableInfo) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetViewRef

`func (o *BTInsertableInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTInsertableInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTInsertableInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTInsertableInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTInsertableInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTInsertableInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTInsertableInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTInsertableInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetWorkspaceName

`func (o *BTInsertableInfo) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *BTInsertableInfo) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *BTInsertableInfo) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *BTInsertableInfo) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


