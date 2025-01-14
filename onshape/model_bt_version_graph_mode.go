/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTVersionGraphMode the model 'BTVersionGraphMode'
type BTVersionGraphMode string

// List of BTVersionGraphMode
const (
	BTVersionGraphModeAllBranches               BTVersionGraphMode = "ALL_BRANCHES"
	BTVersionGraphModeActiveBranch              BTVersionGraphMode = "ACTIVE_BRANCH"
	BTVersionGraphModeAllBranchesWithWorkspaces BTVersionGraphMode = "ALL_BRANCHES_WITH_WORKSPACES"
	BTVersionGraphModeActiveBranchWithParents   BTVersionGraphMode = "ACTIVE_BRANCH_WITH_PARENTS"
)

// All allowed values of BTVersionGraphMode enum
var AllowedBTVersionGraphModeEnumValues = []BTVersionGraphMode{
	"ALL_BRANCHES",
	"ACTIVE_BRANCH",
	"ALL_BRANCHES_WITH_WORKSPACES",
	"ACTIVE_BRANCH_WITH_PARENTS",
}

func (v *BTVersionGraphMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTVersionGraphMode(value)
	for _, existing := range AllowedBTVersionGraphModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTVersionGraphMode", value)
}

// NewBTVersionGraphModeFromValue returns a pointer to a valid BTVersionGraphMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTVersionGraphModeFromValue(v string) (*BTVersionGraphMode, error) {
	ev := BTVersionGraphMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTVersionGraphMode: valid values are %v", v, AllowedBTVersionGraphModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTVersionGraphMode) IsValid() bool {
	for _, existing := range AllowedBTVersionGraphModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTVersionGraphMode value
func (v BTVersionGraphMode) Ptr() *BTVersionGraphMode {
	return &v
}

type NullableBTVersionGraphMode struct {
	value *BTVersionGraphMode
	isSet bool
}

func (v NullableBTVersionGraphMode) Get() *BTVersionGraphMode {
	return v.value
}

func (v *NullableBTVersionGraphMode) Set(val *BTVersionGraphMode) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVersionGraphMode) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVersionGraphMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVersionGraphMode(val *BTVersionGraphMode) *NullableBTVersionGraphMode {
	return &NullableBTVersionGraphMode{value: val, isSet: true}
}

func (v NullableBTVersionGraphMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVersionGraphMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
