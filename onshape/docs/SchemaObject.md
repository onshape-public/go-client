# SchemaObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalItems** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**AdditionalPropertiesField** | Pointer to **map[string]interface{}** |  | [optional] 
**AllOf** | Pointer to [**[]Schema**](Schema.md) |  | [optional] 
**AnyOf** | Pointer to [**[]Schema**](Schema.md) |  | [optional] 
**BooleanSchemaValue** | Pointer to **bool** |  | [optional] 
**Const** | Pointer to **map[string]interface{}** |  | [optional] 
**Contains** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**ContentEncoding** | Pointer to **string** |  | [optional] 
**ContentMediaType** | Pointer to **string** |  | [optional] 
**ContentSchema** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Default** | Pointer to **map[string]interface{}** |  | [optional] 
**DependentRequired** | Pointer to **map[string][]string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Discriminator** | Pointer to [**Discriminator**](Discriminator.md) |  | [optional] 
**Else** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Enum** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**ExampleSetFlag** | Pointer to **bool** |  | [optional] 
**Examples** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ExclusiveMaximum** | Pointer to **bool** |  | [optional] 
**ExclusiveMaximumValue** | Pointer to **float32** |  | [optional] 
**ExclusiveMinimum** | Pointer to **bool** |  | [optional] 
**ExclusiveMinimumValue** | Pointer to **float32** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Getanchor** | Pointer to **string** |  | [optional] 
**Getcomment** | Pointer to **string** |  | [optional] 
**Getid** | Pointer to **string** |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**Getschema** | Pointer to **string** |  | [optional] 
**If** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Items** | Pointer to [**SchemaObject**](SchemaObject.md) |  | [optional] 
**JsonSchema** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**JsonSchemaImpl** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxContains** | Pointer to **int32** |  | [optional] 
**MaxItems** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**MaxProperties** | Pointer to **int32** |  | [optional] 
**Maximum** | Pointer to **float32** |  | [optional] 
**MinContains** | Pointer to **int32** |  | [optional] 
**MinItems** | Pointer to **int32** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MinProperties** | Pointer to **int32** |  | [optional] 
**Minimum** | Pointer to **float32** |  | [optional] 
**MultipleOf** | Pointer to **float32** |  | [optional] 
**Not** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Nullable** | Pointer to **bool** |  | [optional] 
**OneOf** | Pointer to [**[]Schema**](Schema.md) |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**PrefixItems** | Pointer to [**[]Schema**](Schema.md) |  | [optional] 
**PropertyNames** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Then** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**UnevaluatedItems** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**UnevaluatedProperties** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**UniqueItems** | Pointer to **bool** |  | [optional] 
**WriteOnly** | Pointer to **bool** |  | [optional] 
**Xml** | Pointer to [**XML**](XML.md) |  | [optional] 

## Methods

### NewSchemaObject

`func NewSchemaObject() *SchemaObject`

NewSchemaObject instantiates a new SchemaObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaObjectWithDefaults

`func NewSchemaObjectWithDefaults() *SchemaObject`

NewSchemaObjectWithDefaults instantiates a new SchemaObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalItems

`func (o *SchemaObject) GetAdditionalItems() Schema`

GetAdditionalItems returns the AdditionalItems field if non-nil, zero value otherwise.

### GetAdditionalItemsOk

`func (o *SchemaObject) GetAdditionalItemsOk() (*Schema, bool)`

GetAdditionalItemsOk returns a tuple with the AdditionalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalItems

`func (o *SchemaObject) SetAdditionalItems(v Schema)`

SetAdditionalItems sets AdditionalItems field to given value.

### HasAdditionalItems

`func (o *SchemaObject) HasAdditionalItems() bool`

HasAdditionalItems returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *SchemaObject) GetAdditionalPropertiesField() map[string]interface{}`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *SchemaObject) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *SchemaObject) SetAdditionalPropertiesField(v map[string]interface{})`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *SchemaObject) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetAllOf

`func (o *SchemaObject) GetAllOf() []Schema`

GetAllOf returns the AllOf field if non-nil, zero value otherwise.

### GetAllOfOk

