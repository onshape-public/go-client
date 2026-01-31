# BTMDerivedAssemblyMirrorFeature5094

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayParameterFromFeature** | Pointer to [**BTMParameterArray2025**](BTMParameterArray2025.md) |  | [optional] 
**AuxiliaryTreeFeature** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Feature** | Pointer to [**BTMAssemblyFeature887**](BTMAssemblyFeature887.md) |  | [optional] 
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
**ReferenceParameter** | Pointer to [**BTMParameterReferenceWithConfiguration3028**](BTMParameterReferenceWithConfiguration3028.md) |  | [optional] 
**ReturnAfterSubfeatures** | Pointer to **bool** | For internal use only. Should always be &#x60;false&#x60;. | [optional] 
**SubFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) | List of subfeatures belonging to the feature. | [optional] 
**SubFeaturesNotUsedInQuery** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**Suppressed** | Pointer to **bool** | If &#x60;true&#x60;, the feature is suppressed. It will skip regeneration, denoted by a line through the name in the Feature list. | [optional] 
**SuppressionConfigured** | Pointer to **bool** | &#x60;true&#x60; if the suppression is configured in the Part Studio. | [optional] 
**SuppressionState** | Pointer to [**BTMSuppressionState1924**](BTMSuppressionState1924.md) |  | [optional] 
**VariableStudioReference** | Pointer to **bool** | If &#x60;true&#x60;, the feature references a Variable Studio. | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTMDerivedAssemblyMirrorFeature5094

`func NewBTMDerivedAssemblyMirrorFeature5094() *BTMDerivedAssemblyMirrorFeature5094`

NewBTMDerivedAssemblyMirrorFeature5094 instantiates a new BTMDerivedAssemblyMirrorFeature5094 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMDerivedAssemblyMirrorFeature5094WithDefaults

`func NewBTMDerivedAssemblyMirrorFeature5094WithDefaults() *BTMDerivedAssemblyMirrorFeature5094`

NewBTMDerivedAssemblyMirrorFeature5094WithDefaults instantiates a new BTMDerivedAssemblyMirrorFeature5094 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrayParameterFromFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetArrayParameterFromFeature() BTMParameterArray2025`

GetArrayParameterFromFeature returns the ArrayParameterFromFeature field if non-nil, zero value otherwise.

### GetArrayParameterFromFeatureOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetArrayParameterFromFeatureOk() (*BTMParameterArray2025, bool)`

GetArrayParameterFromFeatureOk returns a tuple with the ArrayParameterFromFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayParameterFromFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetArrayParameterFromFeature(v BTMParameterArray2025)`

SetArrayParameterFromFeature sets ArrayParameterFromFeature field to given value.

### HasArrayParameterFromFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasArrayParameterFromFeature() bool`

HasArrayParameterFromFeature returns a boolean if a field has been set.

### GetAuxiliaryTreeFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetAuxiliaryTreeFeature() bool`

GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field if non-nil, zero value otherwise.

### GetAuxiliaryTreeFeatureOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetAuxiliaryTreeFeatureOk() (*bool, bool)`

GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTreeFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetAuxiliaryTreeFeature(v bool)`

SetAuxiliaryTreeFeature sets AuxiliaryTreeFeature field to given value.

### HasAuxiliaryTreeFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasAuxiliaryTreeFeature() bool`

HasAuxiliaryTreeFeature returns a boolean if a field has been set.

### GetBtType

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeature() BTMAssemblyFeature887`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureOk() (*BTMAssemblyFeature887, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetFeature(v BTMAssemblyFeature887)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureFolder

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureFolder() bool`

GetFeatureFolder returns the FeatureFolder field if non-nil, zero value otherwise.

### GetFeatureFolderOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureFolderOk() (*bool, bool)`

GetFeatureFolderOk returns a tuple with the FeatureFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFolder

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetFeatureFolder(v bool)`

SetFeatureFolder sets FeatureFolder field to given value.

### HasFeatureFolder

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasFeatureFolder() bool`

