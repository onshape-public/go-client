# BTDocumentInfo

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
**Trash** | Pointer to **bool** |  | [optional] 
**TrashedAt** | Pointer to **JSONTime** |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**UnparentHref** | Pointer to **string** |  | [optional] 
**UserAccountLimitsBreached** | Pointer to **bool** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTDocumentInfo

`func NewBTDocumentInfo() *BTDocumentInfo`

NewBTDocumentInfo instantiates a new BTDocumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentInfoWithDefaults

`func NewBTDocumentInfoWithDefaults() *BTDocumentInfo`

NewBTDocumentInfoWithDefaults instantiates a new BTDocumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousAccessAllowed

`func (o *BTDocumentInfo) GetAnonymousAccessAllowed() bool`

GetAnonymousAccessAllowed returns the AnonymousAccessAllowed field if non-nil, zero value otherwise.

### GetAnonymousAccessAllowedOk

`func (o *BTDocumentInfo) GetAnonymousAccessAllowedOk() (*bool, bool)`

GetAnonymousAccessAllowedOk returns a tuple with the AnonymousAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessAllowed

`func (o *BTDocumentInfo) SetAnonymousAccessAllowed(v bool)`

SetAnonymousAccessAllowed sets AnonymousAccessAllowed field to given value.

### HasAnonymousAccessAllowed

`func (o *BTDocumentInfo) HasAnonymousAccessAllowed() bool`

HasAnonymousAccessAllowed returns a boolean if a field has been set.

### GetAnonymousAllowsExport

`func (o *BTDocumentInfo) GetAnonymousAllowsExport() bool`

GetAnonymousAllowsExport returns the AnonymousAllowsExport field if non-nil, zero value otherwise.

### GetAnonymousAllowsExportOk

`func (o *BTDocumentInfo) GetAnonymousAllowsExportOk() (*bool, bool)`

GetAnonymousAllowsExportOk returns a tuple with the AnonymousAllowsExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAllowsExport

`func (o *BTDocumentInfo) SetAnonymousAllowsExport(v bool)`

SetAnonymousAllowsExport sets AnonymousAllowsExport field to given value.

### HasAnonymousAllowsExport

`func (o *BTDocumentInfo) HasAnonymousAllowsExport() bool`

HasAnonymousAllowsExport returns a boolean if a field has been set.

### GetCanMove

