# BTPendingUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledTime** | Pointer to **JSONTime** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTPendingUpgradeInfo

`func NewBTPendingUpgradeInfo() *BTPendingUpgradeInfo`

NewBTPendingUpgradeInfo instantiates a new BTPendingUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPendingUpgradeInfoWithDefaults

`func NewBTPendingUpgradeInfoWithDefaults() *BTPendingUpgradeInfo`

NewBTPendingUpgradeInfoWithDefaults instantiates a new BTPendingUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledTime

`func (o *BTPendingUpgradeInfo) GetScheduledTime() JSONTime`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *BTPendingUpgradeInfo) GetScheduledTimeOk() (*JSONTime, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *BTPendingUpgradeInfo) SetScheduledTime(v JSONTime)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *BTPendingUpgradeInfo) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetVersion

`func (o *BTPendingUpgradeInfo) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTPendingUpgradeInfo) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTPendingUpgradeInfo) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTPendingUpgradeInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


