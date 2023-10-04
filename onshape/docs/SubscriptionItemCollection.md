# SubscriptionItemCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]SubscriptionItem**](SubscriptionItem.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**RequestOptions** | Pointer to [**RequestOptions**](RequestOptions.md) |  | [optional] 
**RequestParams** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionItemCollection

`func NewSubscriptionItemCollection() *SubscriptionItemCollection`

NewSubscriptionItemCollection instantiates a new SubscriptionItemCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemCollectionWithDefaults

`func NewSubscriptionItemCollectionWithDefaults() *SubscriptionItemCollection`

NewSubscriptionItemCollectionWithDefaults instantiates a new SubscriptionItemCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SubscriptionItemCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SubscriptionItemCollection) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SubscriptionItemCollection) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SubscriptionItemCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *SubscriptionItemCollection) GetData() []SubscriptionItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionItemCollection) GetDataOk() (*[]SubscriptionItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionItemCollection) SetData(v []SubscriptionItem)`

SetData sets Data field to given value.

### HasData

`func (o *SubscriptionItemCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasMore

`func (o *SubscriptionItemCollection) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SubscriptionItemCollection) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SubscriptionItemCollection) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *SubscriptionItemCollection) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetRequestOptions

`func (o *SubscriptionItemCollection) GetRequestOptions() RequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *SubscriptionItemCollection) GetRequestOptionsOk() (*RequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *SubscriptionItemCollection) SetRequestOptions(v RequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *SubscriptionItemCollection) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetRequestParams

`func (o *SubscriptionItemCollection) GetRequestParams() map[string]map[string]interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *SubscriptionItemCollection) GetRequestParamsOk() (*map[string]map[string]interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *SubscriptionItemCollection) SetRequestParams(v map[string]map[string]interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *SubscriptionItemCollection) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### GetTotalCount

`func (o *SubscriptionItemCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SubscriptionItemCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SubscriptionItemCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SubscriptionItemCollection) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetUrl

`func (o *SubscriptionItemCollection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubscriptionItemCollection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubscriptionItemCollection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SubscriptionItemCollection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


