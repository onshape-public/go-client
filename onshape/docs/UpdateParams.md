# UpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromReference** | Pointer to [**BTUniqueDocumentItemParams**](BTUniqueDocumentItemParams.md) |  | [optional] 
**IdsToUpdate** | Pointer to **[]string** |  | [optional] 
**IgnoreChildren** | Pointer to **bool** |  | [optional] 
**ToReference** | Pointer to [**BTUniqueDocumentItemParams**](BTUniqueDocumentItemParams.md) |  | [optional] 

## Methods

### NewUpdateParams

`func NewUpdateParams() *UpdateParams`

NewUpdateParams instantiates a new UpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateParamsWithDefaults

`func NewUpdateParamsWithDefaults() *UpdateParams`

NewUpdateParamsWithDefaults instantiates a new UpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromReference

`func (o *UpdateParams) GetFromReference() BTUniqueDocumentItemParams`

GetFromReference returns the FromReference field if non-nil, zero value otherwise.

### GetFromReferenceOk

`func (o *UpdateParams) GetFromReferenceOk() (*BTUniqueDocumentItemParams, bool)`

GetFromReferenceOk returns a tuple with the FromReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromReference

`func (o *UpdateParams) SetFromReference(v BTUniqueDocumentItemParams)`

SetFromReference sets FromReference field to given value.

### HasFromReference

`func (o *UpdateParams) HasFromReference() bool`

HasFromReference returns a boolean if a field has been set.

### GetIdsToUpdate

`func (o *UpdateParams) GetIdsToUpdate() []string`

GetIdsToUpdate returns the IdsToUpdate field if non-nil, zero value otherwise.

### GetIdsToUpdateOk

`func (o *UpdateParams) GetIdsToUpdateOk() (*[]string, bool)`

GetIdsToUpdateOk returns a tuple with the IdsToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsToUpdate

`func (o *UpdateParams) SetIdsToUpdate(v []string)`

SetIdsToUpdate sets IdsToUpdate field to given value.

### HasIdsToUpdate

`func (o *UpdateParams) HasIdsToUpdate() bool`

HasIdsToUpdate returns a boolean if a field has been set.

### GetIgnoreChildren

`func (o *UpdateParams) GetIgnoreChildren() bool`

GetIgnoreChildren returns the IgnoreChildren field if non-nil, zero value otherwise.

### GetIgnoreChildrenOk

`func (o *UpdateParams) GetIgnoreChildrenOk() (*bool, bool)`

GetIgnoreChildrenOk returns a tuple with the IgnoreChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChildren

`func (o *UpdateParams) SetIgnoreChildren(v bool)`

SetIgnoreChildren sets IgnoreChildren field to given value.

### HasIgnoreChildren

`func (o *UpdateParams) HasIgnoreChildren() bool`

HasIgnoreChildren returns a boolean if a field has been set.

### GetToReference

`func (o *UpdateParams) GetToReference() BTUniqueDocumentItemParams`

GetToReference returns the ToReference field if non-nil, zero value otherwise.

### GetToReferenceOk

`func (o *UpdateParams) GetToReferenceOk() (*BTUniqueDocumentItemParams, bool)`

GetToReferenceOk returns a tuple with the ToReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToReference

`func (o *UpdateParams) SetToReference(v BTUniqueDocumentItemParams)`

SetToReference sets ToReference field to given value.

### HasToReference

`func (o *UpdateParams) HasToReference() bool`

HasToReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


