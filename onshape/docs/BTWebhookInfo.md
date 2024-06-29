# BTWebhookInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | Company admins can register webhooks to listen to all company events. | [optional] 
**CreatedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DroppedEventCount** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to **[]string** | List of events for which webhook callback is invoked. | [optional] 
**ExternalSessionId** | Pointer to **string** | Applications can pass this parameter as X-Session-ID with every REST call to distinguish webhooks triggered by self. | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsTransient** | Pointer to **bool** | Transient webhooks are automatically cleaned up after a period of inactivity. | [optional] [default to true]
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Options** | Pointer to [**BTWebhookOptions**](BTWebhookOptions.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTWebhookInfo

`func NewBTWebhookInfo() *BTWebhookInfo`

NewBTWebhookInfo instantiates a new BTWebhookInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWebhookInfoWithDefaults

`func NewBTWebhookInfoWithDefaults() *BTWebhookInfo`

NewBTWebhookInfoWithDefaults instantiates a new BTWebhookInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTWebhookInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTWebhookInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTWebhookInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTWebhookInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTWebhookInfo) GetCreatedBy() BTUserSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTWebhookInfo) GetCreatedByOk() (*BTUserSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTWebhookInfo) SetCreatedBy(v BTUserSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTWebhookInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetData

`func (o *BTWebhookInfo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BTWebhookInfo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BTWebhookInfo) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BTWebhookInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *BTWebhookInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWebhookInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWebhookInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWebhookInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDroppedEventCount

`func (o *BTWebhookInfo) GetDroppedEventCount() int32`

GetDroppedEventCount returns the DroppedEventCount field if non-nil, zero value otherwise.

### GetDroppedEventCountOk

`func (o *BTWebhookInfo) GetDroppedEventCountOk() (*int32, bool)`

GetDroppedEventCountOk returns a tuple with the DroppedEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedEventCount

`func (o *BTWebhookInfo) SetDroppedEventCount(v int32)`

SetDroppedEventCount sets DroppedEventCount field to given value.

### HasDroppedEventCount

`func (o *BTWebhookInfo) HasDroppedEventCount() bool`

HasDroppedEventCount returns a boolean if a field has been set.

### GetEvents

`func (o *BTWebhookInfo) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BTWebhookInfo) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BTWebhookInfo) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BTWebhookInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExternalSessionId

`func (o *BTWebhookInfo) GetExternalSessionId() string`

GetExternalSessionId returns the ExternalSessionId field if non-nil, zero value otherwise.

### GetExternalSessionIdOk

`func (o *BTWebhookInfo) GetExternalSessionIdOk() (*string, bool)`

GetExternalSessionIdOk returns a tuple with the ExternalSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSessionId

`func (o *BTWebhookInfo) SetExternalSessionId(v string)`

SetExternalSessionId sets ExternalSessionId field to given value.

### HasExternalSessionId

`func (o *BTWebhookInfo) HasExternalSessionId() bool`

HasExternalSessionId returns a boolean if a field has been set.

### GetFilter

`func (o *BTWebhookInfo) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BTWebhookInfo) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BTWebhookInfo) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BTWebhookInfo) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFolderId

`func (o *BTWebhookInfo) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BTWebhookInfo) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BTWebhookInfo) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BTWebhookInfo) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetHref

`func (o *BTWebhookInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTWebhookInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTWebhookInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTWebhookInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTWebhookInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWebhookInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWebhookInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWebhookInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsTransient

`func (o *BTWebhookInfo) GetIsTransient() bool`

GetIsTransient returns the IsTransient field if non-nil, zero value otherwise.

### GetIsTransientOk

`func (o *BTWebhookInfo) GetIsTransientOk() (*bool, bool)`

GetIsTransientOk returns a tuple with the IsTransient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransient

`func (o *BTWebhookInfo) SetIsTransient(v bool)`

SetIsTransient sets IsTransient field to given value.

### HasIsTransient

`func (o *BTWebhookInfo) HasIsTransient() bool`

HasIsTransient returns a boolean if a field has been set.

### GetName

`func (o *BTWebhookInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWebhookInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWebhookInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWebhookInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *BTWebhookInfo) GetOptions() BTWebhookOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BTWebhookInfo) GetOptionsOk() (*BTWebhookOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BTWebhookInfo) SetOptions(v BTWebhookOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BTWebhookInfo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProjectId

`func (o *BTWebhookInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTWebhookInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTWebhookInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTWebhookInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetUrl

`func (o *BTWebhookInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BTWebhookInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BTWebhookInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BTWebhookInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetViewRef

`func (o *BTWebhookInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTWebhookInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTWebhookInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTWebhookInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


