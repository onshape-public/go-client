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

// BTMParameterQuantity147 - struct for BTMParameterQuantity147
type BTMParameterQuantity147 struct {
	implBTMParameterQuantity147 interface{}
}

// BTMParameterNullableQuantity807AsBTMParameterQuantity147 is a convenience function that returns BTMParameterNullableQuantity807 wrapped in BTMParameterQuantity147
func (o *BTMParameterNullableQuantity807) AsBTMParameterQuantity147() *BTMParameterQuantity147 {
	return &BTMParameterQuantity147{o}
}

// BTMParameterTolerantQuantity2579AsBTMParameterQuantity147 is a convenience function that returns BTMParameterTolerantQuantity2579 wrapped in BTMParameterQuantity147
func (o *BTMParameterTolerantQuantity2579) AsBTMParameterQuantity147() *BTMParameterQuantity147 {
	return &BTMParameterQuantity147{o}
}

// NewBTMParameterQuantity147 instantiates a new BTMParameterQuantity147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterQuantity147() *BTMParameterQuantity147 {
	this := BTMParameterQuantity147{Newbase_BTMParameterQuantity147()}
	return &this
}

// NewBTMParameterQuantity147WithDefaults instantiates a new BTMParameterQuantity147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterQuantity147WithDefaults() *BTMParameterQuantity147 {
	this := BTMParameterQuantity147{Newbase_BTMParameterQuantity147WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetBtType() string {
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
func (o *BTMParameterQuantity147) GetBtTypeOk() (*string, bool) {
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
func (o *BTMParameterQuantity147) HasBtType() bool {
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
func (o *BTMParameterQuantity147) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterQuantity147) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterQuantity147) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetParameterId() string {
	type getResult interface {
		GetParameterId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterId()
	} else {
		var de string
		return de
	}
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetParameterIdOk() (*string, bool) {
	type getResult interface {
		GetParameterIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterIdOk()
	} else {
		return nil, false
	}
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasParameterId() bool {
	type getResult interface {
		HasParameterId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameterId()
	} else {
		return false
	}
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterQuantity147) SetParameterId(v string) {
	type getResult interface {
		SetParameterId(v string)
	}

	o.GetActualInstance().(getResult).SetParameterId(v)
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetValueString() string {
	type getResult interface {
		GetValueString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValueString()
	} else {
		var de string
		return de
	}
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetValueStringOk() (*string, bool) {
	type getResult interface {
		GetValueStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValueStringOk()
	} else {
		return nil, false
	}
}

// HasValueString returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasValueString() bool {
	type getResult interface {
		HasValueString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasValueString()
	} else {
		return false
	}
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *BTMParameterQuantity147) SetValueString(v string) {
	type getResult interface {
		SetValueString(v string)
	}

	o.GetActualInstance().(getResult).SetValueString(v)
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetExpression() string {
	type getResult interface {
		GetExpression() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExpression()
	} else {
		var de string
		return de
	}
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetExpressionOk() (*string, bool) {
	type getResult interface {
		GetExpressionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExpressionOk()
	} else {
		return nil, false
	}
}

// HasExpression returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasExpression() bool {
	type getResult interface {
		HasExpression() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasExpression()
	} else {
		return false
	}
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *BTMParameterQuantity147) SetExpression(v string) {
	type getResult interface {
		SetExpression(v string)
	}

	o.GetActualInstance().(getResult).SetExpression(v)
}

// GetIsInteger returns the IsInteger field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetIsInteger() bool {
	type getResult interface {
		GetIsInteger() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsInteger()
	} else {
		var de bool
		return de
	}
}

// GetIsIntegerOk returns a tuple with the IsInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetIsIntegerOk() (*bool, bool) {
	type getResult interface {
		GetIsIntegerOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsIntegerOk()
	} else {
		return nil, false
	}
}

// HasIsInteger returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasIsInteger() bool {
	type getResult interface {
		HasIsInteger() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsInteger()
	} else {
		return false
	}
}

// SetIsInteger gets a reference to the given bool and assigns it to the IsInteger field.
func (o *BTMParameterQuantity147) SetIsInteger(v bool) {
	type getResult interface {
		SetIsInteger(v bool)
	}

	o.GetActualInstance().(getResult).SetIsInteger(v)
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetUnits() string {
	type getResult interface {
		GetUnits() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUnits()
	} else {
		var de string
		return de
	}
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetUnitsOk() (*string, bool) {
	type getResult interface {
		GetUnitsOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUnitsOk()
	} else {
		return nil, false
	}
}

// HasUnits returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasUnits() bool {
	type getResult interface {
		HasUnits() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasUnits()
	} else {
		return false
	}
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTMParameterQuantity147) SetUnits(v string) {
	type getResult interface {
		SetUnits(v string)
	}

	o.GetActualInstance().(getResult).SetUnits(v)
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMParameterQuantity147) GetValue() float64 {
	type getResult interface {
		GetValue() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValue()
	} else {
		var de float64
		return de
	}
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQuantity147) GetValueOk() (*float64, bool) {
	type getResult interface {
		GetValueOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValueOk()
	} else {
		return nil, false
	}
}

// HasValue returns a boolean if a field has been set.
func (o *BTMParameterQuantity147) HasValue() bool {
	type getResult interface {
		HasValue() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasValue()
	} else {
		return false
	}
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTMParameterQuantity147) SetValue(v float64) {
	type getResult interface {
		SetValue(v float64)
	}

	o.GetActualInstance().(getResult).SetValue(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMParameterQuantity147) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTMParameterNullableQuantity-807'
	if jsonDict["btType"] == "BTMParameterNullableQuantity-807" {
		// try to unmarshal JSON data into BTMParameterNullableQuantity807
		var qr *BTMParameterNullableQuantity807
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterQuantity147 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterQuantity147 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterQuantity147 as BTMParameterNullableQuantity807: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterTolerantQuantity-2579'
	if jsonDict["btType"] == "BTMParameterTolerantQuantity-2579" {
		// try to unmarshal JSON data into BTMParameterTolerantQuantity2579
		var qr *BTMParameterTolerantQuantity2579
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameterQuantity147 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameterQuantity147 = nil
			return fmt.Errorf("failed to unmarshal BTMParameterQuantity147 as BTMParameterTolerantQuantity2579: %s", err.Error())
		}
	}

	var qtx *base_BTMParameterQuantity147
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMParameterQuantity147 = qtx
		return nil // data stored in dst.base_BTMParameterQuantity147, return on the first match
	} else {
		dst.implBTMParameterQuantity147 = nil
		return fmt.Errorf("failed to unmarshal BTMParameterQuantity147 as base_BTMParameterQuantity147: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMParameterQuantity147) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMParameterQuantity147) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMParameterQuantity147
}

type NullableBTMParameterQuantity147 struct {
	value *BTMParameterQuantity147
	isSet bool
}

func (v NullableBTMParameterQuantity147) Get() *BTMParameterQuantity147 {
	return v.value
}

func (v *NullableBTMParameterQuantity147) Set(val *BTMParameterQuantity147) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterQuantity147) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterQuantity147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterQuantity147(val *BTMParameterQuantity147) *NullableBTMParameterQuantity147 {
	return &NullableBTMParameterQuantity147{value: val, isSet: true}
}

func (v NullableBTMParameterQuantity147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterQuantity147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMParameterQuantity147 struct {
	BTMParameter1
	BtType *string `json:"btType,omitempty"`
	// Microversion that resulted from the import.
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	// ID of the parameter's node.
	NodeId *string `json:"nodeId,omitempty"`
	// Unique ID of the parameter.
	ParameterId *string  `json:"parameterId,omitempty"`
	ValueString *string  `json:"valueString,omitempty"`
	Expression  *string  `json:"expression,omitempty"`
	IsInteger   *bool    `json:"isInteger,omitempty"`
	Units       *string  `json:"units,omitempty"`
	Value       *float64 `json:"value,omitempty"`
}

// Newbase_BTMParameterQuantity147 instantiates a new base_BTMParameterQuantity147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMParameterQuantity147() *base_BTMParameterQuantity147 {
	this := base_BTMParameterQuantity147{}
	return &this
}

// Newbase_BTMParameterQuantity147WithDefaults instantiates a new base_BTMParameterQuantity147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMParameterQuantity147WithDefaults() *base_BTMParameterQuantity147 {
	this := base_BTMParameterQuantity147{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMParameterQuantity147) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMParameterQuantity147) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMParameterQuantity147) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *base_BTMParameterQuantity147) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetValueString() string {
	if o == nil || o.ValueString == nil {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetValueStringOk() (*string, bool) {
	if o == nil || o.ValueString == nil {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasValueString() bool {
	if o != nil && o.ValueString != nil {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *base_BTMParameterQuantity147) SetValueString(v string) {
	o.ValueString = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *base_BTMParameterQuantity147) SetExpression(v string) {
	o.Expression = &v
}

// GetIsInteger returns the IsInteger field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetIsInteger() bool {
	if o == nil || o.IsInteger == nil {
		var ret bool
		return ret
	}
	return *o.IsInteger
}

// GetIsIntegerOk returns a tuple with the IsInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetIsIntegerOk() (*bool, bool) {
	if o == nil || o.IsInteger == nil {
		return nil, false
	}
	return o.IsInteger, true
}

// HasIsInteger returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasIsInteger() bool {
	if o != nil && o.IsInteger != nil {
		return true
	}

	return false
}

// SetIsInteger gets a reference to the given bool and assigns it to the IsInteger field.
func (o *base_BTMParameterQuantity147) SetIsInteger(v bool) {
	o.IsInteger = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *base_BTMParameterQuantity147) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *base_BTMParameterQuantity147) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameterQuantity147) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *base_BTMParameterQuantity147) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *base_BTMParameterQuantity147) SetValue(v float64) {
	o.Value = &v
}

func (o base_BTMParameterQuantity147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMParameter1, errBTMParameter1 := json.Marshal(o.BTMParameter1)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	errBTMParameter1 = json.Unmarshal([]byte(serializedBTMParameter1), &toSerialize)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ValueString != nil {
		toSerialize["valueString"] = o.ValueString
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.IsInteger != nil {
		toSerialize["isInteger"] = o.IsInteger
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}
