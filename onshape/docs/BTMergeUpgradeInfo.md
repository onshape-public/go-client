# BTMergeUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingSourceUpgrade** | Pointer to [**BTPendingUpgradeInfo**](BTPendingUpgradeInfo.md) |  | [optional] 
**PendingTargetUpgrade** | Pointer to [**BTPendingUpgradeInfo**](BTPendingUpgradeInfo.md) |  | [optional] 
**RecommendedVersion** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**BTMergeUpgradeType**](BTMergeUpgradeType.md) |  | [optional] 

## Methods

### NewBTMergeUpgradeInfo

`func NewBTMergeUpgradeInfo() *BTMergeUpgradeInfo`

NewBTMergeUpgradeInfo instantiates a new BTMergeUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMergeUpgradeInfoWithDefaults

`func NewBTMergeUpgradeInfoWithDefaults() *BTMergeUpgradeInfo`

NewBTMergeUpgradeInfoWithDefaults instantiates a new BTMergeUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingSourceUpgrade

`func (o *BTMergeUpgradeInfo) GetPendingSourceUpgrade() BTPendingUpgradeInfo`

GetPendingSourceUpgrade returns the PendingSourceUpgrade field if non-nil, zero value otherwise.

### GetPendingSourceUpgradeOk

`func (o *BTMergeUpgradeInfo) GetPendingSourceUpgradeOk() (*BTPendingUpgradeInfo, bool)`

GetPendingSourceUpgradeOk returns a tuple with the PendingSourceUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingSourceUpgrade

`func (o *BTMergeUpgradeInfo) SetPendingSourceUpgrade(v BTPendingUpgradeInfo)`

SetPendingSourceUpgrade sets PendingSourceUpgrade field to given value.

### HasPendingSourceUpgrade

`func (o *BTMergeUpgradeInfo) HasPendingSourceUpgrade() bool`

HasPendingSourceUpgrade returns a boolean if a field has been set.

### GetPendingTargetUpgrade

`func (o *BTMergeUpgradeInfo) GetPendingTargetUpgrade() BTPendingUpgradeInfo`

GetPendingTargetUpgrade returns the PendingTargetUpgrade field if non-nil, zero value otherwise.

### GetPendingTargetUpgradeOk

`func (o *BTMergeUpgradeInfo) GetPendingTargetUpgradeOk() (*BTPendingUpgradeInfo, bool)`

GetPendingTargetUpgradeOk returns a tuple with the PendingTargetUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTargetUpgrade

`func (o *BTMergeUpgradeInfo) SetPendingTargetUpgrade(v BTPendingUpgradeInfo)`

SetPendingTargetUpgrade sets PendingTargetUpgrade field to given value.

### HasPendingTargetUpgrade

`func (o *BTMergeUpgradeInfo) HasPendingTargetUpgrade() bool`

HasPendingTargetUpgrade returns a boolean if a field has been set.

### GetRecommendedVersion

`func (o *BTMergeUpgradeInfo) GetRecommendedVersion() int32`

GetRecommendedVersion returns the RecommendedVersion field if non-nil, zero value otherwise.

### GetRecommendedVersionOk

`func (o *BTMergeUpgradeInfo) GetRecommendedVersionOk() (*int32, bool)`

GetRecommendedVersionOk returns a tuple with the RecommendedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedVersion

`func (o *BTMergeUpgradeInfo) SetRecommendedVersion(v int32)`

SetRecommendedVersion sets RecommendedVersion field to given value.

### HasRecommendedVersion

`func (o *BTMergeUpgradeInfo) HasRecommendedVersion() bool`

HasRecommendedVersion returns a boolean if a field has been set.

### GetType

`func (o *BTMergeUpgradeInfo) GetType() BTMergeUpgradeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTMergeUpgradeInfo) GetTypeOk() (*BTMergeUpgradeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTMergeUpgradeInfo) SetType(v BTMergeUpgradeType)`

SetType sets Type field to given value.

### HasType

`func (o *BTMergeUpgradeInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


