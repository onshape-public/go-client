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

// BTMFeatureInvalid1031 struct for BTMFeatureInvalid1031
type BTMFeatureInvalid1031 struct {
	FeatureId              *string         `json:"featureId,omitempty"`
	FeatureType            *string         `json:"featureType,omitempty"`
	ImportMicroversion     *string         `json:"importMicroversion,omitempty"`
	Name                   *string         `json:"name,omitempty"`
	Namespace              *string         `json:"namespace,omitempty"`
	NodeId                 *string         `json:"nodeId,omitempty"`
	Parameters             []BTMParameter1 `json:"parameters,omitempty"`
	ReturnAfterSubfeatures *bool           `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures            []BTMFeature134 `json:"subFeatures,omitempty"`
	Suppressed             *bool           `json:"suppressed,omitempty"`
	SuppressionConfigured  *bool           `json:"suppressionConfigured,omitempty"`
	BtType                 *string         `json:"btType,omitempty"`
}

// NewBTMFeatureInvalid1031 instantiates a new BTMFeatureInvalid1031 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMFeatureInvalid1031() *BTMFeatureInvalid1031 {
	this := BTMFeatureInvalid1031{}
	return &this
}

// NewBTMFeatureInvalid1031WithDefaults instantiates a new BTMFeatureInvalid1031 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMFeatureInvalid1031WithDefaults() *BTMFeatureInvalid1031 {
	this := BTMFeatureInvalid1031{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMFeatureInvalid1031) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTMFeatureInvalid1031) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMFeatureInvalid1031) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMFeatureInvalid1031) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMFeatureInvalid1031) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMFeatureInvalid1031) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMFeatureInvalid1031) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetReturnAfterSubfeatures() bool {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		var ret bool
		return ret
	}
	return *o.ReturnAfterSubfeatures
}

// GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetReturnAfterSubfeaturesOk() (*bool, bool) {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		return nil, false
	}
	return o.ReturnAfterSubfeatures, true
}

// HasReturnAfterSubfeatures returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasReturnAfterSubfeatures() bool {
	if o != nil && o.ReturnAfterSubfeatures != nil {
		return true
	}

	return false
}

// SetReturnAfterSubfeatures gets a reference to the given bool and assigns it to the ReturnAfterSubfeatures field.
func (o *BTMFeatureInvalid1031) SetReturnAfterSubfeatures(v bool) {
	o.ReturnAfterSubfeatures = &v
}

// GetSubFeatures returns the SubFeatures field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetSubFeatures() []BTMFeature134 {
	if o == nil || o.SubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.SubFeatures == nil {
		return nil, false
	}
	return o.SubFeatures, true
}

// HasSubFeatures returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasSubFeatures() bool {
	if o != nil && o.SubFeatures != nil {
		return true
	}

	return false
}

// SetSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the SubFeatures field.
func (o *BTMFeatureInvalid1031) SetSubFeatures(v []BTMFeature134) {
	o.SubFeatures = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTMFeatureInvalid1031) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTMFeatureInvalid1031) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMFeatureInvalid1031) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFeatureInvalid1031) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMFeatureInvalid1031) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMFeatureInvalid1031) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMFeatureInvalid1031) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.ReturnAfterSubfeatures != nil {
		toSerialize["returnAfterSubfeatures"] = o.ReturnAfterSubfeatures
	}
	if o.SubFeatures != nil {
		toSerialize["subFeatures"] = o.SubFeatures
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMFeatureInvalid1031 struct {
	value *BTMFeatureInvalid1031
	isSet bool
}

func (v NullableBTMFeatureInvalid1031) Get() *BTMFeatureInvalid1031 {
	return v.value
}

func (v *NullableBTMFeatureInvalid1031) Set(val *BTMFeatureInvalid1031) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMFeatureInvalid1031) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMFeatureInvalid1031) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMFeatureInvalid1031(val *BTMFeatureInvalid1031) *NullableBTMFeatureInvalid1031 {
	return &NullableBTMFeatureInvalid1031{value: val, isSet: true}
}

func (v NullableBTMFeatureInvalid1031) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMFeatureInvalid1031) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}