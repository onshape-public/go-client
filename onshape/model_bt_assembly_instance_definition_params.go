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

// BTAssemblyInstanceDefinitionParams struct for BTAssemblyInstanceDefinitionParams
type BTAssemblyInstanceDefinitionParams struct {
	Configuration     *string  `json:"configuration,omitempty"`
	DocumentId        string   `json:"documentId"`
	ElementId         *string  `json:"elementId,omitempty"`
	FeatureId         *string  `json:"featureId,omitempty"`
	IncludePartTypes  []string `json:"includePartTypes,omitempty"`
	IsAssembly        *bool    `json:"isAssembly,omitempty"`
	IsHidden          *bool    `json:"isHidden,omitempty"`
	IsSuppressed      *bool    `json:"isSuppressed,omitempty"`
	IsWholePartStudio *bool    `json:"isWholePartStudio,omitempty"`
	MicroversionId    *string  `json:"microversionId,omitempty"`
	PartId            *string  `json:"partId,omitempty"`
	PartNumber        *string  `json:"partNumber,omitempty"`
	Revision          *string  `json:"revision,omitempty"`
	VersionId         *string  `json:"versionId,omitempty"`
}

// NewBTAssemblyInstanceDefinitionParams instantiates a new BTAssemblyInstanceDefinitionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyInstanceDefinitionParams(documentId string) *BTAssemblyInstanceDefinitionParams {
	this := BTAssemblyInstanceDefinitionParams{}
	this.DocumentId = documentId
	return &this
}

// NewBTAssemblyInstanceDefinitionParamsWithDefaults instantiates a new BTAssemblyInstanceDefinitionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyInstanceDefinitionParamsWithDefaults() *BTAssemblyInstanceDefinitionParams {
	this := BTAssemblyInstanceDefinitionParams{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTAssemblyInstanceDefinitionParams) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value
func (o *BTAssemblyInstanceDefinitionParams) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *BTAssemblyInstanceDefinitionParams) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyInstanceDefinitionParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTAssemblyInstanceDefinitionParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetIncludePartTypes returns the IncludePartTypes field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetIncludePartTypes() []string {
	if o == nil || o.IncludePartTypes == nil {
		var ret []string
		return ret
	}
	return o.IncludePartTypes
}

// GetIncludePartTypesOk returns a tuple with the IncludePartTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetIncludePartTypesOk() ([]string, bool) {
	if o == nil || o.IncludePartTypes == nil {
		return nil, false
	}
	return o.IncludePartTypes, true
}

// HasIncludePartTypes returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasIncludePartTypes() bool {
	if o != nil && o.IncludePartTypes != nil {
		return true
	}

	return false
}

// SetIncludePartTypes gets a reference to the given []string and assigns it to the IncludePartTypes field.
func (o *BTAssemblyInstanceDefinitionParams) SetIncludePartTypes(v []string) {
	o.IncludePartTypes = v
}

// GetIsAssembly returns the IsAssembly field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetIsAssembly() bool {
	if o == nil || o.IsAssembly == nil {
		var ret bool
		return ret
	}
	return *o.IsAssembly
}

// GetIsAssemblyOk returns a tuple with the IsAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetIsAssemblyOk() (*bool, bool) {
	if o == nil || o.IsAssembly == nil {
		return nil, false
	}
	return o.IsAssembly, true
}

// HasIsAssembly returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasIsAssembly() bool {
	if o != nil && o.IsAssembly != nil {
		return true
	}

	return false
}

// SetIsAssembly gets a reference to the given bool and assigns it to the IsAssembly field.
func (o *BTAssemblyInstanceDefinitionParams) SetIsAssembly(v bool) {
	o.IsAssembly = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasIsHidden() bool {
	if o != nil && o.IsHidden != nil {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *BTAssemblyInstanceDefinitionParams) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetIsSuppressed returns the IsSuppressed field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetIsSuppressed() bool {
	if o == nil || o.IsSuppressed == nil {
		var ret bool
		return ret
	}
	return *o.IsSuppressed
}

// GetIsSuppressedOk returns a tuple with the IsSuppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetIsSuppressedOk() (*bool, bool) {
	if o == nil || o.IsSuppressed == nil {
		return nil, false
	}
	return o.IsSuppressed, true
}

// HasIsSuppressed returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasIsSuppressed() bool {
	if o != nil && o.IsSuppressed != nil {
		return true
	}

	return false
}

// SetIsSuppressed gets a reference to the given bool and assigns it to the IsSuppressed field.
func (o *BTAssemblyInstanceDefinitionParams) SetIsSuppressed(v bool) {
	o.IsSuppressed = &v
}

// GetIsWholePartStudio returns the IsWholePartStudio field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetIsWholePartStudio() bool {
	if o == nil || o.IsWholePartStudio == nil {
		var ret bool
		return ret
	}
	return *o.IsWholePartStudio
}

// GetIsWholePartStudioOk returns a tuple with the IsWholePartStudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetIsWholePartStudioOk() (*bool, bool) {
	if o == nil || o.IsWholePartStudio == nil {
		return nil, false
	}
	return o.IsWholePartStudio, true
}

// HasIsWholePartStudio returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasIsWholePartStudio() bool {
	if o != nil && o.IsWholePartStudio != nil {
		return true
	}

	return false
}

// SetIsWholePartStudio gets a reference to the given bool and assigns it to the IsWholePartStudio field.
func (o *BTAssemblyInstanceDefinitionParams) SetIsWholePartStudio(v bool) {
	o.IsWholePartStudio = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTAssemblyInstanceDefinitionParams) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTAssemblyInstanceDefinitionParams) SetPartId(v string) {
	o.PartId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTAssemblyInstanceDefinitionParams) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAssemblyInstanceDefinitionParams) SetRevision(v string) {
	o.Revision = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceDefinitionParams) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceDefinitionParams) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceDefinitionParams) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTAssemblyInstanceDefinitionParams) SetVersionId(v string) {
	o.VersionId = &v
}

func (o BTAssemblyInstanceDefinitionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if true {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.IncludePartTypes != nil {
		toSerialize["includePartTypes"] = o.IncludePartTypes
	}
	if o.IsAssembly != nil {
		toSerialize["isAssembly"] = o.IsAssembly
	}
	if o.IsHidden != nil {
		toSerialize["isHidden"] = o.IsHidden
	}
	if o.IsSuppressed != nil {
		toSerialize["isSuppressed"] = o.IsSuppressed
	}
	if o.IsWholePartStudio != nil {
		toSerialize["isWholePartStudio"] = o.IsWholePartStudio
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyInstanceDefinitionParams struct {
	value *BTAssemblyInstanceDefinitionParams
	isSet bool
}

func (v NullableBTAssemblyInstanceDefinitionParams) Get() *BTAssemblyInstanceDefinitionParams {
	return v.value
}

func (v *NullableBTAssemblyInstanceDefinitionParams) Set(val *BTAssemblyInstanceDefinitionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyInstanceDefinitionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyInstanceDefinitionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyInstanceDefinitionParams(val *BTAssemblyInstanceDefinitionParams) *NullableBTAssemblyInstanceDefinitionParams {
	return &NullableBTAssemblyInstanceDefinitionParams{value: val, isSet: true}
}

func (v NullableBTAssemblyInstanceDefinitionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyInstanceDefinitionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}