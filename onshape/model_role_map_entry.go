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

// RoleMapEntry struct for RoleMapEntry
type RoleMapEntry struct {
	Identities []BTIdentityInfo `json:"identities,omitempty"`
	Role       *BTRbacRoleInfo  `json:"role,omitempty"`
}

// NewRoleMapEntry instantiates a new RoleMapEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMapEntry() *RoleMapEntry {
	this := RoleMapEntry{}
	return &this
}

// NewRoleMapEntryWithDefaults instantiates a new RoleMapEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMapEntryWithDefaults() *RoleMapEntry {
	this := RoleMapEntry{}
	return &this
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *RoleMapEntry) GetIdentities() []BTIdentityInfo {
	if o == nil || o.Identities == nil {
		var ret []BTIdentityInfo
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapEntry) GetIdentitiesOk() ([]BTIdentityInfo, bool) {
	if o == nil || o.Identities == nil {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *RoleMapEntry) HasIdentities() bool {
	if o != nil && o.Identities != nil {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []BTIdentityInfo and assigns it to the Identities field.
func (o *RoleMapEntry) SetIdentities(v []BTIdentityInfo) {
	o.Identities = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleMapEntry) GetRole() BTRbacRoleInfo {
	if o == nil || o.Role == nil {
		var ret BTRbacRoleInfo
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapEntry) GetRoleOk() (*BTRbacRoleInfo, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleMapEntry) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given BTRbacRoleInfo and assigns it to the Role field.
func (o *RoleMapEntry) SetRole(v BTRbacRoleInfo) {
	o.Role = &v
}

func (o RoleMapEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableRoleMapEntry struct {
	value *RoleMapEntry
	isSet bool
}

func (v NullableRoleMapEntry) Get() *RoleMapEntry {
	return v.value
}

func (v *NullableRoleMapEntry) Set(val *RoleMapEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMapEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMapEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMapEntry(val *RoleMapEntry) *NullableRoleMapEntry {
	return &NullableRoleMapEntry{value: val, isSet: true}
}

func (v NullableRoleMapEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMapEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
