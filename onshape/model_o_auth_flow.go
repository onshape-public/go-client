/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// OAuthFlow struct for OAuthFlow
type OAuthFlow struct {
	AuthorizationUrl *string                `json:"authorizationUrl,omitempty"`
	Extensions       map[string]interface{} `json:"extensions,omitempty"`
	RefreshUrl       *string                `json:"refreshUrl,omitempty"`
	Scopes           *map[string]string     `json:"scopes,omitempty"`
	TokenUrl         *string                `json:"tokenUrl,omitempty"`
}

// NewOAuthFlow instantiates a new OAuthFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthFlow() *OAuthFlow {
	this := OAuthFlow{}
	return &this
}

// NewOAuthFlowWithDefaults instantiates a new OAuthFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthFlowWithDefaults() *OAuthFlow {
	this := OAuthFlow{}
	return &this
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise.
func (o *OAuthFlow) GetAuthorizationUrl() string {
	if o == nil || o.AuthorizationUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil || o.AuthorizationUrl == nil {
		return nil, false
	}
	return o.AuthorizationUrl, true
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *OAuthFlow) HasAuthorizationUrl() bool {
	if o != nil && o.AuthorizationUrl != nil {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given string and assigns it to the AuthorizationUrl field.
func (o *OAuthFlow) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *OAuthFlow) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetExtensionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *OAuthFlow) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *OAuthFlow) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetRefreshUrl returns the RefreshUrl field value if set, zero value otherwise.
func (o *OAuthFlow) GetRefreshUrl() string {
	if o == nil || o.RefreshUrl == nil {
		var ret string
		return ret
	}
	return *o.RefreshUrl
}

// GetRefreshUrlOk returns a tuple with the RefreshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetRefreshUrlOk() (*string, bool) {
	if o == nil || o.RefreshUrl == nil {
		return nil, false
	}
	return o.RefreshUrl, true
}

// HasRefreshUrl returns a boolean if a field has been set.
func (o *OAuthFlow) HasRefreshUrl() bool {
	if o != nil && o.RefreshUrl != nil {
		return true
	}

	return false
}

// SetRefreshUrl gets a reference to the given string and assigns it to the RefreshUrl field.
func (o *OAuthFlow) SetRefreshUrl(v string) {
	o.RefreshUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuthFlow) GetScopes() map[string]string {
	if o == nil || o.Scopes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetScopesOk() (*map[string]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuthFlow) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given map[string]string and assigns it to the Scopes field.
func (o *OAuthFlow) SetScopes(v map[string]string) {
	o.Scopes = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *OAuthFlow) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *OAuthFlow) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *OAuthFlow) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

func (o OAuthFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationUrl != nil {
		toSerialize["authorizationUrl"] = o.AuthorizationUrl
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.RefreshUrl != nil {
		toSerialize["refreshUrl"] = o.RefreshUrl
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.TokenUrl != nil {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthFlow struct {
	value *OAuthFlow
	isSet bool
}

func (v NullableOAuthFlow) Get() *OAuthFlow {
	return v.value
}

func (v *NullableOAuthFlow) Set(val *OAuthFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthFlow(val *OAuthFlow) *NullableOAuthFlow {
	return &NullableOAuthFlow{value: val, isSet: true}
}

func (v NullableOAuthFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
