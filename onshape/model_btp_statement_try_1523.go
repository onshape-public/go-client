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

// BTPStatementTry1523 struct for BTPStatementTry1523
type BTPStatementTry1523 struct {
	BTPStatement269
	BtType              *string               `json:"btType,omitempty"`
	Atomic              *bool                 `json:"atomic,omitempty"`
	DocumentationType   *GBTPDefinitionType   `json:"documentationType,omitempty"`
	EndSourceLocation   *int32                `json:"endSourceLocation,omitempty"`
	NodeId              *string               `json:"nodeId,omitempty"`
	ShortDescriptor     *string               `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10           `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10           `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool                 `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32                `json:"startSourceLocation,omitempty"`
	Annotation          *BTPAnnotation231     `json:"annotation,omitempty"`
	Body                *BTPStatementBlock271 `json:"body,omitempty"`
	CatchBlock          *BTPStatementBlock271 `json:"catchBlock,omitempty"`
	CatchVariable       *BTPIdentifier8       `json:"catchVariable,omitempty"`
	Identifier          *BTPIdentifier8       `json:"identifier,omitempty"`
	Silent              *bool                 `json:"silent,omitempty"`
	SpaceAfterCatch     *BTPSpace10           `json:"spaceAfterCatch,omitempty"`
	SpaceBeforeSilent   *BTPSpace10           `json:"spaceBeforeSilent,omitempty"`
	StandardType        *GBTPType             `json:"standardType,omitempty"`
	TypeName            *string               `json:"typeName,omitempty"`
}

// NewBTPStatementTry1523 instantiates a new BTPStatementTry1523 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementTry1523() *BTPStatementTry1523 {
	this := BTPStatementTry1523{}
	return &this
}

// NewBTPStatementTry1523WithDefaults instantiates a new BTPStatementTry1523 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementTry1523WithDefaults() *BTPStatementTry1523 {
	this := BTPStatementTry1523{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementTry1523) SetBtType(v string) {
	o.BtType = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPStatementTry1523) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPStatementTry1523) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPStatementTry1523) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPStatementTry1523) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPStatementTry1523) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPStatementTry1523) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPStatementTry1523) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPStatementTry1523) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPStatementTry1523) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPStatementTry1523) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetBody() BTPStatementBlock271 {
	if o == nil || o.Body == nil {
		var ret BTPStatementBlock271
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetBodyOk() (*BTPStatementBlock271, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatementBlock271 and assigns it to the Body field.
func (o *BTPStatementTry1523) SetBody(v BTPStatementBlock271) {
	o.Body = &v
}

// GetCatchBlock returns the CatchBlock field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetCatchBlock() BTPStatementBlock271 {
	if o == nil || o.CatchBlock == nil {
		var ret BTPStatementBlock271
		return ret
	}
	return *o.CatchBlock
}

// GetCatchBlockOk returns a tuple with the CatchBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetCatchBlockOk() (*BTPStatementBlock271, bool) {
	if o == nil || o.CatchBlock == nil {
		return nil, false
	}
	return o.CatchBlock, true
}

// HasCatchBlock returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasCatchBlock() bool {
	if o != nil && o.CatchBlock != nil {
		return true
	}

	return false
}

// SetCatchBlock gets a reference to the given BTPStatementBlock271 and assigns it to the CatchBlock field.
func (o *BTPStatementTry1523) SetCatchBlock(v BTPStatementBlock271) {
	o.CatchBlock = &v
}

// GetCatchVariable returns the CatchVariable field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetCatchVariable() BTPIdentifier8 {
	if o == nil || o.CatchVariable == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.CatchVariable
}

// GetCatchVariableOk returns a tuple with the CatchVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetCatchVariableOk() (*BTPIdentifier8, bool) {
	if o == nil || o.CatchVariable == nil {
		return nil, false
	}
	return o.CatchVariable, true
}

// HasCatchVariable returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasCatchVariable() bool {
	if o != nil && o.CatchVariable != nil {
		return true
	}

	return false
}

// SetCatchVariable gets a reference to the given BTPIdentifier8 and assigns it to the CatchVariable field.
func (o *BTPStatementTry1523) SetCatchVariable(v BTPIdentifier8) {
	o.CatchVariable = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetIdentifier() BTPIdentifier8 {
	if o == nil || o.Identifier == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetIdentifierOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given BTPIdentifier8 and assigns it to the Identifier field.
func (o *BTPStatementTry1523) SetIdentifier(v BTPIdentifier8) {
	o.Identifier = &v
}

// GetSilent returns the Silent field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetSilent() bool {
	if o == nil || o.Silent == nil {
		var ret bool
		return ret
	}
	return *o.Silent
}

// GetSilentOk returns a tuple with the Silent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetSilentOk() (*bool, bool) {
	if o == nil || o.Silent == nil {
		return nil, false
	}
	return o.Silent, true
}

// HasSilent returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasSilent() bool {
	if o != nil && o.Silent != nil {
		return true
	}

	return false
}

