# BTUnitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultUnits** | Pointer to [**BTDefaultUnitsInfo**](BTDefaultUnitsInfo.md) |  | [optional] 
**UnitsDisplayPrecision** | Pointer to **map[string]int32** | Specifies the display precision for every supported unit. | [optional] 

## Methods

### NewBTUnitInfo

`func NewBTUnitInfo() *BTUnitInfo`

NewBTUnitInfo instantiates a new BTUnitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUnitInfoWithDefaults

`func NewBTUnitInfoWithDefaults() *BTUnitInfo`

NewBTUnitInfoWithDefaults instantiates a new BTUnitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultUnits

`func (o *BTUnitInfo) GetDefaultUnits() BTDefaultUnitsInfo`

GetDefaultUnits returns the DefaultUnits field if non-nil, zero value otherwise.

### GetDefaultUnitsOk

`func (o *BTUnitInfo) GetDefaultUnitsOk() (*BTDefaultUnitsInfo, bool)`

GetDefaultUnitsOk returns a tuple with the DefaultUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUnits

`func (o *BTUnitInfo) SetDefaultUnits(v BTDefaultUnitsInfo)`

SetDefaultUnits sets DefaultUnits field to given value.

### HasDefaultUnits

`func (o *BTUnitInfo) HasDefaultUnits() bool`

HasDefaultUnits returns a boolean if a field has been set.

### GetUnitsDisplayPrecision

`func (o *BTUnitInfo) GetUnitsDisplayPrecision() map[string]int32`

GetUnitsDisplayPrecision returns the UnitsDisplayPrecision field if non-nil, zero value otherwise.

### GetUnitsDisplayPrecisionOk

`func (o *BTUnitInfo) GetUnitsDisplayPrecisionOk() (*map[string]int32, bool)`

GetUnitsDisplayPrecisionOk returns a tuple with the UnitsDisplayPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsDisplayPrecision

`func (o *BTUnitInfo) SetUnitsDisplayPrecision(v map[string]int32)`

SetUnitsDisplayPrecision sets UnitsDisplayPrecision field to given value.

### HasUnitsDisplayPrecision

`func (o *BTUnitInfo) HasUnitsDisplayPrecision() bool`

HasUnitsDisplayPrecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


