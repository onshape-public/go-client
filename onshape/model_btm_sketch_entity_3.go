/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMSketchEntity3 - struct for BTMSketchEntity3
type BTMSketchEntity3 struct {
	implBTMSketchEntity3 interface{}
}

// BTMSketchGeomEntity5AsBTMSketchEntity3 is a convenience function that returns BTMSketchGeomEntity5 wrapped in BTMSketchEntity3
func (o *BTMSketchGeomEntity5) AsBTMSketchEntity3() *BTMSketchEntity3 {
	return &BTMSketchEntity3{o}
}

// BTMSketchConstraint2AsBTMSketchEntity3 is a convenience function that returns BTMSketchConstraint2 wrapped in BTMSketchEntity3
func (o *BTMSketchConstraint2) AsBTMSketchEntity3() *BTMSketchEntity3 {
	return &BTMSketchEntity3{o}
}

// BTMSketchInvalid1601AsBTMSketchEntity3 is a convenience function that returns BTMSketchInvalid1601 wrapped in BTMSketchEntity3
func (o *BTMSketchInvalid1601) AsBTMSketchEntity3() *BTMSketchEntity3 {
	return &BTMSketchEntity3{o}
}

// BTMSketchCompositeEntity893AsBTMSketchEntity3 is a convenience function that returns BTMSketchCompositeEntity893 wrapped in BTMSketchEntity3
func (o *BTMSketchCompositeEntity893) AsBTMSketchEntity3() *BTMSketchEntity3 {
	return &BTMSketchEntity3{o}
}

// NewBTMSketchEntity3 instantiates a new BTMSketchEntity3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchEntity3() *BTMSketchEntity3 {
	this := BTMSketchEntity3{Newbase_BTMSketchEntity3()}
	return &this
}

