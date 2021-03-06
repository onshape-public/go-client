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

// BTMSketchTextEntity1761AllOf struct for BTMSketchTextEntity1761AllOf
type BTMSketchTextEntity1761AllOf struct {
	Ascent *float64 `json:"ascent,omitempty"`
	BaselineDirectionX *float64 `json:"baselineDirectionX,omitempty"`
	BaselineDirectionY *float64 `json:"baselineDirectionY,omitempty"`
	BaselineStartX *float64 `json:"baselineStartX,omitempty"`
	BaselineStartY *float64 `json:"baselineStartY,omitempty"`
	BtType *string `json:"btType,omitempty"`
	FontName *string `json:"fontName,omitempty"`
	Text *string `json:"text,omitempty"`
}

// NewBTMSketchTextEntity1761AllOf instantiates a new BTMSketchTextEntity1761AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchTextEntity1761AllOf() *BTMSketchTextEntity1761AllOf {
	this := BTMSketchTextEntity1761AllOf{}
	return &this
}

// NewBTMSketchTextEntity1761AllOfWithDefaults instantiates a new BTMSketchTextEntity1761AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchTextEntity1761AllOfWithDefaults() *BTMSketchTextEntity1761AllOf {
	this := BTMSketchTextEntity1761AllOf{}
	return &this
}

// GetAscent returns the Ascent field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetAscent() float64 {
	if o == nil || o.Ascent == nil {
		var ret float64
		return ret
	}
	return *o.Ascent
}

// GetAscentOk returns a tuple with the Ascent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetAscentOk() (*float64, bool) {
	if o == nil || o.Ascent == nil {
		return nil, false
	}
	return o.Ascent, true
}

// HasAscent returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasAscent() bool {
	if o != nil && o.Ascent != nil {
		return true
	}

	return false
}

// SetAscent gets a reference to the given float64 and assigns it to the Ascent field.
func (o *BTMSketchTextEntity1761AllOf) SetAscent(v float64) {
	o.Ascent = &v
}

// GetBaselineDirectionX returns the BaselineDirectionX field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineDirectionX() float64 {
	if o == nil || o.BaselineDirectionX == nil {
		var ret float64
		return ret
	}
	return *o.BaselineDirectionX
}

// GetBaselineDirectionXOk returns a tuple with the BaselineDirectionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineDirectionXOk() (*float64, bool) {
	if o == nil || o.BaselineDirectionX == nil {
		return nil, false
	}
	return o.BaselineDirectionX, true
}

// HasBaselineDirectionX returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasBaselineDirectionX() bool {
	if o != nil && o.BaselineDirectionX != nil {
		return true
	}

	return false
}

// SetBaselineDirectionX gets a reference to the given float64 and assigns it to the BaselineDirectionX field.
func (o *BTMSketchTextEntity1761AllOf) SetBaselineDirectionX(v float64) {
	o.BaselineDirectionX = &v
}

// GetBaselineDirectionY returns the BaselineDirectionY field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineDirectionY() float64 {
	if o == nil || o.BaselineDirectionY == nil {
		var ret float64
		return ret
	}
	return *o.BaselineDirectionY
}

// GetBaselineDirectionYOk returns a tuple with the BaselineDirectionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineDirectionYOk() (*float64, bool) {
	if o == nil || o.BaselineDirectionY == nil {
		return nil, false
	}
	return o.BaselineDirectionY, true
}

// HasBaselineDirectionY returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasBaselineDirectionY() bool {
	if o != nil && o.BaselineDirectionY != nil {
		return true
	}

	return false
}

// SetBaselineDirectionY gets a reference to the given float64 and assigns it to the BaselineDirectionY field.
func (o *BTMSketchTextEntity1761AllOf) SetBaselineDirectionY(v float64) {
	o.BaselineDirectionY = &v
}

