# OperationResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Default** | Pointer to [**ApiResponse**](ApiResponse.md) |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 

## Methods

### NewOperationResponses

`func NewOperationResponses() *OperationResponses`

NewOperationResponses instantiates a new OperationResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponsesWithDefaults

`func NewOperationResponsesWithDefaults() *OperationResponses`

NewOperationResponsesWithDefaults instantiates a new OperationResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *OperationResponses) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *OperationResponses) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *OperationResponses) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *OperationResponses) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetDefault

`func (o *OperationResponses) GetDefault() ApiResponse`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *OperationResponses) GetDefaultOk() (*ApiResponse, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *OperationResponses) SetDefault(v ApiResponse)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *OperationResponses) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEmpty

`func (o *OperationResponses) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *OperationResponses) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *OperationResponses) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *OperationResponses) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


