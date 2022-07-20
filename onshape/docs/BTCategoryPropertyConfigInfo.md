# BTCategoryPropertyConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedPropertyConfig** | Pointer to [**BTComputedPropertyConfig**](BTComputedPropertyConfig.md) |  | [optional] 
**ComputedPropertyFunctionName** | Pointer to **string** |  | [optional] 
**ComputedPropertyFunctionNamespace** | Pointer to **string** |  | [optional] 
**ComputedPropertyFunctionURL** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnumValues** | Pointer to [**[]BTMetadataEnumValue**](BTMetadataEnumValue.md) |  | [optional] 
**MaxCount** | Pointer to **int32** |  | [optional] 
**MaxDate** | Pointer to **JSONTime** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**MaxValue** | Pointer to **float64** |  | [optional] 
**MinCount** | Pointer to **int32** |  | [optional] 
**MinDate** | Pointer to **JSONTime** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MinValue** | Pointer to **float64** |  | [optional] 
**Multiline** | Pointer to **bool** |  | [optional] 
**Multivalued** | Pointer to **bool** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**PublishState** | Pointer to **int32** |  | [optional] 
**QuantityType** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTCategoryPropertyConfigInfo

`func NewBTCategoryPropertyConfigInfo() *BTCategoryPropertyConfigInfo`

NewBTCategoryPropertyConfigInfo instantiates a new BTCategoryPropertyConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCategoryPropertyConfigInfoWithDefaults

`func NewBTCategoryPropertyConfigInfoWithDefaults() *BTCategoryPropertyConfigInfo`

NewBTCategoryPropertyConfigInfoWithDefaults instantiates a new BTCategoryPropertyConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedPropertyConfig

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyConfig() BTComputedPropertyConfig`

GetComputedPropertyConfig returns the ComputedPropertyConfig field if non-nil, zero value otherwise.

### GetComputedPropertyConfigOk

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyConfigOk() (*BTComputedPropertyConfig, bool)`

GetComputedPropertyConfigOk returns a tuple with the ComputedPropertyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPropertyConfig

`func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyConfig(v BTComputedPropertyConfig)`

SetComputedPropertyConfig sets ComputedPropertyConfig field to given value.

### HasComputedPropertyConfig

`func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyConfig() bool`

HasComputedPropertyConfig returns a boolean if a field has been set.

### GetComputedPropertyFunctionName

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionName() string`

GetComputedPropertyFunctionName returns the ComputedPropertyFunctionName field if non-nil, zero value otherwise.

### GetComputedPropertyFunctionNameOk

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionNameOk() (*string, bool)`

GetComputedPropertyFunctionNameOk returns a tuple with the ComputedPropertyFunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPropertyFunctionName

`func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyFunctionName(v string)`

SetComputedPropertyFunctionName sets ComputedPropertyFunctionName field to given value.

### HasComputedPropertyFunctionName

`func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyFunctionName() bool`

HasComputedPropertyFunctionName returns a boolean if a field has been set.

### GetComputedPropertyFunctionNamespace

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionNamespace() string`

GetComputedPropertyFunctionNamespace returns the ComputedPropertyFunctionNamespace field if non-nil, zero value otherwise.

### GetComputedPropertyFunctionNamespaceOk

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionNamespaceOk() (*string, bool)`

GetComputedPropertyFunctionNamespaceOk returns a tuple with the ComputedPropertyFunctionNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPropertyFunctionNamespace

`func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyFunctionNamespace(v string)`

SetComputedPropertyFunctionNamespace sets ComputedPropertyFunctionNamespace field to given value.

### HasComputedPropertyFunctionNamespace

`func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyFunctionNamespace() bool`

HasComputedPropertyFunctionNamespace returns a boolean if a field has been set.

### GetComputedPropertyFunctionURL

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionURL() string`

GetComputedPropertyFunctionURL returns the ComputedPropertyFunctionURL field if non-nil, zero value otherwise.

### GetComputedPropertyFunctionURLOk

`func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionURLOk() (*string, bool)`

GetComputedPropertyFunctionURLOk returns a tuple with the ComputedPropertyFunctionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPropertyFunctionURL

`func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyFunctionURL(v string)`

SetComputedPropertyFunctionURL sets ComputedPropertyFunctionURL field to given value.

### HasComputedPropertyFunctionURL

`func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyFunctionURL() bool`

HasComputedPropertyFunctionURL returns a boolean if a field has been set.

### GetDefaultValue

`func (o *BTCategoryPropertyConfigInfo) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTCategoryPropertyConfigInfo) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTCategoryPropertyConfigInfo) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTCategoryPropertyConfigInfo) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDisplayName

`func (o *BTCategoryPropertyConfigInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BTCategoryPropertyConfigInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BTCategoryPropertyConfigInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BTCategoryPropertyConfigInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnumValues

`func (o *BTCategoryPropertyConfigInfo) GetEnumValues() []BTMetadataEnumValue`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *BTCategoryPropertyConfigInfo) GetEnumValuesOk() (*[]BTMetadataEnumValue, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *BTCategoryPropertyConfigInfo) SetEnumValues(v []BTMetadataEnumValue)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *BTCategoryPropertyConfigInfo) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetMaxCount

`func (o *BTCategoryPropertyConfigInfo) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *BTCategoryPropertyConfigInfo) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *BTCategoryPropertyConfigInfo) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *BTCategoryPropertyConfigInfo) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetMaxDate

`func (o *BTCategoryPropertyConfigInfo) GetMaxDate() JSONTime`

GetMaxDate returns the MaxDate field if non-nil, zero value otherwise.

### GetMaxDateOk

`func (o *BTCategoryPropertyConfigInfo) GetMaxDateOk() (*JSONTime, bool)`

GetMaxDateOk returns a tuple with the MaxDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDate

`func (o *BTCategoryPropertyConfigInfo) SetMaxDate(v JSONTime)`

SetMaxDate sets MaxDate field to given value.

### HasMaxDate

`func (o *BTCategoryPropertyConfigInfo) HasMaxDate() bool`

HasMaxDate returns a boolean if a field has been set.

### GetMaxLength

`func (o *BTCategoryPropertyConfigInfo) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *BTCategoryPropertyConfigInfo) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *BTCategoryPropertyConfigInfo) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *BTCategoryPropertyConfigInfo) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMaxValue

`func (o *BTCategoryPropertyConfigInfo) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *BTCategoryPropertyConfigInfo) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *BTCategoryPropertyConfigInfo) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *BTCategoryPropertyConfigInfo) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinCount

`func (o *BTCategoryPropertyConfigInfo) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *BTCategoryPropertyConfigInfo) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *BTCategoryPropertyConfigInfo) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *BTCategoryPropertyConfigInfo) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMinDate

`func (o *BTCategoryPropertyConfigInfo) GetMinDate() JSONTime`

GetMinDate returns the MinDate field if non-nil, zero value otherwise.

### GetMinDateOk

`func (o *BTCategoryPropertyConfigInfo) GetMinDateOk() (*JSONTime, bool)`

GetMinDateOk returns a tuple with the MinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDate

`func (o *BTCategoryPropertyConfigInfo) SetMinDate(v JSONTime)`

SetMinDate sets MinDate field to given value.

### HasMinDate

`func (o *BTCategoryPropertyConfigInfo) HasMinDate() bool`

HasMinDate returns a boolean if a field has been set.

### GetMinLength

`func (o *BTCategoryPropertyConfigInfo) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *BTCategoryPropertyConfigInfo) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *BTCategoryPropertyConfigInfo) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *BTCategoryPropertyConfigInfo) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMinValue

`func (o *BTCategoryPropertyConfigInfo) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *BTCategoryPropertyConfigInfo) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *BTCategoryPropertyConfigInfo) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *BTCategoryPropertyConfigInfo) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMultiline

`func (o *BTCategoryPropertyConfigInfo) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *BTCategoryPropertyConfigInfo) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *BTCategoryPropertyConfigInfo) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *BTCategoryPropertyConfigInfo) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetMultivalued

`func (o *BTCategoryPropertyConfigInfo) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *BTCategoryPropertyConfigInfo) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *BTCategoryPropertyConfigInfo) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *BTCategoryPropertyConfigInfo) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.

### GetPattern

`func (o *BTCategoryPropertyConfigInfo) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *BTCategoryPropertyConfigInfo) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *BTCategoryPropertyConfigInfo) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *BTCategoryPropertyConfigInfo) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPublishState

`func (o *BTCategoryPropertyConfigInfo) GetPublishState() int32`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *BTCategoryPropertyConfigInfo) GetPublishStateOk() (*int32, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *BTCategoryPropertyConfigInfo) SetPublishState(v int32)`

SetPublishState sets PublishState field to given value.

### HasPublishState

`func (o *BTCategoryPropertyConfigInfo) HasPublishState() bool`

HasPublishState returns a boolean if a field has been set.

### GetQuantityType

`func (o *BTCategoryPropertyConfigInfo) GetQuantityType() int32`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *BTCategoryPropertyConfigInfo) GetQuantityTypeOk() (*int32, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *BTCategoryPropertyConfigInfo) SetQuantityType(v int32)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *BTCategoryPropertyConfigInfo) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.

### GetRequired

`func (o *BTCategoryPropertyConfigInfo) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BTCategoryPropertyConfigInfo) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BTCategoryPropertyConfigInfo) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BTCategoryPropertyConfigInfo) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


