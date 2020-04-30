# BTIdentityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdentityType** | Pointer to **int32** |  | [optional] 
**Team** | Pointer to [**BTTeamSummaryInfo**](BTTeamSummaryInfo.md) |  | [optional] 
**User** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTIdentityInfo

`func NewBTIdentityInfo() *BTIdentityInfo`

NewBTIdentityInfo instantiates a new BTIdentityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTIdentityInfoWithDefaults

`func NewBTIdentityInfoWithDefaults() *BTIdentityInfo`

NewBTIdentityInfoWithDefaults instantiates a new BTIdentityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTIdentityInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTIdentityInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTIdentityInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTIdentityInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTIdentityInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTIdentityInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTIdentityInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTIdentityInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityType

`func (o *BTIdentityInfo) GetIdentityType() int32`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *BTIdentityInfo) GetIdentityTypeOk() (*int32, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *BTIdentityInfo) SetIdentityType(v int32)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *BTIdentityInfo) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetTeam

`func (o *BTIdentityInfo) GetTeam() BTTeamSummaryInfo`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BTIdentityInfo) GetTeamOk() (*BTTeamSummaryInfo, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BTIdentityInfo) SetTeam(v BTTeamSummaryInfo)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BTIdentityInfo) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetUser

`func (o *BTIdentityInfo) GetUser() BTUserSummaryInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BTIdentityInfo) GetUserOk() (*BTUserSummaryInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BTIdentityInfo) SetUser(v BTUserSummaryInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *BTIdentityInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetViewRef

`func (o *BTIdentityInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTIdentityInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTIdentityInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTIdentityInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


