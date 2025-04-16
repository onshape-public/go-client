# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callbacks** | Pointer to [**map[string]Callback**](Callback.md) |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]Parameter**](Parameter.md) |  | [optional] 
**RequestBody** | Pointer to [**RequestBody**](RequestBody.md) |  | [optional] 
**Responses** | Pointer to [**map[string]ApiResponse**](ApiResponse.md) |  | [optional] 
**Security** | Pointer to [**[]SecurityRequirement**](SecurityRequirement.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbacks

`func (o *Operation) GetCallbacks() map[string]Callback`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *Operation) GetCallbacksOk() (*map[string]Callback, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *Operation) SetCallbacks(v map[string]Callback)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *Operation) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### GetDeprecated

`func (o *Operation) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Operation) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Operation) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Operation) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *Operation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Operation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Operation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Operation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensions

`func (o *Operation) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Operation) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Operation) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Operation) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExternalDocs

`func (o *Operation) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *Operation) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *Operation) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *Operation) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetOperationId

`func (o *Operation) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *Operation) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *Operation) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *Operation) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetParameters

`func (o *Operation) GetParameters() []Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Operation) GetParametersOk() (*[]Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Operation) SetParameters(v []Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Operation) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRequestBody

`func (o *Operation) GetRequestBody() RequestBody`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *Operation) GetRequestBodyOk() (*RequestBody, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *Operation) SetRequestBody(v RequestBody)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *Operation) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResponses

`func (o *Operation) GetResponses() map[string]ApiResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *Operation) GetResponsesOk() (*map[string]ApiResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *Operation) SetResponses(v map[string]ApiResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *Operation) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetSecurity

`func (o *Operation) GetSecurity() []SecurityRequirement`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *Operation) GetSecurityOk() (*[]SecurityRequirement, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *Operation) SetSecurity(v []SecurityRequirement)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *Operation) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetServers

`func (o *Operation) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Operation) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Operation) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Operation) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSummary

`func (o *Operation) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Operation) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Operation) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Operation) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *Operation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Operation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Operation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Operation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


