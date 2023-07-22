# BTParameterSpec6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalLocalizedStrings** | Pointer to **int32** |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**ColumnName** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to [**BTMParameter1**](BTMParameter1.md) |  | [optional] 
**IconUri** | Pointer to **string** |  | [optional] 
**LocalizableName** | Pointer to **string** |  | [optional] 
**LocalizedName** | Pointer to **string** |  | [optional] 
**ParameterDescription** | Pointer to **string** |  | [optional] 
**ParameterId** | Pointer to **string** |  | [optional] 
**ParameterName** | Pointer to **string** |  | [optional] 
**QuantityType** | Pointer to [**GBTQuantityType**](GBTQuantityType.md) |  | [optional] 
**StringsToLocalize** | Pointer to **[]string** |  | [optional] 
**UiHint** | Pointer to **string** |  | [optional] 
**UiHints** | Pointer to [**[]GBTUIHint**](GBTUIHint.md) |  | [optional] 
**VisibilityCondition** | Pointer to [**BTParameterVisibilityCondition177**](BTParameterVisibilityCondition177.md) |  | [optional] 

## Methods

### NewBTParameterSpec6

`func NewBTParameterSpec6() *BTParameterSpec6`

NewBTParameterSpec6 instantiates a new BTParameterSpec6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTParameterSpec6WithDefaults

`func NewBTParameterSpec6WithDefaults() *BTParameterSpec6`

NewBTParameterSpec6WithDefaults instantiates a new BTParameterSpec6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalLocalizedStrings

`func (o *BTParameterSpec6) GetAdditionalLocalizedStrings() int32`

GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field if non-nil, zero value otherwise.

### GetAdditionalLocalizedStringsOk

`func (o *BTParameterSpec6) GetAdditionalLocalizedStringsOk() (*int32, bool)`

GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLocalizedStrings

`func (o *BTParameterSpec6) SetAdditionalLocalizedStrings(v int32)`

SetAdditionalLocalizedStrings sets AdditionalLocalizedStrings field to given value.

### HasAdditionalLocalizedStrings

`func (o *BTParameterSpec6) HasAdditionalLocalizedStrings() bool`

HasAdditionalLocalizedStrings returns a boolean if a field has been set.

### GetBtType

`func (o *BTParameterSpec6) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTParameterSpec6) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTParameterSpec6) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTParameterSpec6) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetColumnName

`func (o *BTParameterSpec6) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *BTParameterSpec6) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *BTParameterSpec6) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *BTParameterSpec6) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetDefaultValue

`func (o *BTParameterSpec6) GetDefaultValue() BTMParameter1`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTParameterSpec6) GetDefaultValueOk() (*BTMParameter1, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTParameterSpec6) SetDefaultValue(v BTMParameter1)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTParameterSpec6) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIconUri

`func (o *BTParameterSpec6) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *BTParameterSpec6) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *BTParameterSpec6) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *BTParameterSpec6) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.

### GetLocalizableName

`func (o *BTParameterSpec6) GetLocalizableName() string`

GetLocalizableName returns the LocalizableName field if non-nil, zero value otherwise.

### GetLocalizableNameOk

`func (o *BTParameterSpec6) GetLocalizableNameOk() (*string, bool)`

GetLocalizableNameOk returns a tuple with the LocalizableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizableName

`func (o *BTParameterSpec6) SetLocalizableName(v string)`

SetLocalizableName sets LocalizableName field to given value.

### HasLocalizableName

`func (o *BTParameterSpec6) HasLocalizableName() bool`

HasLocalizableName returns a boolean if a field has been set.

### GetLocalizedName

`func (o *BTParameterSpec6) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *BTParameterSpec6) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *BTParameterSpec6) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *BTParameterSpec6) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetParameterDescription

`func (o *BTParameterSpec6) GetParameterDescription() string`

GetParameterDescription returns the ParameterDescription field if non-nil, zero value otherwise.

### GetParameterDescriptionOk

`func (o *BTParameterSpec6) GetParameterDescriptionOk() (*string, bool)`

GetParameterDescriptionOk returns a tuple with the ParameterDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterDescription

`func (o *BTParameterSpec6) SetParameterDescription(v string)`

SetParameterDescription sets ParameterDescription field to given value.

### HasParameterDescription

`func (o *BTParameterSpec6) HasParameterDescription() bool`

HasParameterDescription returns a boolean if a field has been set.

### GetParameterId

`func (o *BTParameterSpec6) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *BTParameterSpec6) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *BTParameterSpec6) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.

### HasParameterId

`func (o *BTParameterSpec6) HasParameterId() bool`

HasParameterId returns a boolean if a field has been set.

### GetParameterName

`func (o *BTParameterSpec6) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *BTParameterSpec6) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *BTParameterSpec6) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *BTParameterSpec6) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.

### GetQuantityType

`func (o *BTParameterSpec6) GetQuantityType() GBTQuantityType`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *BTParameterSpec6) GetQuantityTypeOk() (*GBTQuantityType, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *BTParameterSpec6) SetQuantityType(v GBTQuantityType)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *BTParameterSpec6) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.

### GetStringsToLocalize

`func (o *BTParameterSpec6) GetStringsToLocalize() []string`

GetStringsToLocalize returns the StringsToLocalize field if non-nil, zero value otherwise.

### GetStringsToLocalizeOk

`func (o *BTParameterSpec6) GetStringsToLocalizeOk() (*[]string, bool)`

GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringsToLocalize

`func (o *BTParameterSpec6) SetStringsToLocalize(v []string)`

SetStringsToLocalize sets StringsToLocalize field to given value.

### HasStringsToLocalize

`func (o *BTParameterSpec6) HasStringsToLocalize() bool`

HasStringsToLocalize returns a boolean if a field has been set.

### GetUiHint

`func (o *BTParameterSpec6) GetUiHint() string`

GetUiHint returns the UiHint field if non-nil, zero value otherwise.

### GetUiHintOk

`func (o *BTParameterSpec6) GetUiHintOk() (*string, bool)`

GetUiHintOk returns a tuple with the UiHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHint

`func (o *BTParameterSpec6) SetUiHint(v string)`

SetUiHint sets UiHint field to given value.

### HasUiHint

`func (o *BTParameterSpec6) HasUiHint() bool`

HasUiHint returns a boolean if a field has been set.

### GetUiHints

`func (o *BTParameterSpec6) GetUiHints() []GBTUIHint`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTParameterSpec6) GetUiHintsOk() (*[]GBTUIHint, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTParameterSpec6) SetUiHints(v []GBTUIHint)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTParameterSpec6) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.

### GetVisibilityCondition

`func (o *BTParameterSpec6) GetVisibilityCondition() BTParameterVisibilityCondition177`

GetVisibilityCondition returns the VisibilityCondition field if non-nil, zero value otherwise.

### GetVisibilityConditionOk

`func (o *BTParameterSpec6) GetVisibilityConditionOk() (*BTParameterVisibilityCondition177, bool)`

GetVisibilityConditionOk returns a tuple with the VisibilityCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityCondition

`func (o *BTParameterSpec6) SetVisibilityCondition(v BTParameterVisibilityCondition177)`

SetVisibilityCondition sets VisibilityCondition field to given value.

### HasVisibilityCondition

`func (o *BTParameterSpec6) HasVisibilityCondition() bool`

HasVisibilityCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