HasFeatureFolder returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureListFieldIndex

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureListFieldIndex() int32`

GetFeatureListFieldIndex returns the FeatureListFieldIndex field if non-nil, zero value otherwise.

### GetFeatureListFieldIndexOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureListFieldIndexOk() (*int32, bool)`

GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureListFieldIndex

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetFeatureListFieldIndex(v int32)`

SetFeatureListFieldIndex sets FeatureListFieldIndex field to given value.

### HasFeatureListFieldIndex

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasFeatureListFieldIndex() bool`

HasFeatureListFieldIndex returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFieldIndexForOwnedMateConnectors

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFieldIndexForOwnedMateConnectors() int32`

GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field if non-nil, zero value otherwise.

### GetFieldIndexForOwnedMateConnectorsOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool)`

GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIndexForOwnedMateConnectors

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetFieldIndexForOwnedMateConnectors(v int32)`

SetFieldIndexForOwnedMateConnectors sets FieldIndexForOwnedMateConnectors field to given value.

### HasFieldIndexForOwnedMateConnectors

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasFieldIndexForOwnedMateConnectors() bool`

HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetMateConnectorFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetMateConnectorFeature() bool`

GetMateConnectorFeature returns the MateConnectorFeature field if non-nil, zero value otherwise.

### GetMateConnectorFeatureOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetMateConnectorFeatureOk() (*bool, bool)`

GetMateConnectorFeatureOk returns a tuple with the MateConnectorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectorFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetMateConnectorFeature(v bool)`

SetMateConnectorFeature sets MateConnectorFeature field to given value.

### HasMateConnectorFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasMateConnectorFeature() bool`

HasMateConnectorFeature returns a boolean if a field has been set.

### GetMateConnectors

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetMateConnectors() []BTMMateConnector66`

GetMateConnectors returns the MateConnectors field if non-nil, zero value otherwise.

### GetMateConnectorsOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetMateConnectorsOk() (*[]BTMMateConnector66, bool)`

GetMateConnectorsOk returns a tuple with the MateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectors

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetMateConnectors(v []BTMMateConnector66)`

SetMateConnectors sets MateConnectors field to given value.

### HasMateConnectors

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasMateConnectors() bool`

HasMateConnectors returns a boolean if a field has been set.

### GetName

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeId

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOccurrenceQueriesFromAllConfigurations

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904`

GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field if non-nil, zero value otherwise.

### GetOccurrenceQueriesFromAllConfigurationsOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetOccurrenceQueriesFromAllConfigurationsOk() (*[]BTMIndividualQueryWithOccurrenceBase904, bool)`

GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceQueriesFromAllConfigurations

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904)`

SetOccurrenceQueriesFromAllConfigurations sets OccurrenceQueriesFromAllConfigurations field to given value.

### HasOccurrenceQueriesFromAllConfigurations

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasOccurrenceQueriesFromAllConfigurations() bool`

HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.

### GetParameterLibraries

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetParameterLibraries() []BTMParameter1`

GetParameterLibraries returns the ParameterLibraries field if non-nil, zero value otherwise.

### GetParameterLibrariesOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetParameterLibrariesOk() (*[]BTMParameter1, bool)`

GetParameterLibrariesOk returns a tuple with the ParameterLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLibraries

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetParameterLibraries(v []BTMParameter1)`

SetParameterLibraries sets ParameterLibraries field to given value.

### HasParameterLibraries

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasParameterLibraries() bool`

HasParameterLibraries returns a boolean if a field has been set.

### GetParametricInstanceFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetParametricInstanceFeature() bool`

GetParametricInstanceFeature returns the ParametricInstanceFeature field if non-nil, zero value otherwise.

### GetParametricInstanceFeatureOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetParametricInstanceFeatureOk() (*bool, bool)`

GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstanceFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetParametricInstanceFeature(v bool)`

SetParametricInstanceFeature sets ParametricInstanceFeature field to given value.

### HasParametricInstanceFeature

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasParametricInstanceFeature() bool`

