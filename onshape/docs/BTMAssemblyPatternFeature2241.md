# BTMAssemblyPatternFeature2241

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
**Name** | Pointer to **string** | User-visible name of the feature. | [optional] 
**Namespace** | Pointer to **string** | Indicates where the feature definition lives. Features in the FeatureScript standard library have a namespace value of &#x60;\&quot;\&quot;&#x60;. Custom features identify the Feature Studio that contains the definition. | [optional] 
**NodeId** | Pointer to **string** | ID for the feature node. | [optional] 
**OccurrenceQueriesFromAllConfigurations** | Pointer to [**[]BTMIndividualQueryWithOccurrenceBase904**](BTMIndividualQueryWithOccurrenceBase904.md) |  | [optional] 
**ParametricInstanceFeature** | Pointer to **bool** |  | [optional] 
**PatternType** | Pointer to [**GBTPatternType**](GBTPatternType.md) |  | [optional] 
**ReturnAfterSubfeatures** | Pointer to **bool** | For internal use only. Should always be &#x60;false&#x60;. | [optional] 
**SubFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) | List of subfeatures belonging to the feature. | [optional] 
**Suppressed** | Pointer to **bool** | If &#x60;true&#x60;, the feature is suppressed. It will skip regeneration, denoted by a line through the name in the Feature list. | [optional] 
**SuppressionConfigured** | Pointer to **bool** | &#x60;true&#x60; if the suppression is configured in the Part Studio. | [optional] 
**VariableStudioReference** | Pointer to **bool** | If &#x60;true&#x60;, the feature references a Variable Studio. | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTMAssemblyPatternFeature2241

`func NewBTMAssemblyPatternFeature2241() *BTMAssemblyPatternFeature2241`

NewBTMAssemblyPatternFeature2241 instantiates a new BTMAssemblyPatternFeature2241 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMAssemblyPatternFeature2241WithDefaults

`func NewBTMAssemblyPatternFeature2241WithDefaults() *BTMAssemblyPatternFeature2241`

NewBTMAssemblyPatternFeature2241WithDefaults instantiates a new BTMAssemblyPatternFeature2241 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxiliaryTreeFeature

`func (o *BTMAssemblyPatternFeature2241) GetAuxiliaryTreeFeature() bool`

GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field if non-nil, zero value otherwise.

### GetAuxiliaryTreeFeatureOk

`func (o *BTMAssemblyPatternFeature2241) GetAuxiliaryTreeFeatureOk() (*bool, bool)`

GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTreeFeature

`func (o *BTMAssemblyPatternFeature2241) SetAuxiliaryTreeFeature(v bool)`

SetAuxiliaryTreeFeature sets AuxiliaryTreeFeature field to given value.

### HasAuxiliaryTreeFeature

`func (o *BTMAssemblyPatternFeature2241) HasAuxiliaryTreeFeature() bool`

HasAuxiliaryTreeFeature returns a boolean if a field has been set.

### GetBtType

`func (o *BTMAssemblyPatternFeature2241) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMAssemblyPatternFeature2241) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMAssemblyPatternFeature2241) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMAssemblyPatternFeature2241) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatureFolder

`func (o *BTMAssemblyPatternFeature2241) GetFeatureFolder() bool`

GetFeatureFolder returns the FeatureFolder field if non-nil, zero value otherwise.

### GetFeatureFolderOk

`func (o *BTMAssemblyPatternFeature2241) GetFeatureFolderOk() (*bool, bool)`

GetFeatureFolderOk returns a tuple with the FeatureFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFolder

`func (o *BTMAssemblyPatternFeature2241) SetFeatureFolder(v bool)`

SetFeatureFolder sets FeatureFolder field to given value.

### HasFeatureFolder

`func (o *BTMAssemblyPatternFeature2241) HasFeatureFolder() bool`

HasFeatureFolder returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTMAssemblyPatternFeature2241) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTMAssemblyPatternFeature2241) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTMAssemblyPatternFeature2241) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTMAssemblyPatternFeature2241) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureListFieldIndex

`func (o *BTMAssemblyPatternFeature2241) GetFeatureListFieldIndex() int32`

GetFeatureListFieldIndex returns the FeatureListFieldIndex field if non-nil, zero value otherwise.

### GetFeatureListFieldIndexOk

`func (o *BTMAssemblyPatternFeature2241) GetFeatureListFieldIndexOk() (*int32, bool)`

GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureListFieldIndex

`func (o *BTMAssemblyPatternFeature2241) SetFeatureListFieldIndex(v int32)`

SetFeatureListFieldIndex sets FeatureListFieldIndex field to given value.

### HasFeatureListFieldIndex

`func (o *BTMAssemblyPatternFeature2241) HasFeatureListFieldIndex() bool`

HasFeatureListFieldIndex returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTMAssemblyPatternFeature2241) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTMAssemblyPatternFeature2241) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTMAssemblyPatternFeature2241) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTMAssemblyPatternFeature2241) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFieldIndexForOwnedMateConnectors

`func (o *BTMAssemblyPatternFeature2241) GetFieldIndexForOwnedMateConnectors() int32`

GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field if non-nil, zero value otherwise.

### GetFieldIndexForOwnedMateConnectorsOk

`func (o *BTMAssemblyPatternFeature2241) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool)`

GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIndexForOwnedMateConnectors

`func (o *BTMAssemblyPatternFeature2241) SetFieldIndexForOwnedMateConnectors(v int32)`

SetFieldIndexForOwnedMateConnectors sets FieldIndexForOwnedMateConnectors field to given value.

### HasFieldIndexForOwnedMateConnectors

`func (o *BTMAssemblyPatternFeature2241) HasFieldIndexForOwnedMateConnectors() bool`

HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTMAssemblyPatternFeature2241) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTMAssemblyPatternFeature2241) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTMAssemblyPatternFeature2241) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTMAssemblyPatternFeature2241) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetName

`func (o *BTMAssemblyPatternFeature2241) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMAssemblyPatternFeature2241) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMAssemblyPatternFeature2241) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMAssemblyPatternFeature2241) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTMAssemblyPatternFeature2241) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTMAssemblyPatternFeature2241) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTMAssemblyPatternFeature2241) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTMAssemblyPatternFeature2241) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeId

`func (o *BTMAssemblyPatternFeature2241) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTMAssemblyPatternFeature2241) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTMAssemblyPatternFeature2241) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTMAssemblyPatternFeature2241) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOccurrenceQueriesFromAllConfigurations

`func (o *BTMAssemblyPatternFeature2241) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904`

GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field if non-nil, zero value otherwise.

### GetOccurrenceQueriesFromAllConfigurationsOk

`func (o *BTMAssemblyPatternFeature2241) GetOccurrenceQueriesFromAllConfigurationsOk() (*[]BTMIndividualQueryWithOccurrenceBase904, bool)`

GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceQueriesFromAllConfigurations

`func (o *BTMAssemblyPatternFeature2241) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904)`

SetOccurrenceQueriesFromAllConfigurations sets OccurrenceQueriesFromAllConfigurations field to given value.

### HasOccurrenceQueriesFromAllConfigurations

`func (o *BTMAssemblyPatternFeature2241) HasOccurrenceQueriesFromAllConfigurations() bool`

HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.

### GetParametricInstanceFeature

`func (o *BTMAssemblyPatternFeature2241) GetParametricInstanceFeature() bool`

GetParametricInstanceFeature returns the ParametricInstanceFeature field if non-nil, zero value otherwise.

### GetParametricInstanceFeatureOk

`func (o *BTMAssemblyPatternFeature2241) GetParametricInstanceFeatureOk() (*bool, bool)`

GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstanceFeature

`func (o *BTMAssemblyPatternFeature2241) SetParametricInstanceFeature(v bool)`

SetParametricInstanceFeature sets ParametricInstanceFeature field to given value.

### HasParametricInstanceFeature

`func (o *BTMAssemblyPatternFeature2241) HasParametricInstanceFeature() bool`

HasParametricInstanceFeature returns a boolean if a field has been set.

### GetPatternType

`func (o *BTMAssemblyPatternFeature2241) GetPatternType() GBTPatternType`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *BTMAssemblyPatternFeature2241) GetPatternTypeOk() (*GBTPatternType, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *BTMAssemblyPatternFeature2241) SetPatternType(v GBTPatternType)`

SetPatternType sets PatternType field to given value.

### HasPatternType

`func (o *BTMAssemblyPatternFeature2241) HasPatternType() bool`

HasPatternType returns a boolean if a field has been set.

### GetReturnAfterSubfeatures

`func (o *BTMAssemblyPatternFeature2241) GetReturnAfterSubfeatures() bool`

GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field if non-nil, zero value otherwise.

### GetReturnAfterSubfeaturesOk

`func (o *BTMAssemblyPatternFeature2241) GetReturnAfterSubfeaturesOk() (*bool, bool)`

GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAfterSubfeatures

`func (o *BTMAssemblyPatternFeature2241) SetReturnAfterSubfeatures(v bool)`

SetReturnAfterSubfeatures sets ReturnAfterSubfeatures field to given value.

### HasReturnAfterSubfeatures

`func (o *BTMAssemblyPatternFeature2241) HasReturnAfterSubfeatures() bool`

HasReturnAfterSubfeatures returns a boolean if a field has been set.

### GetSubFeatures

`func (o *BTMAssemblyPatternFeature2241) GetSubFeatures() []BTMFeature134`

GetSubFeatures returns the SubFeatures field if non-nil, zero value otherwise.

### GetSubFeaturesOk

`func (o *BTMAssemblyPatternFeature2241) GetSubFeaturesOk() (*[]BTMFeature134, bool)`

GetSubFeaturesOk returns a tuple with the SubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeatures

`func (o *BTMAssemblyPatternFeature2241) SetSubFeatures(v []BTMFeature134)`

SetSubFeatures sets SubFeatures field to given value.

### HasSubFeatures

`func (o *BTMAssemblyPatternFeature2241) HasSubFeatures() bool`

HasSubFeatures returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTMAssemblyPatternFeature2241) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTMAssemblyPatternFeature2241) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTMAssemblyPatternFeature2241) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTMAssemblyPatternFeature2241) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTMAssemblyPatternFeature2241) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTMAssemblyPatternFeature2241) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTMAssemblyPatternFeature2241) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTMAssemblyPatternFeature2241) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetVariableStudioReference

`func (o *BTMAssemblyPatternFeature2241) GetVariableStudioReference() bool`

GetVariableStudioReference returns the VariableStudioReference field if non-nil, zero value otherwise.

### GetVariableStudioReferenceOk

`func (o *BTMAssemblyPatternFeature2241) GetVariableStudioReferenceOk() (*bool, bool)`

GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudioReference

`func (o *BTMAssemblyPatternFeature2241) SetVariableStudioReference(v bool)`

SetVariableStudioReference sets VariableStudioReference field to given value.

### HasVariableStudioReference

`func (o *BTMAssemblyPatternFeature2241) HasVariableStudioReference() bool`

HasVariableStudioReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTMAssemblyPatternFeature2241) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTMAssemblyPatternFeature2241) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTMAssemblyPatternFeature2241) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTMAssemblyPatternFeature2241) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


