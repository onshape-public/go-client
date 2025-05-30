# BTBEmailExportOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailLink** | **bool** | Use &#39;true&#39; if a link in an email should be sent. | [default to false]
**EmailMessage** | Pointer to **string** | Message to send in the email body along with the download link. | [optional] 
**EmailSubject** | Pointer to **string** | Subject to send the email with. | [optional] [default to "User sent you a file exported from Onshape"]
**EmailTo** | **[]string** | List of emails to send the email to. | 
**FromUserId** | **string** | Id of the user who does the export. | 
**Password** | Pointer to **string** | A password to protect the email with. | [optional] [default to "false"]
**PasswordRequired** | Pointer to **bool** | Use &#39;true&#39; if the email should be protected with a password. | [optional] [default to false]
**SendCopyToMe** | Pointer to **bool** | Use &#39;true&#39; if email copy should be sent to the user who does the export. | [optional] [default to false]
**ValidForDays** | Pointer to **int32** | Number of days to keep the link valid for. | [optional] [default to 3]

## Methods

### NewBTBEmailExportOptions

`func NewBTBEmailExportOptions(emailLink bool, emailTo []string, fromUserId string, ) *BTBEmailExportOptions`

NewBTBEmailExportOptions instantiates a new BTBEmailExportOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBEmailExportOptionsWithDefaults

`func NewBTBEmailExportOptionsWithDefaults() *BTBEmailExportOptions`

NewBTBEmailExportOptionsWithDefaults instantiates a new BTBEmailExportOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailLink

`func (o *BTBEmailExportOptions) GetEmailLink() bool`

GetEmailLink returns the EmailLink field if non-nil, zero value otherwise.

### GetEmailLinkOk

`func (o *BTBEmailExportOptions) GetEmailLinkOk() (*bool, bool)`

GetEmailLinkOk returns a tuple with the EmailLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLink

`func (o *BTBEmailExportOptions) SetEmailLink(v bool)`

SetEmailLink sets EmailLink field to given value.


### GetEmailMessage

`func (o *BTBEmailExportOptions) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *BTBEmailExportOptions) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *BTBEmailExportOptions) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *BTBEmailExportOptions) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetEmailSubject

`func (o *BTBEmailExportOptions) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *BTBEmailExportOptions) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *BTBEmailExportOptions) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *BTBEmailExportOptions) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailTo

`func (o *BTBEmailExportOptions) GetEmailTo() []string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *BTBEmailExportOptions) GetEmailToOk() (*[]string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *BTBEmailExportOptions) SetEmailTo(v []string)`

SetEmailTo sets EmailTo field to given value.


### GetFromUserId

`func (o *BTBEmailExportOptions) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *BTBEmailExportOptions) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *BTBEmailExportOptions) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.


### GetPassword

`func (o *BTBEmailExportOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BTBEmailExportOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BTBEmailExportOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BTBEmailExportOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordRequired

`func (o *BTBEmailExportOptions) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *BTBEmailExportOptions) GetPasswordRequiredOk() (*bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequired

`func (o *BTBEmailExportOptions) SetPasswordRequired(v bool)`

SetPasswordRequired sets PasswordRequired field to given value.

### HasPasswordRequired

`func (o *BTBEmailExportOptions) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### GetSendCopyToMe

`func (o *BTBEmailExportOptions) GetSendCopyToMe() bool`

GetSendCopyToMe returns the SendCopyToMe field if non-nil, zero value otherwise.

### GetSendCopyToMeOk

`func (o *BTBEmailExportOptions) GetSendCopyToMeOk() (*bool, bool)`

GetSendCopyToMeOk returns a tuple with the SendCopyToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCopyToMe

`func (o *BTBEmailExportOptions) SetSendCopyToMe(v bool)`

SetSendCopyToMe sets SendCopyToMe field to given value.

### HasSendCopyToMe

`func (o *BTBEmailExportOptions) HasSendCopyToMe() bool`

HasSendCopyToMe returns a boolean if a field has been set.

### GetValidForDays

`func (o *BTBEmailExportOptions) GetValidForDays() int32`

GetValidForDays returns the ValidForDays field if non-nil, zero value otherwise.

### GetValidForDaysOk

`func (o *BTBEmailExportOptions) GetValidForDaysOk() (*int32, bool)`

GetValidForDaysOk returns a tuple with the ValidForDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidForDays

`func (o *BTBEmailExportOptions) SetValidForDays(v int32)`

SetValidForDays sets ValidForDays field to given value.

### HasValidForDays

`func (o *BTBEmailExportOptions) HasValidForDays() bool`

HasValidForDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


