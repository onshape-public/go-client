# BTDocumentSearchHitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementName** | Pointer to **string** |  | [optional] 
**HighlightedFields** | Pointer to **map[string][]string** |  | [optional] 
**Hit** | Pointer to [**BTLegacySearchHit**](BTLegacySearchHit.md) |  | [optional] 
**HitId** | Pointer to **string** |  | [optional] 
**MeshState** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**SourceMap** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**BTSearchEntityType**](BTSearchEntityType.md) |  | [optional] 
**VersionOrWorkspaceName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTDocumentSearchHitInfo

`func NewBTDocumentSearchHitInfo() *BTDocumentSearchHitInfo`

NewBTDocumentSearchHitInfo instantiates a new BTDocumentSearchHitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentSearchHitInfoWithDefaults

`func NewBTDocumentSearchHitInfoWithDefaults() *BTDocumentSearchHitInfo`

NewBTDocumentSearchHitInfoWithDefaults instantiates a new BTDocumentSearchHitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *BTDocumentSearchHitInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTDocumentSearchHitInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTDocumentSearchHitInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTDocumentSearchHitInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementName

`func (o *BTDocumentSearchHitInfo) GetElementName() string`

GetElementName returns the ElementName field if non-nil, zero value otherwise.

### GetElementNameOk

`func (o *BTDocumentSearchHitInfo) GetElementNameOk() (*string, bool)`

GetElementNameOk returns a tuple with the ElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementName

`func (o *BTDocumentSearchHitInfo) SetElementName(v string)`

SetElementName sets ElementName field to given value.

### HasElementName

`func (o *BTDocumentSearchHitInfo) HasElementName() bool`

HasElementName returns a boolean if a field has been set.

### GetHighlightedFields

`func (o *BTDocumentSearchHitInfo) GetHighlightedFields() map[string][]string`

GetHighlightedFields returns the HighlightedFields field if non-nil, zero value otherwise.

### GetHighlightedFieldsOk

`func (o *BTDocumentSearchHitInfo) GetHighlightedFieldsOk() (*map[string][]string, bool)`

GetHighlightedFieldsOk returns a tuple with the HighlightedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightedFields

`func (o *BTDocumentSearchHitInfo) SetHighlightedFields(v map[string][]string)`

SetHighlightedFields sets HighlightedFields field to given value.

### HasHighlightedFields

`func (o *BTDocumentSearchHitInfo) HasHighlightedFields() bool`

HasHighlightedFields returns a boolean if a field has been set.

### GetHit

`func (o *BTDocumentSearchHitInfo) GetHit() BTLegacySearchHit`

GetHit returns the Hit field if non-nil, zero value otherwise.

### GetHitOk

`func (o *BTDocumentSearchHitInfo) GetHitOk() (*BTLegacySearchHit, bool)`

GetHitOk returns a tuple with the Hit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHit

`func (o *BTDocumentSearchHitInfo) SetHit(v BTLegacySearchHit)`

SetHit sets Hit field to given value.

### HasHit

`func (o *BTDocumentSearchHitInfo) HasHit() bool`

HasHit returns a boolean if a field has been set.

### GetHitId

`func (o *BTDocumentSearchHitInfo) GetHitId() string`

GetHitId returns the HitId field if non-nil, zero value otherwise.

### GetHitIdOk

`func (o *BTDocumentSearchHitInfo) GetHitIdOk() (*string, bool)`

GetHitIdOk returns a tuple with the HitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitId

`func (o *BTDocumentSearchHitInfo) SetHitId(v string)`

SetHitId sets HitId field to given value.

### HasHitId

`func (o *BTDocumentSearchHitInfo) HasHitId() bool`

HasHitId returns a boolean if a field has been set.

### GetMeshState

`func (o *BTDocumentSearchHitInfo) GetMeshState() int32`

GetMeshState returns the MeshState field if non-nil, zero value otherwise.

### GetMeshStateOk

`func (o *BTDocumentSearchHitInfo) GetMeshStateOk() (*int32, bool)`

GetMeshStateOk returns a tuple with the MeshState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshState

`func (o *BTDocumentSearchHitInfo) SetMeshState(v int32)`

SetMeshState sets MeshState field to given value.

### HasMeshState

`func (o *BTDocumentSearchHitInfo) HasMeshState() bool`

HasMeshState returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentSearchHitInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentSearchHitInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentSearchHitInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentSearchHitInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *BTDocumentSearchHitInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTDocumentSearchHitInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTDocumentSearchHitInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTDocumentSearchHitInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSourceMap

`func (o *BTDocumentSearchHitInfo) GetSourceMap() map[string]interface{}`

GetSourceMap returns the SourceMap field if non-nil, zero value otherwise.

### GetSourceMapOk

`func (o *BTDocumentSearchHitInfo) GetSourceMapOk() (*map[string]interface{}, bool)`

GetSourceMapOk returns a tuple with the SourceMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMap

`func (o *BTDocumentSearchHitInfo) SetSourceMap(v map[string]interface{})`

SetSourceMap sets SourceMap field to given value.

### HasSourceMap

`func (o *BTDocumentSearchHitInfo) HasSourceMap() bool`

HasSourceMap returns a boolean if a field has been set.

### GetType

`func (o *BTDocumentSearchHitInfo) GetType() BTSearchEntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTDocumentSearchHitInfo) GetTypeOk() (*BTSearchEntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTDocumentSearchHitInfo) SetType(v BTSearchEntityType)`

SetType sets Type field to given value.

### HasType

`func (o *BTDocumentSearchHitInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersionOrWorkspaceName

`func (o *BTDocumentSearchHitInfo) GetVersionOrWorkspaceName() string`

GetVersionOrWorkspaceName returns the VersionOrWorkspaceName field if non-nil, zero value otherwise.

### GetVersionOrWorkspaceNameOk

`func (o *BTDocumentSearchHitInfo) GetVersionOrWorkspaceNameOk() (*string, bool)`

GetVersionOrWorkspaceNameOk returns a tuple with the VersionOrWorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrWorkspaceName

`func (o *BTDocumentSearchHitInfo) SetVersionOrWorkspaceName(v string)`

SetVersionOrWorkspaceName sets VersionOrWorkspaceName field to given value.

### HasVersionOrWorkspaceName

`func (o *BTDocumentSearchHitInfo) HasVersionOrWorkspaceName() bool`

HasVersionOrWorkspaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


