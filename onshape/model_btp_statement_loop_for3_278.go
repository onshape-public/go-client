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

// BTPStatementLoopFor3278 struct for BTPStatementLoopFor3278
type BTPStatementLoopFor3278 struct {
	BTPStatementLoop277
	BtType                   *string             `json:"btType,omitempty"`
	Atomic                   *bool               `json:"atomic,omitempty"`
	DocumentationType        *GBTPDefinitionType `json:"documentationType,omitempty"`
	EndSourceLocation        *int32              `json:"endSourceLocation,omitempty"`
	NodeId                   *string             `json:"nodeId,omitempty"`
	ShortDescriptor          *string             `json:"shortDescriptor,omitempty"`
	SpaceAfter               *BTPSpace10         `json:"spaceAfter,omitempty"`
	SpaceBefore              *BTPSpace10         `json:"spaceBefore,omitempty"`
	SpaceDefault             *bool               `json:"spaceDefault,omitempty"`
	StartSourceLocation      *int32              `json:"startSourceLocation,omitempty"`
	Annotation               *BTPAnnotation231   `json:"annotation,omitempty"`
	Body                     *BTPStatement269    `json:"body,omitempty"`
	SpaceAfterLoopType       *BTPSpace10         `json:"spaceAfterLoopType,omitempty"`
	Condition                *BTPExpression9     `json:"condition,omitempty"`
	Increment                *BTPStatement269    `json:"increment,omitempty"`
	Initialization           *BTPStatement269    `json:"initialization,omitempty"`
	SpaceAfterInitialization *BTPSpace10         `json:"spaceAfterInitialization,omitempty"`
	SpaceBeforeCondition     *BTPSpace10         `json:"spaceBeforeCondition,omitempty"`
	SpaceBeforeIncrement     *BTPSpace10         `json:"spaceBeforeIncrement,omitempty"`
}

// NewBTPStatementLoopFor3278 instantiates a new BTPStatementLoopFor3278 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoopFor3278() *BTPStatementLoopFor3278 {
	this := BTPStatementLoopFor3278{}
	return &this
}

// NewBTPStatementLoopFor3278WithDefaults instantiates a new BTPStatementLoopFor3278 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoopFor3278WithDefaults() *BTPStatementLoopFor3278 {
	this := BTPStatementLoopFor3278{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoopFor3278) SetBtType(v string) {
	o.BtType = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPStatementLoopFor3278) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPStatementLoopFor3278) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPStatementLoopFor3278) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPStatementLoopFor3278) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPStatementLoopFor3278) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPStatementLoopFor3278) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPStatementLoopFor3278) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPStatementLoopFor3278) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPStatementLoopFor3278) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPStatementLoopFor3278) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetBody() BTPStatement269 {
	if o == nil || o.Body == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetBodyOk() (*BTPStatement269, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatement269 and assigns it to the Body field.
func (o *BTPStatementLoopFor3278) SetBody(v BTPStatement269) {
	o.Body = &v
}

// GetSpaceAfterLoopType returns the SpaceAfterLoopType field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceAfterLoopType() BTPSpace10 {
	if o == nil || o.SpaceAfterLoopType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterLoopType
}

// GetSpaceAfterLoopTypeOk returns a tuple with the SpaceAfterLoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterLoopType == nil {
		return nil, false
	}
	return o.SpaceAfterLoopType, true
}

// HasSpaceAfterLoopType returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceAfterLoopType() bool {
	if o != nil && o.SpaceAfterLoopType != nil {
		return true
	}

	return false
}

