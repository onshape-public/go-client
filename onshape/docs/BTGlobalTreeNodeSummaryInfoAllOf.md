# BTGlobalTreeNodeSummaryInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousAccessAllowed** | Pointer to **bool** |  | [optional] 
**AnonymousAllowsExport** | Pointer to **bool** |  | [optional] 
**CanUnshare** | Pointer to **bool** |  | [optional] 
**CreatedWithEducationPlan** | Pointer to **bool** |  | [optional] 
**DefaultElementId** | Pointer to **string** |  | [optional] 
**DefaultVersionGraphMode** | Pointer to [**BTVersionGraphMode**](BTVersionGraphMode.md) |  | [optional] 
**DefaultVersionGraphShowAutoVersions** | Pointer to **bool** |  | [optional] 
**DefaultVersionGraphShowMerges** | Pointer to **bool** |  | [optional] 
**DefaultWorkspace** | Pointer to [**BTWorkspaceInfo**](BTWorkspaceInfo.md) |  | [optional] 
**DocumentLabels** | Pointer to [**[]BTDocumentLabelInfo**](BTDocumentLabelInfo.md) |  | [optional] 
**DocumentType** | Pointer to **int32** |  | [optional] 
**ElementLibrarySummaryInfo** | Pointer to [**[]BTElementLibrarySummaryInfo**](BTElementLibrarySummaryInfo.md) |  | [optional] 
**ForceExportRules** | Pointer to **bool** |  | [optional] 
**HasReleaseRevisionableObjects** | Pointer to **bool** |  | [optional] 
**HasRelevantInsertables** | Pointer to **bool** |  | [optional] 
**IsOrphaned** | Pointer to **bool** |  | [optional] 
**IsUsingManagedWorkflow** | Pointer to **bool** |  | [optional] 
**LikedByCurrentUser** | Pointer to **bool** |  | [optional] 
**Likes** | Pointer to **int64** |  | [optional] 
**NotRevisionManaged** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NumberOfTimesCopied** | Pointer to **int64** |  | [optional] 
**NumberOfTimesReferenced** | Pointer to **int64** |  | [optional] 
**Permission** | Pointer to [**BTOldPermission**](BTOldPermission.md) |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**PublishedVersionId** | Pointer to **string** |  | [optional] 
**RecentVersion** | Pointer to [**BTBaseInfo**](BTBaseInfo.md) |  | [optional] 
**Sequence** | Pointer to **string** |  | [optional] 
**SupportTeamUserAndShared** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**TotalWorkspacesScheduledForUpdate** | Pointer to **int32** |  | [optional] 
**TotalWorkspacesUpdating** | Pointer to **int32** |  | [optional] 
**Trash** | Pointer to **bool** |  | [optional] 
**TrashedAt** | Pointer to **JSONTime** |  | [optional] 
**UserAccountLimitsBreached** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTGlobalTreeNodeSummaryInfoAllOf

`func NewBTGlobalTreeNodeSummaryInfoAllOf() *BTGlobalTreeNodeSummaryInfoAllOf`

NewBTGlobalTreeNodeSummaryInfoAllOf instantiates a new BTGlobalTreeNodeSummaryInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTGlobalTreeNodeSummaryInfoAllOfWithDefaults

`func NewBTGlobalTreeNodeSummaryInfoAllOfWithDefaults() *BTGlobalTreeNodeSummaryInfoAllOf`

NewBTGlobalTreeNodeSummaryInfoAllOfWithDefaults instantiates a new BTGlobalTreeNodeSummaryInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousAccessAllowed

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetAnonymousAccessAllowed() bool`

GetAnonymousAccessAllowed returns the AnonymousAccessAllowed field if non-nil, zero value otherwise.

### GetAnonymousAccessAllowedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetAnonymousAccessAllowedOk() (*bool, bool)`

GetAnonymousAccessAllowedOk returns a tuple with the AnonymousAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessAllowed

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetAnonymousAccessAllowed(v bool)`

SetAnonymousAccessAllowed sets AnonymousAccessAllowed field to given value.

### HasAnonymousAccessAllowed

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasAnonymousAccessAllowed() bool`

HasAnonymousAccessAllowed returns a boolean if a field has been set.

