# GBTDeleteElementMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppElementSessionId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **int32** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **JSONTime** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewGBTDeleteElementMessageBody

`func NewGBTDeleteElementMessageBody() *GBTDeleteElementMessageBody`

NewGBTDeleteElementMessageBody instantiates a new GBTDeleteElementMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGBTDeleteElementMessageBodyWithDefaults

`func NewGBTDeleteElementMessageBodyWithDefaults() *GBTDeleteElementMessageBody`

NewGBTDeleteElementMessageBodyWithDefaults instantiates a new GBTDeleteElementMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *GBTDeleteElementMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *GBTDeleteElementMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *GBTDeleteElementMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *GBTDeleteElementMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetData

`func (o *GBTDeleteElementMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GBTDeleteElementMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GBTDeleteElementMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GBTDeleteElementMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDocumentId

`func (o *GBTDeleteElementMessageBody) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *GBTDeleteElementMessageBody) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *GBTDeleteElementMessageBody) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *GBTDeleteElementMessageBody) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentType

`func (o *GBTDeleteElementMessageBody) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GBTDeleteElementMessageBody) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GBTDeleteElementMessageBody) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GBTDeleteElementMessageBody) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetElementId

`func (o *GBTDeleteElementMessageBody) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *GBTDeleteElementMessageBody) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *GBTDeleteElementMessageBody) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *GBTDeleteElementMessageBody) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEvent

`func (o *GBTDeleteElementMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GBTDeleteElementMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GBTDeleteElementMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *GBTDeleteElementMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMessageId

`func (o *GBTDeleteElementMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *GBTDeleteElementMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *GBTDeleteElementMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *GBTDeleteElementMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetTimestamp

`func (o *GBTDeleteElementMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GBTDeleteElementMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GBTDeleteElementMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GBTDeleteElementMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetWebhookId

`func (o *GBTDeleteElementMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *GBTDeleteElementMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *GBTDeleteElementMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *GBTDeleteElementMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *GBTDeleteElementMessageBody) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GBTDeleteElementMessageBody) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GBTDeleteElementMessageBody) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *GBTDeleteElementMessageBody) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


