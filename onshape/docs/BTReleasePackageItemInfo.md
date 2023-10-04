# BTReleasePackageItemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanExport** | Pointer to **bool** |  | [optional] 
**ChangeDetectionStatus** | Pointer to **int32** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**ConfigurationKey** | Pointer to **string** |  | [optional] 
**DiffThumbnailConfigurationKey** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]BTReleaseItemErrorInfo**](BTReleaseItemErrorInfo.md) |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsRevisionManaged** | Pointer to **bool** |  | [optional] 
**IsRootItem** | Pointer to **bool** |  | [optional] 
**IsTranslatable** | Pointer to **bool** |  | [optional] 
**MeshState** | Pointer to **int32** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObsoletionRevisionId** | Pointer to **string** |  | [optional] 
**OriginalWorkspaceId** | Pointer to **string** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]BTMetadataPropertyInfo**](BTMetadataPropertyInfo.md) |  | [optional] 
**ReferenceIds** | Pointer to **[]string** |  | [optional] 
**ReferenceIdsFromOriginalWorkspace** | Pointer to **[]string** |  | [optional] 
**Rpid** | Pointer to **string** |  | [optional] 
**SmallThumbnailHref** | Pointer to **string** |  | [optional] 
**SyncedWithPLM** | Pointer to **bool** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTReleasePackageItemInfo

`func NewBTReleasePackageItemInfo() *BTReleasePackageItemInfo`

NewBTReleasePackageItemInfo instantiates a new BTReleasePackageItemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTReleasePackageItemInfoWithDefaults

`func NewBTReleasePackageItemInfoWithDefaults() *BTReleasePackageItemInfo`

NewBTReleasePackageItemInfoWithDefaults instantiates a new BTReleasePackageItemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanExport

`func (o *BTReleasePackageItemInfo) GetCanExport() bool`

GetCanExport returns the CanExport field if non-nil, zero value otherwise.

### GetCanExportOk

`func (o *BTReleasePackageItemInfo) GetCanExportOk() (*bool, bool)`

GetCanExportOk returns a tuple with the CanExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanExport

`func (o *BTReleasePackageItemInfo) SetCanExport(v bool)`

SetCanExport sets CanExport field to given value.

### HasCanExport

`func (o *BTReleasePackageItemInfo) HasCanExport() bool`

HasCanExport returns a boolean if a field has been set.

### GetChangeDetectionStatus

`func (o *BTReleasePackageItemInfo) GetChangeDetectionStatus() int32`

GetChangeDetectionStatus returns the ChangeDetectionStatus field if non-nil, zero value otherwise.

### GetChangeDetectionStatusOk

`func (o *BTReleasePackageItemInfo) GetChangeDetectionStatusOk() (*int32, bool)`

GetChangeDetectionStatusOk returns a tuple with the ChangeDetectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDetectionStatus

`func (o *BTReleasePackageItemInfo) SetChangeDetectionStatus(v int32)`

SetChangeDetectionStatus sets ChangeDetectionStatus field to given value.

### HasChangeDetectionStatus

`func (o *BTReleasePackageItemInfo) HasChangeDetectionStatus() bool`

HasChangeDetectionStatus returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTReleasePackageItemInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTReleasePackageItemInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTReleasePackageItemInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTReleasePackageItemInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTReleasePackageItemInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTReleasePackageItemInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTReleasePackageItemInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTReleasePackageItemInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConfigurationKey

`func (o *BTReleasePackageItemInfo) GetConfigurationKey() string`

GetConfigurationKey returns the ConfigurationKey field if non-nil, zero value otherwise.

### GetConfigurationKeyOk

`func (o *BTReleasePackageItemInfo) GetConfigurationKeyOk() (*string, bool)`

GetConfigurationKeyOk returns a tuple with the ConfigurationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationKey

`func (o *BTReleasePackageItemInfo) SetConfigurationKey(v string)`

SetConfigurationKey sets ConfigurationKey field to given value.

### HasConfigurationKey

`func (o *BTReleasePackageItemInfo) HasConfigurationKey() bool`

HasConfigurationKey returns a boolean if a field has been set.

