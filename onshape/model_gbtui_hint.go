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

// GBTUIHint the model 'GBTUIHint'
type GBTUIHint string

// List of GBTUIHint
const (
	GBTUIHintOppositeDirection                GBTUIHint = "OPPOSITE_DIRECTION"
	GBTUIHintAlwaysHidden                     GBTUIHint = "ALWAYS_HIDDEN"
	GBTUIHintShowCreateSelection              GBTUIHint = "SHOW_CREATE_SELECTION"
	GBTUIHintControlVisibility                GBTUIHint = "CONTROL_VISIBILITY"
	GBTUIHintNoPreviewProvided                GBTUIHint = "NO_PREVIEW_PROVIDED"
	GBTUIHintRememberPreviousValue            GBTUIHint = "REMEMBER_PREVIOUS_VALUE"
	GBTUIHintDisplayShort                     GBTUIHint = "DISPLAY_SHORT"
	GBTUIHintAllowFeatureSelection            GBTUIHint = "ALLOW_FEATURE_SELECTION"
	GBTUIHintMateConnectorAxisType            GBTUIHint = "MATE_CONNECTOR_AXIS_TYPE"
	GBTUIHintPrimaryAxis                      GBTUIHint = "PRIMARY_AXIS"
	GBTUIHintShowExpression                   GBTUIHint = "SHOW_EXPRESSION"
	GBTUIHintOppositeDirectionCircular        GBTUIHint = "OPPOSITE_DIRECTION_CIRCULAR"
	GBTUIHintShowLabel                        GBTUIHint = "SHOW_LABEL"
	GBTUIHintHorizontalEnum                   GBTUIHint = "HORIZONTAL_ENUM"
	GBTUIHintUnconfigurable                   GBTUIHint = "UNCONFIGURABLE"
	GBTUIHintMatchLastArrayItem               GBTUIHint = "MATCH_LAST_ARRAY_ITEM"
	GBTUIHintCollapseArrayItems               GBTUIHint = "COLLAPSE_ARRAY_ITEMS"
	GBTUIHintInitialFocusOnEdit               GBTUIHint = "INITIAL_FOCUS_ON_EDIT"
	GBTUIHintInitialFocus                     GBTUIHint = "INITIAL_FOCUS"
	GBTUIHintDisplayCurrentValueOnly          GBTUIHint = "DISPLAY_CURRENT_VALUE_ONLY"
	GBTUIHintReadOnly                         GBTUIHint = "READ_ONLY"
	GBTUIHintPreventCreatingNewMateConnectors GBTUIHint = "PREVENT_CREATING_NEW_MATE_CONNECTORS"
	GBTUIHintFirstInRow                       GBTUIHint = "FIRST_IN_ROW"
	GBTUIHintAllowQueryOrder                  GBTUIHint = "ALLOW_QUERY_ORDER"
	GBTUIHintPreventArrayReorder              GBTUIHint = "PREVENT_ARRAY_REORDER"
	GBTUIHintVariableName                     GBTUIHint = "VARIABLE_NAME"
	GBTUIHintFocusInnerQuery                  GBTUIHint = "FOCUS_INNER_QUERY"
	GBTUIHintShowTolerance                    GBTUIHint = "SHOW_TOLERANCE"
	GBTUIHintAllowArrayFocus                  GBTUIHint = "ALLOW_ARRAY_FOCUS"
	GBTUIHintShowInlineConfigInputs           GBTUIHint = "SHOW_INLINE_CONFIG_INPUTS"
	GBTUIHintFocusOnVisible                   GBTUIHint = "FOCUS_ON_VISIBLE"
	GBTUIHintCanBeTolerant                    GBTUIHint = "CAN_BE_TOLERANT"
	GBTUIHintCanBeTolerantDiameter            GBTUIHint = "CAN_BE_TOLERANT_DIAMETER"
	GBTUIHintPlusMinus                        GBTUIHint = "PLUS_MINUS"
	GBTUIHintAlwaysUseDependencies            GBTUIHint = "ALWAYS_USE_DEPENDENCIES"
	GBTUIHintUnknown                          GBTUIHint = "UNKNOWN"
)

// All allowed values of GBTUIHint enum
var AllowedGBTUIHintEnumValues = []GBTUIHint{
	"OPPOSITE_DIRECTION",
	"ALWAYS_HIDDEN",
	"SHOW_CREATE_SELECTION",
	"CONTROL_VISIBILITY",
	"NO_PREVIEW_PROVIDED",
	"REMEMBER_PREVIOUS_VALUE",
	"DISPLAY_SHORT",
	"ALLOW_FEATURE_SELECTION",
	"MATE_CONNECTOR_AXIS_TYPE",
	"PRIMARY_AXIS",
	"SHOW_EXPRESSION",
	"OPPOSITE_DIRECTION_CIRCULAR",
	"SHOW_LABEL",
	"HORIZONTAL_ENUM",
	"UNCONFIGURABLE",
	"MATCH_LAST_ARRAY_ITEM",
	"COLLAPSE_ARRAY_ITEMS",
	"INITIAL_FOCUS_ON_EDIT",
	"INITIAL_FOCUS",
	"DISPLAY_CURRENT_VALUE_ONLY",
	"READ_ONLY",
	"PREVENT_CREATING_NEW_MATE_CONNECTORS",
	"FIRST_IN_ROW",
	"ALLOW_QUERY_ORDER",
	"PREVENT_ARRAY_REORDER",
	"VARIABLE_NAME",
	"FOCUS_INNER_QUERY",
	"SHOW_TOLERANCE",
	"ALLOW_ARRAY_FOCUS",
	"SHOW_INLINE_CONFIG_INPUTS",
	"FOCUS_ON_VISIBLE",
	"CAN_BE_TOLERANT",
	"CAN_BE_TOLERANT_DIAMETER",
	"PLUS_MINUS",
	"ALWAYS_USE_DEPENDENCIES",
	"UNKNOWN",
}

func (v *GBTUIHint) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTUIHint(value)
	for _, existing := range AllowedGBTUIHintEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTUIHint", value)
}

// NewGBTUIHintFromValue returns a pointer to a valid GBTUIHint
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTUIHintFromValue(v string) (*GBTUIHint, error) {
	ev := GBTUIHint(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTUIHint: valid values are %v", v, AllowedGBTUIHintEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTUIHint) IsValid() bool {
	for _, existing := range AllowedGBTUIHintEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTUIHint value
func (v GBTUIHint) Ptr() *GBTUIHint {
	return &v
}

type NullableGBTUIHint struct {
	value *GBTUIHint
	isSet bool
}

func (v NullableGBTUIHint) Get() *GBTUIHint {
	return v.value
}

func (v *NullableGBTUIHint) Set(val *GBTUIHint) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTUIHint) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTUIHint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTUIHint(val *GBTUIHint) *NullableGBTUIHint {
	return &NullableGBTUIHint{value: val, isSet: true}
}

func (v NullableGBTUIHint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTUIHint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
