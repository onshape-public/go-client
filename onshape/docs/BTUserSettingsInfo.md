# BTUserSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AxisRotationLock** | Pointer to **bool** |  | [optional] 
**CommonUnits** | Pointer to [**BTCommonUnitsInfo**](BTCommonUnitsInfo.md) |  | [optional] 
**CustomColors** | Pointer to **[]string** |  | [optional] 
**DefaultUnits** | Pointer to [**BTDefaultUnitsInfo**](BTDefaultUnitsInfo.md) |  | [optional] 
**DisplayAssemblyProperties** | Pointer to **bool** |  | [optional] 
**DrawingBackgroundId** | Pointer to **int32** |  | [optional] 
**EnforceApplicationAcl** | Pointer to **bool** |  | [optional] 
**ExportDrawingOptions** | Pointer to **string** |  | [optional] 
**ExportSolidOptions** | Pointer to **string** |  | [optional] 
**GraphicsRenderMode** | Pointer to **string** |  | [optional] 
**GraphicsSmoothEdge** | Pointer to **string** |  | [optional] 
**HighlightLaminarEdges** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImportOptions** | Pointer to **string** |  | [optional] 
**IsolateEnableSelectionDesire** | Pointer to **bool** |  | [optional] 
**IsolateHideTransparent** | Pointer to **string** |  | [optional] 
**IsolateOpacitySliderValue** | Pointer to **int32** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**MakeTransparentEnableSelectionDesire** | Pointer to **bool** |  | [optional] 
**MakeTransparentOpacitySliderValue** | Pointer to **int32** |  | [optional] 
**MaterialLibrarySettings** | Pointer to [**BTMaterialLibrarySettingsInfo**](BTMaterialLibrarySettingsInfo.md) |  | [optional] 
**MiniToolbarSettings** | Pointer to **string** |  | [optional] 
**MouseActions** | Pointer to **string** |  | [optional] 
**PerspectiveModeOn** | Pointer to **string** |  | [optional] 
**PreviousSketchFont** | Pointer to **string** |  | [optional] 
**ReverseScrollWheelZoomDirection** | Pointer to **bool** |  | [optional] 
**SelectItemViewStateInfos** | Pointer to [**[]BTSelectItemViewStateInfo**](BTSelectItemViewStateInfo.md) |  | [optional] 
**StartupPage** | Pointer to **int32** |  | [optional] 
**SubstituteApprovers** | Pointer to [**[]BTSubstituteApproverInfo**](BTSubstituteApproverInfo.md) |  | [optional] 
**Theme** | Pointer to **int32** |  | [optional] 
**UnitsDisplayPrecision** | Pointer to **map[string]int32** |  | [optional] 
**UnitsMaximumDisplayPrecision** | Pointer to [**BTUnitsMaximumDisplayPrecisionInfo**](BTUnitsMaximumDisplayPrecisionInfo.md) |  | [optional] 
**Use24HourTime** | Pointer to **bool** |  | [optional] 
**UseDecimalComma** | Pointer to **bool** |  | [optional] 
**ViewManipulationMouseKeyMapping** | Pointer to [**BTViewManipulationMouseKeyMappingInfo**](BTViewManipulationMouseKeyMappingInfo.md) |  | [optional] 
**ViewMappingId** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTUserSettingsInfo

`func NewBTUserSettingsInfo() *BTUserSettingsInfo`

NewBTUserSettingsInfo instantiates a new BTUserSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserSettingsInfoWithDefaults

`func NewBTUserSettingsInfoWithDefaults() *BTUserSettingsInfo`

NewBTUserSettingsInfoWithDefaults instantiates a new BTUserSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxisRotationLock

`func (o *BTUserSettingsInfo) GetAxisRotationLock() bool`

GetAxisRotationLock returns the AxisRotationLock field if non-nil, zero value otherwise.

### GetAxisRotationLockOk

`func (o *BTUserSettingsInfo) GetAxisRotationLockOk() (*bool, bool)`

GetAxisRotationLockOk returns a tuple with the AxisRotationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxisRotationLock

