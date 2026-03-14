# BTMSimulationFeature5034

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxiliaryTreeFeature** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**FeatureFolder** | Pointer to **bool** |  | [optional] 
**FeatureId** | Pointer to **string** | Unique ID of the feature instance within this Part Studio. | [optional] 
**FeatureListFieldIndex** | Pointer to **int32** |  | [optional] 
**FeatureType** | Pointer to **string** | The name of the feature spec that this feature instantiates. | [optional] 
**FieldIndexForOwnedMateConnectors** | Pointer to **int32** |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Element microversion that is being imported. | [optional] 
**MateConnectorFeature** | Pointer to **bool** |  | [optional] 
**MateConnectors** | Pointer to [**[]BTMMateConnector66**](BTMMateConnector66.md) |  | [optional] 
**Name** | Pointer to **string** | User-visible name of the feature. | [optional] 
**Namespace** | Pointer to **string** | Indicates where the feature definition lives. Features in the FeatureScript standard library have a namespace value of &#x60;\&quot;\&quot;&#x60;. Custom features identify the Feature Studio that contains the definition. | [optional] 
**NodeId** | Pointer to **string** | ID for the feature node. | [optional] 
**OccurrenceQueriesFromAllConfigurations** | Pointer to [**[]BTMIndividualQueryWithOccurrenceBase904**](BTMIndividualQueryWithOccurrenceBase904.md) |  | [optional] 
**ParameterLibraries** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**ParametricInstanceFeature** | Pointer to **bool** |  | [optional] 
**ParentSuppressed** | Pointer to **bool** |  | [optional] 
**ReturnAfterSubfeatures** | Pointer to **bool** | For internal use only. Should always be &#x60;false&#x60;. | [optional] 
**SubFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) | List of subfeatures belonging to the feature. | [optional] 
**SubFeaturesNotUsedInQuery** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**Suppressed** | Pointer to **bool** | If &#x60;true&#x60;, the feature is suppressed. It will skip regeneration, denoted by a line through the name in the Feature list. | [optional] 
**SuppressionConfigured** | Pointer to **bool** | &#x60;true&#x60; if the suppression is configured in the Part Studio. | [optional] 
**SuppressionState** | Pointer to [**BTMSuppressionState1924**](BTMSuppressionState1924.md) |  | [optional] 
**VariableStudioReference** | Pointer to **bool** | If &#x60;true&#x60;, the feature references a Variable Studio. | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTMSimulationFeature5034

`func NewBTMSimulationFeature5034() *BTMSimulationFeature5034`

NewBTMSimulationFeature5034 instantiates a new BTMSimulationFeature5034 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMSimulationFeature5034WithDefaults

`func NewBTMSimulationFeature5034WithDefaults() *BTMSimulationFeature5034`

NewBTMSimulationFeature5034WithDefaults instantiates a new BTMSimulationFeature5034 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxiliaryTreeFeature

`func (o *BTMSimulationFeature5034) GetAuxiliaryTreeFeature() bool`

GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field if non-nil, zero value otherwise.

### GetAuxiliaryTreeFeatureOk

`func (o *BTMSimulationFeature5034) GetAuxiliaryTreeFeatureOk() (*bool, bool)`

GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTreeFeature

`func (o *BTMSimulationFeature5034) SetAuxiliaryTreeFeature(v bool)`

SetAuxiliaryTreeFeature sets AuxiliaryTreeFeature field to given value.

### HasAuxiliaryTreeFeature

`func (o *BTMSimulationFeature5034) HasAuxiliaryTreeFeature() bool`

HasAuxiliaryTreeFeature returns a boolean if a field has been set.

### GetBtType

