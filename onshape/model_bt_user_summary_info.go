/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserSummaryInfo struct for BTUserSummaryInfo
type BTUserSummaryInfo struct {
	Href                      *string               `json:"href,omitempty"`
	Id                        *string               `json:"id,omitempty"`
	Image                     *string               `json:"image,omitempty"`
	Name                      *string               `json:"name,omitempty"`
	State                     *int32                `json:"state,omitempty"`
	ViewRef                   *string               `json:"viewRef,omitempty"`
	Company                   *BTCompanySummaryInfo `json:"company,omitempty"`
	DocumentationName         *string               `json:"documentationName,omitempty"`
	DocumentationNameOverride *string               `json:"documentationNameOverride,omitempty"`
	Email                     *string               `json:"email,omitempty"`
	FirstName                 *string               `json:"firstName,omitempty"`
	GlobalPermissions         *GlobalPermissionInfo `json:"globalPermissions,omitempty"`
	IsGuest                   *bool                 `json:"isGuest,omitempty"`
	IsLight                   *bool                 `json:"isLight,omitempty"`
	LastLoginTime             *JSONTime             `json:"lastLoginTime,omitempty"`
	LastName                  *string               `json:"lastName,omitempty"`
	PersonalMessageAllowed    *bool                 `json:"personalMessageAllowed,omitempty"`
	Source                    *int32                `json:"source,omitempty"`
}

// NewBTUserSummaryInfo instantiates a new BTUserSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserSummaryInfo() *BTUserSummaryInfo {
	this := BTUserSummaryInfo{}
	return &this
}

// NewBTUserSummaryInfoWithDefaults instantiates a new BTUserSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserSummaryInfoWithDefaults() *BTUserSummaryInfo {
	this := BTUserSummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTUserSummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTUserSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetCompany() BTCompanySummaryInfo {
	if o == nil || o.Company == nil {
		var ret BTCompanySummaryInfo
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given BTCompanySummaryInfo and assigns it to the Company field.
func (o *BTUserSummaryInfo) SetCompany(v BTCompanySummaryInfo) {
	o.Company = &v
}

// GetDocumentationName returns the DocumentationName field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetDocumentationName() string {
	if o == nil || o.DocumentationName == nil {
		var ret string
		return ret
	}
	return *o.DocumentationName
}

// GetDocumentationNameOk returns a tuple with the DocumentationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetDocumentationNameOk() (*string, bool) {
	if o == nil || o.DocumentationName == nil {
		return nil, false
	}
	return o.DocumentationName, true
}

// HasDocumentationName returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasDocumentationName() bool {
	if o != nil && o.DocumentationName != nil {
		return true
	}

	return false
}

// SetDocumentationName gets a reference to the given string and assigns it to the DocumentationName field.
func (o *BTUserSummaryInfo) SetDocumentationName(v string) {
	o.DocumentationName = &v
}

// GetDocumentationNameOverride returns the DocumentationNameOverride field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetDocumentationNameOverride() string {
	if o == nil || o.DocumentationNameOverride == nil {
		var ret string
		return ret
	}
	return *o.DocumentationNameOverride
}

// GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetDocumentationNameOverrideOk() (*string, bool) {
	if o == nil || o.DocumentationNameOverride == nil {
		return nil, false
	}
	return o.DocumentationNameOverride, true
}

// HasDocumentationNameOverride returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasDocumentationNameOverride() bool {
	if o != nil && o.DocumentationNameOverride != nil {
		return true
	}

	return false
}

// SetDocumentationNameOverride gets a reference to the given string and assigns it to the DocumentationNameOverride field.
func (o *BTUserSummaryInfo) SetDocumentationNameOverride(v string) {
	o.DocumentationNameOverride = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTUserSummaryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTUserSummaryInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetGlobalPermissions() GlobalPermissionInfo {
	if o == nil || o.GlobalPermissions == nil {
		var ret GlobalPermissionInfo
		return ret
	}
	return *o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given GlobalPermissionInfo and assigns it to the GlobalPermissions field.
func (o *BTUserSummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo) {
	o.GlobalPermissions = &v
}

// GetIsGuest returns the IsGuest field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetIsGuest() bool {
	if o == nil || o.IsGuest == nil {
		var ret bool
		return ret
	}
	return *o.IsGuest
}

// GetIsGuestOk returns a tuple with the IsGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetIsGuestOk() (*bool, bool) {
	if o == nil || o.IsGuest == nil {
		return nil, false
	}
	return o.IsGuest, true
}