// GetBaselineStartX returns the BaselineStartX field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineStartX() float64 {
	if o == nil || o.BaselineStartX == nil {
		var ret float64
		return ret
	}
	return *o.BaselineStartX
}

// GetBaselineStartXOk returns a tuple with the BaselineStartX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineStartXOk() (*float64, bool) {
	if o == nil || o.BaselineStartX == nil {
		return nil, false
	}
	return o.BaselineStartX, true
}

// HasBaselineStartX returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasBaselineStartX() bool {
	if o != nil && o.BaselineStartX != nil {
		return true
	}

	return false
}

// SetBaselineStartX gets a reference to the given float64 and assigns it to the BaselineStartX field.
func (o *BTMSketchTextEntity1761AllOf) SetBaselineStartX(v float64) {
	o.BaselineStartX = &v
}

// GetBaselineStartY returns the BaselineStartY field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineStartY() float64 {
	if o == nil || o.BaselineStartY == nil {
		var ret float64
		return ret
	}
	return *o.BaselineStartY
}

// GetBaselineStartYOk returns a tuple with the BaselineStartY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetBaselineStartYOk() (*float64, bool) {
	if o == nil || o.BaselineStartY == nil {
		return nil, false
	}
	return o.BaselineStartY, true
}

// HasBaselineStartY returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasBaselineStartY() bool {
	if o != nil && o.BaselineStartY != nil {
		return true
	}

	return false
}

// SetBaselineStartY gets a reference to the given float64 and assigns it to the BaselineStartY field.
func (o *BTMSketchTextEntity1761AllOf) SetBaselineStartY(v float64) {
	o.BaselineStartY = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchTextEntity1761AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetFontName returns the FontName field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetFontName() string {
	if o == nil || o.FontName == nil {
		var ret string
		return ret
	}
	return *o.FontName
}

// GetFontNameOk returns a tuple with the FontName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetFontNameOk() (*string, bool) {
	if o == nil || o.FontName == nil {
		return nil, false
	}
	return o.FontName, true
}

// HasFontName returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasFontName() bool {
	if o != nil && o.FontName != nil {
		return true
	}

	return false
}

// SetFontName gets a reference to the given string and assigns it to the FontName field.
func (o *BTMSketchTextEntity1761AllOf) SetFontName(v string) {
	o.FontName = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *BTMSketchTextEntity1761AllOf) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchTextEntity1761AllOf) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *BTMSketchTextEntity1761AllOf) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *BTMSketchTextEntity1761AllOf) SetText(v string) {
	o.Text = &v
}

func (o BTMSketchTextEntity1761AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ascent != nil {
		toSerialize["ascent"] = o.Ascent
	}
	if o.BaselineDirectionX != nil {
		toSerialize["baselineDirectionX"] = o.BaselineDirectionX
	}
	if o.BaselineDirectionY != nil {
		toSerialize["baselineDirectionY"] = o.BaselineDirectionY
	}
	if o.BaselineStartX != nil {
		toSerialize["baselineStartX"] = o.BaselineStartX
	}
	if o.BaselineStartY != nil {
		toSerialize["baselineStartY"] = o.BaselineStartY
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FontName != nil {
		toSerialize["fontName"] = o.FontName
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchTextEntity1761AllOf struct {
	value *BTMSketchTextEntity1761AllOf
	isSet bool
}

func (v NullableBTMSketchTextEntity1761AllOf) Get() *BTMSketchTextEntity1761AllOf {
	return v.value
}

func (v *NullableBTMSketchTextEntity1761AllOf) Set(val *BTMSketchTextEntity1761AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchTextEntity1761AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchTextEntity1761AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchTextEntity1761AllOf(val *BTMSketchTextEntity1761AllOf) *NullableBTMSketchTextEntity1761AllOf {
	return &NullableBTMSketchTextEntity1761AllOf{value: val, isSet: true}
}

func (v NullableBTMSketchTextEntity1761AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchTextEntity1761AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
