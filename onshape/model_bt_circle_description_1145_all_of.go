/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
)

// BTCircleDescription1145AllOf struct for BTCircleDescription1145AllOf
type BTCircleDescription1145AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Normal *BTVector3d389 `json:"normal,omitempty"`
	Origin *BTVector3d389 `json:"origin,omitempty"`
	Radius *float64 `json:"radius,omitempty"`
}

// NewBTCircleDescription1145AllOf instantiates a new BTCircleDescription1145AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCircleDescription1145AllOf() *BTCircleDescription1145AllOf {
	this := BTCircleDescription1145AllOf{}
	return &this
}

// NewBTCircleDescription1145AllOfWithDefaults instantiates a new BTCircleDescription1145AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCircleDescription1145AllOfWithDefaults() *BTCircleDescription1145AllOf {
	this := BTCircleDescription1145AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCircleDescription1145AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCircleDescription1145AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCircleDescription1145AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCircleDescription1145AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTCircleDescription1145AllOf) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCircleDescription1145AllOf) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTCircleDescription1145AllOf) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTCircleDescription1145AllOf) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTCircleDescription1145AllOf) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCircleDescription1145AllOf) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTCircleDescription1145AllOf) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTCircleDescription1145AllOf) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTCircleDescription1145AllOf) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCircleDescription1145AllOf) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTCircleDescription1145AllOf) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTCircleDescription1145AllOf) SetRadius(v float64) {
	o.Radius = &v
}

func (o BTCircleDescription1145AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableBTCircleDescription1145AllOf struct {
	value *BTCircleDescription1145AllOf
	isSet bool
}

func (v NullableBTCircleDescription1145AllOf) Get() *BTCircleDescription1145AllOf {
	return v.value
}

func (v *NullableBTCircleDescription1145AllOf) Set(val *BTCircleDescription1145AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCircleDescription1145AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCircleDescription1145AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCircleDescription1145AllOf(val *BTCircleDescription1145AllOf) *NullableBTCircleDescription1145AllOf {
	return &NullableBTCircleDescription1145AllOf{value: val, isSet: true}
}

func (v NullableBTCircleDescription1145AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCircleDescription1145AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
