# BTGlobalPermissionInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Identities** | Pointer to [**[]BTIdentityInfo**](BTIdentityInfo.md) |  | [optional] 

## Methods

### NewBTGlobalPermissionInfoItem

`func NewBTGlobalPermissionInfoItem() *BTGlobalPermissionInfoItem`

NewBTGlobalPermissionInfoItem instantiates a new BTGlobalPermissionInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTGlobalPermissionInfoItemWithDefaults

`func NewBTGlobalPermissionInfoItemWithDefaults() *BTGlobalPermissionInfoItem`

NewBTGlobalPermissionInfoItemWithDefaults instantiates a new BTGlobalPermissionInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BTGlobalPermissionInfoItem) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BTGlobalPermissionInfoItem) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BTGlobalPermissionInfoItem) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *BTGlobalPermissionInfoItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIdentities

`func (o *BTGlobalPermissionInfoItem) GetIdentities() []BTIdentityInfo`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *BTGlobalPermissionInfoItem) GetIdentitiesOk() (*[]BTIdentityInfo, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *BTGlobalPermissionInfoItem) SetIdentities(v []BTIdentityInfo)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *BTGlobalPermissionInfoItem) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


