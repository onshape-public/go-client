# BTConfiguredValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationParameterId** | Pointer to **string** | The configuration parameter configuring this value, if configured | [optional] 
**ConfigurationToValue** | Pointer to **map[string]string** | Configuration to value, required if configuration parameter id is specified | [optional] 

## Methods

### NewBTConfiguredValue

`func NewBTConfiguredValue() *BTConfiguredValue`

NewBTConfiguredValue instantiates a new BTConfiguredValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTConfiguredValueWithDefaults

`func NewBTConfiguredValueWithDefaults() *BTConfiguredValue`

NewBTConfiguredValueWithDefaults instantiates a new BTConfiguredValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationParameterId

`func (o *BTConfiguredValue) GetConfigurationParameterId() string`

GetConfigurationParameterId returns the ConfigurationParameterId field if non-nil, zero value otherwise.

### GetConfigurationParameterIdOk

`func (o *BTConfiguredValue) GetConfigurationParameterIdOk() (*string, bool)`

GetConfigurationParameterIdOk returns a tuple with the ConfigurationParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParameterId

`func (o *BTConfiguredValue) SetConfigurationParameterId(v string)`

SetConfigurationParameterId sets ConfigurationParameterId field to given value.

### HasConfigurationParameterId

`func (o *BTConfiguredValue) HasConfigurationParameterId() bool`

HasConfigurationParameterId returns a boolean if a field has been set.

### GetConfigurationToValue

`func (o *BTConfiguredValue) GetConfigurationToValue() map[string]string`

GetConfigurationToValue returns the ConfigurationToValue field if non-nil, zero value otherwise.

### GetConfigurationToValueOk

`func (o *BTConfiguredValue) GetConfigurationToValueOk() (*map[string]string, bool)`

GetConfigurationToValueOk returns a tuple with the ConfigurationToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationToValue

`func (o *BTConfiguredValue) SetConfigurationToValue(v map[string]string)`

SetConfigurationToValue sets ConfigurationToValue field to given value.

### HasConfigurationToValue

`func (o *BTConfiguredValue) HasConfigurationToValue() bool`

HasConfigurationToValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


