# BTAppAssociativeDataArrayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** | &#x60;0: OK (healthy) | 1: INFO | 2: WARNING | 3: ERROR (dangling or view generation call failed) | 4: UNKNOWN&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**Items** | Pointer to [**[]BTAssociativeDataInfo**](BTAssociativeDataInfo.md) |  | [optional] 
**ParentChangeId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAppAssociativeDataArrayInfo

`func NewBTAppAssociativeDataArrayInfo() *BTAppAssociativeDataArrayInfo`

NewBTAppAssociativeDataArrayInfo instantiates a new BTAppAssociativeDataArrayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppAssociativeDataArrayInfoWithDefaults

`func NewBTAppAssociativeDataArrayInfoWithDefaults() *BTAppAssociativeDataArrayInfo`

NewBTAppAssociativeDataArrayInfoWithDefaults instantiates a new BTAppAssociativeDataArrayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *BTAppAssociativeDataArrayInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppAssociativeDataArrayInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppAssociativeDataArrayInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *BTAppAssociativeDataArrayInfo) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppAssociativeDataArrayInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppAssociativeDataArrayInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppAssociativeDataArrayInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppAssociativeDataArrayInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppAssociativeDataArrayInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppAssociativeDataArrayInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppAssociativeDataArrayInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppAssociativeDataArrayInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppAssociativeDataArrayInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppAssociativeDataArrayInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppAssociativeDataArrayInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppAssociativeDataArrayInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetItems

`func (o *BTAppAssociativeDataArrayInfo) GetItems() []BTAssociativeDataInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTAppAssociativeDataArrayInfo) GetItemsOk() (*[]BTAssociativeDataInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTAppAssociativeDataArrayInfo) SetItems(v []BTAssociativeDataInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTAppAssociativeDataArrayInfo) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppAssociativeDataArrayInfo) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppAssociativeDataArrayInfo) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppAssociativeDataArrayInfo) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppAssociativeDataArrayInfo) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


