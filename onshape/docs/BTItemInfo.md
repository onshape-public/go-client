# BTItemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | ID of the company, classroom, or enterprise that owns this item. | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Properties** | Pointer to **map[string]string** | Map of the item&#39;s properties and their values. | [optional] 
**PublishState** | Pointer to **int32** | &#x60;0: PENDING | 1: ACTIVE | 2: INACTIVE&#x60; | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTItemInfo

`func NewBTItemInfo() *BTItemInfo`

NewBTItemInfo instantiates a new BTItemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTItemInfoWithDefaults

`func NewBTItemInfoWithDefaults() *BTItemInfo`

NewBTItemInfoWithDefaults instantiates a new BTItemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTItemInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTItemInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTItemInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTItemInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetHref

`func (o *BTItemInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTItemInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTItemInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTItemInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTItemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTItemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTItemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTItemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTItemInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTItemInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTItemInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTItemInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *BTItemInfo) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTItemInfo) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTItemInfo) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTItemInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPublishState

`func (o *BTItemInfo) GetPublishState() int32`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *BTItemInfo) GetPublishStateOk() (*int32, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *BTItemInfo) SetPublishState(v int32)`

SetPublishState sets PublishState field to given value.

### HasPublishState

`func (o *BTItemInfo) HasPublishState() bool`

HasPublishState returns a boolean if a field has been set.

### GetViewRef

`func (o *BTItemInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTItemInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTItemInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTItemInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


