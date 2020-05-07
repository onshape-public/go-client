# BTUserOAuth2SummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**BTCompanySummaryInfo**](BTCompanySummaryInfo.md) |  | [optional] 
**CompanyPlan** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**GlobalPermissions** | Pointer to [**GlobalPermissionInfo**](GlobalPermissionInfo.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**IsGuest** | Pointer to **bool** |  | [optional] 
**IsLight** | Pointer to **bool** |  | [optional] 
**LastLoginTime** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oauth2Scopes** | Pointer to **int64** |  | [optional] 
**PlanGroup** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **int32** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTUserOAuth2SummaryInfo

`func NewBTUserOAuth2SummaryInfo() *BTUserOAuth2SummaryInfo`

NewBTUserOAuth2SummaryInfo instantiates a new BTUserOAuth2SummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserOAuth2SummaryInfoWithDefaults

`func NewBTUserOAuth2SummaryInfoWithDefaults() *BTUserOAuth2SummaryInfo`

NewBTUserOAuth2SummaryInfoWithDefaults instantiates a new BTUserOAuth2SummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BTUserOAuth2SummaryInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BTUserOAuth2SummaryInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BTUserOAuth2SummaryInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BTUserOAuth2SummaryInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCompany

`func (o *BTUserOAuth2SummaryInfo) GetCompany() BTCompanySummaryInfo`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BTUserOAuth2SummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BTUserOAuth2SummaryInfo) SetCompany(v BTCompanySummaryInfo)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BTUserOAuth2SummaryInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyPlan

`func (o *BTUserOAuth2SummaryInfo) GetCompanyPlan() bool`

GetCompanyPlan returns the CompanyPlan field if non-nil, zero value otherwise.

### GetCompanyPlanOk

`func (o *BTUserOAuth2SummaryInfo) GetCompanyPlanOk() (*bool, bool)`

GetCompanyPlanOk returns a tuple with the CompanyPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPlan

`func (o *BTUserOAuth2SummaryInfo) SetCompanyPlan(v bool)`

SetCompanyPlan sets CompanyPlan field to given value.

### HasCompanyPlan

`func (o *BTUserOAuth2SummaryInfo) HasCompanyPlan() bool`

HasCompanyPlan returns a boolean if a field has been set.

### GetEmail

`func (o *BTUserOAuth2SummaryInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BTUserOAuth2SummaryInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BTUserOAuth2SummaryInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BTUserOAuth2SummaryInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *BTUserOAuth2SummaryInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BTUserOAuth2SummaryInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BTUserOAuth2SummaryInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BTUserOAuth2SummaryInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGlobalPermissions

`func (o *BTUserOAuth2SummaryInfo) GetGlobalPermissions() GlobalPermissionInfo`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *BTUserOAuth2SummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *BTUserOAuth2SummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo)`

SetGlobalPermissions sets GlobalPermissions field to given value.

### HasGlobalPermissions

`func (o *BTUserOAuth2SummaryInfo) HasGlobalPermissions() bool`

HasGlobalPermissions returns a boolean if a field has been set.

### GetHref

`func (o *BTUserOAuth2SummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTUserOAuth2SummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTUserOAuth2SummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTUserOAuth2SummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTUserOAuth2SummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTUserOAuth2SummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTUserOAuth2SummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTUserOAuth2SummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *BTUserOAuth2SummaryInfo) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BTUserOAuth2SummaryInfo) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BTUserOAuth2SummaryInfo) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BTUserOAuth2SummaryInfo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIsGuest

`func (o *BTUserOAuth2SummaryInfo) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *BTUserOAuth2SummaryInfo) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *BTUserOAuth2SummaryInfo) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *BTUserOAuth2SummaryInfo) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### GetIsLight

`func (o *BTUserOAuth2SummaryInfo) GetIsLight() bool`

GetIsLight returns the IsLight field if non-nil, zero value otherwise.

### GetIsLightOk

`func (o *BTUserOAuth2SummaryInfo) GetIsLightOk() (*bool, bool)`

GetIsLightOk returns a tuple with the IsLight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLight

`func (o *BTUserOAuth2SummaryInfo) SetIsLight(v bool)`

SetIsLight sets IsLight field to given value.

### HasIsLight

`func (o *BTUserOAuth2SummaryInfo) HasIsLight() bool`

HasIsLight returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *BTUserOAuth2SummaryInfo) GetLastLoginTime() JSONTime`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *BTUserOAuth2SummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *BTUserOAuth2SummaryInfo) SetLastLoginTime(v JSONTime)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *BTUserOAuth2SummaryInfo) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastName

`func (o *BTUserOAuth2SummaryInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BTUserOAuth2SummaryInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BTUserOAuth2SummaryInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BTUserOAuth2SummaryInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetName

`func (o *BTUserOAuth2SummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTUserOAuth2SummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTUserOAuth2SummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTUserOAuth2SummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOauth2Scopes

`func (o *BTUserOAuth2SummaryInfo) GetOauth2Scopes() int64`

GetOauth2Scopes returns the Oauth2Scopes field if non-nil, zero value otherwise.

### GetOauth2ScopesOk

`func (o *BTUserOAuth2SummaryInfo) GetOauth2ScopesOk() (*int64, bool)`

GetOauth2ScopesOk returns a tuple with the Oauth2Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Scopes

`func (o *BTUserOAuth2SummaryInfo) SetOauth2Scopes(v int64)`

SetOauth2Scopes sets Oauth2Scopes field to given value.

### HasOauth2Scopes

`func (o *BTUserOAuth2SummaryInfo) HasOauth2Scopes() bool`

HasOauth2Scopes returns a boolean if a field has been set.

### GetPlanGroup

`func (o *BTUserOAuth2SummaryInfo) GetPlanGroup() string`

GetPlanGroup returns the PlanGroup field if non-nil, zero value otherwise.

### GetPlanGroupOk

`func (o *BTUserOAuth2SummaryInfo) GetPlanGroupOk() (*string, bool)`

GetPlanGroupOk returns a tuple with the PlanGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanGroup

`func (o *BTUserOAuth2SummaryInfo) SetPlanGroup(v string)`

SetPlanGroup sets PlanGroup field to given value.

### HasPlanGroup

`func (o *BTUserOAuth2SummaryInfo) HasPlanGroup() bool`

HasPlanGroup returns a boolean if a field has been set.

### GetRole

`func (o *BTUserOAuth2SummaryInfo) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BTUserOAuth2SummaryInfo) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BTUserOAuth2SummaryInfo) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *BTUserOAuth2SummaryInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *BTUserOAuth2SummaryInfo) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BTUserOAuth2SummaryInfo) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BTUserOAuth2SummaryInfo) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BTUserOAuth2SummaryInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSource

`func (o *BTUserOAuth2SummaryInfo) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BTUserOAuth2SummaryInfo) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BTUserOAuth2SummaryInfo) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *BTUserOAuth2SummaryInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *BTUserOAuth2SummaryInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTUserOAuth2SummaryInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTUserOAuth2SummaryInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTUserOAuth2SummaryInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetViewRef

`func (o *BTUserOAuth2SummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTUserOAuth2SummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTUserOAuth2SummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTUserOAuth2SummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


