# BTReleaseItemMessageBody

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
**TranslatationId** | Pointer to **string** |  | [optional] 
**TranslationId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **int32** |  | [optional] 
**IsFromInitialState** | Pointer to **bool** |  | [optional] 
**IsToTerminalState** | Pointer to **bool** |  | [optional] 
**ReleaseId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TransitionName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTReleaseItemMessageBody

`func NewBTReleaseItemMessageBody() *BTReleaseItemMessageBody`

NewBTReleaseItemMessageBody instantiates a new BTReleaseItemMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTReleaseItemMessageBodyWithDefaults

`func NewBTReleaseItemMessageBodyWithDefaults() *BTReleaseItemMessageBody`

NewBTReleaseItemMessageBodyWithDefaults instantiates a new BTReleaseItemMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *BTReleaseItemMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *BTReleaseItemMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *BTReleaseItemMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *BTReleaseItemMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetCommentId

`func (o *BTReleaseItemMessageBody) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *BTReleaseItemMessageBody) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *BTReleaseItemMessageBody) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *BTReleaseItemMessageBody) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetData

`func (o *BTReleaseItemMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTReleaseItemMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTReleaseItemMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTReleaseItemMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTReleaseItemMessageBody) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTReleaseItemMessageBody) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTReleaseItemMessageBody) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTReleaseItemMessageBody) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentState

`func (o *BTReleaseItemMessageBody) GetDocumentState() string`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *BTReleaseItemMessageBody) GetDocumentStateOk() (*string, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *BTReleaseItemMessageBody) SetDocumentState(v string)`

SetDocumentState sets DocumentState field to given value.

### HasDocumentState

`func (o *BTReleaseItemMessageBody) HasDocumentState() bool`

HasDocumentState returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTReleaseItemMessageBody) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTReleaseItemMessageBody) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTReleaseItemMessageBody) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTReleaseItemMessageBody) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetElementId

`func (o *BTReleaseItemMessageBody) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTReleaseItemMessageBody) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTReleaseItemMessageBody) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTReleaseItemMessageBody) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEvent

`func (o *BTReleaseItemMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BTReleaseItemMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BTReleaseItemMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BTReleaseItemMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMessageId

`func (o *BTReleaseItemMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BTReleaseItemMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BTReleaseItemMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BTReleaseItemMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMetadataObjectType

`func (o *BTReleaseItemMessageBody) GetMetadataObjectType() string`

GetMetadataObjectType returns the MetadataObjectType field if non-nil, zero value otherwise.

### GetMetadataObjectTypeOk

`func (o *BTReleaseItemMessageBody) GetMetadataObjectTypeOk() (*string, bool)`

GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObjectType

`func (o *BTReleaseItemMessageBody) SetMetadataObjectType(v string)`

SetMetadataObjectType sets MetadataObjectType field to given value.

### HasMetadataObjectType

`func (o *BTReleaseItemMessageBody) HasMetadataObjectType() bool`

HasMetadataObjectType returns a boolean if a field has been set.

### GetPartId

`func (o *BTReleaseItemMessageBody) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTReleaseItemMessageBody) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTReleaseItemMessageBody) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTReleaseItemMessageBody) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTReleaseItemMessageBody) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTReleaseItemMessageBody) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTReleaseItemMessageBody) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTReleaseItemMessageBody) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTReleaseItemMessageBody) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTReleaseItemMessageBody) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTReleaseItemMessageBody) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTReleaseItemMessageBody) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetTimestamp

`func (o *BTReleaseItemMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BTReleaseItemMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BTReleaseItemMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BTReleaseItemMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTranslatationId

`func (o *BTReleaseItemMessageBody) GetTranslatationId() string`

GetTranslatationId returns the TranslatationId field if non-nil, zero value otherwise.

### GetTranslatationIdOk

`func (o *BTReleaseItemMessageBody) GetTranslatationIdOk() (*string, bool)`

GetTranslatationIdOk returns a tuple with the TranslatationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatationId

`func (o *BTReleaseItemMessageBody) SetTranslatationId(v string)`

SetTranslatationId sets TranslatationId field to given value.

### HasTranslatationId

`func (o *BTReleaseItemMessageBody) HasTranslatationId() bool`

HasTranslatationId returns a boolean if a field has been set.

### GetTranslationId

`func (o *BTReleaseItemMessageBody) GetTranslationId() string`

GetTranslationId returns the TranslationId field if non-nil, zero value otherwise.

### GetTranslationIdOk

`func (o *BTReleaseItemMessageBody) GetTranslationIdOk() (*string, bool)`

GetTranslationIdOk returns a tuple with the TranslationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationId

`func (o *BTReleaseItemMessageBody) SetTranslationId(v string)`

SetTranslationId sets TranslationId field to given value.

### HasTranslationId

`func (o *BTReleaseItemMessageBody) HasTranslationId() bool`

HasTranslationId returns a boolean if a field has been set.

### GetUserId

`func (o *BTReleaseItemMessageBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTReleaseItemMessageBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTReleaseItemMessageBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTReleaseItemMessageBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTReleaseItemMessageBody) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTReleaseItemMessageBody) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTReleaseItemMessageBody) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTReleaseItemMessageBody) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWebhookId

`func (o *BTReleaseItemMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BTReleaseItemMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BTReleaseItemMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BTReleaseItemMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTReleaseItemMessageBody) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTReleaseItemMessageBody) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTReleaseItemMessageBody) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTReleaseItemMessageBody) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetElementType

`func (o *BTReleaseItemMessageBody) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTReleaseItemMessageBody) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTReleaseItemMessageBody) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTReleaseItemMessageBody) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetIsFromInitialState

`func (o *BTReleaseItemMessageBody) GetIsFromInitialState() bool`

GetIsFromInitialState returns the IsFromInitialState field if non-nil, zero value otherwise.

### GetIsFromInitialStateOk

`func (o *BTReleaseItemMessageBody) GetIsFromInitialStateOk() (*bool, bool)`

GetIsFromInitialStateOk returns a tuple with the IsFromInitialState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromInitialState

`func (o *BTReleaseItemMessageBody) SetIsFromInitialState(v bool)`

SetIsFromInitialState sets IsFromInitialState field to given value.

### HasIsFromInitialState

`func (o *BTReleaseItemMessageBody) HasIsFromInitialState() bool`

HasIsFromInitialState returns a boolean if a field has been set.

### GetIsToTerminalState

`func (o *BTReleaseItemMessageBody) GetIsToTerminalState() bool`

GetIsToTerminalState returns the IsToTerminalState field if non-nil, zero value otherwise.

### GetIsToTerminalStateOk

`func (o *BTReleaseItemMessageBody) GetIsToTerminalStateOk() (*bool, bool)`

GetIsToTerminalStateOk returns a tuple with the IsToTerminalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToTerminalState

`func (o *BTReleaseItemMessageBody) SetIsToTerminalState(v bool)`

SetIsToTerminalState sets IsToTerminalState field to given value.

### HasIsToTerminalState

`func (o *BTReleaseItemMessageBody) HasIsToTerminalState() bool`

HasIsToTerminalState returns a boolean if a field has been set.

### GetReleaseId

`func (o *BTReleaseItemMessageBody) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *BTReleaseItemMessageBody) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *BTReleaseItemMessageBody) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *BTReleaseItemMessageBody) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetStatus

`func (o *BTReleaseItemMessageBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BTReleaseItemMessageBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BTReleaseItemMessageBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BTReleaseItemMessageBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransitionName

`func (o *BTReleaseItemMessageBody) GetTransitionName() string`

GetTransitionName returns the TransitionName field if non-nil, zero value otherwise.

### GetTransitionNameOk

`func (o *BTReleaseItemMessageBody) GetTransitionNameOk() (*string, bool)`

GetTransitionNameOk returns a tuple with the TransitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionName

`func (o *BTReleaseItemMessageBody) SetTransitionName(v string)`

SetTransitionName sets TransitionName field to given value.

### HasTransitionName

`func (o *BTReleaseItemMessageBody) HasTransitionName() bool`

HasTransitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