`func (o *BTDocumentInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTDocumentInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTDocumentInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTDocumentInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetCanUnshare

`func (o *BTDocumentInfo) GetCanUnshare() bool`

GetCanUnshare returns the CanUnshare field if non-nil, zero value otherwise.

### GetCanUnshareOk

`func (o *BTDocumentInfo) GetCanUnshareOk() (*bool, bool)`

GetCanUnshareOk returns a tuple with the CanUnshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnshare

`func (o *BTDocumentInfo) SetCanUnshare(v bool)`

SetCanUnshare sets CanUnshare field to given value.

### HasCanUnshare

`func (o *BTDocumentInfo) HasCanUnshare() bool`

HasCanUnshare returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTDocumentInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTDocumentInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTDocumentInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTDocumentInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTDocumentInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDocumentInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDocumentInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDocumentInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedWithEducationPlan

`func (o *BTDocumentInfo) GetCreatedWithEducationPlan() bool`

GetCreatedWithEducationPlan returns the CreatedWithEducationPlan field if non-nil, zero value otherwise.

### GetCreatedWithEducationPlanOk

`func (o *BTDocumentInfo) GetCreatedWithEducationPlanOk() (*bool, bool)`

GetCreatedWithEducationPlanOk returns a tuple with the CreatedWithEducationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithEducationPlan

`func (o *BTDocumentInfo) SetCreatedWithEducationPlan(v bool)`

SetCreatedWithEducationPlan sets CreatedWithEducationPlan field to given value.

### HasCreatedWithEducationPlan

`func (o *BTDocumentInfo) HasCreatedWithEducationPlan() bool`

HasCreatedWithEducationPlan returns a boolean if a field has been set.

### GetDefaultElementId

`func (o *BTDocumentInfo) GetDefaultElementId() string`

GetDefaultElementId returns the DefaultElementId field if non-nil, zero value otherwise.

### GetDefaultElementIdOk

`func (o *BTDocumentInfo) GetDefaultElementIdOk() (*string, bool)`

GetDefaultElementIdOk returns a tuple with the DefaultElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultElementId

`func (o *BTDocumentInfo) SetDefaultElementId(v string)`

SetDefaultElementId sets DefaultElementId field to given value.

### HasDefaultElementId

`func (o *BTDocumentInfo) HasDefaultElementId() bool`

HasDefaultElementId returns a boolean if a field has been set.

### GetDefaultWorkspace

`func (o *BTDocumentInfo) GetDefaultWorkspace() BTWorkspaceInfo`

GetDefaultWorkspace returns the DefaultWorkspace field if non-nil, zero value otherwise.

### GetDefaultWorkspaceOk

`func (o *BTDocumentInfo) GetDefaultWorkspaceOk() (*BTWorkspaceInfo, bool)`

GetDefaultWorkspaceOk returns a tuple with the DefaultWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkspace

`func (o *BTDocumentInfo) SetDefaultWorkspace(v BTWorkspaceInfo)`

SetDefaultWorkspace sets DefaultWorkspace field to given value.

### HasDefaultWorkspace

`func (o *BTDocumentInfo) HasDefaultWorkspace() bool`

HasDefaultWorkspace returns a boolean if a field has been set.

### GetDescription

`func (o *BTDocumentInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDocumentInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDocumentInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDocumentInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentLabels

`func (o *BTDocumentInfo) GetDocumentLabels() []BTDocumentLabelInfo`

GetDocumentLabels returns the DocumentLabels field if non-nil, zero value otherwise.

### GetDocumentLabelsOk

`func (o *BTDocumentInfo) GetDocumentLabelsOk() (*[]BTDocumentLabelInfo, bool)`

GetDocumentLabelsOk returns a tuple with the DocumentLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLabels

`func (o *BTDocumentInfo) SetDocumentLabels(v []BTDocumentLabelInfo)`

SetDocumentLabels sets DocumentLabels field to given value.

### HasDocumentLabels

`func (o *BTDocumentInfo) HasDocumentLabels() bool`

HasDocumentLabels returns a boolean if a field has been set.

### GetDocumentThumbnailElementId

`func (o *BTDocumentInfo) GetDocumentThumbnailElementId() string`

GetDocumentThumbnailElementId returns the DocumentThumbnailElementId field if non-nil, zero value otherwise.

### GetDocumentThumbnailElementIdOk

`func (o *BTDocumentInfo) GetDocumentThumbnailElementIdOk() (*string, bool)`

GetDocumentThumbnailElementIdOk returns a tuple with the DocumentThumbnailElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentThumbnailElementId

`func (o *BTDocumentInfo) SetDocumentThumbnailElementId(v string)`

SetDocumentThumbnailElementId sets DocumentThumbnailElementId field to given value.

### HasDocumentThumbnailElementId

`func (o *BTDocumentInfo) HasDocumentThumbnailElementId() bool`

HasDocumentThumbnailElementId returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTDocumentInfo) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTDocumentInfo) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTDocumentInfo) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTDocumentInfo) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDuplicateNameViolationError

`func (o *BTDocumentInfo) GetDuplicateNameViolationError() string`

GetDuplicateNameViolationError returns the DuplicateNameViolationError field if non-nil, zero value otherwise.

### GetDuplicateNameViolationErrorOk

`func (o *BTDocumentInfo) GetDuplicateNameViolationErrorOk() (*string, bool)`

GetDuplicateNameViolationErrorOk returns a tuple with the DuplicateNameViolationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateNameViolationError

`func (o *BTDocumentInfo) SetDuplicateNameViolationError(v string)`

SetDuplicateNameViolationError sets DuplicateNameViolationError field to given value.

### HasDuplicateNameViolationError

`func (o *BTDocumentInfo) HasDuplicateNameViolationError() bool`

HasDuplicateNameViolationError returns a boolean if a field has been set.

### GetForceExportRules

`func (o *BTDocumentInfo) GetForceExportRules() bool`

GetForceExportRules returns the ForceExportRules field if non-nil, zero value otherwise.

### GetForceExportRulesOk

`func (o *BTDocumentInfo) GetForceExportRulesOk() (*bool, bool)`

GetForceExportRulesOk returns a tuple with the ForceExportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExportRules

`func (o *BTDocumentInfo) SetForceExportRules(v bool)`

SetForceExportRules sets ForceExportRules field to given value.

### HasForceExportRules

