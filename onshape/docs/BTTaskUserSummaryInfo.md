# BTTaskUserSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acted** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to [**BTCompanySummaryInfo**](BTCompanySummaryInfo.md) |  | [optional] 
**DocumentationName** | Pointer to **string** |  | [optional] 
**DocumentationNameOverride** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**GlobalPermissions** | Pointer to [**GlobalPermissionInfo**](GlobalPermissionInfo.md) |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**InvitationState** | Pointer to **int32** |  | [optional] 
**IsGuest** | Pointer to **bool** |  | [optional] 
**IsLight** | Pointer to **bool** |  | [optional] 
**IsOnshapeSupport** | Pointer to **bool** |  | [optional] 
**LastLoginTime** | Pointer to **JSONTime** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**PersonalMessageAllowed** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTTaskUserSummaryInfo

`func NewBTTaskUserSummaryInfo() *BTTaskUserSummaryInfo`

NewBTTaskUserSummaryInfo instantiates a new BTTaskUserSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTaskUserSummaryInfoWithDefaults

`func NewBTTaskUserSummaryInfoWithDefaults() *BTTaskUserSummaryInfo`

NewBTTaskUserSummaryInfoWithDefaults instantiates a new BTTaskUserSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActed

`func (o *BTTaskUserSummaryInfo) GetActed() bool`

GetActed returns the Acted field if non-nil, zero value otherwise.

### GetActedOk

`func (o *BTTaskUserSummaryInfo) GetActedOk() (*bool, bool)`

GetActedOk returns a tuple with the Acted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActed

`func (o *BTTaskUserSummaryInfo) SetActed(v bool)`

SetActed sets Acted field to given value.

### HasActed

`func (o *BTTaskUserSummaryInfo) HasActed() bool`

HasActed returns a boolean if a field has been set.

### GetCompany

`func (o *BTTaskUserSummaryInfo) GetCompany() BTCompanySummaryInfo`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BTTaskUserSummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BTTaskUserSummaryInfo) SetCompany(v BTCompanySummaryInfo)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BTTaskUserSummaryInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDocumentationName

`func (o *BTTaskUserSummaryInfo) GetDocumentationName() string`

GetDocumentationName returns the DocumentationName field if non-nil, zero value otherwise.

### GetDocumentationNameOk

`func (o *BTTaskUserSummaryInfo) GetDocumentationNameOk() (*string, bool)`

GetDocumentationNameOk returns a tuple with the DocumentationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationName

`func (o *BTTaskUserSummaryInfo) SetDocumentationName(v string)`

SetDocumentationName sets DocumentationName field to given value.

### HasDocumentationName

`func (o *BTTaskUserSummaryInfo) HasDocumentationName() bool`

HasDocumentationName returns a boolean if a field has been set.

### GetDocumentationNameOverride

`func (o *BTTaskUserSummaryInfo) GetDocumentationNameOverride() string`

GetDocumentationNameOverride returns the DocumentationNameOverride field if non-nil, zero value otherwise.

### GetDocumentationNameOverrideOk

`func (o *BTTaskUserSummaryInfo) GetDocumentationNameOverrideOk() (*string, bool)`

GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationNameOverride

`func (o *BTTaskUserSummaryInfo) SetDocumentationNameOverride(v string)`

SetDocumentationNameOverride sets DocumentationNameOverride field to given value.

### HasDocumentationNameOverride

`func (o *BTTaskUserSummaryInfo) HasDocumentationNameOverride() bool`

HasDocumentationNameOverride returns a boolean if a field has been set.

### GetEmail

`func (o *BTTaskUserSummaryInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BTTaskUserSummaryInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BTTaskUserSummaryInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BTTaskUserSummaryInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *BTTaskUserSummaryInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BTTaskUserSummaryInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BTTaskUserSummaryInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BTTaskUserSummaryInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGlobalPermissions

`func (o *BTTaskUserSummaryInfo) GetGlobalPermissions() GlobalPermissionInfo`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *BTTaskUserSummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *BTTaskUserSummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo)`

SetGlobalPermissions sets GlobalPermissions field to given value.

### HasGlobalPermissions

`func (o *BTTaskUserSummaryInfo) HasGlobalPermissions() bool`

HasGlobalPermissions returns a boolean if a field has been set.

### GetHref

