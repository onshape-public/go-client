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

// BTBExportAdvancedParams Advanced element export options.
type BTBExportAdvancedParams struct {
	AssemblyExportParams *BTBAssemblyExportParams `json:"assemblyExportParams,omitempty"`
	// URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details.
	Configuration *string `json:"configuration,omitempty"`
	// An array of element ids for multi-element export.
	ElementIds []string `json:"elementIds,omitempty"`
	// Set to `true` to evaluate the export rule for the given `formatName` and to include an `exportRuleFileName` value in the response.
	EvaluateExportRule *bool `json:"evaluateExportRule,omitempty"`
	// For multiple elements export, use `true` if export rule shouldn't be applied for all elements.
	IgnoreExportRulesForContents *bool `json:"ignoreExportRulesForContents,omitempty"`
	// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
	LinkDocumentId *string `json:"linkDocumentId,omitempty"`
	// The id of the workspace through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
	LinkDocumentWorkspaceId *string `json:"linkDocumentWorkspaceId,omitempty"`
	// IDs of the parts to retrieve. Use comma-separated IDs for multiple parts (example: partIds=JHK,JHD).
	PartIds           *string                  `json:"partIds,omitempty"`
	PartsExportFilter *BTPartsExportFilter4308 `json:"partsExportFilter,omitempty"`
}

// NewBTBExportAdvancedParams instantiates a new BTBExportAdvancedParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBExportAdvancedParams() *BTBExportAdvancedParams {
	this := BTBExportAdvancedParams{}
	var evaluateExportRule bool = false
	this.EvaluateExportRule = &evaluateExportRule
	var ignoreExportRulesForContents bool = false
	this.IgnoreExportRulesForContents = &ignoreExportRulesForContents
	return &this
}

// NewBTBExportAdvancedParamsWithDefaults instantiates a new BTBExportAdvancedParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBExportAdvancedParamsWithDefaults() *BTBExportAdvancedParams {
	this := BTBExportAdvancedParams{}
	var evaluateExportRule bool = false
	this.EvaluateExportRule = &evaluateExportRule
	var ignoreExportRulesForContents bool = false
	this.IgnoreExportRulesForContents = &ignoreExportRulesForContents
	return &this
}

// GetAssemblyExportParams returns the AssemblyExportParams field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetAssemblyExportParams() BTBAssemblyExportParams {
	if o == nil || o.AssemblyExportParams == nil {
		var ret BTBAssemblyExportParams
		return ret
	}
	return *o.AssemblyExportParams
}

// GetAssemblyExportParamsOk returns a tuple with the AssemblyExportParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetAssemblyExportParamsOk() (*BTBAssemblyExportParams, bool) {
	if o == nil || o.AssemblyExportParams == nil {
		return nil, false
	}
	return o.AssemblyExportParams, true
}

// HasAssemblyExportParams returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasAssemblyExportParams() bool {
	if o != nil && o.AssemblyExportParams != nil {
		return true
	}

	return false
}

// SetAssemblyExportParams gets a reference to the given BTBAssemblyExportParams and assigns it to the AssemblyExportParams field.
func (o *BTBExportAdvancedParams) SetAssemblyExportParams(v BTBAssemblyExportParams) {
	o.AssemblyExportParams = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTBExportAdvancedParams) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetElementIds returns the ElementIds field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetElementIds() []string {
	if o == nil || o.ElementIds == nil {
		var ret []string
		return ret
	}
	return o.ElementIds
}

// GetElementIdsOk returns a tuple with the ElementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetElementIdsOk() ([]string, bool) {
	if o == nil || o.ElementIds == nil {
		return nil, false
	}
	return o.ElementIds, true
}

// HasElementIds returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasElementIds() bool {
	if o != nil && o.ElementIds != nil {
		return true
	}

	return false
}

// SetElementIds gets a reference to the given []string and assigns it to the ElementIds field.
func (o *BTBExportAdvancedParams) SetElementIds(v []string) {
	o.ElementIds = v
}

// GetEvaluateExportRule returns the EvaluateExportRule field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetEvaluateExportRule() bool {
	if o == nil || o.EvaluateExportRule == nil {
		var ret bool
		return ret
	}
	return *o.EvaluateExportRule
}

// GetEvaluateExportRuleOk returns a tuple with the EvaluateExportRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetEvaluateExportRuleOk() (*bool, bool) {
	if o == nil || o.EvaluateExportRule == nil {
		return nil, false
	}
	return o.EvaluateExportRule, true
}

// HasEvaluateExportRule returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasEvaluateExportRule() bool {
	if o != nil && o.EvaluateExportRule != nil {
		return true
	}

	return false
}

// SetEvaluateExportRule gets a reference to the given bool and assigns it to the EvaluateExportRule field.
func (o *BTBExportAdvancedParams) SetEvaluateExportRule(v bool) {
	o.EvaluateExportRule = &v
}

// GetIgnoreExportRulesForContents returns the IgnoreExportRulesForContents field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetIgnoreExportRulesForContents() bool {
	if o == nil || o.IgnoreExportRulesForContents == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreExportRulesForContents
}

// GetIgnoreExportRulesForContentsOk returns a tuple with the IgnoreExportRulesForContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetIgnoreExportRulesForContentsOk() (*bool, bool) {
	if o == nil || o.IgnoreExportRulesForContents == nil {
		return nil, false
	}
	return o.IgnoreExportRulesForContents, true
}

// HasIgnoreExportRulesForContents returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasIgnoreExportRulesForContents() bool {
	if o != nil && o.IgnoreExportRulesForContents != nil {
		return true
	}

	return false
}

