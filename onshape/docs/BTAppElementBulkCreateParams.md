# BTAppElementBulkCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The label that will appear in the document&#39;s edit history for this operation. If blank, a value will be auto-generated. | [optional] 
**FormatId** | **string** | The data type of the application. This string allows an application to distinguish their elements from elements of another application. | 
**Location** | Pointer to [**BTElementLocationParams**](BTElementLocationParams.md) |  | [optional] 
**Names** | Pointer to **[]string** | The name of the element being created. If blank, a name will be auto-generated. | [optional] 

## Methods

### NewBTAppElementBulkCreateParams

`func NewBTAppElementBulkCreateParams(formatId string, ) *BTAppElementBulkCreateParams`

NewBTAppElementBulkCreateParams instantiates a new BTAppElementBulkCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementBulkCreateParamsWithDefaults

`func NewBTAppElementBulkCreateParamsWithDefaults() *BTAppElementBulkCreateParams`

NewBTAppElementBulkCreateParamsWithDefaults instantiates a new BTAppElementBulkCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTAppElementBulkCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAppElementBulkCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAppElementBulkCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAppElementBulkCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatId

`func (o *BTAppElementBulkCreateParams) GetFormatId() string`

GetFormatId returns the FormatId field if non-nil, zero value otherwise.

### GetFormatIdOk

`func (o *BTAppElementBulkCreateParams) GetFormatIdOk() (*string, bool)`

GetFormatIdOk returns a tuple with the FormatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatId

`func (o *BTAppElementBulkCreateParams) SetFormatId(v string)`

SetFormatId sets FormatId field to given value.


### GetLocation

`func (o *BTAppElementBulkCreateParams) GetLocation() BTElementLocationParams`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BTAppElementBulkCreateParams) GetLocationOk() (*BTElementLocationParams, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BTAppElementBulkCreateParams) SetLocation(v BTElementLocationParams)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BTAppElementBulkCreateParams) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNames

`func (o *BTAppElementBulkCreateParams) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *BTAppElementBulkCreateParams) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *BTAppElementBulkCreateParams) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *BTAppElementBulkCreateParams) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