`func (o *BTUserSettingsInfo) SetAxisRotationLock(v bool)`

SetAxisRotationLock sets AxisRotationLock field to given value.

### HasAxisRotationLock

`func (o *BTUserSettingsInfo) HasAxisRotationLock() bool`

HasAxisRotationLock returns a boolean if a field has been set.

### GetCommonUnits

`func (o *BTUserSettingsInfo) GetCommonUnits() BTCommonUnitsInfo`

GetCommonUnits returns the CommonUnits field if non-nil, zero value otherwise.

### GetCommonUnitsOk

`func (o *BTUserSettingsInfo) GetCommonUnitsOk() (*BTCommonUnitsInfo, bool)`

GetCommonUnitsOk returns a tuple with the CommonUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonUnits

`func (o *BTUserSettingsInfo) SetCommonUnits(v BTCommonUnitsInfo)`

SetCommonUnits sets CommonUnits field to given value.

### HasCommonUnits

`func (o *BTUserSettingsInfo) HasCommonUnits() bool`

HasCommonUnits returns a boolean if a field has been set.

### GetCustomColors

`func (o *BTUserSettingsInfo) GetCustomColors() []string`

GetCustomColors returns the CustomColors field if non-nil, zero value otherwise.

### GetCustomColorsOk

`func (o *BTUserSettingsInfo) GetCustomColorsOk() (*[]string, bool)`

GetCustomColorsOk returns a tuple with the CustomColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomColors

`func (o *BTUserSettingsInfo) SetCustomColors(v []string)`

SetCustomColors sets CustomColors field to given value.

### HasCustomColors

`func (o *BTUserSettingsInfo) HasCustomColors() bool`

HasCustomColors returns a boolean if a field has been set.

### GetDefaultUnits

`func (o *BTUserSettingsInfo) GetDefaultUnits() BTDefaultUnitsInfo`

GetDefaultUnits returns the DefaultUnits field if non-nil, zero value otherwise.

### GetDefaultUnitsOk

`func (o *BTUserSettingsInfo) GetDefaultUnitsOk() (*BTDefaultUnitsInfo, bool)`

GetDefaultUnitsOk returns a tuple with the DefaultUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUnits

`func (o *BTUserSettingsInfo) SetDefaultUnits(v BTDefaultUnitsInfo)`

SetDefaultUnits sets DefaultUnits field to given value.

### HasDefaultUnits

`func (o *BTUserSettingsInfo) HasDefaultUnits() bool`

HasDefaultUnits returns a boolean if a field has been set.

### GetDisplayAssemblyProperties

`func (o *BTUserSettingsInfo) GetDisplayAssemblyProperties() bool`

GetDisplayAssemblyProperties returns the DisplayAssemblyProperties field if non-nil, zero value otherwise.

### GetDisplayAssemblyPropertiesOk

`func (o *BTUserSettingsInfo) GetDisplayAssemblyPropertiesOk() (*bool, bool)`

GetDisplayAssemblyPropertiesOk returns a tuple with the DisplayAssemblyProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssemblyProperties

`func (o *BTUserSettingsInfo) SetDisplayAssemblyProperties(v bool)`

SetDisplayAssemblyProperties sets DisplayAssemblyProperties field to given value.

### HasDisplayAssemblyProperties

`func (o *BTUserSettingsInfo) HasDisplayAssemblyProperties() bool`

HasDisplayAssemblyProperties returns a boolean if a field has been set.

### GetDrawingBackgroundId

`func (o *BTUserSettingsInfo) GetDrawingBackgroundId() int32`

GetDrawingBackgroundId returns the DrawingBackgroundId field if non-nil, zero value otherwise.

### GetDrawingBackgroundIdOk

`func (o *BTUserSettingsInfo) GetDrawingBackgroundIdOk() (*int32, bool)`

GetDrawingBackgroundIdOk returns a tuple with the DrawingBackgroundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawingBackgroundId

`func (o *BTUserSettingsInfo) SetDrawingBackgroundId(v int32)`

