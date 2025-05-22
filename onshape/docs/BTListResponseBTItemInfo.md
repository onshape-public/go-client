# BTListResponseBTItemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI for current page of resources. | [optional] 
**Items** | Pointer to [**[]BTItemInfo**](BTItemInfo.md) | Array of items in the current page. | [optional] 
**Next** | Pointer to **string** | URI for next page of the resources if more are available. | [optional] 
**Previous** | Pointer to **string** | URI for previous page of the resources. | [optional] 

## Methods

### NewBTListResponseBTItemInfo

`func NewBTListResponseBTItemInfo() *BTListResponseBTItemInfo`

NewBTListResponseBTItemInfo instantiates a new BTListResponseBTItemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTListResponseBTItemInfoWithDefaults

`func NewBTListResponseBTItemInfoWithDefaults() *BTListResponseBTItemInfo`

NewBTListResponseBTItemInfoWithDefaults instantiates a new BTListResponseBTItemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTListResponseBTItemInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTListResponseBTItemInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTListResponseBTItemInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTListResponseBTItemInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTListResponseBTItemInfo) GetItems() []BTItemInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTListResponseBTItemInfo) GetItemsOk() (*[]BTItemInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTListResponseBTItemInfo) SetItems(v []BTItemInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTListResponseBTItemInfo) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTListResponseBTItemInfo) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTListResponseBTItemInfo) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTListResponseBTItemInfo) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTListResponseBTItemInfo) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTListResponseBTItemInfo) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTListResponseBTItemInfo) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTListResponseBTItemInfo) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTListResponseBTItemInfo) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


