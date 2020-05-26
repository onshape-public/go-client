# BTUpdateFeaturesResponse1333

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**FeatureStates** | Pointer to [**map[string]BTFeatureState1688**](BTFeatureState-1688.md) |  | [optional] 
**Features** | Pointer to [**[]BTMFeature134**](BTMFeature-134.md) |  | [optional] 

## Methods

### NewBTUpdateFeaturesResponse1333

`func NewBTUpdateFeaturesResponse1333() *BTUpdateFeaturesResponse1333`

NewBTUpdateFeaturesResponse1333 instantiates a new BTUpdateFeaturesResponse1333 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUpdateFeaturesResponse1333WithDefaults

`func NewBTUpdateFeaturesResponse1333WithDefaults() *BTUpdateFeaturesResponse1333`

NewBTUpdateFeaturesResponse1333WithDefaults instantiates a new BTUpdateFeaturesResponse1333 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTUpdateFeaturesResponse1333) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTUpdateFeaturesResponse1333) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTUpdateFeaturesResponse1333) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTUpdateFeaturesResponse1333) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatureStates

`func (o *BTUpdateFeaturesResponse1333) GetFeatureStates() map[string]BTFeatureState1688`

GetFeatureStates returns the FeatureStates field if non-nil, zero value otherwise.

### GetFeatureStatesOk

`func (o *BTUpdateFeaturesResponse1333) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool)`

GetFeatureStatesOk returns a tuple with the FeatureStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStates

`func (o *BTUpdateFeaturesResponse1333) SetFeatureStates(v map[string]BTFeatureState1688)`

SetFeatureStates sets FeatureStates field to given value.

### HasFeatureStates

`func (o *BTUpdateFeaturesResponse1333) HasFeatureStates() bool`

HasFeatureStates returns a boolean if a field has been set.

### GetFeatures

`func (o *BTUpdateFeaturesResponse1333) GetFeatures() []BTMFeature134`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BTUpdateFeaturesResponse1333) GetFeaturesOk() (*[]BTMFeature134, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BTUpdateFeaturesResponse1333) SetFeatures(v []BTMFeature134)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BTUpdateFeaturesResponse1333) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


