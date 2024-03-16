# BTUserInfo

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
**Discounts** | Pointer to [**[]BTDiscountInfo**](BTDiscountInfo.md) |  | [optional] 
**Enterprise** | Pointer to **bool** |  | [optional] 
**EulaId** | Pointer to **string** |  | [optional] 
**EulaRequired** | Pointer to **bool** |  | [optional] 
**EvalCenter** | Pointer to **bool** |  | [optional] 
**ForumId** | Pointer to **string** |  | [optional] 
**IntendedUse** | Pointer to **int32** |  | [optional] 
**LastTrialInfo** | Pointer to [**BTTrialInfo**](BTTrialInfo.md) |  | [optional] 
**NeedToShowNewWalkthrough** | Pointer to **bool** |  | [optional] 
**OwnPurchase** | Pointer to [**BTPurchaseInfo**](BTPurchaseInfo.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**ProDiscoveryTrialRejected** | Pointer to **bool** |  | [optional] 
**ProductType** | Pointer to **[]string** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **int32** |  | [optional] 
**Roles** | Pointer to [**[]BTRole**](BTRole.md) |  | [optional] 
**ShowRenewStudentPages** | Pointer to **bool** |  | [optional] 
**StartupPage** | Pointer to **int32** |  | [optional] 
**SystemUser** | Pointer to **bool** |  | [optional] 
**TotpEnabled** | Pointer to **bool** |  | [optional] 
**TracingEnabled** | Pointer to **bool** |  | [optional] 
**TrialInfos** | Pointer to [**[]BTTrialInfo**](BTTrialInfo.md) |  | [optional] 

## Methods

### NewBTUserInfo

`func NewBTUserInfo() *BTUserInfo`

NewBTUserInfo instantiates a new BTUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserInfoWithDefaults

`func NewBTUserInfoWithDefaults() *BTUserInfo`

NewBTUserInfoWithDefaults instantiates a new BTUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePlan

`func (o *BTUserInfo) GetActivePlan() BTBillingPlanInfo`

GetActivePlan returns the ActivePlan field if non-nil, zero value otherwise.

### GetActivePlanOk

`func (o *BTUserInfo) GetActivePlanOk() (*BTBillingPlanInfo, bool)`

GetActivePlanOk returns a tuple with the ActivePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePlan

`func (o *BTUserInfo) SetActivePlan(v BTBillingPlanInfo)`

SetActivePlan sets ActivePlan field to given value.

### HasActivePlan

`func (o *BTUserInfo) HasActivePlan() bool`

HasActivePlan returns a boolean if a field has been set.

### GetActivePlanId

`func (o *BTUserInfo) GetActivePlanId() string`

GetActivePlanId returns the ActivePlanId field if non-nil, zero value otherwise.

### GetActivePlanIdOk

`func (o *BTUserInfo) GetActivePlanIdOk() (*string, bool)`

GetActivePlanIdOk returns a tuple with the ActivePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePlanId

`func (o *BTUserInfo) SetActivePlanId(v string)`

SetActivePlanId sets ActivePlanId field to given value.

### HasActivePlanId

`func (o *BTUserInfo) HasActivePlanId() bool`

HasActivePlanId returns a boolean if a field has been set.

### GetActivePurchases

`func (o *BTUserInfo) GetActivePurchases() []BTPurchaseInfo`

GetActivePurchases returns the ActivePurchases field if non-nil, zero value otherwise.

### GetActivePurchasesOk

`func (o *BTUserInfo) GetActivePurchasesOk() (*[]BTPurchaseInfo, bool)`

GetActivePurchasesOk returns a tuple with the ActivePurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePurchases

`func (o *BTUserInfo) SetActivePurchases(v []BTPurchaseInfo)`

SetActivePurchases sets ActivePurchases field to given value.

### HasActivePurchases

`func (o *BTUserInfo) HasActivePurchases() bool`

HasActivePurchases returns a boolean if a field has been set.

### GetActiveTrialInfo

`func (o *BTUserInfo) GetActiveTrialInfo() BTTrialInfo`

GetActiveTrialInfo returns the ActiveTrialInfo field if non-nil, zero value otherwise.

### GetActiveTrialInfoOk

`func (o *BTUserInfo) GetActiveTrialInfoOk() (*BTTrialInfo, bool)`

GetActiveTrialInfoOk returns a tuple with the ActiveTrialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTrialInfo

`func (o *BTUserInfo) SetActiveTrialInfo(v BTTrialInfo)`

SetActiveTrialInfo sets ActiveTrialInfo field to given value.

### HasActiveTrialInfo

`func (o *BTUserInfo) HasActiveTrialInfo() bool`

HasActiveTrialInfo returns a boolean if a field has been set.

### GetB2cId

`func (o *BTUserInfo) GetB2cId() string`

GetB2cId returns the B2cId field if non-nil, zero value otherwise.

### GetB2cIdOk

`func (o *BTUserInfo) GetB2cIdOk() (*string, bool)`

GetB2cIdOk returns a tuple with the B2cId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2cId

`func (o *BTUserInfo) SetB2cId(v string)`

SetB2cId sets B2cId field to given value.

### HasB2cId

`func (o *BTUserInfo) HasB2cId() bool`

HasB2cId returns a boolean if a field has been set.

### GetBillingUpdateRequired

`func (o *BTUserInfo) GetBillingUpdateRequired() bool`

GetBillingUpdateRequired returns the BillingUpdateRequired field if non-nil, zero value otherwise.

### GetBillingUpdateRequiredOk

`func (o *BTUserInfo) GetBillingUpdateRequiredOk() (*bool, bool)`

GetBillingUpdateRequiredOk returns a tuple with the BillingUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUpdateRequired

`func (o *BTUserInfo) SetBillingUpdateRequired(v bool)`

SetBillingUpdateRequired sets BillingUpdateRequired field to given value.

### HasBillingUpdateRequired

`func (o *BTUserInfo) HasBillingUpdateRequired() bool`

HasBillingUpdateRequired returns a boolean if a field has been set.

### GetCadSystemAtSignup

`func (o *BTUserInfo) GetCadSystemAtSignup() string`

GetCadSystemAtSignup returns the CadSystemAtSignup field if non-nil, zero value otherwise.

### GetCadSystemAtSignupOk

`func (o *BTUserInfo) GetCadSystemAtSignupOk() (*string, bool)`

GetCadSystemAtSignupOk returns a tuple with the CadSystemAtSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadSystemAtSignup

`func (o *BTUserInfo) SetCadSystemAtSignup(v string)`

SetCadSystemAtSignup sets CadSystemAtSignup field to given value.

### HasCadSystemAtSignup

`func (o *BTUserInfo) HasCadSystemAtSignup() bool`

HasCadSystemAtSignup returns a boolean if a field has been set.

### GetCountryCode

`func (o *BTUserInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BTUserInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BTUserInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BTUserInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTUserInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTUserInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTUserInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTUserInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredential

`func (o *BTUserInfo) GetCredential() BTSessionCredentialInfo`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *BTUserInfo) GetCredentialOk() (*BTSessionCredentialInfo, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *BTUserInfo) SetCredential(v BTSessionCredentialInfo)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *BTUserInfo) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDefaultCompanyName

`func (o *BTUserInfo) GetDefaultCompanyName() string`

GetDefaultCompanyName returns the DefaultCompanyName field if non-nil, zero value otherwise.

### GetDefaultCompanyNameOk

`func (o *BTUserInfo) GetDefaultCompanyNameOk() (*string, bool)`

GetDefaultCompanyNameOk returns a tuple with the DefaultCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCompanyName

`func (o *BTUserInfo) SetDefaultCompanyName(v string)`

SetDefaultCompanyName sets DefaultCompanyName field to given value.

### HasDefaultCompanyName

`func (o *BTUserInfo) HasDefaultCompanyName() bool`

HasDefaultCompanyName returns a boolean if a field has been set.

### GetDescription

`func (o *BTUserInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTUserInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTUserInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTUserInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *BTUserInfo) GetDeviceInfo() BTDeviceLoginSecretInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *BTUserInfo) GetDeviceInfoOk() (*BTDeviceLoginSecretInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *BTUserInfo) SetDeviceInfo(v BTDeviceLoginSecretInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *BTUserInfo) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetDiscounts

`func (o *BTUserInfo) GetDiscounts() []BTDiscountInfo`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *BTUserInfo) GetDiscountsOk() (*[]BTDiscountInfo, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *BTUserInfo) SetDiscounts(v []BTDiscountInfo)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *BTUserInfo) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetEnterprise

`func (o *BTUserInfo) GetEnterprise() bool`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *BTUserInfo) GetEnterpriseOk() (*bool, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *BTUserInfo) SetEnterprise(v bool)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *BTUserInfo) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.

### GetEulaId

`func (o *BTUserInfo) GetEulaId() string`

GetEulaId returns the EulaId field if non-nil, zero value otherwise.

### GetEulaIdOk

`func (o *BTUserInfo) GetEulaIdOk() (*string, bool)`

GetEulaIdOk returns a tuple with the EulaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaId

`func (o *BTUserInfo) SetEulaId(v string)`

SetEulaId sets EulaId field to given value.

### HasEulaId

`func (o *BTUserInfo) HasEulaId() bool`

HasEulaId returns a boolean if a field has been set.

### GetEulaRequired

`func (o *BTUserInfo) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *BTUserInfo) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *BTUserInfo) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *BTUserInfo) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetEvalCenter

`func (o *BTUserInfo) GetEvalCenter() bool`

GetEvalCenter returns the EvalCenter field if non-nil, zero value otherwise.

### GetEvalCenterOk

`func (o *BTUserInfo) GetEvalCenterOk() (*bool, bool)`

GetEvalCenterOk returns a tuple with the EvalCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalCenter

`func (o *BTUserInfo) SetEvalCenter(v bool)`

SetEvalCenter sets EvalCenter field to given value.

### HasEvalCenter

`func (o *BTUserInfo) HasEvalCenter() bool`

HasEvalCenter returns a boolean if a field has been set.

### GetForumId

`func (o *BTUserInfo) GetForumId() string`

GetForumId returns the ForumId field if non-nil, zero value otherwise.

### GetForumIdOk

`func (o *BTUserInfo) GetForumIdOk() (*string, bool)`

GetForumIdOk returns a tuple with the ForumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumId

`func (o *BTUserInfo) SetForumId(v string)`

SetForumId sets ForumId field to given value.

### HasForumId

`func (o *BTUserInfo) HasForumId() bool`

HasForumId returns a boolean if a field has been set.

### GetIntendedUse

`func (o *BTUserInfo) GetIntendedUse() int32`

GetIntendedUse returns the IntendedUse field if non-nil, zero value otherwise.

### GetIntendedUseOk

`func (o *BTUserInfo) GetIntendedUseOk() (*int32, bool)`

GetIntendedUseOk returns a tuple with the IntendedUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedUse

`func (o *BTUserInfo) SetIntendedUse(v int32)`

SetIntendedUse sets IntendedUse field to given value.

### HasIntendedUse

`func (o *BTUserInfo) HasIntendedUse() bool`

HasIntendedUse returns a boolean if a field has been set.

### GetLastTrialInfo

`func (o *BTUserInfo) GetLastTrialInfo() BTTrialInfo`

GetLastTrialInfo returns the LastTrialInfo field if non-nil, zero value otherwise.

### GetLastTrialInfoOk

`func (o *BTUserInfo) GetLastTrialInfoOk() (*BTTrialInfo, bool)`

GetLastTrialInfoOk returns a tuple with the LastTrialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialInfo

`func (o *BTUserInfo) SetLastTrialInfo(v BTTrialInfo)`

SetLastTrialInfo sets LastTrialInfo field to given value.

### HasLastTrialInfo

`func (o *BTUserInfo) HasLastTrialInfo() bool`

HasLastTrialInfo returns a boolean if a field has been set.

### GetNeedToShowNewWalkthrough

`func (o *BTUserInfo) GetNeedToShowNewWalkthrough() bool`

GetNeedToShowNewWalkthrough returns the NeedToShowNewWalkthrough field if non-nil, zero value otherwise.

### GetNeedToShowNewWalkthroughOk

`func (o *BTUserInfo) GetNeedToShowNewWalkthroughOk() (*bool, bool)`

GetNeedToShowNewWalkthroughOk returns a tuple with the NeedToShowNewWalkthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedToShowNewWalkthrough

`func (o *BTUserInfo) SetNeedToShowNewWalkthrough(v bool)`

SetNeedToShowNewWalkthrough sets NeedToShowNewWalkthrough field to given value.

### HasNeedToShowNewWalkthrough

`func (o *BTUserInfo) HasNeedToShowNewWalkthrough() bool`

HasNeedToShowNewWalkthrough returns a boolean if a field has been set.

### GetOwnPurchase

`func (o *BTUserInfo) GetOwnPurchase() BTPurchaseInfo`

GetOwnPurchase returns the OwnPurchase field if non-nil, zero value otherwise.

### GetOwnPurchaseOk

`func (o *BTUserInfo) GetOwnPurchaseOk() (*BTPurchaseInfo, bool)`

GetOwnPurchaseOk returns a tuple with the OwnPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnPurchase

`func (o *BTUserInfo) SetOwnPurchase(v BTPurchaseInfo)`

SetOwnPurchase sets OwnPurchase field to given value.

### HasOwnPurchase

`func (o *BTUserInfo) HasOwnPurchase() bool`

HasOwnPurchase returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BTUserInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BTUserInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BTUserInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BTUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetProDiscoveryTrialRejected

`func (o *BTUserInfo) GetProDiscoveryTrialRejected() bool`

GetProDiscoveryTrialRejected returns the ProDiscoveryTrialRejected field if non-nil, zero value otherwise.

### GetProDiscoveryTrialRejectedOk

`func (o *BTUserInfo) GetProDiscoveryTrialRejectedOk() (*bool, bool)`

GetProDiscoveryTrialRejectedOk returns a tuple with the ProDiscoveryTrialRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProDiscoveryTrialRejected

`func (o *BTUserInfo) SetProDiscoveryTrialRejected(v bool)`

SetProDiscoveryTrialRejected sets ProDiscoveryTrialRejected field to given value.

### HasProDiscoveryTrialRejected

`func (o *BTUserInfo) HasProDiscoveryTrialRejected() bool`

HasProDiscoveryTrialRejected returns a boolean if a field has been set.

### GetProductType

`func (o *BTUserInfo) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BTUserInfo) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BTUserInfo) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *BTUserInfo) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *BTUserInfo) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *BTUserInfo) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *BTUserInfo) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *BTUserInfo) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetRole

`func (o *BTUserInfo) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BTUserInfo) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BTUserInfo) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *BTUserInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *BTUserInfo) GetRoles() []BTRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BTUserInfo) GetRolesOk() (*[]BTRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BTUserInfo) SetRoles(v []BTRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BTUserInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetShowRenewStudentPages

`func (o *BTUserInfo) GetShowRenewStudentPages() bool`

GetShowRenewStudentPages returns the ShowRenewStudentPages field if non-nil, zero value otherwise.

### GetShowRenewStudentPagesOk

`func (o *BTUserInfo) GetShowRenewStudentPagesOk() (*bool, bool)`

GetShowRenewStudentPagesOk returns a tuple with the ShowRenewStudentPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRenewStudentPages

`func (o *BTUserInfo) SetShowRenewStudentPages(v bool)`

SetShowRenewStudentPages sets ShowRenewStudentPages field to given value.

### HasShowRenewStudentPages

`func (o *BTUserInfo) HasShowRenewStudentPages() bool`

HasShowRenewStudentPages returns a boolean if a field has been set.

### GetStartupPage

`func (o *BTUserInfo) GetStartupPage() int32`

GetStartupPage returns the StartupPage field if non-nil, zero value otherwise.

### GetStartupPageOk

`func (o *BTUserInfo) GetStartupPageOk() (*int32, bool)`

GetStartupPageOk returns a tuple with the StartupPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupPage

`func (o *BTUserInfo) SetStartupPage(v int32)`

SetStartupPage sets StartupPage field to given value.

### HasStartupPage

`func (o *BTUserInfo) HasStartupPage() bool`

HasStartupPage returns a boolean if a field has been set.

### GetSystemUser

`func (o *BTUserInfo) GetSystemUser() bool`

GetSystemUser returns the SystemUser field if non-nil, zero value otherwise.

### GetSystemUserOk

`func (o *BTUserInfo) GetSystemUserOk() (*bool, bool)`

GetSystemUserOk returns a tuple with the SystemUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUser

`func (o *BTUserInfo) SetSystemUser(v bool)`

SetSystemUser sets SystemUser field to given value.

### HasSystemUser

`func (o *BTUserInfo) HasSystemUser() bool`

HasSystemUser returns a boolean if a field has been set.

### GetTotpEnabled

`func (o *BTUserInfo) GetTotpEnabled() bool`

GetTotpEnabled returns the TotpEnabled field if non-nil, zero value otherwise.

### GetTotpEnabledOk

`func (o *BTUserInfo) GetTotpEnabledOk() (*bool, bool)`

GetTotpEnabledOk returns a tuple with the TotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnabled

`func (o *BTUserInfo) SetTotpEnabled(v bool)`

SetTotpEnabled sets TotpEnabled field to given value.

### HasTotpEnabled

`func (o *BTUserInfo) HasTotpEnabled() bool`

HasTotpEnabled returns a boolean if a field has been set.

### GetTracingEnabled

`func (o *BTUserInfo) GetTracingEnabled() bool`

GetTracingEnabled returns the TracingEnabled field if non-nil, zero value otherwise.

### GetTracingEnabledOk

`func (o *BTUserInfo) GetTracingEnabledOk() (*bool, bool)`

GetTracingEnabledOk returns a tuple with the TracingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingEnabled

`func (o *BTUserInfo) SetTracingEnabled(v bool)`

SetTracingEnabled sets TracingEnabled field to given value.

### HasTracingEnabled

`func (o *BTUserInfo) HasTracingEnabled() bool`

HasTracingEnabled returns a boolean if a field has been set.

### GetTrialInfos

`func (o *BTUserInfo) GetTrialInfos() []BTTrialInfo`

GetTrialInfos returns the TrialInfos field if non-nil, zero value otherwise.

### GetTrialInfosOk

`func (o *BTUserInfo) GetTrialInfosOk() (*[]BTTrialInfo, bool)`

GetTrialInfosOk returns a tuple with the TrialInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialInfos

`func (o *BTUserInfo) SetTrialInfos(v []BTTrialInfo)`

SetTrialInfos sets TrialInfos field to given value.

### HasTrialInfos

`func (o *BTUserInfo) HasTrialInfos() bool`

HasTrialInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


