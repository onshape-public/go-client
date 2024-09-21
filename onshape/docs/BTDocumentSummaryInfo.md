# BTDocumentSummaryInfo

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
**ParentId** | Pointer to **string** |  | [optional] 
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

### NewBTDocumentSummaryInfo

`func NewBTDocumentSummaryInfo() *BTDocumentSummaryInfo`

NewBTDocumentSummaryInfo instantiates a new BTDocumentSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentSummaryInfoWithDefaults

`func NewBTDocumentSummaryInfoWithDefaults() *BTDocumentSummaryInfo`

NewBTDocumentSummaryInfoWithDefaults instantiates a new BTDocumentSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousAccessAllowed

`func (o *BTDocumentSummaryInfo) GetAnonymousAccessAllowed() bool`

GetAnonymousAccessAllowed returns the AnonymousAccessAllowed field if non-nil, zero value otherwise.

### GetAnonymousAccessAllowedOk

`func (o *BTDocumentSummaryInfo) GetAnonymousAccessAllowedOk() (*bool, bool)`

GetAnonymousAccessAllowedOk returns a tuple with the AnonymousAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessAllowed

`func (o *BTDocumentSummaryInfo) SetAnonymousAccessAllowed(v bool)`

SetAnonymousAccessAllowed sets AnonymousAccessAllowed field to given value.

### HasAnonymousAccessAllowed

`func (o *BTDocumentSummaryInfo) HasAnonymousAccessAllowed() bool`

HasAnonymousAccessAllowed returns a boolean if a field has been set.

### GetAnonymousAllowsExport

`func (o *BTDocumentSummaryInfo) GetAnonymousAllowsExport() bool`

GetAnonymousAllowsExport returns the AnonymousAllowsExport field if non-nil, zero value otherwise.

### GetAnonymousAllowsExportOk

`func (o *BTDocumentSummaryInfo) GetAnonymousAllowsExportOk() (*bool, bool)`

GetAnonymousAllowsExportOk returns a tuple with the AnonymousAllowsExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAllowsExport

`func (o *BTDocumentSummaryInfo) SetAnonymousAllowsExport(v bool)`

SetAnonymousAllowsExport sets AnonymousAllowsExport field to given value.

### HasAnonymousAllowsExport

`func (o *BTDocumentSummaryInfo) HasAnonymousAllowsExport() bool`

HasAnonymousAllowsExport returns a boolean if a field has been set.

### GetCanUnshare

`func (o *BTDocumentSummaryInfo) GetCanUnshare() bool`

GetCanUnshare returns the CanUnshare field if non-nil, zero value otherwise.

### GetCanUnshareOk

`func (o *BTDocumentSummaryInfo) GetCanUnshareOk() (*bool, bool)`

GetCanUnshareOk returns a tuple with the CanUnshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnshare

`func (o *BTDocumentSummaryInfo) SetCanUnshare(v bool)`

SetCanUnshare sets CanUnshare field to given value.

### HasCanUnshare

`func (o *BTDocumentSummaryInfo) HasCanUnshare() bool`

HasCanUnshare returns a boolean if a field has been set.

### GetCreatedWithEducationPlan

`func (o *BTDocumentSummaryInfo) GetCreatedWithEducationPlan() bool`

GetCreatedWithEducationPlan returns the CreatedWithEducationPlan field if non-nil, zero value otherwise.

### GetCreatedWithEducationPlanOk

`func (o *BTDocumentSummaryInfo) GetCreatedWithEducationPlanOk() (*bool, bool)`

GetCreatedWithEducationPlanOk returns a tuple with the CreatedWithEducationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithEducationPlan

`func (o *BTDocumentSummaryInfo) SetCreatedWithEducationPlan(v bool)`

SetCreatedWithEducationPlan sets CreatedWithEducationPlan field to given value.

### HasCreatedWithEducationPlan

`func (o *BTDocumentSummaryInfo) HasCreatedWithEducationPlan() bool`

