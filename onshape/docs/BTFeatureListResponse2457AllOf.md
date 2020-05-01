# BTFeatureListResponse2457AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**DefaultFeatures** | Pointer to [**[]BTMFeature134**](BTMFeature-134.md) |  | [optional] 
**FeatureStates** | Pointer to [**map[string]BTFeatureState1688**](BTFeatureState-1688.md) |  | [optional] 
**Features** | Pointer to [**[]BTMFeature134**](BTMFeature-134.md) |  | [optional] 
**Imports** | Pointer to [**[]BTMImport136**](BTMImport-136.md) |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] 
**RollbackIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTFeatureListResponse2457AllOf

`func NewBTFeatureListResponse2457AllOf() *BTFeatureListResponse2457AllOf`

NewBTFeatureListResponse2457AllOf instantiates a new BTFeatureListResponse2457AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTFeatureListResponse2457AllOfWithDefaults

`func NewBTFeatureListResponse2457AllOfWithDefaults() *BTFeatureListResponse2457AllOf`

NewBTFeatureListResponse2457AllOfWithDefaults instantiates a new BTFeatureListResponse2457AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTFeatureListResponse2457AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTFeatureListResponse2457AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTFeatureListResponse2457AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTFeatureListResponse2457AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDefaultFeatures

`func (o *BTFeatureListResponse2457AllOf) GetDefaultFeatures() []BTMFeature134`

GetDefaultFeatures returns the DefaultFeatures field if non-nil, zero value otherwise.

### GetDefaultFeaturesOk

`func (o *BTFeatureListResponse2457AllOf) GetDefaultFeaturesOk() (*[]BTMFeature134, bool)`

GetDefaultFeaturesOk returns a tuple with the DefaultFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFeatures

`func (o *BTFeatureListResponse2457AllOf) SetDefaultFeatures(v []BTMFeature134)`

SetDefaultFeatures sets DefaultFeatures field to given value.

### HasDefaultFeatures

`func (o *BTFeatureListResponse2457AllOf) HasDefaultFeatures() bool`

HasDefaultFeatures returns a boolean if a field has been set.

### GetFeatureStates

`func (o *BTFeatureListResponse2457AllOf) GetFeatureStates() map[string]BTFeatureState1688`

GetFeatureStates returns the FeatureStates field if non-nil, zero value otherwise.

### GetFeatureStatesOk

`func (o *BTFeatureListResponse2457AllOf) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool)`

GetFeatureStatesOk returns a tuple with the FeatureStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStates

`func (o *BTFeatureListResponse2457AllOf) SetFeatureStates(v map[string]BTFeatureState1688)`

SetFeatureStates sets FeatureStates field to given value.

### HasFeatureStates

`func (o *BTFeatureListResponse2457AllOf) HasFeatureStates() bool`

HasFeatureStates returns a boolean if a field has been set.

### GetFeatures

`func (o *BTFeatureListResponse2457AllOf) GetFeatures() []BTMFeature134`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTFeatureListResponse2457AllOf) GetFeaturesOk() (*[]BTMFeature134, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTFeatureListResponse2457AllOf) SetFeatures(v []BTMFeature134)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTFeatureListResponse2457AllOf) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetImports

`func (o *BTFeatureListResponse2457AllOf) GetImports() []BTMImport136`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *BTFeatureListResponse2457AllOf) GetImportsOk() (*[]BTMImport136, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *BTFeatureListResponse2457AllOf) SetImports(v []BTMImport136)`

SetImports sets Imports field to given value.

### HasImports

`func (o *BTFeatureListResponse2457AllOf) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetIsComplete

`func (o *BTFeatureListResponse2457AllOf) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *BTFeatureListResponse2457AllOf) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *BTFeatureListResponse2457AllOf) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *BTFeatureListResponse2457AllOf) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetRollbackIndex

`func (o *BTFeatureListResponse2457AllOf) GetRollbackIndex() int32`

GetRollbackIndex returns the RollbackIndex field if non-nil, zero value otherwise.

### GetRollbackIndexOk

`func (o *BTFeatureListResponse2457AllOf) GetRollbackIndexOk() (*int32, bool)`

GetRollbackIndexOk returns a tuple with the RollbackIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackIndex

`func (o *BTFeatureListResponse2457AllOf) SetRollbackIndex(v int32)`

SetRollbackIndex sets RollbackIndex field to given value.

### HasRollbackIndex

`func (o *BTFeatureListResponse2457AllOf) HasRollbackIndex() bool`

HasRollbackIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


