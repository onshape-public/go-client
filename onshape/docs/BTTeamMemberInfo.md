# BTTeamMemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Member** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Team** | Pointer to [**BTTeamSummaryInfo**](BTTeamSummaryInfo.md) |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTTeamMemberInfo

`func NewBTTeamMemberInfo() *BTTeamMemberInfo`

NewBTTeamMemberInfo instantiates a new BTTeamMemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTeamMemberInfoWithDefaults

`func NewBTTeamMemberInfoWithDefaults() *BTTeamMemberInfo`

NewBTTeamMemberInfoWithDefaults instantiates a new BTTeamMemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *BTTeamMemberInfo) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *BTTeamMemberInfo) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *BTTeamMemberInfo) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *BTTeamMemberInfo) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetHref

`func (o *BTTeamMemberInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTeamMemberInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTeamMemberInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTeamMemberInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTeamMemberInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTeamMemberInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTeamMemberInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTeamMemberInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *BTTeamMemberInfo) GetMember() BTUserSummaryInfo`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BTTeamMemberInfo) GetMemberOk() (*BTUserSummaryInfo, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BTTeamMemberInfo) SetMember(v BTUserSummaryInfo)`

SetMember sets Member field to given value.

### HasMember

`func (o *BTTeamMemberInfo) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *BTTeamMemberInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTeamMemberInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTeamMemberInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTeamMemberInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTeam

`func (o *BTTeamMemberInfo) GetTeam() BTTeamSummaryInfo`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BTTeamMemberInfo) GetTeamOk() (*BTTeamSummaryInfo, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BTTeamMemberInfo) SetTeam(v BTTeamSummaryInfo)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BTTeamMemberInfo) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTeamMemberInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTeamMemberInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTeamMemberInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTeamMemberInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


