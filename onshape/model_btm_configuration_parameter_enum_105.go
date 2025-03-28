/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMConfigurationParameterEnum105 struct for BTMConfigurationParameterEnum105
type BTMConfigurationParameterEnum105 struct {
	BTMConfigurationParameter819
	BtType *string `json:"btType,omitempty"`
	// Microversion that resulted from the import.
	ImportMicroversion             *string                                  `json:"importMicroversion,omitempty"`
	NodeId                         *string                                  `json:"nodeId,omitempty"`
	EnumOptionIds                  []string                                 `json:"enumOptionIds,omitempty"`
	GeneratedParameterId           *BTTreeNode20                            `json:"generatedParameterId,omitempty"`
	IsCosmetic                     *bool                                    `json:"isCosmetic,omitempty"`
	ParameterId                    *string                                  `json:"parameterId,omitempty"`
	ParameterName                  *string                                  `json:"parameterName,omitempty"`
	ParameterType                  *GBTConfigurationParameterType           `json:"parameterType,omitempty"`
	Valid                          *bool                                    `json:"valid,omitempty"`
	VisibilityCondition            *BTParameterVisibilityCondition177       `json:"visibilityCondition,omitempty"`
	DefaultValue                   *string                                  `json:"defaultValue,omitempty"`
	EnumName                       *string                                  `json:"enumName,omitempty"`
	EnumOptionVisibilityConditions *BTEnumOptionVisibilityConditionList2936 `json:"enumOptionVisibilityConditions,omitempty"`
	Namespace                      *string                                  `json:"namespace,omitempty"`
	OptionIds                      []string                                 `json:"optionIds,omitempty"`
	Options                        []BTMEnumOption592                       `json:"options,omitempty"`
}

// NewBTMConfigurationParameterEnum105 instantiates a new BTMConfigurationParameterEnum105 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationParameterEnum105() *BTMConfigurationParameterEnum105 {
	this := BTMConfigurationParameterEnum105{}
	return &this
}

// NewBTMConfigurationParameterEnum105WithDefaults instantiates a new BTMConfigurationParameterEnum105 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationParameterEnum105WithDefaults() *BTMConfigurationParameterEnum105 {
	this := BTMConfigurationParameterEnum105{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfigurationParameterEnum105) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfigurationParameterEnum105) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfigurationParameterEnum105) SetNodeId(v string) {
	o.NodeId = &v
}

// GetEnumOptionIds returns the EnumOptionIds field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetEnumOptionIds() []string {
	if o == nil || o.EnumOptionIds == nil {
		var ret []string
		return ret
	}
	return o.EnumOptionIds
}

// GetEnumOptionIdsOk returns a tuple with the EnumOptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetEnumOptionIdsOk() ([]string, bool) {
	if o == nil || o.EnumOptionIds == nil {
		return nil, false
	}
	return o.EnumOptionIds, true
}

// HasEnumOptionIds returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasEnumOptionIds() bool {
	if o != nil && o.EnumOptionIds != nil {
		return true
	}

	return false
}

// SetEnumOptionIds gets a reference to the given []string and assigns it to the EnumOptionIds field.
func (o *BTMConfigurationParameterEnum105) SetEnumOptionIds(v []string) {
	o.EnumOptionIds = v
}

// GetGeneratedParameterId returns the GeneratedParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetGeneratedParameterId() BTTreeNode20 {
	if o == nil || o.GeneratedParameterId == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.GeneratedParameterId
}

// GetGeneratedParameterIdOk returns a tuple with the GeneratedParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetGeneratedParameterIdOk() (*BTTreeNode20, bool) {
	if o == nil || o.GeneratedParameterId == nil {
		return nil, false
	}
	return o.GeneratedParameterId, true
}

// HasGeneratedParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasGeneratedParameterId() bool {
	if o != nil && o.GeneratedParameterId != nil {
		return true
	}

	return false
}

