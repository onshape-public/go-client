# BTWebhookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **string** | Company admins can register webhooks to listen to all company events. | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Webhook description. | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** | List of events for which webhook callback is invoked. | [optional] 
**ExternalSessionId** | Pointer to **string** | Applications can pass this parameter as X-Session-ID with every REST call to distinguish webhooks triggered by self. | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsTransient** | Pointer to **bool** | Transient webhooks are automatically cleaned up after a period of inactivity. | [optional] [default to true]
**LinkDocumentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Webhook name. | [optional] 
**Options** | Pointer to [**BTWebhookOptions**](BTWebhookOptions.md) |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWebhookParams

`func NewBTWebhookParams() *BTWebhookParams`

NewBTWebhookParams instantiates a new BTWebhookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWebhookParamsWithDefaults

`func NewBTWebhookParamsWithDefaults() *BTWebhookParams`

NewBTWebhookParamsWithDefaults instantiates a new BTWebhookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BTWebhookParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BTWebhookParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BTWebhookParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BTWebhookParams) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTWebhookParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTWebhookParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTWebhookParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTWebhookParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetData

`func (o *BTWebhookParams) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTWebhookParams) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTWebhookParams) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTWebhookParams) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *BTWebhookParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWebhookParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWebhookParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWebhookParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTWebhookParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTWebhookParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTWebhookParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTWebhookParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTWebhookParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTWebhookParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTWebhookParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTWebhookParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEvents

`func (o *BTWebhookParams) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BTWebhookParams) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BTWebhookParams) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BTWebhookParams) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExternalSessionId

`func (o *BTWebhookParams) GetExternalSessionId() string`

GetExternalSessionId returns the ExternalSessionId field if non-nil, zero value otherwise.

### GetExternalSessionIdOk

`func (o *BTWebhookParams) GetExternalSessionIdOk() (*string, bool)`

GetExternalSessionIdOk returns a tuple with the ExternalSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSessionId

`func (o *BTWebhookParams) SetExternalSessionId(v string)`

SetExternalSessionId sets ExternalSessionId field to given value.

### HasExternalSessionId

`func (o *BTWebhookParams) HasExternalSessionId() bool`

HasExternalSessionId returns a boolean if a field has been set.

### GetFilter

`func (o *BTWebhookParams) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BTWebhookParams) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BTWebhookParams) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BTWebhookParams) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFolderId

`func (o *BTWebhookParams) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BTWebhookParams) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BTWebhookParams) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BTWebhookParams) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *BTWebhookParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWebhookParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWebhookParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWebhookParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsTransient

`func (o *BTWebhookParams) GetIsTransient() bool`

GetIsTransient returns the IsTransient field if non-nil, zero value otherwise.

### GetIsTransientOk

`func (o *BTWebhookParams) GetIsTransientOk() (*bool, bool)`

GetIsTransientOk returns a tuple with the IsTransient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransient

`func (o *BTWebhookParams) SetIsTransient(v bool)`

SetIsTransient sets IsTransient field to given value.

### HasIsTransient

`func (o *BTWebhookParams) HasIsTransient() bool`

HasIsTransient returns a boolean if a field has been set.

### GetLinkDocumentId

`func (o *BTWebhookParams) GetLinkDocumentId() string`

GetLinkDocumentId returns the LinkDocumentId field if non-nil, zero value otherwise.

### GetLinkDocumentIdOk

`func (o *BTWebhookParams) GetLinkDocumentIdOk() (*string, bool)`

GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentId

`func (o *BTWebhookParams) SetLinkDocumentId(v string)`

SetLinkDocumentId sets LinkDocumentId field to given value.

### HasLinkDocumentId

`func (o *BTWebhookParams) HasLinkDocumentId() bool`

HasLinkDocumentId returns a boolean if a field has been set.

### GetName

`func (o *BTWebhookParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWebhookParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWebhookParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWebhookParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *BTWebhookParams) GetOptions() BTWebhookOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BTWebhookParams) GetOptionsOk() (*BTWebhookOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BTWebhookParams) SetOptions(v BTWebhookOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BTWebhookParams) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPartId

`func (o *BTWebhookParams) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTWebhookParams) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTWebhookParams) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTWebhookParams) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTWebhookParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTWebhookParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTWebhookParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTWebhookParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetUrl

`func (o *BTWebhookParams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BTWebhookParams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BTWebhookParams) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BTWebhookParams) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserId

`func (o *BTWebhookParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTWebhookParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTWebhookParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTWebhookParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTWebhookParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTWebhookParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTWebhookParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTWebhookParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTWebhookParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTWebhookParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTWebhookParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTWebhookParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


