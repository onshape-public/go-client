# BTPrivacyConsentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationsOptInDate** | Pointer to **JSONTime** |  | [optional] 
**CommunicationsOptOutDate** | Pointer to **JSONTime** |  | [optional] 
**CommunicationsStatus** | Pointer to **bool** |  | [optional] 
**ConsentVersion** | Pointer to **string** |  | [optional] 
**DataProcessingOptInDate** | Pointer to **JSONTime** |  | [optional] 
**DataProcessingOptOutDate** | Pointer to **JSONTime** |  | [optional] 
**DataProcessingStatus** | Pointer to **bool** |  | [optional] 
**EulaVersion** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPrivacyConsentInfo

`func NewBTPrivacyConsentInfo() *BTPrivacyConsentInfo`

NewBTPrivacyConsentInfo instantiates a new BTPrivacyConsentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPrivacyConsentInfoWithDefaults

`func NewBTPrivacyConsentInfoWithDefaults() *BTPrivacyConsentInfo`

NewBTPrivacyConsentInfoWithDefaults instantiates a new BTPrivacyConsentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationsOptInDate

`func (o *BTPrivacyConsentInfo) GetCommunicationsOptInDate() JSONTime`

GetCommunicationsOptInDate returns the CommunicationsOptInDate field if non-nil, zero value otherwise.

### GetCommunicationsOptInDateOk

`func (o *BTPrivacyConsentInfo) GetCommunicationsOptInDateOk() (*JSONTime, bool)`

GetCommunicationsOptInDateOk returns a tuple with the CommunicationsOptInDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsOptInDate

`func (o *BTPrivacyConsentInfo) SetCommunicationsOptInDate(v JSONTime)`

SetCommunicationsOptInDate sets CommunicationsOptInDate field to given value.

### HasCommunicationsOptInDate

`func (o *BTPrivacyConsentInfo) HasCommunicationsOptInDate() bool`

HasCommunicationsOptInDate returns a boolean if a field has been set.

### GetCommunicationsOptOutDate

`func (o *BTPrivacyConsentInfo) GetCommunicationsOptOutDate() JSONTime`

GetCommunicationsOptOutDate returns the CommunicationsOptOutDate field if non-nil, zero value otherwise.

### GetCommunicationsOptOutDateOk

`func (o *BTPrivacyConsentInfo) GetCommunicationsOptOutDateOk() (*JSONTime, bool)`

GetCommunicationsOptOutDateOk returns a tuple with the CommunicationsOptOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsOptOutDate

`func (o *BTPrivacyConsentInfo) SetCommunicationsOptOutDate(v JSONTime)`

SetCommunicationsOptOutDate sets CommunicationsOptOutDate field to given value.

### HasCommunicationsOptOutDate

`func (o *BTPrivacyConsentInfo) HasCommunicationsOptOutDate() bool`

HasCommunicationsOptOutDate returns a boolean if a field has been set.

### GetCommunicationsStatus

`func (o *BTPrivacyConsentInfo) GetCommunicationsStatus() bool`

GetCommunicationsStatus returns the CommunicationsStatus field if non-nil, zero value otherwise.

### GetCommunicationsStatusOk

`func (o *BTPrivacyConsentInfo) GetCommunicationsStatusOk() (*bool, bool)`

GetCommunicationsStatusOk returns a tuple with the CommunicationsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsStatus

`func (o *BTPrivacyConsentInfo) SetCommunicationsStatus(v bool)`

SetCommunicationsStatus sets CommunicationsStatus field to given value.

### HasCommunicationsStatus

`func (o *BTPrivacyConsentInfo) HasCommunicationsStatus() bool`

HasCommunicationsStatus returns a boolean if a field has been set.

### GetConsentVersion

`func (o *BTPrivacyConsentInfo) GetConsentVersion() string`

GetConsentVersion returns the ConsentVersion field if non-nil, zero value otherwise.

### GetConsentVersionOk

