# BTSplineDescription2118

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ControlPoints** | Pointer to **[]float64** |  | [optional] 
**Degree** | Pointer to **int32** |  | [optional] 
**Direction** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**DirectionOrientedWithFace** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**IsPeriodic** | Pointer to **bool** |  | [optional] 
**IsRational** | Pointer to **bool** |  | [optional] 
**Knots** | Pointer to **[]float64** |  | [optional] 
**Origin** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**Type** | Pointer to [**GBTCurveTypeEnum**](GBTCurveTypeEnum.md) |  | [optional] 

## Methods

### NewBTSplineDescription2118

`func NewBTSplineDescription2118() *BTSplineDescription2118`

NewBTSplineDescription2118 instantiates a new BTSplineDescription2118 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSplineDescription2118WithDefaults

`func NewBTSplineDescription2118WithDefaults() *BTSplineDescription2118`

NewBTSplineDescription2118WithDefaults instantiates a new BTSplineDescription2118 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTSplineDescription2118) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTSplineDescription2118) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTSplineDescription2118) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTSplineDescription2118) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetControlPoints

`func (o *BTSplineDescription2118) GetControlPoints() []float64`

GetControlPoints returns the ControlPoints field if non-nil, zero value otherwise.

### GetControlPointsOk

`func (o *BTSplineDescription2118) GetControlPointsOk() (*[]float64, bool)`

GetControlPointsOk returns a tuple with the ControlPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPoints

`func (o *BTSplineDescription2118) SetControlPoints(v []float64)`

SetControlPoints sets ControlPoints field to given value.

### HasControlPoints

`func (o *BTSplineDescription2118) HasControlPoints() bool`

HasControlPoints returns a boolean if a field has been set.

### GetDegree

`func (o *BTSplineDescription2118) GetDegree() int32`

GetDegree returns the Degree field if non-nil, zero value otherwise.

### GetDegreeOk

`func (o *BTSplineDescription2118) GetDegreeOk() (*int32, bool)`

GetDegreeOk returns a tuple with the Degree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegree

`func (o *BTSplineDescription2118) SetDegree(v int32)`

SetDegree sets Degree field to given value.

### HasDegree

`func (o *BTSplineDescription2118) HasDegree() bool`

HasDegree returns a boolean if a field has been set.

### GetDirection

`func (o *BTSplineDescription2118) GetDirection() BTVector3d389`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BTSplineDescription2118) GetDirectionOk() (*BTVector3d389, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BTSplineDescription2118) SetDirection(v BTVector3d389)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BTSplineDescription2118) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDirectionOrientedWithFace

`func (o *BTSplineDescription2118) GetDirectionOrientedWithFace() BTVector3d389`

GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field if non-nil, zero value otherwise.

### GetDirectionOrientedWithFaceOk

`func (o *BTSplineDescription2118) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool)`

GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionOrientedWithFace

`func (o *BTSplineDescription2118) SetDirectionOrientedWithFace(v BTVector3d389)`

SetDirectionOrientedWithFace sets DirectionOrientedWithFace field to given value.

### HasDirectionOrientedWithFace

`func (o *BTSplineDescription2118) HasDirectionOrientedWithFace() bool`

HasDirectionOrientedWithFace returns a boolean if a field has been set.

### GetIsPeriodic

`func (o *BTSplineDescription2118) GetIsPeriodic() bool`

GetIsPeriodic returns the IsPeriodic field if non-nil, zero value otherwise.

### GetIsPeriodicOk

`func (o *BTSplineDescription2118) GetIsPeriodicOk() (*bool, bool)`

GetIsPeriodicOk returns a tuple with the IsPeriodic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPeriodic

`func (o *BTSplineDescription2118) SetIsPeriodic(v bool)`

SetIsPeriodic sets IsPeriodic field to given value.

### HasIsPeriodic

`func (o *BTSplineDescription2118) HasIsPeriodic() bool`

HasIsPeriodic returns a boolean if a field has been set.

### GetIsRational

`func (o *BTSplineDescription2118) GetIsRational() bool`

GetIsRational returns the IsRational field if non-nil, zero value otherwise.

### GetIsRationalOk

`func (o *BTSplineDescription2118) GetIsRationalOk() (*bool, bool)`

GetIsRationalOk returns a tuple with the IsRational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRational

`func (o *BTSplineDescription2118) SetIsRational(v bool)`

SetIsRational sets IsRational field to given value.

### HasIsRational

`func (o *BTSplineDescription2118) HasIsRational() bool`

HasIsRational returns a boolean if a field has been set.

### GetKnots

`func (o *BTSplineDescription2118) GetKnots() []float64`

GetKnots returns the Knots field if non-nil, zero value otherwise.

### GetKnotsOk

`func (o *BTSplineDescription2118) GetKnotsOk() (*[]float64, bool)`

GetKnotsOk returns a tuple with the Knots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnots

`func (o *BTSplineDescription2118) SetKnots(v []float64)`

SetKnots sets Knots field to given value.

### HasKnots

`func (o *BTSplineDescription2118) HasKnots() bool`

HasKnots returns a boolean if a field has been set.

### GetOrigin

`func (o *BTSplineDescription2118) GetOrigin() BTVector3d389`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BTSplineDescription2118) GetOriginOk() (*BTVector3d389, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BTSplineDescription2118) SetOrigin(v BTVector3d389)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *BTSplineDescription2118) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetType

`func (o *BTSplineDescription2118) GetType() GBTCurveTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTSplineDescription2118) GetTypeOk() (*GBTCurveTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTSplineDescription2118) SetType(v GBTCurveTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *BTSplineDescription2118) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


