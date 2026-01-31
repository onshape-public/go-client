# BTAssemblySimulation2246

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
**InertialRelief** | Pointer to **bool** |  | [optional] 
**MateConnectorFeature** | Pointer to **bool** |  | [optional] 
**MateConnectors** | Pointer to [**[]BTMMateConnector66**](BTMMateConnector66.md) |  | [optional] 
**Modal** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | User-visible name of the feature. | [optional] 
**Namespace** | Pointer to **string** | Indicates where the feature definition lives. Features in the FeatureScript standard library have a namespace value of &#x60;\&quot;\&quot;&#x60;. Custom features identify the Feature Studio that contains the definition. | [optional] 
**NodeId** | Pointer to **string** | ID for the feature node. | [optional] 
**NumberOfModes** | Pointer to **float64** |  | [optional] 
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
**Type** | Pointer to [**GBTAssemblySimulationType**](GBTAssemblySimulationType.md) |  | [optional] 
**VariableStudioReference** | Pointer to **bool** | If &#x60;true&#x60;, the feature references a Variable Studio. | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTAssemblySimulation2246

`func NewBTAssemblySimulation2246() *BTAssemblySimulation2246`

NewBTAssemblySimulation2246 instantiates a new BTAssemblySimulation2246 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblySimulation2246WithDefaults

`func NewBTAssemblySimulation2246WithDefaults() *BTAssemblySimulation2246`

NewBTAssemblySimulation2246WithDefaults instantiates a new BTAssemblySimulation2246 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxiliaryTreeFeature

`func (o *BTAssemblySimulation2246) GetAuxiliaryTreeFeature() bool`

GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field if non-nil, zero value otherwise.

### GetAuxiliaryTreeFeatureOk

`func (o *BTAssemblySimulation2246) GetAuxiliaryTreeFeatureOk() (*bool, bool)`

GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTreeFeature

`func (o *BTAssemblySimulation2246) SetAuxiliaryTreeFeature(v bool)`

SetAuxiliaryTreeFeature sets AuxiliaryTreeFeature field to given value.

### HasAuxiliaryTreeFeature

`func (o *BTAssemblySimulation2246) HasAuxiliaryTreeFeature() bool`

HasAuxiliaryTreeFeature returns a boolean if a field has been set.

### GetBtType

