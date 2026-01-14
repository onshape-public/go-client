# BTUpdateReleasePackageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Empty** | Pointer to **bool** |  | [optional] 
**ItemIds** | Pointer to **[]string** |  | [optional] 
**Items** | Pointer to [**[]BTReleasePackageItemParams**](BTReleasePackageItemParams.md) |  | [optional] 
**Properties** | Pointer to [**[]BTPropertyValueParam**](BTPropertyValueParam.md) |  | [optional] 

## Methods

### NewBTUpdateReleasePackageParams

`func NewBTUpdateReleasePackageParams() *BTUpdateReleasePackageParams`

NewBTUpdateReleasePackageParams instantiates a new BTUpdateReleasePackageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUpdateReleasePackageParamsWithDefaults

`func NewBTUpdateReleasePackageParamsWithDefaults() *BTUpdateReleasePackageParams`

NewBTUpdateReleasePackageParamsWithDefaults instantiates a new BTUpdateReleasePackageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmpty

`func (o *BTUpdateReleasePackageParams) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *BTUpdateReleasePackageParams) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *BTUpdateReleasePackageParams) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *BTUpdateReleasePackageParams) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetItemIds

`func (o *BTUpdateReleasePackageParams) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *BTUpdateReleasePackageParams) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *BTUpdateReleasePackageParams) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *BTUpdateReleasePackageParams) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetItems

`func (o *BTUpdateReleasePackageParams) GetItems() []BTReleasePackageItemParams`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTUpdateReleasePackageParams) GetItemsOk() (*[]BTReleasePackageItemParams, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTUpdateReleasePackageParams) SetItems(v []BTReleasePackageItemParams)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTUpdateReleasePackageParams) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetProperties

`func (o *BTUpdateReleasePackageParams) GetProperties() []BTPropertyValueParam`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTUpdateReleasePackageParams) GetPropertiesOk() (*[]BTPropertyValueParam, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTUpdateReleasePackageParams) SetProperties(v []BTPropertyValueParam)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTUpdateReleasePackageParams) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


