# BTStandardContentSetCustomParametersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**BTStandardContentSetCustomParametersError**](BTStandardContentSetCustomParametersError.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** | If there was an error, this provides a more detailed description of the problem. | [optional] 

## Methods

### NewBTStandardContentSetCustomParametersResponse

`func NewBTStandardContentSetCustomParametersResponse() *BTStandardContentSetCustomParametersResponse`

NewBTStandardContentSetCustomParametersResponse instantiates a new BTStandardContentSetCustomParametersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTStandardContentSetCustomParametersResponseWithDefaults

`func NewBTStandardContentSetCustomParametersResponseWithDefaults() *BTStandardContentSetCustomParametersResponse`

NewBTStandardContentSetCustomParametersResponseWithDefaults instantiates a new BTStandardContentSetCustomParametersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BTStandardContentSetCustomParametersResponse) GetError() BTStandardContentSetCustomParametersError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BTStandardContentSetCustomParametersResponse) GetErrorOk() (*BTStandardContentSetCustomParametersError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BTStandardContentSetCustomParametersResponse) SetError(v BTStandardContentSetCustomParametersError)`

SetError sets Error field to given value.

### HasError

`func (o *BTStandardContentSetCustomParametersResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BTStandardContentSetCustomParametersResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BTStandardContentSetCustomParametersResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BTStandardContentSetCustomParametersResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BTStandardContentSetCustomParametersResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


