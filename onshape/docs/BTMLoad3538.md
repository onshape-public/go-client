# BTMLoad3538

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxiliaryTreeFeature** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**DefinedByComponents** | Pointer to **bool** |  | [optional] 
**DirectionFlipped** | Pointer to **bool** |  | [optional] 
**FeatureFolder** | Pointer to **bool** |  | [optional] 
**FeatureId** | Pointer to **string** | Unique ID of the feature instance within this Part Studio. | [optional] 
**FeatureListFieldIndex** | Pointer to **int32** |  | [optional] 
**FeatureType** | Pointer to **string** | The name of the feature spec that this feature instantiates. | [optional] 
**FgsBaseUnits** | Pointer to **string** |  | [optional] 
**FieldIndexForOwnedMateConnectors** | Pointer to **int32** |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Element microversion that is being imported. | [optional] 
**LoadComponentParameterIds** | Pointer to **map[string]string** |  | [optional] 
**LoadRegionParameterId** | Pointer to **string** |  | [optional] 
**LoadType** | Pointer to [**GBTLoadType**](GBTLoadType.md) |  | [optional] 
**MagnitudeParameterId** | Pointer to **string** |  | [optional] 
**MagnitudeQuantityType** | Pointer to [**GBTQuantityType**](GBTQuantityType.md) |  | [optional] 
**MateConnectorFeature** | Pointer to **bool** |  | [optional] 
**MateConnectors** | Pointer to [**[]BTMMateConnector66**](BTMMateConnector66.md) |  | [optional] 
**Name** | Pointer to **string** | User-visible name of the feature. | [optional] 
**Namespace** | Pointer to **string** | Indicates where the feature definition lives. Features in the FeatureScript standard library have a namespace value of &#x60;\&quot;\&quot;&#x60;. Custom features identify the Feature Studio that contains the definition. | [optional] 
**NodeId** | Pointer to **string** | ID for the feature node. | [optional] 
**OccurrenceQueriesFromAllConfigurations** | Pointer to [**[]BTMIndividualQueryWithOccurrenceBase904**](BTMIndividualQueryWithOccurrenceBase904.md) |  | [optional] 
**ParameterLibraries** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**ParametricInstanceFeature** | Pointer to **bool** |  | [optional] 
**ReturnAfterSubfeatures** | Pointer to **bool** | For internal use only. Should always be &#x60;false&#x60;. | [optional] 
**StructuralLoad** | Pointer to **bool** |  | [optional] 
**SubFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) | List of subfeatures belonging to the feature. | [optional] 
**SubFeaturesNotUsedInQuery** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**Suppressed** | Pointer to **bool** | If &#x60;true&#x60;, the feature is suppressed. It will skip regeneration, denoted by a line through the name in the Feature list. | [optional] 
**SuppressedInSimulations** | Pointer to **map[string]int32** |  | [optional] 
**SuppressionConfigured** | Pointer to **bool** | &#x60;true&#x60; if the suppression is configured in the Part Studio. | [optional] 
**SuppressionState** | Pointer to [**BTMSuppressionState1924**](BTMSuppressionState1924.md) |  | [optional] 
**VariableStudioReference** | Pointer to **bool** | If &#x60;true&#x60;, the feature references a Variable Studio. | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTMLoad3538

`func NewBTMLoad3538() *BTMLoad3538`

NewBTMLoad3538 instantiates a new BTMLoad3538 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMLoad3538WithDefaults

`func NewBTMLoad3538WithDefaults() *BTMLoad3538`

NewBTMLoad3538WithDefaults instantiates a new BTMLoad3538 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxiliaryTreeFeature

`func (o *BTMLoad3538) GetAuxiliaryTreeFeature() bool`

GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field if non-nil, zero value otherwise.

### GetAuxiliaryTreeFeatureOk

`func (o *BTMLoad3538) GetAuxiliaryTreeFeatureOk() (*bool, bool)`

GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTreeFeature

`func (o *BTMLoad3538) SetAuxiliaryTreeFeature(v bool)`

SetAuxiliaryTreeFeature sets AuxiliaryTreeFeature field to given value.

### HasAuxiliaryTreeFeature

`func (o *BTMLoad3538) HasAuxiliaryTreeFeature() bool`

HasAuxiliaryTreeFeature returns a boolean if a field has been set.

### GetBtType

