# BTAssociativeDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociativeDataId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]BTNameValuePair**](BTNameValuePair.md) |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentMicroversion** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**IdTag** | Pointer to **string** |  | [optional] 
**MicroversionId** | Pointer to **string** |  | [optional] 
**OccurrenceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GBTAppElementAssociativeDataType**](GBTAppElementAssociativeDataType.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAssociativeDataInfo

`func NewBTAssociativeDataInfo() *BTAssociativeDataInfo`

NewBTAssociativeDataInfo instantiates a new BTAssociativeDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssociativeDataInfoWithDefaults

`func NewBTAssociativeDataInfoWithDefaults() *BTAssociativeDataInfo`

NewBTAssociativeDataInfoWithDefaults instantiates a new BTAssociativeDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociativeDataId

`func (o *BTAssociativeDataInfo) GetAssociativeDataId() string`

GetAssociativeDataId returns the AssociativeDataId field if non-nil, zero value otherwise.

### GetAssociativeDataIdOk

`func (o *BTAssociativeDataInfo) GetAssociativeDataIdOk() (*string, bool)`

GetAssociativeDataIdOk returns a tuple with the AssociativeDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociativeDataId

`func (o *BTAssociativeDataInfo) SetAssociativeDataId(v string)`

SetAssociativeDataId sets AssociativeDataId field to given value.

### HasAssociativeDataId

`func (o *BTAssociativeDataInfo) HasAssociativeDataId() bool`

HasAssociativeDataId returns a boolean if a field has been set.

### GetData

`func (o *BTAssociativeDataInfo) GetData() []BTNameValuePair`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTAssociativeDataInfo) GetDataOk() (*[]BTNameValuePair, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTAssociativeDataInfo) SetData(v []BTNameValuePair)`

SetData sets Data field to given value.

### HasData

`func (o *BTAssociativeDataInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTAssociativeDataInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTAssociativeDataInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTAssociativeDataInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTAssociativeDataInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentMicroversion

`func (o *BTAssociativeDataInfo) GetDocumentMicroversion() string`

GetDocumentMicroversion returns the DocumentMicroversion field if non-nil, zero value otherwise.

### GetDocumentMicroversionOk

`func (o *BTAssociativeDataInfo) GetDocumentMicroversionOk() (*string, bool)`

GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversion

`func (o *BTAssociativeDataInfo) SetDocumentMicroversion(v string)`

SetDocumentMicroversion sets DocumentMicroversion field to given value.

### HasDocumentMicroversion

`func (o *BTAssociativeDataInfo) HasDocumentMicroversion() bool`

HasDocumentMicroversion returns a boolean if a field has been set.

### GetElementId

`func (o *BTAssociativeDataInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTAssociativeDataInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTAssociativeDataInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTAssociativeDataInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetError

`func (o *BTAssociativeDataInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BTAssociativeDataInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BTAssociativeDataInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BTAssociativeDataInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdTag

`func (o *BTAssociativeDataInfo) GetIdTag() string`

GetIdTag returns the IdTag field if non-nil, zero value otherwise.

### GetIdTagOk

`func (o *BTAssociativeDataInfo) GetIdTagOk() (*string, bool)`

GetIdTagOk returns a tuple with the IdTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTag

`func (o *BTAssociativeDataInfo) SetIdTag(v string)`

SetIdTag sets IdTag field to given value.

### HasIdTag

`func (o *BTAssociativeDataInfo) HasIdTag() bool`

HasIdTag returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTAssociativeDataInfo) GetMicroversionId() string`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTAssociativeDataInfo) GetMicroversionIdOk() (*string, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTAssociativeDataInfo) SetMicroversionId(v string)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTAssociativeDataInfo) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetOccurrenceId

`func (o *BTAssociativeDataInfo) GetOccurrenceId() string`

GetOccurrenceId returns the OccurrenceId field if non-nil, zero value otherwise.

### GetOccurrenceIdOk

`func (o *BTAssociativeDataInfo) GetOccurrenceIdOk() (*string, bool)`

GetOccurrenceIdOk returns a tuple with the OccurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceId

`func (o *BTAssociativeDataInfo) SetOccurrenceId(v string)`

SetOccurrenceId sets OccurrenceId field to given value.

### HasOccurrenceId

`func (o *BTAssociativeDataInfo) HasOccurrenceId() bool`

HasOccurrenceId returns a boolean if a field has been set.

### GetType

`func (o *BTAssociativeDataInfo) GetType() GBTAppElementAssociativeDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTAssociativeDataInfo) GetTypeOk() (*GBTAppElementAssociativeDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTAssociativeDataInfo) SetType(v GBTAppElementAssociativeDataType)`

SetType sets Type field to given value.

### HasType

`func (o *BTAssociativeDataInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersionId

`func (o *BTAssociativeDataInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTAssociativeDataInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTAssociativeDataInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTAssociativeDataInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


