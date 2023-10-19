# Encoding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowReserved** | Pointer to **bool** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Explode** | Pointer to **bool** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Headers** | Pointer to [**map[string]Header**](Header.md) |  | [optional] 
**Style** | Pointer to [**StyleEnum**](StyleEnum.md) |  | [optional] 

## Methods

### NewEncoding

`func NewEncoding() *Encoding`

NewEncoding instantiates a new Encoding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodingWithDefaults

`func NewEncodingWithDefaults() *Encoding`

NewEncodingWithDefaults instantiates a new Encoding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowReserved

`func (o *Encoding) GetAllowReserved() bool`

GetAllowReserved returns the AllowReserved field if non-nil, zero value otherwise.

### GetAllowReservedOk

`func (o *Encoding) GetAllowReservedOk() (*bool, bool)`

GetAllowReservedOk returns a tuple with the AllowReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReserved

`func (o *Encoding) SetAllowReserved(v bool)`

SetAllowReserved sets AllowReserved field to given value.

### HasAllowReserved

`func (o *Encoding) HasAllowReserved() bool`

HasAllowReserved returns a boolean if a field has been set.

### GetContentType

`func (o *Encoding) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Encoding) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Encoding) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Encoding) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExplode

`func (o *Encoding) GetExplode() bool`

GetExplode returns the Explode field if non-nil, zero value otherwise.

### GetExplodeOk

`func (o *Encoding) GetExplodeOk() (*bool, bool)`

GetExplodeOk returns a tuple with the Explode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplode

`func (o *Encoding) SetExplode(v bool)`

SetExplode sets Explode field to given value.

### HasExplode

`func (o *Encoding) HasExplode() bool`

HasExplode returns a boolean if a field has been set.

### GetExtensions

`func (o *Encoding) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Encoding) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Encoding) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Encoding) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetHeaders

`func (o *Encoding) GetHeaders() map[string]Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Encoding) GetHeadersOk() (*map[string]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Encoding) SetHeaders(v map[string]Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Encoding) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetStyle

`func (o *Encoding) GetStyle() StyleEnum`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *Encoding) GetStyleOk() (*StyleEnum, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *Encoding) SetStyle(v StyleEnum)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *Encoding) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