// SetSilent gets a reference to the given bool and assigns it to the Silent field.
func (o *BTPStatementTry1523) SetSilent(v bool) {
	o.Silent = &v
}

// GetSpaceAfterCatch returns the SpaceAfterCatch field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetSpaceAfterCatch() BTPSpace10 {
	if o == nil || o.SpaceAfterCatch == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterCatch
}

// GetSpaceAfterCatchOk returns a tuple with the SpaceAfterCatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetSpaceAfterCatchOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterCatch == nil {
		return nil, false
	}
	return o.SpaceAfterCatch, true
}

// HasSpaceAfterCatch returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasSpaceAfterCatch() bool {
	if o != nil && o.SpaceAfterCatch != nil {
		return true
	}

	return false
}

// SetSpaceAfterCatch gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterCatch field.
func (o *BTPStatementTry1523) SetSpaceAfterCatch(v BTPSpace10) {
	o.SpaceAfterCatch = &v
}

// GetSpaceBeforeSilent returns the SpaceBeforeSilent field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetSpaceBeforeSilent() BTPSpace10 {
	if o == nil || o.SpaceBeforeSilent == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeSilent
}

// GetSpaceBeforeSilentOk returns a tuple with the SpaceBeforeSilent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetSpaceBeforeSilentOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeSilent == nil {
		return nil, false
	}
	return o.SpaceBeforeSilent, true
}

// HasSpaceBeforeSilent returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasSpaceBeforeSilent() bool {
	if o != nil && o.SpaceBeforeSilent != nil {
		return true
	}

	return false
}

// SetSpaceBeforeSilent gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeSilent field.
func (o *BTPStatementTry1523) SetSpaceBeforeSilent(v BTPSpace10) {
	o.SpaceBeforeSilent = &v
}

// GetStandardType returns the StandardType field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetStandardType() GBTPType {
	if o == nil || o.StandardType == nil {
		var ret GBTPType
		return ret
	}
	return *o.StandardType
}

// GetStandardTypeOk returns a tuple with the StandardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetStandardTypeOk() (*GBTPType, bool) {
	if o == nil || o.StandardType == nil {
		return nil, false
	}
	return o.StandardType, true
}

// HasStandardType returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasStandardType() bool {
	if o != nil && o.StandardType != nil {
		return true
	}

	return false
}

// SetStandardType gets a reference to the given GBTPType and assigns it to the StandardType field.
func (o *BTPStatementTry1523) SetStandardType(v GBTPType) {
	o.StandardType = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *BTPStatementTry1523) GetTypeName() string {
	if o == nil || o.TypeName == nil {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementTry1523) GetTypeNameOk() (*string, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BTPStatementTry1523) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *BTPStatementTry1523) SetTypeName(v string) {
	o.TypeName = &v
}

func (o BTPStatementTry1523) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPStatement269, errBTPStatement269 := json.Marshal(o.BTPStatement269)
	if errBTPStatement269 != nil {
		return []byte{}, errBTPStatement269
	}
	errBTPStatement269 = json.Unmarshal([]byte(serializedBTPStatement269), &toSerialize)
	if errBTPStatement269 != nil {
		return []byte{}, errBTPStatement269
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
	}
	if o.DocumentationType != nil {
		toSerialize["documentationType"] = o.DocumentationType
	}
	if o.EndSourceLocation != nil {
		toSerialize["endSourceLocation"] = o.EndSourceLocation
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ShortDescriptor != nil {
		toSerialize["shortDescriptor"] = o.ShortDescriptor
	}
	if o.SpaceAfter != nil {
		toSerialize["spaceAfter"] = o.SpaceAfter
	}
	if o.SpaceBefore != nil {
		toSerialize["spaceBefore"] = o.SpaceBefore
	}
	if o.SpaceDefault != nil {
		toSerialize["spaceDefault"] = o.SpaceDefault
	}
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.Annotation != nil {
		toSerialize["annotation"] = o.Annotation
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.CatchBlock != nil {
		toSerialize["catchBlock"] = o.CatchBlock
	}
	if o.CatchVariable != nil {
		toSerialize["catchVariable"] = o.CatchVariable
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Silent != nil {
		toSerialize["silent"] = o.Silent
	}
	if o.SpaceAfterCatch != nil {
		toSerialize["spaceAfterCatch"] = o.SpaceAfterCatch
	}
	if o.SpaceBeforeSilent != nil {
		toSerialize["spaceBeforeSilent"] = o.SpaceBeforeSilent
	}
	if o.StandardType != nil {
		toSerialize["standardType"] = o.StandardType
	}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementTry1523 struct {
	value *BTPStatementTry1523
	isSet bool
}

func (v NullableBTPStatementTry1523) Get() *BTPStatementTry1523 {
	return v.value
}

func (v *NullableBTPStatementTry1523) Set(val *BTPStatementTry1523) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementTry1523) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementTry1523) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementTry1523(val *BTPStatementTry1523) *NullableBTPStatementTry1523 {
	return &NullableBTPStatementTry1523{value: val, isSet: true}
}

func (v NullableBTPStatementTry1523) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementTry1523) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
