# BTPModuleId235

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**DbimportString** | Pointer to **string** |  | [optional] 
**ElementImport** | Pointer to **bool** |  | [optional] 
**ExternalDocumentWithVersion** | Pointer to [**BTDocumentWithVersionId**](BTDocumentWithVersionId.md) |  | [optional] 
**ExternalDocumentWithVersionAndElementId** | Pointer to [**BTDocumentWithVersionAndElementId**](BTDocumentWithVersionAndElementId.md) |  | [optional] 
**ExternalImport** | Pointer to **bool** |  | [optional] 
**ImportedDocumentId** | Pointer to **string** |  | [optional] 
**ImportedElementId** | Pointer to **string** |  | [optional] 
**ImportedVersionId** | Pointer to **string** |  | [optional] 
**Legacy** | Pointer to **bool** |  | [optional] 
**Microversion** | Pointer to **string** |  | [optional] 
**Path** | Pointer to [**BTPLiteralString259**](BTPLiteralString-259.md) |  | [optional] 
**PathPotentiallyValid** | Pointer to **bool** |  | [optional] 
**PathVersion** | Pointer to **string** |  | [optional] 
**PotentiallyValid** | Pointer to **bool** |  | [optional] 
**SpaceAfterPath** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceAfterVersion** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceBeforePath** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**SpaceBeforeVersion** | Pointer to [**BTPSpace10**](BTPSpace-10.md) |  | [optional] 
**StandardLibrary** | Pointer to **bool** |  | [optional] 
**StandardLibraryCommon** | Pointer to **bool** |  | [optional] 
**ValidLegacyVersion** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to [**BTPLiteralString259**](BTPLiteralString-259.md) |  | [optional] 
**VersionAndMicroversion** | Pointer to **string** |  | [optional] 
**VersionPotentiallyValid** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTPModuleId235

`func NewBTPModuleId235() *BTPModuleId235`

NewBTPModuleId235 instantiates a new BTPModuleId235 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPModuleId235WithDefaults

`func NewBTPModuleId235WithDefaults() *BTPModuleId235`

NewBTPModuleId235WithDefaults instantiates a new BTPModuleId235 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPModuleId235) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPModuleId235) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPModuleId235) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPModuleId235) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDbimportString

`func (o *BTPModuleId235) GetDbimportString() string`

GetDbimportString returns the DbimportString field if non-nil, zero value otherwise.

### GetDbimportStringOk

`func (o *BTPModuleId235) GetDbimportStringOk() (*string, bool)`

GetDbimportStringOk returns a tuple with the DbimportString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbimportString

`func (o *BTPModuleId235) SetDbimportString(v string)`

SetDbimportString sets DbimportString field to given value.

### HasDbimportString

`func (o *BTPModuleId235) HasDbimportString() bool`

HasDbimportString returns a boolean if a field has been set.

### GetElementImport

`func (o *BTPModuleId235) GetElementImport() bool`

GetElementImport returns the ElementImport field if non-nil, zero value otherwise.

### GetElementImportOk

`func (o *BTPModuleId235) GetElementImportOk() (*bool, bool)`

GetElementImportOk returns a tuple with the ElementImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementImport

`func (o *BTPModuleId235) SetElementImport(v bool)`

SetElementImport sets ElementImport field to given value.

### HasElementImport

`func (o *BTPModuleId235) HasElementImport() bool`

HasElementImport returns a boolean if a field has been set.

### GetExternalDocumentWithVersion

`func (o *BTPModuleId235) GetExternalDocumentWithVersion() BTDocumentWithVersionId`

GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionOk

`func (o *BTPModuleId235) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool)`

GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersion

`func (o *BTPModuleId235) SetExternalDocumentWithVersion(v BTDocumentWithVersionId)`

SetExternalDocumentWithVersion sets ExternalDocumentWithVersion field to given value.

### HasExternalDocumentWithVersion

`func (o *BTPModuleId235) HasExternalDocumentWithVersion() bool`

HasExternalDocumentWithVersion returns a boolean if a field has been set.

### GetExternalDocumentWithVersionAndElementId

`func (o *BTPModuleId235) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId`

GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field if non-nil, zero value otherwise.

### GetExternalDocumentWithVersionAndElementIdOk

`func (o *BTPModuleId235) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool)`

GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentWithVersionAndElementId

`func (o *BTPModuleId235) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId)`

SetExternalDocumentWithVersionAndElementId sets ExternalDocumentWithVersionAndElementId field to given value.

### HasExternalDocumentWithVersionAndElementId

`func (o *BTPModuleId235) HasExternalDocumentWithVersionAndElementId() bool`

HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.

### GetExternalImport

`func (o *BTPModuleId235) GetExternalImport() bool`

GetExternalImport returns the ExternalImport field if non-nil, zero value otherwise.

### GetExternalImportOk

`func (o *BTPModuleId235) GetExternalImportOk() (*bool, bool)`

GetExternalImportOk returns a tuple with the ExternalImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalImport

`func (o *BTPModuleId235) SetExternalImport(v bool)`

SetExternalImport sets ExternalImport field to given value.

### HasExternalImport

`func (o *BTPModuleId235) HasExternalImport() bool`

HasExternalImport returns a boolean if a field has been set.

### GetImportedDocumentId

`func (o *BTPModuleId235) GetImportedDocumentId() string`

GetImportedDocumentId returns the ImportedDocumentId field if non-nil, zero value otherwise.

### GetImportedDocumentIdOk

`func (o *BTPModuleId235) GetImportedDocumentIdOk() (*string, bool)`

GetImportedDocumentIdOk returns a tuple with the ImportedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedDocumentId

`func (o *BTPModuleId235) SetImportedDocumentId(v string)`

SetImportedDocumentId sets ImportedDocumentId field to given value.

### HasImportedDocumentId

`func (o *BTPModuleId235) HasImportedDocumentId() bool`

HasImportedDocumentId returns a boolean if a field has been set.

### GetImportedElementId

`func (o *BTPModuleId235) GetImportedElementId() string`

GetImportedElementId returns the ImportedElementId field if non-nil, zero value otherwise.

### GetImportedElementIdOk

`func (o *BTPModuleId235) GetImportedElementIdOk() (*string, bool)`

GetImportedElementIdOk returns a tuple with the ImportedElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedElementId

`func (o *BTPModuleId235) SetImportedElementId(v string)`

SetImportedElementId sets ImportedElementId field to given value.

### HasImportedElementId

`func (o *BTPModuleId235) HasImportedElementId() bool`

HasImportedElementId returns a boolean if a field has been set.

### GetImportedVersionId

`func (o *BTPModuleId235) GetImportedVersionId() string`

GetImportedVersionId returns the ImportedVersionId field if non-nil, zero value otherwise.

### GetImportedVersionIdOk

`func (o *BTPModuleId235) GetImportedVersionIdOk() (*string, bool)`

GetImportedVersionIdOk returns a tuple with the ImportedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedVersionId

`func (o *BTPModuleId235) SetImportedVersionId(v string)`

SetImportedVersionId sets ImportedVersionId field to given value.

### HasImportedVersionId

`func (o *BTPModuleId235) HasImportedVersionId() bool`

HasImportedVersionId returns a boolean if a field has been set.

### GetLegacy

`func (o *BTPModuleId235) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *BTPModuleId235) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *BTPModuleId235) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *BTPModuleId235) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTPModuleId235) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTPModuleId235) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTPModuleId235) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTPModuleId235) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetPath

`func (o *BTPModuleId235) GetPath() BTPLiteralString259`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BTPModuleId235) GetPathOk() (*BTPLiteralString259, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BTPModuleId235) SetPath(v BTPLiteralString259)`

SetPath sets Path field to given value.

### HasPath

