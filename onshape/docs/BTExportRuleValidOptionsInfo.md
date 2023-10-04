# BTExportRuleValidOptionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConventionPlaceholder** | Pointer to **string** |  | [optional] 
**HardcodedProperties** | Pointer to [**[]BTExportRuleHardcodedPropertyInfo**](BTExportRuleHardcodedPropertyInfo.md) |  | [optional] 
**PropertyContextDisplayMap** | Pointer to **map[string]string** |  | [optional] 
**ValidObjectTypes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBTExportRuleValidOptionsInfo

`func NewBTExportRuleValidOptionsInfo() *BTExportRuleValidOptionsInfo`

NewBTExportRuleValidOptionsInfo instantiates a new BTExportRuleValidOptionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTExportRuleValidOptionsInfoWithDefaults

`func NewBTExportRuleValidOptionsInfoWithDefaults() *BTExportRuleValidOptionsInfo`

NewBTExportRuleValidOptionsInfoWithDefaults instantiates a new BTExportRuleValidOptionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConventionPlaceholder

`func (o *BTExportRuleValidOptionsInfo) GetConventionPlaceholder() string`

GetConventionPlaceholder returns the ConventionPlaceholder field if non-nil, zero value otherwise.

### GetConventionPlaceholderOk

`func (o *BTExportRuleValidOptionsInfo) GetConventionPlaceholderOk() (*string, bool)`

GetConventionPlaceholderOk returns a tuple with the ConventionPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConventionPlaceholder

`func (o *BTExportRuleValidOptionsInfo) SetConventionPlaceholder(v string)`

SetConventionPlaceholder sets ConventionPlaceholder field to given value.

### HasConventionPlaceholder

`func (o *BTExportRuleValidOptionsInfo) HasConventionPlaceholder() bool`

HasConventionPlaceholder returns a boolean if a field has been set.

### GetHardcodedProperties

`func (o *BTExportRuleValidOptionsInfo) GetHardcodedProperties() []BTExportRuleHardcodedPropertyInfo`

GetHardcodedProperties returns the HardcodedProperties field if non-nil, zero value otherwise.

### GetHardcodedPropertiesOk

`func (o *BTExportRuleValidOptionsInfo) GetHardcodedPropertiesOk() (*[]BTExportRuleHardcodedPropertyInfo, bool)`

GetHardcodedPropertiesOk returns a tuple with the HardcodedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardcodedProperties

`func (o *BTExportRuleValidOptionsInfo) SetHardcodedProperties(v []BTExportRuleHardcodedPropertyInfo)`

SetHardcodedProperties sets HardcodedProperties field to given value.

### HasHardcodedProperties

`func (o *BTExportRuleValidOptionsInfo) HasHardcodedProperties() bool`

HasHardcodedProperties returns a boolean if a field has been set.

### GetPropertyContextDisplayMap

`func (o *BTExportRuleValidOptionsInfo) GetPropertyContextDisplayMap() map[string]string`

GetPropertyContextDisplayMap returns the PropertyContextDisplayMap field if non-nil, zero value otherwise.

### GetPropertyContextDisplayMapOk

`func (o *BTExportRuleValidOptionsInfo) GetPropertyContextDisplayMapOk() (*map[string]string, bool)`

GetPropertyContextDisplayMapOk returns a tuple with the PropertyContextDisplayMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyContextDisplayMap

`func (o *BTExportRuleValidOptionsInfo) SetPropertyContextDisplayMap(v map[string]string)`

SetPropertyContextDisplayMap sets PropertyContextDisplayMap field to given value.

### HasPropertyContextDisplayMap

`func (o *BTExportRuleValidOptionsInfo) HasPropertyContextDisplayMap() bool`

HasPropertyContextDisplayMap returns a boolean if a field has been set.

### GetValidObjectTypes

`func (o *BTExportRuleValidOptionsInfo) GetValidObjectTypes() []int32`

GetValidObjectTypes returns the ValidObjectTypes field if non-nil, zero value otherwise.

### GetValidObjectTypesOk

`func (o *BTExportRuleValidOptionsInfo) GetValidObjectTypesOk() (*[]int32, bool)`

GetValidObjectTypesOk returns a tuple with the ValidObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidObjectTypes

`func (o *BTExportRuleValidOptionsInfo) SetValidObjectTypes(v []int32)`

SetValidObjectTypes sets ValidObjectTypes field to given value.

### HasValidObjectTypes

`func (o *BTExportRuleValidOptionsInfo) HasValidObjectTypes() bool`

HasValidObjectTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


