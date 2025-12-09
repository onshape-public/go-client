# BTDimensionDisplayData323

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**CharacteristicId** | Pointer to **string** |  | [optional] 
**CoordinateSystem** | Pointer to [**BTMatrix3x3340**](BTMatrix3x3340.md) |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**FitClass** | Pointer to **string** |  | [optional] 
**HasMaximumLimit** | Pointer to **bool** |  | [optional] 
**HasMinimumLimit** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsAnnotationDimension** | Pointer to **bool** |  | [optional] 
**IsAssociatedWithFlat** | Pointer to **bool** |  | [optional] 
**IsDriven** | Pointer to **bool** |  | [optional] 
**IsOverDefined** | Pointer to **bool** |  | [optional] 
**LowerTolerance** | Pointer to **float64** |  | [optional] 
**MaximumLimit** | Pointer to **float64** |  | [optional] 
**MinimumLimit** | Pointer to **float64** |  | [optional] 
**ParameterId** | Pointer to **string** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PlaneMatrix** | Pointer to [**BTBSMatrix386**](BTBSMatrix386.md) |  | [optional] 
**Precision** | Pointer to [**GBTTolerancePrecision**](GBTTolerancePrecision.md) |  | [optional] 
**ToleranceType** | Pointer to [**GBTToleranceType**](GBTToleranceType.md) |  | [optional] 
**UpperTolerance** | Pointer to **float64** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewBTDimensionDisplayData323

`func NewBTDimensionDisplayData323() *BTDimensionDisplayData323`

NewBTDimensionDisplayData323 instantiates a new BTDimensionDisplayData323 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDimensionDisplayData323WithDefaults

`func NewBTDimensionDisplayData323WithDefaults() *BTDimensionDisplayData323`

NewBTDimensionDisplayData323WithDefaults instantiates a new BTDimensionDisplayData323 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTDimensionDisplayData323) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTDimensionDisplayData323) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTDimensionDisplayData323) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTDimensionDisplayData323) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCharacteristicId

`func (o *BTDimensionDisplayData323) GetCharacteristicId() string`

GetCharacteristicId returns the CharacteristicId field if non-nil, zero value otherwise.

### GetCharacteristicIdOk

`func (o *BTDimensionDisplayData323) GetCharacteristicIdOk() (*string, bool)`

GetCharacteristicIdOk returns a tuple with the CharacteristicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicId

`func (o *BTDimensionDisplayData323) SetCharacteristicId(v string)`

SetCharacteristicId sets CharacteristicId field to given value.

### HasCharacteristicId

`func (o *BTDimensionDisplayData323) HasCharacteristicId() bool`

HasCharacteristicId returns a boolean if a field has been set.

### GetCoordinateSystem

`func (o *BTDimensionDisplayData323) GetCoordinateSystem() BTMatrix3x3340`

GetCoordinateSystem returns the CoordinateSystem field if non-nil, zero value otherwise.

### GetCoordinateSystemOk

`func (o *BTDimensionDisplayData323) GetCoordinateSystemOk() (*BTMatrix3x3340, bool)`

GetCoordinateSystemOk returns a tuple with the CoordinateSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinateSystem

`func (o *BTDimensionDisplayData323) SetCoordinateSystem(v BTMatrix3x3340)`

SetCoordinateSystem sets CoordinateSystem field to given value.

### HasCoordinateSystem

`func (o *BTDimensionDisplayData323) HasCoordinateSystem() bool`

HasCoordinateSystem returns a boolean if a field has been set.

### GetFeatureId

`func (o *BTDimensionDisplayData323) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *BTDimensionDisplayData323) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *BTDimensionDisplayData323) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *BTDimensionDisplayData323) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFitClass

`func (o *BTDimensionDisplayData323) GetFitClass() string`

GetFitClass returns the FitClass field if non-nil, zero value otherwise.

### GetFitClassOk

`func (o *BTDimensionDisplayData323) GetFitClassOk() (*string, bool)`

GetFitClassOk returns a tuple with the FitClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitClass

`func (o *BTDimensionDisplayData323) SetFitClass(v string)`

SetFitClass sets FitClass field to given value.

### HasFitClass

`func (o *BTDimensionDisplayData323) HasFitClass() bool`

HasFitClass returns a boolean if a field has been set.

### GetHasMaximumLimit

`func (o *BTDimensionDisplayData323) GetHasMaximumLimit() bool`

GetHasMaximumLimit returns the HasMaximumLimit field if non-nil, zero value otherwise.

### GetHasMaximumLimitOk

`func (o *BTDimensionDisplayData323) GetHasMaximumLimitOk() (*bool, bool)`

GetHasMaximumLimitOk returns a tuple with the HasMaximumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMaximumLimit

`func (o *BTDimensionDisplayData323) SetHasMaximumLimit(v bool)`

SetHasMaximumLimit sets HasMaximumLimit field to given value.

### HasHasMaximumLimit

`func (o *BTDimensionDisplayData323) HasHasMaximumLimit() bool`

HasHasMaximumLimit returns a boolean if a field has been set.

### GetHasMinimumLimit

`func (o *BTDimensionDisplayData323) GetHasMinimumLimit() bool`

GetHasMinimumLimit returns the HasMinimumLimit field if non-nil, zero value otherwise.

### GetHasMinimumLimitOk

`func (o *BTDimensionDisplayData323) GetHasMinimumLimitOk() (*bool, bool)`

GetHasMinimumLimitOk returns a tuple with the HasMinimumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMinimumLimit

`func (o *BTDimensionDisplayData323) SetHasMinimumLimit(v bool)`

SetHasMinimumLimit sets HasMinimumLimit field to given value.

### HasHasMinimumLimit

`func (o *BTDimensionDisplayData323) HasHasMinimumLimit() bool`

HasHasMinimumLimit returns a boolean if a field has been set.

### GetId

`func (o *BTDimensionDisplayData323) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDimensionDisplayData323) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDimensionDisplayData323) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDimensionDisplayData323) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAnnotationDimension

`func (o *BTDimensionDisplayData323) GetIsAnnotationDimension() bool`

GetIsAnnotationDimension returns the IsAnnotationDimension field if non-nil, zero value otherwise.

### GetIsAnnotationDimensionOk

`func (o *BTDimensionDisplayData323) GetIsAnnotationDimensionOk() (*bool, bool)`

GetIsAnnotationDimensionOk returns a tuple with the IsAnnotationDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnotationDimension

`func (o *BTDimensionDisplayData323) SetIsAnnotationDimension(v bool)`

SetIsAnnotationDimension sets IsAnnotationDimension field to given value.

### HasIsAnnotationDimension

`func (o *BTDimensionDisplayData323) HasIsAnnotationDimension() bool`

HasIsAnnotationDimension returns a boolean if a field has been set.

### GetIsAssociatedWithFlat

`func (o *BTDimensionDisplayData323) GetIsAssociatedWithFlat() bool`

GetIsAssociatedWithFlat returns the IsAssociatedWithFlat field if non-nil, zero value otherwise.

### GetIsAssociatedWithFlatOk

`func (o *BTDimensionDisplayData323) GetIsAssociatedWithFlatOk() (*bool, bool)`

GetIsAssociatedWithFlatOk returns a tuple with the IsAssociatedWithFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssociatedWithFlat

`func (o *BTDimensionDisplayData323) SetIsAssociatedWithFlat(v bool)`

SetIsAssociatedWithFlat sets IsAssociatedWithFlat field to given value.

### HasIsAssociatedWithFlat

`func (o *BTDimensionDisplayData323) HasIsAssociatedWithFlat() bool`

HasIsAssociatedWithFlat returns a boolean if a field has been set.

### GetIsDriven

`func (o *BTDimensionDisplayData323) GetIsDriven() bool`

GetIsDriven returns the IsDriven field if non-nil, zero value otherwise.

### GetIsDrivenOk

`func (o *BTDimensionDisplayData323) GetIsDrivenOk() (*bool, bool)`

GetIsDrivenOk returns a tuple with the IsDriven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDriven

`func (o *BTDimensionDisplayData323) SetIsDriven(v bool)`

SetIsDriven sets IsDriven field to given value.

### HasIsDriven

`func (o *BTDimensionDisplayData323) HasIsDriven() bool`

HasIsDriven returns a boolean if a field has been set.

### GetIsOverDefined

`func (o *BTDimensionDisplayData323) GetIsOverDefined() bool`

GetIsOverDefined returns the IsOverDefined field if non-nil, zero value otherwise.

### GetIsOverDefinedOk

`func (o *BTDimensionDisplayData323) GetIsOverDefinedOk() (*bool, bool)`

GetIsOverDefinedOk returns a tuple with the IsOverDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverDefined

`func (o *BTDimensionDisplayData323) SetIsOverDefined(v bool)`

SetIsOverDefined sets IsOverDefined field to given value.

### HasIsOverDefined

`func (o *BTDimensionDisplayData323) HasIsOverDefined() bool`

HasIsOverDefined returns a boolean if a field has been set.

### GetLowerTolerance

`func (o *BTDimensionDisplayData323) GetLowerTolerance() float64`

GetLowerTolerance returns the LowerTolerance field if non-nil, zero value otherwise.

### GetLowerToleranceOk

`func (o *BTDimensionDisplayData323) GetLowerToleranceOk() (*float64, bool)`

GetLowerToleranceOk returns a tuple with the LowerTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerTolerance

`func (o *BTDimensionDisplayData323) SetLowerTolerance(v float64)`

