# BTLegacySearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**HighlightedFields** | Pointer to **map[string][]string** |  | [optional] 
**HitId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**SourceMap** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**BTSearchEntityType**](BTSearchEntityType.md) |  | [optional] 

## Methods

### NewBTLegacySearchHit

`func NewBTLegacySearchHit() *BTLegacySearchHit`

NewBTLegacySearchHit instantiates a new BTLegacySearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTLegacySearchHitWithDefaults

`func NewBTLegacySearchHitWithDefaults() *BTLegacySearchHit`

NewBTLegacySearchHitWithDefaults instantiates a new BTLegacySearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *BTLegacySearchHit) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTLegacySearchHit) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTLegacySearchHit) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTLegacySearchHit) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetFolderId

`func (o *BTLegacySearchHit) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BTLegacySearchHit) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BTLegacySearchHit) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BTLegacySearchHit) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetHighlightedFields

`func (o *BTLegacySearchHit) GetHighlightedFields() map[string][]string`

GetHighlightedFields returns the HighlightedFields field if non-nil, zero value otherwise.

### GetHighlightedFieldsOk

`func (o *BTLegacySearchHit) GetHighlightedFieldsOk() (*map[string][]string, bool)`

GetHighlightedFieldsOk returns a tuple with the HighlightedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightedFields

`func (o *BTLegacySearchHit) SetHighlightedFields(v map[string][]string)`

SetHighlightedFields sets HighlightedFields field to given value.

### HasHighlightedFields

`func (o *BTLegacySearchHit) HasHighlightedFields() bool`

HasHighlightedFields returns a boolean if a field has been set.

### GetHitId

`func (o *BTLegacySearchHit) GetHitId() string`

GetHitId returns the HitId field if non-nil, zero value otherwise.

### GetHitIdOk

`func (o *BTLegacySearchHit) GetHitIdOk() (*string, bool)`

GetHitIdOk returns a tuple with the HitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitId

`func (o *BTLegacySearchHit) SetHitId(v string)`

SetHitId sets HitId field to given value.

### HasHitId

`func (o *BTLegacySearchHit) HasHitId() bool`

HasHitId returns a boolean if a field has been set.

### GetName

`func (o *BTLegacySearchHit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTLegacySearchHit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTLegacySearchHit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTLegacySearchHit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *BTLegacySearchHit) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTLegacySearchHit) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTLegacySearchHit) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTLegacySearchHit) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSourceMap

`func (o *BTLegacySearchHit) GetSourceMap() map[string]interface{}`

GetSourceMap returns the SourceMap field if non-nil, zero value otherwise.

### GetSourceMapOk

`func (o *BTLegacySearchHit) GetSourceMapOk() (*map[string]interface{}, bool)`

GetSourceMapOk returns a tuple with the SourceMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMap

`func (o *BTLegacySearchHit) SetSourceMap(v map[string]interface{})`

SetSourceMap sets SourceMap field to given value.

### HasSourceMap

`func (o *BTLegacySearchHit) HasSourceMap() bool`

HasSourceMap returns a boolean if a field has been set.

### GetType

`func (o *BTLegacySearchHit) GetType() BTSearchEntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTLegacySearchHit) GetTypeOk() (*BTSearchEntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTLegacySearchHit) SetType(v BTSearchEntityType)`

SetType sets Type field to given value.

### HasType

`func (o *BTLegacySearchHit) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


