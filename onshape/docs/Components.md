# Components

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callbacks** | Pointer to [**map[string]Callback**](Callback.md) |  | [optional] 
**Examples** | Pointer to [**map[string]Example**](Example.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Headers** | Pointer to [**map[string]Header**](Header.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Parameters** | Pointer to [**map[string]Parameter**](Parameter.md) |  | [optional] 
**PathItems** | Pointer to [**map[string]PathItem**](PathItem.md) |  | [optional] 
**RequestBodies** | Pointer to [**map[string]RequestBody**](RequestBody.md) |  | [optional] 
**Responses** | Pointer to [**map[string]ApiResponse**](ApiResponse.md) |  | [optional] 
**Schemas** | Pointer to [**map[string]Schema**](Schema.md) |  | [optional] 
**SecuritySchemes** | Pointer to [**map[string]SecurityScheme**](SecurityScheme.md) |  | [optional] 

## Methods

### NewComponents

`func NewComponents() *Components`

NewComponents instantiates a new Components object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentsWithDefaults

`func NewComponentsWithDefaults() *Components`

NewComponentsWithDefaults instantiates a new Components object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbacks

`func (o *Components) GetCallbacks() map[string]Callback`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *Components) GetCallbacksOk() (*map[string]Callback, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *Components) SetCallbacks(v map[string]Callback)`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *Components) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### GetExamples

`func (o *Components) GetExamples() map[string]Example`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *Components) GetExamplesOk() (*map[string]Example, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *Components) SetExamples(v map[string]Example)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *Components) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetExtensions

`func (o *Components) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Components) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Components) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Components) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetHeaders

`func (o *Components) GetHeaders() map[string]Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Components) GetHeadersOk() (*map[string]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Components) SetHeaders(v map[string]Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Components) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetLinks

`func (o *Components) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Components) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Components) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Components) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetParameters

`func (o *Components) GetParameters() map[string]Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Components) GetParametersOk() (*map[string]Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Components) SetParameters(v map[string]Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Components) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPathItems

`func (o *Components) GetPathItems() map[string]PathItem`

GetPathItems returns the PathItems field if non-nil, zero value otherwise.

### GetPathItemsOk

`func (o *Components) GetPathItemsOk() (*map[string]PathItem, bool)`

GetPathItemsOk returns a tuple with the PathItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathItems

`func (o *Components) SetPathItems(v map[string]PathItem)`

SetPathItems sets PathItems field to given value.

### HasPathItems

`func (o *Components) HasPathItems() bool`

HasPathItems returns a boolean if a field has been set.

### GetRequestBodies

`func (o *Components) GetRequestBodies() map[string]RequestBody`

GetRequestBodies returns the RequestBodies field if non-nil, zero value otherwise.

### GetRequestBodiesOk

`func (o *Components) GetRequestBodiesOk() (*map[string]RequestBody, bool)`

GetRequestBodiesOk returns a tuple with the RequestBodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodies

`func (o *Components) SetRequestBodies(v map[string]RequestBody)`

SetRequestBodies sets RequestBodies field to given value.

### HasRequestBodies

`func (o *Components) HasRequestBodies() bool`

HasRequestBodies returns a boolean if a field has been set.

### GetResponses

`func (o *Components) GetResponses() map[string]ApiResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *Components) GetResponsesOk() (*map[string]ApiResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *Components) SetResponses(v map[string]ApiResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *Components) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetSchemas

`func (o *Components) GetSchemas() map[string]Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Components) GetSchemasOk() (*map[string]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Components) SetSchemas(v map[string]Schema)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *Components) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetSecuritySchemes

`func (o *Components) GetSecuritySchemes() map[string]SecurityScheme`

GetSecuritySchemes returns the SecuritySchemes field if non-nil, zero value otherwise.

### GetSecuritySchemesOk

`func (o *Components) GetSecuritySchemesOk() (*map[string]SecurityScheme, bool)`

GetSecuritySchemesOk returns a tuple with the SecuritySchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySchemes

`func (o *Components) SetSecuritySchemes(v map[string]SecurityScheme)`

SetSecuritySchemes sets SecuritySchemes field to given value.

### HasSecuritySchemes

`func (o *Components) HasSecuritySchemes() bool`

HasSecuritySchemes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


