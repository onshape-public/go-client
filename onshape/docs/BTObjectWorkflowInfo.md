# BTObjectWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBeDiscarded** | Pointer to **bool** | Whether workflowable object can be discarded. | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsDiscarded** | Pointer to **bool** | Whether workflowable object has been discarded. | [optional] 
**IsFrozen** | Pointer to **bool** | Whether workflowable object has reached terminal state and is frozen. | [optional] 
**MetadataState** | Pointer to [**BTMetadataStateType**](BTMetadataStateType.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObjectType** | Pointer to [**BTAPIWorkflowableType**](BTAPIWorkflowableType.md) |  | [optional] 
**StateId** | Pointer to **string** | The current state of object like SETUP, REJECTED etc. Custom workflows can have any declared state. | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkflowId** | Pointer to **string** | The workflow definition id that governs this object&#39;s states and transitions. | [optional] 

## Methods

### NewBTObjectWorkflowInfo

`func NewBTObjectWorkflowInfo() *BTObjectWorkflowInfo`

NewBTObjectWorkflowInfo instantiates a new BTObjectWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTObjectWorkflowInfoWithDefaults

`func NewBTObjectWorkflowInfoWithDefaults() *BTObjectWorkflowInfo`

NewBTObjectWorkflowInfoWithDefaults instantiates a new BTObjectWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanBeDiscarded

`func (o *BTObjectWorkflowInfo) GetCanBeDiscarded() bool`

GetCanBeDiscarded returns the CanBeDiscarded field if non-nil, zero value otherwise.

### GetCanBeDiscardedOk

`func (o *BTObjectWorkflowInfo) GetCanBeDiscardedOk() (*bool, bool)`

GetCanBeDiscardedOk returns a tuple with the CanBeDiscarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeDiscarded

`func (o *BTObjectWorkflowInfo) SetCanBeDiscarded(v bool)`

SetCanBeDiscarded sets CanBeDiscarded field to given value.

### HasCanBeDiscarded

`func (o *BTObjectWorkflowInfo) HasCanBeDiscarded() bool`

HasCanBeDiscarded returns a boolean if a field has been set.

### GetHref

`func (o *BTObjectWorkflowInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTObjectWorkflowInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTObjectWorkflowInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTObjectWorkflowInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTObjectWorkflowInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTObjectWorkflowInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTObjectWorkflowInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTObjectWorkflowInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDiscarded

`func (o *BTObjectWorkflowInfo) GetIsDiscarded() bool`

GetIsDiscarded returns the IsDiscarded field if non-nil, zero value otherwise.

### GetIsDiscardedOk

`func (o *BTObjectWorkflowInfo) GetIsDiscardedOk() (*bool, bool)`

GetIsDiscardedOk returns a tuple with the IsDiscarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscarded

`func (o *BTObjectWorkflowInfo) SetIsDiscarded(v bool)`

SetIsDiscarded sets IsDiscarded field to given value.

### HasIsDiscarded

`func (o *BTObjectWorkflowInfo) HasIsDiscarded() bool`

HasIsDiscarded returns a boolean if a field has been set.

### GetIsFrozen

`func (o *BTObjectWorkflowInfo) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *BTObjectWorkflowInfo) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *BTObjectWorkflowInfo) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.

### HasIsFrozen

`func (o *BTObjectWorkflowInfo) HasIsFrozen() bool`

HasIsFrozen returns a boolean if a field has been set.

### GetMetadataState

`func (o *BTObjectWorkflowInfo) GetMetadataState() BTMetadataStateType`

GetMetadataState returns the MetadataState field if non-nil, zero value otherwise.

### GetMetadataStateOk

`func (o *BTObjectWorkflowInfo) GetMetadataStateOk() (*BTMetadataStateType, bool)`

GetMetadataStateOk returns a tuple with the MetadataState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataState

`func (o *BTObjectWorkflowInfo) SetMetadataState(v BTMetadataStateType)`

SetMetadataState sets MetadataState field to given value.

### HasMetadataState

`func (o *BTObjectWorkflowInfo) HasMetadataState() bool`

HasMetadataState returns a boolean if a field has been set.

### GetName

`func (o *BTObjectWorkflowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTObjectWorkflowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTObjectWorkflowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTObjectWorkflowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *BTObjectWorkflowInfo) GetObjectType() BTAPIWorkflowableType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTObjectWorkflowInfo) GetObjectTypeOk() (*BTAPIWorkflowableType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTObjectWorkflowInfo) SetObjectType(v BTAPIWorkflowableType)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTObjectWorkflowInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStateId

`func (o *BTObjectWorkflowInfo) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *BTObjectWorkflowInfo) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *BTObjectWorkflowInfo) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *BTObjectWorkflowInfo) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTObjectWorkflowInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTObjectWorkflowInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTObjectWorkflowInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTObjectWorkflowInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkflowId

`func (o *BTObjectWorkflowInfo) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *BTObjectWorkflowInfo) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *BTObjectWorkflowInfo) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *BTObjectWorkflowInfo) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


