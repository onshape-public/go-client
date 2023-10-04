# BTCardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**BTAddressInfo**](BTAddressInfo.md) |  | [optional] 
**ExpMonth** | Pointer to **int32** |  | [optional] 
**ExpYear** | Pointer to **int32** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBTCardInfo

`func NewBTCardInfo() *BTCardInfo`

NewBTCardInfo instantiates a new BTCardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCardInfoWithDefaults

`func NewBTCardInfoWithDefaults() *BTCardInfo`

NewBTCardInfoWithDefaults instantiates a new BTCardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *BTCardInfo) GetBillingAddress() BTAddressInfo`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *BTCardInfo) GetBillingAddressOk() (*BTAddressInfo, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *BTCardInfo) SetBillingAddress(v BTAddressInfo)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *BTCardInfo) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetExpMonth

`func (o *BTCardInfo) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *BTCardInfo) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *BTCardInfo) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *BTCardInfo) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *BTCardInfo) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *BTCardInfo) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *BTCardInfo) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *BTCardInfo) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetLast4

`func (o *BTCardInfo) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *BTCardInfo) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *BTCardInfo) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *BTCardInfo) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetType

`func (o *BTCardInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTCardInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTCardInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTCardInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


