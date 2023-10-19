# BTUpdateFeaturesResponse1333

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**FeatureStates** | Pointer to [**map[string]BTFeatureState1688**](BTFeatureState1688.md) |  | [optional] 
**Features** | Pointer to [**[]BTMFeature134**](BTMFeature134.md) |  | [optional] 
**LibraryVersion** | Pointer to **int32** |  | [optional] 
**MicroversionSkew** | Pointer to **bool** |  | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** |  | [optional] 
**SerializationVersion** | Pointer to **string** |  | [optional] 
**SourceMicroversion** | Pointer to **string** |  | [optional] 

## Methods

### NewBTUpdateFeaturesResponse1333

`func NewBTUpdateFeaturesResponse1333() *BTUpdateFeaturesResponse1333`

NewBTUpdateFeaturesResponse1333 instantiates a new BTUpdateFeaturesResponse1333 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUpdateFeaturesResponse1333WithDefaults

`func NewBTUpdateFeaturesResponse1333WithDefaults() *BTUpdateFeaturesResponse1333`

NewBTUpdateFeaturesResponse1333WithDefaults instantiates a new BTUpdateFeaturesResponse1333 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTUpdateFeaturesResponse1333) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTUpdateFeaturesResponse1333) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTUpdateFeaturesResponse1333) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTUpdateFeaturesResponse1333) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatureStates

`func (o *BTUpdateFeaturesResponse1333) GetFeatureStates() map[string]BTFeatureState1688`

GetFeatureStates returns the FeatureStates field if non-nil, zero value otherwise.

### GetFeatureStatesOk

`func (o *BTUpdateFeaturesResponse1333) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool)`

GetFeatureStatesOk returns a tuple with the FeatureStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStates

`func (o *BTUpdateFeaturesResponse1333) SetFeatureStates(v map[string]BTFeatureState1688)`

SetFeatureStates sets FeatureStates field to given value.

### HasFeatureStates

`func (o *BTUpdateFeaturesResponse1333) HasFeatureStates() bool`

HasFeatureStates returns a boolean if a field has been set.

### GetFeatures

`func (o *BTUpdateFeaturesResponse1333) GetFeatures() []BTMFeature134`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTUpdateFeaturesResponse1333) GetFeaturesOk() (*[]BTMFeature134, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTUpdateFeaturesResponse1333) SetFeatures(v []BTMFeature134)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTUpdateFeaturesResponse1333) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTUpdateFeaturesResponse1333) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTUpdateFeaturesResponse1333) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTUpdateFeaturesResponse1333) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTUpdateFeaturesResponse1333) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTUpdateFeaturesResponse1333) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTUpdateFeaturesResponse1333) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTUpdateFeaturesResponse1333) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTUpdateFeaturesResponse1333) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTUpdateFeaturesResponse1333) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTUpdateFeaturesResponse1333) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTUpdateFeaturesResponse1333) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTUpdateFeaturesResponse1333) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTUpdateFeaturesResponse1333) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTUpdateFeaturesResponse1333) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTUpdateFeaturesResponse1333) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTUpdateFeaturesResponse1333) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTUpdateFeaturesResponse1333) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTUpdateFeaturesResponse1333) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTUpdateFeaturesResponse1333) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTUpdateFeaturesResponse1333) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


