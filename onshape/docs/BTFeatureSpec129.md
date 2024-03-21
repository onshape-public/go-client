# BTFeatureSpec129

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalLocalizedStrings** | Pointer to **int32** |  | [optional] 
**AllParameters** | Pointer to [**[]BTParameterSpec6**](BTParameterSpec6.md) |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ComputedPartPropertySpec** | Pointer to **bool** |  | [optional] 
**DescriptionImageUri** | Pointer to **string** |  | [optional] 
**EditingLogic** | Pointer to [**BTEditingLogic2350**](BTEditingLogic2350.md) |  | [optional] 
**FeatureNameTemplate** | Pointer to **string** |  | [optional] 
**FeatureType** | Pointer to **string** |  | [optional] 
**FeatureTypeDescription** | Pointer to **string** |  | [optional] 
**FeatureTypeName** | Pointer to **string** |  | [optional] 
**FilterSelectors** | Pointer to **[]string** |  | [optional] 
**FullFeatureType** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]BTParameterGroupSpec3469**](BTParameterGroupSpec3469.md) |  | [optional] 
**IconUri** | Pointer to **string** |  | [optional] 
**LanguageVersion** | Pointer to **int32** |  | [optional] 
**LinkedLocationName** | Pointer to **string** |  | [optional] 
**LocalizableName** | Pointer to **string** |  | [optional] 
**LocalizedName** | Pointer to **string** |  | [optional] 
**LocationInfos** | Pointer to [**[]BTLocationInfo226**](BTLocationInfo226.md) |  | [optional] 
**ManipulatorChangeFunction** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NamespaceIncludingEnums** | Pointer to **string** |  | [optional] 
**NamespaceTheSource** | Pointer to **bool** |  | [optional] 
**ParameterIdToSpec** | Pointer to [**map[string]BTParameterSpec6**](BTParameterSpec6.md) |  | [optional] 
**Parameters** | Pointer to [**[]BTParameterSpec6**](BTParameterSpec6.md) |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SourceLocation** | Pointer to [**BTLocationInfo226**](BTLocationInfo226.md) |  | [optional] 
**SourceMicroversionId** | Pointer to **string** |  | [optional] 
**StringsToLocalize** | Pointer to **[]string** |  | [optional] 
**TableSpec** | Pointer to **bool** |  | [optional] 
**TooltipTemplate** | Pointer to **string** |  | [optional] 
**UiHints** | Pointer to [**[]GBTUIHint**](GBTUIHint.md) |  | [optional] 

## Methods

### NewBTFeatureSpec129

`func NewBTFeatureSpec129() *BTFeatureSpec129`

NewBTFeatureSpec129 instantiates a new BTFeatureSpec129 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTFeatureSpec129WithDefaults

`func NewBTFeatureSpec129WithDefaults() *BTFeatureSpec129`

NewBTFeatureSpec129WithDefaults instantiates a new BTFeatureSpec129 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalLocalizedStrings

`func (o *BTFeatureSpec129) GetAdditionalLocalizedStrings() int32`

GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field if non-nil, zero value otherwise.

### GetAdditionalLocalizedStringsOk

`func (o *BTFeatureSpec129) GetAdditionalLocalizedStringsOk() (*int32, bool)`

GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLocalizedStrings

`func (o *BTFeatureSpec129) SetAdditionalLocalizedStrings(v int32)`

SetAdditionalLocalizedStrings sets AdditionalLocalizedStrings field to given value.

### HasAdditionalLocalizedStrings

`func (o *BTFeatureSpec129) HasAdditionalLocalizedStrings() bool`

HasAdditionalLocalizedStrings returns a boolean if a field has been set.

### GetAllParameters

`func (o *BTFeatureSpec129) GetAllParameters() []BTParameterSpec6`

GetAllParameters returns the AllParameters field if non-nil, zero value otherwise.

### GetAllParametersOk

`func (o *BTFeatureSpec129) GetAllParametersOk() (*[]BTParameterSpec6, bool)`

GetAllParametersOk returns a tuple with the AllParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllParameters

`func (o *BTFeatureSpec129) SetAllParameters(v []BTParameterSpec6)`

SetAllParameters sets AllParameters field to given value.

### HasAllParameters

`func (o *BTFeatureSpec129) HasAllParameters() bool`

HasAllParameters returns a boolean if a field has been set.

