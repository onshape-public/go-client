# BTPartDisplayData17

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appearance** | Pointer to [**BTGraphicsAppearance1152**](BTGraphicsAppearance1152.md) |  | [optional] 
**AppearanceForNewCell** | Pointer to [**BTGraphicsAppearance1152**](BTGraphicsAppearance1152.md) |  | [optional] 
**CustomProperties** | Pointer to [**BTPartCustomProperties1338**](BTPartCustomProperties1338.md) |  | [optional] 
**DefaultColorHash** | Pointer to **string** |  | [optional] 
**HasFaults** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**HighBoxCorner** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsActiveSheetMetal** | Pointer to **bool** |  | [optional] 
**IsMesh** | Pointer to **bool** |  | [optional] 
**IsModifiable** | Pointer to **bool** |  | [optional] 
**IsSheet** | Pointer to **bool** |  | [optional] 
**IsSolid** | Pointer to **bool** |  | [optional] 
**IsWire** | Pointer to **bool** |  | [optional] 
**LowBoxCorner** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**Material** | Pointer to [**BTPartMaterial1445**](BTPartMaterial1445.md) |  | [optional] 
**MaterialForNewCell** | Pointer to [**BTPartMaterial1445**](BTPartMaterial1445.md) |  | [optional] 
**MeshState** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameForNewCell** | Pointer to **string** |  | [optional] 
**Ordinal** | Pointer to **int32** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PropertyIdToSource** | Pointer to [**map[string]BTPartMetadataSource2895**](BTPartMetadataSource2895.md) |  | [optional] 
**ReferencingConfiguredPartNodeIds** | Pointer to [**[]BTObjectId**](BTObjectId.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPartDisplayData17

`func NewBTPartDisplayData17() *BTPartDisplayData17`

NewBTPartDisplayData17 instantiates a new BTPartDisplayData17 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPartDisplayData17WithDefaults

`func NewBTPartDisplayData17WithDefaults() *BTPartDisplayData17`

NewBTPartDisplayData17WithDefaults instantiates a new BTPartDisplayData17 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearance

`func (o *BTPartDisplayData17) GetAppearance() BTGraphicsAppearance1152`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *BTPartDisplayData17) GetAppearanceOk() (*BTGraphicsAppearance1152, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *BTPartDisplayData17) SetAppearance(v BTGraphicsAppearance1152)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *BTPartDisplayData17) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetAppearanceForNewCell

`func (o *BTPartDisplayData17) GetAppearanceForNewCell() BTGraphicsAppearance1152`

GetAppearanceForNewCell returns the AppearanceForNewCell field if non-nil, zero value otherwise.

### GetAppearanceForNewCellOk

`func (o *BTPartDisplayData17) GetAppearanceForNewCellOk() (*BTGraphicsAppearance1152, bool)`

GetAppearanceForNewCellOk returns a tuple with the AppearanceForNewCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearanceForNewCell

`func (o *BTPartDisplayData17) SetAppearanceForNewCell(v BTGraphicsAppearance1152)`

SetAppearanceForNewCell sets AppearanceForNewCell field to given value.

### HasAppearanceForNewCell

`func (o *BTPartDisplayData17) HasAppearanceForNewCell() bool`

HasAppearanceForNewCell returns a boolean if a field has been set.

### GetCustomProperties

`func (o *BTPartDisplayData17) GetCustomProperties() BTPartCustomProperties1338`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BTPartDisplayData17) GetCustomPropertiesOk() (*BTPartCustomProperties1338, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BTPartDisplayData17) SetCustomProperties(v BTPartCustomProperties1338)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BTPartDisplayData17) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDefaultColorHash

`func (o *BTPartDisplayData17) GetDefaultColorHash() string`

GetDefaultColorHash returns the DefaultColorHash field if non-nil, zero value otherwise.

### GetDefaultColorHashOk

`func (o *BTPartDisplayData17) GetDefaultColorHashOk() (*string, bool)`

