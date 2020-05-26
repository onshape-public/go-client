# FormDataBodyPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDisposition** | Pointer to [**ContentDisposition**](ContentDisposition.md) |  | [optional] 
**Entity** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**FormDataContentDisposition** | Pointer to [**FormDataContentDisposition**](FormDataContentDisposition.md) |  | [optional] 
**Headers** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**MediaType** | Pointer to [**BodyPartMediaType**](BodyPart_mediaType.md) |  | [optional] 
**MessageBodyWorkers** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParameterizedHeaders** | Pointer to [**map[string][]ParameterizedHeader**](array.md) |  | [optional] 
**Parent** | Pointer to [**MultiPart**](MultiPart.md) |  | [optional] 
**Providers** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**Simple** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewFormDataBodyPart

`func NewFormDataBodyPart() *FormDataBodyPart`

NewFormDataBodyPart instantiates a new FormDataBodyPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDataBodyPartWithDefaults

`func NewFormDataBodyPartWithDefaults() *FormDataBodyPart`

NewFormDataBodyPartWithDefaults instantiates a new FormDataBodyPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDisposition

`func (o *FormDataBodyPart) GetContentDisposition() ContentDisposition`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *FormDataBodyPart) GetContentDispositionOk() (*ContentDisposition, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *FormDataBodyPart) SetContentDisposition(v ContentDisposition)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *FormDataBodyPart) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetEntity

`func (o *FormDataBodyPart) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *FormDataBodyPart) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *FormDataBodyPart) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *FormDataBodyPart) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFormDataContentDisposition

`func (o *FormDataBodyPart) GetFormDataContentDisposition() FormDataContentDisposition`

GetFormDataContentDisposition returns the FormDataContentDisposition field if non-nil, zero value otherwise.

### GetFormDataContentDispositionOk

`func (o *FormDataBodyPart) GetFormDataContentDispositionOk() (*FormDataContentDisposition, bool)`

GetFormDataContentDispositionOk returns a tuple with the FormDataContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormDataContentDisposition

`func (o *FormDataBodyPart) SetFormDataContentDisposition(v FormDataContentDisposition)`

SetFormDataContentDisposition sets FormDataContentDisposition field to given value.

### HasFormDataContentDisposition

`func (o *FormDataBodyPart) HasFormDataContentDisposition() bool`

HasFormDataContentDisposition returns a boolean if a field has been set.

### GetHeaders

`func (o *FormDataBodyPart) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FormDataBodyPart) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FormDataBodyPart) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FormDataBodyPart) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMediaType

`func (o *FormDataBodyPart) GetMediaType() BodyPartMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *FormDataBodyPart) GetMediaTypeOk() (*BodyPartMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *FormDataBodyPart) SetMediaType(v BodyPartMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *FormDataBodyPart) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMessageBodyWorkers

`func (o *FormDataBodyPart) GetMessageBodyWorkers() map[string]interface{}`

GetMessageBodyWorkers returns the MessageBodyWorkers field if non-nil, zero value otherwise.

### GetMessageBodyWorkersOk

`func (o *FormDataBodyPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool)`

GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBodyWorkers

`func (o *FormDataBodyPart) SetMessageBodyWorkers(v map[string]interface{})`

SetMessageBodyWorkers sets MessageBodyWorkers field to given value.

### HasMessageBodyWorkers

`func (o *FormDataBodyPart) HasMessageBodyWorkers() bool`

HasMessageBodyWorkers returns a boolean if a field has been set.

### GetName

`func (o *FormDataBodyPart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormDataBodyPart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormDataBodyPart) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormDataBodyPart) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameterizedHeaders

`func (o *FormDataBodyPart) GetParameterizedHeaders() map[string][]ParameterizedHeader`

GetParameterizedHeaders returns the ParameterizedHeaders field if non-nil, zero value otherwise.

### GetParameterizedHeadersOk

`func (o *FormDataBodyPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool)`

GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterizedHeaders

`func (o *FormDataBodyPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader)`

SetParameterizedHeaders sets ParameterizedHeaders field to given value.

### HasParameterizedHeaders

`func (o *FormDataBodyPart) HasParameterizedHeaders() bool`

HasParameterizedHeaders returns a boolean if a field has been set.

### GetParent

`func (o *FormDataBodyPart) GetParent() MultiPart`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FormDataBodyPart) GetParentOk() (*MultiPart, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FormDataBodyPart) SetParent(v MultiPart)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FormDataBodyPart) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProviders

`func (o *FormDataBodyPart) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *FormDataBodyPart) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *FormDataBodyPart) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *FormDataBodyPart) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetSimple

`func (o *FormDataBodyPart) GetSimple() bool`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *FormDataBodyPart) GetSimpleOk() (*bool, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *FormDataBodyPart) SetSimple(v bool)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *FormDataBodyPart) HasSimple() bool`

HasSimple returns a boolean if a field has been set.

### GetValue

`func (o *FormDataBodyPart) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FormDataBodyPart) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FormDataBodyPart) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FormDataBodyPart) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


