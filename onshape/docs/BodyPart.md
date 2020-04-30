# BodyPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDisposition** | Pointer to [**ContentDisposition**](ContentDisposition.md) |  | [optional] 
**Entity** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**Headers** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**MediaType** | Pointer to [**BodyPartMediaType**](BodyPart_mediaType.md) |  | [optional] 
**MessageBodyWorkers** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**ParameterizedHeaders** | Pointer to [**map[string][]ParameterizedHeader**](array.md) |  | [optional] 
**Parent** | Pointer to [**MultiPart**](MultiPart.md) |  | [optional] 
**Providers** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 

## Methods

### NewBodyPart

`func NewBodyPart() *BodyPart`

NewBodyPart instantiates a new BodyPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyPartWithDefaults

`func NewBodyPartWithDefaults() *BodyPart`

NewBodyPartWithDefaults instantiates a new BodyPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDisposition

`func (o *BodyPart) GetContentDisposition() ContentDisposition`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *BodyPart) GetContentDispositionOk() (*ContentDisposition, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *BodyPart) SetContentDisposition(v ContentDisposition)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *BodyPart) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetEntity

`func (o *BodyPart) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *BodyPart) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *BodyPart) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *BodyPart) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetHeaders

`func (o *BodyPart) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BodyPart) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BodyPart) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BodyPart) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMediaType

`func (o *BodyPart) GetMediaType() BodyPartMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *BodyPart) GetMediaTypeOk() (*BodyPartMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *BodyPart) SetMediaType(v BodyPartMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *BodyPart) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMessageBodyWorkers

`func (o *BodyPart) GetMessageBodyWorkers() map[string]interface{}`

GetMessageBodyWorkers returns the MessageBodyWorkers field if non-nil, zero value otherwise.

### GetMessageBodyWorkersOk

`func (o *BodyPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool)`

GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBodyWorkers

`func (o *BodyPart) SetMessageBodyWorkers(v map[string]interface{})`

SetMessageBodyWorkers sets MessageBodyWorkers field to given value.

### HasMessageBodyWorkers

`func (o *BodyPart) HasMessageBodyWorkers() bool`

HasMessageBodyWorkers returns a boolean if a field has been set.

### GetParameterizedHeaders

`func (o *BodyPart) GetParameterizedHeaders() map[string][]ParameterizedHeader`

GetParameterizedHeaders returns the ParameterizedHeaders field if non-nil, zero value otherwise.

### GetParameterizedHeadersOk

`func (o *BodyPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool)`

GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterizedHeaders

`func (o *BodyPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader)`

SetParameterizedHeaders sets ParameterizedHeaders field to given value.

### HasParameterizedHeaders

`func (o *BodyPart) HasParameterizedHeaders() bool`

HasParameterizedHeaders returns a boolean if a field has been set.

### GetParent

`func (o *BodyPart) GetParent() MultiPart`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BodyPart) GetParentOk() (*MultiPart, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BodyPart) SetParent(v MultiPart)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BodyPart) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProviders

`func (o *BodyPart) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *BodyPart) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *BodyPart) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *BodyPart) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