### GetAnonymousAllowsExport

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetAnonymousAllowsExport() bool`

GetAnonymousAllowsExport returns the AnonymousAllowsExport field if non-nil, zero value otherwise.

### GetAnonymousAllowsExportOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetAnonymousAllowsExportOk() (*bool, bool)`

GetAnonymousAllowsExportOk returns a tuple with the AnonymousAllowsExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAllowsExport

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetAnonymousAllowsExport(v bool)`

SetAnonymousAllowsExport sets AnonymousAllowsExport field to given value.

### HasAnonymousAllowsExport

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasAnonymousAllowsExport() bool`

HasAnonymousAllowsExport returns a boolean if a field has been set.

### GetCanUnshare

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetCanUnshare() bool`

GetCanUnshare returns the CanUnshare field if non-nil, zero value otherwise.

### GetCanUnshareOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetCanUnshareOk() (*bool, bool)`

GetCanUnshareOk returns a tuple with the CanUnshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnshare

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetCanUnshare(v bool)`

SetCanUnshare sets CanUnshare field to given value.

### HasCanUnshare

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasCanUnshare() bool`

HasCanUnshare returns a boolean if a field has been set.

### GetCreatedWithEducationPlan

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetCreatedWithEducationPlan() bool`

GetCreatedWithEducationPlan returns the CreatedWithEducationPlan field if non-nil, zero value otherwise.

### GetCreatedWithEducationPlanOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetCreatedWithEducationPlanOk() (*bool, bool)`

GetCreatedWithEducationPlanOk returns a tuple with the CreatedWithEducationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithEducationPlan

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetCreatedWithEducationPlan(v bool)`

SetCreatedWithEducationPlan sets CreatedWithEducationPlan field to given value.

### HasCreatedWithEducationPlan

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasCreatedWithEducationPlan() bool`

HasCreatedWithEducationPlan returns a boolean if a field has been set.

### GetDefaultElementId

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultElementId() string`

GetDefaultElementId returns the DefaultElementId field if non-nil, zero value otherwise.

### GetDefaultElementIdOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultElementIdOk() (*string, bool)`

GetDefaultElementIdOk returns a tuple with the DefaultElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultElementId

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDefaultElementId(v string)`

SetDefaultElementId sets DefaultElementId field to given value.

### HasDefaultElementId

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDefaultElementId() bool`

HasDefaultElementId returns a boolean if a field has been set.

### GetDefaultVersionGraphMode

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultVersionGraphMode() BTVersionGraphMode`

GetDefaultVersionGraphMode returns the DefaultVersionGraphMode field if non-nil, zero value otherwise.

### GetDefaultVersionGraphModeOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultVersionGraphModeOk() (*BTVersionGraphMode, bool)`

GetDefaultVersionGraphModeOk returns a tuple with the DefaultVersionGraphMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionGraphMode

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDefaultVersionGraphMode(v BTVersionGraphMode)`

SetDefaultVersionGraphMode sets DefaultVersionGraphMode field to given value.

### HasDefaultVersionGraphMode

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDefaultVersionGraphMode() bool`

HasDefaultVersionGraphMode returns a boolean if a field has been set.

### GetDefaultVersionGraphShowAutoVersions

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultVersionGraphShowAutoVersions() bool`

GetDefaultVersionGraphShowAutoVersions returns the DefaultVersionGraphShowAutoVersions field if non-nil, zero value otherwise.

### GetDefaultVersionGraphShowAutoVersionsOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultVersionGraphShowAutoVersionsOk() (*bool, bool)`

GetDefaultVersionGraphShowAutoVersionsOk returns a tuple with the DefaultVersionGraphShowAutoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionGraphShowAutoVersions

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDefaultVersionGraphShowAutoVersions(v bool)`

SetDefaultVersionGraphShowAutoVersions sets DefaultVersionGraphShowAutoVersions field to given value.

### HasDefaultVersionGraphShowAutoVersions

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDefaultVersionGraphShowAutoVersions() bool`

HasDefaultVersionGraphShowAutoVersions returns a boolean if a field has been set.

