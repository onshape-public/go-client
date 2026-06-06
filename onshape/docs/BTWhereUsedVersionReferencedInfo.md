# BTWhereUsedVersionReferencedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**[]ConfigInfo**](ConfigInfo.md) | The configuration parameters of the referenced item in this version. | [optional] 
**CreatedAt** | Pointer to **JSONTime** | The timestamp when this version was created. | [optional] 
**Document** | Pointer to [**BTDocumentBaseSummaryInfo**](BTDocumentBaseSummaryInfo.md) |  | [optional] 
**ElementId** | Pointer to **string** | The element ID of the referring element in this version. | [optional] 
**ElementType** | Pointer to **int32** | The element type ordinal of the referring element. | [optional] 
**PartId** | Pointer to **string** | The part ID of the referenced item. | [optional] 
**PartName** | Pointer to **string** | The part name of the referenced item. | [optional] 
**PartNumber** | Pointer to **string** | The part number of the referenced item. | [optional] 
**Revision** | Pointer to **string** | The revision id of the referenced item. | [optional] 
**RevisionObsolete** | Pointer to **bool** | Whether the revision of this item is obsolete. | [optional] 
**Version** | Pointer to [**BTBaseInfo**](BTBaseInfo.md) |  | [optional] 

## Methods

### NewBTWhereUsedVersionReferencedInfo

`func NewBTWhereUsedVersionReferencedInfo() *BTWhereUsedVersionReferencedInfo`

NewBTWhereUsedVersionReferencedInfo instantiates a new BTWhereUsedVersionReferencedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWhereUsedVersionReferencedInfoWithDefaults

`func NewBTWhereUsedVersionReferencedInfoWithDefaults() *BTWhereUsedVersionReferencedInfo`

NewBTWhereUsedVersionReferencedInfoWithDefaults instantiates a new BTWhereUsedVersionReferencedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BTWhereUsedVersionReferencedInfo) GetConfiguration() []ConfigInfo`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTWhereUsedVersionReferencedInfo) GetConfigurationOk() (*[]ConfigInfo, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTWhereUsedVersionReferencedInfo) SetConfiguration(v []ConfigInfo)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTWhereUsedVersionReferencedInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTWhereUsedVersionReferencedInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTWhereUsedVersionReferencedInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTWhereUsedVersionReferencedInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTWhereUsedVersionReferencedInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocument

`func (o *BTWhereUsedVersionReferencedInfo) GetDocument() BTDocumentBaseSummaryInfo`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *BTWhereUsedVersionReferencedInfo) GetDocumentOk() (*BTDocumentBaseSummaryInfo, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *BTWhereUsedVersionReferencedInfo) SetDocument(v BTDocumentBaseSummaryInfo)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *BTWhereUsedVersionReferencedInfo) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetElementId

`func (o *BTWhereUsedVersionReferencedInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTWhereUsedVersionReferencedInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTWhereUsedVersionReferencedInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTWhereUsedVersionReferencedInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTWhereUsedVersionReferencedInfo) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTWhereUsedVersionReferencedInfo) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTWhereUsedVersionReferencedInfo) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTWhereUsedVersionReferencedInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetPartId

`func (o *BTWhereUsedVersionReferencedInfo) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTWhereUsedVersionReferencedInfo) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTWhereUsedVersionReferencedInfo) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTWhereUsedVersionReferencedInfo) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartName

`func (o *BTWhereUsedVersionReferencedInfo) GetPartName() string`

GetPartName returns the PartName field if non-nil, zero value otherwise.

### GetPartNameOk

`func (o *BTWhereUsedVersionReferencedInfo) GetPartNameOk() (*string, bool)`

GetPartNameOk returns a tuple with the PartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartName

`func (o *BTWhereUsedVersionReferencedInfo) SetPartName(v string)`

SetPartName sets PartName field to given value.

### HasPartName

`func (o *BTWhereUsedVersionReferencedInfo) HasPartName() bool`

HasPartName returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTWhereUsedVersionReferencedInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTWhereUsedVersionReferencedInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTWhereUsedVersionReferencedInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTWhereUsedVersionReferencedInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetRevision

`func (o *BTWhereUsedVersionReferencedInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTWhereUsedVersionReferencedInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTWhereUsedVersionReferencedInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTWhereUsedVersionReferencedInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionObsolete

`func (o *BTWhereUsedVersionReferencedInfo) GetRevisionObsolete() bool`

GetRevisionObsolete returns the RevisionObsolete field if non-nil, zero value otherwise.

### GetRevisionObsoleteOk

`func (o *BTWhereUsedVersionReferencedInfo) GetRevisionObsoleteOk() (*bool, bool)`

GetRevisionObsoleteOk returns a tuple with the RevisionObsolete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionObsolete

`func (o *BTWhereUsedVersionReferencedInfo) SetRevisionObsolete(v bool)`

SetRevisionObsolete sets RevisionObsolete field to given value.

### HasRevisionObsolete

`func (o *BTWhereUsedVersionReferencedInfo) HasRevisionObsolete() bool`

HasRevisionObsolete returns a boolean if a field has been set.

### GetVersion

`func (o *BTWhereUsedVersionReferencedInfo) GetVersion() BTBaseInfo`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTWhereUsedVersionReferencedInfo) GetVersionOk() (*BTBaseInfo, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTWhereUsedVersionReferencedInfo) SetVersion(v BTBaseInfo)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTWhereUsedVersionReferencedInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


