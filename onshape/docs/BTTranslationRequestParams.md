# BTTranslationRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowFaultyParts** | Pointer to **bool** |  | [optional] 
**CreateComposite** | Pointer to **bool** |  | [optional] 
**CreateDrawingIfPossible** | Pointer to **bool** |  | [optional] 
**EncodedFilename** | Pointer to **string** |  | [optional] 
**ExtractAssemblyHierarchy** | Pointer to **bool** |  | [optional] 
**File** | Pointer to [***os.File**](*os.File.md) |  | [optional] 
**FileBodyWithDetails** | Pointer to [**FormDataBodyPart**](FormDataBodyPart.md) |  | [optional] 
**FileContentLength** | Pointer to **int64** |  | [optional] 
**FileDetail** | Pointer to [**FormDataContentDisposition**](FormDataContentDisposition.md) |  | [optional] 
**FlattenAssemblies** | Pointer to **bool** |  | [optional] 
**FormatName** | Pointer to **string** |  | [optional] 
**IsyAxisIsUp** | Pointer to **bool** |  | [optional] 
**JoinAdjacentSurfaces** | Pointer to **bool** |  | [optional] 
**LocationElementId** | Pointer to **string** |  | [optional] 
**LocationGroupId** | Pointer to **string** |  | [optional] 
**LocationPosition** | Pointer to **int32** |  | [optional] 
**NotifyUser** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**SplitAssembliesIntoMultipleDocuments** | Pointer to **bool** |  | [optional] 
**StoreInDocument** | Pointer to **bool** |  | [optional] 
**Translate** | Pointer to **bool** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UploadId** | Pointer to **string** |  | [optional] 
**VersionString** | Pointer to **string** |  | [optional] 

## Methods

### NewBTTranslationRequestParams

`func NewBTTranslationRequestParams() *BTTranslationRequestParams`

NewBTTranslationRequestParams instantiates a new BTTranslationRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTranslationRequestParamsWithDefaults

`func NewBTTranslationRequestParamsWithDefaults() *BTTranslationRequestParams`

NewBTTranslationRequestParamsWithDefaults instantiates a new BTTranslationRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowFaultyParts

`func (o *BTTranslationRequestParams) GetAllowFaultyParts() bool`

GetAllowFaultyParts returns the AllowFaultyParts field if non-nil, zero value otherwise.

### GetAllowFaultyPartsOk

`func (o *BTTranslationRequestParams) GetAllowFaultyPartsOk() (*bool, bool)`

GetAllowFaultyPartsOk returns a tuple with the AllowFaultyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFaultyParts

`func (o *BTTranslationRequestParams) SetAllowFaultyParts(v bool)`

SetAllowFaultyParts sets AllowFaultyParts field to given value.

### HasAllowFaultyParts

`func (o *BTTranslationRequestParams) HasAllowFaultyParts() bool`

HasAllowFaultyParts returns a boolean if a field has been set.

### GetCreateComposite

`func (o *BTTranslationRequestParams) GetCreateComposite() bool`

GetCreateComposite returns the CreateComposite field if non-nil, zero value otherwise.

### GetCreateCompositeOk

`func (o *BTTranslationRequestParams) GetCreateCompositeOk() (*bool, bool)`

GetCreateCompositeOk returns a tuple with the CreateComposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateComposite

`func (o *BTTranslationRequestParams) SetCreateComposite(v bool)`

SetCreateComposite sets CreateComposite field to given value.

### HasCreateComposite

`func (o *BTTranslationRequestParams) HasCreateComposite() bool`

HasCreateComposite returns a boolean if a field has been set.

### GetCreateDrawingIfPossible

`func (o *BTTranslationRequestParams) GetCreateDrawingIfPossible() bool`

GetCreateDrawingIfPossible returns the CreateDrawingIfPossible field if non-nil, zero value otherwise.

### GetCreateDrawingIfPossibleOk

`func (o *BTTranslationRequestParams) GetCreateDrawingIfPossibleOk() (*bool, bool)`

GetCreateDrawingIfPossibleOk returns a tuple with the CreateDrawingIfPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDrawingIfPossible

`func (o *BTTranslationRequestParams) SetCreateDrawingIfPossible(v bool)`

