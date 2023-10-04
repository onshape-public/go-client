# BTDocumentSummaryInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousAccessAllowed** | Pointer to **bool** |  | [optional] 
**AnonymousAllowsExport** | Pointer to **bool** |  | [optional] 
**CanUnshare** | Pointer to **bool** |  | [optional] 
**CreatedWithEducationPlan** | Pointer to **bool** |  | [optional] 
**DefaultElementId** | Pointer to **string** |  | [optional] 
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

### NewBTDocumentSummaryInfoAllOf

`func NewBTDocumentSummaryInfoAllOf() *BTDocumentSummaryInfoAllOf`

NewBTDocumentSummaryInfoAllOf instantiates a new BTDocumentSummaryInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentSummaryInfoAllOfWithDefaults

`func NewBTDocumentSummaryInfoAllOfWithDefaults() *BTDocumentSummaryInfoAllOf`

NewBTDocumentSummaryInfoAllOfWithDefaults instantiates a new BTDocumentSummaryInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousAccessAllowed

`func (o *BTDocumentSummaryInfoAllOf) GetAnonymousAccessAllowed() bool`

GetAnonymousAccessAllowed returns the AnonymousAccessAllowed field if non-nil, zero value otherwise.

### GetAnonymousAccessAllowedOk

`func (o *BTDocumentSummaryInfoAllOf) GetAnonymousAccessAllowedOk() (*bool, bool)`

GetAnonymousAccessAllowedOk returns a tuple with the AnonymousAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessAllowed

`func (o *BTDocumentSummaryInfoAllOf) SetAnonymousAccessAllowed(v bool)`

SetAnonymousAccessAllowed sets AnonymousAccessAllowed field to given value.

### HasAnonymousAccessAllowed

`func (o *BTDocumentSummaryInfoAllOf) HasAnonymousAccessAllowed() bool`

HasAnonymousAccessAllowed returns a boolean if a field has been set.

### GetAnonymousAllowsExport

`func (o *BTDocumentSummaryInfoAllOf) GetAnonymousAllowsExport() bool`

GetAnonymousAllowsExport returns the AnonymousAllowsExport field if non-nil, zero value otherwise.

### GetAnonymousAllowsExportOk

`func (o *BTDocumentSummaryInfoAllOf) GetAnonymousAllowsExportOk() (*bool, bool)`

GetAnonymousAllowsExportOk returns a tuple with the AnonymousAllowsExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAllowsExport

`func (o *BTDocumentSummaryInfoAllOf) SetAnonymousAllowsExport(v bool)`

SetAnonymousAllowsExport sets AnonymousAllowsExport field to given value.

### HasAnonymousAllowsExport

`func (o *BTDocumentSummaryInfoAllOf) HasAnonymousAllowsExport() bool`

HasAnonymousAllowsExport returns a boolean if a field has been set.

### GetCanUnshare

`func (o *BTDocumentSummaryInfoAllOf) GetCanUnshare() bool`

GetCanUnshare returns the CanUnshare field if non-nil, zero value otherwise.

### GetCanUnshareOk

`func (o *BTDocumentSummaryInfoAllOf) GetCanUnshareOk() (*bool, bool)`

GetCanUnshareOk returns a tuple with the CanUnshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnshare

`func (o *BTDocumentSummaryInfoAllOf) SetCanUnshare(v bool)`

SetCanUnshare sets CanUnshare field to given value.

### HasCanUnshare

`func (o *BTDocumentSummaryInfoAllOf) HasCanUnshare() bool`

HasCanUnshare returns a boolean if a field has been set.

### GetCreatedWithEducationPlan

`func (o *BTDocumentSummaryInfoAllOf) GetCreatedWithEducationPlan() bool`

GetCreatedWithEducationPlan returns the CreatedWithEducationPlan field if non-nil, zero value otherwise.

### GetCreatedWithEducationPlanOk

`func (o *BTDocumentSummaryInfoAllOf) GetCreatedWithEducationPlanOk() (*bool, bool)`

GetCreatedWithEducationPlanOk returns a tuple with the CreatedWithEducationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithEducationPlan

`func (o *BTDocumentSummaryInfoAllOf) SetCreatedWithEducationPlan(v bool)`

SetCreatedWithEducationPlan sets CreatedWithEducationPlan field to given value.

### HasCreatedWithEducationPlan

`func (o *BTDocumentSummaryInfoAllOf) HasCreatedWithEducationPlan() bool`

