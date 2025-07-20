# BTTaskItemParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyType** | Pointer to **string** | Body type to reference from a task. | [optional] 
**Configuration** | Pointer to **string** | Configuration of reference. Used to restrict a reference to a specific configuration of an element. | [optional] 
**Description** | Pointer to **string** | Description of the reference. | [optional] 
**DocumentId** | Pointer to **string** | Id of a document. Required to reference a document or anything contained within it. | [optional] 
**ElementId** | Pointer to **string** | Id of an element reference. Used when referencing an element. | [optional] 
**ElementType** | Pointer to **int32** | Type of element reference. Options are 0 (Part Studio), 1 (Assembly), 2 (Drawing), 3 (Feature Studio), 4 (Blob), 5 (Application), 6 (Table), 7 (Bill of Materials),  8 (Variable Studio), or 9 (Publication Item). | [optional] 
**MimeType** | Pointer to **string** | Mimetype of reference. Used when referencing blob elements. | [optional] 
**Name** | Pointer to **string** | Name of the reference. | [optional] 
**PartId** | Pointer to **string** | Deterministic Id of a part. Used when referencing parts. | [optional] 
**RevisionId** | Pointer to **string** | Id of a revision to reference from a task. | [optional] 
**VersionId** | Pointer to **string** | Id of a document version. Used when referencing the version itself or an element or part in it. | [optional] 
**WorkspaceId** | Pointer to **string** | Id of a document workspace. Used when referencing the workspace itself or an element or part in it. | [optional] 

## Methods

### NewBTTaskItemParams

`func NewBTTaskItemParams() *BTTaskItemParams`

NewBTTaskItemParams instantiates a new BTTaskItemParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTaskItemParamsWithDefaults

`func NewBTTaskItemParamsWithDefaults() *BTTaskItemParams`

NewBTTaskItemParamsWithDefaults instantiates a new BTTaskItemParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyType

`func (o *BTTaskItemParams) GetBodyType() string`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *BTTaskItemParams) GetBodyTypeOk() (*string, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *BTTaskItemParams) SetBodyType(v string)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *BTTaskItemParams) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTTaskItemParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTTaskItemParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTTaskItemParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTTaskItemParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *BTTaskItemParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTTaskItemParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTTaskItemParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTTaskItemParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTTaskItemParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTTaskItemParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTTaskItemParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTTaskItemParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTTaskItemParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTTaskItemParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTTaskItemParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTTaskItemParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTTaskItemParams) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTTaskItemParams) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTTaskItemParams) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTTaskItemParams) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetMimeType

`func (o *BTTaskItemParams) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTTaskItemParams) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTTaskItemParams) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTTaskItemParams) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *BTTaskItemParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTaskItemParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTaskItemParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTaskItemParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartId

`func (o *BTTaskItemParams) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTTaskItemParams) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTTaskItemParams) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTTaskItemParams) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetRevisionId

`func (o *BTTaskItemParams) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *BTTaskItemParams) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *BTTaskItemParams) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *BTTaskItemParams) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTTaskItemParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTTaskItemParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTTaskItemParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTTaskItemParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTTaskItemParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTTaskItemParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTTaskItemParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTTaskItemParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


