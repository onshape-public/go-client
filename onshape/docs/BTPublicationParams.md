# BTPublicationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]BTPublicationItemParams**](BTPublicationItemParams.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**OldClientNotes** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPublicationParams

`func NewBTPublicationParams() *BTPublicationParams`

NewBTPublicationParams instantiates a new BTPublicationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPublicationParamsWithDefaults

`func NewBTPublicationParamsWithDefaults() *BTPublicationParams`

NewBTPublicationParamsWithDefaults instantiates a new BTPublicationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTPublicationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTPublicationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTPublicationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTPublicationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *BTPublicationParams) GetItems() []BTPublicationItemParams`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTPublicationParams) GetItemsOk() (*[]BTPublicationItemParams, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTPublicationParams) SetItems(v []BTPublicationItemParams)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTPublicationParams) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *BTPublicationParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPublicationParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPublicationParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTPublicationParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *BTPublicationParams) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BTPublicationParams) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BTPublicationParams) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BTPublicationParams) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOldClientNotes

`func (o *BTPublicationParams) GetOldClientNotes() string`

GetOldClientNotes returns the OldClientNotes field if non-nil, zero value otherwise.

### GetOldClientNotesOk

`func (o *BTPublicationParams) GetOldClientNotesOk() (*string, bool)`

GetOldClientNotesOk returns a tuple with the OldClientNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldClientNotes

`func (o *BTPublicationParams) SetOldClientNotes(v string)`

SetOldClientNotes sets OldClientNotes field to given value.

### HasOldClientNotes

`func (o *BTPublicationParams) HasOldClientNotes() bool`

HasOldClientNotes returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTPublicationParams) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTPublicationParams) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTPublicationParams) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTPublicationParams) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTPublicationParams) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTPublicationParams) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTPublicationParams) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTPublicationParams) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetParentId

`func (o *BTPublicationParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTPublicationParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTPublicationParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTPublicationParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTPublicationParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTPublicationParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTPublicationParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTPublicationParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


