# BTReleasePackageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeOrderId** | Pointer to **string** |  | [optional] 
**ColumnNames** | Pointer to **map[string]string** |  | [optional] 
**Comments** | Pointer to [**[]BTCommentInfo**](BTCommentInfo.md) |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Detailed** | Pointer to **bool** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsObsoletion** | Pointer to **bool** |  | [optional] 
**Items** | Pointer to [**[]BTReleasePackageItemInfo**](BTReleasePackageItemInfo.md) |  | [optional] 
**LinkedVersionIds** | Pointer to **[]string** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalWorkspaceId** | Pointer to **string** |  | [optional] 
**PackageThumbnail** | Pointer to **string** |  | [optional] 
**ParentComments** | Pointer to [**[]BTReleaseCommentListInfo**](BTReleaseCommentListInfo.md) |  | [optional] 
**ParentPackages** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to [**[]BTWorkflowPropertyInfo**](BTWorkflowPropertyInfo.md) |  | [optional] 
**RevisionRuleId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 
**Workflow** | Pointer to [**BTWorkflowSnapshotInfo**](BTWorkflowSnapshotInfo.md) |  | [optional] 
**WorkflowError** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to [**BTPublishedWorkflowId**](BTPublishedWorkflowId.md) |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTReleasePackageInfo

`func NewBTReleasePackageInfo() *BTReleasePackageInfo`

NewBTReleasePackageInfo instantiates a new BTReleasePackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTReleasePackageInfoWithDefaults

`func NewBTReleasePackageInfoWithDefaults() *BTReleasePackageInfo`

NewBTReleasePackageInfoWithDefaults instantiates a new BTReleasePackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeOrderId

`func (o *BTReleasePackageInfo) GetChangeOrderId() string`

GetChangeOrderId returns the ChangeOrderId field if non-nil, zero value otherwise.

### GetChangeOrderIdOk

`func (o *BTReleasePackageInfo) GetChangeOrderIdOk() (*string, bool)`

GetChangeOrderIdOk returns a tuple with the ChangeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeOrderId

`func (o *BTReleasePackageInfo) SetChangeOrderId(v string)`

SetChangeOrderId sets ChangeOrderId field to given value.

### HasChangeOrderId

`func (o *BTReleasePackageInfo) HasChangeOrderId() bool`

HasChangeOrderId returns a boolean if a field has been set.

### GetColumnNames

