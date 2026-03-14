# BTMLoad3538AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**DefinedByComponents** | Pointer to **bool** |  | [optional] 
**DirectionFlipped** | Pointer to **bool** |  | [optional] 
**FgsBaseUnits** | Pointer to **string** |  | [optional] 
**LoadComponentParameterIds** | Pointer to **map[string]string** |  | [optional] 
**LoadRegionParameterId** | Pointer to **string** |  | [optional] 
**LoadType** | Pointer to [**GBTLoadType**](GBTLoadType.md) |  | [optional] 
**MagnitudeParameterId** | Pointer to **string** |  | [optional] 
**MagnitudeQuantityType** | Pointer to [**GBTQuantityType**](GBTQuantityType.md) |  | [optional] 
**StructuralLoad** | Pointer to **bool** |  | [optional] 
**SuppressedInSimulations** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewBTMLoad3538AllOf

`func NewBTMLoad3538AllOf() *BTMLoad3538AllOf`

NewBTMLoad3538AllOf instantiates a new BTMLoad3538AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMLoad3538AllOfWithDefaults

`func NewBTMLoad3538AllOfWithDefaults() *BTMLoad3538AllOf`

NewBTMLoad3538AllOfWithDefaults instantiates a new BTMLoad3538AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMLoad3538AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMLoad3538AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMLoad3538AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMLoad3538AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDefinedByComponents

`func (o *BTMLoad3538AllOf) GetDefinedByComponents() bool`

GetDefinedByComponents returns the DefinedByComponents field if non-nil, zero value otherwise.

### GetDefinedByComponentsOk

`func (o *BTMLoad3538AllOf) GetDefinedByComponentsOk() (*bool, bool)`

GetDefinedByComponentsOk returns a tuple with the DefinedByComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedByComponents

`func (o *BTMLoad3538AllOf) SetDefinedByComponents(v bool)`

SetDefinedByComponents sets DefinedByComponents field to given value.

### HasDefinedByComponents

`func (o *BTMLoad3538AllOf) HasDefinedByComponents() bool`

HasDefinedByComponents returns a boolean if a field has been set.

### GetDirectionFlipped

`func (o *BTMLoad3538AllOf) GetDirectionFlipped() bool`

GetDirectionFlipped returns the DirectionFlipped field if non-nil, zero value otherwise.

### GetDirectionFlippedOk

`func (o *BTMLoad3538AllOf) GetDirectionFlippedOk() (*bool, bool)`

GetDirectionFlippedOk returns a tuple with the DirectionFlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionFlipped

`func (o *BTMLoad3538AllOf) SetDirectionFlipped(v bool)`

SetDirectionFlipped sets DirectionFlipped field to given value.

### HasDirectionFlipped

`func (o *BTMLoad3538AllOf) HasDirectionFlipped() bool`

HasDirectionFlipped returns a boolean if a field has been set.

### GetFgsBaseUnits

`func (o *BTMLoad3538AllOf) GetFgsBaseUnits() string`

GetFgsBaseUnits returns the FgsBaseUnits field if non-nil, zero value otherwise.

### GetFgsBaseUnitsOk

`func (o *BTMLoad3538AllOf) GetFgsBaseUnitsOk() (*string, bool)`

GetFgsBaseUnitsOk returns a tuple with the FgsBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgsBaseUnits

`func (o *BTMLoad3538AllOf) SetFgsBaseUnits(v string)`

SetFgsBaseUnits sets FgsBaseUnits field to given value.

### HasFgsBaseUnits

`func (o *BTMLoad3538AllOf) HasFgsBaseUnits() bool`

HasFgsBaseUnits returns a boolean if a field has been set.

### GetLoadComponentParameterIds

`func (o *BTMLoad3538AllOf) GetLoadComponentParameterIds() map[string]string`

GetLoadComponentParameterIds returns the LoadComponentParameterIds field if non-nil, zero value otherwise.

### GetLoadComponentParameterIdsOk

`func (o *BTMLoad3538AllOf) GetLoadComponentParameterIdsOk() (*map[string]string, bool)`

GetLoadComponentParameterIdsOk returns a tuple with the LoadComponentParameterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadComponentParameterIds

`func (o *BTMLoad3538AllOf) SetLoadComponentParameterIds(v map[string]string)`

SetLoadComponentParameterIds sets LoadComponentParameterIds field to given value.

### HasLoadComponentParameterIds

`func (o *BTMLoad3538AllOf) HasLoadComponentParameterIds() bool`

HasLoadComponentParameterIds returns a boolean if a field has been set.

### GetLoadRegionParameterId

`func (o *BTMLoad3538AllOf) GetLoadRegionParameterId() string`