HasCreatedWithEducationPlan returns a boolean if a field has been set.

### GetDefaultElementId

`func (o *BTDocumentSummaryInfoAllOf) GetDefaultElementId() string`

GetDefaultElementId returns the DefaultElementId field if non-nil, zero value otherwise.

### GetDefaultElementIdOk

`func (o *BTDocumentSummaryInfoAllOf) GetDefaultElementIdOk() (*string, bool)`

GetDefaultElementIdOk returns a tuple with the DefaultElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultElementId

`func (o *BTDocumentSummaryInfoAllOf) SetDefaultElementId(v string)`

SetDefaultElementId sets DefaultElementId field to given value.

### HasDefaultElementId

`func (o *BTDocumentSummaryInfoAllOf) HasDefaultElementId() bool`

HasDefaultElementId returns a boolean if a field has been set.

### GetDefaultWorkspace

`func (o *BTDocumentSummaryInfoAllOf) GetDefaultWorkspace() BTWorkspaceInfo`

GetDefaultWorkspace returns the DefaultWorkspace field if non-nil, zero value otherwise.

### GetDefaultWorkspaceOk

`func (o *BTDocumentSummaryInfoAllOf) GetDefaultWorkspaceOk() (*BTWorkspaceInfo, bool)`

GetDefaultWorkspaceOk returns a tuple with the DefaultWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkspace

`func (o *BTDocumentSummaryInfoAllOf) SetDefaultWorkspace(v BTWorkspaceInfo)`

SetDefaultWorkspace sets DefaultWorkspace field to given value.

### HasDefaultWorkspace

`func (o *BTDocumentSummaryInfoAllOf) HasDefaultWorkspace() bool`

HasDefaultWorkspace returns a boolean if a field has been set.

### GetDocumentLabels

`func (o *BTDocumentSummaryInfoAllOf) GetDocumentLabels() []BTDocumentLabelInfo`

GetDocumentLabels returns the DocumentLabels field if non-nil, zero value otherwise.

### GetDocumentLabelsOk

`func (o *BTDocumentSummaryInfoAllOf) GetDocumentLabelsOk() (*[]BTDocumentLabelInfo, bool)`

GetDocumentLabelsOk returns a tuple with the DocumentLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLabels

`func (o *BTDocumentSummaryInfoAllOf) SetDocumentLabels(v []BTDocumentLabelInfo)`

SetDocumentLabels sets DocumentLabels field to given value.

### HasDocumentLabels

`func (o *BTDocumentSummaryInfoAllOf) HasDocumentLabels() bool`

HasDocumentLabels returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTDocumentSummaryInfoAllOf) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTDocumentSummaryInfoAllOf) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTDocumentSummaryInfoAllOf) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTDocumentSummaryInfoAllOf) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetForceExportRules

`func (o *BTDocumentSummaryInfoAllOf) GetForceExportRules() bool`

GetForceExportRules returns the ForceExportRules field if non-nil, zero value otherwise.

### GetForceExportRulesOk

`func (o *BTDocumentSummaryInfoAllOf) GetForceExportRulesOk() (*bool, bool)`

GetForceExportRulesOk returns a tuple with the ForceExportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExportRules

`func (o *BTDocumentSummaryInfoAllOf) SetForceExportRules(v bool)`

SetForceExportRules sets ForceExportRules field to given value.

### HasForceExportRules

`func (o *BTDocumentSummaryInfoAllOf) HasForceExportRules() bool`

HasForceExportRules returns a boolean if a field has been set.

### GetHasReleaseRevisionableObjects

`func (o *BTDocumentSummaryInfoAllOf) GetHasReleaseRevisionableObjects() bool`

GetHasReleaseRevisionableObjects returns the HasReleaseRevisionableObjects field if non-nil, zero value otherwise.

### GetHasReleaseRevisionableObjectsOk

`func (o *BTDocumentSummaryInfoAllOf) GetHasReleaseRevisionableObjectsOk() (*bool, bool)`

GetHasReleaseRevisionableObjectsOk returns a tuple with the HasReleaseRevisionableObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReleaseRevisionableObjects

`func (o *BTDocumentSummaryInfoAllOf) SetHasReleaseRevisionableObjects(v bool)`

SetHasReleaseRevisionableObjects sets HasReleaseRevisionableObjects field to given value.

### HasHasReleaseRevisionableObjects

`func (o *BTDocumentSummaryInfoAllOf) HasHasReleaseRevisionableObjects() bool`

