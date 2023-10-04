# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressCountry** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine1Check** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressState** | Pointer to **string** |  | [optional] 
**AddressZip** | Pointer to **string** |  | [optional] 
**AddressZipCheck** | Pointer to **string** |  | [optional] 
**AvailablePayoutMethods** | Pointer to **[]string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to **string** |  | [optional] 
**CvcCheck** | Pointer to **string** |  | [optional] 
**DefaultForCurrency** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DynamicLast4** | Pointer to **string** |  | [optional] 
**ExpMonth** | Pointer to **int32** |  | [optional] 
**ExpYear** | Pointer to **int32** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Funding** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Iin** | Pointer to **string** |  | [optional] 
**InstanceURL** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ThreeDSecure** | Pointer to [**ThreeDSecure**](ThreeDSecure.md) |  | [optional] 
**TokenizationMethod** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCard

`func NewCard() *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Card) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Card) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Card) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Card) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAddressCity

`func (o *Card) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Card) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Card) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Card) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *Card) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *Card) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *Card) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *Card) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressLine1

`func (o *Card) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Card) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Card) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *Card) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine1Check

`func (o *Card) GetAddressLine1Check() string`

GetAddressLine1Check returns the AddressLine1Check field if non-nil, zero value otherwise.

### GetAddressLine1CheckOk

`func (o *Card) GetAddressLine1CheckOk() (*string, bool)`

GetAddressLine1CheckOk returns a tuple with the AddressLine1Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1Check

`func (o *Card) SetAddressLine1Check(v string)`

SetAddressLine1Check sets AddressLine1Check field to given value.

### HasAddressLine1Check

`func (o *Card) HasAddressLine1Check() bool`

HasAddressLine1Check returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Card) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Card) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Card) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Card) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressState

`func (o *Card) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *Card) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *Card) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *Card) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetAddressZip

`func (o *Card) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *Card) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *Card) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *Card) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### GetAddressZipCheck

`func (o *Card) GetAddressZipCheck() string`

GetAddressZipCheck returns the AddressZipCheck field if non-nil, zero value otherwise.

### GetAddressZipCheckOk

`func (o *Card) GetAddressZipCheckOk() (*string, bool)`

GetAddressZipCheckOk returns a tuple with the AddressZipCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZipCheck

`func (o *Card) SetAddressZipCheck(v string)`

SetAddressZipCheck sets AddressZipCheck field to given value.

### HasAddressZipCheck

`func (o *Card) HasAddressZipCheck() bool`

HasAddressZipCheck returns a boolean if a field has been set.

### GetAvailablePayoutMethods

`func (o *Card) GetAvailablePayoutMethods() []string`

GetAvailablePayoutMethods returns the AvailablePayoutMethods field if non-nil, zero value otherwise.

### GetAvailablePayoutMethodsOk

`func (o *Card) GetAvailablePayoutMethodsOk() (*[]string, bool)`

GetAvailablePayoutMethodsOk returns a tuple with the AvailablePayoutMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePayoutMethods

`func (o *Card) SetAvailablePayoutMethods(v []string)`

SetAvailablePayoutMethods sets AvailablePayoutMethods field to given value.

### HasAvailablePayoutMethods

`func (o *Card) HasAvailablePayoutMethods() bool`

HasAvailablePayoutMethods returns a boolean if a field has been set.

### GetBrand

`func (o *Card) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Card) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Card) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Card) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCountry

`func (o *Card) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Card) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Card) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Card) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *Card) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Card) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Card) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Card) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomer

`func (o *Card) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Card) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Card) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Card) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCvcCheck

`func (o *Card) GetCvcCheck() string`

GetCvcCheck returns the CvcCheck field if non-nil, zero value otherwise.

### GetCvcCheckOk

`func (o *Card) GetCvcCheckOk() (*string, bool)`

GetCvcCheckOk returns a tuple with the CvcCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvcCheck

`func (o *Card) SetCvcCheck(v string)`

SetCvcCheck sets CvcCheck field to given value.

### HasCvcCheck

`func (o *Card) HasCvcCheck() bool`

HasCvcCheck returns a boolean if a field has been set.

### GetDefaultForCurrency

`func (o *Card) GetDefaultForCurrency() bool`

GetDefaultForCurrency returns the DefaultForCurrency field if non-nil, zero value otherwise.

