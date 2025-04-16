# MediaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encoding** | Pointer to [**map[string]Encoding**](Encoding.md) |  | [optional] 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**ExampleSetFlag** | Pointer to **bool** |  | [optional] 
**Examples** | Pointer to [**map[string]Example**](Example.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Schema** | Pointer to [**Schema**](Schema.md) |  | [optional] 

## Methods

### NewMediaType

`func NewMediaType() *MediaType`

NewMediaType instantiates a new MediaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaTypeWithDefaults

`func NewMediaTypeWithDefaults() *MediaType`

NewMediaTypeWithDefaults instantiates a new MediaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncoding

`func (o *MediaType) GetEncoding() map[string]Encoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *MediaType) GetEncodingOk() (*map[string]Encoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *MediaType) SetEncoding(v map[string]Encoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *MediaType) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetExample

`func (o *MediaType) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *MediaType) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *MediaType) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *MediaType) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetExampleSetFlag

`func (o *MediaType) GetExampleSetFlag() bool`

GetExampleSetFlag returns the ExampleSetFlag field if non-nil, zero value otherwise.

### GetExampleSetFlagOk

`func (o *MediaType) GetExampleSetFlagOk() (*bool, bool)`

GetExampleSetFlagOk returns a tuple with the ExampleSetFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleSetFlag

`func (o *MediaType) SetExampleSetFlag(v bool)`

SetExampleSetFlag sets ExampleSetFlag field to given value.

### HasExampleSetFlag

`func (o *MediaType) HasExampleSetFlag() bool`

HasExampleSetFlag returns a boolean if a field has been set.

### GetExamples

`func (o *MediaType) GetExamples() map[string]Example`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *MediaType) GetExamplesOk() (*map[string]Example, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *MediaType) SetExamples(v map[string]Example)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *MediaType) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetExtensions

`func (o *MediaType) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MediaType) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MediaType) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MediaType) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetSchema

`func (o *MediaType) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *MediaType) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *MediaType) SetSchema(v Schema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *MediaType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