`func (o *BTMSimulationFeature5034) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMSimulationFeature5034) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMSimulationFeature5034) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMSimulationFeature5034) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatureFolder

`func (o *BTMSimulationFeature5034) GetFeatureFolder() bool`

GetFeatureFolder returns the FeatureFolder field if non-nil, zero value otherwise.

### GetFeatureFolderOk

`func (o *BTMSimulationFeature5034) GetFeatureFolderOk() (*bool, bool)`

GetFeatureFolderOk returns a tuple with the FeatureFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFolder

`func (o *BTMSimulationFeature5034) SetFeatureFolder(v bool)`

SetFeatureFolder sets FeatureFolder field to given value.

### HasFeatureFolder

`func (o *BTMSimulationFeature5034) HasFeatureFolder() bool`

HasFeatureFolder returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTMSimulationFeature5034) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTMSimulationFeature5034) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTMSimulationFeature5034) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTMSimulationFeature5034) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureListFieldIndex

`func (o *BTMSimulationFeature5034) GetFeatureListFieldIndex() int32`

GetFeatureListFieldIndex returns the FeatureListFieldIndex field if non-nil, zero value otherwise.

### GetFeatureListFieldIndexOk

`func (o *BTMSimulationFeature5034) GetFeatureListFieldIndexOk() (*int32, bool)`

GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureListFieldIndex

`func (o *BTMSimulationFeature5034) SetFeatureListFieldIndex(v int32)`

SetFeatureListFieldIndex sets FeatureListFieldIndex field to given value.

### HasFeatureListFieldIndex

`func (o *BTMSimulationFeature5034) HasFeatureListFieldIndex() bool`

HasFeatureListFieldIndex returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTMSimulationFeature5034) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTMSimulationFeature5034) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTMSimulationFeature5034) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTMSimulationFeature5034) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFieldIndexForOwnedMateConnectors

`func (o *BTMSimulationFeature5034) GetFieldIndexForOwnedMateConnectors() int32`

GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field if non-nil, zero value otherwise.

### GetFieldIndexForOwnedMateConnectorsOk

`func (o *BTMSimulationFeature5034) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool)`

GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIndexForOwnedMateConnectors

`func (o *BTMSimulationFeature5034) SetFieldIndexForOwnedMateConnectors(v int32)`

SetFieldIndexForOwnedMateConnectors sets FieldIndexForOwnedMateConnectors field to given value.

### HasFieldIndexForOwnedMateConnectors

`func (o *BTMSimulationFeature5034) HasFieldIndexForOwnedMateConnectors() bool`

HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTMSimulationFeature5034) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTMSimulationFeature5034) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTMSimulationFeature5034) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTMSimulationFeature5034) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetMateConnectorFeature

`func (o *BTMSimulationFeature5034) GetMateConnectorFeature() bool`

GetMateConnectorFeature returns the MateConnectorFeature field if non-nil, zero value otherwise.

### GetMateConnectorFeatureOk

`func (o *BTMSimulationFeature5034) GetMateConnectorFeatureOk() (*bool, bool)`

GetMateConnectorFeatureOk returns a tuple with the MateConnectorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectorFeature

`func (o *BTMSimulationFeature5034) SetMateConnectorFeature(v bool)`

SetMateConnectorFeature sets MateConnectorFeature field to given value.

### HasMateConnectorFeature

`func (o *BTMSimulationFeature5034) HasMateConnectorFeature() bool`

HasMateConnectorFeature returns a boolean if a field has been set.

### GetMateConnectors

`func (o *BTMSimulationFeature5034) GetMateConnectors() []BTMMateConnector66`

GetMateConnectors returns the MateConnectors field if non-nil, zero value otherwise.

### GetMateConnectorsOk

`func (o *BTMSimulationFeature5034) GetMateConnectorsOk() (*[]BTMMateConnector66, bool)`

GetMateConnectorsOk returns a tuple with the MateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectors

`func (o *BTMSimulationFeature5034) SetMateConnectors(v []BTMMateConnector66)`

SetMateConnectors sets MateConnectors field to given value.

### HasMateConnectors

`func (o *BTMSimulationFeature5034) HasMateConnectors() bool`

HasMateConnectors returns a boolean if a field has been set.

### GetName

`func (o *BTMSimulationFeature5034) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMSimulationFeature5034) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMSimulationFeature5034) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMSimulationFeature5034) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTMSimulationFeature5034) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTMSimulationFeature5034) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTMSimulationFeature5034) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTMSimulationFeature5034) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeId

