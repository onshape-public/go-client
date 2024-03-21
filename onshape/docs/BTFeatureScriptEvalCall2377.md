# BTFeatureScriptEvalCall2377

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**LibraryVersion** | Pointer to **int32** | FeatureScript version used in the Part Studio. Do not modify. | [optional] 
**MicroversionSkew** | Pointer to **bool** | On output, &#x60;true&#x60; indicates a microversion mismatch was encountered. | [optional] 
**Queries** | Pointer to **map[string][]string** |  | [optional] 
**RejectMicroversionSkew** | Pointer to **bool** | If &#x60;true&#x60;, the call will refuse to make the addition if the current microversion for the document does not match the source microversion. If &#x60;false&#x60;, a best-effort attempt is made to re-interpret the feature addition in the context of a newer document microversion. | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**SerializationVersion** | Pointer to **string** | Version of the structure serialization rules used to encode the output. This enables incompatibility detection during software updates. | [optional] 
**SourceMicroversion** | Pointer to **string** | The state from which the result was extracted. Geometry ID interpretation is dependent on this document microversion. | [optional] 

## Methods

### NewBTFeatureScriptEvalCall2377

`func NewBTFeatureScriptEvalCall2377() *BTFeatureScriptEvalCall2377`

NewBTFeatureScriptEvalCall2377 instantiates a new BTFeatureScriptEvalCall2377 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTFeatureScriptEvalCall2377WithDefaults

`func NewBTFeatureScriptEvalCall2377WithDefaults() *BTFeatureScriptEvalCall2377`

NewBTFeatureScriptEvalCall2377WithDefaults instantiates a new BTFeatureScriptEvalCall2377 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTFeatureScriptEvalCall2377) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTFeatureScriptEvalCall2377) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTFeatureScriptEvalCall2377) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTFeatureScriptEvalCall2377) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTFeatureScriptEvalCall2377) GetLibraryVersion() int32`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTFeatureScriptEvalCall2377) GetLibraryVersionOk() (*int32, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTFeatureScriptEvalCall2377) SetLibraryVersion(v int32)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTFeatureScriptEvalCall2377) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetMicroversionSkew

`func (o *BTFeatureScriptEvalCall2377) GetMicroversionSkew() bool`

GetMicroversionSkew returns the MicroversionSkew field if non-nil, zero value otherwise.

### GetMicroversionSkewOk

`func (o *BTFeatureScriptEvalCall2377) GetMicroversionSkewOk() (*bool, bool)`

GetMicroversionSkewOk returns a tuple with the MicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionSkew

`func (o *BTFeatureScriptEvalCall2377) SetMicroversionSkew(v bool)`

SetMicroversionSkew sets MicroversionSkew field to given value.

### HasMicroversionSkew

`func (o *BTFeatureScriptEvalCall2377) HasMicroversionSkew() bool`

HasMicroversionSkew returns a boolean if a field has been set.

### GetQueries

`func (o *BTFeatureScriptEvalCall2377) GetQueries() map[string][]string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *BTFeatureScriptEvalCall2377) GetQueriesOk() (*map[string][]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *BTFeatureScriptEvalCall2377) SetQueries(v map[string][]string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *BTFeatureScriptEvalCall2377) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetRejectMicroversionSkew

`func (o *BTFeatureScriptEvalCall2377) GetRejectMicroversionSkew() bool`

GetRejectMicroversionSkew returns the RejectMicroversionSkew field if non-nil, zero value otherwise.

### GetRejectMicroversionSkewOk

`func (o *BTFeatureScriptEvalCall2377) GetRejectMicroversionSkewOk() (*bool, bool)`

GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMicroversionSkew

`func (o *BTFeatureScriptEvalCall2377) SetRejectMicroversionSkew(v bool)`

SetRejectMicroversionSkew sets RejectMicroversionSkew field to given value.

### HasRejectMicroversionSkew

`func (o *BTFeatureScriptEvalCall2377) HasRejectMicroversionSkew() bool`

HasRejectMicroversionSkew returns a boolean if a field has been set.

### GetScript

`func (o *BTFeatureScriptEvalCall2377) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *BTFeatureScriptEvalCall2377) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *BTFeatureScriptEvalCall2377) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *BTFeatureScriptEvalCall2377) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetSerializationVersion

`func (o *BTFeatureScriptEvalCall2377) GetSerializationVersion() string`

GetSerializationVersion returns the SerializationVersion field if non-nil, zero value otherwise.

### GetSerializationVersionOk

`func (o *BTFeatureScriptEvalCall2377) GetSerializationVersionOk() (*string, bool)`

GetSerializationVersionOk returns a tuple with the SerializationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationVersion

`func (o *BTFeatureScriptEvalCall2377) SetSerializationVersion(v string)`

SetSerializationVersion sets SerializationVersion field to given value.

### HasSerializationVersion

`func (o *BTFeatureScriptEvalCall2377) HasSerializationVersion() bool`

HasSerializationVersion returns a boolean if a field has been set.

### GetSourceMicroversion

`func (o *BTFeatureScriptEvalCall2377) GetSourceMicroversion() string`

GetSourceMicroversion returns the SourceMicroversion field if non-nil, zero value otherwise.

### GetSourceMicroversionOk

`func (o *BTFeatureScriptEvalCall2377) GetSourceMicroversionOk() (*string, bool)`

GetSourceMicroversionOk returns a tuple with the SourceMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversion

`func (o *BTFeatureScriptEvalCall2377) SetSourceMicroversion(v string)`

SetSourceMicroversion sets SourceMicroversion field to given value.

### HasSourceMicroversion

`func (o *BTFeatureScriptEvalCall2377) HasSourceMicroversion() bool`

HasSourceMicroversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


