# BTReleasePackageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeOrderId** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]BTReleasePackageItemParams**](BTReleasePackageItemParams.md) |  | [optional] 

## Methods

### NewBTReleasePackageParams

`func NewBTReleasePackageParams() *BTReleasePackageParams`

NewBTReleasePackageParams instantiates a new BTReleasePackageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTReleasePackageParamsWithDefaults

`func NewBTReleasePackageParamsWithDefaults() *BTReleasePackageParams`

NewBTReleasePackageParamsWithDefaults instantiates a new BTReleasePackageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeOrderId

`func (o *BTReleasePackageParams) GetChangeOrderId() string`

GetChangeOrderId returns the ChangeOrderId field if non-nil, zero value otherwise.

### GetChangeOrderIdOk

`func (o *BTReleasePackageParams) GetChangeOrderIdOk() (*string, bool)`

GetChangeOrderIdOk returns a tuple with the ChangeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeOrderId

`func (o *BTReleasePackageParams) SetChangeOrderId(v string)`

SetChangeOrderId sets ChangeOrderId field to given value.

### HasChangeOrderId

`func (o *BTReleasePackageParams) HasChangeOrderId() bool`

HasChangeOrderId returns a boolean if a field has been set.

### GetItems

`func (o *BTReleasePackageParams) GetItems() []BTReleasePackageItemParams`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTReleasePackageParams) GetItemsOk() (*[]BTReleasePackageItemParams, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTReleasePackageParams) SetItems(v []BTReleasePackageItemParams)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTReleasePackageParams) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


