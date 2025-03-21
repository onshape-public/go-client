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

// BTMassPropertiesInfo Mass properties information.
type BTMassPropertiesInfo struct {
	// Centroid, center of gravity, center of mass
	Centroid []float64 `json:"centroid,omitempty"`
	// `true` if the part has mass.
	HasMass_ *bool `json:"hasMass,omitempty"`
	// Mass moments of inertia
	Inertia []float64 `json:"inertia,omitempty"`
	// Mass
	Mass []float64 `json:"mass,omitempty"`
	// Number of parts without mass.
	MassMissingCount *int32 `json:"massMissingCount,omitempty"`
	// Surface area
	Periphery []float64 `json:"periphery,omitempty"`
	// Vector coordinates of the principal axes. Use `BTVector3d-389` as the `btType`.
	PrincipalAxes []BTVector3d389 `json:"principalAxes,omitempty"`
	// Principal moments of inertia
	PrincipalInertia []float64 `json:"principalInertia,omitempty"`
	// Volume
	Volume []float64 `json:"volume,omitempty"`
}

// NewBTMassPropertiesInfo instantiates a new BTMassPropertiesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMassPropertiesInfo() *BTMassPropertiesInfo {
	this := BTMassPropertiesInfo{}
	return &this
}

// NewBTMassPropertiesInfoWithDefaults instantiates a new BTMassPropertiesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMassPropertiesInfoWithDefaults() *BTMassPropertiesInfo {
	this := BTMassPropertiesInfo{}
	return &this
}

// GetCentroid returns the Centroid field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetCentroid() []float64 {
	if o == nil || o.Centroid == nil {
		var ret []float64
		return ret
	}
	return o.Centroid
}

// GetCentroidOk returns a tuple with the Centroid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetCentroidOk() ([]float64, bool) {
	if o == nil || o.Centroid == nil {
		return nil, false
	}
	return o.Centroid, true
}

// HasCentroid returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasCentroid() bool {
	if o != nil && o.Centroid != nil {
		return true
	}

	return false
}

// SetCentroid gets a reference to the given []float64 and assigns it to the Centroid field.
func (o *BTMassPropertiesInfo) SetCentroid(v []float64) {
	o.Centroid = v
}

// GetHasMass_ returns the HasMass_ field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetHasMass_() bool {
	if o == nil || o.HasMass_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMass_
}

// GetHasMass_Ok returns a tuple with the HasMass_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetHasMass_Ok() (*bool, bool) {
	if o == nil || o.HasMass_ == nil {
		return nil, false
	}
	return o.HasMass_, true
}

// HasHasMass_ returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasHasMass_() bool {
	if o != nil && o.HasMass_ != nil {
		return true
	}

	return false
}

// SetHasMass_ gets a reference to the given bool and assigns it to the HasMass_ field.
func (o *BTMassPropertiesInfo) SetHasMass_(v bool) {
	o.HasMass_ = &v
}

// GetInertia returns the Inertia field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetInertia() []float64 {
	if o == nil || o.Inertia == nil {
		var ret []float64
		return ret
	}
	return o.Inertia
}

// GetInertiaOk returns a tuple with the Inertia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetInertiaOk() ([]float64, bool) {
	if o == nil || o.Inertia == nil {
		return nil, false
	}
	return o.Inertia, true
}

// HasInertia returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasInertia() bool {
	if o != nil && o.Inertia != nil {
		return true
	}

	return false
}

// SetInertia gets a reference to the given []float64 and assigns it to the Inertia field.
func (o *BTMassPropertiesInfo) SetInertia(v []float64) {
	o.Inertia = v
}

// GetMass returns the Mass field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetMass() []float64 {
	if o == nil || o.Mass == nil {
		var ret []float64
		return ret
	}
	return o.Mass
}

// GetMassOk returns a tuple with the Mass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetMassOk() ([]float64, bool) {
	if o == nil || o.Mass == nil {
		return nil, false
	}
	return o.Mass, true
}

// HasMass returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasMass() bool {
	if o != nil && o.Mass != nil {
		return true
	}

	return false
}

// SetMass gets a reference to the given []float64 and assigns it to the Mass field.
func (o *BTMassPropertiesInfo) SetMass(v []float64) {
	o.Mass = v
}

// GetMassMissingCount returns the MassMissingCount field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetMassMissingCount() int32 {
	if o == nil || o.MassMissingCount == nil {
		var ret int32
		return ret
	}
	return *o.MassMissingCount
}

// GetMassMissingCountOk returns a tuple with the MassMissingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetMassMissingCountOk() (*int32, bool) {
	if o == nil || o.MassMissingCount == nil {
		return nil, false
	}
	return o.MassMissingCount, true
}

