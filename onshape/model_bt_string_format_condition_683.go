/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTStringFormatCondition683 - struct for BTStringFormatCondition683
type BTStringFormatCondition683 struct {
	implBTStringFormatCondition683 interface{}
}

// BTStringMaximumLengthPattern2593AsBTStringFormatCondition683 is a convenience function that returns BTStringMaximumLengthPattern2593 wrapped in BTStringFormatCondition683
func (o *BTStringMaximumLengthPattern2593) AsBTStringFormatCondition683() *BTStringFormatCondition683 {
	return &BTStringFormatCondition683{o}
}

// BTStringMinimumLengthPattern895AsBTStringFormatCondition683 is a convenience function that returns BTStringMinimumLengthPattern895 wrapped in BTStringFormatCondition683
func (o *BTStringMinimumLengthPattern895) AsBTStringFormatCondition683() *BTStringFormatCondition683 {
	return &BTStringFormatCondition683{o}
}

// BTStringFormatMatchPattern2446AsBTStringFormatCondition683 is a convenience function that returns BTStringFormatMatchPattern2446 wrapped in BTStringFormatCondition683
func (o *BTStringFormatMatchPattern2446) AsBTStringFormatCondition683() *BTStringFormatCondition683 {
	return &BTStringFormatCondition683{o}
}

// BTStringFormatBlockPattern1755AsBTStringFormatCondition683 is a convenience function that returns BTStringFormatBlockPattern1755 wrapped in BTStringFormatCondition683
func (o *BTStringFormatBlockPattern1755) AsBTStringFormatCondition683() *BTStringFormatCondition683 {
	return &BTStringFormatCondition683{o}
}

// NewBTStringFormatCondition683 instantiates a new BTStringFormatCondition683 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringFormatCondition683() *BTStringFormatCondition683 {
	this := BTStringFormatCondition683{Newbase_BTStringFormatCondition683()}
	return &this
}

// NewBTStringFormatCondition683WithDefaults instantiates a new BTStringFormatCondition683 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringFormatCondition683WithDefaults() *BTStringFormatCondition683 {
	this := BTStringFormatCondition683{Newbase_BTStringFormatCondition683WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringFormatCondition683) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatCondition683) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringFormatCondition683) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringFormatCondition683) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTStringFormatCondition683) GetErrorMessage() string {
	type getResult interface {
		GetErrorMessage() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorMessage()
	} else {
		var de string
		return de
	}
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatCondition683) GetErrorMessageOk() (*string, bool) {
	type getResult interface {
		GetErrorMessageOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorMessageOk()
	} else {
		return nil, false
	}
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTStringFormatCondition683) HasErrorMessage() bool {
	type getResult interface {
		HasErrorMessage() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasErrorMessage()
	} else {
		return false
	}
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTStringFormatCondition683) SetErrorMessage(v string) {
	type getResult interface {
		SetErrorMessage(v string)
	}

	o.GetActualInstance().(getResult).SetErrorMessage(v)
}

// GetShouldResetValueWhenConfirmed returns the ShouldResetValueWhenConfirmed field value if set, zero value otherwise.
func (o *BTStringFormatCondition683) GetShouldResetValueWhenConfirmed() bool {
	type getResult interface {
		GetShouldResetValueWhenConfirmed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetShouldResetValueWhenConfirmed()
	} else {
		var de bool
		return de
	}
}

// GetShouldResetValueWhenConfirmedOk returns a tuple with the ShouldResetValueWhenConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatCondition683) GetShouldResetValueWhenConfirmedOk() (*bool, bool) {
	type getResult interface {
		GetShouldResetValueWhenConfirmedOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetShouldResetValueWhenConfirmedOk()
	} else {
		return nil, false
	}
}

// HasShouldResetValueWhenConfirmed returns a boolean if a field has been set.
func (o *BTStringFormatCondition683) HasShouldResetValueWhenConfirmed() bool {
	type getResult interface {
		HasShouldResetValueWhenConfirmed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasShouldResetValueWhenConfirmed()
	} else {
		return false
	}
}