GetDefaultColorHashOk returns a tuple with the DefaultColorHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColorHash

`func (o *BTPartDisplayData17) SetDefaultColorHash(v string)`

SetDefaultColorHash sets DefaultColorHash field to given value.

### HasDefaultColorHash

`func (o *BTPartDisplayData17) HasDefaultColorHash() bool`

HasDefaultColorHash returns a boolean if a field has been set.

### GetHasFaults

`func (o *BTPartDisplayData17) GetHasFaults() bool`

GetHasFaults returns the HasFaults field if non-nil, zero value otherwise.

### GetHasFaultsOk

`func (o *BTPartDisplayData17) GetHasFaultsOk() (*bool, bool)`

GetHasFaultsOk returns a tuple with the HasFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFaults

`func (o *BTPartDisplayData17) SetHasFaults(v bool)`

SetHasFaults sets HasFaults field to given value.

### HasHasFaults

`func (o *BTPartDisplayData17) HasHasFaults() bool`

HasHasFaults returns a boolean if a field has been set.

### GetHidden

`func (o *BTPartDisplayData17) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BTPartDisplayData17) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BTPartDisplayData17) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BTPartDisplayData17) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHighBoxCorner

`func (o *BTPartDisplayData17) GetHighBoxCorner() BTVector3d389`

GetHighBoxCorner returns the HighBoxCorner field if non-nil, zero value otherwise.

### GetHighBoxCornerOk

`func (o *BTPartDisplayData17) GetHighBoxCornerOk() (*BTVector3d389, bool)`

GetHighBoxCornerOk returns a tuple with the HighBoxCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighBoxCorner

`func (o *BTPartDisplayData17) SetHighBoxCorner(v BTVector3d389)`

SetHighBoxCorner sets HighBoxCorner field to given value.

### HasHighBoxCorner

`func (o *BTPartDisplayData17) HasHighBoxCorner() bool`

HasHighBoxCorner returns a boolean if a field has been set.

### GetId

`func (o *BTPartDisplayData17) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPartDisplayData17) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPartDisplayData17) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPartDisplayData17) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActiveSheetMetal

`func (o *BTPartDisplayData17) GetIsActiveSheetMetal() bool`

GetIsActiveSheetMetal returns the IsActiveSheetMetal field if non-nil, zero value otherwise.

### GetIsActiveSheetMetalOk

`func (o *BTPartDisplayData17) GetIsActiveSheetMetalOk() (*bool, bool)`

GetIsActiveSheetMetalOk returns a tuple with the IsActiveSheetMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveSheetMetal

`func (o *BTPartDisplayData17) SetIsActiveSheetMetal(v bool)`

SetIsActiveSheetMetal sets IsActiveSheetMetal field to given value.

### HasIsActiveSheetMetal

`func (o *BTPartDisplayData17) HasIsActiveSheetMetal() bool`

HasIsActiveSheetMetal returns a boolean if a field has been set.

### GetIsMesh

`func (o *BTPartDisplayData17) GetIsMesh() bool`

GetIsMesh returns the IsMesh field if non-nil, zero value otherwise.

### GetIsMeshOk

`func (o *BTPartDisplayData17) GetIsMeshOk() (*bool, bool)`

GetIsMeshOk returns a tuple with the IsMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMesh

`func (o *BTPartDisplayData17) SetIsMesh(v bool)`

SetIsMesh sets IsMesh field to given value.

### HasIsMesh

`func (o *BTPartDisplayData17) HasIsMesh() bool`

HasIsMesh returns a boolean if a field has been set.

### GetIsModifiable

`func (o *BTPartDisplayData17) GetIsModifiable() bool`

GetIsModifiable returns the IsModifiable field if non-nil, zero value otherwise.

### GetIsModifiableOk

`func (o *BTPartDisplayData17) GetIsModifiableOk() (*bool, bool)`

GetIsModifiableOk returns a tuple with the IsModifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModifiable

`func (o *BTPartDisplayData17) SetIsModifiable(v bool)`

