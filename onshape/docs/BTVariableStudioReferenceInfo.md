# BTVariableStudioReferenceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationIdToValue** | Pointer to [**map[string]BTOptionallyConfiguredValue**](BTOptionallyConfiguredValue.md) | Optional map of configuration parameter id to value | [optional] 
**EntireVariableStudio** | Pointer to **bool** | Whether all variables in the referenced variable studio are included | [optional] 
**ReferenceDocumentId** | Pointer to **string** | DocumentId of referenced variable studio, blank for intra-workspace references | [optional] 
**ReferenceElementId** | **string** | ElementId of referenced variable studio | 
**ReferenceVersionId** | Pointer to **string** | VersionId of referenced variable studio, blank for intra-workspace references | [optional] 
**VariableNames** | Pointer to **[]string** | Optional list of selected variables | [optional] 

## Methods

### NewBTVariableStudioReferenceInfo

`func NewBTVariableStudioReferenceInfo(referenceElementId string, ) *BTVariableStudioReferenceInfo`

NewBTVariableStudioReferenceInfo instantiates a new BTVariableStudioReferenceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVariableStudioReferenceInfoWithDefaults

`func NewBTVariableStudioReferenceInfoWithDefaults() *BTVariableStudioReferenceInfo`

NewBTVariableStudioReferenceInfoWithDefaults instantiates a new BTVariableStudioReferenceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationIdToValue

`func (o *BTVariableStudioReferenceInfo) GetConfigurationIdToValue() map[string]BTOptionallyConfiguredValue`

GetConfigurationIdToValue returns the ConfigurationIdToValue field if non-nil, zero value otherwise.

### GetConfigurationIdToValueOk

`func (o *BTVariableStudioReferenceInfo) GetConfigurationIdToValueOk() (*map[string]BTOptionallyConfiguredValue, bool)`

GetConfigurationIdToValueOk returns a tuple with the ConfigurationIdToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIdToValue

`func (o *BTVariableStudioReferenceInfo) SetConfigurationIdToValue(v map[string]BTOptionallyConfiguredValue)`

SetConfigurationIdToValue sets ConfigurationIdToValue field to given value.

### HasConfigurationIdToValue

`func (o *BTVariableStudioReferenceInfo) HasConfigurationIdToValue() bool`

HasConfigurationIdToValue returns a boolean if a field has been set.

### GetEntireVariableStudio

`func (o *BTVariableStudioReferenceInfo) GetEntireVariableStudio() bool`

GetEntireVariableStudio returns the EntireVariableStudio field if non-nil, zero value otherwise.

### GetEntireVariableStudioOk

`func (o *BTVariableStudioReferenceInfo) GetEntireVariableStudioOk() (*bool, bool)`

GetEntireVariableStudioOk returns a tuple with the EntireVariableStudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntireVariableStudio

`func (o *BTVariableStudioReferenceInfo) SetEntireVariableStudio(v bool)`

SetEntireVariableStudio sets EntireVariableStudio field to given value.

### HasEntireVariableStudio

`func (o *BTVariableStudioReferenceInfo) HasEntireVariableStudio() bool`

HasEntireVariableStudio returns a boolean if a field has been set.

### GetReferenceDocumentId

`func (o *BTVariableStudioReferenceInfo) GetReferenceDocumentId() string`

GetReferenceDocumentId returns the ReferenceDocumentId field if non-nil, zero value otherwise.

### GetReferenceDocumentIdOk

`func (o *BTVariableStudioReferenceInfo) GetReferenceDocumentIdOk() (*string, bool)`

GetReferenceDocumentIdOk returns a tuple with the ReferenceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceDocumentId

`func (o *BTVariableStudioReferenceInfo) SetReferenceDocumentId(v string)`

SetReferenceDocumentId sets ReferenceDocumentId field to given value.

### HasReferenceDocumentId

`func (o *BTVariableStudioReferenceInfo) HasReferenceDocumentId() bool`

HasReferenceDocumentId returns a boolean if a field has been set.

### GetReferenceElementId

`func (o *BTVariableStudioReferenceInfo) GetReferenceElementId() string`

GetReferenceElementId returns the ReferenceElementId field if non-nil, zero value otherwise.

### GetReferenceElementIdOk

`func (o *BTVariableStudioReferenceInfo) GetReferenceElementIdOk() (*string, bool)`

GetReferenceElementIdOk returns a tuple with the ReferenceElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceElementId

`func (o *BTVariableStudioReferenceInfo) SetReferenceElementId(v string)`

SetReferenceElementId sets ReferenceElementId field to given value.


### GetReferenceVersionId

`func (o *BTVariableStudioReferenceInfo) GetReferenceVersionId() string`

GetReferenceVersionId returns the ReferenceVersionId field if non-nil, zero value otherwise.

### GetReferenceVersionIdOk

`func (o *BTVariableStudioReferenceInfo) GetReferenceVersionIdOk() (*string, bool)`

GetReferenceVersionIdOk returns a tuple with the ReferenceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceVersionId

`func (o *BTVariableStudioReferenceInfo) SetReferenceVersionId(v string)`

SetReferenceVersionId sets ReferenceVersionId field to given value.

### HasReferenceVersionId

`func (o *BTVariableStudioReferenceInfo) HasReferenceVersionId() bool`

HasReferenceVersionId returns a boolean if a field has been set.

### GetVariableNames

`func (o *BTVariableStudioReferenceInfo) GetVariableNames() []string`

GetVariableNames returns the VariableNames field if non-nil, zero value otherwise.

### GetVariableNamesOk

`func (o *BTVariableStudioReferenceInfo) GetVariableNamesOk() (*[]string, bool)`

GetVariableNamesOk returns a tuple with the VariableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNames

`func (o *BTVariableStudioReferenceInfo) SetVariableNames(v []string)`

SetVariableNames sets VariableNames field to given value.

### HasVariableNames

`func (o *BTVariableStudioReferenceInfo) HasVariableNames() bool`

HasVariableNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


