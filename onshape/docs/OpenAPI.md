# OpenAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**Components**](Components.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Info** | Pointer to [**Info**](Info.md) |  | [optional] 
**JsonSchemaDialect** | Pointer to **string** |  | [optional] 
**Openapi** | Pointer to **string** |  | [optional] 
**Paths** | Pointer to [**OpenAPIPaths**](OpenAPIPaths.md) |  | [optional] 
**Security** | Pointer to [**[]SecurityRequirement**](SecurityRequirement.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Webhooks** | Pointer to [**map[string]PathItem**](PathItem.md) |  | [optional] 

## Methods

### NewOpenAPI

`func NewOpenAPI() *OpenAPI`

NewOpenAPI instantiates a new OpenAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAPIWithDefaults

`func NewOpenAPIWithDefaults() *OpenAPI`

NewOpenAPIWithDefaults instantiates a new OpenAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *OpenAPI) GetComponents() Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *OpenAPI) GetComponentsOk() (*Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *OpenAPI) SetComponents(v Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *OpenAPI) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetExtensions

`func (o *OpenAPI) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *OpenAPI) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *OpenAPI) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *OpenAPI) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExternalDocs

`func (o *OpenAPI) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *OpenAPI) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *OpenAPI) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *OpenAPI) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetInfo

`func (o *OpenAPI) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OpenAPI) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OpenAPI) SetInfo(v Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OpenAPI) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJsonSchemaDialect

`func (o *OpenAPI) GetJsonSchemaDialect() string`

GetJsonSchemaDialect returns the JsonSchemaDialect field if non-nil, zero value otherwise.

### GetJsonSchemaDialectOk

`func (o *OpenAPI) GetJsonSchemaDialectOk() (*string, bool)`

GetJsonSchemaDialectOk returns a tuple with the JsonSchemaDialect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchemaDialect

`func (o *OpenAPI) SetJsonSchemaDialect(v string)`

SetJsonSchemaDialect sets JsonSchemaDialect field to given value.

### HasJsonSchemaDialect

`func (o *OpenAPI) HasJsonSchemaDialect() bool`

HasJsonSchemaDialect returns a boolean if a field has been set.

### GetOpenapi

`func (o *OpenAPI) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *OpenAPI) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *OpenAPI) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.

### HasOpenapi

`func (o *OpenAPI) HasOpenapi() bool`

HasOpenapi returns a boolean if a field has been set.

### GetPaths

`func (o *OpenAPI) GetPaths() OpenAPIPaths`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *OpenAPI) GetPathsOk() (*OpenAPIPaths, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *OpenAPI) SetPaths(v OpenAPIPaths)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *OpenAPI) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetSecurity

`func (o *OpenAPI) GetSecurity() []SecurityRequirement`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OpenAPI) GetSecurityOk() (*[]SecurityRequirement, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OpenAPI) SetSecurity(v []SecurityRequirement)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OpenAPI) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetServers

`func (o *OpenAPI) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OpenAPI) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OpenAPI) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OpenAPI) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTags

`func (o *OpenAPI) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OpenAPI) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OpenAPI) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OpenAPI) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWebhooks

`func (o *OpenAPI) GetWebhooks() map[string]PathItem`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *OpenAPI) GetWebhooksOk() (*map[string]PathItem, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *OpenAPI) SetWebhooks(v map[string]PathItem)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *OpenAPI) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


