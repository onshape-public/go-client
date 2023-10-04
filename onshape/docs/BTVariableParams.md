# BTVariableParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Variable description | [optional] 
**Expression** | **string** | Variable definition expression | 
**Name** | **string** | Variable name | 
**Type** | **string** | VariableType name, from FeatureScript VariableType | 

## Methods

### NewBTVariableParams

`func NewBTVariableParams(expression string, name string, type_ string, ) *BTVariableParams`

NewBTVariableParams instantiates a new BTVariableParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVariableParamsWithDefaults

`func NewBTVariableParamsWithDefaults() *BTVariableParams`

NewBTVariableParamsWithDefaults instantiates a new BTVariableParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTVariableParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTVariableParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTVariableParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTVariableParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpression

`func (o *BTVariableParams) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *BTVariableParams) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *BTVariableParams) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetName

`func (o *BTVariableParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTVariableParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTVariableParams) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BTVariableParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTVariableParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTVariableParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