SetDrawingBackgroundId sets DrawingBackgroundId field to given value.

### HasDrawingBackgroundId

`func (o *BTUserSettingsInfo) HasDrawingBackgroundId() bool`

HasDrawingBackgroundId returns a boolean if a field has been set.

### GetEnforceApplicationAcl

`func (o *BTUserSettingsInfo) GetEnforceApplicationAcl() bool`

GetEnforceApplicationAcl returns the EnforceApplicationAcl field if non-nil, zero value otherwise.

### GetEnforceApplicationAclOk

`func (o *BTUserSettingsInfo) GetEnforceApplicationAclOk() (*bool, bool)`

GetEnforceApplicationAclOk returns a tuple with the EnforceApplicationAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceApplicationAcl

`func (o *BTUserSettingsInfo) SetEnforceApplicationAcl(v bool)`

SetEnforceApplicationAcl sets EnforceApplicationAcl field to given value.

### HasEnforceApplicationAcl

`func (o *BTUserSettingsInfo) HasEnforceApplicationAcl() bool`

HasEnforceApplicationAcl returns a boolean if a field has been set.

### GetExportDrawingOptions

`func (o *BTUserSettingsInfo) GetExportDrawingOptions() string`

GetExportDrawingOptions returns the ExportDrawingOptions field if non-nil, zero value otherwise.

### GetExportDrawingOptionsOk

`func (o *BTUserSettingsInfo) GetExportDrawingOptionsOk() (*string, bool)`

GetExportDrawingOptionsOk returns a tuple with the ExportDrawingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportDrawingOptions

`func (o *BTUserSettingsInfo) SetExportDrawingOptions(v string)`

SetExportDrawingOptions sets ExportDrawingOptions field to given value.

### HasExportDrawingOptions

`func (o *BTUserSettingsInfo) HasExportDrawingOptions() bool`

HasExportDrawingOptions returns a boolean if a field has been set.

### GetExportSolidOptions

`func (o *BTUserSettingsInfo) GetExportSolidOptions() string`

GetExportSolidOptions returns the ExportSolidOptions field if non-nil, zero value otherwise.

### GetExportSolidOptionsOk

`func (o *BTUserSettingsInfo) GetExportSolidOptionsOk() (*string, bool)`

GetExportSolidOptionsOk returns a tuple with the ExportSolidOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSolidOptions

`func (o *BTUserSettingsInfo) SetExportSolidOptions(v string)`

SetExportSolidOptions sets ExportSolidOptions field to given value.

### HasExportSolidOptions

`func (o *BTUserSettingsInfo) HasExportSolidOptions() bool`

HasExportSolidOptions returns a boolean if a field has been set.

### GetGraphicsRenderMode

`func (o *BTUserSettingsInfo) GetGraphicsRenderMode() string`

GetGraphicsRenderMode returns the GraphicsRenderMode field if non-nil, zero value otherwise.

### GetGraphicsRenderModeOk

`func (o *BTUserSettingsInfo) GetGraphicsRenderModeOk() (*string, bool)`

GetGraphicsRenderModeOk returns a tuple with the GraphicsRenderMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsRenderMode

`func (o *BTUserSettingsInfo) SetGraphicsRenderMode(v string)`

SetGraphicsRenderMode sets GraphicsRenderMode field to given value.

### HasGraphicsRenderMode

`func (o *BTUserSettingsInfo) HasGraphicsRenderMode() bool`

HasGraphicsRenderMode returns a boolean if a field has been set.

### GetGraphicsSmoothEdge

`func (o *BTUserSettingsInfo) GetGraphicsSmoothEdge() string`

GetGraphicsSmoothEdge returns the GraphicsSmoothEdge field if non-nil, zero value otherwise.

### GetGraphicsSmoothEdgeOk

`func (o *BTUserSettingsInfo) GetGraphicsSmoothEdgeOk() (*string, bool)`

GetGraphicsSmoothEdgeOk returns a tuple with the GraphicsSmoothEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsSmoothEdge

`func (o *BTUserSettingsInfo) SetGraphicsSmoothEdge(v string)`

