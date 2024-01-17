# BTDrawingModificationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The label that will appear in the document&#39;s edit history for this operation. If blank, a value will be auto-generated. | [optional] 
**JsonRequests** | Pointer to [**[]BTBDrawingOperationParams**](BTBDrawingOperationParams.md) | Array of drawing modification operations. | [optional] 

## Methods

### NewBTDrawingModificationParams

`func NewBTDrawingModificationParams() *BTDrawingModificationParams`

NewBTDrawingModificationParams instantiates a new BTDrawingModificationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDrawingModificationParamsWithDefaults

`func NewBTDrawingModificationParamsWithDefaults() *BTDrawingModificationParams`

NewBTDrawingModificationParamsWithDefaults instantiates a new BTDrawingModificationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTDrawingModificationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDrawingModificationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDrawingModificationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDrawingModificationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJsonRequests

`func (o *BTDrawingModificationParams) GetJsonRequests() []BTBDrawingOperationParams`

GetJsonRequests returns the JsonRequests field if non-nil, zero value otherwise.

### GetJsonRequestsOk

`func (o *BTDrawingModificationParams) GetJsonRequestsOk() (*[]BTBDrawingOperationParams, bool)`

GetJsonRequestsOk returns a tuple with the JsonRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonRequests

`func (o *BTDrawingModificationParams) SetJsonRequests(v []BTBDrawingOperationParams)`

SetJsonRequests sets JsonRequests field to given value.

### HasJsonRequests

`func (o *BTDrawingModificationParams) HasJsonRequests() bool`

HasJsonRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


