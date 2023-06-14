# BTDocumentSummaryInfo

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
**DocumentType** | Pointer to **int32** |  | [optional] 
**HasReleaseRevisionableObjects** | Pointer to **bool** |  | [optional] 
**HasRelevantInsertables** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsContainer** | Pointer to **bool** |  | [optional] 
**IsEnterpriseOwned** | Pointer to **bool** |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
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

### GetCanMove

`func (o *BTDocumentSummaryInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTDocumentSummaryInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTDocumentSummaryInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTDocumentSummaryInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

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

### GetCreatedAt

`func (o *BTDocumentSummaryInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTDocumentSummaryInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTDocumentSummaryInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTDocumentSummaryInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTDocumentSummaryInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDocumentSummaryInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDocumentSummaryInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDocumentSummaryInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

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

### GetDescription

`func (o *BTDocumentSummaryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDocumentSummaryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDocumentSummaryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDocumentSummaryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### GetHref

`func (o *BTDocumentSummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTDocumentSummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTDocumentSummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTDocumentSummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTDocumentSummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDocumentSummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDocumentSummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDocumentSummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTDocumentSummaryInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTDocumentSummaryInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTDocumentSummaryInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTDocumentSummaryInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTDocumentSummaryInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTDocumentSummaryInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTDocumentSummaryInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTDocumentSummaryInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTDocumentSummaryInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTDocumentSummaryInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTDocumentSummaryInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTDocumentSummaryInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

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

### GetModifiedAt

`func (o *BTDocumentSummaryInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTDocumentSummaryInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTDocumentSummaryInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTDocumentSummaryInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTDocumentSummaryInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTDocumentSummaryInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTDocumentSummaryInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTDocumentSummaryInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentSummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentSummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentSummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentSummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

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

### GetOwner

`func (o *BTDocumentSummaryInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTDocumentSummaryInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTDocumentSummaryInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTDocumentSummaryInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

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

### GetProjectId

`func (o *BTDocumentSummaryInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTDocumentSummaryInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTDocumentSummaryInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTDocumentSummaryInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

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

### GetResourceType

`func (o *BTDocumentSummaryInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTDocumentSummaryInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTDocumentSummaryInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTDocumentSummaryInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

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

### GetTreeHref

`func (o *BTDocumentSummaryInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTDocumentSummaryInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTDocumentSummaryInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTDocumentSummaryInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetUnparentHref

`func (o *BTDocumentSummaryInfo) GetUnparentHref() string`

GetUnparentHref returns the UnparentHref field if non-nil, zero value otherwise.

### GetUnparentHrefOk

`func (o *BTDocumentSummaryInfo) GetUnparentHrefOk() (*string, bool)`

GetUnparentHrefOk returns a tuple with the UnparentHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparentHref

`func (o *BTDocumentSummaryInfo) SetUnparentHref(v string)`

SetUnparentHref sets UnparentHref field to given value.

### HasUnparentHref

`func (o *BTDocumentSummaryInfo) HasUnparentHref() bool`

HasUnparentHref returns a boolean if a field has been set.

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

### GetViewRef

`func (o *BTDocumentSummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTDocumentSummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTDocumentSummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTDocumentSummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


