/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCurveGeometryLine117 struct for BTCurveGeometryLine117
type BTCurveGeometryLine117 struct {
	BtType     *string              `json:"btType,omitempty"`
	EntityType *GBTSketchEntityType `json:"entityType,omitempty"`
	DirX       *float64             `json:"dirX,omitempty"`
	DirY       *float64             `json:"dirY,omitempty"`
	PntX       *float64             `json:"pntX,omitempty"`
	PntY       *float64             `json:"pntY,omitempty"`
}

// NewBTCurveGeometryLine117 instantiates a new BTCurveGeometryLine117 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryLine117() *BTCurveGeometryLine117 {
	this := BTCurveGeometryLine117{}
	return &this
}

// NewBTCurveGeometryLine117WithDefaults instantiates a new BTCurveGeometryLine117 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryLine117WithDefaults() *BTCurveGeometryLine117 {
	this := BTCurveGeometryLine117{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryLine117) SetBtType(v string) {
	o.BtType = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetEntityType() GBTSketchEntityType {
	if o == nil || o.EntityType == nil {
		var ret GBTSketchEntityType
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetEntityTypeOk() (*GBTSketchEntityType, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given GBTSketchEntityType and assigns it to the EntityType field.
func (o *BTCurveGeometryLine117) SetEntityType(v GBTSketchEntityType) {
	o.EntityType = &v
}

// GetDirX returns the DirX field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetDirX() float64 {
	if o == nil || o.DirX == nil {
		var ret float64
		return ret
	}
	return *o.DirX
}

// GetDirXOk returns a tuple with the DirX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetDirXOk() (*float64, bool) {
	if o == nil || o.DirX == nil {
		return nil, false
	}
	return o.DirX, true
}

// HasDirX returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasDirX() bool {
	if o != nil && o.DirX != nil {
		return true
	}

	return false
}

// SetDirX gets a reference to the given float64 and assigns it to the DirX field.
func (o *BTCurveGeometryLine117) SetDirX(v float64) {
	o.DirX = &v
}

// GetDirY returns the DirY field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetDirY() float64 {
	if o == nil || o.DirY == nil {
		var ret float64
		return ret
	}
	return *o.DirY
}

// GetDirYOk returns a tuple with the DirY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetDirYOk() (*float64, bool) {
	if o == nil || o.DirY == nil {
		return nil, false
	}
	return o.DirY, true
}

// HasDirY returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasDirY() bool {
	if o != nil && o.DirY != nil {
		return true
	}

	return false
}

// SetDirY gets a reference to the given float64 and assigns it to the DirY field.
func (o *BTCurveGeometryLine117) SetDirY(v float64) {
	o.DirY = &v
}

// GetPntX returns the PntX field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetPntX() float64 {
	if o == nil || o.PntX == nil {
		var ret float64
		return ret
	}
	return *o.PntX
}

// GetPntXOk returns a tuple with the PntX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetPntXOk() (*float64, bool) {
	if o == nil || o.PntX == nil {
		return nil, false
	}
	return o.PntX, true
}

// HasPntX returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasPntX() bool {
	if o != nil && o.PntX != nil {
		return true
	}

	return false
}

// SetPntX gets a reference to the given float64 and assigns it to the PntX field.
func (o *BTCurveGeometryLine117) SetPntX(v float64) {
	o.PntX = &v
}

// GetPntY returns the PntY field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetPntY() float64 {
	if o == nil || o.PntY == nil {
		var ret float64
		return ret
	}
	return *o.PntY
}

// GetPntYOk returns a tuple with the PntY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetPntYOk() (*float64, bool) {
	if o == nil || o.PntY == nil {
		return nil, false
	}
	return o.PntY, true
}

// HasPntY returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasPntY() bool {
	if o != nil && o.PntY != nil {
		return true
	}

	return false
}

// SetPntY gets a reference to the given float64 and assigns it to the PntY field.
func (o *BTCurveGeometryLine117) SetPntY(v float64) {
	o.PntY = &v
}

func (o BTCurveGeometryLine117) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.DirX != nil {
		toSerialize["dirX"] = o.DirX
	}
	if o.DirY != nil {
		toSerialize["dirY"] = o.DirY
	}
	if o.PntX != nil {
		toSerialize["pntX"] = o.PntX
	}
	if o.PntY != nil {
		toSerialize["pntY"] = o.PntY
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryLine117 struct {
	value *BTCurveGeometryLine117
	isSet bool
}

func (v NullableBTCurveGeometryLine117) Get() *BTCurveGeometryLine117 {
	return v.value
}

func (v *NullableBTCurveGeometryLine117) Set(val *BTCurveGeometryLine117) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryLine117) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryLine117) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryLine117(val *BTCurveGeometryLine117) *NullableBTCurveGeometryLine117 {
	return &NullableBTCurveGeometryLine117{value: val, isSet: true}
}

func (v NullableBTCurveGeometryLine117) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryLine117) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
