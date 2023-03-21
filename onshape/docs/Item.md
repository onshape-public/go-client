# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationTarget** | Pointer to [**BTApplicationTargetInfo**](BTApplicationTargetInfo.md) |  | [optional] 
**BaseHref** | Pointer to **string** |  | [optional] [readonly] 
**DataType** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 
**EncodedConfiguration** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JsonType** | **string** |  | 
**PartId** | Pointer to **string** |  | [optional] 
**PartName** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**RevisionId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 

## Methods

### NewItem

`func NewItem(jsonType string, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationTarget

`func (o *Item) GetApplicationTarget() BTApplicationTargetInfo`

GetApplicationTarget returns the ApplicationTarget field if non-nil, zero value otherwise.

### GetApplicationTargetOk

`func (o *Item) GetApplicationTargetOk() (*BTApplicationTargetInfo, bool)`

GetApplicationTargetOk returns a tuple with the ApplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTarget

`func (o *Item) SetApplicationTarget(v BTApplicationTargetInfo)`

SetApplicationTarget sets ApplicationTarget field to given value.

### HasApplicationTarget

`func (o *Item) HasApplicationTarget() bool`

HasApplicationTarget returns a boolean if a field has been set.

### GetBaseHref

`func (o *Item) GetBaseHref() string`

GetBaseHref returns the BaseHref field if non-nil, zero value otherwise.

### GetBaseHrefOk

`func (o *Item) GetBaseHrefOk() (*string, bool)`

GetBaseHrefOk returns a tuple with the BaseHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseHref

`func (o *Item) SetBaseHref(v string)`

SetBaseHref sets BaseHref field to given value.

### HasBaseHref

`func (o *Item) HasBaseHref() bool`

HasBaseHref returns a boolean if a field has been set.

### GetDataType

`func (o *Item) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Item) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Item) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *Item) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDocumentId

`func (o *Item) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *Item) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *Item) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *Item) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *Item) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *Item) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *Item) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *Item) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *Item) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *Item) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *Item) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *Item) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetEncodedConfiguration

`func (o *Item) GetEncodedConfiguration() string`

GetEncodedConfiguration returns the EncodedConfiguration field if non-nil, zero value otherwise.

### GetEncodedConfigurationOk

`func (o *Item) GetEncodedConfigurationOk() (*string, bool)`

GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedConfiguration

`func (o *Item) SetEncodedConfiguration(v string)`

SetEncodedConfiguration sets EncodedConfiguration field to given value.

### HasEncodedConfiguration

`func (o *Item) HasEncodedConfiguration() bool`

HasEncodedConfiguration returns a boolean if a field has been set.

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Item) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJsonType

`func (o *Item) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *Item) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *Item) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.


### GetPartId

`func (o *Item) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *Item) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *Item) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *Item) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartName

`func (o *Item) GetPartName() string`

GetPartName returns the PartName field if non-nil, zero value otherwise.

### GetPartNameOk

`func (o *Item) GetPartNameOk() (*string, bool)`

GetPartNameOk returns a tuple with the PartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartName

`func (o *Item) SetPartName(v string)`

SetPartName sets PartName field to given value.

### HasPartName

`func (o *Item) HasPartName() bool`

HasPartName returns a boolean if a field has been set.

### GetPartNumber

`func (o *Item) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *Item) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *Item) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *Item) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetRevision

`func (o *Item) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Item) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Item) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Item) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionId

`func (o *Item) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *Item) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *Item) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *Item) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetState

`func (o *Item) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Item) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Item) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *Item) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersionId

`func (o *Item) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Item) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Item) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *Item) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *Item) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *Item) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *Item) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *Item) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


