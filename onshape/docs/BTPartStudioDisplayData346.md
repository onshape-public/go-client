# BTPartStudioDisplayData346

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppearanceIdToAppearanceOverride** | Pointer to [**map[string]BTAppearanceOverride2517**](BTAppearanceOverride2517.md) |  | [optional] 
**AssemblyReferenceDisplayData** | Pointer to [**BTAssemblyReferencesDisplayData1562**](BTAssemblyReferencesDisplayData1562.md) |  | [optional] 
**BodyIdToEntityAppearanceSettings** | Pointer to [**map[string]BTBaseEntityAppearanceSettings1391**](BTBaseEntityAppearanceSettings1391.md) |  | [optional] 
**BodyIdToEntityAppearanceSettingsChanged** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**CacheablePartStudioDisplayDataVersion** | Pointer to [**GBTPartStudioDisplayDataVersion**](GBTPartStudioDisplayDataVersion.md) |  | [optional] 
**DecalIdToDecal** | Pointer to [**map[string]BTDecal2404**](BTDecal2404.md) |  | [optional] 
**DeterministicIdToAssociatedFeatureIds** | Pointer to **map[string][]string** |  | [optional] 
**DeterministicIdToEntity** | Pointer to [**map[string]BTBaseEntityData33**](BTBaseEntityData33.md) |  | [optional] 
**DeterministicIdToPartDisplayData** | Pointer to [**map[string]BTPartDisplayData17**](BTPartDisplayData17.md) |  | [optional] 
**DeterministicPartIdToData** | Pointer to [**map[string]BTPartData16**](BTPartData16.md) |  | [optional] 
**Dimensions** | Pointer to [**[]BTDimensionDisplayData323**](BTDimensionDisplayData323.md) |  | [optional] 
**DisplayStateId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**FeatureIdToOperationIndices** | Pointer to **map[string][]int32** |  | [optional] 
**FromCache** | Pointer to **bool** |  | [optional] 
**FromFullElementId** | Pointer to [**BTFullElementId756**](BTFullElementId756.md) |  | [optional] 
**FullElementId** | Pointer to [**BTFullElementId756**](BTFullElementId756.md) |  | [optional] 
**Incremental** | Pointer to **bool** |  | [optional] 
**InstanceCount** | Pointer to **int32** |  | [optional] 
**IsBase** | Pointer to **bool** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**IsNoop** | Pointer to **bool** |  | [optional] 
**KeepFromMicroversion** | Pointer to **bool** |  | [optional] 
**MicroversionId** | Pointer to [**BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**MicroversionIdAndConfigurationInterval** | Pointer to [**BTMicroversionIdAndConfigurationInterval2364**](BTMicroversionIdAndConfigurationInterval2364.md) |  | [optional] 
**MicroversionInterval** | Pointer to [**BTMicroversionIdInterval367**](BTMicroversionIdInterval367.md) |  | [optional] 
**NumberOfSketchEntities** | Pointer to **int32** |  | [optional] 
**PartColorCycle** | Pointer to [**BTBasePartColorCycle2614**](BTBasePartColorCycle2614.md) |  | [optional] 
**PartDisplayData** | Pointer to [**[]BTPartDisplayData17**](BTPartDisplayData17.md) |  | [optional] 
**SketchImages** | Pointer to [**map[string]map[string]BTSketchImageDisplayData1379**](map.md) |  | [optional] 
**UpdatedParts** | Pointer to **[]string** |  | [optional] 
**Usage** | Pointer to [**GBTDisplayDataUsage**](GBTDisplayDataUsage.md) |  | [optional] 
**UsesMultipleTessellationSettings** | Pointer to **bool** |  | [optional] 
**VersionForRasterization** | Pointer to [**BTElementDisplayData326**](BTElementDisplayData326.md) |  | [optional] 

## Methods

### NewBTPartStudioDisplayData346

`func NewBTPartStudioDisplayData346() *BTPartStudioDisplayData346`

NewBTPartStudioDisplayData346 instantiates a new BTPartStudioDisplayData346 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPartStudioDisplayData346WithDefaults

`func NewBTPartStudioDisplayData346WithDefaults() *BTPartStudioDisplayData346`

NewBTPartStudioDisplayData346WithDefaults instantiates a new BTPartStudioDisplayData346 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearanceIdToAppearanceOverride

`func (o *BTPartStudioDisplayData346) GetAppearanceIdToAppearanceOverride() map[string]BTAppearanceOverride2517`

GetAppearanceIdToAppearanceOverride returns the AppearanceIdToAppearanceOverride field if non-nil, zero value otherwise.

### GetAppearanceIdToAppearanceOverrideOk

`func (o *BTPartStudioDisplayData346) GetAppearanceIdToAppearanceOverrideOk() (*map[string]BTAppearanceOverride2517, bool)`

GetAppearanceIdToAppearanceOverrideOk returns a tuple with the AppearanceIdToAppearanceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearanceIdToAppearanceOverride

`func (o *BTPartStudioDisplayData346) SetAppearanceIdToAppearanceOverride(v map[string]BTAppearanceOverride2517)`

SetAppearanceIdToAppearanceOverride sets AppearanceIdToAppearanceOverride field to given value.

### HasAppearanceIdToAppearanceOverride

`func (o *BTPartStudioDisplayData346) HasAppearanceIdToAppearanceOverride() bool`

HasAppearanceIdToAppearanceOverride returns a boolean if a field has been set.

### GetAssemblyReferenceDisplayData

`func (o *BTPartStudioDisplayData346) GetAssemblyReferenceDisplayData() BTAssemblyReferencesDisplayData1562`

GetAssemblyReferenceDisplayData returns the AssemblyReferenceDisplayData field if non-nil, zero value otherwise.

### GetAssemblyReferenceDisplayDataOk

`func (o *BTPartStudioDisplayData346) GetAssemblyReferenceDisplayDataOk() (*BTAssemblyReferencesDisplayData1562, bool)`

GetAssemblyReferenceDisplayDataOk returns a tuple with the AssemblyReferenceDisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyReferenceDisplayData

`func (o *BTPartStudioDisplayData346) SetAssemblyReferenceDisplayData(v BTAssemblyReferencesDisplayData1562)`

SetAssemblyReferenceDisplayData sets AssemblyReferenceDisplayData field to given value.

### HasAssemblyReferenceDisplayData

`func (o *BTPartStudioDisplayData346) HasAssemblyReferenceDisplayData() bool`

HasAssemblyReferenceDisplayData returns a boolean if a field has been set.

### GetBodyIdToEntityAppearanceSettings

`func (o *BTPartStudioDisplayData346) GetBodyIdToEntityAppearanceSettings() map[string]BTBaseEntityAppearanceSettings1391`

GetBodyIdToEntityAppearanceSettings returns the BodyIdToEntityAppearanceSettings field if non-nil, zero value otherwise.

### GetBodyIdToEntityAppearanceSettingsOk

`func (o *BTPartStudioDisplayData346) GetBodyIdToEntityAppearanceSettingsOk() (*map[string]BTBaseEntityAppearanceSettings1391, bool)`

GetBodyIdToEntityAppearanceSettingsOk returns a tuple with the BodyIdToEntityAppearanceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyIdToEntityAppearanceSettings

`func (o *BTPartStudioDisplayData346) SetBodyIdToEntityAppearanceSettings(v map[string]BTBaseEntityAppearanceSettings1391)`

SetBodyIdToEntityAppearanceSettings sets BodyIdToEntityAppearanceSettings field to given value.

### HasBodyIdToEntityAppearanceSettings

`func (o *BTPartStudioDisplayData346) HasBodyIdToEntityAppearanceSettings() bool`

HasBodyIdToEntityAppearanceSettings returns a boolean if a field has been set.

### GetBodyIdToEntityAppearanceSettingsChanged

`func (o *BTPartStudioDisplayData346) GetBodyIdToEntityAppearanceSettingsChanged() bool`

GetBodyIdToEntityAppearanceSettingsChanged returns the BodyIdToEntityAppearanceSettingsChanged field if non-nil, zero value otherwise.

### GetBodyIdToEntityAppearanceSettingsChangedOk

`func (o *BTPartStudioDisplayData346) GetBodyIdToEntityAppearanceSettingsChangedOk() (*bool, bool)`

GetBodyIdToEntityAppearanceSettingsChangedOk returns a tuple with the BodyIdToEntityAppearanceSettingsChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyIdToEntityAppearanceSettingsChanged

`func (o *BTPartStudioDisplayData346) SetBodyIdToEntityAppearanceSettingsChanged(v bool)`

SetBodyIdToEntityAppearanceSettingsChanged sets BodyIdToEntityAppearanceSettingsChanged field to given value.

### HasBodyIdToEntityAppearanceSettingsChanged

`func (o *BTPartStudioDisplayData346) HasBodyIdToEntityAppearanceSettingsChanged() bool`

HasBodyIdToEntityAppearanceSettingsChanged returns a boolean if a field has been set.

### GetBtType

`func (o *BTPartStudioDisplayData346) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPartStudioDisplayData346) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPartStudioDisplayData346) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPartStudioDisplayData346) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCacheablePartStudioDisplayDataVersion

`func (o *BTPartStudioDisplayData346) GetCacheablePartStudioDisplayDataVersion() GBTPartStudioDisplayDataVersion`

GetCacheablePartStudioDisplayDataVersion returns the CacheablePartStudioDisplayDataVersion field if non-nil, zero value otherwise.

### GetCacheablePartStudioDisplayDataVersionOk

`func (o *BTPartStudioDisplayData346) GetCacheablePartStudioDisplayDataVersionOk() (*GBTPartStudioDisplayDataVersion, bool)`

GetCacheablePartStudioDisplayDataVersionOk returns a tuple with the CacheablePartStudioDisplayDataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheablePartStudioDisplayDataVersion

`func (o *BTPartStudioDisplayData346) SetCacheablePartStudioDisplayDataVersion(v GBTPartStudioDisplayDataVersion)`

SetCacheablePartStudioDisplayDataVersion sets CacheablePartStudioDisplayDataVersion field to given value.

### HasCacheablePartStudioDisplayDataVersion

`func (o *BTPartStudioDisplayData346) HasCacheablePartStudioDisplayDataVersion() bool`

HasCacheablePartStudioDisplayDataVersion returns a boolean if a field has been set.

### GetDecalIdToDecal

`func (o *BTPartStudioDisplayData346) GetDecalIdToDecal() map[string]BTDecal2404`

GetDecalIdToDecal returns the DecalIdToDecal field if non-nil, zero value otherwise.

### GetDecalIdToDecalOk

`func (o *BTPartStudioDisplayData346) GetDecalIdToDecalOk() (*map[string]BTDecal2404, bool)`

GetDecalIdToDecalOk returns a tuple with the DecalIdToDecal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecalIdToDecal

`func (o *BTPartStudioDisplayData346) SetDecalIdToDecal(v map[string]BTDecal2404)`

SetDecalIdToDecal sets DecalIdToDecal field to given value.

### HasDecalIdToDecal

`func (o *BTPartStudioDisplayData346) HasDecalIdToDecal() bool`

HasDecalIdToDecal returns a boolean if a field has been set.

### GetDeterministicIdToAssociatedFeatureIds

`func (o *BTPartStudioDisplayData346) GetDeterministicIdToAssociatedFeatureIds() map[string][]string`

GetDeterministicIdToAssociatedFeatureIds returns the DeterministicIdToAssociatedFeatureIds field if non-nil, zero value otherwise.

### GetDeterministicIdToAssociatedFeatureIdsOk

`func (o *BTPartStudioDisplayData346) GetDeterministicIdToAssociatedFeatureIdsOk() (*map[string][]string, bool)`

GetDeterministicIdToAssociatedFeatureIdsOk returns a tuple with the DeterministicIdToAssociatedFeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicIdToAssociatedFeatureIds

`func (o *BTPartStudioDisplayData346) SetDeterministicIdToAssociatedFeatureIds(v map[string][]string)`

SetDeterministicIdToAssociatedFeatureIds sets DeterministicIdToAssociatedFeatureIds field to given value.

### HasDeterministicIdToAssociatedFeatureIds

`func (o *BTPartStudioDisplayData346) HasDeterministicIdToAssociatedFeatureIds() bool`

HasDeterministicIdToAssociatedFeatureIds returns a boolean if a field has been set.

### GetDeterministicIdToEntity

`func (o *BTPartStudioDisplayData346) GetDeterministicIdToEntity() map[string]BTBaseEntityData33`

GetDeterministicIdToEntity returns the DeterministicIdToEntity field if non-nil, zero value otherwise.

### GetDeterministicIdToEntityOk

`func (o *BTPartStudioDisplayData346) GetDeterministicIdToEntityOk() (*map[string]BTBaseEntityData33, bool)`

GetDeterministicIdToEntityOk returns a tuple with the DeterministicIdToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicIdToEntity

`func (o *BTPartStudioDisplayData346) SetDeterministicIdToEntity(v map[string]BTBaseEntityData33)`

SetDeterministicIdToEntity sets DeterministicIdToEntity field to given value.

### HasDeterministicIdToEntity

`func (o *BTPartStudioDisplayData346) HasDeterministicIdToEntity() bool`

HasDeterministicIdToEntity returns a boolean if a field has been set.

### GetDeterministicIdToPartDisplayData

`func (o *BTPartStudioDisplayData346) GetDeterministicIdToPartDisplayData() map[string]BTPartDisplayData17`

GetDeterministicIdToPartDisplayData returns the DeterministicIdToPartDisplayData field if non-nil, zero value otherwise.

### GetDeterministicIdToPartDisplayDataOk

`func (o *BTPartStudioDisplayData346) GetDeterministicIdToPartDisplayDataOk() (*map[string]BTPartDisplayData17, bool)`

GetDeterministicIdToPartDisplayDataOk returns a tuple with the DeterministicIdToPartDisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicIdToPartDisplayData

`func (o *BTPartStudioDisplayData346) SetDeterministicIdToPartDisplayData(v map[string]BTPartDisplayData17)`

SetDeterministicIdToPartDisplayData sets DeterministicIdToPartDisplayData field to given value.

### HasDeterministicIdToPartDisplayData

`func (o *BTPartStudioDisplayData346) HasDeterministicIdToPartDisplayData() bool`

HasDeterministicIdToPartDisplayData returns a boolean if a field has been set.

### GetDeterministicPartIdToData

`func (o *BTPartStudioDisplayData346) GetDeterministicPartIdToData() map[string]BTPartData16`

GetDeterministicPartIdToData returns the DeterministicPartIdToData field if non-nil, zero value otherwise.

### GetDeterministicPartIdToDataOk

`func (o *BTPartStudioDisplayData346) GetDeterministicPartIdToDataOk() (*map[string]BTPartData16, bool)`

GetDeterministicPartIdToDataOk returns a tuple with the DeterministicPartIdToData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicPartIdToData

`func (o *BTPartStudioDisplayData346) SetDeterministicPartIdToData(v map[string]BTPartData16)`

SetDeterministicPartIdToData sets DeterministicPartIdToData field to given value.

### HasDeterministicPartIdToData

`func (o *BTPartStudioDisplayData346) HasDeterministicPartIdToData() bool`

HasDeterministicPartIdToData returns a boolean if a field has been set.

### GetDimensions

`func (o *BTPartStudioDisplayData346) GetDimensions() []BTDimensionDisplayData323`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *BTPartStudioDisplayData346) GetDimensionsOk() (*[]BTDimensionDisplayData323, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *BTPartStudioDisplayData346) SetDimensions(v []BTDimensionDisplayData323)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *BTPartStudioDisplayData346) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetDisplayStateId

`func (o *BTPartStudioDisplayData346) GetDisplayStateId() string`

GetDisplayStateId returns the DisplayStateId field if non-nil, zero value otherwise.

### GetDisplayStateIdOk

`func (o *BTPartStudioDisplayData346) GetDisplayStateIdOk() (*string, bool)`

GetDisplayStateIdOk returns a tuple with the DisplayStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStateId

`func (o *BTPartStudioDisplayData346) SetDisplayStateId(v string)`

SetDisplayStateId sets DisplayStateId field to given value.

### HasDisplayStateId

`func (o *BTPartStudioDisplayData346) HasDisplayStateId() bool`

HasDisplayStateId returns a boolean if a field has been set.

### GetElementId

`func (o *BTPartStudioDisplayData346) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTPartStudioDisplayData346) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTPartStudioDisplayData346) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTPartStudioDisplayData346) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFeatureIdToOperationIndices

`func (o *BTPartStudioDisplayData346) GetFeatureIdToOperationIndices() map[string][]int32`

GetFeatureIdToOperationIndices returns the FeatureIdToOperationIndices field if non-nil, zero value otherwise.

### GetFeatureIdToOperationIndicesOk

`func (o *BTPartStudioDisplayData346) GetFeatureIdToOperationIndicesOk() (*map[string][]int32, bool)`

GetFeatureIdToOperationIndicesOk returns a tuple with the FeatureIdToOperationIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIdToOperationIndices

`func (o *BTPartStudioDisplayData346) SetFeatureIdToOperationIndices(v map[string][]int32)`

SetFeatureIdToOperationIndices sets FeatureIdToOperationIndices field to given value.

### HasFeatureIdToOperationIndices

`func (o *BTPartStudioDisplayData346) HasFeatureIdToOperationIndices() bool`

HasFeatureIdToOperationIndices returns a boolean if a field has been set.

### GetFromCache

`func (o *BTPartStudioDisplayData346) GetFromCache() bool`

GetFromCache returns the FromCache field if non-nil, zero value otherwise.

### GetFromCacheOk

`func (o *BTPartStudioDisplayData346) GetFromCacheOk() (*bool, bool)`

GetFromCacheOk returns a tuple with the FromCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCache

`func (o *BTPartStudioDisplayData346) SetFromCache(v bool)`

SetFromCache sets FromCache field to given value.

### HasFromCache

`func (o *BTPartStudioDisplayData346) HasFromCache() bool`

HasFromCache returns a boolean if a field has been set.

### GetFromFullElementId

`func (o *BTPartStudioDisplayData346) GetFromFullElementId() BTFullElementId756`

GetFromFullElementId returns the FromFullElementId field if non-nil, zero value otherwise.

### GetFromFullElementIdOk

`func (o *BTPartStudioDisplayData346) GetFromFullElementIdOk() (*BTFullElementId756, bool)`

GetFromFullElementIdOk returns a tuple with the FromFullElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFullElementId

`func (o *BTPartStudioDisplayData346) SetFromFullElementId(v BTFullElementId756)`

SetFromFullElementId sets FromFullElementId field to given value.

### HasFromFullElementId

`func (o *BTPartStudioDisplayData346) HasFromFullElementId() bool`

HasFromFullElementId returns a boolean if a field has been set.

### GetFullElementId

`func (o *BTPartStudioDisplayData346) GetFullElementId() BTFullElementId756`

GetFullElementId returns the FullElementId field if non-nil, zero value otherwise.

### GetFullElementIdOk

`func (o *BTPartStudioDisplayData346) GetFullElementIdOk() (*BTFullElementId756, bool)`

