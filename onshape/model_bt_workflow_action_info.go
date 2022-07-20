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

// BTWorkflowActionInfo struct for BTWorkflowActionInfo
type BTWorkflowActionInfo struct {
	Action             *string  `json:"action,omitempty"`
	AllowIfApprovers   *bool    `json:"allowIfApprovers,omitempty"`
	AllowIfNoApprovers *bool    `json:"allowIfNoApprovers,omitempty"`
	AlwaysAllow        *bool    `json:"alwaysAllow,omitempty"`
	IsAdminOverride    *bool    `json:"isAdminOverride,omitempty"`
	IsApproverAction   *bool    `json:"isApproverAction,omitempty"`
	Label              *string  `json:"label,omitempty"`
	RequiredProperties []string `json:"requiredProperties,omitempty"`
	Tooltip            *string  `json:"tooltip,omitempty"`
	UiHint             *string  `json:"uiHint,omitempty"`
}

// NewBTWorkflowActionInfo instantiates a new BTWorkflowActionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowActionInfo() *BTWorkflowActionInfo {
	this := BTWorkflowActionInfo{}
	return &this
}

// NewBTWorkflowActionInfoWithDefaults instantiates a new BTWorkflowActionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowActionInfoWithDefaults() *BTWorkflowActionInfo {
	this := BTWorkflowActionInfo{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BTWorkflowActionInfo) SetAction(v string) {
	o.Action = &v
}

// GetAllowIfApprovers returns the AllowIfApprovers field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetAllowIfApprovers() bool {
	if o == nil || o.AllowIfApprovers == nil {
		var ret bool
		return ret
	}
	return *o.AllowIfApprovers
}

// GetAllowIfApproversOk returns a tuple with the AllowIfApprovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetAllowIfApproversOk() (*bool, bool) {
	if o == nil || o.AllowIfApprovers == nil {
		return nil, false
	}
	return o.AllowIfApprovers, true
}

// HasAllowIfApprovers returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasAllowIfApprovers() bool {
	if o != nil && o.AllowIfApprovers != nil {
		return true
	}

	return false
}

// SetAllowIfApprovers gets a reference to the given bool and assigns it to the AllowIfApprovers field.
func (o *BTWorkflowActionInfo) SetAllowIfApprovers(v bool) {
	o.AllowIfApprovers = &v
}

// GetAllowIfNoApprovers returns the AllowIfNoApprovers field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetAllowIfNoApprovers() bool {
	if o == nil || o.AllowIfNoApprovers == nil {
		var ret bool
		return ret
	}
	return *o.AllowIfNoApprovers
}

// GetAllowIfNoApproversOk returns a tuple with the AllowIfNoApprovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetAllowIfNoApproversOk() (*bool, bool) {
	if o == nil || o.AllowIfNoApprovers == nil {
		return nil, false
	}
	return o.AllowIfNoApprovers, true
}

// HasAllowIfNoApprovers returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasAllowIfNoApprovers() bool {
	if o != nil && o.AllowIfNoApprovers != nil {
		return true
	}

	return false
}

// SetAllowIfNoApprovers gets a reference to the given bool and assigns it to the AllowIfNoApprovers field.
func (o *BTWorkflowActionInfo) SetAllowIfNoApprovers(v bool) {
	o.AllowIfNoApprovers = &v
}

// GetAlwaysAllow returns the AlwaysAllow field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetAlwaysAllow() bool {
	if o == nil || o.AlwaysAllow == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysAllow
}

// GetAlwaysAllowOk returns a tuple with the AlwaysAllow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetAlwaysAllowOk() (*bool, bool) {
	if o == nil || o.AlwaysAllow == nil {
		return nil, false
	}
	return o.AlwaysAllow, true
}

// HasAlwaysAllow returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasAlwaysAllow() bool {
	if o != nil && o.AlwaysAllow != nil {
		return true
	}

	return false
}

// SetAlwaysAllow gets a reference to the given bool and assigns it to the AlwaysAllow field.
func (o *BTWorkflowActionInfo) SetAlwaysAllow(v bool) {
	o.AlwaysAllow = &v
}

// GetIsAdminOverride returns the IsAdminOverride field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetIsAdminOverride() bool {
	if o == nil || o.IsAdminOverride == nil {
		var ret bool
		return ret
	}
	return *o.IsAdminOverride
}

// GetIsAdminOverrideOk returns a tuple with the IsAdminOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetIsAdminOverrideOk() (*bool, bool) {
	if o == nil || o.IsAdminOverride == nil {
		return nil, false
	}
	return o.IsAdminOverride, true
}

// HasIsAdminOverride returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasIsAdminOverride() bool {
	if o != nil && o.IsAdminOverride != nil {
		return true
	}

	return false
}

