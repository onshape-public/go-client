# BTBExportMeshParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AngularTolerance** | Pointer to **float64** | Determines the maximum angular deviation, between the analytical surface and its triangulation. Lower values result in a finer geometry and higher values result in coarser geometry. | [optional] [default to 0.001]
**DistanceTolerance** | Pointer to **float64** | Determines the maximum distance deviation, between the analytical surface and its triangulation. Lower values result in a finer geometry and higher values result in coarser geometry. | [optional] [default to 0.001]
**MaximumChordLength** | Pointer to **float64** | Determines the maximum of a triangle edge length. Lower values result in a finer geometry and higher values result in coarser geometry. | [optional] [default to 0.01]
**Resolution** | Pointer to [**GBTExportResolution**](GBTExportResolution.md) |  | [optional] [default to GBTExportResolutionFine]
**Unit** | Pointer to [**GBTExportUnit**](GBTExportUnit.md) |  | [optional] [default to GBTExportUnitMeter]

## Methods

### NewBTBExportMeshParams

`func NewBTBExportMeshParams() *BTBExportMeshParams`

NewBTBExportMeshParams instantiates a new BTBExportMeshParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBExportMeshParamsWithDefaults

`func NewBTBExportMeshParamsWithDefaults() *BTBExportMeshParams`

NewBTBExportMeshParamsWithDefaults instantiates a new BTBExportMeshParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngularTolerance

`func (o *BTBExportMeshParams) GetAngularTolerance() float64`

GetAngularTolerance returns the AngularTolerance field if non-nil, zero value otherwise.

### GetAngularToleranceOk

`func (o *BTBExportMeshParams) GetAngularToleranceOk() (*float64, bool)`

GetAngularToleranceOk returns a tuple with the AngularTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngularTolerance

`func (o *BTBExportMeshParams) SetAngularTolerance(v float64)`

SetAngularTolerance sets AngularTolerance field to given value.

### HasAngularTolerance

`func (o *BTBExportMeshParams) HasAngularTolerance() bool`

HasAngularTolerance returns a boolean if a field has been set.

### GetDistanceTolerance

`func (o *BTBExportMeshParams) GetDistanceTolerance() float64`

GetDistanceTolerance returns the DistanceTolerance field if non-nil, zero value otherwise.

### GetDistanceToleranceOk

`func (o *BTBExportMeshParams) GetDistanceToleranceOk() (*float64, bool)`

GetDistanceToleranceOk returns a tuple with the DistanceTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceTolerance

`func (o *BTBExportMeshParams) SetDistanceTolerance(v float64)`

SetDistanceTolerance sets DistanceTolerance field to given value.

### HasDistanceTolerance

`func (o *BTBExportMeshParams) HasDistanceTolerance() bool`

HasDistanceTolerance returns a boolean if a field has been set.

### GetMaximumChordLength

`func (o *BTBExportMeshParams) GetMaximumChordLength() float64`

GetMaximumChordLength returns the MaximumChordLength field if non-nil, zero value otherwise.

### GetMaximumChordLengthOk

`func (o *BTBExportMeshParams) GetMaximumChordLengthOk() (*float64, bool)`

GetMaximumChordLengthOk returns a tuple with the MaximumChordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumChordLength

`func (o *BTBExportMeshParams) SetMaximumChordLength(v float64)`

SetMaximumChordLength sets MaximumChordLength field to given value.

### HasMaximumChordLength

`func (o *BTBExportMeshParams) HasMaximumChordLength() bool`

HasMaximumChordLength returns a boolean if a field has been set.

### GetResolution

`func (o *BTBExportMeshParams) GetResolution() GBTExportResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *BTBExportMeshParams) GetResolutionOk() (*GBTExportResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *BTBExportMeshParams) SetResolution(v GBTExportResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *BTBExportMeshParams) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetUnit

`func (o *BTBExportMeshParams) GetUnit() GBTExportUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BTBExportMeshParams) GetUnitOk() (*GBTExportUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BTBExportMeshParams) SetUnit(v GBTExportUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BTBExportMeshParams) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


