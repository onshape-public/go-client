# BTListResponseBTCompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI for current page of resources. | [optional] 
**Items** | Pointer to [**[]BTCompanyInfo**](BTCompanyInfo.md) | Array of items in the current page. | [optional] 
**Next** | Pointer to **string** | URI for next page of the resources if more are available. | [optional] 
**Previous** | Pointer to **string** | URI for previous page of the resources. | [optional] 

## Methods

### NewBTListResponseBTCompanyInfo

`func NewBTListResponseBTCompanyInfo() *BTListResponseBTCompanyInfo`

NewBTListResponseBTCompanyInfo instantiates a new BTListResponseBTCompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTListResponseBTCompanyInfoWithDefaults

`func NewBTListResponseBTCompanyInfoWithDefaults() *BTListResponseBTCompanyInfo`

NewBTListResponseBTCompanyInfoWithDefaults instantiates a new BTListResponseBTCompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTListResponseBTCompanyInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTListResponseBTCompanyInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTListResponseBTCompanyInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTListResponseBTCompanyInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTListResponseBTCompanyInfo) GetItems() []BTCompanyInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTListResponseBTCompanyInfo) GetItemsOk() (*[]BTCompanyInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTListResponseBTCompanyInfo) SetItems(v []BTCompanyInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTListResponseBTCompanyInfo) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTListResponseBTCompanyInfo) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTListResponseBTCompanyInfo) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTListResponseBTCompanyInfo) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTListResponseBTCompanyInfo) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTListResponseBTCompanyInfo) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTListResponseBTCompanyInfo) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTListResponseBTCompanyInfo) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTListResponseBTCompanyInfo) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


