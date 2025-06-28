# BTNextPartNumberParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]BTCategoryParam**](BTCategoryParam.md) |  | [optional] 
**Configuration** | Pointer to **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [optional] 
**DocumentId** | Pointer to **string** | Document ID | [optional] 
**ElementId** | Pointer to **string** | Element (tab) ID | [optional] 
**ElementType** | Pointer to **int32** | Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**NumberSchemeResourceTypeId** | Pointer to **string** | Must be one of: &#x60;8c96700620f77935a0b2cddc&#x60;: Part Studio, assembly, or drawing, &#x60;29cd738cc6a8819fe84864e0&#x60;: Non-geometric items, &#x60;10f29fc285510ebc648108e6&#x60;: Standard content | [optional] 
**PartId** | Pointer to **string** | Part ID | [optional] 
**PartNumber** | Pointer to **string** | Current part number | [optional] 
**VersionId** | Pointer to **string** | Version ID | [optional] 
**WorkspaceId** | Pointer to **string** | Workspace ID | [optional] 

## Methods

### NewBTNextPartNumberParam

`func NewBTNextPartNumberParam() *BTNextPartNumberParam`

NewBTNextPartNumberParam instantiates a new BTNextPartNumberParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTNextPartNumberParamWithDefaults

`func NewBTNextPartNumberParamWithDefaults() *BTNextPartNumberParam`

NewBTNextPartNumberParamWithDefaults instantiates a new BTNextPartNumberParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *BTNextPartNumberParam) GetCategories() []BTCategoryParam`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BTNextPartNumberParam) GetCategoriesOk() (*[]BTCategoryParam, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BTNextPartNumberParam) SetCategories(v []BTCategoryParam)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *BTNextPartNumberParam) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTNextPartNumberParam) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTNextPartNumberParam) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTNextPartNumberParam) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTNextPartNumberParam) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTNextPartNumberParam) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTNextPartNumberParam) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTNextPartNumberParam) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTNextPartNumberParam) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTNextPartNumberParam) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTNextPartNumberParam) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTNextPartNumberParam) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTNextPartNumberParam) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTNextPartNumberParam) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTNextPartNumberParam) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTNextPartNumberParam) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTNextPartNumberParam) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetId

`func (o *BTNextPartNumberParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTNextPartNumberParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTNextPartNumberParam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTNextPartNumberParam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimeType

`func (o *BTNextPartNumberParam) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTNextPartNumberParam) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTNextPartNumberParam) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTNextPartNumberParam) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetNumberSchemeResourceTypeId

`func (o *BTNextPartNumberParam) GetNumberSchemeResourceTypeId() string`

GetNumberSchemeResourceTypeId returns the NumberSchemeResourceTypeId field if non-nil, zero value otherwise.

### GetNumberSchemeResourceTypeIdOk

`func (o *BTNextPartNumberParam) GetNumberSchemeResourceTypeIdOk() (*string, bool)`

GetNumberSchemeResourceTypeIdOk returns a tuple with the NumberSchemeResourceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberSchemeResourceTypeId

`func (o *BTNextPartNumberParam) SetNumberSchemeResourceTypeId(v string)`

SetNumberSchemeResourceTypeId sets NumberSchemeResourceTypeId field to given value.

### HasNumberSchemeResourceTypeId

`func (o *BTNextPartNumberParam) HasNumberSchemeResourceTypeId() bool`

HasNumberSchemeResourceTypeId returns a boolean if a field has been set.

### GetPartId

`func (o *BTNextPartNumberParam) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTNextPartNumberParam) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTNextPartNumberParam) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTNextPartNumberParam) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTNextPartNumberParam) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTNextPartNumberParam) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTNextPartNumberParam) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTNextPartNumberParam) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetVersionId

`func (o *BTNextPartNumberParam) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTNextPartNumberParam) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTNextPartNumberParam) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTNextPartNumberParam) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTNextPartNumberParam) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTNextPartNumberParam) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTNextPartNumberParam) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTNextPartNumberParam) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


