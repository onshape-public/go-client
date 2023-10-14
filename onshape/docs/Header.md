# Header

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**Examples** | Pointer to [**map[string]Example**](Example.md) |  | [optional] 
**Explode** | Pointer to **bool** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Schema** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Style** | Pointer to [**StyleEnum**](StyleEnum.md) |  | [optional] 

## Methods

### NewHeader

`func NewHeader() *Header`

NewHeader instantiates a new Header object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderWithDefaults

`func NewHeaderWithDefaults() *Header`

NewHeaderWithDefaults instantiates a new Header object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecated

`func (o *Header) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Header) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Header) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Header) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *Header) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Header) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Header) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Header) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *Header) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *Header) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *Header) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *Header) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetExamples

`func (o *Header) GetExamples() map[string]Example`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *Header) GetExamplesOk() (*map[string]Example, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *Header) SetExamples(v map[string]Example)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *Header) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetExplode

`func (o *Header) GetExplode() bool`

GetExplode returns the Explode field if non-nil, zero value otherwise.

### GetExplodeOk

`func (o *Header) GetExplodeOk() (*bool, bool)`

GetExplodeOk returns a tuple with the Explode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplode

`func (o *Header) SetExplode(v bool)`

SetExplode sets Explode field to given value.

### HasExplode

`func (o *Header) HasExplode() bool`

HasExplode returns a boolean if a field has been set.

### GetExtensions

`func (o *Header) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Header) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Header) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Header) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGetref

`func (o *Header) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *Header) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *Header) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *Header) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetRequired

`func (o *Header) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Header) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Header) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Header) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchema

`func (o *Header) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Header) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Header) SetSchema(v Schema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Header) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetStyle

`func (o *Header) GetStyle() StyleEnum`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *Header) GetStyleOk() (*StyleEnum, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *Header) SetStyle(v StyleEnum)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *Header) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


