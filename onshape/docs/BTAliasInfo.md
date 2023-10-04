# BTAliasInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]BTAliasEntryInfo**](BTAliasEntryInfo.md) |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Identities** | Pointer to [**[]BTIdentityInfo**](BTIdentityInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTAliasInfo

`func NewBTAliasInfo() *BTAliasInfo`

NewBTAliasInfo instantiates a new BTAliasInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAliasInfoWithDefaults

`func NewBTAliasInfoWithDefaults() *BTAliasInfo`

NewBTAliasInfoWithDefaults instantiates a new BTAliasInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTAliasInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTAliasInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTAliasInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTAliasInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTAliasInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTAliasInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTAliasInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTAliasInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *BTAliasInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAliasInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAliasInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAliasInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *BTAliasInfo) GetEntries() []BTAliasEntryInfo`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BTAliasInfo) GetEntriesOk() (*[]BTAliasEntryInfo, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BTAliasInfo) SetEntries(v []BTAliasEntryInfo)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BTAliasInfo) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetHref

`func (o *BTAliasInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTAliasInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTAliasInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTAliasInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTAliasInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTAliasInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTAliasInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTAliasInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentities

`func (o *BTAliasInfo) GetIdentities() []BTIdentityInfo`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *BTAliasInfo) GetIdentitiesOk() (*[]BTIdentityInfo, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *BTAliasInfo) SetIdentities(v []BTIdentityInfo)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *BTAliasInfo) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetName

`func (o *BTAliasInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAliasInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAliasInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAliasInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViewRef

`func (o *BTAliasInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTAliasInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTAliasInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTAliasInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


