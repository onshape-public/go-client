/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6820-1bef41f9cc03
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRadialDimensionDisplayData348 struct for BTRadialDimensionDisplayData348
type BTRadialDimensionDisplayData348 struct {
	CoordinateSystem     *BTMatrix3x3340 `json:"coordinateSystem,omitempty"`
	FeatureId            *string         `json:"featureId,omitempty"`
	HasMaximumLimit_     *bool           `json:"hasMaximumLimit,omitempty"`
	HasMinimumLimit_     *bool           `json:"hasMinimumLimit,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	IsAssociatedWithFlat *bool           `json:"isAssociatedWithFlat,omitempty"`
	IsDriven             *bool           `json:"isDriven,omitempty"`
	IsOverDefined        *bool           `json:"isOverDefined,omitempty"`
	MaximumLimit         *float64        `json:"maximumLimit,omitempty"`
	MinimumLimit         *float64        `json:"minimumLimit,omitempty"`
	ParameterId          *string         `json:"parameterId,omitempty"`
	PlaneMatrix          *BTBSMatrix386  `json:"planeMatrix,omitempty"`
	Value                *float64        `json:"value,omitempty"`
	BtType               *string         `json:"btType,omitempty"`
	PositionR            *float64        `json:"positionR,omitempty"`
	PositionT            *float64        `json:"positionT,omitempty"`
	RadiusDisplay        *string         `json:"radiusDisplay,omitempty"`
	RealDiameter         *bool           `json:"realDiameter,omitempty"`
	WitnessEndPoint0r    *float64        `json:"witnessEndPoint0r,omitempty"`
	WitnessEndPoint0t    *float64        `json:"witnessEndPoint0t,omitempty"`
	WitnessEndPoint1r    *float64        `json:"witnessEndPoint1r,omitempty"`
	WitnessEndPoint1t    *float64        `json:"witnessEndPoint1t,omitempty"`
}

// NewBTRadialDimensionDisplayData348 instantiates a new BTRadialDimensionDisplayData348 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRadialDimensionDisplayData348() *BTRadialDimensionDisplayData348 {
	this := BTRadialDimensionDisplayData348{}
	return &this
}

// NewBTRadialDimensionDisplayData348WithDefaults instantiates a new BTRadialDimensionDisplayData348 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRadialDimensionDisplayData348WithDefaults() *BTRadialDimensionDisplayData348 {
	this := BTRadialDimensionDisplayData348{}
	return &this
}

// GetCoordinateSystem returns the CoordinateSystem field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetCoordinateSystem() BTMatrix3x3340 {
	if o == nil || o.CoordinateSystem == nil {
		var ret BTMatrix3x3340
		return ret
	}
	return *o.CoordinateSystem
}

// GetCoordinateSystemOk returns a tuple with the CoordinateSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetCoordinateSystemOk() (*BTMatrix3x3340, bool) {
	if o == nil || o.CoordinateSystem == nil {
		return nil, false
	}
	return o.CoordinateSystem, true
}

// HasCoordinateSystem returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasCoordinateSystem() bool {
	if o != nil && o.CoordinateSystem != nil {
		return true
	}

	return false
}

// SetCoordinateSystem gets a reference to the given BTMatrix3x3340 and assigns it to the CoordinateSystem field.
func (o *BTRadialDimensionDisplayData348) SetCoordinateSystem(v BTMatrix3x3340) {
	o.CoordinateSystem = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTRadialDimensionDisplayData348) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetHasMaximumLimit_ returns the HasMaximumLimit_ field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetHasMaximumLimit_() bool {
	if o == nil || o.HasMaximumLimit_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMaximumLimit_
}

// GetHasMaximumLimit_Ok returns a tuple with the HasMaximumLimit_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetHasMaximumLimit_Ok() (*bool, bool) {
	if o == nil || o.HasMaximumLimit_ == nil {
		return nil, false
	}
	return o.HasMaximumLimit_, true
}

// HasHasMaximumLimit_ returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasHasMaximumLimit_() bool {
	if o != nil && o.HasMaximumLimit_ != nil {
		return true
	}

	return false
}

// SetHasMaximumLimit_ gets a reference to the given bool and assigns it to the HasMaximumLimit_ field.
func (o *BTRadialDimensionDisplayData348) SetHasMaximumLimit_(v bool) {
	o.HasMaximumLimit_ = &v
}

// GetHasMinimumLimit_ returns the HasMinimumLimit_ field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetHasMinimumLimit_() bool {
	if o == nil || o.HasMinimumLimit_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMinimumLimit_
}

// GetHasMinimumLimit_Ok returns a tuple with the HasMinimumLimit_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetHasMinimumLimit_Ok() (*bool, bool) {
	if o == nil || o.HasMinimumLimit_ == nil {
		return nil, false
	}
	return o.HasMinimumLimit_, true
}

// HasHasMinimumLimit_ returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasHasMinimumLimit_() bool {
	if o != nil && o.HasMinimumLimit_ != nil {
		return true
	}

	return false
}

// SetHasMinimumLimit_ gets a reference to the given bool and assigns it to the HasMinimumLimit_ field.
func (o *BTRadialDimensionDisplayData348) SetHasMinimumLimit_(v bool) {
	o.HasMinimumLimit_ = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTRadialDimensionDisplayData348) SetId(v string) {
	o.Id = &v
}

// GetIsAssociatedWithFlat returns the IsAssociatedWithFlat field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetIsAssociatedWithFlat() bool {
	if o == nil || o.IsAssociatedWithFlat == nil {
		var ret bool
		return ret
	}
	return *o.IsAssociatedWithFlat
}

// GetIsAssociatedWithFlatOk returns a tuple with the IsAssociatedWithFlat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetIsAssociatedWithFlatOk() (*bool, bool) {
	if o == nil || o.IsAssociatedWithFlat == nil {
		return nil, false
	}
	return o.IsAssociatedWithFlat, true
}

// HasIsAssociatedWithFlat returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasIsAssociatedWithFlat() bool {
	if o != nil && o.IsAssociatedWithFlat != nil {
		return true
	}

	return false
}

// SetIsAssociatedWithFlat gets a reference to the given bool and assigns it to the IsAssociatedWithFlat field.
func (o *BTRadialDimensionDisplayData348) SetIsAssociatedWithFlat(v bool) {
	o.IsAssociatedWithFlat = &v
}

// GetIsDriven returns the IsDriven field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetIsDriven() bool {
	if o == nil || o.IsDriven == nil {
		var ret bool
		return ret
	}
	return *o.IsDriven
}

// GetIsDrivenOk returns a tuple with the IsDriven field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetIsDrivenOk() (*bool, bool) {
	if o == nil || o.IsDriven == nil {
		return nil, false
	}
	return o.IsDriven, true
}

// HasIsDriven returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasIsDriven() bool {
	if o != nil && o.IsDriven != nil {
		return true
	}

	return false
}

// SetIsDriven gets a reference to the given bool and assigns it to the IsDriven field.
func (o *BTRadialDimensionDisplayData348) SetIsDriven(v bool) {
	o.IsDriven = &v
}

// GetIsOverDefined returns the IsOverDefined field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetIsOverDefined() bool {
	if o == nil || o.IsOverDefined == nil {
		var ret bool
		return ret
	}
	return *o.IsOverDefined
}

// GetIsOverDefinedOk returns a tuple with the IsOverDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetIsOverDefinedOk() (*bool, bool) {
	if o == nil || o.IsOverDefined == nil {
		return nil, false
	}
	return o.IsOverDefined, true
}

// HasIsOverDefined returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasIsOverDefined() bool {
	if o != nil && o.IsOverDefined != nil {
		return true
	}

	return false
}

// SetIsOverDefined gets a reference to the given bool and assigns it to the IsOverDefined field.
func (o *BTRadialDimensionDisplayData348) SetIsOverDefined(v bool) {
	o.IsOverDefined = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetMaximumLimit() float64 {
	if o == nil || o.MaximumLimit == nil {
		var ret float64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetMaximumLimitOk() (*float64, bool) {
	if o == nil || o.MaximumLimit == nil {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasMaximumLimit() bool {
	if o != nil && o.MaximumLimit != nil {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given float64 and assigns it to the MaximumLimit field.
func (o *BTRadialDimensionDisplayData348) SetMaximumLimit(v float64) {
	o.MaximumLimit = &v
}

// GetMinimumLimit returns the MinimumLimit field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetMinimumLimit() float64 {
	if o == nil || o.MinimumLimit == nil {
		var ret float64
		return ret
	}
	return *o.MinimumLimit
}

// GetMinimumLimitOk returns a tuple with the MinimumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetMinimumLimitOk() (*float64, bool) {
	if o == nil || o.MinimumLimit == nil {
		return nil, false
	}
	return o.MinimumLimit, true
}

// HasMinimumLimit returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasMinimumLimit() bool {
	if o != nil && o.MinimumLimit != nil {
		return true
	}

	return false
}

// SetMinimumLimit gets a reference to the given float64 and assigns it to the MinimumLimit field.
func (o *BTRadialDimensionDisplayData348) SetMinimumLimit(v float64) {
	o.MinimumLimit = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTRadialDimensionDisplayData348) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetPlaneMatrix returns the PlaneMatrix field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetPlaneMatrix() BTBSMatrix386 {
	if o == nil || o.PlaneMatrix == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.PlaneMatrix
}

// GetPlaneMatrixOk returns a tuple with the PlaneMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetPlaneMatrixOk() (*BTBSMatrix386, bool) {
	if o == nil || o.PlaneMatrix == nil {
		return nil, false
	}
	return o.PlaneMatrix, true
}

// HasPlaneMatrix returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasPlaneMatrix() bool {
	if o != nil && o.PlaneMatrix != nil {
		return true
	}

	return false
}

// SetPlaneMatrix gets a reference to the given BTBSMatrix386 and assigns it to the PlaneMatrix field.
func (o *BTRadialDimensionDisplayData348) SetPlaneMatrix(v BTBSMatrix386) {
	o.PlaneMatrix = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTRadialDimensionDisplayData348) SetValue(v float64) {
	o.Value = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTRadialDimensionDisplayData348) SetBtType(v string) {
	o.BtType = &v
}

// GetPositionR returns the PositionR field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetPositionR() float64 {
	if o == nil || o.PositionR == nil {
		var ret float64
		return ret
	}
	return *o.PositionR
}

// GetPositionROk returns a tuple with the PositionR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetPositionROk() (*float64, bool) {
	if o == nil || o.PositionR == nil {
		return nil, false
	}
	return o.PositionR, true
}

// HasPositionR returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasPositionR() bool {
	if o != nil && o.PositionR != nil {
		return true
	}

	return false
}

// SetPositionR gets a reference to the given float64 and assigns it to the PositionR field.
func (o *BTRadialDimensionDisplayData348) SetPositionR(v float64) {
	o.PositionR = &v
}

// GetPositionT returns the PositionT field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetPositionT() float64 {
	if o == nil || o.PositionT == nil {
		var ret float64
		return ret
	}
	return *o.PositionT
}

// GetPositionTOk returns a tuple with the PositionT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetPositionTOk() (*float64, bool) {
	if o == nil || o.PositionT == nil {
		return nil, false
	}
	return o.PositionT, true
}

// HasPositionT returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasPositionT() bool {
	if o != nil && o.PositionT != nil {
		return true
	}

	return false
}

// SetPositionT gets a reference to the given float64 and assigns it to the PositionT field.
func (o *BTRadialDimensionDisplayData348) SetPositionT(v float64) {
	o.PositionT = &v
}

// GetRadiusDisplay returns the RadiusDisplay field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetRadiusDisplay() string {
	if o == nil || o.RadiusDisplay == nil {
		var ret string
		return ret
	}
	return *o.RadiusDisplay
}

// GetRadiusDisplayOk returns a tuple with the RadiusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetRadiusDisplayOk() (*string, bool) {
	if o == nil || o.RadiusDisplay == nil {
		return nil, false
	}
	return o.RadiusDisplay, true
}

// HasRadiusDisplay returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasRadiusDisplay() bool {
	if o != nil && o.RadiusDisplay != nil {
		return true
	}

	return false
}

// SetRadiusDisplay gets a reference to the given string and assigns it to the RadiusDisplay field.
func (o *BTRadialDimensionDisplayData348) SetRadiusDisplay(v string) {
	o.RadiusDisplay = &v
}

// GetRealDiameter returns the RealDiameter field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetRealDiameter() bool {
	if o == nil || o.RealDiameter == nil {
		var ret bool
		return ret
	}
	return *o.RealDiameter
}

// GetRealDiameterOk returns a tuple with the RealDiameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetRealDiameterOk() (*bool, bool) {
	if o == nil || o.RealDiameter == nil {
		return nil, false
	}
	return o.RealDiameter, true
}

// HasRealDiameter returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasRealDiameter() bool {
	if o != nil && o.RealDiameter != nil {
		return true
	}

	return false
}

// SetRealDiameter gets a reference to the given bool and assigns it to the RealDiameter field.
func (o *BTRadialDimensionDisplayData348) SetRealDiameter(v bool) {
	o.RealDiameter = &v
}

// GetWitnessEndPoint0r returns the WitnessEndPoint0r field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint0r() float64 {
	if o == nil || o.WitnessEndPoint0r == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint0r
}

// GetWitnessEndPoint0rOk returns a tuple with the WitnessEndPoint0r field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint0rOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint0r == nil {
		return nil, false
	}
	return o.WitnessEndPoint0r, true
}

// HasWitnessEndPoint0r returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasWitnessEndPoint0r() bool {
	if o != nil && o.WitnessEndPoint0r != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint0r gets a reference to the given float64 and assigns it to the WitnessEndPoint0r field.
func (o *BTRadialDimensionDisplayData348) SetWitnessEndPoint0r(v float64) {
	o.WitnessEndPoint0r = &v
}

// GetWitnessEndPoint0t returns the WitnessEndPoint0t field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint0t() float64 {
	if o == nil || o.WitnessEndPoint0t == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint0t
}

// GetWitnessEndPoint0tOk returns a tuple with the WitnessEndPoint0t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint0tOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint0t == nil {
		return nil, false
	}
	return o.WitnessEndPoint0t, true
}

// HasWitnessEndPoint0t returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasWitnessEndPoint0t() bool {
	if o != nil && o.WitnessEndPoint0t != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint0t gets a reference to the given float64 and assigns it to the WitnessEndPoint0t field.
func (o *BTRadialDimensionDisplayData348) SetWitnessEndPoint0t(v float64) {
	o.WitnessEndPoint0t = &v
}

// GetWitnessEndPoint1r returns the WitnessEndPoint1r field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint1r() float64 {
	if o == nil || o.WitnessEndPoint1r == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint1r
}

// GetWitnessEndPoint1rOk returns a tuple with the WitnessEndPoint1r field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint1rOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint1r == nil {
		return nil, false
	}
	return o.WitnessEndPoint1r, true
}

// HasWitnessEndPoint1r returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasWitnessEndPoint1r() bool {
	if o != nil && o.WitnessEndPoint1r != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint1r gets a reference to the given float64 and assigns it to the WitnessEndPoint1r field.
func (o *BTRadialDimensionDisplayData348) SetWitnessEndPoint1r(v float64) {
	o.WitnessEndPoint1r = &v
}

// GetWitnessEndPoint1t returns the WitnessEndPoint1t field value if set, zero value otherwise.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint1t() float64 {
	if o == nil || o.WitnessEndPoint1t == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint1t
}

// GetWitnessEndPoint1tOk returns a tuple with the WitnessEndPoint1t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRadialDimensionDisplayData348) GetWitnessEndPoint1tOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint1t == nil {
		return nil, false
	}
	return o.WitnessEndPoint1t, true
}

// HasWitnessEndPoint1t returns a boolean if a field has been set.
func (o *BTRadialDimensionDisplayData348) HasWitnessEndPoint1t() bool {
	if o != nil && o.WitnessEndPoint1t != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint1t gets a reference to the given float64 and assigns it to the WitnessEndPoint1t field.
func (o *BTRadialDimensionDisplayData348) SetWitnessEndPoint1t(v float64) {
	o.WitnessEndPoint1t = &v
}

func (o BTRadialDimensionDisplayData348) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CoordinateSystem != nil {
		toSerialize["coordinateSystem"] = o.CoordinateSystem
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.HasMaximumLimit_ != nil {
		toSerialize["hasMaximumLimit"] = o.HasMaximumLimit_
	}
	if o.HasMinimumLimit_ != nil {
		toSerialize["hasMinimumLimit"] = o.HasMinimumLimit_
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsAssociatedWithFlat != nil {
		toSerialize["isAssociatedWithFlat"] = o.IsAssociatedWithFlat
	}
	if o.IsDriven != nil {
		toSerialize["isDriven"] = o.IsDriven
	}
	if o.IsOverDefined != nil {
		toSerialize["isOverDefined"] = o.IsOverDefined
	}
	if o.MaximumLimit != nil {
		toSerialize["maximumLimit"] = o.MaximumLimit
	}
	if o.MinimumLimit != nil {
		toSerialize["minimumLimit"] = o.MinimumLimit
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.PlaneMatrix != nil {
		toSerialize["planeMatrix"] = o.PlaneMatrix
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.PositionR != nil {
		toSerialize["positionR"] = o.PositionR
	}
	if o.PositionT != nil {
		toSerialize["positionT"] = o.PositionT
	}
	if o.RadiusDisplay != nil {
		toSerialize["radiusDisplay"] = o.RadiusDisplay
	}
	if o.RealDiameter != nil {
		toSerialize["realDiameter"] = o.RealDiameter
	}
	if o.WitnessEndPoint0r != nil {
		toSerialize["witnessEndPoint0r"] = o.WitnessEndPoint0r
	}
	if o.WitnessEndPoint0t != nil {
		toSerialize["witnessEndPoint0t"] = o.WitnessEndPoint0t
	}
	if o.WitnessEndPoint1r != nil {
		toSerialize["witnessEndPoint1r"] = o.WitnessEndPoint1r
	}
	if o.WitnessEndPoint1t != nil {
		toSerialize["witnessEndPoint1t"] = o.WitnessEndPoint1t
	}
	return json.Marshal(toSerialize)
}

type NullableBTRadialDimensionDisplayData348 struct {
	value *BTRadialDimensionDisplayData348
	isSet bool
}

func (v NullableBTRadialDimensionDisplayData348) Get() *BTRadialDimensionDisplayData348 {
	return v.value
}

func (v *NullableBTRadialDimensionDisplayData348) Set(val *BTRadialDimensionDisplayData348) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRadialDimensionDisplayData348) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRadialDimensionDisplayData348) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRadialDimensionDisplayData348(val *BTRadialDimensionDisplayData348) *NullableBTRadialDimensionDisplayData348 {
	return &NullableBTRadialDimensionDisplayData348{value: val, isSet: true}
}

func (v NullableBTRadialDimensionDisplayData348) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRadialDimensionDisplayData348) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}