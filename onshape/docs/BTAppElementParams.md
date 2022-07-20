# BTAppElementParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The label that will appear in the document&#39;s edit history for this operation. If blank, a value will be auto-generated. | [optional] 
**FormatId** | **string** | The data type of the application. This string allows an application to distinguish their elements from elements of another application. | 
**JsonTree** | Pointer to **map[string]interface{}** | Initialization data for the new element&#39;s json tree. | [optional] 
**Location** | Pointer to [**BTElementLocationParams**](BTElementLocationParams.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the element being created. If blank, a name will be auto-generated. | [optional] 
**Subelements** | Pointer to [**[]BTAppElementChangeParams**](BTAppElementChangeParams.md) | Initialization data for the new element&#39;s subelements. | [optional] 

## Methods

### NewBTAppElementParams

`func NewBTAppElementParams(formatId string, ) *BTAppElementParams`

NewBTAppElementParams instantiates a new BTAppElementParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementParamsWithDefaults

`func NewBTAppElementParamsWithDefaults() *BTAppElementParams`

NewBTAppElementParamsWithDefaults instantiates a new BTAppElementParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTAppElementParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAppElementParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAppElementParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAppElementParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatId

`func (o *BTAppElementParams) GetFormatId() string`

GetFormatId returns the FormatId field if non-nil, zero value otherwise.

### GetFormatIdOk

`func (o *BTAppElementParams) GetFormatIdOk() (*string, bool)`

GetFormatIdOk returns a tuple with the FormatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatId

`func (o *BTAppElementParams) SetFormatId(v string)`

SetFormatId sets FormatId field to given value.


### GetJsonTree

`func (o *BTAppElementParams) GetJsonTree() map[string]interface{}`

GetJsonTree returns the JsonTree field if non-nil, zero value otherwise.

### GetJsonTreeOk

`func (o *BTAppElementParams) GetJsonTreeOk() (*map[string]interface{}, bool)`

GetJsonTreeOk returns a tuple with the JsonTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonTree

`func (o *BTAppElementParams) SetJsonTree(v map[string]interface{})`

SetJsonTree sets JsonTree field to given value.

### HasJsonTree

`func (o *BTAppElementParams) HasJsonTree() bool`

HasJsonTree returns a boolean if a field has been set.

### GetLocation

`func (o *BTAppElementParams) GetLocation() BTElementLocationParams`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BTAppElementParams) GetLocationOk() (*BTElementLocationParams, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BTAppElementParams) SetLocation(v BTElementLocationParams)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BTAppElementParams) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *BTAppElementParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAppElementParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAppElementParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAppElementParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubelements

`func (o *BTAppElementParams) GetSubelements() []BTAppElementChangeParams`

GetSubelements returns the Subelements field if non-nil, zero value otherwise.

### GetSubelementsOk

`func (o *BTAppElementParams) GetSubelementsOk() (*[]BTAppElementChangeParams, bool)`

GetSubelementsOk returns a tuple with the Subelements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelements

`func (o *BTAppElementParams) SetSubelements(v []BTAppElementChangeParams)`

SetSubelements sets Subelements field to given value.

### HasSubelements

`func (o *BTAppElementParams) HasSubelements() bool`

HasSubelements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


