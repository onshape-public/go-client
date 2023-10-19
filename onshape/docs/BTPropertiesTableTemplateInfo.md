# BTPropertiesTableTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsAllCaps** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**PropertyColumns** | Pointer to [**[]BTSimplePropertyInfo**](BTSimplePropertyInfo.md) |  | [optional] 
**TableType** | Pointer to [**BTPropertiesTableTemplateType**](BTPropertiesTableTemplateType.md) |  | [optional] 
**TemplateGroupId** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTPropertiesTableTemplateInfo

`func NewBTPropertiesTableTemplateInfo() *BTPropertiesTableTemplateInfo`

NewBTPropertiesTableTemplateInfo instantiates a new BTPropertiesTableTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPropertiesTableTemplateInfoWithDefaults

`func NewBTPropertiesTableTemplateInfoWithDefaults() *BTPropertiesTableTemplateInfo`

NewBTPropertiesTableTemplateInfoWithDefaults instantiates a new BTPropertiesTableTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTPropertiesTableTemplateInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTPropertiesTableTemplateInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTPropertiesTableTemplateInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTPropertiesTableTemplateInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetHref

`func (o *BTPropertiesTableTemplateInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTPropertiesTableTemplateInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTPropertiesTableTemplateInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTPropertiesTableTemplateInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTPropertiesTableTemplateInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPropertiesTableTemplateInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPropertiesTableTemplateInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPropertiesTableTemplateInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *BTPropertiesTableTemplateInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BTPropertiesTableTemplateInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BTPropertiesTableTemplateInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BTPropertiesTableTemplateInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsAllCaps

`func (o *BTPropertiesTableTemplateInfo) GetIsAllCaps() bool`

GetIsAllCaps returns the IsAllCaps field if non-nil, zero value otherwise.

### GetIsAllCapsOk

`func (o *BTPropertiesTableTemplateInfo) GetIsAllCapsOk() (*bool, bool)`

GetIsAllCapsOk returns a tuple with the IsAllCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllCaps

`func (o *BTPropertiesTableTemplateInfo) SetIsAllCaps(v bool)`

SetIsAllCaps sets IsAllCaps field to given value.

### HasIsAllCaps

`func (o *BTPropertiesTableTemplateInfo) HasIsAllCaps() bool`

HasIsAllCaps returns a boolean if a field has been set.

### GetName

`func (o *BTPropertiesTableTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPropertiesTableTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPropertiesTableTemplateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTPropertiesTableTemplateInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyColumns

`func (o *BTPropertiesTableTemplateInfo) GetPropertyColumns() []BTSimplePropertyInfo`

GetPropertyColumns returns the PropertyColumns field if non-nil, zero value otherwise.

### GetPropertyColumnsOk

`func (o *BTPropertiesTableTemplateInfo) GetPropertyColumnsOk() (*[]BTSimplePropertyInfo, bool)`

GetPropertyColumnsOk returns a tuple with the PropertyColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyColumns

`func (o *BTPropertiesTableTemplateInfo) SetPropertyColumns(v []BTSimplePropertyInfo)`

SetPropertyColumns sets PropertyColumns field to given value.

### HasPropertyColumns

`func (o *BTPropertiesTableTemplateInfo) HasPropertyColumns() bool`

HasPropertyColumns returns a boolean if a field has been set.

### GetTableType

`func (o *BTPropertiesTableTemplateInfo) GetTableType() BTPropertiesTableTemplateType`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *BTPropertiesTableTemplateInfo) GetTableTypeOk() (*BTPropertiesTableTemplateType, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *BTPropertiesTableTemplateInfo) SetTableType(v BTPropertiesTableTemplateType)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *BTPropertiesTableTemplateInfo) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTemplateGroupId

`func (o *BTPropertiesTableTemplateInfo) GetTemplateGroupId() string`

GetTemplateGroupId returns the TemplateGroupId field if non-nil, zero value otherwise.

### GetTemplateGroupIdOk

`func (o *BTPropertiesTableTemplateInfo) GetTemplateGroupIdOk() (*string, bool)`

GetTemplateGroupIdOk returns a tuple with the TemplateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateGroupId

`func (o *BTPropertiesTableTemplateInfo) SetTemplateGroupId(v string)`

SetTemplateGroupId sets TemplateGroupId field to given value.

### HasTemplateGroupId

`func (o *BTPropertiesTableTemplateInfo) HasTemplateGroupId() bool`

HasTemplateGroupId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTPropertiesTableTemplateInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTPropertiesTableTemplateInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTPropertiesTableTemplateInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTPropertiesTableTemplateInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