SetLowerTolerance sets LowerTolerance field to given value.

### HasLowerTolerance

`func (o *BTDimensionDisplayData323) HasLowerTolerance() bool`

HasLowerTolerance returns a boolean if a field has been set.

### GetMaximumLimit

`func (o *BTDimensionDisplayData323) GetMaximumLimit() float64`

GetMaximumLimit returns the MaximumLimit field if non-nil, zero value otherwise.

### GetMaximumLimitOk

`func (o *BTDimensionDisplayData323) GetMaximumLimitOk() (*float64, bool)`

GetMaximumLimitOk returns a tuple with the MaximumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLimit

`func (o *BTDimensionDisplayData323) SetMaximumLimit(v float64)`

SetMaximumLimit sets MaximumLimit field to given value.

### HasMaximumLimit

`func (o *BTDimensionDisplayData323) HasMaximumLimit() bool`

HasMaximumLimit returns a boolean if a field has been set.

### GetMinimumLimit

`func (o *BTDimensionDisplayData323) GetMinimumLimit() float64`

GetMinimumLimit returns the MinimumLimit field if non-nil, zero value otherwise.

### GetMinimumLimitOk

`func (o *BTDimensionDisplayData323) GetMinimumLimitOk() (*float64, bool)`

GetMinimumLimitOk returns a tuple with the MinimumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumLimit

`func (o *BTDimensionDisplayData323) SetMinimumLimit(v float64)`

SetMinimumLimit sets MinimumLimit field to given value.

### HasMinimumLimit

`func (o *BTDimensionDisplayData323) HasMinimumLimit() bool`

HasMinimumLimit returns a boolean if a field has been set.

### GetParameterId

`func (o *BTDimensionDisplayData323) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *BTDimensionDisplayData323) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *BTDimensionDisplayData323) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.

### HasParameterId

`func (o *BTDimensionDisplayData323) HasParameterId() bool`

HasParameterId returns a boolean if a field has been set.

### GetPartId

`func (o *BTDimensionDisplayData323) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTDimensionDisplayData323) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTDimensionDisplayData323) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTDimensionDisplayData323) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPlaneMatrix

`func (o *BTDimensionDisplayData323) GetPlaneMatrix() BTBSMatrix386`

GetPlaneMatrix returns the PlaneMatrix field if non-nil, zero value otherwise.

### GetPlaneMatrixOk

`func (o *BTDimensionDisplayData323) GetPlaneMatrixOk() (*BTBSMatrix386, bool)`

GetPlaneMatrixOk returns a tuple with the PlaneMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaneMatrix

`func (o *BTDimensionDisplayData323) SetPlaneMatrix(v BTBSMatrix386)`

SetPlaneMatrix sets PlaneMatrix field to given value.

### HasPlaneMatrix

`func (o *BTDimensionDisplayData323) HasPlaneMatrix() bool`

HasPlaneMatrix returns a boolean if a field has been set.

### GetPrecision

`func (o *BTDimensionDisplayData323) GetPrecision() GBTTolerancePrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *BTDimensionDisplayData323) GetPrecisionOk() (*GBTTolerancePrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *BTDimensionDisplayData323) SetPrecision(v GBTTolerancePrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *BTDimensionDisplayData323) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetToleranceType

`func (o *BTDimensionDisplayData323) GetToleranceType() GBTToleranceType`

GetToleranceType returns the ToleranceType field if non-nil, zero value otherwise.

### GetToleranceTypeOk

`func (o *BTDimensionDisplayData323) GetToleranceTypeOk() (*GBTToleranceType, bool)`

GetToleranceTypeOk returns a tuple with the ToleranceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToleranceType

`func (o *BTDimensionDisplayData323) SetToleranceType(v GBTToleranceType)`

SetToleranceType sets ToleranceType field to given value.

### HasToleranceType

`func (o *BTDimensionDisplayData323) HasToleranceType() bool`

HasToleranceType returns a boolean if a field has been set.

### GetUpperTolerance

`func (o *BTDimensionDisplayData323) GetUpperTolerance() float64`

GetUpperTolerance returns the UpperTolerance field if non-nil, zero value otherwise.

### GetUpperToleranceOk

`func (o *BTDimensionDisplayData323) GetUpperToleranceOk() (*float64, bool)`

GetUpperToleranceOk returns a tuple with the UpperTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperTolerance

`func (o *BTDimensionDisplayData323) SetUpperTolerance(v float64)`

SetUpperTolerance sets UpperTolerance field to given value.

### HasUpperTolerance

`func (o *BTDimensionDisplayData323) HasUpperTolerance() bool`

HasUpperTolerance returns a boolean if a field has been set.

### GetValue

`func (o *BTDimensionDisplayData323) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTDimensionDisplayData323) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTDimensionDisplayData323) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *BTDimensionDisplayData323) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


