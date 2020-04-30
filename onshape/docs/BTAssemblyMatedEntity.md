# BTAssemblyMatedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MateCS** | Pointer to [**BTMateConnectorCSInfo**](BTMateConnectorCSInfo.md) |  | [optional] 
**MatedOccurrence** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBTAssemblyMatedEntity

`func NewBTAssemblyMatedEntity() *BTAssemblyMatedEntity`

NewBTAssemblyMatedEntity instantiates a new BTAssemblyMatedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyMatedEntityWithDefaults

`func NewBTAssemblyMatedEntityWithDefaults() *BTAssemblyMatedEntity`

NewBTAssemblyMatedEntityWithDefaults instantiates a new BTAssemblyMatedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMateCS

`func (o *BTAssemblyMatedEntity) GetMateCS() BTMateConnectorCSInfo`

GetMateCS returns the MateCS field if non-nil, zero value otherwise.

### GetMateCSOk

`func (o *BTAssemblyMatedEntity) GetMateCSOk() (*BTMateConnectorCSInfo, bool)`

GetMateCSOk returns a tuple with the MateCS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateCS

`func (o *BTAssemblyMatedEntity) SetMateCS(v BTMateConnectorCSInfo)`

SetMateCS sets MateCS field to given value.

### HasMateCS

`func (o *BTAssemblyMatedEntity) HasMateCS() bool`

HasMateCS returns a boolean if a field has been set.

### GetMatedOccurrence

`func (o *BTAssemblyMatedEntity) GetMatedOccurrence() []string`

GetMatedOccurrence returns the MatedOccurrence field if non-nil, zero value otherwise.

### GetMatedOccurrenceOk

`func (o *BTAssemblyMatedEntity) GetMatedOccurrenceOk() (*[]string, bool)`

GetMatedOccurrenceOk returns a tuple with the MatedOccurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatedOccurrence

`func (o *BTAssemblyMatedEntity) SetMatedOccurrence(v []string)`

SetMatedOccurrence sets MatedOccurrence field to given value.

### HasMatedOccurrence

`func (o *BTAssemblyMatedEntity) HasMatedOccurrence() bool`

HasMatedOccurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


