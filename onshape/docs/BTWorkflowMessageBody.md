# BTWorkflowMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppElementSessionId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **JSONTime** |  | [optional] 
**TransitionName** | Pointer to **string** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkflowMessageBody

`func NewBTWorkflowMessageBody() *BTWorkflowMessageBody`

NewBTWorkflowMessageBody instantiates a new BTWorkflowMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowMessageBodyWithDefaults

`func NewBTWorkflowMessageBodyWithDefaults() *BTWorkflowMessageBody`

NewBTWorkflowMessageBodyWithDefaults instantiates a new BTWorkflowMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *BTWorkflowMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *BTWorkflowMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *BTWorkflowMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *BTWorkflowMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetData

`func (o *BTWorkflowMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTWorkflowMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTWorkflowMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTWorkflowMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEvent

`func (o *BTWorkflowMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BTWorkflowMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BTWorkflowMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BTWorkflowMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMessageId

`func (o *BTWorkflowMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BTWorkflowMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BTWorkflowMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BTWorkflowMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetObjectId

`func (o *BTWorkflowMessageBody) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTWorkflowMessageBody) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTWorkflowMessageBody) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTWorkflowMessageBody) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *BTWorkflowMessageBody) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTWorkflowMessageBody) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTWorkflowMessageBody) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTWorkflowMessageBody) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetTimestamp

`func (o *BTWorkflowMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BTWorkflowMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BTWorkflowMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BTWorkflowMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransitionName

`func (o *BTWorkflowMessageBody) GetTransitionName() string`

GetTransitionName returns the TransitionName field if non-nil, zero value otherwise.

### GetTransitionNameOk

`func (o *BTWorkflowMessageBody) GetTransitionNameOk() (*string, bool)`

GetTransitionNameOk returns a tuple with the TransitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionName

`func (o *BTWorkflowMessageBody) SetTransitionName(v string)`

SetTransitionName sets TransitionName field to given value.

### HasTransitionName

`func (o *BTWorkflowMessageBody) HasTransitionName() bool`

HasTransitionName returns a boolean if a field has been set.

### GetWebhookId

`func (o *BTWorkflowMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BTWorkflowMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BTWorkflowMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BTWorkflowMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *BTWorkflowMessageBody) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *BTWorkflowMessageBody) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *BTWorkflowMessageBody) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *BTWorkflowMessageBody) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