SetGraphicsSmoothEdge sets GraphicsSmoothEdge field to given value.

### HasGraphicsSmoothEdge

`func (o *BTUserSettingsInfo) HasGraphicsSmoothEdge() bool`

HasGraphicsSmoothEdge returns a boolean if a field has been set.

### GetHighlightLaminarEdges

`func (o *BTUserSettingsInfo) GetHighlightLaminarEdges() string`

GetHighlightLaminarEdges returns the HighlightLaminarEdges field if non-nil, zero value otherwise.

### GetHighlightLaminarEdgesOk

`func (o *BTUserSettingsInfo) GetHighlightLaminarEdgesOk() (*string, bool)`

GetHighlightLaminarEdgesOk returns a tuple with the HighlightLaminarEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightLaminarEdges

`func (o *BTUserSettingsInfo) SetHighlightLaminarEdges(v string)`

SetHighlightLaminarEdges sets HighlightLaminarEdges field to given value.

### HasHighlightLaminarEdges

`func (o *BTUserSettingsInfo) HasHighlightLaminarEdges() bool`

HasHighlightLaminarEdges returns a boolean if a field has been set.

### GetId

`func (o *BTUserSettingsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTUserSettingsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTUserSettingsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTUserSettingsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImportOptions

`func (o *BTUserSettingsInfo) GetImportOptions() string`

GetImportOptions returns the ImportOptions field if non-nil, zero value otherwise.

### GetImportOptionsOk

`func (o *BTUserSettingsInfo) GetImportOptionsOk() (*string, bool)`

GetImportOptionsOk returns a tuple with the ImportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOptions

`func (o *BTUserSettingsInfo) SetImportOptions(v string)`

SetImportOptions sets ImportOptions field to given value.

### HasImportOptions

`func (o *BTUserSettingsInfo) HasImportOptions() bool`

HasImportOptions returns a boolean if a field has been set.

### GetIsolateEnableSelectionDesire

`func (o *BTUserSettingsInfo) GetIsolateEnableSelectionDesire() bool`

GetIsolateEnableSelectionDesire returns the IsolateEnableSelectionDesire field if non-nil, zero value otherwise.

### GetIsolateEnableSelectionDesireOk

`func (o *BTUserSettingsInfo) GetIsolateEnableSelectionDesireOk() (*bool, bool)`

GetIsolateEnableSelectionDesireOk returns a tuple with the IsolateEnableSelectionDesire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateEnableSelectionDesire

`func (o *BTUserSettingsInfo) SetIsolateEnableSelectionDesire(v bool)`

SetIsolateEnableSelectionDesire sets IsolateEnableSelectionDesire field to given value.

### HasIsolateEnableSelectionDesire

`func (o *BTUserSettingsInfo) HasIsolateEnableSelectionDesire() bool`

HasIsolateEnableSelectionDesire returns a boolean if a field has been set.

### GetIsolateHideTransparent

`func (o *BTUserSettingsInfo) GetIsolateHideTransparent() string`

GetIsolateHideTransparent returns the IsolateHideTransparent field if non-nil, zero value otherwise.

### GetIsolateHideTransparentOk

`func (o *BTUserSettingsInfo) GetIsolateHideTransparentOk() (*string, bool)`

GetIsolateHideTransparentOk returns a tuple with the IsolateHideTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateHideTransparent

`func (o *BTUserSettingsInfo) SetIsolateHideTransparent(v string)`

SetIsolateHideTransparent sets IsolateHideTransparent field to given value.

### HasIsolateHideTransparent

`func (o *BTUserSettingsInfo) HasIsolateHideTransparent() bool`

HasIsolateHideTransparent returns a boolean if a field has been set.

### GetIsolateOpacitySliderValue

`func (o *BTUserSettingsInfo) GetIsolateOpacitySliderValue() int32`

GetIsolateOpacitySliderValue returns the IsolateOpacitySliderValue field if non-nil, zero value otherwise.

### GetIsolateOpacitySliderValueOk

`func (o *BTUserSettingsInfo) GetIsolateOpacitySliderValueOk() (*int32, bool)`

GetIsolateOpacitySliderValueOk returns a tuple with the IsolateOpacitySliderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateOpacitySliderValue

`func (o *BTUserSettingsInfo) SetIsolateOpacitySliderValue(v int32)`

SetIsolateOpacitySliderValue sets IsolateOpacitySliderValue field to given value.

### HasIsolateOpacitySliderValue

`func (o *BTUserSettingsInfo) HasIsolateOpacitySliderValue() bool`

HasIsolateOpacitySliderValue returns a boolean if a field has been set.

### GetLocale

`func (o *BTUserSettingsInfo) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *BTUserSettingsInfo) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *BTUserSettingsInfo) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *BTUserSettingsInfo) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMakeTransparentEnableSelectionDesire

