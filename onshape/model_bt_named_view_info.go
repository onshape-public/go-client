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

// BTNamedViewInfo struct for BTNamedViewInfo
type BTNamedViewInfo struct {
	Angle           *float64                            `json:"angle,omitempty"`
	CameraViewport  []float64                           `json:"cameraViewport,omitempty"`
	Perspective     *bool                               `json:"perspective,omitempty"`
	SectionPlanes   []BTSectionPlaneInfo                `json:"sectionPlanes,omitempty"`
	SectionViewData *BTGraphicsSectionViewStateData4379 `json:"sectionViewData,omitempty"`
	ViewMatrix      []float64                           `json:"viewMatrix,omitempty"`
}

// NewBTNamedViewInfo instantiates a new BTNamedViewInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNamedViewInfo() *BTNamedViewInfo {
	this := BTNamedViewInfo{}
	return &this
}

// NewBTNamedViewInfoWithDefaults instantiates a new BTNamedViewInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNamedViewInfoWithDefaults() *BTNamedViewInfo {
	this := BTNamedViewInfo{}
	return &this
}

// GetAngle returns the Angle field value if set, zero value otherwise.
func (o *BTNamedViewInfo) GetAngle() float64 {
	if o == nil || o.Angle == nil {
		var ret float64
		return ret
	}
	return *o.Angle
}

// GetAngleOk returns a tuple with the Angle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewInfo) GetAngleOk() (*float64, bool) {
	if o == nil || o.Angle == nil {
		return nil, false
	}
	return o.Angle, true
}

// HasAngle returns a boolean if a field has been set.
func (o *BTNamedViewInfo) HasAngle() bool {
	if o != nil && o.Angle != nil {
		return true
	}

	return false
}

// SetAngle gets a reference to the given float64 and assigns it to the Angle field.
func (o *BTNamedViewInfo) SetAngle(v float64) {
	o.Angle = &v
}

// GetCameraViewport returns the CameraViewport field value if set, zero value otherwise.
func (o *BTNamedViewInfo) GetCameraViewport() []float64 {
	if o == nil || o.CameraViewport == nil {
		var ret []float64
		return ret
	}
	return o.CameraViewport
}

// GetCameraViewportOk returns a tuple with the CameraViewport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewInfo) GetCameraViewportOk() ([]float64, bool) {
	if o == nil || o.CameraViewport == nil {
		return nil, false
	}
	return o.CameraViewport, true
}

// HasCameraViewport returns a boolean if a field has been set.
func (o *BTNamedViewInfo) HasCameraViewport() bool {
	if o != nil && o.CameraViewport != nil {
		return true
	}

	return false
}

// SetCameraViewport gets a reference to the given []float64 and assigns it to the CameraViewport field.
func (o *BTNamedViewInfo) SetCameraViewport(v []float64) {
	o.CameraViewport = v
}

// GetPerspective returns the Perspective field value if set, zero value otherwise.
func (o *BTNamedViewInfo) GetPerspective() bool {
	if o == nil || o.Perspective == nil {
		var ret bool
		return ret
	}
	return *o.Perspective
}

// GetPerspectiveOk returns a tuple with the Perspective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewInfo) GetPerspectiveOk() (*bool, bool) {
	if o == nil || o.Perspective == nil {
		return nil, false
	}
	return o.Perspective, true
}

// HasPerspective returns a boolean if a field has been set.
func (o *BTNamedViewInfo) HasPerspective() bool {
	if o != nil && o.Perspective != nil {
		return true
	}

	return false
}

// SetPerspective gets a reference to the given bool and assigns it to the Perspective field.
func (o *BTNamedViewInfo) SetPerspective(v bool) {
	o.Perspective = &v
}

// GetSectionPlanes returns the SectionPlanes field value if set, zero value otherwise.
func (o *BTNamedViewInfo) GetSectionPlanes() []BTSectionPlaneInfo {
	if o == nil || o.SectionPlanes == nil {
		var ret []BTSectionPlaneInfo
		return ret
	}
	return o.SectionPlanes
}

