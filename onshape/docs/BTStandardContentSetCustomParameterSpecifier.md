# BTStandardContentSetCustomParameterSpecifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomParameters** | Pointer to [**[]BTStandardContentPropertyIdAndValue**](BTStandardContentPropertyIdAndValue.md) | Non-driving custom parameters whose values are to be set. | [optional] 
**Parameters** | Pointer to [**[]BTStandardContentParameterIdAndValue**](BTStandardContentParameterIdAndValue.md) | Parameters that drive configuration. Used to specify the standard content component to which the custom parameter values are associated. | [optional] 

## Methods

### NewBTStandardContentSetCustomParameterSpecifier

`func NewBTStandardContentSetCustomParameterSpecifier() *BTStandardContentSetCustomParameterSpecifier`

NewBTStandardContentSetCustomParameterSpecifier instantiates a new BTStandardContentSetCustomParameterSpecifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTStandardContentSetCustomParameterSpecifierWithDefaults

`func NewBTStandardContentSetCustomParameterSpecifierWithDefaults() *BTStandardContentSetCustomParameterSpecifier`

NewBTStandardContentSetCustomParameterSpecifierWithDefaults instantiates a new BTStandardContentSetCustomParameterSpecifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomParameters

`func (o *BTStandardContentSetCustomParameterSpecifier) GetCustomParameters() []BTStandardContentPropertyIdAndValue`

GetCustomParameters returns the CustomParameters field if non-nil, zero value otherwise.

### GetCustomParametersOk

`func (o *BTStandardContentSetCustomParameterSpecifier) GetCustomParametersOk() (*[]BTStandardContentPropertyIdAndValue, bool)`

GetCustomParametersOk returns a tuple with the CustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameters

`func (o *BTStandardContentSetCustomParameterSpecifier) SetCustomParameters(v []BTStandardContentPropertyIdAndValue)`

SetCustomParameters sets CustomParameters field to given value.

### HasCustomParameters

`func (o *BTStandardContentSetCustomParameterSpecifier) HasCustomParameters() bool`

HasCustomParameters returns a boolean if a field has been set.

### GetParameters

`func (o *BTStandardContentSetCustomParameterSpecifier) GetParameters() []BTStandardContentParameterIdAndValue`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTStandardContentSetCustomParameterSpecifier) GetParametersOk() (*[]BTStandardContentParameterIdAndValue, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTStandardContentSetCustomParameterSpecifier) SetParameters(v []BTStandardContentParameterIdAndValue)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTStandardContentSetCustomParameterSpecifier) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


