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

// BTDocumentMessageBody struct for BTDocumentMessageBody
type BTDocumentMessageBody struct {
	AppElementSessionId *string   `json:"appElementSessionId,omitempty"`
	CommentId           *string   `json:"commentId,omitempty"`
	Data                *string   `json:"data,omitempty"`
	DocumentId          *string   `json:"documentId,omitempty"`
	DocumentState       *string   `json:"documentState,omitempty"`
	DocumentType        *int32    `json:"documentType,omitempty"`
	ElementId           *string   `json:"elementId,omitempty"`
	Event               *string   `json:"event,omitempty"`
	MessageId           *string   `json:"messageId,omitempty"`
	MetadataObjectType  *string   `json:"metadataObjectType,omitempty"`
	PartId              *string   `json:"partId,omitempty"`
	PartIdentity        *string   `json:"partIdentity,omitempty"`
	PartNumber          *string   `json:"partNumber,omitempty"`
	Timestamp           *JSONTime `json:"timestamp,omitempty"`
	TranslationId       *string   `json:"translationId,omitempty"`
	UserId              *string   `json:"userId,omitempty"`
	VersionId           *string   `json:"versionId,omitempty"`
	WebhookId           *string   `json:"webhookId,omitempty"`
	WorkspaceId         *string   `json:"workspaceId,omitempty"`
}

// NewBTDocumentMessageBody instantiates a new BTDocumentMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentMessageBody() *BTDocumentMessageBody {
	this := BTDocumentMessageBody{}
	return &this
}

// NewBTDocumentMessageBodyWithDefaults instantiates a new BTDocumentMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentMessageBodyWithDefaults() *BTDocumentMessageBody {
	this := BTDocumentMessageBody{}
	return &this
}

// GetAppElementSessionId returns the AppElementSessionId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetAppElementSessionId() string {
	if o == nil || o.AppElementSessionId == nil {
		var ret string
		return ret
	}
	return *o.AppElementSessionId
}

// GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetAppElementSessionIdOk() (*string, bool) {
	if o == nil || o.AppElementSessionId == nil {
		return nil, false
	}
	return o.AppElementSessionId, true
}

// HasAppElementSessionId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasAppElementSessionId() bool {
	if o != nil && o.AppElementSessionId != nil {
		return true
	}

	return false
}