`func (o *BTTaskUserSummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTaskUserSummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTaskUserSummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTaskUserSummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTaskUserSummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTaskUserSummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTaskUserSummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTaskUserSummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *BTTaskUserSummaryInfo) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BTTaskUserSummaryInfo) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BTTaskUserSummaryInfo) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BTTaskUserSummaryInfo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInvitationState

`func (o *BTTaskUserSummaryInfo) GetInvitationState() int32`

GetInvitationState returns the InvitationState field if non-nil, zero value otherwise.

### GetInvitationStateOk

`func (o *BTTaskUserSummaryInfo) GetInvitationStateOk() (*int32, bool)`

GetInvitationStateOk returns a tuple with the InvitationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationState

`func (o *BTTaskUserSummaryInfo) SetInvitationState(v int32)`

SetInvitationState sets InvitationState field to given value.

### HasInvitationState

`func (o *BTTaskUserSummaryInfo) HasInvitationState() bool`

HasInvitationState returns a boolean if a field has been set.

### GetIsGuest

`func (o *BTTaskUserSummaryInfo) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *BTTaskUserSummaryInfo) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *BTTaskUserSummaryInfo) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *BTTaskUserSummaryInfo) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### GetIsLight

`func (o *BTTaskUserSummaryInfo) GetIsLight() bool`

GetIsLight returns the IsLight field if non-nil, zero value otherwise.

### GetIsLightOk

`func (o *BTTaskUserSummaryInfo) GetIsLightOk() (*bool, bool)`

GetIsLightOk returns a tuple with the IsLight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLight

`func (o *BTTaskUserSummaryInfo) SetIsLight(v bool)`

SetIsLight sets IsLight field to given value.

### HasIsLight

`func (o *BTTaskUserSummaryInfo) HasIsLight() bool`

HasIsLight returns a boolean if a field has been set.

### GetIsOnshapeSupport

`func (o *BTTaskUserSummaryInfo) GetIsOnshapeSupport() bool`

GetIsOnshapeSupport returns the IsOnshapeSupport field if non-nil, zero value otherwise.

### GetIsOnshapeSupportOk

`func (o *BTTaskUserSummaryInfo) GetIsOnshapeSupportOk() (*bool, bool)`

GetIsOnshapeSupportOk returns a tuple with the IsOnshapeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnshapeSupport

`func (o *BTTaskUserSummaryInfo) SetIsOnshapeSupport(v bool)`

SetIsOnshapeSupport sets IsOnshapeSupport field to given value.

### HasIsOnshapeSupport

`func (o *BTTaskUserSummaryInfo) HasIsOnshapeSupport() bool`

HasIsOnshapeSupport returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *BTTaskUserSummaryInfo) GetLastLoginTime() JSONTime`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *BTTaskUserSummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *BTTaskUserSummaryInfo) SetLastLoginTime(v JSONTime)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *BTTaskUserSummaryInfo) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastName

`func (o *BTTaskUserSummaryInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BTTaskUserSummaryInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BTTaskUserSummaryInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BTTaskUserSummaryInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetName

`func (o *BTTaskUserSummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTaskUserSummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTaskUserSummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTaskUserSummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonalMessageAllowed

`func (o *BTTaskUserSummaryInfo) GetPersonalMessageAllowed() bool`

GetPersonalMessageAllowed returns the PersonalMessageAllowed field if non-nil, zero value otherwise.

### GetPersonalMessageAllowedOk

`func (o *BTTaskUserSummaryInfo) GetPersonalMessageAllowedOk() (*bool, bool)`

GetPersonalMessageAllowedOk returns a tuple with the PersonalMessageAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalMessageAllowed

`func (o *BTTaskUserSummaryInfo) SetPersonalMessageAllowed(v bool)`

SetPersonalMessageAllowed sets PersonalMessageAllowed field to given value.

### HasPersonalMessageAllowed

`func (o *BTTaskUserSummaryInfo) HasPersonalMessageAllowed() bool`

HasPersonalMessageAllowed returns a boolean if a field has been set.

### GetSource

`func (o *BTTaskUserSummaryInfo) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BTTaskUserSummaryInfo) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BTTaskUserSummaryInfo) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *BTTaskUserSummaryInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *BTTaskUserSummaryInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTTaskUserSummaryInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTTaskUserSummaryInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTTaskUserSummaryInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTaskUserSummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTaskUserSummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTaskUserSummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTaskUserSummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