// GetSectionPlanesOk returns a tuple with the SectionPlanes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewInfo) GetSectionPlanesOk() ([]BTSectionPlaneInfo, bool) {
	if o == nil || o.SectionPlanes == nil {
		return nil, false
	}
	return o.SectionPlanes, true
}

// HasSectionPlanes returns a boolean if a field has been set.
func (o *BTNamedViewInfo) HasSectionPlanes() bool {
	if o != nil && o.SectionPlanes != nil {
		return true
	}

	return false
}

// SetSectionPlanes gets a reference to the given []BTSectionPlaneInfo and assigns it to the SectionPlanes field.
func (o *BTNamedViewInfo) SetSectionPlanes(v []BTSectionPlaneInfo) {
	o.SectionPlanes = v
}

// GetSectionViewData returns the SectionViewData field value if set, zero value otherwise.
func (o *BTNamedViewInfo) GetSectionViewData() BTGraphicsSectionViewStateData4379 {
	if o == nil || o.SectionViewData == nil {
		var ret BTGraphicsSectionViewStateData4379
		return ret
	}
	return *o.SectionViewData
}

// GetSectionViewDataOk returns a tuple with the SectionViewData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewInfo) GetSectionViewDataOk() (*BTGraphicsSectionViewStateData4379, bool) {
	if o == nil || o.SectionViewData == nil {
		return nil, false
	}
	return o.SectionViewData, true
}

// HasSectionViewData returns a boolean if a field has been set.
func (o *BTNamedViewInfo) HasSectionViewData() bool {
	if o != nil && o.SectionViewData != nil {
		return true
	}

	return false
}

// SetSectionViewData gets a reference to the given BTGraphicsSectionViewStateData4379 and assigns it to the SectionViewData field.
func (o *BTNamedViewInfo) SetSectionViewData(v BTGraphicsSectionViewStateData4379) {
	o.SectionViewData = &v
}

// GetViewMatrix returns the ViewMatrix field value if set, zero value otherwise.
func (o *BTNamedViewInfo) GetViewMatrix() []float64 {
	if o == nil || o.ViewMatrix == nil {
		var ret []float64
		return ret
	}
	return o.ViewMatrix
}

// GetViewMatrixOk returns a tuple with the ViewMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNamedViewInfo) GetViewMatrixOk() ([]float64, bool) {
	if o == nil || o.ViewMatrix == nil {
		return nil, false
	}
	return o.ViewMatrix, true
}

// HasViewMatrix returns a boolean if a field has been set.
func (o *BTNamedViewInfo) HasViewMatrix() bool {
	if o != nil && o.ViewMatrix != nil {
		return true
	}

	return false
}

// SetViewMatrix gets a reference to the given []float64 and assigns it to the ViewMatrix field.
func (o *BTNamedViewInfo) SetViewMatrix(v []float64) {
	o.ViewMatrix = v
}

func (o BTNamedViewInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Angle != nil {
		toSerialize["angle"] = o.Angle
	}
	if o.CameraViewport != nil {
		toSerialize["cameraViewport"] = o.CameraViewport
	}
	if o.Perspective != nil {
		toSerialize["perspective"] = o.Perspective
	}
	if o.SectionPlanes != nil {
		toSerialize["sectionPlanes"] = o.SectionPlanes
	}
	if o.SectionViewData != nil {
		toSerialize["sectionViewData"] = o.SectionViewData
	}
	if o.ViewMatrix != nil {
		toSerialize["viewMatrix"] = o.ViewMatrix
	}
	return json.Marshal(toSerialize)
}

type NullableBTNamedViewInfo struct {
	value *BTNamedViewInfo
	isSet bool
}

func (v NullableBTNamedViewInfo) Get() *BTNamedViewInfo {
	return v.value
}

func (v *NullableBTNamedViewInfo) Set(val *BTNamedViewInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNamedViewInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNamedViewInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNamedViewInfo(val *BTNamedViewInfo) *NullableBTNamedViewInfo {
	return &NullableBTNamedViewInfo{value: val, isSet: true}
}

func (v NullableBTNamedViewInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNamedViewInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}