# BTMSketchEntity3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**CombinedSketchEntityType** | Pointer to [**CombinedSketchEntityType**](CombinedSketchEntityType.md) |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityIdAndReplaceInDependentFields** | Pointer to **string** |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Element microversion that is being imported. | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 

## Methods

### NewBTMSketchEntity3

`func NewBTMSketchEntity3() *BTMSketchEntity3`

NewBTMSketchEntity3 instantiates a new BTMSketchEntity3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMSketchEntity3WithDefaults

`func NewBTMSketchEntity3WithDefaults() *BTMSketchEntity3`

NewBTMSketchEntity3WithDefaults instantiates a new BTMSketchEntity3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMSketchEntity3) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMSketchEntity3) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMSketchEntity3) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMSketchEntity3) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCombinedSketchEntityType

`func (o *BTMSketchEntity3) GetCombinedSketchEntityType() CombinedSketchEntityType`

GetCombinedSketchEntityType returns the CombinedSketchEntityType field if non-nil, zero value otherwise.

### GetCombinedSketchEntityTypeOk

`func (o *BTMSketchEntity3) GetCombinedSketchEntityTypeOk() (*CombinedSketchEntityType, bool)`

GetCombinedSketchEntityTypeOk returns a tuple with the CombinedSketchEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedSketchEntityType

`func (o *BTMSketchEntity3) SetCombinedSketchEntityType(v CombinedSketchEntityType)`

SetCombinedSketchEntityType sets CombinedSketchEntityType field to given value.

### HasCombinedSketchEntityType

`func (o *BTMSketchEntity3) HasCombinedSketchEntityType() bool`

HasCombinedSketchEntityType returns a boolean if a field has been set.

### GetEntityId

`func (o *BTMSketchEntity3) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *BTMSketchEntity3) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *BTMSketchEntity3) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *BTMSketchEntity3) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityIdAndReplaceInDependentFields

`func (o *BTMSketchEntity3) GetEntityIdAndReplaceInDependentFields() string`

GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field if non-nil, zero value otherwise.

### GetEntityIdAndReplaceInDependentFieldsOk

`func (o *BTMSketchEntity3) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool)`

GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIdAndReplaceInDependentFields

`func (o *BTMSketchEntity3) SetEntityIdAndReplaceInDependentFields(v string)`

SetEntityIdAndReplaceInDependentFields sets EntityIdAndReplaceInDependentFields field to given value.

### HasEntityIdAndReplaceInDependentFields

`func (o *BTMSketchEntity3) HasEntityIdAndReplaceInDependentFields() bool`

HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTMSketchEntity3) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTMSketchEntity3) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTMSketchEntity3) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTMSketchEntity3) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetIndex

`func (o *BTMSketchEntity3) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BTMSketchEntity3) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BTMSketchEntity3) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BTMSketchEntity3) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *BTMSketchEntity3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMSketchEntity3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMSketchEntity3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMSketchEntity3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *BTMSketchEntity3) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTMSketchEntity3) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTMSketchEntity3) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTMSketchEntity3) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParameters

`func (o *BTMSketchEntity3) GetParameters() []BTMParameter1`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTMSketchEntity3) GetParametersOk() (*[]BTMParameter1, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTMSketchEntity3) SetParameters(v []BTMParameter1)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTMSketchEntity3) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


