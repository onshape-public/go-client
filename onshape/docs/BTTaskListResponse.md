# BTTaskListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI for current page of resources. | [optional] 
**Items** | Pointer to [**[]BTTaskInfo**](BTTaskInfo.md) | Array of items in the current page. | [optional] 
**Next** | Pointer to **string** | URI for next page of the resources if more are available. | [optional] 
**Previous** | Pointer to **string** | URI for previous page of the resources. | [optional] 
**TaskTypes** | Pointer to [**[]BTTaskTypeInfo**](BTTaskTypeInfo.md) |  | [optional] 

## Methods

### NewBTTaskListResponse

`func NewBTTaskListResponse() *BTTaskListResponse`

NewBTTaskListResponse instantiates a new BTTaskListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTaskListResponseWithDefaults

`func NewBTTaskListResponseWithDefaults() *BTTaskListResponse`

NewBTTaskListResponseWithDefaults instantiates a new BTTaskListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTTaskListResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTaskListResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTaskListResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTaskListResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTTaskListResponse) GetItems() []BTTaskInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTTaskListResponse) GetItemsOk() (*[]BTTaskInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTTaskListResponse) SetItems(v []BTTaskInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTTaskListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTTaskListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTTaskListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTTaskListResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTTaskListResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTTaskListResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTTaskListResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTTaskListResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTTaskListResponse) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetTaskTypes

`func (o *BTTaskListResponse) GetTaskTypes() []BTTaskTypeInfo`

GetTaskTypes returns the TaskTypes field if non-nil, zero value otherwise.

### GetTaskTypesOk

`func (o *BTTaskListResponse) GetTaskTypesOk() (*[]BTTaskTypeInfo, bool)`

GetTaskTypesOk returns a tuple with the TaskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTypes

`func (o *BTTaskListResponse) SetTaskTypes(v []BTTaskTypeInfo)`

SetTaskTypes sets TaskTypes field to given value.

### HasTaskTypes

`func (o *BTTaskListResponse) HasTaskTypes() bool`

HasTaskTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


