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

// BTBEmailExportOptions Options for exporting elements as a link in an email.
type BTBEmailExportOptions struct {
	// Use `true` if a link in an email should be sent.
	EmailLink bool `json:"emailLink"`
	// Message to send in the email body along with the download link.
	EmailMessage *string `json:"emailMessage,omitempty"`
	// Subject to send the email with.
	EmailSubject *string `json:"emailSubject,omitempty"`
	// List of email addresses to send the email to.
	EmailTo []string `json:"emailTo"`
	// Id of the user who does the export.
	FromUserId string `json:"fromUserId"`
	// A password to protect the email with.
	Password *string `json:"password,omitempty"`
	// Use `true` if the email should be protected with a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// Use `true` if email copy should be sent to the user who does the export.
	SendCopyToMe *bool `json:"sendCopyToMe,omitempty"`
	// Number of days to keep the link valid for.
	ValidForDays *int32 `json:"validForDays,omitempty"`
}

// NewBTBEmailExportOptions instantiates a new BTBEmailExportOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBEmailExportOptions(emailLink bool, emailTo []string, fromUserId string) *BTBEmailExportOptions {
	this := BTBEmailExportOptions{}
	this.EmailLink = emailLink
	var emailSubject string = "User sent you a file exported from Onshape"
	this.EmailSubject = &emailSubject
	this.EmailTo = emailTo
	this.FromUserId = fromUserId
	var password string = "false"
	this.Password = &password
	var passwordRequired bool = false
	this.PasswordRequired = &passwordRequired
	var sendCopyToMe bool = false
	this.SendCopyToMe = &sendCopyToMe
	var validForDays int32 = 3
	this.ValidForDays = &validForDays
	return &this
}

// NewBTBEmailExportOptionsWithDefaults instantiates a new BTBEmailExportOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBEmailExportOptionsWithDefaults() *BTBEmailExportOptions {
	this := BTBEmailExportOptions{}
	var emailLink bool = false
	this.EmailLink = emailLink
	var emailSubject string = "User sent you a file exported from Onshape"
	this.EmailSubject = &emailSubject
	var password string = "false"
	this.Password = &password
	var passwordRequired bool = false
	this.PasswordRequired = &passwordRequired
	var sendCopyToMe bool = false
	this.SendCopyToMe = &sendCopyToMe
	var validForDays int32 = 3
	this.ValidForDays = &validForDays
	return &this
}

// GetEmailLink returns the EmailLink field value
func (o *BTBEmailExportOptions) GetEmailLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailLink
}

// GetEmailLinkOk returns a tuple with the EmailLink field value
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetEmailLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailLink, true
}

// SetEmailLink sets field value
func (o *BTBEmailExportOptions) SetEmailLink(v bool) {
	o.EmailLink = v
}

// GetEmailMessage returns the EmailMessage field value if set, zero value otherwise.
func (o *BTBEmailExportOptions) GetEmailMessage() string {
	if o == nil || o.EmailMessage == nil {
		var ret string
		return ret
	}
	return *o.EmailMessage
}

// GetEmailMessageOk returns a tuple with the EmailMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetEmailMessageOk() (*string, bool) {
	if o == nil || o.EmailMessage == nil {
		return nil, false
	}
	return o.EmailMessage, true
}

// HasEmailMessage returns a boolean if a field has been set.
func (o *BTBEmailExportOptions) HasEmailMessage() bool {
	if o != nil && o.EmailMessage != nil {
		return true
	}

	return false
}

// SetEmailMessage gets a reference to the given string and assigns it to the EmailMessage field.
func (o *BTBEmailExportOptions) SetEmailMessage(v string) {
	o.EmailMessage = &v
}

// GetEmailSubject returns the EmailSubject field value if set, zero value otherwise.
func (o *BTBEmailExportOptions) GetEmailSubject() string {
	if o == nil || o.EmailSubject == nil {
		var ret string
		return ret
	}
	return *o.EmailSubject
}

// GetEmailSubjectOk returns a tuple with the EmailSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetEmailSubjectOk() (*string, bool) {
	if o == nil || o.EmailSubject == nil {
		return nil, false
	}
	return o.EmailSubject, true
}

// HasEmailSubject returns a boolean if a field has been set.
func (o *BTBEmailExportOptions) HasEmailSubject() bool {
	if o != nil && o.EmailSubject != nil {
		return true
	}

	return false
}

// SetEmailSubject gets a reference to the given string and assigns it to the EmailSubject field.
func (o *BTBEmailExportOptions) SetEmailSubject(v string) {
	o.EmailSubject = &v
}

// GetEmailTo returns the EmailTo field value
func (o *BTBEmailExportOptions) GetEmailTo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmailTo
}

