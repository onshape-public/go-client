# BTLoadDisplayData837

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ComponentValues** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**DirectionMateConnectorId** | Pointer to **string** |  | [optional] 
**FaceLoadDeterministicIds** | Pointer to **[]string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**IsDerivedFeature** | Pointer to **bool** |  | [optional] 
**IsDirectionFlipped** | Pointer to **bool** |  | [optional] 
**LoadType** | Pointer to [**GBTLoadType**](GBTLoadType.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Occurrence** | Pointer to [**BTOccurrence74**](BTOccurrence74.md) |  | [optional] 
**OwnerOccurrence** | Pointer to [**BTOccurrence74**](BTOccurrence74.md) |  | [optional] 
**Status** | Pointer to [**GBTAssemblyFeatureDisplayStatus**](GBTAssemblyFeatureDisplayStatus.md) |  | [optional] 

## Methods

### NewBTLoadDisplayData837

`func NewBTLoadDisplayData837() *BTLoadDisplayData837`

NewBTLoadDisplayData837 instantiates a new BTLoadDisplayData837 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTLoadDisplayData837WithDefaults

`func NewBTLoadDisplayData837WithDefaults() *BTLoadDisplayData837`

NewBTLoadDisplayData837WithDefaults instantiates a new BTLoadDisplayData837 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTLoadDisplayData837) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTLoadDisplayData837) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTLoadDisplayData837) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTLoadDisplayData837) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetComponentValues

`func (o *BTLoadDisplayData837) GetComponentValues() BTVector3d389`

GetComponentValues returns the ComponentValues field if non-nil, zero value otherwise.

### GetComponentValuesOk

`func (o *BTLoadDisplayData837) GetComponentValuesOk() (*BTVector3d389, bool)`

GetComponentValuesOk returns a tuple with the ComponentValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentValues

`func (o *BTLoadDisplayData837) SetComponentValues(v BTVector3d389)`

SetComponentValues sets ComponentValues field to given value.

### HasComponentValues

`func (o *BTLoadDisplayData837) HasComponentValues() bool`

HasComponentValues returns a boolean if a field has been set.

### GetDirectionMateConnectorId

`func (o *BTLoadDisplayData837) GetDirectionMateConnectorId() string`

GetDirectionMateConnectorId returns the DirectionMateConnectorId field if non-nil, zero value otherwise.

### GetDirectionMateConnectorIdOk

`func (o *BTLoadDisplayData837) GetDirectionMateConnectorIdOk() (*string, bool)`

GetDirectionMateConnectorIdOk returns a tuple with the DirectionMateConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionMateConnectorId

`func (o *BTLoadDisplayData837) SetDirectionMateConnectorId(v string)`

SetDirectionMateConnectorId sets DirectionMateConnectorId field to given value.

### HasDirectionMateConnectorId

`func (o *BTLoadDisplayData837) HasDirectionMateConnectorId() bool`

HasDirectionMateConnectorId returns a boolean if a field has been set.

### GetFaceLoadDeterministicIds

`func (o *BTLoadDisplayData837) GetFaceLoadDeterministicIds() []string`

GetFaceLoadDeterministicIds returns the FaceLoadDeterministicIds field if non-nil, zero value otherwise.

### GetFaceLoadDeterministicIdsOk

`func (o *BTLoadDisplayData837) GetFaceLoadDeterministicIdsOk() (*[]string, bool)`

GetFaceLoadDeterministicIdsOk returns a tuple with the FaceLoadDeterministicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceLoadDeterministicIds

`func (o *BTLoadDisplayData837) SetFaceLoadDeterministicIds(v []string)`

SetFaceLoadDeterministicIds sets FaceLoadDeterministicIds field to given value.

### HasFaceLoadDeterministicIds

`func (o *BTLoadDisplayData837) HasFaceLoadDeterministicIds() bool`

HasFaceLoadDeterministicIds returns a boolean if a field has been set.

### GetHidden