HasHasReleaseRevisionableObjects returns a boolean if a field has been set.

### GetHasRelevantInsertables

`func (o *BTDocumentSummaryInfoAllOf) GetHasRelevantInsertables() bool`

GetHasRelevantInsertables returns the HasRelevantInsertables field if non-nil, zero value otherwise.

### GetHasRelevantInsertablesOk

`func (o *BTDocumentSummaryInfoAllOf) GetHasRelevantInsertablesOk() (*bool, bool)`

GetHasRelevantInsertablesOk returns a tuple with the HasRelevantInsertables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRelevantInsertables

`func (o *BTDocumentSummaryInfoAllOf) SetHasRelevantInsertables(v bool)`

SetHasRelevantInsertables sets HasRelevantInsertables field to given value.

### HasHasRelevantInsertables

`func (o *BTDocumentSummaryInfoAllOf) HasHasRelevantInsertables() bool`

HasHasRelevantInsertables returns a boolean if a field has been set.

### GetIsOrphaned

`func (o *BTDocumentSummaryInfoAllOf) GetIsOrphaned() bool`

GetIsOrphaned returns the IsOrphaned field if non-nil, zero value otherwise.

### GetIsOrphanedOk

`func (o *BTDocumentSummaryInfoAllOf) GetIsOrphanedOk() (*bool, bool)`

GetIsOrphanedOk returns a tuple with the IsOrphaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrphaned

`func (o *BTDocumentSummaryInfoAllOf) SetIsOrphaned(v bool)`

SetIsOrphaned sets IsOrphaned field to given value.

### HasIsOrphaned

`func (o *BTDocumentSummaryInfoAllOf) HasIsOrphaned() bool`

HasIsOrphaned returns a boolean if a field has been set.

### GetIsUsingManagedWorkflow

`func (o *BTDocumentSummaryInfoAllOf) GetIsUsingManagedWorkflow() bool`

GetIsUsingManagedWorkflow returns the IsUsingManagedWorkflow field if non-nil, zero value otherwise.

### GetIsUsingManagedWorkflowOk

`func (o *BTDocumentSummaryInfoAllOf) GetIsUsingManagedWorkflowOk() (*bool, bool)`

GetIsUsingManagedWorkflowOk returns a tuple with the IsUsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingManagedWorkflow

`func (o *BTDocumentSummaryInfoAllOf) SetIsUsingManagedWorkflow(v bool)`

SetIsUsingManagedWorkflow sets IsUsingManagedWorkflow field to given value.

### HasIsUsingManagedWorkflow

`func (o *BTDocumentSummaryInfoAllOf) HasIsUsingManagedWorkflow() bool`

HasIsUsingManagedWorkflow returns a boolean if a field has been set.

### GetLikedByCurrentUser

`func (o *BTDocumentSummaryInfoAllOf) GetLikedByCurrentUser() bool`

GetLikedByCurrentUser returns the LikedByCurrentUser field if non-nil, zero value otherwise.

### GetLikedByCurrentUserOk

`func (o *BTDocumentSummaryInfoAllOf) GetLikedByCurrentUserOk() (*bool, bool)`

GetLikedByCurrentUserOk returns a tuple with the LikedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedByCurrentUser

`func (o *BTDocumentSummaryInfoAllOf) SetLikedByCurrentUser(v bool)`

SetLikedByCurrentUser sets LikedByCurrentUser field to given value.

### HasLikedByCurrentUser

`func (o *BTDocumentSummaryInfoAllOf) HasLikedByCurrentUser() bool`

HasLikedByCurrentUser returns a boolean if a field has been set.

### GetLikes

`func (o *BTDocumentSummaryInfoAllOf) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BTDocumentSummaryInfoAllOf) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BTDocumentSummaryInfoAllOf) SetLikes(v int64)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BTDocumentSummaryInfoAllOf) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTDocumentSummaryInfoAllOf) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTDocumentSummaryInfoAllOf) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTDocumentSummaryInfoAllOf) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTDocumentSummaryInfoAllOf) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetNotes

`func (o *BTDocumentSummaryInfoAllOf) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BTDocumentSummaryInfoAllOf) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BTDocumentSummaryInfoAllOf) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BTDocumentSummaryInfoAllOf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumberOfTimesCopied

`func (o *BTDocumentSummaryInfoAllOf) GetNumberOfTimesCopied() int64`

GetNumberOfTimesCopied returns the NumberOfTimesCopied field if non-nil, zero value otherwise.

