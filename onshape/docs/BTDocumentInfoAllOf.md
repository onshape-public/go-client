# BTDocumentInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousAccessAllowed** | Pointer to **bool** |  | [optional] 
**AnonymousAllowsExport** | Pointer to **bool** |  | [optional] 
**CanMove** | Pointer to **bool** |  | [optional] 
**CanUnshare** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**CreatedWithEducationPlan** | Pointer to **bool** |  | [optional] 
**DefaultElementId** | Pointer to **string** |  | [optional] 
**DefaultWorkspace** | Pointer to [**BTWorkspaceInfo**](BTWorkspaceInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentLabels** | Pointer to [**[]BTDocumentLabelInfo**](BTDocumentLabelInfo.md) |  | [optional] 
**DocumentThumbnailElementId** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **int32** |  | [optional] 
**DuplicateNameViolationError** | Pointer to **string** |  | [optional] 
**ForceExportRules** | Pointer to **bool** |  | [optional] 
**HasReleaseRevisionableObjects** | Pointer to **bool** |  | [optional] 
**HasRelevantInsertables** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsContainer** | Pointer to **bool** |  | [optional] 
**IsEnterpriseOwned** | Pointer to **bool** |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**IsOrphaned** | Pointer to **bool** |  | [optional] 
**IsUpgradedToLatestVersion** | Pointer to **bool** |  | [optional] 
**IsUsingManagedWorkflow** | Pointer to **bool** |  | [optional] 
**LikedByCurrentUser** | Pointer to **bool** |  | [optional] 
**Likes** | Pointer to **int64** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**NotRevisionManaged** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NumberOfTimesCopied** | Pointer to **int64** |  | [optional] 
**NumberOfTimesReferenced** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to [**BTOldPermission**](BTOldPermission.md) |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**PublishedVersionId** | Pointer to **string** |  | [optional] 
**RecentVersion** | Pointer to [**BTBaseInfo**](BTBaseInfo.md) |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Sequence** | Pointer to **string** |  | [optional] 
**SupportTeamUserAndShared** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**TotalWorkspacesScheduledForUpdate** | Pointer to **int32** |  | [optional] 
**TotalWorkspacesUpdating** | Pointer to **int32** |  | [optional] 
**TracingEnabled** | Pointer to **bool** |  | [optional] 
**Trash** | Pointer to **bool** |  | [optional] 
**TrashedAt** | Pointer to **JSONTime** |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**UnparentHref** | Pointer to **string** |  | [optional] 
**UserAccountLimitsBreached** | Pointer to **bool** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTDocumentInfoAllOf

`func NewBTDocumentInfoAllOf() *BTDocumentInfoAllOf`

NewBTDocumentInfoAllOf instantiates a new BTDocumentInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentInfoAllOfWithDefaults

`func NewBTDocumentInfoAllOfWithDefaults() *BTDocumentInfoAllOf`

NewBTDocumentInfoAllOfWithDefaults instantiates a new BTDocumentInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousAccessAllowed

`func (o *BTDocumentInfoAllOf) GetAnonymousAccessAllowed() bool`

GetAnonymousAccessAllowed returns the AnonymousAccessAllowed field if non-nil, zero value otherwise.

### GetAnonymousAccessAllowedOk

`func (o *BTDocumentInfoAllOf) GetAnonymousAccessAllowedOk() (*bool, bool)`

GetAnonymousAccessAllowedOk returns a tuple with the AnonymousAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessAllowed

`func (o *BTDocumentInfoAllOf) SetAnonymousAccessAllowed(v bool)`

SetAnonymousAccessAllowed sets AnonymousAccessAllowed field to given value.

### HasAnonymousAccessAllowed

`func (o *BTDocumentInfoAllOf) HasAnonymousAccessAllowed() bool`

HasAnonymousAccessAllowed returns a boolean if a field has been set.

### GetAnonymousAllowsExport

`func (o *BTDocumentInfoAllOf) GetAnonymousAllowsExport() bool`

GetAnonymousAllowsExport returns the AnonymousAllowsExport field if non-nil, zero value otherwise.

### GetAnonymousAllowsExportOk

`func (o *BTDocumentInfoAllOf) GetAnonymousAllowsExportOk() (*bool, bool)`

GetAnonymousAllowsExportOk returns a tuple with the AnonymousAllowsExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAllowsExport

`func (o *BTDocumentInfoAllOf) SetAnonymousAllowsExport(v bool)`

SetAnonymousAllowsExport sets AnonymousAllowsExport field to given value.

### HasAnonymousAllowsExport

`func (o *BTDocumentInfoAllOf) HasAnonymousAllowsExport() bool`

HasAnonymousAllowsExport returns a boolean if a field has been set.

### GetCanMove

`func (o *BTDocumentInfoAllOf) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTDocumentInfoAllOf) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTDocumentInfoAllOf) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTDocumentInfoAllOf) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetCanUnshare

`func (o *BTDocumentInfoAllOf) GetCanUnshare() bool`

GetCanUnshare returns the CanUnshare field if non-nil, zero value otherwise.

### GetCanUnshareOk

`func (o *BTDocumentInfoAllOf) GetCanUnshareOk() (*bool, bool)`

GetCanUnshareOk returns a tuple with the CanUnshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnshare

`func (o *BTDocumentInfoAllOf) SetCanUnshare(v bool)`

SetCanUnshare sets CanUnshare field to given value.

### HasCanUnshare

`func (o *BTDocumentInfoAllOf) HasCanUnshare() bool`

HasCanUnshare returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTDocumentInfoAllOf) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTDocumentInfoAllOf) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTDocumentInfoAllOf) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTDocumentInfoAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTDocumentInfoAllOf) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDocumentInfoAllOf) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDocumentInfoAllOf) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDocumentInfoAllOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedWithEducationPlan

`func (o *BTDocumentInfoAllOf) GetCreatedWithEducationPlan() bool`

GetCreatedWithEducationPlan returns the CreatedWithEducationPlan field if non-nil, zero value otherwise.

### GetCreatedWithEducationPlanOk

`func (o *BTDocumentInfoAllOf) GetCreatedWithEducationPlanOk() (*bool, bool)`

GetCreatedWithEducationPlanOk returns a tuple with the CreatedWithEducationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithEducationPlan

`func (o *BTDocumentInfoAllOf) SetCreatedWithEducationPlan(v bool)`

SetCreatedWithEducationPlan sets CreatedWithEducationPlan field to given value.

### HasCreatedWithEducationPlan

`func (o *BTDocumentInfoAllOf) HasCreatedWithEducationPlan() bool`

HasCreatedWithEducationPlan returns a boolean if a field has been set.

### GetDefaultElementId

`func (o *BTDocumentInfoAllOf) GetDefaultElementId() string`

GetDefaultElementId returns the DefaultElementId field if non-nil, zero value otherwise.

### GetDefaultElementIdOk

`func (o *BTDocumentInfoAllOf) GetDefaultElementIdOk() (*string, bool)`

GetDefaultElementIdOk returns a tuple with the DefaultElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultElementId

`func (o *BTDocumentInfoAllOf) SetDefaultElementId(v string)`

SetDefaultElementId sets DefaultElementId field to given value.

### HasDefaultElementId

`func (o *BTDocumentInfoAllOf) HasDefaultElementId() bool`

HasDefaultElementId returns a boolean if a field has been set.

### GetDefaultWorkspace

`func (o *BTDocumentInfoAllOf) GetDefaultWorkspace() BTWorkspaceInfo`

GetDefaultWorkspace returns the DefaultWorkspace field if non-nil, zero value otherwise.

### GetDefaultWorkspaceOk

`func (o *BTDocumentInfoAllOf) GetDefaultWorkspaceOk() (*BTWorkspaceInfo, bool)`

GetDefaultWorkspaceOk returns a tuple with the DefaultWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkspace

`func (o *BTDocumentInfoAllOf) SetDefaultWorkspace(v BTWorkspaceInfo)`

SetDefaultWorkspace sets DefaultWorkspace field to given value.

### HasDefaultWorkspace

`func (o *BTDocumentInfoAllOf) HasDefaultWorkspace() bool`

HasDefaultWorkspace returns a boolean if a field has been set.

### GetDescription

`func (o *BTDocumentInfoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDocumentInfoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDocumentInfoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDocumentInfoAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentLabels

`func (o *BTDocumentInfoAllOf) GetDocumentLabels() []BTDocumentLabelInfo`

GetDocumentLabels returns the DocumentLabels field if non-nil, zero value otherwise.

### GetDocumentLabelsOk

`func (o *BTDocumentInfoAllOf) GetDocumentLabelsOk() (*[]BTDocumentLabelInfo, bool)`

GetDocumentLabelsOk returns a tuple with the DocumentLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLabels

`func (o *BTDocumentInfoAllOf) SetDocumentLabels(v []BTDocumentLabelInfo)`

SetDocumentLabels sets DocumentLabels field to given value.

### HasDocumentLabels

`func (o *BTDocumentInfoAllOf) HasDocumentLabels() bool`

HasDocumentLabels returns a boolean if a field has been set.

### GetDocumentThumbnailElementId

`func (o *BTDocumentInfoAllOf) GetDocumentThumbnailElementId() string`

GetDocumentThumbnailElementId returns the DocumentThumbnailElementId field if non-nil, zero value otherwise.

### GetDocumentThumbnailElementIdOk

`func (o *BTDocumentInfoAllOf) GetDocumentThumbnailElementIdOk() (*string, bool)`

GetDocumentThumbnailElementIdOk returns a tuple with the DocumentThumbnailElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentThumbnailElementId

`func (o *BTDocumentInfoAllOf) SetDocumentThumbnailElementId(v string)`

SetDocumentThumbnailElementId sets DocumentThumbnailElementId field to given value.

### HasDocumentThumbnailElementId

`func (o *BTDocumentInfoAllOf) HasDocumentThumbnailElementId() bool`

HasDocumentThumbnailElementId returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTDocumentInfoAllOf) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTDocumentInfoAllOf) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTDocumentInfoAllOf) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTDocumentInfoAllOf) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDuplicateNameViolationError

`func (o *BTDocumentInfoAllOf) GetDuplicateNameViolationError() string`

GetDuplicateNameViolationError returns the DuplicateNameViolationError field if non-nil, zero value otherwise.

### GetDuplicateNameViolationErrorOk

`func (o *BTDocumentInfoAllOf) GetDuplicateNameViolationErrorOk() (*string, bool)`

GetDuplicateNameViolationErrorOk returns a tuple with the DuplicateNameViolationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateNameViolationError

`func (o *BTDocumentInfoAllOf) SetDuplicateNameViolationError(v string)`

SetDuplicateNameViolationError sets DuplicateNameViolationError field to given value.

### HasDuplicateNameViolationError

`func (o *BTDocumentInfoAllOf) HasDuplicateNameViolationError() bool`

HasDuplicateNameViolationError returns a boolean if a field has been set.

### GetForceExportRules

`func (o *BTDocumentInfoAllOf) GetForceExportRules() bool`

GetForceExportRules returns the ForceExportRules field if non-nil, zero value otherwise.

### GetForceExportRulesOk

`func (o *BTDocumentInfoAllOf) GetForceExportRulesOk() (*bool, bool)`

GetForceExportRulesOk returns a tuple with the ForceExportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExportRules

`func (o *BTDocumentInfoAllOf) SetForceExportRules(v bool)`

