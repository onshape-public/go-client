# ServerVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enum** | Pointer to **[]string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewServerVariable

`func NewServerVariable() *ServerVariable`

NewServerVariable instantiates a new ServerVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVariableWithDefaults

`func NewServerVariableWithDefaults() *ServerVariable`

NewServerVariableWithDefaults instantiates a new ServerVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *ServerVariable) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ServerVariable) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ServerVariable) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ServerVariable) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *ServerVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnum

`func (o *ServerVariable) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *ServerVariable) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *ServerVariable) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *ServerVariable) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetExtensions

`func (o *ServerVariable) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ServerVariable) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ServerVariable) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ServerVariable) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


