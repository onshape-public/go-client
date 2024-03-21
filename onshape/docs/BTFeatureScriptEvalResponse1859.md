# BTFeatureScriptEvalResponse1859

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Console** | Pointer to **string** |  | [optional] 
**LibraryVersion** | Pointer to **int32** | FeatureScript version used in the Part Studio. Do not modify. | [optional] 
**MicroversionSkew** | Pointer to **bool** | On output, &#x60;true&#x60; indicates a microversion mismatch was encountered. | [optional] 
**Notices** | Pointer to [**[]BTNotice227**](BTNotice227.md) |  | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** | If &#x60;true&#x60;, the call will refuse to make the addition if the current microversion for the document does not match the source microversion. If &#x60;false&#x60;, a best-effort attempt is made to re-interpret the feature addition in the context of a newer document microversion. | [optional] 
**Result** | Pointer to [**BTFSValue1888**](BTFSValue1888.md) |  | [optional] 
**SerializationVersion** | Pointer to **string** | Version of the structure serialization rules used to encode the output. This enables incompatibility detection during software updates. | [optional] 
**SourceMicroversion** | Pointer to **string** | The state from which the result was extracted. Geometry ID interpretation is dependent on this document microversion. | [optional] 

## Methods

### NewBTFeatureScriptEvalResponse1859

`func NewBTFeatureScriptEvalResponse1859() *BTFeatureScriptEvalResponse1859`

NewBTFeatureScriptEvalResponse1859 instantiates a new BTFeatureScriptEvalResponse1859 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTFeatureScriptEvalResponse1859WithDefaults

`func NewBTFeatureScriptEvalResponse1859WithDefaults() *BTFeatureScriptEvalResponse1859`

NewBTFeatureScriptEvalResponse1859WithDefaults instantiates a new BTFeatureScriptEvalResponse1859 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTFeatureScriptEvalResponse1859) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTFeatureScriptEvalResponse1859) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTFeatureScriptEvalResponse1859) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTFeatureScriptEvalResponse1859) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConsole

`func (o *BTFeatureScriptEvalResponse1859) GetConsole() string`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *BTFeatureScriptEvalResponse1859) GetConsoleOk() (*string, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *BTFeatureScriptEvalResponse1859) SetConsole(v string)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *BTFeatureScriptEvalResponse1859) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTFeatureScriptEvalResponse1859) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTFeatureScriptEvalResponse1859) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTFeatureScriptEvalResponse1859) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTFeatureScriptEvalResponse1859) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTFeatureScriptEvalResponse1859) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTFeatureScriptEvalResponse1859) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTFeatureScriptEvalResponse1859) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTFeatureScriptEvalResponse1859) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetNotices

`func (o *BTFeatureScriptEvalResponse1859) GetNotices() []BTNotice227`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *BTFeatureScriptEvalResponse1859) GetNoticesOk() (*[]BTNotice227, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *BTFeatureScriptEvalResponse1859) SetNotices(v []BTNotice227)`

SetNotices sets Notices field to given value.

### HasNotices

`func (o *BTFeatureScriptEvalResponse1859) HasNotices() bool`

HasNotices returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTFeatureScriptEvalResponse1859) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTFeatureScriptEvalResponse1859) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTFeatureScriptEvalResponse1859) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTFeatureScriptEvalResponse1859) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetResult

`func (o *BTFeatureScriptEvalResponse1859) GetResult() BTFSValue1888`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BTFeatureScriptEvalResponse1859) GetResultOk() (*BTFSValue1888, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BTFeatureScriptEvalResponse1859) SetResult(v BTFSValue1888)`

SetResult sets Result field to given value.

### HasResult

`func (o *BTFeatureScriptEvalResponse1859) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTFeatureScriptEvalResponse1859) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTFeatureScriptEvalResponse1859) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTFeatureScriptEvalResponse1859) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTFeatureScriptEvalResponse1859) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTFeatureScriptEvalResponse1859) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTFeatureScriptEvalResponse1859) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTFeatureScriptEvalResponse1859) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTFeatureScriptEvalResponse1859) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