GetFullElementIdOk returns a tuple with the FullElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullElementId

`func (o *BTPartStudioDisplayData346) SetFullElementId(v BTFullElementId756)`

SetFullElementId sets FullElementId field to given value.

### HasFullElementId

`func (o *BTPartStudioDisplayData346) HasFullElementId() bool`

HasFullElementId returns a boolean if a field has been set.

### GetIncremental

`func (o *BTPartStudioDisplayData346) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *BTPartStudioDisplayData346) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *BTPartStudioDisplayData346) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *BTPartStudioDisplayData346) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.

### GetInstanceCount

`func (o *BTPartStudioDisplayData346) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *BTPartStudioDisplayData346) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *BTPartStudioDisplayData346) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *BTPartStudioDisplayData346) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetIsBase

`func (o *BTPartStudioDisplayData346) GetIsBase() bool`

GetIsBase returns the IsBase field if non-nil, zero value otherwise.

### GetIsBaseOk

`func (o *BTPartStudioDisplayData346) GetIsBaseOk() (*bool, bool)`

GetIsBaseOk returns a tuple with the IsBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBase

`func (o *BTPartStudioDisplayData346) SetIsBase(v bool)`

SetIsBase sets IsBase field to given value.

### HasIsBase

`func (o *BTPartStudioDisplayData346) HasIsBase() bool`