### GetDefaultVersionGraphShowMerges

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultVersionGraphShowMerges() bool`

GetDefaultVersionGraphShowMerges returns the DefaultVersionGraphShowMerges field if non-nil, zero value otherwise.

### GetDefaultVersionGraphShowMergesOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultVersionGraphShowMergesOk() (*bool, bool)`

GetDefaultVersionGraphShowMergesOk returns a tuple with the DefaultVersionGraphShowMerges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionGraphShowMerges

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDefaultVersionGraphShowMerges(v bool)`

SetDefaultVersionGraphShowMerges sets DefaultVersionGraphShowMerges field to given value.

### HasDefaultVersionGraphShowMerges

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDefaultVersionGraphShowMerges() bool`

HasDefaultVersionGraphShowMerges returns a boolean if a field has been set.

### GetDefaultWorkspace

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultWorkspace() BTWorkspaceInfo`

GetDefaultWorkspace returns the DefaultWorkspace field if non-nil, zero value otherwise.

### GetDefaultWorkspaceOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDefaultWorkspaceOk() (*BTWorkspaceInfo, bool)`

GetDefaultWorkspaceOk returns a tuple with the DefaultWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkspace

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDefaultWorkspace(v BTWorkspaceInfo)`

SetDefaultWorkspace sets DefaultWorkspace field to given value.

### HasDefaultWorkspace

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDefaultWorkspace() bool`

HasDefaultWorkspace returns a boolean if a field has been set.

### GetDocumentLabels

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDocumentLabels() []BTDocumentLabelInfo`

GetDocumentLabels returns the DocumentLabels field if non-nil, zero value otherwise.

### GetDocumentLabelsOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDocumentLabelsOk() (*[]BTDocumentLabelInfo, bool)`

GetDocumentLabelsOk returns a tuple with the DocumentLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLabels

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDocumentLabels(v []BTDocumentLabelInfo)`

SetDocumentLabels sets DocumentLabels field to given value.

### HasDocumentLabels

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDocumentLabels() bool`

HasDocumentLabels returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetElementLibrarySummaryInfo

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetElementLibrarySummaryInfo() []BTElementLibrarySummaryInfo`

GetElementLibrarySummaryInfo returns the ElementLibrarySummaryInfo field if non-nil, zero value otherwise.

### GetElementLibrarySummaryInfoOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetElementLibrarySummaryInfoOk() (*[]BTElementLibrarySummaryInfo, bool)`

GetElementLibrarySummaryInfoOk returns a tuple with the ElementLibrarySummaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementLibrarySummaryInfo

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetElementLibrarySummaryInfo(v []BTElementLibrarySummaryInfo)`

SetElementLibrarySummaryInfo sets ElementLibrarySummaryInfo field to given value.

### HasElementLibrarySummaryInfo

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasElementLibrarySummaryInfo() bool`

HasElementLibrarySummaryInfo returns a boolean if a field has been set.

### GetForceExportRules

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetForceExportRules() bool`

GetForceExportRules returns the ForceExportRules field if non-nil, zero value otherwise.

### GetForceExportRulesOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetForceExportRulesOk() (*bool, bool)`

GetForceExportRulesOk returns a tuple with the ForceExportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExportRules

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetForceExportRules(v bool)`

SetForceExportRules sets ForceExportRules field to given value.

### HasForceExportRules

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasForceExportRules() bool`

HasForceExportRules returns a boolean if a field has been set.

### GetHasReleaseRevisionableObjects

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetHasReleaseRevisionableObjects() bool`

GetHasReleaseRevisionableObjects returns the HasReleaseRevisionableObjects field if non-nil, zero value otherwise.

### GetHasReleaseRevisionableObjectsOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetHasReleaseRevisionableObjectsOk() (*bool, bool)`

GetHasReleaseRevisionableObjectsOk returns a tuple with the HasReleaseRevisionableObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReleaseRevisionableObjects

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetHasReleaseRevisionableObjects(v bool)`

SetHasReleaseRevisionableObjects sets HasReleaseRevisionableObjects field to given value.

### HasHasReleaseRevisionableObjects

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasHasReleaseRevisionableObjects() bool`

HasHasReleaseRevisionableObjects returns a boolean if a field has been set.

