# BTRootAssemblyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentMicroversion** | Pointer to **string** |  | [optional] 
**DocumentVersion** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]BTAssemblyFeatureInfo**](BTAssemblyFeatureInfo.md) | List of Assembly features including those are created by replicates. | [optional] 
**FullConfiguration** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]BTAssemblyInstanceInfo**](BTAssemblyInstanceInfo.md) | List of instances including those created by patterns and replicates. | [optional] 
**Occurrences** | Pointer to [**[]BTAssemblyOccurrenceInfo**](BTAssemblyOccurrenceInfo.md) |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Patterns** | Pointer to [**[]BTAssemblyPatternInfo**](BTAssemblyPatternInfo.md) | List of patterns. | [optional] 
**Revision** | Pointer to **string** |  | [optional] 

## Methods

### NewBTRootAssemblyInfo

`func NewBTRootAssemblyInfo() *BTRootAssemblyInfo`

NewBTRootAssemblyInfo instantiates a new BTRootAssemblyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRootAssemblyInfoWithDefaults

`func NewBTRootAssemblyInfoWithDefaults() *BTRootAssemblyInfo`

NewBTRootAssemblyInfoWithDefaults instantiates a new BTRootAssemblyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTRootAssemblyInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTRootAssemblyInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTRootAssemblyInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTRootAssemblyInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTRootAssemblyInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTRootAssemblyInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTRootAssemblyInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTRootAssemblyInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentMicroversion

`func (o *BTRootAssemblyInfo) GetDocumentMicroversion() string`

GetDocumentMicroversion returns the DocumentMicroversion field if non-nil, zero value otherwise.

### GetDocumentMicroversionOk

`func (o *BTRootAssemblyInfo) GetDocumentMicroversionOk() (*string, bool)`

GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversion

`func (o *BTRootAssemblyInfo) SetDocumentMicroversion(v string)`

SetDocumentMicroversion sets DocumentMicroversion field to given value.

### HasDocumentMicroversion

`func (o *BTRootAssemblyInfo) HasDocumentMicroversion() bool`

HasDocumentMicroversion returns a boolean if a field has been set.

### GetDocumentVersion

`func (o *BTRootAssemblyInfo) GetDocumentVersion() string`

GetDocumentVersion returns the DocumentVersion field if non-nil, zero value otherwise.

### GetDocumentVersionOk

`func (o *BTRootAssemblyInfo) GetDocumentVersionOk() (*string, bool)`

GetDocumentVersionOk returns a tuple with the DocumentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersion

`func (o *BTRootAssemblyInfo) SetDocumentVersion(v string)`

SetDocumentVersion sets DocumentVersion field to given value.

### HasDocumentVersion

`func (o *BTRootAssemblyInfo) HasDocumentVersion() bool`

HasDocumentVersion returns a boolean if a field has been set.

### GetElementId

`func (o *BTRootAssemblyInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTRootAssemblyInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTRootAssemblyInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTRootAssemblyInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFeatures

`func (o *BTRootAssemblyInfo) GetFeatures() []BTAssemblyFeatureInfo`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTRootAssemblyInfo) GetFeaturesOk() (*[]BTAssemblyFeatureInfo, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTRootAssemblyInfo) SetFeatures(v []BTAssemblyFeatureInfo)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTRootAssemblyInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFullConfiguration

`func (o *BTRootAssemblyInfo) GetFullConfiguration() string`

GetFullConfiguration returns the FullConfiguration field if non-nil, zero value otherwise.

### GetFullConfigurationOk

`func (o *BTRootAssemblyInfo) GetFullConfigurationOk() (*string, bool)`

GetFullConfigurationOk returns a tuple with the FullConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullConfiguration

`func (o *BTRootAssemblyInfo) SetFullConfiguration(v string)`

SetFullConfiguration sets FullConfiguration field to given value.

### HasFullConfiguration

`func (o *BTRootAssemblyInfo) HasFullConfiguration() bool`

HasFullConfiguration returns a boolean if a field has been set.

### GetInstances

`func (o *BTRootAssemblyInfo) GetInstances() []BTAssemblyInstanceInfo`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BTRootAssemblyInfo) GetInstancesOk() (*[]BTAssemblyInstanceInfo, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BTRootAssemblyInfo) SetInstances(v []BTAssemblyInstanceInfo)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BTRootAssemblyInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetOccurrences

`func (o *BTRootAssemblyInfo) GetOccurrences() []BTAssemblyOccurrenceInfo`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *BTRootAssemblyInfo) GetOccurrencesOk() (*[]BTAssemblyOccurrenceInfo, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *BTRootAssemblyInfo) SetOccurrences(v []BTAssemblyOccurrenceInfo)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *BTRootAssemblyInfo) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTRootAssemblyInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTRootAssemblyInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTRootAssemblyInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTRootAssemblyInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPatterns

`func (o *BTRootAssemblyInfo) GetPatterns() []BTAssemblyPatternInfo`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *BTRootAssemblyInfo) GetPatternsOk() (*[]BTAssemblyPatternInfo, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *BTRootAssemblyInfo) SetPatterns(v []BTAssemblyPatternInfo)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *BTRootAssemblyInfo) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetRevision

`func (o *BTRootAssemblyInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTRootAssemblyInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTRootAssemblyInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTRootAssemblyInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


