# BTAliasEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Identity** | Pointer to [**BTIdentityInfo**](BTIdentityInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTAliasEntryInfo

`func NewBTAliasEntryInfo() *BTAliasEntryInfo`

NewBTAliasEntryInfo instantiates a new BTAliasEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAliasEntryInfoWithDefaults

`func NewBTAliasEntryInfoWithDefaults() *BTAliasEntryInfo`

NewBTAliasEntryInfoWithDefaults instantiates a new BTAliasEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BTAliasEntryInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTAliasEntryInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTAliasEntryInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTAliasEntryInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *BTAliasEntryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTAliasEntryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTAliasEntryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTAliasEntryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTAliasEntryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTAliasEntryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTAliasEntryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTAliasEntryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentity

`func (o *BTAliasEntryInfo) GetIdentity() BTIdentityInfo`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *BTAliasEntryInfo) GetIdentityOk() (*BTIdentityInfo, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *BTAliasEntryInfo) SetIdentity(v BTIdentityInfo)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *BTAliasEntryInfo) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *BTAliasEntryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAliasEntryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAliasEntryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAliasEntryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViewRef

`func (o *BTAliasEntryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTAliasEntryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTAliasEntryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTAliasEntryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


