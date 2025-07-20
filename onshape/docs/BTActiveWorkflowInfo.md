# BTActiveWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowReleaseItemsFromOtherDocuments** | Pointer to **bool** |  | [optional] 
**CanCurrentUserCreateReleases** | Pointer to **bool** |  | [optional] 
**CanCurrentUserEditStandardContent** | Pointer to **bool** |  | [optional] 
**CanCurrentUserManageWorkflows** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSeeArenaItemLink** | Pointer to **bool** | Deprecated, use canCurrentUserSeePLMItemLink | [optional] 
**CanCurrentUserSeePLMItemLink** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncBomToArena** | Pointer to **bool** | Deprecated, use canCurrentUserSyncBomToPLM | [optional] 
**CanCurrentUserSyncBomToPLM** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncRevisionsToArena** | Pointer to **bool** | Deprecated, use canCurrentUserSyncRevisionsToPLM | [optional] 
**CanCurrentUserSyncRevisionsToPLM** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncStandardContentToArena** | Pointer to **bool** | Deprecated, use canCurrentUserSyncStandardContentToPLM | [optional] 
**CanCurrentUserSyncStandardContentToPLM** | Pointer to **bool** |  | [optional] 
**CanCurrentUserSyncToArena** | Pointer to **bool** | Deprecated, use canCurrentUserSyncToPLM | [optional] 
**CanCurrentUserSyncToPLM** | Pointer to **bool** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DrawingCanDuplicatePartNumber** | Pointer to **bool** |  | [optional] 
**EnabledActiveMultipleWorkflows** | Pointer to **bool** | Deprecated, can be determined by checking if the length of releaseWorkflowInfo.pickableWorkflows &gt; 1 | [optional] 
**HasInactiveCustomWorkflows** | Pointer to **bool** | Deprecated, use hasInactiveCustomWorkflows field on the workflowInfo object | [optional] 
**ObsoletionWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ObsoletionWorkflowId** | Pointer to **string** | Deprecated, use obsoletionWorkflowInfo.workflow.id instead | [optional] 
**ObsoletionWorkflowInfo** | Pointer to [**BTActiveWorkflowTypeInfo**](BTActiveWorkflowTypeInfo.md) |  | [optional] 
**OsCategoryIdToArenaNumberFormatId** | Pointer to **map[string]string** | Deprecated, no current alternative | [optional] 
**PLMIntegrationType** | Pointer to **int32** |  | [optional] 
**PLMName** | Pointer to **string** |  | [optional] 
**PartNumberingSchemeId** | Pointer to **string** |  | [optional] 
**PickableWorkflows** | Pointer to [**[]BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) | Deprecated, use the pickableWorkflows field on the workflowInfo object | [optional] 
**ReleaseWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ReleaseWorkflowId** | Pointer to **string** | Deprecated, use releaseWorkflowInfo.workflow.id instead | [optional] 
**ReleaseWorkflowInfo** | Pointer to [**BTActiveWorkflowTypeInfo**](BTActiveWorkflowTypeInfo.md) |  | [optional] 
**ReleaseableApplications** | Pointer to **[]string** | Deprecated, no current alternative | [optional] 
**StandardContentNumberingSchemeId** | Pointer to **string** |  | [optional] 
**StandardContentUsingAutoNumbering** | Pointer to **bool** |  | [optional] 
**StandardContentUsingThirdPartyPartNumbering** | Pointer to **bool** |  | [optional] 
**TaskWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**TaskWorkflowInfo** | Pointer to [**BTActiveWorkflowTypeInfo**](BTActiveWorkflowTypeInfo.md) |  | [optional] 
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

### GetCanCurrentUserSeePLMItemLink

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSeePLMItemLink() bool`

GetCanCurrentUserSeePLMItemLink returns the CanCurrentUserSeePLMItemLink field if non-nil, zero value otherwise.

### GetCanCurrentUserSeePLMItemLinkOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSeePLMItemLinkOk() (*bool, bool)`

