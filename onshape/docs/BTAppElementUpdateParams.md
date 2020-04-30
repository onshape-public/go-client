# BTAppElementUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**[]BTAppElementChangeParams**](BTAppElementChangeParams.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentChangeId** | Pointer to **string** |  | [optional] 
**ReturnError** | Pointer to **bool** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAppElementUpdateParams

`func NewBTAppElementUpdateParams() *BTAppElementUpdateParams`

NewBTAppElementUpdateParams instantiates a new BTAppElementUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementUpdateParamsWithDefaults

`func NewBTAppElementUpdateParamsWithDefaults() *BTAppElementUpdateParams`

NewBTAppElementUpdateParamsWithDefaults instantiates a new BTAppElementUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *BTAppElementUpdateParams) GetChanges() []BTAppElementChangeParams`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *BTAppElementUpdateParams) GetChangesOk() (*[]BTAppElementChangeParams, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *BTAppElementUpdateParams) SetChanges(v []BTAppElementChangeParams)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *BTAppElementUpdateParams) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetDescription

`func (o *BTAppElementUpdateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAppElementUpdateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAppElementUpdateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAppElementUpdateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppElementUpdateParams) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppElementUpdateParams) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppElementUpdateParams) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppElementUpdateParams) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetReturnError

`func (o *BTAppElementUpdateParams) GetReturnError() bool`

GetReturnError returns the ReturnError field if non-nil, zero value otherwise.

### GetReturnErrorOk

`func (o *BTAppElementUpdateParams) GetReturnErrorOk() (*bool, bool)`

GetReturnErrorOk returns a tuple with the ReturnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnError

`func (o *BTAppElementUpdateParams) SetReturnError(v bool)`

SetReturnError sets ReturnError field to given value.

### HasReturnError

`func (o *BTAppElementUpdateParams) HasReturnError() bool`

HasReturnError returns a boolean if a field has been set.

### GetTransactionId

`func (o *BTAppElementUpdateParams) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BTAppElementUpdateParams) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BTAppElementUpdateParams) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BTAppElementUpdateParams) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


