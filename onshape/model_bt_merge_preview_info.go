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

// BTMergePreviewInfo struct for BTMergePreviewInfo
type BTMergePreviewInfo struct {
	BranchPointMicroversionId *string              `json:"branchPointMicroversionId,omitempty"`
	BranchPointVersionId      *string              `json:"branchPointVersionId,omitempty"`
	BranchPointWorkspaceId    *string              `json:"branchPointWorkspaceId,omitempty"`
	Changes                   []BTElementMergeInfo `json:"changes,omitempty"`
	IsBranchPointAtStart      *bool                `json:"isBranchPointAtStart,omitempty"`
	SourceMicroversionId      *string              `json:"sourceMicroversionId,omitempty"`
	TargetMicroversionId      *string              `json:"targetMicroversionId,omitempty"`
	UpgradeInfo               *BTMergeUpgradeInfo  `json:"upgradeInfo,omitempty"`
}

// NewBTMergePreviewInfo instantiates a new BTMergePreviewInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMergePreviewInfo() *BTMergePreviewInfo {
	this := BTMergePreviewInfo{}
	return &this
}

// NewBTMergePreviewInfoWithDefaults instantiates a new BTMergePreviewInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMergePreviewInfoWithDefaults() *BTMergePreviewInfo {
	this := BTMergePreviewInfo{}
	return &this
}

// GetBranchPointMicroversionId returns the BranchPointMicroversionId field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetBranchPointMicroversionId() string {
	if o == nil || o.BranchPointMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.BranchPointMicroversionId
}

// GetBranchPointMicroversionIdOk returns a tuple with the BranchPointMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetBranchPointMicroversionIdOk() (*string, bool) {
	if o == nil || o.BranchPointMicroversionId == nil {
		return nil, false
	}
	return o.BranchPointMicroversionId, true
}

// HasBranchPointMicroversionId returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasBranchPointMicroversionId() bool {
	if o != nil && o.BranchPointMicroversionId != nil {
		return true
	}

	return false
}

// SetBranchPointMicroversionId gets a reference to the given string and assigns it to the BranchPointMicroversionId field.
func (o *BTMergePreviewInfo) SetBranchPointMicroversionId(v string) {
	o.BranchPointMicroversionId = &v
}

// GetBranchPointVersionId returns the BranchPointVersionId field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetBranchPointVersionId() string {
	if o == nil || o.BranchPointVersionId == nil {
		var ret string
		return ret
	}
	return *o.BranchPointVersionId
}

// GetBranchPointVersionIdOk returns a tuple with the BranchPointVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetBranchPointVersionIdOk() (*string, bool) {
	if o == nil || o.BranchPointVersionId == nil {
		return nil, false
	}
	return o.BranchPointVersionId, true
}

// HasBranchPointVersionId returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasBranchPointVersionId() bool {
	if o != nil && o.BranchPointVersionId != nil {
		return true
	}

	return false
}

// SetBranchPointVersionId gets a reference to the given string and assigns it to the BranchPointVersionId field.
func (o *BTMergePreviewInfo) SetBranchPointVersionId(v string) {
	o.BranchPointVersionId = &v
}

// GetBranchPointWorkspaceId returns the BranchPointWorkspaceId field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetBranchPointWorkspaceId() string {
	if o == nil || o.BranchPointWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.BranchPointWorkspaceId
}

// GetBranchPointWorkspaceIdOk returns a tuple with the BranchPointWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetBranchPointWorkspaceIdOk() (*string, bool) {
	if o == nil || o.BranchPointWorkspaceId == nil {
		return nil, false
	}
	return o.BranchPointWorkspaceId, true
}

// HasBranchPointWorkspaceId returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasBranchPointWorkspaceId() bool {
	if o != nil && o.BranchPointWorkspaceId != nil {
		return true
	}

	return false
}

// SetBranchPointWorkspaceId gets a reference to the given string and assigns it to the BranchPointWorkspaceId field.
func (o *BTMergePreviewInfo) SetBranchPointWorkspaceId(v string) {
	o.BranchPointWorkspaceId = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetChanges() []BTElementMergeInfo {
	if o == nil || o.Changes == nil {
		var ret []BTElementMergeInfo
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetChangesOk() ([]BTElementMergeInfo, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []BTElementMergeInfo and assigns it to the Changes field.
func (o *BTMergePreviewInfo) SetChanges(v []BTElementMergeInfo) {
	o.Changes = v
}

// GetIsBranchPointAtStart returns the IsBranchPointAtStart field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetIsBranchPointAtStart() bool {
	if o == nil || o.IsBranchPointAtStart == nil {
		var ret bool
		return ret
	}
	return *o.IsBranchPointAtStart
}

// GetIsBranchPointAtStartOk returns a tuple with the IsBranchPointAtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetIsBranchPointAtStartOk() (*bool, bool) {
	if o == nil || o.IsBranchPointAtStart == nil {
		return nil, false
	}
	return o.IsBranchPointAtStart, true
}

// HasIsBranchPointAtStart returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasIsBranchPointAtStart() bool {
	if o != nil && o.IsBranchPointAtStart != nil {
		return true
	}

	return false
}

// SetIsBranchPointAtStart gets a reference to the given bool and assigns it to the IsBranchPointAtStart field.
func (o *BTMergePreviewInfo) SetIsBranchPointAtStart(v bool) {
	o.IsBranchPointAtStart = &v
}

// GetSourceMicroversionId returns the SourceMicroversionId field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetSourceMicroversionId() string {
	if o == nil || o.SourceMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversionId
}

