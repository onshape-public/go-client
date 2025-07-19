# BTWorkspaceMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppElementSessionId** | Pointer to **string** |  | [optional] 
**CommentId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentState** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **int32** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**MetadataObjectType** | Pointer to **string** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **JSONTime** |  | [optional] 
**TranslationId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**DocumentMicroversionId** | Pointer to **string** | The resultant document microverion if applicable created due to workspace modification. | [optional] 

## Methods

### NewBTWorkspaceMessageBody

`func NewBTWorkspaceMessageBody() *BTWorkspaceMessageBody`

NewBTWorkspaceMessageBody instantiates a new BTWorkspaceMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkspaceMessageBodyWithDefaults

`func NewBTWorkspaceMessageBodyWithDefaults() *BTWorkspaceMessageBody`

NewBTWorkspaceMessageBodyWithDefaults instantiates a new BTWorkspaceMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *BTWorkspaceMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *BTWorkspaceMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *BTWorkspaceMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *BTWorkspaceMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetCommentId

`func (o *BTWorkspaceMessageBody) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *BTWorkspaceMessageBody) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *BTWorkspaceMessageBody) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *BTWorkspaceMessageBody) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetData

`func (o *BTWorkspaceMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTWorkspaceMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTWorkspaceMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTWorkspaceMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTWorkspaceMessageBody) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTWorkspaceMessageBody) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTWorkspaceMessageBody) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTWorkspaceMessageBody) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentState

`func (o *BTWorkspaceMessageBody) GetDocumentState() string`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *BTWorkspaceMessageBody) GetDocumentStateOk() (*string, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *BTWorkspaceMessageBody) SetDocumentState(v string)`

SetDocumentState sets DocumentState field to given value.

### HasDocumentState

`func (o *BTWorkspaceMessageBody) HasDocumentState() bool`

HasDocumentState returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTWorkspaceMessageBody) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTWorkspaceMessageBody) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTWorkspaceMessageBody) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTWorkspaceMessageBody) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetElementId

`func (o *BTWorkspaceMessageBody) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTWorkspaceMessageBody) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTWorkspaceMessageBody) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTWorkspaceMessageBody) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEvent

`func (o *BTWorkspaceMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BTWorkspaceMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BTWorkspaceMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BTWorkspaceMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMessageId

`func (o *BTWorkspaceMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BTWorkspaceMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BTWorkspaceMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BTWorkspaceMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMetadataObjectType

`func (o *BTWorkspaceMessageBody) GetMetadataObjectType() string`

GetMetadataObjectType returns the MetadataObjectType field if non-nil, zero value otherwise.

### GetMetadataObjectTypeOk

`func (o *BTWorkspaceMessageBody) GetMetadataObjectTypeOk() (*string, bool)`

GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObjectType

`func (o *BTWorkspaceMessageBody) SetMetadataObjectType(v string)`

SetMetadataObjectType sets MetadataObjectType field to given value.

### HasMetadataObjectType

`func (o *BTWorkspaceMessageBody) HasMetadataObjectType() bool`

HasMetadataObjectType returns a boolean if a field has been set.

### GetPartId

`func (o *BTWorkspaceMessageBody) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTWorkspaceMessageBody) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTWorkspaceMessageBody) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTWorkspaceMessageBody) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTWorkspaceMessageBody) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTWorkspaceMessageBody) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTWorkspaceMessageBody) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTWorkspaceMessageBody) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTWorkspaceMessageBody) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTWorkspaceMessageBody) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTWorkspaceMessageBody) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTWorkspaceMessageBody) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetTimestamp

`func (o *BTWorkspaceMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BTWorkspaceMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BTWorkspaceMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BTWorkspaceMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTranslationId

`func (o *BTWorkspaceMessageBody) GetTranslationId() string`

GetTranslationId returns the TranslationId field if non-nil, zero value otherwise.

### GetTranslationIdOk

`func (o *BTWorkspaceMessageBody) GetTranslationIdOk() (*string, bool)`

GetTranslationIdOk returns a tuple with the TranslationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationId

`func (o *BTWorkspaceMessageBody) SetTranslationId(v string)`

SetTranslationId sets TranslationId field to given value.

### HasTranslationId

`func (o *BTWorkspaceMessageBody) HasTranslationId() bool`

HasTranslationId returns a boolean if a field has been set.

### GetUserId

`func (o *BTWorkspaceMessageBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTWorkspaceMessageBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTWorkspaceMessageBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTWorkspaceMessageBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTWorkspaceMessageBody) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTWorkspaceMessageBody) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTWorkspaceMessageBody) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTWorkspaceMessageBody) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWebhookId

`func (o *BTWorkspaceMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BTWorkspaceMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BTWorkspaceMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BTWorkspaceMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTWorkspaceMessageBody) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTWorkspaceMessageBody) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTWorkspaceMessageBody) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTWorkspaceMessageBody) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetDocumentMicroversionId

`func (o *BTWorkspaceMessageBody) GetDocumentMicroversionId() string`

GetDocumentMicroversionId returns the DocumentMicroversionId field if non-nil, zero value otherwise.

### GetDocumentMicroversionIdOk

`func (o *BTWorkspaceMessageBody) GetDocumentMicroversionIdOk() (*string, bool)`

GetDocumentMicroversionIdOk returns a tuple with the DocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversionId

`func (o *BTWorkspaceMessageBody) SetDocumentMicroversionId(v string)`

SetDocumentMicroversionId sets DocumentMicroversionId field to given value.

### HasDocumentMicroversionId

`func (o *BTWorkspaceMessageBody) HasDocumentMicroversionId() bool`

HasDocumentMicroversionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


