# BTActiveWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowReleaseItemsFromOtherDocuments** | Pointer to **bool** |  | [optional] 
**CanCurrentUserCreateReleases** | Pointer to **bool** |  | [optional] 
**CanCurrentUserEditStandardContent** | Pointer to **bool** |  | [optional] 
**CanCurrentUserManageWorkflows** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSeeArenaItemLink** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncBomToArena** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncToArena** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncVersionsToArena** | Pointer to **bool** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DrawingCanDuplicatePartNumber** | Pointer to **bool** |  | [optional] 
**EnabledActiveMultipleWorkflows** | Pointer to **bool** |  | [optional] 
**ObsoletionWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ObsoletionWorkflowId** | Pointer to **string** |  | [optional] 
**PartNumberingSchemeId** | Pointer to **string** |  | [optional] 
**PickableWorkflows** | Pointer to [**[]BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ReleaseWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ReleaseWorkflowId** | Pointer to **string** |  | [optional] 
**ReleaseableApplications** | Pointer to **[]string** |  | [optional] 
**StandardContentNumberingSchemeId** | Pointer to **string** |  | [optional] 
**StandardContentUsingAutoNumbering** | Pointer to **bool** |  | [optional] 
**StandardContentUsingThirdPartyPartNumbering** | Pointer to **bool** |  | [optional] 
**UsingAutoPartNumbering** | Pointer to **bool** |  | [optional] 
**UsingManagedWorkflow** | Pointer to **bool** |  | [optional] 
**UsingThirdPartyPartNumbering** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTActiveWorkflowInfo

`func NewBTActiveWorkflowInfo() *BTActiveWorkflowInfo`

NewBTActiveWorkflowInfo instantiates a new BTActiveWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTActiveWorkflowInfoWithDefaults

`func NewBTActiveWorkflowInfoWithDefaults() *BTActiveWorkflowInfo`

NewBTActiveWorkflowInfoWithDefaults instantiates a new BTActiveWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowReleaseItemsFromOtherDocuments

`func (o *BTActiveWorkflowInfo) GetAllowReleaseItemsFromOtherDocuments() bool`

GetAllowReleaseItemsFromOtherDocuments returns the AllowReleaseItemsFromOtherDocuments field if non-nil, zero value otherwise.

### GetAllowReleaseItemsFromOtherDocumentsOk

`func (o *BTActiveWorkflowInfo) GetAllowReleaseItemsFromOtherDocumentsOk() (*bool, bool)`

GetAllowReleaseItemsFromOtherDocumentsOk returns a tuple with the AllowReleaseItemsFromOtherDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReleaseItemsFromOtherDocuments

`func (o *BTActiveWorkflowInfo) SetAllowReleaseItemsFromOtherDocuments(v bool)`

SetAllowReleaseItemsFromOtherDocuments sets AllowReleaseItemsFromOtherDocuments field to given value.

### HasAllowReleaseItemsFromOtherDocuments

`func (o *BTActiveWorkflowInfo) HasAllowReleaseItemsFromOtherDocuments() bool`

HasAllowReleaseItemsFromOtherDocuments returns a boolean if a field has been set.

### GetCanCurrentUserCreateReleases

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleases() bool`

GetCanCurrentUserCreateReleases returns the CanCurrentUserCreateReleases field if non-nil, zero value otherwise.

### GetCanCurrentUserCreateReleasesOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleasesOk() (*bool, bool)`

GetCanCurrentUserCreateReleasesOk returns a tuple with the CanCurrentUserCreateReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserCreateReleases

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserCreateReleases(v bool)`

SetCanCurrentUserCreateReleases sets CanCurrentUserCreateReleases field to given value.

### HasCanCurrentUserCreateReleases

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserCreateReleases() bool`

HasCanCurrentUserCreateReleases returns a boolean if a field has been set.

### GetCanCurrentUserEditStandardContent

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserEditStandardContent() bool`

GetCanCurrentUserEditStandardContent returns the CanCurrentUserEditStandardContent field if non-nil, zero value otherwise.

### GetCanCurrentUserEditStandardContentOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserEditStandardContentOk() (*bool, bool)`

GetCanCurrentUserEditStandardContentOk returns a tuple with the CanCurrentUserEditStandardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserEditStandardContent

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserEditStandardContent(v bool)`

SetCanCurrentUserEditStandardContent sets CanCurrentUserEditStandardContent field to given value.

### HasCanCurrentUserEditStandardContent

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserEditStandardContent() bool`

HasCanCurrentUserEditStandardContent returns a boolean if a field has been set.

### GetCanCurrentUserManageWorkflows

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserManageWorkflows() bool`

GetCanCurrentUserManageWorkflows returns the CanCurrentUserManageWorkflows field if non-nil, zero value otherwise.

### GetCanCurrentUserManageWorkflowsOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserManageWorkflowsOk() (*bool, bool)`

GetCanCurrentUserManageWorkflowsOk returns a tuple with the CanCurrentUserManageWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserManageWorkflows

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserManageWorkflows(v bool)`

SetCanCurrentUserManageWorkflows sets CanCurrentUserManageWorkflows field to given value.

### HasCanCurrentUserManageWorkflows

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserManageWorkflows() bool`

HasCanCurrentUserManageWorkflows returns a boolean if a field has been set.

### GetCanCurrentUserSeeArenaItemLink

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSeeArenaItemLink() bool`

GetCanCurrentUserSeeArenaItemLink returns the CanCurrentUserSeeArenaItemLink field if non-nil, zero value otherwise.

### GetCanCurrentUserSeeArenaItemLinkOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSeeArenaItemLinkOk() (*bool, bool)`

GetCanCurrentUserSeeArenaItemLinkOk returns a tuple with the CanCurrentUserSeeArenaItemLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSeeArenaItemLink

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSeeArenaItemLink(v bool)`

SetCanCurrentUserSeeArenaItemLink sets CanCurrentUserSeeArenaItemLink field to given value.

### HasCanCurrentUserSeeArenaItemLink

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSeeArenaItemLink() bool`

HasCanCurrentUserSeeArenaItemLink returns a boolean if a field has been set.

### GetCanCurrentUserSyncBomToArena

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncBomToArena() bool`

GetCanCurrentUserSyncBomToArena returns the CanCurrentUserSyncBomToArena field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncBomToArenaOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncBomToArenaOk() (*bool, bool)`

GetCanCurrentUserSyncBomToArenaOk returns a tuple with the CanCurrentUserSyncBomToArena field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncBomToArena

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncBomToArena(v bool)`

SetCanCurrentUserSyncBomToArena sets CanCurrentUserSyncBomToArena field to given value.

### HasCanCurrentUserSyncBomToArena

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncBomToArena() bool`

HasCanCurrentUserSyncBomToArena returns a boolean if a field has been set.

### GetCanCurrentUserSyncToArena

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncToArena() bool`

GetCanCurrentUserSyncToArena returns the CanCurrentUserSyncToArena field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncToArenaOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncToArenaOk() (*bool, bool)`

GetCanCurrentUserSyncToArenaOk returns a tuple with the CanCurrentUserSyncToArena field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncToArena

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncToArena(v bool)`

SetCanCurrentUserSyncToArena sets CanCurrentUserSyncToArena field to given value.

### HasCanCurrentUserSyncToArena

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncToArena() bool`

HasCanCurrentUserSyncToArena returns a boolean if a field has been set.

### GetCanCurrentUserSyncVersionsToArena

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncVersionsToArena() bool`

GetCanCurrentUserSyncVersionsToArena returns the CanCurrentUserSyncVersionsToArena field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncVersionsToArenaOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncVersionsToArenaOk() (*bool, bool)`

GetCanCurrentUserSyncVersionsToArenaOk returns a tuple with the CanCurrentUserSyncVersionsToArena field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncVersionsToArena

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncVersionsToArena(v bool)`

SetCanCurrentUserSyncVersionsToArena sets CanCurrentUserSyncVersionsToArena field to given value.

### HasCanCurrentUserSyncVersionsToArena

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncVersionsToArena() bool`

HasCanCurrentUserSyncVersionsToArena returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTActiveWorkflowInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTActiveWorkflowInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTActiveWorkflowInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTActiveWorkflowInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTActiveWorkflowInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTActiveWorkflowInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTActiveWorkflowInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTActiveWorkflowInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDrawingCanDuplicatePartNumber

`func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumber() bool`

GetDrawingCanDuplicatePartNumber returns the DrawingCanDuplicatePartNumber field if non-nil, zero value otherwise.

### GetDrawingCanDuplicatePartNumberOk

`func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumberOk() (*bool, bool)`

GetDrawingCanDuplicatePartNumberOk returns a tuple with the DrawingCanDuplicatePartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawingCanDuplicatePartNumber

`func (o *BTActiveWorkflowInfo) SetDrawingCanDuplicatePartNumber(v bool)`

SetDrawingCanDuplicatePartNumber sets DrawingCanDuplicatePartNumber field to given value.

### HasDrawingCanDuplicatePartNumber

`func (o *BTActiveWorkflowInfo) HasDrawingCanDuplicatePartNumber() bool`

HasDrawingCanDuplicatePartNumber returns a boolean if a field has been set.

### GetEnabledActiveMultipleWorkflows

`func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflows() bool`

GetEnabledActiveMultipleWorkflows returns the EnabledActiveMultipleWorkflows field if non-nil, zero value otherwise.

### GetEnabledActiveMultipleWorkflowsOk

`func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflowsOk() (*bool, bool)`

GetEnabledActiveMultipleWorkflowsOk returns a tuple with the EnabledActiveMultipleWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledActiveMultipleWorkflows

`func (o *BTActiveWorkflowInfo) SetEnabledActiveMultipleWorkflows(v bool)`

SetEnabledActiveMultipleWorkflows sets EnabledActiveMultipleWorkflows field to given value.

### HasEnabledActiveMultipleWorkflows

`func (o *BTActiveWorkflowInfo) HasEnabledActiveMultipleWorkflows() bool`

HasEnabledActiveMultipleWorkflows returns a boolean if a field has been set.

### GetObsoletionWorkflow

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflow() BTPublishedWorkflowInfo`

GetObsoletionWorkflow returns the ObsoletionWorkflow field if non-nil, zero value otherwise.

### GetObsoletionWorkflowOk

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowOk() (*BTPublishedWorkflowInfo, bool)`

GetObsoletionWorkflowOk returns a tuple with the ObsoletionWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionWorkflow

`func (o *BTActiveWorkflowInfo) SetObsoletionWorkflow(v BTPublishedWorkflowInfo)`

SetObsoletionWorkflow sets ObsoletionWorkflow field to given value.

### HasObsoletionWorkflow

`func (o *BTActiveWorkflowInfo) HasObsoletionWorkflow() bool`

HasObsoletionWorkflow returns a boolean if a field has been set.

### GetObsoletionWorkflowId

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowId() string`

GetObsoletionWorkflowId returns the ObsoletionWorkflowId field if non-nil, zero value otherwise.

### GetObsoletionWorkflowIdOk

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowIdOk() (*string, bool)`

GetObsoletionWorkflowIdOk returns a tuple with the ObsoletionWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionWorkflowId

`func (o *BTActiveWorkflowInfo) SetObsoletionWorkflowId(v string)`

SetObsoletionWorkflowId sets ObsoletionWorkflowId field to given value.

### HasObsoletionWorkflowId

`func (o *BTActiveWorkflowInfo) HasObsoletionWorkflowId() bool`

HasObsoletionWorkflowId returns a boolean if a field has been set.

### GetPartNumberingSchemeId

`func (o *BTActiveWorkflowInfo) GetPartNumberingSchemeId() string`

GetPartNumberingSchemeId returns the PartNumberingSchemeId field if non-nil, zero value otherwise.

### GetPartNumberingSchemeIdOk

`func (o *BTActiveWorkflowInfo) GetPartNumberingSchemeIdOk() (*string, bool)`

GetPartNumberingSchemeIdOk returns a tuple with the PartNumberingSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumberingSchemeId

`func (o *BTActiveWorkflowInfo) SetPartNumberingSchemeId(v string)`

SetPartNumberingSchemeId sets PartNumberingSchemeId field to given value.

### HasPartNumberingSchemeId

`func (o *BTActiveWorkflowInfo) HasPartNumberingSchemeId() bool`

HasPartNumberingSchemeId returns a boolean if a field has been set.

### GetPickableWorkflows

`func (o *BTActiveWorkflowInfo) GetPickableWorkflows() []BTPublishedWorkflowInfo`

GetPickableWorkflows returns the PickableWorkflows field if non-nil, zero value otherwise.

### GetPickableWorkflowsOk

`func (o *BTActiveWorkflowInfo) GetPickableWorkflowsOk() (*[]BTPublishedWorkflowInfo, bool)`

GetPickableWorkflowsOk returns a tuple with the PickableWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickableWorkflows

`func (o *BTActiveWorkflowInfo) SetPickableWorkflows(v []BTPublishedWorkflowInfo)`

SetPickableWorkflows sets PickableWorkflows field to given value.

### HasPickableWorkflows

`func (o *BTActiveWorkflowInfo) HasPickableWorkflows() bool`

HasPickableWorkflows returns a boolean if a field has been set.

### GetReleaseWorkflow

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflow() BTPublishedWorkflowInfo`

GetReleaseWorkflow returns the ReleaseWorkflow field if non-nil, zero value otherwise.

### GetReleaseWorkflowOk

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowOk() (*BTPublishedWorkflowInfo, bool)`

GetReleaseWorkflowOk returns a tuple with the ReleaseWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseWorkflow

`func (o *BTActiveWorkflowInfo) SetReleaseWorkflow(v BTPublishedWorkflowInfo)`

SetReleaseWorkflow sets ReleaseWorkflow field to given value.

### HasReleaseWorkflow

`func (o *BTActiveWorkflowInfo) HasReleaseWorkflow() bool`

HasReleaseWorkflow returns a boolean if a field has been set.

### GetReleaseWorkflowId

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowId() string`

GetReleaseWorkflowId returns the ReleaseWorkflowId field if non-nil, zero value otherwise.

### GetReleaseWorkflowIdOk

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowIdOk() (*string, bool)`

GetReleaseWorkflowIdOk returns a tuple with the ReleaseWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseWorkflowId

`func (o *BTActiveWorkflowInfo) SetReleaseWorkflowId(v string)`

SetReleaseWorkflowId sets ReleaseWorkflowId field to given value.

### HasReleaseWorkflowId

`func (o *BTActiveWorkflowInfo) HasReleaseWorkflowId() bool`

HasReleaseWorkflowId returns a boolean if a field has been set.

### GetReleaseableApplications

`func (o *BTActiveWorkflowInfo) GetReleaseableApplications() []string`

GetReleaseableApplications returns the ReleaseableApplications field if non-nil, zero value otherwise.

### GetReleaseableApplicationsOk

`func (o *BTActiveWorkflowInfo) GetReleaseableApplicationsOk() (*[]string, bool)`

GetReleaseableApplicationsOk returns a tuple with the ReleaseableApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseableApplications

`func (o *BTActiveWorkflowInfo) SetReleaseableApplications(v []string)`

SetReleaseableApplications sets ReleaseableApplications field to given value.

### HasReleaseableApplications

`func (o *BTActiveWorkflowInfo) HasReleaseableApplications() bool`

HasReleaseableApplications returns a boolean if a field has been set.

### GetStandardContentNumberingSchemeId

`func (o *BTActiveWorkflowInfo) GetStandardContentNumberingSchemeId() string`

GetStandardContentNumberingSchemeId returns the StandardContentNumberingSchemeId field if non-nil, zero value otherwise.

### GetStandardContentNumberingSchemeIdOk

`func (o *BTActiveWorkflowInfo) GetStandardContentNumberingSchemeIdOk() (*string, bool)`

GetStandardContentNumberingSchemeIdOk returns a tuple with the StandardContentNumberingSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentNumberingSchemeId

`func (o *BTActiveWorkflowInfo) SetStandardContentNumberingSchemeId(v string)`

SetStandardContentNumberingSchemeId sets StandardContentNumberingSchemeId field to given value.

### HasStandardContentNumberingSchemeId

`func (o *BTActiveWorkflowInfo) HasStandardContentNumberingSchemeId() bool`

HasStandardContentNumberingSchemeId returns a boolean if a field has been set.

### GetStandardContentUsingAutoNumbering

`func (o *BTActiveWorkflowInfo) GetStandardContentUsingAutoNumbering() bool`

GetStandardContentUsingAutoNumbering returns the StandardContentUsingAutoNumbering field if non-nil, zero value otherwise.

### GetStandardContentUsingAutoNumberingOk

`func (o *BTActiveWorkflowInfo) GetStandardContentUsingAutoNumberingOk() (*bool, bool)`

GetStandardContentUsingAutoNumberingOk returns a tuple with the StandardContentUsingAutoNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentUsingAutoNumbering

`func (o *BTActiveWorkflowInfo) SetStandardContentUsingAutoNumbering(v bool)`

SetStandardContentUsingAutoNumbering sets StandardContentUsingAutoNumbering field to given value.

### HasStandardContentUsingAutoNumbering

`func (o *BTActiveWorkflowInfo) HasStandardContentUsingAutoNumbering() bool`

HasStandardContentUsingAutoNumbering returns a boolean if a field has been set.

### GetStandardContentUsingThirdPartyPartNumbering

`func (o *BTActiveWorkflowInfo) GetStandardContentUsingThirdPartyPartNumbering() bool`

GetStandardContentUsingThirdPartyPartNumbering returns the StandardContentUsingThirdPartyPartNumbering field if non-nil, zero value otherwise.

### GetStandardContentUsingThirdPartyPartNumberingOk

`func (o *BTActiveWorkflowInfo) GetStandardContentUsingThirdPartyPartNumberingOk() (*bool, bool)`

GetStandardContentUsingThirdPartyPartNumberingOk returns a tuple with the StandardContentUsingThirdPartyPartNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentUsingThirdPartyPartNumbering

`func (o *BTActiveWorkflowInfo) SetStandardContentUsingThirdPartyPartNumbering(v bool)`

SetStandardContentUsingThirdPartyPartNumbering sets StandardContentUsingThirdPartyPartNumbering field to given value.

### HasStandardContentUsingThirdPartyPartNumbering

`func (o *BTActiveWorkflowInfo) HasStandardContentUsingThirdPartyPartNumbering() bool`

HasStandardContentUsingThirdPartyPartNumbering returns a boolean if a field has been set.

### GetUsingAutoPartNumbering

`func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumbering() bool`

GetUsingAutoPartNumbering returns the UsingAutoPartNumbering field if non-nil, zero value otherwise.

### GetUsingAutoPartNumberingOk

`func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingOk() (*bool, bool)`

GetUsingAutoPartNumberingOk returns a tuple with the UsingAutoPartNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAutoPartNumbering

`func (o *BTActiveWorkflowInfo) SetUsingAutoPartNumbering(v bool)`

SetUsingAutoPartNumbering sets UsingAutoPartNumbering field to given value.

### HasUsingAutoPartNumbering

`func (o *BTActiveWorkflowInfo) HasUsingAutoPartNumbering() bool`

HasUsingAutoPartNumbering returns a boolean if a field has been set.

### GetUsingManagedWorkflow

`func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflow() bool`

GetUsingManagedWorkflow returns the UsingManagedWorkflow field if non-nil, zero value otherwise.

### GetUsingManagedWorkflowOk

`func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflowOk() (*bool, bool)`

GetUsingManagedWorkflowOk returns a tuple with the UsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingManagedWorkflow

`func (o *BTActiveWorkflowInfo) SetUsingManagedWorkflow(v bool)`

SetUsingManagedWorkflow sets UsingManagedWorkflow field to given value.

### HasUsingManagedWorkflow

`func (o *BTActiveWorkflowInfo) HasUsingManagedWorkflow() bool`

HasUsingManagedWorkflow returns a boolean if a field has been set.

### GetUsingThirdPartyPartNumbering

`func (o *BTActiveWorkflowInfo) GetUsingThirdPartyPartNumbering() bool`

GetUsingThirdPartyPartNumbering returns the UsingThirdPartyPartNumbering field if non-nil, zero value otherwise.

### GetUsingThirdPartyPartNumberingOk

`func (o *BTActiveWorkflowInfo) GetUsingThirdPartyPartNumberingOk() (*bool, bool)`

GetUsingThirdPartyPartNumberingOk returns a tuple with the UsingThirdPartyPartNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingThirdPartyPartNumbering

`func (o *BTActiveWorkflowInfo) SetUsingThirdPartyPartNumbering(v bool)`

SetUsingThirdPartyPartNumbering sets UsingThirdPartyPartNumbering field to given value.

### HasUsingThirdPartyPartNumbering

`func (o *BTActiveWorkflowInfo) HasUsingThirdPartyPartNumbering() bool`

HasUsingThirdPartyPartNumbering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