`func (o *BTMLoad3538) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMLoad3538) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMLoad3538) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMLoad3538) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDefinedByComponents

`func (o *BTMLoad3538) GetDefinedByComponents() bool`

GetDefinedByComponents returns the DefinedByComponents field if non-nil, zero value otherwise.

### GetDefinedByComponentsOk

`func (o *BTMLoad3538) GetDefinedByComponentsOk() (*bool, bool)`

GetDefinedByComponentsOk returns a tuple with the DefinedByComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedByComponents

`func (o *BTMLoad3538) SetDefinedByComponents(v bool)`

SetDefinedByComponents sets DefinedByComponents field to given value.

### HasDefinedByComponents

`func (o *BTMLoad3538) HasDefinedByComponents() bool`

HasDefinedByComponents returns a boolean if a field has been set.

### GetDirectionFlipped

`func (o *BTMLoad3538) GetDirectionFlipped() bool`

GetDirectionFlipped returns the DirectionFlipped field if non-nil, zero value otherwise.

### GetDirectionFlippedOk

`func (o *BTMLoad3538) GetDirectionFlippedOk() (*bool, bool)`

GetDirectionFlippedOk returns a tuple with the DirectionFlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionFlipped

`func (o *BTMLoad3538) SetDirectionFlipped(v bool)`

SetDirectionFlipped sets DirectionFlipped field to given value.

### HasDirectionFlipped

`func (o *BTMLoad3538) HasDirectionFlipped() bool`

HasDirectionFlipped returns a boolean if a field has been set.

### GetFeatureFolder

`func (o *BTMLoad3538) GetFeatureFolder() bool`

GetFeatureFolder returns the FeatureFolder field if non-nil, zero value otherwise.

### GetFeatureFolderOk

`func (o *BTMLoad3538) GetFeatureFolderOk() (*bool, bool)`

GetFeatureFolderOk returns a tuple with the FeatureFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFolder

`func (o *BTMLoad3538) SetFeatureFolder(v bool)`

SetFeatureFolder sets FeatureFolder field to given value.

### HasFeatureFolder

`func (o *BTMLoad3538) HasFeatureFolder() bool`

HasFeatureFolder returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTMLoad3538) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTMLoad3538) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTMLoad3538) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTMLoad3538) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureListFieldIndex

`func (o *BTMLoad3538) GetFeatureListFieldIndex() int32`

GetFeatureListFieldIndex returns the FeatureListFieldIndex field if non-nil, zero value otherwise.

### GetFeatureListFieldIndexOk

`func (o *BTMLoad3538) GetFeatureListFieldIndexOk() (*int32, bool)`

GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureListFieldIndex

`func (o *BTMLoad3538) SetFeatureListFieldIndex(v int32)`

SetFeatureListFieldIndex sets FeatureListFieldIndex field to given value.

### HasFeatureListFieldIndex

`func (o *BTMLoad3538) HasFeatureListFieldIndex() bool`

HasFeatureListFieldIndex returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTMLoad3538) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTMLoad3538) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTMLoad3538) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTMLoad3538) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFgsBaseUnits

`func (o *BTMLoad3538) GetFgsBaseUnits() string`

GetFgsBaseUnits returns the FgsBaseUnits field if non-nil, zero value otherwise.

### GetFgsBaseUnitsOk

`func (o *BTMLoad3538) GetFgsBaseUnitsOk() (*string, bool)`

GetFgsBaseUnitsOk returns a tuple with the FgsBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgsBaseUnits

`func (o *BTMLoad3538) SetFgsBaseUnits(v string)`

SetFgsBaseUnits sets FgsBaseUnits field to given value.

### HasFgsBaseUnits

`func (o *BTMLoad3538) HasFgsBaseUnits() bool`

HasFgsBaseUnits returns a boolean if a field has been set.

### GetFieldIndexForOwnedMateConnectors

`func (o *BTMLoad3538) GetFieldIndexForOwnedMateConnectors() int32`

GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field if non-nil, zero value otherwise.

### GetFieldIndexForOwnedMateConnectorsOk

`func (o *BTMLoad3538) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool)`

GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIndexForOwnedMateConnectors

`func (o *BTMLoad3538) SetFieldIndexForOwnedMateConnectors(v int32)`

SetFieldIndexForOwnedMateConnectors sets FieldIndexForOwnedMateConnectors field to given value.

### HasFieldIndexForOwnedMateConnectors

`func (o *BTMLoad3538) HasFieldIndexForOwnedMateConnectors() bool`

HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTMLoad3538) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTMLoad3538) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTMLoad3538) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTMLoad3538) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetLoadComponentParameterIds

`func (o *BTMLoad3538) GetLoadComponentParameterIds() map[string]string`

GetLoadComponentParameterIds returns the LoadComponentParameterIds field if non-nil, zero value otherwise.

### GetLoadComponentParameterIdsOk

`func (o *BTMLoad3538) GetLoadComponentParameterIdsOk() (*map[string]string, bool)`

GetLoadComponentParameterIdsOk returns a tuple with the LoadComponentParameterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadComponentParameterIds

`func (o *BTMLoad3538) SetLoadComponentParameterIds(v map[string]string)`

SetLoadComponentParameterIds sets LoadComponentParameterIds field to given value.

### HasLoadComponentParameterIds

`func (o *BTMLoad3538) HasLoadComponentParameterIds() bool`

HasLoadComponentParameterIds returns a boolean if a field has been set.

### GetLoadRegionParameterId

`func (o *BTMLoad3538) GetLoadRegionParameterId() string`

GetLoadRegionParameterId returns the LoadRegionParameterId field if non-nil, zero value otherwise.

### GetLoadRegionParameterIdOk

`func (o *BTMLoad3538) GetLoadRegionParameterIdOk() (*string, bool)`

GetLoadRegionParameterIdOk returns a tuple with the LoadRegionParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadRegionParameterId

`func (o *BTMLoad3538) SetLoadRegionParameterId(v string)`

SetLoadRegionParameterId sets LoadRegionParameterId field to given value.

### HasLoadRegionParameterId

`func (o *BTMLoad3538) HasLoadRegionParameterId() bool`

HasLoadRegionParameterId returns a boolean if a field has been set.

### GetLoadType

`func (o *BTMLoad3538) GetLoadType() GBTLoadType`

GetLoadType returns the LoadType field if non-nil, zero value otherwise.

### GetLoadTypeOk

`func (o *BTMLoad3538) GetLoadTypeOk() (*GBTLoadType, bool)`

GetLoadTypeOk returns a tuple with the LoadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadType

`func (o *BTMLoad3538) SetLoadType(v GBTLoadType)`

SetLoadType sets LoadType field to given value.

### HasLoadType

`func (o *BTMLoad3538) HasLoadType() bool`

HasLoadType returns a boolean if a field has been set.

### GetMagnitudeParameterId

`func (o *BTMLoad3538) GetMagnitudeParameterId() string`

GetMagnitudeParameterId returns the MagnitudeParameterId field if non-nil, zero value otherwise.

### GetMagnitudeParameterIdOk

`func (o *BTMLoad3538) GetMagnitudeParameterIdOk() (*string, bool)`

GetMagnitudeParameterIdOk returns a tuple with the MagnitudeParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeParameterId

`func (o *BTMLoad3538) SetMagnitudeParameterId(v string)`

SetMagnitudeParameterId sets MagnitudeParameterId field to given value.

### HasMagnitudeParameterId

`func (o *BTMLoad3538) HasMagnitudeParameterId() bool`

HasMagnitudeParameterId returns a boolean if a field has been set.

### GetMagnitudeQuantityType

`func (o *BTMLoad3538) GetMagnitudeQuantityType() GBTQuantityType`

GetMagnitudeQuantityType returns the MagnitudeQuantityType field if non-nil, zero value otherwise.

### GetMagnitudeQuantityTypeOk

`func (o *BTMLoad3538) GetMagnitudeQuantityTypeOk() (*GBTQuantityType, bool)`

GetMagnitudeQuantityTypeOk returns a tuple with the MagnitudeQuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeQuantityType

`func (o *BTMLoad3538) SetMagnitudeQuantityType(v GBTQuantityType)`

SetMagnitudeQuantityType sets MagnitudeQuantityType field to given value.

### HasMagnitudeQuantityType

`func (o *BTMLoad3538) HasMagnitudeQuantityType() bool`

HasMagnitudeQuantityType returns a boolean if a field has been set.

### GetMateConnectorFeature

`func (o *BTMLoad3538) GetMateConnectorFeature() bool`

GetMateConnectorFeature returns the MateConnectorFeature field if non-nil, zero value otherwise.

### GetMateConnectorFeatureOk

`func (o *BTMLoad3538) GetMateConnectorFeatureOk() (*bool, bool)`

GetMateConnectorFeatureOk returns a tuple with the MateConnectorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectorFeature

`func (o *BTMLoad3538) SetMateConnectorFeature(v bool)`

SetMateConnectorFeature sets MateConnectorFeature field to given value.

### HasMateConnectorFeature

`func (o *BTMLoad3538) HasMateConnectorFeature() bool`

HasMateConnectorFeature returns a boolean if a field has been set.

### GetMateConnectors

`func (o *BTMLoad3538) GetMateConnectors() []BTMMateConnector66`

GetMateConnectors returns the MateConnectors field if non-nil, zero value otherwise.

### GetMateConnectorsOk

`func (o *BTMLoad3538) GetMateConnectorsOk() (*[]BTMMateConnector66, bool)`

GetMateConnectorsOk returns a tuple with the MateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectors

`func (o *BTMLoad3538) SetMateConnectors(v []BTMMateConnector66)`

SetMateConnectors sets MateConnectors field to given value.

### HasMateConnectors

`func (o *BTMLoad3538) HasMateConnectors() bool`

HasMateConnectors returns a boolean if a field has been set.

### GetName

`func (o *BTMLoad3538) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMLoad3538) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMLoad3538) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMLoad3538) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTMLoad3538) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTMLoad3538) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTMLoad3538) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTMLoad3538) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeId

