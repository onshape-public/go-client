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

// BTExportTessellatedFacesBody1321 struct for BTExportTessellatedFacesBody1321
type BTExportTessellatedFacesBody1321 struct {
	Constituents []string                           `json:"constituents,omitempty"`
	Id           *string                            `json:"id,omitempty"`
	Name         *string                            `json:"name,omitempty"`
	Appearance   *BTGraphicsAppearance1152          `json:"appearance,omitempty"`
	BodyType     *string                            `json:"bodyType,omitempty"`
	BtType       *string                            `json:"btType,omitempty"`
	Faces        []BTExportTessellatedFacesFace1192 `json:"faces,omitempty"`
	FacetPoints  []BTVector3d389                    `json:"facetPoints,omitempty"`
}

// NewBTExportTessellatedFacesBody1321 instantiates a new BTExportTessellatedFacesBody1321 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedFacesBody1321() *BTExportTessellatedFacesBody1321 {
	this := BTExportTessellatedFacesBody1321{}
	return &this
}

// NewBTExportTessellatedFacesBody1321WithDefaults instantiates a new BTExportTessellatedFacesBody1321 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedFacesBody1321WithDefaults() *BTExportTessellatedFacesBody1321 {
	this := BTExportTessellatedFacesBody1321{}
	return &this
}

// GetConstituents returns the Constituents field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetConstituents() []string {
	if o == nil || o.Constituents == nil {
		var ret []string
		return ret
	}
	return o.Constituents
}

// GetConstituentsOk returns a tuple with the Constituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetConstituentsOk() ([]string, bool) {
	if o == nil || o.Constituents == nil {
		return nil, false
	}
	return o.Constituents, true
}

// HasConstituents returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasConstituents() bool {
	if o != nil && o.Constituents != nil {
		return true
	}

	return false
}

// SetConstituents gets a reference to the given []string and assigns it to the Constituents field.
func (o *BTExportTessellatedFacesBody1321) SetConstituents(v []string) {
	o.Constituents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportTessellatedFacesBody1321) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTExportTessellatedFacesBody1321) SetName(v string) {
	o.Name = &v
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTExportTessellatedFacesBody1321) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *BTExportTessellatedFacesBody1321) SetBodyType(v string) {
	o.BodyType = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportTessellatedFacesBody1321) SetBtType(v string) {
	o.BtType = &v
}

// GetFaces returns the Faces field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetFaces() []BTExportTessellatedFacesFace1192 {
	if o == nil || o.Faces == nil {
		var ret []BTExportTessellatedFacesFace1192
		return ret
	}
	return o.Faces
}

// GetFacesOk returns a tuple with the Faces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetFacesOk() ([]BTExportTessellatedFacesFace1192, bool) {
	if o == nil || o.Faces == nil {
		return nil, false
	}
	return o.Faces, true
}

// HasFaces returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasFaces() bool {
	if o != nil && o.Faces != nil {
		return true
	}

	return false
}

// SetFaces gets a reference to the given []BTExportTessellatedFacesFace1192 and assigns it to the Faces field.
func (o *BTExportTessellatedFacesBody1321) SetFaces(v []BTExportTessellatedFacesFace1192) {
	o.Faces = v
}

// GetFacetPoints returns the FacetPoints field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesBody1321) GetFacetPoints() []BTVector3d389 {
	if o == nil || o.FacetPoints == nil {
		var ret []BTVector3d389
		return ret
	}
	return o.FacetPoints
}

// GetFacetPointsOk returns a tuple with the FacetPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesBody1321) GetFacetPointsOk() ([]BTVector3d389, bool) {
	if o == nil || o.FacetPoints == nil {
		return nil, false
	}
	return o.FacetPoints, true
}

// HasFacetPoints returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesBody1321) HasFacetPoints() bool {
	if o != nil && o.FacetPoints != nil {
		return true
	}

	return false
}

// SetFacetPoints gets a reference to the given []BTVector3d389 and assigns it to the FacetPoints field.
func (o *BTExportTessellatedFacesBody1321) SetFacetPoints(v []BTVector3d389) {
	o.FacetPoints = v
}

func (o BTExportTessellatedFacesBody1321) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constituents != nil {
		toSerialize["constituents"] = o.Constituents
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Faces != nil {
		toSerialize["faces"] = o.Faces
	}
	if o.FacetPoints != nil {
		toSerialize["facetPoints"] = o.FacetPoints
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedFacesBody1321 struct {
	value *BTExportTessellatedFacesBody1321
	isSet bool
}

func (v NullableBTExportTessellatedFacesBody1321) Get() *BTExportTessellatedFacesBody1321 {
	return v.value
}

func (v *NullableBTExportTessellatedFacesBody1321) Set(val *BTExportTessellatedFacesBody1321) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedFacesBody1321) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedFacesBody1321) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedFacesBody1321(val *BTExportTessellatedFacesBody1321) *NullableBTExportTessellatedFacesBody1321 {
	return &NullableBTExportTessellatedFacesBody1321{value: val, isSet: true}
}

func (v NullableBTExportTessellatedFacesBody1321) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedFacesBody1321) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}