### GetHasRelevantInsertables

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetHasRelevantInsertables() bool`

GetHasRelevantInsertables returns the HasRelevantInsertables field if non-nil, zero value otherwise.

### GetHasRelevantInsertablesOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetHasRelevantInsertablesOk() (*bool, bool)`

GetHasRelevantInsertablesOk returns a tuple with the HasRelevantInsertables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRelevantInsertables

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetHasRelevantInsertables(v bool)`

SetHasRelevantInsertables sets HasRelevantInsertables field to given value.

### HasHasRelevantInsertables

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasHasRelevantInsertables() bool`

HasHasRelevantInsertables returns a boolean if a field has been set.

### GetIsOrphaned

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetIsOrphaned() bool`

GetIsOrphaned returns the IsOrphaned field if non-nil, zero value otherwise.

### GetIsOrphanedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetIsOrphanedOk() (*bool, bool)`

GetIsOrphanedOk returns a tuple with the IsOrphaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrphaned

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetIsOrphaned(v bool)`

SetIsOrphaned sets IsOrphaned field to given value.

### HasIsOrphaned

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasIsOrphaned() bool`

HasIsOrphaned returns a boolean if a field has been set.

### GetIsUsingManagedWorkflow

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetIsUsingManagedWorkflow() bool`

GetIsUsingManagedWorkflow returns the IsUsingManagedWorkflow field if non-nil, zero value otherwise.

### GetIsUsingManagedWorkflowOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetIsUsingManagedWorkflowOk() (*bool, bool)`

GetIsUsingManagedWorkflowOk returns a tuple with the IsUsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingManagedWorkflow

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetIsUsingManagedWorkflow(v bool)`

SetIsUsingManagedWorkflow sets IsUsingManagedWorkflow field to given value.

### HasIsUsingManagedWorkflow

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasIsUsingManagedWorkflow() bool`

HasIsUsingManagedWorkflow returns a boolean if a field has been set.

### GetLikedByCurrentUser

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetLikedByCurrentUser() bool`

GetLikedByCurrentUser returns the LikedByCurrentUser field if non-nil, zero value otherwise.

### GetLikedByCurrentUserOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetLikedByCurrentUserOk() (*bool, bool)`

GetLikedByCurrentUserOk returns a tuple with the LikedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedByCurrentUser

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetLikedByCurrentUser(v bool)`

SetLikedByCurrentUser sets LikedByCurrentUser field to given value.

### HasLikedByCurrentUser

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasLikedByCurrentUser() bool`

HasLikedByCurrentUser returns a boolean if a field has been set.

### GetLikes

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetLikes(v int64)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetNotes

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumberOfTimesCopied

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNumberOfTimesCopied() int64`

GetNumberOfTimesCopied returns the NumberOfTimesCopied field if non-nil, zero value otherwise.

### GetNumberOfTimesCopiedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNumberOfTimesCopiedOk() (*int64, bool)`

GetNumberOfTimesCopiedOk returns a tuple with the NumberOfTimesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesCopied

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetNumberOfTimesCopied(v int64)`

SetNumberOfTimesCopied sets NumberOfTimesCopied field to given value.

### HasNumberOfTimesCopied

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasNumberOfTimesCopied() bool`

HasNumberOfTimesCopied returns a boolean if a field has been set.

### GetNumberOfTimesReferenced

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNumberOfTimesReferenced() int64`

GetNumberOfTimesReferenced returns the NumberOfTimesReferenced field if non-nil, zero value otherwise.

### GetNumberOfTimesReferencedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetNumberOfTimesReferencedOk() (*int64, bool)`

GetNumberOfTimesReferencedOk returns a tuple with the NumberOfTimesReferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesReferenced

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetNumberOfTimesReferenced(v int64)`

SetNumberOfTimesReferenced sets NumberOfTimesReferenced field to given value.

### HasNumberOfTimesReferenced

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasNumberOfTimesReferenced() bool`

HasNumberOfTimesReferenced returns a boolean if a field has been set.

