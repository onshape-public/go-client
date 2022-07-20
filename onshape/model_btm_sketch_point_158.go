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
)

// BTMSketchPoint158 struct for BTMSketchPoint158
type BTMSketchPoint158 struct {
	BtType                              *string         `json:"btType,omitempty"`
	ControlBoxIds                       []string        `json:"controlBoxIds,omitempty"`
	EntityId                            *string         `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string         `json:"entityIdAndReplaceInDependentFields,omitempty"`
	ImportMicroversion                  *string         `json:"importMicroversion,omitempty"`
	IsConstruction                      *bool           `json:"isConstruction,omitempty"`
	IsFromEndpointSplineHandle          *bool           `json:"isFromEndpointSplineHandle,omitempty"`
	IsFromSplineControlPolygon          *bool           `json:"isFromSplineControlPolygon,omitempty"`
	IsFromSplineHandle                  *bool           `json:"isFromSplineHandle,omitempty"`
	Namespace                           *string         `json:"namespace,omitempty"`
	NodeId                              *string         `json:"nodeId,omitempty"`
	Parameters                          []BTMParameter1 `json:"parameters,omitempty"`
	IsUserPoint                         *bool           `json:"isUserPoint,omitempty"`
	X                                   *float64        `json:"x,omitempty"`
	Y                                   *float64        `json:"y,omitempty"`
}

// NewBTMSketchPoint158 instantiates a new BTMSketchPoint158 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchPoint158() *BTMSketchPoint158 {
	this := BTMSketchPoint158{}
	return &this
}

// NewBTMSketchPoint158WithDefaults instantiates a new BTMSketchPoint158 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchPoint158WithDefaults() *BTMSketchPoint158 {
	this := BTMSketchPoint158{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchPoint158) SetBtType(v string) {
	o.BtType = &v
}

// GetControlBoxIds returns the ControlBoxIds field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetControlBoxIds() []string {
	if o == nil || o.ControlBoxIds == nil {
		var ret []string
		return ret
	}
	return o.ControlBoxIds
}

// GetControlBoxIdsOk returns a tuple with the ControlBoxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetControlBoxIdsOk() ([]string, bool) {
	if o == nil || o.ControlBoxIds == nil {
		return nil, false
	}
	return o.ControlBoxIds, true
}

// HasControlBoxIds returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasControlBoxIds() bool {
	if o != nil && o.ControlBoxIds != nil {
		return true
	}

	return false
}

// SetControlBoxIds gets a reference to the given []string and assigns it to the ControlBoxIds field.
func (o *BTMSketchPoint158) SetControlBoxIds(v []string) {
	o.ControlBoxIds = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchPoint158) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchPoint158) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSketchPoint158) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTMSketchPoint158) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

// GetIsFromEndpointSplineHandle returns the IsFromEndpointSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetIsFromEndpointSplineHandle() bool {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromEndpointSplineHandle
}

// GetIsFromEndpointSplineHandleOk returns a tuple with the IsFromEndpointSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetIsFromEndpointSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		return nil, false
	}
	return o.IsFromEndpointSplineHandle, true
}

// HasIsFromEndpointSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasIsFromEndpointSplineHandle() bool {
	if o != nil && o.IsFromEndpointSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromEndpointSplineHandle gets a reference to the given bool and assigns it to the IsFromEndpointSplineHandle field.
func (o *BTMSketchPoint158) SetIsFromEndpointSplineHandle(v bool) {
	o.IsFromEndpointSplineHandle = &v
}

// GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetIsFromSplineControlPolygon() bool {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineControlPolygon
}

// GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetIsFromSplineControlPolygonOk() (*bool, bool) {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		return nil, false
	}
	return o.IsFromSplineControlPolygon, true
}

// HasIsFromSplineControlPolygon returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasIsFromSplineControlPolygon() bool {
	if o != nil && o.IsFromSplineControlPolygon != nil {
		return true
	}

	return false
}

// SetIsFromSplineControlPolygon gets a reference to the given bool and assigns it to the IsFromSplineControlPolygon field.
func (o *BTMSketchPoint158) SetIsFromSplineControlPolygon(v bool) {
	o.IsFromSplineControlPolygon = &v
}

// GetIsFromSplineHandle returns the IsFromSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetIsFromSplineHandle() bool {
	if o == nil || o.IsFromSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineHandle
}

// GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetIsFromSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromSplineHandle == nil {
		return nil, false
	}
	return o.IsFromSplineHandle, true
}

// HasIsFromSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasIsFromSplineHandle() bool {
	if o != nil && o.IsFromSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromSplineHandle gets a reference to the given bool and assigns it to the IsFromSplineHandle field.
func (o *BTMSketchPoint158) SetIsFromSplineHandle(v bool) {
	o.IsFromSplineHandle = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchPoint158) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSketchPoint158) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchPoint158) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetIsUserPoint returns the IsUserPoint field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetIsUserPoint() bool {
	if o == nil || o.IsUserPoint == nil {
		var ret bool
		return ret
	}
	return *o.IsUserPoint
}

// GetIsUserPointOk returns a tuple with the IsUserPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetIsUserPointOk() (*bool, bool) {
	if o == nil || o.IsUserPoint == nil {
		return nil, false
	}
	return o.IsUserPoint, true
}

// HasIsUserPoint returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasIsUserPoint() bool {
	if o != nil && o.IsUserPoint != nil {
		return true
	}

	return false
}

// SetIsUserPoint gets a reference to the given bool and assigns it to the IsUserPoint field.
func (o *BTMSketchPoint158) SetIsUserPoint(v bool) {
	o.IsUserPoint = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *BTMSketchPoint158) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *BTMSketchPoint158) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchPoint158) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *BTMSketchPoint158) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *BTMSketchPoint158) SetY(v float64) {
	o.Y = &v
}

func (o BTMSketchPoint158) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ControlBoxIds != nil {
		toSerialize["controlBoxIds"] = o.ControlBoxIds
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	if o.IsFromEndpointSplineHandle != nil {
		toSerialize["isFromEndpointSplineHandle"] = o.IsFromEndpointSplineHandle
	}
	if o.IsFromSplineControlPolygon != nil {
		toSerialize["isFromSplineControlPolygon"] = o.IsFromSplineControlPolygon
	}
	if o.IsFromSplineHandle != nil {
		toSerialize["isFromSplineHandle"] = o.IsFromSplineHandle
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.IsUserPoint != nil {
		toSerialize["isUserPoint"] = o.IsUserPoint
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchPoint158 struct {
	value *BTMSketchPoint158
	isSet bool
}

func (v NullableBTMSketchPoint158) Get() *BTMSketchPoint158 {
	return v.value
}

func (v *NullableBTMSketchPoint158) Set(val *BTMSketchPoint158) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchPoint158) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchPoint158) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchPoint158(val *BTMSketchPoint158) *NullableBTMSketchPoint158 {
	return &NullableBTMSketchPoint158{value: val, isSet: true}
}

func (v NullableBTMSketchPoint158) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchPoint158) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}