`func (o *BTUserSettingsInfo) GetMakeTransparentEnableSelectionDesire() bool`

GetMakeTransparentEnableSelectionDesire returns the MakeTransparentEnableSelectionDesire field if non-nil, zero value otherwise.

### GetMakeTransparentEnableSelectionDesireOk

`func (o *BTUserSettingsInfo) GetMakeTransparentEnableSelectionDesireOk() (*bool, bool)`

GetMakeTransparentEnableSelectionDesireOk returns a tuple with the MakeTransparentEnableSelectionDesire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeTransparentEnableSelectionDesire

`func (o *BTUserSettingsInfo) SetMakeTransparentEnableSelectionDesire(v bool)`

SetMakeTransparentEnableSelectionDesire sets MakeTransparentEnableSelectionDesire field to given value.

### HasMakeTransparentEnableSelectionDesire

`func (o *BTUserSettingsInfo) HasMakeTransparentEnableSelectionDesire() bool`

HasMakeTransparentEnableSelectionDesire returns a boolean if a field has been set.

### GetMakeTransparentOpacitySliderValue

`func (o *BTUserSettingsInfo) GetMakeTransparentOpacitySliderValue() int32`

GetMakeTransparentOpacitySliderValue returns the MakeTransparentOpacitySliderValue field if non-nil, zero value otherwise.

### GetMakeTransparentOpacitySliderValueOk

`func (o *BTUserSettingsInfo) GetMakeTransparentOpacitySliderValueOk() (*int32, bool)`

GetMakeTransparentOpacitySliderValueOk returns a tuple with the MakeTransparentOpacitySliderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeTransparentOpacitySliderValue

`func (o *BTUserSettingsInfo) SetMakeTransparentOpacitySliderValue(v int32)`

SetMakeTransparentOpacitySliderValue sets MakeTransparentOpacitySliderValue field to given value.

### HasMakeTransparentOpacitySliderValue

`func (o *BTUserSettingsInfo) HasMakeTransparentOpacitySliderValue() bool`

HasMakeTransparentOpacitySliderValue returns a boolean if a field has been set.

### GetMaterialLibrarySettings

`func (o *BTUserSettingsInfo) GetMaterialLibrarySettings() BTMaterialLibrarySettingsInfo`

GetMaterialLibrarySettings returns the MaterialLibrarySettings field if non-nil, zero value otherwise.

### GetMaterialLibrarySettingsOk

`func (o *BTUserSettingsInfo) GetMaterialLibrarySettingsOk() (*BTMaterialLibrarySettingsInfo, bool)`

GetMaterialLibrarySettingsOk returns a tuple with the MaterialLibrarySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialLibrarySettings

`func (o *BTUserSettingsInfo) SetMaterialLibrarySettings(v BTMaterialLibrarySettingsInfo)`

SetMaterialLibrarySettings sets MaterialLibrarySettings field to given value.

### HasMaterialLibrarySettings

`func (o *BTUserSettingsInfo) HasMaterialLibrarySettings() bool`

HasMaterialLibrarySettings returns a boolean if a field has been set.

### GetMiniToolbarSettings

`func (o *BTUserSettingsInfo) GetMiniToolbarSettings() string`

