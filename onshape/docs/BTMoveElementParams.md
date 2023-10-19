# BTMoveElementParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnchorElementId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementOriginalToNewMap** | Pointer to **map[string]string** |  | [optional] 
**Elements** | Pointer to **[]string** |  | [optional] 
**GenerateUnknownMessages** | Pointer to **bool** |  | [optional] 
**ImportData** | Pointer to **[]string** |  | [optional] 
**IsCopy** | Pointer to **bool** |  | [optional] 
**IsDeepCopy** | Pointer to **bool** |  | [optional] 
**IsGroupAnchor** | Pointer to **bool** |  | [optional] 
**IsNewDocument** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**IsSelectivePartOut** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedNewVersion** | Pointer to **bool** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**SelectedGroupIds** | Pointer to **[]string** |  | [optional] 
**SourceDocumentId** | Pointer to **string** |  | [optional] 
**SourceVersionId** | Pointer to **string** |  | [optional] 
**SourceWorkspaceId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TargetDocumentId** | Pointer to **string** |  | [optional] 
**TargetWorkspaceId** | Pointer to **string** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMoveElementParams

`func NewBTMoveElementParams() *BTMoveElementParams`

NewBTMoveElementParams instantiates a new BTMoveElementParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMoveElementParamsWithDefaults

`func NewBTMoveElementParamsWithDefaults() *BTMoveElementParams`

NewBTMoveElementParamsWithDefaults instantiates a new BTMoveElementParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchorElementId

`func (o *BTMoveElementParams) GetAnchorElementId() string`

GetAnchorElementId returns the AnchorElementId field if non-nil, zero value otherwise.

### GetAnchorElementIdOk

`func (o *BTMoveElementParams) GetAnchorElementIdOk() (*string, bool)`

GetAnchorElementIdOk returns a tuple with the AnchorElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorElementId

`func (o *BTMoveElementParams) SetAnchorElementId(v string)`

SetAnchorElementId sets AnchorElementId field to given value.

### HasAnchorElementId

`func (o *BTMoveElementParams) HasAnchorElementId() bool`

HasAnchorElementId returns a boolean if a field has been set.

### GetDescription

`func (o *BTMoveElementParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTMoveElementParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTMoveElementParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTMoveElementParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementOriginalToNewMap

`func (o *BTMoveElementParams) GetElementOriginalToNewMap() map[string]string`

GetElementOriginalToNewMap returns the ElementOriginalToNewMap field if non-nil, zero value otherwise.

### GetElementOriginalToNewMapOk

`func (o *BTMoveElementParams) GetElementOriginalToNewMapOk() (*map[string]string, bool)`

GetElementOriginalToNewMapOk returns a tuple with the ElementOriginalToNewMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementOriginalToNewMap

`func (o *BTMoveElementParams) SetElementOriginalToNewMap(v map[string]string)`

SetElementOriginalToNewMap sets ElementOriginalToNewMap field to given value.

### HasElementOriginalToNewMap

`func (o *BTMoveElementParams) HasElementOriginalToNewMap() bool`

HasElementOriginalToNewMap returns a boolean if a field has been set.

### GetElements

`func (o *BTMoveElementParams) GetElements() []string`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *BTMoveElementParams) GetElementsOk() (*[]string, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *BTMoveElementParams) SetElements(v []string)`

SetElements sets Elements field to given value.

### HasElements

