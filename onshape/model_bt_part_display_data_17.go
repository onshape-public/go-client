/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartDisplayData17 struct for BTPartDisplayData17
type BTPartDisplayData17 struct {
	Appearance           *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	AppearanceForNewCell *BTGraphicsAppearance1152 `json:"appearanceForNewCell,omitempty"`
	// Type of JSON object.
	BtType                           *string                              `json:"btType,omitempty"`
	CustomProperties                 *BTPartCustomProperties1338          `json:"customProperties,omitempty"`
	DefaultColorHash                 *string                              `json:"defaultColorHash,omitempty"`
	HasFaults                        *bool                                `json:"hasFaults,omitempty"`
	Hidden                           *bool                                `json:"hidden,omitempty"`
	HighBoxCorner                    *BTVector3d389                       `json:"highBoxCorner,omitempty"`
	Id                               *string                              `json:"id,omitempty"`
	IsActiveSheetMetal               *bool                                `json:"isActiveSheetMetal,omitempty"`
	IsMesh                           *bool                                `json:"isMesh,omitempty"`
	IsModifiable                     *bool                                `json:"isModifiable,omitempty"`
	IsSheet                          *bool                                `json:"isSheet,omitempty"`
	IsSolid                          *bool                                `json:"isSolid,omitempty"`
	IsWire                           *bool                                `json:"isWire,omitempty"`
	LowBoxCorner                     *BTVector3d389                       `json:"lowBoxCorner,omitempty"`
	Material                         *BTPartMaterial1445                  `json:"material,omitempty"`
	MaterialForNewCell               *BTPartMaterial1445                  `json:"materialForNewCell,omitempty"`
	MeshState                        *GBTMeshState                        `json:"meshState,omitempty"`
	Name                             *string                              `json:"name,omitempty"`
	NameForNewCell                   *string                              `json:"nameForNewCell,omitempty"`
	Ordinal                          *int32                               `json:"ordinal,omitempty"`
	PartId                           *string                              `json:"partId,omitempty"`
	PropertyIdToSource               *map[string]BTPartMetadataSource2895 `json:"propertyIdToSource,omitempty"`
	ReferencingConfiguredPartNodeIds []BTObjectId                         `json:"referencingConfiguredPartNodeIds,omitempty"`
	Visibility                       *GBTPartVisibility                   `json:"visibility,omitempty"`
}

// NewBTPartDisplayData17 instantiates a new BTPartDisplayData17 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartDisplayData17() *BTPartDisplayData17 {
	this := BTPartDisplayData17{}
	return &this
}

// NewBTPartDisplayData17WithDefaults instantiates a new BTPartDisplayData17 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartDisplayData17WithDefaults() *BTPartDisplayData17 {
	this := BTPartDisplayData17{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTPartDisplayData17) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetAppearanceForNewCell returns the AppearanceForNewCell field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetAppearanceForNewCell() BTGraphicsAppearance1152 {
	if o == nil || o.AppearanceForNewCell == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.AppearanceForNewCell
}

// GetAppearanceForNewCellOk returns a tuple with the AppearanceForNewCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetAppearanceForNewCellOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.AppearanceForNewCell == nil {
		return nil, false
	}
	return o.AppearanceForNewCell, true
}

// HasAppearanceForNewCell returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasAppearanceForNewCell() bool {
	if o != nil && o.AppearanceForNewCell != nil {
		return true
	}

	return false
}

