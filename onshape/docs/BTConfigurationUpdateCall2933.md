# BTConfigurationUpdateCall2933

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ConfigurationParameters** | Pointer to [**[]BTMConfigurationParameter819**](BTMConfigurationParameter819.md) |  | [optional] 
**CurrentConfiguration** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 
**LibraryVersion** | Pointer to **int32** | FeatureScript version used in the Part Studio. Do not modify. | [optional] 
**MicroversionSkew** | Pointer to **bool** | On output, &#x60;true&#x60; indicates a microversion mismatch was encountered. | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** | If &#x60;true&#x60;, the call will refuse to make the addition if the current microversion for the document does not match the source microversion. If &#x60;false&#x60;, a best-effort attempt is made to re-interpret the feature addition in the context of a newer document microversion. | [optional] 
**SerializationVersion** | Pointer to **string** | Version of the structure serialization rules used to encode the output. This enables incompatibility detection during software updates. | [optional] 
**SourceMicroversion** | Pointer to **string** | The state from which the result was extracted. Geometry ID interpretation is dependent on this document microversion. | [optional] 

## Methods

### NewBTConfigurationUpdateCall2933

`func NewBTConfigurationUpdateCall2933() *BTConfigurationUpdateCall2933`

NewBTConfigurationUpdateCall2933 instantiates a new BTConfigurationUpdateCall2933 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTConfigurationUpdateCall2933WithDefaults

`func NewBTConfigurationUpdateCall2933WithDefaults() *BTConfigurationUpdateCall2933`

NewBTConfigurationUpdateCall2933WithDefaults instantiates a new BTConfigurationUpdateCall2933 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTConfigurationUpdateCall2933) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTConfigurationUpdateCall2933) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTConfigurationUpdateCall2933) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTConfigurationUpdateCall2933) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConfigurationParameters

`func (o *BTConfigurationUpdateCall2933) GetConfigurationParameters() []BTMConfigurationParameter819`

GetConfigurationParameters returns the ConfigurationParameters field if non-nil, zero value otherwise.

### GetConfigurationParametersOk

`func (o *BTConfigurationUpdateCall2933) GetConfigurationParametersOk() (*[]BTMConfigurationParameter819, bool)`

GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameters

`func (o *BTConfigurationUpdateCall2933) SetConfigurationParameters(v []BTMConfigurationParameter819)`

SetConfigurationParameters sets ConfigurationParameters field to given value.

### HasConfigurationParameters

`func (o *BTConfigurationUpdateCall2933) HasConfigurationParameters() bool`

HasConfigurationParameters returns a boolean if a field has been set.

### GetCurrentConfiguration

`func (o *BTConfigurationUpdateCall2933) GetCurrentConfiguration() []BTMParameter1`

GetCurrentConfiguration returns the CurrentConfiguration field if non-nil, zero value otherwise.

### GetCurrentConfigurationOk

`func (o *BTConfigurationUpdateCall2933) GetCurrentConfigurationOk() (*[]BTMParameter1, bool)`

GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfiguration

`func (o *BTConfigurationUpdateCall2933) SetCurrentConfiguration(v []BTMParameter1)`

SetCurrentConfiguration sets CurrentConfiguration field to given value.

### HasCurrentConfiguration

`func (o *BTConfigurationUpdateCall2933) HasCurrentConfiguration() bool`

HasCurrentConfiguration returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTConfigurationUpdateCall2933) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTConfigurationUpdateCall2933) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTConfigurationUpdateCall2933) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTConfigurationUpdateCall2933) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTConfigurationUpdateCall2933) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTConfigurationUpdateCall2933) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTConfigurationUpdateCall2933) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTConfigurationUpdateCall2933) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTConfigurationUpdateCall2933) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTConfigurationUpdateCall2933) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTConfigurationUpdateCall2933) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTConfigurationUpdateCall2933) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTConfigurationUpdateCall2933) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTConfigurationUpdateCall2933) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTConfigurationUpdateCall2933) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTConfigurationUpdateCall2933) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTConfigurationUpdateCall2933) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTConfigurationUpdateCall2933) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTConfigurationUpdateCall2933) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTConfigurationUpdateCall2933) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


