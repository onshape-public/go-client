# BTInsertablesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSaveVersion** | Pointer to **bool** |  | [optional] 
**ChangesSinceVersionSave** | Pointer to **int32** |  | [optional] 
**HasMultipleVersions** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]BTInsertableInfo**](BTInsertableInfo.md) |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 

## Methods

### NewBTInsertablesListResponse

`func NewBTInsertablesListResponse() *BTInsertablesListResponse`

NewBTInsertablesListResponse instantiates a new BTInsertablesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTInsertablesListResponseWithDefaults

`func NewBTInsertablesListResponseWithDefaults() *BTInsertablesListResponse`

NewBTInsertablesListResponseWithDefaults instantiates a new BTInsertablesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSaveVersion

`func (o *BTInsertablesListResponse) GetCanSaveVersion() bool`

GetCanSaveVersion returns the CanSaveVersion field if non-nil, zero value otherwise.

### GetCanSaveVersionOk

`func (o *BTInsertablesListResponse) GetCanSaveVersionOk() (*bool, bool)`

GetCanSaveVersionOk returns a tuple with the CanSaveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSaveVersion

`func (o *BTInsertablesListResponse) SetCanSaveVersion(v bool)`

SetCanSaveVersion sets CanSaveVersion field to given value.

### HasCanSaveVersion

`func (o *BTInsertablesListResponse) HasCanSaveVersion() bool`

HasCanSaveVersion returns a boolean if a field has been set.

### GetChangesSinceVersionSave

`func (o *BTInsertablesListResponse) GetChangesSinceVersionSave() int32`

GetChangesSinceVersionSave returns the ChangesSinceVersionSave field if non-nil, zero value otherwise.

### GetChangesSinceVersionSaveOk

`func (o *BTInsertablesListResponse) GetChangesSinceVersionSaveOk() (*int32, bool)`

GetChangesSinceVersionSaveOk returns a tuple with the ChangesSinceVersionSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesSinceVersionSave

`func (o *BTInsertablesListResponse) SetChangesSinceVersionSave(v int32)`

SetChangesSinceVersionSave sets ChangesSinceVersionSave field to given value.

### HasChangesSinceVersionSave

`func (o *BTInsertablesListResponse) HasChangesSinceVersionSave() bool`

HasChangesSinceVersionSave returns a boolean if a field has been set.

### GetHasMultipleVersions

`func (o *BTInsertablesListResponse) GetHasMultipleVersions() bool`

GetHasMultipleVersions returns the HasMultipleVersions field if non-nil, zero value otherwise.

### GetHasMultipleVersionsOk

`func (o *BTInsertablesListResponse) GetHasMultipleVersionsOk() (*bool, bool)`

GetHasMultipleVersionsOk returns a tuple with the HasMultipleVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMultipleVersions

`func (o *BTInsertablesListResponse) SetHasMultipleVersions(v bool)`

SetHasMultipleVersions sets HasMultipleVersions field to given value.

### HasHasMultipleVersions

`func (o *BTInsertablesListResponse) HasHasMultipleVersions() bool`

HasHasMultipleVersions returns a boolean if a field has been set.

### GetHref

`func (o *BTInsertablesListResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTInsertablesListResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTInsertablesListResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTInsertablesListResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTInsertablesListResponse) GetItems() []BTInsertableInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTInsertablesListResponse) GetItemsOk() (*[]BTInsertableInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTInsertablesListResponse) SetItems(v []BTInsertableInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTInsertablesListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTInsertablesListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTInsertablesListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTInsertablesListResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTInsertablesListResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTInsertablesListResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTInsertablesListResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTInsertablesListResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTInsertablesListResponse) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