// SetAppearanceForNewCell gets a reference to the given BTGraphicsAppearance1152 and assigns it to the AppearanceForNewCell field.
func (o *BTPartDisplayData17) SetAppearanceForNewCell(v BTGraphicsAppearance1152) {
	o.AppearanceForNewCell = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartDisplayData17) SetBtType(v string) {
	o.BtType = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetCustomProperties() BTPartCustomProperties1338 {
	if o == nil || o.CustomProperties == nil {
		var ret BTPartCustomProperties1338
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetCustomPropertiesOk() (*BTPartCustomProperties1338, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given BTPartCustomProperties1338 and assigns it to the CustomProperties field.
func (o *BTPartDisplayData17) SetCustomProperties(v BTPartCustomProperties1338) {
	o.CustomProperties = &v
}

// GetDefaultColorHash returns the DefaultColorHash field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetDefaultColorHash() string {
	if o == nil || o.DefaultColorHash == nil {
		var ret string
		return ret
	}
	return *o.DefaultColorHash
}

// GetDefaultColorHashOk returns a tuple with the DefaultColorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetDefaultColorHashOk() (*string, bool) {
	if o == nil || o.DefaultColorHash == nil {
		return nil, false
	}
	return o.DefaultColorHash, true
}

// HasDefaultColorHash returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasDefaultColorHash() bool {
	if o != nil && o.DefaultColorHash != nil {
		return true
	}

	return false
}

// SetDefaultColorHash gets a reference to the given string and assigns it to the DefaultColorHash field.
func (o *BTPartDisplayData17) SetDefaultColorHash(v string) {
	o.DefaultColorHash = &v
}

// GetHasFaults returns the HasFaults field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetHasFaults() bool {
	if o == nil || o.HasFaults == nil {
		var ret bool
		return ret
	}
	return *o.HasFaults
}

// GetHasFaultsOk returns a tuple with the HasFaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetHasFaultsOk() (*bool, bool) {
	if o == nil || o.HasFaults == nil {
		return nil, false
	}
	return o.HasFaults, true
}

// HasHasFaults returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasHasFaults() bool {
	if o != nil && o.HasFaults != nil {
		return true
	}

	return false
}

// SetHasFaults gets a reference to the given bool and assigns it to the HasFaults field.
func (o *BTPartDisplayData17) SetHasFaults(v bool) {
	o.HasFaults = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTPartDisplayData17) SetHidden(v bool) {
	o.Hidden = &v
}

// GetHighBoxCorner returns the HighBoxCorner field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetHighBoxCorner() BTVector3d389 {
	if o == nil || o.HighBoxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.HighBoxCorner
}

// GetHighBoxCornerOk returns a tuple with the HighBoxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetHighBoxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.HighBoxCorner == nil {
		return nil, false
	}
	return o.HighBoxCorner, true
}

// HasHighBoxCorner returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasHighBoxCorner() bool {
	if o != nil && o.HighBoxCorner != nil {
		return true
	}

	return false
}

// SetHighBoxCorner gets a reference to the given BTVector3d389 and assigns it to the HighBoxCorner field.
func (o *BTPartDisplayData17) SetHighBoxCorner(v BTVector3d389) {
	o.HighBoxCorner = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPartDisplayData17) SetId(v string) {
	o.Id = &v
}

// GetIsActiveSheetMetal returns the IsActiveSheetMetal field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetIsActiveSheetMetal() bool {
	if o == nil || o.IsActiveSheetMetal == nil {
		var ret bool
		return ret
	}
	return *o.IsActiveSheetMetal
}

// GetIsActiveSheetMetalOk returns a tuple with the IsActiveSheetMetal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIsActiveSheetMetalOk() (*bool, bool) {
	if o == nil || o.IsActiveSheetMetal == nil {
		return nil, false
	}
	return o.IsActiveSheetMetal, true
}

// HasIsActiveSheetMetal returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasIsActiveSheetMetal() bool {
	if o != nil && o.IsActiveSheetMetal != nil {
		return true
	}

	return false
}

// SetIsActiveSheetMetal gets a reference to the given bool and assigns it to the IsActiveSheetMetal field.
func (o *BTPartDisplayData17) SetIsActiveSheetMetal(v bool) {
	o.IsActiveSheetMetal = &v
}

// GetIsMesh returns the IsMesh field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetIsMesh() bool {
	if o == nil || o.IsMesh == nil {
		var ret bool
		return ret
	}
	return *o.IsMesh
}

// GetIsMeshOk returns a tuple with the IsMesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIsMeshOk() (*bool, bool) {
	if o == nil || o.IsMesh == nil {
		return nil, false
	}
	return o.IsMesh, true
}

// HasIsMesh returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasIsMesh() bool {
	if o != nil && o.IsMesh != nil {
		return true
	}

	return false
}

// SetIsMesh gets a reference to the given bool and assigns it to the IsMesh field.
func (o *BTPartDisplayData17) SetIsMesh(v bool) {
	o.IsMesh = &v
}

// GetIsModifiable returns the IsModifiable field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetIsModifiable() bool {
	if o == nil || o.IsModifiable == nil {
		var ret bool
		return ret
	}
	return *o.IsModifiable
}