GetCanCurrentUserSeePLMItemLinkOk returns a tuple with the CanCurrentUserSeePLMItemLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSeePLMItemLink

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSeePLMItemLink(v bool)`

SetCanCurrentUserSeePLMItemLink sets CanCurrentUserSeePLMItemLink field to given value.

### HasCanCurrentUserSeePLMItemLink

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSeePLMItemLink() bool`

HasCanCurrentUserSeePLMItemLink returns a boolean if a field has been set.

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

### GetCanCurrentUserSyncBomToPLM

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncBomToPLM() bool`

GetCanCurrentUserSyncBomToPLM returns the CanCurrentUserSyncBomToPLM field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncBomToPLMOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncBomToPLMOk() (*bool, bool)`

GetCanCurrentUserSyncBomToPLMOk returns a tuple with the CanCurrentUserSyncBomToPLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncBomToPLM

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncBomToPLM(v bool)`

SetCanCurrentUserSyncBomToPLM sets CanCurrentUserSyncBomToPLM field to given value.

### HasCanCurrentUserSyncBomToPLM

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncBomToPLM() bool`

HasCanCurrentUserSyncBomToPLM returns a boolean if a field has been set.

### GetCanCurrentUserSyncRevisionsToArena

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncRevisionsToArena() bool`

GetCanCurrentUserSyncRevisionsToArena returns the CanCurrentUserSyncRevisionsToArena field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncRevisionsToArenaOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncRevisionsToArenaOk() (*bool, bool)`

GetCanCurrentUserSyncRevisionsToArenaOk returns a tuple with the CanCurrentUserSyncRevisionsToArena field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncRevisionsToArena

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncRevisionsToArena(v bool)`

SetCanCurrentUserSyncRevisionsToArena sets CanCurrentUserSyncRevisionsToArena field to given value.

### HasCanCurrentUserSyncRevisionsToArena

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncRevisionsToArena() bool`

HasCanCurrentUserSyncRevisionsToArena returns a boolean if a field has been set.

### GetCanCurrentUserSyncRevisionsToPLM

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncRevisionsToPLM() bool`

GetCanCurrentUserSyncRevisionsToPLM returns the CanCurrentUserSyncRevisionsToPLM field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncRevisionsToPLMOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncRevisionsToPLMOk() (*bool, bool)`

GetCanCurrentUserSyncRevisionsToPLMOk returns a tuple with the CanCurrentUserSyncRevisionsToPLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncRevisionsToPLM

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncRevisionsToPLM(v bool)`

SetCanCurrentUserSyncRevisionsToPLM sets CanCurrentUserSyncRevisionsToPLM field to given value.

### HasCanCurrentUserSyncRevisionsToPLM

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncRevisionsToPLM() bool`

HasCanCurrentUserSyncRevisionsToPLM returns a boolean if a field has been set.

### GetCanCurrentUserSyncStandardContentToArena

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncStandardContentToArena() bool`

GetCanCurrentUserSyncStandardContentToArena returns the CanCurrentUserSyncStandardContentToArena field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncStandardContentToArenaOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncStandardContentToArenaOk() (*bool, bool)`

GetCanCurrentUserSyncStandardContentToArenaOk returns a tuple with the CanCurrentUserSyncStandardContentToArena field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncStandardContentToArena

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncStandardContentToArena(v bool)`

SetCanCurrentUserSyncStandardContentToArena sets CanCurrentUserSyncStandardContentToArena field to given value.

### HasCanCurrentUserSyncStandardContentToArena

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncStandardContentToArena() bool`

HasCanCurrentUserSyncStandardContentToArena returns a boolean if a field has been set.

### GetCanCurrentUserSyncStandardContentToPLM

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncStandardContentToPLM() bool`

GetCanCurrentUserSyncStandardContentToPLM returns the CanCurrentUserSyncStandardContentToPLM field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncStandardContentToPLMOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncStandardContentToPLMOk() (*bool, bool)`

GetCanCurrentUserSyncStandardContentToPLMOk returns a tuple with the CanCurrentUserSyncStandardContentToPLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncStandardContentToPLM

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncStandardContentToPLM(v bool)`

SetCanCurrentUserSyncStandardContentToPLM sets CanCurrentUserSyncStandardContentToPLM field to given value.

### HasCanCurrentUserSyncStandardContentToPLM

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncStandardContentToPLM() bool`