`func (o *SchemaObject) GetAllOfOk() (*[]Schema, bool)`

GetAllOfOk returns a tuple with the AllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOf

`func (o *SchemaObject) SetAllOf(v []Schema)`

SetAllOf sets AllOf field to given value.

### HasAllOf

`func (o *SchemaObject) HasAllOf() bool`

HasAllOf returns a boolean if a field has been set.

### GetAnyOf

`func (o *SchemaObject) GetAnyOf() []Schema`

GetAnyOf returns the AnyOf field if non-nil, zero value otherwise.

### GetAnyOfOk

`func (o *SchemaObject) GetAnyOfOk() (*[]Schema, bool)`

GetAnyOfOk returns a tuple with the AnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyOf

`func (o *SchemaObject) SetAnyOf(v []Schema)`

SetAnyOf sets AnyOf field to given value.

### HasAnyOf

`func (o *SchemaObject) HasAnyOf() bool`

HasAnyOf returns a boolean if a field has been set.

### GetBooleanSchemaValue

`func (o *SchemaObject) GetBooleanSchemaValue() bool`

GetBooleanSchemaValue returns the BooleanSchemaValue field if non-nil, zero value otherwise.

### GetBooleanSchemaValueOk

`func (o *SchemaObject) GetBooleanSchemaValueOk() (*bool, bool)`

GetBooleanSchemaValueOk returns a tuple with the BooleanSchemaValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanSchemaValue

`func (o *SchemaObject) SetBooleanSchemaValue(v bool)`

SetBooleanSchemaValue sets BooleanSchemaValue field to given value.

### HasBooleanSchemaValue

`func (o *SchemaObject) HasBooleanSchemaValue() bool`

HasBooleanSchemaValue returns a boolean if a field has been set.

### GetConst

`func (o *SchemaObject) GetConst() map[string]interface{}`

GetConst returns the Const field if non-nil, zero value otherwise.

### GetConstOk

`func (o *SchemaObject) GetConstOk() (*map[string]interface{}, bool)`

GetConstOk returns a tuple with the Const field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConst

`func (o *SchemaObject) SetConst(v map[string]interface{})`

SetConst sets Const field to given value.

### HasConst

`func (o *SchemaObject) HasConst() bool`

HasConst returns a boolean if a field has been set.

### GetContains

`func (o *SchemaObject) GetContains() Schema`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *SchemaObject) GetContainsOk() (*Schema, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *SchemaObject) SetContains(v Schema)`

SetContains sets Contains field to given value.

### HasContains

`func (o *SchemaObject) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetContentEncoding

`func (o *SchemaObject) GetContentEncoding() string`

GetContentEncoding returns the ContentEncoding field if non-nil, zero value otherwise.

### GetContentEncodingOk

`func (o *SchemaObject) GetContentEncodingOk() (*string, bool)`

GetContentEncodingOk returns a tuple with the ContentEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEncoding

`func (o *SchemaObject) SetContentEncoding(v string)`

SetContentEncoding sets ContentEncoding field to given value.

### HasContentEncoding

`func (o *SchemaObject) HasContentEncoding() bool`

HasContentEncoding returns a boolean if a field has been set.

### GetContentMediaType

`func (o *SchemaObject) GetContentMediaType() string`

GetContentMediaType returns the ContentMediaType field if non-nil, zero value otherwise.

### GetContentMediaTypeOk

`func (o *SchemaObject) GetContentMediaTypeOk() (*string, bool)`

GetContentMediaTypeOk returns a tuple with the ContentMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMediaType

`func (o *SchemaObject) SetContentMediaType(v string)`

SetContentMediaType sets ContentMediaType field to given value.

### HasContentMediaType

`func (o *SchemaObject) HasContentMediaType() bool`

HasContentMediaType returns a boolean if a field has been set.

### GetContentSchema

`func (o *SchemaObject) GetContentSchema() Schema`

GetContentSchema returns the ContentSchema field if non-nil, zero value otherwise.

### GetContentSchemaOk

`func (o *SchemaObject) GetContentSchemaOk() (*Schema, bool)`

GetContentSchemaOk returns a tuple with the ContentSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSchema

`func (o *SchemaObject) SetContentSchema(v Schema)`

SetContentSchema sets ContentSchema field to given value.

### HasContentSchema

`func (o *SchemaObject) HasContentSchema() bool`

HasContentSchema returns a boolean if a field has been set.

### GetDefault

`func (o *SchemaObject) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SchemaObject) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SchemaObject) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SchemaObject) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDependentRequired

`func (o *SchemaObject) GetDependentRequired() map[string][]string`

GetDependentRequired returns the DependentRequired field if non-nil, zero value otherwise.

### GetDependentRequiredOk

`func (o *SchemaObject) GetDependentRequiredOk() (*map[string][]string, bool)`

GetDependentRequiredOk returns a tuple with the DependentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentRequired

`func (o *SchemaObject) SetDependentRequired(v map[string][]string)`

SetDependentRequired sets DependentRequired field to given value.

### HasDependentRequired

`func (o *SchemaObject) HasDependentRequired() bool`

HasDependentRequired returns a boolean if a field has been set.

### GetDeprecated

`func (o *SchemaObject) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *SchemaObject) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *SchemaObject) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *SchemaObject) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscriminator

`func (o *SchemaObject) GetDiscriminator() Discriminator`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *SchemaObject) GetDiscriminatorOk() (*Discriminator, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *SchemaObject) SetDiscriminator(v Discriminator)`

SetDiscriminator sets Discriminator field to given value.

### HasDiscriminator

`func (o *SchemaObject) HasDiscriminator() bool`

HasDiscriminator returns a boolean if a field has been set.

### GetElse

`func (o *SchemaObject) GetElse() Schema`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *SchemaObject) GetElseOk() (*Schema, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *SchemaObject) SetElse(v Schema)`

SetElse sets Else field to given value.

### HasElse

`func (o *SchemaObject) HasElse() bool`

HasElse returns a boolean if a field has been set.

### GetEnum

`func (o *SchemaObject) GetEnum() []map[string]interface{}`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *SchemaObject) GetEnumOk() (*[]map[string]interface{}, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *SchemaObject) SetEnum(v []map[string]interface{})`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *SchemaObject) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetExample

`func (o *SchemaObject) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *SchemaObject) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *SchemaObject) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *SchemaObject) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetExampleSetFlag

`func (o *SchemaObject) GetExampleSetFlag() bool`

GetExampleSetFlag returns the ExampleSetFlag field if non-nil, zero value otherwise.

### GetExampleSetFlagOk

`func (o *SchemaObject) GetExampleSetFlagOk() (*bool, bool)`

GetExampleSetFlagOk returns a tuple with the ExampleSetFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleSetFlag

`func (o *SchemaObject) SetExampleSetFlag(v bool)`

SetExampleSetFlag sets ExampleSetFlag field to given value.

### HasExampleSetFlag

`func (o *SchemaObject) HasExampleSetFlag() bool`

HasExampleSetFlag returns a boolean if a field has been set.

### GetExamples