// GetSourceMicroversionIdOk returns a tuple with the SourceMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetSourceMicroversionIdOk() (*string, bool) {
	if o == nil || o.SourceMicroversionId == nil {
		return nil, false
	}
	return o.SourceMicroversionId, true
}

// HasSourceMicroversionId returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasSourceMicroversionId() bool {
	if o != nil && o.SourceMicroversionId != nil {
		return true
	}

	return false
}

// SetSourceMicroversionId gets a reference to the given string and assigns it to the SourceMicroversionId field.
func (o *BTMergePreviewInfo) SetSourceMicroversionId(v string) {
	o.SourceMicroversionId = &v
}

// GetTargetMicroversionId returns the TargetMicroversionId field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetTargetMicroversionId() string {
	if o == nil || o.TargetMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.TargetMicroversionId
}

// GetTargetMicroversionIdOk returns a tuple with the TargetMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetTargetMicroversionIdOk() (*string, bool) {
	if o == nil || o.TargetMicroversionId == nil {
		return nil, false
	}
	return o.TargetMicroversionId, true
}

// HasTargetMicroversionId returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasTargetMicroversionId() bool {
	if o != nil && o.TargetMicroversionId != nil {
		return true
	}

	return false
}

// SetTargetMicroversionId gets a reference to the given string and assigns it to the TargetMicroversionId field.
func (o *BTMergePreviewInfo) SetTargetMicroversionId(v string) {
	o.TargetMicroversionId = &v
}

// GetUpgradeInfo returns the UpgradeInfo field value if set, zero value otherwise.
func (o *BTMergePreviewInfo) GetUpgradeInfo() BTMergeUpgradeInfo {
	if o == nil || o.UpgradeInfo == nil {
		var ret BTMergeUpgradeInfo
		return ret
	}
	return *o.UpgradeInfo
}

// GetUpgradeInfoOk returns a tuple with the UpgradeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMergePreviewInfo) GetUpgradeInfoOk() (*BTMergeUpgradeInfo, bool) {
	if o == nil || o.UpgradeInfo == nil {
		return nil, false
	}
	return o.UpgradeInfo, true
}

// HasUpgradeInfo returns a boolean if a field has been set.
func (o *BTMergePreviewInfo) HasUpgradeInfo() bool {
	if o != nil && o.UpgradeInfo != nil {
		return true
	}

	return false
}

// SetUpgradeInfo gets a reference to the given BTMergeUpgradeInfo and assigns it to the UpgradeInfo field.
func (o *BTMergePreviewInfo) SetUpgradeInfo(v BTMergeUpgradeInfo) {
	o.UpgradeInfo = &v
}

func (o BTMergePreviewInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BranchPointMicroversionId != nil {
		toSerialize["branchPointMicroversionId"] = o.BranchPointMicroversionId
	}
	if o.BranchPointVersionId != nil {
		toSerialize["branchPointVersionId"] = o.BranchPointVersionId
	}
	if o.BranchPointWorkspaceId != nil {
		toSerialize["branchPointWorkspaceId"] = o.BranchPointWorkspaceId
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.IsBranchPointAtStart != nil {
		toSerialize["isBranchPointAtStart"] = o.IsBranchPointAtStart
	}
	if o.SourceMicroversionId != nil {
		toSerialize["sourceMicroversionId"] = o.SourceMicroversionId
	}
	if o.TargetMicroversionId != nil {
		toSerialize["targetMicroversionId"] = o.TargetMicroversionId
	}
	if o.UpgradeInfo != nil {
		toSerialize["upgradeInfo"] = o.UpgradeInfo
	}
	return json.Marshal(toSerialize)
}

type NullableBTMergePreviewInfo struct {
	value *BTMergePreviewInfo
	isSet bool
}

func (v NullableBTMergePreviewInfo) Get() *BTMergePreviewInfo {
	return v.value
}

func (v *NullableBTMergePreviewInfo) Set(val *BTMergePreviewInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMergePreviewInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMergePreviewInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMergePreviewInfo(val *BTMergePreviewInfo) *NullableBTMergePreviewInfo {
	return &NullableBTMergePreviewInfo{value: val, isSet: true}
}

func (v NullableBTMergePreviewInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMergePreviewInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