GetMiniToolbarSettings returns the MiniToolbarSettings field if non-nil, zero value otherwise.

### GetMiniToolbarSettingsOk

`func (o *BTUserSettingsInfo) GetMiniToolbarSettingsOk() (*string, bool)`

GetMiniToolbarSettingsOk returns a tuple with the MiniToolbarSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiniToolbarSettings

`func (o *BTUserSettingsInfo) SetMiniToolbarSettings(v string)`

SetMiniToolbarSettings sets MiniToolbarSettings field to given value.

### HasMiniToolbarSettings

`func (o *BTUserSettingsInfo) HasMiniToolbarSettings() bool`

HasMiniToolbarSettings returns a boolean if a field has been set.

### GetMouseActions

`func (o *BTUserSettingsInfo) GetMouseActions() string`

GetMouseActions returns the MouseActions field if non-nil, zero value otherwise.

### GetMouseActionsOk

`func (o *BTUserSettingsInfo) GetMouseActionsOk() (*string, bool)`

GetMouseActionsOk returns a tuple with the MouseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseActions

`func (o *BTUserSettingsInfo) SetMouseActions(v string)`

SetMouseActions sets MouseActions field to given value.

### HasMouseActions

`func (o *BTUserSettingsInfo) HasMouseActions() bool`

HasMouseActions returns a boolean if a field has been set.

### GetPerspectiveModeOn

`func (o *BTUserSettingsInfo) GetPerspectiveModeOn() string`

GetPerspectiveModeOn returns the PerspectiveModeOn field if non-nil, zero value otherwise.

### GetPerspectiveModeOnOk

`func (o *BTUserSettingsInfo) GetPerspectiveModeOnOk() (*string, bool)`

GetPerspectiveModeOnOk returns a tuple with the PerspectiveModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerspectiveModeOn

`func (o *BTUserSettingsInfo) SetPerspectiveModeOn(v string)`

SetPerspectiveModeOn sets PerspectiveModeOn field to given value.

### HasPerspectiveModeOn

`func (o *BTUserSettingsInfo) HasPerspectiveModeOn() bool`

HasPerspectiveModeOn returns a boolean if a field has been set.

### GetPreviousSketchFont

`func (o *BTUserSettingsInfo) GetPreviousSketchFont() string`

GetPreviousSketchFont returns the PreviousSketchFont field if non-nil, zero value otherwise.

### GetPreviousSketchFontOk

`func (o *BTUserSettingsInfo) GetPreviousSketchFontOk() (*string, bool)`

GetPreviousSketchFontOk returns a tuple with the PreviousSketchFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSketchFont

`func (o *BTUserSettingsInfo) SetPreviousSketchFont(v string)`

SetPreviousSketchFont sets PreviousSketchFont field to given value.

### HasPreviousSketchFont

`func (o *BTUserSettingsInfo) HasPreviousSketchFont() bool`

HasPreviousSketchFont returns a boolean if a field has been set.

### GetReverseScrollWheelZoomDirection

`func (o *BTUserSettingsInfo) GetReverseScrollWheelZoomDirection() bool`

GetReverseScrollWheelZoomDirection returns the ReverseScrollWheelZoomDirection field if non-nil, zero value otherwise.

### GetReverseScrollWheelZoomDirectionOk

`func (o *BTUserSettingsInfo) GetReverseScrollWheelZoomDirectionOk() (*bool, bool)`

GetReverseScrollWheelZoomDirectionOk returns a tuple with the ReverseScrollWheelZoomDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseScrollWheelZoomDirection

`func (o *BTUserSettingsInfo) SetReverseScrollWheelZoomDirection(v bool)`

SetReverseScrollWheelZoomDirection sets ReverseScrollWheelZoomDirection field to given value.

### HasReverseScrollWheelZoomDirection

`func (o *BTUserSettingsInfo) HasReverseScrollWheelZoomDirection() bool`

HasReverseScrollWheelZoomDirection returns a boolean if a field has been set.

### GetSelectItemViewStateInfos

`func (o *BTUserSettingsInfo) GetSelectItemViewStateInfos() []BTSelectItemViewStateInfo`

