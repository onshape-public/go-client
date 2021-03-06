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

// BTPlaneDescription692AllOf struct for BTPlaneDescription692AllOf
type BTPlaneDescription692AllOf struct {
	BtType *string `json:"btType,omitempty"`
	IsOrientedWithFace *bool `json:"isOrientedWithFace,omitempty"`
	Normal *BTVector3d389 `json:"normal,omitempty"`
	Origin *BTVector3d389 `json:"origin,omitempty"`
}

// NewBTPlaneDescription692AllOf instantiates a new BTPlaneDescription692AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPlaneDescription692AllOf() *BTPlaneDescription692AllOf {
	this := BTPlaneDescription692AllOf{}
	return &this
}

// NewBTPlaneDescription692AllOfWithDefaults instantiates a new BTPlaneDescription692AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPlaneDescription692AllOfWithDefaults() *BTPlaneDescription692AllOf {
	this := BTPlaneDescription692AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPlaneDescription692AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPlaneDescription692AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPlaneDescription692AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetIsOrientedWithFace returns the IsOrientedWithFace field value if set, zero value otherwise.
func (o *BTPlaneDescription692AllOf) GetIsOrientedWithFace() bool {
	if o == nil || o.IsOrientedWithFace == nil {
		var ret bool
		return ret
	}
	return *o.IsOrientedWithFace
}

// GetIsOrientedWithFaceOk returns a tuple with the IsOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692AllOf) GetIsOrientedWithFaceOk() (*bool, bool) {
	if o == nil || o.IsOrientedWithFace == nil {
		return nil, false
	}
	return o.IsOrientedWithFace, true
}

// HasIsOrientedWithFace returns a boolean if a field has been set.
func (o *BTPlaneDescription692AllOf) HasIsOrientedWithFace() bool {
	if o != nil && o.IsOrientedWithFace != nil {
		return true
	}

	return false
}

// SetIsOrientedWithFace gets a reference to the given bool and assigns it to the IsOrientedWithFace field.
func (o *BTPlaneDescription692AllOf) SetIsOrientedWithFace(v bool) {
	o.IsOrientedWithFace = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTPlaneDescription692AllOf) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692AllOf) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTPlaneDescription692AllOf) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTPlaneDescription692AllOf) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTPlaneDescription692AllOf) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPlaneDescription692AllOf) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTPlaneDescription692AllOf) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTPlaneDescription692AllOf) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

func (o BTPlaneDescription692AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsOrientedWithFace != nil {
		toSerialize["isOrientedWithFace"] = o.IsOrientedWithFace
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableBTPlaneDescription692AllOf struct {
	value *BTPlaneDescription692AllOf
	isSet bool
}

func (v NullableBTPlaneDescription692AllOf) Get() *BTPlaneDescription692AllOf {
	return v.value
}

func (v *NullableBTPlaneDescription692AllOf) Set(val *BTPlaneDescription692AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPlaneDescription692AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPlaneDescription692AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPlaneDescription692AllOf(val *BTPlaneDescription692AllOf) *NullableBTPlaneDescription692AllOf {
	return &NullableBTPlaneDescription692AllOf{value: val, isSet: true}
}

func (v NullableBTPlaneDescription692AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPlaneDescription692AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
