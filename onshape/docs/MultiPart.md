# MultiPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyParts** | Pointer to [**[]BodyPart**](BodyPart.md) |  | [optional] 
**ContentDisposition** | Pointer to [**ContentDisposition**](ContentDisposition.md) |  | [optional] 
**Entity** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**Headers** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**MediaType** | Pointer to [**BodyPartMediaType**](BodyPart_mediaType.md) |  | [optional] 
**MessageBodyWorkers** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**ParameterizedHeaders** | Pointer to [**map[string][]ParameterizedHeader**](array.md) |  | [optional] 
**Parent** | Pointer to [**MultiPart**](MultiPart.md) |  | [optional] 
**Providers** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 

## Methods

### NewMultiPart

`func NewMultiPart() *MultiPart`

NewMultiPart instantiates a new MultiPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiPartWithDefaults

`func NewMultiPartWithDefaults() *MultiPart`

NewMultiPartWithDefaults instantiates a new MultiPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyParts

`func (o *MultiPart) GetBodyParts() []BodyPart`

GetBodyParts returns the BodyParts field if non-nil, zero value otherwise.

### GetBodyPartsOk

`func (o *MultiPart) GetBodyPartsOk() (*[]BodyPart, bool)`

GetBodyPartsOk returns a tuple with the BodyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyParts

`func (o *MultiPart) SetBodyParts(v []BodyPart)`

SetBodyParts sets BodyParts field to given value.

### HasBodyParts

`func (o *MultiPart) HasBodyParts() bool`

HasBodyParts returns a boolean if a field has been set.

### GetContentDisposition

`func (o *MultiPart) GetContentDisposition() ContentDisposition`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *MultiPart) GetContentDispositionOk() (*ContentDisposition, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *MultiPart) SetContentDisposition(v ContentDisposition)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *MultiPart) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetEntity

`func (o *MultiPart) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *MultiPart) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *MultiPart) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *MultiPart) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetHeaders

`func (o *MultiPart) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MultiPart) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MultiPart) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MultiPart) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMediaType

`func (o *MultiPart) GetMediaType() BodyPartMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MultiPart) GetMediaTypeOk() (*BodyPartMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MultiPart) SetMediaType(v BodyPartMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *MultiPart) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMessageBodyWorkers

`func (o *MultiPart) GetMessageBodyWorkers() map[string]interface{}`

GetMessageBodyWorkers returns the MessageBodyWorkers field if non-nil, zero value otherwise.

### GetMessageBodyWorkersOk

`func (o *MultiPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool)`

GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBodyWorkers

`func (o *MultiPart) SetMessageBodyWorkers(v map[string]interface{})`

SetMessageBodyWorkers sets MessageBodyWorkers field to given value.

### HasMessageBodyWorkers

`func (o *MultiPart) HasMessageBodyWorkers() bool`

HasMessageBodyWorkers returns a boolean if a field has been set.

### GetParameterizedHeaders

`func (o *MultiPart) GetParameterizedHeaders() map[string][]ParameterizedHeader`

GetParameterizedHeaders returns the ParameterizedHeaders field if non-nil, zero value otherwise.

### GetParameterizedHeadersOk

`func (o *MultiPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool)`

GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterizedHeaders

`func (o *MultiPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader)`

SetParameterizedHeaders sets ParameterizedHeaders field to given value.

### HasParameterizedHeaders

`func (o *MultiPart) HasParameterizedHeaders() bool`

HasParameterizedHeaders returns a boolean if a field has been set.

### GetParent

`func (o *MultiPart) GetParent() MultiPart`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MultiPart) GetParentOk() (*MultiPart, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MultiPart) SetParent(v MultiPart)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MultiPart) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProviders

`func (o *MultiPart) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *MultiPart) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *MultiPart) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *MultiPart) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


