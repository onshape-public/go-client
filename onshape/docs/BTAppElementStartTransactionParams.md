# BTAppElementStartTransactionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The label that will appear in the document&#39;s edit history for this operation. If blank, a value will be auto-generated. | [optional] 
**ParentChangeId** | Pointer to **string** |  | [optional] 
**ReturnError** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTAppElementStartTransactionParams

`func NewBTAppElementStartTransactionParams() *BTAppElementStartTransactionParams`

NewBTAppElementStartTransactionParams instantiates a new BTAppElementStartTransactionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementStartTransactionParamsWithDefaults

`func NewBTAppElementStartTransactionParamsWithDefaults() *BTAppElementStartTransactionParams`

NewBTAppElementStartTransactionParamsWithDefaults instantiates a new BTAppElementStartTransactionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTAppElementStartTransactionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAppElementStartTransactionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAppElementStartTransactionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAppElementStartTransactionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppElementStartTransactionParams) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppElementStartTransactionParams) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppElementStartTransactionParams) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppElementStartTransactionParams) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetReturnError

`func (o *BTAppElementStartTransactionParams) GetReturnError() bool`

GetReturnError returns the ReturnError field if non-nil, zero value otherwise.

### GetReturnErrorOk

`func (o *BTAppElementStartTransactionParams) GetReturnErrorOk() (*bool, bool)`

GetReturnErrorOk returns a tuple with the ReturnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnError

`func (o *BTAppElementStartTransactionParams) SetReturnError(v bool)`

SetReturnError sets ReturnError field to given value.

### HasReturnError

`func (o *BTAppElementStartTransactionParams) HasReturnError() bool`

HasReturnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


