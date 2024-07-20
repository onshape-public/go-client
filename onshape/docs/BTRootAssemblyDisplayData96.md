# BTRootAssemblyDisplayData96

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationsForElement** | Pointer to [**BTAnnotationElementDisplayData894**](BTAnnotationElementDisplayData894.md) |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**BuildDurationMillis** | Pointer to **float64** |  | [optional] 
**DeletedGeometryMateIds** | Pointer to **[]string** |  | [optional] 
**DeletedLoads** | Pointer to **[]string** |  | [optional] 
**DeletedMateConnectorIds** | Pointer to **[]string** |  | [optional] 
**DeletedMateGroupIds** | Pointer to **[]string** |  | [optional] 
**DeletedMateIds** | Pointer to **[]string** |  | [optional] 
**DeletedOccurrences** | Pointer to [**[]BTOccurrence74**](BTOccurrence74.md) |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**FromFullElementId** | Pointer to [**BTFullElementId756**](BTFullElementId756.md) |  | [optional] 
**FullElementId** | Pointer to [**BTFullElementId756**](BTFullElementId756.md) |  | [optional] 
**GeometryMates** | Pointer to [**[]BTGeometryMateDisplayData1050**](BTGeometryMateDisplayData1050.md) |  | [optional] 
**Incremental** | Pointer to **bool** |  | [optional] 
**InstanceCount** | Pointer to **int32** |  | [optional] 
**IsCollapsible** | Pointer to **bool** |  | [optional] 
**IsForInContext** | Pointer to **bool** |  | [optional] 
**KeepFromMicroversion** | Pointer to **bool** |  | [optional] 
**Loads** | Pointer to [**[]BTLoadDisplayData837**](BTLoadDisplayData837.md) |  | [optional] 
**MateConnectors** | Pointer to [**[]BTMateConnectorDisplayData94**](BTMateConnectorDisplayData94.md) |  | [optional] 
**MateGroups** | Pointer to [**[]BTMateGroupDisplayData1990**](BTMateGroupDisplayData1990.md) |  | [optional] 
**Mates** | Pointer to [**[]BTMateDisplayData1358**](BTMateDisplayData1358.md) |  | [optional] 
**MicroversionId** | Pointer to [**BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**MicroversionIdAndConfigurationInterval** | Pointer to [**BTMicroversionIdAndConfigurationInterval2364**](BTMicroversionIdAndConfigurationInterval2364.md) |  | [optional] 
**MicroversionInterval** | Pointer to [**BTMicroversionIdInterval367**](BTMicroversionIdInterval367.md) |  | [optional] 
**Occurrences** | Pointer to [**[]BTOccurrenceDisplayData95**](BTOccurrenceDisplayData95.md) |  | [optional] 
**OriginDisplayData** | Pointer to [**BTOriginDisplayData934**](BTOriginDisplayData934.md) |  | [optional] 
**PartStudioDisplayData** | Pointer to [**[]BTPartStudioDisplayDataBase2751**](BTPartStudioDisplayDataBase2751.md) |  | [optional] 
**QuickSummary** | Pointer to **string** |  | [optional] 
**SentTimeISO** | Pointer to **string** |  | [optional] 
**VersionForRasterization** | Pointer to [**BTElementDisplayData326**](BTElementDisplayData326.md) |  | [optional] 

## Methods

### NewBTRootAssemblyDisplayData96

`func NewBTRootAssemblyDisplayData96() *BTRootAssemblyDisplayData96`

NewBTRootAssemblyDisplayData96 instantiates a new BTRootAssemblyDisplayData96 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRootAssemblyDisplayData96WithDefaults

`func NewBTRootAssemblyDisplayData96WithDefaults() *BTRootAssemblyDisplayData96`

NewBTRootAssemblyDisplayData96WithDefaults instantiates a new BTRootAssemblyDisplayData96 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotationsForElement

`func (o *BTRootAssemblyDisplayData96) GetAnnotationsForElement() BTAnnotationElementDisplayData894`

GetAnnotationsForElement returns the AnnotationsForElement field if non-nil, zero value otherwise.

### GetAnnotationsForElementOk

`func (o *BTRootAssemblyDisplayData96) GetAnnotationsForElementOk() (*BTAnnotationElementDisplayData894, bool)`

GetAnnotationsForElementOk returns a tuple with the AnnotationsForElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsForElement

`func (o *BTRootAssemblyDisplayData96) SetAnnotationsForElement(v BTAnnotationElementDisplayData894)`

SetAnnotationsForElement sets AnnotationsForElement field to given value.

### HasAnnotationsForElement

`func (o *BTRootAssemblyDisplayData96) HasAnnotationsForElement() bool`

HasAnnotationsForElement returns a boolean if a field has been set.

### GetBtType

`func (o *BTRootAssemblyDisplayData96) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTRootAssemblyDisplayData96) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTRootAssemblyDisplayData96) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTRootAssemblyDisplayData96) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetBuildDurationMillis

`func (o *BTRootAssemblyDisplayData96) GetBuildDurationMillis() float64`

GetBuildDurationMillis returns the BuildDurationMillis field if non-nil, zero value otherwise.

### GetBuildDurationMillisOk

`func (o *BTRootAssemblyDisplayData96) GetBuildDurationMillisOk() (*float64, bool)`

GetBuildDurationMillisOk returns a tuple with the BuildDurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDurationMillis

`func (o *BTRootAssemblyDisplayData96) SetBuildDurationMillis(v float64)`

SetBuildDurationMillis sets BuildDurationMillis field to given value.

### HasBuildDurationMillis

`func (o *BTRootAssemblyDisplayData96) HasBuildDurationMillis() bool`

HasBuildDurationMillis returns a boolean if a field has been set.

### GetDeletedGeometryMateIds

`func (o *BTRootAssemblyDisplayData96) GetDeletedGeometryMateIds() []string`

GetDeletedGeometryMateIds returns the DeletedGeometryMateIds field if non-nil, zero value otherwise.

### GetDeletedGeometryMateIdsOk

`func (o *BTRootAssemblyDisplayData96) GetDeletedGeometryMateIdsOk() (*[]string, bool)`

GetDeletedGeometryMateIdsOk returns a tuple with the DeletedGeometryMateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedGeometryMateIds

`func (o *BTRootAssemblyDisplayData96) SetDeletedGeometryMateIds(v []string)`

SetDeletedGeometryMateIds sets DeletedGeometryMateIds field to given value.

### HasDeletedGeometryMateIds

`func (o *BTRootAssemblyDisplayData96) HasDeletedGeometryMateIds() bool`

HasDeletedGeometryMateIds returns a boolean if a field has been set.

### GetDeletedLoads

`func (o *BTRootAssemblyDisplayData96) GetDeletedLoads() []string`

GetDeletedLoads returns the DeletedLoads field if non-nil, zero value otherwise.

### GetDeletedLoadsOk

`func (o *BTRootAssemblyDisplayData96) GetDeletedLoadsOk() (*[]string, bool)`

GetDeletedLoadsOk returns a tuple with the DeletedLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedLoads

`func (o *BTRootAssemblyDisplayData96) SetDeletedLoads(v []string)`

SetDeletedLoads sets DeletedLoads field to given value.

### HasDeletedLoads

`func (o *BTRootAssemblyDisplayData96) HasDeletedLoads() bool`

HasDeletedLoads returns a boolean if a field has been set.

### GetDeletedMateConnectorIds

`func (o *BTRootAssemblyDisplayData96) GetDeletedMateConnectorIds() []string`

GetDeletedMateConnectorIds returns the DeletedMateConnectorIds field if non-nil, zero value otherwise.

### GetDeletedMateConnectorIdsOk

`func (o *BTRootAssemblyDisplayData96) GetDeletedMateConnectorIdsOk() (*[]string, bool)`

GetDeletedMateConnectorIdsOk returns a tuple with the DeletedMateConnectorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMateConnectorIds

`func (o *BTRootAssemblyDisplayData96) SetDeletedMateConnectorIds(v []string)`

SetDeletedMateConnectorIds sets DeletedMateConnectorIds field to given value.

### HasDeletedMateConnectorIds

`func (o *BTRootAssemblyDisplayData96) HasDeletedMateConnectorIds() bool`

HasDeletedMateConnectorIds returns a boolean if a field has been set.

### GetDeletedMateGroupIds

`func (o *BTRootAssemblyDisplayData96) GetDeletedMateGroupIds() []string`

GetDeletedMateGroupIds returns the DeletedMateGroupIds field if non-nil, zero value otherwise.

### GetDeletedMateGroupIdsOk

`func (o *BTRootAssemblyDisplayData96) GetDeletedMateGroupIdsOk() (*[]string, bool)`

GetDeletedMateGroupIdsOk returns a tuple with the DeletedMateGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMateGroupIds

`func (o *BTRootAssemblyDisplayData96) SetDeletedMateGroupIds(v []string)`

SetDeletedMateGroupIds sets DeletedMateGroupIds field to given value.

### HasDeletedMateGroupIds

`func (o *BTRootAssemblyDisplayData96) HasDeletedMateGroupIds() bool`

HasDeletedMateGroupIds returns a boolean if a field has been set.

### GetDeletedMateIds

`func (o *BTRootAssemblyDisplayData96) GetDeletedMateIds() []string`

GetDeletedMateIds returns the DeletedMateIds field if non-nil, zero value otherwise.

### GetDeletedMateIdsOk

`func (o *BTRootAssemblyDisplayData96) GetDeletedMateIdsOk() (*[]string, bool)`

GetDeletedMateIdsOk returns a tuple with the DeletedMateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMateIds

`func (o *BTRootAssemblyDisplayData96) SetDeletedMateIds(v []string)`

SetDeletedMateIds sets DeletedMateIds field to given value.

### HasDeletedMateIds

`func (o *BTRootAssemblyDisplayData96) HasDeletedMateIds() bool`

HasDeletedMateIds returns a boolean if a field has been set.

### GetDeletedOccurrences

`func (o *BTRootAssemblyDisplayData96) GetDeletedOccurrences() []BTOccurrence74`

GetDeletedOccurrences returns the DeletedOccurrences field if non-nil, zero value otherwise.

### GetDeletedOccurrencesOk

`func (o *BTRootAssemblyDisplayData96) GetDeletedOccurrencesOk() (*[]BTOccurrence74, bool)`

GetDeletedOccurrencesOk returns a tuple with the DeletedOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedOccurrences

`func (o *BTRootAssemblyDisplayData96) SetDeletedOccurrences(v []BTOccurrence74)`

SetDeletedOccurrences sets DeletedOccurrences field to given value.

### HasDeletedOccurrences

`func (o *BTRootAssemblyDisplayData96) HasDeletedOccurrences() bool`

HasDeletedOccurrences returns a boolean if a field has been set.

### GetElementId

`func (o *BTRootAssemblyDisplayData96) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTRootAssemblyDisplayData96) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTRootAssemblyDisplayData96) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTRootAssemblyDisplayData96) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFromFullElementId

`func (o *BTRootAssemblyDisplayData96) GetFromFullElementId() BTFullElementId756`

GetFromFullElementId returns the FromFullElementId field if non-nil, zero value otherwise.

### GetFromFullElementIdOk

`func (o *BTRootAssemblyDisplayData96) GetFromFullElementIdOk() (*BTFullElementId756, bool)`

GetFromFullElementIdOk returns a tuple with the FromFullElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFullElementId

`func (o *BTRootAssemblyDisplayData96) SetFromFullElementId(v BTFullElementId756)`

SetFromFullElementId sets FromFullElementId field to given value.

### HasFromFullElementId

`func (o *BTRootAssemblyDisplayData96) HasFromFullElementId() bool`

HasFromFullElementId returns a boolean if a field has been set.

### GetFullElementId

`func (o *BTRootAssemblyDisplayData96) GetFullElementId() BTFullElementId756`

GetFullElementId returns the FullElementId field if non-nil, zero value otherwise.

### GetFullElementIdOk

`func (o *BTRootAssemblyDisplayData96) GetFullElementIdOk() (*BTFullElementId756, bool)`

GetFullElementIdOk returns a tuple with the FullElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullElementId

`func (o *BTRootAssemblyDisplayData96) SetFullElementId(v BTFullElementId756)`