// NewBTMSketchEntity3WithDefaults instantiates a new BTMSketchEntity3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchEntity3WithDefaults() *BTMSketchEntity3 {
	this := BTMSketchEntity3{Newbase_BTMSketchEntity3WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetBtType() string {
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
func (o *BTMSketchEntity3) GetBtTypeOk() (*string, bool) {
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
func (o *BTMSketchEntity3) HasBtType() bool {
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
func (o *BTMSketchEntity3) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetImportMicroversion() string {
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
func (o *BTMSketchEntity3) GetImportMicroversionOk() (*string, bool) {
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
func (o *BTMSketchEntity3) HasImportMicroversion() bool {
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
func (o *BTMSketchEntity3) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetNodeId() string {
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
func (o *BTMSketchEntity3) GetNodeIdOk() (*string, bool) {
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
func (o *BTMSketchEntity3) HasNodeId() bool {
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
func (o *BTMSketchEntity3) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetEntityId() string {
	type getResult interface {
		GetEntityId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEntityId()
	} else {
		var de string
		return de
	}
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchEntity3) GetEntityIdOk() (*string, bool) {
	type getResult interface {
		GetEntityIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEntityIdOk()
	} else {
		return nil, false
	}
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchEntity3) HasEntityId() bool {
	type getResult interface {
		HasEntityId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEntityId()
	} else {
		return false
	}
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchEntity3) SetEntityId(v string) {
	type getResult interface {
		SetEntityId(v string)
	}

	o.GetActualInstance().(getResult).SetEntityId(v)
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetEntityIdAndReplaceInDependentFields() string {
	type getResult interface {
		GetEntityIdAndReplaceInDependentFields() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEntityIdAndReplaceInDependentFields()
	} else {
		var de string
		return de
	}
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchEntity3) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	type getResult interface {
		GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEntityIdAndReplaceInDependentFieldsOk()
	} else {
		return nil, false
	}
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchEntity3) HasEntityIdAndReplaceInDependentFields() bool {
	type getResult interface {
		HasEntityIdAndReplaceInDependentFields() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEntityIdAndReplaceInDependentFields()
	} else {
		return false
	}
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchEntity3) SetEntityIdAndReplaceInDependentFields(v string) {
	type getResult interface {
		SetEntityIdAndReplaceInDependentFields(v string)
	}

	o.GetActualInstance().(getResult).SetEntityIdAndReplaceInDependentFields(v)
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetNamespace() string {
	type getResult interface {
		GetNamespace() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNamespace()
	} else {
		var de string
		return de
	}
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchEntity3) GetNamespaceOk() (*string, bool) {
	type getResult interface {
		GetNamespaceOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNamespaceOk()
	} else {
		return nil, false
	}
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchEntity3) HasNamespace() bool {
	type getResult interface {
		HasNamespace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNamespace()
	} else {
		return false
	}
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchEntity3) SetNamespace(v string) {
	type getResult interface {
		SetNamespace(v string)
	}

	o.GetActualInstance().(getResult).SetNamespace(v)
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchEntity3) GetParameters() []BTMParameter1 {
	type getResult interface {
		GetParameters() []BTMParameter1
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameters()
	} else {
		var de []BTMParameter1
		return de
	}
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchEntity3) GetParametersOk() ([]BTMParameter1, bool) {
	type getResult interface {
		GetParametersOk() ([]BTMParameter1, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParametersOk()
	} else {
		return nil, false
	}
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchEntity3) HasParameters() bool {
	type getResult interface {
		HasParameters() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameters()
	} else {
		return false
	}
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchEntity3) SetParameters(v []BTMParameter1) {
	type getResult interface {
		SetParameters(v []BTMParameter1)
	}

	o.GetActualInstance().(getResult).SetParameters(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMSketchEntity3) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMSketchCompositeEntity-893'
	if jsonDict["btType"] == "BTMSketchCompositeEntity-893" {
		// try to unmarshal JSON data into BTMSketchCompositeEntity893
		var qr *BTMSketchCompositeEntity893
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMSketchEntity3 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMSketchEntity3 = nil
			return fmt.Errorf("Failed to unmarshal BTMSketchEntity3 as BTMSketchCompositeEntity893: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSketchConstraint-2'
	if jsonDict["btType"] == "BTMSketchConstraint-2" {
		// try to unmarshal JSON data into BTMSketchConstraint2
		var qr *BTMSketchConstraint2
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMSketchEntity3 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMSketchEntity3 = nil
			return fmt.Errorf("Failed to unmarshal BTMSketchEntity3 as BTMSketchConstraint2: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSketchGeomEntity-5'
	if jsonDict["btType"] == "BTMSketchGeomEntity-5" {
		// try to unmarshal JSON data into BTMSketchGeomEntity5
		var qr *BTMSketchGeomEntity5
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMSketchEntity3 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMSketchEntity3 = nil
			return fmt.Errorf("Failed to unmarshal BTMSketchEntity3 as BTMSketchGeomEntity5: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMSketchInvalid-1601'
	if jsonDict["btType"] == "BTMSketchInvalid-1601" {
		// try to unmarshal JSON data into BTMSketchInvalid1601
		var qr *BTMSketchInvalid1601
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMSketchEntity3 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMSketchEntity3 = nil
			return fmt.Errorf("Failed to unmarshal BTMSketchEntity3 as BTMSketchInvalid1601: %s", err.Error())
		}
	}

	var qtx *base_BTMSketchEntity3
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMSketchEntity3 = qtx
		return nil // data stored in dst.base_BTMSketchEntity3, return on the first match
	} else {
		dst.implBTMSketchEntity3 = nil
		return fmt.Errorf("Failed to unmarshal BTMSketchEntity3 as base_BTMSketchEntity3: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMSketchEntity3) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMSketchEntity3) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMSketchEntity3
}

type NullableBTMSketchEntity3 struct {
	value *BTMSketchEntity3
	isSet bool
}

func (v NullableBTMSketchEntity3) Get() *BTMSketchEntity3 {
	return v.value
}

func (v *NullableBTMSketchEntity3) Set(val *BTMSketchEntity3) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchEntity3) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchEntity3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchEntity3(val *BTMSketchEntity3) *NullableBTMSketchEntity3 {
	return &NullableBTMSketchEntity3{value: val, isSet: true}
}

func (v NullableBTMSketchEntity3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchEntity3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMSketchEntity3 struct {
	BtType                              *string         `json:"btType,omitempty"`
	ImportMicroversion                  *string         `json:"importMicroversion,omitempty"`
	NodeId                              *string         `json:"nodeId,omitempty"`
	EntityId                            *string         `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string         `json:"entityIdAndReplaceInDependentFields,omitempty"`
	Namespace                           *string         `json:"namespace,omitempty"`
	Parameters                          []BTMParameter1 `json:"parameters,omitempty"`
}

// Newbase_BTMSketchEntity3 instantiates a new base_BTMSketchEntity3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMSketchEntity3() *base_BTMSketchEntity3 {
	this := base_BTMSketchEntity3{}
	return &this
}

// Newbase_BTMSketchEntity3WithDefaults instantiates a new base_BTMSketchEntity3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMSketchEntity3WithDefaults() *base_BTMSketchEntity3 {
	this := base_BTMSketchEntity3{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMSketchEntity3) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMSketchEntity3) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMSketchEntity3) SetNodeId(v string) {
	o.NodeId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *base_BTMSketchEntity3) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *base_BTMSketchEntity3) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *base_BTMSketchEntity3) SetNamespace(v string) {
	o.Namespace = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *base_BTMSketchEntity3) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSketchEntity3) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *base_BTMSketchEntity3) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *base_BTMSketchEntity3) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

func (o base_BTMSketchEntity3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}