### GetNumberOfTimesCopiedOk

`func (o *BTDocumentSummaryInfoAllOf) GetNumberOfTimesCopiedOk() (*int64, bool)`

GetNumberOfTimesCopiedOk returns a tuple with the NumberOfTimesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesCopied

`func (o *BTDocumentSummaryInfoAllOf) SetNumberOfTimesCopied(v int64)`

SetNumberOfTimesCopied sets NumberOfTimesCopied field to given value.

### HasNumberOfTimesCopied

`func (o *BTDocumentSummaryInfoAllOf) HasNumberOfTimesCopied() bool`

HasNumberOfTimesCopied returns a boolean if a field has been set.

### GetNumberOfTimesReferenced

`func (o *BTDocumentSummaryInfoAllOf) GetNumberOfTimesReferenced() int64`

GetNumberOfTimesReferenced returns the NumberOfTimesReferenced field if non-nil, zero value otherwise.

### GetNumberOfTimesReferencedOk

`func (o *BTDocumentSummaryInfoAllOf) GetNumberOfTimesReferencedOk() (*int64, bool)`

GetNumberOfTimesReferencedOk returns a tuple with the NumberOfTimesReferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesReferenced

`func (o *BTDocumentSummaryInfoAllOf) SetNumberOfTimesReferenced(v int64)`

SetNumberOfTimesReferenced sets NumberOfTimesReferenced field to given value.

### HasNumberOfTimesReferenced

`func (o *BTDocumentSummaryInfoAllOf) HasNumberOfTimesReferenced() bool`

HasNumberOfTimesReferenced returns a boolean if a field has been set.

### GetParentId

`func (o *BTDocumentSummaryInfoAllOf) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTDocumentSummaryInfoAllOf) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTDocumentSummaryInfoAllOf) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTDocumentSummaryInfoAllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermission

`func (o *BTDocumentSummaryInfoAllOf) GetPermission() BTOldPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTDocumentSummaryInfoAllOf) GetPermissionOk() (*BTOldPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTDocumentSummaryInfoAllOf) SetPermission(v BTOldPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTDocumentSummaryInfoAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTDocumentSummaryInfoAllOf) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTDocumentSummaryInfoAllOf) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTDocumentSummaryInfoAllOf) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTDocumentSummaryInfoAllOf) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetPublic

`func (o *BTDocumentSummaryInfoAllOf) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTDocumentSummaryInfoAllOf) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTDocumentSummaryInfoAllOf) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTDocumentSummaryInfoAllOf) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *BTDocumentSummaryInfoAllOf) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *BTDocumentSummaryInfoAllOf) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *BTDocumentSummaryInfoAllOf) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *BTDocumentSummaryInfoAllOf) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetRecentVersion

`func (o *BTDocumentSummaryInfoAllOf) GetRecentVersion() BTBaseInfo`

GetRecentVersion returns the RecentVersion field if non-nil, zero value otherwise.

### GetRecentVersionOk

`func (o *BTDocumentSummaryInfoAllOf) GetRecentVersionOk() (*BTBaseInfo, bool)`

GetRecentVersionOk returns a tuple with the RecentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVersion

`func (o *BTDocumentSummaryInfoAllOf) SetRecentVersion(v BTBaseInfo)`

SetRecentVersion sets RecentVersion field to given value.

### HasRecentVersion

`func (o *BTDocumentSummaryInfoAllOf) HasRecentVersion() bool`

HasRecentVersion returns a boolean if a field has been set.

### GetSequence

`func (o *BTDocumentSummaryInfoAllOf) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BTDocumentSummaryInfoAllOf) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BTDocumentSummaryInfoAllOf) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BTDocumentSummaryInfoAllOf) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSupportTeamUserAndShared

`func (o *BTDocumentSummaryInfoAllOf) GetSupportTeamUserAndShared() bool`

GetSupportTeamUserAndShared returns the SupportTeamUserAndShared field if non-nil, zero value otherwise.

### GetSupportTeamUserAndSharedOk

`func (o *BTDocumentSummaryInfoAllOf) GetSupportTeamUserAndSharedOk() (*bool, bool)`

GetSupportTeamUserAndSharedOk returns a tuple with the SupportTeamUserAndShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeamUserAndShared

`func (o *BTDocumentSummaryInfoAllOf) SetSupportTeamUserAndShared(v bool)`

SetSupportTeamUserAndShared sets SupportTeamUserAndShared field to given value.