// SetShouldResetValueWhenConfirmed gets a reference to the given bool and assigns it to the ShouldResetValueWhenConfirmed field.
func (o *BTStringFormatCondition683) SetShouldResetValueWhenConfirmed(v bool) {
	type getResult interface {
		SetShouldResetValueWhenConfirmed(v bool)
	}

	o.GetActualInstance().(getResult).SetShouldResetValueWhenConfirmed(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTStringFormatCondition683) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTStringFormatBlockPattern-1755'
	if jsonDict["btType"] == "BTStringFormatBlockPattern-1755" {
		// try to unmarshal JSON data into BTStringFormatBlockPattern1755
		var qr *BTStringFormatBlockPattern1755
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTStringFormatCondition683 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTStringFormatCondition683 = nil
			return fmt.Errorf("Failed to unmarshal BTStringFormatCondition683 as BTStringFormatBlockPattern1755: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTStringFormatMatchPattern-2446'
	if jsonDict["btType"] == "BTStringFormatMatchPattern-2446" {
		// try to unmarshal JSON data into BTStringFormatMatchPattern2446
		var qr *BTStringFormatMatchPattern2446
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTStringFormatCondition683 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTStringFormatCondition683 = nil
			return fmt.Errorf("Failed to unmarshal BTStringFormatCondition683 as BTStringFormatMatchPattern2446: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTStringMaximumLengthPattern-2593'
	if jsonDict["btType"] == "BTStringMaximumLengthPattern-2593" {
		// try to unmarshal JSON data into BTStringMaximumLengthPattern2593
		var qr *BTStringMaximumLengthPattern2593
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTStringFormatCondition683 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTStringFormatCondition683 = nil
			return fmt.Errorf("Failed to unmarshal BTStringFormatCondition683 as BTStringMaximumLengthPattern2593: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTStringMinimumLengthPattern-895'
	if jsonDict["btType"] == "BTStringMinimumLengthPattern-895" {
		// try to unmarshal JSON data into BTStringMinimumLengthPattern895
		var qr *BTStringMinimumLengthPattern895
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTStringFormatCondition683 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTStringFormatCondition683 = nil
			return fmt.Errorf("Failed to unmarshal BTStringFormatCondition683 as BTStringMinimumLengthPattern895: %s", err.Error())
		}
	}

	var qtx *base_BTStringFormatCondition683
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTStringFormatCondition683 = qtx
		return nil // data stored in dst.base_BTStringFormatCondition683, return on the first match
	} else {
		dst.implBTStringFormatCondition683 = nil
		return fmt.Errorf("Failed to unmarshal BTStringFormatCondition683 as base_BTStringFormatCondition683: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTStringFormatCondition683) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTStringFormatCondition683) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTStringFormatCondition683
}

type NullableBTStringFormatCondition683 struct {
	value *BTStringFormatCondition683
	isSet bool
}

func (v NullableBTStringFormatCondition683) Get() *BTStringFormatCondition683 {
	return v.value
}

func (v *NullableBTStringFormatCondition683) Set(val *BTStringFormatCondition683) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringFormatCondition683) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringFormatCondition683) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringFormatCondition683(val *BTStringFormatCondition683) *NullableBTStringFormatCondition683 {
	return &NullableBTStringFormatCondition683{value: val, isSet: true}
}

func (v NullableBTStringFormatCondition683) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringFormatCondition683) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTStringFormatCondition683 struct {
	BtType                        *string `json:"btType,omitempty"`
	ErrorMessage                  *string `json:"errorMessage,omitempty"`
	ShouldResetValueWhenConfirmed *bool   `json:"shouldResetValueWhenConfirmed,omitempty"`
}

// Newbase_BTStringFormatCondition683 instantiates a new base_BTStringFormatCondition683 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTStringFormatCondition683() *base_BTStringFormatCondition683 {
	this := base_BTStringFormatCondition683{}
	return &this
}

// Newbase_BTStringFormatCondition683WithDefaults instantiates a new base_BTStringFormatCondition683 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTStringFormatCondition683WithDefaults() *base_BTStringFormatCondition683 {
	this := base_BTStringFormatCondition683{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTStringFormatCondition683) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTStringFormatCondition683) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTStringFormatCondition683) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTStringFormatCondition683) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *base_BTStringFormatCondition683) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTStringFormatCondition683) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *base_BTStringFormatCondition683) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *base_BTStringFormatCondition683) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetShouldResetValueWhenConfirmed returns the ShouldResetValueWhenConfirmed field value if set, zero value otherwise.
func (o *base_BTStringFormatCondition683) GetShouldResetValueWhenConfirmed() bool {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.ShouldResetValueWhenConfirmed
}

// GetShouldResetValueWhenConfirmedOk returns a tuple with the ShouldResetValueWhenConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTStringFormatCondition683) GetShouldResetValueWhenConfirmedOk() (*bool, bool) {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		return nil, false
	}
	return o.ShouldResetValueWhenConfirmed, true
}

// HasShouldResetValueWhenConfirmed returns a boolean if a field has been set.
func (o *base_BTStringFormatCondition683) HasShouldResetValueWhenConfirmed() bool {
	if o != nil && o.ShouldResetValueWhenConfirmed != nil {
		return true
	}

	return false
}

// SetShouldResetValueWhenConfirmed gets a reference to the given bool and assigns it to the ShouldResetValueWhenConfirmed field.
func (o *base_BTStringFormatCondition683) SetShouldResetValueWhenConfirmed(v bool) {
	o.ShouldResetValueWhenConfirmed = &v
}

func (o base_BTStringFormatCondition683) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ShouldResetValueWhenConfirmed != nil {
		toSerialize["shouldResetValueWhenConfirmed"] = o.ShouldResetValueWhenConfirmed
	}
	return json.Marshal(toSerialize)
}