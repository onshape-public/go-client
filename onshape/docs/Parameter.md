# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEmptyValue** | Pointer to **bool** |  | [optional] 
**AllowReserved** | Pointer to **bool** |  | [optional] 
**Content** | Pointer to [**map[string]MediaType**](MediaType.md) |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**Examples** | Pointer to [**map[string]Example**](Example.md) |  | [optional] 
**Explode** | Pointer to **bool** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**In** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Schema** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Style** | Pointer to [**StyleEnum**](StyleEnum.md) |  | [optional] 

## Methods

### NewParameter

`func NewParameter() *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEmptyValue

`func (o *Parameter) GetAllowEmptyValue() bool`

GetAllowEmptyValue returns the AllowEmptyValue field if non-nil, zero value otherwise.

### GetAllowEmptyValueOk

`func (o *Parameter) GetAllowEmptyValueOk() (*bool, bool)`

GetAllowEmptyValueOk returns a tuple with the AllowEmptyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyValue

`func (o *Parameter) SetAllowEmptyValue(v bool)`

SetAllowEmptyValue sets AllowEmptyValue field to given value.

### HasAllowEmptyValue

`func (o *Parameter) HasAllowEmptyValue() bool`

HasAllowEmptyValue returns a boolean if a field has been set.

### GetAllowReserved

`func (o *Parameter) GetAllowReserved() bool`

GetAllowReserved returns the AllowReserved field if non-nil, zero value otherwise.

### GetAllowReservedOk

`func (o *Parameter) GetAllowReservedOk() (*bool, bool)`

GetAllowReservedOk returns a tuple with the AllowReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReserved

`func (o *Parameter) SetAllowReserved(v bool)`

SetAllowReserved sets AllowReserved field to given value.

### HasAllowReserved

`func (o *Parameter) HasAllowReserved() bool`

HasAllowReserved returns a boolean if a field has been set.

### GetContent

`func (o *Parameter) GetContent() map[string]MediaType`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Parameter) GetContentOk() (*map[string]MediaType, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Parameter) SetContent(v map[string]MediaType)`

SetContent sets Content field to given value.

### HasContent

`func (o *Parameter) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDeprecated

`func (o *Parameter) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Parameter) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Parameter) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Parameter) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *Parameter) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *Parameter) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *Parameter) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *Parameter) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetExamples

`func (o *Parameter) GetExamples() map[string]Example`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *Parameter) GetExamplesOk() (*map[string]Example, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *Parameter) SetExamples(v map[string]Example)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *Parameter) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetExplode

`func (o *Parameter) GetExplode() bool`

GetExplode returns the Explode field if non-nil, zero value otherwise.

### GetExplodeOk

`func (o *Parameter) GetExplodeOk() (*bool, bool)`

GetExplodeOk returns a tuple with the Explode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplode

`func (o *Parameter) SetExplode(v bool)`

SetExplode sets Explode field to given value.

### HasExplode

`func (o *Parameter) HasExplode() bool`

HasExplode returns a boolean if a field has been set.

### GetExtensions

`func (o *Parameter) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Parameter) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Parameter) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Parameter) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGetref

`func (o *Parameter) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *Parameter) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *Parameter) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *Parameter) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetIn

`func (o *Parameter) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *Parameter) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *Parameter) SetIn(v string)`

SetIn sets In field to given value.

### HasIn

`func (o *Parameter) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetName

`func (o *Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Parameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Parameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *Parameter) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Parameter) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Parameter) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Parameter) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchema

`func (o *Parameter) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Parameter) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Parameter) SetSchema(v Schema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Parameter) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetStyle

`func (o *Parameter) GetStyle() StyleEnum`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *Parameter) GetStyleOk() (*StyleEnum, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *Parameter) SetStyle(v StyleEnum)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *Parameter) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