SetForceExportRules sets ForceExportRules field to given value.

### HasForceExportRules

`func (o *BTDocumentInfoAllOf) HasForceExportRules() bool`

HasForceExportRules returns a boolean if a field has been set.

### GetHasReleaseRevisionableObjects

`func (o *BTDocumentInfoAllOf) GetHasReleaseRevisionableObjects() bool`

GetHasReleaseRevisionableObjects returns the HasReleaseRevisionableObjects field if non-nil, zero value otherwise.

### GetHasReleaseRevisionableObjectsOk

`func (o *BTDocumentInfoAllOf) GetHasReleaseRevisionableObjectsOk() (*bool, bool)`

GetHasReleaseRevisionableObjectsOk returns a tuple with the HasReleaseRevisionableObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReleaseRevisionableObjects

`func (o *BTDocumentInfoAllOf) SetHasReleaseRevisionableObjects(v bool)`

SetHasReleaseRevisionableObjects sets HasReleaseRevisionableObjects field to given value.

### HasHasReleaseRevisionableObjects

`func (o *BTDocumentInfoAllOf) HasHasReleaseRevisionableObjects() bool`

HasHasReleaseRevisionableObjects returns a boolean if a field has been set.

### GetHasRelevantInsertables

`func (o *BTDocumentInfoAllOf) GetHasRelevantInsertables() bool`