`func (o *SchemaObject) GetExamples() []map[string]interface{}`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *SchemaObject) GetExamplesOk() (*[]map[string]interface{}, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *SchemaObject) SetExamples(v []map[string]interface{})`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *SchemaObject) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetExclusiveMaximum

`func (o *SchemaObject) GetExclusiveMaximum() bool`

GetExclusiveMaximum returns the ExclusiveMaximum field if non-nil, zero value otherwise.

### GetExclusiveMaximumOk

`func (o *SchemaObject) GetExclusiveMaximumOk() (*bool, bool)`

GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMaximum

`func (o *SchemaObject) SetExclusiveMaximum(v bool)`

SetExclusiveMaximum sets ExclusiveMaximum field to given value.

### HasExclusiveMaximum

`func (o *SchemaObject) HasExclusiveMaximum() bool`

HasExclusiveMaximum returns a boolean if a field has been set.

### GetExclusiveMaximumValue

`func (o *SchemaObject) GetExclusiveMaximumValue() float32`

GetExclusiveMaximumValue returns the ExclusiveMaximumValue field if non-nil, zero value otherwise.

### GetExclusiveMaximumValueOk

`func (o *SchemaObject) GetExclusiveMaximumValueOk() (*float32, bool)`

GetExclusiveMaximumValueOk returns a tuple with the ExclusiveMaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMaximumValue

`func (o *SchemaObject) SetExclusiveMaximumValue(v float32)`

SetExclusiveMaximumValue sets ExclusiveMaximumValue field to given value.

### HasExclusiveMaximumValue

`func (o *SchemaObject) HasExclusiveMaximumValue() bool`

HasExclusiveMaximumValue returns a boolean if a field has been set.

### GetExclusiveMinimum

`func (o *SchemaObject) GetExclusiveMinimum() bool`

GetExclusiveMinimum returns the ExclusiveMinimum field if non-nil, zero value otherwise.

### GetExclusiveMinimumOk

`func (o *SchemaObject) GetExclusiveMinimumOk() (*bool, bool)`

GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimum

`func (o *SchemaObject) SetExclusiveMinimum(v bool)`

SetExclusiveMinimum sets ExclusiveMinimum field to given value.

### HasExclusiveMinimum

`func (o *SchemaObject) HasExclusiveMinimum() bool`

HasExclusiveMinimum returns a boolean if a field has been set.

### GetExclusiveMinimumValue

`func (o *SchemaObject) GetExclusiveMinimumValue() float32`

GetExclusiveMinimumValue returns the ExclusiveMinimumValue field if non-nil, zero value otherwise.

### GetExclusiveMinimumValueOk

`func (o *SchemaObject) GetExclusiveMinimumValueOk() (*float32, bool)`

GetExclusiveMinimumValueOk returns a tuple with the ExclusiveMinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimumValue

`func (o *SchemaObject) SetExclusiveMinimumValue(v float32)`

SetExclusiveMinimumValue sets ExclusiveMinimumValue field to given value.

### HasExclusiveMinimumValue

`func (o *SchemaObject) HasExclusiveMinimumValue() bool`

HasExclusiveMinimumValue returns a boolean if a field has been set.

### GetExtensions

`func (o *SchemaObject) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SchemaObject) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SchemaObject) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SchemaObject) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExternalDocs

`func (o *SchemaObject) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *SchemaObject) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *SchemaObject) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *SchemaObject) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetFormat

`func (o *SchemaObject) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SchemaObject) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SchemaObject) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SchemaObject) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGetanchor

`func (o *SchemaObject) GetGetanchor() string`

GetGetanchor returns the Getanchor field if non-nil, zero value otherwise.

### GetGetanchorOk

`func (o *SchemaObject) GetGetanchorOk() (*string, bool)`

GetGetanchorOk returns a tuple with the Getanchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetanchor

`func (o *SchemaObject) SetGetanchor(v string)`

SetGetanchor sets Getanchor field to given value.

### HasGetanchor

`func (o *SchemaObject) HasGetanchor() bool`

HasGetanchor returns a boolean if a field has been set.

### GetGetcomment

`func (o *SchemaObject) GetGetcomment() string`

GetGetcomment returns the Getcomment field if non-nil, zero value otherwise.

### GetGetcommentOk

`func (o *SchemaObject) GetGetcommentOk() (*string, bool)`

GetGetcommentOk returns a tuple with the Getcomment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetcomment

`func (o *SchemaObject) SetGetcomment(v string)`

SetGetcomment sets Getcomment field to given value.

### HasGetcomment

`func (o *SchemaObject) HasGetcomment() bool`

HasGetcomment returns a boolean if a field has been set.

### GetGetid

`func (o *SchemaObject) GetGetid() string`

GetGetid returns the Getid field if non-nil, zero value otherwise.

### GetGetidOk

`func (o *SchemaObject) GetGetidOk() (*string, bool)`

GetGetidOk returns a tuple with the Getid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetid

`func (o *SchemaObject) SetGetid(v string)`

SetGetid sets Getid field to given value.

### HasGetid

`func (o *SchemaObject) HasGetid() bool`

HasGetid returns a boolean if a field has been set.

### GetGetref

