# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BufferView** | Pointer to **int32** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBufferView

`func (o *Image) GetBufferView() int32`

GetBufferView returns the BufferView field if non-nil, zero value otherwise.

### GetBufferViewOk

`func (o *Image) GetBufferViewOk() (*int32, bool)`

GetBufferViewOk returns a tuple with the BufferView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferView

`func (o *Image) SetBufferView(v int32)`

SetBufferView sets BufferView field to given value.

### HasBufferView

`func (o *Image) HasBufferView() bool`

HasBufferView returns a boolean if a field has been set.

### GetExtensions

`func (o *Image) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Image) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Image) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Image) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Image) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Image) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Image) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Image) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetMimeType

`func (o *Image) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Image) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Image) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Image) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Image) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUri

`func (o *Image) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Image) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Image) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Image) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


