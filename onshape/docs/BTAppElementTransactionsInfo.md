# BTAppElementTransactionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementTransactions** | Pointer to [**[]BTElementTransaction**](BTElementTransaction.md) |  | [optional] 
**ErrorCode** | Pointer to **int32** | The numeric code identifying the error that occurred, if one occurred. | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 

## Methods

### NewBTAppElementTransactionsInfo

`func NewBTAppElementTransactionsInfo() *BTAppElementTransactionsInfo`

NewBTAppElementTransactionsInfo instantiates a new BTAppElementTransactionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementTransactionsInfoWithDefaults

`func NewBTAppElementTransactionsInfoWithDefaults() *BTAppElementTransactionsInfo`

NewBTAppElementTransactionsInfoWithDefaults instantiates a new BTAppElementTransactionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementTransactions

`func (o *BTAppElementTransactionsInfo) GetElementTransactions() []BTElementTransaction`

GetElementTransactions returns the ElementTransactions field if non-nil, zero value otherwise.

### GetElementTransactionsOk

`func (o *BTAppElementTransactionsInfo) GetElementTransactionsOk() (*[]BTElementTransaction, bool)`

GetElementTransactionsOk returns a tuple with the ElementTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementTransactions

`func (o *BTAppElementTransactionsInfo) SetElementTransactions(v []BTElementTransaction)`

SetElementTransactions sets ElementTransactions field to given value.

### HasElementTransactions

`func (o *BTAppElementTransactionsInfo) HasElementTransactions() bool`

HasElementTransactions returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementTransactionsInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementTransactionsInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementTransactionsInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementTransactionsInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementTransactionsInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementTransactionsInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementTransactionsInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementTransactionsInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementTransactionsInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementTransactionsInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementTransactionsInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementTransactionsInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