// HasMassMissingCount returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasMassMissingCount() bool {
	if o != nil && o.MassMissingCount != nil {
		return true
	}

	return false
}

// SetMassMissingCount gets a reference to the given int32 and assigns it to the MassMissingCount field.
func (o *BTMassPropertiesInfo) SetMassMissingCount(v int32) {
	o.MassMissingCount = &v
}

// GetPeriphery returns the Periphery field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetPeriphery() []float64 {
	if o == nil || o.Periphery == nil {
		var ret []float64
		return ret
	}
	return o.Periphery
}

// GetPeripheryOk returns a tuple with the Periphery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetPeripheryOk() ([]float64, bool) {
	if o == nil || o.Periphery == nil {
		return nil, false
	}
	return o.Periphery, true
}

// HasPeriphery returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasPeriphery() bool {
	if o != nil && o.Periphery != nil {
		return true
	}

	return false
}

// SetPeriphery gets a reference to the given []float64 and assigns it to the Periphery field.
func (o *BTMassPropertiesInfo) SetPeriphery(v []float64) {
	o.Periphery = v
}

// GetPrincipalAxes returns the PrincipalAxes field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetPrincipalAxes() []BTVector3d389 {
	if o == nil || o.PrincipalAxes == nil {
		var ret []BTVector3d389
		return ret
	}
	return o.PrincipalAxes
}

// GetPrincipalAxesOk returns a tuple with the PrincipalAxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetPrincipalAxesOk() ([]BTVector3d389, bool) {
	if o == nil || o.PrincipalAxes == nil {
		return nil, false
	}
	return o.PrincipalAxes, true
}

// HasPrincipalAxes returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasPrincipalAxes() bool {
	if o != nil && o.PrincipalAxes != nil {
		return true
	}

	return false
}

// SetPrincipalAxes gets a reference to the given []BTVector3d389 and assigns it to the PrincipalAxes field.
func (o *BTMassPropertiesInfo) SetPrincipalAxes(v []BTVector3d389) {
	o.PrincipalAxes = v
}

// GetPrincipalInertia returns the PrincipalInertia field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetPrincipalInertia() []float64 {
	if o == nil || o.PrincipalInertia == nil {
		var ret []float64
		return ret
	}
	return o.PrincipalInertia
}

// GetPrincipalInertiaOk returns a tuple with the PrincipalInertia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetPrincipalInertiaOk() ([]float64, bool) {
	if o == nil || o.PrincipalInertia == nil {
		return nil, false
	}
	return o.PrincipalInertia, true
}

// HasPrincipalInertia returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasPrincipalInertia() bool {
	if o != nil && o.PrincipalInertia != nil {
		return true
	}

	return false
}

// SetPrincipalInertia gets a reference to the given []float64 and assigns it to the PrincipalInertia field.
func (o *BTMassPropertiesInfo) SetPrincipalInertia(v []float64) {
	o.PrincipalInertia = v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *BTMassPropertiesInfo) GetVolume() []float64 {
	if o == nil || o.Volume == nil {
		var ret []float64
		return ret
	}
	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesInfo) GetVolumeOk() ([]float64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *BTMassPropertiesInfo) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given []float64 and assigns it to the Volume field.
func (o *BTMassPropertiesInfo) SetVolume(v []float64) {
	o.Volume = v
}

func (o BTMassPropertiesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Centroid != nil {
		toSerialize["centroid"] = o.Centroid
	}
	if o.HasMass_ != nil {
		toSerialize["hasMass"] = o.HasMass_
	}
	if o.Inertia != nil {
		toSerialize["inertia"] = o.Inertia
	}
	if o.Mass != nil {
		toSerialize["mass"] = o.Mass
	}
	if o.MassMissingCount != nil {
		toSerialize["massMissingCount"] = o.MassMissingCount
	}
	if o.Periphery != nil {
		toSerialize["periphery"] = o.Periphery
	}
	if o.PrincipalAxes != nil {
		toSerialize["principalAxes"] = o.PrincipalAxes
	}
	if o.PrincipalInertia != nil {
		toSerialize["principalInertia"] = o.PrincipalInertia
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableBTMassPropertiesInfo struct {
	value *BTMassPropertiesInfo
	isSet bool
}

func (v NullableBTMassPropertiesInfo) Get() *BTMassPropertiesInfo {
	return v.value
}

func (v *NullableBTMassPropertiesInfo) Set(val *BTMassPropertiesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMassPropertiesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMassPropertiesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMassPropertiesInfo(val *BTMassPropertiesInfo) *NullableBTMassPropertiesInfo {
	return &NullableBTMassPropertiesInfo{value: val, isSet: true}
}

func (v NullableBTMassPropertiesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMassPropertiesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
