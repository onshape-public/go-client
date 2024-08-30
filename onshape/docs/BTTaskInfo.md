# BTTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ApproverRole** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**[]BTCommentInfo**](BTCommentInfo.md) |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Creator** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentName** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **int32** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]BTMetadataPropertyInfo**](BTMetadataPropertyInfo.md) |  | [optional] 
**ResolvedAt** | Pointer to **JSONTime** |  | [optional] 
**ResolvedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Roles** | Pointer to [**[]BTTaskRbacRoleInfo**](BTTaskRbacRoleInfo.md) |  | [optional] 
**SimpleName** | Pointer to **string** |  | [optional] 
**SourceWorkspaceOrVersionName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**TaskItems** | Pointer to [**[]BTTaskItemInfo**](BTTaskItemInfo.md) |  | [optional] 
**TaskTemplate** | Pointer to **int32** |  | [optional] 
**TaskType** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**[]BTTaskTeamSummaryInfo**](BTTaskTeamSummaryInfo.md) |  | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**Transition** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]BTTaskUserSummaryInfo**](BTTaskUserSummaryInfo.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkflowInfo** | Pointer to [**BTWorkflowableObjectInfo**](BTWorkflowableObjectInfo.md) |  | [optional] 
**WorkflowableObjectType** | Pointer to **int32** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTTaskInfo

`func NewBTTaskInfo() *BTTaskInfo`

NewBTTaskInfo instantiates a new BTTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTaskInfoWithDefaults

`func NewBTTaskInfoWithDefaults() *BTTaskInfo`

NewBTTaskInfoWithDefaults instantiates a new BTTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BTTaskInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BTTaskInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BTTaskInfo) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BTTaskInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApproverRole

`func (o *BTTaskInfo) GetApproverRole() string`

GetApproverRole returns the ApproverRole field if non-nil, zero value otherwise.

### GetApproverRoleOk

`func (o *BTTaskInfo) GetApproverRoleOk() (*string, bool)`

GetApproverRoleOk returns a tuple with the ApproverRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRole

`func (o *BTTaskInfo) SetApproverRole(v string)`

SetApproverRole sets ApproverRole field to given value.

### HasApproverRole

`func (o *BTTaskInfo) HasApproverRole() bool`

HasApproverRole returns a boolean if a field has been set.

### GetComments

`func (o *BTTaskInfo) GetComments() []BTCommentInfo`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BTTaskInfo) GetCommentsOk() (*[]BTCommentInfo, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BTTaskInfo) SetComments(v []BTCommentInfo)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BTTaskInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTTaskInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTTaskInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTTaskInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTTaskInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTTaskInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTTaskInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTTaskInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTTaskInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BTTaskInfo) GetCreator() BTUserSummaryInfo`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BTTaskInfo) GetCreatorOk() (*BTUserSummaryInfo, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BTTaskInfo) SetCreator(v BTUserSummaryInfo)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BTTaskInfo) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDeletable

`func (o *BTTaskInfo) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *BTTaskInfo) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *BTTaskInfo) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *BTTaskInfo) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDescription

`func (o *BTTaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTTaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTTaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTTaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTTaskInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTTaskInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTTaskInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTTaskInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentName

`func (o *BTTaskInfo) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *BTTaskInfo) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *BTTaskInfo) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *BTTaskInfo) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetDocumentType

`func (o *BTTaskInfo) GetDocumentType() int32`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *BTTaskInfo) GetDocumentTypeOk() (*int32, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *BTTaskInfo) SetDocumentType(v int32)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *BTTaskInfo) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetEditable

`func (o *BTTaskInfo) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *BTTaskInfo) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *BTTaskInfo) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *BTTaskInfo) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetElementId

`func (o *BTTaskInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTTaskInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTTaskInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTTaskInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetHref

`func (o *BTTaskInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTaskInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTaskInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTaskInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *BTTaskInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTTaskInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTTaskInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTTaskInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetProperties

`func (o *BTTaskInfo) GetProperties() []BTMetadataPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTTaskInfo) GetPropertiesOk() (*[]BTMetadataPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTTaskInfo) SetProperties(v []BTMetadataPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTTaskInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResolvedAt

`func (o *BTTaskInfo) GetResolvedAt() JSONTime`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *BTTaskInfo) GetResolvedAtOk() (*JSONTime, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *BTTaskInfo) SetResolvedAt(v JSONTime)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *BTTaskInfo) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetResolvedBy

`func (o *BTTaskInfo) GetResolvedBy() BTUserSummaryInfo`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *BTTaskInfo) GetResolvedByOk() (*BTUserSummaryInfo, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *BTTaskInfo) SetResolvedBy(v BTUserSummaryInfo)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *BTTaskInfo) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### GetRoles

`func (o *BTTaskInfo) GetRoles() []BTTaskRbacRoleInfo`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BTTaskInfo) GetRolesOk() (*[]BTTaskRbacRoleInfo, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BTTaskInfo) SetRoles(v []BTTaskRbacRoleInfo)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BTTaskInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSimpleName

`func (o *BTTaskInfo) GetSimpleName() string`

GetSimpleName returns the SimpleName field if non-nil, zero value otherwise.

### GetSimpleNameOk

