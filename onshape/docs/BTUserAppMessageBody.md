# BTUserAppMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppElementSessionId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**IdentityId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**SettingType** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **JSONTime** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTUserAppMessageBody

`func NewBTUserAppMessageBody() *BTUserAppMessageBody`

NewBTUserAppMessageBody instantiates a new BTUserAppMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserAppMessageBodyWithDefaults

`func NewBTUserAppMessageBodyWithDefaults() *BTUserAppMessageBody`

NewBTUserAppMessageBodyWithDefaults instantiates a new BTUserAppMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *BTUserAppMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *BTUserAppMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *BTUserAppMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *BTUserAppMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetClientId

`func (o *BTUserAppMessageBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BTUserAppMessageBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BTUserAppMessageBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BTUserAppMessageBody) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetData

`func (o *BTUserAppMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTUserAppMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTUserAppMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTUserAppMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEvent

`func (o *BTUserAppMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BTUserAppMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BTUserAppMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BTUserAppMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetIdentityId

`func (o *BTUserAppMessageBody) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *BTUserAppMessageBody) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *BTUserAppMessageBody) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *BTUserAppMessageBody) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetMessageId

`func (o *BTUserAppMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BTUserAppMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BTUserAppMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BTUserAppMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetSettingType

`func (o *BTUserAppMessageBody) GetSettingType() string`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *BTUserAppMessageBody) GetSettingTypeOk() (*string, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *BTUserAppMessageBody) SetSettingType(v string)`

SetSettingType sets SettingType field to given value.

### HasSettingType

`func (o *BTUserAppMessageBody) HasSettingType() bool`

HasSettingType returns a boolean if a field has been set.

### GetTimestamp

`func (o *BTUserAppMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BTUserAppMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BTUserAppMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BTUserAppMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetWebhookId

`func (o *BTUserAppMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BTUserAppMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BTUserAppMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BTUserAppMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


