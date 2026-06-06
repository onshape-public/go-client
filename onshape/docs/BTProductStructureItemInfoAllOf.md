# BTProductStructureItemInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**[]ConfigInfo**](ConfigInfo.md) | The configuration parameters of the referring element. | [optional] 
**DocumentId** | Pointer to **string** | The document ID of the referring element. | [optional] 
**DocumentName** | Pointer to **string** | The name of the document containing the referring element. | [optional] 
**DocumentState** | Pointer to **int32** |  | [optional] 
**ElementId** | Pointer to **string** | The element ID of the referring element. | [optional] 
**ElementType** | Pointer to **int32** | The element type ordinal of the referring element. | [optional] 
**FlattenedBody** | Pointer to **bool** | Whether the item represents a flattened sheet metal body. | [optional] 
**FolderId** | Pointer to **string** | The folder ID containing the document. | [optional] 
**HasDrawing** | Pointer to **bool** | Whether the item has an associated drawing. | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**LatestRevision** | Pointer to **bool** | Whether this item is the latest revision. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
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
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WvCreatedAt** | Pointer to **JSONTime** | The timestamp when the version or workspace was created. | [optional] 

## Methods

### NewBTProductStructureItemInfoAllOf

`func NewBTProductStructureItemInfoAllOf() *BTProductStructureItemInfoAllOf`

NewBTProductStructureItemInfoAllOf instantiates a new BTProductStructureItemInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTProductStructureItemInfoAllOfWithDefaults

`func NewBTProductStructureItemInfoAllOfWithDefaults() *BTProductStructureItemInfoAllOf`

NewBTProductStructureItemInfoAllOfWithDefaults instantiates a new BTProductStructureItemInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTProductStructureItemInfoAllOf) GetConfiguration() []ConfigInfo`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTProductStructureItemInfoAllOf) GetConfigurationOk() (*[]ConfigInfo, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTProductStructureItemInfoAllOf) SetConfiguration(v []ConfigInfo)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTProductStructureItemInfoAllOf) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTProductStructureItemInfoAllOf) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTProductStructureItemInfoAllOf) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTProductStructureItemInfoAllOf) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTProductStructureItemInfoAllOf) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentName

`func (o *BTProductStructureItemInfoAllOf) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *BTProductStructureItemInfoAllOf) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *BTProductStructureItemInfoAllOf) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *BTProductStructureItemInfoAllOf) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetDocumentState

`func (o *BTProductStructureItemInfoAllOf) GetDocumentState() int32`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *BTProductStructureItemInfoAllOf) GetDocumentStateOk() (*int32, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *BTProductStructureItemInfoAllOf) SetDocumentState(v int32)`

SetDocumentState sets DocumentState field to given value.

### HasDocumentState

`func (o *BTProductStructureItemInfoAllOf) HasDocumentState() bool`

HasDocumentState returns a boolean if a field has been set.

### GetElementId

`func (o *BTProductStructureItemInfoAllOf) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTProductStructureItemInfoAllOf) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTProductStructureItemInfoAllOf) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTProductStructureItemInfoAllOf) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTProductStructureItemInfoAllOf) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTProductStructureItemInfoAllOf) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTProductStructureItemInfoAllOf) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTProductStructureItemInfoAllOf) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetFlattenedBody

`func (o *BTProductStructureItemInfoAllOf) GetFlattenedBody() bool`

GetFlattenedBody returns the FlattenedBody field if non-nil, zero value otherwise.

### GetFlattenedBodyOk

`func (o *BTProductStructureItemInfoAllOf) GetFlattenedBodyOk() (*bool, bool)`

GetFlattenedBodyOk returns a tuple with the FlattenedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenedBody

`func (o *BTProductStructureItemInfoAllOf) SetFlattenedBody(v bool)`

SetFlattenedBody sets FlattenedBody field to given value.

### HasFlattenedBody

`func (o *BTProductStructureItemInfoAllOf) HasFlattenedBody() bool`

HasFlattenedBody returns a boolean if a field has been set.

### GetFolderId

`func (o *BTProductStructureItemInfoAllOf) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BTProductStructureItemInfoAllOf) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BTProductStructureItemInfoAllOf) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BTProductStructureItemInfoAllOf) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetHasDrawing

`func (o *BTProductStructureItemInfoAllOf) GetHasDrawing() bool`

GetHasDrawing returns the HasDrawing field if non-nil, zero value otherwise.

### GetHasDrawingOk

`func (o *BTProductStructureItemInfoAllOf) GetHasDrawingOk() (*bool, bool)`

GetHasDrawingOk returns a tuple with the HasDrawing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDrawing

`func (o *BTProductStructureItemInfoAllOf) SetHasDrawing(v bool)`

SetHasDrawing sets HasDrawing field to given value.

### HasHasDrawing

`func (o *BTProductStructureItemInfoAllOf) HasHasDrawing() bool`

HasHasDrawing returns a boolean if a field has been set.

### GetHref

`func (o *BTProductStructureItemInfoAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTProductStructureItemInfoAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTProductStructureItemInfoAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTProductStructureItemInfoAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTProductStructureItemInfoAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTProductStructureItemInfoAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTProductStructureItemInfoAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTProductStructureItemInfoAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestRevision

`func (o *BTProductStructureItemInfoAllOf) GetLatestRevision() bool`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *BTProductStructureItemInfoAllOf) GetLatestRevisionOk() (*bool, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *BTProductStructureItemInfoAllOf) SetLatestRevision(v bool)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *BTProductStructureItemInfoAllOf) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetName

`func (o *BTProductStructureItemInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTProductStructureItemInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTProductStructureItemInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTProductStructureItemInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTProductStructureItemInfoAllOf) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTProductStructureItemInfoAllOf) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTProductStructureItemInfoAllOf) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTProductStructureItemInfoAllOf) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTProductStructureItemInfoAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTProductStructureItemInfoAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTProductStructureItemInfoAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTProductStructureItemInfoAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProjectId

`func (o *BTProductStructureItemInfoAllOf) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTProductStructureItemInfoAllOf) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTProductStructureItemInfoAllOf) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTProductStructureItemInfoAllOf) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProperties

`func (o *BTProductStructureItemInfoAllOf) GetProperties() []BTMetadataValueInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTProductStructureItemInfoAllOf) GetPropertiesOk() (*[]BTMetadataValueInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTProductStructureItemInfoAllOf) SetProperties(v []BTMetadataValueInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTProductStructureItemInfoAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResourceType

`func (o *BTProductStructureItemInfoAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTProductStructureItemInfoAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTProductStructureItemInfoAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTProductStructureItemInfoAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRevision

`func (o *BTProductStructureItemInfoAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTProductStructureItemInfoAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTProductStructureItemInfoAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTProductStructureItemInfoAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionObsolete

`func (o *BTProductStructureItemInfoAllOf) GetRevisionObsolete() bool`

GetRevisionObsolete returns the RevisionObsolete field if non-nil, zero value otherwise.

### GetRevisionObsoleteOk

`func (o *BTProductStructureItemInfoAllOf) GetRevisionObsoleteOk() (*bool, bool)`

GetRevisionObsoleteOk returns a tuple with the RevisionObsolete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionObsolete

`func (o *BTProductStructureItemInfoAllOf) SetRevisionObsolete(v bool)`

SetRevisionObsolete sets RevisionObsolete field to given value.

### HasRevisionObsolete

`func (o *BTProductStructureItemInfoAllOf) HasRevisionObsolete() bool`

HasRevisionObsolete returns a boolean if a field has been set.

### GetStandardProperties

`func (o *BTProductStructureItemInfoAllOf) GetStandardProperties() []Property`

GetStandardProperties returns the StandardProperties field if non-nil, zero value otherwise.

### GetStandardPropertiesOk

`func (o *BTProductStructureItemInfoAllOf) GetStandardPropertiesOk() (*[]Property, bool)`

GetStandardPropertiesOk returns a tuple with the StandardProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardProperties

`func (o *BTProductStructureItemInfoAllOf) SetStandardProperties(v []Property)`

SetStandardProperties sets StandardProperties field to given value.

### HasStandardProperties

`func (o *BTProductStructureItemInfoAllOf) HasStandardProperties() bool`

HasStandardProperties returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTProductStructureItemInfoAllOf) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTProductStructureItemInfoAllOf) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTProductStructureItemInfoAllOf) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTProductStructureItemInfoAllOf) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetThumbnailHref

`func (o *BTProductStructureItemInfoAllOf) GetThumbnailHref() string`

GetThumbnailHref returns the ThumbnailHref field if non-nil, zero value otherwise.

### GetThumbnailHrefOk

`func (o *BTProductStructureItemInfoAllOf) GetThumbnailHrefOk() (*string, bool)`

GetThumbnailHrefOk returns a tuple with the ThumbnailHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHref

`func (o *BTProductStructureItemInfoAllOf) SetThumbnailHref(v string)`

SetThumbnailHref sets ThumbnailHref field to given value.

### HasThumbnailHref

`func (o *BTProductStructureItemInfoAllOf) HasThumbnailHref() bool`

HasThumbnailHref returns a boolean if a field has been set.

### GetVersionId

`func (o *BTProductStructureItemInfoAllOf) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTProductStructureItemInfoAllOf) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTProductStructureItemInfoAllOf) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTProductStructureItemInfoAllOf) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *BTProductStructureItemInfoAllOf) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *BTProductStructureItemInfoAllOf) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *BTProductStructureItemInfoAllOf) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *BTProductStructureItemInfoAllOf) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetViewRef

`func (o *BTProductStructureItemInfoAllOf) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTProductStructureItemInfoAllOf) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTProductStructureItemInfoAllOf) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTProductStructureItemInfoAllOf) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWvCreatedAt

`func (o *BTProductStructureItemInfoAllOf) GetWvCreatedAt() JSONTime`

GetWvCreatedAt returns the WvCreatedAt field if non-nil, zero value otherwise.

### GetWvCreatedAtOk

`func (o *BTProductStructureItemInfoAllOf) GetWvCreatedAtOk() (*JSONTime, bool)`

GetWvCreatedAtOk returns a tuple with the WvCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWvCreatedAt

`func (o *BTProductStructureItemInfoAllOf) SetWvCreatedAt(v JSONTime)`

SetWvCreatedAt sets WvCreatedAt field to given value.

### HasWvCreatedAt

`func (o *BTProductStructureItemInfoAllOf) HasWvCreatedAt() bool`

HasWvCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


