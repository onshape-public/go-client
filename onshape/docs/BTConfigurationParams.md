# BTConfigurationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]ConfigurationEntry**](ConfigurationEntry.md) |  | [optional] 
**StandardContentParametersId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTConfigurationParams

`func NewBTConfigurationParams() *BTConfigurationParams`

NewBTConfigurationParams instantiates a new BTConfigurationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTConfigurationParamsWithDefaults

`func NewBTConfigurationParamsWithDefaults() *BTConfigurationParams`

NewBTConfigurationParamsWithDefaults instantiates a new BTConfigurationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *BTConfigurationParams) GetParameters() []ConfigurationEntry`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTConfigurationParams) GetParametersOk() (*[]ConfigurationEntry, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTConfigurationParams) SetParameters(v []ConfigurationEntry)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTConfigurationParams) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetStandardContentParametersId

`func (o *BTConfigurationParams) GetStandardContentParametersId() string`

GetStandardContentParametersId returns the StandardContentParametersId field if non-nil, zero value otherwise.

### GetStandardContentParametersIdOk

`func (o *BTConfigurationParams) GetStandardContentParametersIdOk() (*string, bool)`

GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContentParametersId

`func (o *BTConfigurationParams) SetStandardContentParametersId(v string)`

SetStandardContentParametersId sets StandardContentParametersId field to given value.

### HasStandardContentParametersId

`func (o *BTConfigurationParams) HasStandardContentParametersId() bool`

HasStandardContentParametersId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