`func (o *BTReleasePackageInfo) GetColumnNames() map[string]string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *BTReleasePackageInfo) GetColumnNamesOk() (*map[string]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *BTReleasePackageInfo) SetColumnNames(v map[string]string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *BTReleasePackageInfo) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetComments

`func (o *BTReleasePackageInfo) GetComments() []BTCommentInfo`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BTReleasePackageInfo) GetCommentsOk() (*[]BTCommentInfo, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BTReleasePackageInfo) SetComments(v []BTCommentInfo)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BTReleasePackageInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTReleasePackageInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTReleasePackageInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTReleasePackageInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTReleasePackageInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTReleasePackageInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTReleasePackageInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTReleasePackageInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTReleasePackageInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTReleasePackageInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTReleasePackageInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTReleasePackageInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTReleasePackageInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *BTReleasePackageInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTReleasePackageInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTReleasePackageInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTReleasePackageInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetailed

`func (o *BTReleasePackageInfo) GetDetailed() bool`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *BTReleasePackageInfo) GetDetailedOk() (*bool, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *BTReleasePackageInfo) SetDetailed(v bool)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *BTReleasePackageInfo) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTReleasePackageInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTReleasePackageInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTReleasePackageInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTReleasePackageInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetHref

`func (o *BTReleasePackageInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTReleasePackageInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTReleasePackageInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTReleasePackageInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTReleasePackageInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTReleasePackageInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTReleasePackageInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTReleasePackageInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsObsoletion

`func (o *BTReleasePackageInfo) GetIsObsoletion() bool`

GetIsObsoletion returns the IsObsoletion field if non-nil, zero value otherwise.

### GetIsObsoletionOk

`func (o *BTReleasePackageInfo) GetIsObsoletionOk() (*bool, bool)`

GetIsObsoletionOk returns a tuple with the IsObsoletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObsoletion

`func (o *BTReleasePackageInfo) SetIsObsoletion(v bool)`

SetIsObsoletion sets IsObsoletion field to given value.

### HasIsObsoletion

`func (o *BTReleasePackageInfo) HasIsObsoletion() bool`

HasIsObsoletion returns a boolean if a field has been set.

### GetItems

`func (o *BTReleasePackageInfo) GetItems() []BTReleasePackageItemInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTReleasePackageInfo) GetItemsOk() (*[]BTReleasePackageItemInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTReleasePackageInfo) SetItems(v []BTReleasePackageItemInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTReleasePackageInfo) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinkedVersionIds

`func (o *BTReleasePackageInfo) GetLinkedVersionIds() []string`

GetLinkedVersionIds returns the LinkedVersionIds field if non-nil, zero value otherwise.

### GetLinkedVersionIdsOk

`func (o *BTReleasePackageInfo) GetLinkedVersionIdsOk() (*[]string, bool)`

GetLinkedVersionIdsOk returns a tuple with the LinkedVersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedVersionIds

`func (o *BTReleasePackageInfo) SetLinkedVersionIds(v []string)`

SetLinkedVersionIds sets LinkedVersionIds field to given value.

### HasLinkedVersionIds

`func (o *BTReleasePackageInfo) HasLinkedVersionIds() bool`

HasLinkedVersionIds returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTReleasePackageInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTReleasePackageInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTReleasePackageInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTReleasePackageInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTReleasePackageInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTReleasePackageInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTReleasePackageInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTReleasePackageInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTReleasePackageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTReleasePackageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTReleasePackageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTReleasePackageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalWorkspaceId

`func (o *BTReleasePackageInfo) GetOriginalWorkspaceId() string`

GetOriginalWorkspaceId returns the OriginalWorkspaceId field if non-nil, zero value otherwise.

### GetOriginalWorkspaceIdOk

`func (o *BTReleasePackageInfo) GetOriginalWorkspaceIdOk() (*string, bool)`

GetOriginalWorkspaceIdOk returns a tuple with the OriginalWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWorkspaceId

`func (o *BTReleasePackageInfo) SetOriginalWorkspaceId(v string)`

SetOriginalWorkspaceId sets OriginalWorkspaceId field to given value.

### HasOriginalWorkspaceId

`func (o *BTReleasePackageInfo) HasOriginalWorkspaceId() bool`

HasOriginalWorkspaceId returns a boolean if a field has been set.

### GetPackageThumbnail

`func (o *BTReleasePackageInfo) GetPackageThumbnail() string`

GetPackageThumbnail returns the PackageThumbnail field if non-nil, zero value otherwise.

### GetPackageThumbnailOk

`func (o *BTReleasePackageInfo) GetPackageThumbnailOk() (*string, bool)`

GetPackageThumbnailOk returns a tuple with the PackageThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageThumbnail

`func (o *BTReleasePackageInfo) SetPackageThumbnail(v string)`

SetPackageThumbnail sets PackageThumbnail field to given value.

### HasPackageThumbnail

`func (o *BTReleasePackageInfo) HasPackageThumbnail() bool`

HasPackageThumbnail returns a boolean if a field has been set.

### GetParentComments

`func (o *BTReleasePackageInfo) GetParentComments() []BTReleaseCommentListInfo`

GetParentComments returns the ParentComments field if non-nil, zero value otherwise.

### GetParentCommentsOk

`func (o *BTReleasePackageInfo) GetParentCommentsOk() (*[]BTReleaseCommentListInfo, bool)`

GetParentCommentsOk returns a tuple with the ParentComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentComments

`func (o *BTReleasePackageInfo) SetParentComments(v []BTReleaseCommentListInfo)`

SetParentComments sets ParentComments field to given value.

### HasParentComments

`func (o *BTReleasePackageInfo) HasParentComments() bool`

HasParentComments returns a boolean if a field has been set.

### GetParentPackages

`func (o *BTReleasePackageInfo) GetParentPackages() []string`

GetParentPackages returns the ParentPackages field if non-nil, zero value otherwise.

### GetParentPackagesOk

`func (o *BTReleasePackageInfo) GetParentPackagesOk() (*[]string, bool)`

GetParentPackagesOk returns a tuple with the ParentPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPackages

`func (o *BTReleasePackageInfo) SetParentPackages(v []string)`

SetParentPackages sets ParentPackages field to given value.

### HasParentPackages

`func (o *BTReleasePackageInfo) HasParentPackages() bool`

HasParentPackages returns a boolean if a field has been set.

### GetProperties

`func (o *BTReleasePackageInfo) GetProperties() []BTWorkflowPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTReleasePackageInfo) GetPropertiesOk() (*[]BTWorkflowPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTReleasePackageInfo) SetProperties(v []BTWorkflowPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTReleasePackageInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRevisionRuleId

`func (o *BTReleasePackageInfo) GetRevisionRuleId() string`

GetRevisionRuleId returns the RevisionRuleId field if non-nil, zero value otherwise.

### GetRevisionRuleIdOk

`func (o *BTReleasePackageInfo) GetRevisionRuleIdOk() (*string, bool)`

GetRevisionRuleIdOk returns a tuple with the RevisionRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionRuleId

`func (o *BTReleasePackageInfo) SetRevisionRuleId(v string)`

SetRevisionRuleId sets RevisionRuleId field to given value.

### HasRevisionRuleId

`func (o *BTReleasePackageInfo) HasRevisionRuleId() bool`

HasRevisionRuleId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTReleasePackageInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTReleasePackageInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTReleasePackageInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTReleasePackageInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTReleasePackageInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTReleasePackageInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTReleasePackageInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTReleasePackageInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkflow

`func (o *BTReleasePackageInfo) GetWorkflow() BTWorkflowSnapshotInfo`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BTReleasePackageInfo) GetWorkflowOk() (*BTWorkflowSnapshotInfo, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BTReleasePackageInfo) SetWorkflow(v BTWorkflowSnapshotInfo)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *BTReleasePackageInfo) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetWorkflowError

`func (o *BTReleasePackageInfo) GetWorkflowError() string`

GetWorkflowError returns the WorkflowError field if non-nil, zero value otherwise.

### GetWorkflowErrorOk

`func (o *BTReleasePackageInfo) GetWorkflowErrorOk() (*string, bool)`

GetWorkflowErrorOk returns a tuple with the WorkflowError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowError

`func (o *BTReleasePackageInfo) SetWorkflowError(v string)`

SetWorkflowError sets WorkflowError field to given value.

### HasWorkflowError

`func (o *BTReleasePackageInfo) HasWorkflowError() bool`

HasWorkflowError returns a boolean if a field has been set.

### GetWorkflowId

`func (o *BTReleasePackageInfo) GetWorkflowId() BTPublishedWorkflowId`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *BTReleasePackageInfo) GetWorkflowIdOk() (*BTPublishedWorkflowId, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *BTReleasePackageInfo) SetWorkflowId(v BTPublishedWorkflowId)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *BTReleasePackageInfo) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTReleasePackageInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTReleasePackageInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTReleasePackageInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTReleasePackageInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