// SetIgnoreExportRulesForContents gets a reference to the given bool and assigns it to the IgnoreExportRulesForContents field.
func (o *BTBExportAdvancedParams) SetIgnoreExportRulesForContents(v bool) {
	o.IgnoreExportRulesForContents = &v
}

// GetLinkDocumentId returns the LinkDocumentId field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetLinkDocumentId() string {
	if o == nil || o.LinkDocumentId == nil {
		var ret string
		return ret
	}
	return *o.LinkDocumentId
}

// GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetLinkDocumentIdOk() (*string, bool) {
	if o == nil || o.LinkDocumentId == nil {
		return nil, false
	}
	return o.LinkDocumentId, true
}

// HasLinkDocumentId returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasLinkDocumentId() bool {
	if o != nil && o.LinkDocumentId != nil {
		return true
	}

	return false
}

// SetLinkDocumentId gets a reference to the given string and assigns it to the LinkDocumentId field.
func (o *BTBExportAdvancedParams) SetLinkDocumentId(v string) {
	o.LinkDocumentId = &v
}

// GetLinkDocumentWorkspaceId returns the LinkDocumentWorkspaceId field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetLinkDocumentWorkspaceId() string {
	if o == nil || o.LinkDocumentWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.LinkDocumentWorkspaceId
}

// GetLinkDocumentWorkspaceIdOk returns a tuple with the LinkDocumentWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetLinkDocumentWorkspaceIdOk() (*string, bool) {
	if o == nil || o.LinkDocumentWorkspaceId == nil {
		return nil, false
	}
	return o.LinkDocumentWorkspaceId, true
}

// HasLinkDocumentWorkspaceId returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasLinkDocumentWorkspaceId() bool {
	if o != nil && o.LinkDocumentWorkspaceId != nil {
		return true
	}

	return false
}

// SetLinkDocumentWorkspaceId gets a reference to the given string and assigns it to the LinkDocumentWorkspaceId field.
func (o *BTBExportAdvancedParams) SetLinkDocumentWorkspaceId(v string) {
	o.LinkDocumentWorkspaceId = &v
}

// GetPartIds returns the PartIds field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetPartIds() string {
	if o == nil || o.PartIds == nil {
		var ret string
		return ret
	}
	return *o.PartIds
}

// GetPartIdsOk returns a tuple with the PartIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetPartIdsOk() (*string, bool) {
	if o == nil || o.PartIds == nil {
		return nil, false
	}
	return o.PartIds, true
}

// HasPartIds returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasPartIds() bool {
	if o != nil && o.PartIds != nil {
		return true
	}

	return false
}

// SetPartIds gets a reference to the given string and assigns it to the PartIds field.
func (o *BTBExportAdvancedParams) SetPartIds(v string) {
	o.PartIds = &v
}

// GetPartsExportFilter returns the PartsExportFilter field value if set, zero value otherwise.
func (o *BTBExportAdvancedParams) GetPartsExportFilter() BTPartsExportFilter4308 {
	if o == nil || o.PartsExportFilter == nil {
		var ret BTPartsExportFilter4308
		return ret
	}
	return *o.PartsExportFilter
}

// GetPartsExportFilterOk returns a tuple with the PartsExportFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBExportAdvancedParams) GetPartsExportFilterOk() (*BTPartsExportFilter4308, bool) {
	if o == nil || o.PartsExportFilter == nil {
		return nil, false
	}
	return o.PartsExportFilter, true
}

// HasPartsExportFilter returns a boolean if a field has been set.
func (o *BTBExportAdvancedParams) HasPartsExportFilter() bool {
	if o != nil && o.PartsExportFilter != nil {
		return true
	}

	return false
}

// SetPartsExportFilter gets a reference to the given BTPartsExportFilter4308 and assigns it to the PartsExportFilter field.
func (o *BTBExportAdvancedParams) SetPartsExportFilter(v BTPartsExportFilter4308) {
	o.PartsExportFilter = &v
}

func (o BTBExportAdvancedParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyExportParams != nil {
		toSerialize["assemblyExportParams"] = o.AssemblyExportParams
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ElementIds != nil {
		toSerialize["elementIds"] = o.ElementIds
	}
	if o.EvaluateExportRule != nil {
		toSerialize["evaluateExportRule"] = o.EvaluateExportRule
	}
	if o.IgnoreExportRulesForContents != nil {
		toSerialize["ignoreExportRulesForContents"] = o.IgnoreExportRulesForContents
	}
	if o.LinkDocumentId != nil {
		toSerialize["linkDocumentId"] = o.LinkDocumentId
	}
	if o.LinkDocumentWorkspaceId != nil {
		toSerialize["linkDocumentWorkspaceId"] = o.LinkDocumentWorkspaceId
	}
	if o.PartIds != nil {
		toSerialize["partIds"] = o.PartIds
	}
	if o.PartsExportFilter != nil {
		toSerialize["partsExportFilter"] = o.PartsExportFilter
	}
	return json.Marshal(toSerialize)
}

type NullableBTBExportAdvancedParams struct {
	value *BTBExportAdvancedParams
	isSet bool
}

func (v NullableBTBExportAdvancedParams) Get() *BTBExportAdvancedParams {
	return v.value
}

func (v *NullableBTBExportAdvancedParams) Set(val *BTBExportAdvancedParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBExportAdvancedParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBExportAdvancedParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBExportAdvancedParams(val *BTBExportAdvancedParams) *NullableBTBExportAdvancedParams {
	return &NullableBTBExportAdvancedParams{value: val, isSet: true}
}

func (v NullableBTBExportAdvancedParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBExportAdvancedParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