GetHasRelevantInsertables returns the HasRelevantInsertables field if non-nil, zero value otherwise.

### GetHasRelevantInsertablesOk

`func (o *BTDocumentInfoAllOf) GetHasRelevantInsertablesOk() (*bool, bool)`

GetHasRelevantInsertablesOk returns a tuple with the HasRelevantInsertables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRelevantInsertables

`func (o *BTDocumentInfoAllOf) SetHasRelevantInsertables(v bool)`

SetHasRelevantInsertables sets HasRelevantInsertables field to given value.

### HasHasRelevantInsertables

`func (o *BTDocumentInfoAllOf) HasHasRelevantInsertables() bool`

HasHasRelevantInsertables returns a boolean if a field has been set.

### GetHref

`func (o *BTDocumentInfoAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTDocumentInfoAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTDocumentInfoAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTDocumentInfoAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTDocumentInfoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDocumentInfoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDocumentInfoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDocumentInfoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTDocumentInfoAllOf) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTDocumentInfoAllOf) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTDocumentInfoAllOf) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTDocumentInfoAllOf) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTDocumentInfoAllOf) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTDocumentInfoAllOf) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTDocumentInfoAllOf) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTDocumentInfoAllOf) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTDocumentInfoAllOf) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTDocumentInfoAllOf) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTDocumentInfoAllOf) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTDocumentInfoAllOf) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetIsOrphaned

`func (o *BTDocumentInfoAllOf) GetIsOrphaned() bool`

GetIsOrphaned returns the IsOrphaned field if non-nil, zero value otherwise.

### GetIsOrphanedOk

`func (o *BTDocumentInfoAllOf) GetIsOrphanedOk() (*bool, bool)`

GetIsOrphanedOk returns a tuple with the IsOrphaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrphaned

`func (o *BTDocumentInfoAllOf) SetIsOrphaned(v bool)`

SetIsOrphaned sets IsOrphaned field to given value.

### HasIsOrphaned

`func (o *BTDocumentInfoAllOf) HasIsOrphaned() bool`

HasIsOrphaned returns a boolean if a field has been set.

### GetIsUpgradedToLatestVersion

`func (o *BTDocumentInfoAllOf) GetIsUpgradedToLatestVersion() bool`

GetIsUpgradedToLatestVersion returns the IsUpgradedToLatestVersion field if non-nil, zero value otherwise.

### GetIsUpgradedToLatestVersionOk

`func (o *BTDocumentInfoAllOf) GetIsUpgradedToLatestVersionOk() (*bool, bool)`

GetIsUpgradedToLatestVersionOk returns a tuple with the IsUpgradedToLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradedToLatestVersion

`func (o *BTDocumentInfoAllOf) SetIsUpgradedToLatestVersion(v bool)`

SetIsUpgradedToLatestVersion sets IsUpgradedToLatestVersion field to given value.

### HasIsUpgradedToLatestVersion

`func (o *BTDocumentInfoAllOf) HasIsUpgradedToLatestVersion() bool`

HasIsUpgradedToLatestVersion returns a boolean if a field has been set.

### GetIsUsingManagedWorkflow

`func (o *BTDocumentInfoAllOf) GetIsUsingManagedWorkflow() bool`

GetIsUsingManagedWorkflow returns the IsUsingManagedWorkflow field if non-nil, zero value otherwise.

### GetIsUsingManagedWorkflowOk

`func (o *BTDocumentInfoAllOf) GetIsUsingManagedWorkflowOk() (*bool, bool)`

GetIsUsingManagedWorkflowOk returns a tuple with the IsUsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingManagedWorkflow

`func (o *BTDocumentInfoAllOf) SetIsUsingManagedWorkflow(v bool)`

SetIsUsingManagedWorkflow sets IsUsingManagedWorkflow field to given value.

### HasIsUsingManagedWorkflow

`func (o *BTDocumentInfoAllOf) HasIsUsingManagedWorkflow() bool`

HasIsUsingManagedWorkflow returns a boolean if a field has been set.

### GetLikedByCurrentUser

`func (o *BTDocumentInfoAllOf) GetLikedByCurrentUser() bool`

GetLikedByCurrentUser returns the LikedByCurrentUser field if non-nil, zero value otherwise.

### GetLikedByCurrentUserOk

`func (o *BTDocumentInfoAllOf) GetLikedByCurrentUserOk() (*bool, bool)`

GetLikedByCurrentUserOk returns a tuple with the LikedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedByCurrentUser

`func (o *BTDocumentInfoAllOf) SetLikedByCurrentUser(v bool)`

SetLikedByCurrentUser sets LikedByCurrentUser field to given value.

### HasLikedByCurrentUser

`func (o *BTDocumentInfoAllOf) HasLikedByCurrentUser() bool`

HasLikedByCurrentUser returns a boolean if a field has been set.

### GetLikes

`func (o *BTDocumentInfoAllOf) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BTDocumentInfoAllOf) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BTDocumentInfoAllOf) SetLikes(v int64)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BTDocumentInfoAllOf) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTDocumentInfoAllOf) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTDocumentInfoAllOf) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTDocumentInfoAllOf) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTDocumentInfoAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTDocumentInfoAllOf) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTDocumentInfoAllOf) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTDocumentInfoAllOf) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTDocumentInfoAllOf) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTDocumentInfoAllOf) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTDocumentInfoAllOf) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTDocumentInfoAllOf) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTDocumentInfoAllOf) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetNotes

`func (o *BTDocumentInfoAllOf) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BTDocumentInfoAllOf) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BTDocumentInfoAllOf) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BTDocumentInfoAllOf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumberOfTimesCopied

`func (o *BTDocumentInfoAllOf) GetNumberOfTimesCopied() int64`

GetNumberOfTimesCopied returns the NumberOfTimesCopied field if non-nil, zero value otherwise.

### GetNumberOfTimesCopiedOk

`func (o *BTDocumentInfoAllOf) GetNumberOfTimesCopiedOk() (*int64, bool)`

GetNumberOfTimesCopiedOk returns a tuple with the NumberOfTimesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesCopied

`func (o *BTDocumentInfoAllOf) SetNumberOfTimesCopied(v int64)`

SetNumberOfTimesCopied sets NumberOfTimesCopied field to given value.

### HasNumberOfTimesCopied

`func (o *BTDocumentInfoAllOf) HasNumberOfTimesCopied() bool`

HasNumberOfTimesCopied returns a boolean if a field has been set.

### GetNumberOfTimesReferenced

`func (o *BTDocumentInfoAllOf) GetNumberOfTimesReferenced() int64`

GetNumberOfTimesReferenced returns the NumberOfTimesReferenced field if non-nil, zero value otherwise.

### GetNumberOfTimesReferencedOk

`func (o *BTDocumentInfoAllOf) GetNumberOfTimesReferencedOk() (*int64, bool)`

GetNumberOfTimesReferencedOk returns a tuple with the NumberOfTimesReferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesReferenced

`func (o *BTDocumentInfoAllOf) SetNumberOfTimesReferenced(v int64)`

SetNumberOfTimesReferenced sets NumberOfTimesReferenced field to given value.

### HasNumberOfTimesReferenced

`func (o *BTDocumentInfoAllOf) HasNumberOfTimesReferenced() bool`

HasNumberOfTimesReferenced returns a boolean if a field has been set.

### GetOwner

`func (o *BTDocumentInfoAllOf) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTDocumentInfoAllOf) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTDocumentInfoAllOf) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTDocumentInfoAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentId

`func (o *BTDocumentInfoAllOf) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTDocumentInfoAllOf) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTDocumentInfoAllOf) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTDocumentInfoAllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermission

`func (o *BTDocumentInfoAllOf) GetPermission() BTOldPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTDocumentInfoAllOf) GetPermissionOk() (*BTOldPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTDocumentInfoAllOf) SetPermission(v BTOldPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTDocumentInfoAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTDocumentInfoAllOf) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTDocumentInfoAllOf) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTDocumentInfoAllOf) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTDocumentInfoAllOf) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetProjectId

`func (o *BTDocumentInfoAllOf) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTDocumentInfoAllOf) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTDocumentInfoAllOf) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTDocumentInfoAllOf) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPublic

`func (o *BTDocumentInfoAllOf) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTDocumentInfoAllOf) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTDocumentInfoAllOf) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTDocumentInfoAllOf) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *BTDocumentInfoAllOf) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *BTDocumentInfoAllOf) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *BTDocumentInfoAllOf) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *BTDocumentInfoAllOf) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetRecentVersion

`func (o *BTDocumentInfoAllOf) GetRecentVersion() BTBaseInfo`

GetRecentVersion returns the RecentVersion field if non-nil, zero value otherwise.

### GetRecentVersionOk

`func (o *BTDocumentInfoAllOf) GetRecentVersionOk() (*BTBaseInfo, bool)`

GetRecentVersionOk returns a tuple with the RecentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVersion

`func (o *BTDocumentInfoAllOf) SetRecentVersion(v BTBaseInfo)`

SetRecentVersion sets RecentVersion field to given value.

### HasRecentVersion

`func (o *BTDocumentInfoAllOf) HasRecentVersion() bool`

HasRecentVersion returns a boolean if a field has been set.

### GetResourceType

`func (o *BTDocumentInfoAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTDocumentInfoAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTDocumentInfoAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTDocumentInfoAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSequence

`func (o *BTDocumentInfoAllOf) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BTDocumentInfoAllOf) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BTDocumentInfoAllOf) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BTDocumentInfoAllOf) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSupportTeamUserAndShared

`func (o *BTDocumentInfoAllOf) GetSupportTeamUserAndShared() bool`

GetSupportTeamUserAndShared returns the SupportTeamUserAndShared field if non-nil, zero value otherwise.

### GetSupportTeamUserAndSharedOk

`func (o *BTDocumentInfoAllOf) GetSupportTeamUserAndSharedOk() (*bool, bool)`

GetSupportTeamUserAndSharedOk returns a tuple with the SupportTeamUserAndShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeamUserAndShared

`func (o *BTDocumentInfoAllOf) SetSupportTeamUserAndShared(v bool)`

SetSupportTeamUserAndShared sets SupportTeamUserAndShared field to given value.

### HasSupportTeamUserAndShared

`func (o *BTDocumentInfoAllOf) HasSupportTeamUserAndShared() bool`

HasSupportTeamUserAndShared returns a boolean if a field has been set.

### GetTags

`func (o *BTDocumentInfoAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTDocumentInfoAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTDocumentInfoAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTDocumentInfoAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTDocumentInfoAllOf) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTDocumentInfoAllOf) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTDocumentInfoAllOf) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTDocumentInfoAllOf) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentInfoAllOf) GetTotalWorkspacesScheduledForUpdate() int32`

GetTotalWorkspacesScheduledForUpdate returns the TotalWorkspacesScheduledForUpdate field if non-nil, zero value otherwise.

### GetTotalWorkspacesScheduledForUpdateOk

`func (o *BTDocumentInfoAllOf) GetTotalWorkspacesScheduledForUpdateOk() (*int32, bool)`

