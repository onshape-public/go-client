# BTAppElementReferenceResolveInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** | &#x60;0: OK (healthy) | 1: INFO | 2: WARNING | 3: ERROR (dangling or view generation call failed) | 4: UNKNOWN&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**IdTag** | Pointer to **string** |  | [optional] 
**IdTagIsValid** | Pointer to **bool** |  | [optional] 
**IsConfigurable** | Pointer to **bool** |  | [optional] 
**IsFlattenedPart** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsSketchOnly** | Pointer to **bool** |  | [optional] 
**IsSurface** | Pointer to **bool** |  | [optional] 
**LatestElementMicroversionId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ReferenceType** | Pointer to **int32** |  | [optional] 
**ResolvedDocumentMicroversionId** | Pointer to **string** |  | [optional] 
**ResolvedElementMicroversionId** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**SketchIds** | Pointer to **[]string** |  | [optional] 
**SourceElementId** | Pointer to **string** |  | [optional] 
**TargetConfiguration** | Pointer to **string** |  | [optional] 
**TargetDocumentId** | Pointer to **string** |  | [optional] 
**TargetDocumentMicroversionId** | Pointer to **string** |  | [optional] 
**TargetElementId** | Pointer to **string** |  | [optional] 
**TargetElementMicroversionId** | Pointer to **string** |  | [optional] 
**TargetVersionId** | Pointer to **string** | Reference to a part or assembly in a version; &#x60;NULL&#x60; when reference is in a workspace. | [optional] 
**TrackNewVersions** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTAppElementReferenceResolveInfo

`func NewBTAppElementReferenceResolveInfo() *BTAppElementReferenceResolveInfo`

NewBTAppElementReferenceResolveInfo instantiates a new BTAppElementReferenceResolveInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementReferenceResolveInfoWithDefaults

`func NewBTAppElementReferenceResolveInfoWithDefaults() *BTAppElementReferenceResolveInfo`

NewBTAppElementReferenceResolveInfoWithDefaults instantiates a new BTAppElementReferenceResolveInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *BTAppElementReferenceResolveInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppElementReferenceResolveInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppElementReferenceResolveInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *BTAppElementReferenceResolveInfo) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementReferenceResolveInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementReferenceResolveInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementReferenceResolveInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementReferenceResolveInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementReferenceResolveInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementReferenceResolveInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementReferenceResolveInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementReferenceResolveInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementReferenceResolveInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementReferenceResolveInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementReferenceResolveInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementReferenceResolveInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetIdTag

`func (o *BTAppElementReferenceResolveInfo) GetIdTag() string`

GetIdTag returns the IdTag field if non-nil, zero value otherwise.

### GetIdTagOk

`func (o *BTAppElementReferenceResolveInfo) GetIdTagOk() (*string, bool)`

GetIdTagOk returns a tuple with the IdTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTag

`func (o *BTAppElementReferenceResolveInfo) SetIdTag(v string)`

SetIdTag sets IdTag field to given value.

### HasIdTag

`func (o *BTAppElementReferenceResolveInfo) HasIdTag() bool`

HasIdTag returns a boolean if a field has been set.

### GetIdTagIsValid

`func (o *BTAppElementReferenceResolveInfo) GetIdTagIsValid() bool`

GetIdTagIsValid returns the IdTagIsValid field if non-nil, zero value otherwise.

### GetIdTagIsValidOk

`func (o *BTAppElementReferenceResolveInfo) GetIdTagIsValidOk() (*bool, bool)`

GetIdTagIsValidOk returns a tuple with the IdTagIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTagIsValid

`func (o *BTAppElementReferenceResolveInfo) SetIdTagIsValid(v bool)`

SetIdTagIsValid sets IdTagIsValid field to given value.

### HasIdTagIsValid

`func (o *BTAppElementReferenceResolveInfo) HasIdTagIsValid() bool`

HasIdTagIsValid returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *BTAppElementReferenceResolveInfo) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *BTAppElementReferenceResolveInfo) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *BTAppElementReferenceResolveInfo) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.

### HasIsConfigurable

`func (o *BTAppElementReferenceResolveInfo) HasIsConfigurable() bool`

