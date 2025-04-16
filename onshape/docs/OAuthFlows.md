# OAuthFlows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationCode** | Pointer to [**OAuthFlow**](OAuthFlow.md) |  | [optional] 
**ClientCredentials** | Pointer to [**OAuthFlow**](OAuthFlow.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Implicit** | Pointer to [**OAuthFlow**](OAuthFlow.md) |  | [optional] 
**Password** | Pointer to [**OAuthFlow**](OAuthFlow.md) |  | [optional] 

## Methods

### NewOAuthFlows

`func NewOAuthFlows() *OAuthFlows`

NewOAuthFlows instantiates a new OAuthFlows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthFlowsWithDefaults

`func NewOAuthFlowsWithDefaults() *OAuthFlows`

NewOAuthFlowsWithDefaults instantiates a new OAuthFlows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationCode

`func (o *OAuthFlows) GetAuthorizationCode() OAuthFlow`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *OAuthFlows) GetAuthorizationCodeOk() (*OAuthFlow, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *OAuthFlows) SetAuthorizationCode(v OAuthFlow)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *OAuthFlows) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetClientCredentials

`func (o *OAuthFlows) GetClientCredentials() OAuthFlow`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *OAuthFlows) GetClientCredentialsOk() (*OAuthFlow, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *OAuthFlows) SetClientCredentials(v OAuthFlow)`

SetClientCredentials sets ClientCredentials field to given value.

### HasClientCredentials

`func (o *OAuthFlows) HasClientCredentials() bool`

HasClientCredentials returns a boolean if a field has been set.

### GetExtensions

`func (o *OAuthFlows) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *OAuthFlows) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *OAuthFlows) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *OAuthFlows) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetImplicit

`func (o *OAuthFlows) GetImplicit() OAuthFlow`

GetImplicit returns the Implicit field if non-nil, zero value otherwise.

### GetImplicitOk

`func (o *OAuthFlows) GetImplicitOk() (*OAuthFlow, bool)`

GetImplicitOk returns a tuple with the Implicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicit

`func (o *OAuthFlows) SetImplicit(v OAuthFlow)`

SetImplicit sets Implicit field to given value.

### HasImplicit

`func (o *OAuthFlows) HasImplicit() bool`

HasImplicit returns a boolean if a field has been set.

### GetPassword

`func (o *OAuthFlows) GetPassword() OAuthFlow`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OAuthFlows) GetPasswordOk() (*OAuthFlow, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OAuthFlows) SetPassword(v OAuthFlow)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OAuthFlows) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


