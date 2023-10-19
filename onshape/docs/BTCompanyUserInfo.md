# BTCompanyUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to [**BTCompanySummaryInfo**](BTCompanySummaryInfo.md) |  | [optional] 
**DocumentationNameOverride** | Pointer to **string** |  | [optional] 
**Guest** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**LastLoginTime** | Pointer to **JSONTime** |  | [optional] 
**Light** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**User** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTCompanyUserInfo

`func NewBTCompanyUserInfo() *BTCompanyUserInfo`

NewBTCompanyUserInfo instantiates a new BTCompanyUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCompanyUserInfoWithDefaults

`func NewBTCompanyUserInfoWithDefaults() *BTCompanyUserInfo`

NewBTCompanyUserInfoWithDefaults instantiates a new BTCompanyUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *BTCompanyUserInfo) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *BTCompanyUserInfo) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *BTCompanyUserInfo) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *BTCompanyUserInfo) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCompany

`func (o *BTCompanyUserInfo) GetCompany() BTCompanySummaryInfo`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BTCompanyUserInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BTCompanyUserInfo) SetCompany(v BTCompanySummaryInfo)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BTCompanyUserInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDocumentationNameOverride

`func (o *BTCompanyUserInfo) GetDocumentationNameOverride() string`

GetDocumentationNameOverride returns the DocumentationNameOverride field if non-nil, zero value otherwise.

### GetDocumentationNameOverrideOk

`func (o *BTCompanyUserInfo) GetDocumentationNameOverrideOk() (*string, bool)`

GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationNameOverride

`func (o *BTCompanyUserInfo) SetDocumentationNameOverride(v string)`

SetDocumentationNameOverride sets DocumentationNameOverride field to given value.

### HasDocumentationNameOverride

`func (o *BTCompanyUserInfo) HasDocumentationNameOverride() bool`

HasDocumentationNameOverride returns a boolean if a field has been set.

### GetGuest

`func (o *BTCompanyUserInfo) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *BTCompanyUserInfo) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *BTCompanyUserInfo) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *BTCompanyUserInfo) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHref

`func (o *BTCompanyUserInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTCompanyUserInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTCompanyUserInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTCompanyUserInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTCompanyUserInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTCompanyUserInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTCompanyUserInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTCompanyUserInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *BTCompanyUserInfo) GetLastLoginTime() JSONTime`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *BTCompanyUserInfo) GetLastLoginTimeOk() (*JSONTime, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *BTCompanyUserInfo) SetLastLoginTime(v JSONTime)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *BTCompanyUserInfo) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLight

`func (o *BTCompanyUserInfo) GetLight() bool`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *BTCompanyUserInfo) GetLightOk() (*bool, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *BTCompanyUserInfo) SetLight(v bool)`

SetLight sets Light field to given value.

### HasLight

`func (o *BTCompanyUserInfo) HasLight() bool`

HasLight returns a boolean if a field has been set.

### GetName

`func (o *BTCompanyUserInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTCompanyUserInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTCompanyUserInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTCompanyUserInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *BTCompanyUserInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTCompanyUserInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTCompanyUserInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTCompanyUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUser

`func (o *BTCompanyUserInfo) GetUser() BTUserBasicSummaryInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BTCompanyUserInfo) GetUserOk() (*BTUserBasicSummaryInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BTCompanyUserInfo) SetUser(v BTUserBasicSummaryInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *BTCompanyUserInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetViewRef

`func (o *BTCompanyUserInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTCompanyUserInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTCompanyUserInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTCompanyUserInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


