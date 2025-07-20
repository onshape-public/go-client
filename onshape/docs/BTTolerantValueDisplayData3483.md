# BTTolerantValueDisplayData3483

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**FitClass** | Pointer to **string** |  | [optional] 
**IsAngle** | Pointer to **bool** |  | [optional] 
**IsDefined** | Pointer to **bool** |  | [optional] 
**LowerTolerance** | Pointer to **float64** |  | [optional] 
**NominalValue** | Pointer to **float64** |  | [optional] 
**Precision** | Pointer to [**GBTTolerancePrecision**](GBTTolerancePrecision.md) |  | [optional] 
**ToleranceType** | Pointer to [**GBTToleranceType**](GBTToleranceType.md) |  | [optional] 
**UpperTolerance** | Pointer to **float64** |  | [optional] 

## Methods

### NewBTTolerantValueDisplayData3483

`func NewBTTolerantValueDisplayData3483() *BTTolerantValueDisplayData3483`

NewBTTolerantValueDisplayData3483 instantiates a new BTTolerantValueDisplayData3483 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTolerantValueDisplayData3483WithDefaults

`func NewBTTolerantValueDisplayData3483WithDefaults() *BTTolerantValueDisplayData3483`

NewBTTolerantValueDisplayData3483WithDefaults instantiates a new BTTolerantValueDisplayData3483 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTTolerantValueDisplayData3483) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTTolerantValueDisplayData3483) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTTolerantValueDisplayData3483) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTTolerantValueDisplayData3483) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFitClass

`func (o *BTTolerantValueDisplayData3483) GetFitClass() string`

GetFitClass returns the FitClass field if non-nil, zero value otherwise.

### GetFitClassOk

`func (o *BTTolerantValueDisplayData3483) GetFitClassOk() (*string, bool)`

GetFitClassOk returns a tuple with the FitClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitClass

`func (o *BTTolerantValueDisplayData3483) SetFitClass(v string)`

SetFitClass sets FitClass field to given value.

### HasFitClass

`func (o *BTTolerantValueDisplayData3483) HasFitClass() bool`

HasFitClass returns a boolean if a field has been set.

### GetIsAngle

`func (o *BTTolerantValueDisplayData3483) GetIsAngle() bool`

GetIsAngle returns the IsAngle field if non-nil, zero value otherwise.

### GetIsAngleOk

`func (o *BTTolerantValueDisplayData3483) GetIsAngleOk() (*bool, bool)`

GetIsAngleOk returns a tuple with the IsAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAngle

`func (o *BTTolerantValueDisplayData3483) SetIsAngle(v bool)`

SetIsAngle sets IsAngle field to given value.

### HasIsAngle

`func (o *BTTolerantValueDisplayData3483) HasIsAngle() bool`

HasIsAngle returns a boolean if a field has been set.

### GetIsDefined

`func (o *BTTolerantValueDisplayData3483) GetIsDefined() bool`

GetIsDefined returns the IsDefined field if non-nil, zero value otherwise.

### GetIsDefinedOk

`func (o *BTTolerantValueDisplayData3483) GetIsDefinedOk() (*bool, bool)`

GetIsDefinedOk returns a tuple with the IsDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefined

`func (o *BTTolerantValueDisplayData3483) SetIsDefined(v bool)`

SetIsDefined sets IsDefined field to given value.

### HasIsDefined

`func (o *BTTolerantValueDisplayData3483) HasIsDefined() bool`

HasIsDefined returns a boolean if a field has been set.

### GetLowerTolerance

`func (o *BTTolerantValueDisplayData3483) GetLowerTolerance() float64`

GetLowerTolerance returns the LowerTolerance field if non-nil, zero value otherwise.

### GetLowerToleranceOk

`func (o *BTTolerantValueDisplayData3483) GetLowerToleranceOk() (*float64, bool)`

GetLowerToleranceOk returns a tuple with the LowerTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerTolerance

`func (o *BTTolerantValueDisplayData3483) SetLowerTolerance(v float64)`

SetLowerTolerance sets LowerTolerance field to given value.

### HasLowerTolerance

`func (o *BTTolerantValueDisplayData3483) HasLowerTolerance() bool`

HasLowerTolerance returns a boolean if a field has been set.

### GetNominalValue

`func (o *BTTolerantValueDisplayData3483) GetNominalValue() float64`

GetNominalValue returns the NominalValue field if non-nil, zero value otherwise.

### GetNominalValueOk

`func (o *BTTolerantValueDisplayData3483) GetNominalValueOk() (*float64, bool)`

GetNominalValueOk returns a tuple with the NominalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominalValue

`func (o *BTTolerantValueDisplayData3483) SetNominalValue(v float64)`

SetNominalValue sets NominalValue field to given value.

### HasNominalValue

`func (o *BTTolerantValueDisplayData3483) HasNominalValue() bool`

HasNominalValue returns a boolean if a field has been set.

### GetPrecision

`func (o *BTTolerantValueDisplayData3483) GetPrecision() GBTTolerancePrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *BTTolerantValueDisplayData3483) GetPrecisionOk() (*GBTTolerancePrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *BTTolerantValueDisplayData3483) SetPrecision(v GBTTolerancePrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *BTTolerantValueDisplayData3483) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetToleranceType

`func (o *BTTolerantValueDisplayData3483) GetToleranceType() GBTToleranceType`

GetToleranceType returns the ToleranceType field if non-nil, zero value otherwise.

### GetToleranceTypeOk

`func (o *BTTolerantValueDisplayData3483) GetToleranceTypeOk() (*GBTToleranceType, bool)`

GetToleranceTypeOk returns a tuple with the ToleranceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToleranceType

`func (o *BTTolerantValueDisplayData3483) SetToleranceType(v GBTToleranceType)`

SetToleranceType sets ToleranceType field to given value.

### HasToleranceType

`func (o *BTTolerantValueDisplayData3483) HasToleranceType() bool`

HasToleranceType returns a boolean if a field has been set.

### GetUpperTolerance

`func (o *BTTolerantValueDisplayData3483) GetUpperTolerance() float64`

GetUpperTolerance returns the UpperTolerance field if non-nil, zero value otherwise.

### GetUpperToleranceOk

`func (o *BTTolerantValueDisplayData3483) GetUpperToleranceOk() (*float64, bool)`

GetUpperToleranceOk returns a tuple with the UpperTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperTolerance

`func (o *BTTolerantValueDisplayData3483) SetUpperTolerance(v float64)`

SetUpperTolerance sets UpperTolerance field to given value.

### HasUpperTolerance

`func (o *BTTolerantValueDisplayData3483) HasUpperTolerance() bool`

HasUpperTolerance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


