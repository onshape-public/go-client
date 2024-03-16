# BTFeatureListResponse2457

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**DefaultFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) | List of Onshape-defined features instantiated within the Part Studio. | [optional] 
**FeatureStates** | Pointer to [**map[string]BTFeatureState1688**](BTFeatureState1688.md) | State of each feature, indicating if the feature is valid. Incorrectly defined features will still appear in the Feature list. | [optional] 
**Features** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) | List of user-defined features instantiated within the Part Studio. | [optional] 
**Imports** | Pointer to [**[]BTMImport136**](BTMImport136.md) | Internal only. Do not modify. | [optional] 
**IsComplete** | Pointer to **bool** | &#x60;true&#x60; if the features represent the entire part studio or &#x60;false&#x60; for a filtered subset. | [optional] 
**LibraryVersion** | Pointer to **int32** | FeatureScript version used in the Part Studio. Do not modify. | [optional] 
**MicroversionSkew** | Pointer to **bool** | On output, &#x60;true&#x60; indicates a microversion mismatch was encountered. | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** | If &#x60;true&#x60;, the call will refuse to make the addition if the current microversion for the document does not match the source microversion. If &#x60;false&#x60;, a best-effort attempt is made to re-interpret the feature addition in the context of a newer document microversion. | [optional] 
**RollbackIndex** | Pointer to **int32** | Index of the rollback bar location. &#x60;-1&#x60; indicates the bar is at the end of the Feature List. | [optional] 
**SerializationVersion** | Pointer to **string** | Version of the structure serialization rules used to encode the output. This enables incompatibility detection during software updates. | [optional] 
**SourceMicroversion** | Pointer to **string** | The document microversion from which the result was extracted. Part, face, edge, and vertex IDs are only valid for the same microversion. | [optional] 

## Methods

### NewBTFeatureListResponse2457

`func NewBTFeatureListResponse2457() *BTFeatureListResponse2457`

NewBTFeatureListResponse2457 instantiates a new BTFeatureListResponse2457 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTFeatureListResponse2457WithDefaults

`func NewBTFeatureListResponse2457WithDefaults() *BTFeatureListResponse2457`

NewBTFeatureListResponse2457WithDefaults instantiates a new BTFeatureListResponse2457 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTFeatureListResponse2457) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTFeatureListResponse2457) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTFeatureListResponse2457) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTFeatureListResponse2457) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDefaultFeatures

`func (o *BTFeatureListResponse2457) GetDefaultFeatures() []BTMFeature134`

GetDefaultFeatures returns the DefaultFeatures field if non-nil, zero value otherwise.

### GetDefaultFeaturesOk

`func (o *BTFeatureListResponse2457) GetDefaultFeaturesOk() (*[]BTMFeature134, bool)`

GetDefaultFeaturesOk returns a tuple with the DefaultFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFeatures

`func (o *BTFeatureListResponse2457) SetDefaultFeatures(v []BTMFeature134)`

SetDefaultFeatures sets DefaultFeatures field to given value.

### HasDefaultFeatures

`func (o *BTFeatureListResponse2457) HasDefaultFeatures() bool`

HasDefaultFeatures returns a boolean if a field has been set.

### GetFeatureStates

`func (o *BTFeatureListResponse2457) GetFeatureStates() map[string]BTFeatureState1688`

GetFeatureStates returns the FeatureStates field if non-nil, zero value otherwise.

### GetFeatureStatesOk

`func (o *BTFeatureListResponse2457) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool)`

GetFeatureStatesOk returns a tuple with the FeatureStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStates

`func (o *BTFeatureListResponse2457) SetFeatureStates(v map[string]BTFeatureState1688)`

SetFeatureStates sets FeatureStates field to given value.

### HasFeatureStates

`func (o *BTFeatureListResponse2457) HasFeatureStates() bool`

HasFeatureStates returns a boolean if a field has been set.

### GetFeatures

`func (o *BTFeatureListResponse2457) GetFeatures() []BTMFeature134`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTFeatureListResponse2457) GetFeaturesOk() (*[]BTMFeature134, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTFeatureListResponse2457) SetFeatures(v []BTMFeature134)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTFeatureListResponse2457) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetImports

`func (o *BTFeatureListResponse2457) GetImports() []BTMImport136`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *BTFeatureListResponse2457) GetImportsOk() (*[]BTMImport136, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *BTFeatureListResponse2457) SetImports(v []BTMImport136)`

SetImports sets Imports field to given value.

### HasImports

`func (o *BTFeatureListResponse2457) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetIsComplete

`func (o *BTFeatureListResponse2457) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *BTFeatureListResponse2457) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *BTFeatureListResponse2457) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *BTFeatureListResponse2457) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTFeatureListResponse2457) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTFeatureListResponse2457) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTFeatureListResponse2457) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTFeatureListResponse2457) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTFeatureListResponse2457) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTFeatureListResponse2457) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTFeatureListResponse2457) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTFeatureListResponse2457) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTFeatureListResponse2457) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTFeatureListResponse2457) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTFeatureListResponse2457) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTFeatureListResponse2457) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetRollbackIndex

`func (o *BTFeatureListResponse2457) GetRollbackIndex() int32`

GetRollbackIndex returns the RollbackIndex field if non-nil, zero value otherwise.

### GetRollbackIndexOk

`func (o *BTFeatureListResponse2457) GetRollbackIndexOk() (*int32, bool)`

GetRollbackIndexOk returns a tuple with the RollbackIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackIndex

`func (o *BTFeatureListResponse2457) SetRollbackIndex(v int32)`

SetRollbackIndex sets RollbackIndex field to given value.

### HasRollbackIndex

`func (o *BTFeatureListResponse2457) HasRollbackIndex() bool`

HasRollbackIndex returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTFeatureListResponse2457) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTFeatureListResponse2457) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTFeatureListResponse2457) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTFeatureListResponse2457) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTFeatureListResponse2457) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTFeatureListResponse2457) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTFeatureListResponse2457) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTFeatureListResponse2457) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