HasIsBase returns a boolean if a field has been set.

### GetIsExternal

`func (o *BTPartStudioDisplayData346) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *BTPartStudioDisplayData346) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *BTPartStudioDisplayData346) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *BTPartStudioDisplayData346) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetIsNoop

`func (o *BTPartStudioDisplayData346) GetIsNoop() bool`

GetIsNoop returns the IsNoop field if non-nil, zero value otherwise.

### GetIsNoopOk

`func (o *BTPartStudioDisplayData346) GetIsNoopOk() (*bool, bool)`

GetIsNoopOk returns a tuple with the IsNoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNoop

`func (o *BTPartStudioDisplayData346) SetIsNoop(v bool)`

SetIsNoop sets IsNoop field to given value.

### HasIsNoop

`func (o *BTPartStudioDisplayData346) HasIsNoop() bool`

HasIsNoop returns a boolean if a field has been set.

### GetKeepFromMicroversion

`func (o *BTPartStudioDisplayData346) GetKeepFromMicroversion() bool`

GetKeepFromMicroversion returns the KeepFromMicroversion field if non-nil, zero value otherwise.

### GetKeepFromMicroversionOk

`func (o *BTPartStudioDisplayData346) GetKeepFromMicroversionOk() (*bool, bool)`

GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFromMicroversion

`func (o *BTPartStudioDisplayData346) SetKeepFromMicroversion(v bool)`

SetKeepFromMicroversion sets KeepFromMicroversion field to given value.

### HasKeepFromMicroversion

`func (o *BTPartStudioDisplayData346) HasKeepFromMicroversion() bool`

HasKeepFromMicroversion returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTPartStudioDisplayData346) GetMicroversionId() BTMicroversionId366`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTPartStudioDisplayData346) GetMicroversionIdOk() (*BTMicroversionId366, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTPartStudioDisplayData346) SetMicroversionId(v BTMicroversionId366)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTPartStudioDisplayData346) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetMicroversionIdAndConfigurationInterval

`func (o *BTPartStudioDisplayData346) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364`

GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field if non-nil, zero value otherwise.

### GetMicroversionIdAndConfigurationIntervalOk

`func (o *BTPartStudioDisplayData346) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool)`

GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionIdAndConfigurationInterval

`func (o *BTPartStudioDisplayData346) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364)`

SetMicroversionIdAndConfigurationInterval sets MicroversionIdAndConfigurationInterval field to given value.

### HasMicroversionIdAndConfigurationInterval

`func (o *BTPartStudioDisplayData346) HasMicroversionIdAndConfigurationInterval() bool`

HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.

### GetMicroversionInterval

`func (o *BTPartStudioDisplayData346) GetMicroversionInterval() BTMicroversionIdInterval367`

GetMicroversionInterval returns the MicroversionInterval field if non-nil, zero value otherwise.

### GetMicroversionIntervalOk

`func (o *BTPartStudioDisplayData346) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool)`

GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionInterval

`func (o *BTPartStudioDisplayData346) SetMicroversionInterval(v BTMicroversionIdInterval367)`

SetMicroversionInterval sets MicroversionInterval field to given value.

### HasMicroversionInterval

`func (o *BTPartStudioDisplayData346) HasMicroversionInterval() bool`

HasMicroversionInterval returns a boolean if a field has been set.

### GetNumberOfSketchEntities

`func (o *BTPartStudioDisplayData346) GetNumberOfSketchEntities() int32`

GetNumberOfSketchEntities returns the NumberOfSketchEntities field if non-nil, zero value otherwise.

### GetNumberOfSketchEntitiesOk

`func (o *BTPartStudioDisplayData346) GetNumberOfSketchEntitiesOk() (*int32, bool)`

GetNumberOfSketchEntitiesOk returns a tuple with the NumberOfSketchEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSketchEntities

`func (o *BTPartStudioDisplayData346) SetNumberOfSketchEntities(v int32)`

SetNumberOfSketchEntities sets NumberOfSketchEntities field to given value.

### HasNumberOfSketchEntities

`func (o *BTPartStudioDisplayData346) HasNumberOfSketchEntities() bool`

HasNumberOfSketchEntities returns a boolean if a field has been set.

### GetPartColorCycle

`func (o *BTPartStudioDisplayData346) GetPartColorCycle() BTBasePartColorCycle2614`

GetPartColorCycle returns the PartColorCycle field if non-nil, zero value otherwise.

### GetPartColorCycleOk

`func (o *BTPartStudioDisplayData346) GetPartColorCycleOk() (*BTBasePartColorCycle2614, bool)`

GetPartColorCycleOk returns a tuple with the PartColorCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartColorCycle

`func (o *BTPartStudioDisplayData346) SetPartColorCycle(v BTBasePartColorCycle2614)`

SetPartColorCycle sets PartColorCycle field to given value.

### HasPartColorCycle

`func (o *BTPartStudioDisplayData346) HasPartColorCycle() bool`

HasPartColorCycle returns a boolean if a field has been set.

### GetPartDisplayData

