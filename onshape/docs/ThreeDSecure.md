# ThreeDSecure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**Authenticated** | Pointer to **bool** |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**RedirectURL** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewThreeDSecure

`func NewThreeDSecure() *ThreeDSecure`

NewThreeDSecure instantiates a new ThreeDSecure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureWithDefaults

`func NewThreeDSecureWithDefaults() *ThreeDSecure`

NewThreeDSecureWithDefaults instantiates a new ThreeDSecure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ThreeDSecure) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ThreeDSecure) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ThreeDSecure) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ThreeDSecure) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthenticated

`func (o *ThreeDSecure) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *ThreeDSecure) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *ThreeDSecure) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *ThreeDSecure) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetCard

`func (o *ThreeDSecure) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *ThreeDSecure) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *ThreeDSecure) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *ThreeDSecure) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCreated

`func (o *ThreeDSecure) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ThreeDSecure) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ThreeDSecure) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ThreeDSecure) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *ThreeDSecure) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ThreeDSecure) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ThreeDSecure) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ThreeDSecure) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *ThreeDSecure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreeDSecure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreeDSecure) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreeDSecure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *ThreeDSecure) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ThreeDSecure) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ThreeDSecure) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ThreeDSecure) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *ThreeDSecure) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ThreeDSecure) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ThreeDSecure) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ThreeDSecure) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRedirectURL

`func (o *ThreeDSecure) GetRedirectURL() string`

GetRedirectURL returns the RedirectURL field if non-nil, zero value otherwise.

### GetRedirectURLOk

`func (o *ThreeDSecure) GetRedirectURLOk() (*string, bool)`

GetRedirectURLOk returns a tuple with the RedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectURL

`func (o *ThreeDSecure) SetRedirectURL(v string)`

SetRedirectURL sets RedirectURL field to given value.

### HasRedirectURL

`func (o *ThreeDSecure) HasRedirectURL() bool`

HasRedirectURL returns a boolean if a field has been set.

### GetStatus

`func (o *ThreeDSecure) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ThreeDSecure) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ThreeDSecure) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ThreeDSecure) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


