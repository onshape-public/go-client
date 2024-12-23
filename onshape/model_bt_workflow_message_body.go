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

// BTWorkflowMessageBody struct for BTWorkflowMessageBody
type BTWorkflowMessageBody struct {
	AppElementSessionId *string   `json:"appElementSessionId,omitempty"`
	Data                *string   `json:"data,omitempty"`
	Event               *string   `json:"event,omitempty"`
	MessageId           *string   `json:"messageId,omitempty"`
	ObjectId            *string   `json:"objectId,omitempty"`
	ObjectType          *string   `json:"objectType,omitempty"`
	Timestamp           *JSONTime `json:"timestamp,omitempty"`
	TransitionName      *string   `json:"transitionName,omitempty"`
	WebhookId           *string   `json:"webhookId,omitempty"`
	WorkflowId          *string   `json:"workflowId,omitempty"`
}

// NewBTWorkflowMessageBody instantiates a new BTWorkflowMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowMessageBody() *BTWorkflowMessageBody {
	this := BTWorkflowMessageBody{}
	return &this
}

// NewBTWorkflowMessageBodyWithDefaults instantiates a new BTWorkflowMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowMessageBodyWithDefaults() *BTWorkflowMessageBody {
	this := BTWorkflowMessageBody{}
	return &this
}

// GetAppElementSessionId returns the AppElementSessionId field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetAppElementSessionId() string {
	if o == nil || o.AppElementSessionId == nil {
		var ret string
		return ret
	}
	return *o.AppElementSessionId
}

// GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetAppElementSessionIdOk() (*string, bool) {
	if o == nil || o.AppElementSessionId == nil {
		return nil, false
	}
	return o.AppElementSessionId, true
}

// HasAppElementSessionId returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasAppElementSessionId() bool {
	if o != nil && o.AppElementSessionId != nil {
		return true
	}

	return false
}

// SetAppElementSessionId gets a reference to the given string and assigns it to the AppElementSessionId field.
func (o *BTWorkflowMessageBody) SetAppElementSessionId(v string) {
	o.AppElementSessionId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTWorkflowMessageBody) SetData(v string) {
	o.Data = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *BTWorkflowMessageBody) SetEvent(v string) {
	o.Event = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *BTWorkflowMessageBody) SetMessageId(v string) {
	o.MessageId = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTWorkflowMessageBody) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *BTWorkflowMessageBody) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetTimestamp() JSONTime {
	if o == nil || o.Timestamp == nil {
		var ret JSONTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetTimestampOk() (*JSONTime, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given JSONTime and assigns it to the Timestamp field.
func (o *BTWorkflowMessageBody) SetTimestamp(v JSONTime) {
	o.Timestamp = &v
}

// GetTransitionName returns the TransitionName field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetTransitionName() string {
	if o == nil || o.TransitionName == nil {
		var ret string
		return ret
	}
	return *o.TransitionName
}

// GetTransitionNameOk returns a tuple with the TransitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetTransitionNameOk() (*string, bool) {
	if o == nil || o.TransitionName == nil {
		return nil, false
	}
	return o.TransitionName, true
}

// HasTransitionName returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasTransitionName() bool {
	if o != nil && o.TransitionName != nil {
		return true
	}

	return false
}

// SetTransitionName gets a reference to the given string and assigns it to the TransitionName field.
func (o *BTWorkflowMessageBody) SetTransitionName(v string) {
	o.TransitionName = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *BTWorkflowMessageBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *BTWorkflowMessageBody) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowMessageBody) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *BTWorkflowMessageBody) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *BTWorkflowMessageBody) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

func (o BTWorkflowMessageBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppElementSessionId != nil {
		toSerialize["appElementSessionId"] = o.AppElementSessionId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TransitionName != nil {
		toSerialize["transitionName"] = o.TransitionName
	}
	if o.WebhookId != nil {
		toSerialize["webhookId"] = o.WebhookId
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowMessageBody struct {
	value *BTWorkflowMessageBody
	isSet bool
}

func (v NullableBTWorkflowMessageBody) Get() *BTWorkflowMessageBody {
	return v.value
}

func (v *NullableBTWorkflowMessageBody) Set(val *BTWorkflowMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowMessageBody(val *BTWorkflowMessageBody) *NullableBTWorkflowMessageBody {
	return &NullableBTWorkflowMessageBody{value: val, isSet: true}
}

func (v NullableBTWorkflowMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
