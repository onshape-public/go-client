# BTConfigurationResponse2019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationParameters** | Pointer to [**[]BTMConfigurationParameter819**](BTMConfigurationParameter819.md) |  | [optional] 
**CurrentConfiguration** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**LibraryVersion** | Pointer to **int32** |  | [optional] 
**MicroversionSkew** | Pointer to **bool** |  | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** |  | [optional] 
**SerializationVersion** | Pointer to **string** |  | [optional] 
**SourceMicroversion** | Pointer to **string** |  | [optional] 

## Methods

### NewBTConfigurationResponse2019

`func NewBTConfigurationResponse2019() *BTConfigurationResponse2019`

NewBTConfigurationResponse2019 instantiates a new BTConfigurationResponse2019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTConfigurationResponse2019WithDefaults

`func NewBTConfigurationResponse2019WithDefaults() *BTConfigurationResponse2019`

NewBTConfigurationResponse2019WithDefaults instantiates a new BTConfigurationResponse2019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationParameters

`func (o *BTConfigurationResponse2019) GetConfigurationParameters() []BTMConfigurationParameter819`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *BTConfigurationResponse2019) GetConfigurationParametersOk() (*[]BTMConfigurationParameter819, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *BTConfigurationResponse2019) SetConfigurationParameters(v []BTMConfigurationParameter819)`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *BTConfigurationResponse2019) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### GetCurrentConfiguration

`func (o *BTConfigurationResponse2019) GetCurrentConfiguration() []BTMParameter1`

GetCurrentConfiguration returns the CurrentConfiguration field if non-nil, zero value otherwise.

### GetCurrentConfigurationOk

`func (o *BTConfigurationResponse2019) GetCurrentConfigurationOk() (*[]BTMParameter1, bool)`

GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfiguration

`func (o *BTConfigurationResponse2019) SetCurrentConfiguration(v []BTMParameter1)`

SetCurrentConfiguration sets CurrentConfiguration field to given value.

### HasCurrentConfiguration

`func (o *BTConfigurationResponse2019) HasCurrentConfiguration() bool`

HasCurrentConfiguration returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTConfigurationResponse2019) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTConfigurationResponse2019) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTConfigurationResponse2019) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTConfigurationResponse2019) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTConfigurationResponse2019) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTConfigurationResponse2019) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTConfigurationResponse2019) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTConfigurationResponse2019) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTConfigurationResponse2019) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTConfigurationResponse2019) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTConfigurationResponse2019) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTConfigurationResponse2019) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTConfigurationResponse2019) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTConfigurationResponse2019) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTConfigurationResponse2019) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTConfigurationResponse2019) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTConfigurationResponse2019) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTConfigurationResponse2019) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTConfigurationResponse2019) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTConfigurationResponse2019) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


