# BTMConfigurationData1560AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**ConfigurationParameters** | Pointer to [**[]BTMConfigurationParameter819**](BTMConfigurationParameter819.md) |  | [optional] 
**CosmeticParameterIds** | Pointer to **[]string** |  | [optional] 
**CurrentConfiguration** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**CurrentFSValues** | Pointer to [**map[string]BTFSValue1888**](BTFSValue1888.md) |  | [optional] 
**DefaultConfigurationValues** | Pointer to [**map[string]BTFSValue1888**](BTFSValue1888.md) |  | [optional] 
**SyncAndPassthroughReferenceNodeId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMConfigurationData1560AllOf

`func NewBTMConfigurationData1560AllOf() *BTMConfigurationData1560AllOf`

NewBTMConfigurationData1560AllOf instantiates a new BTMConfigurationData1560AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMConfigurationData1560AllOfWithDefaults

`func NewBTMConfigurationData1560AllOfWithDefaults() *BTMConfigurationData1560AllOf`

NewBTMConfigurationData1560AllOfWithDefaults instantiates a new BTMConfigurationData1560AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMConfigurationData1560AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMConfigurationData1560AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMConfigurationData1560AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMConfigurationData1560AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *BTMConfigurationData1560AllOf) GetConfigurationParameters() []BTMConfigurationParameter819`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *BTMConfigurationData1560AllOf) GetConfigurationParametersOk() (*[]BTMConfigurationParameter819, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *BTMConfigurationData1560AllOf) SetConfigurationParameters(v []BTMConfigurationParameter819)`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *BTMConfigurationData1560AllOf) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### GetCosmeticParameterIds

`func (o *BTMConfigurationData1560AllOf) GetCosmeticParameterIds() []string`

GetCosmeticParameterIds returns the CosmeticParameterIds field if non-nil, zero value otherwise.

### GetCosmeticParameterIdsOk

`func (o *BTMConfigurationData1560AllOf) GetCosmeticParameterIdsOk() (*[]string, bool)`

GetCosmeticParameterIdsOk returns a tuple with the CosmeticParameterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmeticParameterIds

`func (o *BTMConfigurationData1560AllOf) SetCosmeticParameterIds(v []string)`

SetCosmeticParameterIds sets CosmeticParameterIds field to given value.

### HasCosmeticParameterIds

`func (o *BTMConfigurationData1560AllOf) HasCosmeticParameterIds() bool`

HasCosmeticParameterIds returns a boolean if a field has been set.

### GetCurrentConfiguration

`func (o *BTMConfigurationData1560AllOf) GetCurrentConfiguration() []BTMParameter1`

GetCurrentConfiguration returns the CurrentConfiguration field if non-nil, zero value otherwise.

### GetCurrentConfigurationOk

`func (o *BTMConfigurationData1560AllOf) GetCurrentConfigurationOk() (*[]BTMParameter1, bool)`

GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfiguration

`func (o *BTMConfigurationData1560AllOf) SetCurrentConfiguration(v []BTMParameter1)`

SetCurrentConfiguration sets CurrentConfiguration field to given value.

### HasCurrentConfiguration

`func (o *BTMConfigurationData1560AllOf) HasCurrentConfiguration() bool`

HasCurrentConfiguration returns a boolean if a field has been set.

### GetCurrentFSValues

`func (o *BTMConfigurationData1560AllOf) GetCurrentFSValues() map[string]BTFSValue1888`

GetCurrentFSValues returns the CurrentFSValues field if non-nil, zero value otherwise.

### GetCurrentFSValuesOk

`func (o *BTMConfigurationData1560AllOf) GetCurrentFSValuesOk() (*map[string]BTFSValue1888, bool)`

GetCurrentFSValuesOk returns a tuple with the CurrentFSValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFSValues

`func (o *BTMConfigurationData1560AllOf) SetCurrentFSValues(v map[string]BTFSValue1888)`

SetCurrentFSValues sets CurrentFSValues field to given value.

### HasCurrentFSValues

`func (o *BTMConfigurationData1560AllOf) HasCurrentFSValues() bool`

HasCurrentFSValues returns a boolean if a field has been set.

### GetDefaultConfigurationValues

`func (o *BTMConfigurationData1560AllOf) GetDefaultConfigurationValues() map[string]BTFSValue1888`

GetDefaultConfigurationValues returns the DefaultConfigurationValues field if non-nil, zero value otherwise.

### GetDefaultConfigurationValuesOk

`func (o *BTMConfigurationData1560AllOf) GetDefaultConfigurationValuesOk() (*map[string]BTFSValue1888, bool)`

GetDefaultConfigurationValuesOk returns a tuple with the DefaultConfigurationValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfigurationValues

`func (o *BTMConfigurationData1560AllOf) SetDefaultConfigurationValues(v map[string]BTFSValue1888)`

SetDefaultConfigurationValues sets DefaultConfigurationValues field to given value.

### HasDefaultConfigurationValues

`func (o *BTMConfigurationData1560AllOf) HasDefaultConfigurationValues() bool`

HasDefaultConfigurationValues returns a boolean if a field has been set.

### GetSyncAndPassthroughReferenceNodeId

`func (o *BTMConfigurationData1560AllOf) GetSyncAndPassthroughReferenceNodeId() string`

GetSyncAndPassthroughReferenceNodeId returns the SyncAndPassthroughReferenceNodeId field if non-nil, zero value otherwise.

### GetSyncAndPassthroughReferenceNodeIdOk

`func (o *BTMConfigurationData1560AllOf) GetSyncAndPassthroughReferenceNodeIdOk() (*string, bool)`

GetSyncAndPassthroughReferenceNodeIdOk returns a tuple with the SyncAndPassthroughReferenceNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAndPassthroughReferenceNodeId

`func (o *BTMConfigurationData1560AllOf) SetSyncAndPassthroughReferenceNodeId(v string)`

SetSyncAndPassthroughReferenceNodeId sets SyncAndPassthroughReferenceNodeId field to given value.

### HasSyncAndPassthroughReferenceNodeId

`func (o *BTMConfigurationData1560AllOf) HasSyncAndPassthroughReferenceNodeId() bool`

HasSyncAndPassthroughReferenceNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