HasCreatedWithEducationPlan returns a boolean if a field has been set.

### GetDefaultElementId

`func (o *BTDocumentSummaryInfo) GetDefaultElementId() string`

GetDefaultElementId returns the DefaultElementId field if non-nil, zero value otherwise.

### GetDefaultElementIdOk

`func (o *BTDocumentSummaryInfo) GetDefaultElementIdOk() (*string, bool)`

GetDefaultElementIdOk returns a tuple with the DefaultElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultElementId

`func (o *BTDocumentSummaryInfo) SetDefaultElementId(v string)`

SetDefaultElementId sets DefaultElementId field to given value.

### HasDefaultElementId

`func (o *BTDocumentSummaryInfo) HasDefaultElementId() bool`

HasDefaultElementId returns a boolean if a field has been set.

### GetDefaultVersionGraphMode

`func (o *BTDocumentSummaryInfo) GetDefaultVersionGraphMode() BTVersionGraphMode`

GetDefaultVersionGraphMode returns the DefaultVersionGraphMode field if non-nil, zero value otherwise.

### GetDefaultVersionGraphModeOk

`func (o *BTDocumentSummaryInfo) GetDefaultVersionGraphModeOk() (*BTVersionGraphMode, bool)`

GetDefaultVersionGraphModeOk returns a tuple with the DefaultVersionGraphMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionGraphMode

`func (o *BTDocumentSummaryInfo) SetDefaultVersionGraphMode(v BTVersionGraphMode)`

SetDefaultVersionGraphMode sets DefaultVersionGraphMode field to given value.

### HasDefaultVersionGraphMode

`func (o *BTDocumentSummaryInfo) HasDefaultVersionGraphMode() bool`

HasDefaultVersionGraphMode returns a boolean if a field has been set.

### GetDefaultVersionGraphShowAutoVersions

`func (o *BTDocumentSummaryInfo) GetDefaultVersionGraphShowAutoVersions() bool`

GetDefaultVersionGraphShowAutoVersions returns the DefaultVersionGraphShowAutoVersions field if non-nil, zero value otherwise.

### GetDefaultVersionGraphShowAutoVersionsOk

`func (o *BTDocumentSummaryInfo) GetDefaultVersionGraphShowAutoVersionsOk() (*bool, bool)`

GetDefaultVersionGraphShowAutoVersionsOk returns a tuple with the DefaultVersionGraphShowAutoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionGraphShowAutoVersions

`func (o *BTDocumentSummaryInfo) SetDefaultVersionGraphShowAutoVersions(v bool)`

SetDefaultVersionGraphShowAutoVersions sets DefaultVersionGraphShowAutoVersions field to given value.

### HasDefaultVersionGraphShowAutoVersions

`func (o *BTDocumentSummaryInfo) HasDefaultVersionGraphShowAutoVersions() bool`

HasDefaultVersionGraphShowAutoVersions returns a boolean if a field has been set.

### GetDefaultVersionGraphShowMerges

`func (o *BTDocumentSummaryInfo) GetDefaultVersionGraphShowMerges() bool`

GetDefaultVersionGraphShowMerges returns the DefaultVersionGraphShowMerges field if non-nil, zero value otherwise.

### GetDefaultVersionGraphShowMergesOk

`func (o *BTDocumentSummaryInfo) GetDefaultVersionGraphShowMergesOk() (*bool, bool)`

GetDefaultVersionGraphShowMergesOk returns a tuple with the DefaultVersionGraphShowMerges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionGraphShowMerges

`func (o *BTDocumentSummaryInfo) SetDefaultVersionGraphShowMerges(v bool)`

SetDefaultVersionGraphShowMerges sets DefaultVersionGraphShowMerges field to given value.

### HasDefaultVersionGraphShowMerges

`func (o *BTDocumentSummaryInfo) HasDefaultVersionGraphShowMerges() bool`

HasDefaultVersionGraphShowMerges returns a boolean if a field has been set.

