/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.5998-d3227e94fd7e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNonAlignedBoundingBox4180 struct for BTNonAlignedBoundingBox4180
type BTNonAlignedBoundingBox4180 struct {
	MaxCorner        *BTVector3d389         `json:"maxCorner,omitempty"`
	MinCorner        *BTVector3d389         `json:"minCorner,omitempty"`
	Valid            *bool                  `json:"valid,omitempty"`
	BtType           *string                `json:"btType,omitempty"`
	CoordinateSystem *BTCoordinateSystem387 `json:"coordinateSystem,omitempty"`
}

// NewBTNonAlignedBoundingBox4180 instantiates a new BTNonAlignedBoundingBox4180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNonAlignedBoundingBox4180() *BTNonAlignedBoundingBox4180 {
	this := BTNonAlignedBoundingBox4180{}
	return &this
}

// NewBTNonAlignedBoundingBox4180WithDefaults instantiates a new BTNonAlignedBoundingBox4180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNonAlignedBoundingBox4180WithDefaults() *BTNonAlignedBoundingBox4180 {
	this := BTNonAlignedBoundingBox4180{}
	return &this
}

// GetMaxCorner returns the MaxCorner field value if set, zero value otherwise.
func (o *BTNonAlignedBoundingBox4180) GetMaxCorner() BTVector3d389 {
	if o == nil || o.MaxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MaxCorner
}

// GetMaxCornerOk returns a tuple with the MaxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNonAlignedBoundingBox4180) GetMaxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.MaxCorner == nil {
		return nil, false
	}
	return o.MaxCorner, true
}

// HasMaxCorner returns a boolean if a field has been set.
func (o *BTNonAlignedBoundingBox4180) HasMaxCorner() bool {
	if o != nil && o.MaxCorner != nil {
		return true
	}

	return false
}

// SetMaxCorner gets a reference to the given BTVector3d389 and assigns it to the MaxCorner field.
func (o *BTNonAlignedBoundingBox4180) SetMaxCorner(v BTVector3d389) {
	o.MaxCorner = &v
}

// GetMinCorner returns the MinCorner field value if set, zero value otherwise.
func (o *BTNonAlignedBoundingBox4180) GetMinCorner() BTVector3d389 {
	if o == nil || o.MinCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MinCorner
}

// GetMinCornerOk returns a tuple with the MinCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNonAlignedBoundingBox4180) GetMinCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.MinCorner == nil {
		return nil, false
	}
	return o.MinCorner, true
}

// HasMinCorner returns a boolean if a field has been set.
func (o *BTNonAlignedBoundingBox4180) HasMinCorner() bool {
	if o != nil && o.MinCorner != nil {
		return true
	}

	return false
}

// SetMinCorner gets a reference to the given BTVector3d389 and assigns it to the MinCorner field.
func (o *BTNonAlignedBoundingBox4180) SetMinCorner(v BTVector3d389) {
	o.MinCorner = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTNonAlignedBoundingBox4180) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNonAlignedBoundingBox4180) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BTNonAlignedBoundingBox4180) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTNonAlignedBoundingBox4180) SetValid(v bool) {
	o.Valid = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNonAlignedBoundingBox4180) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNonAlignedBoundingBox4180) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNonAlignedBoundingBox4180) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNonAlignedBoundingBox4180) SetBtType(v string) {
	o.BtType = &v
}

// GetCoordinateSystem returns the CoordinateSystem field value if set, zero value otherwise.
func (o *BTNonAlignedBoundingBox4180) GetCoordinateSystem() BTCoordinateSystem387 {
	if o == nil || o.CoordinateSystem == nil {
		var ret BTCoordinateSystem387
		return ret
	}
	return *o.CoordinateSystem
}

// GetCoordinateSystemOk returns a tuple with the CoordinateSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNonAlignedBoundingBox4180) GetCoordinateSystemOk() (*BTCoordinateSystem387, bool) {
	if o == nil || o.CoordinateSystem == nil {
		return nil, false
	}
	return o.CoordinateSystem, true
}

// HasCoordinateSystem returns a boolean if a field has been set.
func (o *BTNonAlignedBoundingBox4180) HasCoordinateSystem() bool {
	if o != nil && o.CoordinateSystem != nil {
		return true
	}

	return false
}

// SetCoordinateSystem gets a reference to the given BTCoordinateSystem387 and assigns it to the CoordinateSystem field.
func (o *BTNonAlignedBoundingBox4180) SetCoordinateSystem(v BTCoordinateSystem387) {
	o.CoordinateSystem = &v
}

func (o BTNonAlignedBoundingBox4180) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxCorner != nil {
		toSerialize["maxCorner"] = o.MaxCorner
	}
	if o.MinCorner != nil {
		toSerialize["minCorner"] = o.MinCorner
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CoordinateSystem != nil {
		toSerialize["coordinateSystem"] = o.CoordinateSystem
	}
	return json.Marshal(toSerialize)
}

type NullableBTNonAlignedBoundingBox4180 struct {
	value *BTNonAlignedBoundingBox4180
	isSet bool
}

func (v NullableBTNonAlignedBoundingBox4180) Get() *BTNonAlignedBoundingBox4180 {
	return v.value
}

func (v *NullableBTNonAlignedBoundingBox4180) Set(val *BTNonAlignedBoundingBox4180) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNonAlignedBoundingBox4180) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNonAlignedBoundingBox4180) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNonAlignedBoundingBox4180(val *BTNonAlignedBoundingBox4180) *NullableBTNonAlignedBoundingBox4180 {
	return &NullableBTNonAlignedBoundingBox4180{value: val, isSet: true}
}

func (v NullableBTNonAlignedBoundingBox4180) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNonAlignedBoundingBox4180) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}