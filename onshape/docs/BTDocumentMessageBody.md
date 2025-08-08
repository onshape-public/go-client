# BTDocumentMessageBody

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
**EntryId** | Pointer to **string** |  | [optional] 
**EntryType** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**MetadataObjectType** | Pointer to **string** |  | [optional] 
**NewPermissionSet** | Pointer to **[]string** |  | [optional] 
**OldPermissionSet** | Pointer to **[]string** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ShareAction** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **JSONTime** |  | [optional] 
**TranslationId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTDocumentMessageBody

`func NewBTDocumentMessageBody() *BTDocumentMessageBody`

NewBTDocumentMessageBody instantiates a new BTDocumentMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentMessageBodyWithDefaults

`func NewBTDocumentMessageBodyWithDefaults() *BTDocumentMessageBody`

NewBTDocumentMessageBodyWithDefaults instantiates a new BTDocumentMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppElementSessionId

`func (o *BTDocumentMessageBody) GetAppElementSessionId() string`

GetAppElementSessionId returns the AppElementSessionId field if non-nil, zero value otherwise.

### GetAppElementSessionIdOk

`func (o *BTDocumentMessageBody) GetAppElementSessionIdOk() (*string, bool)`

GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppElementSessionId

`func (o *BTDocumentMessageBody) SetAppElementSessionId(v string)`

SetAppElementSessionId sets AppElementSessionId field to given value.

### HasAppElementSessionId

`func (o *BTDocumentMessageBody) HasAppElementSessionId() bool`

HasAppElementSessionId returns a boolean if a field has been set.

### GetCommentId

`func (o *BTDocumentMessageBody) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *BTDocumentMessageBody) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *BTDocumentMessageBody) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *BTDocumentMessageBody) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetData

`func (o *BTDocumentMessageBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTDocumentMessageBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTDocumentMessageBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTDocumentMessageBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTDocumentMessageBody) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTDocumentMessageBody) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTDocumentMessageBody) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTDocumentMessageBody) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentState

`func (o *BTDocumentMessageBody) GetDocumentState() string`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *BTDocumentMessageBody) GetDocumentStateOk() (*string, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *BTDocumentMessageBody) SetDocumentState(v string)`

SetDocumentState sets DocumentState field to given value.

### HasDocumentState

`func (o *BTDocumentMessageBody) HasDocumentState() bool`

HasDocumentState returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTDocumentMessageBody) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTDocumentMessageBody) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTDocumentMessageBody) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTDocumentMessageBody) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetElementId

`func (o *BTDocumentMessageBody) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTDocumentMessageBody) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTDocumentMessageBody) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTDocumentMessageBody) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEntryId

`func (o *BTDocumentMessageBody) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BTDocumentMessageBody) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BTDocumentMessageBody) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *BTDocumentMessageBody) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEntryType

`func (o *BTDocumentMessageBody) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BTDocumentMessageBody) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BTDocumentMessageBody) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *BTDocumentMessageBody) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetEvent

`func (o *BTDocumentMessageBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BTDocumentMessageBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BTDocumentMessageBody) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BTDocumentMessageBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMessageId

`func (o *BTDocumentMessageBody) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BTDocumentMessageBody) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BTDocumentMessageBody) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BTDocumentMessageBody) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMetadataObjectType

`func (o *BTDocumentMessageBody) GetMetadataObjectType() string`

GetMetadataObjectType returns the MetadataObjectType field if non-nil, zero value otherwise.

### GetMetadataObjectTypeOk

`func (o *BTDocumentMessageBody) GetMetadataObjectTypeOk() (*string, bool)`

GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObjectType

`func (o *BTDocumentMessageBody) SetMetadataObjectType(v string)`

SetMetadataObjectType sets MetadataObjectType field to given value.

### HasMetadataObjectType

`func (o *BTDocumentMessageBody) HasMetadataObjectType() bool`

HasMetadataObjectType returns a boolean if a field has been set.

### GetNewPermissionSet