### GetPermission

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPermission() BTOldPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPermissionOk() (*BTOldPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetPermission(v BTOldPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetPublic

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetRecentVersion

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetRecentVersion() BTBaseInfo`

GetRecentVersion returns the RecentVersion field if non-nil, zero value otherwise.

### GetRecentVersionOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetRecentVersionOk() (*BTBaseInfo, bool)`

GetRecentVersionOk returns a tuple with the RecentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVersion

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetRecentVersion(v BTBaseInfo)`

SetRecentVersion sets RecentVersion field to given value.

### HasRecentVersion

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasRecentVersion() bool`

HasRecentVersion returns a boolean if a field has been set.

### GetSequence

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSupportTeamUserAndShared

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetSupportTeamUserAndShared() bool`

GetSupportTeamUserAndShared returns the SupportTeamUserAndShared field if non-nil, zero value otherwise.

### GetSupportTeamUserAndSharedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetSupportTeamUserAndSharedOk() (*bool, bool)`

GetSupportTeamUserAndSharedOk returns a tuple with the SupportTeamUserAndShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeamUserAndShared

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetSupportTeamUserAndShared(v bool)`

SetSupportTeamUserAndShared sets SupportTeamUserAndShared field to given value.

### HasSupportTeamUserAndShared

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasSupportTeamUserAndShared() bool`

HasSupportTeamUserAndShared returns a boolean if a field has been set.

### GetTags

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTotalWorkspacesScheduledForUpdate

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTotalWorkspacesScheduledForUpdate() int32`

GetTotalWorkspacesScheduledForUpdate returns the TotalWorkspacesScheduledForUpdate field if non-nil, zero value otherwise.

### GetTotalWorkspacesScheduledForUpdateOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTotalWorkspacesScheduledForUpdateOk() (*int32, bool)`

GetTotalWorkspacesScheduledForUpdateOk returns a tuple with the TotalWorkspacesScheduledForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesScheduledForUpdate

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetTotalWorkspacesScheduledForUpdate(v int32)`

SetTotalWorkspacesScheduledForUpdate sets TotalWorkspacesScheduledForUpdate field to given value.

### HasTotalWorkspacesScheduledForUpdate

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasTotalWorkspacesScheduledForUpdate() bool`

HasTotalWorkspacesScheduledForUpdate returns a boolean if a field has been set.

### GetTotalWorkspacesUpdating

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTotalWorkspacesUpdating() int32`

GetTotalWorkspacesUpdating returns the TotalWorkspacesUpdating field if non-nil, zero value otherwise.

### GetTotalWorkspacesUpdatingOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTotalWorkspacesUpdatingOk() (*int32, bool)`

GetTotalWorkspacesUpdatingOk returns a tuple with the TotalWorkspacesUpdating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesUpdating

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetTotalWorkspacesUpdating(v int32)`

SetTotalWorkspacesUpdating sets TotalWorkspacesUpdating field to given value.

### HasTotalWorkspacesUpdating

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasTotalWorkspacesUpdating() bool`

HasTotalWorkspacesUpdating returns a boolean if a field has been set.

### GetTrash

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetTrashedAt

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTrashedAt() JSONTime`

GetTrashedAt returns the TrashedAt field if non-nil, zero value otherwise.

### GetTrashedAtOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetTrashedAtOk() (*JSONTime, bool)`

GetTrashedAtOk returns a tuple with the TrashedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedAt

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetTrashedAt(v JSONTime)`

SetTrashedAt sets TrashedAt field to given value.

### HasTrashedAt

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasTrashedAt() bool`

HasTrashedAt returns a boolean if a field has been set.

### GetUserAccountLimitsBreached

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetUserAccountLimitsBreached() bool`

GetUserAccountLimitsBreached returns the UserAccountLimitsBreached field if non-nil, zero value otherwise.

### GetUserAccountLimitsBreachedOk

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) GetUserAccountLimitsBreachedOk() (*bool, bool)`

GetUserAccountLimitsBreachedOk returns a tuple with the UserAccountLimitsBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountLimitsBreached

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) SetUserAccountLimitsBreached(v bool)`

SetUserAccountLimitsBreached sets UserAccountLimitsBreached field to given value.

### HasUserAccountLimitsBreached

`func (o *BTGlobalTreeNodeSummaryInfoAllOf) HasUserAccountLimitsBreached() bool`

HasUserAccountLimitsBreached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