SetIsModifiable sets IsModifiable field to given value.

### HasIsModifiable

`func (o *BTPartDisplayData17) HasIsModifiable() bool`

HasIsModifiable returns a boolean if a field has been set.

### GetIsSheet

`func (o *BTPartDisplayData17) GetIsSheet() bool`

GetIsSheet returns the IsSheet field if non-nil, zero value otherwise.

### GetIsSheetOk

`func (o *BTPartDisplayData17) GetIsSheetOk() (*bool, bool)`

GetIsSheetOk returns a tuple with the IsSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSheet

`func (o *BTPartDisplayData17) SetIsSheet(v bool)`

SetIsSheet sets IsSheet field to given value.

### HasIsSheet

`func (o *BTPartDisplayData17) HasIsSheet() bool`

HasIsSheet returns a boolean if a field has been set.

### GetIsSolid

`func (o *BTPartDisplayData17) GetIsSolid() bool`

GetIsSolid returns the IsSolid field if non-nil, zero value otherwise.

### GetIsSolidOk

`func (o *BTPartDisplayData17) GetIsSolidOk() (*bool, bool)`

GetIsSolidOk returns a tuple with the IsSolid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSolid

`func (o *BTPartDisplayData17) SetIsSolid(v bool)`

SetIsSolid sets IsSolid field to given value.

### HasIsSolid

`func (o *BTPartDisplayData17) HasIsSolid() bool`

HasIsSolid returns a boolean if a field has been set.

### GetIsWire

`func (o *BTPartDisplayData17) GetIsWire() bool`

GetIsWire returns the IsWire field if non-nil, zero value otherwise.

### GetIsWireOk

`func (o *BTPartDisplayData17) GetIsWireOk() (*bool, bool)`

GetIsWireOk returns a tuple with the IsWire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWire

`func (o *BTPartDisplayData17) SetIsWire(v bool)`

SetIsWire sets IsWire field to given value.

### HasIsWire

`func (o *BTPartDisplayData17) HasIsWire() bool`

HasIsWire returns a boolean if a field has been set.

### GetLowBoxCorner

`func (o *BTPartDisplayData17) GetLowBoxCorner() BTVector3d389`

GetLowBoxCorner returns the LowBoxCorner field if non-nil, zero value otherwise.

### GetLowBoxCornerOk

`func (o *BTPartDisplayData17) GetLowBoxCornerOk() (*BTVector3d389, bool)`

GetLowBoxCornerOk returns a tuple with the LowBoxCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBoxCorner

`func (o *BTPartDisplayData17) SetLowBoxCorner(v BTVector3d389)`

SetLowBoxCorner sets LowBoxCorner field to given value.

### HasLowBoxCorner

`func (o *BTPartDisplayData17) HasLowBoxCorner() bool`

HasLowBoxCorner returns a boolean if a field has been set.

### GetMaterial

`func (o *BTPartDisplayData17) GetMaterial() BTPartMaterial1445`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *BTPartDisplayData17) GetMaterialOk() (*BTPartMaterial1445, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *BTPartDisplayData17) SetMaterial(v BTPartMaterial1445)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *BTPartDisplayData17) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetMaterialForNewCell

`func (o *BTPartDisplayData17) GetMaterialForNewCell() BTPartMaterial1445`

GetMaterialForNewCell returns the MaterialForNewCell field if non-nil, zero value otherwise.

### GetMaterialForNewCellOk

`func (o *BTPartDisplayData17) GetMaterialForNewCellOk() (*BTPartMaterial1445, bool)`

GetMaterialForNewCellOk returns a tuple with the MaterialForNewCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialForNewCell

`func (o *BTPartDisplayData17) SetMaterialForNewCell(v BTPartMaterial1445)`

SetMaterialForNewCell sets MaterialForNewCell field to given value.

### HasMaterialForNewCell

`func (o *BTPartDisplayData17) HasMaterialForNewCell() bool`

HasMaterialForNewCell returns a boolean if a field has been set.

