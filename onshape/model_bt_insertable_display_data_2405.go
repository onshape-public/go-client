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

// BTInsertableDisplayData2405 - struct for BTInsertableDisplayData2405
type BTInsertableDisplayData2405 struct {
	implBTInsertableDisplayData2405 interface{}
}

// BTInsertableSketchDisplayData3775AsBTInsertableDisplayData2405 is a convenience function that returns BTInsertableSketchDisplayData3775 wrapped in BTInsertableDisplayData2405
func (o *BTInsertableSketchDisplayData3775) AsBTInsertableDisplayData2405() *BTInsertableDisplayData2405 {
	return &BTInsertableDisplayData2405{o}
}

// BTInsertablePartDisplayData3103AsBTInsertableDisplayData2405 is a convenience function that returns BTInsertablePartDisplayData3103 wrapped in BTInsertableDisplayData2405
func (o *BTInsertablePartDisplayData3103) AsBTInsertableDisplayData2405() *BTInsertableDisplayData2405 {
	return &BTInsertableDisplayData2405{o}
}

// NewBTInsertableDisplayData2405 instantiates a new BTInsertableDisplayData2405 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInsertableDisplayData2405() *BTInsertableDisplayData2405 {
	this := BTInsertableDisplayData2405{Newbase_BTInsertableDisplayData2405()}
	return &this
}

// NewBTInsertableDisplayData2405WithDefaults instantiates a new BTInsertableDisplayData2405 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInsertableDisplayData2405WithDefaults() *BTInsertableDisplayData2405 {
	this := BTInsertableDisplayData2405{Newbase_BTInsertableDisplayData2405WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetBtType() string {
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
func (o *BTInsertableDisplayData2405) GetBtTypeOk() (*string, bool) {
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
func (o *BTInsertableDisplayData2405) HasBtType() bool {
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
func (o *BTInsertableDisplayData2405) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetFullElementId() BTFullElementId756 {
	type getResult interface {
		GetFullElementId() BTFullElementId756
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullElementId()
	} else {
		var de BTFullElementId756
		return de
	}
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableDisplayData2405) GetFullElementIdOk() (*BTFullElementId756, bool) {
	type getResult interface {
		GetFullElementIdOk() (*BTFullElementId756, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullElementIdOk()
	} else {
		return nil, false
	}
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTInsertableDisplayData2405) HasFullElementId() bool {
	type getResult interface {
		HasFullElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFullElementId()
	} else {
		return false
	}
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTInsertableDisplayData2405) SetFullElementId(v BTFullElementId756) {
	type getResult interface {
		SetFullElementId(v BTFullElementId756)
	}

	o.GetActualInstance().(getResult).SetFullElementId(v)
}

// GetGraphicsBuffers returns the GraphicsBuffers field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetGraphicsBuffers() map[string]map[string]BTGraphicsBuffer2668 {
	type getResult interface {
		GetGraphicsBuffers() map[string]map[string]BTGraphicsBuffer2668
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetGraphicsBuffers()
	} else {
		var de map[string]map[string]BTGraphicsBuffer2668
		return de
	}
}

// GetGraphicsBuffersOk returns a tuple with the GraphicsBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableDisplayData2405) GetGraphicsBuffersOk() (*map[string]map[string]BTGraphicsBuffer2668, bool) {
	type getResult interface {
		GetGraphicsBuffersOk() (*map[string]map[string]BTGraphicsBuffer2668, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetGraphicsBuffersOk()
	} else {
		return nil, false
	}
}

// HasGraphicsBuffers returns a boolean if a field has been set.
func (o *BTInsertableDisplayData2405) HasGraphicsBuffers() bool {
	type getResult interface {
		HasGraphicsBuffers() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasGraphicsBuffers()
	} else {
		return false
	}
}

// SetGraphicsBuffers gets a reference to the given map[string]map[string]BTGraphicsBuffer2668 and assigns it to the GraphicsBuffers field.
func (o *BTInsertableDisplayData2405) SetGraphicsBuffers(v map[string]map[string]BTGraphicsBuffer2668) {
	type getResult interface {
		SetGraphicsBuffers(v map[string]map[string]BTGraphicsBuffer2668)
	}

	o.GetActualInstance().(getResult).SetGraphicsBuffers(v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableDisplayData2405) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *BTInsertableDisplayData2405) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTInsertableDisplayData2405) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetPart() bool {
	type getResult interface {
		GetPart() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPart()
	} else {
		var de bool
		return de
	}
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableDisplayData2405) GetPartOk() (*bool, bool) {
	type getResult interface {
		GetPartOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartOk()
	} else {
		return nil, false
	}
}

// HasPart returns a boolean if a field has been set.
func (o *BTInsertableDisplayData2405) HasPart() bool {
	type getResult interface {
		HasPart() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPart()
	} else {
		return false
	}
}

// SetPart gets a reference to the given bool and assigns it to the Part field.
func (o *BTInsertableDisplayData2405) SetPart(v bool) {
	type getResult interface {
		SetPart(v bool)
	}

	o.GetActualInstance().(getResult).SetPart(v)
}

// GetSketchFeature returns the SketchFeature field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetSketchFeature() bool {
	type getResult interface {
		GetSketchFeature() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSketchFeature()
	} else {
		var de bool
		return de
	}
}

// GetSketchFeatureOk returns a tuple with the SketchFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableDisplayData2405) GetSketchFeatureOk() (*bool, bool) {
	type getResult interface {
		GetSketchFeatureOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSketchFeatureOk()
	} else {
		return nil, false
	}
}

// HasSketchFeature returns a boolean if a field has been set.
func (o *BTInsertableDisplayData2405) HasSketchFeature() bool {
	type getResult interface {
		HasSketchFeature() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSketchFeature()
	} else {
		return false
	}
}

// SetSketchFeature gets a reference to the given bool and assigns it to the SketchFeature field.
func (o *BTInsertableDisplayData2405) SetSketchFeature(v bool) {
	type getResult interface {
		SetSketchFeature(v bool)
	}

	o.GetActualInstance().(getResult).SetSketchFeature(v)
}

// GetTessellationSettingIndex returns the TessellationSettingIndex field value if set, zero value otherwise.
func (o *BTInsertableDisplayData2405) GetTessellationSettingIndex() int32 {
	type getResult interface {
		GetTessellationSettingIndex() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTessellationSettingIndex()
	} else {
		var de int32
		return de
	}
}

// GetTessellationSettingIndexOk returns a tuple with the TessellationSettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertableDisplayData2405) GetTessellationSettingIndexOk() (*int32, bool) {
	type getResult interface {
		GetTessellationSettingIndexOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTessellationSettingIndexOk()
	} else {
		return nil, false
	}
}