HasParametricInstanceFeature returns a boolean if a field has been set.

### GetParentSuppressed

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetParentSuppressed() bool`

GetParentSuppressed returns the ParentSuppressed field if non-nil, zero value otherwise.

### GetParentSuppressedOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetParentSuppressedOk() (*bool, bool)`

GetParentSuppressedOk returns a tuple with the ParentSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSuppressed

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetParentSuppressed(v bool)`

SetParentSuppressed sets ParentSuppressed field to given value.

### HasParentSuppressed

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasParentSuppressed() bool`

HasParentSuppressed returns a boolean if a field has been set.

### GetReferenceParameter

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetReferenceParameter() BTMParameterReferenceWithConfiguration3028`

GetReferenceParameter returns the ReferenceParameter field if non-nil, zero value otherwise.

### GetReferenceParameterOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetReferenceParameterOk() (*BTMParameterReferenceWithConfiguration3028, bool)`

GetReferenceParameterOk returns a tuple with the ReferenceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceParameter

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetReferenceParameter(v BTMParameterReferenceWithConfiguration3028)`

SetReferenceParameter sets ReferenceParameter field to given value.

### HasReferenceParameter

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasReferenceParameter() bool`

HasReferenceParameter returns a boolean if a field has been set.

### GetReturnAfterSubfeatures

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetReturnAfterSubfeatures() bool`

GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field if non-nil, zero value otherwise.

### GetReturnAfterSubfeaturesOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetReturnAfterSubfeaturesOk() (*bool, bool)`

GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAfterSubfeatures

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetReturnAfterSubfeatures(v bool)`

SetReturnAfterSubfeatures sets ReturnAfterSubfeatures field to given value.

### HasReturnAfterSubfeatures

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasReturnAfterSubfeatures() bool`

HasReturnAfterSubfeatures returns a boolean if a field has been set.

### GetSubFeatures

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSubFeatures() []BTMFeature134`

GetSubFeatures returns the SubFeatures field if non-nil, zero value otherwise.

### GetSubFeaturesOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSubFeaturesOk() (*[]BTMFeature134, bool)`

GetSubFeaturesOk returns a tuple with the SubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeatures

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetSubFeatures(v []BTMFeature134)`

SetSubFeatures sets SubFeatures field to given value.

### HasSubFeatures

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasSubFeatures() bool`

HasSubFeatures returns a boolean if a field has been set.

### GetSubFeaturesNotUsedInQuery

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSubFeaturesNotUsedInQuery() []BTMFeature134`

GetSubFeaturesNotUsedInQuery returns the SubFeaturesNotUsedInQuery field if non-nil, zero value otherwise.

### GetSubFeaturesNotUsedInQueryOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSubFeaturesNotUsedInQueryOk() (*[]BTMFeature134, bool)`

GetSubFeaturesNotUsedInQueryOk returns a tuple with the SubFeaturesNotUsedInQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeaturesNotUsedInQuery

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetSubFeaturesNotUsedInQuery(v []BTMFeature134)`

SetSubFeaturesNotUsedInQuery sets SubFeaturesNotUsedInQuery field to given value.

### HasSubFeaturesNotUsedInQuery

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasSubFeaturesNotUsedInQuery() bool`

HasSubFeaturesNotUsedInQuery returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.

### GetVariableStudioReference

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetVariableStudioReference() bool`

GetVariableStudioReference returns the VariableStudioReference field if non-nil, zero value otherwise.

### GetVariableStudioReferenceOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetVariableStudioReferenceOk() (*bool, bool)`

GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudioReference

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetVariableStudioReference(v bool)`

SetVariableStudioReference sets VariableStudioReference field to given value.

### HasVariableStudioReference

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasVariableStudioReference() bool`

HasVariableStudioReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTMDerivedAssemblyMirrorFeature5094) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTMDerivedAssemblyMirrorFeature5094) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTMDerivedAssemblyMirrorFeature5094) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


