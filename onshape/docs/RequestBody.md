# RequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**RequestBodyContent**](RequestBodyContent.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewRequestBody

`func NewRequestBody() *RequestBody`

NewRequestBody instantiates a new RequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBodyWithDefaults

`func NewRequestBodyWithDefaults() *RequestBody`

NewRequestBodyWithDefaults instantiates a new RequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *RequestBody) GetContent() RequestBodyContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RequestBody) GetContentOk() (*RequestBodyContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RequestBody) SetContent(v RequestBodyContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *RequestBody) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDescription

`func (o *RequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensions

`func (o *RequestBody) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *RequestBody) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *RequestBody) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *RequestBody) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGetref

`func (o *RequestBody) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *RequestBody) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *RequestBody) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *RequestBody) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetRequired

`func (o *RequestBody) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RequestBody) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RequestBody) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RequestBody) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