### GetDiffThumbnailConfigurationKey

`func (o *BTReleasePackageItemInfo) GetDiffThumbnailConfigurationKey() string`

GetDiffThumbnailConfigurationKey returns the DiffThumbnailConfigurationKey field if non-nil, zero value otherwise.

### GetDiffThumbnailConfigurationKeyOk

`func (o *BTReleasePackageItemInfo) GetDiffThumbnailConfigurationKeyOk() (*string, bool)`

GetDiffThumbnailConfigurationKeyOk returns a tuple with the DiffThumbnailConfigurationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffThumbnailConfigurationKey

`func (o *BTReleasePackageItemInfo) SetDiffThumbnailConfigurationKey(v string)`

SetDiffThumbnailConfigurationKey sets DiffThumbnailConfigurationKey field to given value.

### HasDiffThumbnailConfigurationKey

`func (o *BTReleasePackageItemInfo) HasDiffThumbnailConfigurationKey() bool`

HasDiffThumbnailConfigurationKey returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTReleasePackageItemInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTReleasePackageItemInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTReleasePackageItemInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTReleasePackageItemInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTReleasePackageItemInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTReleasePackageItemInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTReleasePackageItemInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTReleasePackageItemInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTReleasePackageItemInfo) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTReleasePackageItemInfo) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTReleasePackageItemInfo) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTReleasePackageItemInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetErrors

`func (o *BTReleasePackageItemInfo) GetErrors() []BTReleaseItemErrorInfo`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BTReleasePackageItemInfo) GetErrorsOk() (*[]BTReleaseItemErrorInfo, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BTReleasePackageItemInfo) SetErrors(v []BTReleaseItemErrorInfo)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BTReleasePackageItemInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetHref

`func (o *BTReleasePackageItemInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTReleasePackageItemInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTReleasePackageItemInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTReleasePackageItemInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTReleasePackageItemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTReleasePackageItemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTReleasePackageItemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTReleasePackageItemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRevisionManaged

`func (o *BTReleasePackageItemInfo) GetIsRevisionManaged() bool`

GetIsRevisionManaged returns the IsRevisionManaged field if non-nil, zero value otherwise.

### GetIsRevisionManagedOk

`func (o *BTReleasePackageItemInfo) GetIsRevisionManagedOk() (*bool, bool)`

GetIsRevisionManagedOk returns a tuple with the IsRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevisionManaged

`func (o *BTReleasePackageItemInfo) SetIsRevisionManaged(v bool)`

SetIsRevisionManaged sets IsRevisionManaged field to given value.

### HasIsRevisionManaged

`func (o *BTReleasePackageItemInfo) HasIsRevisionManaged() bool`

HasIsRevisionManaged returns a boolean if a field has been set.

### GetIsRootItem

`func (o *BTReleasePackageItemInfo) GetIsRootItem() bool`

GetIsRootItem returns the IsRootItem field if non-nil, zero value otherwise.

### GetIsRootItemOk

`func (o *BTReleasePackageItemInfo) GetIsRootItemOk() (*bool, bool)`

GetIsRootItemOk returns a tuple with the IsRootItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootItem

`func (o *BTReleasePackageItemInfo) SetIsRootItem(v bool)`

SetIsRootItem sets IsRootItem field to given value.

### HasIsRootItem

`func (o *BTReleasePackageItemInfo) HasIsRootItem() bool`

HasIsRootItem returns a boolean if a field has been set.

### GetIsTranslatable

`func (o *BTReleasePackageItemInfo) GetIsTranslatable() bool`

GetIsTranslatable returns the IsTranslatable field if non-nil, zero value otherwise.

### GetIsTranslatableOk

`func (o *BTReleasePackageItemInfo) GetIsTranslatableOk() (*bool, bool)`

GetIsTranslatableOk returns a tuple with the IsTranslatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranslatable

`func (o *BTReleasePackageItemInfo) SetIsTranslatable(v bool)`

SetIsTranslatable sets IsTranslatable field to given value.

### HasIsTranslatable

`func (o *BTReleasePackageItemInfo) HasIsTranslatable() bool`

HasIsTranslatable returns a boolean if a field has been set.

### GetMeshState

`func (o *BTReleasePackageItemInfo) GetMeshState() int32`

GetMeshState returns the MeshState field if non-nil, zero value otherwise.

### GetMeshStateOk

`func (o *BTReleasePackageItemInfo) GetMeshStateOk() (*int32, bool)`

GetMeshStateOk returns a tuple with the MeshState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshState

`func (o *BTReleasePackageItemInfo) SetMeshState(v int32)`

SetMeshState sets MeshState field to given value.

### HasMeshState

`func (o *BTReleasePackageItemInfo) HasMeshState() bool`

HasMeshState returns a boolean if a field has been set.

### GetMimeType

`func (o *BTReleasePackageItemInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTReleasePackageItemInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTReleasePackageItemInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTReleasePackageItemInfo) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *BTReleasePackageItemInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTReleasePackageItemInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTReleasePackageItemInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTReleasePackageItemInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObsoletionRevisionId

`func (o *BTReleasePackageItemInfo) GetObsoletionRevisionId() string`

GetObsoletionRevisionId returns the ObsoletionRevisionId field if non-nil, zero value otherwise.

### GetObsoletionRevisionIdOk

`func (o *BTReleasePackageItemInfo) GetObsoletionRevisionIdOk() (*string, bool)`

GetObsoletionRevisionIdOk returns a tuple with the ObsoletionRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionRevisionId

`func (o *BTReleasePackageItemInfo) SetObsoletionRevisionId(v string)`

SetObsoletionRevisionId sets ObsoletionRevisionId field to given value.

### HasObsoletionRevisionId

`func (o *BTReleasePackageItemInfo) HasObsoletionRevisionId() bool`

HasObsoletionRevisionId returns a boolean if a field has been set.

### GetOriginalWorkspaceId

`func (o *BTReleasePackageItemInfo) GetOriginalWorkspaceId() string`

GetOriginalWorkspaceId returns the OriginalWorkspaceId field if non-nil, zero value otherwise.

### GetOriginalWorkspaceIdOk

`func (o *BTReleasePackageItemInfo) GetOriginalWorkspaceIdOk() (*string, bool)`

GetOriginalWorkspaceIdOk returns a tuple with the OriginalWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWorkspaceId

`func (o *BTReleasePackageItemInfo) SetOriginalWorkspaceId(v string)`

SetOriginalWorkspaceId sets OriginalWorkspaceId field to given value.

### HasOriginalWorkspaceId

`func (o *BTReleasePackageItemInfo) HasOriginalWorkspaceId() bool`

HasOriginalWorkspaceId returns a boolean if a field has been set.

### GetPartId