GetLoadRegionParameterId returns the LoadRegionParameterId field if non-nil, zero value otherwise.

### GetLoadRegionParameterIdOk

`func (o *BTMLoad3538AllOf) GetLoadRegionParameterIdOk() (*string, bool)`

GetLoadRegionParameterIdOk returns a tuple with the LoadRegionParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadRegionParameterId

`func (o *BTMLoad3538AllOf) SetLoadRegionParameterId(v string)`

SetLoadRegionParameterId sets LoadRegionParameterId field to given value.

### HasLoadRegionParameterId

`func (o *BTMLoad3538AllOf) HasLoadRegionParameterId() bool`

HasLoadRegionParameterId returns a boolean if a field has been set.

### GetLoadType

`func (o *BTMLoad3538AllOf) GetLoadType() GBTLoadType`

GetLoadType returns the LoadType field if non-nil, zero value otherwise.

### GetLoadTypeOk

`func (o *BTMLoad3538AllOf) GetLoadTypeOk() (*GBTLoadType, bool)`

GetLoadTypeOk returns a tuple with the LoadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadType

`func (o *BTMLoad3538AllOf) SetLoadType(v GBTLoadType)`

SetLoadType sets LoadType field to given value.

### HasLoadType

`func (o *BTMLoad3538AllOf) HasLoadType() bool`

HasLoadType returns a boolean if a field has been set.

### GetMagnitudeParameterId

`func (o *BTMLoad3538AllOf) GetMagnitudeParameterId() string`

GetMagnitudeParameterId returns the MagnitudeParameterId field if non-nil, zero value otherwise.

### GetMagnitudeParameterIdOk

`func (o *BTMLoad3538AllOf) GetMagnitudeParameterIdOk() (*string, bool)`

GetMagnitudeParameterIdOk returns a tuple with the MagnitudeParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeParameterId

`func (o *BTMLoad3538AllOf) SetMagnitudeParameterId(v string)`

SetMagnitudeParameterId sets MagnitudeParameterId field to given value.

### HasMagnitudeParameterId

`func (o *BTMLoad3538AllOf) HasMagnitudeParameterId() bool`

HasMagnitudeParameterId returns a boolean if a field has been set.

### GetMagnitudeQuantityType

`func (o *BTMLoad3538AllOf) GetMagnitudeQuantityType() GBTQuantityType`

GetMagnitudeQuantityType returns the MagnitudeQuantityType field if non-nil, zero value otherwise.

### GetMagnitudeQuantityTypeOk

`func (o *BTMLoad3538AllOf) GetMagnitudeQuantityTypeOk() (*GBTQuantityType, bool)`

GetMagnitudeQuantityTypeOk returns a tuple with the MagnitudeQuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeQuantityType

`func (o *BTMLoad3538AllOf) SetMagnitudeQuantityType(v GBTQuantityType)`

SetMagnitudeQuantityType sets MagnitudeQuantityType field to given value.

### HasMagnitudeQuantityType

`func (o *BTMLoad3538AllOf) HasMagnitudeQuantityType() bool`

HasMagnitudeQuantityType returns a boolean if a field has been set.

### GetStructuralLoad

`func (o *BTMLoad3538AllOf) GetStructuralLoad() bool`

GetStructuralLoad returns the StructuralLoad field if non-nil, zero value otherwise.

### GetStructuralLoadOk

`func (o *BTMLoad3538AllOf) GetStructuralLoadOk() (*bool, bool)`

GetStructuralLoadOk returns a tuple with the StructuralLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLoad

`func (o *BTMLoad3538AllOf) SetStructuralLoad(v bool)`

SetStructuralLoad sets StructuralLoad field to given value.

### HasStructuralLoad

`func (o *BTMLoad3538AllOf) HasStructuralLoad() bool`

HasStructuralLoad returns a boolean if a field has been set.

### GetSuppressedInSimulations

`func (o *BTMLoad3538AllOf) GetSuppressedInSimulations() map[string]int32`

GetSuppressedInSimulations returns the SuppressedInSimulations field if non-nil, zero value otherwise.

### GetSuppressedInSimulationsOk

`func (o *BTMLoad3538AllOf) GetSuppressedInSimulationsOk() (*map[string]int32, bool)`

GetSuppressedInSimulationsOk returns a tuple with the SuppressedInSimulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedInSimulations

`func (o *BTMLoad3538AllOf) SetSuppressedInSimulations(v map[string]int32)`

SetSuppressedInSimulations sets SuppressedInSimulations field to given value.

### HasSuppressedInSimulations

`func (o *BTMLoad3538AllOf) HasSuppressedInSimulations() bool`

HasSuppressedInSimulations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


