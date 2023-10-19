# BTGlobalTreeNodeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Requested Document URL | [optional] 
**Items** | Pointer to [**[]BTGlobalTreeNodeInfo**](BTGlobalTreeNodeInfo.md) | Document Items array. Array entries are the same as that returned from \&quot;/api/documents/{did}\&quot;. | [optional] 
**Next** | Pointer to **string** | The URL for the next page of items. Responses are limited to 20 items per page. | [optional] 
**Previous** | Pointer to **string** | The URL for the previous page of items. Responses are limited to 20 items per page. | [optional] 

## Methods

### NewBTGlobalTreeNodeListResponse

`func NewBTGlobalTreeNodeListResponse() *BTGlobalTreeNodeListResponse`

NewBTGlobalTreeNodeListResponse instantiates a new BTGlobalTreeNodeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTGlobalTreeNodeListResponseWithDefaults

`func NewBTGlobalTreeNodeListResponseWithDefaults() *BTGlobalTreeNodeListResponse`

NewBTGlobalTreeNodeListResponseWithDefaults instantiates a new BTGlobalTreeNodeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTGlobalTreeNodeListResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTGlobalTreeNodeListResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTGlobalTreeNodeListResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTGlobalTreeNodeListResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTGlobalTreeNodeListResponse) GetItems() []BTGlobalTreeNodeInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTGlobalTreeNodeListResponse) GetItemsOk() (*[]BTGlobalTreeNodeInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTGlobalTreeNodeListResponse) SetItems(v []BTGlobalTreeNodeInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTGlobalTreeNodeListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTGlobalTreeNodeListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTGlobalTreeNodeListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTGlobalTreeNodeListResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTGlobalTreeNodeListResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTGlobalTreeNodeListResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTGlobalTreeNodeListResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTGlobalTreeNodeListResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTGlobalTreeNodeListResponse) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


