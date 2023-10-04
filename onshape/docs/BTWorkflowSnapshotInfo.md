# BTWorkflowSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]BTWorkflowActionInfo**](BTWorkflowActionInfo.md) |  | [optional] 
**ApproverIds** | Pointer to **[]string** |  | [optional] 
**CanBeDiscarded** | Pointer to **bool** |  | [optional] 
**DebugMicroversionId** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**IsCreator** | Pointer to **bool** |  | [optional] 
**IsDiscarded** | Pointer to **bool** |  | [optional] 
**IsFrozen** | Pointer to **bool** |  | [optional] 
**IsSetup** | Pointer to **bool** |  | [optional] 
**MetadataState** | Pointer to **string** |  | [optional] 
**NotifierIds** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to [**BTWorkflowStateInfo**](BTWorkflowStateInfo.md) |  | [optional] 
**UsesExternalPlm** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTWorkflowSnapshotInfo

`func NewBTWorkflowSnapshotInfo() *BTWorkflowSnapshotInfo`

NewBTWorkflowSnapshotInfo instantiates a new BTWorkflowSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowSnapshotInfoWithDefaults

`func NewBTWorkflowSnapshotInfoWithDefaults() *BTWorkflowSnapshotInfo`

NewBTWorkflowSnapshotInfoWithDefaults instantiates a new BTWorkflowSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *BTWorkflowSnapshotInfo) GetActions() []BTWorkflowActionInfo`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BTWorkflowSnapshotInfo) GetActionsOk() (*[]BTWorkflowActionInfo, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BTWorkflowSnapshotInfo) SetActions(v []BTWorkflowActionInfo)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BTWorkflowSnapshotInfo) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApproverIds

`func (o *BTWorkflowSnapshotInfo) GetApproverIds() []string`

GetApproverIds returns the ApproverIds field if non-nil, zero value otherwise.

### GetApproverIdsOk

`func (o *BTWorkflowSnapshotInfo) GetApproverIdsOk() (*[]string, bool)`

GetApproverIdsOk returns a tuple with the ApproverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverIds

`func (o *BTWorkflowSnapshotInfo) SetApproverIds(v []string)`

SetApproverIds sets ApproverIds field to given value.

### HasApproverIds

`func (o *BTWorkflowSnapshotInfo) HasApproverIds() bool`

HasApproverIds returns a boolean if a field has been set.

### GetCanBeDiscarded

`func (o *BTWorkflowSnapshotInfo) GetCanBeDiscarded() bool`

GetCanBeDiscarded returns the CanBeDiscarded field if non-nil, zero value otherwise.

### GetCanBeDiscardedOk

`func (o *BTWorkflowSnapshotInfo) GetCanBeDiscardedOk() (*bool, bool)`

GetCanBeDiscardedOk returns a tuple with the CanBeDiscarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeDiscarded

`func (o *BTWorkflowSnapshotInfo) SetCanBeDiscarded(v bool)`

SetCanBeDiscarded sets CanBeDiscarded field to given value.

### HasCanBeDiscarded

`func (o *BTWorkflowSnapshotInfo) HasCanBeDiscarded() bool`

HasCanBeDiscarded returns a boolean if a field has been set.

### GetDebugMicroversionId

`func (o *BTWorkflowSnapshotInfo) GetDebugMicroversionId() string`

GetDebugMicroversionId returns the DebugMicroversionId field if non-nil, zero value otherwise.

### GetDebugMicroversionIdOk

`func (o *BTWorkflowSnapshotInfo) GetDebugMicroversionIdOk() (*string, bool)`

GetDebugMicroversionIdOk returns a tuple with the DebugMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMicroversionId

`func (o *BTWorkflowSnapshotInfo) SetDebugMicroversionId(v string)`

SetDebugMicroversionId sets DebugMicroversionId field to given value.

### HasDebugMicroversionId

`func (o *BTWorkflowSnapshotInfo) HasDebugMicroversionId() bool`

HasDebugMicroversionId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BTWorkflowSnapshotInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BTWorkflowSnapshotInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BTWorkflowSnapshotInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BTWorkflowSnapshotInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetIsCreator

`func (o *BTWorkflowSnapshotInfo) GetIsCreator() bool`

GetIsCreator returns the IsCreator field if non-nil, zero value otherwise.

### GetIsCreatorOk

`func (o *BTWorkflowSnapshotInfo) GetIsCreatorOk() (*bool, bool)`

GetIsCreatorOk returns a tuple with the IsCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreator

`func (o *BTWorkflowSnapshotInfo) SetIsCreator(v bool)`

SetIsCreator sets IsCreator field to given value.

### HasIsCreator

`func (o *BTWorkflowSnapshotInfo) HasIsCreator() bool`

HasIsCreator returns a boolean if a field has been set.

### GetIsDiscarded

`func (o *BTWorkflowSnapshotInfo) GetIsDiscarded() bool`

GetIsDiscarded returns the IsDiscarded field if non-nil, zero value otherwise.

### GetIsDiscardedOk

`func (o *BTWorkflowSnapshotInfo) GetIsDiscardedOk() (*bool, bool)`

GetIsDiscardedOk returns a tuple with the IsDiscarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscarded

`func (o *BTWorkflowSnapshotInfo) SetIsDiscarded(v bool)`

SetIsDiscarded sets IsDiscarded field to given value.

### HasIsDiscarded

`func (o *BTWorkflowSnapshotInfo) HasIsDiscarded() bool`

HasIsDiscarded returns a boolean if a field has been set.

### GetIsFrozen

`func (o *BTWorkflowSnapshotInfo) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *BTWorkflowSnapshotInfo) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *BTWorkflowSnapshotInfo) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.