// SetGeneratedParameterId gets a reference to the given BTTreeNode20 and assigns it to the GeneratedParameterId field.
func (o *BTMConfigurationParameterEnum105) SetGeneratedParameterId(v BTTreeNode20) {
	o.GeneratedParameterId = &v
}

// GetIsCosmetic returns the IsCosmetic field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetIsCosmetic() bool {
	if o == nil || o.IsCosmetic == nil {
		var ret bool
		return ret
	}
	return *o.IsCosmetic
}

// GetIsCosmeticOk returns a tuple with the IsCosmetic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetIsCosmeticOk() (*bool, bool) {
	if o == nil || o.IsCosmetic == nil {
		return nil, false
	}
	return o.IsCosmetic, true
}

// HasIsCosmetic returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasIsCosmetic() bool {
	if o != nil && o.IsCosmetic != nil {
		return true
	}

	return false
}

// SetIsCosmetic gets a reference to the given bool and assigns it to the IsCosmetic field.
func (o *BTMConfigurationParameterEnum105) SetIsCosmetic(v bool) {
	o.IsCosmetic = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMConfigurationParameterEnum105) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTMConfigurationParameterEnum105) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetParameterType() GBTConfigurationParameterType {
	if o == nil || o.ParameterType == nil {
		var ret GBTConfigurationParameterType
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetParameterTypeOk() (*GBTConfigurationParameterType, bool) {
	if o == nil || o.ParameterType == nil {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasParameterType() bool {
	if o != nil && o.ParameterType != nil {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given GBTConfigurationParameterType and assigns it to the ParameterType field.
func (o *BTMConfigurationParameterEnum105) SetParameterType(v GBTConfigurationParameterType) {
	o.ParameterType = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTMConfigurationParameterEnum105) SetValid(v bool) {
	o.Valid = &v
}

// GetVisibilityCondition returns the VisibilityCondition field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetVisibilityCondition() BTParameterVisibilityCondition177 {
	if o == nil || o.VisibilityCondition == nil {
		var ret BTParameterVisibilityCondition177
		return ret
	}
	return *o.VisibilityCondition
}

// GetVisibilityConditionOk returns a tuple with the VisibilityCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetVisibilityConditionOk() (*BTParameterVisibilityCondition177, bool) {
	if o == nil || o.VisibilityCondition == nil {
		return nil, false
	}
	return o.VisibilityCondition, true
}

// HasVisibilityCondition returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasVisibilityCondition() bool {
	if o != nil && o.VisibilityCondition != nil {
		return true
	}

	return false
}

// SetVisibilityCondition gets a reference to the given BTParameterVisibilityCondition177 and assigns it to the VisibilityCondition field.
func (o *BTMConfigurationParameterEnum105) SetVisibilityCondition(v BTParameterVisibilityCondition177) {
	o.VisibilityCondition = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *BTMConfigurationParameterEnum105) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetEnumName returns the EnumName field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetEnumName() string {
	if o == nil || o.EnumName == nil {
		var ret string
		return ret
	}
	return *o.EnumName
}

// GetEnumNameOk returns a tuple with the EnumName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetEnumNameOk() (*string, bool) {
	if o == nil || o.EnumName == nil {
		return nil, false
	}
	return o.EnumName, true
}

// HasEnumName returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasEnumName() bool {
	if o != nil && o.EnumName != nil {
		return true
	}

	return false
}

// SetEnumName gets a reference to the given string and assigns it to the EnumName field.
func (o *BTMConfigurationParameterEnum105) SetEnumName(v string) {
	o.EnumName = &v
}

// GetEnumOptionVisibilityConditions returns the EnumOptionVisibilityConditions field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetEnumOptionVisibilityConditions() BTEnumOptionVisibilityConditionList2936 {
	if o == nil || o.EnumOptionVisibilityConditions == nil {
		var ret BTEnumOptionVisibilityConditionList2936
		return ret
	}
	return *o.EnumOptionVisibilityConditions
}

