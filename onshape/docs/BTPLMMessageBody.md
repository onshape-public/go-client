# BTPLMMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppElementSessionId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** | Background PLM job&#39;s document ID. | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** | ID of the background PLM job that was created. | [optional] 
**JobType** | Pointer to [**JobType**](JobType.md) |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**SettingsDisabled** | Pointer to **bool** | Whether PLM integration was disabled. | [optional] 
**SettingsModified** | Pointer to **bool** | Whether PLM integration settings parameters were modified. | [optional] 
**Timestamp** | Pointer to **JSONTime** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPLMMessageBody

`func NewBTPLMMessageBody() *BTPLMMessageBody`

NewBTPLMMessageBody instantiates a new BTPLMMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPLMMessageBodyWithDefaults

`func NewBTPLMMessageBodyWithDefaults() *BTPLMMessageBody`

NewBTPLMMessageBodyWithDefaults instantiates a new BTPLMMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *BTPLMMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *BTPLMMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *BTPLMMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *BTPLMMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetData

`func (o *BTPLMMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTPLMMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTPLMMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTPLMMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTPLMMessageBody) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTPLMMessageBody) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTPLMMessageBody) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTPLMMessageBody) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetEvent

`func (o *BTPLMMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BTPLMMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BTPLMMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BTPLMMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetJobId

`func (o *BTPLMMessageBody) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BTPLMMessageBody) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BTPLMMessageBody) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BTPLMMessageBody) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobType

`func (o *BTPLMMessageBody) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BTPLMMessageBody) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BTPLMMessageBody) SetJobType(v JobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BTPLMMessageBody) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetMessageId

`func (o *BTPLMMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BTPLMMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BTPLMMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BTPLMMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetSettingsDisabled

`func (o *BTPLMMessageBody) GetSettingsDisabled() bool`

GetSettingsDisabled returns the SettingsDisabled field if non-nil, zero value otherwise.

### GetSettingsDisabledOk

`func (o *BTPLMMessageBody) GetSettingsDisabledOk() (*bool, bool)`

GetSettingsDisabledOk returns a tuple with the SettingsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsDisabled

`func (o *BTPLMMessageBody) SetSettingsDisabled(v bool)`

SetSettingsDisabled sets SettingsDisabled field to given value.

### HasSettingsDisabled

`func (o *BTPLMMessageBody) HasSettingsDisabled() bool`

HasSettingsDisabled returns a boolean if a field has been set.

### GetSettingsModified

`func (o *BTPLMMessageBody) GetSettingsModified() bool`

GetSettingsModified returns the SettingsModified field if non-nil, zero value otherwise.

### GetSettingsModifiedOk

`func (o *BTPLMMessageBody) GetSettingsModifiedOk() (*bool, bool)`

GetSettingsModifiedOk returns a tuple with the SettingsModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsModified

`func (o *BTPLMMessageBody) SetSettingsModified(v bool)`

SetSettingsModified sets SettingsModified field to given value.

### HasSettingsModified

`func (o *BTPLMMessageBody) HasSettingsModified() bool`

HasSettingsModified returns a boolean if a field has been set.

### GetTimestamp

`func (o *BTPLMMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BTPLMMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BTPLMMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BTPLMMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetWebhookId

`func (o *BTPLMMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BTPLMMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BTPLMMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BTPLMMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