### HasSupportTeamUserAndShared

`func (o *BTDocumentSummaryInfoAllOf) HasSupportTeamUserAndShared() bool`

HasSupportTeamUserAndShared returns a boolean if a field has been set.

### GetTags

`func (o *BTDocumentSummaryInfoAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTDocumentSummaryInfoAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTDocumentSummaryInfoAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTDocumentSummaryInfoAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTDocumentSummaryInfoAllOf) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTDocumentSummaryInfoAllOf) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTDocumentSummaryInfoAllOf) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTDocumentSummaryInfoAllOf) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentSummaryInfoAllOf) GetTotalWorkspacesScheduledForUpdate() int32`

GetTotalWorkspacesScheduledForUpdate returns the TotalWorkspacesScheduledForUpdate field if non-nil, zero value otherwise.

### GetTotalWorkspacesScheduledForUpdateOk

`func (o *BTDocumentSummaryInfoAllOf) GetTotalWorkspacesScheduledForUpdateOk() (*int32, bool)`

GetTotalWorkspacesScheduledForUpdateOk returns a tuple with the TotalWorkspacesScheduledForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentSummaryInfoAllOf) SetTotalWorkspacesScheduledForUpdate(v int32)`

SetTotalWorkspacesScheduledForUpdate sets TotalWorkspacesScheduledForUpdate field to given value.

### HasTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentSummaryInfoAllOf) HasTotalWorkspacesScheduledForUpdate() bool`

HasTotalWorkspacesScheduledForUpdate returns a boolean if a field has been set.

### GetTotalWorkspacesUpdating

`func (o *BTDocumentSummaryInfoAllOf) GetTotalWorkspacesUpdating() int32`

GetTotalWorkspacesUpdating returns the TotalWorkspacesUpdating field if non-nil, zero value otherwise.

### GetTotalWorkspacesUpdatingOk

`func (o *BTDocumentSummaryInfoAllOf) GetTotalWorkspacesUpdatingOk() (*int32, bool)`

GetTotalWorkspacesUpdatingOk returns a tuple with the TotalWorkspacesUpdating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesUpdating

`func (o *BTDocumentSummaryInfoAllOf) SetTotalWorkspacesUpdating(v int32)`

SetTotalWorkspacesUpdating sets TotalWorkspacesUpdating field to given value.

### HasTotalWorkspacesUpdating

`func (o *BTDocumentSummaryInfoAllOf) HasTotalWorkspacesUpdating() bool`

HasTotalWorkspacesUpdating returns a boolean if a field has been set.

### GetTrash

`func (o *BTDocumentSummaryInfoAllOf) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTDocumentSummaryInfoAllOf) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTDocumentSummaryInfoAllOf) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTDocumentSummaryInfoAllOf) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetTrashedAt

`func (o *BTDocumentSummaryInfoAllOf) GetTrashedAt() JSONTime`

GetTrashedAt returns the TrashedAt field if non-nil, zero value otherwise.

### GetTrashedAtOk

`func (o *BTDocumentSummaryInfoAllOf) GetTrashedAtOk() (*JSONTime, bool)`

GetTrashedAtOk returns a tuple with the TrashedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedAt

`func (o *BTDocumentSummaryInfoAllOf) SetTrashedAt(v JSONTime)`

SetTrashedAt sets TrashedAt field to given value.

### HasTrashedAt

`func (o *BTDocumentSummaryInfoAllOf) HasTrashedAt() bool`

HasTrashedAt returns a boolean if a field has been set.

### GetUserAccountLimitsBreached

`func (o *BTDocumentSummaryInfoAllOf) GetUserAccountLimitsBreached() bool`

GetUserAccountLimitsBreached returns the UserAccountLimitsBreached field if non-nil, zero value otherwise.

### GetUserAccountLimitsBreachedOk

`func (o *BTDocumentSummaryInfoAllOf) GetUserAccountLimitsBreachedOk() (*bool, bool)`

GetUserAccountLimitsBreachedOk returns a tuple with the UserAccountLimitsBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountLimitsBreached

`func (o *BTDocumentSummaryInfoAllOf) SetUserAccountLimitsBreached(v bool)`

SetUserAccountLimitsBreached sets UserAccountLimitsBreached field to given value.

### HasUserAccountLimitsBreached

`func (o *BTDocumentSummaryInfoAllOf) HasUserAccountLimitsBreached() bool`

HasUserAccountLimitsBreached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


