# BTListResponseBTCommentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI for current page of resources. | [optional] 
**Items** | Pointer to [**[]BTCommentInfo**](BTCommentInfo.md) | Array of items in the current page. | [optional] 
**Next** | Pointer to **string** | URI for next page of the resources if more are available. | [optional] 
**Previous** | Pointer to **string** | URI for previous page of the resources. | [optional] 

## Methods

### NewBTListResponseBTCommentInfo

`func NewBTListResponseBTCommentInfo() *BTListResponseBTCommentInfo`

NewBTListResponseBTCommentInfo instantiates a new BTListResponseBTCommentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTListResponseBTCommentInfoWithDefaults

`func NewBTListResponseBTCommentInfoWithDefaults() *BTListResponseBTCommentInfo`

NewBTListResponseBTCommentInfoWithDefaults instantiates a new BTListResponseBTCommentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTListResponseBTCommentInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTListResponseBTCommentInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTListResponseBTCommentInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTListResponseBTCommentInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTListResponseBTCommentInfo) GetItems() []BTCommentInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTListResponseBTCommentInfo) GetItemsOk() (*[]BTCommentInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTListResponseBTCommentInfo) SetItems(v []BTCommentInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTListResponseBTCommentInfo) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTListResponseBTCommentInfo) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTListResponseBTCommentInfo) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTListResponseBTCommentInfo) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTListResponseBTCommentInfo) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTListResponseBTCommentInfo) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTListResponseBTCommentInfo) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTListResponseBTCommentInfo) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTListResponseBTCommentInfo) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


