# BTProjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionScheme** | Pointer to [**BTRbacPermissionSchemeInfo**](BTRbacPermissionSchemeInfo.md) |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**PlmContext** | Pointer to [**BTPlmContextInfo**](BTPlmContextInfo.md) |  | [optional] 
**RoleMapEntries** | Pointer to [**[]RoleMapEntry**](RoleMapEntry.md) |  | [optional] 
**Trash** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTProjectInfo

`func NewBTProjectInfo() *BTProjectInfo`

NewBTProjectInfo instantiates a new BTProjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTProjectInfoWithDefaults

`func NewBTProjectInfoWithDefaults() *BTProjectInfo`

NewBTProjectInfoWithDefaults instantiates a new BTProjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionScheme

`func (o *BTProjectInfo) GetPermissionScheme() BTRbacPermissionSchemeInfo`

GetPermissionScheme returns the PermissionScheme field if non-nil, zero value otherwise.

### GetPermissionSchemeOk

`func (o *BTProjectInfo) GetPermissionSchemeOk() (*BTRbacPermissionSchemeInfo, bool)`

GetPermissionSchemeOk returns a tuple with the PermissionScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionScheme

`func (o *BTProjectInfo) SetPermissionScheme(v BTRbacPermissionSchemeInfo)`

SetPermissionScheme sets PermissionScheme field to given value.

### HasPermissionScheme

`func (o *BTProjectInfo) HasPermissionScheme() bool`

HasPermissionScheme returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTProjectInfo) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTProjectInfo) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTProjectInfo) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTProjectInfo) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetPlmContext

`func (o *BTProjectInfo) GetPlmContext() BTPlmContextInfo`

GetPlmContext returns the PlmContext field if non-nil, zero value otherwise.

### GetPlmContextOk

`func (o *BTProjectInfo) GetPlmContextOk() (*BTPlmContextInfo, bool)`

GetPlmContextOk returns a tuple with the PlmContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmContext

`func (o *BTProjectInfo) SetPlmContext(v BTPlmContextInfo)`

SetPlmContext sets PlmContext field to given value.

### HasPlmContext

`func (o *BTProjectInfo) HasPlmContext() bool`

HasPlmContext returns a boolean if a field has been set.

### GetRoleMapEntries

`func (o *BTProjectInfo) GetRoleMapEntries() []RoleMapEntry`

GetRoleMapEntries returns the RoleMapEntries field if non-nil, zero value otherwise.

### GetRoleMapEntriesOk

`func (o *BTProjectInfo) GetRoleMapEntriesOk() (*[]RoleMapEntry, bool)`

GetRoleMapEntriesOk returns a tuple with the RoleMapEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapEntries

`func (o *BTProjectInfo) SetRoleMapEntries(v []RoleMapEntry)`

SetRoleMapEntries sets RoleMapEntries field to given value.

### HasRoleMapEntries

`func (o *BTProjectInfo) HasRoleMapEntries() bool`

HasRoleMapEntries returns a boolean if a field has been set.

### GetTrash

`func (o *BTProjectInfo) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *BTProjectInfo) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *BTProjectInfo) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *BTProjectInfo) HasTrash() bool`

HasTrash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


