# BTBDrawingOperationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Operation description | [optional] 
**FormatVersion** | **string** | Version of drawing entity format. | 
**MessageName** | **string** | Type of drawing modification operation: &#x60;onshapeCreateAnnotations&#x60; | &#x60;onshapeEditAnnotations&#x60; | 
**OtherFields** | Pointer to **map[string]map[string]interface{}** | Other entity creation or modification parameters. | [optional] 

## Methods

### NewBTBDrawingOperationParams

`func NewBTBDrawingOperationParams(formatVersion string, messageName string, ) *BTBDrawingOperationParams`

NewBTBDrawingOperationParams instantiates a new BTBDrawingOperationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBDrawingOperationParamsWithDefaults

`func NewBTBDrawingOperationParamsWithDefaults() *BTBDrawingOperationParams`

NewBTBDrawingOperationParamsWithDefaults instantiates a new BTBDrawingOperationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTBDrawingOperationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTBDrawingOperationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTBDrawingOperationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTBDrawingOperationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatVersion

`func (o *BTBDrawingOperationParams) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *BTBDrawingOperationParams) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *BTBDrawingOperationParams) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.


### GetMessageName

`func (o *BTBDrawingOperationParams) GetMessageName() string`

GetMessageName returns the MessageName field if non-nil, zero value otherwise.

### GetMessageNameOk

`func (o *BTBDrawingOperationParams) GetMessageNameOk() (*string, bool)`

GetMessageNameOk returns a tuple with the MessageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageName

`func (o *BTBDrawingOperationParams) SetMessageName(v string)`

SetMessageName sets MessageName field to given value.


### GetOtherFields

`func (o *BTBDrawingOperationParams) GetOtherFields() map[string]interface{}`

GetOtherFields returns the OtherFields field if non-nil, zero value otherwise.

### GetOtherFieldsOk

`func (o *BTBDrawingOperationParams) GetOtherFieldsOk() (*map[string]interface{}, bool)`

GetOtherFieldsOk returns a tuple with the OtherFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFields

`func (o *BTBDrawingOperationParams) SetOtherFields(v map[string]interface{})`

SetOtherFields sets OtherFields field to given value.

### HasOtherFields

`func (o *BTBDrawingOperationParams) HasOtherFields() bool`

HasOtherFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


