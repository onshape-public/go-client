# BTInheritedAclInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]BTAclEntryInfo**](BTAclEntryInfo.md) |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectName** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**SharedWithSupport** | Pointer to **bool** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewBTInheritedAclInfo

`func NewBTInheritedAclInfo() *BTInheritedAclInfo`

NewBTInheritedAclInfo instantiates a new BTInheritedAclInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInheritedAclInfoWithDefaults

`func NewBTInheritedAclInfoWithDefaults() *BTInheritedAclInfo`

NewBTInheritedAclInfoWithDefaults instantiates a new BTInheritedAclInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *BTInheritedAclInfo) GetEntries() []BTAclEntryInfo`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BTInheritedAclInfo) GetEntriesOk() (*[]BTAclEntryInfo, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BTInheritedAclInfo) SetEntries(v []BTAclEntryInfo)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BTInheritedAclInfo) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetHref

`func (o *BTInheritedAclInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTInheritedAclInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTInheritedAclInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTInheritedAclInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTInheritedAclInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTInheritedAclInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTInheritedAclInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTInheritedAclInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTInheritedAclInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTInheritedAclInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTInheritedAclInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTInheritedAclInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *BTInheritedAclInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTInheritedAclInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTInheritedAclInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTInheritedAclInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectName

`func (o *BTInheritedAclInfo) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *BTInheritedAclInfo) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *BTInheritedAclInfo) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *BTInheritedAclInfo) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *BTInheritedAclInfo) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTInheritedAclInfo) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTInheritedAclInfo) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTInheritedAclInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOwner

`func (o *BTInheritedAclInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTInheritedAclInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTInheritedAclInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTInheritedAclInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublic

`func (o *BTInheritedAclInfo) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BTInheritedAclInfo) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BTInheritedAclInfo) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *BTInheritedAclInfo) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSharedWithSupport

`func (o *BTInheritedAclInfo) GetSharedWithSupport() bool`

GetSharedWithSupport returns the SharedWithSupport field if non-nil, zero value otherwise.

### GetSharedWithSupportOk

`func (o *BTInheritedAclInfo) GetSharedWithSupportOk() (*bool, bool)`

GetSharedWithSupportOk returns a tuple with the SharedWithSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithSupport

`func (o *BTInheritedAclInfo) SetSharedWithSupport(v bool)`

SetSharedWithSupport sets SharedWithSupport field to given value.

### HasSharedWithSupport

`func (o *BTInheritedAclInfo) HasSharedWithSupport() bool`

HasSharedWithSupport returns a boolean if a field has been set.

### GetViewRef

`func (o *BTInheritedAclInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTInheritedAclInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTInheritedAclInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTInheritedAclInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetVisibility

`func (o *BTInheritedAclInfo) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BTInheritedAclInfo) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BTInheritedAclInfo) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BTInheritedAclInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


