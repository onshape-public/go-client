# BTEntityEdge30AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**CompressedPoints** | Pointer to [**BTImmutableByteArray**](BTImmutableByteArray.md) |  | [optional] 
**EdgeSmoothnessStatus** | Pointer to [**GBTEntityEdgeSmoothnessStatus**](GBTEntityEdgeSmoothnessStatus.md) |  | [optional] 
**EdgeType** | Pointer to [**GBTEdgeType**](GBTEdgeType.md) |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**IsInternalEdge** | Pointer to **bool** |  | [optional] 
**Points** | Pointer to [**BTImmutableFloatArray**](BTImmutableFloatArray.md) |  | [optional] 

## Methods

### NewBTEntityEdge30AllOf

`func NewBTEntityEdge30AllOf() *BTEntityEdge30AllOf`

NewBTEntityEdge30AllOf instantiates a new BTEntityEdge30AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTEntityEdge30AllOfWithDefaults

`func NewBTEntityEdge30AllOfWithDefaults() *BTEntityEdge30AllOf`

NewBTEntityEdge30AllOfWithDefaults instantiates a new BTEntityEdge30AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTEntityEdge30AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTEntityEdge30AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTEntityEdge30AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTEntityEdge30AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCompressedPoints

`func (o *BTEntityEdge30AllOf) GetCompressedPoints() BTImmutableByteArray`

GetCompressedPoints returns the CompressedPoints field if non-nil, zero value otherwise.

### GetCompressedPointsOk

`func (o *BTEntityEdge30AllOf) GetCompressedPointsOk() (*BTImmutableByteArray, bool)`

GetCompressedPointsOk returns a tuple with the CompressedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedPoints

`func (o *BTEntityEdge30AllOf) SetCompressedPoints(v BTImmutableByteArray)`

SetCompressedPoints sets CompressedPoints field to given value.

### HasCompressedPoints

`func (o *BTEntityEdge30AllOf) HasCompressedPoints() bool`

HasCompressedPoints returns a boolean if a field has been set.

### GetEdgeSmoothnessStatus

`func (o *BTEntityEdge30AllOf) GetEdgeSmoothnessStatus() GBTEntityEdgeSmoothnessStatus`

GetEdgeSmoothnessStatus returns the EdgeSmoothnessStatus field if non-nil, zero value otherwise.

### GetEdgeSmoothnessStatusOk

`func (o *BTEntityEdge30AllOf) GetEdgeSmoothnessStatusOk() (*GBTEntityEdgeSmoothnessStatus, bool)`

GetEdgeSmoothnessStatusOk returns a tuple with the EdgeSmoothnessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeSmoothnessStatus

`func (o *BTEntityEdge30AllOf) SetEdgeSmoothnessStatus(v GBTEntityEdgeSmoothnessStatus)`

SetEdgeSmoothnessStatus sets EdgeSmoothnessStatus field to given value.

### HasEdgeSmoothnessStatus

`func (o *BTEntityEdge30AllOf) HasEdgeSmoothnessStatus() bool`

HasEdgeSmoothnessStatus returns a boolean if a field has been set.

### GetEdgeType

`func (o *BTEntityEdge30AllOf) GetEdgeType() GBTEdgeType`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *BTEntityEdge30AllOf) GetEdgeTypeOk() (*GBTEdgeType, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *BTEntityEdge30AllOf) SetEdgeType(v GBTEdgeType)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *BTEntityEdge30AllOf) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### GetIsClosed

`func (o *BTEntityEdge30AllOf) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *BTEntityEdge30AllOf) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *BTEntityEdge30AllOf) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *BTEntityEdge30AllOf) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsInternalEdge

`func (o *BTEntityEdge30AllOf) GetIsInternalEdge() bool`

GetIsInternalEdge returns the IsInternalEdge field if non-nil, zero value otherwise.

### GetIsInternalEdgeOk

`func (o *BTEntityEdge30AllOf) GetIsInternalEdgeOk() (*bool, bool)`

GetIsInternalEdgeOk returns a tuple with the IsInternalEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalEdge

`func (o *BTEntityEdge30AllOf) SetIsInternalEdge(v bool)`

SetIsInternalEdge sets IsInternalEdge field to given value.

### HasIsInternalEdge

`func (o *BTEntityEdge30AllOf) HasIsInternalEdge() bool`

HasIsInternalEdge returns a boolean if a field has been set.

### GetPoints

`func (o *BTEntityEdge30AllOf) GetPoints() BTImmutableFloatArray`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *BTEntityEdge30AllOf) GetPointsOk() (*BTImmutableFloatArray, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *BTEntityEdge30AllOf) SetPoints(v BTImmutableFloatArray)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *BTEntityEdge30AllOf) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