GetSelectItemViewStateInfos returns the SelectItemViewStateInfos field if non-nil, zero value otherwise.

### GetSelectItemViewStateInfosOk

`func (o *BTUserSettingsInfo) GetSelectItemViewStateInfosOk() (*[]BTSelectItemViewStateInfo, bool)`

GetSelectItemViewStateInfosOk returns a tuple with the SelectItemViewStateInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectItemViewStateInfos

`func (o *BTUserSettingsInfo) SetSelectItemViewStateInfos(v []BTSelectItemViewStateInfo)`

SetSelectItemViewStateInfos sets SelectItemViewStateInfos field to given value.

### HasSelectItemViewStateInfos

`func (o *BTUserSettingsInfo) HasSelectItemViewStateInfos() bool`

HasSelectItemViewStateInfos returns a boolean if a field has been set.

### GetStartupPage

`func (o *BTUserSettingsInfo) GetStartupPage() int32`

GetStartupPage returns the StartupPage field if non-nil, zero value otherwise.

### GetStartupPageOk

`func (o *BTUserSettingsInfo) GetStartupPageOk() (*int32, bool)`

GetStartupPageOk returns a tuple with the StartupPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupPage

`func (o *BTUserSettingsInfo) SetStartupPage(v int32)`

SetStartupPage sets StartupPage field to given value.

### HasStartupPage

`func (o *BTUserSettingsInfo) HasStartupPage() bool`

HasStartupPage returns a boolean if a field has been set.

### GetSubstituteApprovers

`func (o *BTUserSettingsInfo) GetSubstituteApprovers() []BTSubstituteApproverInfo`

GetSubstituteApprovers returns the SubstituteApprovers field if non-nil, zero value otherwise.

### GetSubstituteApproversOk

`func (o *BTUserSettingsInfo) GetSubstituteApproversOk() (*[]BTSubstituteApproverInfo, bool)`

GetSubstituteApproversOk returns a tuple with the SubstituteApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstituteApprovers

`func (o *BTUserSettingsInfo) SetSubstituteApprovers(v []BTSubstituteApproverInfo)`

SetSubstituteApprovers sets SubstituteApprovers field to given value.

### HasSubstituteApprovers

`func (o *BTUserSettingsInfo) HasSubstituteApprovers() bool`

HasSubstituteApprovers returns a boolean if a field has been set.

### GetTheme

`func (o *BTUserSettingsInfo) GetTheme() int32`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *BTUserSettingsInfo) GetThemeOk() (*int32, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *BTUserSettingsInfo) SetTheme(v int32)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *BTUserSettingsInfo) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetUnitsDisplayPrecision

`func (o *BTUserSettingsInfo) GetUnitsDisplayPrecision() map[string]int32`

GetUnitsDisplayPrecision returns the UnitsDisplayPrecision field if non-nil, zero value otherwise.

### GetUnitsDisplayPrecisionOk

`func (o *BTUserSettingsInfo) GetUnitsDisplayPrecisionOk() (*map[string]int32, bool)`

GetUnitsDisplayPrecisionOk returns a tuple with the UnitsDisplayPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsDisplayPrecision

`func (o *BTUserSettingsInfo) SetUnitsDisplayPrecision(v map[string]int32)`

SetUnitsDisplayPrecision sets UnitsDisplayPrecision field to given value.

### HasUnitsDisplayPrecision

`func (o *BTUserSettingsInfo) HasUnitsDisplayPrecision() bool`

HasUnitsDisplayPrecision returns a boolean if a field has been set.

### GetUnitsMaximumDisplayPrecision

`func (o *BTUserSettingsInfo) GetUnitsMaximumDisplayPrecision() BTUnitsMaximumDisplayPrecisionInfo`

GetUnitsMaximumDisplayPrecision returns the UnitsMaximumDisplayPrecision field if non-nil, zero value otherwise.

### GetUnitsMaximumDisplayPrecisionOk

