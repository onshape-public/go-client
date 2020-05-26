# BTUserAppSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**[]BTSettingInfo**](BTSettingInfo.md) |  | [optional] 

## Methods

### NewBTUserAppSettingsInfo

`func NewBTUserAppSettingsInfo() *BTUserAppSettingsInfo`

NewBTUserAppSettingsInfo instantiates a new BTUserAppSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserAppSettingsInfoWithDefaults

`func NewBTUserAppSettingsInfoWithDefaults() *BTUserAppSettingsInfo`

NewBTUserAppSettingsInfoWithDefaults instantiates a new BTUserAppSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *BTUserAppSettingsInfo) GetSettings() []BTSettingInfo`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BTUserAppSettingsInfo) GetSettingsOk() (*[]BTSettingInfo, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BTUserAppSettingsInfo) SetSettings(v []BTSettingInfo)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BTUserAppSettingsInfo) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


