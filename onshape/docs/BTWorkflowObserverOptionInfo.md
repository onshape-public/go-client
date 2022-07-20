# BTWorkflowObserverOptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to [**BTAliasInfo**](BTAliasInfo.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdentityType** | Pointer to **int32** |  | [optional] 
**Role** | Pointer to [**BTRbacRoleInfo**](BTRbacRoleInfo.md) |  | [optional] 
**Team** | Pointer to [**BTTeamSummaryInfo**](BTTeamSummaryInfo.md) |  | [optional] 
**User** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkflowObserverOptionInfo

`func NewBTWorkflowObserverOptionInfo() *BTWorkflowObserverOptionInfo`

NewBTWorkflowObserverOptionInfo instantiates a new BTWorkflowObserverOptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowObserverOptionInfoWithDefaults

`func NewBTWorkflowObserverOptionInfoWithDefaults() *BTWorkflowObserverOptionInfo`

NewBTWorkflowObserverOptionInfoWithDefaults instantiates a new BTWorkflowObserverOptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *BTWorkflowObserverOptionInfo) GetAlias() BTAliasInfo`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BTWorkflowObserverOptionInfo) GetAliasOk() (*BTAliasInfo, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BTWorkflowObserverOptionInfo) SetAlias(v BTAliasInfo)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *BTWorkflowObserverOptionInfo) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetHref

`func (o *BTWorkflowObserverOptionInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTWorkflowObserverOptionInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTWorkflowObserverOptionInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTWorkflowObserverOptionInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTWorkflowObserverOptionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWorkflowObserverOptionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWorkflowObserverOptionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWorkflowObserverOptionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityType

`func (o *BTWorkflowObserverOptionInfo) GetIdentityType() int32`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *BTWorkflowObserverOptionInfo) GetIdentityTypeOk() (*int32, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *BTWorkflowObserverOptionInfo) SetIdentityType(v int32)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *BTWorkflowObserverOptionInfo) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetRole

`func (o *BTWorkflowObserverOptionInfo) GetRole() BTRbacRoleInfo`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BTWorkflowObserverOptionInfo) GetRoleOk() (*BTRbacRoleInfo, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BTWorkflowObserverOptionInfo) SetRole(v BTRbacRoleInfo)`

SetRole sets Role field to given value.

### HasRole

`func (o *BTWorkflowObserverOptionInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeam

`func (o *BTWorkflowObserverOptionInfo) GetTeam() BTTeamSummaryInfo`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BTWorkflowObserverOptionInfo) GetTeamOk() (*BTTeamSummaryInfo, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BTWorkflowObserverOptionInfo) SetTeam(v BTTeamSummaryInfo)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BTWorkflowObserverOptionInfo) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetUser

`func (o *BTWorkflowObserverOptionInfo) GetUser() BTUserSummaryInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BTWorkflowObserverOptionInfo) GetUserOk() (*BTUserSummaryInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BTWorkflowObserverOptionInfo) SetUser(v BTUserSummaryInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *BTWorkflowObserverOptionInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetViewRef

`func (o *BTWorkflowObserverOptionInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTWorkflowObserverOptionInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTWorkflowObserverOptionInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTWorkflowObserverOptionInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


