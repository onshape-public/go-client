# BTFeatureDefinitionResponse1617

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | Pointer to [**BTMFeature134**](BTMFeature134.md) |  | [optional] 
**FeatureState** | Pointer to [**BTFeatureState1688**](BTFeatureState1688.md) |  | [optional] 
**LibraryVersion** | Pointer to **int32** |  | [optional] 
**MicroversionSkew** | Pointer to **bool** |  | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** |  | [optional] 
**SerializationVersion** | Pointer to **string** |  | [optional] 
**SourceMicroversion** | Pointer to **string** |  | [optional] 

## Methods

### NewBTFeatureDefinitionResponse1617

`func NewBTFeatureDefinitionResponse1617() *BTFeatureDefinitionResponse1617`

NewBTFeatureDefinitionResponse1617 instantiates a new BTFeatureDefinitionResponse1617 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTFeatureDefinitionResponse1617WithDefaults

`func NewBTFeatureDefinitionResponse1617WithDefaults() *BTFeatureDefinitionResponse1617`

NewBTFeatureDefinitionResponse1617WithDefaults instantiates a new BTFeatureDefinitionResponse1617 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *BTFeatureDefinitionResponse1617) GetFeature() BTMFeature134`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *BTFeatureDefinitionResponse1617) GetFeatureOk() (*BTMFeature134, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *BTFeatureDefinitionResponse1617) SetFeature(v BTMFeature134)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *BTFeatureDefinitionResponse1617) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureState

`func (o *BTFeatureDefinitionResponse1617) GetFeatureState() BTFeatureState1688`

GetFeatureState returns the FeatureState field if non-nil, zero value otherwise.

### GetFeatureStateOk

`func (o *BTFeatureDefinitionResponse1617) GetFeatureStateOk() (*BTFeatureState1688, bool)`

GetFeatureStateOk returns a tuple with the FeatureState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureState

`func (o *BTFeatureDefinitionResponse1617) SetFeatureState(v BTFeatureState1688)`

SetFeatureState sets FeatureState field to given value.

### HasFeatureState

`func (o *BTFeatureDefinitionResponse1617) HasFeatureState() bool`

HasFeatureState returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTFeatureDefinitionResponse1617) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTFeatureDefinitionResponse1617) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTFeatureDefinitionResponse1617) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTFeatureDefinitionResponse1617) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTFeatureDefinitionResponse1617) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTFeatureDefinitionResponse1617) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTFeatureDefinitionResponse1617) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTFeatureDefinitionResponse1617) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTFeatureDefinitionResponse1617) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTFeatureDefinitionResponse1617) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTFeatureDefinitionResponse1617) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTFeatureDefinitionResponse1617) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTFeatureDefinitionResponse1617) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTFeatureDefinitionResponse1617) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTFeatureDefinitionResponse1617) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTFeatureDefinitionResponse1617) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTFeatureDefinitionResponse1617) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTFeatureDefinitionResponse1617) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTFeatureDefinitionResponse1617) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTFeatureDefinitionResponse1617) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


