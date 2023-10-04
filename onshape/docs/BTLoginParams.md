# BTLoginParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EnableTotp** | Pointer to **bool** |  | [optional] 
**IsRecaptchaV3** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**RandomToken** | Pointer to **string** |  | [optional] 
**RecaptchaToken** | Pointer to **string** |  | [optional] 
**RememberTotp** | Pointer to **bool** |  | [optional] 
**RendererPerformanceMeasurement** | Pointer to [**BTWebRendererPerformanceMeasurementParams**](BTWebRendererPerformanceMeasurementParams.md) |  | [optional] 
**Totp** | Pointer to **string** |  | [optional] 
**WebClientCapabilities** | Pointer to [**BTWebClientCapabilitiesParams**](BTWebClientCapabilitiesParams.md) |  | [optional] 

## Methods

### NewBTLoginParams

`func NewBTLoginParams() *BTLoginParams`

NewBTLoginParams instantiates a new BTLoginParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTLoginParamsWithDefaults

`func NewBTLoginParamsWithDefaults() *BTLoginParams`

NewBTLoginParamsWithDefaults instantiates a new BTLoginParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *BTLoginParams) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BTLoginParams) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BTLoginParams) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BTLoginParams) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEmail

`func (o *BTLoginParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BTLoginParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BTLoginParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BTLoginParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableTotp

`func (o *BTLoginParams) GetEnableTotp() bool`

GetEnableTotp returns the EnableTotp field if non-nil, zero value otherwise.

### GetEnableTotpOk

`func (o *BTLoginParams) GetEnableTotpOk() (*bool, bool)`

GetEnableTotpOk returns a tuple with the EnableTotp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTotp

`func (o *BTLoginParams) SetEnableTotp(v bool)`

SetEnableTotp sets EnableTotp field to given value.

### HasEnableTotp

`func (o *BTLoginParams) HasEnableTotp() bool`

HasEnableTotp returns a boolean if a field has been set.

### GetIsRecaptchaV3

`func (o *BTLoginParams) GetIsRecaptchaV3() bool`

GetIsRecaptchaV3 returns the IsRecaptchaV3 field if non-nil, zero value otherwise.

### GetIsRecaptchaV3Ok

`func (o *BTLoginParams) GetIsRecaptchaV3Ok() (*bool, bool)`

GetIsRecaptchaV3Ok returns a tuple with the IsRecaptchaV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecaptchaV3

`func (o *BTLoginParams) SetIsRecaptchaV3(v bool)`

SetIsRecaptchaV3 sets IsRecaptchaV3 field to given value.

### HasIsRecaptchaV3

`func (o *BTLoginParams) HasIsRecaptchaV3() bool`

HasIsRecaptchaV3 returns a boolean if a field has been set.

### GetPassword

`func (o *BTLoginParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BTLoginParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BTLoginParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BTLoginParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRandomToken

`func (o *BTLoginParams) GetRandomToken() string`

GetRandomToken returns the RandomToken field if non-nil, zero value otherwise.

### GetRandomTokenOk

`func (o *BTLoginParams) GetRandomTokenOk() (*string, bool)`

GetRandomTokenOk returns a tuple with the RandomToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomToken

`func (o *BTLoginParams) SetRandomToken(v string)`

SetRandomToken sets RandomToken field to given value.

### HasRandomToken

`func (o *BTLoginParams) HasRandomToken() bool`

HasRandomToken returns a boolean if a field has been set.

### GetRecaptchaToken

`func (o *BTLoginParams) GetRecaptchaToken() string`

GetRecaptchaToken returns the RecaptchaToken field if non-nil, zero value otherwise.

### GetRecaptchaTokenOk

`func (o *BTLoginParams) GetRecaptchaTokenOk() (*string, bool)`

GetRecaptchaTokenOk returns a tuple with the RecaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaToken

`func (o *BTLoginParams) SetRecaptchaToken(v string)`

SetRecaptchaToken sets RecaptchaToken field to given value.

### HasRecaptchaToken

`func (o *BTLoginParams) HasRecaptchaToken() bool`

HasRecaptchaToken returns a boolean if a field has been set.

### GetRememberTotp

`func (o *BTLoginParams) GetRememberTotp() bool`

GetRememberTotp returns the RememberTotp field if non-nil, zero value otherwise.

### GetRememberTotpOk

`func (o *BTLoginParams) GetRememberTotpOk() (*bool, bool)`

GetRememberTotpOk returns a tuple with the RememberTotp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberTotp

`func (o *BTLoginParams) SetRememberTotp(v bool)`

SetRememberTotp sets RememberTotp field to given value.

### HasRememberTotp

`func (o *BTLoginParams) HasRememberTotp() bool`

HasRememberTotp returns a boolean if a field has been set.

### GetRendererPerformanceMeasurement

`func (o *BTLoginParams) GetRendererPerformanceMeasurement() BTWebRendererPerformanceMeasurementParams`

GetRendererPerformanceMeasurement returns the RendererPerformanceMeasurement field if non-nil, zero value otherwise.

### GetRendererPerformanceMeasurementOk

`func (o *BTLoginParams) GetRendererPerformanceMeasurementOk() (*BTWebRendererPerformanceMeasurementParams, bool)`

GetRendererPerformanceMeasurementOk returns a tuple with the RendererPerformanceMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendererPerformanceMeasurement

`func (o *BTLoginParams) SetRendererPerformanceMeasurement(v BTWebRendererPerformanceMeasurementParams)`

SetRendererPerformanceMeasurement sets RendererPerformanceMeasurement field to given value.

### HasRendererPerformanceMeasurement

`func (o *BTLoginParams) HasRendererPerformanceMeasurement() bool`

HasRendererPerformanceMeasurement returns a boolean if a field has been set.

### GetTotp

`func (o *BTLoginParams) GetTotp() string`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *BTLoginParams) GetTotpOk() (*string, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *BTLoginParams) SetTotp(v string)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *BTLoginParams) HasTotp() bool`

HasTotp returns a boolean if a field has been set.

### GetWebClientCapabilities

`func (o *BTLoginParams) GetWebClientCapabilities() BTWebClientCapabilitiesParams`

GetWebClientCapabilities returns the WebClientCapabilities field if non-nil, zero value otherwise.

### GetWebClientCapabilitiesOk

`func (o *BTLoginParams) GetWebClientCapabilitiesOk() (*BTWebClientCapabilitiesParams, bool)`

GetWebClientCapabilitiesOk returns a tuple with the WebClientCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebClientCapabilities

`func (o *BTLoginParams) SetWebClientCapabilities(v BTWebClientCapabilitiesParams)`

SetWebClientCapabilities sets WebClientCapabilities field to given value.

### HasWebClientCapabilities

`func (o *BTLoginParams) HasWebClientCapabilities() bool`

HasWebClientCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


