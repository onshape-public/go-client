# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**Role** | Pointer to [**BTRbacRoleInfo**](BTRbacRoleInfo.md) |  | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionSet

`func (o *Entry) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *Entry) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *Entry) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *Entry) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetRole

`func (o *Entry) GetRole() BTRbacRoleInfo`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Entry) GetRoleOk() (*BTRbacRoleInfo, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Entry) SetRole(v BTRbacRoleInfo)`

SetRole sets Role field to given value.

### HasRole

`func (o *Entry) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


