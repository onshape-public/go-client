# BTCommonUnitsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuantityTypeToBaseUnits** | Pointer to **map[string]map[string]int32** |  | [optional] 
**Units** | Pointer to [**[]BTCommonUnitInfo**](BTCommonUnitInfo.md) |  | [optional] 

## Methods

### NewBTCommonUnitsInfo

`func NewBTCommonUnitsInfo() *BTCommonUnitsInfo`

NewBTCommonUnitsInfo instantiates a new BTCommonUnitsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCommonUnitsInfoWithDefaults

`func NewBTCommonUnitsInfoWithDefaults() *BTCommonUnitsInfo`

NewBTCommonUnitsInfoWithDefaults instantiates a new BTCommonUnitsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantityTypeToBaseUnits

`func (o *BTCommonUnitsInfo) GetQuantityTypeToBaseUnits() map[string]map[string]int32`

GetQuantityTypeToBaseUnits returns the QuantityTypeToBaseUnits field if non-nil, zero value otherwise.

### GetQuantityTypeToBaseUnitsOk

`func (o *BTCommonUnitsInfo) GetQuantityTypeToBaseUnitsOk() (*map[string]map[string]int32, bool)`

GetQuantityTypeToBaseUnitsOk returns a tuple with the QuantityTypeToBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityTypeToBaseUnits

`func (o *BTCommonUnitsInfo) SetQuantityTypeToBaseUnits(v map[string]map[string]int32)`

SetQuantityTypeToBaseUnits sets QuantityTypeToBaseUnits field to given value.

### HasQuantityTypeToBaseUnits

`func (o *BTCommonUnitsInfo) HasQuantityTypeToBaseUnits() bool`

HasQuantityTypeToBaseUnits returns a boolean if a field has been set.

### GetUnits

`func (o *BTCommonUnitsInfo) GetUnits() []BTCommonUnitInfo`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BTCommonUnitsInfo) GetUnitsOk() (*[]BTCommonUnitInfo, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BTCommonUnitsInfo) SetUnits(v []BTCommonUnitInfo)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BTCommonUnitsInfo) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