// GetIsModifiableOk returns a tuple with the IsModifiable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIsModifiableOk() (*bool, bool) {
	if o == nil || o.IsModifiable == nil {
		return nil, false
	}
	return o.IsModifiable, true
}

// HasIsModifiable returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasIsModifiable() bool {
	if o != nil && o.IsModifiable != nil {
		return true
	}

	return false
}

// SetIsModifiable gets a reference to the given bool and assigns it to the IsModifiable field.
func (o *BTPartDisplayData17) SetIsModifiable(v bool) {
	o.IsModifiable = &v
}

// GetIsSheet returns the IsSheet field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetIsSheet() bool {
	if o == nil || o.IsSheet == nil {
		var ret bool
		return ret
	}
	return *o.IsSheet
}

// GetIsSheetOk returns a tuple with the IsSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIsSheetOk() (*bool, bool) {
	if o == nil || o.IsSheet == nil {
		return nil, false
	}
	return o.IsSheet, true
}

// HasIsSheet returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasIsSheet() bool {
	if o != nil && o.IsSheet != nil {
		return true
	}

	return false
}

// SetIsSheet gets a reference to the given bool and assigns it to the IsSheet field.
func (o *BTPartDisplayData17) SetIsSheet(v bool) {
	o.IsSheet = &v
}

// GetIsSolid returns the IsSolid field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetIsSolid() bool {
	if o == nil || o.IsSolid == nil {
		var ret bool
		return ret
	}
	return *o.IsSolid
}

// GetIsSolidOk returns a tuple with the IsSolid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIsSolidOk() (*bool, bool) {
	if o == nil || o.IsSolid == nil {
		return nil, false
	}
	return o.IsSolid, true
}

// HasIsSolid returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasIsSolid() bool {
	if o != nil && o.IsSolid != nil {
		return true
	}

	return false
}

// SetIsSolid gets a reference to the given bool and assigns it to the IsSolid field.
func (o *BTPartDisplayData17) SetIsSolid(v bool) {
	o.IsSolid = &v
}

// GetIsWire returns the IsWire field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetIsWire() bool {
	if o == nil || o.IsWire == nil {
		var ret bool
		return ret
	}
	return *o.IsWire
}

// GetIsWireOk returns a tuple with the IsWire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetIsWireOk() (*bool, bool) {
	if o == nil || o.IsWire == nil {
		return nil, false
	}
	return o.IsWire, true
}

// HasIsWire returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasIsWire() bool {
	if o != nil && o.IsWire != nil {
		return true
	}

	return false
}

// SetIsWire gets a reference to the given bool and assigns it to the IsWire field.
func (o *BTPartDisplayData17) SetIsWire(v bool) {
	o.IsWire = &v
}

// GetLowBoxCorner returns the LowBoxCorner field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetLowBoxCorner() BTVector3d389 {
	if o == nil || o.LowBoxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.LowBoxCorner
}

// GetLowBoxCornerOk returns a tuple with the LowBoxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetLowBoxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.LowBoxCorner == nil {
		return nil, false
	}
	return o.LowBoxCorner, true
}

// HasLowBoxCorner returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasLowBoxCorner() bool {
	if o != nil && o.LowBoxCorner != nil {
		return true
	}

	return false
}

