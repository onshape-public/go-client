# BTTaskItemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyFeatures** | Pointer to **[]string** |  | [optional] 
**BodyType** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementFeature** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementOccurrences** | Pointer to **[]string** |  | [optional] 
**ElementQuery** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**ReleaseState** | Pointer to **int32** |  | [optional] 
**RevisionId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ViewData** | Pointer to [**BTViewDataInfo**](BTViewDataInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTTaskItemInfo

`func NewBTTaskItemInfo() *BTTaskItemInfo`

NewBTTaskItemInfo instantiates a new BTTaskItemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTaskItemInfoWithDefaults

`func NewBTTaskItemInfoWithDefaults() *BTTaskItemInfo`

NewBTTaskItemInfoWithDefaults instantiates a new BTTaskItemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyFeatures

`func (o *BTTaskItemInfo) GetAssemblyFeatures() []string`

GetAssemblyFeatures returns the AssemblyFeatures field if non-nil, zero value otherwise.

### GetAssemblyFeaturesOk

`func (o *BTTaskItemInfo) GetAssemblyFeaturesOk() (*[]string, bool)`

GetAssemblyFeaturesOk returns a tuple with the AssemblyFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyFeatures

`func (o *BTTaskItemInfo) SetAssemblyFeatures(v []string)`

SetAssemblyFeatures sets AssemblyFeatures field to given value.

### HasAssemblyFeatures

`func (o *BTTaskItemInfo) HasAssemblyFeatures() bool`

HasAssemblyFeatures returns a boolean if a field has been set.

### GetBodyType

`func (o *BTTaskItemInfo) GetBodyType() string`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *BTTaskItemInfo) GetBodyTypeOk() (*string, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *BTTaskItemInfo) SetBodyType(v string)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *BTTaskItemInfo) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTTaskItemInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTTaskItemInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTTaskItemInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTTaskItemInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDataType

`func (o *BTTaskItemInfo) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BTTaskItemInfo) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BTTaskItemInfo) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *BTTaskItemInfo) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTTaskItemInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTTaskItemInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTTaskItemInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTTaskItemInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementFeature

`func (o *BTTaskItemInfo) GetElementFeature() string`

GetElementFeature returns the ElementFeature field if non-nil, zero value otherwise.

### GetElementFeatureOk

`func (o *BTTaskItemInfo) GetElementFeatureOk() (*string, bool)`

GetElementFeatureOk returns a tuple with the ElementFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementFeature

`func (o *BTTaskItemInfo) SetElementFeature(v string)`

SetElementFeature sets ElementFeature field to given value.

### HasElementFeature

`func (o *BTTaskItemInfo) HasElementFeature() bool`

HasElementFeature returns a boolean if a field has been set.

### GetElementId

`func (o *BTTaskItemInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTTaskItemInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTTaskItemInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTTaskItemInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementOccurrences

`func (o *BTTaskItemInfo) GetElementOccurrences() []string`

GetElementOccurrences returns the ElementOccurrences field if non-nil, zero value otherwise.

### GetElementOccurrencesOk

`func (o *BTTaskItemInfo) GetElementOccurrencesOk() (*[]string, bool)`

GetElementOccurrencesOk returns a tuple with the ElementOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementOccurrences

`func (o *BTTaskItemInfo) SetElementOccurrences(v []string)`

SetElementOccurrences sets ElementOccurrences field to given value.

### HasElementOccurrences

`func (o *BTTaskItemInfo) HasElementOccurrences() bool`

HasElementOccurrences returns a boolean if a field has been set.

### GetElementQuery

`func (o *BTTaskItemInfo) GetElementQuery() string`

GetElementQuery returns the ElementQuery field if non-nil, zero value otherwise.

### GetElementQueryOk

`func (o *BTTaskItemInfo) GetElementQueryOk() (*string, bool)`

GetElementQueryOk returns a tuple with the ElementQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementQuery

`func (o *BTTaskItemInfo) SetElementQuery(v string)`

SetElementQuery sets ElementQuery field to given value.

### HasElementQuery

`func (o *BTTaskItemInfo) HasElementQuery() bool`

HasElementQuery returns a boolean if a field has been set.

### GetElementType

`func (o *BTTaskItemInfo) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTTaskItemInfo) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTTaskItemInfo) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTTaskItemInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetFileName

`func (o *BTTaskItemInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BTTaskItemInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BTTaskItemInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BTTaskItemInfo) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetHref

`func (o *BTTaskItemInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTaskItemInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTaskItemInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTaskItemInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTaskItemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTaskItemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTaskItemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTaskItemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimeType

`func (o *BTTaskItemInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTTaskItemInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTTaskItemInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTTaskItemInfo) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *BTTaskItemInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTaskItemInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTaskItemInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTaskItemInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartId

`func (o *BTTaskItemInfo) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTTaskItemInfo) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTTaskItemInfo) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTTaskItemInfo) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetReleaseState

`func (o *BTTaskItemInfo) GetReleaseState() int32`

GetReleaseState returns the ReleaseState field if non-nil, zero value otherwise.

### GetReleaseStateOk

`func (o *BTTaskItemInfo) GetReleaseStateOk() (*int32, bool)`

GetReleaseStateOk returns a tuple with the ReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseState

`func (o *BTTaskItemInfo) SetReleaseState(v int32)`

SetReleaseState sets ReleaseState field to given value.

### HasReleaseState

`func (o *BTTaskItemInfo) HasReleaseState() bool`

HasReleaseState returns a boolean if a field has been set.

### GetRevisionId

`func (o *BTTaskItemInfo) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *BTTaskItemInfo) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *BTTaskItemInfo) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *BTTaskItemInfo) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTTaskItemInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTTaskItemInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTTaskItemInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTTaskItemInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetViewData

`func (o *BTTaskItemInfo) GetViewData() BTViewDataInfo`

GetViewData returns the ViewData field if non-nil, zero value otherwise.

### GetViewDataOk

`func (o *BTTaskItemInfo) GetViewDataOk() (*BTViewDataInfo, bool)`

GetViewDataOk returns a tuple with the ViewData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewData

`func (o *BTTaskItemInfo) SetViewData(v BTViewDataInfo)`

SetViewData sets ViewData field to given value.

### HasViewData

`func (o *BTTaskItemInfo) HasViewData() bool`

HasViewData returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTaskItemInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTaskItemInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTaskItemInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTaskItemInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTTaskItemInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTTaskItemInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTTaskItemInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTTaskItemInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