HasCanCurrentUserSyncStandardContentToPLM returns a boolean if a field has been set.

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

### GetCanCurrentUserSyncToPLM

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncToPLM() bool`

GetCanCurrentUserSyncToPLM returns the CanCurrentUserSyncToPLM field if non-nil, zero value otherwise.

### GetCanCurrentUserSyncToPLMOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserSyncToPLMOk() (*bool, bool)`

GetCanCurrentUserSyncToPLMOk returns a tuple with the CanCurrentUserSyncToPLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserSyncToPLM

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserSyncToPLM(v bool)`

SetCanCurrentUserSyncToPLM sets CanCurrentUserSyncToPLM field to given value.

### HasCanCurrentUserSyncToPLM

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserSyncToPLM() bool`

HasCanCurrentUserSyncToPLM returns a boolean if a field has been set.

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

### GetHasInactiveCustomWorkflows

`func (o *BTActiveWorkflowInfo) GetHasInactiveCustomWorkflows() bool`

GetHasInactiveCustomWorkflows returns the HasInactiveCustomWorkflows field if non-nil, zero value otherwise.

### GetHasInactiveCustomWorkflowsOk

`func (o *BTActiveWorkflowInfo) GetHasInactiveCustomWorkflowsOk() (*bool, bool)`

GetHasInactiveCustomWorkflowsOk returns a tuple with the HasInactiveCustomWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInactiveCustomWorkflows

`func (o *BTActiveWorkflowInfo) SetHasInactiveCustomWorkflows(v bool)`

SetHasInactiveCustomWorkflows sets HasInactiveCustomWorkflows field to given value.

### HasHasInactiveCustomWorkflows

`func (o *BTActiveWorkflowInfo) HasHasInactiveCustomWorkflows() bool`

HasHasInactiveCustomWorkflows returns a boolean if a field has been set.

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

### GetObsoletionWorkflowInfo

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowInfo() BTActiveWorkflowTypeInfo`

GetObsoletionWorkflowInfo returns the ObsoletionWorkflowInfo field if non-nil, zero value otherwise.

### GetObsoletionWorkflowInfoOk

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowInfoOk() (*BTActiveWorkflowTypeInfo, bool)`

GetObsoletionWorkflowInfoOk returns a tuple with the ObsoletionWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionWorkflowInfo

`func (o *BTActiveWorkflowInfo) SetObsoletionWorkflowInfo(v BTActiveWorkflowTypeInfo)`

SetObsoletionWorkflowInfo sets ObsoletionWorkflowInfo field to given value.

### HasObsoletionWorkflowInfo

`func (o *BTActiveWorkflowInfo) HasObsoletionWorkflowInfo() bool`

HasObsoletionWorkflowInfo returns a boolean if a field has been set.

### GetOsCategoryIdToArenaNumberFormatId

`func (o *BTActiveWorkflowInfo) GetOsCategoryIdToArenaNumberFormatId() map[string]string`

GetOsCategoryIdToArenaNumberFormatId returns the OsCategoryIdToArenaNumberFormatId field if non-nil, zero value otherwise.

### GetOsCategoryIdToArenaNumberFormatIdOk

`func (o *BTActiveWorkflowInfo) GetOsCategoryIdToArenaNumberFormatIdOk() (*map[string]string, bool)`

GetOsCategoryIdToArenaNumberFormatIdOk returns a tuple with the OsCategoryIdToArenaNumberFormatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCategoryIdToArenaNumberFormatId

`func (o *BTActiveWorkflowInfo) SetOsCategoryIdToArenaNumberFormatId(v map[string]string)`

SetOsCategoryIdToArenaNumberFormatId sets OsCategoryIdToArenaNumberFormatId field to given value.

### HasOsCategoryIdToArenaNumberFormatId

`func (o *BTActiveWorkflowInfo) HasOsCategoryIdToArenaNumberFormatId() bool`

HasOsCategoryIdToArenaNumberFormatId returns a boolean if a field has been set.