`func (o *BTDocumentMessageBody) GetNewPermissionSet() []string`

GetNewPermissionSet returns the NewPermissionSet field if non-nil, zero value otherwise.

### GetNewPermissionSetOk

`func (o *BTDocumentMessageBody) GetNewPermissionSetOk() (*[]string, bool)`

GetNewPermissionSetOk returns a tuple with the NewPermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissionSet

`func (o *BTDocumentMessageBody) SetNewPermissionSet(v []string)`

SetNewPermissionSet sets NewPermissionSet field to given value.

### HasNewPermissionSet

`func (o *BTDocumentMessageBody) HasNewPermissionSet() bool`

HasNewPermissionSet returns a boolean if a field has been set.

### GetOldPermissionSet

`func (o *BTDocumentMessageBody) GetOldPermissionSet() []string`

GetOldPermissionSet returns the OldPermissionSet field if non-nil, zero value otherwise.

### GetOldPermissionSetOk

`func (o *BTDocumentMessageBody) GetOldPermissionSetOk() (*[]string, bool)`

GetOldPermissionSetOk returns a tuple with the OldPermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPermissionSet

`func (o *BTDocumentMessageBody) SetOldPermissionSet(v []string)`

SetOldPermissionSet sets OldPermissionSet field to given value.

### HasOldPermissionSet

`func (o *BTDocumentMessageBody) HasOldPermissionSet() bool`

HasOldPermissionSet returns a boolean if a field has been set.

### GetPartId

`func (o *BTDocumentMessageBody) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTDocumentMessageBody) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTDocumentMessageBody) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTDocumentMessageBody) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTDocumentMessageBody) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTDocumentMessageBody) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTDocumentMessageBody) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTDocumentMessageBody) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTDocumentMessageBody) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTDocumentMessageBody) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTDocumentMessageBody) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTDocumentMessageBody) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetResourceType

`func (o *BTDocumentMessageBody) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTDocumentMessageBody) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTDocumentMessageBody) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTDocumentMessageBody) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetShareAction

`func (o *BTDocumentMessageBody) GetShareAction() string`

GetShareAction returns the ShareAction field if non-nil, zero value otherwise.

### GetShareActionOk

`func (o *BTDocumentMessageBody) GetShareActionOk() (*string, bool)`

GetShareActionOk returns a tuple with the ShareAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAction

`func (o *BTDocumentMessageBody) SetShareAction(v string)`

SetShareAction sets ShareAction field to given value.

### HasShareAction

`func (o *BTDocumentMessageBody) HasShareAction() bool`

HasShareAction returns a boolean if a field has been set.

### GetTimestamp

`func (o *BTDocumentMessageBody) GetTimestamp() JSONTime`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BTDocumentMessageBody) GetTimestampOk() (*JSONTime, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BTDocumentMessageBody) SetTimestamp(v JSONTime)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BTDocumentMessageBody) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTranslationId

`func (o *BTDocumentMessageBody) GetTranslationId() string`

GetTranslationId returns the TranslationId field if non-nil, zero value otherwise.

### GetTranslationIdOk

`func (o *BTDocumentMessageBody) GetTranslationIdOk() (*string, bool)`

GetTranslationIdOk returns a tuple with the TranslationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationId

`func (o *BTDocumentMessageBody) SetTranslationId(v string)`

SetTranslationId sets TranslationId field to given value.

### HasTranslationId

`func (o *BTDocumentMessageBody) HasTranslationId() bool`

HasTranslationId returns a boolean if a field has been set.

### GetUserId

`func (o *BTDocumentMessageBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTDocumentMessageBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTDocumentMessageBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTDocumentMessageBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTDocumentMessageBody) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTDocumentMessageBody) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTDocumentMessageBody) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTDocumentMessageBody) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWebhookId

`func (o *BTDocumentMessageBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BTDocumentMessageBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BTDocumentMessageBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BTDocumentMessageBody) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTDocumentMessageBody) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTDocumentMessageBody) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTDocumentMessageBody) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTDocumentMessageBody) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


