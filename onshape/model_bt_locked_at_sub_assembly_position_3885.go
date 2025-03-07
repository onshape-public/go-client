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

// BTLockedAtSubAssemblyPosition3885 struct for BTLockedAtSubAssemblyPosition3885
type BTLockedAtSubAssemblyPosition3885 struct {
	BTLockedSubAssembly4590
	BtType                      *string                                   `json:"btType,omitempty"`
	LockType                    *GBTSubAssemblyLockType                   `json:"lockType,omitempty"`
	LockedSubAssemblyOutputInfo *BTRigidOrLockedSubAssemblyOutputInfo3860 `json:"lockedSubAssemblyOutputInfo,omitempty"`
}

// NewBTLockedAtSubAssemblyPosition3885 instantiates a new BTLockedAtSubAssemblyPosition3885 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLockedAtSubAssemblyPosition3885() *BTLockedAtSubAssemblyPosition3885 {
	this := BTLockedAtSubAssemblyPosition3885{}
	return &this
}

// NewBTLockedAtSubAssemblyPosition3885WithDefaults instantiates a new BTLockedAtSubAssemblyPosition3885 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLockedAtSubAssemblyPosition3885WithDefaults() *BTLockedAtSubAssemblyPosition3885 {
	this := BTLockedAtSubAssemblyPosition3885{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLockedAtSubAssemblyPosition3885) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtSubAssemblyPosition3885) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLockedAtSubAssemblyPosition3885) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLockedAtSubAssemblyPosition3885) SetBtType(v string) {
	o.BtType = &v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *BTLockedAtSubAssemblyPosition3885) GetLockType() GBTSubAssemblyLockType {
	if o == nil || o.LockType == nil {
		var ret GBTSubAssemblyLockType
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtSubAssemblyPosition3885) GetLockTypeOk() (*GBTSubAssemblyLockType, bool) {
	if o == nil || o.LockType == nil {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *BTLockedAtSubAssemblyPosition3885) HasLockType() bool {
	if o != nil && o.LockType != nil {
		return true
	}

	return false
}

// SetLockType gets a reference to the given GBTSubAssemblyLockType and assigns it to the LockType field.
func (o *BTLockedAtSubAssemblyPosition3885) SetLockType(v GBTSubAssemblyLockType) {
	o.LockType = &v
}

// GetLockedSubAssemblyOutputInfo returns the LockedSubAssemblyOutputInfo field value if set, zero value otherwise.
func (o *BTLockedAtSubAssemblyPosition3885) GetLockedSubAssemblyOutputInfo() BTRigidOrLockedSubAssemblyOutputInfo3860 {
	if o == nil || o.LockedSubAssemblyOutputInfo == nil {
		var ret BTRigidOrLockedSubAssemblyOutputInfo3860
		return ret
	}
	return *o.LockedSubAssemblyOutputInfo
}

// GetLockedSubAssemblyOutputInfoOk returns a tuple with the LockedSubAssemblyOutputInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLockedAtSubAssemblyPosition3885) GetLockedSubAssemblyOutputInfoOk() (*BTRigidOrLockedSubAssemblyOutputInfo3860, bool) {
	if o == nil || o.LockedSubAssemblyOutputInfo == nil {
		return nil, false
	}
	return o.LockedSubAssemblyOutputInfo, true
}

// HasLockedSubAssemblyOutputInfo returns a boolean if a field has been set.
func (o *BTLockedAtSubAssemblyPosition3885) HasLockedSubAssemblyOutputInfo() bool {
	if o != nil && o.LockedSubAssemblyOutputInfo != nil {
		return true
	}

	return false
}

// SetLockedSubAssemblyOutputInfo gets a reference to the given BTRigidOrLockedSubAssemblyOutputInfo3860 and assigns it to the LockedSubAssemblyOutputInfo field.
func (o *BTLockedAtSubAssemblyPosition3885) SetLockedSubAssemblyOutputInfo(v BTRigidOrLockedSubAssemblyOutputInfo3860) {
	o.LockedSubAssemblyOutputInfo = &v
}

func (o BTLockedAtSubAssemblyPosition3885) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTLockedSubAssembly4590, errBTLockedSubAssembly4590 := json.Marshal(o.BTLockedSubAssembly4590)
	if errBTLockedSubAssembly4590 != nil {
		return []byte{}, errBTLockedSubAssembly4590
	}
	errBTLockedSubAssembly4590 = json.Unmarshal([]byte(serializedBTLockedSubAssembly4590), &toSerialize)
	if errBTLockedSubAssembly4590 != nil {
		return []byte{}, errBTLockedSubAssembly4590
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.LockType != nil {
		toSerialize["lockType"] = o.LockType
	}
	if o.LockedSubAssemblyOutputInfo != nil {
		toSerialize["lockedSubAssemblyOutputInfo"] = o.LockedSubAssemblyOutputInfo
	}
	return json.Marshal(toSerialize)
}

type NullableBTLockedAtSubAssemblyPosition3885 struct {
	value *BTLockedAtSubAssemblyPosition3885
	isSet bool
}

func (v NullableBTLockedAtSubAssemblyPosition3885) Get() *BTLockedAtSubAssemblyPosition3885 {
	return v.value
}

func (v *NullableBTLockedAtSubAssemblyPosition3885) Set(val *BTLockedAtSubAssemblyPosition3885) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLockedAtSubAssemblyPosition3885) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLockedAtSubAssemblyPosition3885) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLockedAtSubAssemblyPosition3885(val *BTLockedAtSubAssemblyPosition3885) *NullableBTLockedAtSubAssemblyPosition3885 {
	return &NullableBTLockedAtSubAssemblyPosition3885{value: val, isSet: true}
}

func (v NullableBTLockedAtSubAssemblyPosition3885) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLockedAtSubAssemblyPosition3885) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
