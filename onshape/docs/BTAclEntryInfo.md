# BTAclEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptOwnerTransfer** | Pointer to **bool** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**ConnectionName** | Pointer to **string** |  | [optional] 
**ConnectionUser** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EnterpriseMember** | Pointer to **bool** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**EntryState** | Pointer to [**BTUserState**](BTUserState.md) |  | [optional] 
**EntryType** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**PendingOwnerTransfer** | Pointer to **bool** |  | [optional] 
**Permission** | Pointer to **int64** |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**ProCompanySubtype** | Pointer to **int32** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAclEntryInfo

`func NewBTAclEntryInfo() *BTAclEntryInfo`

NewBTAclEntryInfo instantiates a new BTAclEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAclEntryInfoWithDefaults

`func NewBTAclEntryInfoWithDefaults() *BTAclEntryInfo`

NewBTAclEntryInfoWithDefaults instantiates a new BTAclEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptOwnerTransfer

`func (o *BTAclEntryInfo) GetAcceptOwnerTransfer() bool`

GetAcceptOwnerTransfer returns the AcceptOwnerTransfer field if non-nil, zero value otherwise.

### GetAcceptOwnerTransferOk

`func (o *BTAclEntryInfo) GetAcceptOwnerTransferOk() (*bool, bool)`

GetAcceptOwnerTransferOk returns a tuple with the AcceptOwnerTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptOwnerTransfer

`func (o *BTAclEntryInfo) SetAcceptOwnerTransfer(v bool)`

SetAcceptOwnerTransfer sets AcceptOwnerTransfer field to given value.

### HasAcceptOwnerTransfer

`func (o *BTAclEntryInfo) HasAcceptOwnerTransfer() bool`

HasAcceptOwnerTransfer returns a boolean if a field has been set.

### GetCompanyName

`func (o *BTAclEntryInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BTAclEntryInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BTAclEntryInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BTAclEntryInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetConnectionId

`func (o *BTAclEntryInfo) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BTAclEntryInfo) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BTAclEntryInfo) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *BTAclEntryInfo) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionName

`func (o *BTAclEntryInfo) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *BTAclEntryInfo) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *BTAclEntryInfo) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *BTAclEntryInfo) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetConnectionUser

`func (o *BTAclEntryInfo) GetConnectionUser() bool`

GetConnectionUser returns the ConnectionUser field if non-nil, zero value otherwise.

### GetConnectionUserOk

`func (o *BTAclEntryInfo) GetConnectionUserOk() (*bool, bool)`

GetConnectionUserOk returns a tuple with the ConnectionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUser

`func (o *BTAclEntryInfo) SetConnectionUser(v bool)`

SetConnectionUser sets ConnectionUser field to given value.

### HasConnectionUser

`func (o *BTAclEntryInfo) HasConnectionUser() bool`

HasConnectionUser returns a boolean if a field has been set.

### GetEmail

`func (o *BTAclEntryInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BTAclEntryInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BTAclEntryInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BTAclEntryInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnterpriseMember

`func (o *BTAclEntryInfo) GetEnterpriseMember() bool`

GetEnterpriseMember returns the EnterpriseMember field if non-nil, zero value otherwise.

### GetEnterpriseMemberOk

`func (o *BTAclEntryInfo) GetEnterpriseMemberOk() (*bool, bool)`

GetEnterpriseMemberOk returns a tuple with the EnterpriseMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseMember

`func (o *BTAclEntryInfo) SetEnterpriseMember(v bool)`

SetEnterpriseMember sets EnterpriseMember field to given value.

### HasEnterpriseMember

`func (o *BTAclEntryInfo) HasEnterpriseMember() bool`

HasEnterpriseMember returns a boolean if a field has been set.

### GetEntryId

`func (o *BTAclEntryInfo) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BTAclEntryInfo) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BTAclEntryInfo) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *BTAclEntryInfo) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEntryState

`func (o *BTAclEntryInfo) GetEntryState() BTUserState`

GetEntryState returns the EntryState field if non-nil, zero value otherwise.

### GetEntryStateOk

`func (o *BTAclEntryInfo) GetEntryStateOk() (*BTUserState, bool)`

GetEntryStateOk returns a tuple with the EntryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryState

`func (o *BTAclEntryInfo) SetEntryState(v BTUserState)`

SetEntryState sets EntryState field to given value.

### HasEntryState

`func (o *BTAclEntryInfo) HasEntryState() bool`

HasEntryState returns a boolean if a field has been set.

### GetEntryType

`func (o *BTAclEntryInfo) GetEntryType() int32`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BTAclEntryInfo) GetEntryTypeOk() (*int32, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BTAclEntryInfo) SetEntryType(v int32)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *BTAclEntryInfo) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetImage

`func (o *BTAclEntryInfo) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BTAclEntryInfo) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BTAclEntryInfo) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BTAclEntryInfo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *BTAclEntryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAclEntryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAclEntryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAclEntryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *BTAclEntryInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTAclEntryInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTAclEntryInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTAclEntryInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPendingOwnerTransfer

`func (o *BTAclEntryInfo) GetPendingOwnerTransfer() bool`

GetPendingOwnerTransfer returns the PendingOwnerTransfer field if non-nil, zero value otherwise.

### GetPendingOwnerTransferOk

`func (o *BTAclEntryInfo) GetPendingOwnerTransferOk() (*bool, bool)`

GetPendingOwnerTransferOk returns a tuple with the PendingOwnerTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOwnerTransfer

`func (o *BTAclEntryInfo) SetPendingOwnerTransfer(v bool)`

SetPendingOwnerTransfer sets PendingOwnerTransfer field to given value.

### HasPendingOwnerTransfer

`func (o *BTAclEntryInfo) HasPendingOwnerTransfer() bool`

HasPendingOwnerTransfer returns a boolean if a field has been set.

### GetPermission

`func (o *BTAclEntryInfo) GetPermission() int64`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTAclEntryInfo) GetPermissionOk() (*int64, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTAclEntryInfo) SetPermission(v int64)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTAclEntryInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTAclEntryInfo) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTAclEntryInfo) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTAclEntryInfo) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTAclEntryInfo) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetProCompanySubtype

`func (o *BTAclEntryInfo) GetProCompanySubtype() int32`

GetProCompanySubtype returns the ProCompanySubtype field if non-nil, zero value otherwise.

### GetProCompanySubtypeOk

`func (o *BTAclEntryInfo) GetProCompanySubtypeOk() (*int32, bool)`

GetProCompanySubtypeOk returns a tuple with the ProCompanySubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProCompanySubtype

`func (o *BTAclEntryInfo) SetProCompanySubtype(v int32)`

SetProCompanySubtype sets ProCompanySubtype field to given value.

### HasProCompanySubtype

`func (o *BTAclEntryInfo) HasProCompanySubtype() bool`

HasProCompanySubtype returns a boolean if a field has been set.

### GetTeamName

`func (o *BTAclEntryInfo) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *BTAclEntryInfo) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *BTAclEntryInfo) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *BTAclEntryInfo) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


