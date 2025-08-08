# BTExternalConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**ContactUser** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**InvitedCompany** | Pointer to [**BTCompanySummaryInfo**](BTCompanySummaryInfo.md) |  | [optional] 
**IsOwnerEnterpriseEdu** | Pointer to **bool** |  | [optional] 
**Member** | Pointer to **bool** |  | [optional] 
**NumberOfMembers** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTExternalConnectionInfo

`func NewBTExternalConnectionInfo() *BTExternalConnectionInfo`

NewBTExternalConnectionInfo instantiates a new BTExternalConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTExternalConnectionInfoWithDefaults

`func NewBTExternalConnectionInfoWithDefaults() *BTExternalConnectionInfo`

NewBTExternalConnectionInfoWithDefaults instantiates a new BTExternalConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedBy

`func (o *BTExternalConnectionInfo) GetApprovedBy() BTUserSummaryInfo`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *BTExternalConnectionInfo) GetApprovedByOk() (*BTUserSummaryInfo, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *BTExternalConnectionInfo) SetApprovedBy(v BTUserSummaryInfo)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *BTExternalConnectionInfo) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetContactUser

`func (o *BTExternalConnectionInfo) GetContactUser() BTUserSummaryInfo`

GetContactUser returns the ContactUser field if non-nil, zero value otherwise.

### GetContactUserOk

`func (o *BTExternalConnectionInfo) GetContactUserOk() (*BTUserSummaryInfo, bool)`

GetContactUserOk returns a tuple with the ContactUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactUser

`func (o *BTExternalConnectionInfo) SetContactUser(v BTUserSummaryInfo)`

SetContactUser sets ContactUser field to given value.

### HasContactUser

`func (o *BTExternalConnectionInfo) HasContactUser() bool`

HasContactUser returns a boolean if a field has been set.

### GetIcon

`func (o *BTExternalConnectionInfo) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BTExternalConnectionInfo) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BTExternalConnectionInfo) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *BTExternalConnectionInfo) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetInvitedCompany

`func (o *BTExternalConnectionInfo) GetInvitedCompany() BTCompanySummaryInfo`

GetInvitedCompany returns the InvitedCompany field if non-nil, zero value otherwise.

### GetInvitedCompanyOk

`func (o *BTExternalConnectionInfo) GetInvitedCompanyOk() (*BTCompanySummaryInfo, bool)`

GetInvitedCompanyOk returns a tuple with the InvitedCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedCompany

`func (o *BTExternalConnectionInfo) SetInvitedCompany(v BTCompanySummaryInfo)`

SetInvitedCompany sets InvitedCompany field to given value.

### HasInvitedCompany

`func (o *BTExternalConnectionInfo) HasInvitedCompany() bool`

HasInvitedCompany returns a boolean if a field has been set.

### GetIsOwnerEnterpriseEdu

`func (o *BTExternalConnectionInfo) GetIsOwnerEnterpriseEdu() bool`

GetIsOwnerEnterpriseEdu returns the IsOwnerEnterpriseEdu field if non-nil, zero value otherwise.

### GetIsOwnerEnterpriseEduOk

`func (o *BTExternalConnectionInfo) GetIsOwnerEnterpriseEduOk() (*bool, bool)`

GetIsOwnerEnterpriseEduOk returns a tuple with the IsOwnerEnterpriseEdu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwnerEnterpriseEdu

`func (o *BTExternalConnectionInfo) SetIsOwnerEnterpriseEdu(v bool)`

SetIsOwnerEnterpriseEdu sets IsOwnerEnterpriseEdu field to given value.

### HasIsOwnerEnterpriseEdu

`func (o *BTExternalConnectionInfo) HasIsOwnerEnterpriseEdu() bool`

HasIsOwnerEnterpriseEdu returns a boolean if a field has been set.

### GetMember

`func (o *BTExternalConnectionInfo) GetMember() bool`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BTExternalConnectionInfo) GetMemberOk() (*bool, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BTExternalConnectionInfo) SetMember(v bool)`

SetMember sets Member field to given value.

### HasMember

`func (o *BTExternalConnectionInfo) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNumberOfMembers

`func (o *BTExternalConnectionInfo) GetNumberOfMembers() int64`

GetNumberOfMembers returns the NumberOfMembers field if non-nil, zero value otherwise.

### GetNumberOfMembersOk

`func (o *BTExternalConnectionInfo) GetNumberOfMembersOk() (*int64, bool)`

GetNumberOfMembersOk returns a tuple with the NumberOfMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMembers

`func (o *BTExternalConnectionInfo) SetNumberOfMembers(v int64)`

SetNumberOfMembers sets NumberOfMembers field to given value.

### HasNumberOfMembers

`func (o *BTExternalConnectionInfo) HasNumberOfMembers() bool`

HasNumberOfMembers returns a boolean if a field has been set.

### GetState

`func (o *BTExternalConnectionInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTExternalConnectionInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTExternalConnectionInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTExternalConnectionInfo) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