### GetDefaultWorkspace

`func (o *BTDocumentSummaryInfo) GetDefaultWorkspace() BTWorkspaceInfo`

GetDefaultWorkspace returns the DefaultWorkspace field if non-nil, zero value otherwise.

### GetDefaultWorkspaceOk

`func (o *BTDocumentSummaryInfo) GetDefaultWorkspaceOk() (*BTWorkspaceInfo, bool)`

GetDefaultWorkspaceOk returns a tuple with the DefaultWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkspace

`func (o *BTDocumentSummaryInfo) SetDefaultWorkspace(v BTWorkspaceInfo)`

SetDefaultWorkspace sets DefaultWorkspace field to given value.

### HasDefaultWorkspace

`func (o *BTDocumentSummaryInfo) HasDefaultWorkspace() bool`

HasDefaultWorkspace returns a boolean if a field has been set.

### GetDocumentLabels

`func (o *BTDocumentSummaryInfo) GetDocumentLabels() []BTDocumentLabelInfo`

GetDocumentLabels returns the DocumentLabels field if non-nil, zero value otherwise.

### GetDocumentLabelsOk

`func (o *BTDocumentSummaryInfo) GetDocumentLabelsOk() (*[]BTDocumentLabelInfo, bool)`

GetDocumentLabelsOk returns a tuple with the DocumentLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLabels

`func (o *BTDocumentSummaryInfo) SetDocumentLabels(v []BTDocumentLabelInfo)`

SetDocumentLabels sets DocumentLabels field to given value.

### HasDocumentLabels

`func (o *BTDocumentSummaryInfo) HasDocumentLabels() bool`

HasDocumentLabels returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTDocumentSummaryInfo) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTDocumentSummaryInfo) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTDocumentSummaryInfo) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTDocumentSummaryInfo) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetForceExportRules

`func (o *BTDocumentSummaryInfo) GetForceExportRules() bool`

GetForceExportRules returns the ForceExportRules field if non-nil, zero value otherwise.

### GetForceExportRulesOk

`func (o *BTDocumentSummaryInfo) GetForceExportRulesOk() (*bool, bool)`

GetForceExportRulesOk returns a tuple with the ForceExportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExportRules

`func (o *BTDocumentSummaryInfo) SetForceExportRules(v bool)`

SetForceExportRules sets ForceExportRules field to given value.

### HasForceExportRules

`func (o *BTDocumentSummaryInfo) HasForceExportRules() bool`

HasForceExportRules returns a boolean if a field has been set.

### GetHasReleaseRevisionableObjects

`func (o *BTDocumentSummaryInfo) GetHasReleaseRevisionableObjects() bool`

GetHasReleaseRevisionableObjects returns the HasReleaseRevisionableObjects field if non-nil, zero value otherwise.

### GetHasReleaseRevisionableObjectsOk

`func (o *BTDocumentSummaryInfo) GetHasReleaseRevisionableObjectsOk() (*bool, bool)`

GetHasReleaseRevisionableObjectsOk returns a tuple with the HasReleaseRevisionableObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReleaseRevisionableObjects

`func (o *BTDocumentSummaryInfo) SetHasReleaseRevisionableObjects(v bool)`

SetHasReleaseRevisionableObjects sets HasReleaseRevisionableObjects field to given value.

### HasHasReleaseRevisionableObjects

`func (o *BTDocumentSummaryInfo) HasHasReleaseRevisionableObjects() bool`

HasHasReleaseRevisionableObjects returns a boolean if a field has been set.

### GetHasRelevantInsertables

`func (o *BTDocumentSummaryInfo) GetHasRelevantInsertables() bool`

GetHasRelevantInsertables returns the HasRelevantInsertables field if non-nil, zero value otherwise.

### GetHasRelevantInsertablesOk

`func (o *BTDocumentSummaryInfo) GetHasRelevantInsertablesOk() (*bool, bool)`

GetHasRelevantInsertablesOk returns a tuple with the HasRelevantInsertables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRelevantInsertables

`func (o *BTDocumentSummaryInfo) SetHasRelevantInsertables(v bool)`

SetHasRelevantInsertables sets HasRelevantInsertables field to given value.

### HasHasRelevantInsertables

`func (o *BTDocumentSummaryInfo) HasHasRelevantInsertables() bool`

HasHasRelevantInsertables returns a boolean if a field has been set.

### GetIsOrphaned

`func (o *BTDocumentSummaryInfo) GetIsOrphaned() bool`

GetIsOrphaned returns the IsOrphaned field if non-nil, zero value otherwise.

### GetIsOrphanedOk

`func (o *BTDocumentSummaryInfo) GetIsOrphanedOk() (*bool, bool)`

GetIsOrphanedOk returns a tuple with the IsOrphaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrphaned

`func (o *BTDocumentSummaryInfo) SetIsOrphaned(v bool)`

SetIsOrphaned sets IsOrphaned field to given value.

### HasIsOrphaned

`func (o *BTDocumentSummaryInfo) HasIsOrphaned() bool`

HasIsOrphaned returns a boolean if a field has been set.

### GetIsUsingManagedWorkflow

`func (o *BTDocumentSummaryInfo) GetIsUsingManagedWorkflow() bool`

GetIsUsingManagedWorkflow returns the IsUsingManagedWorkflow field if non-nil, zero value otherwise.

### GetIsUsingManagedWorkflowOk

`func (o *BTDocumentSummaryInfo) GetIsUsingManagedWorkflowOk() (*bool, bool)`

GetIsUsingManagedWorkflowOk returns a tuple with the IsUsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingManagedWorkflow

`func (o *BTDocumentSummaryInfo) SetIsUsingManagedWorkflow(v bool)`

SetIsUsingManagedWorkflow sets IsUsingManagedWorkflow field to given value.

### HasIsUsingManagedWorkflow

`func (o *BTDocumentSummaryInfo) HasIsUsingManagedWorkflow() bool`

HasIsUsingManagedWorkflow returns a boolean if a field has been set.

### GetLikedByCurrentUser

`func (o *BTDocumentSummaryInfo) GetLikedByCurrentUser() bool`

GetLikedByCurrentUser returns the LikedByCurrentUser field if non-nil, zero value otherwise.

### GetLikedByCurrentUserOk

`func (o *BTDocumentSummaryInfo) GetLikedByCurrentUserOk() (*bool, bool)`

GetLikedByCurrentUserOk returns a tuple with the LikedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedByCurrentUser

`func (o *BTDocumentSummaryInfo) SetLikedByCurrentUser(v bool)`

SetLikedByCurrentUser sets LikedByCurrentUser field to given value.

### HasLikedByCurrentUser

`func (o *BTDocumentSummaryInfo) HasLikedByCurrentUser() bool`

HasLikedByCurrentUser returns a boolean if a field has been set.

### GetLikes

`func (o *BTDocumentSummaryInfo) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BTDocumentSummaryInfo) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BTDocumentSummaryInfo) SetLikes(v int64)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BTDocumentSummaryInfo) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTDocumentSummaryInfo) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTDocumentSummaryInfo) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTDocumentSummaryInfo) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTDocumentSummaryInfo) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetNotes

`func (o *BTDocumentSummaryInfo) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BTDocumentSummaryInfo) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BTDocumentSummaryInfo) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BTDocumentSummaryInfo) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumberOfTimesCopied

`func (o *BTDocumentSummaryInfo) GetNumberOfTimesCopied() int64`

GetNumberOfTimesCopied returns the NumberOfTimesCopied field if non-nil, zero value otherwise.

### GetNumberOfTimesCopiedOk

`func (o *BTDocumentSummaryInfo) GetNumberOfTimesCopiedOk() (*int64, bool)`

GetNumberOfTimesCopiedOk returns a tuple with the NumberOfTimesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesCopied

`func (o *BTDocumentSummaryInfo) SetNumberOfTimesCopied(v int64)`

SetNumberOfTimesCopied sets NumberOfTimesCopied field to given value.

### HasNumberOfTimesCopied

`func (o *BTDocumentSummaryInfo) HasNumberOfTimesCopied() bool`

HasNumberOfTimesCopied returns a boolean if a field has been set.

### GetNumberOfTimesReferenced

`func (o *BTDocumentSummaryInfo) GetNumberOfTimesReferenced() int64`

GetNumberOfTimesReferenced returns the NumberOfTimesReferenced field if non-nil, zero value otherwise.

### GetNumberOfTimesReferencedOk

`func (o *BTDocumentSummaryInfo) GetNumberOfTimesReferencedOk() (*int64, bool)`

GetNumberOfTimesReferencedOk returns a tuple with the NumberOfTimesReferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesReferenced

`func (o *BTDocumentSummaryInfo) SetNumberOfTimesReferenced(v int64)`

SetNumberOfTimesReferenced sets NumberOfTimesReferenced field to given value.

### HasNumberOfTimesReferenced

`func (o *BTDocumentSummaryInfo) HasNumberOfTimesReferenced() bool`

HasNumberOfTimesReferenced returns a boolean if a field has been set.

### GetParentId

`func (o *BTDocumentSummaryInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTDocumentSummaryInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTDocumentSummaryInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTDocumentSummaryInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermission

`func (o *BTDocumentSummaryInfo) GetPermission() BTOldPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTDocumentSummaryInfo) GetPermissionOk() (*BTOldPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTDocumentSummaryInfo) SetPermission(v BTOldPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTDocumentSummaryInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTDocumentSummaryInfo) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTDocumentSummaryInfo) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTDocumentSummaryInfo) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTDocumentSummaryInfo) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetPublic

`func (o *BTDocumentSummaryInfo) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTDocumentSummaryInfo) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTDocumentSummaryInfo) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTDocumentSummaryInfo) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *BTDocumentSummaryInfo) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *BTDocumentSummaryInfo) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *BTDocumentSummaryInfo) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *BTDocumentSummaryInfo) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetRecentVersion

`func (o *BTDocumentSummaryInfo) GetRecentVersion() BTBaseInfo`

GetRecentVersion returns the RecentVersion field if non-nil, zero value otherwise.

### GetRecentVersionOk

`func (o *BTDocumentSummaryInfo) GetRecentVersionOk() (*BTBaseInfo, bool)`

GetRecentVersionOk returns a tuple with the RecentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVersion

`func (o *BTDocumentSummaryInfo) SetRecentVersion(v BTBaseInfo)`

SetRecentVersion sets RecentVersion field to given value.

### HasRecentVersion

`func (o *BTDocumentSummaryInfo) HasRecentVersion() bool`

HasRecentVersion returns a boolean if a field has been set.

### GetSequence

`func (o *BTDocumentSummaryInfo) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BTDocumentSummaryInfo) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BTDocumentSummaryInfo) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BTDocumentSummaryInfo) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSupportTeamUserAndShared

`func (o *BTDocumentSummaryInfo) GetSupportTeamUserAndShared() bool`

GetSupportTeamUserAndShared returns the SupportTeamUserAndShared field if non-nil, zero value otherwise.

### GetSupportTeamUserAndSharedOk

`func (o *BTDocumentSummaryInfo) GetSupportTeamUserAndSharedOk() (*bool, bool)`

GetSupportTeamUserAndSharedOk returns a tuple with the SupportTeamUserAndShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeamUserAndShared

`func (o *BTDocumentSummaryInfo) SetSupportTeamUserAndShared(v bool)`

SetSupportTeamUserAndShared sets SupportTeamUserAndShared field to given value.

### HasSupportTeamUserAndShared

`func (o *BTDocumentSummaryInfo) HasSupportTeamUserAndShared() bool`

HasSupportTeamUserAndShared returns a boolean if a field has been set.