SetCreateDrawingIfPossible sets CreateDrawingIfPossible field to given value.

### HasCreateDrawingIfPossible

`func (o *BTTranslationRequestParams) HasCreateDrawingIfPossible() bool`

HasCreateDrawingIfPossible returns a boolean if a field has been set.

### GetEncodedFilename

`func (o *BTTranslationRequestParams) GetEncodedFilename() string`

GetEncodedFilename returns the EncodedFilename field if non-nil, zero value otherwise.

### GetEncodedFilenameOk

`func (o *BTTranslationRequestParams) GetEncodedFilenameOk() (*string, bool)`

GetEncodedFilenameOk returns a tuple with the EncodedFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedFilename

`func (o *BTTranslationRequestParams) SetEncodedFilename(v string)`

SetEncodedFilename sets EncodedFilename field to given value.

### HasEncodedFilename

`func (o *BTTranslationRequestParams) HasEncodedFilename() bool`

HasEncodedFilename returns a boolean if a field has been set.

### GetExtractAssemblyHierarchy

`func (o *BTTranslationRequestParams) GetExtractAssemblyHierarchy() bool`

GetExtractAssemblyHierarchy returns the ExtractAssemblyHierarchy field if non-nil, zero value otherwise.

### GetExtractAssemblyHierarchyOk

`func (o *BTTranslationRequestParams) GetExtractAssemblyHierarchyOk() (*bool, bool)`

GetExtractAssemblyHierarchyOk returns a tuple with the ExtractAssemblyHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractAssemblyHierarchy

`func (o *BTTranslationRequestParams) SetExtractAssemblyHierarchy(v bool)`

SetExtractAssemblyHierarchy sets ExtractAssemblyHierarchy field to given value.

### HasExtractAssemblyHierarchy

`func (o *BTTranslationRequestParams) HasExtractAssemblyHierarchy() bool`

HasExtractAssemblyHierarchy returns a boolean if a field has been set.

### GetFile

`func (o *BTTranslationRequestParams) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *BTTranslationRequestParams) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *BTTranslationRequestParams) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *BTTranslationRequestParams) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileBodyWithDetails

`func (o *BTTranslationRequestParams) GetFileBodyWithDetails() FormDataBodyPart`

GetFileBodyWithDetails returns the FileBodyWithDetails field if non-nil, zero value otherwise.

### GetFileBodyWithDetailsOk

`func (o *BTTranslationRequestParams) GetFileBodyWithDetailsOk() (*FormDataBodyPart, bool)`

GetFileBodyWithDetailsOk returns a tuple with the FileBodyWithDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBodyWithDetails

`func (o *BTTranslationRequestParams) SetFileBodyWithDetails(v FormDataBodyPart)`

SetFileBodyWithDetails sets FileBodyWithDetails field to given value.

### HasFileBodyWithDetails

`func (o *BTTranslationRequestParams) HasFileBodyWithDetails() bool`

HasFileBodyWithDetails returns a boolean if a field has been set.

### GetFileContentLength

`func (o *BTTranslationRequestParams) GetFileContentLength() int64`

GetFileContentLength returns the FileContentLength field if non-nil, zero value otherwise.

### GetFileContentLengthOk

`func (o *BTTranslationRequestParams) GetFileContentLengthOk() (*int64, bool)`

GetFileContentLengthOk returns a tuple with the FileContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContentLength

`func (o *BTTranslationRequestParams) SetFileContentLength(v int64)`

SetFileContentLength sets FileContentLength field to given value.

### HasFileContentLength

`func (o *BTTranslationRequestParams) HasFileContentLength() bool`

HasFileContentLength returns a boolean if a field has been set.

### GetFileDetail

`func (o *BTTranslationRequestParams) GetFileDetail() FormDataContentDisposition`

GetFileDetail returns the FileDetail field if non-nil, zero value otherwise.

### GetFileDetailOk

`func (o *BTTranslationRequestParams) GetFileDetailOk() (*FormDataContentDisposition, bool)`

GetFileDetailOk returns a tuple with the FileDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDetail

`func (o *BTTranslationRequestParams) SetFileDetail(v FormDataContentDisposition)`

SetFileDetail sets FileDetail field to given value.

### HasFileDetail

`func (o *BTTranslationRequestParams) HasFileDetail() bool`

HasFileDetail returns a boolean if a field has been set.

### GetFlattenAssemblies

`func (o *BTTranslationRequestParams) GetFlattenAssemblies() bool`

GetFlattenAssemblies returns the FlattenAssemblies field if non-nil, zero value otherwise.

### GetFlattenAssembliesOk

`func (o *BTTranslationRequestParams) GetFlattenAssembliesOk() (*bool, bool)`

GetFlattenAssembliesOk returns a tuple with the FlattenAssemblies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenAssemblies

`func (o *BTTranslationRequestParams) SetFlattenAssemblies(v bool)`

SetFlattenAssemblies sets FlattenAssemblies field to given value.

### HasFlattenAssemblies

`func (o *BTTranslationRequestParams) HasFlattenAssemblies() bool`

HasFlattenAssemblies returns a boolean if a field has been set.

### GetFormatName

`func (o *BTTranslationRequestParams) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *BTTranslationRequestParams) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *BTTranslationRequestParams) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.

