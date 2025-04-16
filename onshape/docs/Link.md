# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**map[string]Header**](Header.md) |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] 
**OperationRef** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**RequestBody** | Pointer to **map[string]interface{}** |  | [optional] 
**Server** | Pointer to [**Server**](Server.md) |  | [optional] 

## Methods

### NewLink

`func NewLink() *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Link) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Link) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Link) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Link) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensions

`func (o *Link) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Link) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Link) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Link) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGetref

`func (o *Link) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *Link) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *Link) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *Link) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetHeaders

`func (o *Link) GetHeaders() map[string]Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Link) GetHeadersOk() (*map[string]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Link) SetHeaders(v map[string]Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Link) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetOperationId

`func (o *Link) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *Link) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *Link) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *Link) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetOperationRef

`func (o *Link) GetOperationRef() string`

GetOperationRef returns the OperationRef field if non-nil, zero value otherwise.

### GetOperationRefOk

`func (o *Link) GetOperationRefOk() (*string, bool)`

GetOperationRefOk returns a tuple with the OperationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationRef

`func (o *Link) SetOperationRef(v string)`

SetOperationRef sets OperationRef field to given value.

### HasOperationRef

`func (o *Link) HasOperationRef() bool`

HasOperationRef returns a boolean if a field has been set.

### GetParameters

`func (o *Link) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Link) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Link) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Link) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRequestBody

`func (o *Link) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *Link) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *Link) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *Link) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetServer

`func (o *Link) GetServer() Server`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Link) GetServerOk() (*Server, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Link) SetServer(v Server)`

SetServer sets Server field to given value.

### HasServer

`func (o *Link) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