// SetAppElementSessionId gets a reference to the given string and assigns it to the AppElementSessionId field.
func (o *BTDocumentMessageBody) SetAppElementSessionId(v string) {
	o.AppElementSessionId = &v
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetCommentId() string {
	if o == nil || o.CommentId == nil {
		var ret string
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetCommentIdOk() (*string, bool) {
	if o == nil || o.CommentId == nil {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasCommentId() bool {
	if o != nil && o.CommentId != nil {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given string and assigns it to the CommentId field.
func (o *BTDocumentMessageBody) SetCommentId(v string) {
	o.CommentId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTDocumentMessageBody) SetData(v string) {
	o.Data = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTDocumentMessageBody) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentState returns the DocumentState field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetDocumentState() string {
	if o == nil || o.DocumentState == nil {
		var ret string
		return ret
	}
	return *o.DocumentState
}

// GetDocumentStateOk returns a tuple with the DocumentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetDocumentStateOk() (*string, bool) {
	if o == nil || o.DocumentState == nil {
		return nil, false
	}
	return o.DocumentState, true
}

// HasDocumentState returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasDocumentState() bool {
	if o != nil && o.DocumentState != nil {
		return true
	}

	return false
}

// SetDocumentState gets a reference to the given string and assigns it to the DocumentState field.
func (o *BTDocumentMessageBody) SetDocumentState(v string) {
	o.DocumentState = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetDocumentType() int32 {
	if o == nil || o.DocumentType == nil {
		var ret int32
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetDocumentTypeOk() (*int32, bool) {
	if o == nil || o.DocumentType == nil {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasDocumentType() bool {
	if o != nil && o.DocumentType != nil {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given int32 and assigns it to the DocumentType field.
func (o *BTDocumentMessageBody) SetDocumentType(v int32) {
	o.DocumentType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTDocumentMessageBody) SetElementId(v string) {
	o.ElementId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *BTDocumentMessageBody) SetEvent(v string) {
	o.Event = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *BTDocumentMessageBody) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMetadataObjectType returns the MetadataObjectType field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetMetadataObjectType() string {
	if o == nil || o.MetadataObjectType == nil {
		var ret string
		return ret
	}
	return *o.MetadataObjectType
}

// GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetMetadataObjectTypeOk() (*string, bool) {
	if o == nil || o.MetadataObjectType == nil {
		return nil, false
	}
	return o.MetadataObjectType, true
}

// HasMetadataObjectType returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasMetadataObjectType() bool {
	if o != nil && o.MetadataObjectType != nil {
		return true
	}

	return false
}

// SetMetadataObjectType gets a reference to the given string and assigns it to the MetadataObjectType field.
func (o *BTDocumentMessageBody) SetMetadataObjectType(v string) {
	o.MetadataObjectType = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTDocumentMessageBody) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetPartIdentity() string {
	if o == nil || o.PartIdentity == nil {
		var ret string
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetPartIdentityOk() (*string, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given string and assigns it to the PartIdentity field.
func (o *BTDocumentMessageBody) SetPartIdentity(v string) {
	o.PartIdentity = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTDocumentMessageBody) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetTimestamp() JSONTime {
	if o == nil || o.Timestamp == nil {
		var ret JSONTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetTimestampOk() (*JSONTime, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given JSONTime and assigns it to the Timestamp field.
func (o *BTDocumentMessageBody) SetTimestamp(v JSONTime) {
	o.Timestamp = &v
}

// GetTranslationId returns the TranslationId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetTranslationId() string {
	if o == nil || o.TranslationId == nil {
		var ret string
		return ret
	}
	return *o.TranslationId
}

// GetTranslationIdOk returns a tuple with the TranslationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetTranslationIdOk() (*string, bool) {
	if o == nil || o.TranslationId == nil {
		return nil, false
	}
	return o.TranslationId, true
}

// HasTranslationId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasTranslationId() bool {
	if o != nil && o.TranslationId != nil {
		return true
	}

	return false
}

// SetTranslationId gets a reference to the given string and assigns it to the TranslationId field.
func (o *BTDocumentMessageBody) SetTranslationId(v string) {
	o.TranslationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTDocumentMessageBody) SetUserId(v string) {
	o.UserId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTDocumentMessageBody) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *BTDocumentMessageBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTDocumentMessageBody) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMessageBody) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTDocumentMessageBody) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTDocumentMessageBody) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTDocumentMessageBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppElementSessionId != nil {
		toSerialize["appElementSessionId"] = o.AppElementSessionId
	}
	if o.CommentId != nil {
		toSerialize["commentId"] = o.CommentId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentState != nil {
		toSerialize["documentState"] = o.DocumentState
	}
	if o.DocumentType != nil {
		toSerialize["documentType"] = o.DocumentType
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.MetadataObjectType != nil {
		toSerialize["metadataObjectType"] = o.MetadataObjectType
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartIdentity != nil {
		toSerialize["partIdentity"] = o.PartIdentity
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TranslationId != nil {
		toSerialize["translationId"] = o.TranslationId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WebhookId != nil {
		toSerialize["webhookId"] = o.WebhookId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentMessageBody struct {
	value *BTDocumentMessageBody
	isSet bool
}

func (v NullableBTDocumentMessageBody) Get() *BTDocumentMessageBody {
	return v.value
}

func (v *NullableBTDocumentMessageBody) Set(val *BTDocumentMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentMessageBody(val *BTDocumentMessageBody) *NullableBTDocumentMessageBody {
	return &NullableBTDocumentMessageBody{value: val, isSet: true}
}

func (v NullableBTDocumentMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