GetTotalWorkspacesScheduledForUpdateOk returns a tuple with the TotalWorkspacesScheduledForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentInfoAllOf) SetTotalWorkspacesScheduledForUpdate(v int32)`

SetTotalWorkspacesScheduledForUpdate sets TotalWorkspacesScheduledForUpdate field to given value.

### HasTotalWorkspacesScheduledForUpdate

`func (o *BTDocumentInfoAllOf) HasTotalWorkspacesScheduledForUpdate() bool`

HasTotalWorkspacesScheduledForUpdate returns a boolean if a field has been set.

### GetTotalWorkspacesUpdating

`func (o *BTDocumentInfoAllOf) GetTotalWorkspacesUpdating() int32`

GetTotalWorkspacesUpdating returns the TotalWorkspacesUpdating field if non-nil, zero value otherwise.

### GetTotalWorkspacesUpdatingOk

`func (o *BTDocumentInfoAllOf) GetTotalWorkspacesUpdatingOk() (*int32, bool)`

GetTotalWorkspacesUpdatingOk returns a tuple with the TotalWorkspacesUpdating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspacesUpdating

`func (o *BTDocumentInfoAllOf) SetTotalWorkspacesUpdating(v int32)`

SetTotalWorkspacesUpdating sets TotalWorkspacesUpdating field to given value.

### HasTotalWorkspacesUpdating

`func (o *BTDocumentInfoAllOf) HasTotalWorkspacesUpdating() bool`

HasTotalWorkspacesUpdating returns a boolean if a field has been set.

### GetTracingEnabled

`func (o *BTDocumentInfoAllOf) GetTracingEnabled() bool`

GetTracingEnabled returns the TracingEnabled field if non-nil, zero value otherwise.

### GetTracingEnabledOk

`func (o *BTDocumentInfoAllOf) GetTracingEnabledOk() (*bool, bool)`

GetTracingEnabledOk returns a tuple with the TracingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingEnabled

`func (o *BTDocumentInfoAllOf) SetTracingEnabled(v bool)`

SetTracingEnabled sets TracingEnabled field to given value.

### HasTracingEnabled

`func (o *BTDocumentInfoAllOf) HasTracingEnabled() bool`

HasTracingEnabled returns a boolean if a field has been set.

### GetTrash

`func (o *BTDocumentInfoAllOf) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTDocumentInfoAllOf) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTDocumentInfoAllOf) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTDocumentInfoAllOf) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetTrashedAt

`func (o *BTDocumentInfoAllOf) GetTrashedAt() JSONTime`

GetTrashedAt returns the TrashedAt field if non-nil, zero value otherwise.

### GetTrashedAtOk

`func (o *BTDocumentInfoAllOf) GetTrashedAtOk() (*JSONTime, bool)`

GetTrashedAtOk returns a tuple with the TrashedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedAt

`func (o *BTDocumentInfoAllOf) SetTrashedAt(v JSONTime)`

SetTrashedAt sets TrashedAt field to given value.

### HasTrashedAt

`func (o *BTDocumentInfoAllOf) HasTrashedAt() bool`

HasTrashedAt returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTDocumentInfoAllOf) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTDocumentInfoAllOf) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTDocumentInfoAllOf) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTDocumentInfoAllOf) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetUnparentHref

`func (o *BTDocumentInfoAllOf) GetUnparentHref() string`

GetUnparentHref returns the UnparentHref field if non-nil, zero value otherwise.

### GetUnparentHrefOk

`func (o *BTDocumentInfoAllOf) GetUnparentHrefOk() (*string, bool)`

GetUnparentHrefOk returns a tuple with the UnparentHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparentHref

`func (o *BTDocumentInfoAllOf) SetUnparentHref(v string)`

SetUnparentHref sets UnparentHref field to given value.

### HasUnparentHref

`func (o *BTDocumentInfoAllOf) HasUnparentHref() bool`

HasUnparentHref returns a boolean if a field has been set.

### GetUserAccountLimitsBreached

`func (o *BTDocumentInfoAllOf) GetUserAccountLimitsBreached() bool`

GetUserAccountLimitsBreached returns the UserAccountLimitsBreached field if non-nil, zero value otherwise.

### GetUserAccountLimitsBreachedOk

`func (o *BTDocumentInfoAllOf) GetUserAccountLimitsBreachedOk() (*bool, bool)`

GetUserAccountLimitsBreachedOk returns a tuple with the UserAccountLimitsBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountLimitsBreached

`func (o *BTDocumentInfoAllOf) SetUserAccountLimitsBreached(v bool)`

SetUserAccountLimitsBreached sets UserAccountLimitsBreached field to given value.

### HasUserAccountLimitsBreached

`func (o *BTDocumentInfoAllOf) HasUserAccountLimitsBreached() bool`

HasUserAccountLimitsBreached returns a boolean if a field has been set.

### GetViewRef

`func (o *BTDocumentInfoAllOf) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTDocumentInfoAllOf) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTDocumentInfoAllOf) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTDocumentInfoAllOf) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


