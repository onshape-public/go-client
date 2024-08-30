# BTAppElementIdsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** | &#x60;0: OK (healthy) | 1: INFO | 2: WARNING | 3: ERROR (dangling or view generation call failed) | 4: UNKNOWN&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**SubelementIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBTAppElementIdsInfo

`func NewBTAppElementIdsInfo() *BTAppElementIdsInfo`

NewBTAppElementIdsInfo instantiates a new BTAppElementIdsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementIdsInfoWithDefaults

`func NewBTAppElementIdsInfoWithDefaults() *BTAppElementIdsInfo`

NewBTAppElementIdsInfoWithDefaults instantiates a new BTAppElementIdsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *BTAppElementIdsInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppElementIdsInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppElementIdsInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *BTAppElementIdsInfo) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementIdsInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementIdsInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementIdsInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementIdsInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementIdsInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementIdsInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementIdsInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementIdsInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementIdsInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementIdsInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementIdsInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementIdsInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetSubelementIds

`func (o *BTAppElementIdsInfo) GetSubelementIds() []string`

GetSubelementIds returns the SubelementIds field if non-nil, zero value otherwise.

### GetSubelementIdsOk

`func (o *BTAppElementIdsInfo) GetSubelementIdsOk() (*[]string, bool)`

GetSubelementIdsOk returns a tuple with the SubelementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementIds

`func (o *BTAppElementIdsInfo) SetSubelementIds(v []string)`

SetSubelementIds sets SubelementIds field to given value.

### HasSubelementIds

`func (o *BTAppElementIdsInfo) HasSubelementIds() bool`

HasSubelementIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


