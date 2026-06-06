# BTProductStructureItemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonType** | **string** |  | 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**Configuration** | Pointer to [**[]ConfigInfo**](ConfigInfo.md) | The configuration parameters of the referring element. | [optional] 
**DocumentId** | Pointer to **string** | The document ID of the referring element. | [optional] 
**DocumentName** | Pointer to **string** | The name of the document containing the referring element. | [optional] 
**DocumentState** | Pointer to **int32** |  | [optional] 
**ElementId** | Pointer to **string** | The element ID of the referring element. | [optional] 
**ElementType** | Pointer to **int32** | The element type ordinal of the referring element. | [optional] 
**FlattenedBody** | Pointer to **bool** | Whether the item represents a flattened sheet metal body. | [optional] 
**FolderId** | Pointer to **string** | The folder ID containing the document. | [optional] 
**HasDrawing** | Pointer to **bool** | Whether the item has an associated drawing. | [optional] 
**LatestRevision** | Pointer to **bool** | Whether this item is the latest revision. | [optional] 
**NotRevisionManaged** | Pointer to **bool** | Whether this item is not revision-managed. | [optional] 
**PartNumber** | Pointer to **string** | The part number of the referring element. | [optional] 
**ProjectId** | Pointer to **string** | The project ID associated with the document. | [optional] 
**Properties** | Pointer to [**[]BTMetadataValueInfo**](BTMetadataValueInfo.md) | Custom metadata properties of the item. Only populated when includeProperties is true. | [optional] 
**ResourceType** | Pointer to **string** | The resource type of this item. | [optional] 
**Revision** | Pointer to **string** | The revision of the referring element. | [optional] 
**RevisionObsolete** | Pointer to **bool** | Whether the revision is obsolete. | [optional] 
**StandardProperties** | Pointer to [**[]Property**](Property.md) | Standard metadata properties of the item. | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**ThumbnailHref** | Pointer to **string** | The thumbnail href URI for the referring element. | [optional] 
**VersionId** | Pointer to **string** | The version ID of the document containing the referring element. | [optional] 
**VersionName** | Pointer to **string** | The version name of the document containing the referring element. | [optional] 
**WvCreatedAt** | Pointer to **JSONTime** | The timestamp when the version or workspace was created. | [optional] 

## Methods

### NewBTProductStructureItemInfo

`func NewBTProductStructureItemInfo(jsonType string, ) *BTProductStructureItemInfo`

NewBTProductStructureItemInfo instantiates a new BTProductStructureItemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTProductStructureItemInfoWithDefaults

`func NewBTProductStructureItemInfoWithDefaults() *BTProductStructureItemInfo`

NewBTProductStructureItemInfoWithDefaults instantiates a new BTProductStructureItemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonType

`func (o *BTProductStructureItemInfo) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *BTProductStructureItemInfo) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *BTProductStructureItemInfo) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.


### GetHref

`func (o *BTProductStructureItemInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTProductStructureItemInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTProductStructureItemInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTProductStructureItemInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTProductStructureItemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTProductStructureItemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTProductStructureItemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTProductStructureItemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTProductStructureItemInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTProductStructureItemInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTProductStructureItemInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTProductStructureItemInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViewRef

`func (o *BTProductStructureItemInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTProductStructureItemInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTProductStructureItemInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTProductStructureItemInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTProductStructureItemInfo) GetConfiguration() []ConfigInfo`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTProductStructureItemInfo) GetConfigurationOk() (*[]ConfigInfo, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTProductStructureItemInfo) SetConfiguration(v []ConfigInfo)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTProductStructureItemInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTProductStructureItemInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTProductStructureItemInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTProductStructureItemInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTProductStructureItemInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentName

`func (o *BTProductStructureItemInfo) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *BTProductStructureItemInfo) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *BTProductStructureItemInfo) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *BTProductStructureItemInfo) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetDocumentState

`func (o *BTProductStructureItemInfo) GetDocumentState() int32`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *BTProductStructureItemInfo) GetDocumentStateOk() (*int32, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *BTProductStructureItemInfo) SetDocumentState(v int32)`

SetDocumentState sets DocumentState field to given value.

### HasDocumentState

`func (o *BTProductStructureItemInfo) HasDocumentState() bool`

HasDocumentState returns a boolean if a field has been set.

### GetElementId

`func (o *BTProductStructureItemInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTProductStructureItemInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTProductStructureItemInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTProductStructureItemInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTProductStructureItemInfo) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTProductStructureItemInfo) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTProductStructureItemInfo) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTProductStructureItemInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetFlattenedBody

`func (o *BTProductStructureItemInfo) GetFlattenedBody() bool`

GetFlattenedBody returns the FlattenedBody field if non-nil, zero value otherwise.

### GetFlattenedBodyOk

`func (o *BTProductStructureItemInfo) GetFlattenedBodyOk() (*bool, bool)`

GetFlattenedBodyOk returns a tuple with the FlattenedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenedBody

`func (o *BTProductStructureItemInfo) SetFlattenedBody(v bool)`

SetFlattenedBody sets FlattenedBody field to given value.

### HasFlattenedBody

`func (o *BTProductStructureItemInfo) HasFlattenedBody() bool`

HasFlattenedBody returns a boolean if a field has been set.

### GetFolderId

`func (o *BTProductStructureItemInfo) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BTProductStructureItemInfo) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BTProductStructureItemInfo) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BTProductStructureItemInfo) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetHasDrawing

`func (o *BTProductStructureItemInfo) GetHasDrawing() bool`

GetHasDrawing returns the HasDrawing field if non-nil, zero value otherwise.

### GetHasDrawingOk

`func (o *BTProductStructureItemInfo) GetHasDrawingOk() (*bool, bool)`

GetHasDrawingOk returns a tuple with the HasDrawing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDrawing

`func (o *BTProductStructureItemInfo) SetHasDrawing(v bool)`

SetHasDrawing sets HasDrawing field to given value.

### HasHasDrawing

`func (o *BTProductStructureItemInfo) HasHasDrawing() bool`

HasHasDrawing returns a boolean if a field has been set.

### GetLatestRevision

`func (o *BTProductStructureItemInfo) GetLatestRevision() bool`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *BTProductStructureItemInfo) GetLatestRevisionOk() (*bool, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *BTProductStructureItemInfo) SetLatestRevision(v bool)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *BTProductStructureItemInfo) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTProductStructureItemInfo) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTProductStructureItemInfo) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTProductStructureItemInfo) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTProductStructureItemInfo) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTProductStructureItemInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTProductStructureItemInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTProductStructureItemInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTProductStructureItemInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProjectId

`func (o *BTProductStructureItemInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTProductStructureItemInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTProductStructureItemInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTProductStructureItemInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProperties

`func (o *BTProductStructureItemInfo) GetProperties() []BTMetadataValueInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTProductStructureItemInfo) GetPropertiesOk() (*[]BTMetadataValueInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTProductStructureItemInfo) SetProperties(v []BTMetadataValueInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTProductStructureItemInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResourceType

`func (o *BTProductStructureItemInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTProductStructureItemInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTProductStructureItemInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTProductStructureItemInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRevision

`func (o *BTProductStructureItemInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTProductStructureItemInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTProductStructureItemInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTProductStructureItemInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionObsolete

`func (o *BTProductStructureItemInfo) GetRevisionObsolete() bool`

GetRevisionObsolete returns the RevisionObsolete field if non-nil, zero value otherwise.

### GetRevisionObsoleteOk

`func (o *BTProductStructureItemInfo) GetRevisionObsoleteOk() (*bool, bool)`

GetRevisionObsoleteOk returns a tuple with the RevisionObsolete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionObsolete

`func (o *BTProductStructureItemInfo) SetRevisionObsolete(v bool)`

SetRevisionObsolete sets RevisionObsolete field to given value.

### HasRevisionObsolete

`func (o *BTProductStructureItemInfo) HasRevisionObsolete() bool`

HasRevisionObsolete returns a boolean if a field has been set.

### GetStandardProperties

`func (o *BTProductStructureItemInfo) GetStandardProperties() []Property`

GetStandardProperties returns the StandardProperties field if non-nil, zero value otherwise.

### GetStandardPropertiesOk

`func (o *BTProductStructureItemInfo) GetStandardPropertiesOk() (*[]Property, bool)`

GetStandardPropertiesOk returns a tuple with the StandardProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardProperties

`func (o *BTProductStructureItemInfo) SetStandardProperties(v []Property)`

SetStandardProperties sets StandardProperties field to given value.

### HasStandardProperties

`func (o *BTProductStructureItemInfo) HasStandardProperties() bool`

HasStandardProperties returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTProductStructureItemInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTProductStructureItemInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTProductStructureItemInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTProductStructureItemInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetThumbnailHref

`func (o *BTProductStructureItemInfo) GetThumbnailHref() string`

GetThumbnailHref returns the ThumbnailHref field if non-nil, zero value otherwise.

### GetThumbnailHrefOk

`func (o *BTProductStructureItemInfo) GetThumbnailHrefOk() (*string, bool)`

GetThumbnailHrefOk returns a tuple with the ThumbnailHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHref

`func (o *BTProductStructureItemInfo) SetThumbnailHref(v string)`

SetThumbnailHref sets ThumbnailHref field to given value.

### HasThumbnailHref

`func (o *BTProductStructureItemInfo) HasThumbnailHref() bool`

HasThumbnailHref returns a boolean if a field has been set.

### GetVersionId

`func (o *BTProductStructureItemInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTProductStructureItemInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTProductStructureItemInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTProductStructureItemInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *BTProductStructureItemInfo) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *BTProductStructureItemInfo) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *BTProductStructureItemInfo) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *BTProductStructureItemInfo) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetWvCreatedAt

`func (o *BTProductStructureItemInfo) GetWvCreatedAt() JSONTime`

GetWvCreatedAt returns the WvCreatedAt field if non-nil, zero value otherwise.

### GetWvCreatedAtOk

`func (o *BTProductStructureItemInfo) GetWvCreatedAtOk() (*JSONTime, bool)`

GetWvCreatedAtOk returns a tuple with the WvCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWvCreatedAt

`func (o *BTProductStructureItemInfo) SetWvCreatedAt(v JSONTime)`

SetWvCreatedAt sets WvCreatedAt field to given value.

### HasWvCreatedAt

`func (o *BTProductStructureItemInfo) HasWvCreatedAt() bool`

HasWvCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


