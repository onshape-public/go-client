# BTAppElementReferenceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | **string** | The latest change id for the element, after the edit was committed. | 
**ErrorCode** | Pointer to **int32** | The numeric code identifying the error that occurred, if one occurred. | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**ParentChangeId** | Pointer to **string** | The latest change id for the element, before the edit was made. | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** | The id of the transaction in which the edit was applied. | [optional] 

## Methods

### NewBTAppElementReferenceInfo

`func NewBTAppElementReferenceInfo(changeId string, ) *BTAppElementReferenceInfo`

NewBTAppElementReferenceInfo instantiates a new BTAppElementReferenceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementReferenceInfoWithDefaults

`func NewBTAppElementReferenceInfoWithDefaults() *BTAppElementReferenceInfo`

NewBTAppElementReferenceInfoWithDefaults instantiates a new BTAppElementReferenceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *BTAppElementReferenceInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppElementReferenceInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppElementReferenceInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.


### GetErrorCode

`func (o *BTAppElementReferenceInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementReferenceInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementReferenceInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementReferenceInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementReferenceInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementReferenceInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementReferenceInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementReferenceInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementReferenceInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementReferenceInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementReferenceInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementReferenceInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppElementReferenceInfo) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppElementReferenceInfo) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppElementReferenceInfo) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppElementReferenceInfo) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetReferenceId

`func (o *BTAppElementReferenceInfo) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *BTAppElementReferenceInfo) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *BTAppElementReferenceInfo) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *BTAppElementReferenceInfo) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetTransactionId

`func (o *BTAppElementReferenceInfo) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BTAppElementReferenceInfo) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BTAppElementReferenceInfo) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BTAppElementReferenceInfo) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