`func (o *BTMoveElementParams) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetGenerateUnknownMessages

`func (o *BTMoveElementParams) GetGenerateUnknownMessages() bool`

GetGenerateUnknownMessages returns the GenerateUnknownMessages field if non-nil, zero value otherwise.

### GetGenerateUnknownMessagesOk

`func (o *BTMoveElementParams) GetGenerateUnknownMessagesOk() (*bool, bool)`

GetGenerateUnknownMessagesOk returns a tuple with the GenerateUnknownMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateUnknownMessages

`func (o *BTMoveElementParams) SetGenerateUnknownMessages(v bool)`

SetGenerateUnknownMessages sets GenerateUnknownMessages field to given value.

### HasGenerateUnknownMessages

`func (o *BTMoveElementParams) HasGenerateUnknownMessages() bool`

HasGenerateUnknownMessages returns a boolean if a field has been set.

### GetImportData

`func (o *BTMoveElementParams) GetImportData() []string`

GetImportData returns the ImportData field if non-nil, zero value otherwise.

### GetImportDataOk

`func (o *BTMoveElementParams) GetImportDataOk() (*[]string, bool)`

GetImportDataOk returns a tuple with the ImportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportData

`func (o *BTMoveElementParams) SetImportData(v []string)`

SetImportData sets ImportData field to given value.

### HasImportData

`func (o *BTMoveElementParams) HasImportData() bool`

HasImportData returns a boolean if a field has been set.

### GetIsCopy

`func (o *BTMoveElementParams) GetIsCopy() bool`

GetIsCopy returns the IsCopy field if non-nil, zero value otherwise.

### GetIsCopyOk

`func (o *BTMoveElementParams) GetIsCopyOk() (*bool, bool)`

GetIsCopyOk returns a tuple with the IsCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopy

`func (o *BTMoveElementParams) SetIsCopy(v bool)`

SetIsCopy sets IsCopy field to given value.

### HasIsCopy

`func (o *BTMoveElementParams) HasIsCopy() bool`

HasIsCopy returns a boolean if a field has been set.

### GetIsDeepCopy

`func (o *BTMoveElementParams) GetIsDeepCopy() bool`

GetIsDeepCopy returns the IsDeepCopy field if non-nil, zero value otherwise.

### GetIsDeepCopyOk

`func (o *BTMoveElementParams) GetIsDeepCopyOk() (*bool, bool)`

GetIsDeepCopyOk returns a tuple with the IsDeepCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeepCopy

`func (o *BTMoveElementParams) SetIsDeepCopy(v bool)`

SetIsDeepCopy sets IsDeepCopy field to given value.

### HasIsDeepCopy

`func (o *BTMoveElementParams) HasIsDeepCopy() bool`

HasIsDeepCopy returns a boolean if a field has been set.

### GetIsGroupAnchor

`func (o *BTMoveElementParams) GetIsGroupAnchor() bool`

GetIsGroupAnchor returns the IsGroupAnchor field if non-nil, zero value otherwise.

### GetIsGroupAnchorOk

`func (o *BTMoveElementParams) GetIsGroupAnchorOk() (*bool, bool)`

GetIsGroupAnchorOk returns a tuple with the IsGroupAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroupAnchor

`func (o *BTMoveElementParams) SetIsGroupAnchor(v bool)`

SetIsGroupAnchor sets IsGroupAnchor field to given value.

### HasIsGroupAnchor

`func (o *BTMoveElementParams) HasIsGroupAnchor() bool`

HasIsGroupAnchor returns a boolean if a field has been set.

### GetIsNewDocument

`func (o *BTMoveElementParams) GetIsNewDocument() bool`

GetIsNewDocument returns the IsNewDocument field if non-nil, zero value otherwise.

### GetIsNewDocumentOk

`func (o *BTMoveElementParams) GetIsNewDocumentOk() (*bool, bool)`

GetIsNewDocumentOk returns a tuple with the IsNewDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewDocument

`func (o *BTMoveElementParams) SetIsNewDocument(v bool)`

SetIsNewDocument sets IsNewDocument field to given value.

### HasIsNewDocument

`func (o *BTMoveElementParams) HasIsNewDocument() bool`

HasIsNewDocument returns a boolean if a field has been set.

### GetIsPublic

`func (o *BTMoveElementParams) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *BTMoveElementParams) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *BTMoveElementParams) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *BTMoveElementParams) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsSelectivePartOut

`func (o *BTMoveElementParams) GetIsSelectivePartOut() bool`

GetIsSelectivePartOut returns the IsSelectivePartOut field if non-nil, zero value otherwise.

### GetIsSelectivePartOutOk

`func (o *BTMoveElementParams) GetIsSelectivePartOutOk() (*bool, bool)`

GetIsSelectivePartOutOk returns a tuple with the IsSelectivePartOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelectivePartOut

`func (o *BTMoveElementParams) SetIsSelectivePartOut(v bool)`

SetIsSelectivePartOut sets IsSelectivePartOut field to given value.

### HasIsSelectivePartOut

`func (o *BTMoveElementParams) HasIsSelectivePartOut() bool`

HasIsSelectivePartOut returns a boolean if a field has been set.

### GetName