### HasFormatName

`func (o *BTTranslationRequestParams) HasFormatName() bool`

HasFormatName returns a boolean if a field has been set.

### GetIsyAxisIsUp

`func (o *BTTranslationRequestParams) GetIsyAxisIsUp() bool`

GetIsyAxisIsUp returns the IsyAxisIsUp field if non-nil, zero value otherwise.

### GetIsyAxisIsUpOk

`func (o *BTTranslationRequestParams) GetIsyAxisIsUpOk() (*bool, bool)`

GetIsyAxisIsUpOk returns a tuple with the IsyAxisIsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsyAxisIsUp

`func (o *BTTranslationRequestParams) SetIsyAxisIsUp(v bool)`

SetIsyAxisIsUp sets IsyAxisIsUp field to given value.

### HasIsyAxisIsUp

`func (o *BTTranslationRequestParams) HasIsyAxisIsUp() bool`

HasIsyAxisIsUp returns a boolean if a field has been set.

### GetJoinAdjacentSurfaces

`func (o *BTTranslationRequestParams) GetJoinAdjacentSurfaces() bool`

GetJoinAdjacentSurfaces returns the JoinAdjacentSurfaces field if non-nil, zero value otherwise.

### GetJoinAdjacentSurfacesOk

`func (o *BTTranslationRequestParams) GetJoinAdjacentSurfacesOk() (*bool, bool)`

GetJoinAdjacentSurfacesOk returns a tuple with the JoinAdjacentSurfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAdjacentSurfaces

`func (o *BTTranslationRequestParams) SetJoinAdjacentSurfaces(v bool)`

SetJoinAdjacentSurfaces sets JoinAdjacentSurfaces field to given value.

### HasJoinAdjacentSurfaces

`func (o *BTTranslationRequestParams) HasJoinAdjacentSurfaces() bool`

HasJoinAdjacentSurfaces returns a boolean if a field has been set.

### GetLocationElementId

`func (o *BTTranslationRequestParams) GetLocationElementId() string`

GetLocationElementId returns the LocationElementId field if non-nil, zero value otherwise.

### GetLocationElementIdOk

`func (o *BTTranslationRequestParams) GetLocationElementIdOk() (*string, bool)`

GetLocationElementIdOk returns a tuple with the LocationElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationElementId

`func (o *BTTranslationRequestParams) SetLocationElementId(v string)`

SetLocationElementId sets LocationElementId field to given value.

### HasLocationElementId

`func (o *BTTranslationRequestParams) HasLocationElementId() bool`

HasLocationElementId returns a boolean if a field has been set.

### GetLocationGroupId

`func (o *BTTranslationRequestParams) GetLocationGroupId() string`

GetLocationGroupId returns the LocationGroupId field if non-nil, zero value otherwise.

### GetLocationGroupIdOk

`func (o *BTTranslationRequestParams) GetLocationGroupIdOk() (*string, bool)`

GetLocationGroupIdOk returns a tuple with the LocationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationGroupId

`func (o *BTTranslationRequestParams) SetLocationGroupId(v string)`

SetLocationGroupId sets LocationGroupId field to given value.

### HasLocationGroupId

`func (o *BTTranslationRequestParams) HasLocationGroupId() bool`

