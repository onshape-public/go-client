# ServerVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 

## Methods

### NewServerVariables

`func NewServerVariables() *ServerVariables`

NewServerVariables instantiates a new ServerVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVariablesWithDefaults

`func NewServerVariablesWithDefaults() *ServerVariables`

NewServerVariablesWithDefaults instantiates a new ServerVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *ServerVariables) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ServerVariables) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ServerVariables) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ServerVariables) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetEmpty

`func (o *ServerVariables) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *ServerVariables) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *ServerVariables) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *ServerVariables) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


