# BTAssemblyFeatureDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MateType** | Pointer to **string** |  | [optional] 
**MatedEntities** | Pointer to [**[]BTAssemblyMatedEntity**](BTAssemblyMatedEntity.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAssemblyFeatureDataInfo

`func NewBTAssemblyFeatureDataInfo() *BTAssemblyFeatureDataInfo`

NewBTAssemblyFeatureDataInfo instantiates a new BTAssemblyFeatureDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyFeatureDataInfoWithDefaults

`func NewBTAssemblyFeatureDataInfoWithDefaults() *BTAssemblyFeatureDataInfo`

NewBTAssemblyFeatureDataInfoWithDefaults instantiates a new BTAssemblyFeatureDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMateType

`func (o *BTAssemblyFeatureDataInfo) GetMateType() string`

GetMateType returns the MateType field if non-nil, zero value otherwise.

### GetMateTypeOk

`func (o *BTAssemblyFeatureDataInfo) GetMateTypeOk() (*string, bool)`

GetMateTypeOk returns a tuple with the MateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMateType

`func (o *BTAssemblyFeatureDataInfo) SetMateType(v string)`

SetMateType sets MateType field to given value.

### HasMateType

`func (o *BTAssemblyFeatureDataInfo) HasMateType() bool`

HasMateType returns a boolean if a field has been set.

### GetMatedEntities

`func (o *BTAssemblyFeatureDataInfo) GetMatedEntities() []BTAssemblyMatedEntity`

GetMatedEntities returns the MatedEntities field if non-nil, zero value otherwise.

### GetMatedEntitiesOk

`func (o *BTAssemblyFeatureDataInfo) GetMatedEntitiesOk() (*[]BTAssemblyMatedEntity, bool)`

GetMatedEntitiesOk returns a tuple with the MatedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatedEntities

`func (o *BTAssemblyFeatureDataInfo) SetMatedEntities(v []BTAssemblyMatedEntity)`

SetMatedEntities sets MatedEntities field to given value.

### HasMatedEntities

`func (o *BTAssemblyFeatureDataInfo) HasMatedEntities() bool`

HasMatedEntities returns a boolean if a field has been set.

### GetName

`func (o *BTAssemblyFeatureDataInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAssemblyFeatureDataInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAssemblyFeatureDataInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAssemblyFeatureDataInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


