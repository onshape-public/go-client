# BTUserOAuth2SummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**CompanyPlan** | Pointer to **bool** |  | [optional] 
**Oauth2Scopes** | Pointer to **int64** |  | [optional] 
**PlanGroup** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **int32** |  | [optional] 
**Roles** | Pointer to [**[]BTRole**](BTRole.md) |  | [optional] 

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

`func (o *BTUserOAuth2SummaryInfo) GetRoles() []BTRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BTUserOAuth2SummaryInfo) GetRolesOk() (*[]BTRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BTUserOAuth2SummaryInfo) SetRoles(v []BTRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BTUserOAuth2SummaryInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


