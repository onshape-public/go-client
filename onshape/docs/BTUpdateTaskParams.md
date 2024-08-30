# BTUpdateTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | Use to transfer task ownership to the company. | [optional] 
**DeleteItemIds** | Pointer to **[]string** | References to remove from the task. | [optional] 
**DescriptionParamValue** | Pointer to **string** |  | [optional] 
**ItemParams** | Pointer to [**[]BTTaskItemParams**](BTTaskItemParams.md) | References to add to the task. | [optional] 
**NameParamValue** | Pointer to **string** |  | [optional] 
**PropertyValues** | Pointer to [**[]BTPropertyValueParam**](BTPropertyValueParam.md) | Task metadata properties. | [optional] 

## Methods

### NewBTUpdateTaskParams

`func NewBTUpdateTaskParams() *BTUpdateTaskParams`

NewBTUpdateTaskParams instantiates a new BTUpdateTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUpdateTaskParamsWithDefaults

`func NewBTUpdateTaskParamsWithDefaults() *BTUpdateTaskParams`

NewBTUpdateTaskParamsWithDefaults instantiates a new BTUpdateTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTUpdateTaskParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTUpdateTaskParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTUpdateTaskParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTUpdateTaskParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDeleteItemIds

`func (o *BTUpdateTaskParams) GetDeleteItemIds() []string`

GetDeleteItemIds returns the DeleteItemIds field if non-nil, zero value otherwise.

### GetDeleteItemIdsOk

`func (o *BTUpdateTaskParams) GetDeleteItemIdsOk() (*[]string, bool)`

GetDeleteItemIdsOk returns a tuple with the DeleteItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteItemIds

`func (o *BTUpdateTaskParams) SetDeleteItemIds(v []string)`

SetDeleteItemIds sets DeleteItemIds field to given value.

### HasDeleteItemIds

`func (o *BTUpdateTaskParams) HasDeleteItemIds() bool`

HasDeleteItemIds returns a boolean if a field has been set.

### GetDescriptionParamValue

`func (o *BTUpdateTaskParams) GetDescriptionParamValue() string`

GetDescriptionParamValue returns the DescriptionParamValue field if non-nil, zero value otherwise.

### GetDescriptionParamValueOk

`func (o *BTUpdateTaskParams) GetDescriptionParamValueOk() (*string, bool)`

GetDescriptionParamValueOk returns a tuple with the DescriptionParamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionParamValue

`func (o *BTUpdateTaskParams) SetDescriptionParamValue(v string)`

SetDescriptionParamValue sets DescriptionParamValue field to given value.

### HasDescriptionParamValue

`func (o *BTUpdateTaskParams) HasDescriptionParamValue() bool`

HasDescriptionParamValue returns a boolean if a field has been set.

### GetItemParams

`func (o *BTUpdateTaskParams) GetItemParams() []BTTaskItemParams`

GetItemParams returns the ItemParams field if non-nil, zero value otherwise.

### GetItemParamsOk

`func (o *BTUpdateTaskParams) GetItemParamsOk() (*[]BTTaskItemParams, bool)`

GetItemParamsOk returns a tuple with the ItemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemParams

`func (o *BTUpdateTaskParams) SetItemParams(v []BTTaskItemParams)`

SetItemParams sets ItemParams field to given value.

### HasItemParams

`func (o *BTUpdateTaskParams) HasItemParams() bool`

HasItemParams returns a boolean if a field has been set.

### GetNameParamValue

`func (o *BTUpdateTaskParams) GetNameParamValue() string`

GetNameParamValue returns the NameParamValue field if non-nil, zero value otherwise.

### GetNameParamValueOk

`func (o *BTUpdateTaskParams) GetNameParamValueOk() (*string, bool)`

GetNameParamValueOk returns a tuple with the NameParamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameParamValue

`func (o *BTUpdateTaskParams) SetNameParamValue(v string)`

SetNameParamValue sets NameParamValue field to given value.

### HasNameParamValue

`func (o *BTUpdateTaskParams) HasNameParamValue() bool`

HasNameParamValue returns a boolean if a field has been set.

### GetPropertyValues

`func (o *BTUpdateTaskParams) GetPropertyValues() []BTPropertyValueParam`

GetPropertyValues returns the PropertyValues field if non-nil, zero value otherwise.

### GetPropertyValuesOk

`func (o *BTUpdateTaskParams) GetPropertyValuesOk() (*[]BTPropertyValueParam, bool)`

GetPropertyValuesOk returns a tuple with the PropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValues

`func (o *BTUpdateTaskParams) SetPropertyValues(v []BTPropertyValueParam)`

SetPropertyValues sets PropertyValues field to given value.

### HasPropertyValues

`func (o *BTUpdateTaskParams) HasPropertyValues() bool`

HasPropertyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


