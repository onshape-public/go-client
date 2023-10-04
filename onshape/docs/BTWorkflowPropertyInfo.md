# BTWorkflowPropertyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedAssemblyProperty** | Pointer to **bool** |  | [optional] 
**ComputedProperty** | Pointer to **bool** |  | [optional] 
**ComputedPropertyError** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **map[string]interface{}** |  | [optional] 
**Dirty** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EditableInUi** | Pointer to **bool** |  | [optional] 
**EnumValues** | Pointer to [**[]BTMetadataEnumValueInfo**](BTMetadataEnumValueInfo.md) |  | [optional] 
**HideInUi** | Pointer to **bool** |  | [optional] 
**InitialValue** | Pointer to **map[string]interface{}** |  | [optional] 
**IsApproverProperty** | Pointer to **bool** |  | [optional] 
**IsNotifierProperty** | Pointer to **bool** |  | [optional] 
**Multivalued** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Observers** | Pointer to [**[]BTWorkflowableObjectObserver**](BTWorkflowableObjectObserver.md) |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**PropertySource** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**TeamsOnly** | Pointer to **bool** |  | [optional] 
**UiHints** | Pointer to [**BTMetadataPropertyUiHintsInfo**](BTMetadataPropertyUiHintsInfo.md) |  | [optional] 
**UsersOnly** | Pointer to **bool** |  | [optional] 
**Validator** | Pointer to [**BTMetadataPropertyValidatorInfo**](BTMetadataPropertyValidatorInfo.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkflowPropertyInfo

`func NewBTWorkflowPropertyInfo() *BTWorkflowPropertyInfo`

NewBTWorkflowPropertyInfo instantiates a new BTWorkflowPropertyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowPropertyInfoWithDefaults

`func NewBTWorkflowPropertyInfoWithDefaults() *BTWorkflowPropertyInfo`

NewBTWorkflowPropertyInfoWithDefaults instantiates a new BTWorkflowPropertyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedAssemblyProperty

`func (o *BTWorkflowPropertyInfo) GetComputedAssemblyProperty() bool`

GetComputedAssemblyProperty returns the ComputedAssemblyProperty field if non-nil, zero value otherwise.

### GetComputedAssemblyPropertyOk

`func (o *BTWorkflowPropertyInfo) GetComputedAssemblyPropertyOk() (*bool, bool)`

GetComputedAssemblyPropertyOk returns a tuple with the ComputedAssemblyProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedAssemblyProperty

`func (o *BTWorkflowPropertyInfo) SetComputedAssemblyProperty(v bool)`

SetComputedAssemblyProperty sets ComputedAssemblyProperty field to given value.

### HasComputedAssemblyProperty

`func (o *BTWorkflowPropertyInfo) HasComputedAssemblyProperty() bool`

HasComputedAssemblyProperty returns a boolean if a field has been set.

### GetComputedProperty

`func (o *BTWorkflowPropertyInfo) GetComputedProperty() bool`

GetComputedProperty returns the ComputedProperty field if non-nil, zero value otherwise.

### GetComputedPropertyOk

`func (o *BTWorkflowPropertyInfo) GetComputedPropertyOk() (*bool, bool)`

GetComputedPropertyOk returns a tuple with the ComputedProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedProperty

`func (o *BTWorkflowPropertyInfo) SetComputedProperty(v bool)`

SetComputedProperty sets ComputedProperty field to given value.

### HasComputedProperty

`func (o *BTWorkflowPropertyInfo) HasComputedProperty() bool`

HasComputedProperty returns a boolean if a field has been set.

### GetComputedPropertyError

`func (o *BTWorkflowPropertyInfo) GetComputedPropertyError() string`

GetComputedPropertyError returns the ComputedPropertyError field if non-nil, zero value otherwise.

### GetComputedPropertyErrorOk

`func (o *BTWorkflowPropertyInfo) GetComputedPropertyErrorOk() (*string, bool)`

GetComputedPropertyErrorOk returns a tuple with the ComputedPropertyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPropertyError

`func (o *BTWorkflowPropertyInfo) SetComputedPropertyError(v string)`

SetComputedPropertyError sets ComputedPropertyError field to given value.

### HasComputedPropertyError

`func (o *BTWorkflowPropertyInfo) HasComputedPropertyError() bool`

HasComputedPropertyError returns a boolean if a field has been set.

### GetDefaultValue

`func (o *BTWorkflowPropertyInfo) GetDefaultValue() map[string]interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTWorkflowPropertyInfo) GetDefaultValueOk() (*map[string]interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTWorkflowPropertyInfo) SetDefaultValue(v map[string]interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTWorkflowPropertyInfo) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDirty

`func (o *BTWorkflowPropertyInfo) GetDirty() bool`

GetDirty returns the Dirty field if non-nil, zero value otherwise.

### GetDirtyOk

`func (o *BTWorkflowPropertyInfo) GetDirtyOk() (*bool, bool)`

GetDirtyOk returns a tuple with the Dirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirty

`func (o *BTWorkflowPropertyInfo) SetDirty(v bool)`

SetDirty sets Dirty field to given value.

### HasDirty

`func (o *BTWorkflowPropertyInfo) HasDirty() bool`

HasDirty returns a boolean if a field has been set.

### GetEditable

`func (o *BTWorkflowPropertyInfo) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *BTWorkflowPropertyInfo) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *BTWorkflowPropertyInfo) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *BTWorkflowPropertyInfo) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEditableInUi

`func (o *BTWorkflowPropertyInfo) GetEditableInUi() bool`

GetEditableInUi returns the EditableInUi field if non-nil, zero value otherwise.

### GetEditableInUiOk

`func (o *BTWorkflowPropertyInfo) GetEditableInUiOk() (*bool, bool)`

GetEditableInUiOk returns a tuple with the EditableInUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInUi

`func (o *BTWorkflowPropertyInfo) SetEditableInUi(v bool)`

SetEditableInUi sets EditableInUi field to given value.

### HasEditableInUi

`func (o *BTWorkflowPropertyInfo) HasEditableInUi() bool`

HasEditableInUi returns a boolean if a field has been set.

### GetEnumValues

`func (o *BTWorkflowPropertyInfo) GetEnumValues() []BTMetadataEnumValueInfo`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *BTWorkflowPropertyInfo) GetEnumValuesOk() (*[]BTMetadataEnumValueInfo, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *BTWorkflowPropertyInfo) SetEnumValues(v []BTMetadataEnumValueInfo)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *BTWorkflowPropertyInfo) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetHideInUi

`func (o *BTWorkflowPropertyInfo) GetHideInUi() bool`

GetHideInUi returns the HideInUi field if non-nil, zero value otherwise.

### GetHideInUiOk

`func (o *BTWorkflowPropertyInfo) GetHideInUiOk() (*bool, bool)`

GetHideInUiOk returns a tuple with the HideInUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideInUi

`func (o *BTWorkflowPropertyInfo) SetHideInUi(v bool)`

SetHideInUi sets HideInUi field to given value.

### HasHideInUi

`func (o *BTWorkflowPropertyInfo) HasHideInUi() bool`

HasHideInUi returns a boolean if a field has been set.

### GetInitialValue

`func (o *BTWorkflowPropertyInfo) GetInitialValue() map[string]interface{}`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *BTWorkflowPropertyInfo) GetInitialValueOk() (*map[string]interface{}, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *BTWorkflowPropertyInfo) SetInitialValue(v map[string]interface{})`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *BTWorkflowPropertyInfo) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetIsApproverProperty

`func (o *BTWorkflowPropertyInfo) GetIsApproverProperty() bool`

GetIsApproverProperty returns the IsApproverProperty field if non-nil, zero value otherwise.

### GetIsApproverPropertyOk

`func (o *BTWorkflowPropertyInfo) GetIsApproverPropertyOk() (*bool, bool)`

GetIsApproverPropertyOk returns a tuple with the IsApproverProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproverProperty

`func (o *BTWorkflowPropertyInfo) SetIsApproverProperty(v bool)`

SetIsApproverProperty sets IsApproverProperty field to given value.

### HasIsApproverProperty

`func (o *BTWorkflowPropertyInfo) HasIsApproverProperty() bool`

HasIsApproverProperty returns a boolean if a field has been set.

### GetIsNotifierProperty

`func (o *BTWorkflowPropertyInfo) GetIsNotifierProperty() bool`

GetIsNotifierProperty returns the IsNotifierProperty field if non-nil, zero value otherwise.

### GetIsNotifierPropertyOk

`func (o *BTWorkflowPropertyInfo) GetIsNotifierPropertyOk() (*bool, bool)`

GetIsNotifierPropertyOk returns a tuple with the IsNotifierProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotifierProperty

`func (o *BTWorkflowPropertyInfo) SetIsNotifierProperty(v bool)`

SetIsNotifierProperty sets IsNotifierProperty field to given value.

### HasIsNotifierProperty

`func (o *BTWorkflowPropertyInfo) HasIsNotifierProperty() bool`

HasIsNotifierProperty returns a boolean if a field has been set.

### GetMultivalued

`func (o *BTWorkflowPropertyInfo) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *BTWorkflowPropertyInfo) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *BTWorkflowPropertyInfo) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *BTWorkflowPropertyInfo) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.

