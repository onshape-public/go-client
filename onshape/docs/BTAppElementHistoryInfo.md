# BTAppElementHistoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**[]BTAppElementHistoryEntryInfo**](BTAppElementHistoryEntryInfo.md) |  | [optional] 
**ErrorCode** | Pointer to **int32** | The numeric code identifying the error that occurred, if one occurred. | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 

## Methods

### NewBTAppElementHistoryInfo

`func NewBTAppElementHistoryInfo() *BTAppElementHistoryInfo`

NewBTAppElementHistoryInfo instantiates a new BTAppElementHistoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementHistoryInfoWithDefaults

`func NewBTAppElementHistoryInfoWithDefaults() *BTAppElementHistoryInfo`

NewBTAppElementHistoryInfoWithDefaults instantiates a new BTAppElementHistoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *BTAppElementHistoryInfo) GetChanges() []BTAppElementHistoryEntryInfo`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *BTAppElementHistoryInfo) GetChangesOk() (*[]BTAppElementHistoryEntryInfo, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *BTAppElementHistoryInfo) SetChanges(v []BTAppElementHistoryEntryInfo)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *BTAppElementHistoryInfo) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementHistoryInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementHistoryInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementHistoryInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementHistoryInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementHistoryInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementHistoryInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementHistoryInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementHistoryInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementHistoryInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementHistoryInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementHistoryInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementHistoryInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


