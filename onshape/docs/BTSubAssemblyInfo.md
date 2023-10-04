# BTSubAssemblyInfo

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
**PartNumber** | Pointer to **string** |  | [optional] 
**Patterns** | Pointer to [**[]BTAssemblyPatternInfo**](BTAssemblyPatternInfo.md) | List of patterns. | [optional] 
**Revision** | Pointer to **string** |  | [optional] 

## Methods

### NewBTSubAssemblyInfo

`func NewBTSubAssemblyInfo() *BTSubAssemblyInfo`

NewBTSubAssemblyInfo instantiates a new BTSubAssemblyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSubAssemblyInfoWithDefaults

`func NewBTSubAssemblyInfoWithDefaults() *BTSubAssemblyInfo`

NewBTSubAssemblyInfoWithDefaults instantiates a new BTSubAssemblyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTSubAssemblyInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTSubAssemblyInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTSubAssemblyInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTSubAssemblyInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTSubAssemblyInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTSubAssemblyInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTSubAssemblyInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTSubAssemblyInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentMicroversion

`func (o *BTSubAssemblyInfo) GetDocumentMicroversion() string`

GetDocumentMicroversion returns the DocumentMicroversion field if non-nil, zero value otherwise.

### GetDocumentMicroversionOk

`func (o *BTSubAssemblyInfo) GetDocumentMicroversionOk() (*string, bool)`

GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversion

`func (o *BTSubAssemblyInfo) SetDocumentMicroversion(v string)`

SetDocumentMicroversion sets DocumentMicroversion field to given value.

### HasDocumentMicroversion

`func (o *BTSubAssemblyInfo) HasDocumentMicroversion() bool`

HasDocumentMicroversion returns a boolean if a field has been set.

### GetDocumentVersion

`func (o *BTSubAssemblyInfo) GetDocumentVersion() string`

GetDocumentVersion returns the DocumentVersion field if non-nil, zero value otherwise.

### GetDocumentVersionOk

`func (o *BTSubAssemblyInfo) GetDocumentVersionOk() (*string, bool)`

GetDocumentVersionOk returns a tuple with the DocumentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersion

`func (o *BTSubAssemblyInfo) SetDocumentVersion(v string)`

SetDocumentVersion sets DocumentVersion field to given value.

### HasDocumentVersion

`func (o *BTSubAssemblyInfo) HasDocumentVersion() bool`

HasDocumentVersion returns a boolean if a field has been set.

### GetElementId

`func (o *BTSubAssemblyInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTSubAssemblyInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTSubAssemblyInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTSubAssemblyInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFeatures

`func (o *BTSubAssemblyInfo) GetFeatures() []BTAssemblyFeatureInfo`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTSubAssemblyInfo) GetFeaturesOk() (*[]BTAssemblyFeatureInfo, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTSubAssemblyInfo) SetFeatures(v []BTAssemblyFeatureInfo)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTSubAssemblyInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFullConfiguration

`func (o *BTSubAssemblyInfo) GetFullConfiguration() string`

GetFullConfiguration returns the FullConfiguration field if non-nil, zero value otherwise.

### GetFullConfigurationOk

`func (o *BTSubAssemblyInfo) GetFullConfigurationOk() (*string, bool)`

GetFullConfigurationOk returns a tuple with the FullConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullConfiguration

`func (o *BTSubAssemblyInfo) SetFullConfiguration(v string)`

SetFullConfiguration sets FullConfiguration field to given value.

### HasFullConfiguration

`func (o *BTSubAssemblyInfo) HasFullConfiguration() bool`

HasFullConfiguration returns a boolean if a field has been set.

### GetInstances

`func (o *BTSubAssemblyInfo) GetInstances() []BTAssemblyInstanceInfo`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BTSubAssemblyInfo) GetInstancesOk() (*[]BTAssemblyInstanceInfo, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BTSubAssemblyInfo) SetInstances(v []BTAssemblyInstanceInfo)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BTSubAssemblyInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTSubAssemblyInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTSubAssemblyInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTSubAssemblyInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTSubAssemblyInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPatterns

`func (o *BTSubAssemblyInfo) GetPatterns() []BTAssemblyPatternInfo`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *BTSubAssemblyInfo) GetPatternsOk() (*[]BTAssemblyPatternInfo, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *BTSubAssemblyInfo) SetPatterns(v []BTAssemblyPatternInfo)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *BTSubAssemblyInfo) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetRevision

`func (o *BTSubAssemblyInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTSubAssemblyInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTSubAssemblyInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTSubAssemblyInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


