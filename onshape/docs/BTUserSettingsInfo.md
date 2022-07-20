# BTUserSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonUnits** | Pointer to [**BTCommonUnitsInfo**](BTCommonUnitsInfo.md) |  | [optional] 
**CustomColors** | Pointer to **[]string** |  | [optional] 
**DefaultUnits** | Pointer to [**BTDefaultUnitsInfo**](BTDefaultUnitsInfo.md) |  | [optional] 
**DrawingBackgroundId** | Pointer to **int32** |  | [optional] 
**EnforceApplicationAcl** | Pointer to **bool** |  | [optional] 
**ExportDrawingOptions** | Pointer to **string** |  | [optional] 
**ExportSolidOptions** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImportOptions** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**MaterialLibrarySettings** | Pointer to [**BTMaterialLibrarySettingsInfo**](BTMaterialLibrarySettingsInfo.md) |  | [optional] 
**MiniToolbarSettings** | Pointer to **string** |  | [optional] 
**MouseActions** | Pointer to **string** |  | [optional] 
**ReverseScrollWheelZoomDirection** | Pointer to **bool** |  | [optional] 
**SelectItemViewStateInfos** | Pointer to [**[]BTSelectItemViewStateInfo**](BTSelectItemViewStateInfo.md) |  | [optional] 
**StartupPage** | Pointer to **int32** |  | [optional] 
**SubstituteApprovers** | Pointer to [**[]BTSubstituteApproverInfo**](BTSubstituteApproverInfo.md) |  | [optional] 
**UnitsDisplayPrecision** | Pointer to **map[string]int32** |  | [optional] 
**UnitsMaximumDisplayPrecision** | Pointer to [**BTUnitsMaximumDisplayPrecisionInfo**](BTUnitsMaximumDisplayPrecisionInfo.md) |  | [optional] 
**Use24HourTime** | Pointer to **bool** |  | [optional] 
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