`func (o *BTMoveElementParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMoveElementParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMoveElementParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMoveElementParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedNewVersion

`func (o *BTMoveElementParams) GetNeedNewVersion() bool`

GetNeedNewVersion returns the NeedNewVersion field if non-nil, zero value otherwise.

### GetNeedNewVersionOk

`func (o *BTMoveElementParams) GetNeedNewVersionOk() (*bool, bool)`

GetNeedNewVersionOk returns a tuple with the NeedNewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedNewVersion

`func (o *BTMoveElementParams) SetNeedNewVersion(v bool)`

SetNeedNewVersion sets NeedNewVersion field to given value.

### HasNeedNewVersion

`func (o *BTMoveElementParams) HasNeedNewVersion() bool`

HasNeedNewVersion returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *BTMoveElementParams) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *BTMoveElementParams) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *BTMoveElementParams) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *BTMoveElementParams) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTMoveElementParams) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTMoveElementParams) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTMoveElementParams) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTMoveElementParams) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTMoveElementParams) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTMoveElementParams) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTMoveElementParams) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTMoveElementParams) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetParentId

`func (o *BTMoveElementParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTMoveElementParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTMoveElementParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTMoveElementParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTMoveElementParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTMoveElementParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTMoveElementParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTMoveElementParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSelectedGroupIds

`func (o *BTMoveElementParams) GetSelectedGroupIds() []string`

GetSelectedGroupIds returns the SelectedGroupIds field if non-nil, zero value otherwise.

### GetSelectedGroupIdsOk

`func (o *BTMoveElementParams) GetSelectedGroupIdsOk() (*[]string, bool)`

GetSelectedGroupIdsOk returns a tuple with the SelectedGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedGroupIds

`func (o *BTMoveElementParams) SetSelectedGroupIds(v []string)`

SetSelectedGroupIds sets SelectedGroupIds field to given value.

### HasSelectedGroupIds

`func (o *BTMoveElementParams) HasSelectedGroupIds() bool`

HasSelectedGroupIds returns a boolean if a field has been set.

### GetSourceDocumentId

`func (o *BTMoveElementParams) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *BTMoveElementParams) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *BTMoveElementParams) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *BTMoveElementParams) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### GetSourceVersionId

`func (o *BTMoveElementParams) GetSourceVersionId() string`

GetSourceVersionId returns the SourceVersionId field if non-nil, zero value otherwise.

### GetSourceVersionIdOk

`func (o *BTMoveElementParams) GetSourceVersionIdOk() (*string, bool)`

GetSourceVersionIdOk returns a tuple with the SourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersionId

`func (o *BTMoveElementParams) SetSourceVersionId(v string)`

SetSourceVersionId sets SourceVersionId field to given value.

### HasSourceVersionId

`func (o *BTMoveElementParams) HasSourceVersionId() bool`

HasSourceVersionId returns a boolean if a field has been set.

### GetSourceWorkspaceId

`func (o *BTMoveElementParams) GetSourceWorkspaceId() string`

GetSourceWorkspaceId returns the SourceWorkspaceId field if non-nil, zero value otherwise.

### GetSourceWorkspaceIdOk

`func (o *BTMoveElementParams) GetSourceWorkspaceIdOk() (*string, bool)`

GetSourceWorkspaceIdOk returns a tuple with the SourceWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWorkspaceId

`func (o *BTMoveElementParams) SetSourceWorkspaceId(v string)`

SetSourceWorkspaceId sets SourceWorkspaceId field to given value.

### HasSourceWorkspaceId

`func (o *BTMoveElementParams) HasSourceWorkspaceId() bool`

HasSourceWorkspaceId returns a boolean if a field has been set.

### GetTags

`func (o *BTMoveElementParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTMoveElementParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTMoveElementParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTMoveElementParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetDocumentId

`func (o *BTMoveElementParams) GetTargetDocumentId() string`

GetTargetDocumentId returns the TargetDocumentId field if non-nil, zero value otherwise.

### GetTargetDocumentIdOk

`func (o *BTMoveElementParams) GetTargetDocumentIdOk() (*string, bool)`

GetTargetDocumentIdOk returns a tuple with the TargetDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocumentId

`func (o *BTMoveElementParams) SetTargetDocumentId(v string)`

SetTargetDocumentId sets TargetDocumentId field to given value.

### HasTargetDocumentId

`func (o *BTMoveElementParams) HasTargetDocumentId() bool`

HasTargetDocumentId returns a boolean if a field has been set.

### GetTargetWorkspaceId

`func (o *BTMoveElementParams) GetTargetWorkspaceId() string`

GetTargetWorkspaceId returns the TargetWorkspaceId field if non-nil, zero value otherwise.

### GetTargetWorkspaceIdOk

`func (o *BTMoveElementParams) GetTargetWorkspaceIdOk() (*string, bool)`

GetTargetWorkspaceIdOk returns a tuple with the TargetWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWorkspaceId

`func (o *BTMoveElementParams) SetTargetWorkspaceId(v string)`

SetTargetWorkspaceId sets TargetWorkspaceId field to given value.

### HasTargetWorkspaceId

`func (o *BTMoveElementParams) HasTargetWorkspaceId() bool`

HasTargetWorkspaceId returns a boolean if a field has been set.

### GetVersionName

`func (o *BTMoveElementParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *BTMoveElementParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *BTMoveElementParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *BTMoveElementParams) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