`func (o *BTLoadDisplayData837) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BTLoadDisplayData837) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BTLoadDisplayData837) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BTLoadDisplayData837) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIsDerivedFeature

`func (o *BTLoadDisplayData837) GetIsDerivedFeature() bool`

GetIsDerivedFeature returns the IsDerivedFeature field if non-nil, zero value otherwise.

### GetIsDerivedFeatureOk

`func (o *BTLoadDisplayData837) GetIsDerivedFeatureOk() (*bool, bool)`

GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerivedFeature

`func (o *BTLoadDisplayData837) SetIsDerivedFeature(v bool)`

SetIsDerivedFeature sets IsDerivedFeature field to given value.

### HasIsDerivedFeature

`func (o *BTLoadDisplayData837) HasIsDerivedFeature() bool`

HasIsDerivedFeature returns a boolean if a field has been set.

### GetIsDirectionFlipped

`func (o *BTLoadDisplayData837) GetIsDirectionFlipped() bool`

GetIsDirectionFlipped returns the IsDirectionFlipped field if non-nil, zero value otherwise.

### GetIsDirectionFlippedOk

`func (o *BTLoadDisplayData837) GetIsDirectionFlippedOk() (*bool, bool)`

GetIsDirectionFlippedOk returns a tuple with the IsDirectionFlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectionFlipped

`func (o *BTLoadDisplayData837) SetIsDirectionFlipped(v bool)`

SetIsDirectionFlipped sets IsDirectionFlipped field to given value.

### HasIsDirectionFlipped

`func (o *BTLoadDisplayData837) HasIsDirectionFlipped() bool`

HasIsDirectionFlipped returns a boolean if a field has been set.

### GetLoadType

`func (o *BTLoadDisplayData837) GetLoadType() GBTLoadType`

GetLoadType returns the LoadType field if non-nil, zero value otherwise.

### GetLoadTypeOk

`func (o *BTLoadDisplayData837) GetLoadTypeOk() (*GBTLoadType, bool)`

GetLoadTypeOk returns a tuple with the LoadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadType

`func (o *BTLoadDisplayData837) SetLoadType(v GBTLoadType)`

SetLoadType sets LoadType field to given value.

### HasLoadType

`func (o *BTLoadDisplayData837) HasLoadType() bool`

HasLoadType returns a boolean if a field has been set.

### GetNodeId

`func (o *BTLoadDisplayData837) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTLoadDisplayData837) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTLoadDisplayData837) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTLoadDisplayData837) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOccurrence

`func (o *BTLoadDisplayData837) GetOccurrence() BTOccurrence74`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *BTLoadDisplayData837) GetOccurrenceOk() (*BTOccurrence74, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *BTLoadDisplayData837) SetOccurrence(v BTOccurrence74)`

SetOccurrence sets Occurrence field to given value.

### HasOccurrence

`func (o *BTLoadDisplayData837) HasOccurrence() bool`

HasOccurrence returns a boolean if a field has been set.

### GetOwnerOccurrence

`func (o *BTLoadDisplayData837) GetOwnerOccurrence() BTOccurrence74`

GetOwnerOccurrence returns the OwnerOccurrence field if non-nil, zero value otherwise.

### GetOwnerOccurrenceOk

`func (o *BTLoadDisplayData837) GetOwnerOccurrenceOk() (*BTOccurrence74, bool)`

GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOccurrence

`func (o *BTLoadDisplayData837) SetOwnerOccurrence(v BTOccurrence74)`

SetOwnerOccurrence sets OwnerOccurrence field to given value.

### HasOwnerOccurrence

`func (o *BTLoadDisplayData837) HasOwnerOccurrence() bool`

HasOwnerOccurrence returns a boolean if a field has been set.

### GetStatus

`func (o *BTLoadDisplayData837) GetStatus() GBTAssemblyFeatureDisplayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BTLoadDisplayData837) GetStatusOk() (*GBTAssemblyFeatureDisplayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BTLoadDisplayData837) SetStatus(v GBTAssemblyFeatureDisplayStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BTLoadDisplayData837) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


