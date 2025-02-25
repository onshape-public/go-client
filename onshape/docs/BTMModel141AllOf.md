# BTMModel141AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**AllFeaturesAndOtherReferences** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**AllFeaturesAndSubFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**ChildNodeIdToIndex** | Pointer to **map[string]int32** |  | [optional] 
**ConfigurableTreeNodes** | Pointer to [**[]BTConfigurableTreeNode**](BTConfigurableTreeNode.md) |  | [optional] 
**ConfigurationData** | Pointer to [**BTMConfigurationData1560**](BTMConfigurationData1560.md) |  | [optional] 
**Configured** | Pointer to **bool** |  | [optional] 
**DeepImports** | Pointer to [**map[string][]BTImport**](array.md) |  | [optional] 
**DefaultFeatures** | Pointer to [**BTDefaultFeatures119**](BTDefaultFeatures119.md) |  | [optional] 
**DefaultUnits** | Pointer to [**BTMUnitsDefault160**](BTMUnitsDefault160.md) |  | [optional] 
**FeatureImports** | Pointer to [**map[string][]BTImport**](array.md) |  | [optional] 
**FirstRollbackIndex** | Pointer to **int32** |  | [optional] 
**ImportSet** | Pointer to [**[]BTPModuleId235**](BTPModuleId235.md) |  | [optional] 
**Imports** | Pointer to [**[]BTMImport136**](BTMImport136.md) |  | [optional] 
**IsVariableStudio** | Pointer to **bool** |  | [optional] 
**LastFeatureBeforeRollBack** | Pointer to [**BTMFeature134**](BTMFeature134.md) |  | [optional] 
**ModelAnnotations** | Pointer to [**BTModelAnnotations3945**](BTModelAnnotations3945.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartProperties** | Pointer to [**BTPartProperties293**](BTPartProperties293.md) |  | [optional] 
**PathToCache** | Pointer to [**BTCacheDataPath191**](BTCacheDataPath191.md) |  | [optional] 
**Properties** | Pointer to [**BTModelProperties1258**](BTModelProperties1258.md) |  | [optional] 
**RollbackBar** | Pointer to [**BTMRollback150**](BTMRollback150.md) |  | [optional] 
**RolledBackToEnd** | Pointer to **bool** |  | [optional] 
**VariableStudios** | Pointer to [**[]BTMVariableStudioReference2764**](BTMVariableStudioReference2764.md) |  | [optional] 

## Methods

### NewBTMModel141AllOf

`func NewBTMModel141AllOf() *BTMModel141AllOf`

NewBTMModel141AllOf instantiates a new BTMModel141AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMModel141AllOfWithDefaults

`func NewBTMModel141AllOfWithDefaults() *BTMModel141AllOf`

NewBTMModel141AllOfWithDefaults instantiates a new BTMModel141AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFeatures

`func (o *BTMModel141AllOf) GetAllFeatures() []BTMFeature134`

GetAllFeatures returns the AllFeatures field if non-nil, zero value otherwise.

### GetAllFeaturesOk

`func (o *BTMModel141AllOf) GetAllFeaturesOk() (*[]BTMFeature134, bool)`

GetAllFeaturesOk returns a tuple with the AllFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFeatures

`func (o *BTMModel141AllOf) SetAllFeatures(v []BTMFeature134)`

SetAllFeatures sets AllFeatures field to given value.

### HasAllFeatures

`func (o *BTMModel141AllOf) HasAllFeatures() bool`

HasAllFeatures returns a boolean if a field has been set.

### GetAllFeaturesAndOtherReferences

`func (o *BTMModel141AllOf) GetAllFeaturesAndOtherReferences() []BTMFeature134`

GetAllFeaturesAndOtherReferences returns the AllFeaturesAndOtherReferences field if non-nil, zero value otherwise.

### GetAllFeaturesAndOtherReferencesOk

`func (o *BTMModel141AllOf) GetAllFeaturesAndOtherReferencesOk() (*[]BTMFeature134, bool)`

GetAllFeaturesAndOtherReferencesOk returns a tuple with the AllFeaturesAndOtherReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFeaturesAndOtherReferences

`func (o *BTMModel141AllOf) SetAllFeaturesAndOtherReferences(v []BTMFeature134)`

SetAllFeaturesAndOtherReferences sets AllFeaturesAndOtherReferences field to given value.

### HasAllFeaturesAndOtherReferences

`func (o *BTMModel141AllOf) HasAllFeaturesAndOtherReferences() bool`

HasAllFeaturesAndOtherReferences returns a boolean if a field has been set.

### GetAllFeaturesAndSubFeatures

`func (o *BTMModel141AllOf) GetAllFeaturesAndSubFeatures() []BTMFeature134`

GetAllFeaturesAndSubFeatures returns the AllFeaturesAndSubFeatures field if non-nil, zero value otherwise.

### GetAllFeaturesAndSubFeaturesOk

`func (o *BTMModel141AllOf) GetAllFeaturesAndSubFeaturesOk() (*[]BTMFeature134, bool)`

GetAllFeaturesAndSubFeaturesOk returns a tuple with the AllFeaturesAndSubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFeaturesAndSubFeatures

`func (o *BTMModel141AllOf) SetAllFeaturesAndSubFeatures(v []BTMFeature134)`

SetAllFeaturesAndSubFeatures sets AllFeaturesAndSubFeatures field to given value.

### HasAllFeaturesAndSubFeatures

`func (o *BTMModel141AllOf) HasAllFeaturesAndSubFeatures() bool`

HasAllFeaturesAndSubFeatures returns a boolean if a field has been set.

### GetBtType

`func (o *BTMModel141AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMModel141AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMModel141AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMModel141AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetChildNodeIdToIndex

`func (o *BTMModel141AllOf) GetChildNodeIdToIndex() map[string]int32`

GetChildNodeIdToIndex returns the ChildNodeIdToIndex field if non-nil, zero value otherwise.

### GetChildNodeIdToIndexOk

`func (o *BTMModel141AllOf) GetChildNodeIdToIndexOk() (*map[string]int32, bool)`

GetChildNodeIdToIndexOk returns a tuple with the ChildNodeIdToIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildNodeIdToIndex

`func (o *BTMModel141AllOf) SetChildNodeIdToIndex(v map[string]int32)`

SetChildNodeIdToIndex sets ChildNodeIdToIndex field to given value.

### HasChildNodeIdToIndex

`func (o *BTMModel141AllOf) HasChildNodeIdToIndex() bool`

HasChildNodeIdToIndex returns a boolean if a field has been set.

### GetConfigurableTreeNodes

`func (o *BTMModel141AllOf) GetConfigurableTreeNodes() []BTConfigurableTreeNode`

GetConfigurableTreeNodes returns the ConfigurableTreeNodes field if non-nil, zero value otherwise.

### GetConfigurableTreeNodesOk

`func (o *BTMModel141AllOf) GetConfigurableTreeNodesOk() (*[]BTConfigurableTreeNode, bool)`

GetConfigurableTreeNodesOk returns a tuple with the ConfigurableTreeNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableTreeNodes

`func (o *BTMModel141AllOf) SetConfigurableTreeNodes(v []BTConfigurableTreeNode)`

SetConfigurableTreeNodes sets ConfigurableTreeNodes field to given value.

### HasConfigurableTreeNodes

`func (o *BTMModel141AllOf) HasConfigurableTreeNodes() bool`

HasConfigurableTreeNodes returns a boolean if a field has been set.

### GetConfigurationData

`func (o *BTMModel141AllOf) GetConfigurationData() BTMConfigurationData1560`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *BTMModel141AllOf) GetConfigurationDataOk() (*BTMConfigurationData1560, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *BTMModel141AllOf) SetConfigurationData(v BTMConfigurationData1560)`

SetConfigurationData sets ConfigurationData field to given value.

### HasConfigurationData

`func (o *BTMModel141AllOf) HasConfigurationData() bool`

HasConfigurationData returns a boolean if a field has been set.

### GetConfigured

`func (o *BTMModel141AllOf) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *BTMModel141AllOf) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *BTMModel141AllOf) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *BTMModel141AllOf) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDeepImports

`func (o *BTMModel141AllOf) GetDeepImports() map[string][]BTImport`

GetDeepImports returns the DeepImports field if non-nil, zero value otherwise.

### GetDeepImportsOk

`func (o *BTMModel141AllOf) GetDeepImportsOk() (*map[string][]BTImport, bool)`

GetDeepImportsOk returns a tuple with the DeepImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepImports

`func (o *BTMModel141AllOf) SetDeepImports(v map[string][]BTImport)`

SetDeepImports sets DeepImports field to given value.

### HasDeepImports

`func (o *BTMModel141AllOf) HasDeepImports() bool`

HasDeepImports returns a boolean if a field has been set.

### GetDefaultFeatures

`func (o *BTMModel141AllOf) GetDefaultFeatures() BTDefaultFeatures119`

GetDefaultFeatures returns the DefaultFeatures field if non-nil, zero value otherwise.

### GetDefaultFeaturesOk

`func (o *BTMModel141AllOf) GetDefaultFeaturesOk() (*BTDefaultFeatures119, bool)`

GetDefaultFeaturesOk returns a tuple with the DefaultFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFeatures

`func (o *BTMModel141AllOf) SetDefaultFeatures(v BTDefaultFeatures119)`

SetDefaultFeatures sets DefaultFeatures field to given value.

### HasDefaultFeatures

`func (o *BTMModel141AllOf) HasDefaultFeatures() bool`

HasDefaultFeatures returns a boolean if a field has been set.

### GetDefaultUnits

`func (o *BTMModel141AllOf) GetDefaultUnits() BTMUnitsDefault160`

GetDefaultUnits returns the DefaultUnits field if non-nil, zero value otherwise.

### GetDefaultUnitsOk

`func (o *BTMModel141AllOf) GetDefaultUnitsOk() (*BTMUnitsDefault160, bool)`

GetDefaultUnitsOk returns a tuple with the DefaultUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUnits

`func (o *BTMModel141AllOf) SetDefaultUnits(v BTMUnitsDefault160)`

SetDefaultUnits sets DefaultUnits field to given value.

### HasDefaultUnits

`func (o *BTMModel141AllOf) HasDefaultUnits() bool`

HasDefaultUnits returns a boolean if a field has been set.

### GetFeatureImports

`func (o *BTMModel141AllOf) GetFeatureImports() map[string][]BTImport`

GetFeatureImports returns the FeatureImports field if non-nil, zero value otherwise.

### GetFeatureImportsOk

`func (o *BTMModel141AllOf) GetFeatureImportsOk() (*map[string][]BTImport, bool)`

GetFeatureImportsOk returns a tuple with the FeatureImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureImports

`func (o *BTMModel141AllOf) SetFeatureImports(v map[string][]BTImport)`

SetFeatureImports sets FeatureImports field to given value.

### HasFeatureImports

`func (o *BTMModel141AllOf) HasFeatureImports() bool`

HasFeatureImports returns a boolean if a field has been set.

### GetFirstRollbackIndex

`func (o *BTMModel141AllOf) GetFirstRollbackIndex() int32`

GetFirstRollbackIndex returns the FirstRollbackIndex field if non-nil, zero value otherwise.

### GetFirstRollbackIndexOk

`func (o *BTMModel141AllOf) GetFirstRollbackIndexOk() (*int32, bool)`

GetFirstRollbackIndexOk returns a tuple with the FirstRollbackIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRollbackIndex

`func (o *BTMModel141AllOf) SetFirstRollbackIndex(v int32)`

SetFirstRollbackIndex sets FirstRollbackIndex field to given value.

### HasFirstRollbackIndex

`func (o *BTMModel141AllOf) HasFirstRollbackIndex() bool`

HasFirstRollbackIndex returns a boolean if a field has been set.

### GetImportSet

`func (o *BTMModel141AllOf) GetImportSet() []BTPModuleId235`

GetImportSet returns the ImportSet field if non-nil, zero value otherwise.

### GetImportSetOk

`func (o *BTMModel141AllOf) GetImportSetOk() (*[]BTPModuleId235, bool)`

GetImportSetOk returns a tuple with the ImportSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSet

`func (o *BTMModel141AllOf) SetImportSet(v []BTPModuleId235)`

SetImportSet sets ImportSet field to given value.

### HasImportSet

`func (o *BTMModel141AllOf) HasImportSet() bool`

HasImportSet returns a boolean if a field has been set.

### GetImports

`func (o *BTMModel141AllOf) GetImports() []BTMImport136`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *BTMModel141AllOf) GetImportsOk() (*[]BTMImport136, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *BTMModel141AllOf) SetImports(v []BTMImport136)`

SetImports sets Imports field to given value.

### HasImports

`func (o *BTMModel141AllOf) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetIsVariableStudio

`func (o *BTMModel141AllOf) GetIsVariableStudio() bool`

GetIsVariableStudio returns the IsVariableStudio field if non-nil, zero value otherwise.

### GetIsVariableStudioOk

`func (o *BTMModel141AllOf) GetIsVariableStudioOk() (*bool, bool)`

GetIsVariableStudioOk returns a tuple with the IsVariableStudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariableStudio

`func (o *BTMModel141AllOf) SetIsVariableStudio(v bool)`

SetIsVariableStudio sets IsVariableStudio field to given value.

### HasIsVariableStudio

`func (o *BTMModel141AllOf) HasIsVariableStudio() bool`

HasIsVariableStudio returns a boolean if a field has been set.

### GetLastFeatureBeforeRollBack

`func (o *BTMModel141AllOf) GetLastFeatureBeforeRollBack() BTMFeature134`

GetLastFeatureBeforeRollBack returns the LastFeatureBeforeRollBack field if non-nil, zero value otherwise.

### GetLastFeatureBeforeRollBackOk

`func (o *BTMModel141AllOf) GetLastFeatureBeforeRollBackOk() (*BTMFeature134, bool)`

GetLastFeatureBeforeRollBackOk returns a tuple with the LastFeatureBeforeRollBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFeatureBeforeRollBack

`func (o *BTMModel141AllOf) SetLastFeatureBeforeRollBack(v BTMFeature134)`

SetLastFeatureBeforeRollBack sets LastFeatureBeforeRollBack field to given value.

### HasLastFeatureBeforeRollBack

`func (o *BTMModel141AllOf) HasLastFeatureBeforeRollBack() bool`

HasLastFeatureBeforeRollBack returns a boolean if a field has been set.

### GetModelAnnotations

`func (o *BTMModel141AllOf) GetModelAnnotations() BTModelAnnotations3945`

GetModelAnnotations returns the ModelAnnotations field if non-nil, zero value otherwise.

### GetModelAnnotationsOk

`func (o *BTMModel141AllOf) GetModelAnnotationsOk() (*BTModelAnnotations3945, bool)`

GetModelAnnotationsOk returns a tuple with the ModelAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelAnnotations

`func (o *BTMModel141AllOf) SetModelAnnotations(v BTModelAnnotations3945)`

SetModelAnnotations sets ModelAnnotations field to given value.

### HasModelAnnotations

`func (o *BTMModel141AllOf) HasModelAnnotations() bool`

HasModelAnnotations returns a boolean if a field has been set.

### GetName

`func (o *BTMModel141AllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMModel141AllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMModel141AllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMModel141AllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartProperties

`func (o *BTMModel141AllOf) GetPartProperties() BTPartProperties293`

GetPartProperties returns the PartProperties field if non-nil, zero value otherwise.

### GetPartPropertiesOk

`func (o *BTMModel141AllOf) GetPartPropertiesOk() (*BTPartProperties293, bool)`

GetPartPropertiesOk returns a tuple with the PartProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartProperties

`func (o *BTMModel141AllOf) SetPartProperties(v BTPartProperties293)`

SetPartProperties sets PartProperties field to given value.

### HasPartProperties

`func (o *BTMModel141AllOf) HasPartProperties() bool`

HasPartProperties returns a boolean if a field has been set.

### GetPathToCache

`func (o *BTMModel141AllOf) GetPathToCache() BTCacheDataPath191`

GetPathToCache returns the PathToCache field if non-nil, zero value otherwise.

### GetPathToCacheOk

`func (o *BTMModel141AllOf) GetPathToCacheOk() (*BTCacheDataPath191, bool)`

GetPathToCacheOk returns a tuple with the PathToCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathToCache

`func (o *BTMModel141AllOf) SetPathToCache(v BTCacheDataPath191)`

SetPathToCache sets PathToCache field to given value.

### HasPathToCache

`func (o *BTMModel141AllOf) HasPathToCache() bool`

HasPathToCache returns a boolean if a field has been set.

### GetProperties

`func (o *BTMModel141AllOf) GetProperties() BTModelProperties1258`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTMModel141AllOf) GetPropertiesOk() (*BTModelProperties1258, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTMModel141AllOf) SetProperties(v BTModelProperties1258)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTMModel141AllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRollbackBar

`func (o *BTMModel141AllOf) GetRollbackBar() BTMRollback150`

GetRollbackBar returns the RollbackBar field if non-nil, zero value otherwise.

### GetRollbackBarOk

`func (o *BTMModel141AllOf) GetRollbackBarOk() (*BTMRollback150, bool)`

GetRollbackBarOk returns a tuple with the RollbackBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackBar

`func (o *BTMModel141AllOf) SetRollbackBar(v BTMRollback150)`

SetRollbackBar sets RollbackBar field to given value.

### HasRollbackBar

`func (o *BTMModel141AllOf) HasRollbackBar() bool`

HasRollbackBar returns a boolean if a field has been set.

### GetRolledBackToEnd

`func (o *BTMModel141AllOf) GetRolledBackToEnd() bool`

GetRolledBackToEnd returns the RolledBackToEnd field if non-nil, zero value otherwise.

### GetRolledBackToEndOk

`func (o *BTMModel141AllOf) GetRolledBackToEndOk() (*bool, bool)`

GetRolledBackToEndOk returns a tuple with the RolledBackToEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledBackToEnd

`func (o *BTMModel141AllOf) SetRolledBackToEnd(v bool)`

SetRolledBackToEnd sets RolledBackToEnd field to given value.

### HasRolledBackToEnd

`func (o *BTMModel141AllOf) HasRolledBackToEnd() bool`

HasRolledBackToEnd returns a boolean if a field has been set.

### GetVariableStudios

`func (o *BTMModel141AllOf) GetVariableStudios() []BTMVariableStudioReference2764`

GetVariableStudios returns the VariableStudios field if non-nil, zero value otherwise.

### GetVariableStudiosOk

`func (o *BTMModel141AllOf) GetVariableStudiosOk() (*[]BTMVariableStudioReference2764, bool)`

GetVariableStudiosOk returns a tuple with the VariableStudios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStudios

`func (o *BTMModel141AllOf) SetVariableStudios(v []BTMVariableStudioReference2764)`

SetVariableStudios sets VariableStudios field to given value.

### HasVariableStudios

`func (o *BTMModel141AllOf) HasVariableStudios() bool`

HasVariableStudios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