### GetName

`func (o *BTWorkflowPropertyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWorkflowPropertyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWorkflowPropertyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWorkflowPropertyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObservers

`func (o *BTWorkflowPropertyInfo) GetObservers() []BTWorkflowableObjectObserver`

GetObservers returns the Observers field if non-nil, zero value otherwise.

### GetObserversOk

`func (o *BTWorkflowPropertyInfo) GetObserversOk() (*[]BTWorkflowableObjectObserver, bool)`

GetObserversOk returns a tuple with the Observers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservers

`func (o *BTWorkflowPropertyInfo) SetObservers(v []BTWorkflowableObjectObserver)`

SetObservers sets Observers field to given value.

### HasObservers

`func (o *BTWorkflowPropertyInfo) HasObservers() bool`

HasObservers returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTWorkflowPropertyInfo) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTWorkflowPropertyInfo) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTWorkflowPropertyInfo) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTWorkflowPropertyInfo) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertySource

`func (o *BTWorkflowPropertyInfo) GetPropertySource() int32`

GetPropertySource returns the PropertySource field if non-nil, zero value otherwise.

### GetPropertySourceOk

`func (o *BTWorkflowPropertyInfo) GetPropertySourceOk() (*int32, bool)`

GetPropertySourceOk returns a tuple with the PropertySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySource

`func (o *BTWorkflowPropertyInfo) SetPropertySource(v int32)`

SetPropertySource sets PropertySource field to given value.

### HasPropertySource

`func (o *BTWorkflowPropertyInfo) HasPropertySource() bool`

HasPropertySource returns a boolean if a field has been set.

### GetRequired

`func (o *BTWorkflowPropertyInfo) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BTWorkflowPropertyInfo) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BTWorkflowPropertyInfo) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BTWorkflowPropertyInfo) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchemaId

`func (o *BTWorkflowPropertyInfo) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *BTWorkflowPropertyInfo) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *BTWorkflowPropertyInfo) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *BTWorkflowPropertyInfo) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTeamsOnly

`func (o *BTWorkflowPropertyInfo) GetTeamsOnly() bool`

GetTeamsOnly returns the TeamsOnly field if non-nil, zero value otherwise.

### GetTeamsOnlyOk

`func (o *BTWorkflowPropertyInfo) GetTeamsOnlyOk() (*bool, bool)`

GetTeamsOnlyOk returns a tuple with the TeamsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsOnly

`func (o *BTWorkflowPropertyInfo) SetTeamsOnly(v bool)`

SetTeamsOnly sets TeamsOnly field to given value.

### HasTeamsOnly

`func (o *BTWorkflowPropertyInfo) HasTeamsOnly() bool`

HasTeamsOnly returns a boolean if a field has been set.

### GetUiHints

`func (o *BTWorkflowPropertyInfo) GetUiHints() BTMetadataPropertyUiHintsInfo`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTWorkflowPropertyInfo) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTWorkflowPropertyInfo) SetUiHints(v BTMetadataPropertyUiHintsInfo)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTWorkflowPropertyInfo) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.

### GetUsersOnly

`func (o *BTWorkflowPropertyInfo) GetUsersOnly() bool`

GetUsersOnly returns the UsersOnly field if non-nil, zero value otherwise.

### GetUsersOnlyOk

`func (o *BTWorkflowPropertyInfo) GetUsersOnlyOk() (*bool, bool)`

GetUsersOnlyOk returns a tuple with the UsersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersOnly

`func (o *BTWorkflowPropertyInfo) SetUsersOnly(v bool)`

SetUsersOnly sets UsersOnly field to given value.

### HasUsersOnly

`func (o *BTWorkflowPropertyInfo) HasUsersOnly() bool`

HasUsersOnly returns a boolean if a field has been set.

### GetValidator

`func (o *BTWorkflowPropertyInfo) GetValidator() BTMetadataPropertyValidatorInfo`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BTWorkflowPropertyInfo) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BTWorkflowPropertyInfo) SetValidator(v BTMetadataPropertyValidatorInfo)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BTWorkflowPropertyInfo) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetValue

`func (o *BTWorkflowPropertyInfo) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTWorkflowPropertyInfo) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTWorkflowPropertyInfo) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BTWorkflowPropertyInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *BTWorkflowPropertyInfo) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTWorkflowPropertyInfo) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTWorkflowPropertyInfo) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BTWorkflowPropertyInfo) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


