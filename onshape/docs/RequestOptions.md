# RequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**ConnectTimeout** | Pointer to **int32** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**ReadTimeout** | Pointer to **int32** |  | [optional] 
**StripeAccount** | Pointer to **string** |  | [optional] 
**StripeVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestOptions

`func NewRequestOptions() *RequestOptions`

NewRequestOptions instantiates a new RequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestOptionsWithDefaults

`func NewRequestOptionsWithDefaults() *RequestOptions`

NewRequestOptionsWithDefaults instantiates a new RequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *RequestOptions) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *RequestOptions) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *RequestOptions) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *RequestOptions) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *RequestOptions) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *RequestOptions) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *RequestOptions) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *RequestOptions) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *RequestOptions) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *RequestOptions) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *RequestOptions) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *RequestOptions) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetReadTimeout

`func (o *RequestOptions) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *RequestOptions) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *RequestOptions) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *RequestOptions) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetStripeAccount

`func (o *RequestOptions) GetStripeAccount() string`

GetStripeAccount returns the StripeAccount field if non-nil, zero value otherwise.

### GetStripeAccountOk

`func (o *RequestOptions) GetStripeAccountOk() (*string, bool)`

GetStripeAccountOk returns a tuple with the StripeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccount

`func (o *RequestOptions) SetStripeAccount(v string)`

SetStripeAccount sets StripeAccount field to given value.

### HasStripeAccount

`func (o *RequestOptions) HasStripeAccount() bool`

HasStripeAccount returns a boolean if a field has been set.

### GetStripeVersion

`func (o *RequestOptions) GetStripeVersion() string`

GetStripeVersion returns the StripeVersion field if non-nil, zero value otherwise.

### GetStripeVersionOk

`func (o *RequestOptions) GetStripeVersionOk() (*string, bool)`

GetStripeVersionOk returns a tuple with the StripeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeVersion

`func (o *RequestOptions) SetStripeVersion(v string)`

SetStripeVersion sets StripeVersion field to given value.

### HasStripeVersion

`func (o *RequestOptions) HasStripeVersion() bool`

HasStripeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