SetFullElementId sets FullElementId field to given value.

### HasFullElementId

`func (o *BTRootAssemblyDisplayData96) HasFullElementId() bool`

HasFullElementId returns a boolean if a field has been set.

### GetGeometryMates

`func (o *BTRootAssemblyDisplayData96) GetGeometryMates() []BTGeometryMateDisplayData1050`

GetGeometryMates returns the GeometryMates field if non-nil, zero value otherwise.

### GetGeometryMatesOk

`func (o *BTRootAssemblyDisplayData96) GetGeometryMatesOk() (*[]BTGeometryMateDisplayData1050, bool)`

GetGeometryMatesOk returns a tuple with the GeometryMates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometryMates

`func (o *BTRootAssemblyDisplayData96) SetGeometryMates(v []BTGeometryMateDisplayData1050)`

SetGeometryMates sets GeometryMates field to given value.

### HasGeometryMates

`func (o *BTRootAssemblyDisplayData96) HasGeometryMates() bool`

HasGeometryMates returns a boolean if a field has been set.

### GetIncremental

`func (o *BTRootAssemblyDisplayData96) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *BTRootAssemblyDisplayData96) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *BTRootAssemblyDisplayData96) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *BTRootAssemblyDisplayData96) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.

### GetInstanceCount

`func (o *BTRootAssemblyDisplayData96) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *BTRootAssemblyDisplayData96) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *BTRootAssemblyDisplayData96) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *BTRootAssemblyDisplayData96) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetIsCollapsible

`func (o *BTRootAssemblyDisplayData96) GetIsCollapsible() bool`

GetIsCollapsible returns the IsCollapsible field if non-nil, zero value otherwise.

### GetIsCollapsibleOk

`func (o *BTRootAssemblyDisplayData96) GetIsCollapsibleOk() (*bool, bool)`

GetIsCollapsibleOk returns a tuple with the IsCollapsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollapsible

`func (o *BTRootAssemblyDisplayData96) SetIsCollapsible(v bool)`

SetIsCollapsible sets IsCollapsible field to given value.

### HasIsCollapsible

`func (o *BTRootAssemblyDisplayData96) HasIsCollapsible() bool`

HasIsCollapsible returns a boolean if a field has been set.

### GetIsForInContext

`func (o *BTRootAssemblyDisplayData96) GetIsForInContext() bool`

GetIsForInContext returns the IsForInContext field if non-nil, zero value otherwise.

### GetIsForInContextOk

`func (o *BTRootAssemblyDisplayData96) GetIsForInContextOk() (*bool, bool)`

GetIsForInContextOk returns a tuple with the IsForInContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForInContext

`func (o *BTRootAssemblyDisplayData96) SetIsForInContext(v bool)`

SetIsForInContext sets IsForInContext field to given value.

### HasIsForInContext

`func (o *BTRootAssemblyDisplayData96) HasIsForInContext() bool`

HasIsForInContext returns a boolean if a field has been set.

### GetKeepFromMicroversion

`func (o *BTRootAssemblyDisplayData96) GetKeepFromMicroversion() bool`

GetKeepFromMicroversion returns the KeepFromMicroversion field if non-nil, zero value otherwise.

### GetKeepFromMicroversionOk

`func (o *BTRootAssemblyDisplayData96) GetKeepFromMicroversionOk() (*bool, bool)`

GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFromMicroversion

`func (o *BTRootAssemblyDisplayData96) SetKeepFromMicroversion(v bool)`

SetKeepFromMicroversion sets KeepFromMicroversion field to given value.

### HasKeepFromMicroversion

`func (o *BTRootAssemblyDisplayData96) HasKeepFromMicroversion() bool`

HasKeepFromMicroversion returns a boolean if a field has been set.

### GetLoads