HasLocationGroupId returns a boolean if a field has been set.

### GetLocationPosition

`func (o *BTTranslationRequestParams) GetLocationPosition() int32`

GetLocationPosition returns the LocationPosition field if non-nil, zero value otherwise.

### GetLocationPositionOk

`func (o *BTTranslationRequestParams) GetLocationPositionOk() (*int32, bool)`

GetLocationPositionOk returns a tuple with the LocationPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationPosition

`func (o *BTTranslationRequestParams) SetLocationPosition(v int32)`

SetLocationPosition sets LocationPosition field to given value.

### HasLocationPosition

`func (o *BTTranslationRequestParams) HasLocationPosition() bool`

HasLocationPosition returns a boolean if a field has been set.

### GetNotifyUser

`func (o *BTTranslationRequestParams) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *BTTranslationRequestParams) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *BTTranslationRequestParams) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *BTTranslationRequestParams) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTTranslationRequestParams) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTTranslationRequestParams) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTTranslationRequestParams) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTTranslationRequestParams) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTTranslationRequestParams) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTTranslationRequestParams) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTTranslationRequestParams) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTTranslationRequestParams) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetParentId

`func (o *BTTranslationRequestParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTTranslationRequestParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTTranslationRequestParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTTranslationRequestParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTTranslationRequestParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTTranslationRequestParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTTranslationRequestParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTTranslationRequestParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPublic

`func (o *BTTranslationRequestParams) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTTranslationRequestParams) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTTranslationRequestParams) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTTranslationRequestParams) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSplitAssembliesIntoMultipleDocuments

`func (o *BTTranslationRequestParams) GetSplitAssembliesIntoMultipleDocuments() bool`

GetSplitAssembliesIntoMultipleDocuments returns the SplitAssembliesIntoMultipleDocuments field if non-nil, zero value otherwise.

### GetSplitAssembliesIntoMultipleDocumentsOk

`func (o *BTTranslationRequestParams) GetSplitAssembliesIntoMultipleDocumentsOk() (*bool, bool)`

GetSplitAssembliesIntoMultipleDocumentsOk returns a tuple with the SplitAssembliesIntoMultipleDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAssembliesIntoMultipleDocuments

`func (o *BTTranslationRequestParams) SetSplitAssembliesIntoMultipleDocuments(v bool)`

SetSplitAssembliesIntoMultipleDocuments sets SplitAssembliesIntoMultipleDocuments field to given value.

### HasSplitAssembliesIntoMultipleDocuments

`func (o *BTTranslationRequestParams) HasSplitAssembliesIntoMultipleDocuments() bool`

HasSplitAssembliesIntoMultipleDocuments returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTTranslationRequestParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTTranslationRequestParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTTranslationRequestParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTTranslationRequestParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTranslate

`func (o *BTTranslationRequestParams) GetTranslate() bool`

GetTranslate returns the Translate field if non-nil, zero value otherwise.

### GetTranslateOk

`func (o *BTTranslationRequestParams) GetTranslateOk() (*bool, bool)`

GetTranslateOk returns a tuple with the Translate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslate

`func (o *BTTranslationRequestParams) SetTranslate(v bool)`

SetTranslate sets Translate field to given value.

### HasTranslate

`func (o *BTTranslationRequestParams) HasTranslate() bool`

HasTranslate returns a boolean if a field has been set.

### GetUnit

`func (o *BTTranslationRequestParams) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BTTranslationRequestParams) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BTTranslationRequestParams) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BTTranslationRequestParams) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUploadId

`func (o *BTTranslationRequestParams) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BTTranslationRequestParams) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BTTranslationRequestParams) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *BTTranslationRequestParams) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetVersionString

`func (o *BTTranslationRequestParams) GetVersionString() string`

GetVersionString returns the VersionString field if non-nil, zero value otherwise.

### GetVersionStringOk

`func (o *BTTranslationRequestParams) GetVersionStringOk() (*string, bool)`

GetVersionStringOk returns a tuple with the VersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionString

`func (o *BTTranslationRequestParams) SetVersionString(v string)`

SetVersionString sets VersionString field to given value.

### HasVersionString

`func (o *BTTranslationRequestParams) HasVersionString() bool`

HasVersionString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