### GetBtType

`func (o *BTFeatureSpec129) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTFeatureSpec129) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTFeatureSpec129) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTFeatureSpec129) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetComputedPartPropertySpec

`func (o *BTFeatureSpec129) GetComputedPartPropertySpec() bool`

GetComputedPartPropertySpec returns the ComputedPartPropertySpec field if non-nil, zero value otherwise.

### GetComputedPartPropertySpecOk

`func (o *BTFeatureSpec129) GetComputedPartPropertySpecOk() (*bool, bool)`

GetComputedPartPropertySpecOk returns a tuple with the ComputedPartPropertySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPartPropertySpec

`func (o *BTFeatureSpec129) SetComputedPartPropertySpec(v bool)`

SetComputedPartPropertySpec sets ComputedPartPropertySpec field to given value.

### HasComputedPartPropertySpec

`func (o *BTFeatureSpec129) HasComputedPartPropertySpec() bool`

HasComputedPartPropertySpec returns a boolean if a field has been set.

### GetDescriptionImageUri

`func (o *BTFeatureSpec129) GetDescriptionImageUri() string`

GetDescriptionImageUri returns the DescriptionImageUri field if non-nil, zero value otherwise.

### GetDescriptionImageUriOk

`func (o *BTFeatureSpec129) GetDescriptionImageUriOk() (*string, bool)`

GetDescriptionImageUriOk returns a tuple with the DescriptionImageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionImageUri

`func (o *BTFeatureSpec129) SetDescriptionImageUri(v string)`

SetDescriptionImageUri sets DescriptionImageUri field to given value.

### HasDescriptionImageUri

`func (o *BTFeatureSpec129) HasDescriptionImageUri() bool`

HasDescriptionImageUri returns a boolean if a field has been set.

### GetEditingLogic

`func (o *BTFeatureSpec129) GetEditingLogic() BTEditingLogic2350`

GetEditingLogic returns the EditingLogic field if non-nil, zero value otherwise.

### GetEditingLogicOk

`func (o *BTFeatureSpec129) GetEditingLogicOk() (*BTEditingLogic2350, bool)`

GetEditingLogicOk returns a tuple with the EditingLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditingLogic

`func (o *BTFeatureSpec129) SetEditingLogic(v BTEditingLogic2350)`

SetEditingLogic sets EditingLogic field to given value.

### HasEditingLogic

`func (o *BTFeatureSpec129) HasEditingLogic() bool`

HasEditingLogic returns a boolean if a field has been set.

### GetFeatureNameTemplate

`func (o *BTFeatureSpec129) GetFeatureNameTemplate() string`

GetFeatureNameTemplate returns the FeatureNameTemplate field if non-nil, zero value otherwise.

### GetFeatureNameTemplateOk

`func (o *BTFeatureSpec129) GetFeatureNameTemplateOk() (*string, bool)`

GetFeatureNameTemplateOk returns a tuple with the FeatureNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureNameTemplate

`func (o *BTFeatureSpec129) SetFeatureNameTemplate(v string)`

SetFeatureNameTemplate sets FeatureNameTemplate field to given value.

### HasFeatureNameTemplate

`func (o *BTFeatureSpec129) HasFeatureNameTemplate() bool`

HasFeatureNameTemplate returns a boolean if a field has been set.

### GetFeatureType

`func (o *BTFeatureSpec129) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *BTFeatureSpec129) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *BTFeatureSpec129) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *BTFeatureSpec129) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFeatureTypeDescription

`func (o *BTFeatureSpec129) GetFeatureTypeDescription() string`

GetFeatureTypeDescription returns the FeatureTypeDescription field if non-nil, zero value otherwise.

### GetFeatureTypeDescriptionOk

`func (o *BTFeatureSpec129) GetFeatureTypeDescriptionOk() (*string, bool)`

GetFeatureTypeDescriptionOk returns a tuple with the FeatureTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTypeDescription

`func (o *BTFeatureSpec129) SetFeatureTypeDescription(v string)`

SetFeatureTypeDescription sets FeatureTypeDescription field to given value.

### HasFeatureTypeDescription

`func (o *BTFeatureSpec129) HasFeatureTypeDescription() bool`

HasFeatureTypeDescription returns a boolean if a field has been set.

### GetFeatureTypeName