// SetLowBoxCorner gets a reference to the given BTVector3d389 and assigns it to the LowBoxCorner field.
func (o *BTPartDisplayData17) SetLowBoxCorner(v BTVector3d389) {
	o.LowBoxCorner = &v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetMaterial() BTPartMaterial1445 {
	if o == nil || o.Material == nil {
		var ret BTPartMaterial1445
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetMaterialOk() (*BTPartMaterial1445, bool) {
	if o == nil || o.Material == nil {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasMaterial() bool {
	if o != nil && o.Material != nil {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given BTPartMaterial1445 and assigns it to the Material field.
func (o *BTPartDisplayData17) SetMaterial(v BTPartMaterial1445) {
	o.Material = &v
}

// GetMaterialForNewCell returns the MaterialForNewCell field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetMaterialForNewCell() BTPartMaterial1445 {
	if o == nil || o.MaterialForNewCell == nil {
		var ret BTPartMaterial1445
		return ret
	}
	return *o.MaterialForNewCell
}

// GetMaterialForNewCellOk returns a tuple with the MaterialForNewCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetMaterialForNewCellOk() (*BTPartMaterial1445, bool) {
	if o == nil || o.MaterialForNewCell == nil {
		return nil, false
	}
	return o.MaterialForNewCell, true
}

// HasMaterialForNewCell returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasMaterialForNewCell() bool {
	if o != nil && o.MaterialForNewCell != nil {
		return true
	}

	return false
}

// SetMaterialForNewCell gets a reference to the given BTPartMaterial1445 and assigns it to the MaterialForNewCell field.
func (o *BTPartDisplayData17) SetMaterialForNewCell(v BTPartMaterial1445) {
	o.MaterialForNewCell = &v
}

// GetMeshState returns the MeshState field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetMeshState() GBTMeshState {
	if o == nil || o.MeshState == nil {
		var ret GBTMeshState
		return ret
	}
	return *o.MeshState
}

// GetMeshStateOk returns a tuple with the MeshState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetMeshStateOk() (*GBTMeshState, bool) {
	if o == nil || o.MeshState == nil {
		return nil, false
	}
	return o.MeshState, true
}

// HasMeshState returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasMeshState() bool {
	if o != nil && o.MeshState != nil {
		return true
	}

	return false
}

// SetMeshState gets a reference to the given GBTMeshState and assigns it to the MeshState field.
func (o *BTPartDisplayData17) SetMeshState(v GBTMeshState) {
	o.MeshState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPartDisplayData17) SetName(v string) {
	o.Name = &v
}

// GetNameForNewCell returns the NameForNewCell field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetNameForNewCell() string {
	if o == nil || o.NameForNewCell == nil {
		var ret string
		return ret
	}
	return *o.NameForNewCell
}

// GetNameForNewCellOk returns a tuple with the NameForNewCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetNameForNewCellOk() (*string, bool) {
	if o == nil || o.NameForNewCell == nil {
		return nil, false
	}
	return o.NameForNewCell, true
}

// HasNameForNewCell returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasNameForNewCell() bool {
	if o != nil && o.NameForNewCell != nil {
		return true
	}

	return false
}

// SetNameForNewCell gets a reference to the given string and assigns it to the NameForNewCell field.
func (o *BTPartDisplayData17) SetNameForNewCell(v string) {
	o.NameForNewCell = &v
}

// GetOrdinal returns the Ordinal field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetOrdinal() int32 {
	if o == nil || o.Ordinal == nil {
		var ret int32
		return ret
	}
	return *o.Ordinal
}

// GetOrdinalOk returns a tuple with the Ordinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetOrdinalOk() (*int32, bool) {
	if o == nil || o.Ordinal == nil {
		return nil, false
	}
	return o.Ordinal, true
}

// HasOrdinal returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasOrdinal() bool {
	if o != nil && o.Ordinal != nil {
		return true
	}

	return false
}

// SetOrdinal gets a reference to the given int32 and assigns it to the Ordinal field.
func (o *BTPartDisplayData17) SetOrdinal(v int32) {
	o.Ordinal = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTPartDisplayData17) SetPartId(v string) {
	o.PartId = &v
}

// GetPropertyIdToSource returns the PropertyIdToSource field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetPropertyIdToSource() map[string]BTPartMetadataSource2895 {
	if o == nil || o.PropertyIdToSource == nil {
		var ret map[string]BTPartMetadataSource2895
		return ret
	}
	return *o.PropertyIdToSource
}

// GetPropertyIdToSourceOk returns a tuple with the PropertyIdToSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetPropertyIdToSourceOk() (*map[string]BTPartMetadataSource2895, bool) {
	if o == nil || o.PropertyIdToSource == nil {
		return nil, false
	}
	return o.PropertyIdToSource, true
}

// HasPropertyIdToSource returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasPropertyIdToSource() bool {
	if o != nil && o.PropertyIdToSource != nil {
		return true
	}

	return false
}

// SetPropertyIdToSource gets a reference to the given map[string]BTPartMetadataSource2895 and assigns it to the PropertyIdToSource field.
func (o *BTPartDisplayData17) SetPropertyIdToSource(v map[string]BTPartMetadataSource2895) {
	o.PropertyIdToSource = &v
}