### GetPLMIntegrationType

`func (o *BTActiveWorkflowInfo) GetPLMIntegrationType() int32`

GetPLMIntegrationType returns the PLMIntegrationType field if non-nil, zero value otherwise.

### GetPLMIntegrationTypeOk

`func (o *BTActiveWorkflowInfo) GetPLMIntegrationTypeOk() (*int32, bool)`

GetPLMIntegrationTypeOk returns a tuple with the PLMIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPLMIntegrationType

`func (o *BTActiveWorkflowInfo) SetPLMIntegrationType(v int32)`

SetPLMIntegrationType sets PLMIntegrationType field to given value.

### HasPLMIntegrationType

`func (o *BTActiveWorkflowInfo) HasPLMIntegrationType() bool`

HasPLMIntegrationType returns a boolean if a field has been set.

### GetPLMName

`func (o *BTActiveWorkflowInfo) GetPLMName() string`

GetPLMName returns the PLMName field if non-nil, zero value otherwise.

### GetPLMNameOk

`func (o *BTActiveWorkflowInfo) GetPLMNameOk() (*string, bool)`

GetPLMNameOk returns a tuple with the PLMName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPLMName

`func (o *BTActiveWorkflowInfo) SetPLMName(v string)`

SetPLMName sets PLMName field to given value.

### HasPLMName

`func (o *BTActiveWorkflowInfo) HasPLMName() bool`

HasPLMName returns a boolean if a field has been set.

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

### GetReleaseWorkflowInfo

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowInfo() BTActiveWorkflowTypeInfo`

GetReleaseWorkflowInfo returns the ReleaseWorkflowInfo field if non-nil, zero value otherwise.

### GetReleaseWorkflowInfoOk

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowInfoOk() (*BTActiveWorkflowTypeInfo, bool)`

GetReleaseWorkflowInfoOk returns a tuple with the ReleaseWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseWorkflowInfo

`func (o *BTActiveWorkflowInfo) SetReleaseWorkflowInfo(v BTActiveWorkflowTypeInfo)`

SetReleaseWorkflowInfo sets ReleaseWorkflowInfo field to given value.

### HasReleaseWorkflowInfo

`func (o *BTActiveWorkflowInfo) HasReleaseWorkflowInfo() bool`

HasReleaseWorkflowInfo returns a boolean if a field has been set.

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

### GetTaskWorkflow

`func (o *BTActiveWorkflowInfo) GetTaskWorkflow() BTPublishedWorkflowInfo`

GetTaskWorkflow returns the TaskWorkflow field if non-nil, zero value otherwise.

### GetTaskWorkflowOk

`func (o *BTActiveWorkflowInfo) GetTaskWorkflowOk() (*BTPublishedWorkflowInfo, bool)`

GetTaskWorkflowOk returns a tuple with the TaskWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflow

`func (o *BTActiveWorkflowInfo) SetTaskWorkflow(v BTPublishedWorkflowInfo)`

SetTaskWorkflow sets TaskWorkflow field to given value.

### HasTaskWorkflow

`func (o *BTActiveWorkflowInfo) HasTaskWorkflow() bool`

HasTaskWorkflow returns a boolean if a field has been set.

### GetTaskWorkflowInfo

`func (o *BTActiveWorkflowInfo) GetTaskWorkflowInfo() BTActiveWorkflowTypeInfo`

GetTaskWorkflowInfo returns the TaskWorkflowInfo field if non-nil, zero value otherwise.

### GetTaskWorkflowInfoOk

`func (o *BTActiveWorkflowInfo) GetTaskWorkflowInfoOk() (*BTActiveWorkflowTypeInfo, bool)`

GetTaskWorkflowInfoOk returns a tuple with the TaskWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowInfo

`func (o *BTActiveWorkflowInfo) SetTaskWorkflowInfo(v BTActiveWorkflowTypeInfo)`

SetTaskWorkflowInfo sets TaskWorkflowInfo field to given value.

### HasTaskWorkflowInfo

`func (o *BTActiveWorkflowInfo) HasTaskWorkflowInfo() bool`

HasTaskWorkflowInfo returns a boolean if a field has been set.

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