HasIsConfigurable returns a boolean if a field has been set.

### GetIsFlattenedPart

`func (o *BTAppElementReferenceResolveInfo) GetIsFlattenedPart() bool`

GetIsFlattenedPart returns the IsFlattenedPart field if non-nil, zero value otherwise.

### GetIsFlattenedPartOk

`func (o *BTAppElementReferenceResolveInfo) GetIsFlattenedPartOk() (*bool, bool)`

GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedPart

`func (o *BTAppElementReferenceResolveInfo) SetIsFlattenedPart(v bool)`

SetIsFlattenedPart sets IsFlattenedPart field to given value.

### HasIsFlattenedPart

`func (o *BTAppElementReferenceResolveInfo) HasIsFlattenedPart() bool`

HasIsFlattenedPart returns a boolean if a field has been set.

### GetIsLocked

`func (o *BTAppElementReferenceResolveInfo) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *BTAppElementReferenceResolveInfo) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *BTAppElementReferenceResolveInfo) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *BTAppElementReferenceResolveInfo) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsSketchOnly

`func (o *BTAppElementReferenceResolveInfo) GetIsSketchOnly() bool`

GetIsSketchOnly returns the IsSketchOnly field if non-nil, zero value otherwise.

### GetIsSketchOnlyOk

`func (o *BTAppElementReferenceResolveInfo) GetIsSketchOnlyOk() (*bool, bool)`

GetIsSketchOnlyOk returns a tuple with the IsSketchOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSketchOnly

`func (o *BTAppElementReferenceResolveInfo) SetIsSketchOnly(v bool)`

SetIsSketchOnly sets IsSketchOnly field to given value.

### HasIsSketchOnly

`func (o *BTAppElementReferenceResolveInfo) HasIsSketchOnly() bool`

HasIsSketchOnly returns a boolean if a field has been set.

### GetIsSurface

`func (o *BTAppElementReferenceResolveInfo) GetIsSurface() bool`

GetIsSurface returns the IsSurface field if non-nil, zero value otherwise.

### GetIsSurfaceOk

`func (o *BTAppElementReferenceResolveInfo) GetIsSurfaceOk() (*bool, bool)`

GetIsSurfaceOk returns a tuple with the IsSurface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurface

`func (o *BTAppElementReferenceResolveInfo) SetIsSurface(v bool)`

SetIsSurface sets IsSurface field to given value.

### HasIsSurface

`func (o *BTAppElementReferenceResolveInfo) HasIsSurface() bool`

HasIsSurface returns a boolean if a field has been set.

### GetLatestElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) GetLatestElementMicroversionId() string`

GetLatestElementMicroversionId returns the LatestElementMicroversionId field if non-nil, zero value otherwise.

### GetLatestElementMicroversionIdOk

`func (o *BTAppElementReferenceResolveInfo) GetLatestElementMicroversionIdOk() (*string, bool)`

GetLatestElementMicroversionIdOk returns a tuple with the LatestElementMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) SetLatestElementMicroversionId(v string)`

SetLatestElementMicroversionId sets LatestElementMicroversionId field to given value.

### HasLatestElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) HasLatestElementMicroversionId() bool`

HasLatestElementMicroversionId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTAppElementReferenceResolveInfo) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTAppElementReferenceResolveInfo) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTAppElementReferenceResolveInfo) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTAppElementReferenceResolveInfo) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTAppElementReferenceResolveInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTAppElementReferenceResolveInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTAppElementReferenceResolveInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTAppElementReferenceResolveInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetReferenceId

`func (o *BTAppElementReferenceResolveInfo) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *BTAppElementReferenceResolveInfo) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *BTAppElementReferenceResolveInfo) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *BTAppElementReferenceResolveInfo) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceType

`func (o *BTAppElementReferenceResolveInfo) GetReferenceType() int32`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *BTAppElementReferenceResolveInfo) GetReferenceTypeOk() (*int32, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *BTAppElementReferenceResolveInfo) SetReferenceType(v int32)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *BTAppElementReferenceResolveInfo) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetResolvedDocumentMicroversionId

`func (o *BTAppElementReferenceResolveInfo) GetResolvedDocumentMicroversionId() string`

GetResolvedDocumentMicroversionId returns the ResolvedDocumentMicroversionId field if non-nil, zero value otherwise.

### GetResolvedDocumentMicroversionIdOk

`func (o *BTAppElementReferenceResolveInfo) GetResolvedDocumentMicroversionIdOk() (*string, bool)`

GetResolvedDocumentMicroversionIdOk returns a tuple with the ResolvedDocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDocumentMicroversionId

`func (o *BTAppElementReferenceResolveInfo) SetResolvedDocumentMicroversionId(v string)`

SetResolvedDocumentMicroversionId sets ResolvedDocumentMicroversionId field to given value.

### HasResolvedDocumentMicroversionId

`func (o *BTAppElementReferenceResolveInfo) HasResolvedDocumentMicroversionId() bool`

HasResolvedDocumentMicroversionId returns a boolean if a field has been set.

### GetResolvedElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) GetResolvedElementMicroversionId() string`

GetResolvedElementMicroversionId returns the ResolvedElementMicroversionId field if non-nil, zero value otherwise.

### GetResolvedElementMicroversionIdOk

`func (o *BTAppElementReferenceResolveInfo) GetResolvedElementMicroversionIdOk() (*string, bool)`

GetResolvedElementMicroversionIdOk returns a tuple with the ResolvedElementMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) SetResolvedElementMicroversionId(v string)`

SetResolvedElementMicroversionId sets ResolvedElementMicroversionId field to given value.

### HasResolvedElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) HasResolvedElementMicroversionId() bool`

HasResolvedElementMicroversionId returns a boolean if a field has been set.

### GetRevision

`func (o *BTAppElementReferenceResolveInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTAppElementReferenceResolveInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTAppElementReferenceResolveInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTAppElementReferenceResolveInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSketchIds

`func (o *BTAppElementReferenceResolveInfo) GetSketchIds() []string`

GetSketchIds returns the SketchIds field if non-nil, zero value otherwise.

### GetSketchIdsOk

`func (o *BTAppElementReferenceResolveInfo) GetSketchIdsOk() (*[]string, bool)`

GetSketchIdsOk returns a tuple with the SketchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchIds

`func (o *BTAppElementReferenceResolveInfo) SetSketchIds(v []string)`

SetSketchIds sets SketchIds field to given value.

### HasSketchIds

`func (o *BTAppElementReferenceResolveInfo) HasSketchIds() bool`

HasSketchIds returns a boolean if a field has been set.

### GetSourceElementId

`func (o *BTAppElementReferenceResolveInfo) GetSourceElementId() string`

GetSourceElementId returns the SourceElementId field if non-nil, zero value otherwise.

### GetSourceElementIdOk

`func (o *BTAppElementReferenceResolveInfo) GetSourceElementIdOk() (*string, bool)`

GetSourceElementIdOk returns a tuple with the SourceElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceElementId

`func (o *BTAppElementReferenceResolveInfo) SetSourceElementId(v string)`

SetSourceElementId sets SourceElementId field to given value.

### HasSourceElementId

`func (o *BTAppElementReferenceResolveInfo) HasSourceElementId() bool`

HasSourceElementId returns a boolean if a field has been set.

### GetTargetConfiguration

`func (o *BTAppElementReferenceResolveInfo) GetTargetConfiguration() string`

GetTargetConfiguration returns the TargetConfiguration field if non-nil, zero value otherwise.

### GetTargetConfigurationOk

`func (o *BTAppElementReferenceResolveInfo) GetTargetConfigurationOk() (*string, bool)`

GetTargetConfigurationOk returns a tuple with the TargetConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfiguration

`func (o *BTAppElementReferenceResolveInfo) SetTargetConfiguration(v string)`

SetTargetConfiguration sets TargetConfiguration field to given value.

### HasTargetConfiguration

`func (o *BTAppElementReferenceResolveInfo) HasTargetConfiguration() bool`

HasTargetConfiguration returns a boolean if a field has been set.

### GetTargetDocumentId

`func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentId() string`

GetTargetDocumentId returns the TargetDocumentId field if non-nil, zero value otherwise.

### GetTargetDocumentIdOk

`func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentIdOk() (*string, bool)`

GetTargetDocumentIdOk returns a tuple with the TargetDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocumentId

`func (o *BTAppElementReferenceResolveInfo) SetTargetDocumentId(v string)`

SetTargetDocumentId sets TargetDocumentId field to given value.

### HasTargetDocumentId

`func (o *BTAppElementReferenceResolveInfo) HasTargetDocumentId() bool`

HasTargetDocumentId returns a boolean if a field has been set.

### GetTargetDocumentMicroversionId

`func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentMicroversionId() string`

GetTargetDocumentMicroversionId returns the TargetDocumentMicroversionId field if non-nil, zero value otherwise.

### GetTargetDocumentMicroversionIdOk

`func (o *BTAppElementReferenceResolveInfo) GetTargetDocumentMicroversionIdOk() (*string, bool)`

GetTargetDocumentMicroversionIdOk returns a tuple with the TargetDocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocumentMicroversionId

`func (o *BTAppElementReferenceResolveInfo) SetTargetDocumentMicroversionId(v string)`

SetTargetDocumentMicroversionId sets TargetDocumentMicroversionId field to given value.

### HasTargetDocumentMicroversionId

`func (o *BTAppElementReferenceResolveInfo) HasTargetDocumentMicroversionId() bool`

HasTargetDocumentMicroversionId returns a boolean if a field has been set.

### GetTargetElementId

`func (o *BTAppElementReferenceResolveInfo) GetTargetElementId() string`

GetTargetElementId returns the TargetElementId field if non-nil, zero value otherwise.

### GetTargetElementIdOk

`func (o *BTAppElementReferenceResolveInfo) GetTargetElementIdOk() (*string, bool)`

GetTargetElementIdOk returns a tuple with the TargetElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetElementId

`func (o *BTAppElementReferenceResolveInfo) SetTargetElementId(v string)`

SetTargetElementId sets TargetElementId field to given value.

### HasTargetElementId

`func (o *BTAppElementReferenceResolveInfo) HasTargetElementId() bool`

HasTargetElementId returns a boolean if a field has been set.

### GetTargetElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) GetTargetElementMicroversionId() string`

