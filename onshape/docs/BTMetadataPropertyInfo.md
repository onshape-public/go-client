# BTMetadataPropertyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedAssemblyProperty** | Pointer to **bool** |  | [optional] 
**ComputedProperty** | Pointer to **bool** |  | [optional] 
**ComputedPropertyError** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**Dirty** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EditableInUi** | Pointer to **bool** |  | [optional] 
**EnumValues** | Pointer to [**[]BTMetadataEnumValueInfo**](BTMetadataEnumValueInfo.md) |  | [optional] 
**InitialValue** | Pointer to **map[string]interface{}** |  | [optional] 
**Multivalued** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**PropertySource** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**UiHints** | Pointer to [**BTMetadataPropertyUiHintsInfo**](BTMetadataPropertyUiHintsInfo.md) |  | [optional] 
**Validator** | Pointer to [**BTMetadataPropertyValidatorInfo**](BTMetadataPropertyValidatorInfo.md) |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMetadataPropertyInfo

`func NewBTMetadataPropertyInfo() *BTMetadataPropertyInfo`

NewBTMetadataPropertyInfo instantiates a new BTMetadataPropertyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataPropertyInfoWithDefaults

`func NewBTMetadataPropertyInfoWithDefaults() *BTMetadataPropertyInfo`

NewBTMetadataPropertyInfoWithDefaults instantiates a new BTMetadataPropertyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedAssemblyProperty

`func (o *BTMetadataPropertyInfo) GetComputedAssemblyProperty() bool`

GetComputedAssemblyProperty returns the ComputedAssemblyProperty field if non-nil, zero value otherwise.

### GetComputedAssemblyPropertyOk

`func (o *BTMetadataPropertyInfo) GetComputedAssemblyPropertyOk() (*bool, bool)`

GetComputedAssemblyPropertyOk returns a tuple with the ComputedAssemblyProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedAssemblyProperty

`func (o *BTMetadataPropertyInfo) SetComputedAssemblyProperty(v bool)`

SetComputedAssemblyProperty sets ComputedAssemblyProperty field to given value.

### HasComputedAssemblyProperty

`func (o *BTMetadataPropertyInfo) HasComputedAssemblyProperty() bool`

HasComputedAssemblyProperty returns a boolean if a field has been set.

### GetComputedProperty

`func (o *BTMetadataPropertyInfo) GetComputedProperty() bool`

GetComputedProperty returns the ComputedProperty field if non-nil, zero value otherwise.

### GetComputedPropertyOk

`func (o *BTMetadataPropertyInfo) GetComputedPropertyOk() (*bool, bool)`

GetComputedPropertyOk returns a tuple with the ComputedProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedProperty

`func (o *BTMetadataPropertyInfo) SetComputedProperty(v bool)`

SetComputedProperty sets ComputedProperty field to given value.

### HasComputedProperty

`func (o *BTMetadataPropertyInfo) HasComputedProperty() bool`

HasComputedProperty returns a boolean if a field has been set.

### GetComputedPropertyError

`func (o *BTMetadataPropertyInfo) GetComputedPropertyError() string`

GetComputedPropertyError returns the ComputedPropertyError field if non-nil, zero value otherwise.

### GetComputedPropertyErrorOk

`func (o *BTMetadataPropertyInfo) GetComputedPropertyErrorOk() (*string, bool)`

GetComputedPropertyErrorOk returns a tuple with the ComputedPropertyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPropertyError

`func (o *BTMetadataPropertyInfo) SetComputedPropertyError(v string)`

SetComputedPropertyError sets ComputedPropertyError field to given value.

### HasComputedPropertyError

`func (o *BTMetadataPropertyInfo) HasComputedPropertyError() bool`

HasComputedPropertyError returns a boolean if a field has been set.

### GetDefaultValue

`func (o *BTMetadataPropertyInfo) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTMetadataPropertyInfo) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTMetadataPropertyInfo) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTMetadataPropertyInfo) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *BTMetadataPropertyInfo) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *BTMetadataPropertyInfo) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDirty

`func (o *BTMetadataPropertyInfo) GetDirty() bool`

GetDirty returns the Dirty field if non-nil, zero value otherwise.

### GetDirtyOk

`func (o *BTMetadataPropertyInfo) GetDirtyOk() (*bool, bool)`

GetDirtyOk returns a tuple with the Dirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirty

`func (o *BTMetadataPropertyInfo) SetDirty(v bool)`

SetDirty sets Dirty field to given value.

### HasDirty

`func (o *BTMetadataPropertyInfo) HasDirty() bool`

HasDirty returns a boolean if a field has been set.

### GetEditable

`func (o *BTMetadataPropertyInfo) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *BTMetadataPropertyInfo) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *BTMetadataPropertyInfo) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *BTMetadataPropertyInfo) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEditableInUi

`func (o *BTMetadataPropertyInfo) GetEditableInUi() bool`

GetEditableInUi returns the EditableInUi field if non-nil, zero value otherwise.

### GetEditableInUiOk

`func (o *BTMetadataPropertyInfo) GetEditableInUiOk() (*bool, bool)`

GetEditableInUiOk returns a tuple with the EditableInUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInUi

`func (o *BTMetadataPropertyInfo) SetEditableInUi(v bool)`

SetEditableInUi sets EditableInUi field to given value.

### HasEditableInUi

`func (o *BTMetadataPropertyInfo) HasEditableInUi() bool`

HasEditableInUi returns a boolean if a field has been set.

### GetEnumValues

`func (o *BTMetadataPropertyInfo) GetEnumValues() []BTMetadataEnumValueInfo`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *BTMetadataPropertyInfo) GetEnumValuesOk() (*[]BTMetadataEnumValueInfo, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *BTMetadataPropertyInfo) SetEnumValues(v []BTMetadataEnumValueInfo)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *BTMetadataPropertyInfo) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetInitialValue

`func (o *BTMetadataPropertyInfo) GetInitialValue() map[string]interface{}`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *BTMetadataPropertyInfo) GetInitialValueOk() (*map[string]interface{}, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *BTMetadataPropertyInfo) SetInitialValue(v map[string]interface{})`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *BTMetadataPropertyInfo) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetMultivalued

`func (o *BTMetadataPropertyInfo) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *BTMetadataPropertyInfo) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *BTMetadataPropertyInfo) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *BTMetadataPropertyInfo) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataPropertyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataPropertyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataPropertyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataPropertyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTMetadataPropertyInfo) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTMetadataPropertyInfo) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTMetadataPropertyInfo) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTMetadataPropertyInfo) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertySource

`func (o *BTMetadataPropertyInfo) GetPropertySource() int32`

GetPropertySource returns the PropertySource field if non-nil, zero value otherwise.

### GetPropertySourceOk

`func (o *BTMetadataPropertyInfo) GetPropertySourceOk() (*int32, bool)`

GetPropertySourceOk returns a tuple with the PropertySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySource

`func (o *BTMetadataPropertyInfo) SetPropertySource(v int32)`

SetPropertySource sets PropertySource field to given value.

### HasPropertySource

`func (o *BTMetadataPropertyInfo) HasPropertySource() bool`

HasPropertySource returns a boolean if a field has been set.

### GetRequired

`func (o *BTMetadataPropertyInfo) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BTMetadataPropertyInfo) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BTMetadataPropertyInfo) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BTMetadataPropertyInfo) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchemaId

`func (o *BTMetadataPropertyInfo) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *BTMetadataPropertyInfo) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *BTMetadataPropertyInfo) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *BTMetadataPropertyInfo) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetUiHints

`func (o *BTMetadataPropertyInfo) GetUiHints() BTMetadataPropertyUiHintsInfo`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTMetadataPropertyInfo) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTMetadataPropertyInfo) SetUiHints(v BTMetadataPropertyUiHintsInfo)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTMetadataPropertyInfo) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.

### GetValidator

`func (o *BTMetadataPropertyInfo) GetValidator() BTMetadataPropertyValidatorInfo`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BTMetadataPropertyInfo) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BTMetadataPropertyInfo) SetValidator(v BTMetadataPropertyValidatorInfo)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BTMetadataPropertyInfo) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetValue

`func (o *BTMetadataPropertyInfo) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTMetadataPropertyInfo) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTMetadataPropertyInfo) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BTMetadataPropertyInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BTMetadataPropertyInfo) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BTMetadataPropertyInfo) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueType

`func (o *BTMetadataPropertyInfo) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTMetadataPropertyInfo) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTMetadataPropertyInfo) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BTMetadataPropertyInfo) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


