# BTProjectInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionScheme** | Pointer to [**BTRbacPermissionSchemeInfo**](BTRbacPermissionSchemeInfo.md) |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**PlmContext** | Pointer to [**BTPlmContextInfo**](BTPlmContextInfo.md) |  | [optional] 
**RoleMapEntries** | Pointer to [**[]RoleMapEntry**](RoleMapEntry.md) |  | [optional] 
**Trash** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTProjectInfoAllOf

`func NewBTProjectInfoAllOf() *BTProjectInfoAllOf`

NewBTProjectInfoAllOf instantiates a new BTProjectInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTProjectInfoAllOfWithDefaults

`func NewBTProjectInfoAllOfWithDefaults() *BTProjectInfoAllOf`

NewBTProjectInfoAllOfWithDefaults instantiates a new BTProjectInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionScheme

`func (o *BTProjectInfoAllOf) GetPermissionScheme() BTRbacPermissionSchemeInfo`

GetPermissionScheme returns the PermissionScheme field if non-nil, zero value otherwise.

### GetPermissionSchemeOk

`func (o *BTProjectInfoAllOf) GetPermissionSchemeOk() (*BTRbacPermissionSchemeInfo, bool)`

GetPermissionSchemeOk returns a tuple with the PermissionScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionScheme

`func (o *BTProjectInfoAllOf) SetPermissionScheme(v BTRbacPermissionSchemeInfo)`

SetPermissionScheme sets PermissionScheme field to given value.

### HasPermissionScheme

`func (o *BTProjectInfoAllOf) HasPermissionScheme() bool`

HasPermissionScheme returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTProjectInfoAllOf) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTProjectInfoAllOf) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTProjectInfoAllOf) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTProjectInfoAllOf) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetPlmContext

`func (o *BTProjectInfoAllOf) GetPlmContext() BTPlmContextInfo`

GetPlmContext returns the PlmContext field if non-nil, zero value otherwise.

### GetPlmContextOk

`func (o *BTProjectInfoAllOf) GetPlmContextOk() (*BTPlmContextInfo, bool)`

GetPlmContextOk returns a tuple with the PlmContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmContext

`func (o *BTProjectInfoAllOf) SetPlmContext(v BTPlmContextInfo)`

SetPlmContext sets PlmContext field to given value.

### HasPlmContext

`func (o *BTProjectInfoAllOf) HasPlmContext() bool`

HasPlmContext returns a boolean if a field has been set.

### GetRoleMapEntries

`func (o *BTProjectInfoAllOf) GetRoleMapEntries() []RoleMapEntry`

GetRoleMapEntries returns the RoleMapEntries field if non-nil, zero value otherwise.

### GetRoleMapEntriesOk

`func (o *BTProjectInfoAllOf) GetRoleMapEntriesOk() (*[]RoleMapEntry, bool)`

GetRoleMapEntriesOk returns a tuple with the RoleMapEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapEntries

`func (o *BTProjectInfoAllOf) SetRoleMapEntries(v []RoleMapEntry)`

SetRoleMapEntries sets RoleMapEntries field to given value.

### HasRoleMapEntries

`func (o *BTProjectInfoAllOf) HasRoleMapEntries() bool`

HasRoleMapEntries returns a boolean if a field has been set.

### GetTrash

`func (o *BTProjectInfoAllOf) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTProjectInfoAllOf) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTProjectInfoAllOf) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTProjectInfoAllOf) HasTrash() bool`

HasTrash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