GetTargetElementMicroversionId returns the TargetElementMicroversionId field if non-nil, zero value otherwise.

### GetTargetElementMicroversionIdOk

`func (o *BTAppElementReferenceResolveInfo) GetTargetElementMicroversionIdOk() (*string, bool)`

GetTargetElementMicroversionIdOk returns a tuple with the TargetElementMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) SetTargetElementMicroversionId(v string)`

SetTargetElementMicroversionId sets TargetElementMicroversionId field to given value.

### HasTargetElementMicroversionId

`func (o *BTAppElementReferenceResolveInfo) HasTargetElementMicroversionId() bool`

HasTargetElementMicroversionId returns a boolean if a field has been set.

### GetTargetVersionId

`func (o *BTAppElementReferenceResolveInfo) GetTargetVersionId() string`

GetTargetVersionId returns the TargetVersionId field if non-nil, zero value otherwise.

### GetTargetVersionIdOk

`func (o *BTAppElementReferenceResolveInfo) GetTargetVersionIdOk() (*string, bool)`

GetTargetVersionIdOk returns a tuple with the TargetVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersionId

`func (o *BTAppElementReferenceResolveInfo) SetTargetVersionId(v string)`

SetTargetVersionId sets TargetVersionId field to given value.

### HasTargetVersionId

`func (o *BTAppElementReferenceResolveInfo) HasTargetVersionId() bool`

HasTargetVersionId returns a boolean if a field has been set.

### GetTrackNewVersions

`func (o *BTAppElementReferenceResolveInfo) GetTrackNewVersions() bool`

GetTrackNewVersions returns the TrackNewVersions field if non-nil, zero value otherwise.

### GetTrackNewVersionsOk

`func (o *BTAppElementReferenceResolveInfo) GetTrackNewVersionsOk() (*bool, bool)`

GetTrackNewVersionsOk returns a tuple with the TrackNewVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNewVersions

`func (o *BTAppElementReferenceResolveInfo) SetTrackNewVersions(v bool)`

SetTrackNewVersions sets TrackNewVersions field to given value.

### HasTrackNewVersions

`func (o *BTAppElementReferenceResolveInfo) HasTrackNewVersions() bool`

HasTrackNewVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


