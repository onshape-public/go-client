# BTUserInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePlan** | Pointer to [**BTBillingPlanInfo**](BTBillingPlanInfo.md) |  | [optional] 
**ActivePlanId** | Pointer to **string** |  | [optional] 
**ActivePurchases** | Pointer to [**[]BTPurchaseInfo**](BTPurchaseInfo.md) |  | [optional] 
**ActiveTrialInfo** | Pointer to [**BTTrialInfo**](BTTrialInfo.md) |  | [optional] 
**B2cId** | Pointer to **string** |  | [optional] 
**BillingUpdateRequired** | Pointer to **bool** |  | [optional] 
**CadSystemAtSignup** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Credential** | Pointer to [**BTSessionCredentialInfo**](BTSessionCredentialInfo.md) |  | [optional] 
**DefaultCompanyName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceInfo** | Pointer to [**BTDeviceLoginSecretInfo**](BTDeviceLoginSecretInfo.md) |  | [optional] 
**Enterprise** | Pointer to **bool** |  | [optional] 
**EulaId** | Pointer to **string** |  | [optional] 
**EulaRequired** | Pointer to **bool** |  | [optional] 
**EvalCenter** | Pointer to **bool** |  | [optional] 
**ForumId** | Pointer to **string** |  | [optional] 
**LastTrialInfo** | Pointer to [**BTTrialInfo**](BTTrialInfo.md) |  | [optional] 
**NeedToShowNewWalkthrough** | Pointer to **bool** |  | [optional] 
**OwnPurchase** | Pointer to [**BTPurchaseInfo**](BTPurchaseInfo.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**ProductType** | Pointer to **[]string** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **int32** |  | [optional] 
**Roles** | Pointer to [**[]BTRole**](BTRole.md) |  | [optional] 
**StartupPage** | Pointer to **int32** |  | [optional] 
**SystemUser** | Pointer to **bool** |  | [optional] 
**TotpEnabled** | Pointer to **bool** |  | [optional] 
**TracingEnabled** | Pointer to **bool** |  | [optional] 
**TrialInfos** | Pointer to [**[]BTTrialInfo**](BTTrialInfo.md) |  | [optional] 

## Methods

### NewBTUserInfoAllOf

`func NewBTUserInfoAllOf() *BTUserInfoAllOf`

NewBTUserInfoAllOf instantiates a new BTUserInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserInfoAllOfWithDefaults

`func NewBTUserInfoAllOfWithDefaults() *BTUserInfoAllOf`

NewBTUserInfoAllOfWithDefaults instantiates a new BTUserInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePlan

`func (o *BTUserInfoAllOf) GetActivePlan() BTBillingPlanInfo`

GetActivePlan returns the ActivePlan field if non-nil, zero value otherwise.

### GetActivePlanOk

`func (o *BTUserInfoAllOf) GetActivePlanOk() (*BTBillingPlanInfo, bool)`

GetActivePlanOk returns a tuple with the ActivePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePlan

`func (o *BTUserInfoAllOf) SetActivePlan(v BTBillingPlanInfo)`

SetActivePlan sets ActivePlan field to given value.

### HasActivePlan

`func (o *BTUserInfoAllOf) HasActivePlan() bool`

HasActivePlan returns a boolean if a field has been set.

### GetActivePlanId

`func (o *BTUserInfoAllOf) GetActivePlanId() string`

GetActivePlanId returns the ActivePlanId field if non-nil, zero value otherwise.

### GetActivePlanIdOk

`func (o *BTUserInfoAllOf) GetActivePlanIdOk() (*string, bool)`

GetActivePlanIdOk returns a tuple with the ActivePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePlanId

`func (o *BTUserInfoAllOf) SetActivePlanId(v string)`

SetActivePlanId sets ActivePlanId field to given value.

### HasActivePlanId

`func (o *BTUserInfoAllOf) HasActivePlanId() bool`

HasActivePlanId returns a boolean if a field has been set.

### GetActivePurchases

`func (o *BTUserInfoAllOf) GetActivePurchases() []BTPurchaseInfo`

GetActivePurchases returns the ActivePurchases field if non-nil, zero value otherwise.

### GetActivePurchasesOk

`func (o *BTUserInfoAllOf) GetActivePurchasesOk() (*[]BTPurchaseInfo, bool)`

GetActivePurchasesOk returns a tuple with the ActivePurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePurchases

`func (o *BTUserInfoAllOf) SetActivePurchases(v []BTPurchaseInfo)`

SetActivePurchases sets ActivePurchases field to given value.

### HasActivePurchases

`func (o *BTUserInfoAllOf) HasActivePurchases() bool`

HasActivePurchases returns a boolean if a field has been set.

### GetActiveTrialInfo

`func (o *BTUserInfoAllOf) GetActiveTrialInfo() BTTrialInfo`

GetActiveTrialInfo returns the ActiveTrialInfo field if non-nil, zero value otherwise.

### GetActiveTrialInfoOk

`func (o *BTUserInfoAllOf) GetActiveTrialInfoOk() (*BTTrialInfo, bool)`

GetActiveTrialInfoOk returns a tuple with the ActiveTrialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTrialInfo

`func (o *BTUserInfoAllOf) SetActiveTrialInfo(v BTTrialInfo)`

SetActiveTrialInfo sets ActiveTrialInfo field to given value.

### HasActiveTrialInfo

`func (o *BTUserInfoAllOf) HasActiveTrialInfo() bool`

HasActiveTrialInfo returns a boolean if a field has been set.

### GetB2cId

`func (o *BTUserInfoAllOf) GetB2cId() string`

GetB2cId returns the B2cId field if non-nil, zero value otherwise.

### GetB2cIdOk

`func (o *BTUserInfoAllOf) GetB2cIdOk() (*string, bool)`

GetB2cIdOk returns a tuple with the B2cId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2cId

`func (o *BTUserInfoAllOf) SetB2cId(v string)`

SetB2cId sets B2cId field to given value.

### HasB2cId

`func (o *BTUserInfoAllOf) HasB2cId() bool`

HasB2cId returns a boolean if a field has been set.

### GetBillingUpdateRequired

`func (o *BTUserInfoAllOf) GetBillingUpdateRequired() bool`

GetBillingUpdateRequired returns the BillingUpdateRequired field if non-nil, zero value otherwise.

### GetBillingUpdateRequiredOk

`func (o *BTUserInfoAllOf) GetBillingUpdateRequiredOk() (*bool, bool)`

GetBillingUpdateRequiredOk returns a tuple with the BillingUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUpdateRequired

`func (o *BTUserInfoAllOf) SetBillingUpdateRequired(v bool)`

SetBillingUpdateRequired sets BillingUpdateRequired field to given value.

### HasBillingUpdateRequired

`func (o *BTUserInfoAllOf) HasBillingUpdateRequired() bool`

HasBillingUpdateRequired returns a boolean if a field has been set.

### GetCadSystemAtSignup

`func (o *BTUserInfoAllOf) GetCadSystemAtSignup() string`

GetCadSystemAtSignup returns the CadSystemAtSignup field if non-nil, zero value otherwise.

### GetCadSystemAtSignupOk

`func (o *BTUserInfoAllOf) GetCadSystemAtSignupOk() (*string, bool)`

GetCadSystemAtSignupOk returns a tuple with the CadSystemAtSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadSystemAtSignup

`func (o *BTUserInfoAllOf) SetCadSystemAtSignup(v string)`

SetCadSystemAtSignup sets CadSystemAtSignup field to given value.

### HasCadSystemAtSignup

`func (o *BTUserInfoAllOf) HasCadSystemAtSignup() bool`

HasCadSystemAtSignup returns a boolean if a field has been set.

### GetCountryCode

`func (o *BTUserInfoAllOf) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BTUserInfoAllOf) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BTUserInfoAllOf) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BTUserInfoAllOf) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTUserInfoAllOf) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTUserInfoAllOf) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTUserInfoAllOf) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTUserInfoAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredential

`func (o *BTUserInfoAllOf) GetCredential() BTSessionCredentialInfo`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *BTUserInfoAllOf) GetCredentialOk() (*BTSessionCredentialInfo, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *BTUserInfoAllOf) SetCredential(v BTSessionCredentialInfo)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *BTUserInfoAllOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDefaultCompanyName

`func (o *BTUserInfoAllOf) GetDefaultCompanyName() string`

GetDefaultCompanyName returns the DefaultCompanyName field if non-nil, zero value otherwise.

### GetDefaultCompanyNameOk

`func (o *BTUserInfoAllOf) GetDefaultCompanyNameOk() (*string, bool)`

GetDefaultCompanyNameOk returns a tuple with the DefaultCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCompanyName

`func (o *BTUserInfoAllOf) SetDefaultCompanyName(v string)`

SetDefaultCompanyName sets DefaultCompanyName field to given value.

### HasDefaultCompanyName

`func (o *BTUserInfoAllOf) HasDefaultCompanyName() bool`

HasDefaultCompanyName returns a boolean if a field has been set.

### GetDescription

`func (o *BTUserInfoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTUserInfoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTUserInfoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTUserInfoAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *BTUserInfoAllOf) GetDeviceInfo() BTDeviceLoginSecretInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *BTUserInfoAllOf) GetDeviceInfoOk() (*BTDeviceLoginSecretInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *BTUserInfoAllOf) SetDeviceInfo(v BTDeviceLoginSecretInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *BTUserInfoAllOf) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnterprise

`func (o *BTUserInfoAllOf) GetEnterprise() bool`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *BTUserInfoAllOf) GetEnterpriseOk() (*bool, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *BTUserInfoAllOf) SetEnterprise(v bool)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *BTUserInfoAllOf) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.

### GetEulaId

`func (o *BTUserInfoAllOf) GetEulaId() string`

GetEulaId returns the EulaId field if non-nil, zero value otherwise.

### GetEulaIdOk

`func (o *BTUserInfoAllOf) GetEulaIdOk() (*string, bool)`

GetEulaIdOk returns a tuple with the EulaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaId

`func (o *BTUserInfoAllOf) SetEulaId(v string)`

SetEulaId sets EulaId field to given value.

### HasEulaId

`func (o *BTUserInfoAllOf) HasEulaId() bool`

HasEulaId returns a boolean if a field has been set.

### GetEulaRequired

`func (o *BTUserInfoAllOf) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *BTUserInfoAllOf) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *BTUserInfoAllOf) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *BTUserInfoAllOf) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetEvalCenter

`func (o *BTUserInfoAllOf) GetEvalCenter() bool`

GetEvalCenter returns the EvalCenter field if non-nil, zero value otherwise.

### GetEvalCenterOk

`func (o *BTUserInfoAllOf) GetEvalCenterOk() (*bool, bool)`

GetEvalCenterOk returns a tuple with the EvalCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalCenter

`func (o *BTUserInfoAllOf) SetEvalCenter(v bool)`

SetEvalCenter sets EvalCenter field to given value.

### HasEvalCenter

`func (o *BTUserInfoAllOf) HasEvalCenter() bool`

HasEvalCenter returns a boolean if a field has been set.

### GetForumId

`func (o *BTUserInfoAllOf) GetForumId() string`

GetForumId returns the ForumId field if non-nil, zero value otherwise.

### GetForumIdOk

`func (o *BTUserInfoAllOf) GetForumIdOk() (*string, bool)`

GetForumIdOk returns a tuple with the ForumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumId

`func (o *BTUserInfoAllOf) SetForumId(v string)`

SetForumId sets ForumId field to given value.

### HasForumId

`func (o *BTUserInfoAllOf) HasForumId() bool`

HasForumId returns a boolean if a field has been set.

### GetLastTrialInfo

`func (o *BTUserInfoAllOf) GetLastTrialInfo() BTTrialInfo`

GetLastTrialInfo returns the LastTrialInfo field if non-nil, zero value otherwise.

### GetLastTrialInfoOk

`func (o *BTUserInfoAllOf) GetLastTrialInfoOk() (*BTTrialInfo, bool)`

GetLastTrialInfoOk returns a tuple with the LastTrialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialInfo

`func (o *BTUserInfoAllOf) SetLastTrialInfo(v BTTrialInfo)`

SetLastTrialInfo sets LastTrialInfo field to given value.

### HasLastTrialInfo

`func (o *BTUserInfoAllOf) HasLastTrialInfo() bool`

HasLastTrialInfo returns a boolean if a field has been set.

### GetNeedToShowNewWalkthrough

`func (o *BTUserInfoAllOf) GetNeedToShowNewWalkthrough() bool`

GetNeedToShowNewWalkthrough returns the NeedToShowNewWalkthrough field if non-nil, zero value otherwise.

### GetNeedToShowNewWalkthroughOk

`func (o *BTUserInfoAllOf) GetNeedToShowNewWalkthroughOk() (*bool, bool)`

GetNeedToShowNewWalkthroughOk returns a tuple with the NeedToShowNewWalkthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedToShowNewWalkthrough

`func (o *BTUserInfoAllOf) SetNeedToShowNewWalkthrough(v bool)`

SetNeedToShowNewWalkthrough sets NeedToShowNewWalkthrough field to given value.

### HasNeedToShowNewWalkthrough

`func (o *BTUserInfoAllOf) HasNeedToShowNewWalkthrough() bool`

HasNeedToShowNewWalkthrough returns a boolean if a field has been set.

### GetOwnPurchase

`func (o *BTUserInfoAllOf) GetOwnPurchase() BTPurchaseInfo`

GetOwnPurchase returns the OwnPurchase field if non-nil, zero value otherwise.

### GetOwnPurchaseOk

`func (o *BTUserInfoAllOf) GetOwnPurchaseOk() (*BTPurchaseInfo, bool)`

GetOwnPurchaseOk returns a tuple with the OwnPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnPurchase

`func (o *BTUserInfoAllOf) SetOwnPurchase(v BTPurchaseInfo)`

SetOwnPurchase sets OwnPurchase field to given value.

### HasOwnPurchase

`func (o *BTUserInfoAllOf) HasOwnPurchase() bool`

HasOwnPurchase returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BTUserInfoAllOf) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BTUserInfoAllOf) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BTUserInfoAllOf) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BTUserInfoAllOf) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetProductType

`func (o *BTUserInfoAllOf) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BTUserInfoAllOf) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BTUserInfoAllOf) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *BTUserInfoAllOf) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *BTUserInfoAllOf) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *BTUserInfoAllOf) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *BTUserInfoAllOf) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *BTUserInfoAllOf) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetRole

`func (o *BTUserInfoAllOf) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BTUserInfoAllOf) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BTUserInfoAllOf) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *BTUserInfoAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *BTUserInfoAllOf) GetRoles() []BTRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BTUserInfoAllOf) GetRolesOk() (*[]BTRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BTUserInfoAllOf) SetRoles(v []BTRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BTUserInfoAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStartupPage

`func (o *BTUserInfoAllOf) GetStartupPage() int32`

GetStartupPage returns the StartupPage field if non-nil, zero value otherwise.

### GetStartupPageOk

`func (o *BTUserInfoAllOf) GetStartupPageOk() (*int32, bool)`

GetStartupPageOk returns a tuple with the StartupPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupPage

`func (o *BTUserInfoAllOf) SetStartupPage(v int32)`

SetStartupPage sets StartupPage field to given value.

### HasStartupPage

`func (o *BTUserInfoAllOf) HasStartupPage() bool`

HasStartupPage returns a boolean if a field has been set.

### GetSystemUser

`func (o *BTUserInfoAllOf) GetSystemUser() bool`

GetSystemUser returns the SystemUser field if non-nil, zero value otherwise.

### GetSystemUserOk

`func (o *BTUserInfoAllOf) GetSystemUserOk() (*bool, bool)`

GetSystemUserOk returns a tuple with the SystemUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUser

`func (o *BTUserInfoAllOf) SetSystemUser(v bool)`

SetSystemUser sets SystemUser field to given value.

### HasSystemUser

`func (o *BTUserInfoAllOf) HasSystemUser() bool`

HasSystemUser returns a boolean if a field has been set.

### GetTotpEnabled

`func (o *BTUserInfoAllOf) GetTotpEnabled() bool`

GetTotpEnabled returns the TotpEnabled field if non-nil, zero value otherwise.

### GetTotpEnabledOk

`func (o *BTUserInfoAllOf) GetTotpEnabledOk() (*bool, bool)`

GetTotpEnabledOk returns a tuple with the TotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnabled

`func (o *BTUserInfoAllOf) SetTotpEnabled(v bool)`

SetTotpEnabled sets TotpEnabled field to given value.

### HasTotpEnabled

`func (o *BTUserInfoAllOf) HasTotpEnabled() bool`

HasTotpEnabled returns a boolean if a field has been set.

### GetTracingEnabled

`func (o *BTUserInfoAllOf) GetTracingEnabled() bool`

GetTracingEnabled returns the TracingEnabled field if non-nil, zero value otherwise.

### GetTracingEnabledOk

`func (o *BTUserInfoAllOf) GetTracingEnabledOk() (*bool, bool)`

GetTracingEnabledOk returns a tuple with the TracingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingEnabled

`func (o *BTUserInfoAllOf) SetTracingEnabled(v bool)`

SetTracingEnabled sets TracingEnabled field to given value.

### HasTracingEnabled

`func (o *BTUserInfoAllOf) HasTracingEnabled() bool`

HasTracingEnabled returns a boolean if a field has been set.

### GetTrialInfos

`func (o *BTUserInfoAllOf) GetTrialInfos() []BTTrialInfo`

GetTrialInfos returns the TrialInfos field if non-nil, zero value otherwise.

### GetTrialInfosOk

`func (o *BTUserInfoAllOf) GetTrialInfosOk() (*[]BTTrialInfo, bool)`

GetTrialInfosOk returns a tuple with the TrialInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialInfos

`func (o *BTUserInfoAllOf) SetTrialInfos(v []BTTrialInfo)`

SetTrialInfos sets TrialInfos field to given value.

### HasTrialInfos

`func (o *BTUserInfoAllOf) HasTrialInfos() bool`

HasTrialInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