`func (o *BTPrivacyConsentInfo) GetConsentVersionOk() (*string, bool)`

GetConsentVersionOk returns a tuple with the ConsentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentVersion

`func (o *BTPrivacyConsentInfo) SetConsentVersion(v string)`

SetConsentVersion sets ConsentVersion field to given value.

### HasConsentVersion

`func (o *BTPrivacyConsentInfo) HasConsentVersion() bool`

HasConsentVersion returns a boolean if a field has been set.

### GetDataProcessingOptInDate

`func (o *BTPrivacyConsentInfo) GetDataProcessingOptInDate() JSONTime`

GetDataProcessingOptInDate returns the DataProcessingOptInDate field if non-nil, zero value otherwise.

### GetDataProcessingOptInDateOk

`func (o *BTPrivacyConsentInfo) GetDataProcessingOptInDateOk() (*JSONTime, bool)`

GetDataProcessingOptInDateOk returns a tuple with the DataProcessingOptInDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessingOptInDate

`func (o *BTPrivacyConsentInfo) SetDataProcessingOptInDate(v JSONTime)`

SetDataProcessingOptInDate sets DataProcessingOptInDate field to given value.

### HasDataProcessingOptInDate

`func (o *BTPrivacyConsentInfo) HasDataProcessingOptInDate() bool`

HasDataProcessingOptInDate returns a boolean if a field has been set.

### GetDataProcessingOptOutDate

`func (o *BTPrivacyConsentInfo) GetDataProcessingOptOutDate() JSONTime`

GetDataProcessingOptOutDate returns the DataProcessingOptOutDate field if non-nil, zero value otherwise.

### GetDataProcessingOptOutDateOk

`func (o *BTPrivacyConsentInfo) GetDataProcessingOptOutDateOk() (*JSONTime, bool)`

GetDataProcessingOptOutDateOk returns a tuple with the DataProcessingOptOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessingOptOutDate

`func (o *BTPrivacyConsentInfo) SetDataProcessingOptOutDate(v JSONTime)`

SetDataProcessingOptOutDate sets DataProcessingOptOutDate field to given value.

### HasDataProcessingOptOutDate

`func (o *BTPrivacyConsentInfo) HasDataProcessingOptOutDate() bool`

HasDataProcessingOptOutDate returns a boolean if a field has been set.

### GetDataProcessingStatus

`func (o *BTPrivacyConsentInfo) GetDataProcessingStatus() bool`

GetDataProcessingStatus returns the DataProcessingStatus field if non-nil, zero value otherwise.

### GetDataProcessingStatusOk

`func (o *BTPrivacyConsentInfo) GetDataProcessingStatusOk() (*bool, bool)`

GetDataProcessingStatusOk returns a tuple with the DataProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessingStatus

`func (o *BTPrivacyConsentInfo) SetDataProcessingStatus(v bool)`

SetDataProcessingStatus sets DataProcessingStatus field to given value.

### HasDataProcessingStatus

`func (o *BTPrivacyConsentInfo) HasDataProcessingStatus() bool`

HasDataProcessingStatus returns a boolean if a field has been set.

### GetEulaVersion

`func (o *BTPrivacyConsentInfo) GetEulaVersion() int64`

GetEulaVersion returns the EulaVersion field if non-nil, zero value otherwise.

### GetEulaVersionOk

`func (o *BTPrivacyConsentInfo) GetEulaVersionOk() (*int64, bool)`

GetEulaVersionOk returns a tuple with the EulaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaVersion

`func (o *BTPrivacyConsentInfo) SetEulaVersion(v int64)`

SetEulaVersion sets EulaVersion field to given value.

### HasEulaVersion

`func (o *BTPrivacyConsentInfo) HasEulaVersion() bool`

HasEulaVersion returns a boolean if a field has been set.

### GetUserId

`func (o *BTPrivacyConsentInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTPrivacyConsentInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTPrivacyConsentInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTPrivacyConsentInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