`func (o *BTMLoad3538) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTMLoad3538) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTMLoad3538) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTMLoad3538) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOccurrenceQueriesFromAllConfigurations

`func (o *BTMLoad3538) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904`

GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field if non-nil, zero value otherwise.

### GetOccurrenceQueriesFromAllConfigurationsOk

`func (o *BTMLoad3538) GetOccurrenceQueriesFromAllConfigurationsOk() (*[]BTMIndividualQueryWithOccurrenceBase904, bool)`

GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceQueriesFromAllConfigurations

`func (o *BTMLoad3538) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904)`

SetOccurrenceQueriesFromAllConfigurations sets OccurrenceQueriesFromAllConfigurations field to given value.

### HasOccurrenceQueriesFromAllConfigurations

`func (o *BTMLoad3538) HasOccurrenceQueriesFromAllConfigurations() bool`

HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.

### GetParameterLibraries

`func (o *BTMLoad3538) GetParameterLibraries() []BTMParameter1`

GetParameterLibraries returns the ParameterLibraries field if non-nil, zero value otherwise.

### GetParameterLibrariesOk

`func (o *BTMLoad3538) GetParameterLibrariesOk() (*[]BTMParameter1, bool)`

GetParameterLibrariesOk returns a tuple with the ParameterLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLibraries

`func (o *BTMLoad3538) SetParameterLibraries(v []BTMParameter1)`

SetParameterLibraries sets ParameterLibraries field to given value.

### HasParameterLibraries

`func (o *BTMLoad3538) HasParameterLibraries() bool`

HasParameterLibraries returns a boolean if a field has been set.

### GetParametricInstanceFeature

`func (o *BTMLoad3538) GetParametricInstanceFeature() bool`

GetParametricInstanceFeature returns the ParametricInstanceFeature field if non-nil, zero value otherwise.

### GetParametricInstanceFeatureOk

`func (o *BTMLoad3538) GetParametricInstanceFeatureOk() (*bool, bool)`

GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstanceFeature

`func (o *BTMLoad3538) SetParametricInstanceFeature(v bool)`

SetParametricInstanceFeature sets ParametricInstanceFeature field to given value.

### HasParametricInstanceFeature

`func (o *BTMLoad3538) HasParametricInstanceFeature() bool`

HasParametricInstanceFeature returns a boolean if a field has been set.

### GetReturnAfterSubfeatures

`func (o *BTMLoad3538) GetReturnAfterSubfeatures() bool`

GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field if non-nil, zero value otherwise.

### GetReturnAfterSubfeaturesOk

`func (o *BTMLoad3538) GetReturnAfterSubfeaturesOk() (*bool, bool)`

GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAfterSubfeatures

`func (o *BTMLoad3538) SetReturnAfterSubfeatures(v bool)`

SetReturnAfterSubfeatures sets ReturnAfterSubfeatures field to given value.

### HasReturnAfterSubfeatures

`func (o *BTMLoad3538) HasReturnAfterSubfeatures() bool`

HasReturnAfterSubfeatures returns a boolean if a field has been set.

### GetStructuralLoad

`func (o *BTMLoad3538) GetStructuralLoad() bool`

GetStructuralLoad returns the StructuralLoad field if non-nil, zero value otherwise.

### GetStructuralLoadOk

`func (o *BTMLoad3538) GetStructuralLoadOk() (*bool, bool)`

GetStructuralLoadOk returns a tuple with the StructuralLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLoad

`func (o *BTMLoad3538) SetStructuralLoad(v bool)`

SetStructuralLoad sets StructuralLoad field to given value.

### HasStructuralLoad

`func (o *BTMLoad3538) HasStructuralLoad() bool`

HasStructuralLoad returns a boolean if a field has been set.

### GetSubFeatures

`func (o *BTMLoad3538) GetSubFeatures() []BTMFeature134`

GetSubFeatures returns the SubFeatures field if non-nil, zero value otherwise.

### GetSubFeaturesOk

`func (o *BTMLoad3538) GetSubFeaturesOk() (*[]BTMFeature134, bool)`

GetSubFeaturesOk returns a tuple with the SubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeatures

`func (o *BTMLoad3538) SetSubFeatures(v []BTMFeature134)`

SetSubFeatures sets SubFeatures field to given value.

### HasSubFeatures

`func (o *BTMLoad3538) HasSubFeatures() bool`

HasSubFeatures returns a boolean if a field has been set.

### GetSubFeaturesNotUsedInQuery

`func (o *BTMLoad3538) GetSubFeaturesNotUsedInQuery() []BTMFeature134`

GetSubFeaturesNotUsedInQuery returns the SubFeaturesNotUsedInQuery field if non-nil, zero value otherwise.

### GetSubFeaturesNotUsedInQueryOk

`func (o *BTMLoad3538) GetSubFeaturesNotUsedInQueryOk() (*[]BTMFeature134, bool)`

GetSubFeaturesNotUsedInQueryOk returns a tuple with the SubFeaturesNotUsedInQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeaturesNotUsedInQuery

`func (o *BTMLoad3538) SetSubFeaturesNotUsedInQuery(v []BTMFeature134)`

SetSubFeaturesNotUsedInQuery sets SubFeaturesNotUsedInQuery field to given value.

### HasSubFeaturesNotUsedInQuery

`func (o *BTMLoad3538) HasSubFeaturesNotUsedInQuery() bool`

HasSubFeaturesNotUsedInQuery returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTMLoad3538) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTMLoad3538) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTMLoad3538) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTMLoad3538) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressedInSimulations

`func (o *BTMLoad3538) GetSuppressedInSimulations() map[string]int32`

GetSuppressedInSimulations returns the SuppressedInSimulations field if non-nil, zero value otherwise.

### GetSuppressedInSimulationsOk

`func (o *BTMLoad3538) GetSuppressedInSimulationsOk() (*map[string]int32, bool)`

GetSuppressedInSimulationsOk returns a tuple with the SuppressedInSimulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedInSimulations

`func (o *BTMLoad3538) SetSuppressedInSimulations(v map[string]int32)`

SetSuppressedInSimulations sets SuppressedInSimulations field to given value.

### HasSuppressedInSimulations

`func (o *BTMLoad3538) HasSuppressedInSimulations() bool`

HasSuppressedInSimulations returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTMLoad3538) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTMLoad3538) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTMLoad3538) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTMLoad3538) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTMLoad3538) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTMLoad3538) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTMLoad3538) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTMLoad3538) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.

### GetVariableStudioReference

`func (o *BTMLoad3538) GetVariableStudioReference() bool`

GetVariableStudioReference returns the VariableStudioReference field if non-nil, zero value otherwise.

### GetVariableStudioReferenceOk

`func (o *BTMLoad3538) GetVariableStudioReferenceOk() (*bool, bool)`

GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudioReference

`func (o *BTMLoad3538) SetVariableStudioReference(v bool)`

SetVariableStudioReference sets VariableStudioReference field to given value.

### HasVariableStudioReference

`func (o *BTMLoad3538) HasVariableStudioReference() bool`

HasVariableStudioReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTMLoad3538) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTMLoad3538) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTMLoad3538) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTMLoad3538) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


