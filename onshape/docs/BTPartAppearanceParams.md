# BTPartAppearanceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to [**BTColorParams**](BTColorParams.md) |  | [optional] 
**Opacity** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTPartAppearanceParams

`func NewBTPartAppearanceParams() *BTPartAppearanceParams`

NewBTPartAppearanceParams instantiates a new BTPartAppearanceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPartAppearanceParamsWithDefaults

`func NewBTPartAppearanceParamsWithDefaults() *BTPartAppearanceParams`

NewBTPartAppearanceParamsWithDefaults instantiates a new BTPartAppearanceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *BTPartAppearanceParams) GetColor() BTColorParams`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BTPartAppearanceParams) GetColorOk() (*BTColorParams, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BTPartAppearanceParams) SetColor(v BTColorParams)`

SetColor sets Color field to given value.

### HasColor

`func (o *BTPartAppearanceParams) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetOpacity

`func (o *BTPartAppearanceParams) GetOpacity() int32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *BTPartAppearanceParams) GetOpacityOk() (*int32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *BTPartAppearanceParams) SetOpacity(v int32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *BTPartAppearanceParams) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


