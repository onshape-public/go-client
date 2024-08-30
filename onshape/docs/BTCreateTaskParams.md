# BTCreateTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **string** | Id of the company that owns the task. | 
**Description** | Pointer to **string** | Description of the task. | [optional] 
**DescriptionParamValue** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** | Id of a document to associate the task to. | [optional] 
**ItemParams** | Pointer to [**[]BTTaskItemParams**](BTTaskItemParams.md) | References to include in the task. | [optional] 
**Name** | Pointer to **string** | Name of the task. | [optional] 
**NameParamValue** | Pointer to **string** |  | [optional] 
**PropertyValues** | Pointer to [**[]BTPropertyValueParam**](BTPropertyValueParam.md) | Set Task metadata properties. | [optional] 

## Methods

### NewBTCreateTaskParams

`func NewBTCreateTaskParams(companyId string, ) *BTCreateTaskParams`

NewBTCreateTaskParams instantiates a new BTCreateTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCreateTaskParamsWithDefaults

`func NewBTCreateTaskParamsWithDefaults() *BTCreateTaskParams`

NewBTCreateTaskParamsWithDefaults instantiates a new BTCreateTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTCreateTaskParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTCreateTaskParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTCreateTaskParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetDescription

`func (o *BTCreateTaskParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTCreateTaskParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTCreateTaskParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTCreateTaskParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionParamValue

`func (o *BTCreateTaskParams) GetDescriptionParamValue() string`

GetDescriptionParamValue returns the DescriptionParamValue field if non-nil, zero value otherwise.

### GetDescriptionParamValueOk

`func (o *BTCreateTaskParams) GetDescriptionParamValueOk() (*string, bool)`

GetDescriptionParamValueOk returns a tuple with the DescriptionParamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionParamValue

`func (o *BTCreateTaskParams) SetDescriptionParamValue(v string)`

SetDescriptionParamValue sets DescriptionParamValue field to given value.

### HasDescriptionParamValue

`func (o *BTCreateTaskParams) HasDescriptionParamValue() bool`

HasDescriptionParamValue returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTCreateTaskParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTCreateTaskParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTCreateTaskParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTCreateTaskParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetItemParams

`func (o *BTCreateTaskParams) GetItemParams() []BTTaskItemParams`

GetItemParams returns the ItemParams field if non-nil, zero value otherwise.

### GetItemParamsOk

`func (o *BTCreateTaskParams) GetItemParamsOk() (*[]BTTaskItemParams, bool)`

GetItemParamsOk returns a tuple with the ItemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemParams

`func (o *BTCreateTaskParams) SetItemParams(v []BTTaskItemParams)`

SetItemParams sets ItemParams field to given value.

### HasItemParams

`func (o *BTCreateTaskParams) HasItemParams() bool`

HasItemParams returns a boolean if a field has been set.

### GetName

`func (o *BTCreateTaskParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTCreateTaskParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTCreateTaskParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTCreateTaskParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameParamValue

`func (o *BTCreateTaskParams) GetNameParamValue() string`

GetNameParamValue returns the NameParamValue field if non-nil, zero value otherwise.

### GetNameParamValueOk

`func (o *BTCreateTaskParams) GetNameParamValueOk() (*string, bool)`

GetNameParamValueOk returns a tuple with the NameParamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameParamValue

`func (o *BTCreateTaskParams) SetNameParamValue(v string)`

SetNameParamValue sets NameParamValue field to given value.

### HasNameParamValue

`func (o *BTCreateTaskParams) HasNameParamValue() bool`

HasNameParamValue returns a boolean if a field has been set.

### GetPropertyValues

`func (o *BTCreateTaskParams) GetPropertyValues() []BTPropertyValueParam`

GetPropertyValues returns the PropertyValues field if non-nil, zero value otherwise.

### GetPropertyValuesOk

`func (o *BTCreateTaskParams) GetPropertyValuesOk() (*[]BTPropertyValueParam, bool)`

GetPropertyValuesOk returns a tuple with the PropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValues

`func (o *BTCreateTaskParams) SetPropertyValues(v []BTPropertyValueParam)`

SetPropertyValues sets PropertyValues field to given value.

### HasPropertyValues

`func (o *BTCreateTaskParams) HasPropertyValues() bool`

HasPropertyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


