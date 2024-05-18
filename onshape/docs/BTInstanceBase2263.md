# BTInstanceBase2263

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**AssemblyInstance** | Pointer to **bool** |  | [optional] 
**AssemblyMirror** | Pointer to **bool** |  | [optional] 
**AssemblyPattern** | Pointer to **bool** |  | [optional] 
**AssemblyReplicate** | Pointer to **bool** |  | [optional] 
**ClonedInstance** | Pointer to **bool** |  | [optional] 
**CustomData** | Pointer to [**map[string]BTReferenceCustomData1551**](BTReferenceCustomData1551.md) |  | [optional] 
**InstanceFolder** | Pointer to **bool** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**IsFlattenedPart** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**ParametricInstance** | Pointer to **bool** |  | [optional] 
**ParametricOutputInstance** | Pointer to **bool** |  | [optional] 
**ParametricPartStudioChildInstance** | Pointer to **bool** |  | [optional] 
**ParametricPartStudioInstance** | Pointer to **bool** |  | [optional] 
**PartInstance** | Pointer to **bool** |  | [optional] 
**Releasable** | Pointer to **bool** |  | [optional] 
**RevisionCustomData** | Pointer to [**BTRevisionCustomData2090**](BTRevisionCustomData2090.md) |  | [optional] 
**StandardContent** | Pointer to **bool** |  | [optional] 
**StandardContentParametersId** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 
**SuppressedFieldIndex** | Pointer to **int32** |  | [optional] 
**SuppressionConfigured** | Pointer to **bool** | &#x60;true&#x60; if the suppression is configured in the Part Studio. | [optional] 
**SuppressionState** | Pointer to [**BTMSuppressionState1924**](BTMSuppressionState1924.md) |  | [optional] 
**ValidRevisionReference** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTInstanceBase2263

`func NewBTInstanceBase2263() *BTInstanceBase2263`

NewBTInstanceBase2263 instantiates a new BTInstanceBase2263 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInstanceBase2263WithDefaults

`func NewBTInstanceBase2263WithDefaults() *BTInstanceBase2263`

NewBTInstanceBase2263WithDefaults instantiates a new BTInstanceBase2263 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTInstanceBase2263) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTInstanceBase2263) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTInstanceBase2263) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTInstanceBase2263) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetAssemblyInstance

`func (o *BTInstanceBase2263) GetAssemblyInstance() bool`

GetAssemblyInstance returns the AssemblyInstance field if non-nil, zero value otherwise.

### GetAssemblyInstanceOk

`func (o *BTInstanceBase2263) GetAssemblyInstanceOk() (*bool, bool)`

GetAssemblyInstanceOk returns a tuple with the AssemblyInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyInstance

`func (o *BTInstanceBase2263) SetAssemblyInstance(v bool)`

SetAssemblyInstance sets AssemblyInstance field to given value.

### HasAssemblyInstance

`func (o *BTInstanceBase2263) HasAssemblyInstance() bool`

HasAssemblyInstance returns a boolean if a field has been set.

### GetAssemblyMirror

`func (o *BTInstanceBase2263) GetAssemblyMirror() bool`

GetAssemblyMirror returns the AssemblyMirror field if non-nil, zero value otherwise.

### GetAssemblyMirrorOk

`func (o *BTInstanceBase2263) GetAssemblyMirrorOk() (*bool, bool)`

GetAssemblyMirrorOk returns a tuple with the AssemblyMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyMirror

`func (o *BTInstanceBase2263) SetAssemblyMirror(v bool)`

SetAssemblyMirror sets AssemblyMirror field to given value.

### HasAssemblyMirror

`func (o *BTInstanceBase2263) HasAssemblyMirror() bool`

HasAssemblyMirror returns a boolean if a field has been set.

### GetAssemblyPattern

`func (o *BTInstanceBase2263) GetAssemblyPattern() bool`

GetAssemblyPattern returns the AssemblyPattern field if non-nil, zero value otherwise.

### GetAssemblyPatternOk

`func (o *BTInstanceBase2263) GetAssemblyPatternOk() (*bool, bool)`

GetAssemblyPatternOk returns a tuple with the AssemblyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyPattern

`func (o *BTInstanceBase2263) SetAssemblyPattern(v bool)`

SetAssemblyPattern sets AssemblyPattern field to given value.

### HasAssemblyPattern

`func (o *BTInstanceBase2263) HasAssemblyPattern() bool`

HasAssemblyPattern returns a boolean if a field has been set.

### GetAssemblyReplicate

`func (o *BTInstanceBase2263) GetAssemblyReplicate() bool`

GetAssemblyReplicate returns the AssemblyReplicate field if non-nil, zero value otherwise.

### GetAssemblyReplicateOk

`func (o *BTInstanceBase2263) GetAssemblyReplicateOk() (*bool, bool)`

GetAssemblyReplicateOk returns a tuple with the AssemblyReplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyReplicate

`func (o *BTInstanceBase2263) SetAssemblyReplicate(v bool)`

SetAssemblyReplicate sets AssemblyReplicate field to given value.

### HasAssemblyReplicate

`func (o *BTInstanceBase2263) HasAssemblyReplicate() bool`

HasAssemblyReplicate returns a boolean if a field has been set.

### GetClonedInstance

`func (o *BTInstanceBase2263) GetClonedInstance() bool`

GetClonedInstance returns the ClonedInstance field if non-nil, zero value otherwise.

### GetClonedInstanceOk

`func (o *BTInstanceBase2263) GetClonedInstanceOk() (*bool, bool)`

GetClonedInstanceOk returns a tuple with the ClonedInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedInstance

`func (o *BTInstanceBase2263) SetClonedInstance(v bool)`

SetClonedInstance sets ClonedInstance field to given value.

### HasClonedInstance

`func (o *BTInstanceBase2263) HasClonedInstance() bool`

HasClonedInstance returns a boolean if a field has been set.

### GetCustomData

`func (o *BTInstanceBase2263) GetCustomData() map[string]BTReferenceCustomData1551`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *BTInstanceBase2263) GetCustomDataOk() (*map[string]BTReferenceCustomData1551, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *BTInstanceBase2263) SetCustomData(v map[string]BTReferenceCustomData1551)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *BTInstanceBase2263) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetInstanceFolder

`func (o *BTInstanceBase2263) GetInstanceFolder() bool`

GetInstanceFolder returns the InstanceFolder field if non-nil, zero value otherwise.

### GetInstanceFolderOk

`func (o *BTInstanceBase2263) GetInstanceFolderOk() (*bool, bool)`

GetInstanceFolderOk returns a tuple with the InstanceFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFolder

`func (o *BTInstanceBase2263) SetInstanceFolder(v bool)`

SetInstanceFolder sets InstanceFolder field to given value.

### HasInstanceFolder

`func (o *BTInstanceBase2263) HasInstanceFolder() bool`

HasInstanceFolder returns a boolean if a field has been set.

### GetInstanceName

`func (o *BTInstanceBase2263) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BTInstanceBase2263) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BTInstanceBase2263) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BTInstanceBase2263) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetIsFlattenedPart

`func (o *BTInstanceBase2263) GetIsFlattenedPart() bool`

GetIsFlattenedPart returns the IsFlattenedPart field if non-nil, zero value otherwise.

### GetIsFlattenedPartOk

`func (o *BTInstanceBase2263) GetIsFlattenedPartOk() (*bool, bool)`

GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedPart

`func (o *BTInstanceBase2263) SetIsFlattenedPart(v bool)`

SetIsFlattenedPart sets IsFlattenedPart field to given value.

### HasIsFlattenedPart

`func (o *BTInstanceBase2263) HasIsFlattenedPart() bool`

HasIsFlattenedPart returns a boolean if a field has been set.

### GetLocked

`func (o *BTInstanceBase2263) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BTInstanceBase2263) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BTInstanceBase2263) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *BTInstanceBase2263) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetParametricInstance

`func (o *BTInstanceBase2263) GetParametricInstance() bool`

GetParametricInstance returns the ParametricInstance field if non-nil, zero value otherwise.

### GetParametricInstanceOk

`func (o *BTInstanceBase2263) GetParametricInstanceOk() (*bool, bool)`