`func (o *BTDocumentInfo) HasForceExportRules() bool`

HasForceExportRules returns a boolean if a field has been set.

### GetHasReleaseRevisionableObjects

`func (o *BTDocumentInfo) GetHasReleaseRevisionableObjects() bool`

GetHasReleaseRevisionableObjects returns the HasReleaseRevisionableObjects field if non-nil, zero value otherwise.

### GetHasReleaseRevisionableObjectsOk

`func (o *BTDocumentInfo) GetHasReleaseRevisionableObjectsOk() (*bool, bool)`

GetHasReleaseRevisionableObjectsOk returns a tuple with the HasReleaseRevisionableObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReleaseRevisionableObjects

`func (o *BTDocumentInfo) SetHasReleaseRevisionableObjects(v bool)`

SetHasReleaseRevisionableObjects sets HasReleaseRevisionableObjects field to given value.

### HasHasReleaseRevisionableObjects

`func (o *BTDocumentInfo) HasHasReleaseRevisionableObjects() bool`

HasHasReleaseRevisionableObjects returns a boolean if a field has been set.

### GetHasRelevantInsertables

`func (o *BTDocumentInfo) GetHasRelevantInsertables() bool`

GetHasRelevantInsertables returns the HasRelevantInsertables field if non-nil, zero value otherwise.

### GetHasRelevantInsertablesOk

`func (o *BTDocumentInfo) GetHasRelevantInsertablesOk() (*bool, bool)`

GetHasRelevantInsertablesOk returns a tuple with the HasRelevantInsertables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRelevantInsertables

`func (o *BTDocumentInfo) SetHasRelevantInsertables(v bool)`

SetHasRelevantInsertables sets HasRelevantInsertables field to given value.

### HasHasRelevantInsertables

`func (o *BTDocumentInfo) HasHasRelevantInsertables() bool`

HasHasRelevantInsertables returns a boolean if a field has been set.

### GetHref

`func (o *BTDocumentInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTDocumentInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTDocumentInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTDocumentInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTDocumentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDocumentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDocumentInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDocumentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTDocumentInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTDocumentInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTDocumentInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTDocumentInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTDocumentInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTDocumentInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTDocumentInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTDocumentInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTDocumentInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTDocumentInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTDocumentInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTDocumentInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetIsUpgradedToLatestVersion

`func (o *BTDocumentInfo) GetIsUpgradedToLatestVersion() bool`

GetIsUpgradedToLatestVersion returns the IsUpgradedToLatestVersion field if non-nil, zero value otherwise.

### GetIsUpgradedToLatestVersionOk

`func (o *BTDocumentInfo) GetIsUpgradedToLatestVersionOk() (*bool, bool)`

GetIsUpgradedToLatestVersionOk returns a tuple with the IsUpgradedToLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradedToLatestVersion

`func (o *BTDocumentInfo) SetIsUpgradedToLatestVersion(v bool)`

SetIsUpgradedToLatestVersion sets IsUpgradedToLatestVersion field to given value.

### HasIsUpgradedToLatestVersion

`func (o *BTDocumentInfo) HasIsUpgradedToLatestVersion() bool`

HasIsUpgradedToLatestVersion returns a boolean if a field has been set.

### GetIsUsingManagedWorkflow

`func (o *BTDocumentInfo) GetIsUsingManagedWorkflow() bool`

GetIsUsingManagedWorkflow returns the IsUsingManagedWorkflow field if non-nil, zero value otherwise.

### GetIsUsingManagedWorkflowOk

`func (o *BTDocumentInfo) GetIsUsingManagedWorkflowOk() (*bool, bool)`

GetIsUsingManagedWorkflowOk returns a tuple with the IsUsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingManagedWorkflow

`func (o *BTDocumentInfo) SetIsUsingManagedWorkflow(v bool)`

SetIsUsingManagedWorkflow sets IsUsingManagedWorkflow field to given value.

### HasIsUsingManagedWorkflow

`func (o *BTDocumentInfo) HasIsUsingManagedWorkflow() bool`

HasIsUsingManagedWorkflow returns a boolean if a field has been set.

### GetLikedByCurrentUser

`func (o *BTDocumentInfo) GetLikedByCurrentUser() bool`

GetLikedByCurrentUser returns the LikedByCurrentUser field if non-nil, zero value otherwise.

### GetLikedByCurrentUserOk

`func (o *BTDocumentInfo) GetLikedByCurrentUserOk() (*bool, bool)`

GetLikedByCurrentUserOk returns a tuple with the LikedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedByCurrentUser

`func (o *BTDocumentInfo) SetLikedByCurrentUser(v bool)`

SetLikedByCurrentUser sets LikedByCurrentUser field to given value.

### HasLikedByCurrentUser

`func (o *BTDocumentInfo) HasLikedByCurrentUser() bool`

HasLikedByCurrentUser returns a boolean if a field has been set.

### GetLikes

`func (o *BTDocumentInfo) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BTDocumentInfo) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BTDocumentInfo) SetLikes(v int64)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BTDocumentInfo) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTDocumentInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTDocumentInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTDocumentInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTDocumentInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTDocumentInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTDocumentInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTDocumentInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTDocumentInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTDocumentInfo) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTDocumentInfo) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTDocumentInfo) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTDocumentInfo) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetNotes

`func (o *BTDocumentInfo) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BTDocumentInfo) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BTDocumentInfo) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BTDocumentInfo) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumberOfTimesCopied

`func (o *BTDocumentInfo) GetNumberOfTimesCopied() int64`

GetNumberOfTimesCopied returns the NumberOfTimesCopied field if non-nil, zero value otherwise.

### GetNumberOfTimesCopiedOk

`func (o *BTDocumentInfo) GetNumberOfTimesCopiedOk() (*int64, bool)`

GetNumberOfTimesCopiedOk returns a tuple with the NumberOfTimesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesCopied

`func (o *BTDocumentInfo) SetNumberOfTimesCopied(v int64)`

SetNumberOfTimesCopied sets NumberOfTimesCopied field to given value.

### HasNumberOfTimesCopied

`func (o *BTDocumentInfo) HasNumberOfTimesCopied() bool`

HasNumberOfTimesCopied returns a boolean if a field has been set.

### GetNumberOfTimesReferenced

`func (o *BTDocumentInfo) GetNumberOfTimesReferenced() int64`

GetNumberOfTimesReferenced returns the NumberOfTimesReferenced field if non-nil, zero value otherwise.

### GetNumberOfTimesReferencedOk

`func (o *BTDocumentInfo) GetNumberOfTimesReferencedOk() (*int64, bool)`

GetNumberOfTimesReferencedOk returns a tuple with the NumberOfTimesReferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTimesReferenced

`func (o *BTDocumentInfo) SetNumberOfTimesReferenced(v int64)`

SetNumberOfTimesReferenced sets NumberOfTimesReferenced field to given value.

### HasNumberOfTimesReferenced

`func (o *BTDocumentInfo) HasNumberOfTimesReferenced() bool`

HasNumberOfTimesReferenced returns a boolean if a field has been set.

### GetOwner

`func (o *BTDocumentInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTDocumentInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTDocumentInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTDocumentInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentId

`func (o *BTDocumentInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTDocumentInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTDocumentInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTDocumentInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermission

`func (o *BTDocumentInfo) GetPermission() BTOldPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTDocumentInfo) GetPermissionOk() (*BTOldPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTDocumentInfo) SetPermission(v BTOldPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTDocumentInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTDocumentInfo) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTDocumentInfo) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTDocumentInfo) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTDocumentInfo) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetProjectId

`func (o *BTDocumentInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTDocumentInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTDocumentInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTDocumentInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPublic

`func (o *BTDocumentInfo) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTDocumentInfo) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTDocumentInfo) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTDocumentInfo) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPublishedVersionId

`func (o *BTDocumentInfo) GetPublishedVersionId() string`

GetPublishedVersionId returns the PublishedVersionId field if non-nil, zero value otherwise.

### GetPublishedVersionIdOk

`func (o *BTDocumentInfo) GetPublishedVersionIdOk() (*string, bool)`

GetPublishedVersionIdOk returns a tuple with the PublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersionId

`func (o *BTDocumentInfo) SetPublishedVersionId(v string)`

SetPublishedVersionId sets PublishedVersionId field to given value.

### HasPublishedVersionId

`func (o *BTDocumentInfo) HasPublishedVersionId() bool`

HasPublishedVersionId returns a boolean if a field has been set.

### GetRecentVersion

`func (o *BTDocumentInfo) GetRecentVersion() BTBaseInfo`