`func (o *BTReleasePackageItemInfo) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTReleasePackageItemInfo) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTReleasePackageItemInfo) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTReleasePackageItemInfo) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTReleasePackageItemInfo) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTReleasePackageItemInfo) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTReleasePackageItemInfo) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTReleasePackageItemInfo) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartType

`func (o *BTReleasePackageItemInfo) GetPartType() string`

GetPartType returns the PartType field if non-nil, zero value otherwise.

### GetPartTypeOk

`func (o *BTReleasePackageItemInfo) GetPartTypeOk() (*string, bool)`

GetPartTypeOk returns a tuple with the PartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartType

`func (o *BTReleasePackageItemInfo) SetPartType(v string)`

SetPartType sets PartType field to given value.

### HasPartType

`func (o *BTReleasePackageItemInfo) HasPartType() bool`

HasPartType returns a boolean if a field has been set.

### GetProperties

`func (o *BTReleasePackageItemInfo) GetProperties() []BTMetadataPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTReleasePackageItemInfo) GetPropertiesOk() (*[]BTMetadataPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTReleasePackageItemInfo) SetProperties(v []BTMetadataPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTReleasePackageItemInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetReferenceIds

`func (o *BTReleasePackageItemInfo) GetReferenceIds() []string`

GetReferenceIds returns the ReferenceIds field if non-nil, zero value otherwise.

### GetReferenceIdsOk

`func (o *BTReleasePackageItemInfo) GetReferenceIdsOk() (*[]string, bool)`

GetReferenceIdsOk returns a tuple with the ReferenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIds

`func (o *BTReleasePackageItemInfo) SetReferenceIds(v []string)`

SetReferenceIds sets ReferenceIds field to given value.

### HasReferenceIds

`func (o *BTReleasePackageItemInfo) HasReferenceIds() bool`

HasReferenceIds returns a boolean if a field has been set.

### GetReferenceIdsFromOriginalWorkspace

`func (o *BTReleasePackageItemInfo) GetReferenceIdsFromOriginalWorkspace() []string`

GetReferenceIdsFromOriginalWorkspace returns the ReferenceIdsFromOriginalWorkspace field if non-nil, zero value otherwise.

### GetReferenceIdsFromOriginalWorkspaceOk

`func (o *BTReleasePackageItemInfo) GetReferenceIdsFromOriginalWorkspaceOk() (*[]string, bool)`

GetReferenceIdsFromOriginalWorkspaceOk returns a tuple with the ReferenceIdsFromOriginalWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIdsFromOriginalWorkspace

`func (o *BTReleasePackageItemInfo) SetReferenceIdsFromOriginalWorkspace(v []string)`

SetReferenceIdsFromOriginalWorkspace sets ReferenceIdsFromOriginalWorkspace field to given value.

### HasReferenceIdsFromOriginalWorkspace

`func (o *BTReleasePackageItemInfo) HasReferenceIdsFromOriginalWorkspace() bool`

HasReferenceIdsFromOriginalWorkspace returns a boolean if a field has been set.

### GetRpid

`func (o *BTReleasePackageItemInfo) GetRpid() string`

GetRpid returns the Rpid field if non-nil, zero value otherwise.

### GetRpidOk

`func (o *BTReleasePackageItemInfo) GetRpidOk() (*string, bool)`

GetRpidOk returns a tuple with the Rpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpid

`func (o *BTReleasePackageItemInfo) SetRpid(v string)`

SetRpid sets Rpid field to given value.

### HasRpid

`func (o *BTReleasePackageItemInfo) HasRpid() bool`

HasRpid returns a boolean if a field has been set.

### GetSmallThumbnailHref

`func (o *BTReleasePackageItemInfo) GetSmallThumbnailHref() string`

GetSmallThumbnailHref returns the SmallThumbnailHref field if non-nil, zero value otherwise.

### GetSmallThumbnailHrefOk

`func (o *BTReleasePackageItemInfo) GetSmallThumbnailHrefOk() (*string, bool)`

GetSmallThumbnailHrefOk returns a tuple with the SmallThumbnailHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallThumbnailHref

`func (o *BTReleasePackageItemInfo) SetSmallThumbnailHref(v string)`

SetSmallThumbnailHref sets SmallThumbnailHref field to given value.

### HasSmallThumbnailHref

`func (o *BTReleasePackageItemInfo) HasSmallThumbnailHref() bool`

HasSmallThumbnailHref returns a boolean if a field has been set.

### GetSyncedWithPLM

`func (o *BTReleasePackageItemInfo) GetSyncedWithPLM() bool`

GetSyncedWithPLM returns the SyncedWithPLM field if non-nil, zero value otherwise.

### GetSyncedWithPLMOk

`func (o *BTReleasePackageItemInfo) GetSyncedWithPLMOk() (*bool, bool)`

GetSyncedWithPLMOk returns a tuple with the SyncedWithPLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedWithPLM

`func (o *BTReleasePackageItemInfo) SetSyncedWithPLM(v bool)`

SetSyncedWithPLM sets SyncedWithPLM field to given value.

### HasSyncedWithPLM

`func (o *BTReleasePackageItemInfo) HasSyncedWithPLM() bool`

HasSyncedWithPLM returns a boolean if a field has been set.

### GetVersionId

`func (o *BTReleasePackageItemInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTReleasePackageItemInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTReleasePackageItemInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTReleasePackageItemInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTReleasePackageItemInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTReleasePackageItemInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTReleasePackageItemInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTReleasePackageItemInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTReleasePackageItemInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTReleasePackageItemInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTReleasePackageItemInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTReleasePackageItemInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


