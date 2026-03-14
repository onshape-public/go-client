# BTReleasePackageItemParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **string** | Encoded configuration string for the item. See [encodeConfigurationMap](#/Element/encodeConfigurationMap). | [optional] 
**DocumentId** | Pointer to **string** | Document containing the item. | [optional] 
**ElementId** | Pointer to **string** | Element containing the item. | [optional] 
**ElementType** | Pointer to **int32** | Type of element, which can be: &#x60;0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob&#x60; | [optional] 
**FlatPartId** | Pointer to **string** | Flat part ID. | [optional] 
**Href** | Pointer to **string** | Href link to the item. | [optional] 
**Id** | Pointer to **string** | ID of the item. Use this to reference the item when updating it. See [getReleasePackage](#/ReleasePackage/getReleasePackage). | [optional] 
**ParentId** | Pointer to **string** | ID of the parent item, if any. | [optional] 
**PartId** | Pointer to **string** | ID of the part to add to the release candidate. | [optional] 
**PartIdentity** | Pointer to **string** | Part identity string for the part to add to the release candidate. | [optional] 
**PartNumber** | Pointer to **string** | Part number. See [nextNumbers](#/NumberingScheme/nextNumbers). | [optional] 
**Properties** | Pointer to [**[]BTPropertyValueParam**](BTPropertyValueParam.md) | List of property values associated with the item. | [optional] 
**RevisionId** | Pointer to **string** | Revision ID of the item. Use when adding items to an obsoletion candidate. | [optional] 
**VersionId** | Pointer to **string** | Version containing the item. | [optional] 
**WorkspaceId** | Pointer to **string** | Workspace containing the item. | [optional] 

## Methods

### NewBTReleasePackageItemParams

`func NewBTReleasePackageItemParams() *BTReleasePackageItemParams`

NewBTReleasePackageItemParams instantiates a new BTReleasePackageItemParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTReleasePackageItemParamsWithDefaults

`func NewBTReleasePackageItemParamsWithDefaults() *BTReleasePackageItemParams`

NewBTReleasePackageItemParamsWithDefaults instantiates a new BTReleasePackageItemParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTReleasePackageItemParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTReleasePackageItemParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTReleasePackageItemParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTReleasePackageItemParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTReleasePackageItemParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTReleasePackageItemParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTReleasePackageItemParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTReleasePackageItemParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTReleasePackageItemParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTReleasePackageItemParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTReleasePackageItemParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTReleasePackageItemParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTReleasePackageItemParams) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTReleasePackageItemParams) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTReleasePackageItemParams) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTReleasePackageItemParams) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetFlatPartId

`func (o *BTReleasePackageItemParams) GetFlatPartId() string`

GetFlatPartId returns the FlatPartId field if non-nil, zero value otherwise.

### GetFlatPartIdOk

`func (o *BTReleasePackageItemParams) GetFlatPartIdOk() (*string, bool)`

GetFlatPartIdOk returns a tuple with the FlatPartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatPartId

`func (o *BTReleasePackageItemParams) SetFlatPartId(v string)`

SetFlatPartId sets FlatPartId field to given value.

### HasFlatPartId

`func (o *BTReleasePackageItemParams) HasFlatPartId() bool`

HasFlatPartId returns a boolean if a field has been set.

### GetHref

`func (o *BTReleasePackageItemParams) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTReleasePackageItemParams) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTReleasePackageItemParams) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTReleasePackageItemParams) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTReleasePackageItemParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTReleasePackageItemParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTReleasePackageItemParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTReleasePackageItemParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *BTReleasePackageItemParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTReleasePackageItemParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTReleasePackageItemParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTReleasePackageItemParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartId

`func (o *BTReleasePackageItemParams) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTReleasePackageItemParams) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTReleasePackageItemParams) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTReleasePackageItemParams) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTReleasePackageItemParams) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTReleasePackageItemParams) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTReleasePackageItemParams) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTReleasePackageItemParams) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTReleasePackageItemParams) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTReleasePackageItemParams) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTReleasePackageItemParams) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTReleasePackageItemParams) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProperties

`func (o *BTReleasePackageItemParams) GetProperties() []BTPropertyValueParam`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTReleasePackageItemParams) GetPropertiesOk() (*[]BTPropertyValueParam, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTReleasePackageItemParams) SetProperties(v []BTPropertyValueParam)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTReleasePackageItemParams) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRevisionId

`func (o *BTReleasePackageItemParams) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *BTReleasePackageItemParams) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *BTReleasePackageItemParams) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *BTReleasePackageItemParams) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTReleasePackageItemParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTReleasePackageItemParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTReleasePackageItemParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTReleasePackageItemParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTReleasePackageItemParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTReleasePackageItemParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTReleasePackageItemParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTReleasePackageItemParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


