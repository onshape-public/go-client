# BTRevisionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI for current page of resources. | [optional] 
**Items** | Pointer to [**[]BTRevisionInfo**](BTRevisionInfo.md) | Array of items in the current page. | [optional] 
**Next** | Pointer to **string** | URI for next page of the resources if more are available. | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** | URI for previous page of the resources. | [optional] 

## Methods

### NewBTRevisionListResponse

`func NewBTRevisionListResponse() *BTRevisionListResponse`

NewBTRevisionListResponse instantiates a new BTRevisionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRevisionListResponseWithDefaults

`func NewBTRevisionListResponseWithDefaults() *BTRevisionListResponse`

NewBTRevisionListResponseWithDefaults instantiates a new BTRevisionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTRevisionListResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTRevisionListResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTRevisionListResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTRevisionListResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTRevisionListResponse) GetItems() []BTRevisionInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTRevisionListResponse) GetItemsOk() (*[]BTRevisionInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTRevisionListResponse) SetItems(v []BTRevisionInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTRevisionListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTRevisionListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTRevisionListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTRevisionListResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTRevisionListResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTRevisionListResponse) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTRevisionListResponse) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTRevisionListResponse) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTRevisionListResponse) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPrevious

`func (o *BTRevisionListResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTRevisionListResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTRevisionListResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTRevisionListResponse) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