`func (o *BTPartStudioDisplayData346) GetPartDisplayData() []BTPartDisplayData17`

GetPartDisplayData returns the PartDisplayData field if non-nil, zero value otherwise.

### GetPartDisplayDataOk

`func (o *BTPartStudioDisplayData346) GetPartDisplayDataOk() (*[]BTPartDisplayData17, bool)`

GetPartDisplayDataOk returns a tuple with the PartDisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartDisplayData

`func (o *BTPartStudioDisplayData346) SetPartDisplayData(v []BTPartDisplayData17)`

SetPartDisplayData sets PartDisplayData field to given value.

### HasPartDisplayData

`func (o *BTPartStudioDisplayData346) HasPartDisplayData() bool`

HasPartDisplayData returns a boolean if a field has been set.

### GetSketchImages

`func (o *BTPartStudioDisplayData346) GetSketchImages() map[string]map[string]BTSketchImageDisplayData1379`

GetSketchImages returns the SketchImages field if non-nil, zero value otherwise.

### GetSketchImagesOk

`func (o *BTPartStudioDisplayData346) GetSketchImagesOk() (*map[string]map[string]BTSketchImageDisplayData1379, bool)`

GetSketchImagesOk returns a tuple with the SketchImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchImages

`func (o *BTPartStudioDisplayData346) SetSketchImages(v map[string]map[string]BTSketchImageDisplayData1379)`

SetSketchImages sets SketchImages field to given value.

### HasSketchImages

`func (o *BTPartStudioDisplayData346) HasSketchImages() bool`

HasSketchImages returns a boolean if a field has been set.

### GetUpdatedParts

`func (o *BTPartStudioDisplayData346) GetUpdatedParts() []string`

GetUpdatedParts returns the UpdatedParts field if non-nil, zero value otherwise.

### GetUpdatedPartsOk

`func (o *BTPartStudioDisplayData346) GetUpdatedPartsOk() (*[]string, bool)`

GetUpdatedPartsOk returns a tuple with the UpdatedParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedParts

`func (o *BTPartStudioDisplayData346) SetUpdatedParts(v []string)`

SetUpdatedParts sets UpdatedParts field to given value.

### HasUpdatedParts

`func (o *BTPartStudioDisplayData346) HasUpdatedParts() bool`

HasUpdatedParts returns a boolean if a field has been set.

### GetUsage

`func (o *BTPartStudioDisplayData346) GetUsage() GBTDisplayDataUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BTPartStudioDisplayData346) GetUsageOk() (*GBTDisplayDataUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BTPartStudioDisplayData346) SetUsage(v GBTDisplayDataUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *BTPartStudioDisplayData346) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsesMultipleTessellationSettings

`func (o *BTPartStudioDisplayData346) GetUsesMultipleTessellationSettings() bool`

GetUsesMultipleTessellationSettings returns the UsesMultipleTessellationSettings field if non-nil, zero value otherwise.

### GetUsesMultipleTessellationSettingsOk

`func (o *BTPartStudioDisplayData346) GetUsesMultipleTessellationSettingsOk() (*bool, bool)`

GetUsesMultipleTessellationSettingsOk returns a tuple with the UsesMultipleTessellationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesMultipleTessellationSettings

`func (o *BTPartStudioDisplayData346) SetUsesMultipleTessellationSettings(v bool)`

SetUsesMultipleTessellationSettings sets UsesMultipleTessellationSettings field to given value.

### HasUsesMultipleTessellationSettings

`func (o *BTPartStudioDisplayData346) HasUsesMultipleTessellationSettings() bool`

HasUsesMultipleTessellationSettings returns a boolean if a field has been set.

### GetVersionForRasterization

`func (o *BTPartStudioDisplayData346) GetVersionForRasterization() BTElementDisplayData326`

GetVersionForRasterization returns the VersionForRasterization field if non-nil, zero value otherwise.

### GetVersionForRasterizationOk

`func (o *BTPartStudioDisplayData346) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool)`

GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionForRasterization

`func (o *BTPartStudioDisplayData346) SetVersionForRasterization(v BTElementDisplayData326)`

SetVersionForRasterization sets VersionForRasterization field to given value.

### HasVersionForRasterization

`func (o *BTPartStudioDisplayData346) HasVersionForRasterization() bool`

HasVersionForRasterization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