### GetTags

`func (o *BTDocumentSummaryInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTDocumentSummaryInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTDocumentSummaryInfo) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTDocumentSummaryInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTDocumentSummaryInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTDocumentSummaryInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTDocumentSummaryInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTDocumentSummaryInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentSummaryInfo) GetTotalWorkspacesScheduledForUpdate() int32`

GetTotalWorkspacesScheduledForUpdate returns the TotalWorkspacesScheduledForUpdate field if non-nil, zero value otherwise.

### GetTotalWorkspacesScheduledForUpdateOk

`func (o *BTDocumentSummaryInfo) GetTotalWorkspacesScheduledForUpdateOk() (*int32, bool)`

GetTotalWorkspacesScheduledForUpdateOk returns a tuple with the TotalWorkspacesScheduledForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentSummaryInfo) SetTotalWorkspacesScheduledForUpdate(v int32)`

SetTotalWorkspacesScheduledForUpdate sets TotalWorkspacesScheduledForUpdate field to given value.

### HasTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentSummaryInfo) HasTotalWorkspacesScheduledForUpdate() bool`

HasTotalWorkspacesScheduledForUpdate returns a boolean if a field has been set.

### GetTotalWorkspacesUpdating

`func (o *BTDocumentSummaryInfo) GetTotalWorkspacesUpdating() int32`

GetTotalWorkspacesUpdating returns the TotalWorkspacesUpdating field if non-nil, zero value otherwise.

### GetTotalWorkspacesUpdatingOk

`func (o *BTDocumentSummaryInfo) GetTotalWorkspacesUpdatingOk() (*int32, bool)`

GetTotalWorkspacesUpdatingOk returns a tuple with the TotalWorkspacesUpdating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesUpdating

`func (o *BTDocumentSummaryInfo) SetTotalWorkspacesUpdating(v int32)`

SetTotalWorkspacesUpdating sets TotalWorkspacesUpdating field to given value.

### HasTotalWorkspacesUpdating

`func (o *BTDocumentSummaryInfo) HasTotalWorkspacesUpdating() bool`

HasTotalWorkspacesUpdating returns a boolean if a field has been set.

### GetTrash

`func (o *BTDocumentSummaryInfo) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTDocumentSummaryInfo) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTDocumentSummaryInfo) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTDocumentSummaryInfo) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetTrashedAt

`func (o *BTDocumentSummaryInfo) GetTrashedAt() JSONTime`

GetTrashedAt returns the TrashedAt field if non-nil, zero value otherwise.

### GetTrashedAtOk

`func (o *BTDocumentSummaryInfo) GetTrashedAtOk() (*JSONTime, bool)`

GetTrashedAtOk returns a tuple with the TrashedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedAt

`func (o *BTDocumentSummaryInfo) SetTrashedAt(v JSONTime)`

SetTrashedAt sets TrashedAt field to given value.

### HasTrashedAt

`func (o *BTDocumentSummaryInfo) HasTrashedAt() bool`

HasTrashedAt returns a boolean if a field has been set.

### GetUserAccountLimitsBreached

`func (o *BTDocumentSummaryInfo) GetUserAccountLimitsBreached() bool`

GetUserAccountLimitsBreached returns the UserAccountLimitsBreached field if non-nil, zero value otherwise.

### GetUserAccountLimitsBreachedOk

`func (o *BTDocumentSummaryInfo) GetUserAccountLimitsBreachedOk() (*bool, bool)`

GetUserAccountLimitsBreachedOk returns a tuple with the UserAccountLimitsBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountLimitsBreached

`func (o *BTDocumentSummaryInfo) SetUserAccountLimitsBreached(v bool)`

SetUserAccountLimitsBreached sets UserAccountLimitsBreached field to given value.

### HasUserAccountLimitsBreached

`func (o *BTDocumentSummaryInfo) HasUserAccountLimitsBreached() bool`

HasUserAccountLimitsBreached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