`func (o *BTUserSettingsInfo) GetUnitsMaximumDisplayPrecisionOk() (*BTUnitsMaximumDisplayPrecisionInfo, bool)`

GetUnitsMaximumDisplayPrecisionOk returns a tuple with the UnitsMaximumDisplayPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsMaximumDisplayPrecision

`func (o *BTUserSettingsInfo) SetUnitsMaximumDisplayPrecision(v BTUnitsMaximumDisplayPrecisionInfo)`

SetUnitsMaximumDisplayPrecision sets UnitsMaximumDisplayPrecision field to given value.

### HasUnitsMaximumDisplayPrecision

`func (o *BTUserSettingsInfo) HasUnitsMaximumDisplayPrecision() bool`

HasUnitsMaximumDisplayPrecision returns a boolean if a field has been set.

### GetUse24HourTime

`func (o *BTUserSettingsInfo) GetUse24HourTime() bool`

GetUse24HourTime returns the Use24HourTime field if non-nil, zero value otherwise.

### GetUse24HourTimeOk

`func (o *BTUserSettingsInfo) GetUse24HourTimeOk() (*bool, bool)`

GetUse24HourTimeOk returns a tuple with the Use24HourTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse24HourTime

`func (o *BTUserSettingsInfo) SetUse24HourTime(v bool)`

SetUse24HourTime sets Use24HourTime field to given value.

### HasUse24HourTime

`func (o *BTUserSettingsInfo) HasUse24HourTime() bool`

HasUse24HourTime returns a boolean if a field has been set.

### GetUseDecimalComma

`func (o *BTUserSettingsInfo) GetUseDecimalComma() bool`

GetUseDecimalComma returns the UseDecimalComma field if non-nil, zero value otherwise.

### GetUseDecimalCommaOk

`func (o *BTUserSettingsInfo) GetUseDecimalCommaOk() (*bool, bool)`

GetUseDecimalCommaOk returns a tuple with the UseDecimalComma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDecimalComma

`func (o *BTUserSettingsInfo) SetUseDecimalComma(v bool)`

SetUseDecimalComma sets UseDecimalComma field to given value.

### HasUseDecimalComma

`func (o *BTUserSettingsInfo) HasUseDecimalComma() bool`

HasUseDecimalComma returns a boolean if a field has been set.

### GetViewManipulationMouseKeyMapping

`func (o *BTUserSettingsInfo) GetViewManipulationMouseKeyMapping() BTViewManipulationMouseKeyMappingInfo`

GetViewManipulationMouseKeyMapping returns the ViewManipulationMouseKeyMapping field if non-nil, zero value otherwise.

### GetViewManipulationMouseKeyMappingOk

`func (o *BTUserSettingsInfo) GetViewManipulationMouseKeyMappingOk() (*BTViewManipulationMouseKeyMappingInfo, bool)`

GetViewManipulationMouseKeyMappingOk returns a tuple with the ViewManipulationMouseKeyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewManipulationMouseKeyMapping

`func (o *BTUserSettingsInfo) SetViewManipulationMouseKeyMapping(v BTViewManipulationMouseKeyMappingInfo)`

SetViewManipulationMouseKeyMapping sets ViewManipulationMouseKeyMapping field to given value.

### HasViewManipulationMouseKeyMapping

`func (o *BTUserSettingsInfo) HasViewManipulationMouseKeyMapping() bool`

HasViewManipulationMouseKeyMapping returns a boolean if a field has been set.

### GetViewMappingId

`func (o *BTUserSettingsInfo) GetViewMappingId() int32`

GetViewMappingId returns the ViewMappingId field if non-nil, zero value otherwise.

### GetViewMappingIdOk

`func (o *BTUserSettingsInfo) GetViewMappingIdOk() (*int32, bool)`

GetViewMappingIdOk returns a tuple with the ViewMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMappingId

`func (o *BTUserSettingsInfo) SetViewMappingId(v int32)`

SetViewMappingId sets ViewMappingId field to given value.

### HasViewMappingId

`func (o *BTUserSettingsInfo) HasViewMappingId() bool`

HasViewMappingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


