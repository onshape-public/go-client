# BTInstanceBase2263AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyInstance** | Pointer to **bool** |  | [optional] 
**AssemblyMirror** | Pointer to **bool** |  | [optional] 
**AssemblyPattern** | Pointer to **bool** |  | [optional] 
**AssemblyReplicate** | Pointer to **bool** |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**ClonedInstance** | Pointer to **bool** |  | [optional] 
**CustomData** | Pointer to [**map[string]BTReferenceCustomData1551**](BTReferenceCustomData1551.md) |  | [optional] 
**DerivedAssemblyMirror** | Pointer to **bool** |  | [optional] 
**InstanceFolder** | Pointer to **bool** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**IsFlattenedPart** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**ParametricInstance** | Pointer to **bool** |  | [optional] 
**ParametricOutputInstance** | Pointer to **bool** |  | [optional] 
**ParametricPartStudioChildInstance** | Pointer to **bool** |  | [optional] 
**ParametricPartStudioInstance** | Pointer to **bool** |  | [optional] 
**ParentSuppressed** | Pointer to **bool** |  | [optional] 
**PartInstance** | Pointer to **bool** |  | [optional] 
**Releasable** | Pointer to **bool** |  | [optional] 
**RevisionCustomData** | Pointer to [**BTRevisionCustomData2090**](BTRevisionCustomData2090.md) |  | [optional] 
**StandardContent** | Pointer to **bool** |  | [optional] 
**StandardContentParametersId** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 
**SuppressedFieldIndex** | Pointer to **int32** |  | [optional] 
**SuppressionConfigured** | Pointer to **bool** | &#x60;true&#x60; if the suppression is configured in the Part Studio. | [optional] 
**SuppressionState** | Pointer to [**BTMSuppressionState1924**](BTMSuppressionState1924.md) |  | [optional] 
**SuppressionStateFieldIndex** | Pointer to **int32** |  | [optional] 
**ValidRevisionReference** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTInstanceBase2263AllOf

`func NewBTInstanceBase2263AllOf() *BTInstanceBase2263AllOf`

NewBTInstanceBase2263AllOf instantiates a new BTInstanceBase2263AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInstanceBase2263AllOfWithDefaults

`func NewBTInstanceBase2263AllOfWithDefaults() *BTInstanceBase2263AllOf`

NewBTInstanceBase2263AllOfWithDefaults instantiates a new BTInstanceBase2263AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyInstance

`func (o *BTInstanceBase2263AllOf) GetAssemblyInstance() bool`

GetAssemblyInstance returns the AssemblyInstance field if non-nil, zero value otherwise.

### GetAssemblyInstanceOk

`func (o *BTInstanceBase2263AllOf) GetAssemblyInstanceOk() (*bool, bool)`

GetAssemblyInstanceOk returns a tuple with the AssemblyInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyInstance

`func (o *BTInstanceBase2263AllOf) SetAssemblyInstance(v bool)`

SetAssemblyInstance sets AssemblyInstance field to given value.

### HasAssemblyInstance

`func (o *BTInstanceBase2263AllOf) HasAssemblyInstance() bool`

HasAssemblyInstance returns a boolean if a field has been set.

### GetAssemblyMirror

`func (o *BTInstanceBase2263AllOf) GetAssemblyMirror() bool`

GetAssemblyMirror returns the AssemblyMirror field if non-nil, zero value otherwise.

### GetAssemblyMirrorOk

`func (o *BTInstanceBase2263AllOf) GetAssemblyMirrorOk() (*bool, bool)`

GetAssemblyMirrorOk returns a tuple with the AssemblyMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyMirror

`func (o *BTInstanceBase2263AllOf) SetAssemblyMirror(v bool)`

SetAssemblyMirror sets AssemblyMirror field to given value.

### HasAssemblyMirror

`func (o *BTInstanceBase2263AllOf) HasAssemblyMirror() bool`

HasAssemblyMirror returns a boolean if a field has been set.

### GetAssemblyPattern

`func (o *BTInstanceBase2263AllOf) GetAssemblyPattern() bool`

GetAssemblyPattern returns the AssemblyPattern field if non-nil, zero value otherwise.

### GetAssemblyPatternOk

`func (o *BTInstanceBase2263AllOf) GetAssemblyPatternOk() (*bool, bool)`

GetAssemblyPatternOk returns a tuple with the AssemblyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyPattern

`func (o *BTInstanceBase2263AllOf) SetAssemblyPattern(v bool)`

SetAssemblyPattern sets AssemblyPattern field to given value.

### HasAssemblyPattern

`func (o *BTInstanceBase2263AllOf) HasAssemblyPattern() bool`

HasAssemblyPattern returns a boolean if a field has been set.

### GetAssemblyReplicate

`func (o *BTInstanceBase2263AllOf) GetAssemblyReplicate() bool`

GetAssemblyReplicate returns the AssemblyReplicate field if non-nil, zero value otherwise.

### GetAssemblyReplicateOk

`func (o *BTInstanceBase2263AllOf) GetAssemblyReplicateOk() (*bool, bool)`

GetAssemblyReplicateOk returns a tuple with the AssemblyReplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyReplicate

`func (o *BTInstanceBase2263AllOf) SetAssemblyReplicate(v bool)`

SetAssemblyReplicate sets AssemblyReplicate field to given value.

### HasAssemblyReplicate

`func (o *BTInstanceBase2263AllOf) HasAssemblyReplicate() bool`

HasAssemblyReplicate returns a boolean if a field has been set.

### GetBtType

`func (o *BTInstanceBase2263AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTInstanceBase2263AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTInstanceBase2263AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTInstanceBase2263AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetClonedInstance

`func (o *BTInstanceBase2263AllOf) GetClonedInstance() bool`

GetClonedInstance returns the ClonedInstance field if non-nil, zero value otherwise.

### GetClonedInstanceOk

`func (o *BTInstanceBase2263AllOf) GetClonedInstanceOk() (*bool, bool)`

GetClonedInstanceOk returns a tuple with the ClonedInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedInstance

`func (o *BTInstanceBase2263AllOf) SetClonedInstance(v bool)`

SetClonedInstance sets ClonedInstance field to given value.

### HasClonedInstance

`func (o *BTInstanceBase2263AllOf) HasClonedInstance() bool`

HasClonedInstance returns a boolean if a field has been set.

### GetCustomData

`func (o *BTInstanceBase2263AllOf) GetCustomData() map[string]BTReferenceCustomData1551`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *BTInstanceBase2263AllOf) GetCustomDataOk() (*map[string]BTReferenceCustomData1551, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *BTInstanceBase2263AllOf) SetCustomData(v map[string]BTReferenceCustomData1551)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *BTInstanceBase2263AllOf) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetDerivedAssemblyMirror

`func (o *BTInstanceBase2263AllOf) GetDerivedAssemblyMirror() bool`

GetDerivedAssemblyMirror returns the DerivedAssemblyMirror field if non-nil, zero value otherwise.

### GetDerivedAssemblyMirrorOk

`func (o *BTInstanceBase2263AllOf) GetDerivedAssemblyMirrorOk() (*bool, bool)`

GetDerivedAssemblyMirrorOk returns a tuple with the DerivedAssemblyMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedAssemblyMirror

`func (o *BTInstanceBase2263AllOf) SetDerivedAssemblyMirror(v bool)`

SetDerivedAssemblyMirror sets DerivedAssemblyMirror field to given value.

### HasDerivedAssemblyMirror

`func (o *BTInstanceBase2263AllOf) HasDerivedAssemblyMirror() bool`

HasDerivedAssemblyMirror returns a boolean if a field has been set.

### GetInstanceFolder

`func (o *BTInstanceBase2263AllOf) GetInstanceFolder() bool`

GetInstanceFolder returns the InstanceFolder field if non-nil, zero value otherwise.

### GetInstanceFolderOk

`func (o *BTInstanceBase2263AllOf) GetInstanceFolderOk() (*bool, bool)`

GetInstanceFolderOk returns a tuple with the InstanceFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFolder

`func (o *BTInstanceBase2263AllOf) SetInstanceFolder(v bool)`

SetInstanceFolder sets InstanceFolder field to given value.

### HasInstanceFolder

`func (o *BTInstanceBase2263AllOf) HasInstanceFolder() bool`

HasInstanceFolder returns a boolean if a field has been set.

### GetInstanceName

`func (o *BTInstanceBase2263AllOf) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BTInstanceBase2263AllOf) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BTInstanceBase2263AllOf) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BTInstanceBase2263AllOf) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetIsFlattenedPart

`func (o *BTInstanceBase2263AllOf) GetIsFlattenedPart() bool`

GetIsFlattenedPart returns the IsFlattenedPart field if non-nil, zero value otherwise.

### GetIsFlattenedPartOk

`func (o *BTInstanceBase2263AllOf) GetIsFlattenedPartOk() (*bool, bool)`

GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedPart

`func (o *BTInstanceBase2263AllOf) SetIsFlattenedPart(v bool)`

SetIsFlattenedPart sets IsFlattenedPart field to given value.

### HasIsFlattenedPart

`func (o *BTInstanceBase2263AllOf) HasIsFlattenedPart() bool`

HasIsFlattenedPart returns a boolean if a field has been set.

### GetLocked

`func (o *BTInstanceBase2263AllOf) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BTInstanceBase2263AllOf) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BTInstanceBase2263AllOf) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *BTInstanceBase2263AllOf) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetParametricInstance

`func (o *BTInstanceBase2263AllOf) GetParametricInstance() bool`

GetParametricInstance returns the ParametricInstance field if non-nil, zero value otherwise.

### GetParametricInstanceOk

`func (o *BTInstanceBase2263AllOf) GetParametricInstanceOk() (*bool, bool)`

GetParametricInstanceOk returns a tuple with the ParametricInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricInstance

`func (o *BTInstanceBase2263AllOf) SetParametricInstance(v bool)`

SetParametricInstance sets ParametricInstance field to given value.

### HasParametricInstance

`func (o *BTInstanceBase2263AllOf) HasParametricInstance() bool`

HasParametricInstance returns a boolean if a field has been set.

### GetParametricOutputInstance

`func (o *BTInstanceBase2263AllOf) GetParametricOutputInstance() bool`

GetParametricOutputInstance returns the ParametricOutputInstance field if non-nil, zero value otherwise.

### GetParametricOutputInstanceOk

`func (o *BTInstanceBase2263AllOf) GetParametricOutputInstanceOk() (*bool, bool)`

GetParametricOutputInstanceOk returns a tuple with the ParametricOutputInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricOutputInstance

`func (o *BTInstanceBase2263AllOf) SetParametricOutputInstance(v bool)`

SetParametricOutputInstance sets ParametricOutputInstance field to given value.

### HasParametricOutputInstance

`func (o *BTInstanceBase2263AllOf) HasParametricOutputInstance() bool`

HasParametricOutputInstance returns a boolean if a field has been set.

### GetParametricPartStudioChildInstance

`func (o *BTInstanceBase2263AllOf) GetParametricPartStudioChildInstance() bool`

GetParametricPartStudioChildInstance returns the ParametricPartStudioChildInstance field if non-nil, zero value otherwise.

### GetParametricPartStudioChildInstanceOk

`func (o *BTInstanceBase2263AllOf) GetParametricPartStudioChildInstanceOk() (*bool, bool)`

GetParametricPartStudioChildInstanceOk returns a tuple with the ParametricPartStudioChildInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricPartStudioChildInstance

`func (o *BTInstanceBase2263AllOf) SetParametricPartStudioChildInstance(v bool)`

SetParametricPartStudioChildInstance sets ParametricPartStudioChildInstance field to given value.

### HasParametricPartStudioChildInstance

`func (o *BTInstanceBase2263AllOf) HasParametricPartStudioChildInstance() bool`

HasParametricPartStudioChildInstance returns a boolean if a field has been set.

### GetParametricPartStudioInstance

`func (o *BTInstanceBase2263AllOf) GetParametricPartStudioInstance() bool`

GetParametricPartStudioInstance returns the ParametricPartStudioInstance field if non-nil, zero value otherwise.

### GetParametricPartStudioInstanceOk

`func (o *BTInstanceBase2263AllOf) GetParametricPartStudioInstanceOk() (*bool, bool)`

GetParametricPartStudioInstanceOk returns a tuple with the ParametricPartStudioInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametricPartStudioInstance

`func (o *BTInstanceBase2263AllOf) SetParametricPartStudioInstance(v bool)`

SetParametricPartStudioInstance sets ParametricPartStudioInstance field to given value.

### HasParametricPartStudioInstance

`func (o *BTInstanceBase2263AllOf) HasParametricPartStudioInstance() bool`

HasParametricPartStudioInstance returns a boolean if a field has been set.

### GetParentSuppressed

`func (o *BTInstanceBase2263AllOf) GetParentSuppressed() bool`

GetParentSuppressed returns the ParentSuppressed field if non-nil, zero value otherwise.

### GetParentSuppressedOk

`func (o *BTInstanceBase2263AllOf) GetParentSuppressedOk() (*bool, bool)`

GetParentSuppressedOk returns a tuple with the ParentSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSuppressed

`func (o *BTInstanceBase2263AllOf) SetParentSuppressed(v bool)`

SetParentSuppressed sets ParentSuppressed field to given value.

### HasParentSuppressed

`func (o *BTInstanceBase2263AllOf) HasParentSuppressed() bool`

HasParentSuppressed returns a boolean if a field has been set.

### GetPartInstance

`func (o *BTInstanceBase2263AllOf) GetPartInstance() bool`

GetPartInstance returns the PartInstance field if non-nil, zero value otherwise.

### GetPartInstanceOk

`func (o *BTInstanceBase2263AllOf) GetPartInstanceOk() (*bool, bool)`

GetPartInstanceOk returns a tuple with the PartInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartInstance

`func (o *BTInstanceBase2263AllOf) SetPartInstance(v bool)`

SetPartInstance sets PartInstance field to given value.

### HasPartInstance

`func (o *BTInstanceBase2263AllOf) HasPartInstance() bool`

HasPartInstance returns a boolean if a field has been set.

### GetReleasable

`func (o *BTInstanceBase2263AllOf) GetReleasable() bool`

GetReleasable returns the Releasable field if non-nil, zero value otherwise.

### GetReleasableOk

`func (o *BTInstanceBase2263AllOf) GetReleasableOk() (*bool, bool)`

GetReleasableOk returns a tuple with the Releasable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasable

`func (o *BTInstanceBase2263AllOf) SetReleasable(v bool)`

SetReleasable sets Releasable field to given value.

### HasReleasable

`func (o *BTInstanceBase2263AllOf) HasReleasable() bool`

HasReleasable returns a boolean if a field has been set.

### GetRevisionCustomData

`func (o *BTInstanceBase2263AllOf) GetRevisionCustomData() BTRevisionCustomData2090`

GetRevisionCustomData returns the RevisionCustomData field if non-nil, zero value otherwise.

### GetRevisionCustomDataOk

`func (o *BTInstanceBase2263AllOf) GetRevisionCustomDataOk() (*BTRevisionCustomData2090, bool)`

GetRevisionCustomDataOk returns a tuple with the RevisionCustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionCustomData

`func (o *BTInstanceBase2263AllOf) SetRevisionCustomData(v BTRevisionCustomData2090)`

SetRevisionCustomData sets RevisionCustomData field to given value.

### HasRevisionCustomData

`func (o *BTInstanceBase2263AllOf) HasRevisionCustomData() bool`

HasRevisionCustomData returns a boolean if a field has been set.

### GetStandardContent

`func (o *BTInstanceBase2263AllOf) GetStandardContent() bool`

GetStandardContent returns the StandardContent field if non-nil, zero value otherwise.

### GetStandardContentOk

`func (o *BTInstanceBase2263AllOf) GetStandardContentOk() (*bool, bool)`

GetStandardContentOk returns a tuple with the StandardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContent

`func (o *BTInstanceBase2263AllOf) SetStandardContent(v bool)`

SetStandardContent sets StandardContent field to given value.

### HasStandardContent

`func (o *BTInstanceBase2263AllOf) HasStandardContent() bool`

HasStandardContent returns a boolean if a field has been set.

### GetStandardContentParametersId

`func (o *BTInstanceBase2263AllOf) GetStandardContentParametersId() string`

GetStandardContentParametersId returns the StandardContentParametersId field if non-nil, zero value otherwise.

### GetStandardContentParametersIdOk

`func (o *BTInstanceBase2263AllOf) GetStandardContentParametersIdOk() (*string, bool)`

GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentParametersId

`func (o *BTInstanceBase2263AllOf) SetStandardContentParametersId(v string)`

SetStandardContentParametersId sets StandardContentParametersId field to given value.

### HasStandardContentParametersId

`func (o *BTInstanceBase2263AllOf) HasStandardContentParametersId() bool`

HasStandardContentParametersId returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTInstanceBase2263AllOf) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTInstanceBase2263AllOf) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTInstanceBase2263AllOf) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTInstanceBase2263AllOf) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressedFieldIndex

`func (o *BTInstanceBase2263AllOf) GetSuppressedFieldIndex() int32`

GetSuppressedFieldIndex returns the SuppressedFieldIndex field if non-nil, zero value otherwise.

### GetSuppressedFieldIndexOk

`func (o *BTInstanceBase2263AllOf) GetSuppressedFieldIndexOk() (*int32, bool)`

GetSuppressedFieldIndexOk returns a tuple with the SuppressedFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedFieldIndex

`func (o *BTInstanceBase2263AllOf) SetSuppressedFieldIndex(v int32)`

SetSuppressedFieldIndex sets SuppressedFieldIndex field to given value.

### HasSuppressedFieldIndex

`func (o *BTInstanceBase2263AllOf) HasSuppressedFieldIndex() bool`

HasSuppressedFieldIndex returns a boolean if a field has been set.

### GetSuppressionConfigured

`func (o *BTInstanceBase2263AllOf) GetSuppressionConfigured() bool`

GetSuppressionConfigured returns the SuppressionConfigured field if non-nil, zero value otherwise.

### GetSuppressionConfiguredOk

`func (o *BTInstanceBase2263AllOf) GetSuppressionConfiguredOk() (*bool, bool)`

GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionConfigured

`func (o *BTInstanceBase2263AllOf) SetSuppressionConfigured(v bool)`

SetSuppressionConfigured sets SuppressionConfigured field to given value.

### HasSuppressionConfigured

`func (o *BTInstanceBase2263AllOf) HasSuppressionConfigured() bool`

HasSuppressionConfigured returns a boolean if a field has been set.

### GetSuppressionState

`func (o *BTInstanceBase2263AllOf) GetSuppressionState() BTMSuppressionState1924`

GetSuppressionState returns the SuppressionState field if non-nil, zero value otherwise.

### GetSuppressionStateOk

`func (o *BTInstanceBase2263AllOf) GetSuppressionStateOk() (*BTMSuppressionState1924, bool)`

GetSuppressionStateOk returns a tuple with the SuppressionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionState

`func (o *BTInstanceBase2263AllOf) SetSuppressionState(v BTMSuppressionState1924)`

SetSuppressionState sets SuppressionState field to given value.

### HasSuppressionState

`func (o *BTInstanceBase2263AllOf) HasSuppressionState() bool`

HasSuppressionState returns a boolean if a field has been set.

### GetSuppressionStateFieldIndex

`func (o *BTInstanceBase2263AllOf) GetSuppressionStateFieldIndex() int32`

GetSuppressionStateFieldIndex returns the SuppressionStateFieldIndex field if non-nil, zero value otherwise.

### GetSuppressionStateFieldIndexOk

`func (o *BTInstanceBase2263AllOf) GetSuppressionStateFieldIndexOk() (*int32, bool)`

GetSuppressionStateFieldIndexOk returns a tuple with the SuppressionStateFieldIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionStateFieldIndex

`func (o *BTInstanceBase2263AllOf) SetSuppressionStateFieldIndex(v int32)`

SetSuppressionStateFieldIndex sets SuppressionStateFieldIndex field to given value.

### HasSuppressionStateFieldIndex

`func (o *BTInstanceBase2263AllOf) HasSuppressionStateFieldIndex() bool`

HasSuppressionStateFieldIndex returns a boolean if a field has been set.

### GetValidRevisionReference

`func (o *BTInstanceBase2263AllOf) GetValidRevisionReference() bool`

GetValidRevisionReference returns the ValidRevisionReference field if non-nil, zero value otherwise.

### GetValidRevisionReferenceOk

`func (o *BTInstanceBase2263AllOf) GetValidRevisionReferenceOk() (*bool, bool)`

GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRevisionReference

`func (o *BTInstanceBase2263AllOf) SetValidRevisionReference(v bool)`

SetValidRevisionReference sets ValidRevisionReference field to given value.

### HasValidRevisionReference

`func (o *BTInstanceBase2263AllOf) HasValidRevisionReference() bool`

HasValidRevisionReference returns a boolean if a field has been set.

### GetVersion

`func (o *BTInstanceBase2263AllOf) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTInstanceBase2263AllOf) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTInstanceBase2263AllOf) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTInstanceBase2263AllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