// GetReferencingConfiguredPartNodeIds returns the ReferencingConfiguredPartNodeIds field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetReferencingConfiguredPartNodeIds() []BTObjectId {
	if o == nil || o.ReferencingConfiguredPartNodeIds == nil {
		var ret []BTObjectId
		return ret
	}
	return o.ReferencingConfiguredPartNodeIds
}

// GetReferencingConfiguredPartNodeIdsOk returns a tuple with the ReferencingConfiguredPartNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetReferencingConfiguredPartNodeIdsOk() ([]BTObjectId, bool) {
	if o == nil || o.ReferencingConfiguredPartNodeIds == nil {
		return nil, false
	}
	return o.ReferencingConfiguredPartNodeIds, true
}

// HasReferencingConfiguredPartNodeIds returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasReferencingConfiguredPartNodeIds() bool {
	if o != nil && o.ReferencingConfiguredPartNodeIds != nil {
		return true
	}

	return false
}

// SetReferencingConfiguredPartNodeIds gets a reference to the given []BTObjectId and assigns it to the ReferencingConfiguredPartNodeIds field.
func (o *BTPartDisplayData17) SetReferencingConfiguredPartNodeIds(v []BTObjectId) {
	o.ReferencingConfiguredPartNodeIds = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTPartDisplayData17) GetVisibility() GBTPartVisibility {
	if o == nil || o.Visibility == nil {
		var ret GBTPartVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartDisplayData17) GetVisibilityOk() (*GBTPartVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTPartDisplayData17) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given GBTPartVisibility and assigns it to the Visibility field.
func (o *BTPartDisplayData17) SetVisibility(v GBTPartVisibility) {
	o.Visibility = &v
}

func (o BTPartDisplayData17) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.AppearanceForNewCell != nil {
		toSerialize["appearanceForNewCell"] = o.AppearanceForNewCell
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if o.DefaultColorHash != nil {
		toSerialize["defaultColorHash"] = o.DefaultColorHash
	}
	if o.HasFaults != nil {
		toSerialize["hasFaults"] = o.HasFaults
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.HighBoxCorner != nil {
		toSerialize["highBoxCorner"] = o.HighBoxCorner
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsActiveSheetMetal != nil {
		toSerialize["isActiveSheetMetal"] = o.IsActiveSheetMetal
	}
	if o.IsMesh != nil {
		toSerialize["isMesh"] = o.IsMesh
	}
	if o.IsModifiable != nil {
		toSerialize["isModifiable"] = o.IsModifiable
	}
	if o.IsSheet != nil {
		toSerialize["isSheet"] = o.IsSheet
	}
	if o.IsSolid != nil {
		toSerialize["isSolid"] = o.IsSolid
	}
	if o.IsWire != nil {
		toSerialize["isWire"] = o.IsWire
	}
	if o.LowBoxCorner != nil {
		toSerialize["lowBoxCorner"] = o.LowBoxCorner
	}
	if o.Material != nil {
		toSerialize["material"] = o.Material
	}
	if o.MaterialForNewCell != nil {
		toSerialize["materialForNewCell"] = o.MaterialForNewCell
	}
	if o.MeshState != nil {
		toSerialize["meshState"] = o.MeshState
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameForNewCell != nil {
		toSerialize["nameForNewCell"] = o.NameForNewCell
	}
	if o.Ordinal != nil {
		toSerialize["ordinal"] = o.Ordinal
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PropertyIdToSource != nil {
		toSerialize["propertyIdToSource"] = o.PropertyIdToSource
	}
	if o.ReferencingConfiguredPartNodeIds != nil {
		toSerialize["referencingConfiguredPartNodeIds"] = o.ReferencingConfiguredPartNodeIds
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartDisplayData17 struct {
	value *BTPartDisplayData17
	isSet bool
}

func (v NullableBTPartDisplayData17) Get() *BTPartDisplayData17 {
	return v.value
}

func (v *NullableBTPartDisplayData17) Set(val *BTPartDisplayData17) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartDisplayData17) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartDisplayData17) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartDisplayData17(val *BTPartDisplayData17) *NullableBTPartDisplayData17 {
	return &NullableBTPartDisplayData17{value: val, isSet: true}
}

func (v NullableBTPartDisplayData17) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartDisplayData17) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
