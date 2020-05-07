# ShippingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewShippingDetails

`func NewShippingDetails() *ShippingDetails`

NewShippingDetails instantiates a new ShippingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingDetailsWithDefaults

`func NewShippingDetailsWithDefaults() *ShippingDetails`

NewShippingDetailsWithDefaults instantiates a new ShippingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ShippingDetails) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ShippingDetails) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ShippingDetails) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ShippingDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *ShippingDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ShippingDetails) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ShippingDetails) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ShippingDetails) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ShippingDetails) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