// SetSpaceAfterLoopType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterLoopType field.
func (o *BTPStatementLoopFor3278) SetSpaceAfterLoopType(v BTPSpace10) {
	o.SpaceAfterLoopType = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetCondition() BTPExpression9 {
	if o == nil || o.Condition == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetConditionOk() (*BTPExpression9, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given BTPExpression9 and assigns it to the Condition field.
func (o *BTPStatementLoopFor3278) SetCondition(v BTPExpression9) {
	o.Condition = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetIncrement() BTPStatement269 {
	if o == nil || o.Increment == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetIncrementOk() (*BTPStatement269, bool) {
	if o == nil || o.Increment == nil {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasIncrement() bool {
	if o != nil && o.Increment != nil {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given BTPStatement269 and assigns it to the Increment field.
func (o *BTPStatementLoopFor3278) SetIncrement(v BTPStatement269) {
	o.Increment = &v
}

// GetInitialization returns the Initialization field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetInitialization() BTPStatement269 {
	if o == nil || o.Initialization == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Initialization
}

// GetInitializationOk returns a tuple with the Initialization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetInitializationOk() (*BTPStatement269, bool) {
	if o == nil || o.Initialization == nil {
		return nil, false
	}
	return o.Initialization, true
}

// HasInitialization returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasInitialization() bool {
	if o != nil && o.Initialization != nil {
		return true
	}

	return false
}

// SetInitialization gets a reference to the given BTPStatement269 and assigns it to the Initialization field.
func (o *BTPStatementLoopFor3278) SetInitialization(v BTPStatement269) {
	o.Initialization = &v
}

// GetSpaceAfterInitialization returns the SpaceAfterInitialization field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceAfterInitialization() BTPSpace10 {
	if o == nil || o.SpaceAfterInitialization == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterInitialization
}

// GetSpaceAfterInitializationOk returns a tuple with the SpaceAfterInitialization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceAfterInitializationOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterInitialization == nil {
		return nil, false
	}
	return o.SpaceAfterInitialization, true
}

// HasSpaceAfterInitialization returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceAfterInitialization() bool {
	if o != nil && o.SpaceAfterInitialization != nil {
		return true
	}

	return false
}

// SetSpaceAfterInitialization gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterInitialization field.
func (o *BTPStatementLoopFor3278) SetSpaceAfterInitialization(v BTPSpace10) {
	o.SpaceAfterInitialization = &v
}

// GetSpaceBeforeCondition returns the SpaceBeforeCondition field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceBeforeCondition() BTPSpace10 {
	if o == nil || o.SpaceBeforeCondition == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeCondition
}

// GetSpaceBeforeConditionOk returns a tuple with the SpaceBeforeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceBeforeConditionOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeCondition == nil {
		return nil, false
	}
	return o.SpaceBeforeCondition, true
}

// HasSpaceBeforeCondition returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceBeforeCondition() bool {
	if o != nil && o.SpaceBeforeCondition != nil {
		return true
	}

	return false
}

// SetSpaceBeforeCondition gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeCondition field.
func (o *BTPStatementLoopFor3278) SetSpaceBeforeCondition(v BTPSpace10) {
	o.SpaceBeforeCondition = &v
}

// GetSpaceBeforeIncrement returns the SpaceBeforeIncrement field value if set, zero value otherwise.
func (o *BTPStatementLoopFor3278) GetSpaceBeforeIncrement() BTPSpace10 {
	if o == nil || o.SpaceBeforeIncrement == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeIncrement
}

// GetSpaceBeforeIncrementOk returns a tuple with the SpaceBeforeIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopFor3278) GetSpaceBeforeIncrementOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeIncrement == nil {
		return nil, false
	}
	return o.SpaceBeforeIncrement, true
}

// HasSpaceBeforeIncrement returns a boolean if a field has been set.
func (o *BTPStatementLoopFor3278) HasSpaceBeforeIncrement() bool {
	if o != nil && o.SpaceBeforeIncrement != nil {
		return true
	}

	return false
}

// SetSpaceBeforeIncrement gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeIncrement field.
func (o *BTPStatementLoopFor3278) SetSpaceBeforeIncrement(v BTPSpace10) {
	o.SpaceBeforeIncrement = &v
}

func (o BTPStatementLoopFor3278) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPStatementLoop277, errBTPStatementLoop277 := json.Marshal(o.BTPStatementLoop277)
	if errBTPStatementLoop277 != nil {
		return []byte{}, errBTPStatementLoop277
	}
	errBTPStatementLoop277 = json.Unmarshal([]byte(serializedBTPStatementLoop277), &toSerialize)
	if errBTPStatementLoop277 != nil {
		return []byte{}, errBTPStatementLoop277
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
	if o.SpaceAfterLoopType != nil {
		toSerialize["spaceAfterLoopType"] = o.SpaceAfterLoopType
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Increment != nil {
		toSerialize["increment"] = o.Increment
	}
	if o.Initialization != nil {
		toSerialize["initialization"] = o.Initialization
	}
	if o.SpaceAfterInitialization != nil {
		toSerialize["spaceAfterInitialization"] = o.SpaceAfterInitialization
	}
	if o.SpaceBeforeCondition != nil {
		toSerialize["spaceBeforeCondition"] = o.SpaceBeforeCondition
	}
	if o.SpaceBeforeIncrement != nil {
		toSerialize["spaceBeforeIncrement"] = o.SpaceBeforeIncrement
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementLoopFor3278 struct {
	value *BTPStatementLoopFor3278
	isSet bool
}

func (v NullableBTPStatementLoopFor3278) Get() *BTPStatementLoopFor3278 {
	return v.value
}

func (v *NullableBTPStatementLoopFor3278) Set(val *BTPStatementLoopFor3278) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoopFor3278) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoopFor3278) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoopFor3278(val *BTPStatementLoopFor3278) *NullableBTPStatementLoopFor3278 {
	return &NullableBTPStatementLoopFor3278{value: val, isSet: true}
}

func (v NullableBTPStatementLoopFor3278) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoopFor3278) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