`func (o *BTRootAssemblyDisplayData96) GetLoads() []BTLoadDisplayData837`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *BTRootAssemblyDisplayData96) GetLoadsOk() (*[]BTLoadDisplayData837, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *BTRootAssemblyDisplayData96) SetLoads(v []BTLoadDisplayData837)`

SetLoads sets Loads field to given value.

### HasLoads

`func (o *BTRootAssemblyDisplayData96) HasLoads() bool`

HasLoads returns a boolean if a field has been set.

### GetMateConnectors

`func (o *BTRootAssemblyDisplayData96) GetMateConnectors() []BTMateConnectorDisplayData94`

GetMateConnectors returns the MateConnectors field if non-nil, zero value otherwise.

### GetMateConnectorsOk

`func (o *BTRootAssemblyDisplayData96) GetMateConnectorsOk() (*[]BTMateConnectorDisplayData94, bool)`

GetMateConnectorsOk returns a tuple with the MateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectors

`func (o *BTRootAssemblyDisplayData96) SetMateConnectors(v []BTMateConnectorDisplayData94)`

SetMateConnectors sets MateConnectors field to given value.

### HasMateConnectors

`func (o *BTRootAssemblyDisplayData96) HasMateConnectors() bool`

HasMateConnectors returns a boolean if a field has been set.

### GetMateGroups

`func (o *BTRootAssemblyDisplayData96) GetMateGroups() []BTMateGroupDisplayData1990`

GetMateGroups returns the MateGroups field if non-nil, zero value otherwise.

### GetMateGroupsOk

`func (o *BTRootAssemblyDisplayData96) GetMateGroupsOk() (*[]BTMateGroupDisplayData1990, bool)`

GetMateGroupsOk returns a tuple with the MateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateGroups

`func (o *BTRootAssemblyDisplayData96) SetMateGroups(v []BTMateGroupDisplayData1990)`

SetMateGroups sets MateGroups field to given value.

### HasMateGroups

`func (o *BTRootAssemblyDisplayData96) HasMateGroups() bool`

HasMateGroups returns a boolean if a field has been set.

### GetMates

`func (o *BTRootAssemblyDisplayData96) GetMates() []BTMateDisplayData1358`

GetMates returns the Mates field if non-nil, zero value otherwise.

### GetMatesOk

`func (o *BTRootAssemblyDisplayData96) GetMatesOk() (*[]BTMateDisplayData1358, bool)`

GetMatesOk returns a tuple with the Mates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMates

`func (o *BTRootAssemblyDisplayData96) SetMates(v []BTMateDisplayData1358)`

SetMates sets Mates field to given value.

### HasMates

`func (o *BTRootAssemblyDisplayData96) HasMates() bool`

HasMates returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTRootAssemblyDisplayData96) GetMicroversionId() BTMicroversionId366`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTRootAssemblyDisplayData96) GetMicroversionIdOk() (*BTMicroversionId366, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTRootAssemblyDisplayData96) SetMicroversionId(v BTMicroversionId366)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTRootAssemblyDisplayData96) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetMicroversionIdAndConfigurationInterval

`func (o *BTRootAssemblyDisplayData96) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364`

GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field if non-nil, zero value otherwise.

### GetMicroversionIdAndConfigurationIntervalOk

`func (o *BTRootAssemblyDisplayData96) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool)`

GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionIdAndConfigurationInterval

`func (o *BTRootAssemblyDisplayData96) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364)`

SetMicroversionIdAndConfigurationInterval sets MicroversionIdAndConfigurationInterval field to given value.

### HasMicroversionIdAndConfigurationInterval

`func (o *BTRootAssemblyDisplayData96) HasMicroversionIdAndConfigurationInterval() bool`

HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.

### GetMicroversionInterval

`func (o *BTRootAssemblyDisplayData96) GetMicroversionInterval() BTMicroversionIdInterval367`

GetMicroversionInterval returns the MicroversionInterval field if non-nil, zero value otherwise.

### GetMicroversionIntervalOk

`func (o *BTRootAssemblyDisplayData96) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool)`

GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionInterval

`func (o *BTRootAssemblyDisplayData96) SetMicroversionInterval(v BTMicroversionIdInterval367)`

SetMicroversionInterval sets MicroversionInterval field to given value.

### HasMicroversionInterval

`func (o *BTRootAssemblyDisplayData96) HasMicroversionInterval() bool`

HasMicroversionInterval returns a boolean if a field has been set.

### GetOccurrences

`func (o *BTRootAssemblyDisplayData96) GetOccurrences() []BTOccurrenceDisplayData95`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *BTRootAssemblyDisplayData96) GetOccurrencesOk() (*[]BTOccurrenceDisplayData95, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *BTRootAssemblyDisplayData96) SetOccurrences(v []BTOccurrenceDisplayData95)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *BTRootAssemblyDisplayData96) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.