GetRecentVersion returns the RecentVersion field if non-nil, zero value otherwise.

### GetRecentVersionOk

`func (o *BTDocumentInfo) GetRecentVersionOk() (*BTBaseInfo, bool)`

GetRecentVersionOk returns a tuple with the RecentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVersion

`func (o *BTDocumentInfo) SetRecentVersion(v BTBaseInfo)`

SetRecentVersion sets RecentVersion field to given value.

### HasRecentVersion

`func (o *BTDocumentInfo) HasRecentVersion() bool`

HasRecentVersion returns a boolean if a field has been set.

### GetResourceType

`func (o *BTDocumentInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTDocumentInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTDocumentInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTDocumentInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSequence

`func (o *BTDocumentInfo) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BTDocumentInfo) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BTDocumentInfo) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BTDocumentInfo) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSupportTeamUserAndShared

`func (o *BTDocumentInfo) GetSupportTeamUserAndShared() bool`

GetSupportTeamUserAndShared returns the SupportTeamUserAndShared field if non-nil, zero value otherwise.

### GetSupportTeamUserAndSharedOk

`func (o *BTDocumentInfo) GetSupportTeamUserAndSharedOk() (*bool, bool)`

GetSupportTeamUserAndSharedOk returns a tuple with the SupportTeamUserAndShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeamUserAndShared

`func (o *BTDocumentInfo) SetSupportTeamUserAndShared(v bool)`

SetSupportTeamUserAndShared sets SupportTeamUserAndShared field to given value.

### HasSupportTeamUserAndShared

`func (o *BTDocumentInfo) HasSupportTeamUserAndShared() bool`

HasSupportTeamUserAndShared returns a boolean if a field has been set.

### GetTags

`func (o *BTDocumentInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTDocumentInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTDocumentInfo) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTDocumentInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTDocumentInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTDocumentInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTDocumentInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTDocumentInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTrash

`func (o *BTDocumentInfo) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTDocumentInfo) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTDocumentInfo) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTDocumentInfo) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetTrashedAt

`func (o *BTDocumentInfo) GetTrashedAt() JSONTime`

GetTrashedAt returns the TrashedAt field if non-nil, zero value otherwise.

### GetTrashedAtOk

`func (o *BTDocumentInfo) GetTrashedAtOk() (*JSONTime, bool)`

GetTrashedAtOk returns a tuple with the TrashedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedAt

`func (o *BTDocumentInfo) SetTrashedAt(v JSONTime)`

SetTrashedAt sets TrashedAt field to given value.

### HasTrashedAt

`func (o *BTDocumentInfo) HasTrashedAt() bool`

HasTrashedAt returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTDocumentInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTDocumentInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTDocumentInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTDocumentInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetUnparentHref

`func (o *BTDocumentInfo) GetUnparentHref() string`

GetUnparentHref returns the UnparentHref field if non-nil, zero value otherwise.

### GetUnparentHrefOk

`func (o *BTDocumentInfo) GetUnparentHrefOk() (*string, bool)`

GetUnparentHrefOk returns a tuple with the UnparentHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparentHref

`func (o *BTDocumentInfo) SetUnparentHref(v string)`

SetUnparentHref sets UnparentHref field to given value.

### HasUnparentHref

`func (o *BTDocumentInfo) HasUnparentHref() bool`

HasUnparentHref returns a boolean if a field has been set.

### GetUserAccountLimitsBreached

`func (o *BTDocumentInfo) GetUserAccountLimitsBreached() bool`

GetUserAccountLimitsBreached returns the UserAccountLimitsBreached field if non-nil, zero value otherwise.

### GetUserAccountLimitsBreachedOk

`func (o *BTDocumentInfo) GetUserAccountLimitsBreachedOk() (*bool, bool)`

GetUserAccountLimitsBreachedOk returns a tuple with the UserAccountLimitsBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountLimitsBreached

`func (o *BTDocumentInfo) SetUserAccountLimitsBreached(v bool)`

SetUserAccountLimitsBreached sets UserAccountLimitsBreached field to given value.

### HasUserAccountLimitsBreached

`func (o *BTDocumentInfo) HasUserAccountLimitsBreached() bool`

HasUserAccountLimitsBreached returns a boolean if a field has been set.

### GetViewRef

`func (o *BTDocumentInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTDocumentInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTDocumentInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTDocumentInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


