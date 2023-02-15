# BTCommentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyFeatures** | Pointer to **[]string** |  | [optional] 
**AssignedAt** | Pointer to **JSONTime** |  | [optional] 
**Assignee** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Attachment** | Pointer to [**BTCommentAttachmentInfo**](BTCommentAttachmentInfo.md) |  | [optional] 
**CanDelete** | Pointer to **bool** |  | [optional] 
**CanResolveOrReopen** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementFeature** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementOccurrences** | Pointer to **[]string** |  | [optional] 
**ElementQuery** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ReleasePackageId** | Pointer to **string** |  | [optional] 
**ReopenedAt** | Pointer to **JSONTime** |  | [optional] 
**ReopenedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**ReplyCount** | Pointer to **int64** |  | [optional] 
**ResolvedAt** | Pointer to **JSONTime** |  | [optional] 
**ResolvedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**Thumbnail** | Pointer to [**BTCommentAttachmentInfo**](BTCommentAttachmentInfo.md) |  | [optional] 
**TopLevel** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ViewData** | Pointer to [**BTViewDataInfo**](BTViewDataInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTCommentInfo

`func NewBTCommentInfo() *BTCommentInfo`

NewBTCommentInfo instantiates a new BTCommentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCommentInfoWithDefaults

`func NewBTCommentInfoWithDefaults() *BTCommentInfo`

NewBTCommentInfoWithDefaults instantiates a new BTCommentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyFeatures

`func (o *BTCommentInfo) GetAssemblyFeatures() []string`

GetAssemblyFeatures returns the AssemblyFeatures field if non-nil, zero value otherwise.

### GetAssemblyFeaturesOk

`func (o *BTCommentInfo) GetAssemblyFeaturesOk() (*[]string, bool)`

GetAssemblyFeaturesOk returns a tuple with the AssemblyFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyFeatures

`func (o *BTCommentInfo) SetAssemblyFeatures(v []string)`

SetAssemblyFeatures sets AssemblyFeatures field to given value.

### HasAssemblyFeatures

`func (o *BTCommentInfo) HasAssemblyFeatures() bool`

HasAssemblyFeatures returns a boolean if a field has been set.

### GetAssignedAt

`func (o *BTCommentInfo) GetAssignedAt() JSONTime`

GetAssignedAt returns the AssignedAt field if non-nil, zero value otherwise.

### GetAssignedAtOk

`func (o *BTCommentInfo) GetAssignedAtOk() (*JSONTime, bool)`

GetAssignedAtOk returns a tuple with the AssignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAt

`func (o *BTCommentInfo) SetAssignedAt(v JSONTime)`

SetAssignedAt sets AssignedAt field to given value.

### HasAssignedAt

`func (o *BTCommentInfo) HasAssignedAt() bool`

HasAssignedAt returns a boolean if a field has been set.

### GetAssignee

`func (o *BTCommentInfo) GetAssignee() BTUserSummaryInfo`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *BTCommentInfo) GetAssigneeOk() (*BTUserSummaryInfo, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *BTCommentInfo) SetAssignee(v BTUserSummaryInfo)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *BTCommentInfo) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAttachment

`func (o *BTCommentInfo) GetAttachment() BTCommentAttachmentInfo`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *BTCommentInfo) GetAttachmentOk() (*BTCommentAttachmentInfo, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *BTCommentInfo) SetAttachment(v BTCommentAttachmentInfo)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *BTCommentInfo) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetCanDelete

`func (o *BTCommentInfo) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *BTCommentInfo) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *BTCommentInfo) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *BTCommentInfo) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### GetCanResolveOrReopen

`func (o *BTCommentInfo) GetCanResolveOrReopen() bool`

GetCanResolveOrReopen returns the CanResolveOrReopen field if non-nil, zero value otherwise.

### GetCanResolveOrReopenOk

`func (o *BTCommentInfo) GetCanResolveOrReopenOk() (*bool, bool)`

GetCanResolveOrReopenOk returns a tuple with the CanResolveOrReopen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResolveOrReopen

`func (o *BTCommentInfo) SetCanResolveOrReopen(v bool)`

SetCanResolveOrReopen sets CanResolveOrReopen field to given value.

### HasCanResolveOrReopen

`func (o *BTCommentInfo) HasCanResolveOrReopen() bool`

HasCanResolveOrReopen returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTCommentInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTCommentInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTCommentInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTCommentInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTCommentInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTCommentInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTCommentInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTCommentInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementFeature

`func (o *BTCommentInfo) GetElementFeature() string`

GetElementFeature returns the ElementFeature field if non-nil, zero value otherwise.

### GetElementFeatureOk

`func (o *BTCommentInfo) GetElementFeatureOk() (*string, bool)`

GetElementFeatureOk returns a tuple with the ElementFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementFeature

`func (o *BTCommentInfo) SetElementFeature(v string)`

SetElementFeature sets ElementFeature field to given value.

### HasElementFeature

`func (o *BTCommentInfo) HasElementFeature() bool`

HasElementFeature returns a boolean if a field has been set.

### GetElementId

`func (o *BTCommentInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTCommentInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTCommentInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTCommentInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementOccurrences

`func (o *BTCommentInfo) GetElementOccurrences() []string`

GetElementOccurrences returns the ElementOccurrences field if non-nil, zero value otherwise.

### GetElementOccurrencesOk

`func (o *BTCommentInfo) GetElementOccurrencesOk() (*[]string, bool)`

GetElementOccurrencesOk returns a tuple with the ElementOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementOccurrences

`func (o *BTCommentInfo) SetElementOccurrences(v []string)`

SetElementOccurrences sets ElementOccurrences field to given value.

### HasElementOccurrences

`func (o *BTCommentInfo) HasElementOccurrences() bool`

HasElementOccurrences returns a boolean if a field has been set.

### GetElementQuery

`func (o *BTCommentInfo) GetElementQuery() string`

GetElementQuery returns the ElementQuery field if non-nil, zero value otherwise.

### GetElementQueryOk

`func (o *BTCommentInfo) GetElementQueryOk() (*string, bool)`

GetElementQueryOk returns a tuple with the ElementQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementQuery

`func (o *BTCommentInfo) SetElementQuery(v string)`

SetElementQuery sets ElementQuery field to given value.

### HasElementQuery

`func (o *BTCommentInfo) HasElementQuery() bool`

HasElementQuery returns a boolean if a field has been set.

### GetHref

`func (o *BTCommentInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTCommentInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTCommentInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTCommentInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTCommentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTCommentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTCommentInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTCommentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *BTCommentInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BTCommentInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BTCommentInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BTCommentInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *BTCommentInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTCommentInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTCommentInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTCommentInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *BTCommentInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTCommentInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTCommentInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTCommentInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *BTCommentInfo) GetObjectType() int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTCommentInfo) GetObjectTypeOk() (*int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTCommentInfo) SetObjectType(v int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTCommentInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetParentId

`func (o *BTCommentInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTCommentInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTCommentInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTCommentInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReleasePackageId

`func (o *BTCommentInfo) GetReleasePackageId() string`

GetReleasePackageId returns the ReleasePackageId field if non-nil, zero value otherwise.

### GetReleasePackageIdOk

`func (o *BTCommentInfo) GetReleasePackageIdOk() (*string, bool)`

GetReleasePackageIdOk returns a tuple with the ReleasePackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePackageId

`func (o *BTCommentInfo) SetReleasePackageId(v string)`

SetReleasePackageId sets ReleasePackageId field to given value.

### HasReleasePackageId

`func (o *BTCommentInfo) HasReleasePackageId() bool`

HasReleasePackageId returns a boolean if a field has been set.

### GetReopenedAt

`func (o *BTCommentInfo) GetReopenedAt() JSONTime`

GetReopenedAt returns the ReopenedAt field if non-nil, zero value otherwise.

### GetReopenedAtOk

`func (o *BTCommentInfo) GetReopenedAtOk() (*JSONTime, bool)`

GetReopenedAtOk returns a tuple with the ReopenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedAt

`func (o *BTCommentInfo) SetReopenedAt(v JSONTime)`

SetReopenedAt sets ReopenedAt field to given value.

### HasReopenedAt

`func (o *BTCommentInfo) HasReopenedAt() bool`

HasReopenedAt returns a boolean if a field has been set.

### GetReopenedBy

`func (o *BTCommentInfo) GetReopenedBy() BTUserSummaryInfo`

GetReopenedBy returns the ReopenedBy field if non-nil, zero value otherwise.

### GetReopenedByOk

`func (o *BTCommentInfo) GetReopenedByOk() (*BTUserSummaryInfo, bool)`

GetReopenedByOk returns a tuple with the ReopenedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedBy

`func (o *BTCommentInfo) SetReopenedBy(v BTUserSummaryInfo)`

SetReopenedBy sets ReopenedBy field to given value.

### HasReopenedBy

`func (o *BTCommentInfo) HasReopenedBy() bool`

HasReopenedBy returns a boolean if a field has been set.

### GetReplyCount

`func (o *BTCommentInfo) GetReplyCount() int64`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *BTCommentInfo) GetReplyCountOk() (*int64, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *BTCommentInfo) SetReplyCount(v int64)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *BTCommentInfo) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetResolvedAt

`func (o *BTCommentInfo) GetResolvedAt() JSONTime`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *BTCommentInfo) GetResolvedAtOk() (*JSONTime, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *BTCommentInfo) SetResolvedAt(v JSONTime)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *BTCommentInfo) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetResolvedBy

`func (o *BTCommentInfo) GetResolvedBy() BTUserSummaryInfo`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *BTCommentInfo) GetResolvedByOk() (*BTUserSummaryInfo, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *BTCommentInfo) SetResolvedBy(v BTUserSummaryInfo)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *BTCommentInfo) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### GetState

`func (o *BTCommentInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTCommentInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTCommentInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTCommentInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTCommentInfo) GetThumbnail() BTCommentAttachmentInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTCommentInfo) GetThumbnailOk() (*BTCommentAttachmentInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTCommentInfo) SetThumbnail(v BTCommentAttachmentInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTCommentInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTopLevel

`func (o *BTCommentInfo) GetTopLevel() bool`

GetTopLevel returns the TopLevel field if non-nil, zero value otherwise.

### GetTopLevelOk

`func (o *BTCommentInfo) GetTopLevelOk() (*bool, bool)`

GetTopLevelOk returns a tuple with the TopLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevel

`func (o *BTCommentInfo) SetTopLevel(v bool)`

SetTopLevel sets TopLevel field to given value.

### HasTopLevel

`func (o *BTCommentInfo) HasTopLevel() bool`

HasTopLevel returns a boolean if a field has been set.

### GetUser

`func (o *BTCommentInfo) GetUser() BTUserSummaryInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BTCommentInfo) GetUserOk() (*BTUserSummaryInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BTCommentInfo) SetUser(v BTUserSummaryInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *BTCommentInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersionId

`func (o *BTCommentInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTCommentInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTCommentInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTCommentInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetViewData

`func (o *BTCommentInfo) GetViewData() BTViewDataInfo`

GetViewData returns the ViewData field if non-nil, zero value otherwise.

### GetViewDataOk

`func (o *BTCommentInfo) GetViewDataOk() (*BTViewDataInfo, bool)`

GetViewDataOk returns a tuple with the ViewData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewData

`func (o *BTCommentInfo) SetViewData(v BTViewDataInfo)`

SetViewData sets ViewData field to given value.

### HasViewData

`func (o *BTCommentInfo) HasViewData() bool`

HasViewData returns a boolean if a field has been set.

### GetViewRef

`func (o *BTCommentInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTCommentInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTCommentInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTCommentInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTCommentInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTCommentInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTCommentInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTCommentInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


