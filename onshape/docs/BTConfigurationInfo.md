# BTConfigurationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsStandardContent** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**[]ConfigurationInfoEntry**](ConfigurationInfoEntry.md) |  | [optional] 

## Methods

### NewBTConfigurationInfo

`func NewBTConfigurationInfo() *BTConfigurationInfo`

NewBTConfigurationInfo instantiates a new BTConfigurationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTConfigurationInfoWithDefaults

`func NewBTConfigurationInfoWithDefaults() *BTConfigurationInfo`

NewBTConfigurationInfoWithDefaults instantiates a new BTConfigurationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsStandardContent

`func (o *BTConfigurationInfo) GetIsStandardContent() bool`

GetIsStandardContent returns the IsStandardContent field if non-nil, zero value otherwise.

### GetIsStandardContentOk

`func (o *BTConfigurationInfo) GetIsStandardContentOk() (*bool, bool)`

GetIsStandardContentOk returns a tuple with the IsStandardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandardContent

`func (o *BTConfigurationInfo) SetIsStandardContent(v bool)`

SetIsStandardContent sets IsStandardContent field to given value.

### HasIsStandardContent

`func (o *BTConfigurationInfo) HasIsStandardContent() bool`

HasIsStandardContent returns a boolean if a field has been set.

### GetParameters

`func (o *BTConfigurationInfo) GetParameters() []ConfigurationInfoEntry`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTConfigurationInfo) GetParametersOk() (*[]ConfigurationInfoEntry, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTConfigurationInfo) SetParameters(v []ConfigurationInfoEntry)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTConfigurationInfo) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