`func (o *BTMSimulationFeature5034) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTMSimulationFeature5034) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTMSimulationFeature5034) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTMSimulationFeature5034) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOccurrenceQueriesFromAllConfigurations

`func (o *BTMSimulationFeature5034) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904`

GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field if non-nil, zero value otherwise.

### GetOccurrenceQueriesFromAllConfigurationsOk

`func (o *BTMSimulationFeature5034) GetOccurrenceQueriesFromAllConfigurationsOk() (*[]BTMIndividualQueryWithOccurrenceBase904, bool)`

GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceQueriesFromAllConfigurations

`func (o *BTMSimulationFeature5034) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904)`

SetOccurrenceQueriesFromAllConfigurations sets OccurrenceQueriesFromAllConfigurations field to given value.

### HasOccurrenceQueriesFromAllConfigurations

`func (o *BTMSimulationFeature5034) HasOccurrenceQueriesFromAllConfigurations() bool`

HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.

### GetParameterLibraries

`func (o *BTMSimulationFeature5034) GetParameterLibraries() []BTMParameter1`

GetParameterLibraries returns the ParameterLibraries field if non-nil, zero value otherwise.

### GetParameterLibrariesOk

`func (o *BTMSimulationFeature5034) GetParameterLibrariesOk() (*[]BTMParameter1, bool)`

GetParameterLibrariesOk returns a tuple with the ParameterLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLibraries

`func (o *BTMSimulationFeature5034) SetParameterLibraries(v []BTMParameter1)`

SetParameterLibraries sets ParameterLibraries field to given value.

### HasParameterLibraries

`func (o *BTMSimulationFeature5034) HasParameterLibraries() bool`

HasParameterLibraries returns a boolean if a field has been set.

### GetParametricInstanceFeature

`func (o *BTMSimulationFeature5034) GetParametricInstanceFeature() bool`

GetParametricInstanceFeature returns the ParametricInstanceFeature field if non-nil, zero value otherwise.

### GetParametricInstanceFeatureOk

`func (o *BTMSimulationFeature5034) GetParametricInstanceFeatureOk() (*bool, bool)`

GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstanceFeature

`func (o *BTMSimulationFeature5034) SetParametricInstanceFeature(v bool)`

SetParametricInstanceFeature sets ParametricInstanceFeature field to given value.

### HasParametricInstanceFeature

`func (o *BTMSimulationFeature5034) HasParametricInstanceFeature() bool`

HasParametricInstanceFeature returns a boolean if a field has been set.

### GetParentSuppressed

`func (o *BTMSimulationFeature5034) GetParentSuppressed() bool`

GetParentSuppressed returns the ParentSuppressed field if non-nil, zero value otherwise.

### GetParentSuppressedOk

`func (o *BTMSimulationFeature5034) GetParentSuppressedOk() (*bool, bool)`

GetParentSuppressedOk returns a tuple with the ParentSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSuppressed

`func (o *BTMSimulationFeature5034) SetParentSuppressed(v bool)`

SetParentSuppressed sets ParentSuppressed field to given value.

### HasParentSuppressed

`func (o *BTMSimulationFeature5034) HasParentSuppressed() bool`

HasParentSuppressed returns a boolean if a field has been set.

### GetReturnAfterSubfeatures

`func (o *BTMSimulationFeature5034) GetReturnAfterSubfeatures() bool`

GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field if non-nil, zero value otherwise.

### GetReturnAfterSubfeaturesOk

`func (o *BTMSimulationFeature5034) GetReturnAfterSubfeaturesOk() (*bool, bool)`

GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAfterSubfeatures

`func (o *BTMSimulationFeature5034) SetReturnAfterSubfeatures(v bool)`

SetReturnAfterSubfeatures sets ReturnAfterSubfeatures field to given value.

### HasReturnAfterSubfeatures

`func (o *BTMSimulationFeature5034) HasReturnAfterSubfeatures() bool`

HasReturnAfterSubfeatures returns a boolean if a field has been set.

### GetSubFeatures

`func (o *BTMSimulationFeature5034) GetSubFeatures() []BTMFeature134`

GetSubFeatures returns the SubFeatures field if non-nil, zero value otherwise.

### GetSubFeaturesOk

`func (o *BTMSimulationFeature5034) GetSubFeaturesOk() (*[]BTMFeature134, bool)`

GetSubFeaturesOk returns a tuple with the SubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeatures

`func (o *BTMSimulationFeature5034) SetSubFeatures(v []BTMFeature134)`

SetSubFeatures sets SubFeatures field to given value.

### HasSubFeatures

`func (o *BTMSimulationFeature5034) HasSubFeatures() bool`

HasSubFeatures returns a boolean if a field has been set.

### GetSubFeaturesNotUsedInQuery

`func (o *BTMSimulationFeature5034) GetSubFeaturesNotUsedInQuery() []BTMFeature134`

GetSubFeaturesNotUsedInQuery returns the SubFeaturesNotUsedInQuery field if non-nil, zero value otherwise.

### GetSubFeaturesNotUsedInQueryOk

`func (o *BTMSimulationFeature5034) GetSubFeaturesNotUsedInQueryOk() (*[]BTMFeature134, bool)`

GetSubFeaturesNotUsedInQueryOk returns a tuple with the SubFeaturesNotUsedInQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeaturesNotUsedInQuery

`func (o *BTMSimulationFeature5034) SetSubFeaturesNotUsedInQuery(v []BTMFeature134)`

SetSubFeaturesNotUsedInQuery sets SubFeaturesNotUsedInQuery field to given value.

### HasSubFeaturesNotUsedInQuery

`func (o *BTMSimulationFeature5034) HasSubFeaturesNotUsedInQuery() bool`

HasSubFeaturesNotUsedInQuery returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTMSimulationFeature5034) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTMSimulationFeature5034) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTMSimulationFeature5034) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTMSimulationFeature5034) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTMSimulationFeature5034) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTMSimulationFeature5034) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTMSimulationFeature5034) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTMSimulationFeature5034) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTMSimulationFeature5034) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTMSimulationFeature5034) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTMSimulationFeature5034) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTMSimulationFeature5034) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.

### GetVariableStudioReference

`func (o *BTMSimulationFeature5034) GetVariableStudioReference() bool`

GetVariableStudioReference returns the VariableStudioReference field if non-nil, zero value otherwise.

### GetVariableStudioReferenceOk

`func (o *BTMSimulationFeature5034) GetVariableStudioReferenceOk() (*bool, bool)`

GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudioReference

`func (o *BTMSimulationFeature5034) SetVariableStudioReference(v bool)`

SetVariableStudioReference sets VariableStudioReference field to given value.

### HasVariableStudioReference

`func (o *BTMSimulationFeature5034) HasVariableStudioReference() bool`

HasVariableStudioReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTMSimulationFeature5034) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTMSimulationFeature5034) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTMSimulationFeature5034) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTMSimulationFeature5034) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