// GetEmailToOk returns a tuple with the EmailTo field value
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetEmailToOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailTo, true
}

// SetEmailTo sets field value
func (o *BTBEmailExportOptions) SetEmailTo(v []string) {
	o.EmailTo = v
}

// GetFromUserId returns the FromUserId field value
func (o *BTBEmailExportOptions) GetFromUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetFromUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromUserId, true
}

// SetFromUserId sets field value
func (o *BTBEmailExportOptions) SetFromUserId(v string) {
	o.FromUserId = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *BTBEmailExportOptions) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *BTBEmailExportOptions) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *BTBEmailExportOptions) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordRequired returns the PasswordRequired field value if set, zero value otherwise.
func (o *BTBEmailExportOptions) GetPasswordRequired() bool {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret
	}
	return *o.PasswordRequired
}

// GetPasswordRequiredOk returns a tuple with the PasswordRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetPasswordRequiredOk() (*bool, bool) {
	if o == nil || o.PasswordRequired == nil {
		return nil, false
	}
	return o.PasswordRequired, true
}

// HasPasswordRequired returns a boolean if a field has been set.
func (o *BTBEmailExportOptions) HasPasswordRequired() bool {
	if o != nil && o.PasswordRequired != nil {
		return true
	}

	return false
}

// SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.
func (o *BTBEmailExportOptions) SetPasswordRequired(v bool) {
	o.PasswordRequired = &v
}

// GetSendCopyToMe returns the SendCopyToMe field value if set, zero value otherwise.
func (o *BTBEmailExportOptions) GetSendCopyToMe() bool {
	if o == nil || o.SendCopyToMe == nil {
		var ret bool
		return ret
	}
	return *o.SendCopyToMe
}

// GetSendCopyToMeOk returns a tuple with the SendCopyToMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetSendCopyToMeOk() (*bool, bool) {
	if o == nil || o.SendCopyToMe == nil {
		return nil, false
	}
	return o.SendCopyToMe, true
}

// HasSendCopyToMe returns a boolean if a field has been set.
func (o *BTBEmailExportOptions) HasSendCopyToMe() bool {
	if o != nil && o.SendCopyToMe != nil {
		return true
	}

	return false
}

// SetSendCopyToMe gets a reference to the given bool and assigns it to the SendCopyToMe field.
func (o *BTBEmailExportOptions) SetSendCopyToMe(v bool) {
	o.SendCopyToMe = &v
}

// GetValidForDays returns the ValidForDays field value if set, zero value otherwise.
func (o *BTBEmailExportOptions) GetValidForDays() int32 {
	if o == nil || o.ValidForDays == nil {
		var ret int32
		return ret
	}
	return *o.ValidForDays
}

// GetValidForDaysOk returns a tuple with the ValidForDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBEmailExportOptions) GetValidForDaysOk() (*int32, bool) {
	if o == nil || o.ValidForDays == nil {
		return nil, false
	}
	return o.ValidForDays, true
}

// HasValidForDays returns a boolean if a field has been set.
func (o *BTBEmailExportOptions) HasValidForDays() bool {
	if o != nil && o.ValidForDays != nil {
		return true
	}

	return false
}

// SetValidForDays gets a reference to the given int32 and assigns it to the ValidForDays field.
func (o *BTBEmailExportOptions) SetValidForDays(v int32) {
	o.ValidForDays = &v
}

func (o BTBEmailExportOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["emailLink"] = o.EmailLink
	}
	if o.EmailMessage != nil {
		toSerialize["emailMessage"] = o.EmailMessage
	}
	if o.EmailSubject != nil {
		toSerialize["emailSubject"] = o.EmailSubject
	}
	if true {
		toSerialize["emailTo"] = o.EmailTo
	}
	if true {
		toSerialize["fromUserId"] = o.FromUserId
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordRequired != nil {
		toSerialize["passwordRequired"] = o.PasswordRequired
	}
	if o.SendCopyToMe != nil {
		toSerialize["sendCopyToMe"] = o.SendCopyToMe
	}
	if o.ValidForDays != nil {
		toSerialize["validForDays"] = o.ValidForDays
	}
	return json.Marshal(toSerialize)
}

type NullableBTBEmailExportOptions struct {
	value *BTBEmailExportOptions
	isSet bool
}

func (v NullableBTBEmailExportOptions) Get() *BTBEmailExportOptions {
	return v.value
}

func (v *NullableBTBEmailExportOptions) Set(val *BTBEmailExportOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBEmailExportOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBEmailExportOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBEmailExportOptions(val *BTBEmailExportOptions) *NullableBTBEmailExportOptions {
	return &NullableBTBEmailExportOptions{value: val, isSet: true}
}

func (v NullableBTBEmailExportOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBEmailExportOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
