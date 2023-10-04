# BTCompanyUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Indicates the user is an admin if true. | [optional] 
**CompanyId** | Pointer to **string** | Company ID of the user. | [optional] 
**DocumentationNameOverride** | Pointer to **string** | String to override documentation name. | [optional] 
**Email** | Pointer to **string** | Email ID of the company user. | [optional] 
**GlobalPermissions** | Pointer to **[]int32** | List of global permissions to grant. | [optional] 
**Guest** | Pointer to **bool** | Indicates the user is a guest user if true. | [optional] 
**Light** | Pointer to **bool** | Indicates the user is a light user if true. | [optional] 

## Methods

### NewBTCompanyUserParams

`func NewBTCompanyUserParams() *BTCompanyUserParams`

NewBTCompanyUserParams instantiates a new BTCompanyUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCompanyUserParamsWithDefaults

`func NewBTCompanyUserParamsWithDefaults() *BTCompanyUserParams`

NewBTCompanyUserParamsWithDefaults instantiates a new BTCompanyUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *BTCompanyUserParams) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *BTCompanyUserParams) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *BTCompanyUserParams) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *BTCompanyUserParams) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTCompanyUserParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTCompanyUserParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTCompanyUserParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTCompanyUserParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDocumentationNameOverride

`func (o *BTCompanyUserParams) GetDocumentationNameOverride() string`

GetDocumentationNameOverride returns the DocumentationNameOverride field if non-nil, zero value otherwise.

### GetDocumentationNameOverrideOk

`func (o *BTCompanyUserParams) GetDocumentationNameOverrideOk() (*string, bool)`

GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationNameOverride

`func (o *BTCompanyUserParams) SetDocumentationNameOverride(v string)`

SetDocumentationNameOverride sets DocumentationNameOverride field to given value.

### HasDocumentationNameOverride

`func (o *BTCompanyUserParams) HasDocumentationNameOverride() bool`

HasDocumentationNameOverride returns a boolean if a field has been set.

### GetEmail

`func (o *BTCompanyUserParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BTCompanyUserParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BTCompanyUserParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BTCompanyUserParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGlobalPermissions

`func (o *BTCompanyUserParams) GetGlobalPermissions() []int32`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *BTCompanyUserParams) GetGlobalPermissionsOk() (*[]int32, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *BTCompanyUserParams) SetGlobalPermissions(v []int32)`

SetGlobalPermissions sets GlobalPermissions field to given value.

### HasGlobalPermissions

`func (o *BTCompanyUserParams) HasGlobalPermissions() bool`

HasGlobalPermissions returns a boolean if a field has been set.

### GetGuest

`func (o *BTCompanyUserParams) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *BTCompanyUserParams) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *BTCompanyUserParams) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *BTCompanyUserParams) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetLight

`func (o *BTCompanyUserParams) GetLight() bool`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *BTCompanyUserParams) GetLightOk() (*bool, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *BTCompanyUserParams) SetLight(v bool)`

SetLight sets Light field to given value.

### HasLight

`func (o *BTCompanyUserParams) HasLight() bool`

HasLight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


