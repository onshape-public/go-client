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

// BTSweepDescription1473 struct for BTSweepDescription1473
type BTSweepDescription1473 struct {
	BTSurfaceDescription1564
	BtType                    *string                  `json:"btType,omitempty"`
	Direction                 *BTVector3d389           `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389           `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389           `json:"origin,omitempty"`
	Type                      *GBTSurfaceTypeEnum      `json:"type,omitempty"`
	Profile                   *BTSplineDescription2118 `json:"profile,omitempty"`
}

// NewBTSweepDescription1473 instantiates a new BTSweepDescription1473 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSweepDescription1473() *BTSweepDescription1473 {
	this := BTSweepDescription1473{}
	return &this
}

// NewBTSweepDescription1473WithDefaults instantiates a new BTSweepDescription1473 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSweepDescription1473WithDefaults() *BTSweepDescription1473 {
	this := BTSweepDescription1473{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSweepDescription1473) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSweepDescription1473) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSweepDescription1473) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSweepDescription1473) SetBtType(v string) {
	o.BtType = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTSweepDescription1473) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSweepDescription1473) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTSweepDescription1473) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTSweepDescription1473) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTSweepDescription1473) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSweepDescription1473) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTSweepDescription1473) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTSweepDescription1473) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTSweepDescription1473) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSweepDescription1473) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTSweepDescription1473) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTSweepDescription1473) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTSweepDescription1473) GetType() GBTSurfaceTypeEnum {
	if o == nil || o.Type == nil {
		var ret GBTSurfaceTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSweepDescription1473) GetTypeOk() (*GBTSurfaceTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTSweepDescription1473) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTSurfaceTypeEnum and assigns it to the Type field.
func (o *BTSweepDescription1473) SetType(v GBTSurfaceTypeEnum) {
	o.Type = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *BTSweepDescription1473) GetProfile() BTSplineDescription2118 {
	if o == nil || o.Profile == nil {
		var ret BTSplineDescription2118
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSweepDescription1473) GetProfileOk() (*BTSplineDescription2118, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *BTSweepDescription1473) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given BTSplineDescription2118 and assigns it to the Profile field.
func (o *BTSweepDescription1473) SetProfile(v BTSplineDescription2118) {
	o.Profile = &v
}

func (o BTSweepDescription1473) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTSurfaceDescription1564, errBTSurfaceDescription1564 := json.Marshal(o.BTSurfaceDescription1564)
	if errBTSurfaceDescription1564 != nil {
		return []byte{}, errBTSurfaceDescription1564
	}
	errBTSurfaceDescription1564 = json.Unmarshal([]byte(serializedBTSurfaceDescription1564), &toSerialize)
	if errBTSurfaceDescription1564 != nil {
		return []byte{}, errBTSurfaceDescription1564
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.DirectionOrientedWithFace != nil {
		toSerialize["directionOrientedWithFace"] = o.DirectionOrientedWithFace
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableBTSweepDescription1473 struct {
	value *BTSweepDescription1473
	isSet bool
}

func (v NullableBTSweepDescription1473) Get() *BTSweepDescription1473 {
	return v.value
}

func (v *NullableBTSweepDescription1473) Set(val *BTSweepDescription1473) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSweepDescription1473) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSweepDescription1473) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSweepDescription1473(val *BTSweepDescription1473) *NullableBTSweepDescription1473 {
	return &NullableBTSweepDescription1473{value: val, isSet: true}
}

func (v NullableBTSweepDescription1473) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSweepDescription1473) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