// HasTessellationSettingIndex returns a boolean if a field has been set.
func (o *BTInsertableDisplayData2405) HasTessellationSettingIndex() bool {
	type getResult interface {
		HasTessellationSettingIndex() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTessellationSettingIndex()
	} else {
		return false
	}
}

// SetTessellationSettingIndex gets a reference to the given int32 and assigns it to the TessellationSettingIndex field.
func (o *BTInsertableDisplayData2405) SetTessellationSettingIndex(v int32) {
	type getResult interface {
		SetTessellationSettingIndex(v int32)
	}

	o.GetActualInstance().(getResult).SetTessellationSettingIndex(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTInsertableDisplayData2405) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTInsertablePartDisplayData-3103'
	if jsonDict["btType"] == "BTInsertablePartDisplayData-3103" {
		// try to unmarshal JSON data into BTInsertablePartDisplayData3103
		var qr *BTInsertablePartDisplayData3103
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTInsertableDisplayData2405 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTInsertableDisplayData2405 = nil
			return fmt.Errorf("failed to unmarshal BTInsertableDisplayData2405 as BTInsertablePartDisplayData3103: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTInsertableSketchDisplayData-3775'
	if jsonDict["btType"] == "BTInsertableSketchDisplayData-3775" {
		// try to unmarshal JSON data into BTInsertableSketchDisplayData3775
		var qr *BTInsertableSketchDisplayData3775
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTInsertableDisplayData2405 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTInsertableDisplayData2405 = nil
			return fmt.Errorf("failed to unmarshal BTInsertableDisplayData2405 as BTInsertableSketchDisplayData3775: %s", err.Error())
		}
	}

	var qtx *base_BTInsertableDisplayData2405
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTInsertableDisplayData2405 = qtx
		return nil // data stored in dst.base_BTInsertableDisplayData2405, return on the first match
	} else {
		dst.implBTInsertableDisplayData2405 = nil
		return fmt.Errorf("failed to unmarshal BTInsertableDisplayData2405 as base_BTInsertableDisplayData2405: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTInsertableDisplayData2405) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTInsertableDisplayData2405) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTInsertableDisplayData2405
}

type NullableBTInsertableDisplayData2405 struct {
	value *BTInsertableDisplayData2405
	isSet bool
}

func (v NullableBTInsertableDisplayData2405) Get() *BTInsertableDisplayData2405 {
	return v.value
}

func (v *NullableBTInsertableDisplayData2405) Set(val *BTInsertableDisplayData2405) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInsertableDisplayData2405) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInsertableDisplayData2405) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInsertableDisplayData2405(val *BTInsertableDisplayData2405) *NullableBTInsertableDisplayData2405 {
	return &NullableBTInsertableDisplayData2405{value: val, isSet: true}
}

func (v NullableBTInsertableDisplayData2405) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInsertableDisplayData2405) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTInsertableDisplayData2405 struct {
	// Type of JSON object.
	BtType                   *string                                     `json:"btType,omitempty"`
	FullElementId            *BTFullElementId756                         `json:"fullElementId,omitempty"`
	GraphicsBuffers          *map[string]map[string]BTGraphicsBuffer2668 `json:"graphicsBuffers,omitempty"`
	Id                       *string                                     `json:"id,omitempty"`
	Part                     *bool                                       `json:"part,omitempty"`
	SketchFeature            *bool                                       `json:"sketchFeature,omitempty"`
	TessellationSettingIndex *int32                                      `json:"tessellationSettingIndex,omitempty"`
}

