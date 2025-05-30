/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTReferenceCustomData1551 - struct for BTReferenceCustomData1551
type BTReferenceCustomData1551 struct {
	implBTReferenceCustomData1551 interface{}
}

// BTInstanceStandardContentData2081AsBTReferenceCustomData1551 is a convenience function that returns BTInstanceStandardContentData2081 wrapped in BTReferenceCustomData1551
func (o *BTInstanceStandardContentData2081) AsBTReferenceCustomData1551() *BTReferenceCustomData1551 {
	return &BTReferenceCustomData1551{o}
}

// BTRevisionCustomData2090AsBTReferenceCustomData1551 is a convenience function that returns BTRevisionCustomData2090 wrapped in BTReferenceCustomData1551
func (o *BTRevisionCustomData2090) AsBTReferenceCustomData1551() *BTReferenceCustomData1551 {
	return &BTReferenceCustomData1551{o}
}

// NewBTReferenceCustomData1551 instantiates a new BTReferenceCustomData1551 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReferenceCustomData1551() *BTReferenceCustomData1551 {
	this := BTReferenceCustomData1551{Newbase_BTReferenceCustomData1551()}
	return &this
}

// NewBTReferenceCustomData1551WithDefaults instantiates a new BTReferenceCustomData1551 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReferenceCustomData1551WithDefaults() *BTReferenceCustomData1551 {
	this := BTReferenceCustomData1551{Newbase_BTReferenceCustomData1551WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTReferenceCustomData1551) GetBtType() string {
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
func (o *BTReferenceCustomData1551) GetBtTypeOk() (*string, bool) {
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
func (o *BTReferenceCustomData1551) HasBtType() bool {
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
func (o *BTReferenceCustomData1551) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTReferenceCustomData1551) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTInstanceStandardContentData-2081'
	if jsonDict["btType"] == "BTInstanceStandardContentData-2081" {
		// try to unmarshal JSON data into BTInstanceStandardContentData2081
		var qr *BTInstanceStandardContentData2081
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTReferenceCustomData1551 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTReferenceCustomData1551 = nil
			return fmt.Errorf("failed to unmarshal BTReferenceCustomData1551 as BTInstanceStandardContentData2081: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTRevisionCustomData-2090'
	if jsonDict["btType"] == "BTRevisionCustomData-2090" {
		// try to unmarshal JSON data into BTRevisionCustomData2090
		var qr *BTRevisionCustomData2090
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTReferenceCustomData1551 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTReferenceCustomData1551 = nil
			return fmt.Errorf("failed to unmarshal BTReferenceCustomData1551 as BTRevisionCustomData2090: %s", err.Error())
		}
	}

	var qtx *base_BTReferenceCustomData1551
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTReferenceCustomData1551 = qtx
		return nil // data stored in dst.base_BTReferenceCustomData1551, return on the first match
	} else {
		dst.implBTReferenceCustomData1551 = nil
		return fmt.Errorf("failed to unmarshal BTReferenceCustomData1551 as base_BTReferenceCustomData1551: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTReferenceCustomData1551) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTReferenceCustomData1551) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTReferenceCustomData1551
}

type NullableBTReferenceCustomData1551 struct {
	value *BTReferenceCustomData1551
	isSet bool
}

func (v NullableBTReferenceCustomData1551) Get() *BTReferenceCustomData1551 {
	return v.value
}

func (v *NullableBTReferenceCustomData1551) Set(val *BTReferenceCustomData1551) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReferenceCustomData1551) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReferenceCustomData1551) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReferenceCustomData1551(val *BTReferenceCustomData1551) *NullableBTReferenceCustomData1551 {
	return &NullableBTReferenceCustomData1551{value: val, isSet: true}
}

func (v NullableBTReferenceCustomData1551) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReferenceCustomData1551) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTReferenceCustomData1551 struct {
	// Type of JSON object.
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTReferenceCustomData1551 instantiates a new base_BTReferenceCustomData1551 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTReferenceCustomData1551() *base_BTReferenceCustomData1551 {
	this := base_BTReferenceCustomData1551{}
	return &this
}

// Newbase_BTReferenceCustomData1551WithDefaults instantiates a new base_BTReferenceCustomData1551 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTReferenceCustomData1551WithDefaults() *base_BTReferenceCustomData1551 {
	this := base_BTReferenceCustomData1551{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTReferenceCustomData1551) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTReferenceCustomData1551) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTReferenceCustomData1551) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTReferenceCustomData1551) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTReferenceCustomData1551) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
