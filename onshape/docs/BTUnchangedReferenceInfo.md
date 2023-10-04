# BTUnchangedReferenceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataUnchanged** | Pointer to **bool** |  | [optional] 
**NodeIds** | Pointer to **[]string** |  | [optional] 
**ToRevision** | Pointer to [**BTRevisionInfo**](BTRevisionInfo.md) |  | [optional] 

## Methods

### NewBTUnchangedReferenceInfo

`func NewBTUnchangedReferenceInfo() *BTUnchangedReferenceInfo`

NewBTUnchangedReferenceInfo instantiates a new BTUnchangedReferenceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUnchangedReferenceInfoWithDefaults

`func NewBTUnchangedReferenceInfoWithDefaults() *BTUnchangedReferenceInfo`

NewBTUnchangedReferenceInfoWithDefaults instantiates a new BTUnchangedReferenceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataUnchanged

`func (o *BTUnchangedReferenceInfo) GetMetadataUnchanged() bool`

GetMetadataUnchanged returns the MetadataUnchanged field if non-nil, zero value otherwise.

### GetMetadataUnchangedOk

`func (o *BTUnchangedReferenceInfo) GetMetadataUnchangedOk() (*bool, bool)`

GetMetadataUnchangedOk returns a tuple with the MetadataUnchanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUnchanged

`func (o *BTUnchangedReferenceInfo) SetMetadataUnchanged(v bool)`

SetMetadataUnchanged sets MetadataUnchanged field to given value.

### HasMetadataUnchanged

`func (o *BTUnchangedReferenceInfo) HasMetadataUnchanged() bool`

HasMetadataUnchanged returns a boolean if a field has been set.

### GetNodeIds

`func (o *BTUnchangedReferenceInfo) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *BTUnchangedReferenceInfo) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *BTUnchangedReferenceInfo) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *BTUnchangedReferenceInfo) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### GetToRevision

`func (o *BTUnchangedReferenceInfo) GetToRevision() BTRevisionInfo`

GetToRevision returns the ToRevision field if non-nil, zero value otherwise.

### GetToRevisionOk

`func (o *BTUnchangedReferenceInfo) GetToRevisionOk() (*BTRevisionInfo, bool)`

GetToRevisionOk returns a tuple with the ToRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRevision

`func (o *BTUnchangedReferenceInfo) SetToRevision(v BTRevisionInfo)`

SetToRevision sets ToRevision field to given value.

### HasToRevision

`func (o *BTUnchangedReferenceInfo) HasToRevision() bool`

HasToRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


