# BTAppElementModifyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | **string** | The latest change id for the element, after the edit was committed. | 
**ElementId** | Pointer to **string** | The id of the edited element. | [optional] 
**ElementIds** | Pointer to **[]string** | The ids of the edited elements, if multiple elements were edited. | [optional] 
**ErrorCode** | Pointer to **int32** | The numeric code identifying the error that occurred, if one occurred. | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**JsonDifference** | Pointer to [**BTDiffJsonResponse2725**](BTDiffJsonResponse2725.md) |  | [optional] 
**ParentChangeId** | Pointer to **string** | The latest change id for the element, before the edit was made. | [optional] 
**PropertyEditsMerged** | Pointer to **bool** | When committing a transaction, this field indicates if the properties of the application element were changed after the transaction was created. | [optional] 
**TransactionId** | Pointer to **string** | The id of the transaction in which the edit was applied. | [optional] 

## Methods

### NewBTAppElementModifyInfo

`func NewBTAppElementModifyInfo(changeId string, ) *BTAppElementModifyInfo`

NewBTAppElementModifyInfo instantiates a new BTAppElementModifyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementModifyInfoWithDefaults

`func NewBTAppElementModifyInfoWithDefaults() *BTAppElementModifyInfo`

NewBTAppElementModifyInfoWithDefaults instantiates a new BTAppElementModifyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *BTAppElementModifyInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppElementModifyInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppElementModifyInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.


### GetElementId

`func (o *BTAppElementModifyInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTAppElementModifyInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTAppElementModifyInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTAppElementModifyInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementIds

`func (o *BTAppElementModifyInfo) GetElementIds() []string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTAppElementModifyInfo) GetElementIdsOk() (*[]string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTAppElementModifyInfo) SetElementIds(v []string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTAppElementModifyInfo) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementModifyInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementModifyInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementModifyInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementModifyInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementModifyInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementModifyInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementModifyInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementModifyInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementModifyInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementModifyInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementModifyInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementModifyInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetJsonDifference

`func (o *BTAppElementModifyInfo) GetJsonDifference() BTDiffJsonResponse2725`

GetJsonDifference returns the JsonDifference field if non-nil, zero value otherwise.

### GetJsonDifferenceOk

`func (o *BTAppElementModifyInfo) GetJsonDifferenceOk() (*BTDiffJsonResponse2725, bool)`

GetJsonDifferenceOk returns a tuple with the JsonDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonDifference

`func (o *BTAppElementModifyInfo) SetJsonDifference(v BTDiffJsonResponse2725)`

SetJsonDifference sets JsonDifference field to given value.

### HasJsonDifference

`func (o *BTAppElementModifyInfo) HasJsonDifference() bool`

HasJsonDifference returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppElementModifyInfo) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppElementModifyInfo) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppElementModifyInfo) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppElementModifyInfo) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetPropertyEditsMerged

`func (o *BTAppElementModifyInfo) GetPropertyEditsMerged() bool`

GetPropertyEditsMerged returns the PropertyEditsMerged field if non-nil, zero value otherwise.

### GetPropertyEditsMergedOk

`func (o *BTAppElementModifyInfo) GetPropertyEditsMergedOk() (*bool, bool)`

GetPropertyEditsMergedOk returns a tuple with the PropertyEditsMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEditsMerged

`func (o *BTAppElementModifyInfo) SetPropertyEditsMerged(v bool)`

SetPropertyEditsMerged sets PropertyEditsMerged field to given value.

### HasPropertyEditsMerged

`func (o *BTAppElementModifyInfo) HasPropertyEditsMerged() bool`

HasPropertyEditsMerged returns a boolean if a field has been set.

### GetTransactionId

`func (o *BTAppElementModifyInfo) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BTAppElementModifyInfo) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BTAppElementModifyInfo) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BTAppElementModifyInfo) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


