# BTAssemblyPartInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyType** | Pointer to [**BTAssemblyPartBodyType**](BTAssemblyPartBodyType.md) |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentMicroversion** | Pointer to **string** |  | [optional] 
**DocumentVersion** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**FullConfiguration** | Pointer to **string** |  | [optional] 
**IsStandardContent** | Pointer to **bool** |  | [optional] 
**MateConnectors** | Pointer to [**[]BTAssemblyMateConnectorInfo**](BTAssemblyMateConnectorInfo.md) |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAssemblyPartInfo

`func NewBTAssemblyPartInfo() *BTAssemblyPartInfo`

NewBTAssemblyPartInfo instantiates a new BTAssemblyPartInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyPartInfoWithDefaults

`func NewBTAssemblyPartInfoWithDefaults() *BTAssemblyPartInfo`

NewBTAssemblyPartInfoWithDefaults instantiates a new BTAssemblyPartInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyType

`func (o *BTAssemblyPartInfo) GetBodyType() BTAssemblyPartBodyType`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *BTAssemblyPartInfo) GetBodyTypeOk() (*BTAssemblyPartBodyType, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *BTAssemblyPartInfo) SetBodyType(v BTAssemblyPartBodyType)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *BTAssemblyPartInfo) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTAssemblyPartInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTAssemblyPartInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTAssemblyPartInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTAssemblyPartInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTAssemblyPartInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTAssemblyPartInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTAssemblyPartInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTAssemblyPartInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentMicroversion

`func (o *BTAssemblyPartInfo) GetDocumentMicroversion() string`

GetDocumentMicroversion returns the DocumentMicroversion field if non-nil, zero value otherwise.

### GetDocumentMicroversionOk

`func (o *BTAssemblyPartInfo) GetDocumentMicroversionOk() (*string, bool)`

GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversion

`func (o *BTAssemblyPartInfo) SetDocumentMicroversion(v string)`

SetDocumentMicroversion sets DocumentMicroversion field to given value.

### HasDocumentMicroversion

`func (o *BTAssemblyPartInfo) HasDocumentMicroversion() bool`

HasDocumentMicroversion returns a boolean if a field has been set.

### GetDocumentVersion

`func (o *BTAssemblyPartInfo) GetDocumentVersion() string`

GetDocumentVersion returns the DocumentVersion field if non-nil, zero value otherwise.

### GetDocumentVersionOk

`func (o *BTAssemblyPartInfo) GetDocumentVersionOk() (*string, bool)`

GetDocumentVersionOk returns a tuple with the DocumentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersion

`func (o *BTAssemblyPartInfo) SetDocumentVersion(v string)`

SetDocumentVersion sets DocumentVersion field to given value.

### HasDocumentVersion

`func (o *BTAssemblyPartInfo) HasDocumentVersion() bool`

HasDocumentVersion returns a boolean if a field has been set.

### GetElementId

`func (o *BTAssemblyPartInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTAssemblyPartInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTAssemblyPartInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTAssemblyPartInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFullConfiguration

`func (o *BTAssemblyPartInfo) GetFullConfiguration() string`

GetFullConfiguration returns the FullConfiguration field if non-nil, zero value otherwise.

### GetFullConfigurationOk

`func (o *BTAssemblyPartInfo) GetFullConfigurationOk() (*string, bool)`

GetFullConfigurationOk returns a tuple with the FullConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullConfiguration

`func (o *BTAssemblyPartInfo) SetFullConfiguration(v string)`

SetFullConfiguration sets FullConfiguration field to given value.

### HasFullConfiguration

`func (o *BTAssemblyPartInfo) HasFullConfiguration() bool`

HasFullConfiguration returns a boolean if a field has been set.

### GetIsStandardContent

`func (o *BTAssemblyPartInfo) GetIsStandardContent() bool`

GetIsStandardContent returns the IsStandardContent field if non-nil, zero value otherwise.

### GetIsStandardContentOk

`func (o *BTAssemblyPartInfo) GetIsStandardContentOk() (*bool, bool)`

GetIsStandardContentOk returns a tuple with the IsStandardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandardContent

`func (o *BTAssemblyPartInfo) SetIsStandardContent(v bool)`

SetIsStandardContent sets IsStandardContent field to given value.

### HasIsStandardContent

`func (o *BTAssemblyPartInfo) HasIsStandardContent() bool`

HasIsStandardContent returns a boolean if a field has been set.

### GetMateConnectors

`func (o *BTAssemblyPartInfo) GetMateConnectors() []BTAssemblyMateConnectorInfo`

GetMateConnectors returns the MateConnectors field if non-nil, zero value otherwise.

### GetMateConnectorsOk

`func (o *BTAssemblyPartInfo) GetMateConnectorsOk() (*[]BTAssemblyMateConnectorInfo, bool)`

GetMateConnectorsOk returns a tuple with the MateConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateConnectors

`func (o *BTAssemblyPartInfo) SetMateConnectors(v []BTAssemblyMateConnectorInfo)`

SetMateConnectors sets MateConnectors field to given value.

### HasMateConnectors

`func (o *BTAssemblyPartInfo) HasMateConnectors() bool`

HasMateConnectors returns a boolean if a field has been set.

### GetPartId

`func (o *BTAssemblyPartInfo) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTAssemblyPartInfo) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTAssemblyPartInfo) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTAssemblyPartInfo) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTAssemblyPartInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTAssemblyPartInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTAssemblyPartInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTAssemblyPartInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetRevision

`func (o *BTAssemblyPartInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTAssemblyPartInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTAssemblyPartInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTAssemblyPartInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


