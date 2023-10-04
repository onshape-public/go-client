# CustomerSubscriptionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Subscription**](Subscription.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**RequestOptions** | Pointer to [**RequestOptions**](RequestOptions.md) |  | [optional] 
**RequestParams** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerSubscriptionCollection

`func NewCustomerSubscriptionCollection() *CustomerSubscriptionCollection`

NewCustomerSubscriptionCollection instantiates a new CustomerSubscriptionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSubscriptionCollectionWithDefaults

`func NewCustomerSubscriptionCollectionWithDefaults() *CustomerSubscriptionCollection`

NewCustomerSubscriptionCollectionWithDefaults instantiates a new CustomerSubscriptionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CustomerSubscriptionCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CustomerSubscriptionCollection) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CustomerSubscriptionCollection) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CustomerSubscriptionCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *CustomerSubscriptionCollection) GetData() []Subscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerSubscriptionCollection) GetDataOk() (*[]Subscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerSubscriptionCollection) SetData(v []Subscription)`

SetData sets Data field to given value.

### HasData

`func (o *CustomerSubscriptionCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasMore

`func (o *CustomerSubscriptionCollection) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CustomerSubscriptionCollection) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CustomerSubscriptionCollection) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *CustomerSubscriptionCollection) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetRequestOptions

`func (o *CustomerSubscriptionCollection) GetRequestOptions() RequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *CustomerSubscriptionCollection) GetRequestOptionsOk() (*RequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *CustomerSubscriptionCollection) SetRequestOptions(v RequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *CustomerSubscriptionCollection) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetRequestParams

`func (o *CustomerSubscriptionCollection) GetRequestParams() map[string]map[string]interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *CustomerSubscriptionCollection) GetRequestParamsOk() (*map[string]map[string]interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *CustomerSubscriptionCollection) SetRequestParams(v map[string]map[string]interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *CustomerSubscriptionCollection) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### GetTotalCount

`func (o *CustomerSubscriptionCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CustomerSubscriptionCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CustomerSubscriptionCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CustomerSubscriptionCollection) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetUrl

`func (o *CustomerSubscriptionCollection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomerSubscriptionCollection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomerSubscriptionCollection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CustomerSubscriptionCollection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


