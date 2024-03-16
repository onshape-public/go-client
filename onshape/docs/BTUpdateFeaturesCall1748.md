# BTUpdateFeaturesCall1748

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Features** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**LibraryVersion** | Pointer to **int32** | FeatureScript version used in the Part Studio. Do not modify. | [optional] 
**MicroversionSkew** | Pointer to **bool** | On output, &#x60;true&#x60; indicates a microversion mismatch was encountered. | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** | If &#x60;true&#x60;, the call will refuse to make the addition if the current microversion for the document does not match the source microversion. If &#x60;false&#x60;, a best-effort attempt is made to re-interpret the feature addition in the context of a newer document microversion. | [optional] 
**SerializationVersion** | Pointer to **string** | Version of the structure serialization rules used to encode the output. This enables incompatibility detection during software updates. | [optional] 
**SourceMicroversion** | Pointer to **string** | The state from which the result was extracted. Geometry ID interpretation is dependent on this document microversion. | [optional] 
**UpdateSuppressionAttributes** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTUpdateFeaturesCall1748

`func NewBTUpdateFeaturesCall1748() *BTUpdateFeaturesCall1748`

NewBTUpdateFeaturesCall1748 instantiates a new BTUpdateFeaturesCall1748 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUpdateFeaturesCall1748WithDefaults

`func NewBTUpdateFeaturesCall1748WithDefaults() *BTUpdateFeaturesCall1748`

NewBTUpdateFeaturesCall1748WithDefaults instantiates a new BTUpdateFeaturesCall1748 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTUpdateFeaturesCall1748) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTUpdateFeaturesCall1748) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTUpdateFeaturesCall1748) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTUpdateFeaturesCall1748) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatures

`func (o *BTUpdateFeaturesCall1748) GetFeatures() []BTMFeature134`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTUpdateFeaturesCall1748) GetFeaturesOk() (*[]BTMFeature134, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTUpdateFeaturesCall1748) SetFeatures(v []BTMFeature134)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTUpdateFeaturesCall1748) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTUpdateFeaturesCall1748) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTUpdateFeaturesCall1748) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTUpdateFeaturesCall1748) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTUpdateFeaturesCall1748) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTUpdateFeaturesCall1748) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTUpdateFeaturesCall1748) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTUpdateFeaturesCall1748) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTUpdateFeaturesCall1748) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTUpdateFeaturesCall1748) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTUpdateFeaturesCall1748) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTUpdateFeaturesCall1748) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTUpdateFeaturesCall1748) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTUpdateFeaturesCall1748) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTUpdateFeaturesCall1748) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTUpdateFeaturesCall1748) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTUpdateFeaturesCall1748) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTUpdateFeaturesCall1748) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTUpdateFeaturesCall1748) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTUpdateFeaturesCall1748) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTUpdateFeaturesCall1748) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.

### GetUpdateSuppressionAttributes

`func (o *BTUpdateFeaturesCall1748) GetUpdateSuppressionAttributes() bool`

GetUpdateSuppressionAttributes returns the UpdateSuppressionAttributes field if non-nil, zero value otherwise.

### GetUpdateSuppressionAttributesOk

`func (o *BTUpdateFeaturesCall1748) GetUpdateSuppressionAttributesOk() (*bool, bool)`

GetUpdateSuppressionAttributesOk returns a tuple with the UpdateSuppressionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSuppressionAttributes

`func (o *BTUpdateFeaturesCall1748) SetUpdateSuppressionAttributes(v bool)`

SetUpdateSuppressionAttributes sets UpdateSuppressionAttributes field to given value.

### HasUpdateSuppressionAttributes

`func (o *BTUpdateFeaturesCall1748) HasUpdateSuppressionAttributes() bool`

HasUpdateSuppressionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


