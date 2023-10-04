# ConfigurationEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParameterId** | Pointer to **string** |  | [optional] 
**ParameterValue** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigurationEntry

`func NewConfigurationEntry() *ConfigurationEntry`

NewConfigurationEntry instantiates a new ConfigurationEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationEntryWithDefaults

`func NewConfigurationEntryWithDefaults() *ConfigurationEntry`

NewConfigurationEntryWithDefaults instantiates a new ConfigurationEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameterId

`func (o *ConfigurationEntry) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *ConfigurationEntry) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *ConfigurationEntry) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.

### HasParameterId

`func (o *ConfigurationEntry) HasParameterId() bool`

HasParameterId returns a boolean if a field has been set.

### GetParameterValue

`func (o *ConfigurationEntry) GetParameterValue() string`

GetParameterValue returns the ParameterValue field if non-nil, zero value otherwise.

### GetParameterValueOk

`func (o *ConfigurationEntry) GetParameterValueOk() (*string, bool)`

GetParameterValueOk returns a tuple with the ParameterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterValue

`func (o *ConfigurationEntry) SetParameterValue(v string)`

SetParameterValue sets ParameterValue field to given value.

### HasParameterValue

`func (o *ConfigurationEntry) HasParameterValue() bool`

HasParameterValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