`func (o *BTPModuleId235) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPathPotentiallyValid

`func (o *BTPModuleId235) GetPathPotentiallyValid() bool`

GetPathPotentiallyValid returns the PathPotentiallyValid field if non-nil, zero value otherwise.

### GetPathPotentiallyValidOk

`func (o *BTPModuleId235) GetPathPotentiallyValidOk() (*bool, bool)`

GetPathPotentiallyValidOk returns a tuple with the PathPotentiallyValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPotentiallyValid

`func (o *BTPModuleId235) SetPathPotentiallyValid(v bool)`

SetPathPotentiallyValid sets PathPotentiallyValid field to given value.

### HasPathPotentiallyValid

`func (o *BTPModuleId235) HasPathPotentiallyValid() bool`

HasPathPotentiallyValid returns a boolean if a field has been set.

### GetPathVersion

`func (o *BTPModuleId235) GetPathVersion() string`

GetPathVersion returns the PathVersion field if non-nil, zero value otherwise.

### GetPathVersionOk

`func (o *BTPModuleId235) GetPathVersionOk() (*string, bool)`

GetPathVersionOk returns a tuple with the PathVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathVersion

`func (o *BTPModuleId235) SetPathVersion(v string)`

SetPathVersion sets PathVersion field to given value.

### HasPathVersion

`func (o *BTPModuleId235) HasPathVersion() bool`

HasPathVersion returns a boolean if a field has been set.

### GetPotentiallyValid

`func (o *BTPModuleId235) GetPotentiallyValid() bool`

GetPotentiallyValid returns the PotentiallyValid field if non-nil, zero value otherwise.

### GetPotentiallyValidOk

`func (o *BTPModuleId235) GetPotentiallyValidOk() (*bool, bool)`

GetPotentiallyValidOk returns a tuple with the PotentiallyValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyValid

`func (o *BTPModuleId235) SetPotentiallyValid(v bool)`

SetPotentiallyValid sets PotentiallyValid field to given value.

### HasPotentiallyValid

`func (o *BTPModuleId235) HasPotentiallyValid() bool`

HasPotentiallyValid returns a boolean if a field has been set.

### GetSpaceAfterPath

`func (o *BTPModuleId235) GetSpaceAfterPath() BTPSpace10`

GetSpaceAfterPath returns the SpaceAfterPath field if non-nil, zero value otherwise.

### GetSpaceAfterPathOk

`func (o *BTPModuleId235) GetSpaceAfterPathOk() (*BTPSpace10, bool)`

GetSpaceAfterPathOk returns a tuple with the SpaceAfterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterPath

`func (o *BTPModuleId235) SetSpaceAfterPath(v BTPSpace10)`

SetSpaceAfterPath sets SpaceAfterPath field to given value.

### HasSpaceAfterPath

`func (o *BTPModuleId235) HasSpaceAfterPath() bool`

HasSpaceAfterPath returns a boolean if a field has been set.

### GetSpaceAfterVersion

`func (o *BTPModuleId235) GetSpaceAfterVersion() BTPSpace10`

GetSpaceAfterVersion returns the SpaceAfterVersion field if non-nil, zero value otherwise.

### GetSpaceAfterVersionOk

`func (o *BTPModuleId235) GetSpaceAfterVersionOk() (*BTPSpace10, bool)`

GetSpaceAfterVersionOk returns a tuple with the SpaceAfterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAfterVersion

`func (o *BTPModuleId235) SetSpaceAfterVersion(v BTPSpace10)`

SetSpaceAfterVersion sets SpaceAfterVersion field to given value.

### HasSpaceAfterVersion

`func (o *BTPModuleId235) HasSpaceAfterVersion() bool`

HasSpaceAfterVersion returns a boolean if a field has been set.

### GetSpaceBeforePath

`func (o *BTPModuleId235) GetSpaceBeforePath() BTPSpace10`

GetSpaceBeforePath returns the SpaceBeforePath field if non-nil, zero value otherwise.

### GetSpaceBeforePathOk

`func (o *BTPModuleId235) GetSpaceBeforePathOk() (*BTPSpace10, bool)`

GetSpaceBeforePathOk returns a tuple with the SpaceBeforePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforePath

`func (o *BTPModuleId235) SetSpaceBeforePath(v BTPSpace10)`

SetSpaceBeforePath sets SpaceBeforePath field to given value.

### HasSpaceBeforePath

`func (o *BTPModuleId235) HasSpaceBeforePath() bool`

HasSpaceBeforePath returns a boolean if a field has been set.

### GetSpaceBeforeVersion

`func (o *BTPModuleId235) GetSpaceBeforeVersion() BTPSpace10`

GetSpaceBeforeVersion returns the SpaceBeforeVersion field if non-nil, zero value otherwise.

### GetSpaceBeforeVersionOk

`func (o *BTPModuleId235) GetSpaceBeforeVersionOk() (*BTPSpace10, bool)`

GetSpaceBeforeVersionOk returns a tuple with the SpaceBeforeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBeforeVersion

`func (o *BTPModuleId235) SetSpaceBeforeVersion(v BTPSpace10)`

SetSpaceBeforeVersion sets SpaceBeforeVersion field to given value.

### HasSpaceBeforeVersion

`func (o *BTPModuleId235) HasSpaceBeforeVersion() bool`

HasSpaceBeforeVersion returns a boolean if a field has been set.

### GetStandardLibrary

`func (o *BTPModuleId235) GetStandardLibrary() bool`

GetStandardLibrary returns the StandardLibrary field if non-nil, zero value otherwise.

### GetStandardLibraryOk

`func (o *BTPModuleId235) GetStandardLibraryOk() (*bool, bool)`

GetStandardLibraryOk returns a tuple with the StandardLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLibrary

`func (o *BTPModuleId235) SetStandardLibrary(v bool)`

SetStandardLibrary sets StandardLibrary field to given value.

### HasStandardLibrary

`func (o *BTPModuleId235) HasStandardLibrary() bool`

HasStandardLibrary returns a boolean if a field has been set.

### GetStandardLibraryCommon

`func (o *BTPModuleId235) GetStandardLibraryCommon() bool`

GetStandardLibraryCommon returns the StandardLibraryCommon field if non-nil, zero value otherwise.

### GetStandardLibraryCommonOk

`func (o *BTPModuleId235) GetStandardLibraryCommonOk() (*bool, bool)`

GetStandardLibraryCommonOk returns a tuple with the StandardLibraryCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLibraryCommon

`func (o *BTPModuleId235) SetStandardLibraryCommon(v bool)`

SetStandardLibraryCommon sets StandardLibraryCommon field to given value.

### HasStandardLibraryCommon

`func (o *BTPModuleId235) HasStandardLibraryCommon() bool`

HasStandardLibraryCommon returns a boolean if a field has been set.

### GetValidLegacyVersion

`func (o *BTPModuleId235) GetValidLegacyVersion() bool`

GetValidLegacyVersion returns the ValidLegacyVersion field if non-nil, zero value otherwise.

### GetValidLegacyVersionOk

`func (o *BTPModuleId235) GetValidLegacyVersionOk() (*bool, bool)`

GetValidLegacyVersionOk returns a tuple with the ValidLegacyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLegacyVersion

`func (o *BTPModuleId235) SetValidLegacyVersion(v bool)`

SetValidLegacyVersion sets ValidLegacyVersion field to given value.

### HasValidLegacyVersion

`func (o *BTPModuleId235) HasValidLegacyVersion() bool`

HasValidLegacyVersion returns a boolean if a field has been set.

### GetVersion

`func (o *BTPModuleId235) GetVersion() BTPLiteralString259`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTPModuleId235) GetVersionOk() (*BTPLiteralString259, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTPModuleId235) SetVersion(v BTPLiteralString259)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTPModuleId235) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionAndMicroversion

`func (o *BTPModuleId235) GetVersionAndMicroversion() string`

GetVersionAndMicroversion returns the VersionAndMicroversion field if non-nil, zero value otherwise.

### GetVersionAndMicroversionOk

`func (o *BTPModuleId235) GetVersionAndMicroversionOk() (*string, bool)`

GetVersionAndMicroversionOk returns a tuple with the VersionAndMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionAndMicroversion

`func (o *BTPModuleId235) SetVersionAndMicroversion(v string)`

SetVersionAndMicroversion sets VersionAndMicroversion field to given value.

### HasVersionAndMicroversion

`func (o *BTPModuleId235) HasVersionAndMicroversion() bool`

HasVersionAndMicroversion returns a boolean if a field has been set.

### GetVersionPotentiallyValid

`func (o *BTPModuleId235) GetVersionPotentiallyValid() bool`

GetVersionPotentiallyValid returns the VersionPotentiallyValid field if non-nil, zero value otherwise.

### GetVersionPotentiallyValidOk

`func (o *BTPModuleId235) GetVersionPotentiallyValidOk() (*bool, bool)`

GetVersionPotentiallyValidOk returns a tuple with the VersionPotentiallyValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPotentiallyValid

`func (o *BTPModuleId235) SetVersionPotentiallyValid(v bool)`

SetVersionPotentiallyValid sets VersionPotentiallyValid field to given value.

### HasVersionPotentiallyValid

`func (o *BTPModuleId235) HasVersionPotentiallyValid() bool`

HasVersionPotentiallyValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


