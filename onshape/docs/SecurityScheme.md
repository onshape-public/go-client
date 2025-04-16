# SecurityScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BearerFormat** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Flows** | Pointer to [**OAuthFlows**](OAuthFlows.md) |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**In** | Pointer to [**In**](In.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpenIdConnectUrl** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**Type**](Type.md) |  | [optional] 

## Methods

### NewSecurityScheme

`func NewSecurityScheme() *SecurityScheme`

NewSecurityScheme instantiates a new SecurityScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemeWithDefaults

`func NewSecuritySchemeWithDefaults() *SecurityScheme`

NewSecuritySchemeWithDefaults instantiates a new SecurityScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBearerFormat

`func (o *SecurityScheme) GetBearerFormat() string`

GetBearerFormat returns the BearerFormat field if non-nil, zero value otherwise.

### GetBearerFormatOk

`func (o *SecurityScheme) GetBearerFormatOk() (*string, bool)`

GetBearerFormatOk returns a tuple with the BearerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerFormat

`func (o *SecurityScheme) SetBearerFormat(v string)`

SetBearerFormat sets BearerFormat field to given value.

### HasBearerFormat

`func (o *SecurityScheme) HasBearerFormat() bool`

HasBearerFormat returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensions

`func (o *SecurityScheme) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SecurityScheme) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SecurityScheme) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SecurityScheme) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetFlows

`func (o *SecurityScheme) GetFlows() OAuthFlows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *SecurityScheme) GetFlowsOk() (*OAuthFlows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *SecurityScheme) SetFlows(v OAuthFlows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *SecurityScheme) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetGetref

`func (o *SecurityScheme) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *SecurityScheme) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *SecurityScheme) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *SecurityScheme) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetIn

`func (o *SecurityScheme) GetIn() In`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *SecurityScheme) GetInOk() (*In, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *SecurityScheme) SetIn(v In)`

SetIn sets In field to given value.

### HasIn

`func (o *SecurityScheme) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetName

`func (o *SecurityScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityScheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityScheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenIdConnectUrl

`func (o *SecurityScheme) GetOpenIdConnectUrl() string`

GetOpenIdConnectUrl returns the OpenIdConnectUrl field if non-nil, zero value otherwise.

### GetOpenIdConnectUrlOk

`func (o *SecurityScheme) GetOpenIdConnectUrlOk() (*string, bool)`

GetOpenIdConnectUrlOk returns a tuple with the OpenIdConnectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIdConnectUrl

`func (o *SecurityScheme) SetOpenIdConnectUrl(v string)`

SetOpenIdConnectUrl sets OpenIdConnectUrl field to given value.

### HasOpenIdConnectUrl

`func (o *SecurityScheme) HasOpenIdConnectUrl() bool`

HasOpenIdConnectUrl returns a boolean if a field has been set.

### GetScheme

`func (o *SecurityScheme) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *SecurityScheme) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *SecurityScheme) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *SecurityScheme) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetType

`func (o *SecurityScheme) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityScheme) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityScheme) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityScheme) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