`func (o *SchemaObject) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *SchemaObject) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *SchemaObject) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *SchemaObject) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetGetschema

`func (o *SchemaObject) GetGetschema() string`

GetGetschema returns the Getschema field if non-nil, zero value otherwise.

### GetGetschemaOk

`func (o *SchemaObject) GetGetschemaOk() (*string, bool)`

GetGetschemaOk returns a tuple with the Getschema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetschema

`func (o *SchemaObject) SetGetschema(v string)`

SetGetschema sets Getschema field to given value.

### HasGetschema

`func (o *SchemaObject) HasGetschema() bool`

HasGetschema returns a boolean if a field has been set.

### GetIf

`func (o *SchemaObject) GetIf() Schema`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *SchemaObject) GetIfOk() (*Schema, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *SchemaObject) SetIf(v Schema)`

SetIf sets If field to given value.

### HasIf

`func (o *SchemaObject) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetItems

`func (o *SchemaObject) GetItems() SchemaObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SchemaObject) GetItemsOk() (*SchemaObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SchemaObject) SetItems(v SchemaObject)`

SetItems sets Items field to given value.

### HasItems

`func (o *SchemaObject) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetJsonSchema

`func (o *SchemaObject) GetJsonSchema() map[string]map[string]interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *SchemaObject) GetJsonSchemaOk() (*map[string]map[string]interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *SchemaObject) SetJsonSchema(v map[string]map[string]interface{})`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *SchemaObject) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.

### GetJsonSchemaImpl

`func (o *SchemaObject) GetJsonSchemaImpl() map[string]interface{}`

GetJsonSchemaImpl returns the JsonSchemaImpl field if non-nil, zero value otherwise.

### GetJsonSchemaImplOk

`func (o *SchemaObject) GetJsonSchemaImplOk() (*map[string]interface{}, bool)`

GetJsonSchemaImplOk returns a tuple with the JsonSchemaImpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchemaImpl

`func (o *SchemaObject) SetJsonSchemaImpl(v map[string]interface{})`

SetJsonSchemaImpl sets JsonSchemaImpl field to given value.

### HasJsonSchemaImpl

`func (o *SchemaObject) HasJsonSchemaImpl() bool`

HasJsonSchemaImpl returns a boolean if a field has been set.

### GetMaxContains

`func (o *SchemaObject) GetMaxContains() int32`

GetMaxContains returns the MaxContains field if non-nil, zero value otherwise.

### GetMaxContainsOk

`func (o *SchemaObject) GetMaxContainsOk() (*int32, bool)`

GetMaxContainsOk returns a tuple with the MaxContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContains

`func (o *SchemaObject) SetMaxContains(v int32)`

SetMaxContains sets MaxContains field to given value.

### HasMaxContains

`func (o *SchemaObject) HasMaxContains() bool`

HasMaxContains returns a boolean if a field has been set.

### GetMaxItems

`func (o *SchemaObject) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *SchemaObject) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *SchemaObject) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *SchemaObject) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### GetMaxLength

`func (o *SchemaObject) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *SchemaObject) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *SchemaObject) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *SchemaObject) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMaxProperties

`func (o *SchemaObject) GetMaxProperties() int32`

GetMaxProperties returns the MaxProperties field if non-nil, zero value otherwise.

### GetMaxPropertiesOk

`func (o *SchemaObject) GetMaxPropertiesOk() (*int32, bool)`

GetMaxPropertiesOk returns a tuple with the MaxProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProperties

`func (o *SchemaObject) SetMaxProperties(v int32)`

SetMaxProperties sets MaxProperties field to given value.

### HasMaxProperties

`func (o *SchemaObject) HasMaxProperties() bool`

HasMaxProperties returns a boolean if a field has been set.

### GetMaximum

`func (o *SchemaObject) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *SchemaObject) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *SchemaObject) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *SchemaObject) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinContains

`func (o *SchemaObject) GetMinContains() int32`

GetMinContains returns the MinContains field if non-nil, zero value otherwise.

### GetMinContainsOk

`func (o *SchemaObject) GetMinContainsOk() (*int32, bool)`

GetMinContainsOk returns a tuple with the MinContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinContains

`func (o *SchemaObject) SetMinContains(v int32)`

SetMinContains sets MinContains field to given value.

### HasMinContains

`func (o *SchemaObject) HasMinContains() bool`

HasMinContains returns a boolean if a field has been set.

### GetMinItems

`func (o *SchemaObject) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *SchemaObject) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *SchemaObject) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *SchemaObject) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### GetMinLength

