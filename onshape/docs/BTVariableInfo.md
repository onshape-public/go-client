# BTVariableInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Variable description | [optional] 
**Expression** | **string** | Variable expression | 
**Name** | **string** | Variable name | 
**Type** | [**GBTVariableType**](GBTVariableType.md) |  | 
**Value** | **string** | Variable formatted value | 

## Methods

### NewBTVariableInfo

`func NewBTVariableInfo(expression string, name string, type_ GBTVariableType, value string, ) *BTVariableInfo`

NewBTVariableInfo instantiates a new BTVariableInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVariableInfoWithDefaults

`func NewBTVariableInfoWithDefaults() *BTVariableInfo`

NewBTVariableInfoWithDefaults instantiates a new BTVariableInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTVariableInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTVariableInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTVariableInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTVariableInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpression

`func (o *BTVariableInfo) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *BTVariableInfo) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *BTVariableInfo) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetName

`func (o *BTVariableInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTVariableInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTVariableInfo) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BTVariableInfo) GetType() GBTVariableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTVariableInfo) GetTypeOk() (*GBTVariableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTVariableInfo) SetType(v GBTVariableType)`

SetType sets Type field to given value.


### GetValue

`func (o *BTVariableInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTVariableInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTVariableInfo) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


