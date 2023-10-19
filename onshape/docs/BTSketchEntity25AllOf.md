# BTSketchEntity25AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**IsConstruction** | Pointer to **bool** |  | [optional] 
**IsFromSplineControlPolygon** | Pointer to **bool** |  | [optional] 
**IsFromSplineHandle** | Pointer to **bool** |  | [optional] 
**IsTextStroke** | Pointer to **bool** |  | [optional] 
**IsUserPoint** | Pointer to **bool** |  | [optional] 
**SketchCurveType** | Pointer to [**GBTSketchCurveType**](GBTSketchCurveType.md) |  | [optional] 
**SketchEntityId** | Pointer to **string** |  | [optional] 
**SolveStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTSketchEntity25AllOf

`func NewBTSketchEntity25AllOf() *BTSketchEntity25AllOf`

NewBTSketchEntity25AllOf instantiates a new BTSketchEntity25AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSketchEntity25AllOfWithDefaults

`func NewBTSketchEntity25AllOfWithDefaults() *BTSketchEntity25AllOf`

NewBTSketchEntity25AllOfWithDefaults instantiates a new BTSketchEntity25AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTSketchEntity25AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTSketchEntity25AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTSketchEntity25AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTSketchEntity25AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetIsConstruction

`func (o *BTSketchEntity25AllOf) GetIsConstruction() bool`

GetIsConstruction returns the IsConstruction field if non-nil, zero value otherwise.

### GetIsConstructionOk

`func (o *BTSketchEntity25AllOf) GetIsConstructionOk() (*bool, bool)`

GetIsConstructionOk returns a tuple with the IsConstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstruction

`func (o *BTSketchEntity25AllOf) SetIsConstruction(v bool)`

SetIsConstruction sets IsConstruction field to given value.

### HasIsConstruction

`func (o *BTSketchEntity25AllOf) HasIsConstruction() bool`

HasIsConstruction returns a boolean if a field has been set.

### GetIsFromSplineControlPolygon

`func (o *BTSketchEntity25AllOf) GetIsFromSplineControlPolygon() bool`

GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field if non-nil, zero value otherwise.

### GetIsFromSplineControlPolygonOk

`func (o *BTSketchEntity25AllOf) GetIsFromSplineControlPolygonOk() (*bool, bool)`

GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSplineControlPolygon

`func (o *BTSketchEntity25AllOf) SetIsFromSplineControlPolygon(v bool)`

SetIsFromSplineControlPolygon sets IsFromSplineControlPolygon field to given value.

### HasIsFromSplineControlPolygon

`func (o *BTSketchEntity25AllOf) HasIsFromSplineControlPolygon() bool`

HasIsFromSplineControlPolygon returns a boolean if a field has been set.

### GetIsFromSplineHandle

`func (o *BTSketchEntity25AllOf) GetIsFromSplineHandle() bool`

GetIsFromSplineHandle returns the IsFromSplineHandle field if non-nil, zero value otherwise.

### GetIsFromSplineHandleOk

`func (o *BTSketchEntity25AllOf) GetIsFromSplineHandleOk() (*bool, bool)`

GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSplineHandle

`func (o *BTSketchEntity25AllOf) SetIsFromSplineHandle(v bool)`

SetIsFromSplineHandle sets IsFromSplineHandle field to given value.

### HasIsFromSplineHandle

`func (o *BTSketchEntity25AllOf) HasIsFromSplineHandle() bool`

HasIsFromSplineHandle returns a boolean if a field has been set.

### GetIsTextStroke

`func (o *BTSketchEntity25AllOf) GetIsTextStroke() bool`

GetIsTextStroke returns the IsTextStroke field if non-nil, zero value otherwise.

### GetIsTextStrokeOk

`func (o *BTSketchEntity25AllOf) GetIsTextStrokeOk() (*bool, bool)`

GetIsTextStrokeOk returns a tuple with the IsTextStroke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTextStroke

`func (o *BTSketchEntity25AllOf) SetIsTextStroke(v bool)`

SetIsTextStroke sets IsTextStroke field to given value.

### HasIsTextStroke

`func (o *BTSketchEntity25AllOf) HasIsTextStroke() bool`

HasIsTextStroke returns a boolean if a field has been set.

### GetIsUserPoint

`func (o *BTSketchEntity25AllOf) GetIsUserPoint() bool`

GetIsUserPoint returns the IsUserPoint field if non-nil, zero value otherwise.

### GetIsUserPointOk

`func (o *BTSketchEntity25AllOf) GetIsUserPointOk() (*bool, bool)`

GetIsUserPointOk returns a tuple with the IsUserPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserPoint

`func (o *BTSketchEntity25AllOf) SetIsUserPoint(v bool)`

SetIsUserPoint sets IsUserPoint field to given value.

### HasIsUserPoint

`func (o *BTSketchEntity25AllOf) HasIsUserPoint() bool`

HasIsUserPoint returns a boolean if a field has been set.

### GetSketchCurveType

`func (o *BTSketchEntity25AllOf) GetSketchCurveType() GBTSketchCurveType`

GetSketchCurveType returns the SketchCurveType field if non-nil, zero value otherwise.

### GetSketchCurveTypeOk

`func (o *BTSketchEntity25AllOf) GetSketchCurveTypeOk() (*GBTSketchCurveType, bool)`

GetSketchCurveTypeOk returns a tuple with the SketchCurveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchCurveType

`func (o *BTSketchEntity25AllOf) SetSketchCurveType(v GBTSketchCurveType)`

SetSketchCurveType sets SketchCurveType field to given value.

### HasSketchCurveType

`func (o *BTSketchEntity25AllOf) HasSketchCurveType() bool`

HasSketchCurveType returns a boolean if a field has been set.

### GetSketchEntityId

`func (o *BTSketchEntity25AllOf) GetSketchEntityId() string`

GetSketchEntityId returns the SketchEntityId field if non-nil, zero value otherwise.

### GetSketchEntityIdOk

`func (o *BTSketchEntity25AllOf) GetSketchEntityIdOk() (*string, bool)`

GetSketchEntityIdOk returns a tuple with the SketchEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchEntityId

`func (o *BTSketchEntity25AllOf) SetSketchEntityId(v string)`

SetSketchEntityId sets SketchEntityId field to given value.

### HasSketchEntityId

`func (o *BTSketchEntity25AllOf) HasSketchEntityId() bool`

HasSketchEntityId returns a boolean if a field has been set.

### GetSolveStatus

`func (o *BTSketchEntity25AllOf) GetSolveStatus() int32`

GetSolveStatus returns the SolveStatus field if non-nil, zero value otherwise.

### GetSolveStatusOk

`func (o *BTSketchEntity25AllOf) GetSolveStatusOk() (*int32, bool)`

GetSolveStatusOk returns a tuple with the SolveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolveStatus

`func (o *BTSketchEntity25AllOf) SetSolveStatus(v int32)`

SetSolveStatus sets SolveStatus field to given value.

### HasSolveStatus

`func (o *BTSketchEntity25AllOf) HasSolveStatus() bool`

HasSolveStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