### GetMeshState

`func (o *BTPartDisplayData17) GetMeshState() string`

GetMeshState returns the MeshState field if non-nil, zero value otherwise.

### GetMeshStateOk

`func (o *BTPartDisplayData17) GetMeshStateOk() (*string, bool)`

GetMeshStateOk returns a tuple with the MeshState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshState

`func (o *BTPartDisplayData17) SetMeshState(v string)`

SetMeshState sets MeshState field to given value.

### HasMeshState

`func (o *BTPartDisplayData17) HasMeshState() bool`

HasMeshState returns a boolean if a field has been set.

### GetName

`func (o *BTPartDisplayData17) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPartDisplayData17) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPartDisplayData17) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTPartDisplayData17) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameForNewCell

`func (o *BTPartDisplayData17) GetNameForNewCell() string`

GetNameForNewCell returns the NameForNewCell field if non-nil, zero value otherwise.

### GetNameForNewCellOk

`func (o *BTPartDisplayData17) GetNameForNewCellOk() (*string, bool)`

GetNameForNewCellOk returns a tuple with the NameForNewCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameForNewCell

`func (o *BTPartDisplayData17) SetNameForNewCell(v string)`

SetNameForNewCell sets NameForNewCell field to given value.

### HasNameForNewCell

`func (o *BTPartDisplayData17) HasNameForNewCell() bool`

HasNameForNewCell returns a boolean if a field has been set.

### GetOrdinal

`func (o *BTPartDisplayData17) GetOrdinal() int32`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *BTPartDisplayData17) GetOrdinalOk() (*int32, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *BTPartDisplayData17) SetOrdinal(v int32)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *BTPartDisplayData17) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### GetPartId

`func (o *BTPartDisplayData17) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTPartDisplayData17) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTPartDisplayData17) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTPartDisplayData17) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPropertyIdToSource

`func (o *BTPartDisplayData17) GetPropertyIdToSource() map[string]BTPartMetadataSource2895`

GetPropertyIdToSource returns the PropertyIdToSource field if non-nil, zero value otherwise.

### GetPropertyIdToSourceOk

`func (o *BTPartDisplayData17) GetPropertyIdToSourceOk() (*map[string]BTPartMetadataSource2895, bool)`

GetPropertyIdToSourceOk returns a tuple with the PropertyIdToSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIdToSource

`func (o *BTPartDisplayData17) SetPropertyIdToSource(v map[string]BTPartMetadataSource2895)`

SetPropertyIdToSource sets PropertyIdToSource field to given value.

### HasPropertyIdToSource

`func (o *BTPartDisplayData17) HasPropertyIdToSource() bool`

HasPropertyIdToSource returns a boolean if a field has been set.

### GetReferencingConfiguredPartNodeIds

`func (o *BTPartDisplayData17) GetReferencingConfiguredPartNodeIds() []BTObjectId`

GetReferencingConfiguredPartNodeIds returns the ReferencingConfiguredPartNodeIds field if non-nil, zero value otherwise.

### GetReferencingConfiguredPartNodeIdsOk

`func (o *BTPartDisplayData17) GetReferencingConfiguredPartNodeIdsOk() (*[]BTObjectId, bool)`

GetReferencingConfiguredPartNodeIdsOk returns a tuple with the ReferencingConfiguredPartNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingConfiguredPartNodeIds

`func (o *BTPartDisplayData17) SetReferencingConfiguredPartNodeIds(v []BTObjectId)`

SetReferencingConfiguredPartNodeIds sets ReferencingConfiguredPartNodeIds field to given value.

### HasReferencingConfiguredPartNodeIds

`func (o *BTPartDisplayData17) HasReferencingConfiguredPartNodeIds() bool`

HasReferencingConfiguredPartNodeIds returns a boolean if a field has been set.

### GetVisibility

`func (o *BTPartDisplayData17) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BTPartDisplayData17) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BTPartDisplayData17) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BTPartDisplayData17) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