// GetEnumOptionVisibilityConditionsOk returns a tuple with the EnumOptionVisibilityConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetEnumOptionVisibilityConditionsOk() (*BTEnumOptionVisibilityConditionList2936, bool) {
	if o == nil || o.EnumOptionVisibilityConditions == nil {
		return nil, false
	}
	return o.EnumOptionVisibilityConditions, true
}

// HasEnumOptionVisibilityConditions returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasEnumOptionVisibilityConditions() bool {
	if o != nil && o.EnumOptionVisibilityConditions != nil {
		return true
	}

	return false
}

// SetEnumOptionVisibilityConditions gets a reference to the given BTEnumOptionVisibilityConditionList2936 and assigns it to the EnumOptionVisibilityConditions field.
func (o *BTMConfigurationParameterEnum105) SetEnumOptionVisibilityConditions(v BTEnumOptionVisibilityConditionList2936) {
	o.EnumOptionVisibilityConditions = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMConfigurationParameterEnum105) SetNamespace(v string) {
	o.Namespace = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetOptionIds() []string {
	if o == nil || o.OptionIds == nil {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetOptionIdsOk() ([]string, bool) {
	if o == nil || o.OptionIds == nil {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasOptionIds() bool {
	if o != nil && o.OptionIds != nil {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *BTMConfigurationParameterEnum105) SetOptionIds(v []string) {
	o.OptionIds = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *BTMConfigurationParameterEnum105) GetOptions() []BTMEnumOption592 {
	if o == nil || o.Options == nil {
		var ret []BTMEnumOption592
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterEnum105) GetOptionsOk() ([]BTMEnumOption592, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *BTMConfigurationParameterEnum105) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []BTMEnumOption592 and assigns it to the Options field.
func (o *BTMConfigurationParameterEnum105) SetOptions(v []BTMEnumOption592) {
	o.Options = v
}

func (o BTMConfigurationParameterEnum105) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMConfigurationParameter819, errBTMConfigurationParameter819 := json.Marshal(o.BTMConfigurationParameter819)
	if errBTMConfigurationParameter819 != nil {
		return []byte{}, errBTMConfigurationParameter819
	}
	errBTMConfigurationParameter819 = json.Unmarshal([]byte(serializedBTMConfigurationParameter819), &toSerialize)
	if errBTMConfigurationParameter819 != nil {
		return []byte{}, errBTMConfigurationParameter819
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.EnumOptionIds != nil {
		toSerialize["enumOptionIds"] = o.EnumOptionIds
	}
	if o.GeneratedParameterId != nil {
		toSerialize["generatedParameterId"] = o.GeneratedParameterId
	}
	if o.IsCosmetic != nil {
		toSerialize["isCosmetic"] = o.IsCosmetic
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.ParameterType != nil {
		toSerialize["parameterType"] = o.ParameterType
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.VisibilityCondition != nil {
		toSerialize["visibilityCondition"] = o.VisibilityCondition
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.EnumName != nil {
		toSerialize["enumName"] = o.EnumName
	}
	if o.EnumOptionVisibilityConditions != nil {
		toSerialize["enumOptionVisibilityConditions"] = o.EnumOptionVisibilityConditions
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.OptionIds != nil {
		toSerialize["optionIds"] = o.OptionIds
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfigurationParameterEnum105 struct {
	value *BTMConfigurationParameterEnum105
	isSet bool
}

func (v NullableBTMConfigurationParameterEnum105) Get() *BTMConfigurationParameterEnum105 {
	return v.value
}

func (v *NullableBTMConfigurationParameterEnum105) Set(val *BTMConfigurationParameterEnum105) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationParameterEnum105) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationParameterEnum105) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationParameterEnum105(val *BTMConfigurationParameterEnum105) *NullableBTMConfigurationParameterEnum105 {
	return &NullableBTMConfigurationParameterEnum105{value: val, isSet: true}
}

func (v NullableBTMConfigurationParameterEnum105) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationParameterEnum105) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
