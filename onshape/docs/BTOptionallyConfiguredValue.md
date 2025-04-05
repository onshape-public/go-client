# BTOptionallyConfiguredValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredValue** | Pointer to [**BTConfiguredValue**](BTConfiguredValue.md) |  | [optional] 
**Value** | Pointer to **string** | The string value, if not configured | [optional] 

## Methods

### NewBTOptionallyConfiguredValue

`func NewBTOptionallyConfiguredValue() *BTOptionallyConfiguredValue`

NewBTOptionallyConfiguredValue instantiates a new BTOptionallyConfiguredValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTOptionallyConfiguredValueWithDefaults

`func NewBTOptionallyConfiguredValueWithDefaults() *BTOptionallyConfiguredValue`

NewBTOptionallyConfiguredValueWithDefaults instantiates a new BTOptionallyConfiguredValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguredValue

`func (o *BTOptionallyConfiguredValue) GetConfiguredValue() BTConfiguredValue`

GetConfiguredValue returns the ConfiguredValue field if non-nil, zero value otherwise.

### GetConfiguredValueOk

`func (o *BTOptionallyConfiguredValue) GetConfiguredValueOk() (*BTConfiguredValue, bool)`

GetConfiguredValueOk returns a tuple with the ConfiguredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredValue

`func (o *BTOptionallyConfiguredValue) SetConfiguredValue(v BTConfiguredValue)`

SetConfiguredValue sets ConfiguredValue field to given value.

### HasConfiguredValue

`func (o *BTOptionallyConfiguredValue) HasConfiguredValue() bool`

HasConfiguredValue returns a boolean if a field has been set.

### GetValue

`func (o *BTOptionallyConfiguredValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTOptionallyConfiguredValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTOptionallyConfiguredValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BTOptionallyConfiguredValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


