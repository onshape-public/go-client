# BTSetFeatureRollbackResponse1042

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LibraryVersion** | Pointer to **int32** |  | [optional] 
**MicroversionId** | Pointer to [**BTMicroversionId366**](BTMicroversionId-366.md) |  | [optional] 
**MicroversionSkew** | Pointer to **bool** |  | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** |  | [optional] 
**RollbackIndex** | Pointer to **int32** |  | [optional] 
**SerializationVersion** | Pointer to **string** |  | [optional] 
**SourceMicroversion** | Pointer to **string** |  | [optional] 

## Methods

### NewBTSetFeatureRollbackResponse1042

`func NewBTSetFeatureRollbackResponse1042() *BTSetFeatureRollbackResponse1042`

NewBTSetFeatureRollbackResponse1042 instantiates a new BTSetFeatureRollbackResponse1042 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSetFeatureRollbackResponse1042WithDefaults

`func NewBTSetFeatureRollbackResponse1042WithDefaults() *BTSetFeatureRollbackResponse1042`

NewBTSetFeatureRollbackResponse1042WithDefaults instantiates a new BTSetFeatureRollbackResponse1042 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibraryVersion

`func (o *BTSetFeatureRollbackResponse1042) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTSetFeatureRollbackResponse1042) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTSetFeatureRollbackResponse1042) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTSetFeatureRollbackResponse1042) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTSetFeatureRollbackResponse1042) GetMicroversionId() BTMicroversionId366`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTSetFeatureRollbackResponse1042) GetMicroversionIdOk() (*BTMicroversionId366, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTSetFeatureRollbackResponse1042) SetMicroversionId(v BTMicroversionId366)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTSetFeatureRollbackResponse1042) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTSetFeatureRollbackResponse1042) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTSetFeatureRollbackResponse1042) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTSetFeatureRollbackResponse1042) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTSetFeatureRollbackResponse1042) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTSetFeatureRollbackResponse1042) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTSetFeatureRollbackResponse1042) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTSetFeatureRollbackResponse1042) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTSetFeatureRollbackResponse1042) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetRollbackIndex

`func (o *BTSetFeatureRollbackResponse1042) GetRollbackIndex() int32`

GetRollbackIndex returns the RollbackIndex field if non-nil, zero value otherwise.

### GetRollbackIndexOk

`func (o *BTSetFeatureRollbackResponse1042) GetRollbackIndexOk() (*int32, bool)`

GetRollbackIndexOk returns a tuple with the RollbackIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackIndex

`func (o *BTSetFeatureRollbackResponse1042) SetRollbackIndex(v int32)`

SetRollbackIndex sets RollbackIndex field to given value.

### HasRollbackIndex

`func (o *BTSetFeatureRollbackResponse1042) HasRollbackIndex() bool`

HasRollbackIndex returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTSetFeatureRollbackResponse1042) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTSetFeatureRollbackResponse1042) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTSetFeatureRollbackResponse1042) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTSetFeatureRollbackResponse1042) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTSetFeatureRollbackResponse1042) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTSetFeatureRollbackResponse1042) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTSetFeatureRollbackResponse1042) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTSetFeatureRollbackResponse1042) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


