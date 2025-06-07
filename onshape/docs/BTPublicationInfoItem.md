# BTPublicationInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationTarget** | Pointer to [**BTApplicationTargetInfo**](BTApplicationTargetInfo.md) |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to [**GBTElementType**](GBTElementType.md) |  | [optional] 
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

### NewBTPublicationInfoItem

`func NewBTPublicationInfoItem(jsonType string, ) *BTPublicationInfoItem`

NewBTPublicationInfoItem instantiates a new BTPublicationInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPublicationInfoItemWithDefaults

`func NewBTPublicationInfoItemWithDefaults() *BTPublicationInfoItem`

NewBTPublicationInfoItemWithDefaults instantiates a new BTPublicationInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationTarget

`func (o *BTPublicationInfoItem) GetApplicationTarget() BTApplicationTargetInfo`

GetApplicationTarget returns the ApplicationTarget field if non-nil, zero value otherwise.

### GetApplicationTargetOk

`func (o *BTPublicationInfoItem) GetApplicationTargetOk() (*BTApplicationTargetInfo, bool)`

GetApplicationTargetOk returns a tuple with the ApplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTarget

`func (o *BTPublicationInfoItem) SetApplicationTarget(v BTApplicationTargetInfo)`

SetApplicationTarget sets ApplicationTarget field to given value.

### HasApplicationTarget

`func (o *BTPublicationInfoItem) HasApplicationTarget() bool`

HasApplicationTarget returns a boolean if a field has been set.

### GetDataType

`func (o *BTPublicationInfoItem) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BTPublicationInfoItem) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BTPublicationInfoItem) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *BTPublicationInfoItem) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTPublicationInfoItem) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTPublicationInfoItem) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTPublicationInfoItem) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTPublicationInfoItem) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTPublicationInfoItem) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTPublicationInfoItem) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTPublicationInfoItem) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTPublicationInfoItem) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTPublicationInfoItem) GetElementType() GBTElementType`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTPublicationInfoItem) GetElementTypeOk() (*GBTElementType, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTPublicationInfoItem) SetElementType(v GBTElementType)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTPublicationInfoItem) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetEncodedConfiguration

`func (o *BTPublicationInfoItem) GetEncodedConfiguration() string`

GetEncodedConfiguration returns the EncodedConfiguration field if non-nil, zero value otherwise.

### GetEncodedConfigurationOk

`func (o *BTPublicationInfoItem) GetEncodedConfigurationOk() (*string, bool)`

GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedConfiguration

`func (o *BTPublicationInfoItem) SetEncodedConfiguration(v string)`

SetEncodedConfiguration sets EncodedConfiguration field to given value.

### HasEncodedConfiguration

`func (o *BTPublicationInfoItem) HasEncodedConfiguration() bool`

HasEncodedConfiguration returns a boolean if a field has been set.

### GetId

`func (o *BTPublicationInfoItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPublicationInfoItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPublicationInfoItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPublicationInfoItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJsonType

`func (o *BTPublicationInfoItem) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *BTPublicationInfoItem) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *BTPublicationInfoItem) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.


### GetPartId

`func (o *BTPublicationInfoItem) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTPublicationInfoItem) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTPublicationInfoItem) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTPublicationInfoItem) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartName

`func (o *BTPublicationInfoItem) GetPartName() string`

GetPartName returns the PartName field if non-nil, zero value otherwise.

### GetPartNameOk

`func (o *BTPublicationInfoItem) GetPartNameOk() (*string, bool)`

GetPartNameOk returns a tuple with the PartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartName

`func (o *BTPublicationInfoItem) SetPartName(v string)`

SetPartName sets PartName field to given value.

### HasPartName

`func (o *BTPublicationInfoItem) HasPartName() bool`

HasPartName returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTPublicationInfoItem) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTPublicationInfoItem) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTPublicationInfoItem) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTPublicationInfoItem) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetRevision

`func (o *BTPublicationInfoItem) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTPublicationInfoItem) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTPublicationInfoItem) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTPublicationInfoItem) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionId

`func (o *BTPublicationInfoItem) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *BTPublicationInfoItem) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *BTPublicationInfoItem) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *BTPublicationInfoItem) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetState

`func (o *BTPublicationInfoItem) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTPublicationInfoItem) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTPublicationInfoItem) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTPublicationInfoItem) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersionId

`func (o *BTPublicationInfoItem) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTPublicationInfoItem) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTPublicationInfoItem) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTPublicationInfoItem) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *BTPublicationInfoItem) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *BTPublicationInfoItem) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *BTPublicationInfoItem) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *BTPublicationInfoItem) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


