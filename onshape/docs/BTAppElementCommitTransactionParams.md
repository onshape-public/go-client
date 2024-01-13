# BTAppElementCommitTransactionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMerge** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** | The label that will appear in the document&#39;s edit history for this operation. If blank, a value will be auto-generated. | [optional] 
**ReturnError** | Pointer to **bool** |  | [optional] 
**TransactionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBTAppElementCommitTransactionParams

`func NewBTAppElementCommitTransactionParams() *BTAppElementCommitTransactionParams`

NewBTAppElementCommitTransactionParams instantiates a new BTAppElementCommitTransactionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementCommitTransactionParamsWithDefaults

`func NewBTAppElementCommitTransactionParamsWithDefaults() *BTAppElementCommitTransactionParams`

NewBTAppElementCommitTransactionParamsWithDefaults instantiates a new BTAppElementCommitTransactionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMerge

`func (o *BTAppElementCommitTransactionParams) GetAllowMerge() bool`

GetAllowMerge returns the AllowMerge field if non-nil, zero value otherwise.

### GetAllowMergeOk

`func (o *BTAppElementCommitTransactionParams) GetAllowMergeOk() (*bool, bool)`

GetAllowMergeOk returns a tuple with the AllowMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMerge

`func (o *BTAppElementCommitTransactionParams) SetAllowMerge(v bool)`

SetAllowMerge sets AllowMerge field to given value.

### HasAllowMerge

`func (o *BTAppElementCommitTransactionParams) HasAllowMerge() bool`

HasAllowMerge returns a boolean if a field has been set.

### GetDescription

`func (o *BTAppElementCommitTransactionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAppElementCommitTransactionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAppElementCommitTransactionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAppElementCommitTransactionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReturnError

`func (o *BTAppElementCommitTransactionParams) GetReturnError() bool`

GetReturnError returns the ReturnError field if non-nil, zero value otherwise.

### GetReturnErrorOk

`func (o *BTAppElementCommitTransactionParams) GetReturnErrorOk() (*bool, bool)`

GetReturnErrorOk returns a tuple with the ReturnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnError

`func (o *BTAppElementCommitTransactionParams) SetReturnError(v bool)`

SetReturnError sets ReturnError field to given value.

### HasReturnError

`func (o *BTAppElementCommitTransactionParams) HasReturnError() bool`

HasReturnError returns a boolean if a field has been set.

### GetTransactionIds

`func (o *BTAppElementCommitTransactionParams) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *BTAppElementCommitTransactionParams) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *BTAppElementCommitTransactionParams) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *BTAppElementCommitTransactionParams) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


