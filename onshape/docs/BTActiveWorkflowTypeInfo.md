# BTActiveWorkflowTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasInactiveCustomWorkflows** | Pointer to **bool** |  | [optional] 
**PickableWorkflows** | Pointer to [**[]BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**Workflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 

## Methods

### NewBTActiveWorkflowTypeInfo

`func NewBTActiveWorkflowTypeInfo() *BTActiveWorkflowTypeInfo`

NewBTActiveWorkflowTypeInfo instantiates a new BTActiveWorkflowTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTActiveWorkflowTypeInfoWithDefaults

`func NewBTActiveWorkflowTypeInfoWithDefaults() *BTActiveWorkflowTypeInfo`

NewBTActiveWorkflowTypeInfoWithDefaults instantiates a new BTActiveWorkflowTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasInactiveCustomWorkflows

`func (o *BTActiveWorkflowTypeInfo) GetHasInactiveCustomWorkflows() bool`

GetHasInactiveCustomWorkflows returns the HasInactiveCustomWorkflows field if non-nil, zero value otherwise.

### GetHasInactiveCustomWorkflowsOk

`func (o *BTActiveWorkflowTypeInfo) GetHasInactiveCustomWorkflowsOk() (*bool, bool)`

GetHasInactiveCustomWorkflowsOk returns a tuple with the HasInactiveCustomWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInactiveCustomWorkflows

`func (o *BTActiveWorkflowTypeInfo) SetHasInactiveCustomWorkflows(v bool)`

SetHasInactiveCustomWorkflows sets HasInactiveCustomWorkflows field to given value.

### HasHasInactiveCustomWorkflows

`func (o *BTActiveWorkflowTypeInfo) HasHasInactiveCustomWorkflows() bool`

HasHasInactiveCustomWorkflows returns a boolean if a field has been set.

### GetPickableWorkflows

`func (o *BTActiveWorkflowTypeInfo) GetPickableWorkflows() []BTPublishedWorkflowInfo`

GetPickableWorkflows returns the PickableWorkflows field if non-nil, zero value otherwise.

### GetPickableWorkflowsOk

`func (o *BTActiveWorkflowTypeInfo) GetPickableWorkflowsOk() (*[]BTPublishedWorkflowInfo, bool)`

GetPickableWorkflowsOk returns a tuple with the PickableWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickableWorkflows

`func (o *BTActiveWorkflowTypeInfo) SetPickableWorkflows(v []BTPublishedWorkflowInfo)`

SetPickableWorkflows sets PickableWorkflows field to given value.

### HasPickableWorkflows

`func (o *BTActiveWorkflowTypeInfo) HasPickableWorkflows() bool`

HasPickableWorkflows returns a boolean if a field has been set.

### GetWorkflow

`func (o *BTActiveWorkflowTypeInfo) GetWorkflow() BTPublishedWorkflowInfo`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BTActiveWorkflowTypeInfo) GetWorkflowOk() (*BTPublishedWorkflowInfo, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BTActiveWorkflowTypeInfo) SetWorkflow(v BTPublishedWorkflowInfo)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *BTActiveWorkflowTypeInfo) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