### GetDefaultForCurrencyOk

`func (o *Card) GetDefaultForCurrencyOk() (*bool, bool)`

GetDefaultForCurrencyOk returns a tuple with the DefaultForCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForCurrency

`func (o *Card) SetDefaultForCurrency(v bool)`

SetDefaultForCurrency sets DefaultForCurrency field to given value.

### HasDefaultForCurrency

`func (o *Card) HasDefaultForCurrency() bool`

HasDefaultForCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *Card) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Card) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Card) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Card) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDynamicLast4

`func (o *Card) GetDynamicLast4() string`

GetDynamicLast4 returns the DynamicLast4 field if non-nil, zero value otherwise.

### GetDynamicLast4Ok

`func (o *Card) GetDynamicLast4Ok() (*string, bool)`

GetDynamicLast4Ok returns a tuple with the DynamicLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicLast4

`func (o *Card) SetDynamicLast4(v string)`

SetDynamicLast4 sets DynamicLast4 field to given value.

### HasDynamicLast4

`func (o *Card) HasDynamicLast4() bool`

HasDynamicLast4 returns a boolean if a field has been set.

### GetExpMonth

`func (o *Card) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *Card) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *Card) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *Card) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *Card) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *Card) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *Card) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *Card) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetFingerprint

`func (o *Card) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Card) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Card) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Card) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFunding

`func (o *Card) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *Card) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *Card) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *Card) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetId

`func (o *Card) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Card) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Card) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Card) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIin

`func (o *Card) GetIin() string`

GetIin returns the Iin field if non-nil, zero value otherwise.

### GetIinOk

`func (o *Card) GetIinOk() (*string, bool)`

GetIinOk returns a tuple with the Iin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIin

`func (o *Card) SetIin(v string)`

SetIin sets Iin field to given value.

### HasIin

`func (o *Card) HasIin() bool`

HasIin returns a boolean if a field has been set.

### GetInstanceURL

`func (o *Card) GetInstanceURL() string`

GetInstanceURL returns the InstanceURL field if non-nil, zero value otherwise.

### GetInstanceURLOk

`func (o *Card) GetInstanceURLOk() (*string, bool)`

GetInstanceURLOk returns a tuple with the InstanceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceURL

`func (o *Card) SetInstanceURL(v string)`

SetInstanceURL sets InstanceURL field to given value.

### HasInstanceURL

`func (o *Card) HasInstanceURL() bool`

HasInstanceURL returns a boolean if a field has been set.

### GetIssuer

`func (o *Card) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Card) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Card) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Card) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLast4

`func (o *Card) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *Card) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *Card) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *Card) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetMetadata

`func (o *Card) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Card) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Card) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Card) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Card) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Card) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Card) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Card) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *Card) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Card) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Card) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Card) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRecipient

`func (o *Card) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Card) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Card) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Card) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetStatus

`func (o *Card) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Card) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Card) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Card) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThreeDSecure

`func (o *Card) GetThreeDSecure() ThreeDSecure`

GetThreeDSecure returns the ThreeDSecure field if non-nil, zero value otherwise.

### GetThreeDSecureOk

`func (o *Card) GetThreeDSecureOk() (*ThreeDSecure, bool)`

GetThreeDSecureOk returns a tuple with the ThreeDSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecure

`func (o *Card) SetThreeDSecure(v ThreeDSecure)`

SetThreeDSecure sets ThreeDSecure field to given value.

### HasThreeDSecure

`func (o *Card) HasThreeDSecure() bool`

HasThreeDSecure returns a boolean if a field has been set.

### GetTokenizationMethod

`func (o *Card) GetTokenizationMethod() string`

GetTokenizationMethod returns the TokenizationMethod field if non-nil, zero value otherwise.

### GetTokenizationMethodOk

`func (o *Card) GetTokenizationMethodOk() (*string, bool)`

GetTokenizationMethodOk returns a tuple with the TokenizationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizationMethod

`func (o *Card) SetTokenizationMethod(v string)`

SetTokenizationMethod sets TokenizationMethod field to given value.

### HasTokenizationMethod

`func (o *Card) HasTokenizationMethod() bool`

HasTokenizationMethod returns a boolean if a field has been set.

### GetType

`func (o *Card) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Card) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Card) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Card) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