// Newbase_BTInsertableDisplayData2405 instantiates a new base_BTInsertableDisplayData2405 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTInsertableDisplayData2405() *base_BTInsertableDisplayData2405 {
	this := base_BTInsertableDisplayData2405{}
	return &this
}

// Newbase_BTInsertableDisplayData2405WithDefaults instantiates a new base_BTInsertableDisplayData2405 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTInsertableDisplayData2405WithDefaults() *base_BTInsertableDisplayData2405 {
	this := base_BTInsertableDisplayData2405{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTInsertableDisplayData2405) SetBtType(v string) {
	o.BtType = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *base_BTInsertableDisplayData2405) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetGraphicsBuffers returns the GraphicsBuffers field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetGraphicsBuffers() map[string]map[string]BTGraphicsBuffer2668 {
	if o == nil || o.GraphicsBuffers == nil {
		var ret map[string]map[string]BTGraphicsBuffer2668
		return ret
	}
	return *o.GraphicsBuffers
}

// GetGraphicsBuffersOk returns a tuple with the GraphicsBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetGraphicsBuffersOk() (*map[string]map[string]BTGraphicsBuffer2668, bool) {
	if o == nil || o.GraphicsBuffers == nil {
		return nil, false
	}
	return o.GraphicsBuffers, true
}

// HasGraphicsBuffers returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasGraphicsBuffers() bool {
	if o != nil && o.GraphicsBuffers != nil {
		return true
	}

	return false
}

// SetGraphicsBuffers gets a reference to the given map[string]map[string]BTGraphicsBuffer2668 and assigns it to the GraphicsBuffers field.
func (o *base_BTInsertableDisplayData2405) SetGraphicsBuffers(v map[string]map[string]BTGraphicsBuffer2668) {
	o.GraphicsBuffers = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_BTInsertableDisplayData2405) SetId(v string) {
	o.Id = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetPart() bool {
	if o == nil || o.Part == nil {
		var ret bool
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetPartOk() (*bool, bool) {
	if o == nil || o.Part == nil {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasPart() bool {
	if o != nil && o.Part != nil {
		return true
	}

	return false
}

// SetPart gets a reference to the given bool and assigns it to the Part field.
func (o *base_BTInsertableDisplayData2405) SetPart(v bool) {
	o.Part = &v
}

// GetSketchFeature returns the SketchFeature field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetSketchFeature() bool {
	if o == nil || o.SketchFeature == nil {
		var ret bool
		return ret
	}
	return *o.SketchFeature
}

// GetSketchFeatureOk returns a tuple with the SketchFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetSketchFeatureOk() (*bool, bool) {
	if o == nil || o.SketchFeature == nil {
		return nil, false
	}
	return o.SketchFeature, true
}

// HasSketchFeature returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasSketchFeature() bool {
	if o != nil && o.SketchFeature != nil {
		return true
	}

	return false
}

// SetSketchFeature gets a reference to the given bool and assigns it to the SketchFeature field.
func (o *base_BTInsertableDisplayData2405) SetSketchFeature(v bool) {
	o.SketchFeature = &v
}

// GetTessellationSettingIndex returns the TessellationSettingIndex field value if set, zero value otherwise.
func (o *base_BTInsertableDisplayData2405) GetTessellationSettingIndex() int32 {
	if o == nil || o.TessellationSettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.TessellationSettingIndex
}

// GetTessellationSettingIndexOk returns a tuple with the TessellationSettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInsertableDisplayData2405) GetTessellationSettingIndexOk() (*int32, bool) {
	if o == nil || o.TessellationSettingIndex == nil {
		return nil, false
	}
	return o.TessellationSettingIndex, true
}

// HasTessellationSettingIndex returns a boolean if a field has been set.
func (o *base_BTInsertableDisplayData2405) HasTessellationSettingIndex() bool {
	if o != nil && o.TessellationSettingIndex != nil {
		return true
	}

	return false
}

// SetTessellationSettingIndex gets a reference to the given int32 and assigns it to the TessellationSettingIndex field.
func (o *base_BTInsertableDisplayData2405) SetTessellationSettingIndex(v int32) {
	o.TessellationSettingIndex = &v
}

func (o base_BTInsertableDisplayData2405) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.GraphicsBuffers != nil {
		toSerialize["graphicsBuffers"] = o.GraphicsBuffers
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Part != nil {
		toSerialize["part"] = o.Part
	}
	if o.SketchFeature != nil {
		toSerialize["sketchFeature"] = o.SketchFeature
	}
	if o.TessellationSettingIndex != nil {
		toSerialize["tessellationSettingIndex"] = o.TessellationSettingIndex
	}
	return json.Marshal(toSerialize)
}