// SetIsAdminOverride gets a reference to the given bool and assigns it to the IsAdminOverride field.
func (o *BTWorkflowActionInfo) SetIsAdminOverride(v bool) {
	o.IsAdminOverride = &v
}

// GetIsApproverAction returns the IsApproverAction field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetIsApproverAction() bool {
	if o == nil || o.IsApproverAction == nil {
		var ret bool
		return ret
	}
	return *o.IsApproverAction
}

// GetIsApproverActionOk returns a tuple with the IsApproverAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetIsApproverActionOk() (*bool, bool) {
	if o == nil || o.IsApproverAction == nil {
		return nil, false
	}
	return o.IsApproverAction, true
}

// HasIsApproverAction returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasIsApproverAction() bool {
	if o != nil && o.IsApproverAction != nil {
		return true
	}

	return false
}

// SetIsApproverAction gets a reference to the given bool and assigns it to the IsApproverAction field.
func (o *BTWorkflowActionInfo) SetIsApproverAction(v bool) {
	o.IsApproverAction = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BTWorkflowActionInfo) SetLabel(v string) {
	o.Label = &v
}

// GetRequiredProperties returns the RequiredProperties field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetRequiredProperties() []string {
	if o == nil || o.RequiredProperties == nil {
		var ret []string
		return ret
	}
	return o.RequiredProperties
}

// GetRequiredPropertiesOk returns a tuple with the RequiredProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetRequiredPropertiesOk() ([]string, bool) {
	if o == nil || o.RequiredProperties == nil {
		return nil, false
	}
	return o.RequiredProperties, true
}

// HasRequiredProperties returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasRequiredProperties() bool {
	if o != nil && o.RequiredProperties != nil {
		return true
	}

	return false
}

// SetRequiredProperties gets a reference to the given []string and assigns it to the RequiredProperties field.
func (o *BTWorkflowActionInfo) SetRequiredProperties(v []string) {
	o.RequiredProperties = v
}

// GetTooltip returns the Tooltip field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetTooltip() string {
	if o == nil || o.Tooltip == nil {
		var ret string
		return ret
	}
	return *o.Tooltip
}

// GetTooltipOk returns a tuple with the Tooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetTooltipOk() (*string, bool) {
	if o == nil || o.Tooltip == nil {
		return nil, false
	}
	return o.Tooltip, true
}

// HasTooltip returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasTooltip() bool {
	if o != nil && o.Tooltip != nil {
		return true
	}

	return false
}

// SetTooltip gets a reference to the given string and assigns it to the Tooltip field.
func (o *BTWorkflowActionInfo) SetTooltip(v string) {
	o.Tooltip = &v
}

// GetUiHint returns the UiHint field value if set, zero value otherwise.
func (o *BTWorkflowActionInfo) GetUiHint() string {
	if o == nil || o.UiHint == nil {
		var ret string
		return ret
	}
	return *o.UiHint
}

// GetUiHintOk returns a tuple with the UiHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowActionInfo) GetUiHintOk() (*string, bool) {
	if o == nil || o.UiHint == nil {
		return nil, false
	}
	return o.UiHint, true
}

// HasUiHint returns a boolean if a field has been set.
func (o *BTWorkflowActionInfo) HasUiHint() bool {
	if o != nil && o.UiHint != nil {
		return true
	}

	return false
}

// SetUiHint gets a reference to the given string and assigns it to the UiHint field.
func (o *BTWorkflowActionInfo) SetUiHint(v string) {
	o.UiHint = &v
}

func (o BTWorkflowActionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.AllowIfApprovers != nil {
		toSerialize["allowIfApprovers"] = o.AllowIfApprovers
	}
	if o.AllowIfNoApprovers != nil {
		toSerialize["allowIfNoApprovers"] = o.AllowIfNoApprovers
	}
	if o.AlwaysAllow != nil {
		toSerialize["alwaysAllow"] = o.AlwaysAllow
	}
	if o.IsAdminOverride != nil {
		toSerialize["isAdminOverride"] = o.IsAdminOverride
	}
	if o.IsApproverAction != nil {
		toSerialize["isApproverAction"] = o.IsApproverAction
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.RequiredProperties != nil {
		toSerialize["requiredProperties"] = o.RequiredProperties
	}
	if o.Tooltip != nil {
		toSerialize["tooltip"] = o.Tooltip
	}
	if o.UiHint != nil {
		toSerialize["uiHint"] = o.UiHint
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowActionInfo struct {
	value *BTWorkflowActionInfo
	isSet bool
}

func (v NullableBTWorkflowActionInfo) Get() *BTWorkflowActionInfo {
	return v.value
}

func (v *NullableBTWorkflowActionInfo) Set(val *BTWorkflowActionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowActionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowActionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowActionInfo(val *BTWorkflowActionInfo) *NullableBTWorkflowActionInfo {
	return &NullableBTWorkflowActionInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowActionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowActionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}