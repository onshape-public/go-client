# BTUserAdminInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | Pointer to [**BTDiscount**](BTDiscount.md) |  | [optional] 
**InvitationId** | Pointer to **string** |  | [optional] 
**InvitedByEmail** | Pointer to **string** |  | [optional] 
**InvitedDocumentId** | Pointer to **string** |  | [optional] 
**IsTrialRequest** | Pointer to **bool** |  | [optional] 
**PrivacyConsents** | Pointer to [**[]BTPrivacyConsentInfo**](BTPrivacyConsentInfo.md) |  | [optional] 
**UserMetrics** | Pointer to [**BTUserMetricsInfo**](BTUserMetricsInfo.md) |  | [optional] 

## Methods

### NewBTUserAdminInfo

`func NewBTUserAdminInfo() *BTUserAdminInfo`

NewBTUserAdminInfo instantiates a new BTUserAdminInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserAdminInfoWithDefaults

`func NewBTUserAdminInfoWithDefaults() *BTUserAdminInfo`

NewBTUserAdminInfoWithDefaults instantiates a new BTUserAdminInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *BTUserAdminInfo) GetDiscount() BTDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BTUserAdminInfo) GetDiscountOk() (*BTDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BTUserAdminInfo) SetDiscount(v BTDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BTUserAdminInfo) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetInvitationId

`func (o *BTUserAdminInfo) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *BTUserAdminInfo) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *BTUserAdminInfo) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *BTUserAdminInfo) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### GetInvitedByEmail

`func (o *BTUserAdminInfo) GetInvitedByEmail() string`

GetInvitedByEmail returns the InvitedByEmail field if non-nil, zero value otherwise.

### GetInvitedByEmailOk

`func (o *BTUserAdminInfo) GetInvitedByEmailOk() (*string, bool)`

GetInvitedByEmailOk returns a tuple with the InvitedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedByEmail

`func (o *BTUserAdminInfo) SetInvitedByEmail(v string)`

SetInvitedByEmail sets InvitedByEmail field to given value.

### HasInvitedByEmail

`func (o *BTUserAdminInfo) HasInvitedByEmail() bool`

HasInvitedByEmail returns a boolean if a field has been set.

### GetInvitedDocumentId

`func (o *BTUserAdminInfo) GetInvitedDocumentId() string`

GetInvitedDocumentId returns the InvitedDocumentId field if non-nil, zero value otherwise.

### GetInvitedDocumentIdOk

`func (o *BTUserAdminInfo) GetInvitedDocumentIdOk() (*string, bool)`

GetInvitedDocumentIdOk returns a tuple with the InvitedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedDocumentId

`func (o *BTUserAdminInfo) SetInvitedDocumentId(v string)`

SetInvitedDocumentId sets InvitedDocumentId field to given value.

### HasInvitedDocumentId

`func (o *BTUserAdminInfo) HasInvitedDocumentId() bool`

HasInvitedDocumentId returns a boolean if a field has been set.

### GetIsTrialRequest

`func (o *BTUserAdminInfo) GetIsTrialRequest() bool`

GetIsTrialRequest returns the IsTrialRequest field if non-nil, zero value otherwise.

### GetIsTrialRequestOk

`func (o *BTUserAdminInfo) GetIsTrialRequestOk() (*bool, bool)`

GetIsTrialRequestOk returns a tuple with the IsTrialRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialRequest

`func (o *BTUserAdminInfo) SetIsTrialRequest(v bool)`

SetIsTrialRequest sets IsTrialRequest field to given value.

### HasIsTrialRequest

`func (o *BTUserAdminInfo) HasIsTrialRequest() bool`

HasIsTrialRequest returns a boolean if a field has been set.

### GetPrivacyConsents

`func (o *BTUserAdminInfo) GetPrivacyConsents() []BTPrivacyConsentInfo`

GetPrivacyConsents returns the PrivacyConsents field if non-nil, zero value otherwise.

### GetPrivacyConsentsOk

`func (o *BTUserAdminInfo) GetPrivacyConsentsOk() (*[]BTPrivacyConsentInfo, bool)`

GetPrivacyConsentsOk returns a tuple with the PrivacyConsents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyConsents

`func (o *BTUserAdminInfo) SetPrivacyConsents(v []BTPrivacyConsentInfo)`

SetPrivacyConsents sets PrivacyConsents field to given value.

### HasPrivacyConsents

`func (o *BTUserAdminInfo) HasPrivacyConsents() bool`

HasPrivacyConsents returns a boolean if a field has been set.

### GetUserMetrics

`func (o *BTUserAdminInfo) GetUserMetrics() BTUserMetricsInfo`

GetUserMetrics returns the UserMetrics field if non-nil, zero value otherwise.

### GetUserMetricsOk

`func (o *BTUserAdminInfo) GetUserMetricsOk() (*BTUserMetricsInfo, bool)`

GetUserMetricsOk returns a tuple with the UserMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetrics

`func (o *BTUserAdminInfo) SetUserMetrics(v BTUserMetricsInfo)`

SetUserMetrics sets UserMetrics field to given value.

### HasUserMetrics

`func (o *BTUserAdminInfo) HasUserMetrics() bool`

HasUserMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