`func (o *BTTaskInfo) GetSimpleNameOk() (*string, bool)`

GetSimpleNameOk returns a tuple with the SimpleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleName

`func (o *BTTaskInfo) SetSimpleName(v string)`

SetSimpleName sets SimpleName field to given value.

### HasSimpleName

`func (o *BTTaskInfo) HasSimpleName() bool`

HasSimpleName returns a boolean if a field has been set.

### GetSourceWorkspaceOrVersionName

`func (o *BTTaskInfo) GetSourceWorkspaceOrVersionName() string`

GetSourceWorkspaceOrVersionName returns the SourceWorkspaceOrVersionName field if non-nil, zero value otherwise.

### GetSourceWorkspaceOrVersionNameOk

`func (o *BTTaskInfo) GetSourceWorkspaceOrVersionNameOk() (*string, bool)`

GetSourceWorkspaceOrVersionNameOk returns a tuple with the SourceWorkspaceOrVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWorkspaceOrVersionName

`func (o *BTTaskInfo) SetSourceWorkspaceOrVersionName(v string)`

SetSourceWorkspaceOrVersionName sets SourceWorkspaceOrVersionName field to given value.

### HasSourceWorkspaceOrVersionName

`func (o *BTTaskInfo) HasSourceWorkspaceOrVersionName() bool`

HasSourceWorkspaceOrVersionName returns a boolean if a field has been set.

### GetState

`func (o *BTTaskInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTTaskInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTTaskInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BTTaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *BTTaskInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BTTaskInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BTTaskInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BTTaskInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskItems

`func (o *BTTaskInfo) GetTaskItems() []BTTaskItemInfo`

GetTaskItems returns the TaskItems field if non-nil, zero value otherwise.

### GetTaskItemsOk

`func (o *BTTaskInfo) GetTaskItemsOk() (*[]BTTaskItemInfo, bool)`

GetTaskItemsOk returns a tuple with the TaskItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskItems

`func (o *BTTaskInfo) SetTaskItems(v []BTTaskItemInfo)`

SetTaskItems sets TaskItems field to given value.

### HasTaskItems

`func (o *BTTaskInfo) HasTaskItems() bool`

HasTaskItems returns a boolean if a field has been set.

### GetTaskTemplate

`func (o *BTTaskInfo) GetTaskTemplate() int32`

GetTaskTemplate returns the TaskTemplate field if non-nil, zero value otherwise.

### GetTaskTemplateOk

`func (o *BTTaskInfo) GetTaskTemplateOk() (*int32, bool)`

GetTaskTemplateOk returns a tuple with the TaskTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTemplate

`func (o *BTTaskInfo) SetTaskTemplate(v int32)`

SetTaskTemplate sets TaskTemplate field to given value.

### HasTaskTemplate

`func (o *BTTaskInfo) HasTaskTemplate() bool`

HasTaskTemplate returns a boolean if a field has been set.

### GetTaskType

`func (o *BTTaskInfo) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *BTTaskInfo) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *BTTaskInfo) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *BTTaskInfo) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetTeams

`func (o *BTTaskInfo) GetTeams() []BTTaskTeamSummaryInfo`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *BTTaskInfo) GetTeamsOk() (*[]BTTaskTeamSummaryInfo, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *BTTaskInfo) SetTeams(v []BTTaskTeamSummaryInfo)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *BTTaskInfo) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTTaskInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTTaskInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTTaskInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTTaskInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTransition

`func (o *BTTaskInfo) GetTransition() string`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *BTTaskInfo) GetTransitionOk() (*string, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *BTTaskInfo) SetTransition(v string)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *BTTaskInfo) HasTransition() bool`

HasTransition returns a boolean if a field has been set.

### GetUsers

`func (o *BTTaskInfo) GetUsers() []BTTaskUserSummaryInfo`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BTTaskInfo) GetUsersOk() (*[]BTTaskUserSummaryInfo, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BTTaskInfo) SetUsers(v []BTTaskUserSummaryInfo)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BTTaskInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetVersionId

`func (o *BTTaskInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTTaskInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTTaskInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTTaskInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTaskInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTaskInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTaskInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTaskInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *BTTaskInfo) GetWorkflowInfo() BTWorkflowableObjectInfo`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *BTTaskInfo) GetWorkflowInfoOk() (*BTWorkflowableObjectInfo, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *BTTaskInfo) SetWorkflowInfo(v BTWorkflowableObjectInfo)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *BTTaskInfo) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### GetWorkflowableObjectType

`func (o *BTTaskInfo) GetWorkflowableObjectType() int32`

GetWorkflowableObjectType returns the WorkflowableObjectType field if non-nil, zero value otherwise.

### GetWorkflowableObjectTypeOk

`func (o *BTTaskInfo) GetWorkflowableObjectTypeOk() (*int32, bool)`

GetWorkflowableObjectTypeOk returns a tuple with the WorkflowableObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowableObjectType

`func (o *BTTaskInfo) SetWorkflowableObjectType(v int32)`

SetWorkflowableObjectType sets WorkflowableObjectType field to given value.

### HasWorkflowableObjectType

`func (o *BTTaskInfo) HasWorkflowableObjectType() bool`

HasWorkflowableObjectType returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTTaskInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTTaskInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTTaskInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTTaskInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