GetParametricInstanceOk returns a tuple with the ParametricInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstance

`func (o *BTInstanceBase2263) SetParametricInstance(v bool)`

SetParametricInstance sets ParametricInstance field to given value.

### HasParametricInstance

`func (o *BTInstanceBase2263) HasParametricInstance() bool`

HasParametricInstance returns a boolean if a field has been set.

### GetParametricOutputInstance

`func (o *BTInstanceBase2263) GetParametricOutputInstance() bool`

GetParametricOutputInstance returns the ParametricOutputInstance field if non-nil, zero value otherwise.

### GetParametricOutputInstanceOk

`func (o *BTInstanceBase2263) GetParametricOutputInstanceOk() (*bool, bool)`

GetParametricOutputInstanceOk returns a tuple with the ParametricOutputInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricOutputInstance

`func (o *BTInstanceBase2263) SetParametricOutputInstance(v bool)`

SetParametricOutputInstance sets ParametricOutputInstance field to given value.

### HasParametricOutputInstance

`func (o *BTInstanceBase2263) HasParametricOutputInstance() bool`

HasParametricOutputInstance returns a boolean if a field has been set.

### GetParametricPartStudioChildInstance

`func (o *BTInstanceBase2263) GetParametricPartStudioChildInstance() bool`

GetParametricPartStudioChildInstance returns the ParametricPartStudioChildInstance field if non-nil, zero value otherwise.

### GetParametricPartStudioChildInstanceOk

`func (o *BTInstanceBase2263) GetParametricPartStudioChildInstanceOk() (*bool, bool)`

GetParametricPartStudioChildInstanceOk returns a tuple with the ParametricPartStudioChildInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricPartStudioChildInstance

`func (o *BTInstanceBase2263) SetParametricPartStudioChildInstance(v bool)`

SetParametricPartStudioChildInstance sets ParametricPartStudioChildInstance field to given value.

### HasParametricPartStudioChildInstance

`func (o *BTInstanceBase2263) HasParametricPartStudioChildInstance() bool`

HasParametricPartStudioChildInstance returns a boolean if a field has been set.

### GetParametricPartStudioInstance

`func (o *BTInstanceBase2263) GetParametricPartStudioInstance() bool`

GetParametricPartStudioInstance returns the ParametricPartStudioInstance field if non-nil, zero value otherwise.

### GetParametricPartStudioInstanceOk

`func (o *BTInstanceBase2263) GetParametricPartStudioInstanceOk() (*bool, bool)`

GetParametricPartStudioInstanceOk returns a tuple with the ParametricPartStudioInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricPartStudioInstance

`func (o *BTInstanceBase2263) SetParametricPartStudioInstance(v bool)`

SetParametricPartStudioInstance sets ParametricPartStudioInstance field to given value.

### HasParametricPartStudioInstance

`func (o *BTInstanceBase2263) HasParametricPartStudioInstance() bool`

HasParametricPartStudioInstance returns a boolean if a field has been set.

### GetPartInstance

`func (o *BTInstanceBase2263) GetPartInstance() bool`

GetPartInstance returns the PartInstance field if non-nil, zero value otherwise.

### GetPartInstanceOk

`func (o *BTInstanceBase2263) GetPartInstanceOk() (*bool, bool)`

GetPartInstanceOk returns a tuple with the PartInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartInstance

`func (o *BTInstanceBase2263) SetPartInstance(v bool)`

SetPartInstance sets PartInstance field to given value.

### HasPartInstance

`func (o *BTInstanceBase2263) HasPartInstance() bool`

HasPartInstance returns a boolean if a field has been set.

### GetReleasable

`func (o *BTInstanceBase2263) GetReleasable() bool`

GetReleasable returns the Releasable field if non-nil, zero value otherwise.

### GetReleasableOk

`func (o *BTInstanceBase2263) GetReleasableOk() (*bool, bool)`

GetReleasableOk returns a tuple with the Releasable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasable

`func (o *BTInstanceBase2263) SetReleasable(v bool)`

SetReleasable sets Releasable field to given value.

### HasReleasable