`func (o *SchemaObject) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SchemaObject) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SchemaObject) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SchemaObject) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMinProperties

`func (o *SchemaObject) GetMinProperties() int32`

GetMinProperties returns the MinProperties field if non-nil, zero value otherwise.

### GetMinPropertiesOk

`func (o *SchemaObject) GetMinPropertiesOk() (*int32, bool)`

GetMinPropertiesOk returns a tuple with the MinProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProperties

`func (o *SchemaObject) SetMinProperties(v int32)`

SetMinProperties sets MinProperties field to given value.

### HasMinProperties

`func (o *SchemaObject) HasMinProperties() bool`

HasMinProperties returns a boolean if a field has been set.

### GetMinimum

`func (o *SchemaObject) GetMinimum() float32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *SchemaObject) GetMinimumOk() (*float32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *SchemaObject) SetMinimum(v float32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *SchemaObject) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetMultipleOf

`func (o *SchemaObject) GetMultipleOf() float32`

GetMultipleOf returns the MultipleOf field if non-nil, zero value otherwise.

### GetMultipleOfOk

`func (o *SchemaObject) GetMultipleOfOk() (*float32, bool)`

GetMultipleOfOk returns a tuple with the MultipleOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleOf

`func (o *SchemaObject) SetMultipleOf(v float32)`

SetMultipleOf sets MultipleOf field to given value.

### HasMultipleOf

`func (o *SchemaObject) HasMultipleOf() bool`

HasMultipleOf returns a boolean if a field has been set.

### GetNot

`func (o *SchemaObject) GetNot() Schema`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *SchemaObject) GetNotOk() (*Schema, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *SchemaObject) SetNot(v Schema)`

SetNot sets Not field to given value.

### HasNot

`func (o *SchemaObject) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetNullable

`func (o *SchemaObject) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *SchemaObject) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *SchemaObject) SetNullable(v bool)`

SetNullable sets Nullable field to given value.

### HasNullable

`func (o *SchemaObject) HasNullable() bool`

HasNullable returns a boolean if a field has been set.

### GetOneOf

`func (o *SchemaObject) GetOneOf() []Schema`

GetOneOf returns the OneOf field if non-nil, zero value otherwise.

### GetOneOfOk

`func (o *SchemaObject) GetOneOfOk() (*[]Schema, bool)`

GetOneOfOk returns a tuple with the OneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOf

`func (o *SchemaObject) SetOneOf(v []Schema)`

SetOneOf sets OneOf field to given value.

### HasOneOf

`func (o *SchemaObject) HasOneOf() bool`

HasOneOf returns a boolean if a field has been set.

### GetPattern

`func (o *SchemaObject) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SchemaObject) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SchemaObject) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SchemaObject) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPrefixItems

`func (o *SchemaObject) GetPrefixItems() []Schema`

GetPrefixItems returns the PrefixItems field if non-nil, zero value otherwise.

### GetPrefixItemsOk

`func (o *SchemaObject) GetPrefixItemsOk() (*[]Schema, bool)`

GetPrefixItemsOk returns a tuple with the PrefixItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixItems

`func (o *SchemaObject) SetPrefixItems(v []Schema)`

SetPrefixItems sets PrefixItems field to given value.

### HasPrefixItems

`func (o *SchemaObject) HasPrefixItems() bool`

HasPrefixItems returns a boolean if a field has been set.

### GetPropertyNames

`func (o *SchemaObject) GetPropertyNames() Schema`

GetPropertyNames returns the PropertyNames field if non-nil, zero value otherwise.

### GetPropertyNamesOk

`func (o *SchemaObject) GetPropertyNamesOk() (*Schema, bool)`

GetPropertyNamesOk returns a tuple with the PropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNames

`func (o *SchemaObject) SetPropertyNames(v Schema)`

SetPropertyNames sets PropertyNames field to given value.

### HasPropertyNames

`func (o *SchemaObject) HasPropertyNames() bool`

HasPropertyNames returns a boolean if a field has been set.

### GetReadOnly

`func (o *SchemaObject) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *SchemaObject) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *SchemaObject) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *SchemaObject) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRequired

`func (o *SchemaObject) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SchemaObject) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SchemaObject) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SchemaObject) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetThen

`func (o *SchemaObject) GetThen() Schema`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *SchemaObject) GetThenOk() (*Schema, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *SchemaObject) SetThen(v Schema)`

SetThen sets Then field to given value.

### HasThen

`func (o *SchemaObject) HasThen() bool`

HasThen returns a boolean if a field has been set.

### GetTitle

`func (o *SchemaObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SchemaObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SchemaObject) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SchemaObject) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *SchemaObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SchemaObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypes

`func (o *SchemaObject) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SchemaObject) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SchemaObject) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SchemaObject) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUnevaluatedItems

`func (o *SchemaObject) GetUnevaluatedItems() Schema`

GetUnevaluatedItems returns the UnevaluatedItems field if non-nil, zero value otherwise.

### GetUnevaluatedItemsOk

`func (o *SchemaObject) GetUnevaluatedItemsOk() (*Schema, bool)`

GetUnevaluatedItemsOk returns a tuple with the UnevaluatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnevaluatedItems

`func (o *SchemaObject) SetUnevaluatedItems(v Schema)`

SetUnevaluatedItems sets UnevaluatedItems field to given value.

### HasUnevaluatedItems

`func (o *SchemaObject) HasUnevaluatedItems() bool`

HasUnevaluatedItems returns a boolean if a field has been set.

### GetUnevaluatedProperties

`func (o *SchemaObject) GetUnevaluatedProperties() Schema`

GetUnevaluatedProperties returns the UnevaluatedProperties field if non-nil, zero value otherwise.

### GetUnevaluatedPropertiesOk

`func (o *SchemaObject) GetUnevaluatedPropertiesOk() (*Schema, bool)`

GetUnevaluatedPropertiesOk returns a tuple with the UnevaluatedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnevaluatedProperties

`func (o *SchemaObject) SetUnevaluatedProperties(v Schema)`

SetUnevaluatedProperties sets UnevaluatedProperties field to given value.

### HasUnevaluatedProperties

`func (o *SchemaObject) HasUnevaluatedProperties() bool`

HasUnevaluatedProperties returns a boolean if a field has been set.

### GetUniqueItems

`func (o *SchemaObject) GetUniqueItems() bool`

GetUniqueItems returns the UniqueItems field if non-nil, zero value otherwise.

### GetUniqueItemsOk

`func (o *SchemaObject) GetUniqueItemsOk() (*bool, bool)`

GetUniqueItemsOk returns a tuple with the UniqueItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueItems

`func (o *SchemaObject) SetUniqueItems(v bool)`

SetUniqueItems sets UniqueItems field to given value.

### HasUniqueItems

`func (o *SchemaObject) HasUniqueItems() bool`

HasUniqueItems returns a boolean if a field has been set.

### GetWriteOnly

`func (o *SchemaObject) GetWriteOnly() bool`

GetWriteOnly returns the WriteOnly field if non-nil, zero value otherwise.

### GetWriteOnlyOk

`func (o *SchemaObject) GetWriteOnlyOk() (*bool, bool)`

GetWriteOnlyOk returns a tuple with the WriteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOnly

`func (o *SchemaObject) SetWriteOnly(v bool)`

SetWriteOnly sets WriteOnly field to given value.

### HasWriteOnly

`func (o *SchemaObject) HasWriteOnly() bool`

HasWriteOnly returns a boolean if a field has been set.

### GetXml

`func (o *SchemaObject) GetXml() XML`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *SchemaObject) GetXmlOk() (*XML, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *SchemaObject) SetXml(v XML)`

SetXml sets Xml field to given value.

### HasXml

`func (o *SchemaObject) HasXml() bool`

HasXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


