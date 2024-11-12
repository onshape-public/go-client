# ApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**map[string]MediaType**](MediaType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**map[string]Header**](Header.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewApiResponse

`func NewApiResponse() *ApiResponse`

NewApiResponse instantiates a new ApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponseWithDefaults

`func NewApiResponseWithDefaults() *ApiResponse`

NewApiResponseWithDefaults instantiates a new ApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ApiResponse) GetContent() map[string]MediaType`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ApiResponse) GetContentOk() (*map[string]MediaType, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ApiResponse) SetContent(v map[string]MediaType)`

SetContent sets Content field to given value.

### HasContent

`func (o *ApiResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDescription

`func (o *ApiResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensions

`func (o *ApiResponse) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ApiResponse) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ApiResponse) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ApiResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGetref

`func (o *ApiResponse) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *ApiResponse) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *ApiResponse) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *ApiResponse) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetHeaders

`func (o *ApiResponse) GetHeaders() map[string]Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApiResponse) GetHeadersOk() (*map[string]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApiResponse) SetHeaders(v map[string]Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ApiResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetLinks

`func (o *ApiResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


