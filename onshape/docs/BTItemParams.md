# BTItemParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | ID of the company, classroom, or enterprise that owns this item. | [optional] 
**Name** | Pointer to **string** | Item name. | [optional] 
**PublishState** | Pointer to **int32** | &#x60;0: PENDING | 1: ACTIVE | 2: INACTIVE&#x60; | [optional] 

## Methods

### NewBTItemParams

`func NewBTItemParams() *BTItemParams`

NewBTItemParams instantiates a new BTItemParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTItemParamsWithDefaults

`func NewBTItemParamsWithDefaults() *BTItemParams`

NewBTItemParamsWithDefaults instantiates a new BTItemParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTItemParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTItemParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTItemParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTItemParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetName

`func (o *BTItemParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTItemParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTItemParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTItemParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishState

`func (o *BTItemParams) GetPublishState() int32`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *BTItemParams) GetPublishStateOk() (*int32, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *BTItemParams) SetPublishState(v int32)`

SetPublishState sets PublishState field to given value.

### HasPublishState

`func (o *BTItemParams) HasPublishState() bool`

HasPublishState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