// HasIsGuest returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasIsGuest() bool {
	if o != nil && o.IsGuest != nil {
		return true
	}

	return false
}

// SetIsGuest gets a reference to the given bool and assigns it to the IsGuest field.
func (o *BTUserSummaryInfo) SetIsGuest(v bool) {
	o.IsGuest = &v
}

// GetIsLight returns the IsLight field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetIsLight() bool {
	if o == nil || o.IsLight == nil {
		var ret bool
		return ret
	}
	return *o.IsLight
}

// GetIsLightOk returns a tuple with the IsLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetIsLightOk() (*bool, bool) {
	if o == nil || o.IsLight == nil {
		return nil, false
	}
	return o.IsLight, true
}

// HasIsLight returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasIsLight() bool {
	if o != nil && o.IsLight != nil {
		return true
	}

	return false
}

// SetIsLight gets a reference to the given bool and assigns it to the IsLight field.
func (o *BTUserSummaryInfo) SetIsLight(v bool) {
	o.IsLight = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetLastLoginTime() JSONTime {
	if o == nil || o.LastLoginTime == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given JSONTime and assigns it to the LastLoginTime field.
func (o *BTUserSummaryInfo) SetLastLoginTime(v JSONTime) {
	o.LastLoginTime = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTUserSummaryInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetPersonalMessageAllowed returns the PersonalMessageAllowed field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetPersonalMessageAllowed() bool {
	if o == nil || o.PersonalMessageAllowed == nil {
		var ret bool
		return ret
	}
	return *o.PersonalMessageAllowed
}

// GetPersonalMessageAllowedOk returns a tuple with the PersonalMessageAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetPersonalMessageAllowedOk() (*bool, bool) {
	if o == nil || o.PersonalMessageAllowed == nil {
		return nil, false
	}
	return o.PersonalMessageAllowed, true
}

// HasPersonalMessageAllowed returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasPersonalMessageAllowed() bool {
	if o != nil && o.PersonalMessageAllowed != nil {
		return true
	}

	return false
}

// SetPersonalMessageAllowed gets a reference to the given bool and assigns it to the PersonalMessageAllowed field.
func (o *BTUserSummaryInfo) SetPersonalMessageAllowed(v bool) {
	o.PersonalMessageAllowed = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetSource() int32 {
	if o == nil || o.Source == nil {
		var ret int32
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetSourceOk() (*int32, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given int32 and assigns it to the Source field.
func (o *BTUserSummaryInfo) SetSource(v int32) {
	o.Source = &v
}

func (o BTUserSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.DocumentationName != nil {
		toSerialize["documentationName"] = o.DocumentationName
	}
	if o.DocumentationNameOverride != nil {
		toSerialize["documentationNameOverride"] = o.DocumentationNameOverride
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.IsGuest != nil {
		toSerialize["isGuest"] = o.IsGuest
	}
	if o.IsLight != nil {
		toSerialize["isLight"] = o.IsLight
	}
	if o.LastLoginTime != nil {
		toSerialize["lastLoginTime"] = o.LastLoginTime
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.PersonalMessageAllowed != nil {
		toSerialize["personalMessageAllowed"] = o.PersonalMessageAllowed
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserSummaryInfo struct {
	value *BTUserSummaryInfo
	isSet bool
}

func (v NullableBTUserSummaryInfo) Get() *BTUserSummaryInfo {
	return v.value
}

func (v *NullableBTUserSummaryInfo) Set(val *BTUserSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserSummaryInfo(val *BTUserSummaryInfo) *NullableBTUserSummaryInfo {
	return &NullableBTUserSummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}