# RoleMapEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identities** | Pointer to [**[]BTIdentityInfo**](BTIdentityInfo.md) |  | [optional] 
**Role** | Pointer to [**BTRbacRoleInfo**](BTRbacRoleInfo.md) |  | [optional] 

## Methods

### NewRoleMapEntry

`func NewRoleMapEntry() *RoleMapEntry`

NewRoleMapEntry instantiates a new RoleMapEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMapEntryWithDefaults

`func NewRoleMapEntryWithDefaults() *RoleMapEntry`

NewRoleMapEntryWithDefaults instantiates a new RoleMapEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentities

`func (o *RoleMapEntry) GetIdentities() []BTIdentityInfo`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *RoleMapEntry) GetIdentitiesOk() (*[]BTIdentityInfo, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *RoleMapEntry) SetIdentities(v []BTIdentityInfo)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *RoleMapEntry) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetRole

`func (o *RoleMapEntry) GetRole() BTRbacRoleInfo`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleMapEntry) GetRoleOk() (*BTRbacRoleInfo, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleMapEntry) SetRole(v BTRbacRoleInfo)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleMapEntry) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


