# BTAclInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** |  | [optional] 
**Entries** | Pointer to [**[]BTAclEntryInfo**](BTAclEntryInfo.md) |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**InheritedAcls** | Pointer to [**[]BTInheritedAclInfo**](BTInheritedAclInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**SharedWithSupport** | Pointer to **bool** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAclInfo

`func NewBTAclInfo() *BTAclInfo`

NewBTAclInfo instantiates a new BTAclInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAclInfoWithDefaults

`func NewBTAclInfoWithDefaults() *BTAclInfo`

NewBTAclInfoWithDefaults instantiates a new BTAclInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *BTAclInfo) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *BTAclInfo) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *BTAclInfo) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *BTAclInfo) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetEntries

`func (o *BTAclInfo) GetEntries() []BTAclEntryInfo`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BTAclInfo) GetEntriesOk() (*[]BTAclEntryInfo, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BTAclInfo) SetEntries(v []BTAclEntryInfo)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BTAclInfo) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetHref

`func (o *BTAclInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTAclInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTAclInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTAclInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTAclInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTAclInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTAclInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTAclInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritedAcls

`func (o *BTAclInfo) GetInheritedAcls() []BTInheritedAclInfo`

GetInheritedAcls returns the InheritedAcls field if non-nil, zero value otherwise.

### GetInheritedAclsOk

`func (o *BTAclInfo) GetInheritedAclsOk() (*[]BTInheritedAclInfo, bool)`

GetInheritedAclsOk returns a tuple with the InheritedAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedAcls

`func (o *BTAclInfo) SetInheritedAcls(v []BTInheritedAclInfo)`

SetInheritedAcls sets InheritedAcls field to given value.

### HasInheritedAcls

`func (o *BTAclInfo) HasInheritedAcls() bool`

HasInheritedAcls returns a boolean if a field has been set.

### GetName

`func (o *BTAclInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAclInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAclInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAclInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *BTAclInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTAclInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTAclInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTAclInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *BTAclInfo) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTAclInfo) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTAclInfo) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTAclInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOwner

`func (o *BTAclInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTAclInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTAclInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTAclInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublic

`func (o *BTAclInfo) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTAclInfo) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTAclInfo) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTAclInfo) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSharedWithSupport

`func (o *BTAclInfo) GetSharedWithSupport() bool`

GetSharedWithSupport returns the SharedWithSupport field if non-nil, zero value otherwise.

### GetSharedWithSupportOk

`func (o *BTAclInfo) GetSharedWithSupportOk() (*bool, bool)`

GetSharedWithSupportOk returns a tuple with the SharedWithSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithSupport

`func (o *BTAclInfo) SetSharedWithSupport(v bool)`

SetSharedWithSupport sets SharedWithSupport field to given value.

### HasSharedWithSupport

`func (o *BTAclInfo) HasSharedWithSupport() bool`

HasSharedWithSupport returns a boolean if a field has been set.

### GetViewRef

`func (o *BTAclInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTAclInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTAclInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTAclInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetVisibility

`func (o *BTAclInfo) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BTAclInfo) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BTAclInfo) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BTAclInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