### HasIsFrozen

`func (o *BTWorkflowSnapshotInfo) HasIsFrozen() bool`

HasIsFrozen returns a boolean if a field has been set.

### GetIsSetup

`func (o *BTWorkflowSnapshotInfo) GetIsSetup() bool`

GetIsSetup returns the IsSetup field if non-nil, zero value otherwise.

### GetIsSetupOk

`func (o *BTWorkflowSnapshotInfo) GetIsSetupOk() (*bool, bool)`

GetIsSetupOk returns a tuple with the IsSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetup

`func (o *BTWorkflowSnapshotInfo) SetIsSetup(v bool)`

SetIsSetup sets IsSetup field to given value.

### HasIsSetup

`func (o *BTWorkflowSnapshotInfo) HasIsSetup() bool`

HasIsSetup returns a boolean if a field has been set.

### GetMetadataState

`func (o *BTWorkflowSnapshotInfo) GetMetadataState() string`

GetMetadataState returns the MetadataState field if non-nil, zero value otherwise.

### GetMetadataStateOk

`func (o *BTWorkflowSnapshotInfo) GetMetadataStateOk() (*string, bool)`

GetMetadataStateOk returns a tuple with the MetadataState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataState

`func (o *BTWorkflowSnapshotInfo) SetMetadataState(v string)`

SetMetadataState sets MetadataState field to given value.

### HasMetadataState

`func (o *BTWorkflowSnapshotInfo) HasMetadataState() bool`

HasMetadataState returns a boolean if a field has been set.

### GetNotifierIds

`func (o *BTWorkflowSnapshotInfo) GetNotifierIds() []string`

GetNotifierIds returns the NotifierIds field if non-nil, zero value otherwise.

### GetNotifierIdsOk

`func (o *BTWorkflowSnapshotInfo) GetNotifierIdsOk() (*[]string, bool)`

GetNotifierIdsOk returns a tuple with the NotifierIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierIds

`func (o *BTWorkflowSnapshotInfo) SetNotifierIds(v []string)`

SetNotifierIds sets NotifierIds field to given value.

### HasNotifierIds

`func (o *BTWorkflowSnapshotInfo) HasNotifierIds() bool`

HasNotifierIds returns a boolean if a field has been set.

### GetState

`func (o *BTWorkflowSnapshotInfo) GetState() BTWorkflowStateInfo`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTWorkflowSnapshotInfo) GetStateOk() (*BTWorkflowStateInfo, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTWorkflowSnapshotInfo) SetState(v BTWorkflowStateInfo)`

SetState sets State field to given value.

### HasState

`func (o *BTWorkflowSnapshotInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsesExternalPlm

`func (o *BTWorkflowSnapshotInfo) GetUsesExternalPlm() bool`

GetUsesExternalPlm returns the UsesExternalPlm field if non-nil, zero value otherwise.

### GetUsesExternalPlmOk

`func (o *BTWorkflowSnapshotInfo) GetUsesExternalPlmOk() (*bool, bool)`

GetUsesExternalPlmOk returns a tuple with the UsesExternalPlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExternalPlm

`func (o *BTWorkflowSnapshotInfo) SetUsesExternalPlm(v bool)`

SetUsesExternalPlm sets UsesExternalPlm field to given value.

### HasUsesExternalPlm

`func (o *BTWorkflowSnapshotInfo) HasUsesExternalPlm() bool`

HasUsesExternalPlm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