`func (o *BTFeatureSpec129) GetFeatureTypeName() string`

GetFeatureTypeName returns the FeatureTypeName field if non-nil, zero value otherwise.

### GetFeatureTypeNameOk

`func (o *BTFeatureSpec129) GetFeatureTypeNameOk() (*string, bool)`

GetFeatureTypeNameOk returns a tuple with the FeatureTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTypeName

`func (o *BTFeatureSpec129) SetFeatureTypeName(v string)`

SetFeatureTypeName sets FeatureTypeName field to given value.

### HasFeatureTypeName

`func (o *BTFeatureSpec129) HasFeatureTypeName() bool`

HasFeatureTypeName returns a boolean if a field has been set.

### GetFilterSelectors

`func (o *BTFeatureSpec129) GetFilterSelectors() []string`

GetFilterSelectors returns the FilterSelectors field if non-nil, zero value otherwise.

### GetFilterSelectorsOk

`func (o *BTFeatureSpec129) GetFilterSelectorsOk() (*[]string, bool)`

GetFilterSelectorsOk returns a tuple with the FilterSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSelectors

`func (o *BTFeatureSpec129) SetFilterSelectors(v []string)`

SetFilterSelectors sets FilterSelectors field to given value.

### HasFilterSelectors

`func (o *BTFeatureSpec129) HasFilterSelectors() bool`

HasFilterSelectors returns a boolean if a field has been set.

### GetFullFeatureType

`func (o *BTFeatureSpec129) GetFullFeatureType() string`

GetFullFeatureType returns the FullFeatureType field if non-nil, zero value otherwise.

### GetFullFeatureTypeOk

`func (o *BTFeatureSpec129) GetFullFeatureTypeOk() (*string, bool)`

GetFullFeatureTypeOk returns a tuple with the FullFeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullFeatureType

`func (o *BTFeatureSpec129) SetFullFeatureType(v string)`

SetFullFeatureType sets FullFeatureType field to given value.

### HasFullFeatureType

`func (o *BTFeatureSpec129) HasFullFeatureType() bool`

HasFullFeatureType returns a boolean if a field has been set.

### GetGroups

`func (o *BTFeatureSpec129) GetGroups() []BTParameterGroupSpec3469`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BTFeatureSpec129) GetGroupsOk() (*[]BTParameterGroupSpec3469, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BTFeatureSpec129) SetGroups(v []BTParameterGroupSpec3469)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BTFeatureSpec129) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIconUri

`func (o *BTFeatureSpec129) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *BTFeatureSpec129) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *BTFeatureSpec129) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *BTFeatureSpec129) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.

### GetLanguageVersion

`func (o *BTFeatureSpec129) GetLanguageVersion() int32`

GetLanguageVersion returns the LanguageVersion field if non-nil, zero value otherwise.

### GetLanguageVersionOk

`func (o *BTFeatureSpec129) GetLanguageVersionOk() (*int32, bool)`

GetLanguageVersionOk returns a tuple with the LanguageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageVersion

`func (o *BTFeatureSpec129) SetLanguageVersion(v int32)`

SetLanguageVersion sets LanguageVersion field to given value.

### HasLanguageVersion

`func (o *BTFeatureSpec129) HasLanguageVersion() bool`

HasLanguageVersion returns a boolean if a field has been set.

### GetLinkedLocationName

`func (o *BTFeatureSpec129) GetLinkedLocationName() string`

GetLinkedLocationName returns the LinkedLocationName field if non-nil, zero value otherwise.

### GetLinkedLocationNameOk

`func (o *BTFeatureSpec129) GetLinkedLocationNameOk() (*string, bool)`

GetLinkedLocationNameOk returns a tuple with the LinkedLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedLocationName

`func (o *BTFeatureSpec129) SetLinkedLocationName(v string)`

SetLinkedLocationName sets LinkedLocationName field to given value.

### HasLinkedLocationName

`func (o *BTFeatureSpec129) HasLinkedLocationName() bool`

HasLinkedLocationName returns a boolean if a field has been set.

### GetLocalizableName

`func (o *BTFeatureSpec129) GetLocalizableName() string`

GetLocalizableName returns the LocalizableName field if non-nil, zero value otherwise.

### GetLocalizableNameOk

`func (o *BTFeatureSpec129) GetLocalizableNameOk() (*string, bool)`

GetLocalizableNameOk returns a tuple with the LocalizableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizableName

`func (o *BTFeatureSpec129) SetLocalizableName(v string)`

SetLocalizableName sets LocalizableName field to given value.

### HasLocalizableName

`func (o *BTFeatureSpec129) HasLocalizableName() bool`

HasLocalizableName returns a boolean if a field has been set.

### GetLocalizedName

`func (o *BTFeatureSpec129) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *BTFeatureSpec129) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *BTFeatureSpec129) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *BTFeatureSpec129) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetLocationInfos

`func (o *BTFeatureSpec129) GetLocationInfos() []BTLocationInfo226`

GetLocationInfos returns the LocationInfos field if non-nil, zero value otherwise.

### GetLocationInfosOk

`func (o *BTFeatureSpec129) GetLocationInfosOk() (*[]BTLocationInfo226, bool)`

GetLocationInfosOk returns a tuple with the LocationInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInfos

`func (o *BTFeatureSpec129) SetLocationInfos(v []BTLocationInfo226)`

SetLocationInfos sets LocationInfos field to given value.

### HasLocationInfos

`func (o *BTFeatureSpec129) HasLocationInfos() bool`

HasLocationInfos returns a boolean if a field has been set.

### GetManipulatorChangeFunction

`func (o *BTFeatureSpec129) GetManipulatorChangeFunction() string`

GetManipulatorChangeFunction returns the ManipulatorChangeFunction field if non-nil, zero value otherwise.

### GetManipulatorChangeFunctionOk

`func (o *BTFeatureSpec129) GetManipulatorChangeFunctionOk() (*string, bool)`

GetManipulatorChangeFunctionOk returns a tuple with the ManipulatorChangeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManipulatorChangeFunction

`func (o *BTFeatureSpec129) SetManipulatorChangeFunction(v string)`

SetManipulatorChangeFunction sets ManipulatorChangeFunction field to given value.

### HasManipulatorChangeFunction

`func (o *BTFeatureSpec129) HasManipulatorChangeFunction() bool`

HasManipulatorChangeFunction returns a boolean if a field has been set.

### GetNamespace

`func (o *BTFeatureSpec129) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BTFeatureSpec129) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BTFeatureSpec129) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BTFeatureSpec129) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceIncludingEnums

`func (o *BTFeatureSpec129) GetNamespaceIncludingEnums() string`

GetNamespaceIncludingEnums returns the NamespaceIncludingEnums field if non-nil, zero value otherwise.

### GetNamespaceIncludingEnumsOk

`func (o *BTFeatureSpec129) GetNamespaceIncludingEnumsOk() (*string, bool)`

GetNamespaceIncludingEnumsOk returns a tuple with the NamespaceIncludingEnums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceIncludingEnums

`func (o *BTFeatureSpec129) SetNamespaceIncludingEnums(v string)`

SetNamespaceIncludingEnums sets NamespaceIncludingEnums field to given value.

### HasNamespaceIncludingEnums

`func (o *BTFeatureSpec129) HasNamespaceIncludingEnums() bool`

HasNamespaceIncludingEnums returns a boolean if a field has been set.

### GetNamespaceTheSource

`func (o *BTFeatureSpec129) GetNamespaceTheSource() bool`

GetNamespaceTheSource returns the NamespaceTheSource field if non-nil, zero value otherwise.

### GetNamespaceTheSourceOk

`func (o *BTFeatureSpec129) GetNamespaceTheSourceOk() (*bool, bool)`

GetNamespaceTheSourceOk returns a tuple with the NamespaceTheSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceTheSource

`func (o *BTFeatureSpec129) SetNamespaceTheSource(v bool)`

SetNamespaceTheSource sets NamespaceTheSource field to given value.

### HasNamespaceTheSource

`func (o *BTFeatureSpec129) HasNamespaceTheSource() bool`

HasNamespaceTheSource returns a boolean if a field has been set.

### GetParameterIdToSpec

`func (o *BTFeatureSpec129) GetParameterIdToSpec() map[string]BTParameterSpec6`

GetParameterIdToSpec returns the ParameterIdToSpec field if non-nil, zero value otherwise.

### GetParameterIdToSpecOk

`func (o *BTFeatureSpec129) GetParameterIdToSpecOk() (*map[string]BTParameterSpec6, bool)`

GetParameterIdToSpecOk returns a tuple with the ParameterIdToSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterIdToSpec

`func (o *BTFeatureSpec129) SetParameterIdToSpec(v map[string]BTParameterSpec6)`

SetParameterIdToSpec sets ParameterIdToSpec field to given value.

### HasParameterIdToSpec

`func (o *BTFeatureSpec129) HasParameterIdToSpec() bool`

HasParameterIdToSpec returns a boolean if a field has been set.

### GetParameters

`func (o *BTFeatureSpec129) GetParameters() []BTParameterSpec6`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTFeatureSpec129) GetParametersOk() (*[]BTParameterSpec6, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTFeatureSpec129) SetParameters(v []BTParameterSpec6)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTFeatureSpec129) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSignature

`func (o *BTFeatureSpec129) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *BTFeatureSpec129) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *BTFeatureSpec129) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *BTFeatureSpec129) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSourceLocation

`func (o *BTFeatureSpec129) GetSourceLocation() BTLocationInfo226`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *BTFeatureSpec129) GetSourceLocationOk() (*BTLocationInfo226, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *BTFeatureSpec129) SetSourceLocation(v BTLocationInfo226)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *BTFeatureSpec129) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.

### GetSourceMicroversionId

`func (o *BTFeatureSpec129) GetSourceMicroversionId() string`

GetSourceMicroversionId returns the SourceMicroversionId field if non-nil, zero value otherwise.

### GetSourceMicroversionIdOk

`func (o *BTFeatureSpec129) GetSourceMicroversionIdOk() (*string, bool)`

GetSourceMicroversionIdOk returns a tuple with the SourceMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversionId

`func (o *BTFeatureSpec129) SetSourceMicroversionId(v string)`

SetSourceMicroversionId sets SourceMicroversionId field to given value.

### HasSourceMicroversionId

`func (o *BTFeatureSpec129) HasSourceMicroversionId() bool`

HasSourceMicroversionId returns a boolean if a field has been set.

### GetStringsToLocalize

`func (o *BTFeatureSpec129) GetStringsToLocalize() []string`

GetStringsToLocalize returns the StringsToLocalize field if non-nil, zero value otherwise.

### GetStringsToLocalizeOk

`func (o *BTFeatureSpec129) GetStringsToLocalizeOk() (*[]string, bool)`

GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringsToLocalize

`func (o *BTFeatureSpec129) SetStringsToLocalize(v []string)`

SetStringsToLocalize sets StringsToLocalize field to given value.

### HasStringsToLocalize

`func (o *BTFeatureSpec129) HasStringsToLocalize() bool`

HasStringsToLocalize returns a boolean if a field has been set.

### GetTableSpec

`func (o *BTFeatureSpec129) GetTableSpec() bool`

GetTableSpec returns the TableSpec field if non-nil, zero value otherwise.

### GetTableSpecOk

`func (o *BTFeatureSpec129) GetTableSpecOk() (*bool, bool)`

GetTableSpecOk returns a tuple with the TableSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableSpec

`func (o *BTFeatureSpec129) SetTableSpec(v bool)`

SetTableSpec sets TableSpec field to given value.

### HasTableSpec

`func (o *BTFeatureSpec129) HasTableSpec() bool`

HasTableSpec returns a boolean if a field has been set.

### GetTooltipTemplate

`func (o *BTFeatureSpec129) GetTooltipTemplate() string`

GetTooltipTemplate returns the TooltipTemplate field if non-nil, zero value otherwise.

### GetTooltipTemplateOk

`func (o *BTFeatureSpec129) GetTooltipTemplateOk() (*string, bool)`

GetTooltipTemplateOk returns a tuple with the TooltipTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipTemplate

`func (o *BTFeatureSpec129) SetTooltipTemplate(v string)`

SetTooltipTemplate sets TooltipTemplate field to given value.

### HasTooltipTemplate

`func (o *BTFeatureSpec129) HasTooltipTemplate() bool`

HasTooltipTemplate returns a boolean if a field has been set.

### GetUiHints

`func (o *BTFeatureSpec129) GetUiHints() []GBTUIHint`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTFeatureSpec129) GetUiHintsOk() (*[]GBTUIHint, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTFeatureSpec129) SetUiHints(v []GBTUIHint)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTFeatureSpec129) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


