# ExternalAccountCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]ExternalAccount**](ExternalAccount.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**RequestOptions** | Pointer to [**RequestOptions**](RequestOptions.md) |  | [optional] 
**RequestParams** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalAccountCollection

`func NewExternalAccountCollection() *ExternalAccountCollection`

NewExternalAccountCollection instantiates a new ExternalAccountCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountCollectionWithDefaults

`func NewExternalAccountCollectionWithDefaults() *ExternalAccountCollection`

NewExternalAccountCollectionWithDefaults instantiates a new ExternalAccountCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ExternalAccountCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ExternalAccountCollection) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ExternalAccountCollection) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ExternalAccountCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *ExternalAccountCollection) GetData() []ExternalAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalAccountCollection) GetDataOk() (*[]ExternalAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalAccountCollection) SetData(v []ExternalAccount)`

SetData sets Data field to given value.

### HasData

`func (o *ExternalAccountCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasMore

`func (o *ExternalAccountCollection) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ExternalAccountCollection) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ExternalAccountCollection) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ExternalAccountCollection) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetRequestOptions

`func (o *ExternalAccountCollection) GetRequestOptions() RequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *ExternalAccountCollection) GetRequestOptionsOk() (*RequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *ExternalAccountCollection) SetRequestOptions(v RequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *ExternalAccountCollection) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetRequestParams

`func (o *ExternalAccountCollection) GetRequestParams() map[string]map[string]interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *ExternalAccountCollection) GetRequestParamsOk() (*map[string]map[string]interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *ExternalAccountCollection) SetRequestParams(v map[string]map[string]interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *ExternalAccountCollection) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### GetTotalCount

`func (o *ExternalAccountCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExternalAccountCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExternalAccountCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ExternalAccountCollection) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetUrl

`func (o *ExternalAccountCollection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExternalAccountCollection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExternalAccountCollection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ExternalAccountCollection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


