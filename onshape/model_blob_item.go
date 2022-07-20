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

// BlobItem struct for BlobItem
type BlobItem struct {
	BaseHref             *string `json:"baseHref,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	DocumentId           *string `json:"documentId,omitempty"`
	ElementId            *string `json:"elementId,omitempty"`
	ElementType          *string `json:"elementType,omitempty"`
	EncodedConfiguration *string `json:"encodedConfiguration,omitempty"`
	Id                   *string `json:"id,omitempty"`
	JsonType             string  `json:"jsonType"`
	PartId               *string `json:"partId,omitempty"`
	PartName             *string `json:"partName,omitempty"`
	PartNumber           *string `json:"partNumber,omitempty"`
	Revision             *string `json:"revision,omitempty"`
	RevisionId           *string `json:"revisionId,omitempty"`
	State                *int32  `json:"state,omitempty"`
	VersionId            *string `json:"versionId,omitempty"`
	VersionName          *string `json:"versionName,omitempty"`
	DataTypeForResponse  *string `json:"dataTypeForResponse,omitempty"`
	Filename             *string `json:"filename,omitempty"`
	ForeignDataId        *string `json:"foreignDataId,omitempty"`
	Href                 *string `json:"href,omitempty"`
	PrettyType           *string `json:"prettyType,omitempty"`
	SafeToShow           *bool   `json:"safeToShow,omitempty"`
	SpecifiedUnit        *string `json:"specifiedUnit,omitempty"`
	Unupdatable          *bool   `json:"unupdatable,omitempty"`
}

// NewBlobItem instantiates a new BlobItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobItem(jsonType string) *BlobItem {
	this := BlobItem{}
	this.JsonType = jsonType
	return &this
}

// NewBlobItemWithDefaults instantiates a new BlobItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobItemWithDefaults() *BlobItem {
	this := BlobItem{}
	return &this
}

// GetBaseHref returns the BaseHref field value if set, zero value otherwise.
func (o *BlobItem) GetBaseHref() string {
	if o == nil || o.BaseHref == nil {
		var ret string
		return ret
	}
	return *o.BaseHref
}

// GetBaseHrefOk returns a tuple with the BaseHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetBaseHrefOk() (*string, bool) {
	if o == nil || o.BaseHref == nil {
		return nil, false
	}
	return o.BaseHref, true
}

// HasBaseHref returns a boolean if a field has been set.
func (o *BlobItem) HasBaseHref() bool {
	if o != nil && o.BaseHref != nil {
		return true
	}

	return false
}

// SetBaseHref gets a reference to the given string and assigns it to the BaseHref field.
func (o *BlobItem) SetBaseHref(v string) {
	o.BaseHref = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *BlobItem) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *BlobItem) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *BlobItem) SetDataType(v string) {
	o.DataType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BlobItem) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BlobItem) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BlobItem) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BlobItem) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BlobItem) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BlobItem) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BlobItem) GetElementType() string {
	if o == nil || o.ElementType == nil {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetElementTypeOk() (*string, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BlobItem) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *BlobItem) SetElementType(v string) {
	o.ElementType = &v
}

// GetEncodedConfiguration returns the EncodedConfiguration field value if set, zero value otherwise.
func (o *BlobItem) GetEncodedConfiguration() string {
	if o == nil || o.EncodedConfiguration == nil {
		var ret string
		return ret
	}
	return *o.EncodedConfiguration
}

// GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetEncodedConfigurationOk() (*string, bool) {
	if o == nil || o.EncodedConfiguration == nil {
		return nil, false
	}
	return o.EncodedConfiguration, true
}

// HasEncodedConfiguration returns a boolean if a field has been set.
func (o *BlobItem) HasEncodedConfiguration() bool {
	if o != nil && o.EncodedConfiguration != nil {
		return true
	}

	return false
}

// SetEncodedConfiguration gets a reference to the given string and assigns it to the EncodedConfiguration field.
func (o *BlobItem) SetEncodedConfiguration(v string) {
	o.EncodedConfiguration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlobItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlobItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BlobItem) SetId(v string) {
	o.Id = &v
}

// GetJsonType returns the JsonType field value
func (o *BlobItem) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BlobItem) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BlobItem) SetJsonType(v string) {
	o.JsonType = v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BlobItem) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BlobItem) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BlobItem) SetPartId(v string) {
	o.PartId = &v
}

// GetPartName returns the PartName field value if set, zero value otherwise.
func (o *BlobItem) GetPartName() string {
	if o == nil || o.PartName == nil {
		var ret string
		return ret
	}
	return *o.PartName
}

// GetPartNameOk returns a tuple with the PartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetPartNameOk() (*string, bool) {
	if o == nil || o.PartName == nil {
		return nil, false
	}
	return o.PartName, true
}

// HasPartName returns a boolean if a field has been set.
func (o *BlobItem) HasPartName() bool {
	if o != nil && o.PartName != nil {
		return true
	}

	return false
}

// SetPartName gets a reference to the given string and assigns it to the PartName field.
func (o *BlobItem) SetPartName(v string) {
	o.PartName = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BlobItem) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BlobItem) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BlobItem) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BlobItem) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BlobItem) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BlobItem) SetRevision(v string) {
	o.Revision = &v
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *BlobItem) GetRevisionId() string {
	if o == nil || o.RevisionId == nil {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetRevisionIdOk() (*string, bool) {
	if o == nil || o.RevisionId == nil {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *BlobItem) HasRevisionId() bool {
	if o != nil && o.RevisionId != nil {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *BlobItem) SetRevisionId(v string) {
	o.RevisionId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BlobItem) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BlobItem) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BlobItem) SetState(v int32) {
	o.State = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BlobItem) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BlobItem) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BlobItem) SetVersionId(v string) {
	o.VersionId = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *BlobItem) GetVersionName() string {
	if o == nil || o.VersionName == nil {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetVersionNameOk() (*string, bool) {
	if o == nil || o.VersionName == nil {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *BlobItem) HasVersionName() bool {
	if o != nil && o.VersionName != nil {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *BlobItem) SetVersionName(v string) {
	o.VersionName = &v
}

// GetDataTypeForResponse returns the DataTypeForResponse field value if set, zero value otherwise.
func (o *BlobItem) GetDataTypeForResponse() string {
	if o == nil || o.DataTypeForResponse == nil {
		var ret string
		return ret
	}
	return *o.DataTypeForResponse
}

// GetDataTypeForResponseOk returns a tuple with the DataTypeForResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetDataTypeForResponseOk() (*string, bool) {
	if o == nil || o.DataTypeForResponse == nil {
		return nil, false
	}
	return o.DataTypeForResponse, true
}

// HasDataTypeForResponse returns a boolean if a field has been set.
func (o *BlobItem) HasDataTypeForResponse() bool {
	if o != nil && o.DataTypeForResponse != nil {
		return true
	}

	return false
}

// SetDataTypeForResponse gets a reference to the given string and assigns it to the DataTypeForResponse field.
func (o *BlobItem) SetDataTypeForResponse(v string) {
	o.DataTypeForResponse = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *BlobItem) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *BlobItem) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *BlobItem) SetFilename(v string) {
	o.Filename = &v
}

// GetForeignDataId returns the ForeignDataId field value if set, zero value otherwise.
func (o *BlobItem) GetForeignDataId() string {
	if o == nil || o.ForeignDataId == nil {
		var ret string
		return ret
	}
	return *o.ForeignDataId
}

// GetForeignDataIdOk returns a tuple with the ForeignDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetForeignDataIdOk() (*string, bool) {
	if o == nil || o.ForeignDataId == nil {
		return nil, false
	}
	return o.ForeignDataId, true
}

// HasForeignDataId returns a boolean if a field has been set.
func (o *BlobItem) HasForeignDataId() bool {
	if o != nil && o.ForeignDataId != nil {
		return true
	}

	return false
}

// SetForeignDataId gets a reference to the given string and assigns it to the ForeignDataId field.
func (o *BlobItem) SetForeignDataId(v string) {
	o.ForeignDataId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BlobItem) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BlobItem) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BlobItem) SetHref(v string) {
	o.Href = &v
}

// GetPrettyType returns the PrettyType field value if set, zero value otherwise.
func (o *BlobItem) GetPrettyType() string {
	if o == nil || o.PrettyType == nil {
		var ret string
		return ret
	}
	return *o.PrettyType
}

// GetPrettyTypeOk returns a tuple with the PrettyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetPrettyTypeOk() (*string, bool) {
	if o == nil || o.PrettyType == nil {
		return nil, false
	}
	return o.PrettyType, true
}

// HasPrettyType returns a boolean if a field has been set.
func (o *BlobItem) HasPrettyType() bool {
	if o != nil && o.PrettyType != nil {
		return true
	}

	return false
}

// SetPrettyType gets a reference to the given string and assigns it to the PrettyType field.
func (o *BlobItem) SetPrettyType(v string) {
	o.PrettyType = &v
}

// GetSafeToShow returns the SafeToShow field value if set, zero value otherwise.
func (o *BlobItem) GetSafeToShow() bool {
	if o == nil || o.SafeToShow == nil {
		var ret bool
		return ret
	}
	return *o.SafeToShow
}

// GetSafeToShowOk returns a tuple with the SafeToShow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetSafeToShowOk() (*bool, bool) {
	if o == nil || o.SafeToShow == nil {
		return nil, false
	}
	return o.SafeToShow, true
}

// HasSafeToShow returns a boolean if a field has been set.
func (o *BlobItem) HasSafeToShow() bool {
	if o != nil && o.SafeToShow != nil {
		return true
	}

	return false
}

// SetSafeToShow gets a reference to the given bool and assigns it to the SafeToShow field.
func (o *BlobItem) SetSafeToShow(v bool) {
	o.SafeToShow = &v
}

// GetSpecifiedUnit returns the SpecifiedUnit field value if set, zero value otherwise.
func (o *BlobItem) GetSpecifiedUnit() string {
	if o == nil || o.SpecifiedUnit == nil {
		var ret string
		return ret
	}
	return *o.SpecifiedUnit
}

// GetSpecifiedUnitOk returns a tuple with the SpecifiedUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetSpecifiedUnitOk() (*string, bool) {
	if o == nil || o.SpecifiedUnit == nil {
		return nil, false
	}
	return o.SpecifiedUnit, true
}

// HasSpecifiedUnit returns a boolean if a field has been set.
func (o *BlobItem) HasSpecifiedUnit() bool {
	if o != nil && o.SpecifiedUnit != nil {
		return true
	}

	return false
}

// SetSpecifiedUnit gets a reference to the given string and assigns it to the SpecifiedUnit field.
func (o *BlobItem) SetSpecifiedUnit(v string) {
	o.SpecifiedUnit = &v
}

// GetUnupdatable returns the Unupdatable field value if set, zero value otherwise.
func (o *BlobItem) GetUnupdatable() bool {
	if o == nil || o.Unupdatable == nil {
		var ret bool
		return ret
	}
	return *o.Unupdatable
}

// GetUnupdatableOk returns a tuple with the Unupdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobItem) GetUnupdatableOk() (*bool, bool) {
	if o == nil || o.Unupdatable == nil {
		return nil, false
	}
	return o.Unupdatable, true
}

// HasUnupdatable returns a boolean if a field has been set.
func (o *BlobItem) HasUnupdatable() bool {
	if o != nil && o.Unupdatable != nil {
		return true
	}

	return false
}

// SetUnupdatable gets a reference to the given bool and assigns it to the Unupdatable field.
func (o *BlobItem) SetUnupdatable(v bool) {
	o.Unupdatable = &v
}

func (o BlobItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseHref != nil {
		toSerialize["baseHref"] = o.BaseHref
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.EncodedConfiguration != nil {
		toSerialize["encodedConfiguration"] = o.EncodedConfiguration
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartName != nil {
		toSerialize["partName"] = o.PartName
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.RevisionId != nil {
		toSerialize["revisionId"] = o.RevisionId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.VersionName != nil {
		toSerialize["versionName"] = o.VersionName
	}
	if o.DataTypeForResponse != nil {
		toSerialize["dataTypeForResponse"] = o.DataTypeForResponse
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.ForeignDataId != nil {
		toSerialize["foreignDataId"] = o.ForeignDataId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.PrettyType != nil {
		toSerialize["prettyType"] = o.PrettyType
	}
	if o.SafeToShow != nil {
		toSerialize["safeToShow"] = o.SafeToShow
	}
	if o.SpecifiedUnit != nil {
		toSerialize["specifiedUnit"] = o.SpecifiedUnit
	}
	if o.Unupdatable != nil {
		toSerialize["unupdatable"] = o.Unupdatable
	}
	return json.Marshal(toSerialize)
}

type NullableBlobItem struct {
	value *BlobItem
	isSet bool
}

func (v NullableBlobItem) Get() *BlobItem {
	return v.value
}

func (v *NullableBlobItem) Set(val *BlobItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobItem(val *BlobItem) *NullableBlobItem {
	return &NullableBlobItem{value: val, isSet: true}
}

func (v NullableBlobItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}