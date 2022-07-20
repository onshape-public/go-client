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

// BTFeatureListResponse2457 struct for BTFeatureListResponse2457
type BTFeatureListResponse2457 struct {
	DefaultFeatures        []BTMFeature134                `json:"defaultFeatures,omitempty"`
	FeatureStates          *map[string]BTFeatureState1688 `json:"featureStates,omitempty"`
	Features               []BTMFeature134                `json:"features,omitempty"`
	Imports                []BTMImport136                 `json:"imports,omitempty"`
	IsComplete             *bool                          `json:"isComplete,omitempty"`
	LibraryVersion         *int32                         `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool                          `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew *bool                          `json:"rejectMicroversionSkew,omitempty"`
	RollbackIndex          *int32                         `json:"rollbackIndex,omitempty"`
	SerializationVersion   *string                        `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string                        `json:"sourceMicroversion,omitempty"`
}

// NewBTFeatureListResponse2457 instantiates a new BTFeatureListResponse2457 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureListResponse2457() *BTFeatureListResponse2457 {
	this := BTFeatureListResponse2457{}
	return &this
}

// NewBTFeatureListResponse2457WithDefaults instantiates a new BTFeatureListResponse2457 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureListResponse2457WithDefaults() *BTFeatureListResponse2457 {
	this := BTFeatureListResponse2457{}
	return &this
}

// GetDefaultFeatures returns the DefaultFeatures field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetDefaultFeatures() []BTMFeature134 {
	if o == nil || o.DefaultFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.DefaultFeatures
}

// GetDefaultFeaturesOk returns a tuple with the DefaultFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetDefaultFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.DefaultFeatures == nil {
		return nil, false
	}
	return o.DefaultFeatures, true
}

// HasDefaultFeatures returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasDefaultFeatures() bool {
	if o != nil && o.DefaultFeatures != nil {
		return true
	}

	return false
}

// SetDefaultFeatures gets a reference to the given []BTMFeature134 and assigns it to the DefaultFeatures field.
func (o *BTFeatureListResponse2457) SetDefaultFeatures(v []BTMFeature134) {
	o.DefaultFeatures = v
}

// GetFeatureStates returns the FeatureStates field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetFeatureStates() map[string]BTFeatureState1688 {
	if o == nil || o.FeatureStates == nil {
		var ret map[string]BTFeatureState1688
		return ret
	}
	return *o.FeatureStates
}

// GetFeatureStatesOk returns a tuple with the FeatureStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool) {
	if o == nil || o.FeatureStates == nil {
		return nil, false
	}
	return o.FeatureStates, true
}

// HasFeatureStates returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasFeatureStates() bool {
	if o != nil && o.FeatureStates != nil {
		return true
	}

	return false
}

// SetFeatureStates gets a reference to the given map[string]BTFeatureState1688 and assigns it to the FeatureStates field.
func (o *BTFeatureListResponse2457) SetFeatureStates(v map[string]BTFeatureState1688) {
	o.FeatureStates = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetFeatures() []BTMFeature134 {
	if o == nil || o.Features == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTMFeature134 and assigns it to the Features field.
func (o *BTFeatureListResponse2457) SetFeatures(v []BTMFeature134) {
	o.Features = v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetImports() []BTMImport136 {
	if o == nil || o.Imports == nil {
		var ret []BTMImport136
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetImportsOk() ([]BTMImport136, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []BTMImport136 and assigns it to the Imports field.
func (o *BTFeatureListResponse2457) SetImports(v []BTMImport136) {
	o.Imports = v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetIsComplete() bool {
	if o == nil || o.IsComplete == nil {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetIsCompleteOk() (*bool, bool) {
	if o == nil || o.IsComplete == nil {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasIsComplete() bool {
	if o != nil && o.IsComplete != nil {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *BTFeatureListResponse2457) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTFeatureListResponse2457) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTFeatureListResponse2457) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTFeatureListResponse2457) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetRollbackIndex returns the RollbackIndex field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetRollbackIndex() int32 {
	if o == nil || o.RollbackIndex == nil {
		var ret int32
		return ret
	}
	return *o.RollbackIndex
}

// GetRollbackIndexOk returns a tuple with the RollbackIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetRollbackIndexOk() (*int32, bool) {
	if o == nil || o.RollbackIndex == nil {
		return nil, false
	}
	return o.RollbackIndex, true
}

// HasRollbackIndex returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasRollbackIndex() bool {
	if o != nil && o.RollbackIndex != nil {
		return true
	}

	return false
}

// SetRollbackIndex gets a reference to the given int32 and assigns it to the RollbackIndex field.
func (o *BTFeatureListResponse2457) SetRollbackIndex(v int32) {
	o.RollbackIndex = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTFeatureListResponse2457) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTFeatureListResponse2457) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureListResponse2457) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTFeatureListResponse2457) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTFeatureListResponse2457) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTFeatureListResponse2457) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultFeatures != nil {
		toSerialize["defaultFeatures"] = o.DefaultFeatures
	}
	if o.FeatureStates != nil {
		toSerialize["featureStates"] = o.FeatureStates
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.IsComplete != nil {
		toSerialize["isComplete"] = o.IsComplete
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.RollbackIndex != nil {
		toSerialize["rollbackIndex"] = o.RollbackIndex
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureListResponse2457 struct {
	value *BTFeatureListResponse2457
	isSet bool
}

func (v NullableBTFeatureListResponse2457) Get() *BTFeatureListResponse2457 {
	return v.value
}

func (v *NullableBTFeatureListResponse2457) Set(val *BTFeatureListResponse2457) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureListResponse2457) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureListResponse2457) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureListResponse2457(val *BTFeatureListResponse2457) *NullableBTFeatureListResponse2457 {
	return &NullableBTFeatureListResponse2457{value: val, isSet: true}
}

func (v NullableBTFeatureListResponse2457) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureListResponse2457) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}