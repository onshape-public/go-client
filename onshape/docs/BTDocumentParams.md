# BTDocumentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetaCapabilityIds** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Elements** | Pointer to [**[]BTDocumentElementCreationDescriptor**](BTDocumentElementCreationDescriptor.md) |  | [optional] 
**GenerateUnknownMessages** | Pointer to **bool** |  | [optional] 
**IsEmptyContent** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotRevisionManaged** | Pointer to **bool** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBTDocumentParams

`func NewBTDocumentParams() *BTDocumentParams`

NewBTDocumentParams instantiates a new BTDocumentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentParamsWithDefaults

`func NewBTDocumentParamsWithDefaults() *BTDocumentParams`

NewBTDocumentParamsWithDefaults instantiates a new BTDocumentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetaCapabilityIds

`func (o *BTDocumentParams) GetBetaCapabilityIds() []string`

GetBetaCapabilityIds returns the BetaCapabilityIds field if non-nil, zero value otherwise.

### GetBetaCapabilityIdsOk

`func (o *BTDocumentParams) GetBetaCapabilityIdsOk() (*[]string, bool)`

GetBetaCapabilityIdsOk returns a tuple with the BetaCapabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaCapabilityIds

`func (o *BTDocumentParams) SetBetaCapabilityIds(v []string)`

SetBetaCapabilityIds sets BetaCapabilityIds field to given value.

### HasBetaCapabilityIds

`func (o *BTDocumentParams) HasBetaCapabilityIds() bool`

HasBetaCapabilityIds returns a boolean if a field has been set.

### GetDescription

`func (o *BTDocumentParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDocumentParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDocumentParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDocumentParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElements

`func (o *BTDocumentParams) GetElements() []BTDocumentElementCreationDescriptor`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *BTDocumentParams) GetElementsOk() (*[]BTDocumentElementCreationDescriptor, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *BTDocumentParams) SetElements(v []BTDocumentElementCreationDescriptor)`

SetElements sets Elements field to given value.

### HasElements

`func (o *BTDocumentParams) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetGenerateUnknownMessages

`func (o *BTDocumentParams) GetGenerateUnknownMessages() bool`

GetGenerateUnknownMessages returns the GenerateUnknownMessages field if non-nil, zero value otherwise.

### GetGenerateUnknownMessagesOk

`func (o *BTDocumentParams) GetGenerateUnknownMessagesOk() (*bool, bool)`

GetGenerateUnknownMessagesOk returns a tuple with the GenerateUnknownMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateUnknownMessages

`func (o *BTDocumentParams) SetGenerateUnknownMessages(v bool)`

SetGenerateUnknownMessages sets GenerateUnknownMessages field to given value.

### HasGenerateUnknownMessages

`func (o *BTDocumentParams) HasGenerateUnknownMessages() bool`

HasGenerateUnknownMessages returns a boolean if a field has been set.

### GetIsEmptyContent

`func (o *BTDocumentParams) GetIsEmptyContent() bool`

GetIsEmptyContent returns the IsEmptyContent field if non-nil, zero value otherwise.

### GetIsEmptyContentOk

`func (o *BTDocumentParams) GetIsEmptyContentOk() (*bool, bool)`

GetIsEmptyContentOk returns a tuple with the IsEmptyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyContent

`func (o *BTDocumentParams) SetIsEmptyContent(v bool)`

SetIsEmptyContent sets IsEmptyContent field to given value.

### HasIsEmptyContent

`func (o *BTDocumentParams) HasIsEmptyContent() bool`

HasIsEmptyContent returns a boolean if a field has been set.

### GetIsPublic

`func (o *BTDocumentParams) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *BTDocumentParams) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *BTDocumentParams) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *BTDocumentParams) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotRevisionManaged

`func (o *BTDocumentParams) GetNotRevisionManaged() bool`

GetNotRevisionManaged returns the NotRevisionManaged field if non-nil, zero value otherwise.

### GetNotRevisionManagedOk

`func (o *BTDocumentParams) GetNotRevisionManagedOk() (*bool, bool)`

GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRevisionManaged

`func (o *BTDocumentParams) SetNotRevisionManaged(v bool)`

SetNotRevisionManaged sets NotRevisionManaged field to given value.

### HasNotRevisionManaged

`func (o *BTDocumentParams) HasNotRevisionManaged() bool`

HasNotRevisionManaged returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *BTDocumentParams) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *BTDocumentParams) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *BTDocumentParams) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *BTDocumentParams) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTDocumentParams) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTDocumentParams) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTDocumentParams) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTDocumentParams) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTDocumentParams) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTDocumentParams) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTDocumentParams) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTDocumentParams) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetParentId

`func (o *BTDocumentParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTDocumentParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTDocumentParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTDocumentParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTDocumentParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTDocumentParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTDocumentParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTDocumentParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTags

`func (o *BTDocumentParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BTDocumentParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BTDocumentParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BTDocumentParams) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