`func (o *BTInstanceBase2263) HasReleasable() bool`

HasReleasable returns a boolean if a field has been set.

### GetRevisionCustomData

`func (o *BTInstanceBase2263) GetRevisionCustomData() BTRevisionCustomData2090`

GetRevisionCustomData returns the RevisionCustomData field if non-nil, zero value otherwise.

### GetRevisionCustomDataOk

`func (o *BTInstanceBase2263) GetRevisionCustomDataOk() (*BTRevisionCustomData2090, bool)`

GetRevisionCustomDataOk returns a tuple with the RevisionCustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionCustomData

`func (o *BTInstanceBase2263) SetRevisionCustomData(v BTRevisionCustomData2090)`

SetRevisionCustomData sets RevisionCustomData field to given value.

### HasRevisionCustomData

`func (o *BTInstanceBase2263) HasRevisionCustomData() bool`

HasRevisionCustomData returns a boolean if a field has been set.

### GetStandardContent

`func (o *BTInstanceBase2263) GetStandardContent() bool`

GetStandardContent returns the StandardContent field if non-nil, zero value otherwise.

### GetStandardContentOk

`func (o *BTInstanceBase2263) GetStandardContentOk() (*bool, bool)`

GetStandardContentOk returns a tuple with the StandardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContent

`func (o *BTInstanceBase2263) SetStandardContent(v bool)`

SetStandardContent sets StandardContent field to given value.

### HasStandardContent

`func (o *BTInstanceBase2263) HasStandardContent() bool`

HasStandardContent returns a boolean if a field has been set.

### GetStandardContentParametersId

`func (o *BTInstanceBase2263) GetStandardContentParametersId() string`

GetStandardContentParametersId returns the StandardContentParametersId field if non-nil, zero value otherwise.

### GetStandardContentParametersIdOk

`func (o *BTInstanceBase2263) GetStandardContentParametersIdOk() (*string, bool)`

GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentParametersId

`func (o *BTInstanceBase2263) SetStandardContentParametersId(v string)`

SetStandardContentParametersId sets StandardContentParametersId field to given value.

### HasStandardContentParametersId

`func (o *BTInstanceBase2263) HasStandardContentParametersId() bool`

HasStandardContentParametersId returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTInstanceBase2263) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTInstanceBase2263) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTInstanceBase2263) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTInstanceBase2263) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressedFieldIndex

`func (o *BTInstanceBase2263) GetSuppressedFieldIndex() int32`

GetSuppressedFieldIndex returns the SuppressedFieldIndex field if non-nil, zero value otherwise.

### GetSuppressedFieldIndexOk

`func (o *BTInstanceBase2263) GetSuppressedFieldIndexOk() (*int32, bool)`

GetSuppressedFieldIndexOk returns a tuple with the SuppressedFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedFieldIndex

`func (o *BTInstanceBase2263) SetSuppressedFieldIndex(v int32)`

SetSuppressedFieldIndex sets SuppressedFieldIndex field to given value.

### HasSuppressedFieldIndex

`func (o *BTInstanceBase2263) HasSuppressedFieldIndex() bool`

HasSuppressedFieldIndex returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTInstanceBase2263) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTInstanceBase2263) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTInstanceBase2263) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTInstanceBase2263) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTInstanceBase2263) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTInstanceBase2263) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTInstanceBase2263) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTInstanceBase2263) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.

### GetValidRevisionReference

`func (o *BTInstanceBase2263) GetValidRevisionReference() bool`

GetValidRevisionReference returns the ValidRevisionReference field if non-nil, zero value otherwise.

### GetValidRevisionReferenceOk

`func (o *BTInstanceBase2263) GetValidRevisionReferenceOk() (*bool, bool)`

GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRevisionReference

`func (o *BTInstanceBase2263) SetValidRevisionReference(v bool)`

SetValidRevisionReference sets ValidRevisionReference field to given value.

### HasValidRevisionReference

`func (o *BTInstanceBase2263) HasValidRevisionReference() bool`

HasValidRevisionReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTInstanceBase2263) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTInstanceBase2263) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTInstanceBase2263) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTInstanceBase2263) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


