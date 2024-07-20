# BTMSketchConstraint2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ConstraintType** | Pointer to [**GBTConstraintType**](GBTConstraintType.md) |  | [optional] 
**DrivenDimension** | Pointer to **bool** |  | [optional] 
**HasOffsetData1** | Pointer to **bool** |  | [optional] 
**HasOffsetData2** | Pointer to **bool** |  | [optional] 
**HasPierceParameter** | Pointer to **bool** |  | [optional] 
**HelpParameters** | Pointer to **[]float64** |  | [optional] 
**OffsetDistance1** | Pointer to **float64** |  | [optional] 
**OffsetDistance2** | Pointer to **float64** |  | [optional] 
**OffsetOrientation1** | Pointer to **bool** |  | [optional] 
**OffsetOrientation2** | Pointer to **bool** |  | [optional] 
**PierceParameter** | Pointer to **float64** |  | [optional] 

## Methods

### NewBTMSketchConstraint2

`func NewBTMSketchConstraint2() *BTMSketchConstraint2`

NewBTMSketchConstraint2 instantiates a new BTMSketchConstraint2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMSketchConstraint2WithDefaults

`func NewBTMSketchConstraint2WithDefaults() *BTMSketchConstraint2`

NewBTMSketchConstraint2WithDefaults instantiates a new BTMSketchConstraint2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMSketchConstraint2) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMSketchConstraint2) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMSketchConstraint2) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMSketchConstraint2) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConstraintType

`func (o *BTMSketchConstraint2) GetConstraintType() GBTConstraintType`

GetConstraintType returns the ConstraintType field if non-nil, zero value otherwise.

### GetConstraintTypeOk

`func (o *BTMSketchConstraint2) GetConstraintTypeOk() (*GBTConstraintType, bool)`

GetConstraintTypeOk returns a tuple with the ConstraintType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintType

`func (o *BTMSketchConstraint2) SetConstraintType(v GBTConstraintType)`

SetConstraintType sets ConstraintType field to given value.

### HasConstraintType

`func (o *BTMSketchConstraint2) HasConstraintType() bool`

HasConstraintType returns a boolean if a field has been set.

### GetDrivenDimension

`func (o *BTMSketchConstraint2) GetDrivenDimension() bool`

GetDrivenDimension returns the DrivenDimension field if non-nil, zero value otherwise.

### GetDrivenDimensionOk

`func (o *BTMSketchConstraint2) GetDrivenDimensionOk() (*bool, bool)`

GetDrivenDimensionOk returns a tuple with the DrivenDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivenDimension

`func (o *BTMSketchConstraint2) SetDrivenDimension(v bool)`

SetDrivenDimension sets DrivenDimension field to given value.

### HasDrivenDimension

`func (o *BTMSketchConstraint2) HasDrivenDimension() bool`

HasDrivenDimension returns a boolean if a field has been set.

### GetHasOffsetData1

`func (o *BTMSketchConstraint2) GetHasOffsetData1() bool`

GetHasOffsetData1 returns the HasOffsetData1 field if non-nil, zero value otherwise.

### GetHasOffsetData1Ok

`func (o *BTMSketchConstraint2) GetHasOffsetData1Ok() (*bool, bool)`

GetHasOffsetData1Ok returns a tuple with the HasOffsetData1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOffsetData1

`func (o *BTMSketchConstraint2) SetHasOffsetData1(v bool)`

SetHasOffsetData1 sets HasOffsetData1 field to given value.

### HasHasOffsetData1

`func (o *BTMSketchConstraint2) HasHasOffsetData1() bool`

HasHasOffsetData1 returns a boolean if a field has been set.

### GetHasOffsetData2

`func (o *BTMSketchConstraint2) GetHasOffsetData2() bool`

GetHasOffsetData2 returns the HasOffsetData2 field if non-nil, zero value otherwise.

### GetHasOffsetData2Ok

`func (o *BTMSketchConstraint2) GetHasOffsetData2Ok() (*bool, bool)`

GetHasOffsetData2Ok returns a tuple with the HasOffsetData2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOffsetData2

`func (o *BTMSketchConstraint2) SetHasOffsetData2(v bool)`

SetHasOffsetData2 sets HasOffsetData2 field to given value.

### HasHasOffsetData2

`func (o *BTMSketchConstraint2) HasHasOffsetData2() bool`

HasHasOffsetData2 returns a boolean if a field has been set.

### GetHasPierceParameter

`func (o *BTMSketchConstraint2) GetHasPierceParameter() bool`

GetHasPierceParameter returns the HasPierceParameter field if non-nil, zero value otherwise.

### GetHasPierceParameterOk

`func (o *BTMSketchConstraint2) GetHasPierceParameterOk() (*bool, bool)`

GetHasPierceParameterOk returns a tuple with the HasPierceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPierceParameter

`func (o *BTMSketchConstraint2) SetHasPierceParameter(v bool)`

SetHasPierceParameter sets HasPierceParameter field to given value.

### HasHasPierceParameter

`func (o *BTMSketchConstraint2) HasHasPierceParameter() bool`

HasHasPierceParameter returns a boolean if a field has been set.

### GetHelpParameters

`func (o *BTMSketchConstraint2) GetHelpParameters() []float64`

GetHelpParameters returns the HelpParameters field if non-nil, zero value otherwise.

### GetHelpParametersOk

`func (o *BTMSketchConstraint2) GetHelpParametersOk() (*[]float64, bool)`

GetHelpParametersOk returns a tuple with the HelpParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpParameters

`func (o *BTMSketchConstraint2) SetHelpParameters(v []float64)`

SetHelpParameters sets HelpParameters field to given value.

### HasHelpParameters

`func (o *BTMSketchConstraint2) HasHelpParameters() bool`

HasHelpParameters returns a boolean if a field has been set.

### GetOffsetDistance1

`func (o *BTMSketchConstraint2) GetOffsetDistance1() float64`

GetOffsetDistance1 returns the OffsetDistance1 field if non-nil, zero value otherwise.

### GetOffsetDistance1Ok

`func (o *BTMSketchConstraint2) GetOffsetDistance1Ok() (*float64, bool)`

GetOffsetDistance1Ok returns a tuple with the OffsetDistance1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDistance1

`func (o *BTMSketchConstraint2) SetOffsetDistance1(v float64)`

SetOffsetDistance1 sets OffsetDistance1 field to given value.

### HasOffsetDistance1

`func (o *BTMSketchConstraint2) HasOffsetDistance1() bool`

HasOffsetDistance1 returns a boolean if a field has been set.

### GetOffsetDistance2

`func (o *BTMSketchConstraint2) GetOffsetDistance2() float64`

GetOffsetDistance2 returns the OffsetDistance2 field if non-nil, zero value otherwise.

### GetOffsetDistance2Ok

`func (o *BTMSketchConstraint2) GetOffsetDistance2Ok() (*float64, bool)`

GetOffsetDistance2Ok returns a tuple with the OffsetDistance2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDistance2

`func (o *BTMSketchConstraint2) SetOffsetDistance2(v float64)`

SetOffsetDistance2 sets OffsetDistance2 field to given value.

### HasOffsetDistance2

`func (o *BTMSketchConstraint2) HasOffsetDistance2() bool`

HasOffsetDistance2 returns a boolean if a field has been set.

### GetOffsetOrientation1

`func (o *BTMSketchConstraint2) GetOffsetOrientation1() bool`

GetOffsetOrientation1 returns the OffsetOrientation1 field if non-nil, zero value otherwise.

### GetOffsetOrientation1Ok

`func (o *BTMSketchConstraint2) GetOffsetOrientation1Ok() (*bool, bool)`

GetOffsetOrientation1Ok returns a tuple with the OffsetOrientation1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetOrientation1

`func (o *BTMSketchConstraint2) SetOffsetOrientation1(v bool)`

SetOffsetOrientation1 sets OffsetOrientation1 field to given value.

### HasOffsetOrientation1

`func (o *BTMSketchConstraint2) HasOffsetOrientation1() bool`

HasOffsetOrientation1 returns a boolean if a field has been set.

### GetOffsetOrientation2

`func (o *BTMSketchConstraint2) GetOffsetOrientation2() bool`

GetOffsetOrientation2 returns the OffsetOrientation2 field if non-nil, zero value otherwise.

### GetOffsetOrientation2Ok

`func (o *BTMSketchConstraint2) GetOffsetOrientation2Ok() (*bool, bool)`

GetOffsetOrientation2Ok returns a tuple with the OffsetOrientation2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetOrientation2

`func (o *BTMSketchConstraint2) SetOffsetOrientation2(v bool)`

SetOffsetOrientation2 sets OffsetOrientation2 field to given value.

### HasOffsetOrientation2

`func (o *BTMSketchConstraint2) HasOffsetOrientation2() bool`

HasOffsetOrientation2 returns a boolean if a field has been set.

### GetPierceParameter

`func (o *BTMSketchConstraint2) GetPierceParameter() float64`

GetPierceParameter returns the PierceParameter field if non-nil, zero value otherwise.

### GetPierceParameterOk

`func (o *BTMSketchConstraint2) GetPierceParameterOk() (*float64, bool)`

GetPierceParameterOk returns a tuple with the PierceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPierceParameter

`func (o *BTMSketchConstraint2) SetPierceParameter(v float64)`

SetPierceParameter sets PierceParameter field to given value.

### HasPierceParameter

`func (o *BTMSketchConstraint2) HasPierceParameter() bool`

HasPierceParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


