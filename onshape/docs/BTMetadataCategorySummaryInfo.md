# BTMetadataCategorySummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultObjectType** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ObjectTypes** | Pointer to **[]int32** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**PublishState** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTMetadataCategorySummaryInfo

`func NewBTMetadataCategorySummaryInfo() *BTMetadataCategorySummaryInfo`

NewBTMetadataCategorySummaryInfo instantiates a new BTMetadataCategorySummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataCategorySummaryInfoWithDefaults

`func NewBTMetadataCategorySummaryInfoWithDefaults() *BTMetadataCategorySummaryInfo`

NewBTMetadataCategorySummaryInfoWithDefaults instantiates a new BTMetadataCategorySummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultObjectType

`func (o *BTMetadataCategorySummaryInfo) GetDefaultObjectType() int32`

GetDefaultObjectType returns the DefaultObjectType field if non-nil, zero value otherwise.

### GetDefaultObjectTypeOk

`func (o *BTMetadataCategorySummaryInfo) GetDefaultObjectTypeOk() (*int32, bool)`

GetDefaultObjectTypeOk returns a tuple with the DefaultObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultObjectType

`func (o *BTMetadataCategorySummaryInfo) SetDefaultObjectType(v int32)`

SetDefaultObjectType sets DefaultObjectType field to given value.

### HasDefaultObjectType

`func (o *BTMetadataCategorySummaryInfo) HasDefaultObjectType() bool`

HasDefaultObjectType returns a boolean if a field has been set.

### GetDescription

`func (o *BTMetadataCategorySummaryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTMetadataCategorySummaryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTMetadataCategorySummaryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTMetadataCategorySummaryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *BTMetadataCategorySummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTMetadataCategorySummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTMetadataCategorySummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTMetadataCategorySummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTMetadataCategorySummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTMetadataCategorySummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTMetadataCategorySummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTMetadataCategorySummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataCategorySummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataCategorySummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataCategorySummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataCategorySummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectTypes

`func (o *BTMetadataCategorySummaryInfo) GetObjectTypes() []int32`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BTMetadataCategorySummaryInfo) GetObjectTypesOk() (*[]int32, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BTMetadataCategorySummaryInfo) SetObjectTypes(v []int32)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *BTMetadataCategorySummaryInfo) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTMetadataCategorySummaryInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTMetadataCategorySummaryInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTMetadataCategorySummaryInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTMetadataCategorySummaryInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTMetadataCategorySummaryInfo) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTMetadataCategorySummaryInfo) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTMetadataCategorySummaryInfo) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTMetadataCategorySummaryInfo) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPublishState

`func (o *BTMetadataCategorySummaryInfo) GetPublishState() int32`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *BTMetadataCategorySummaryInfo) GetPublishStateOk() (*int32, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *BTMetadataCategorySummaryInfo) SetPublishState(v int32)`

SetPublishState sets PublishState field to given value.

### HasPublishState

`func (o *BTMetadataCategorySummaryInfo) HasPublishState() bool`

HasPublishState returns a boolean if a field has been set.

### GetViewRef

`func (o *BTMetadataCategorySummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTMetadataCategorySummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTMetadataCategorySummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTMetadataCategorySummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


