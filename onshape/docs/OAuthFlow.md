# OAuthFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUrl** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**RefreshUrl** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to [**OAuthFlowScopes**](OAuthFlowScopes.md) |  | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuthFlow

`func NewOAuthFlow() *OAuthFlow`

NewOAuthFlow instantiates a new OAuthFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthFlowWithDefaults

`func NewOAuthFlowWithDefaults() *OAuthFlow`

NewOAuthFlowWithDefaults instantiates a new OAuthFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUrl

`func (o *OAuthFlow) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *OAuthFlow) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *OAuthFlow) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *OAuthFlow) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### GetExtensions

`func (o *OAuthFlow) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *OAuthFlow) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *OAuthFlow) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *OAuthFlow) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetRefreshUrl

`func (o *OAuthFlow) GetRefreshUrl() string`

GetRefreshUrl returns the RefreshUrl field if non-nil, zero value otherwise.

### GetRefreshUrlOk

`func (o *OAuthFlow) GetRefreshUrlOk() (*string, bool)`

GetRefreshUrlOk returns a tuple with the RefreshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUrl

`func (o *OAuthFlow) SetRefreshUrl(v string)`

SetRefreshUrl sets RefreshUrl field to given value.

### HasRefreshUrl

`func (o *OAuthFlow) HasRefreshUrl() bool`

HasRefreshUrl returns a boolean if a field has been set.

### GetScopes

`func (o *OAuthFlow) GetScopes() OAuthFlowScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthFlow) GetScopesOk() (*OAuthFlowScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthFlow) SetScopes(v OAuthFlowScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthFlow) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenUrl

`func (o *OAuthFlow) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuthFlow) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuthFlow) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *OAuthFlow) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