### GetOriginDisplayData

`func (o *BTRootAssemblyDisplayData96) GetOriginDisplayData() BTOriginDisplayData934`

GetOriginDisplayData returns the OriginDisplayData field if non-nil, zero value otherwise.

### GetOriginDisplayDataOk

`func (o *BTRootAssemblyDisplayData96) GetOriginDisplayDataOk() (*BTOriginDisplayData934, bool)`

GetOriginDisplayDataOk returns a tuple with the OriginDisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDisplayData

`func (o *BTRootAssemblyDisplayData96) SetOriginDisplayData(v BTOriginDisplayData934)`

SetOriginDisplayData sets OriginDisplayData field to given value.

### HasOriginDisplayData

`func (o *BTRootAssemblyDisplayData96) HasOriginDisplayData() bool`

HasOriginDisplayData returns a boolean if a field has been set.

### GetPartStudioDisplayData

`func (o *BTRootAssemblyDisplayData96) GetPartStudioDisplayData() []BTPartStudioDisplayDataBase2751`

GetPartStudioDisplayData returns the PartStudioDisplayData field if non-nil, zero value otherwise.

### GetPartStudioDisplayDataOk

`func (o *BTRootAssemblyDisplayData96) GetPartStudioDisplayDataOk() (*[]BTPartStudioDisplayDataBase2751, bool)`

GetPartStudioDisplayDataOk returns a tuple with the PartStudioDisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartStudioDisplayData

`func (o *BTRootAssemblyDisplayData96) SetPartStudioDisplayData(v []BTPartStudioDisplayDataBase2751)`

SetPartStudioDisplayData sets PartStudioDisplayData field to given value.

### HasPartStudioDisplayData

`func (o *BTRootAssemblyDisplayData96) HasPartStudioDisplayData() bool`

HasPartStudioDisplayData returns a boolean if a field has been set.

### GetQuickSummary

`func (o *BTRootAssemblyDisplayData96) GetQuickSummary() string`

GetQuickSummary returns the QuickSummary field if non-nil, zero value otherwise.

### GetQuickSummaryOk

`func (o *BTRootAssemblyDisplayData96) GetQuickSummaryOk() (*string, bool)`

GetQuickSummaryOk returns a tuple with the QuickSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickSummary

`func (o *BTRootAssemblyDisplayData96) SetQuickSummary(v string)`

SetQuickSummary sets QuickSummary field to given value.

### HasQuickSummary

`func (o *BTRootAssemblyDisplayData96) HasQuickSummary() bool`

HasQuickSummary returns a boolean if a field has been set.

### GetSentTimeISO

`func (o *BTRootAssemblyDisplayData96) GetSentTimeISO() string`

GetSentTimeISO returns the SentTimeISO field if non-nil, zero value otherwise.

### GetSentTimeISOOk

`func (o *BTRootAssemblyDisplayData96) GetSentTimeISOOk() (*string, bool)`

GetSentTimeISOOk returns a tuple with the SentTimeISO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTimeISO

`func (o *BTRootAssemblyDisplayData96) SetSentTimeISO(v string)`

SetSentTimeISO sets SentTimeISO field to given value.

### HasSentTimeISO

`func (o *BTRootAssemblyDisplayData96) HasSentTimeISO() bool`

HasSentTimeISO returns a boolean if a field has been set.

### GetVersionForRasterization

`func (o *BTRootAssemblyDisplayData96) GetVersionForRasterization() BTElementDisplayData326`

GetVersionForRasterization returns the VersionForRasterization field if non-nil, zero value otherwise.

### GetVersionForRasterizationOk

`func (o *BTRootAssemblyDisplayData96) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool)`

GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionForRasterization

`func (o *BTRootAssemblyDisplayData96) SetVersionForRasterization(v BTElementDisplayData326)`

SetVersionForRasterization sets VersionForRasterization field to given value.

### HasVersionForRasterization

`func (o *BTRootAssemblyDisplayData96) HasVersionForRasterization() bool`

HasVersionForRasterization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