`func (o *BTAssemblySimulation2246) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTAssemblySimulation2246) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTAssemblySimulation2246) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTAssemblySimulation2246) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatureFolder

`func (o *BTAssemblySimulation2246) GetFeatureFolder() bool`

GetFeatureFolder returns the FeatureFolder field if non-nil, zero value otherwise.

### GetFeatureFolderOk

`func (o *BTAssemblySimulation2246) GetFeatureFolderOk() (*bool, bool)`

GetFeatureFolderOk returns a tuple with the FeatureFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFolder

`func (o *BTAssemblySimulation2246) SetFeatureFolder(v bool)`

SetFeatureFolder sets FeatureFolder field to given value.

### HasFeatureFolder

`func (o *BTAssemblySimulation2246) HasFeatureFolder() bool`

HasFeatureFolder returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTAssemblySimulation2246) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTAssemblySimulation2246) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTAssemblySimulation2246) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTAssemblySimulation2246) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureListFieldIndex

`func (o *BTAssemblySimulation2246) GetFeatureListFieldIndex() int32`

GetFeatureListFieldIndex returns the FeatureListFieldIndex field if non-nil, zero value otherwise.

### GetFeatureListFieldIndexOk

`func (o *BTAssemblySimulation2246) GetFeatureListFieldIndexOk() (*int32, bool)`

GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureListFieldIndex

`func (o *BTAssemblySimulation2246) SetFeatureListFieldIndex(v int32)`

SetFeatureListFieldIndex sets FeatureListFieldIndex field to given value.

### HasFeatureListFieldIndex

`func (o *BTAssemblySimulation2246) HasFeatureListFieldIndex() bool`

HasFeatureListFieldIndex returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTAssemblySimulation2246) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTAssemblySimulation2246) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTAssemblySimulation2246) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTAssemblySimulation2246) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFieldIndexForOwnedMateConnectors

`func (o *BTAssemblySimulation2246) GetFieldIndexForOwnedMateConnectors() int32`

GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field if non-nil, zero value otherwise.

### GetFieldIndexForOwnedMateConnectorsOk

`func (o *BTAssemblySimulation2246) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool)`

GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIndexForOwnedMateConnectors

`func (o *BTAssemblySimulation2246) SetFieldIndexForOwnedMateConnectors(v int32)`

SetFieldIndexForOwnedMateConnectors sets FieldIndexForOwnedMateConnectors field to given value.

### HasFieldIndexForOwnedMateConnectors

`func (o *BTAssemblySimulation2246) HasFieldIndexForOwnedMateConnectors() bool`

HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTAssemblySimulation2246) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTAssemblySimulation2246) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTAssemblySimulation2246) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTAssemblySimulation2246) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetInertialRelief

`func (o *BTAssemblySimulation2246) GetInertialRelief() bool`

GetInertialRelief returns the InertialRelief field if non-nil, zero value otherwise.

### GetInertialReliefOk

`func (o *BTAssemblySimulation2246) GetInertialReliefOk() (*bool, bool)`

GetInertialReliefOk returns a tuple with the InertialRelief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInertialRelief

`func (o *BTAssemblySimulation2246) SetInertialRelief(v bool)`

SetInertialRelief sets InertialRelief field to given value.

### HasInertialRelief

`func (o *BTAssemblySimulation2246) HasInertialRelief() bool`

HasInertialRelief returns a boolean if a field has been set.

### GetMateConnectorFeature

`func (o *BTAssemblySimulation2246) GetMateConnectorFeature() bool`

GetMateConnectorFeature returns the MateConnectorFeature field if non-nil, zero value otherwise.

### GetMateConnectorFeatureOk

`func (o *BTAssemblySimulation2246) GetMateConnectorFeatureOk() (*bool, bool)`

GetMateConnectorFeatureOk returns a tuple with the MateConnectorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectorFeature

`func (o *BTAssemblySimulation2246) SetMateConnectorFeature(v bool)`

SetMateConnectorFeature sets MateConnectorFeature field to given value.

### HasMateConnectorFeature

`func (o *BTAssemblySimulation2246) HasMateConnectorFeature() bool`

HasMateConnectorFeature returns a boolean if a field has been set.

### GetMateConnectors

`func (o *BTAssemblySimulation2246) GetMateConnectors() []BTMMateConnector66`

GetMateConnectors returns the MateConnectors field if non-nil, zero value otherwise.

### GetMateConnectorsOk

`func (o *BTAssemblySimulation2246) GetMateConnectorsOk() (*[]BTMMateConnector66, bool)`

GetMateConnectorsOk returns a tuple with the MateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectors

`func (o *BTAssemblySimulation2246) SetMateConnectors(v []BTMMateConnector66)`

SetMateConnectors sets MateConnectors field to given value.

### HasMateConnectors

`func (o *BTAssemblySimulation2246) HasMateConnectors() bool`

HasMateConnectors returns a boolean if a field has been set.

### GetModal

`func (o *BTAssemblySimulation2246) GetModal() bool`

GetModal returns the Modal field if non-nil, zero value otherwise.

### GetModalOk

`func (o *BTAssemblySimulation2246) GetModalOk() (*bool, bool)`

GetModalOk returns a tuple with the Modal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModal

`func (o *BTAssemblySimulation2246) SetModal(v bool)`

SetModal sets Modal field to given value.

### HasModal

`func (o *BTAssemblySimulation2246) HasModal() bool`

HasModal returns a boolean if a field has been set.

### GetName

`func (o *BTAssemblySimulation2246) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAssemblySimulation2246) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAssemblySimulation2246) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAssemblySimulation2246) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTAssemblySimulation2246) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTAssemblySimulation2246) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTAssemblySimulation2246) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTAssemblySimulation2246) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeId

`func (o *BTAssemblySimulation2246) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTAssemblySimulation2246) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTAssemblySimulation2246) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTAssemblySimulation2246) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNumberOfModes

`func (o *BTAssemblySimulation2246) GetNumberOfModes() float64`

GetNumberOfModes returns the NumberOfModes field if non-nil, zero value otherwise.

### GetNumberOfModesOk

`func (o *BTAssemblySimulation2246) GetNumberOfModesOk() (*float64, bool)`

GetNumberOfModesOk returns a tuple with the NumberOfModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfModes

`func (o *BTAssemblySimulation2246) SetNumberOfModes(v float64)`

SetNumberOfModes sets NumberOfModes field to given value.

### HasNumberOfModes

`func (o *BTAssemblySimulation2246) HasNumberOfModes() bool`

HasNumberOfModes returns a boolean if a field has been set.

### GetOccurrenceQueriesFromAllConfigurations

`func (o *BTAssemblySimulation2246) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904`

GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field if non-nil, zero value otherwise.

### GetOccurrenceQueriesFromAllConfigurationsOk

`func (o *BTAssemblySimulation2246) GetOccurrenceQueriesFromAllConfigurationsOk() (*[]BTMIndividualQueryWithOccurrenceBase904, bool)`

GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceQueriesFromAllConfigurations

`func (o *BTAssemblySimulation2246) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904)`

SetOccurrenceQueriesFromAllConfigurations sets OccurrenceQueriesFromAllConfigurations field to given value.

### HasOccurrenceQueriesFromAllConfigurations

`func (o *BTAssemblySimulation2246) HasOccurrenceQueriesFromAllConfigurations() bool`

HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.

### GetParameterLibraries

`func (o *BTAssemblySimulation2246) GetParameterLibraries() []BTMParameter1`

GetParameterLibraries returns the ParameterLibraries field if non-nil, zero value otherwise.

### GetParameterLibrariesOk

`func (o *BTAssemblySimulation2246) GetParameterLibrariesOk() (*[]BTMParameter1, bool)`

GetParameterLibrariesOk returns a tuple with the ParameterLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLibraries

`func (o *BTAssemblySimulation2246) SetParameterLibraries(v []BTMParameter1)`

SetParameterLibraries sets ParameterLibraries field to given value.

### HasParameterLibraries

`func (o *BTAssemblySimulation2246) HasParameterLibraries() bool`

HasParameterLibraries returns a boolean if a field has been set.

### GetParametricInstanceFeature

`func (o *BTAssemblySimulation2246) GetParametricInstanceFeature() bool`

GetParametricInstanceFeature returns the ParametricInstanceFeature field if non-nil, zero value otherwise.

### GetParametricInstanceFeatureOk

`func (o *BTAssemblySimulation2246) GetParametricInstanceFeatureOk() (*bool, bool)`

GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstanceFeature

`func (o *BTAssemblySimulation2246) SetParametricInstanceFeature(v bool)`

SetParametricInstanceFeature sets ParametricInstanceFeature field to given value.

### HasParametricInstanceFeature

`func (o *BTAssemblySimulation2246) HasParametricInstanceFeature() bool`

HasParametricInstanceFeature returns a boolean if a field has been set.

### GetParentSuppressed

`func (o *BTAssemblySimulation2246) GetParentSuppressed() bool`

GetParentSuppressed returns the ParentSuppressed field if non-nil, zero value otherwise.

### GetParentSuppressedOk

`func (o *BTAssemblySimulation2246) GetParentSuppressedOk() (*bool, bool)`

GetParentSuppressedOk returns a tuple with the ParentSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSuppressed

`func (o *BTAssemblySimulation2246) SetParentSuppressed(v bool)`

SetParentSuppressed sets ParentSuppressed field to given value.

### HasParentSuppressed

`func (o *BTAssemblySimulation2246) HasParentSuppressed() bool`

HasParentSuppressed returns a boolean if a field has been set.

### GetReturnAfterSubfeatures

`func (o *BTAssemblySimulation2246) GetReturnAfterSubfeatures() bool`

GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field if non-nil, zero value otherwise.

### GetReturnAfterSubfeaturesOk

`func (o *BTAssemblySimulation2246) GetReturnAfterSubfeaturesOk() (*bool, bool)`

GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAfterSubfeatures

`func (o *BTAssemblySimulation2246) SetReturnAfterSubfeatures(v bool)`

SetReturnAfterSubfeatures sets ReturnAfterSubfeatures field to given value.

### HasReturnAfterSubfeatures

`func (o *BTAssemblySimulation2246) HasReturnAfterSubfeatures() bool`

HasReturnAfterSubfeatures returns a boolean if a field has been set.

### GetSubFeatures

`func (o *BTAssemblySimulation2246) GetSubFeatures() []BTMFeature134`

GetSubFeatures returns the SubFeatures field if non-nil, zero value otherwise.

### GetSubFeaturesOk

`func (o *BTAssemblySimulation2246) GetSubFeaturesOk() (*[]BTMFeature134, bool)`

GetSubFeaturesOk returns a tuple with the SubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeatures

`func (o *BTAssemblySimulation2246) SetSubFeatures(v []BTMFeature134)`

SetSubFeatures sets SubFeatures field to given value.

### HasSubFeatures

`func (o *BTAssemblySimulation2246) HasSubFeatures() bool`

HasSubFeatures returns a boolean if a field has been set.

### GetSubFeaturesNotUsedInQuery

`func (o *BTAssemblySimulation2246) GetSubFeaturesNotUsedInQuery() []BTMFeature134`

GetSubFeaturesNotUsedInQuery returns the SubFeaturesNotUsedInQuery field if non-nil, zero value otherwise.

### GetSubFeaturesNotUsedInQueryOk

`func (o *BTAssemblySimulation2246) GetSubFeaturesNotUsedInQueryOk() (*[]BTMFeature134, bool)`

GetSubFeaturesNotUsedInQueryOk returns a tuple with the SubFeaturesNotUsedInQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeaturesNotUsedInQuery

`func (o *BTAssemblySimulation2246) SetSubFeaturesNotUsedInQuery(v []BTMFeature134)`

SetSubFeaturesNotUsedInQuery sets SubFeaturesNotUsedInQuery field to given value.

### HasSubFeaturesNotUsedInQuery

`func (o *BTAssemblySimulation2246) HasSubFeaturesNotUsedInQuery() bool`

HasSubFeaturesNotUsedInQuery returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTAssemblySimulation2246) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTAssemblySimulation2246) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTAssemblySimulation2246) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTAssemblySimulation2246) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTAssemblySimulation2246) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTAssemblySimulation2246) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTAssemblySimulation2246) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTAssemblySimulation2246) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTAssemblySimulation2246) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTAssemblySimulation2246) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTAssemblySimulation2246) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTAssemblySimulation2246) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.

### GetType

`func (o *BTAssemblySimulation2246) GetType() GBTAssemblySimulationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTAssemblySimulation2246) GetTypeOk() (*GBTAssemblySimulationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTAssemblySimulation2246) SetType(v GBTAssemblySimulationType)`

SetType sets Type field to given value.

### HasType

`func (o *BTAssemblySimulation2246) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariableStudioReference

`func (o *BTAssemblySimulation2246) GetVariableStudioReference() bool`

GetVariableStudioReference returns the VariableStudioReference field if non-nil, zero value otherwise.

### GetVariableStudioReferenceOk

`func (o *BTAssemblySimulation2246) GetVariableStudioReferenceOk() (*bool, bool)`

GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudioReference

`func (o *BTAssemblySimulation2246) SetVariableStudioReference(v bool)`

SetVariableStudioReference sets VariableStudioReference field to given value.

### HasVariableStudioReference

`func (o *BTAssemblySimulation2246) HasVariableStudioReference() bool`

HasVariableStudioReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTAssemblySimulation2246) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTAssemblySimulation2246) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTAssemblySimulation2246) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTAssemblySimulation2246) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


