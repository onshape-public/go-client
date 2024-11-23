# BTAssemblyMateValuesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MateValues** | Pointer to [**[]BTAssemblyMateValueInfo**](BTAssemblyMateValueInfo.md) | Describes the relative position of the first mate connector with respect to the second along the designated degrees of freedom (DOF) for mates in the specified assembly. | [optional] 

## Methods

### NewBTAssemblyMateValuesInfo

`func NewBTAssemblyMateValuesInfo() *BTAssemblyMateValuesInfo`

NewBTAssemblyMateValuesInfo instantiates a new BTAssemblyMateValuesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyMateValuesInfoWithDefaults

`func NewBTAssemblyMateValuesInfoWithDefaults() *BTAssemblyMateValuesInfo`

NewBTAssemblyMateValuesInfoWithDefaults instantiates a new BTAssemblyMateValuesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMateValues

`func (o *BTAssemblyMateValuesInfo) GetMateValues() []BTAssemblyMateValueInfo`

GetMateValues returns the MateValues field if non-nil, zero value otherwise.

### GetMateValuesOk

`func (o *BTAssemblyMateValuesInfo) GetMateValuesOk() (*[]BTAssemblyMateValueInfo, bool)`

GetMateValuesOk returns a tuple with the MateValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateValues

`func (o *BTAssemblyMateValuesInfo) SetMateValues(v []BTAssemblyMateValueInfo)`

SetMateValues sets MateValues field to given value.

### HasMateValues

`func (o *BTAssemblyMateValuesInfo) HasMateValues() bool`

HasMateValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


